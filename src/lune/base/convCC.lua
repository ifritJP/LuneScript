--lune/base/convCC.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@convCC'
local _lune = {}
if _lune1 then
   _lune = _lune1
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

if not _lune1 then
   _lune1 = _lune
end

local Ver = _lune.loadModule( 'lune.base.Ver' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Util = _lune.loadModule( 'lune.base.Util' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local LuaMod = _lune.loadModule( 'lune.base.LuaMod' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Parser = _lune.loadModule( 'lune.base.Parser' )

local cTypeInt = "lune_int_t"
local cTypeReal = "lune_real_t"
local cTypeBool = "lune_bool_t"
local cTypeStem = "lune_stem_t"
local cTypeAny = "lune_any_t"
local cTypeAnyP = "lune_any_t *"
local cTypeAnyPP = "lune_any_t **"
local cTypeEnvP = "lune_env_t *"
local cTypeVarP = "lune_var_t *"
local cValNil = "lune_global.nilStem"
local cValNone = "lune_global.noneStem"

local accessAny = ".val.pAny"

local stepIndent = 3

local scopeAccess = Ast.ScopeAccess.Full

local invalidSymbolId = -1

local function getSymbolName( symbolInfo )

   if Ast.isPubToExternal( symbolInfo:get_accessMode() ) then
      return symbolInfo:get_name()
   end
   
   do
      local _switchExp = symbolInfo:get_kind()
      if _switchExp == Ast.SymbolKind.Var then
         if symbolInfo:get_symbolId() == invalidSymbolId then
            return symbolInfo:get_name()
         end
         
         return string.format( "%s_%d", symbolInfo:get_name(), symbolInfo:get_symbolId())
      end
   end
   
   return symbolInfo:get_name()
end

local function getBelongClassType( node )

   if node:get_expType():get_kind() ~= Ast.TypeInfoKind.Method then
      return nil
   end
   
   local fieldNode = _lune.__Cast( node, 3, Nodes.RefFieldNode )
   if  nil == fieldNode then
      local _fieldNode = fieldNode
   
      return nil
   end
   
   return fieldNode:get_prefix():get_expType()
end

local function isClosure( funcType )

   do
      local scope = funcType:get_scope()
      if scope ~= nil then
         return #scope:get_closureSymList() > 0
      end
   end
   
   return false
end

local PubVarInfo = {}
function PubVarInfo.setmeta( obj )
  setmetatable( obj, { __index = PubVarInfo  } )
end
function PubVarInfo.new( staticFlag, accessMode, mutable, typeInfo )
   local obj = {}
   PubVarInfo.setmeta( obj )
   if obj.__init then
      obj:__init( staticFlag, accessMode, mutable, typeInfo )
   end
   return obj
end
function PubVarInfo:__init( staticFlag, accessMode, mutable, typeInfo )

   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.mutable = mutable
   self.typeInfo = typeInfo
end



local PubFuncInfo = {}
function PubFuncInfo.setmeta( obj )
  setmetatable( obj, { __index = PubFuncInfo  } )
end
function PubFuncInfo.new( accessMode, typeInfo )
   local obj = {}
   PubFuncInfo.setmeta( obj )
   if obj.__init then
      obj:__init( accessMode, typeInfo )
   end
   return obj
end
function PubFuncInfo:__init( accessMode, typeInfo )

   self.accessMode = accessMode
   self.typeInfo = typeInfo
end


local ConvMode = {}
_moduleObj.ConvMode = ConvMode
ConvMode._val2NameMap = {}
function ConvMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ConvMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ConvMode._from( val )
   if ConvMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ConvMode.__allList = {}
function ConvMode.get__allList()
   return ConvMode.__allList
end

ConvMode.Exec = 0
ConvMode._val2NameMap[0] = 'Exec'
ConvMode.__allList[1] = ConvMode.Exec
ConvMode.Convert = 1
ConvMode._val2NameMap[1] = 'Convert'
ConvMode.__allList[2] = ConvMode.Convert
ConvMode.ConvMeta = 2
ConvMode._val2NameMap[2] = 'ConvMeta'
ConvMode.__allList[3] = ConvMode.ConvMeta


local ModuleInfo = {}
setmetatable( ModuleInfo, { ifList = {Ast.ModuleInfoIF,} } )
function ModuleInfo.setmeta( obj )
  setmetatable( obj, { __index = ModuleInfo  } )
end
function ModuleInfo.new( assignName, modulePath )
   local obj = {}
   ModuleInfo.setmeta( obj )
   if obj.__init then
      obj:__init( assignName, modulePath )
   end
   return obj
end
function ModuleInfo:__init( assignName, modulePath )

   self.assignName = assignName
   self.modulePath = modulePath
end
function ModuleInfo:get_assignName()
   return self.assignName
end
function ModuleInfo:get_modulePath()
   return self.modulePath
end


local Opt = {}
_moduleObj.Opt = Opt
function Opt.setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
end
function Opt.new( node )
   local obj = {}
   Opt.setmeta( obj )
   if obj.__init then
      obj:__init( node )
   end
   return obj
end
function Opt:__init( node )

   self.node = node
end


local RoutineInfo = {}
function RoutineInfo.new( funcInfo )
   local obj = {}
   RoutineInfo.setmeta( obj )
   if obj.__init then obj:__init( funcInfo ); end
   return obj
end
function RoutineInfo:__init(funcInfo) 
   self.funcInfo = funcInfo
   self.blockDepth = 1
end
function RoutineInfo:pushDepth(  )

   self.blockDepth = self.blockDepth + 1
end
function RoutineInfo:popDepth(  )

   self.blockDepth = self.blockDepth - 1
end
function RoutineInfo.setmeta( obj )
  setmetatable( obj, { __index = RoutineInfo  } )
end
function RoutineInfo:get_funcInfo()
   return self.funcInfo
end
function RoutineInfo:get_blockDepth()
   return self.blockDepth
end


local ValKind = {}
ValKind._val2NameMap = {}
function ValKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ValKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ValKind._from( val )
   if ValKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ValKind.__allList = {}
function ValKind.get__allList()
   return ValKind.__allList
end

ValKind.Prim = 0
ValKind._val2NameMap[0] = 'Prim'
ValKind.__allList[1] = ValKind.Prim
ValKind.Any = 1
ValKind._val2NameMap[1] = 'Any'
ValKind.__allList[2] = ValKind.Any
ValKind.StemWork = 2
ValKind._val2NameMap[2] = 'StemWork'
ValKind.__allList[3] = ValKind.StemWork
ValKind.Stem = 3
ValKind._val2NameMap[3] = 'Stem'
ValKind.__allList[4] = ValKind.Stem
ValKind.Var = 4
ValKind._val2NameMap[4] = 'Var'
ValKind.__allList[5] = ValKind.Var
ValKind.Other = 5
ValKind._val2NameMap[5] = 'Other'
ValKind.__allList[6] = ValKind.Other


local function getValKind( valType )

   local expType = valType:get_srcTypeInfo()
   if expType:get_nilable() then
      return ValKind.Stem
   end
   
   do
      local _switchExp = expType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Alternate or _switchExp == Ast.TypeInfoKind.Stem or _switchExp == Ast.TypeInfoKind.DDD or _switchExp == Ast.TypeInfoKind.Alge then
         return ValKind.Stem
      end
   end
   
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         return ValKind.Prim
      elseif _switchExp == Ast.builtinTypeReal then
         return ValKind.Prim
      elseif _switchExp == Ast.builtinTypeBool then
         return ValKind.Prim
      else 
         
            
            do
               local enumType = _lune.__Cast( expType, 3, Ast.EnumTypeInfo )
               if enumType ~= nil then
                  return getValKind( enumType:get_valTypeInfo() )
               end
            end
            
            return ValKind.Any
      end
   end
   
end

local function isStemType( valType )

   return getValKind( valType ) == ValKind.Stem
end

local function getRetKind( retTypeList )

   do
      local _switchExp = #retTypeList
      if _switchExp == 0 then
         return ValKind.Other
      elseif _switchExp == 1 then
         return getValKind( retTypeList[1] )
      end
   end
   
   return ValKind.Stem
end

local function isStemRet( retTypeList )

   return getRetKind( retTypeList ) == ValKind.Stem
end

local function getCType( valType )

   local expType = valType:get_srcTypeInfo()
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         return cTypeInt
      elseif _switchExp == Ast.builtinTypeReal then
         return cTypeReal
      elseif _switchExp == Ast.builtinTypeBool then
         return cTypeBool
      else 
         
            do
               local enumType = _lune.__Cast( expType, 3, Ast.EnumTypeInfo )
               if enumType ~= nil then
                  return getCType( enumType:get_valTypeInfo() )
               end
            end
            
            
            if getValKind( valType ) == ValKind.Any then
               return cTypeAnyP
            end
            
            
            return cTypeStem
      end
   end
   
end

local function getCRetType( retTypeList )

   do
      local _switchExp = #retTypeList
      if _switchExp == 0 then
         return "void"
      elseif _switchExp == 1 then
         return getCType( retTypeList[1] )
      end
   end
   
   return cTypeStem
end

local function getBlockName( scope )

   return string.format( "pBlock_%X", scope:get_scopeId())
end

local ProcessMode = {}
ProcessMode._val2NameMap = {}
function ProcessMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ProcessMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ProcessMode._from( val )
   if ProcessMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ProcessMode.__allList = {}
function ProcessMode.get__allList()
   return ProcessMode.__allList
end

ProcessMode.Include = 0
ProcessMode._val2NameMap[0] = 'Include'
ProcessMode.__allList[1] = ProcessMode.Include
ProcessMode.Prototype = 1
ProcessMode._val2NameMap[1] = 'Prototype'
ProcessMode.__allList[2] = ProcessMode.Prototype
ProcessMode.WideScopeVer = 2
ProcessMode._val2NameMap[2] = 'WideScopeVer'
ProcessMode.__allList[3] = ProcessMode.WideScopeVer
ProcessMode.InitModule = 3
ProcessMode._val2NameMap[3] = 'InitModule'
ProcessMode.__allList[4] = ProcessMode.InitModule
ProcessMode.Intermediate = 4
ProcessMode._val2NameMap[4] = 'Intermediate'
ProcessMode.__allList[5] = ProcessMode.Intermediate
ProcessMode.StringFormat = 5
ProcessMode._val2NameMap[5] = 'StringFormat'
ProcessMode.__allList[6] = ProcessMode.StringFormat
ProcessMode.DefClass = 6
ProcessMode._val2NameMap[6] = 'DefClass'
ProcessMode.__allList[7] = ProcessMode.DefClass
ProcessMode.Form = 7
ProcessMode._val2NameMap[7] = 'Form'
ProcessMode.__allList[8] = ProcessMode.Form
ProcessMode.Immediate = 8
ProcessMode._val2NameMap[8] = 'Immediate'
ProcessMode.__allList[9] = ProcessMode.Immediate


local ModuleCtrl = {}
function ModuleCtrl.new( typeNameCtrl, moduleInfoManager )
   local obj = {}
   ModuleCtrl.setmeta( obj )
   if obj.__init then obj:__init( typeNameCtrl, moduleInfoManager ); end
   return obj
end
function ModuleCtrl:__init(typeNameCtrl, moduleInfoManager) 
   self.typeNameCtrl = typeNameCtrl
   self.moduleInfoManager = moduleInfoManager
end
function ModuleCtrl:getFilePath( moduleTypeInfo )

   local workName = moduleTypeInfo:getFullName( self.typeNameCtrl, self.moduleInfoManager, false )
   local fullName = string.format( "%s", (workName:gsub( "[&@]", "" ):gsub( "%.", "/" ) ))
   return fullName
end
function ModuleCtrl:getFullName( typeInfo )

   typeInfo = typeInfo:get_srcTypeInfo()
   local workName = typeInfo:getFullName( self.typeNameCtrl, self.moduleInfoManager, false )
   
   local fullName = string.format( "%s", (workName:gsub( "[&@]", "" ):gsub( "%.", "_" ) ))
   if Ast.isPubToExternal( typeInfo:get_accessMode() ) then
      return fullName
   end
   
   return string.format( "_%d_%s", typeInfo:get_typeId(), fullName)
end
function ModuleCtrl:getAlgeCName( algeType )

   return self:getFullName( algeType )
end
function ModuleCtrl:getAlgeEnumCName( algeType )

   return string.format( "lune_algeType_%s", self:getAlgeCName( algeType ))
end
function ModuleCtrl:getAlgeValCName( algeType, valName )

   return string.format( "lune_alge_%s_%s", self:getFullName( algeType ), valName)
end
function ModuleCtrl:getAlgeValStrCName( algeType, valName )

   return string.format( "lune_alge_%s_%s_t", self:getFullName( algeType ), valName)
end
function ModuleCtrl:getNewAlgeCName( algeType, valName )

   return string.format( "lune_new_alge_%s_%s", self:getFullName( algeType ), valName)
end
function ModuleCtrl:getAlgeInitCName( algeType )

   return string.format( "lune_init_alge_%s", self:getAlgeCName( algeType ))
end
function ModuleCtrl:getEnumTypeName( typeInfo )

   local srcType = typeInfo:get_srcTypeInfo()
   local fullName = self:getFullName( srcType )
   
   if Ast.isPubToExternal( typeInfo:get_accessMode() ) then
      return fullName
   end
   
   return string.format( "%s_%d", fullName, srcType:get_typeId())
end
function ModuleCtrl:getEnumValCName( typeInfo, valName )

   return string.format( "%s__%s", self:getEnumTypeName( typeInfo ), valName)
end
function ModuleCtrl:getEnumVal2NameMapName( enumType )

   return string.format( "%s_val2NameMap", self:getEnumTypeName( enumType ))
end
function ModuleCtrl:getFuncName( typeInfo )

   if typeInfo:get_rawTxt() == "" then
      return string.format( "_anonymous_%d", typeInfo:get_typeId())
   end
   
   do
      local _switchExp = typeInfo:get_accessMode()
      if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
         local builtinFunc = TransUnit.getBuiltinFunc(  )
         if typeInfo == builtinFunc.lune_print then
            return "lune_print"
         end
         
         return self:getFullName( typeInfo )
      end
   end
   
   return string.format( "lune_f_%d_%s", typeInfo:get_typeId(), typeInfo:get_rawTxt())
end
function ModuleCtrl:getClassCName( classType )

   return self:getFullName( classType )
end
function ModuleCtrl:getMethodCName( methodTypeInfo )

   return string.format( "u_mtd_%s_%s", self:getClassCName( methodTypeInfo:get_parentInfo() ), methodTypeInfo:get_rawTxt())
end
function ModuleCtrl:getCallMethodCName( methodTypeInfo )

   do
      local _switchExp = methodTypeInfo:get_parentInfo():get_kind()
      if _switchExp == Ast.TypeInfoKind.List then
         return string.format( "lune_mtd_List_%s", methodTypeInfo:get_rawTxt())
      elseif _switchExp == Ast.TypeInfoKind.Array then
         return string.format( "lune_mtd_Array_%s", methodTypeInfo:get_rawTxt())
      elseif _switchExp == Ast.TypeInfoKind.Set then
         return string.format( "lune_mtd_Set_%s", methodTypeInfo:get_rawTxt())
      elseif _switchExp == Ast.TypeInfoKind.Map then
         return string.format( "lune_mtd_Map_%s", methodTypeInfo:get_rawTxt())
      end
   end
   
   return string.format( "u_call_mtd_%s_%s", self:getClassCName( methodTypeInfo:get_parentInfo() ), methodTypeInfo:get_rawTxt())
end
function ModuleCtrl:getClassMemberName( symbolInfo )

   local classTypeInfo = symbolInfo:get_scope():getClassTypeInfo(  )
   return string.format( "lune_var_%s_%s", self:getClassCName( classTypeInfo ), symbolInfo:get_name())
end
function ModuleCtrl:getFormName( typeInfo )

   return string.format( "lune_form_%s", self:getFullName( typeInfo ))
end
function ModuleCtrl:getCallFormName( typeInfo )

   return string.format( "lune_call_formFunc_%s", self:getFullName( typeInfo ))
end
function ModuleCtrl.setmeta( obj )
  setmetatable( obj, { __index = ModuleCtrl  } )
end


local SymbolParam = {}
function SymbolParam.setmeta( obj )
  setmetatable( obj, { __index = SymbolParam  } )
end
function SymbolParam.new( kind, index, typeTxt )
   local obj = {}
   SymbolParam.setmeta( obj )
   if obj.__init then
      obj:__init( kind, index, typeTxt )
   end
   return obj
end
function SymbolParam:__init( kind, index, typeTxt )

   self.kind = kind
   self.index = index
   self.typeTxt = typeTxt
end


local WorkSymbol = {}
setmetatable( WorkSymbol, { ifList = {Ast.LowSymbol,} } )
function WorkSymbol:get_symbolId(  )

   return invalidSymbolId
end
function WorkSymbol:get_hasAccessFromClosure(  )

   return false
end
function WorkSymbol.setmeta( obj )
  setmetatable( obj, { __index = WorkSymbol  } )
end
function WorkSymbol.new( scope, accessMode, name, typeInfo, kind, convModuleParam )
   local obj = {}
   WorkSymbol.setmeta( obj )
   if obj.__init then
      obj:__init( scope, accessMode, name, typeInfo, kind, convModuleParam )
   end
   return obj
end
function WorkSymbol:__init( scope, accessMode, name, typeInfo, kind, convModuleParam )

   self.scope = scope
   self.accessMode = accessMode
   self.name = name
   self.typeInfo = typeInfo
   self.kind = kind
   self.convModuleParam = convModuleParam
end
function WorkSymbol:get_scope()
   return self.scope
end
function WorkSymbol:get_accessMode()
   return self.accessMode
end
function WorkSymbol:get_name()
   return self.name
end
function WorkSymbol:get_typeInfo()
   return self.typeInfo
end
function WorkSymbol:get_kind()
   return self.kind
end
function WorkSymbol:get_convModuleParam()
   return self.convModuleParam
end


local ScopeInfo = {}
function ScopeInfo.setmeta( obj )
  setmetatable( obj, { __index = ScopeInfo  } )
end
function ScopeInfo.new( anyNum, stemNum, varNum )
   local obj = {}
   ScopeInfo.setmeta( obj )
   if obj.__init then
      obj:__init( anyNum, stemNum, varNum )
   end
   return obj
end
function ScopeInfo:__init( anyNum, stemNum, varNum )

   self.anyNum = anyNum
   self.stemNum = stemNum
   self.varNum = varNum
end


local function getOrgTypeInfo( typeInfo )

   do
      local enumType = _lune.__Cast( typeInfo:get_srcTypeInfo():get_nonnilableType(), 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         return enumType:get_valTypeInfo()
      end
   end
   
   return typeInfo:get_srcTypeInfo():get_nonnilableType()
end

local function getAccessPrimValFromSymbolDirect( symName, valKind, symType )

   local txt = symName
   do
      local _switchExp = valKind
      if _switchExp == ValKind.Var then
         txt = txt .. "->stem"
      elseif _switchExp == ValKind.Stem then
      elseif _switchExp == ValKind.Prim then
         return txt
      end
   end
   
   
   do
      local _switchExp = getOrgTypeInfo( symType )
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         txt = txt .. ".val.intVal"
      elseif _switchExp == Ast.builtinTypeReal then
         txt = txt .. ".val.realVal"
      elseif _switchExp == Ast.builtinTypeBool then
         txt = txt .. ".val.boolVal"
      end
   end
   
   return txt
end

local function createSymbolParam( symbol, valKind, cTypeTxt )

   do
      local _switchExp = valKind
      if _switchExp == ValKind.Stem then
         return SymbolParam.new(ValKind.Stem, 0, cTypeStem)
      elseif _switchExp == ValKind.Any then
         return SymbolParam.new(ValKind.Any, 0, cTypeAnyP)
      elseif _switchExp == ValKind.Prim then
         return SymbolParam.new(ValKind.Prim, 0, cTypeTxt)
      elseif _switchExp == ValKind.Other then
         return SymbolParam.new(ValKind.Other, 0, "void")
      else 
         
            Util.err( string.format( "not support %s:%s", symbol:get_name(), ValKind:_getTxt( valKind)
            ) )
      end
   end
   
end

local ScopeMgr = {}
function ScopeMgr.new(  )
   local obj = {}
   ScopeMgr.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function ScopeMgr:__init() 
   self.scope2InfoMap = {}
end
function ScopeMgr:setupScopeParam( scope )

   do
      local scopeInfo = self.scope2InfoMap[scope]
      if scopeInfo ~= nil then
         return scopeInfo.anyNum, scopeInfo.stemNum, scopeInfo.varNum
      end
   end
   
   
   local scopeInfo
   
   if not _lune.nilacc( scope:get_ownerTypeInfo(), 'get_externalFlag', 'callmtd' ) then
      local varNum = 0
      local stemNum = 0
      local anyNum = 0
      
      if _lune.nilacc( scope:get_ownerTypeInfo(), 'isModule', 'callmtd'  ) then
         anyNum = anyNum + 2
      end
      
      
      do
         local __sorted = {}
         local __map = scope:get_symbol2SymbolInfoMap()
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, __key in ipairs( __sorted ) do
            local symbol = __map[ __key ]
            do
               local param
               
               if symbol:get_name() ~= "__func__" then
                  do
                     local _switchExp = symbol:get_kind()
                     if _switchExp == Ast.SymbolKind.Var or _switchExp == Ast.SymbolKind.Arg then
                        if symbol:get_hasAccessFromClosure() then
                           param = SymbolParam.new(ValKind.Var, varNum, cTypeVarP)
                           varNum = varNum + 1
                        else
                         
                           do
                              local _switchExp = getValKind( symbol:get_typeInfo() )
                              if _switchExp == ValKind.Stem then
                                 param = SymbolParam.new(ValKind.Stem, stemNum, cTypeStem)
                                 stemNum = stemNum + 1
                              elseif _switchExp == ValKind.Any then
                                 if symbol:get_name() == "self" then
                                    param = SymbolParam.new(ValKind.Any, 0, cTypeAnyP)
                                 else
                                  
                                    if symbol:get_kind() == Ast.SymbolKind.Var then
                                       param = SymbolParam.new(ValKind.Any, anyNum, cTypeAnyPP)
                                    else
                                     
                                       param = SymbolParam.new(ValKind.Any, anyNum, cTypeAnyP)
                                    end
                                    
                                    anyNum = anyNum + 1
                                 end
                                 
                              elseif _switchExp == ValKind.Prim then
                                 param = SymbolParam.new(ValKind.Prim, 0, getCType( symbol:get_typeInfo() ))
                              else 
                                 
                                    Util.err( string.format( "not support %s", symbol:get_typeInfo():getTxt(  )) )
                              end
                           end
                           
                        end
                        
                     elseif _switchExp == Ast.SymbolKind.Mtd then
                        local retTypeList = symbol:get_typeInfo():get_retTypeInfoList()
                        param = createSymbolParam( symbol, getRetKind( retTypeList ), getCRetType( retTypeList ) )
                     elseif _switchExp == Ast.SymbolKind.Mbr then
                        param = createSymbolParam( symbol, getValKind( symbol:get_typeInfo() ), getCType( symbol:get_typeInfo() ) )
                     else 
                        
                           param = SymbolParam.new(ValKind.Other, 0, cTypeStem)
                     end
                  end
                  
               else
                
                  param = SymbolParam.new(ValKind.Any, 0, cTypeAnyP)
               end
               
               symbol:set_convModuleParam( param )
            end
         end
      end
      
      
      scopeInfo = ScopeInfo.new(anyNum, stemNum, varNum)
   else
    
      scopeInfo = ScopeInfo.new(0, 0, 0)
   end
   
   self.scope2InfoMap[scope] = scopeInfo
   
   return scopeInfo.anyNum, scopeInfo.stemNum, scopeInfo.varNum
end
function ScopeMgr:getSymbolParam( symbol )

   do
      local param = symbol:get_convModuleParam()
      if param ~= nil then
         return param
      end
   end
   
   
   local scope = symbol:get_scope()
   
   if not self.scope2InfoMap[scope] then
      self:setupScopeParam( scope )
      do
         local param = symbol:get_convModuleParam()
         if param ~= nil then
            return param
         end
      end
      
      
   end
   
   Util.err( string.format( "illegal symbol -- %s %d", symbol:get_name(), 673) )
end
function ScopeMgr:getSymbolValKind( symbol )

   local symbolParam = self:getSymbolParam( symbol )
   return symbolParam.kind
end
function ScopeMgr:getCTypeForSym( symbol )

   local param = self:getSymbolParam( symbol )
   return param.typeTxt, param.kind
end
function ScopeMgr:symbol2Any( symbol )

   local name = getSymbolName( symbol )
   
   do
      local _switchExp = self:getSymbolValKind( symbol )
      if _switchExp == ValKind.Var then
         return name
      else 
         
            Util.err( string.format( "not support -- %s", symbol:get_typeInfo():getTxt(  )) )
      end
   end
   
end
function ScopeMgr:getAccessPrimValFromSymbol( symbolInfo )

   return getAccessPrimValFromSymbolDirect( getSymbolName( symbolInfo ), self:getSymbolValKind( symbolInfo ), symbolInfo:get_typeInfo() )
end
function ScopeMgr:getAccessPrimValFromSymbolOnly( symbolInfo )

   return getAccessPrimValFromSymbolDirect( "", self:getSymbolValKind( symbolInfo ), symbolInfo:get_typeInfo() )
end
function ScopeMgr.setmeta( obj )
  setmetatable( obj, { __index = ScopeMgr  } )
end


local Out2HMode = {}
Out2HMode._val2NameMap = {}
function Out2HMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Out2HMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function Out2HMode._from( val )
   if Out2HMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
Out2HMode.__allList = {}
function Out2HMode.get__allList()
   return Out2HMode.__allList
end

Out2HMode.HeaderPub = 0
Out2HMode._val2NameMap[0] = 'HeaderPub'
Out2HMode.__allList[1] = Out2HMode.HeaderPub
Out2HMode.HeaderPri = 1
Out2HMode._val2NameMap[1] = 'HeaderPri'
Out2HMode.__allList[2] = Out2HMode.HeaderPri
Out2HMode.SourcePub = 2
Out2HMode._val2NameMap[2] = 'SourcePub'
Out2HMode.__allList[3] = Out2HMode.SourcePub
Out2HMode.SourcePri = 3
Out2HMode._val2NameMap[3] = 'SourcePri'
Out2HMode.__allList[4] = Out2HMode.SourcePri


local function getOut2HeaderPrefix( mode )

   do
      local _switchExp = mode
      if _switchExp == Out2HMode.HeaderPub then
         return "extern "
      elseif _switchExp == Out2HMode.SourcePri then
         return "static "
      end
   end
   
   return ""
end



local convFilter = {}
setmetatable( convFilter, { __index = Nodes.Filter } )
function convFilter.new( enableTest, outputBuiltin, streamName, stream, headerStream, ast )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( enableTest, outputBuiltin, streamName, stream, headerStream, ast ); end
   return obj
end
function convFilter:__init(enableTest, outputBuiltin, streamName, stream, headerStream, ast) 
   Nodes.Filter.__init( self,ast:get_moduleTypeInfo(), ast:get_moduleTypeInfo():get_scope())
   
   
   self.enableTest = enableTest
   self.outputBuiltinFlag = outputBuiltin
   self.scopeMgr = ScopeMgr.new()
   self.processingNode = nil
   self.processedNodeSet = {}
   self.accessSymbolSet = Util.OrderedSet.new()
   self.literalNode2AccessSymbolSet = {}
   
   self.duringDeclFunc = false
   
   self.processMode = ProcessMode.Prototype
   
   self.routineInfoQueue = {}
   self.currentRoutineInfo = RoutineInfo.new(Ast.builtinTypeNone)
   
   self.moduleTypeInfo = ast:get_moduleTypeInfo()
   self.ast = _lune.unwrap( _lune.__Cast( ast:get_node(), 3, Nodes.RootNode ))
   self.streamName = streamName
   self.streamQueue = {}
   self.pubVarName2InfoMap = {}
   self.pubFuncName2InfoMap = {}
   self.moduleCtrl = ModuleCtrl.new(self:get_typeNameCtrl(), self:get_moduleInfoManager())
   
   self.currentRoutineInfo = RoutineInfo.new(ast:get_moduleTypeInfo())
   table.insert( self.routineInfoQueue, self.currentRoutineInfo )
   
   self.stream = Util.SimpleSourceOStream.new(stream, headerStream, stepIndent)
end
function convFilter:pushStream(  )

   table.insert( self.streamQueue, self.stream )
   local stream = Util.memStream.new()
   self.stream = Util.SimpleSourceOStream.new(stream, nil, stepIndent)
   return stream
end
function convFilter:popStream(  )

   if #self.streamQueue == 0 then
      Util.err( "streamQueue is empty." )
   end
   
   self.stream = self.streamQueue[#self.streamQueue]
   table.remove( self.streamQueue )
end
function convFilter:getFullName( typeInfo )

   return self.moduleCtrl:getFullName( typeInfo )
end
function convFilter:processAddModuleGlobal( valName )

   self.stream:writeln( string.format( "lune_mtd_List_insert( _pEnv, *lune_module_globalStemList, %s );", valName) )
end
function convFilter.setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end
function convFilter:write( ... )
   return self.stream:write( ... )
end

function convFilter:writeln( ... )
   return self.stream:writeln( ... )
end

function convFilter:pushIndent( ... )
   return self.stream:pushIndent( ... )
end

function convFilter:popIndent( ... )
   return self.stream:popIndent( ... )
end

function convFilter:switchToHeader( ... )
   return self.stream:switchToHeader( ... )
end

function convFilter:returnToSource( ... )
   return self.stream:returnToSource( ... )
end



local function filter( node, filter, parent )

   node:processFilter( filter, Opt.new(parent) )
end

function convFilter:processNone( node, opt )

end



function convFilter:processImport( node, opt )

   if self.processMode == ProcessMode.Include then
      local function process( out2HMode )
      
         do
            local _switchExp = out2HMode
            if _switchExp == Out2HMode.HeaderPub then
               if node:get_symbolInfo():get_scope() ~= node:get_moduleTypeInfo():get_scope() then
                  return 
               end
               
            elseif _switchExp == Out2HMode.SourcePub or _switchExp == Out2HMode.SourcePri then
            else 
               
                  return 
            end
         end
         
         self:writeln( string.format( "#include<%s.h>", (node:get_modulePath():gsub( "%.", "/" ) )) )
      end
      do
         local function processwork( out2HMode )
         
            process( out2HMode )
            
            
         end
         if true then
            self.stream:switchToHeader(  )
            processwork( Out2HMode.HeaderPub )
            self.stream:returnToSource(  )
            
            processwork( Out2HMode.SourcePub )
         else
          
            processwork( Out2HMode.SourcePri )
         end
         
      end
      
      
   else
    
      self:writeln( string.format( "lune_init_%s( _pEnv );", self.moduleCtrl:getFullName( node:get_moduleTypeInfo() )) )
   end
   
end



local function getSymbolIndex( symbol )

   local param = symbol:get_convModuleParam()
   if  nil == param then
      local _param = param
   
      return 0
   end
   
   return (param ).index
end

function convFilter:processInitModule( node )

   local anyNum, stemNum, varNum = self.scopeMgr:setupScopeParam( self.ast:get_moduleScope() )
   
   self.processMode = ProcessMode.InitModule
   
   if self.outputBuiltinFlag then
      self:writeln( "static void lune_init_builtinSub( lune_env_t * _pEnv );" )
   end
   
   
   local function process( out2HMode )
   
      self:write( string.format( "%svoid lune_init_%s( %s _pEnv )", getOut2HeaderPrefix( out2HMode ), self.moduleCtrl:getFullName( node:get_moduleTypeInfo() ), cTypeEnvP) )
      
      do
         local _switchExp = out2HMode
         if _switchExp == Out2HMode.HeaderPub then
            self:writeln( ";" )
         elseif _switchExp == Out2HMode.SourcePub then
            self:writeln( "{" )
         end
      end
      
   end
   
   do
      local function processwork( out2HMode )
      
         process( out2HMode )
         
         
      end
      if true then
         self.stream:switchToHeader(  )
         processwork( Out2HMode.HeaderPub )
         self.stream:returnToSource(  )
         
         processwork( Out2HMode.SourcePub )
      else
       
         processwork( Out2HMode.SourcePri )
      end
      
   end
   
   
   
   self:pushIndent(  )
   
   local moduneBlockName = getBlockName( self.ast:get_moduleScope() )
   self:writeln( string.format( "lune_block_t * %s = lune_enter_module( %d, %d, %d );", moduneBlockName, anyNum, stemNum, varNum) )
   
   self:writeln( string.format( "lune_set_block_any( %s, 0, lune_module_globalStemList);", moduneBlockName) )
   self:writeln( "lune_setQ_any( lune_module_globalStemList, lune_class_List_new( _pEnv ));" )
   
   self:writeln( string.format( "lune_set_block_any( %s, 1, lune_module_path);", moduneBlockName) )
   self:writeln( string.format( 'lune_setQ_any( lune_module_path, lune_litStr2any( _pEnv, "%s"));', self.moduleCtrl:getFullName( node:get_moduleTypeInfo() )) )
   
   self:writeln( "lune_enter_block( _pEnv, 0, 0, 0 );" )
   
   if self.outputBuiltinFlag then
      self:writeln( "lune_init_builtinSub( _pEnv );" )
   else
    
      for __index, declAlgeNode in pairs( node:get_nodeManager():getDeclAlgeNodeList(  ) ) do
         filter( declAlgeNode, self, node )
      end
      
      
      for __index, child in pairs( node:get_children() ) do
         do
            local _switchExp = child:get_kind()
            if _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_DeclFunc() or _switchExp == Nodes.NodeKind.get_DeclMacro() then
            else 
               
                  filter( child, self, node )
                  self:writeln( "" )
            end
         end
         
      end
      
   end
   
   
   self:writeln( "lune_leave_block( _pEnv );" )
   
   self:popIndent(  )
   
   self:writeln( "}" )
end


local BuiltinArgSymbolInfo = {}
setmetatable( BuiltinArgSymbolInfo, { __index = Ast.SymbolInfo } )
function BuiltinArgSymbolInfo:get_canBeLeft(  )

   return false
end
function BuiltinArgSymbolInfo:get_canBeRight(  )

   return true
end
function BuiltinArgSymbolInfo:get_symbolId(  )

   return 0
end
function BuiltinArgSymbolInfo:get_accessMode(  )

   return Ast.AccessMode.Pub
end
function BuiltinArgSymbolInfo:get_staticFlag(  )

   return false
end
function BuiltinArgSymbolInfo:get_kind(  )

   return Ast.SymbolKind.Arg
end
function BuiltinArgSymbolInfo:get_pos(  )

   return nil
end
function BuiltinArgSymbolInfo:get_mutable(  )

   return true
end
function BuiltinArgSymbolInfo:get_mutMode(  )

   return Ast.MutMode.Mut
end
function BuiltinArgSymbolInfo:get_hasValueFlag(  )

   return true
end
function BuiltinArgSymbolInfo:set_hasValueFlag( arg )

end
function BuiltinArgSymbolInfo:get_hasAccessFromClosure(  )

   return false
end
function BuiltinArgSymbolInfo:set_hasAccessFromClosure( flag )

end
function BuiltinArgSymbolInfo:canAccess( fromScope, access )

   return self
end
function BuiltinArgSymbolInfo:getOrg(  )

   return self
end
function BuiltinArgSymbolInfo.setmeta( obj )
  setmetatable( obj, { __index = BuiltinArgSymbolInfo  } )
end
function BuiltinArgSymbolInfo.new( scope, name, typeInfo, convModuleParam, namespaceTypeInfo )
   local obj = {}
   BuiltinArgSymbolInfo.setmeta( obj )
   if obj.__init then
      obj:__init( scope, name, typeInfo, convModuleParam, namespaceTypeInfo )
   end
   return obj
end
function BuiltinArgSymbolInfo:__init( scope, name, typeInfo, convModuleParam, namespaceTypeInfo )

   Ast.SymbolInfo.__init( self)
   self.scope = scope
   self.name = name
   self.typeInfo = typeInfo
   self.convModuleParam = convModuleParam
   self.namespaceTypeInfo = namespaceTypeInfo
end
function BuiltinArgSymbolInfo:get_scope()
   return self.scope
end
function BuiltinArgSymbolInfo:get_name()
   return self.name
end
function BuiltinArgSymbolInfo:get_typeInfo()
   return self.typeInfo
end
function BuiltinArgSymbolInfo:set_typeInfo( typeInfo )
   self.typeInfo = typeInfo
end
function BuiltinArgSymbolInfo:get_convModuleParam()
   return self.convModuleParam
end
function BuiltinArgSymbolInfo:set_convModuleParam( convModuleParam )
   self.convModuleParam = convModuleParam
end
function BuiltinArgSymbolInfo:get_namespaceTypeInfo()
   return self.namespaceTypeInfo
end


local function registerBuiltin(  )

   local builtin = TransUnit.getBuiltinFunc(  )
   for symbol, __val in pairs( builtin:get_allSymbol() ) do
      local param
      
      do
         local _switchExp = symbol:get_kind()
         if _switchExp == Ast.SymbolKind.Mtd or _switchExp == Ast.SymbolKind.Fun then
            local retTypeList = symbol:get_typeInfo():get_retTypeInfoList()
            param = createSymbolParam( symbol, getRetKind( retTypeList ), getCRetType( retTypeList ) )
         elseif _switchExp == Ast.SymbolKind.Mbr then
            param = createSymbolParam( symbol, getValKind( symbol:get_typeInfo() ), getCType( symbol:get_typeInfo() ) )
         else 
            
               Util.err( string.format( "illeal symbol -- %s %d", symbol:get_name(), 1167) )
         end
      end
      
      symbol:set_convModuleParam( param )
   end
   
end

function convFilter:processBuiltin(  )

   local nodeManager = Nodes.NodeManager.new()
   local dummyPos = Parser.Position.new(0, 0)
   
   local builtin = TransUnit.getBuiltinFunc(  )
   
   for __index, classInfo in pairs( builtin:get_allClass() ) do
      do
         local _switchExp = classInfo:get_kind()
         if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Map or _switchExp == Ast.TypeInfoKind.Box then
         else 
            
               if classInfo ~= Ast.builtinTypeString then
                  print( classInfo:getTxt(  ) )
                  
                  local classScope = _lune.unwrap( classInfo:get_scope())
                  
                  local fieldList = {}
                  
                  local declClassNode = Nodes.DeclClassNode.create( nodeManager, dummyPos, {classInfo}, Ast.AccessMode.Pub, Parser.Token.new(Parser.TokenKind.Symb, classInfo:get_rawTxt(), dummyPos, false), classInfo:get_rawTxt(), {}, fieldList, nil, {}, classScope, Nodes.ClassInitBlockInfo.new(nil), {}, {}, {} )
                  
                  do
                     local __sorted = {}
                     local __map = classScope:get_symbol2SymbolInfoMap()
                     for __key in pairs( __map ) do
                        table.insert( __sorted, __key )
                     end
                     table.sort( __sorted )
                     for __index, __key in ipairs( __sorted ) do
                        local field = __map[ __key ]
                        do
                           local token = Parser.Token.new(Parser.TokenKind.Symb, field:get_name(), dummyPos, false)
                           do
                              local _switchExp = field:get_kind()
                              if _switchExp == Ast.SymbolKind.Mtd or _switchExp == Ast.SymbolKind.Fun then
                                 local argList = {}
                                 for index, argType in pairs( field:get_typeInfo():get_argTypeInfoList() ) do
                                    local argToken = Parser.Token.new(Parser.TokenKind.Symb, string.format( "arg%d", index), dummyPos, false)
                                    local dummyScope = Ast.Scope.new(nil, false)
                                    local argSym = BuiltinArgSymbolInfo.new(dummyScope, argToken.txt, argType, nil, field:get_typeInfo())
                                    
                                    table.insert( argList, Nodes.DeclArgNode.create( nodeManager, dummyPos, {argType}, argToken, argSym ) )
                                 end
                                 
                                 
                                 local declMethodNode = Nodes.DeclMethodNode.create( nodeManager, dummyPos, {field:get_typeInfo()}, Nodes.DeclFuncInfo.new(Nodes.FuncKind.Mtd, classInfo, token, argList, false, Ast.AccessMode.Pub, nil, field:get_typeInfo():get_retTypeInfoList(), false) )
                                 table.insert( fieldList, declMethodNode )
                              end
                           end
                           
                        end
                     end
                  end
                  
               end
               
         end
      end
      
   end
   
   return nodeManager
end


function convFilter:processRoot( node, opt )

   local nodeManager
   
   
   registerBuiltin(  )
   if self.outputBuiltinFlag then
      nodeManager = self:processBuiltin(  )
   else
    
      nodeManager = node:get_nodeManager()
   end
   
   
   self.scopeMgr:setupScopeParam( self.ast:get_moduleScope() )
   
   Ast.pushProcessInfo( node:get_processInfo() )
   
   self.stream:switchToHeader(  )
   local ifdefname = self.moduleCtrl:getFilePath( self.moduleTypeInfo ):gsub( "/", "_" )
   self:writeln( string.format( [==[
#ifndef __%s__
#define __%s__
       ]==], ifdefname, ifdefname) )
   self.stream:returnToSource(  )
   
   self:writeln( string.format( "// %s", self.streamName) )
   
   self:writeln( "#include <lunescript.h>" )
   self:writeln( string.format( "#include <%s.h>", self.moduleCtrl:getFilePath( node:get_moduleTypeInfo() )) )
   
   local children = node:get_children(  )
   
   self.processMode = ProcessMode.Include
   for __index, importNode in pairs( nodeManager:getImportNodeList(  ) ) do
      filter( importNode, self, node )
   end
   
   
   self:writeln( string.format( "static %s lune_module_globalStemList;", cTypeAnyPP) )
   self:writeln( string.format( "static %s lune_module_path = NULL;", cTypeAnyPP) )
   
   self.processMode = ProcessMode.Prototype
   for __index, declEnumNode in pairs( nodeManager:getDeclEnumNodeList(  ) ) do
      filter( declEnumNode, self, node )
   end
   
   for __index, declFormNode in pairs( nodeManager:getDeclFormNodeList(  ) ) do
      filter( declFormNode, self, node )
   end
   
   for __index, declFuncNode in pairs( nodeManager:getDeclFuncNodeList(  ) ) do
      filter( declFuncNode, self, node )
   end
   
   for __index, declAlgeNode in pairs( nodeManager:getDeclAlgeNodeList(  ) ) do
      filter( declAlgeNode, self, node )
   end
   
   for __index, declClassNode in pairs( nodeManager:getDeclClassNodeList(  ) ) do
      filter( declClassNode, self, node )
   end
   
   for __index, declConstrNode in pairs( nodeManager:getDeclConstrNodeList(  ) ) do
      filter( declConstrNode, self, node )
   end
   
   for __index, declMethodNode in pairs( nodeManager:getDeclMethodNodeList(  ) ) do
      filter( declMethodNode, self, node )
   end
   
   for __index, dddNode in pairs( nodeManager:getExpToDDDNodeList(  ) ) do
      filter( dddNode, self, node )
   end
   
   for __index, litStr in pairs( nodeManager:getLiteralStringNodeList(  ) ) do
      filter( litStr, self, node )
   end
   
   
   self.processMode = ProcessMode.WideScopeVer
   for __index, child in pairs( children ) do
      if child:get_kind() == Nodes.NodeKind.get_DeclVar() then
         filter( child, self, node )
      end
      
   end
   
   for __index, declAlgeNode in pairs( nodeManager:getDeclAlgeNodeList(  ) ) do
      filter( declAlgeNode, self, node )
   end
   
   for __index, declClassNode in pairs( nodeManager:getDeclClassNodeList(  ) ) do
      filter( declClassNode, self, node )
   end
   
   
   self.processMode = ProcessMode.StringFormat
   for __index, litStr in pairs( nodeManager:getLiteralStringNodeList(  ) ) do
      filter( litStr, self, node )
   end
   
   
   self.processMode = ProcessMode.Immediate
   
   self.processedNodeSet = {}
   
   local function procssLiteralCtor( literalNodeList )
   
      for __index, literalNode in pairs( literalNodeList ) do
         self.processingNode = literalNode
         if not _lune._Set_has(self.processedNodeSet, literalNode ) then
            self.accessSymbolSet = Util.OrderedSet.new()
            filter( literalNode, self, node )
            self.processedNodeSet[node]= true
         end
         
      end
      
   end
   procssLiteralCtor( nodeManager:getLiteralListNodeList(  ) )
   procssLiteralCtor( nodeManager:getLiteralArrayNodeList(  ) )
   procssLiteralCtor( nodeManager:getLiteralSetNodeList(  ) )
   procssLiteralCtor( nodeManager:getLiteralMapNodeList(  ) )
   self.processingNode = nil
   
   self.processMode = ProcessMode.Intermediate
   for __index, callNode in pairs( nodeManager:getExpCallNodeList(  ) ) do
      filter( callNode, self, node )
   end
   
   for __index, dddNode in pairs( nodeManager:getExpToDDDNodeList(  ) ) do
      filter( dddNode, self, node )
   end
   
   
   self.processMode = ProcessMode.DefClass
   for __index, declClassNode in pairs( nodeManager:getDeclClassNodeList(  ) ) do
      filter( declClassNode, self, node )
   end
   
   
   self.processMode = ProcessMode.Form
   for __index, declEnumNode in pairs( nodeManager:getDeclEnumNodeList(  ) ) do
      filter( declEnumNode, self, node )
   end
   
   for __index, declAlgeNode in pairs( nodeManager:getDeclAlgeNodeList(  ) ) do
      filter( declAlgeNode, self, node )
   end
   
   for __index, declConstrNode in pairs( nodeManager:getDeclConstrNodeList(  ) ) do
      filter( declConstrNode, self, node )
   end
   
   for __index, declMethodNode in pairs( nodeManager:getDeclMethodNodeList(  ) ) do
      filter( declMethodNode, self, node )
   end
   
   for __index, declFormNode in pairs( nodeManager:getDeclFormNodeList(  ) ) do
      filter( declFormNode, self, node )
   end
   
   for __index, declFuncNode in pairs( nodeManager:getDeclFuncNodeList(  ) ) do
      self.duringDeclFunc = false
      filter( declFuncNode, self, node )
   end
   
   
   self:processInitModule( node )
   
   if self.outputBuiltinFlag then
      self:writeln( '#include "lns_builtinInc.c"' )
   end
   
   
   self.stream:switchToHeader(  )
   self:writeln( "#endif" )
   self.stream:returnToSource(  )
   
   Ast.popProcessInfo(  )
end



function convFilter:processSubfile( node, opt )

end


local function getAccessPrimValFromStem( dddFlag, typeInfo, index )

   local txt = ""
   if dddFlag then
      txt = string.format( ".val.pAny->val.ddd.stemList[ %d ]", index)
   end
   
   
   local expType
   
   
   do
      local enumType = _lune.__Cast( typeInfo:get_srcTypeInfo(), 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         expType = enumType:get_valTypeInfo()
      else
         expType = typeInfo:get_srcTypeInfo()
      end
   end
   
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         txt = txt .. ".val.intVal"
      elseif _switchExp == Ast.builtinTypeReal then
         txt = txt .. ".val.realVal"
      elseif _switchExp == Ast.builtinTypeBool then
         txt = txt .. ".val.boolVal"
      else 
         
            if getValKind( expType ) == ValKind.Any then
               txt = txt .. ".val.pAny"
            end
            
      end
   end
   
   return txt
end

local function getAccessValFromStem( typeInfo )

   local txt
   
   
   local expType
   
   
   do
      local enumType = _lune.__Cast( typeInfo:get_srcTypeInfo(), 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         expType = enumType:get_valTypeInfo()
      else
         expType = typeInfo:get_srcTypeInfo()
      end
   end
   
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         txt = ".val.intVal"
      elseif _switchExp == Ast.builtinTypeReal then
         txt = ".val.realVal"
      elseif _switchExp == Ast.builtinTypeBool then
         txt = ".val.boolVal"
      else 
         
            if getValKind( typeInfo ) == ValKind.Any then
               txt = accessAny
            else
             
               txt = ""
            end
            
      end
   end
   
   return txt
end

function convFilter:processBlockPreProcess( blockNode )

   self:pushIndent(  )
   local anyNum, stemNum, varNum = self.scopeMgr:setupScopeParam( blockNode:get_scope() )
   self:writeln( string.format( "lune_block_t * %s = lune_enter_block( _pEnv, %d, %d, %d );", getBlockName( blockNode:get_scope() ), anyNum, stemNum, varNum) )
   self.currentRoutineInfo:pushDepth(  )
end


function convFilter:processBlockPostProcess(  )

   self.currentRoutineInfo:popDepth(  )
   self:writeln( "lune_leave_block( _pEnv );" )
   self:popIndent(  )
end


function convFilter:pushRoutine( funcType, blockNode )

   self:processBlockPreProcess( blockNode )
   self.currentRoutineInfo = RoutineInfo.new(funcType)
   table.insert( self.routineInfoQueue, self.currentRoutineInfo )
end


function convFilter:popRoutine(  )

   self.currentRoutineInfo = self.routineInfoQueue[#self.routineInfoQueue - 1]
   table.remove( self.routineInfoQueue )
   self:processBlockPostProcess(  )
end


function convFilter:processLoopPreProcess( blockNode )

   self:processBlockPreProcess( blockNode )
end


function convFilter:processLoopPostProcess(  )

   self:processBlockPostProcess(  )
end


function convFilter:processBlockSub( node, opt )

   self.scopeMgr:setupScopeParam( node:get_scope() )
   
   local scope = node:get_scope()
   do
      local __sorted = {}
      local __map = scope:get_closureSymMap()
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local symbol = __map[ __key ]
         do
            local typeTxt, valKind = self.scopeMgr:getCTypeForSym( symbol )
            self:write( string.format( "%s %s = lune_form_closure_var( _pForm, %d )", typeTxt, getSymbolName( symbol ), _lune.unwrap( scope:get_closureSym2NumMap()[symbol])) )
            self:writeln( ";" )
         end
      end
   end
   
   
   local loopFlag = false
   local readyBlock = false
   
   local word = ""
   do
      local _switchExp = node:get_blockKind(  )
      if _switchExp == Nodes.BlockKind.If or _switchExp == Nodes.BlockKind.Elseif then
         word = "{"
      elseif _switchExp == Nodes.BlockKind.Else then
         word = ""
      elseif _switchExp == Nodes.BlockKind.While then
         word = "{"
         loopFlag = true
      elseif _switchExp == Nodes.BlockKind.Repeat then
         word = ""
         loopFlag = true
      elseif _switchExp == Nodes.BlockKind.For then
         word = ""
         loopFlag = true
      elseif _switchExp == Nodes.BlockKind.Apply then
         word = "{"
         loopFlag = true
      elseif _switchExp == Nodes.BlockKind.Foreach then
         word = ""
         loopFlag = true
      elseif _switchExp == Nodes.BlockKind.Macro then
         word = ""
      elseif _switchExp == Nodes.BlockKind.Func then
         readyBlock = true
         word = ""
      elseif _switchExp == Nodes.BlockKind.Default then
         word = ""
      elseif _switchExp == Nodes.BlockKind.Block then
         word = "{"
      elseif _switchExp == Nodes.BlockKind.Macro then
         word = ""
      elseif _switchExp == Nodes.BlockKind.LetUnwrap then
         word = ""
      elseif _switchExp == Nodes.BlockKind.IfUnwrap then
         readyBlock = true
         word = ""
      elseif _switchExp == Nodes.BlockKind.When then
         readyBlock = true
         word = ""
      end
   end
   
   if loopFlag then
      readyBlock = true
   end
   
   
   self:writeln( string.format( "%s // %d", word, node:get_pos().lineNo) )
   
   if not readyBlock then
      self:processBlockPreProcess( node )
   end
   
   local stmtList = node:get_stmtList(  )
   for __index, statement in pairs( stmtList ) do
      filter( statement, self, node )
      self:writeln( "" )
   end
   
   
   if not readyBlock then
      self:processBlockPostProcess(  )
   end
   
   
   if node:get_blockKind(  ) == Nodes.BlockKind.Block then
      self:writeln( "}" )
   end
   
end



function convFilter:processStmtExp( node, opt )

   filter( node:get_exp(), self, node )
   self:write( string.format( "; // %d", node:get_pos().lineNo) )
end



local function getLiteral2Stem( valTxt, typeInfo )

   do
      local _switchExp = typeInfo:get_srcTypeInfo()
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         return string.format( "LUNE_STEM_INT( %s )", valTxt)
      elseif _switchExp == Ast.builtinTypeReal then
         return string.format( "LUNE_STEM_REAL( %s )", valTxt)
      elseif _switchExp == Ast.builtinTypeBool then
         return string.format( "LUNE_STEM_BOOL( %s )", valTxt)
      else 
         
            return "NULL"
      end
   end
   
end

local function getStemTypeId( typeInfo )

   do
      local _switchExp = getOrgTypeInfo( typeInfo )
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         return "lune_stem_type_int"
      elseif _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeReal then
         return "lune_stem_type_real"
      else 
         
            return "lune_stem_type_any"
      end
   end
   
end

function convFilter:getPrepareClosure( funcName, argNum, hasDDD, symList )

   local txt
   
   
   txt = string.format( "lune_func2any( _pEnv, (lune_closure_t *)%s, %d, %s, %d", funcName, argNum, hasDDD, #symList)
   for __index, symbolInfo in pairs( symList ) do
      txt = txt .. ", "
      txt = txt .. self.scopeMgr:symbol2Any( symbolInfo )
   end
   
   txt = txt .. ")"
   return txt
end


function convFilter:getFunc2stem( funcType )

   local argList = funcType:get_argTypeInfoList()
   local hasDDD = #argList > 0 and argList[#argList]:get_kind() == Ast.TypeInfoKind.DDD or false
   
   return self:getPrepareClosure( self.moduleCtrl:getFuncName( funcType ), #funcType:get_argTypeInfoList(), hasDDD, (_lune.unwrap( funcType:get_scope()) ):get_closureSymList() )
end


function convFilter:processSym2stem( symbolInfo )

   local valKind = self.scopeMgr:getSymbolValKind( symbolInfo )
   do
      local _switchExp = valKind
      if _switchExp == ValKind.Any then
         self:write( "LUNE_STEM_ANY( " )
         if symbolInfo:get_kind() == Ast.SymbolKind.Var then
            self:write( "*" )
         end
         
         self:write( getSymbolName( symbolInfo ) )
         self:write( ")" )
         return 
      elseif _switchExp == ValKind.Var then
         self:write( getSymbolName( symbolInfo ) )
         self:write( "->stem" )
         return 
      elseif _switchExp == ValKind.Stem then
         self:write( getSymbolName( symbolInfo ) )
         return 
      end
   end
   
   
   local expType = symbolInfo:get_typeInfo():get_srcTypeInfo()
   do
      local enumType = _lune.__Cast( expType, 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         expType = enumType:get_valTypeInfo()
      end
   end
   
   
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         self:write( "LUNE_STEM_INT( " )
         self:write( "" )
         self:write( self.scopeMgr:getAccessPrimValFromSymbol( symbolInfo ) )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeReal then
         self:write( "LUNE_STEM_REAL( " )
         self:write( self.scopeMgr:getAccessPrimValFromSymbol( symbolInfo ) )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeBool then
         self:write( "LUNE_STEM_BOOL( " )
         self:write( self.scopeMgr:getAccessPrimValFromSymbol( symbolInfo ) )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeStem or _switchExp == Ast.builtinTypeStem_ then
         self:write( getSymbolName( symbolInfo ) )
      else 
         
            do
               local _switchExp = expType:get_kind()
               if _switchExp == Ast.TypeInfoKind.DDD then
                  self:write( "_pDDD" )
               elseif _switchExp == Ast.TypeInfoKind.Func then
                  do
                     local scope = expType:get_scope()
                     if scope ~= nil then
                        self:write( "LUNE_STEM_ANY(" )
                        self:write( self:getFunc2stem( expType ) )
                        self:write( ")" )
                     else
                        Util.err( "illegal func" )
                     end
                  end
                  
               else 
                  
                     self:write( getSymbolName( symbolInfo ) )
               end
            end
            
      end
   end
   
end


function convFilter:processDeclEnum( node, opt )

   local enumType = _lune.unwrap( _lune.__Cast( node:get_expType(), 3, Ast.EnumTypeInfo ))
   
   local enumFullName = self.moduleCtrl:getEnumTypeName( enumType )
   local fullName = self:getFullName( enumType )
   
   local isStrEnum = enumType:get_valTypeInfo():equals( Ast.builtinTypeString )
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         local function process( out2HMode )
         
            local prefix = getOut2HeaderPrefix( out2HMode )
            for index, valName in pairs( node:get_valueNameList() ) do
               local valInfo = _lune.unwrap( enumType:getEnumValInfo( valName.txt ))
               local enumValName = self.moduleCtrl:getEnumValCName( enumType, valName.txt )
               if isStrEnum then
                  self:writeln( string.format( "%s%s %s;", prefix, cTypeAnyP, enumValName) )
               else
                
                  do
                     local _switchExp = out2HMode
                     if _switchExp == Out2HMode.HeaderPub or _switchExp == Out2HMode.SourcePri then
                        local valTxt = string.format( "%s", Ast.getEnumLiteralVal( valInfo:get_val() ))
                        self:writeln( string.format( "#define %s %s", enumValName, valTxt) )
                     end
                  end
                  
               end
               
            end
            
            do
               local _switchExp = out2HMode
               if _switchExp == Out2HMode.HeaderPub or _switchExp == Out2HMode.SourcePri then
                  self:writeln( string.format( "%s%s %s_get__allList( lune_env_t * _pEnv );", prefix, cTypeAnyP, enumFullName) )
                  self:writeln( string.format( "%s%s %s_get__txt( %s _pEnv, %s val );", prefix, cTypeAnyP, enumFullName, cTypeEnvP, getCType( enumType:get_valTypeInfo() )) )
               end
            end
            
         end
         do
            local function processwork( out2HMode )
            
               process( out2HMode )
               
               
            end
            if Ast.isPubToExternal( enumType:get_accessMode() ) then
               self.stream:switchToHeader(  )
               processwork( Out2HMode.HeaderPub )
               self.stream:returnToSource(  )
               
               processwork( Out2HMode.SourcePub )
            else
             
               processwork( Out2HMode.SourcePri )
            end
            
         end
         
         
         
         self:writeln( string.format( "static %s %s_val2NameMap;", cTypeAnyP, enumFullName) )
         self:writeln( string.format( "static %s %s_allList;", cTypeAnyP, enumFullName) )
      elseif _switchExp == ProcessMode.Form then
         if not Ast.isPubToExternal( enumType:get_accessMode() ) then
            self:write( "static " )
         end
         
         self:writeln( string.format( "%s %s_get__allList( lune_env_t * _pEnv )", cTypeAnyP, enumFullName) )
         self:writeln( "{" )
         self:writeln( string.format( "    return %s_allList;", enumFullName) )
         self:writeln( "}" )
         
         if not isStrEnum then
            local typeTxt
            
            if enumType:get_valTypeInfo():get_srcTypeInfo() == Ast.builtinTypeReal then
               typeTxt = "real"
            else
             
               typeTxt = "int"
            end
            
         end
         
         
         if not Ast.isPubToExternal( enumType:get_accessMode() ) then
            self:write( "static " )
         end
         
         self:writeln( string.format( "%s %s_get__txt( %s _pEnv, %s val ) {", cTypeAnyP, enumFullName, cTypeEnvP, getCType( enumType:get_valTypeInfo() )) )
         self:pushIndent(  )
         
         self:write( string.format( "%s _work =  lune_mtd_Map_get( _pEnv, %s, ", cTypeStem, self.moduleCtrl:getEnumVal2NameMapName( enumType )) )
         local workSym = WorkSymbol.new(_lune.unwrap( self.moduleTypeInfo:get_scope()), Ast.AccessMode.Local, "val", enumType:get_valTypeInfo(), Ast.SymbolKind.Arg, SymbolParam.new(getValKind( enumType:get_valTypeInfo() ), 0, getCType( enumType:get_valTypeInfo() )))
         
         self:processSym2stem( workSym )
         self:writeln( ");" )
         self:writeln( string.format( "return _work%s;", accessAny) )
         
         self:popIndent(  )
         self:writeln( "}" )
         
         self:writeln( string.format( "static void init_%s( lune_env_t * _pEnv )", enumFullName) )
         self:writeln( "{" )
         self:pushIndent(  )
         
         local anyVarList = {}
         
         if isStrEnum then
            for index, valName in pairs( node:get_valueNameList() ) do
               local valInfo = _lune.unwrap( enumType:getEnumValInfo( valName.txt ))
               local valTxt = string.format( '"%s"', Ast.getEnumLiteralVal( valInfo:get_val() ))
               local anyVar = self.moduleCtrl:getEnumValCName( enumType, valName.txt )
               table.insert( anyVarList, string.format( "LUNE_STEM_ANY( %s )", anyVar) )
               
               self:writeln( string.format( "%s = lune_litStr2any( _pEnv, %s );", anyVar, valTxt) )
            end
            
         else
          
            for index, valName in pairs( node:get_valueNameList() ) do
               local valInfo = _lune.unwrap( enumType:getEnumValInfo( valName.txt ))
               local valTxt = string.format( '%s', Ast.getEnumLiteralVal( valInfo:get_val() ))
               local anyVar = string.format( "_%s", valName.txt)
               table.insert( anyVarList, anyVar )
               self:write( string.format( "%s %s = ", cTypeStem, anyVar) )
               self:write( getLiteral2Stem( valTxt, enumType:get_valTypeInfo() ) )
               self:writeln( ";" )
            end
            
         end
         
         
         local allListName = string.format( "%s_allList", enumFullName)
         self:write( allListName )
         self:writeln( " = lune_class_List_new( _pEnv );" )
         self:processAddModuleGlobal( string.format( "LUNE_STEM_ANY( %s )", allListName) )
         
         for __index, anyVar in pairs( anyVarList ) do
            self:writeln( string.format( "lune_mtd_List_insert( _pEnv, %s_allList, %s );", enumFullName, anyVar) )
         end
         
         
         local val2NameMapName = string.format( "%s_val2NameMap", enumFullName)
         self:write( val2NameMapName )
         self:writeln( " = lune_class_Map_new( _pEnv );" )
         self:processAddModuleGlobal( string.format( "LUNE_STEM_ANY( %s )", val2NameMapName) )
         
         for index, anyVar in pairs( anyVarList ) do
            self:writeln( string.format( "lune_mtd_Map_add( _pEnv, %s_val2NameMap, %s, ", enumFullName, anyVar) )
            
            self:writeln( string.format( '  LUNE_STEM_ANY( lune_litStr2any( _pEnv, "%s.%s" ) ) );', fullName, node:get_valueNameList()[index].txt) )
         end
         
         
         self:popIndent(  )
         self:writeln( "}" )
      elseif _switchExp == ProcessMode.InitModule then
         self:writeln( string.format( "init_%s( _pEnv );", enumFullName) )
      end
   end
   
end


local function isGenericType( typeInfo )

   if Ast.isGenericType( typeInfo ) then
      return true
   end
   
   do
      local _switchExp = typeInfo:get_kind()
      if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
         if #typeInfo:get_itemTypeInfoList() > 0 then
            return true
         end
         
      end
   end
   
   return false
end

local function processAlgeNewProto( stream, moduleCtrl, typeInfo, valInfo )

   stream:write( string.format( "%s %s( %s _pEnv", cTypeStem, moduleCtrl:getNewAlgeCName( typeInfo, valInfo:get_name() ), cTypeEnvP) )
   
   for index, typeInfo in pairs( valInfo:get_typeList() ) do
      stream:write( string.format( ", %s _val%d", getCType( typeInfo ), index) )
   end
   
   stream:write( ")" )
end

local function processAlgePrototype( stream, moduleCtrl, node )

   local algeType = node:get_algeType()
   local valList = {}
   
   do
      local __sorted = {}
      local __map = algeType:get_valInfoMap()
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local valInfo = __map[ __key ]
         do
            table.insert( valList, valInfo )
         end
      end
   end
   
   
   local function process(  )
   
      stream:writeln( "typedef enum {" )
      stream:pushIndent(  )
      local algeTypeName = moduleCtrl:getAlgeCName( node:get_expType() )
      local enumName = moduleCtrl:getAlgeEnumCName( node:get_expType() )
      
      for index, valInfo in pairs( valList ) do
         if index > 1 then
            stream:writeln( "," )
         end
         
         stream:write( string.format( "%s_%s", enumName, valInfo:get_name()) )
      end
      
      stream:writeln( "" )
      stream:popIndent(  )
      stream:writeln( string.format( "} %s;", enumName) )
      
      for __index, valInfo in pairs( valList ) do
         if #valInfo:get_typeList() > 0 then
            stream:writeln( "typedef struct {" )
            stream:pushIndent(  )
            for index, typeInfo in pairs( valInfo:get_typeList() ) do
               stream:writeln( string.format( "%s _val%d;", getCType( typeInfo ), index) )
            end
            
            stream:popIndent(  )
            stream:writeln( string.format( "} %s;", moduleCtrl:getAlgeValStrCName( node:get_expType(), valInfo:get_name() )) )
         end
         
      end
      
      
      for __index, valInfo in pairs( valList ) do
         if #valInfo:get_typeList() > 0 then
            processAlgeNewProto( stream, moduleCtrl, node:get_expType(), valInfo )
            stream:writeln( ";" )
         end
         
      end
      
   end
   
   do
      local function processwork( out2HMode )
      
         do
            local _switchExp = out2HMode
            if _switchExp == Out2HMode.HeaderPub or _switchExp == Out2HMode.SourcePri then
               process(  )
            end
         end
         
         
         
      end
      if Ast.isPubToExternal( node:get_expType():get_accessMode() ) then
         stream:switchToHeader(  )
         processwork( Out2HMode.HeaderPub )
         stream:returnToSource(  )
         
         processwork( Out2HMode.SourcePub )
      else
       
         processwork( Out2HMode.SourcePri )
      end
      
   end
   
   
   
   stream:writeln( string.format( "static void %s( %s _pEnv );", moduleCtrl:getAlgeInitCName( node:get_expType() ), cTypeEnvP) )
end

local function processAlgeWideScope( stream, moduleCtrl, node )

   local algeType = node:get_algeType()
   local valList = {}
   
   do
      local __sorted = {}
      local __map = algeType:get_valInfoMap()
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local valInfo = __map[ __key ]
         do
            table.insert( valList, valInfo )
         end
      end
   end
   
   
   local algeTypeName = moduleCtrl:getAlgeCName( node:get_expType() )
   
   local function process( out2HMode )
   
      local prefix = getOut2HeaderPrefix( out2HMode )
      
      for index, valInfo in pairs( valList ) do
         if #valInfo:get_typeList() == 0 then
            local varName = moduleCtrl:getAlgeValCName( node:get_expType(), valInfo:get_name() )
            stream:writeln( string.format( "%s%s %s;", prefix, cTypeStem, varName) )
            stream:writeln( string.format( "%s%s %s_any;", prefix, cTypeAny, varName) )
         end
         
      end
      
   end
   
   do
      local function processwork( out2HMode )
      
         do
            local _switchExp = out2HMode
            if _switchExp == Out2HMode.HeaderPub or _switchExp == Out2HMode.SourcePri or _switchExp == Out2HMode.SourcePub then
               process( out2HMode )
            end
         end
         
         
         
      end
      if Ast.isPubToExternal( node:get_expType():get_accessMode() ) then
         stream:switchToHeader(  )
         processwork( Out2HMode.HeaderPub )
         stream:returnToSource(  )
         
         processwork( Out2HMode.SourcePub )
      else
       
         processwork( Out2HMode.SourcePri )
      end
      
   end
   
   
end

local function processAlgeForm( stream, moduleCtrl, node )

   local algeType = node:get_algeType()
   local valList = {}
   
   do
      local __sorted = {}
      local __map = algeType:get_valInfoMap()
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local valInfo = __map[ __key ]
         do
            table.insert( valList, valInfo )
         end
      end
   end
   
   
   stream:writeln( string.format( "static void %s( %s _pEnv ) {", moduleCtrl:getAlgeInitCName( algeType ), cTypeEnvP) )
   stream:pushIndent(  )
   
   local enumName = moduleCtrl:getAlgeEnumCName( algeType )
   for index, valInfo in pairs( valList ) do
      if #valInfo:get_typeList() == 0 then
         local varName = moduleCtrl:getAlgeValCName( algeType, valInfo:get_name() )
         stream:writeln( string.format( "lune_init_alge( &%s, &%s_any, %s_%s );", varName, varName, enumName, valInfo:get_name()) )
      end
      
   end
   
   
   stream:popIndent(  )
   stream:writeln( "}" )
   
   local algeName = moduleCtrl:getAlgeCName( algeType )
   
   for index, valInfo in pairs( valList ) do
      if #valInfo:get_typeList() > 0 then
         local hasAnyFlag = false
         for paramIndex, valType in pairs( valInfo:get_typeList() ) do
            if isStemType( valType ) then
               hasAnyFlag = true
               break
            end
            
         end
         
         
         local valStruct = moduleCtrl:getAlgeValStrCName( algeType, valInfo:get_name() )
         local gcTxt
         
         
         if hasAnyFlag then
            gcTxt = string.format( "lune_gc_alge_%s_%s", algeName, valInfo:get_name())
            stream:writeln( string.format( "static void %s( %s _pEnv, void * pVal ) {", gcTxt, cTypeEnvP) )
            stream:pushIndent(  )
            stream:writeln( string.format( "%s *pWorkVal = (%s *)pVal;", valStruct, valStruct) )
            
            for paramIndex, valType in pairs( valInfo:get_typeList() ) do
               if isStemType( valType ) then
                  stream:writeln( string.format( "lune_decre_ref( _pEnv, pWorkVal->_val%d.val.pAny );", paramIndex) )
               end
               
            end
            
            stream:popIndent(  )
            stream:writeln( "}" )
         else
          
            gcTxt = "NULL"
         end
         
         
         processAlgeNewProto( stream, moduleCtrl, algeType, valInfo )
         stream:writeln( "{" )
         stream:pushIndent(  )
         
         stream:writeln( string.format( "%s pAny = lune_alge_new( _pEnv, %s_%s, sizeof( %s ), %s );", cTypeAnyP, enumName, valInfo:get_name(), valStruct, gcTxt) )
         stream:writeln( string.format( "%s *pVal = pAny->val.alge.pVal;", valStruct) )
         
         for paramIndex, valType in pairs( valInfo:get_typeList() ) do
            if isStemType( valType ) then
               stream:writeln( string.format( "lune_setQ( pVal->_val%d, _val%d );", paramIndex, paramIndex) )
            else
             
               stream:writeln( string.format( "pVal->_val%d = _val%d;", paramIndex, paramIndex) )
            end
            
         end
         
         
         stream:writeln( "return LUNE_STEM_ANY( pAny );" )
         
         stream:popIndent(  )
         stream:writeln( "}" )
      end
      
   end
   
end

function convFilter:processDeclAlge( node, opt )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         processAlgePrototype( self.stream, self.moduleCtrl, node )
      elseif _switchExp == ProcessMode.WideScopeVer then
         processAlgeWideScope( self.stream, self.moduleCtrl, node )
      elseif _switchExp == ProcessMode.Form then
         processAlgeForm( self.stream, self.moduleCtrl, node )
      elseif _switchExp == ProcessMode.InitModule then
         self:writeln( string.format( "%s( _pEnv );", self.moduleCtrl:getAlgeInitCName( node:get_expType() )) )
      end
   end
   
end


function convFilter:processNewAlgeVal( node, opt )

   local valInfo = node:get_valInfo()
   if #valInfo:get_typeList() == 0 then
      local valName = self.moduleCtrl:getAlgeValCName( node:get_algeTypeInfo(), valInfo:get_name() )
      self:write( string.format( "%s", valName) )
   else
    
      self:write( self.moduleCtrl:getNewAlgeCName( node:get_algeTypeInfo(), valInfo:get_name() ) )
      self:write( "( _pEnv" )
      
      for __index, arg in pairs( node:get_paramList() ) do
         self:write( "," )
         filter( arg, self, node )
      end
      
      self:write( ")" )
   end
   
end


function convFilter:outputAlter2MapFunc( stream, alt2Map )

end


local function getMethodTypeTxt( retTypeList )

   if #retTypeList == 1 then
      do
         local _switchExp = retTypeList[1]:get_srcTypeInfo()
         if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
            return "lune_method_int_t"
         elseif _switchExp == Ast.builtinTypeReal then
            return "lune_method_real_t"
         elseif _switchExp == Ast.builtinTypeBool then
            return "lune_method_bool_t"
         end
      end
      
      if getValKind( retTypeList[1] ) == ValKind.Any then
         return "lune_method_any_t"
      end
      
   end
   
   return "lune_method_t"
end

local function processNewConstrProto( stream, moduleCtrl, node, out2HMode, outputBuiltinFlag )

   local className = moduleCtrl:getClassCName( node:get_expType() )
   
   stream:write( getOut2HeaderPrefix( out2HMode ) )
   stream:write( string.format( "%s lune_class_%s_new( %s _pEnv", cTypeAnyP, className, cTypeEnvP) )
   
   if not outputBuiltinFlag then
      local scope = _lune.unwrap( node:get_expType():get_scope())
      local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, scopeAccess ))
      for index, argType in pairs( initFuncType:get_argTypeInfoList() ) do
         stream:write( string.format( ", %s arg%d", getCType( argType ), index) )
      end
      
   end
   
   stream:write( ")" )
end

local function processMethodDeclTxt( stream, moduleCtrl, callFlag, methodTypeInfo, argList )

   if methodTypeInfo:get_rawTxt() ~= "__init" and not callFlag then
      if not methodTypeInfo:get_staticFlag() or not Ast.isPubToExternal( methodTypeInfo:get_accessMode() ) or not Ast.isPubToExternal( methodTypeInfo:get_parentInfo():get_accessMode() ) then
         stream:write( "static " )
      end
      
   end
   
   
   stream:write( string.format( "%s %s( %s _pEnv", getCRetType( methodTypeInfo:get_retTypeInfoList() ), callFlag and moduleCtrl:getCallMethodCName( methodTypeInfo ) or moduleCtrl:getMethodCName( methodTypeInfo ), cTypeEnvP) )
   if methodTypeInfo:get_staticFlag() then
      if isClosure( methodTypeInfo ) then
         stream:write( string.format( ", %s _pForm", cTypeAnyP) )
      end
      
   else
    
      stream:write( string.format( ", %s pObj", cTypeAnyP) )
   end
   
   
   if argList ~= nil then
      for index, argNode in pairs( argList ) do
         do
            local declArgNode = _lune.__Cast( argNode, 3, Nodes.DeclArgNode )
            if declArgNode ~= nil then
               stream:write( string.format( ", %s %s", getCType( declArgNode:get_expType() ), declArgNode:get_name().txt) )
            end
         end
         
      end
      
   else
      for index, arg in pairs( methodTypeInfo:get_argTypeInfoList() ) do
         stream:write( string.format( ", %s arg%d", getCType( arg ), index) )
      end
      
   end
   
   stream:write( ")" )
end

local function processDeclMethodTable( stream, classTypeInfo )

   local function outputField( name, retTypeList )
   
      local methodType = getMethodTypeTxt( retTypeList )
      stream:writeln( string.format( "%s * %s;", methodType, name) )
   end
   
   local function outputVal( scope )
   
      do
         local inherit = scope:get_inherit()
         if inherit ~= nil then
            outputVal( inherit )
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
                  if _switchExp == Ast.SymbolKind.Mtd then
                     if symbolInfo:get_name() ~= "__init" then
                        outputField( symbolInfo:get_name(), symbolInfo:get_typeInfo():get_retTypeInfoList() )
                     end
                     
                  end
               end
               
            end
         end
      end
      
   end
   outputVal( _lune.unwrap( classTypeInfo:get_scope()) )
end

local function processDeclMemberTable( stream, classTypeInfo )

   local function outputVal( scope )
   
      do
         local inherit = scope:get_inherit()
         if inherit ~= nil then
            outputVal( inherit )
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
                  if _switchExp == Ast.SymbolKind.Mbr then
                     if not symbolInfo:get_staticFlag() then
                        stream:writeln( string.format( "%s %s;", getCType( symbolInfo:get_typeInfo() ), symbolInfo:get_name()) )
                     end
                     
                  end
               end
               
            end
         end
      end
      
   end
   stream:writeln( "// member" )
   outputVal( _lune.unwrap( classTypeInfo:get_scope()) )
end

local function hasGC( classTypeInfo )

   do
      local scope = classTypeInfo:get_scope()
      if scope ~= nil then
         if scope:getSymbolInfoField( "_gc", true, scope, scopeAccess ) then
            return true
         end
         
      end
   end
   
   local workInfo = classTypeInfo
   while workInfo:hasBase(  ) do
      workInfo = workInfo:get_baseTypeInfo()
      do
         local scope = classTypeInfo:get_scope()
         if scope ~= nil then
            if scope:getSymbolInfoField( "_gc", true, scope, scopeAccess ) then
               return true
            end
            
         end
      end
      
   end
   
   return false
end

local function processDeclClassPrototype( stream, moduleCtrl, scopeMgr, node )

   local className = moduleCtrl:getClassCName( node:get_expType() )
   
   stream:writeln( string.format( "static void u_mtd_%s__del( lune_env_t * _pEnv, %s pObj );", className, cTypeAnyP) )
   
   if hasGC( node:get_expType() ) then
      stream:writeln( string.format( "static void u_mtd_%s__gc( lune_env_t * _pEnv, %s pObj );", className, cTypeAnyP) )
   end
   
   
   for __index, member in pairs( node:get_memberList() ) do
      local memberName = member:get_name().txt
      if member:get_getterMode() ~= Ast.AccessMode.None then
         local getterType = _lune.unwrap( node:get_scope():getTypeInfoField( string.format( "get_%s", memberName), true, node:get_scope(), scopeAccess ))
         processMethodDeclTxt( stream, moduleCtrl, false, getterType )
         stream:writeln( ";" )
      end
      
      if member:get_setterMode() ~= Ast.AccessMode.None then
         local setterType = _lune.unwrap( node:get_scope():getTypeInfoField( string.format( "set_%s", memberName), true, node:get_scope(), scopeAccess ))
         processMethodDeclTxt( stream, moduleCtrl, false, setterType )
         stream:writeln( ";" )
      end
      
   end
   
end

local function processDefaultCtor( stream, moduleCtrl, scopeMgr, node )

   local className = moduleCtrl:getClassCName( node:get_expType() )
   
   if not node:hasUserInit(  ) then
      stream:write( string.format( "static void u_mtd_%s___init( lune_env_t * _pEnv, %s pAny", className, cTypeAnyP) )
      
      local ctorType = _lune.unwrap( node:get_scope():getTypeInfoField( "__init", true, node:get_scope(), scopeAccess ))
      
      for index, argType in pairs( ctorType:get_argTypeInfoList() ) do
         stream:write( string.format( ", %s _arg%d", getCType( argType ), index) )
      end
      
      stream:writeln( ") {" )
      stream:pushIndent(  )
      
      do
         local baseScope = node:get_scope():get_inherit()
         if baseScope ~= nil then
            local superInitType = _lune.unwrap( baseScope:getTypeInfoField( "__init", true, baseScope, scopeAccess ))
            stream:write( string.format( "u_mtd_%s___init( _pEnv, pAny", moduleCtrl:getClassCName( node:get_expType():get_baseTypeInfo() )) )
            for index, argType in pairs( superInitType:get_argTypeInfoList() ) do
               stream:write( string.format( ", _arg%d", index) )
            end
            
            stream:writeln( ");" )
         end
      end
      
      
      stream:writeln( string.format( "%s * pObj = lune_obj_%s( pAny );", className, className) )
      for index, member in pairs( node:get_memberList() ) do
         if not member:get_staticFlag() then
            local valKind = scopeMgr:getSymbolValKind( member:get_symbolInfo() )
            do
               local _switchExp = valKind
               if _switchExp == ValKind.Stem then
                  stream:writeln( string.format( "lune_setQ( pObj->%s, _arg%d );", member:get_name().txt, index) )
               elseif _switchExp == ValKind.Any then
                  stream:writeln( string.format( "lune_setQ_any( &pObj->%s, _arg%d );", member:get_name().txt, index) )
               elseif _switchExp == ValKind.Prim then
                  stream:writeln( string.format( "pObj->%s = _arg%d;", member:get_name().txt, index) )
               else 
                  
                     Util.err( string.format( "no support -- %s:%s:%d", member:get_name().txt, ValKind:_getTxt( valKind)
                     , 2702) )
               end
            end
            
         end
         
      end
      
      stream:popIndent(  )
      stream:writeln( "}" )
   end
   
end

local function processIFObjDecl( stream, moduleCtrl, classType )

   if classType:hasBase(  ) then
      processIFObjDecl( stream, moduleCtrl, classType:get_baseTypeInfo() )
   end
   
   
   for __index, ifType in pairs( classType:get_interfaceList() ) do
      stream:writeln( string.format( "%s %s;", cTypeAny, moduleCtrl:getClassCName( ifType )) )
   end
   
end

local function processIFObjInit( stream, moduleCtrl, classType, impClassType )

   if classType:hasBase(  ) then
      processIFObjInit( stream, moduleCtrl, classType:get_baseTypeInfo(), impClassType )
   end
   
   
   local className = moduleCtrl:getClassCName( impClassType )
   for __index, ifType in pairs( classType:get_interfaceList() ) do
      local ifName = moduleCtrl:getClassCName( ifType )
      stream:writeln( string.format( "lune_init_if( &pObj->imp.%s, _pEnv, pAny, &lune_if_%s_imp_%s, %s );", ifName, className, ifName, ifName) )
   end
   
end

function convFilter:processDeclClassNodePrototype( node )

   local className = self.moduleCtrl:getClassCName( node:get_expType() )
   
   local kind = node:get_expType():get_kind()
   
   local function process( out2HMode )
   
      self:writeln( string.format( "typedef struct lune_mtd_%s_t {", className) )
      self:pushIndent(  )
      
      if kind == Ast.TypeInfoKind.Class then
         self:writeln( "lune_del_t * _del;" )
         self:writeln( "lune_gc_t * _gc;" )
      end
      
      
      processDeclMethodTable( self.stream, node:get_expType() )
      
      self:popIndent(  )
      self:writeln( string.format( "} lune_mtd_%s_t;", className) )
      
      if kind == Ast.TypeInfoKind.Class then
         self:writeln( string.format( "typedef struct u_if_imp_%s_t {", className) )
         self:pushIndent(  )
         
         processIFObjDecl( self.stream, self.moduleCtrl, node:get_expType() )
         
         self:writeln( string.format( "%s sentinel;", cTypeAny) )
         self:popIndent(  )
         self:writeln( string.format( "} u_if_imp_%s_t;", className) )
      end
      
      
      self:writeln( string.format( "typedef struct %s {", className) )
      self:pushIndent(  )
      
      self:writeln( "lune_type_meta_t * pMeta;" )
      do
         local _switchExp = kind
         if _switchExp == Ast.TypeInfoKind.Class then
            self:writeln( string.format( "u_if_imp_%s_t * pImp;", className) )
            self:writeln( string.format( "lune_mtd_%s_t * pMtd;", className) )
            processDeclMemberTable( self.stream, node:get_expType() )
            self:writeln( "// interface implements" )
            self:writeln( string.format( "u_if_imp_%s_t imp;", className) )
         elseif _switchExp == Ast.TypeInfoKind.IF then
            self:writeln( string.format( "%s pObj;", cTypeAnyP) )
            self:writeln( string.format( "lune_mtd_%s_t * pMtd;", className) )
         end
      end
      
      
      self:popIndent(  )
      self:writeln( string.format( "} %s;", className) )
      
      do
         local _switchExp = kind
         if _switchExp == Ast.TypeInfoKind.Class then
            self:writeln( string.format( [==[#define lune_mtd_%s( OBJ )                     \
                (((%s*)OBJ->val.classVal)->pMtd )]==], className, className) )
            self:writeln( string.format( "#define lune_obj_%s( OBJ ) ((%s*)OBJ->val.classVal)", className, className) )
            self:writeln( string.format( "#define lune_if_%s( OBJ ) ((%s*)OBJ->val.classVal)->pImp", className, className) )
            
            processNewConstrProto( self.stream, self.moduleCtrl, node, out2HMode, self.outputBuiltinFlag )
            self.stream:writeln( ";" )
         elseif _switchExp == Ast.TypeInfoKind.IF then
            self:writeln( string.format( [==[#define lune_mtd_%s( OBJ )                     \
          ((%s*)&OBJ->val.ifVal)->pMtd]==], className, className) )
         end
      end
      
   end
   
   do
      local function processwork( out2HMode )
      
         do
            local _switchExp = out2HMode
            if _switchExp == Out2HMode.HeaderPub or _switchExp == Out2HMode.SourcePri then
               process( out2HMode )
            end
         end
         
         
         
      end
      if Ast.isPubToExternal( node:get_expType():get_accessMode() ) then
         self.stream:switchToHeader(  )
         processwork( Out2HMode.HeaderPub )
         self.stream:returnToSource(  )
         
         processwork( Out2HMode.SourcePub )
      else
       
         processwork( Out2HMode.SourcePri )
      end
      
   end
   
   
   
   if kind == Ast.TypeInfoKind.Class then
      processDeclClassPrototype( self.stream, self.moduleCtrl, self.scopeMgr, node )
      
      if not self.outputBuiltinFlag then
         processDefaultCtor( self.stream, self.moduleCtrl, self.scopeMgr, node )
      end
      
   end
   
end


function convFilter:processDeclClassDef( node )

   local className = self.moduleCtrl:getClassCName( node:get_expType() )
   
   self:writeln( string.format( "static void u_mtd_%s__del( lune_env_t * _pEnv, %s pObj ) {", className, cTypeAnyP) )
   self:pushIndent(  )
   
   if node:get_expType():hasBase(  ) then
      self:writeln( string.format( "u_mtd_%s__del( _pEnv, pObj );", self.moduleCtrl:getClassCName( node:get_expType():get_baseTypeInfo() )) )
   end
   
   
   for __index, member in pairs( node:get_memberList() ) do
      local valKind = self.scopeMgr:getSymbolValKind( member:get_symbolInfo() )
      do
         local _switchExp = valKind
         if _switchExp == ValKind.Stem or _switchExp == ValKind.Any then
            local access
            
            if member:get_expType():get_kind() == Ast.TypeInfoKind.Alternate then
               self:writeln( "lune_decre_ref_alter" )
               access = ""
            else
             
               self:writeln( "lune_decre_ref" )
               if valKind == ValKind.Stem then
                  access = accessAny
               else
                
                  access = ""
               end
               
            end
            
            self:writeln( string.format( "( _pEnv, lune_obj_%s( pObj )->%s%s );", className, member:get_name().txt, access) )
         end
      end
      
   end
   
   
   self:popIndent(  )
   self:writeln( "}" )
   
   processNewConstrProto( self.stream, self.moduleCtrl, node, Out2HMode.SourcePub, self.outputBuiltinFlag )
   self:writeln( "{" )
   self:pushIndent(  )
   
   self:writeln( string.format( "lune_class_new_( _pEnv, %s, pAny, pObj );", className) )
   
   if not self.outputBuiltinFlag then
      self:write( string.format( "u_mtd_%s___init( _pEnv, pAny", className) )
      
      local scope = _lune.unwrap( node:get_expType():get_scope())
      if not self.outputBuiltinFlag then
         local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, scopeAccess ))
         for index, argType in pairs( initFuncType:get_argTypeInfoList() ) do
            self:write( string.format( ", arg%d", index) )
         end
         
      end
      
      self:writeln( ");" )
   end
   
   
   self:writeln( "pObj->pImp = &pObj->imp;" )
   self:writeln( "pObj->imp.sentinel.type = lune_value_type_none;" )
   
   processIFObjInit( self.stream, self.moduleCtrl, node:get_expType(), node:get_expType() )
   
   self:writeln( "return pAny;" )
   self:popIndent(  )
   self:writeln( "}" )
   
   for __index, member in pairs( node:get_memberList() ) do
      local memberName = member:get_name().txt
      if member:get_getterMode() ~= Ast.AccessMode.None then
         local getterType = _lune.unwrap( node:get_scope():getTypeInfoField( string.format( "get_%s", memberName), true, node:get_scope(), scopeAccess ))
         if getterType:get_autoFlag() then
            processMethodDeclTxt( self.stream, self.moduleCtrl, false, getterType )
            self:writeln( "{" )
            self:pushIndent(  )
            self:writeln( string.format( "return lune_obj_%s(pObj)->%s;", className, memberName) )
            self:popIndent(  )
            self:writeln( "}" )
            
            processMethodDeclTxt( self.stream, self.moduleCtrl, true, getterType )
            self:writeln( "{" )
            self:pushIndent(  )
            self:writeln( string.format( "return lune_mtd_%s( pObj )->get_%s( _pEnv, pObj );", className, memberName) )
            self:popIndent(  )
            self:writeln( "}" )
         end
         
      end
      
      if member:get_setterMode() ~= Ast.AccessMode.None then
         local setterType = _lune.unwrap( node:get_scope():getTypeInfoField( string.format( "set_%s", memberName), true, node:get_scope(), scopeAccess ))
         if setterType:get_autoFlag() then
            processMethodDeclTxt( self.stream, self.moduleCtrl, false, setterType )
            self:writeln( "{" )
            self:pushIndent(  )
            local valKind = self.scopeMgr:getSymbolValKind( member:get_symbolInfo() )
            do
               local _switchExp = valKind
               if _switchExp == ValKind.Stem then
                  self:writeln( string.format( 'lune_setq( _pEnv, lune_obj_%s(pObj)->%s, arg1 );', className, memberName) )
               elseif _switchExp == ValKind.Any then
                  self:writeln( string.format( 'lune_setq_any( _pEnv, &lune_obj_%s(pObj)->%s, arg1 );', className, memberName) )
               elseif _switchExp == ValKind.Prim then
                  self:writeln( string.format( "lune_obj_%s(pObj)->%s = arg1;", className, memberName) )
               else 
                  
                     Util.err( string.format( "no support -- %s:%s:%d", member:get_symbolInfo():get_name(), ValKind:_getTxt( valKind)
                     , 2967) )
               end
            end
            
            self:popIndent(  )
            self:writeln( "}" )
            
            processMethodDeclTxt( self.stream, self.moduleCtrl, true, setterType )
            self:writeln( "{" )
            self:pushIndent(  )
            self:writeln( string.format( "lune_mtd_%s( pObj )->set_%s( _pEnv, pObj, arg1 );", className, memberName) )
            self:popIndent(  )
            self:writeln( "}" )
         end
         
      end
      
   end
   
end


local function processInitMethodTable( stream, moduleCtrl, classTypeInfo )

   local function outputField( name, retTypeList )
   
      local methodType = getMethodTypeTxt( retTypeList )
      stream:writeln( string.format( "(%s *)%s,", methodType, name) )
   end
   
   local function outputVal( scope )
   
      do
         local inherit = scope:get_inherit()
         if inherit ~= nil then
            outputVal( inherit )
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
                  if _switchExp == Ast.SymbolKind.Mtd then
                     if symbolInfo:get_name() ~= "__init" then
                        outputField( moduleCtrl:getMethodCName( symbolInfo:get_typeInfo() ), symbolInfo:get_typeInfo():get_retTypeInfoList() )
                     end
                     
                  end
               end
               
            end
         end
      end
      
   end
   outputVal( _lune.unwrap( classTypeInfo:get_scope()) )
end

local function processInitIFMethodTable( stream, moduleCtrl, ifType, classTypeInfo )

   local function outputField( name, retTypeList )
   
      local methodType = getMethodTypeTxt( retTypeList )
      stream:writeln( string.format( "(%s *)%s,", methodType, name) )
   end
   
   local function outputVal( scope, impScope )
   
      do
         local inherit = scope:get_inherit()
         if inherit ~= nil then
            outputVal( inherit, impScope )
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
                  if _switchExp == Ast.SymbolKind.Mtd then
                     if symbolInfo:get_name() ~= "__init" then
                        local impMethodType = _lune.unwrap( impScope:getTypeInfoField( symbolInfo:get_name(), true, impScope, scopeAccess ))
                        outputField( moduleCtrl:getMethodCName( impMethodType ), impMethodType:get_retTypeInfoList() )
                     end
                     
                  end
               end
               
            end
         end
      end
      
   end
   outputVal( _lune.unwrap( ifType:get_scope()), _lune.unwrap( classTypeInfo:get_scope()) )
end

local function processIFMethodDataInit( stream, moduleCtrl, classType, orgClassType )

   local className = moduleCtrl:getClassCName( orgClassType )
   for __index, ifType in pairs( classType:get_interfaceList() ) do
      local ifName = moduleCtrl:getClassCName( ifType )
      stream:writeln( string.format( "static lune_mtd_%s_t lune_if_%s_imp_%s = {", ifName, className, ifName) )
      stream:pushIndent(  )
      
      processInitIFMethodTable( stream, moduleCtrl, ifType, orgClassType )
      
      stream:popIndent(  )
      stream:writeln( "};" )
   end
   
end

local function processClassDataInit( stream, moduleCtrl, classTypeInfo )

   
   local className = moduleCtrl:getClassCName( classTypeInfo )
   
   if not Ast.isPubToExternal( classTypeInfo:get_accessMode() ) then
      stream:write( "static " )
   end
   
   stream:writeln( string.format( 'lune_type_meta_t lune_type_meta_%s = { "%s" };', className, className) )
   
   if not Ast.isPubToExternal( classTypeInfo:get_accessMode() ) then
      stream:write( "static " )
   end
   
   stream:writeln( string.format( "lune_mtd_%s_t lune_mtd_%s = {", className, className) )
   stream:pushIndent(  )
   
   stream:writeln( string.format( "u_mtd_%s__del,", className) )
   
   if hasGC( classTypeInfo ) then
      stream:writeln( string.format( "u_mtd_%s__gc,", className) )
   else
    
      stream:writeln( "NULL," )
   end
   
   
   processInitMethodTable( stream, moduleCtrl, classTypeInfo )
   stream:popIndent(  )
   stream:writeln( "};" )
   
   local function process( out2HMode, symbolInfo )
   
      do
         local _switchExp = out2HMode
         if _switchExp == Out2HMode.HeaderPub or _switchExp == Out2HMode.SourcePub or _switchExp == Out2HMode.SourcePri then
            stream:writeln( string.format( "%s%s %s;", getOut2HeaderPrefix( out2HMode ), getCType( symbolInfo:get_typeInfo() ), moduleCtrl:getClassMemberName( symbolInfo )) )
         end
      end
      
   end
   
   for __index, symbolInfo in pairs( (_lune.unwrap( classTypeInfo:get_scope()) ):get_symbol2SymbolInfoMap() ) do
      if symbolInfo:get_kind() == Ast.SymbolKind.Mbr and symbolInfo:get_staticFlag() then
         do
            local function processwork( out2HMode )
            
               
               process( out2HMode, symbolInfo )
               
               
            end
            if Ast.isPubToExternal( classTypeInfo:get_accessMode() ) and Ast.isPubToExternal( symbolInfo:get_accessMode() ) then
               stream:switchToHeader(  )
               processwork( Out2HMode.HeaderPub )
               stream:returnToSource(  )
               
               processwork( Out2HMode.SourcePub )
            else
             
               processwork( Out2HMode.SourcePri )
            end
            
         end
         
         
      end
      
   end
   
end

function convFilter:processDeclMember( node, opt )

end



function convFilter:processExpMacroExp( node, opt )

   for __index, stmt in pairs( node:get_stmtList() ) do
      filter( stmt, self, node )
      self:writeln( "" )
   end
   
end





function convFilter:outputDeclMacro( name, argNameList, callback )

end


function convFilter:processDeclMacro( node, opt )

end



function convFilter:processExpMacroStat( node, opt )

end



function convFilter:process__func__symbol( has__func__Symbol, classType, funcName )

end


function convFilter:processDeclMethodInfo( declInfo, funcTypeInfo, parent )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         local function processHeader( out2HMode )
         
            if out2HMode == Out2HMode.HeaderPub then
               self:write( "extern " )
            else
             
               self:write( "static " )
            end
            
            if not funcTypeInfo:get_staticFlag() then
               processMethodDeclTxt( self.stream, self.moduleCtrl, true, funcTypeInfo, declInfo:get_argList() )
               self:writeln( ";" )
            else
             
               processMethodDeclTxt( self.stream, self.moduleCtrl, false, funcTypeInfo, declInfo:get_argList() )
               self:writeln( ";" )
            end
            
         end
         
         if funcTypeInfo:get_parentInfo():get_kind() == Ast.TypeInfoKind.Class and funcTypeInfo:get_rawTxt() == "__init" then
            processMethodDeclTxt( self.stream, self.moduleCtrl, false, funcTypeInfo, declInfo:get_argList() )
            self:writeln( ";" )
         else
          
            do
               local function processwork( out2HMode )
               
                  do
                     local _switchExp = out2HMode
                     if _switchExp == Out2HMode.SourcePri then
                        if funcTypeInfo:get_parentInfo():get_kind() == Ast.TypeInfoKind.Class then
                           processMethodDeclTxt( self.stream, self.moduleCtrl, false, funcTypeInfo, declInfo:get_argList() )
                           self:writeln( ";" )
                        end
                        
                     elseif _switchExp == Out2HMode.SourcePub then
                        if funcTypeInfo:get_parentInfo():get_kind() == Ast.TypeInfoKind.Class then
                           if not funcTypeInfo:get_staticFlag() then
                              processMethodDeclTxt( self.stream, self.moduleCtrl, false, funcTypeInfo, declInfo:get_argList() )
                              self:writeln( ";" )
                           end
                           
                        end
                        
                     elseif _switchExp == Out2HMode.HeaderPub then
                        processHeader( out2HMode )
                     end
                  end
                  
                  
                  
               end
               if Ast.isPubToExternal( funcTypeInfo:get_parentInfo():get_accessMode() ) and Ast.isPubToExternal( funcTypeInfo:get_accessMode() ) then
                  self.stream:switchToHeader(  )
                  processwork( Out2HMode.HeaderPub )
                  self.stream:returnToSource(  )
                  
                  processwork( Out2HMode.SourcePub )
               else
                
                  processwork( Out2HMode.SourcePri )
               end
               
            end
            
            
         end
         
      elseif _switchExp == ProcessMode.Form then
         local className = self.moduleCtrl:getClassCName( _lune.unwrap( declInfo:get_classTypeInfo()) )
         do
            local body = declInfo:get_body()
            if body ~= nil then
               local methodNodeToken = _lune.unwrap( declInfo:get_name(  ))
               local methodName = methodNodeToken.txt
               processMethodDeclTxt( self.stream, self.moduleCtrl, false, funcTypeInfo, declInfo:get_argList() )
               self:writeln( "{" )
               
               self:pushIndent(  )
               if not funcTypeInfo:get_staticFlag() then
                  self:writeln( string.format( "%s self = pObj;", cTypeAnyP) )
               end
               
               
               self:pushRoutine( funcTypeInfo, body )
               
               filter( body, self, parent )
               self:popRoutine(  )
               
               self:popIndent(  )
               self:writeln( "}" )
            end
         end
         
         if funcTypeInfo:get_rawTxt() ~= "__init" and not funcTypeInfo:get_staticFlag() then
            processMethodDeclTxt( self.stream, self.moduleCtrl, true, funcTypeInfo, declInfo:get_argList() )
            self:writeln( "{" )
            
            if #funcTypeInfo:get_retTypeInfoList() ~= 0 then
               self:write( "return " )
            end
            
            
            self:write( string.format( "lune_mtd_%s( pObj )->%s( _pEnv, ", className, funcTypeInfo:get_rawTxt()) )
            if _lune.nilacc( declInfo:get_classTypeInfo(), 'get_kind', 'callmtd' ) == Ast.TypeInfoKind.IF then
               self:write( "lune_getImpObj( pObj ) " )
            else
             
               self:write( "pObj " )
            end
            
            for __index, argNode in pairs( declInfo:get_argList() ) do
               do
                  local declArgNode = _lune.__Cast( argNode, 3, Nodes.DeclArgNode )
                  if declArgNode ~= nil then
                     self:write( string.format( ", %s", declArgNode:get_name().txt) )
                  end
               end
               
            end
            
            self:writeln( ");" )
            
            self:writeln( "}" )
         end
         
      end
   end
   
end


function convFilter:processDeclConstr( node, opt )

   self:processDeclMethodInfo( node:get_declInfo(), node:get_expType(), node )
end



function convFilter:processDeclDestr( node, opt )

end


function convFilter:processExpCallSuper( node, opt )

end



function convFilter:processDeclMethod( node, opt )

   self:processDeclMethodInfo( node:get_declInfo(), node:get_expType(), node )
end



function convFilter:processUnwrapSet( node, opt )

end


function convFilter:accessPrimValFromAny( dddFlag, typeInfo, index )

   self:write( getAccessPrimValFromStem( dddFlag, typeInfo, index ) )
end


function convFilter:isStemSym( symbolInfo )

   return self.scopeMgr:getSymbolValKind( symbolInfo ) ~= ValKind.Prim
end


function convFilter:getValKindOfNode( node )

   if #node:get_expTypeList() > 1 then
      return ValKind.Stem
   end
   
   
   local symbolList = node:getSymbolInfo(  )
   if #symbolList > 0 then
      return self.scopeMgr:getSymbolValKind( symbolList[1] )
   end
   
   return getValKind( node:get_expType() )
end


function convFilter:isStemVal( node )

   if #node:get_expTypeList() > 1 then
      return false
   end
   
   
   local symbolList = node:getSymbolInfo(  )
   if #symbolList > 0 then
      return self.scopeMgr:getSymbolValKind( symbolList[1] ) ~= ValKind.Prim
   end
   
   return isStemType( node:get_expType() )
end


function convFilter:accessPrimVal( exp, parent )

   
   do
      local _switchExp = self:getValKindOfNode( exp )
      if _switchExp == ValKind.Prim then
         filter( exp, self, parent )
      elseif _switchExp == ValKind.Stem then
         filter( exp, self, parent )
         self:accessPrimValFromAny( #exp:get_expTypeList() > 1, exp:get_expType(), 0 )
      elseif _switchExp == ValKind.Any then
         filter( exp, self, parent )
      end
   end
   
end


function convFilter:processSym2Any( symbol )

   local valKind = self.scopeMgr:getSymbolValKind( symbol )
   do
      local _switchExp = valKind
      if _switchExp == ValKind.Stem then
         self:write( getSymbolName( symbol ) )
         self:write( accessAny )
      elseif _switchExp == ValKind.Any then
         self:write( getSymbolName( symbol ) )
      else 
         
            Util.err( string.format( "not suppport -- %s, %d", ValKind:_getTxt( valKind)
            , 3558) )
      end
   end
   
end


function convFilter:processVal2stem( node, parent )

   local valKind = self:getValKindOfNode( node )
   
   do
      local _switchExp = valKind
      if _switchExp == ValKind.Stem then
         filter( node, self, parent )
      else 
         
            local expType = node:get_expType():get_srcTypeInfo()
            do
               local enumType = _lune.__Cast( expType, 3, Ast.EnumTypeInfo )
               if enumType ~= nil then
                  expType = enumType:get_valTypeInfo()
               end
            end
            
            
            do
               local _switchExp = expType
               if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                  self:write( "LUNE_STEM_INT( " )
                  filter( node, self, parent )
                  self:write( ")" )
               elseif _switchExp == Ast.builtinTypeReal then
                  self:write( "LUNE_STEM_REAL( " )
                  filter( node, self, parent )
                  self:write( ")" )
               elseif _switchExp == Ast.builtinTypeBool then
                  self:write( "LUNE_STEM_BOOL( " )
                  filter( node, self, parent )
                  self:write( ")" )
               else 
                  
                     do
                        local _switchExp = expType:get_kind()
                        if _switchExp == Ast.TypeInfoKind.DDD then
                           self:write( "_pDDD" )
                        elseif _switchExp == Ast.TypeInfoKind.Func then
                           do
                              local scope = expType:get_scope()
                              if scope ~= nil then
                                 self:write( "LUNE_STEM_ANY(" )
                                 self:write( self:getFunc2stem( expType ) )
                                 self:write( ")" )
                              else
                                 Util.err( "illegal func" )
                              end
                           end
                           
                        else 
                           
                              if valKind == ValKind.Var and getValKind( expType ) == ValKind.Stem then
                                 filter( node, self, parent )
                              else
                               
                                 self:write( "LUNE_STEM_ANY( " )
                                 filter( node, self, parent )
                                 self:write( ")" )
                              end
                              
                        end
                     end
                     
               end
            end
            
      end
   end
   
end


function convFilter:processVal2any( node, parent )

   local valKind = self:getValKindOfNode( node )
   do
      local _switchExp = valKind
      if _switchExp == ValKind.Stem then
         filter( node, self, parent )
         self:write( accessAny )
      elseif _switchExp == ValKind.Any then
         filter( node, self, parent )
      elseif _switchExp == ValKind.Var then
         filter( node, self, parent )
      else 
         
            Util.err( string.format( "not suppport -- %s, %d", ValKind:_getTxt( valKind)
            , 3674) )
      end
   end
   
end


function convFilter:processSetValSingleDirect( parent, node, var, initFlag, expValKind, index, firstMRet, processVal )

   local valKind = self.scopeMgr:getSymbolValKind( var )
   
   local varName = getSymbolName( var )
   if valKind == ValKind.Any and var:get_kind() == Ast.SymbolKind.Var then
      varName = string.format( "(*%s)", varName)
   end
   
   
   local processPrefix = nil
   do
      local fieldNode = _lune.__Cast( node, 3, Nodes.RefFieldNode )
      if fieldNode ~= nil then
         if _lune.nilacc( fieldNode:get_symbolInfo(), 'get_staticFlag', 'callmtd' ) then
            varName = self.moduleCtrl:getClassMemberName( _lune.unwrap( fieldNode:get_symbolInfo()) )
         else
          
            processPrefix = function (  )
            
               local prefixNode = fieldNode:get_prefix()
               local className = self.moduleCtrl:getClassCName( prefixNode:get_expType() )
               self:write( string.format( "lune_obj_%s( ", className) )
               self:processVal2any( prefixNode, fieldNode )
               self:write( ")->" )
            end
         end
         
      end
   end
   
   
   if var:get_symbolId() == -1 then
      if valKind == expValKind then
         self:write( string.format( "%s = ", varName) )
         processVal(  )
         self:write( ";" )
         return 
      end
      
      Util.err( string.format( "illegal %s %s -- %d", ValKind:_getTxt( valKind)
      , ValKind:_getTxt( expValKind)
      , 3728) )
   end
   
   
   do
      local _switchExp = valKind
      if _switchExp == ValKind.Var then
         do
            local _switchExp = expValKind
            if _switchExp == ValKind.Stem then
               if initFlag then
                  self:write( string.format( "lune_setQ( %s->stem, ", varName) )
               else
                
                  self:write( string.format( "lune_setq( _pEnv, %s->stem, ", varName) )
               end
               
               processVal(  )
               self:writeln( " );" )
            elseif _switchExp == ValKind.Any then
               if initFlag then
                  self:write( string.format( "lune_setQ_any( &%s->stem.val.pAny, ", varName) )
               else
                
                  self:write( string.format( "lune_setq_any( _pEnv, &%s->stem.val.pAny, ", varName) )
               end
               
               if processPrefix ~= nil then
                  processPrefix(  )
               end
               
               processVal(  )
               self:write( " );" )
            elseif _switchExp == ValKind.Prim then
               self:write( string.format( "%s->stem", varName) )
               self:write( getAccessValFromStem( var:get_typeInfo() ) )
               self:write( " = " )
               processVal(  )
               self:write( ";" )
            end
         end
         
      elseif _switchExp == ValKind.Stem then
         do
            local _switchExp = expValKind
            if _switchExp == ValKind.Stem then
               if initFlag then
                  self:write( "lune_setQ( " )
               else
                
                  self:write( "lune_setq( _pEnv, " )
               end
               
               if processPrefix ~= nil then
                  processPrefix(  )
               end
               
               self:write( string.format( "%s, ", varName) )
               processVal(  )
               self:write( " );" )
            elseif _switchExp == ValKind.Any then
               if initFlag then
                  self:write( string.format( "lune_setQ_any( &%s->val.pAny, ", varName) )
               else
                
                  self:write( string.format( "lune_setq_any( _pEnv, &%s->val.pAny, ", varName) )
               end
               
               if processPrefix ~= nil then
                  processPrefix(  )
               end
               
               self:write( string.format( "%s, ", varName) )
               processVal(  )
               self:write( " );" )
            elseif _switchExp == ValKind.Prim then
               if processPrefix ~= nil then
                  processPrefix(  )
               end
               
               self:write( self.scopeMgr:getAccessPrimValFromSymbol( var ) )
               self:write( " = " )
               processVal(  )
               self:write( ";" )
            end
         end
         
      elseif _switchExp == ValKind.Any then
         if initFlag then
            self:write( "lune_setQ_any( &" )
         else
          
            self:write( "lune_setq_any( _pEnv, &" )
         end
         
         if processPrefix ~= nil then
            processPrefix(  )
         end
         
         self:write( string.format( "%s, ", varName) )
         processVal(  )
         self:write( " );" )
      else 
         
            if processPrefix ~= nil then
               processPrefix(  )
            end
            
            self:write( string.format( "%s = ", varName) )
            processVal(  )
            self:write( ";" )
      end
   end
   
end


function convFilter:processSymForSetOp( parent, dstKind, dstTypeInfo, symbol )

   local srcKind = self.scopeMgr:getSymbolValKind( symbol )
   local isStemExp = srcKind ~= ValKind.Prim
   
   if dstKind ~= srcKind then
      do
         local _switchExp = dstKind
         if _switchExp == ValKind.Prim then
            self:write( self.scopeMgr:getAccessPrimValFromSymbol( symbol ) )
            return 
         elseif _switchExp == ValKind.Stem then
            self:processSym2stem( symbol )
            return 
         elseif _switchExp == ValKind.Var then
            self:processSym2stem( symbol )
            return 
         elseif _switchExp == ValKind.Any then
            self:processSym2Any( symbol )
            return 
         else 
            
               Util.err( string.format( "not support -- %s", ValKind:_getTxt( dstKind)
               ) )
         end
      end
      
   end
   
   self:write( getSymbolName( symbol ) )
end


local function processToIF( stream, moduleCtrl, expType, process )

   if expType:get_kind() == Ast.TypeInfoKind.IF then
      stream:write( "lune_toIF( _pEnv, " )
      process(  )
      stream:write( string.format( ", &lune_type_meta_%s )", moduleCtrl:getClassCName( expType )) )
      stream:write( accessAny )
   else
    
      process(  )
   end
   
end

local function processGetMRet( stream, moduleCtrl, typeInfo, index )

   local function process(  )
   
      stream:write( string.format( "lune_getMRet( _pEnv, %d )", index) )
      stream:write( getAccessValFromStem( typeInfo ) )
   end
   if typeInfo:get_kind() == Ast.TypeInfoKind.IF then
      processToIF( stream, moduleCtrl, typeInfo, process )
   else
    
      process(  )
   end
   
end

function convFilter:processValForSetOp( parent, dstKind, dstTypeInfo, exp, index, firstMRet )

   local valKind = self:getValKindOfNode( exp )
   
   local function accessVal(  )
   
      if firstMRet then
         processGetMRet( self.stream, self.moduleCtrl, exp:get_expType(), 0 )
      else
       
         filter( exp, self, parent )
      end
      
   end
   
   local function processVal(  )
   
      local setValTxt = ""
      if firstMRet then
         accessVal(  )
      elseif not firstMRet and Nodes.hasMultiValNode( exp ) then
         self:write( "lune_fromDDD( " )
         accessVal(  )
         self:write( accessAny )
         self:write( string.format( ", %d )", index) )
         self:write( getAccessValFromStem( exp:get_expType() ) )
      else
       
         if dstKind == ValKind.Stem then
            self:processVal2stem( exp, parent )
         elseif dstKind == ValKind.Var and valKind ~= ValKind.Stem then
            accessVal(  )
         else
          
            accessVal(  )
         end
         
      end
      
   end
   if dstKind == ValKind.Prim and valKind == ValKind.Stem then
      local expSymList = exp:getSymbolInfo(  )
      
      if #expSymList > 0 then
         self:write( self.scopeMgr:getAccessPrimValFromSymbol( expSymList[1] ) )
      else
       
         processVal(  )
      end
      
   else
    
      processVal(  )
   end
   
end


local function processAlterAccessVal( stream, srcTypeList, dstTypeList )

   if #dstTypeList == 1 then
      if srcTypeList[1]:get_kind() == Ast.TypeInfoKind.Alternate then
         stream:write( getAccessValFromStem( dstTypeList[1] ) )
      end
      
   end
   
end

local function processAlterToActualType( stream, moduleCtrl, fromType, toType, process )

   if fromType:get_kind() == Ast.TypeInfoKind.Alternate then
      if toType:get_kind() == Ast.TypeInfoKind.IF then
         processToIF( stream, moduleCtrl, toType, process )
      else
       
         process(  )
      end
      
   else
    
      process(  )
   end
   
end

function convFilter:processSetValSingle( parent, node, var, initFlag, exp, index, firstMRet )

   local expValKind
   
   if firstMRet then
      expValKind = getValKind( exp:get_expType() )
   else
    
      expValKind = self:getValKindOfNode( exp )
   end
   
   
   self:processSetValSingleDirect( parent, node, var, initFlag, expValKind, index, firstMRet, function (  )
   
      self:processValForSetOp( parent, self.scopeMgr:getSymbolValKind( var ), var:get_typeInfo(), exp, index, firstMRet )
   end )
end


function convFilter:processSetSymSingle( parent, node, var, initFlag, symbol, toIF )

   local function process(  )
   
      self:processSymForSetOp( parent, self.scopeMgr:getSymbolValKind( var ), var:get_typeInfo(), symbol )
   end
   
   self:processSetValSingleDirect( parent, node, var, initFlag, self.scopeMgr:getSymbolValKind( symbol ), 1, false, function (  )
   
      if toIF then
         processToIF( self.stream, self.moduleCtrl, symbol:get_typeInfo(), process )
      else
       
         process(  )
      end
      
   end )
end


function convFilter:processSetValSingleNode( parent, var, initFlag, exp, index, firstMRet )

   local symbolList = var:getSymbolInfo(  )
   if #symbolList > 0 then
      self:processSetValSingle( parent, var, symbolList[1], initFlag, exp, index, firstMRet )
      return 
   end
   
   
   do
      local _switchExp = var:get_kind()
      if _switchExp == Nodes.NodeKind.get_ExpRefItem() then
         do
            local refItemNode = _lune.__Cast( var, 3, Nodes.ExpRefItemNode )
            if refItemNode ~= nil then
               local dstType = refItemNode:get_val():get_expType()
               if dstType:get_kind() == Ast.TypeInfoKind.Map then
                  self:write( "lune_mtd_Map_add( _pEnv, " )
                  self:processVal2any( refItemNode:get_val(), var )
                  self:write( ", " )
                  do
                     local indexNode = refItemNode:get_index()
                     if indexNode ~= nil then
                        self:processVal2stem( indexNode, var )
                     end
                  end
                  
                  self:write( ", " )
                  self:processValForSetOp( parent, ValKind.Stem, dstType:get_itemTypeInfoList()[2], exp, index, firstMRet )
                  self:write( ")" )
               end
               
            end
         end
         
      else 
         
            Util.err( string.format( "not support -- %s", Nodes.getNodeKindName( var:get_kind() )) )
      end
   end
   
end


local DstInfo = {}
DstInfo._name2Val = {}
function DstInfo:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "DstInfo.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function DstInfo._from( val )
   return _lune._AlgeFrom( DstInfo, val )
end

DstInfo.Node = { "Node", {{ func=Nodes.Node._fromMap, nilable=false, child={} }}}
DstInfo._name2Val["Node"] = DstInfo.Node
DstInfo.Symbol = { "Symbol", {{ func=Ast.LowSymbol._fromMap, nilable=false, child={} },{ func=Nodes.Node._fromMap, nilable=true, child={} }}}
DstInfo._name2Val["Symbol"] = DstInfo.Symbol


function convFilter:processSetValToDst( parent, dstList, initFlag, expList, mRetExp )

   local mRetIndex = _lune.nilacc( mRetExp, 'get_index', 'callmtd' )
   
   for index, exp in pairs( expList ) do
      local is1stMRet = index == mRetIndex
      if is1stMRet then
         if mRetExp ~= nil then
            self:write( "lune_setMRet( _pEnv, " )
            filter( mRetExp:get_exp(), self, parent )
            self:write( accessAny )
            self:writeln( ");" )
         end
         
      end
      
      
      if index > #dstList then
         return 
      end
      
      
      if index == #expList then
         
         for dstIndex = index, #dstList do
            local accessIndex
            
            if mRetIndex ~= nil then
               accessIndex = index - mRetIndex
            else
               accessIndex = 0
            end
            
            do
               local _matchExp = dstList[dstIndex]
               if _matchExp[1] == DstInfo.Symbol[1] then
                  local symbolInfo = _matchExp[2][1]
                  local dstNode = _matchExp[2][2]
               
                  self:processSetValSingle( parent, dstNode, symbolInfo, initFlag, exp, accessIndex, is1stMRet and dstIndex == index )
               elseif _matchExp[1] == DstInfo.Node[1] then
                  local dstNode = _matchExp[2][1]
               
                  self:processSetValSingleNode( parent, dstNode, initFlag, exp, accessIndex, is1stMRet and dstIndex == index )
               end
            end
            
            self:writeln( "" )
         end
         
      else
       
         local accessIndex
         
         if mRetIndex ~= nil then
            accessIndex = index - mRetIndex
         else
            accessIndex = 0
         end
         
         do
            local _matchExp = dstList[index]
            if _matchExp[1] == DstInfo.Symbol[1] then
               local symbolInfo = _matchExp[2][1]
               local dstNode = _matchExp[2][2]
            
               self:processSetValSingle( parent, dstNode, symbolInfo, initFlag, exp, accessIndex, is1stMRet )
            elseif _matchExp[1] == DstInfo.Node[1] then
               local dstNode = _matchExp[2][1]
            
               self:processSetValSingleNode( parent, dstNode, initFlag, exp, accessIndex, is1stMRet )
            end
         end
         
         self:writeln( "" )
      end
      
   end
   
end


function convFilter:processSetValToSym( parent, varSymList, initFlag, expList, varNode, mRetExp )

   local varNodeList
   
   do
      local expListNode = _lune.__Cast( varNode, 3, Nodes.ExpListNode )
      if expListNode ~= nil then
         varNodeList = expListNode:get_expList()
      else
         if varNode ~= nil then
            varNodeList = {varNode}
         else
            varNodeList = {}
         end
         
      end
   end
   
   
   local dstList = {}
   for index, symbol in pairs( varSymList ) do
      local node
      
      if index <= #varNodeList then
         node = varNodeList[index]
      else
       
         node = nil
      end
      
      table.insert( dstList, _lune.newAlge( DstInfo.Symbol, {symbol,node}) )
   end
   
   self:processSetValToDst( parent, dstList, initFlag, expList, mRetExp )
end


function convFilter:processSetValToNode( parent, dstNode, expList, mRetExp )

   local dstList = {}
   do
      local expListNode = _lune.__Cast( dstNode, 3, Nodes.ExpListNode )
      if expListNode ~= nil then
         for __index, node in pairs( expListNode:get_expList() ) do
            table.insert( dstList, _lune.newAlge( DstInfo.Node, {node}) )
         end
         
      else
         table.insert( dstList, _lune.newAlge( DstInfo.Node, {dstNode}) )
      end
   end
   
   
   self:processSetValToDst( parent, dstList, false, expList, mRetExp )
end


function convFilter:processDeclVarC( declFlag, var, init0 )

   if declFlag then
      local typeTxt = self.scopeMgr:getCTypeForSym( var )
      self:writeln( string.format( "%s %s;", typeTxt, getSymbolName( var )) )
   end
   
   
   local valKind = self.scopeMgr:getSymbolValKind( var )
   if valKind == ValKind.Prim then
      if init0 then
         self:writeln( string.format( "%s = 0;", getSymbolName( var )) )
      end
      
      return 
   end
   
   
   local initVal
   
   if not init0 or isStemType( var:get_typeInfo() ) then
      
      do
         local symbolInfo = _lune.__Cast( var, 3, Ast.SymbolInfo )
         if symbolInfo ~= nil then
            do
               local _switchExp = valKind
               if _switchExp == ValKind.Any then
                  self:write( "lune_set_block_any" )
                  self:writeln( string.format( "( %s, %d, %s );", getBlockName( var:get_scope() ), getSymbolIndex( symbolInfo ), getSymbolName( var )) )
               elseif _switchExp == ValKind.Stem then
                  self:write( "lune_set_block_stem" )
                  self:writeln( string.format( "( %s, %d, %s );", getBlockName( var:get_scope() ), getSymbolIndex( symbolInfo ), getSymbolName( var )) )
               elseif _switchExp == ValKind.Var then
                  local typeTxt = getStemTypeId( var:get_typeInfo() )
                  self:writeln( string.format( "lune_set_block_var( %s, %d, %s, %s );", getBlockName( var:get_scope() ), getSymbolIndex( symbolInfo ), typeTxt, getSymbolName( var )) )
               end
            end
            
         end
      end
      
   else
    
      initVal = getLiteral2Stem( "0", var:get_typeInfo() )
      
      do
         local symbolInfo = _lune.__Cast( var, 3, Ast.SymbolInfo )
         if symbolInfo ~= nil then
            do
               local _switchExp = valKind
               if _switchExp == ValKind.Stem then
                  self:write( "lune_initVal_stem" )
               elseif _switchExp == ValKind.Any then
                  self:write( "lune_initVal_any" )
               elseif _switchExp == ValKind.Var then
                  self:write( "lune_initVal_var" )
               else 
                  
                     Util.err( string.format( "not support -- %s", ValKind:_getTxt( valKind)
                     ) )
               end
            end
            
            self:writeln( string.format( "( %s, %s, %d, %s );", getSymbolName( var ), getBlockName( var:get_scope() ), getSymbolIndex( symbolInfo ), initVal) )
         else
            self:writeln( string.format( "%s = %s;", getSymbolName( var ), initVal) )
         end
      end
      
   end
   
end


function convFilter:processDeclVarAndSet( varSymList, expListNode )

   for index, var in pairs( varSymList ) do
      local typeTxt, valKind = self.scopeMgr:getCTypeForSym( var )
      if valKind ~= ValKind.Prim then
         local declVarFlag
         
         if varSymList[1]:get_scope() ~= self.ast:get_moduleScope() then
            declVarFlag = true
         else
          
            declVarFlag = false
         end
         
         self:processDeclVarC( declVarFlag, var, valKind ~= ValKind.Var )
      else
       
         if varSymList[1]:get_scope() ~= self.ast:get_moduleScope() then
            self:writeln( string.format( "%s %s;", typeTxt, getSymbolName( var )) )
         end
         
      end
      
   end
   
   
   if expListNode ~= nil then
      self:processSetValToSym( expListNode, varSymList, true, expListNode:get_expList(), nil, expListNode:get_mRetExp() )
   end
   
end


function convFilter:processIfUnwrap( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   
   local expListNode = node:get_expList()
   local workSymList = {}
   for index, varSym in pairs( node:get_varSymList() ) do
      local workSymbol = WorkSymbol.new(varSym:get_scope(), varSym:get_accessMode(), string.format( "_%s", varSym:get_name()), varSym:get_typeInfo():get_nilableTypeInfo(), varSym:get_kind(), SymbolParam.new(ValKind.Stem, 0, cTypeStem))
      table.insert( workSymList, workSymbol )
   end
   
   self:processDeclVarAndSet( workSymList, expListNode )
   
   self:write( "if ( " )
   for index, workSym in pairs( workSymList ) do
      self:write( string.format( "%s.type != lune_stem_type_nil", getSymbolName( workSym )) )
      if index ~= #workSymList then
         self:write( " && " )
      end
      
   end
   
   self:writeln( ") {" )
   
   self:processBlockPreProcess( node:get_block() )
   
   for index, varSym in pairs( node:get_varSymList() ) do
      self:processDeclVarC( true, varSym, false )
      self:processSetSymSingle( node, nil, varSym, true, workSymList[index], false )
      self:writeln( "" )
   end
   
   filter( node:get_block(), self, node )
   
   self:processBlockPostProcess(  )
   self:writeln( "}" )
   
   do
      local _exp = node:get_nilBlock()
      if _exp ~= nil then
         self:writeln( "else {" )
         filter( _exp, self, node )
         self:writeln( "}" )
      end
   end
   
   
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processDeclVar( node, opt )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.WideScopeVer then
         local varSymList = node:get_symbolInfoList()
         for index, var in pairs( varSymList ) do
            local typeTxt = self.scopeMgr:getCTypeForSym( var )
            do
               local _switchExp = var:get_accessMode()
               if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
               else 
                  
                     self:write( "static " )
               end
            end
            
            self:writeln( string.format( "%s %s;", typeTxt, getSymbolName( var )) )
         end
         
         return 
      elseif _switchExp == ProcessMode.InitModule or _switchExp == ProcessMode.Form then
      else 
         
            return 
      end
   end
   
   
   if node:get_syncBlock() then
      self:writeln( "{" )
      self:pushIndent(  )
      for __index, varInfo in pairs( node:get_syncVarList() ) do
         self:writeln( string.format( "_sync_%s", varInfo:get_name().txt) )
      end
      
      self:writeln( "{" )
      self:pushIndent(  )
   end
   
   
   local varSymList = node:get_symbolInfoList()
   
   self:processDeclVarAndSet( varSymList, node:get_expList() )
   
   do
      local _exp = node:get_unwrapBlock()
      if _exp ~= nil then
         self:writeln( "" )
         self:write( "if " )
         local firstFlag = true
         for __index, var in pairs( varSymList ) do
            if var:get_typeInfo():get_nilable() then
               if firstFlag then
                  firstFlag = false
               else
                
                  self:write( " || " )
               end
               
               self:write( string.format( "lune_stem_type_nil == %s.type", getSymbolName( var )) )
            end
            
         end
         
         self:writeln( " {" )
         self:pushIndent(  )
         
         for index, var in pairs( varSymList ) do
            self:writeln( string.format( "local _%s = %s", getSymbolName( var ), getSymbolName( var )) )
         end
         
         self:popIndent(  )
         
         filter( _exp, self, node )
         
         do
            local thenBlock = node:get_thenBlock()
            if thenBlock ~= nil then
               self:writeln( "else {" )
               self:pushIndent(  )
               filter( thenBlock, self, node )
               self:popIndent(  )
            end
         end
         
         
         
         self:writeln( "}" )
      end
   end
   
   
   do
      local _exp = node:get_syncBlock()
      if _exp ~= nil then
         filter( _exp, self, node )
         
         for __index, varInfo in pairs( node:get_syncVarList() ) do
            self:writeln( string.format( "_sync_%s = %s", varInfo:get_name().txt, varInfo:get_name().txt) )
         end
         
         self:popIndent(  )
         self:writeln( "}" )
         
         for __index, varInfo in pairs( node:get_syncVarList() ) do
            self:writeln( string.format( "%s = _sync_%s", varInfo:get_name().txt, varInfo:get_name().txt) )
         end
         
         self:popIndent(  )
         self:writeln( "}" )
      end
   end
   
   
   if node:get_accessMode(  ) == Ast.AccessMode.Pub then
      self:writeln( "" )
      for index, var in pairs( varSymList ) do
         local name = getSymbolName( var )
         self.pubVarName2InfoMap[name] = PubVarInfo.new(node:get_staticFlag(), node:get_accessMode(), node:get_symbolInfoList()[index]:get_mutable(), node:get_typeInfoList()[index])
      end
      
   end
   
end



function convFilter:processWhen( node, opt )

   self:write( "if ( " )
   for index, symPair in pairs( node:get_symPairList() ) do
      self:write( string.format( "%s.type != lune_stem_type_nil", getSymbolName( symPair:get_src() )) )
      if index ~= #node:get_symPairList() then
         self:write( " && " )
      end
      
   end
   
   self:writeln( " ) " )
   self:writeln( "{" )
   
   self:processBlockPreProcess( node:get_block() )
   
   for __index, symPair in pairs( node:get_symPairList() ) do
      local srcSymbol = symPair:get_src()
      local dstSymbol = symPair:get_dst()
      local srcTypeTxt = self.scopeMgr:getCTypeForSym( srcSymbol )
      local dstTypeTxt = self.scopeMgr:getCTypeForSym( dstSymbol )
      if srcTypeTxt ~= dstTypeTxt then
         self:processDeclVarC( true, dstSymbol, false )
         
         self:processSetSymSingle( node, nil, dstSymbol, true, srcSymbol, false )
         self:writeln( "" )
      else
       
         self:writeln( string.format( "%s %s = %s;", dstTypeTxt, getSymbolName( dstSymbol ), getSymbolName( srcSymbol )) )
      end
      
   end
   
   
   filter( node:get_block(), self, node )
   
   self:processBlockPostProcess(  )
   
   do
      local _exp = node:get_elseBlock()
      if _exp ~= nil then
         self:write( "} else {" )
         filter( _exp, self, node )
      end
   end
   
   self:writeln( "}" )
end


function convFilter:processDeclArg( node, opt )

   self:write( getCType( node:get_expType() ) )
   if node:get_symbolInfo():get_hasAccessFromClosure() then
      self:write( ' _' )
   else
    
      self:write( ' ' )
   end
   
   self:write( node:get_name(  ).txt )
end



function convFilter:processDeclArgDDD( node, opt )

   self:write( string.format( "%s _pDDD", cTypeStem) )
end



function convFilter:processExpDDD( node, opt )

end



function convFilter:processExpSubDDD( node, opt )

   self:write( string.format( "lune_createSubDDD( _pEnv, %d, LUNE_STEM_ANY( _pEnv->pMRet ) )", node:get_remainIndex()) )
end



function convFilter:processPrototype( parent, accessMode, needFormVal, name, retType, argList )

   local letTxt = ""
   
   if not Ast.isPubToExternal( accessMode ) then
      letTxt = "static "
   end
   
   self:write( string.format( "%s%s %s( %s _pEnv", letTxt, retType, name, cTypeEnvP ) )
   
   if needFormVal then
      self:write( string.format( ", %s _pForm", cTypeAnyP ) )
   end
   
   
   for index, arg in pairs( argList ) do
      self:write( ", " )
      filter( arg, self, parent )
   end
   
   self:write( " )" )
end


function convFilter:processDeclForm( node, opt )

   local formType = node:get_expType()
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         self:processPrototype( node, formType:get_accessMode(), true, self.moduleCtrl:getCallFormName( formType ), getCRetType( formType:get_retTypeInfoList() ), node:get_argList() )
         
         self:writeln( ";" )
      elseif _switchExp == ProcessMode.Form then
         self:processPrototype( node, formType:get_accessMode(), true, self.moduleCtrl:getCallFormName( formType ), getCRetType( formType:get_retTypeInfoList() ), node:get_argList() )
         
         local function process( prefix )
         
            self:pushIndent(  )
            self:write( prefix )
            for index, arg in pairs( node:get_argList() ) do
               self:write( ", " )
               do
                  local workArg = _lune.__Cast( arg, 3, Nodes.DeclArgNode )
                  if workArg ~= nil then
                     self:write( workArg:get_name().txt )
                  else
                     do
                        local workArg = _lune.__Cast( arg, 3, Nodes.DeclArgDDDNode )
                        if workArg ~= nil then
                           self:write( "_pDDD" )
                        end
                     end
                     
                  end
               end
               
            end
            
            self:writeln( ");" )
            self:popIndent(  )
         end
         
         self:writeln( "{" )
         self:pushIndent(  )
         
         self:writeln( "if lune_isClosure( _pForm ) {" )
         do
            local _switchExp = getCRetType( formType:get_retTypeInfoList() )
            if _switchExp == "void" then
               process( "lune_closure( _pForm )( _pEnv, _pForm" )
            elseif _switchExp == cTypeAny then
               process( "return lune_closure_any( _pForm )( _pEnv, _pForm" )
            elseif _switchExp == cTypeInt then
               process( "return lune_closure_int( _pForm )( _pEnv, _pForm" )
            elseif _switchExp == cTypeReal then
               process( "return lune_closure_real( _pForm )( _pEnv, _pForm" )
            elseif _switchExp == cTypeBool then
               process( "return lune_closure_bool( _pForm )( _pEnv, _pForm" )
            else 
               
                  process( "return lune_closure( _pForm )( _pEnv, _pForm" )
            end
         end
         
         self:writeln( "}" )
         self:writeln( "else {" )
         do
            local _switchExp = getCRetType( formType:get_retTypeInfoList() )
            if _switchExp == "void" then
               process( "lune_func( _pForm )( _pEnv" )
            elseif _switchExp == cTypeAnyP then
               process( "return lune_func_any( _pForm )( _pEnv" )
            elseif _switchExp == cTypeInt then
               process( "return lune_func_int( _pForm )( _pEnv" )
            elseif _switchExp == cTypeReal then
               process( "return lune_func_real( _pForm )( _pEnv" )
            elseif _switchExp == cTypeBool then
               process( "return lune_func_bool( _pForm )( _pEnv" )
            else 
               
                  process( "return lune_func( _pForm )( _pEnv" )
            end
         end
         
         self:writeln( "}" )
         
         self:popIndent(  )
         self:writeln( "}" )
      end
   end
   
end


function convFilter:processDeclFunc( node, opt )

   local declInfo = node:get_declInfo(  )
   local name = self.moduleCtrl:getFuncName( node:get_expType() )
   
   local function processPrototype(  )
   
      self:processPrototype( node, declInfo:get_accessMode(), isClosure( node:get_expType() ), name, getCRetType( node:get_expType():get_retTypeInfoList() ), declInfo:get_argList() )
   end
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Form then
      elseif _switchExp == ProcessMode.Prototype then
         processPrototype(  )
         self:writeln( ";" )
         return 
      else 
         
            do
               local _switchExp = opt.node:get_kind()
               if _switchExp == Nodes.NodeKind.get_Block() or _switchExp == Nodes.NodeKind.get_ExpMacroExp() then
               else 
                  
                     self:write( self:getFunc2stem( node:get_expType() ) )
               end
            end
            
            return 
      end
   end
   
   if self.duringDeclFunc then
      if opt.node:get_kind() == Nodes.NodeKind.get_Block() then
         return 
      end
      
      self:write( self.moduleCtrl:getFuncName( node:get_expType() ) )
      return 
   end
   
   
   local body = declInfo:get_body()
   if  nil == body then
      local _body = body
   
      return 
   end
   
   
   self.duringDeclFunc = true
   
   processPrototype(  )
   self:writeln( "" )
   self:writeln( "{" )
   
   self:pushRoutine( node:get_expType(), body )
   
   for __index, argNode in pairs( node:get_declInfo():get_argList() ) do
      do
         local declArg = _lune.__Cast( argNode, 3, Nodes.DeclArgNode )
         if declArg ~= nil then
            local symbolInfo = declArg:get_symbolInfo()
            if symbolInfo:get_hasAccessFromClosure() then
               local symbolParam = self.scopeMgr:getSymbolParam( symbolInfo )
               self:writeln( string.format( "%s %s;", symbolParam.typeTxt, getSymbolName( symbolInfo )) )
               
               local skipFlag = false
               
               if symbolParam.kind == ValKind.Stem then
                  self:write( "lune_initVal_any(" )
                  self:write( string.format( " %s, %s, %d, ", getSymbolName( symbolInfo ), getBlockName( symbolInfo:get_scope() ), symbolParam.index) )
               else
                
                  self:write( "lune_initVal_var(" )
                  self:write( string.format( " %s, %s, %d, ", getSymbolName( symbolInfo ), getBlockName( symbolInfo:get_scope() ), symbolParam.index) )
                  
                  local valKind
                  
                  if isStemType( symbolInfo:get_typeInfo() ) then
                     valKind = ValKind.Stem
                  else
                   
                     valKind = ValKind.Prim
                  end
                  
                  local workSymbol = WorkSymbol.new(symbolInfo:get_scope(), symbolInfo:get_accessMode(), string.format( "_%s", symbolInfo:get_name()), symbolInfo:get_typeInfo(), symbolInfo:get_kind(), SymbolParam.new(valKind, 0, cTypeStem))
                  
                  self:processSym2stem( workSymbol )
                  self:writeln( ");" )
                  skipFlag = true
               end
               
               if not skipFlag then
                  local workArgName = "_" .. getSymbolName( symbolInfo )
                  if not isStemType( symbolInfo:get_typeInfo() ) then
                     self:write( getLiteral2Stem( workArgName, symbolInfo:get_typeInfo() ) )
                  else
                   
                     self:write( workArgName )
                  end
                  
                  self:writeln( ");" )
               end
               
            end
            
         end
      end
      
   end
   
   
   local breakKind = Nodes.BreakKind.None
   
   self:process__func__symbol( declInfo:get_has__func__Symbol(), node:get_expType():get_parentInfo(), name )
   
   filter( body, self, node )
   
   self:popRoutine(  )
   
   breakKind = body:getBreakKind( Nodes.CheckBreakMode.Normal )
   
   do
      local _switchExp = breakKind
      if _switchExp == Nodes.BreakKind.Return or _switchExp == Nodes.BreakKind.NeverRet then
      else 
         
      end
   end
   
   
   self:writeln( "}" )
   
   local expType = node:get_expType(  )
   if expType:get_accessMode(  ) == Ast.AccessMode.Pub then
      self.pubFuncName2InfoMap[name] = PubFuncInfo.new(declInfo:get_accessMode(  ), node:get_expType(  ))
   end
   
end



function convFilter:processRefType( node, opt )

end



function convFilter:processIf( node, opt )

   local valList = node:get_stmtList(  )
   for index, val in pairs( valList ) do
      if index == 1 then
         self:write( "if ( lune_isCondTrue( " )
         self:processVal2stem( val:get_exp(), node )
         self:write( ") )" )
      elseif val:get_kind() == Nodes.IfKind.ElseIf then
         self:write( "else if ( lune_isCondTrue( " )
         self:processVal2stem( val:get_exp(), node )
         self:write( ") )" )
      else
       
         self:writeln( "else {" )
      end
      
      self:write( " " )
      filter( val:get_block(), self, node )
      self:write( "}" )
   end
   
end



function convFilter:processSwitch( node, opt )

end



function convFilter:processMatch( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   self:write( string.format( "%s _matchExp = ", cTypeAnyP) )
   filter( node:get_val(), self, node )
   self:write( accessAny )
   self:writeln( ";" )
   
   self:writeln( "switch( _matchExp->val.alge.type ) {" )
   local algeType = node:get_algeTypeInfo()
   local enumName = self.moduleCtrl:getAlgeEnumCName( algeType )
   for index, caseInfo in pairs( node:get_caseList() ) do
      local valInfo = caseInfo:get_valInfo()
      self:writeln( string.format( "case %s_%s:", enumName, valInfo:get_name()) )
      self:pushIndent(  )
      self:writeln( "{" )
      self:pushIndent(  )
      if #valInfo:get_typeList() > 0 then
         local structTxt = self.moduleCtrl:getAlgeValStrCName( algeType, valInfo:get_name() )
         self:writeln( string.format( "%s * _pVal = (%s *)_matchExp->val.alge.pVal;", structTxt, structTxt) )
         for paramIndex, paramType in pairs( valInfo:get_typeList() ) do
            local paramName = caseInfo:get_valParamNameList()[paramIndex]
            self:writeln( string.format( "%s %s = _pVal->_val%d;", getCType( paramType ), paramName, paramIndex) )
         end
         
      end
      
      self:popIndent(  )
      
      filter( caseInfo:get_block(), self, node )
      
      self:writeln( "}" )
      self:writeln( "break;" )
      self:popIndent(  )
   end
   
   self:writeln( "}" )
   self:popIndent(  )
   
   self:writeln( "}" )
end



function convFilter:processWhile( node, opt )

end



function convFilter:processRepeat( node, opt )

end



function convFilter:processFor( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   self:writeln( string.format( "%s _to;", cTypeInt) )
   self:writeln( string.format( "%s _inc;", cTypeInt) )
   self:writeln( string.format( "%s %s;", cTypeInt, getSymbolName( node:get_val() )) )
   
   self:processSetValSingle( node, nil, node:get_val(), true, node:get_to(), 0, false )
   self:writeln( "" )
   self:writeln( string.format( "_to = %s;", getSymbolName( node:get_val() )) )
   do
      local _exp = node:get_delta(  )
      if _exp ~= nil then
         self:processSetValToSym( node, {node:get_val()}, true, {_exp} )
         self:writeln( string.format( "_inc = %s;", getSymbolName( node:get_val() )) )
      else
         self:writeln( "_inc = 1;" )
      end
   end
   
   self:processSetValToSym( node, {node:get_val()}, true, {node:get_init()} )
   self:writeln( "" )
   
   self:processLoopPreProcess( node:get_block() )
   
   self:writeln( string.format( "for (; %s <= _to; %s += _inc ) {", getSymbolName( node:get_val() ), getSymbolName( node:get_val() )) )
   self:writeln( "lune_reset_block( _pEnv );" )
   filter( node:get_block(), self, node )
   self:writeln( "}" )
   
   self:processLoopPostProcess(  )
   
   self:popIndent(  )
   self:writeln( "}" )
end



function convFilter:processApply( node, opt )

end



function convFilter:processForeachSetupVal( parent, scope, workTxt, symTxt, symType )

   local symbolInfo = scope:getSymbolInfoChild( symTxt )
   if  nil == symbolInfo then
      local _symbolInfo = symbolInfo
   
      Util.err( string.format( "not found symTxt -- %s", symTxt) )
   end
   
   
   self:processDeclVarC( true, symbolInfo, false )
   
   local srcSymbol = WorkSymbol.new(symbolInfo:get_scope(), symbolInfo:get_accessMode(), workTxt, symbolInfo:get_typeInfo(), symbolInfo:get_kind(), SymbolParam.new(ValKind.Stem, 0, cTypeStem))
   
   self:processSetSymSingle( parent, nil, symbolInfo, true, srcSymbol, true )
end


function convFilter:processPoolForeachSetupVal( parent, loopType, scope, keyToken, valToken )

   
   local valType = loopType:get_itemTypeInfoList()[1]
   local valSymTxt
   
   if loopType:get_kind() == Ast.TypeInfoKind.Set then
      do
         local _exp = keyToken
         if _exp ~= nil then
            valSymTxt = _exp.txt
         else
            Util.err( "keyToken is nil" )
         end
      end
      
   else
    
      do
         local _exp = valToken
         if _exp ~= nil then
            valSymTxt = _exp.txt
         else
            Util.err( "valToken is nil" )
         end
      end
      
   end
   
   
   self:processForeachSetupVal( parent, scope, "_val", valSymTxt, valType )
end


function convFilter:processMapForeachSetupVal( parent, loopType, scope, keyToken, valToken, keyTxt, valTxt )

   if keyToken ~= nil then
      self:processForeachSetupVal( parent, scope, keyTxt, keyToken.txt, loopType:get_itemTypeInfoList()[1] )
   end
   
   self:writeln( "" )
   if valToken ~= nil then
      self:processForeachSetupVal( parent, scope, valTxt, valToken.txt, loopType:get_itemTypeInfoList()[2] )
   end
   
end


local CollectionKind = {}
CollectionKind._name2Val = {}
function CollectionKind:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "CollectionKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function CollectionKind._from( val )
   return _lune._AlgeFrom( CollectionKind, val )
end

CollectionKind.Array = { "Array"}
CollectionKind._name2Val["Array"] = CollectionKind.Array
CollectionKind.ExtArray = { "ExtArray", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
CollectionKind._name2Val["ExtArray"] = CollectionKind.ExtArray
CollectionKind.ExtList = { "ExtList", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
CollectionKind._name2Val["ExtList"] = CollectionKind.ExtList
CollectionKind.ExtMap = { "ExtMap", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
CollectionKind._name2Val["ExtMap"] = CollectionKind.ExtMap
CollectionKind.ExtSet = { "ExtSet", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
CollectionKind._name2Val["ExtSet"] = CollectionKind.ExtSet
CollectionKind.List = { "List"}
CollectionKind._name2Val["List"] = CollectionKind.List
CollectionKind.Map = { "Map"}
CollectionKind._name2Val["Map"] = CollectionKind.Map
CollectionKind.Set = { "Set"}
CollectionKind._name2Val["Set"] = CollectionKind.Set


local function getCollectionKind( typeInfo )

   do
      local _switchExp = typeInfo:get_kind()
      if _switchExp == Ast.TypeInfoKind.List then
         return _lune.newAlge( CollectionKind.List)
      elseif _switchExp == Ast.TypeInfoKind.Array then
         return _lune.newAlge( CollectionKind.Array)
      elseif _switchExp == Ast.TypeInfoKind.Set then
         return _lune.newAlge( CollectionKind.Set)
      elseif _switchExp == Ast.TypeInfoKind.Map then
         return _lune.newAlge( CollectionKind.Map)
      elseif _switchExp == Ast.TypeInfoKind.Ext then
         local extType = _lune.unwrap( _lune.__Cast( typeInfo:get_srcTypeInfo(), 3, Ast.ExtTypeInfo ))
         local luneType = extType:get_luneType()
         do
            local _switchExp = luneType:get_kind()
            if _switchExp == Ast.TypeInfoKind.List then
               return _lune.newAlge( CollectionKind.ExtList, {luneType})
            elseif _switchExp == Ast.TypeInfoKind.Array then
               return _lune.newAlge( CollectionKind.ExtArray, {luneType})
            elseif _switchExp == Ast.TypeInfoKind.Set then
               return _lune.newAlge( CollectionKind.ExtSet, {luneType})
            elseif _switchExp == Ast.TypeInfoKind.Map then
               return _lune.newAlge( CollectionKind.ExtMap, {luneType})
            end
         end
         
      end
   end
   
   Util.err( string.format( "unknown collection type -- %s", typeInfo:getTxt(  )) )
end

function convFilter:processForeach( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   
   self:write( string.format( "%s _obj = ", cTypeAnyP) )
   self:processVal2any( node:get_exp(), node )
   self:writeln( ";" )
   
   local indexSymbol
   
   local loopType = node:get_exp():get_expType()
   
   local collectionKind = getCollectionKind( loopType )
   
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         self:writeln( string.format( "%s _itAny = lune_itList_new( _pEnv, _obj );", cTypeAnyP) )
         do
            local keyToken = node:get_key()
            if keyToken ~= nil then
               local workSymbol = node:get_block():get_scope():getSymbolInfoChild( keyToken.txt )
               if  nil == workSymbol then
                  local _workSymbol = workSymbol
               
                  Util.err( string.format( "not found symbol -- %s", keyToken.txt) )
               end
               
               indexSymbol = workSymbol
               
               if self.scopeMgr:getSymbolValKind( workSymbol ) ~= ValKind.Prim then
                  self:writeln( string.format( "int _%s = 0;", keyToken.txt) )
               else
                
                  self:processDeclVarC( true, workSymbol, true )
               end
               
            else
               indexSymbol = nil
            end
         end
         
         self:writeln( string.format( "%s _val;", cTypeStem) )
      elseif _switchExp == Ast.TypeInfoKind.Set then
         self:writeln( string.format( "%s _itAny = lune_itSet_new( _pEnv, _obj );", cTypeAnyP) )
         indexSymbol = nil
         self:writeln( string.format( "%s _val;", cTypeStem) )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( string.format( "%s _itAny = lune_itMap_new( _pEnv, _obj );", cTypeAnyP) )
         indexSymbol = nil
         self:writeln( "lune_Map_entry_t _entry;" )
      elseif _switchExp == Ast.TypeInfoKind.Ext then
         indexSymbol = nil
         do
            local _matchExp = collectionKind
            if _matchExp[1] == CollectionKind.ExtMap[1] then
               local luneType = _matchExp[2][1]
            
               self:writeln( string.format( "%s _itAny = lune_lua_itMap_new( _pEnv, _obj );", cTypeAnyP) )
               self:writeln( "lune_Map_entry_t _entry;" )
            end
         end
         
      else 
         
            Util.err( string.format( "illegal kind -- %s", Ast.TypeInfoKind:_getTxt( loopType:get_kind())
            ) )
      end
   end
   
   
   self:processLoopPreProcess( node:get_block() )
   
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         self:writeln( "for ( ; lune_itList_hasNext( _pEnv, _itAny, &_val );" )
         self:writeln( "      lune_itList_inc( _pEnv, _itAny ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Set then
         self:writeln( "for ( ; lune_itSet_hasNext( _pEnv, _itAny, &_val );" )
         self:writeln( "      lune_itSet_inc( _pEnv, _itAny ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( "for ( ; lune_itMap_hasNext( _pEnv, _itAny, &_entry );" )
         self:writeln( "      lune_itMap_inc( _pEnv, _itAny ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Ext then
         do
            local _matchExp = collectionKind
            if _matchExp[1] == CollectionKind.ExtMap[1] then
               local luneType = _matchExp[2][1]
            
               self:writeln( "while ( lune_lua_itMap_hasNext( _pEnv, _itAny ) )" )
            end
         end
         
      end
   end
   
   self:writeln( "{" )
   self:pushIndent(  )
   self:writeln( "lune_reset_block( _pEnv );" )
   
   if indexSymbol ~= nil then
      if self.scopeMgr:getSymbolValKind( indexSymbol ) ~= ValKind.Prim then
         self:writeln( string.format( "_%s++;", getSymbolName( indexSymbol )) )
         self:processDeclVarC( true, indexSymbol, true )
         self:processSetValSingleDirect( node, nil, indexSymbol, true, ValKind.Prim, 0, false, function (  )
         
            self:write( string.format( "_%s;", getSymbolName( indexSymbol )) )
         end )
         self:writeln( "" )
      else
       
         self:writeln( string.format( "%s++;", getSymbolName( indexSymbol )) )
      end
      
   end
   
   
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Array then
         self:processPoolForeachSetupVal( node, loopType, node:get_block():get_scope(), node:get_key(), node:get_val() )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:processMapForeachSetupVal( node, loopType, node:get_block():get_scope(), node:get_key(), node:get_val(), "_entry.key", "_entry.val" )
      else 
         
            do
               local _matchExp = collectionKind
               if _matchExp[1] == CollectionKind.ExtMap[1] then
                  local luneType = _matchExp[2][1]
               
                  self:writeln( "lune_lua_itMap_getEntry( _pEnv, _itAny, &_entry );" )
                  self:processMapForeachSetupVal( node, luneType, node:get_block():get_scope(), node:get_key(), node:get_val(), "_entry.key", "_entry.val" )
               end
            end
            
      end
   end
   
   
   filter( node:get_block(), self, node )
   
   self:popIndent(  )
   self:writeln( "}" )
   
   self:processLoopPostProcess(  )
   self:popIndent(  )
   self:writeln( "}" )
end



function convFilter:processForsort( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   
   self:write( string.format( "%s _obj = ", cTypeAnyP) )
   self:processVal2any( node:get_exp(), node )
   self:writeln( ";" )
   
   local loopType = node:get_exp():get_expType()
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Set then
         self:writeln( string.format( "%s _pList = lune_mtd_Map_createKeyList( _pEnv, _obj );", cTypeAnyP) )
         self:writeln( string.format( "lune_mtd_List( _pList )->sort( _pEnv, _pList, %s );", cValNil) )
         self:writeln( string.format( "%s _itAny = lune_itList_new( _pEnv, _pList );", cTypeAnyP) )
         self:writeln( string.format( "%s _val;", cTypeStem) )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( string.format( "%s _pKeyList = lune_mtd_Map_createKeyList( _pEnv, _obj );", cTypeAnyP) )
         self:writeln( string.format( "lune_mtd_List( _pKeyList )->sort( _pEnv, _pKeyList, %s );", cValNil) )
         self:writeln( string.format( "%s _itAny = lune_itList_new( _pEnv, _pKeyList );", cTypeAnyP) )
         self:writeln( string.format( "%s _key;", cTypeStem) )
      else 
         
            Util.err( string.format( "illegal kind -- %s", Ast.TypeInfoKind:_getTxt( loopType:get_kind())
            ) )
      end
   end
   
   
   self:processLoopPreProcess( node:get_block() )
   
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Set then
         self:writeln( "for ( ; lune_itList_hasNext( _pEnv, _itAny, &_val );" )
         self:writeln( "      lune_itList_inc( _pEnv, _itAny ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( "for ( ; lune_itList_hasNext( _pEnv, _itAny, &_key );" )
         self:writeln( "      lune_itList_inc( _pEnv, _itAny ) )" )
      end
   end
   
   self:writeln( "{" )
   
   self:writeln( "lune_reset_block( _pEnv );" )
   
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Set then
         self:processPoolForeachSetupVal( node, loopType, node:get_block():get_scope(), node:get_key(), node:get_val() )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:processMapForeachSetupVal( node, loopType, node:get_block():get_scope(), node:get_key(), node:get_val(), "_key", "lune_mtd_Map_get( _pEnv, _obj, _key )" )
      else 
         
      end
   end
   
   
   filter( node:get_block(), self, node )
   self:writeln( "}" )
   
   self:processLoopPostProcess(  )
   self:writeln( "}" )
   self:popIndent(  )
end



function convFilter:processExpUnwrap( node, opt )

   local function processUnwrap( typeTxt )
   
      do
         local defVal = node:get_default()
         if defVal ~= nil then
            self:write( string.format( "lune_unwrap_%sDefault( ", typeTxt) )
            self:processVal2stem( node:get_exp(), node )
            self:write( "," )
            self:accessPrimVal( defVal, node )
            self:write( ")" )
         else
            self:write( string.format( "lune_unwrap_%s( ", typeTxt) )
            self:processVal2stem( node:get_exp(), node )
            self:write( ")" )
         end
      end
      
   end
   
   do
      local _switchExp = node:get_expType():get_srcTypeInfo()
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         processUnwrap( "int" )
      elseif _switchExp == Ast.builtinTypeReal then
         processUnwrap( "real" )
      else 
         
            do
               local _switchExp = self:getValKindOfNode( node )
               if _switchExp == ValKind.Stem then
                  self:write( "lune_unwrap_stem( " )
               elseif _switchExp == ValKind.Any then
                  self:write( "lune_unwrap_any( " )
               else 
                  
                     Util.err( string.format( "no support -- %d", 5671) )
               end
            end
            
            
            self:processVal2stem( node:get_exp(), node )
            do
               local defVal = node:get_default()
               if defVal ~= nil then
                  self:write( "," )
                  self:processVal2stem( defVal, node )
                  self:write( ")" )
               else
                  self:write( ", NULL )" )
               end
            end
            
      end
   end
   
end


function convFilter:processCreateMRet( retTypeList, expList, parent )

   if expList[1]:get_expType():get_kind() == Ast.TypeInfoKind.DDD and #expList == 1 then
      self:write( "_pDDD" )
      return 
   end
   
   
   self:write( "lune_createMRet" )
   
   local lastExp = expList[#expList]
   self:write( string.format( "( _pEnv, %s, %d", Nodes.hasMultiValNode( lastExp ), #expList) )
   
   for expIndex, exp in pairs( expList ) do
      self:write( ", " )
      self:processVal2stem( exp, parent )
   end
   
   self:write( ")" )
end


local MRetInfo = {}
MRetInfo._name2Val = {}
function MRetInfo:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "MRetInfo.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function MRetInfo._from( val )
   return _lune._AlgeFrom( MRetInfo, val )
end

MRetInfo.DDD = { "DDD", {{ func=Nodes.ExpListNode._fromMap, nilable=false, child={} }}}
MRetInfo._name2Val["DDD"] = MRetInfo.DDD
MRetInfo.Form = { "Form"}
MRetInfo._name2Val["Form"] = MRetInfo.Form
MRetInfo.FormFunc = { "FormFunc", {{ func=Nodes.ExpRefNode._fromMap, nilable=false, child={} }}}
MRetInfo._name2Val["FormFunc"] = MRetInfo.FormFunc
MRetInfo.Format = { "Format", {{ func=_lune._toStr, nilable=false, child={} },{ func=Nodes.ExpListNode._fromMap, nilable=false, child={} }}}
MRetInfo._name2Val["Format"] = MRetInfo.Format
MRetInfo.Func = { "Func", {{ func=Nodes.Node._fromMap, nilable=false, child={} }}}
MRetInfo._name2Val["Func"] = MRetInfo.Func
MRetInfo.Method = { "Method", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
MRetInfo._name2Val["Method"] = MRetInfo.Method


function convFilter:processCallWithMRet( parent, mRetFuncName, retTypeName, funcArgNum, mRetInfo, argList )

   local mRetExp = argList:get_mRetExp()
   if  nil == mRetExp then
      local _mRetExp = mRetExp
   
      return 
   end
   
   
   local function processDeclMRetProto(  )
   
      self:write( string.format( "static %s %s( %s _pEnv", retTypeName, mRetFuncName, cTypeEnvP) )
      
      local function processArgs(  )
      
         for index, argNode in pairs( argList:get_expList() ) do
            if index >= mRetExp:get_index() then
               break
            end
            
            local argType = argNode:get_expType()
            self:write( string.format( ", %s arg%d", getCType( argType ), index) )
         end
         
         self:write( string.format( ", %s pMRet )", cTypeStem) )
      end
      
      do
         local _matchExp = mRetInfo
         if _matchExp[1] == MRetInfo.Method[1] then
            local funcType = _matchExp[2][1]
         
            self:write( string.format( ", %s _pObj", cTypeAnyP) )
            processArgs(  )
         elseif _matchExp[1] == MRetInfo.Form[1] then
         
            self:write( string.format( ", %s _pForm", cTypeAnyP) )
            processArgs(  )
         elseif _matchExp[1] == MRetInfo.FormFunc[1] then
            local funcNode = _matchExp[2][1]
         
            self:write( string.format( ", %s _pForm", cTypeAnyP) )
            processArgs(  )
         elseif _matchExp[1] == MRetInfo.DDD[1] then
            local node = _matchExp[2][1]
         
            processArgs(  )
         elseif _matchExp[1] == MRetInfo.Func[1] then
            local funcNode = _matchExp[2][1]
         
            processArgs(  )
         elseif _matchExp[1] == MRetInfo.Format[1] then
            local format = _matchExp[2][1]
            local node = _matchExp[2][2]
         
            processArgs(  )
         end
      end
      
   end
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Intermediate or _switchExp == ProcessMode.StringFormat then
         processDeclMRetProto(  )
         self:writeln( string.format( "// %d", parent:get_pos().lineNo) )
         self:writeln( "{" )
         
         self:pushIndent(  )
         
         local argTypeList = {}
         
         local function processSetArg( primFlag )
         
            for index, argNode in pairs( argList:get_expList() ) do
               local argType = argNode:get_expType()
               if index == mRetExp:get_index() then
                  self:writeln( string.format( "lune_setMRet( _pEnv, pMRet%s );", accessAny) )
               end
               
               if index >= mRetExp:get_index() then
                  do
                     local _switchExp = argNode:get_kind()
                     if _switchExp == Nodes.NodeKind.get_ExpToDDD() then
                        local toDDDNode = _lune.unwrap( _lune.__Cast( argNode, 3, Nodes.ExpToDDDNode ))
                        self:write( string.format( "%s arg%d = ", cTypeStem, index) )
                        self:write( "lune_createDDD" )
                        local expList = toDDDNode:get_expList():get_expList()
                        local lastExp = expList[#expList]
                        self:write( string.format( "( _pEnv, %s, %d", Nodes.hasMultiValNode( lastExp ), #expList) )
                        for workIndex, exp in pairs( expList ) do
                           self:write( string.format( ", lune_getMRet( _pEnv, %d )", workIndex + index - 2) )
                        end
                        
                        self:write( ")" )
                        table.insert( argTypeList, Ast.builtinTypeDDD )
                     elseif _switchExp == Nodes.NodeKind.get_ExpSubDDD() then
                        self:write( string.format( "%s arg%d = ", cTypeStem, index) )
                        filter( argNode, self, parent )
                     else 
                        
                           do
                              local castNode = _lune.__Cast( argNode, 3, Nodes.ExpCastNode )
                              if castNode ~= nil then
                                 argType = castNode:get_castType()
                              end
                           end
                           
                           
                           local typeTxt
                           
                           if primFlag then
                              typeTxt = getCType( argType )
                              table.insert( argTypeList, argType:get_srcTypeInfo() )
                           else
                            
                              typeTxt = cTypeStem
                              table.insert( argTypeList, Ast.builtinTypeStem )
                           end
                           
                           self:write( string.format( "%s arg%d = ", typeTxt, index) )
                           if argType:get_kind() == Ast.TypeInfoKind.DDD then
                              self:write( "pMRet" )
                           else
                            
                              self:write( string.format( "lune_getMRet( _pEnv, %d )", index - mRetExp:get_index()) )
                           end
                           
                           if primFlag then
                              self:write( getAccessValFromStem( argType ) )
                           else
                            
                              self:write( "->val.pAny" )
                           end
                           
                     end
                  end
                  
                  self:writeln( string.format( "; // %s", argType:getTxt( self:get_typeNameCtrl() )) )
               else
                
                  table.insert( argTypeList, argNode:get_expType() )
               end
               
            end
            
            if retTypeName ~= "void" then
               self:write( "return " )
            end
            
         end
         
         local wroteArgFlag = false
         
         local function processCreateDDD( expList )
         
            self:write( "lune_createDDD" )
            local lastExp = expList[#expList]
            self:write( string.format( "( _pEnv, %s, %d", Nodes.hasMultiValNode( lastExp ), #expList) )
            
            for index = 1, #expList do
               local workExp = expList[index]
               self:write( ", " )
               do
                  local _switchExp = argTypeList[index]
                  if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                     self:write( "LUNE_STEM_INT(" )
                     self:write( string.format( "arg%d )", index) )
                  elseif _switchExp == Ast.builtinTypeReal then
                     self:write( "LUNE_STEM_REAL(" )
                     self:write( string.format( "arg%d )", index) )
                  elseif _switchExp == Ast.builtinTypeBool then
                     self:write( "LUNE_STEM_BOOL(" )
                     self:write( string.format( "arg%d )", index) )
                  else 
                     
                        if getValKind( argTypeList[index] ) == ValKind.Any then
                           self:write( "LUNE_STEM_ANY(" )
                           self:write( string.format( "arg%d )", index) )
                        else
                         
                           self:write( string.format( "arg%d", index) )
                        end
                        
                  end
               end
               
            end
            
         end
         
         do
            local _matchExp = mRetInfo
            if _matchExp[1] == MRetInfo.Method[1] then
               local funcType = _matchExp[2][1]
            
               processSetArg( true )
               self:write( string.format( "%s( _pEnv, _pObj", self.moduleCtrl:getCallMethodCName( funcType )) )
            elseif _matchExp[1] == MRetInfo.Form[1] then
            
               processSetArg( true )
               self:write( "lune_closure( _pForm )( _pEnv, pForm" )
               
               wroteArgFlag = true
               processCreateDDD( argList:get_expList() )
            elseif _matchExp[1] == MRetInfo.FormFunc[1] then
               local funcNode = _matchExp[2][1]
            
               processSetArg( true )
               self:write( string.format( "%s( _pEnv, _pForm", self.moduleCtrl:getCallFormName( funcNode:get_expType() )) )
            elseif _matchExp[1] == MRetInfo.Func[1] then
               local funcNode = _matchExp[2][1]
            
               processSetArg( true )
               local wroteFuncFlag = false
               local builtinFunc = TransUnit.getBuiltinFunc(  )
               
               if funcNode:get_expType() == builtinFunc.lune_print then
                  wroteFuncFlag = true
                  self:write( "lune_print(" )
               end
               
               
               if not wroteFuncFlag then
                  filter( funcNode, self, parent )
                  self:write( "(" )
               end
               
               self:write( " _pEnv" )
            elseif _matchExp[1] == MRetInfo.DDD[1] then
               local expListNode = _matchExp[2][1]
            
               processSetArg( true )
               wroteArgFlag = true
               processCreateDDD( expListNode:get_expList() )
            elseif _matchExp[1] == MRetInfo.Format[1] then
               local format = _matchExp[2][1]
               local expListNode = _matchExp[2][2]
            
               processSetArg( true )
               wroteArgFlag = true
               self:write( "u_mtd_string_format( _pEnv, " )
               self:write( "lune_litStr2any( _pEnv," )
               self:write( format )
               self:write( "), " )
               processCreateDDD( expListNode:get_expList() )
               self:write( ")" )
            end
         end
         
         
         if not wroteArgFlag then
            for index = 1, funcArgNum do
               self:write( string.format( ", arg%d", index) )
            end
            
         end
         
         
         self:popIndent(  )
         self:writeln( ");" )
         self:writeln( "}" )
      elseif _switchExp == ProcessMode.Prototype then
         processDeclMRetProto(  )
         self:writeln( string.format( "; // %d", argList:get_pos().lineNo) )
      end
   end
   
end


local function getMRetFuncName( node )

   return string.format( "lune_call_mret_%d", node:get_id())
end

function convFilter:processExpToDDD( node, opt )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Intermediate or _switchExp == ProcessMode.Prototype then
         do
            local mRetExp = node:get_expList():get_mRetExp()
            if mRetExp ~= nil then
               if mRetExp:get_index() > 0 then
                  self:processCallWithMRet( node, getMRetFuncName( node ), cTypeStem, #node:get_expList():get_expList(), _lune.newAlge( MRetInfo.DDD, {node:get_expList()}), node:get_expList() )
               end
               
            end
         end
         
         return 
      end
   end
   
   
   local expList = node:get_expList():get_expList()
   
   do
      local mRetExp = node:get_expList():get_mRetExp()
      if mRetExp ~= nil then
         self:write( string.format( "%s( _pEnv", getMRetFuncName( node )) )
         for index, exp in pairs( expList ) do
            if index > mRetExp:get_index() then
               break
            end
            
            self:write( ", " )
            filter( exp, self, node )
         end
         
      else
         self:write( "lune_createDDD" )
         local lastExp = expList[#expList]
         self:write( string.format( "( _pEnv, %s, %d", Nodes.hasMultiValNode( lastExp ), #expList) )
         
         for __index, exp in pairs( expList ) do
            self:write( ", " )
            self:processVal2stem( exp, node )
         end
         
      end
   end
   
   self:write( ")" )
end


function convFilter:processCallArgList( funcArgTypeList, expListNode )

   if expListNode ~= nil then
      local expList = expListNode:get_expList()
      for index, funcArgType in pairs( funcArgTypeList ) do
         if index ~= 1 then
            self:write( ", " )
         end
         
         if #expList >= index then
            if funcArgType:get_kind() == Ast.TypeInfoKind.DDD then
               if expList[index]:get_kind() == Nodes.NodeKind.get_Abbr() then
                  self:write( "lune_global.ddd0" )
               else
                
                  filter( expList[index], self, expListNode )
               end
               
               return 
            else
             
               
               if isStemType( funcArgType ) then
                  self:processVal2stem( expList[index], expListNode )
               else
                
                  filter( expList[index], self, expListNode )
               end
               
            end
            
         else
          
            if funcArgType:get_kind() == Ast.TypeInfoKind.DDD then
               self:write( "lune_global.ddd0" )
            else
             
               self:write( cValNone )
            end
            
         end
         
      end
      
   end
   
end


function convFilter:processExpNew( node, opt )

   local classFullName = self:getFullName( node:get_symbol():get_expType() )
   self:write( string.format( "lune_class_%s_new( _pEnv", classFullName) )
   
   do
      local _exp = node:get_argList()
      if _exp ~= nil then
         self:write( ", " )
         self:processCallArgList( node:get_ctorTypeInfo():get_argTypeInfoList(), _exp )
      end
   end
   
   self:write( ")" )
end



function convFilter:processCall( funcType, setArgFlag, argList )

   if not setArgFlag then
      self:write( "_pEnv" )
      
      do
         local scope = funcType:get_scope()
         if scope ~= nil then
            
            if #scope:get_closureSymList() > 0 then
               self:write( ", " )
               self:write( self:getPrepareClosure( "NULL", 0, false, scope:get_closureSymList() ) )
            end
            
         end
      end
      
   end
   
   
   if argList ~= nil then
      local expList = {}
      for __index, expNode in pairs( argList:get_expList() ) do
         if expNode:get_expType():get_kind() ~= Ast.TypeInfoKind.Abbr then
            table.insert( expList, expNode )
         end
         
      end
      
      
      if #expList > 0 then
         self:write( ", " )
         self:processCallArgList( funcType:get_argTypeInfoList(), argList )
      end
      
   end
   
   self:write( " )" )
end


function convFilter:processDeclClass( node, opt )

   local classType = node:get_expType()
   local className = self.moduleCtrl:getClassCName( classType )
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         self:processDeclClassNodePrototype( node )
      elseif _switchExp == ProcessMode.WideScopeVer then
         do
            local _switchExp = classType:get_kind()
            if _switchExp == Ast.TypeInfoKind.Class then
               processIFMethodDataInit( self.stream, self.moduleCtrl, classType, classType )
               
               processClassDataInit( self.stream, self.moduleCtrl, classType )
            elseif _switchExp == Ast.TypeInfoKind.IF then
               self:writeln( string.format( 'lune_type_meta_t lune_type_meta_%s = { "%s" };', className, className) )
            end
         end
         
      elseif _switchExp == ProcessMode.DefClass then
         if classType:get_kind() == Ast.TypeInfoKind.Class then
            self:processDeclClassDef( node )
         end
         
      elseif _switchExp == ProcessMode.Form or _switchExp == ProcessMode.InitModule then
         
         do
            local initBlockNode = node:get_initBlock():get_func()
            if initBlockNode ~= nil then
               self:write( string.format( "%s(", self.moduleCtrl:getMethodCName( initBlockNode:get_expType() )) )
               self:processCall( initBlockNode:get_expType(), false, nil )
               self:writeln( ";" )
            end
         end
         
      end
   end
   
end



function convFilter:processExpCall( node, opt )

   local funcType = node:get_func():get_expType()
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Intermediate or _switchExp == ProcessMode.Prototype then
         do
            local argList = node:get_argList()
            if argList ~= nil then
               
               local funcNode = node:get_func()
               
               local mRetInfo
               
               do
                  local _switchExp = funcNode:get_expType():get_kind()
                  if _switchExp == Ast.TypeInfoKind.Method then
                     mRetInfo = _lune.newAlge( MRetInfo.Method, {funcType})
                  elseif _switchExp == Ast.TypeInfoKind.Form then
                     mRetInfo = _lune.newAlge( MRetInfo.Form)
                  elseif _switchExp == Ast.TypeInfoKind.FormFunc then
                     mRetInfo = _lune.newAlge( MRetInfo.FormFunc, {_lune.unwrap( _lune.__Cast( node:get_func(), 3, Nodes.ExpRefNode ))})
                  else 
                     
                        mRetInfo = _lune.newAlge( MRetInfo.Func, {node:get_func()})
                  end
               end
               
               
               self:processCallWithMRet( node, getMRetFuncName( node ), getCRetType( node:get_expTypeList() ), #funcType:get_argTypeInfoList(), mRetInfo, argList )
            end
         end
         
         return 
      end
   end
   
   
   local classTypeInfo = getBelongClassType( node:get_func() )
   
   local function process(  )
   
      if funcType:get_kind() == Ast.TypeInfoKind.Form then
         self:write( 'lune_call_form( _pEnv, ' )
         self:processVal2any( node:get_func(), node )
         do
            local argList = node:get_argList()
            if argList ~= nil then
               if #argList:get_expList() > 0 then
                  self:write( ', ' )
                  self:processCallArgList( funcType:get_argTypeInfoList(), argList )
               end
               
            else
               self:write( ', lune_global.ddd0' )
            end
         end
         
         self:write( ' )' )
         
         return 
      end
      
      
      local wroteFuncFlag = false
      local setArgFlag = false
      
      local function fieldCall(  )
      
         
         local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
         if  nil == fieldNode then
            local _fieldNode = fieldNode
         
            return true
         end
         
         local prefixNode = fieldNode:get_prefix()
         
         local prefixType = prefixNode:get_expType()
         
         local function processEnumAlge(  )
         
         end
         
         if node:get_nilAccess() then
         else
          
            
            do
               local _switchExp = prefixType:get_kind()
               if _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.Alge then
                  processEnumAlge(  )
               elseif _switchExp == Ast.TypeInfoKind.Class then
                  
                  if prefixType:equals( Ast.builtinTypeString ) then
                     wroteFuncFlag = true
                     setArgFlag = true
                     self:write( string.format( "u_mtd_string_%s( _pEnv, ", fieldNode:get_field().txt) )
                     filter( prefixNode, self, fieldNode )
                  end
                  
               end
            end
            
         end
         
         return true
      end
      
      if not fieldCall(  ) then
         return 
      end
      
      
      do
         local refNode = _lune.__Cast( node:get_func(), 3, Nodes.ExpRefNode )
         if refNode ~= nil then
            local builtinFunc = TransUnit.getBuiltinFunc(  )
            
            do
               local _switchExp = refNode:get_expType()
               if _switchExp == builtinFunc.lune_print then
                  wroteFuncFlag = true
                  self:write( "lune_print(" )
               end
            end
            
         end
      end
      
      
      if not wroteFuncFlag then
         local funcSymList = node:get_func():getSymbolInfo(  )
         if #funcSymList > 0 then
            local builtinFunc = TransUnit.getBuiltinFunc(  )
            local funcSym = funcSymList[1]:getOrg(  )
            if funcSym == builtinFunc.string_format_sym then
               wroteFuncFlag = true
               self:write( "u_mtd_string_format(" )
            end
            
         end
         
      end
      
      
      if not wroteFuncFlag then
         do
            local _switchExp = funcType:get_kind()
            if _switchExp == Ast.TypeInfoKind.Method then
               
               do
                  local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
                  if fieldNode ~= nil then
                     self:write( string.format( "%s( _pEnv, ", self.moduleCtrl:getCallMethodCName( funcType )) )
                     self:processVal2any( fieldNode:get_prefix(), fieldNode )
                  end
               end
               
               wroteFuncFlag = true
               setArgFlag = true
            elseif _switchExp == Ast.TypeInfoKind.FormFunc then
               self:write( string.format( "%s( _pEnv, ", self.moduleCtrl:getCallFormName( funcType )) )
               self:processVal2any( node:get_func(), node )
               wroteFuncFlag = true
               setArgFlag = true
            end
         end
         
      end
      
      
      if not wroteFuncFlag then
         filter( node:get_func(), self, node )
         
         self:write( "( " )
      end
      
      
      self:processCall( funcType, setArgFlag, node:get_argList() )
   end
   
   local function call(  )
   
      local isMret = false
      do
         local argList = node:get_argList()
         if argList ~= nil then
            do
               local mRetExp = argList:get_mRetExp()
               if mRetExp ~= nil then
                  isMret = true
                  self:write( string.format( "%s( _pEnv", getMRetFuncName( node )) )
                  
                  local funcNode = node:get_func()
                  do
                     local _switchExp = funcNode:get_expType():get_kind()
                     if _switchExp == Ast.TypeInfoKind.Method then
                        do
                           local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
                           if fieldNode ~= nil then
                              self:write( ", " )
                              self:processVal2any( fieldNode:get_prefix(), fieldNode )
                           end
                        end
                        
                     elseif _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
                        self:write( ", " )
                        self:processVal2any( node:get_func(), node )
                     end
                  end
                  
                  for index, argNode in pairs( argList:get_expList() ) do
                     if index <= mRetExp:get_index() then
                        self:write( ", " )
                        filter( argNode, self, argList )
                     end
                     
                  end
                  
                  self:write( ")" )
               end
            end
            
         end
      end
      
      if not isMret then
         process(  )
      end
      
      
      processAlterAccessVal( self.stream, funcType:get_retTypeInfoList(), node:get_expTypeList() )
   end
   
   local retTypeInfoList = funcType:get_retTypeInfoList()
   if #retTypeInfoList == 1 then
      processAlterToActualType( self.stream, self.moduleCtrl, retTypeInfoList[1], node:get_expType(), call )
   else
    
      call(  )
   end
   
end



function convFilter:processExpAccessMRet( node, opt )

   processGetMRet( self.stream, self.moduleCtrl, node:get_expType(), node:get_index() - 1 )
end



function convFilter:processExpList( node, opt )

   local expList = node:get_expList(  )
   
   for index, exp in pairs( expList ) do
      if exp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
         break
      end
      
      if index > 1 then
         self:write( ", " )
      end
      
      filter( exp, self, node )
   end
   
end



function convFilter:processExpOp1( node, opt )

   local op = node:get_op().txt
   
   do
      local _switchExp = op
      if _switchExp == "~" or _switchExp == "+" or _switchExp == "-" then
         self:write( op )
         self:accessPrimVal( node:get_exp(), node )
      elseif _switchExp == "not" then
         self:write( "lune_op_not( _pEnv, " )
         self:processVal2stem( node:get_exp(), node )
         self:write( ")" )
      else 
         
            Util.err( string.format( "not support op -- %s", op) )
      end
   end
   
end



function convFilter:processExpMultiTo1( node, opt )

   self:write( "lune_fromDDD( " )
   filter( node:get_exp(), self, node )
   self:write( accessAny )
   self:write( ", 0 )" )
   if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.DDD and Ast.isNumberType( node:get_expType():get_srcTypeInfo():get_nonnilableType() ) then
      self:write( accessAny )
   else
    
      self:write( getAccessValFromStem( node:get_exp():get_expType() ) )
   end
   
end


function convFilter:processExpCast( node, opt )

   local exp = node:get_exp()
   local expType = exp:get_expType()
   local nodeExpType = node:get_expType()
   local castType = node:get_castType()
   do
      local _switchExp = node:get_castKind()
      if _switchExp == Nodes.CastKind.Implicit then
         do
            local _switchExp = castType:get_kind()
            if _switchExp == Ast.TypeInfoKind.IF then
               if expType:get_kind() == Ast.TypeInfoKind.Class then
                  self:write( string.format( "lune_getIF( _pEnv, &lune_if_%s( ", self.moduleCtrl:getClassCName( expType )) )
                  self:processVal2any( node:get_exp(), node )
                  self:write( string.format( ")->%s )", self.moduleCtrl:getClassCName( castType )) )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.FormFunc then
               self:write( self:getFunc2stem( expType ) )
            elseif _switchExp == Ast.TypeInfoKind.Form then
               self:write( self:getFunc2stem( expType ) )
            else 
               
                  filter( exp, self, node )
            end
         end
         
      elseif _switchExp == Nodes.CastKind.Force then
         if isStemType( expType ) then
            do
               local _switchExp = castType
               if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                  self:write( "lune_stem2int( " )
                  filter( exp, self, node )
                  self:write( ")" )
               elseif _switchExp == Ast.builtinTypeReal then
                  self:write( "lune_stem2real( " )
                  filter( exp, self, node )
                  self:write( ")" )
               elseif _switchExp == Ast.builtinTypeBool then
                  self:write( "lune_stem2bool( " )
                  filter( exp, self, node )
                  self:write( ")" )
               else 
                  
                     
                     if getValKind( castType ) == ValKind.Any then
                        self:processVal2any( exp, node )
                     else
                      
                        filter( exp, self, node )
                     end
                     
               end
            end
            
         else
          
            filter( exp, self, node )
         end
         
      elseif _switchExp == Nodes.CastKind.Normal then
         error( "not support cast" )
      end
   end
   
end



function convFilter:processExpParen( node, opt )

   if #node:get_exp():get_expTypeList() == 1 then
      self:write( "(" )
      self:accessPrimVal( node:get_exp(), node )
      self:write( " )" )
   else
    
      processToIF( self.stream, self.moduleCtrl, node:get_expType(), function (  )
      
         self:accessPrimVal( node:get_exp(), node )
      end )
   end
   
end



function convFilter:processWrapForm2Func( funcType )

   self:write( string.format( "static %s _wrap_%s_%d( %s _pEnv, %s _pForm, ", cTypeStem, funcType:get_rawTxt(), funcType:get_typeId(), cTypeEnvP, cTypeAnyP) )
   for index, argType in pairs( funcType:get_argTypeInfoList() ) do
      self:write( string.format( ", %s arg%d", getCType( argType ), index) )
   end
   
   self:writeln( ")" )
   self:writeln( "{" )
   self:writeln( 'return %s( _pEnv, _pForm' )
   for index, argType in pairs( funcType:get_argTypeInfoList() ) do
   end
   
   self:writeln( "}" )
end


function convFilter:processAndOr( node, opTxt, parent )

   local function isAndOr( exp )
   
      do
         local parentNode = _lune.__Cast( exp, 3, Nodes.ExpOp2Node )
         if parentNode ~= nil then
            do
               local _switchExp = parentNode:get_op().txt
               if _switchExp == "and" or _switchExp == "or" then
                  return true
               end
            end
            
         end
      end
      
      return false
   end
   
   local firstFlag = not isAndOr( parent )
   if firstFlag then
      self:writeln( "lune_popVal( _pEnv, lune_incStack( _pEnv ) ||" )
      self:pushIndent(  )
   end
   
   
   local opCC
   
   if opTxt == "and" then
      opCC = "&&"
   else
    
      opCC = "||"
   end
   
   
   if isAndOr( node:get_exp1() ) then
      filter( node:get_exp1(), self, node )
   else
    
      self:write( "lune_setStackVal( _pEnv, " )
      self:processVal2stem( node:get_exp1(), node )
      self:write( ") " )
   end
   
   if isAndOr( node:get_exp2() ) then
      filter( node:get_exp2(), self, node )
   else
    
      self:writeln( opCC )
      self:write( "lune_setStackVal( _pEnv, " )
      self:processVal2stem( node:get_exp2(), node )
      self:write( ") " )
   end
   
   
   if firstFlag then
      self:write( ")" )
      
      if not isStemType( node:get_expType() ) then
         self:write( getAccessPrimValFromStem( false, node:get_expType(), 0 ) )
      end
      
      
      self:popIndent(  )
   end
   
end


function convFilter:processExpOp2( node, opt )

   local opTxt = node:get_op().txt
   
   do
      local _switchExp = opTxt
      if _switchExp == "=" then
         local workParent
         
         local expList
         
         local mRetExp
         
         do
            local expListNode = _lune.__Cast( node:get_exp2(), 3, Nodes.ExpListNode )
            if expListNode ~= nil then
               expList = expListNode:get_expList()
               mRetExp = expListNode:get_mRetExp()
               workParent = expListNode
            else
               expList = {node:get_exp2()}
               mRetExp = nil
               workParent = node
            end
         end
         
         
         self:processSetValToNode( node, node:get_exp1(), expList, mRetExp )
      elseif _switchExp == "and" or _switchExp == "or" then
         self:processAndOr( node, opTxt, opt.node )
      else 
         
            do
               local _exp = Ast.bitBinOpMap[opTxt]
               if _exp ~= nil then
                  
                  do
                     local _switchExp = _exp
                     if _switchExp == Ast.BitOpKind.LShift then
                        opTxt = "<<"
                     elseif _switchExp == Ast.BitOpKind.RShift then
                        opTxt = ">>"
                     end
                  end
                  
                  filter( node:get_exp1(), self, node )
                  self:write( " " .. opTxt .. " " )
                  filter( node:get_exp2(), self, node )
               else
                  if _lune._Set_has(Ast.compOpSet, opTxt ) then
                     local valKind1 = getValKind( node:get_exp1():get_expType() )
                     local valKind2 = getValKind( node:get_exp2():get_expType() )
                     
                     if valKind1 == ValKind.Stem or valKind2 == ValKind.Stem then
                        if opTxt ~= "==" then
                           self:write( "!" )
                        end
                        
                        self:write( "lune_equals( " )
                        self:processVal2stem( node:get_exp1(), node )
                        self:write( "," )
                        self:processVal2stem( node:get_exp2(), node )
                        self:write( ")" )
                     elseif valKind1 == ValKind.Any and valKind2 == ValKind.Any then
                        self:write( "lune_equals_any( " )
                        self:processVal2any( node:get_exp1(), node )
                        self:write( "," )
                        self:processVal2any( node:get_exp2(), node )
                        self:write( ")" )
                     else
                      
                        self:accessPrimVal( node:get_exp1(), node )
                        self:write( " " .. opTxt .. " " )
                        self:accessPrimVal( node:get_exp2(), node )
                     end
                     
                  elseif _lune._Set_has(Ast.mathCompOpSet, opTxt ) then
                     self:accessPrimVal( node:get_exp1(), node )
                     self:write( " " .. opTxt .. " " )
                     self:accessPrimVal( node:get_exp2(), node )
                  else
                   
                     filter( node:get_exp1(), self, node )
                     self:write( " " .. opTxt .. " " )
                     filter( node:get_exp2(), self, node )
                  end
                  
               end
            end
            
      end
   end
   
end



function convFilter:processExpRef( node, opt )

   if self.processMode == ProcessMode.Immediate then
      self.accessSymbolSet:add( node:get_symbolInfo() )
   end
   
   
   if node:get_token().txt == "super" then
      local funcType = node:get_expType()
      self:write( string.format( "%s.%s", self:getFullName( funcType:get_parentInfo() ), funcType:get_rawTxt()) )
   elseif node:get_token().txt == "__mod__" then
      self:write( "*lune_module_path" )
   elseif node:get_token().txt == "..." then
      self:write( "_pDDD" )
   elseif node:get_expType():equals( TransUnit.getBuiltinFunc(  ).lune__load ) then
      self:write( "lune__load" )
   else
    
      local symbolInfo = node:get_symbolInfo()
      local valKind = self.scopeMgr:getSymbolValKind( symbolInfo )
      if valKind == ValKind.Var then
         self:write( string.format( "%s->stem", getSymbolName( symbolInfo )) )
         self:write( getAccessValFromStem( symbolInfo:get_typeInfo() ) )
      else
       
         
         if symbolInfo:get_kind() == Ast.SymbolKind.Fun or symbolInfo:get_typeInfo():get_kind() == Ast.TypeInfoKind.Func then
            self:write( self.moduleCtrl:getFuncName( symbolInfo:get_typeInfo() ) )
         else
          
            
            if valKind == ValKind.Any and symbolInfo:get_kind() == Ast.SymbolKind.Var then
               self:write( string.format( "(*%s)", getSymbolName( symbolInfo )) )
            else
             
               self:write( getSymbolName( symbolInfo ) )
            end
            
         end
         
      end
      
   end
   
end



function convFilter:processExpRefItem( node, opt )

   local val = node:get_val()
   local valType = val:get_expType()
   
   if valType:equals( Ast.builtinTypeString ) then
      self:accessPrimVal( val, node )
      self:write( "->val.str.pStr[" )
      do
         local indexNode = node:get_index()
         if indexNode ~= nil then
            filter( indexNode, self, node )
         else
            error( "index is nil" )
         end
      end
      
      self:write( "- 1 ]" )
   else
    
      processToIF( self.stream, self.moduleCtrl, node:get_expType(), function (  )
      
         do
            local _switchExp = valType:get_kind()
            if _switchExp == Ast.TypeInfoKind.List then
               self:write( "lune_mtd_List_refAt( _pEnv, " )
               filter( val, self, node )
               self:write( ", " )
               self:accessPrimVal( _lune.unwrap( node:get_index()), node )
               self:write( ")" )
               self:write( getAccessValFromStem( valType:get_itemTypeInfoList()[1] ) )
            end
         end
         
      end )
   end
   
end



function convFilter:processRefField( node, opt )

   
   do
      local symbolInfo = node:get_symbolInfo()
      if symbolInfo ~= nil then
         if symbolInfo:get_typeInfo():get_kind() == Ast.TypeInfoKind.Enum then
            if symbolInfo:get_kind() == Ast.SymbolKind.Mbr then
               self:write( self.moduleCtrl:getEnumTypeName( symbolInfo:get_typeInfo() ) )
               self:write( string.format( "__%s", getSymbolName( symbolInfo )) )
               return 
            end
            
            Util.err( "illegal access" )
         end
         
         
         do
            local _switchExp = symbolInfo:get_kind()
            if _switchExp == Ast.SymbolKind.Mbr then
               if node:get_prefix():get_expType():get_kind() == Ast.TypeInfoKind.Class then
                  if symbolInfo:get_staticFlag() then
                     self:write( self.moduleCtrl:getClassMemberName( symbolInfo ) )
                  else
                   
                     local className = self.moduleCtrl:getClassCName( node:get_prefix():get_expType() )
                     self:write( string.format( "lune_obj_%s( ", className) )
                     self:processVal2any( node:get_prefix(), node )
                     self:write( string.format( ")->%s", node:get_field().txt) )
                  end
                  
               end
               
            elseif _switchExp == Ast.SymbolKind.Mtd then
               if not symbolInfo:get_staticFlag() then
                  Util.err( "not support yet. instanse method." )
               end
               
               self:write( self.moduleCtrl:getMethodCName( symbolInfo:get_typeInfo() ) )
            end
         end
         
      end
   end
   
end



function convFilter:processExpOmitEnum( node, opt )

   self:write( self.moduleCtrl:getEnumValCName( node:get_expType(), node:get_valInfo():get_name() ) )
end



function convFilter:processGetField( node, opt )

   local prefixNode = node:get_prefix(  )
   local prefixType = prefixNode:get_expType()
   local fieldTxt = node:get_field(  ).txt
   
   if prefixType:get_kind() == Ast.TypeInfoKind.Enum then
      local enumFullName = self.moduleCtrl:getEnumTypeName( prefixType )
      do
         local _switchExp = fieldTxt
         if _switchExp == "_allList" then
            self:write( string.format( "%s_get__allList( _pEnv )", enumFullName) )
         elseif _switchExp == "_txt" then
            self:write( string.format( "%s_get__txt( _pEnv, ", enumFullName) )
            filter( prefixNode, self, node )
            self:write( ")" )
         end
      end
      
   elseif prefixType:get_kind() == Ast.TypeInfoKind.Class or prefixType:get_kind() == Ast.TypeInfoKind.IF then
      local getterType = _lune.unwrap( _lune.nilacc( prefixType:get_scope(), 'getTypeInfoField', 'callmtd' , string.format( "get_%s", fieldTxt), true, _lune.unwrap( prefixType:get_scope()), scopeAccess ))
      
      local function process(  )
      
         self:write( string.format( "%s( _pEnv, ", self.moduleCtrl:getCallMethodCName( getterType )) )
         self:processVal2any( prefixNode, node )
         self:write( ")" )
         
         processAlterAccessVal( self.stream, getterType:get_retTypeInfoList(), node:get_expTypeList() )
      end
      
      if #node:get_expTypeList() == 1 then
         processAlterToActualType( self.stream, self.moduleCtrl, getterType:get_retTypeInfoList()[1], node:get_expType(), process )
      else
       
         process(  )
      end
      
   end
   
end



function convFilter:processReturn( node, opt )

   local retTypeInfoList = self.currentRoutineInfo:get_funcInfo():get_retTypeInfoList()
   
   local blockStart
   
   do
      local expListNode = node:get_expList()
      if expListNode ~= nil then
         local expList = expListNode:get_expList()
         local retKind = getRetKind( retTypeInfoList )
         local needSetRet = true
         self:writeln( "{" )
         blockStart = true
         self:pushIndent(  )
         self:write( string.format( "%s _ret = ", getCRetType( retTypeInfoList )) )
         
         if #retTypeInfoList >= 2 then
            self:processCreateMRet( retTypeInfoList, expList, node )
         elseif #retTypeInfoList == 1 then
            do
               local _switchExp = retKind
               if _switchExp == ValKind.Stem then
                  self:processVal2stem( expList[1], node )
               elseif _switchExp == ValKind.Any then
                  self:processVal2any( expList[1], node )
               elseif _switchExp == ValKind.Prim then
                  filter( expList[1], self, node )
               else 
                  
                     Util.err( string.format( "no support -- %d", 7515) )
               end
            end
            
         else
          
         end
         
         self:writeln( ";" )
         if needSetRet then
            do
               local _switchExp = retKind
               if _switchExp == ValKind.Stem then
                  self:writeln( "lune_setRet( _pEnv, _ret );" )
               elseif _switchExp == ValKind.Any then
                  self:writeln( "lune_setRet( _pEnv, LUNE_STEM_ANY( _ret ) );" )
               elseif _switchExp == ValKind.Prim then
               else 
                  
                     Util.err( string.format( "no support -- %d", 7534) )
               end
            end
            
         end
         
      else
         blockStart = false
      end
   end
   
   
   if self.currentRoutineInfo:get_blockDepth() == 1 then
      self:writeln( "lune_leave_block( _pEnv );" )
   else
    
      self:writeln( string.format( "lune_leave_blockMulti( _pEnv, %d );", self.currentRoutineInfo:get_blockDepth()) )
   end
   
   
   if #retTypeInfoList ~= 0 then
      self:writeln( "return _ret;" )
   else
    
      self:writeln( "return;" )
   end
   
   
   if blockStart then
      self:popIndent(  )
      self:writeln( "}" )
   end
   
end



function convFilter:processTestBlock( node, opt )

   if self.enableTest then
      self:writeln( string.format( 'printf( "%s:\\n" );', node:get_name().txt) )
      filter( node:get_block(), self, node )
   end
   
end


function convFilter:processProvide( node, opt )

end


function convFilter:processAlias( node, opt )

end


function convFilter:processBoxing( node, opt )

end


function convFilter:processUnboxing( node, opt )

end


function convFilter:processLiteralVal( exp, parent )

   if self.processMode ~= ProcessMode.Immediate then
      local symbolList = exp:getSymbolInfo(  )
      if #symbolList > 0 then
         local work, valKind = self.scopeMgr:getCTypeForSym( symbolList[1] )
         if valKind ~= ValKind.Prim then
            self:processVal2stem( exp, parent )
            return 
         end
         
      end
      
   end
   
   local valType = exp:get_expType():get_srcTypeInfo()
   do
      local enumType = _lune.__Cast( valType, 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         valType = enumType:get_valTypeInfo()
      end
   end
   
   
   do
      local _switchExp = valType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         self:write( "lune_imdInt( " )
         filter( exp, self, parent )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeReal then
         self:write( "lune_imdReal( " )
         filter( exp, self, parent )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeBool then
         self:write( "lune_imdBool( " )
         filter( exp, self, parent )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeString then
         do
            local strNode = _lune.__Cast( exp, 3, Nodes.LiteralStringNode )
            if strNode ~= nil then
               if not strNode:get_expList() then
                  self:write( string.format( "lune_imdStr( %s )", strNode:get_token().txt) )
                  return 
               end
               
            end
         end
         
         self:write( "lune_imdAny( " )
         filter( exp, self, parent )
         self:write( ")" )
      else 
         
            do
               local _switchExp = valType:get_kind()
               if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Map or _switchExp == Ast.TypeInfoKind.Array or _switchExp == Ast.TypeInfoKind.Class then
                  self:write( "lune_imdAny( " )
                  filter( exp, self, parent )
                  self:write( ")" )
               else 
                  
                     Util.err( string.format( "illegal type -- %s", valType:getTxt(  )) )
               end
            end
            
      end
   end
   
end


local function getLiteralListFuncName( node )

   return string.format( "lune_list_%X", node:get_id())
end

function convFilter:processLiteralNode( exp, parent )

   do
      local _switchExp = exp:get_kind()
      if _switchExp == Nodes.NodeKind.get_LiteralList() or _switchExp == Nodes.NodeKind.get_LiteralMap() or _switchExp == Nodes.NodeKind.get_LiteralArray() or _switchExp == Nodes.NodeKind.get_LiteralSet() then
         self.processingNode = exp
         filter( exp, self, parent )
      else 
         
            self:pushStream(  )
            filter( exp, self, parent )
            self:popStream(  )
      end
   end
   
end


function convFilter:processLiteralListSub( collectionType, node, expListNodeOrg, literalFuncName )

   if _lune._Set_has(self.processedNodeSet, node ) then
      do
         local set = self.literalNode2AccessSymbolSet[node]
         if set ~= nil then
            for __index, symbol in pairs( set:get_list() ) do
               self.accessSymbolSet:add( symbol )
            end
            
         end
      end
      
      return 
   end
   
   self.processedNodeSet[node]= true
   local expListNode = expListNodeOrg
   if  nil == expListNode then
      local _expListNode = expListNode
   
      return 
   end
   
   if #expListNode:get_expList() == 0 then
      return 
   end
   
   
   for __index, exp in pairs( expListNode:get_expList() ) do
      self:processLiteralNode( exp, node )
   end
   
   self.processingNode = node
   
   self:write( string.format( "static %s %s( %s _pEnv", cTypeAnyP, literalFuncName, cTypeEnvP) )
   for __index, symbol in pairs( self.accessSymbolSet:get_list() ) do
      local valKind = self.scopeMgr:getSymbolValKind( symbol )
      local ctype
      
      ctype = self.scopeMgr:getCTypeForSym( symbol )
      self:write( string.format( ", %s %s", ctype, getSymbolName( symbol )) )
   end
   
   self:writeln( ")" )
   self:writeln( "{" )
   
   self:pushIndent(  )
   self:write( string.format( "lune_imd%s( list", collectionType) )
   self:pushIndent(  )
   for __index, exp in pairs( expListNode:get_expList() ) do
      self:write( ", " )
      self:processLiteralVal( exp, node )
   end
   
   self:popIndent(  )
   self:writeln( ");" )
   self:writeln( string.format( "return lune_create%s( _pEnv, list );", collectionType) )
   self:popIndent(  )
   self:writeln( "}" )
   
   self.literalNode2AccessSymbolSet[node] = self.accessSymbolSet:clone(  )
end


function convFilter:processLiteralList( node, opt )

   if self.processMode == ProcessMode.Immediate and self.processingNode == node then
      self:processLiteralListSub( "List", node, node:get_expList(), getLiteralListFuncName( node ) )
   else
    
      if node:get_expList() then
         self:write( string.format( "%s( _pEnv", getLiteralListFuncName( node )) )
         local symbolSet = self.literalNode2AccessSymbolSet[node]
         if  nil == symbolSet then
            local _symbolSet = symbolSet
         
            return 
         end
         
         for __index, symbol in pairs( symbolSet:get_list() ) do
            self:write( string.format( ", %s", getSymbolName( symbol )) )
         end
         
         self:write( ")" )
      else
       
         self:write( "lune_class_List_new( _pEnv )" )
      end
      
   end
   
end



local function getLiteralSetFuncName( node )

   return string.format( "lune_set_%X", node:get_id())
end

function convFilter:processLiteralSet( node, opt )

   if self.processMode == ProcessMode.Immediate and self.processingNode == node then
      self:processLiteralListSub( "Set", node, node:get_expList(), getLiteralSetFuncName( node ) )
   else
    
      if node:get_expList() then
         self:write( string.format( "%s( _pEnv", getLiteralSetFuncName( node )) )
         local symbolSet = self.literalNode2AccessSymbolSet[node]
         if  nil == symbolSet then
            local _symbolSet = symbolSet
         
            return 
         end
         
         for __index, symbol in pairs( symbolSet:get_list() ) do
            self:write( string.format( ", %s", getSymbolName( symbol )) )
         end
         
         self:write( ")" )
      else
       
         self:write( "lune_class_Set_new( _pEnv )" )
      end
      
   end
   
end



local function getLiteralMapFuncName( node )

   return string.format( "lune_map_%X", node:get_id())
end

function convFilter:processLiteralMapSub( node )

   if _lune._Set_has(self.processedNodeSet, node ) then
      do
         local set = self.literalNode2AccessSymbolSet[node]
         if set ~= nil then
            for __index, symbol in pairs( set:get_list() ) do
               self.accessSymbolSet:add( symbol )
            end
            
         end
      end
      
      return 
   end
   
   self.processedNodeSet[node]= true
   local pairList = node:get_pairList()
   if #pairList == 0 then
      return 
   end
   
   
   for __index, pair in pairs( pairList ) do
      self:processLiteralNode( pair:get_key(), node )
      self:processLiteralNode( pair:get_val(), node )
   end
   
   self.processingNode = node
   
   self:write( string.format( "static %s %s( %s _pEnv", cTypeAnyP, getLiteralMapFuncName( node ), cTypeEnvP) )
   for __index, symbol in pairs( self.accessSymbolSet:get_list() ) do
      self:write( string.format( ", %s %s", self.scopeMgr:getCTypeForSym( symbol ), getSymbolName( symbol )) )
   end
   
   self:writeln( ")" )
   self:writeln( "{" )
   
   self:pushIndent(  )
   self:write( "lune_imdMap( list" )
   self:pushIndent(  )
   for __index, pair in pairs( pairList ) do
      self:writeln( ", " )
      self:write( "{ " )
      self:processLiteralVal( pair:get_key(), node )
      self:write( ", " )
      self:processLiteralVal( pair:get_val(), node )
      self:write( "} " )
   end
   
   self:popIndent(  )
   self:writeln( ");" )
   self:writeln( "return lune_createMap( _pEnv, list );" )
   self:popIndent(  )
   self:writeln( "}" )
   
   self.literalNode2AccessSymbolSet[node] = self.accessSymbolSet:clone(  )
end


function convFilter:processLiteralMap( node, opt )

   if self.processMode == ProcessMode.Immediate and self.processingNode == node then
      self:processLiteralMapSub( node )
   else
    
      if #node:get_pairList() > 0 then
         self:write( string.format( "%s( _pEnv", getLiteralMapFuncName( node )) )
         local symbolSet = self.literalNode2AccessSymbolSet[node]
         if  nil == symbolSet then
            local _symbolSet = symbolSet
         
            return 
         end
         
         for __index, symbol in pairs( symbolSet:get_list() ) do
            self:write( string.format( ", %s", getSymbolName( symbol )) )
         end
         
         self:write( ")" )
      else
       
         self:write( "lune_class_Map_new( _pEnv )" )
      end
      
   end
   
end



function convFilter:processLiteralArray( node, opt )

end



function convFilter:processLiteralChar( node, opt )

   self:write( string.format( "%d", node:get_num() ) )
end



function convFilter:processLiteralInt( node, opt )

   self:write( node:get_token().txt )
end



function convFilter:processLiteralReal( node, opt )

   self:write( node:get_token().txt )
end



function convFilter:processLiteralString( node, opt )

   local txt = node:get_token().txt
   if string.find( txt, '^```' ) then
      txt = (string.format( "%q", txt:sub( 4, -4 )) ):gsub( "\\\n", "\\n" )
   elseif string.find( txt, "^'" ) then
      txt = string.format( '%q', txt:sub( 2, -2 ))
   end
   
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         do
            local expListNode = node:get_expList()
            if expListNode ~= nil then
               do
                  local mRetExp = expListNode:get_mRetExp()
                  if mRetExp ~= nil then
                     self:processCallWithMRet( node, getMRetFuncName( node ), cTypeAnyP, #expListNode:get_expList(), _lune.newAlge( MRetInfo.Format, {txt,expListNode}), expListNode )
                  else
                     self:write( string.format( "static %s lune_litstr_%d( %s _pEnv", cTypeAnyP, node:get_id(), cTypeEnvP) )
                     for index, exp in pairs( expListNode:get_expList() ) do
                        self:write( string.format( ", %s arg%d", cTypeStem, index) )
                     end
                     
                     self:writeln( ");" )
                  end
               end
               
            end
         end
         
         return 
      elseif _switchExp == ProcessMode.StringFormat then
         do
            local expListNode = node:get_expList()
            if expListNode ~= nil then
               do
                  local mRetExp = expListNode:get_mRetExp()
                  if mRetExp ~= nil then
                     self:processCallWithMRet( node, getMRetFuncName( node ), cTypeAnyP, #expListNode:get_expList(), _lune.newAlge( MRetInfo.Format, {txt,expListNode}), expListNode )
                  else
                     self:write( string.format( "static %s lune_litstr_%d( %s _pEnv", cTypeAnyP, node:get_id(), cTypeEnvP) )
                     for index, exp in pairs( expListNode:get_expList() ) do
                        self:write( string.format( ", %s arg%d", cTypeStem, index) )
                     end
                     
                     self:writeln( ")" )
                     self:writeln( "{" )
                     self:pushIndent(  )
                     
                     self:write( "return u_mtd_string_format( _pEnv, " )
                     self:write( "lune_litStr2any( _pEnv, " )
                     self:write( txt )
                     self:write( "), " )
                     
                     local expList = expListNode:get_expList()
                     self:write( "lune_createDDD" )
                     local lastExp = expList[#expList]
                     self:write( string.format( "( _pEnv, %s, %d", Nodes.hasMultiValNode( lastExp ), #expList) )
                     
                     for index = 1, #expList do
                        self:write( string.format( ", arg%d", index) )
                     end
                     
                     self:writeln( ") );" )
                     
                     self:popIndent(  )
                     self:writeln( "}" )
                  end
               end
               
            end
         end
         
         return 
      end
   end
   
   
   do
      local expListNode = node:get_expList()
      if expListNode ~= nil then
         do
            local mRetExp = expListNode:get_mRetExp()
            if mRetExp ~= nil then
               self:write( string.format( "%s( _pEnv", getMRetFuncName( node )) )
               for index, exp in pairs( expListNode:get_expList() ) do
                  if index > mRetExp:get_index() then
                     break
                  end
                  
                  self:write( ", " )
                  filter( exp, self, node )
               end
               
               self:write( ")" )
            else
               self:write( string.format( "lune_litstr_%d( _pEnv ", node:get_id()) )
               for __index, exp in pairs( expListNode:get_expList() ) do
                  self:write( ", " )
                  self:processVal2stem( exp, node )
               end
               
               self:write( ")" )
            end
         end
         
      else
         local opList = TransUnit.findForm( txt )
         
         self:write( "lune_litStr2any( _pEnv, " )
         self:write( txt )
         
         self:write( ")" )
      end
   end
   
end



function convFilter:processLiteralBool( node, opt )

   if node:get_token().txt == "true" then
      self:write( "true" )
   else
    
      self:write( "false" )
   end
   
end



function convFilter:processLiteralNil( node, opt )

   self:write( cValNil )
end



function convFilter:processBreak( node, opt )

   self:write( "break" )
end



function convFilter:processLiteralSymbol( node, opt )

end



function convFilter:processAbbr( node, opt )

   self:write( cValNone )
end



local function createFilter( enableTest, outputBuiltin, streamName, stream, headerStream, ast )

   return convFilter.new(enableTest, outputBuiltin, streamName, stream, headerStream, ast)
end
_moduleObj.createFilter = createFilter

local function outputBootcode( stream, launchModuleName )

   local srcStream = Util.SimpleSourceOStream.new(stream, nil, stepIndent)
   
   local launchModulePath = launchModuleName:gsub( "%.", "/" )
   local initFuncName = string.format( "lune_init_%s", (launchModuleName:gsub( "%.", "_" ) ))
   srcStream:writeln( string.format( [==[
#include <lunescript.h>
#include <%s.h>
    
void lune_run_module( lune_env_t * _pEnv ) {
   %s( _pEnv );    
}
]==], launchModulePath, initFuncName) )
end
_moduleObj.outputBootcode = outputBootcode

return _moduleObj
