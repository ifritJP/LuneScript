--lune/base/convCC.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@convCC'
local _lune = {}
if _lune2 then
   _lune = _lune2
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

if not _lune2 then
   _lune2 = _lune
end

local Ver = _lune.loadModule( 'lune.base.Ver' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Util = _lune.loadModule( 'lune.base.Util' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local LuaMod = _lune.loadModule( 'lune.base.LuaMod' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )

local cTypeInt = "lns_int_t"
local cTypeReal = "lns_real_t"
local cTypeBool = "lns_bool_t"
local cTypeStem = "lns_stem_t"
local cTypeAny = "lns_any_t"
local cTypeAnyP = "lns_any_t *"
local cTypeAnyPP = "lns_any_t **"
local cTypeEnvP = "lns_env_t *"
local cTypeVarP = "lns_closureVar_t *"
local cTypeMod = "lns_module_t"
local cTypeModP = "lns_module_t *"
local cTypeBlockP = "lns_block_t *"
local cValNil = "lns_global.nilStem"
local cValNone = "lns_global.noneStem"
local cValDDD0 = "lns_global.ddd0"

local accessAny = ".val.pAny"

local stepIndent = 4

local scopeAccess = Ast.ScopeAccess.Full

local invalidSymbolId = -1

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


local DepthInfo = {}
function DepthInfo.new(  )
   local obj = {}
   DepthInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function DepthInfo:__init() 
   self.blockDepth = 1
end
function DepthInfo:pushDepth(  )

   self.blockDepth = self.blockDepth + 1
end
function DepthInfo:popDepth(  )

   self.blockDepth = self.blockDepth - 1
end
function DepthInfo.setmeta( obj )
  setmetatable( obj, { __index = DepthInfo  } )
end
function DepthInfo:get_blockDepth()
   return self.blockDepth
end


local DepthStack = {}
function DepthStack.new(  )
   local obj = {}
   DepthStack.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function DepthStack:__init() 
   self.stack = {}
end
function DepthStack:newInfo( info )

   table.insert( self.stack, info )
end
function DepthStack:delInfo(  )

   table.remove( self.stack )
end
function DepthStack:current(  )

   if #self.stack == 0 then
      Util.err( "stack empty" )
   end
   
   return self.stack[#self.stack]
end
function DepthStack:currentR(  )

   if #self.stack == 0 then
      Util.err( "stack empty" )
   end
   
   return self.stack[#self.stack]
end
function DepthStack:pushDepth(  )

   self:current(  ):pushDepth(  )
end
function DepthStack:popDepth(  )

   self:current(  ):popDepth(  )
end
function DepthStack:get_blockDepth(  )

   return self:currentR(  ):get_blockDepth()
end
function DepthStack.setmeta( obj )
  setmetatable( obj, { __index = DepthStack  } )
end


local RoutineInfo = {}
setmetatable( RoutineInfo, { __index = DepthInfo } )
function RoutineInfo.new( funcInfo )
   local obj = {}
   RoutineInfo.setmeta( obj )
   if obj.__init then obj:__init( funcInfo ); end
   return obj
end
function RoutineInfo:__init(funcInfo) 
   DepthInfo.__init( self)
   
   self.funcInfo = funcInfo
end
function RoutineInfo.setmeta( obj )
  setmetatable( obj, { __index = RoutineInfo  } )
end
function RoutineInfo:get_funcInfo()
   return self.funcInfo
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
ProcessMode.InitFuncSym = 9
ProcessMode._val2NameMap[9] = 'InitFuncSym'
ProcessMode.__allList[10] = ProcessMode.InitFuncSym
ProcessMode.DefWrap = 10
ProcessMode._val2NameMap[10] = 'DefWrap'
ProcessMode.__allList[11] = ProcessMode.DefWrap


local function isClassMember( symbol )

   if _lune.nilacc( symbol:get_scope():get_ownerTypeInfo(), 'get_kind', 'callmtd' ) == Ast.TypeInfoKind.Class and symbol:get_kind() == Ast.SymbolKind.Mbr and symbol:get_staticFlag() then
      return true
   end
   
   return false
end

local function str2cstr( txt )

   local work = txt
   if string.find( work, '^```' ) then
      work = (string.format( "%q", work:sub( 4, -4 )) ):gsub( "\\\n", "\\n" )
   elseif string.find( work, "^'" ) then
      work = string.format( '"%s"', ((string.format( "%s", work:sub( 2, -2 )) ):gsub( '"', '\\"' ) ))
   end
   
   work = work:gsub( "\\9", "\\t" )
   return work
end

local ModuleCtrl = {}
function ModuleCtrl.setupBuiltinSym(  )

   local builtinFunc = TransUnit.getBuiltinFunc(  )
   
   local symMap = {}
   local typeMap = {}
   
   
   
   
   typeMap[builtinFunc.lns_type] = "lns_f_" .. 'type'
   
   
   typeMap[builtinFunc.lns_error] = "lns_f_" .. 'error'
   
   
   typeMap[builtinFunc.lns_print] = "lns_f_" .. 'print'
   
   
   typeMap[builtinFunc.lns_tonumber] = "lns_f_" .. 'tonumber'
   
   
   typeMap[builtinFunc.lns_tostring] = "lns_f_" .. 'tostring'
   
   
   typeMap[builtinFunc.lns_load] = "lns_f_" .. 'load'
   
   
   typeMap[builtinFunc.lns_loadfile] = "lns_f_" .. 'loadfile'
   
   
   typeMap[builtinFunc.lns_require] = "lns_f_" .. 'require'
   
   
   typeMap[builtinFunc.lns_collectgarbage] = "lns_f_" .. 'collectgarbage'
   
   
   typeMap[builtinFunc.lns__fcall] = "lns_f_" .. '_fcall'
   
   
   typeMap[builtinFunc.lns__load] = "lns_f_" .. '_load'
   
   
   typeMap[builtinFunc.lns__kind] = "lns_f_" .. '_kind'
   
   
   typeMap[builtinFunc.io_open] = "mtd_lns_" .. 'io_open'
   
   
   typeMap[builtinFunc.io_popen] = "mtd_lns_" .. 'io_popen'
   
   
   typeMap[builtinFunc.package_searchpath] = "mtd_lns_" .. 'package_searchpath'
   
   
   typeMap[builtinFunc.os_clock] = "mtd_lns_" .. 'os_clock'
   
   
   typeMap[builtinFunc.os_date] = "mtd_lns_" .. 'os_date'
   
   
   typeMap[builtinFunc.os_difftime] = "mtd_lns_" .. 'os_difftime'
   
   
   typeMap[builtinFunc.os_exit] = "mtd_lns_" .. 'os_exit'
   
   
   typeMap[builtinFunc.os_remove] = "mtd_lns_" .. 'os_remove'
   
   
   typeMap[builtinFunc.os_rename] = "mtd_lns_" .. 'os_rename'
   
   
   typeMap[builtinFunc.os_time] = "mtd_lns_" .. 'os_time'
   
   
   typeMap[builtinFunc.string_byte] = "mtd_lns_" .. 'string_byte'
   
   
   typeMap[builtinFunc.string_dump] = "mtd_lns_" .. 'string_dump'
   
   
   typeMap[builtinFunc.string_find] = "mtd_lns_" .. 'string_find'
   
   
   typeMap[builtinFunc.string_format] = "mtd_lns_" .. 'string_format'
   
   
   typeMap[builtinFunc.string_gmatch] = "mtd_lns_" .. 'string_gmatch'
   
   
   typeMap[builtinFunc.string_gsub] = "mtd_lns_" .. 'string_gsub'
   
   
   typeMap[builtinFunc.string_lower] = "mtd_lns_" .. 'string_lower'
   
   
   typeMap[builtinFunc.string_rep] = "mtd_lns_" .. 'string_rep'
   
   
   typeMap[builtinFunc.string_reverse] = "mtd_lns_" .. 'string_reverse'
   
   
   typeMap[builtinFunc.string_sub] = "mtd_lns_" .. 'string_sub'
   
   
   typeMap[builtinFunc.string_upper] = "mtd_lns_" .. 'string_upper'
   
   
   typeMap[builtinFunc.math_random] = "mtd_lns_" .. 'math_random'
   
   
   typeMap[builtinFunc.math_randomseed] = "mtd_lns_" .. 'math_randomseed'
   
   
   typeMap[builtinFunc.debug_getinfo] = "mtd_lns_" .. 'debug_getinfo'
   
   
   typeMap[builtinFunc.debug_getlocal] = "mtd_lns_" .. 'debug_getlocal'
   
   return symMap, typeMap
end
function ModuleCtrl.new( typeNameCtrl, moduleInfoManager )
   local obj = {}
   ModuleCtrl.setmeta( obj )
   if obj.__init then obj:__init( typeNameCtrl, moduleInfoManager ); end
   return obj
end
function ModuleCtrl:__init(typeNameCtrl, moduleInfoManager) 
   self.builtinSym2CFuncMap, self.builtinType2CFuncMap = ModuleCtrl.setupBuiltinSym(  )
   self.typeNameCtrl = typeNameCtrl
   self.moduleInfoManager = moduleInfoManager
end
function ModuleCtrl:getBuiltinFuncName( symbol )

   return self.builtinSym2CFuncMap[symbol]
end
function ModuleCtrl:getBuiltinFuncNameFromType( typeInfo )

   return self.builtinType2CFuncMap[typeInfo]
end
function ModuleCtrl:getFilePath( moduleTypeInfo )

   local workName = moduleTypeInfo:getFullName( self.typeNameCtrl, self.moduleInfoManager, false )
   local fullName = string.format( "%s", (workName:gsub( "[&@]", "" ):gsub( "%.", "/" ) ))
   return fullName
end
function ModuleCtrl:getCanonicalName( typeInfo )

   return typeInfo:getFullName( self.typeNameCtrl, self.moduleInfoManager, false )
end
function ModuleCtrl:getFullName( typeInfo )

   do
      local alterType = _lune.__Cast( typeInfo:get_srcTypeInfo(), 3, Ast.AlternateTypeInfo )
      if alterType ~= nil then
         if alterType:hasBase(  ) then
            typeInfo = alterType:get_baseTypeInfo()
         end
         
      end
   end
   
   typeInfo = typeInfo:get_srcTypeInfo():get_genSrcTypeInfo()
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

   return string.format( "lns_algeType_%s", self:getAlgeCName( algeType ))
end
function ModuleCtrl:getAlgeValCName( algeType, valName )

   return string.format( "lns__alge_%s_%s", self:getFullName( algeType ), valName)
end
function ModuleCtrl:getAlgeValStrCName( algeType, valName )

   return string.format( "lns__alge_%s_%s_t", self:getFullName( algeType ), valName)
end
function ModuleCtrl:getNewAlgeCName( algeType, valName )

   return string.format( "lns__new_alge_%s_%s", self:getFullName( algeType ), valName)
end
function ModuleCtrl:getAlgeInitCName( algeType )

   return string.format( "lns__init_alge_%s", self:getAlgeCName( algeType ))
end
function ModuleCtrl:getEnumTypeName( typeInfo )

   local srcType = typeInfo:get_srcTypeInfo()
   local fullName = self:getFullName( srcType )
   
   if Ast.isPubToExternal( typeInfo:get_accessMode() ) then
      return fullName
   end
   
   return string.format( "e_%s", fullName)
end
function ModuleCtrl:getEnumValCName( typeInfo, valName )

   return string.format( "%s__%s", self:getEnumTypeName( typeInfo ), valName)
end
function ModuleCtrl:getEnumVal2NameMapName( enumType )

   return string.format( "%s_val2NameMap", self:getEnumTypeName( enumType ))
end
function ModuleCtrl:getClassCName( classType )

   return "lns_" .. self:getFullName( classType )
end
function ModuleCtrl:getNewName( classType )

   return string.format( "lns_class_%s_new", self:getClassCName( classType ))
end
function ModuleCtrl:getCtorName( classType )

   return string.format( "mtd_%s___init", self:getClassCName( classType ))
end
function ModuleCtrl:getClassMetaName( classType )

   if classType:get_srcTypeInfo() == Ast.headTypeInfo then
      return "lns_type_meta_lns__root"
   end
   
   return string.format( "lns_type_meta_%s", self:getClassCName( classType ))
end
function ModuleCtrl:getMethodCName( methodTypeInfo )

   return string.format( "mtd_%s_%s", self:getClassCName( methodTypeInfo:get_parentInfo() ), methodTypeInfo:get_rawTxt())
end
function ModuleCtrl:getFuncName( typeInfo )

   if typeInfo:get_rawTxt() == "" then
      return string.format( "lns_anonymous_%d", typeInfo:get_typeId())
   end
   
   do
      local _switchExp = typeInfo:get_accessMode()
      if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
         do
            local cFuncName = self:getBuiltinFuncNameFromType( typeInfo )
            if cFuncName ~= nil then
               return cFuncName
            end
         end
         
         do
            local _switchExp = typeInfo:get_parentInfo():get_kind()
            if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Enum then
               return self:getMethodCName( typeInfo )
            end
         end
         
         return self:getFullName( typeInfo )
      end
   end
   
   
   if typeInfo:get_parentInfo():get_kind() == Ast.TypeInfoKind.Class then
      return self:getMethodCName( typeInfo )
   end
   
   
   return string.format( "lns_f_%d_%s", typeInfo:get_typeId(), typeInfo:get_rawTxt())
end
function ModuleCtrl:getNilMethodCName( methodTypeInfo )

   return string.format( "l_nil_mtd_%s_%s", self:getClassCName( methodTypeInfo:get_parentInfo() ), methodTypeInfo:get_rawTxt())
end
function ModuleCtrl:getCallMethodCName( methodTypeInfo )

   do
      local _switchExp = methodTypeInfo:get_parentInfo():get_kind()
      if _switchExp == Ast.TypeInfoKind.List then
         return string.format( "lns_mtd_List_%s", methodTypeInfo:get_rawTxt())
      elseif _switchExp == Ast.TypeInfoKind.Array then
         return string.format( "lns_mtd_Array_%s", methodTypeInfo:get_rawTxt())
      elseif _switchExp == Ast.TypeInfoKind.Set then
         return string.format( "lns_mtd_Set_%s", methodTypeInfo:get_rawTxt())
      elseif _switchExp == Ast.TypeInfoKind.Map then
         return string.format( "lns_mtd_Map_%s", methodTypeInfo:get_rawTxt())
      end
   end
   
   return string.format( "l_call_mtd_%s_%s", self:getClassCName( methodTypeInfo:get_parentInfo() ), methodTypeInfo:get_rawTxt())
end
function ModuleCtrl:getClassMemberName( symbolInfo )

   local classTypeInfo = symbolInfo:get_scope():getClassTypeInfo(  )
   return string.format( "l_var_%s_%s", self:getClassCName( classTypeInfo ), symbolInfo:get_name())
end
function ModuleCtrl:getSymbolName( symbolInfo )

   if symbolInfo:get_typeInfo():get_kind() == Ast.TypeInfoKind.DDD then
      return "_pDDD"
   end
   
   if symbolInfo:get_kind() == Ast.SymbolKind.Mbr then
      if isClassMember( symbolInfo ) then
         return self:getClassMemberName( symbolInfo )
      end
      
      return symbolInfo:get_name()
   end
   
   if Ast.isPubToExternal( symbolInfo:get_accessMode() ) then
      if symbolInfo:get_accessMode() == Ast.AccessMode.Global then
         return "lns_" .. symbolInfo:get_name()
      end
      
      local moduleType = symbolInfo:get_scope():getNamespaceTypeInfo(  ):getModule(  )
      return string.format( "lns_%s_%s", self:getFullName( moduleType ), symbolInfo:get_name())
   end
   
   do
      local _switchExp = symbolInfo:get_kind()
      if _switchExp == Ast.SymbolKind.Var then
         if symbolInfo:get_symbolId() == invalidSymbolId then
            return symbolInfo:get_name()
         end
         
         return string.format( "lns_%s_%d", symbolInfo:get_name(), symbolInfo:get_symbolId())
      end
   end
   
   return symbolInfo:get_name()
end
function ModuleCtrl:getFormName( typeInfo )

   return string.format( "l_form_%s", self:getFullName( typeInfo ))
end
function ModuleCtrl:getCallFormName( typeInfo )

   return string.format( "lns_call_formFunc_%s", self:getFullName( typeInfo ))
end
function ModuleCtrl:getFuncCastWrapName( orgFunc, castType )

   return string.format( "wrap_%s_2_%s", self:getFuncName( orgFunc ), self:getFuncName( castType ))
end
function ModuleCtrl:getEnumFuncName( enumType, name )

   local scope = _lune.unwrap( enumType:get_scope())
   return self:getFuncName( (_lune.unwrap( scope:getSymbolInfoChild( name )) ):get_typeInfo() )
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
function WorkSymbol:get_mutable(  )

   return false
end
function WorkSymbol:get_symbolId(  )

   return invalidSymbolId
end
function WorkSymbol:get_hasAccessFromClosure(  )

   return false
end
function WorkSymbol.setmeta( obj )
  setmetatable( obj, { __index = WorkSymbol  } )
end
function WorkSymbol.new( scope, accessMode, name, typeInfo, kind, staticFlag, convModuleParam )
   local obj = {}
   WorkSymbol.setmeta( obj )
   if obj.__init then
      obj:__init( scope, accessMode, name, typeInfo, kind, staticFlag, convModuleParam )
   end
   return obj
end
function WorkSymbol:__init( scope, accessMode, name, typeInfo, kind, staticFlag, convModuleParam )

   self.scope = scope
   self.accessMode = accessMode
   self.name = name
   self.typeInfo = typeInfo
   self.kind = kind
   self.staticFlag = staticFlag
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
function WorkSymbol:get_staticFlag()
   return self.staticFlag
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

local function createSymbolParam( name, valKind, cTypeTxt )

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
         
            Util.err( string.format( "not support %s:%s", name, ValKind:_getTxt( valKind)
            ) )
      end
   end
   
end

local ScopeMgr = {}
function ScopeMgr.new( moduleCtrl )
   local obj = {}
   ScopeMgr.setmeta( obj )
   if obj.__init then obj:__init( moduleCtrl ); end
   return obj
end
function ScopeMgr:__init(moduleCtrl) 
   self.scope2InfoMap = {}
   self.moduleCtrl = moduleCtrl
   self.numOf__func__ = 0
   self.moduleBlockAnyNum = 0
end
function ScopeMgr.setSymbolParam( scopeInfo, symbol )

   local param
   
   do
      local _switchExp = getValKind( symbol:get_typeInfo() )
      if _switchExp == ValKind.Stem then
         param = SymbolParam.new(ValKind.Stem, scopeInfo.stemNum, cTypeStem)
         scopeInfo.stemNum = scopeInfo.stemNum + 1
      elseif _switchExp == ValKind.Any then
         if symbol:get_name() == "self" then
            param = SymbolParam.new(ValKind.Any, 0, cTypeAnyP)
         else
          
            if symbol:get_kind() == Ast.SymbolKind.Var or symbol:get_kind() == Ast.SymbolKind.Mbr then
               param = SymbolParam.new(ValKind.Any, scopeInfo.anyNum, cTypeAnyPP)
               scopeInfo.anyNum = scopeInfo.anyNum + 1
            else
             
               if symbol:get_mutable() then
                  param = SymbolParam.new(ValKind.Any, scopeInfo.anyNum, cTypeAnyPP)
                  scopeInfo.anyNum = scopeInfo.anyNum + 1
               else
                
                  param = SymbolParam.new(ValKind.Any, 0, cTypeAnyP)
               end
               
            end
            
         end
         
      elseif _switchExp == ValKind.Prim then
         param = SymbolParam.new(ValKind.Prim, 0, getCType( symbol:get_typeInfo() ))
      else 
         
            Util.err( string.format( "not support %s", symbol:get_typeInfo():getTxt(  )) )
      end
   end
   
   return param
end
function ScopeMgr:setupScopeParamSub( scope )

   do
      local scopeInfo = self.scope2InfoMap[scope]
      if scopeInfo ~= nil then
         return scopeInfo
      end
   end
   
   
   local scopeInfo
   
   
   scopeInfo = ScopeInfo.new(0, 0, 0)
   
   if _lune.nilacc( scope:get_ownerTypeInfo(), 'isModule', 'callmtd'  ) then
      
      scopeInfo.anyNum = 2
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
            if not symbol:get_convModuleParam() then
               local param
               
               if symbol:get_name() ~= "__func__" then
                  do
                     local _switchExp = symbol:get_kind()
                     if _switchExp == Ast.SymbolKind.Var or _switchExp == Ast.SymbolKind.Arg then
                        if symbol:get_hasAccessFromClosure() then
                           
                           param = SymbolParam.new(ValKind.Var, scopeInfo.varNum, cTypeVarP)
                           scopeInfo.varNum = scopeInfo.varNum + 1
                        else
                         
                           param = ScopeMgr.setSymbolParam( scopeInfo, symbol )
                        end
                        
                     elseif _switchExp == Ast.SymbolKind.Fun then
                        if symbol:get_hasAccessFromClosure() then
                           
                           param = SymbolParam.new(ValKind.Var, scopeInfo.varNum, cTypeVarP)
                           scopeInfo.varNum = scopeInfo.varNum + 1
                        else
                         
                           param = createSymbolParam( symbol:get_name(), getValKind( symbol:get_typeInfo() ), getCType( symbol:get_typeInfo() ) )
                        end
                        
                     elseif _switchExp == Ast.SymbolKind.Mtd then
                        local retTypeList = symbol:get_typeInfo():get_retTypeInfoList()
                        param = createSymbolParam( symbol:get_name(), getRetKind( retTypeList ), getCRetType( retTypeList ) )
                     elseif _switchExp == Ast.SymbolKind.Mbr then
                        if isClassMember( symbol ) then
                           param = (_lune.unwrap( symbol:get_convModuleParam()) )
                        else
                         
                           param = createSymbolParam( symbol:get_name(), getValKind( symbol:get_typeInfo() ), getCType( symbol:get_typeInfo() ) )
                        end
                        
                     else 
                        
                           param = SymbolParam.new(ValKind.Other, 0, cTypeStem)
                     end
                  end
                  
               else
                
                  param = SymbolParam.new(ValKind.Any, self.numOf__func__ + self.moduleBlockAnyNum, cTypeAnyP)
                  self.numOf__func__ = self.numOf__func__ + 1
               end
               
               symbol:set_convModuleParam( param )
            end
            
         end
      end
   end
   
   
   self.scope2InfoMap[scope] = scopeInfo
   
   return scopeInfo
end
function ScopeMgr:setup( scope, declMemberList )

   self:setupScopeParamSub( scope )
   
   for __index, declMember in pairs( declMemberList ) do
      local scopeInfo = _lune.unwrap( self.scope2InfoMap[scope])
      local symbol = declMember:get_symbolInfo()
      if isClassMember( symbol ) then
         symbol:set_convModuleParam( ScopeMgr.setSymbolParam( scopeInfo, symbol ) )
      end
      
   end
   
end
function ScopeMgr.create( moduleCtrl, initBlockScope )

   local scopeMgr = ScopeMgr.new(moduleCtrl)
   local param = scopeMgr:setupScopeParamSub( initBlockScope )
   scopeMgr.moduleBlockAnyNum = param.anyNum
   return scopeMgr
end
function ScopeMgr:setupScopeParam( scope )

   local scopeInfo = self:setupScopeParamSub( scope )
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
   
   Util.err( string.format( "illegal symbol -- %s %s %s %d", symbol:get_name(), Ast.SymbolKind:_getTxt( symbol:get_kind())
   , self.moduleCtrl:getCanonicalName( symbol:get_scope():getNamespaceTypeInfo(  ) ), 952) )
end
function ScopeMgr:getSymbolValKind( symbol )

   local symbolParam = self:getSymbolParam( symbol )
   return symbolParam.kind
end
function ScopeMgr:getCTypeForSym( symbol )

   local param = self:getSymbolParam( symbol )
   return param.typeTxt, param.kind
end
function ScopeMgr:getAccessPrimValFromSymbol( symbolInfo )

   return getAccessPrimValFromSymbolDirect( self.moduleCtrl:getSymbolName( symbolInfo ), self:getSymbolValKind( symbolInfo ), symbolInfo:get_typeInfo() )
end
function ScopeMgr:getAccessPrimValFromSymbolOnly( symbolInfo )

   return getAccessPrimValFromSymbolDirect( "", self:getSymbolValKind( symbolInfo ), symbolInfo:get_typeInfo() )
end
function ScopeMgr.setmeta( obj )
  setmetatable( obj, { __index = ScopeMgr  } )
end
function ScopeMgr:get_numOf__func__()
   return self.numOf__func__
end


local function getLiteralStrAny( txt )

   return string.format( 'lns_litStr2any( _pEnv, %s )', txt)
end

local function getLiteralStrStem( txt )

   return string.format( 'LNS_STEM_ANY( %s )', getLiteralStrAny( txt ))
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
function convFilter:createRefNodeFromSym( symbol )

   return Nodes.ExpRefNode.create( self.dummyNodeManager, _lune.unwrap( symbol:get_pos()), false, {symbol:get_typeInfo()}, symbol )
end
function convFilter.new( enableTest, outputBuiltin, streamName, stream, headerStream, ast )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( enableTest, outputBuiltin, streamName, stream, headerStream, ast ); end
   return obj
end
function convFilter:__init(enableTest, outputBuiltin, streamName, stream, headerStream, ast) 
   Nodes.Filter.__init( self,ast:get_moduleTypeInfo(), ast:get_moduleTypeInfo():get_scope())
   
   
   self.dummyNodeManager = Nodes.NodeManager.new()
   self.canConv = true
   self.enableTest = enableTest
   self.outputBuiltinFlag = outputBuiltin
   self.processingNode = nil
   self.processedNodeSet = {}
   self.accessSymbolSet = Util.OrderedSet.new()
   self.literalNode2AccessSymbolSet = {}
   
   self.duringDeclFunc = false
   
   self.processMode = ProcessMode.Prototype
   
   self.moduleTypeInfo = ast:get_moduleTypeInfo()
   
   self.routineInfoStack = DepthStack.new()
   self.routineInfoStack:newInfo( RoutineInfo.new(Ast.builtinTypeNone) )
   self.routineInfoStack:newInfo( RoutineInfo.new(ast:get_moduleTypeInfo()) )
   
   self.loopInfoStack = DepthStack.new()
   self.loopInfoStack:newInfo( DepthInfo.new() )
   self.loopInfoStack:newInfo( DepthInfo.new() )
   
   self.ast = _lune.unwrap( _lune.__Cast( ast:get_node(), 3, Nodes.RootNode ))
   
   self.streamName = streamName
   
   self.streamQueue = {}
   
   self.pubVarName2InfoMap = {}
   self.pubFuncName2InfoMap = {}
   
   self.moduleCtrl = ModuleCtrl.new(self:get_typeNameCtrl(), self:get_moduleInfoManager())
   self.scopeMgr = ScopeMgr.create( self.moduleCtrl, _lune.unwrap( ast:get_moduleTypeInfo():get_scope()) )
   
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
function convFilter.setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end
function convFilter:popIndent( ... )
   return self.stream:popIndent( ... )
end

function convFilter:pushIndent( ... )
   return self.stream:pushIndent( ... )
end

function convFilter:returnToSource( ... )
   return self.stream:returnToSource( ... )
end

function convFilter:switchToHeader( ... )
   return self.stream:switchToHeader( ... )
end

function convFilter:write( ... )
   return self.stream:write( ... )
end

function convFilter:writeln( ... )
   return self.stream:writeln( ... )
end



local function processAddModuleGlobal( stream, valName )

   stream:writeln( string.format( "lns_mtd_List_insert( _pEnv, *lns_module_globalStemList, %s );", valName) )
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
    
      self:writeln( string.format( "lns_init_%s( _pEnv );", self.moduleCtrl:getFullName( node:get_moduleTypeInfo() )) )
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
   
   local moduleFullName = self.moduleCtrl:getFullName( node:get_moduleTypeInfo() )
   local moduleInfoName = string.format( "lns_moduleInfo_%s", moduleFullName)
   
   if self.outputBuiltinFlag then
      self:writeln( "static void lns_init_lns_builtin_Sub( lns_env_t * _pEnv );" )
   elseif not self.canConv then
      self:writeln( string.format( "extern void lns_init_%s_Sub( lns_env_t * _pEnv );", moduleFullName) )
   end
   
   
   local function process( out2HMode )
   
      self:write( string.format( "%s%s lns_init_%s( %s _pEnv )", getOut2HeaderPrefix( out2HMode ), cTypeModP, moduleFullName, cTypeEnvP) )
      
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
   
   self:writeln( string.format( "if ( %s.readyFlag ) {", moduleInfoName) )
   self:pushIndent(  )
   self:writeln( string.format( "return &%s;", moduleInfoName) )
   self:popIndent(  )
   self:writeln( "}" )
   self:writeln( string.format( "%s.readyFlag = true;", moduleInfoName) )
   self:writeln( string.format( "lns_add2list( &_pEnv->loadModuleTop, &%s);", moduleInfoName) )
   self:writeln( "" )
   
   local moduleBlockName = getBlockName( self.ast:get_moduleScope() )
   self:writeln( string.format( "lns_block_t * %s = lns_enter_module( _pEnv, %d, %d, %d );", moduleBlockName, anyNum + self.scopeMgr:get_numOf__func__(), stemNum, varNum) )
   self:writeln( string.format( "%s.pBlock = %s;", moduleInfoName, moduleBlockName) )
   
   self:writeln( string.format( "lns_set_block_any( %s, 0, lns_module_globalStemList);", moduleBlockName) )
   self:writeln( "lns_setQ_any( lns_module_globalStemList, lns_class_List_new( _pEnv ));" )
   
   self:writeln( string.format( "lns_set_block_any( %s, 1, lns_module_path);", moduleBlockName) )
   self:writeln( string.format( 'lns_setQ_any( lns_module_path, lns_litStr2any( _pEnv, "%s"));', node:get_moduleTypeInfo():getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager() )) )
   
   for __index, blockNode in pairs( node:get_nodeManager():getBlockNodeList(  ) ) do
      self.scopeMgr:setupScopeParam( blockNode:get_scope() )
   end
   
   self:writeln( "lns_enter_block( _pEnv, 0, 0, 0 );" )
   if self.canConv then
      self:writeln( string.format( "initFuncSym( _pEnv, %s );", moduleBlockName) )
   end
   
   self:writeln( "" )
   
   if self.outputBuiltinFlag then
      self:writeln( "lns_init_lns_builtin_Sub( _pEnv );" )
   end
   
   if not self.canConv then
      self:writeln( string.format( "lns_init_%s_Sub( _pEnv );", moduleFullName) )
   else
    
      for __index, declAlgeNode in pairs( node:get_nodeManager():getDeclAlgeNodeList(  ) ) do
         filter( declAlgeNode, self, node )
      end
      
      
      for __index, child in pairs( node:get_children() ) do
         do
            local _switchExp = child:get_kind()
            if _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_DeclFunc() or _switchExp == Nodes.NodeKind.get_DeclMacro() or _switchExp == Nodes.NodeKind.get_TestBlock() then
            else 
               
                  filter( child, self, node )
                  self:writeln( "" )
            end
         end
         
      end
      
   end
   
   
   self:writeln( "lns_leave_block( _pEnv );" )
   
   self:writeln( string.format( "return &%s;", moduleInfoName) )
   
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

   return false
end
function BuiltinArgSymbolInfo:get_mutMode(  )

   return Ast.MutMode.IMut
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
   for __index, symbol in pairs( builtin:get_allSymbol() ) do
      local param
      
      do
         local _switchExp = symbol:get_kind()
         if _switchExp == Ast.SymbolKind.Mtd or _switchExp == Ast.SymbolKind.Fun then
            local retTypeList = symbol:get_typeInfo():get_retTypeInfoList()
            param = createSymbolParam( symbol:get_name(), getRetKind( retTypeList ), getCRetType( retTypeList ) )
         elseif _switchExp == Ast.SymbolKind.Mbr or _switchExp == Ast.SymbolKind.Var then
            param = createSymbolParam( symbol:get_name(), getValKind( symbol:get_typeInfo() ), getCType( symbol:get_typeInfo() ) )
         else 
            
               Util.err( string.format( "illeal symbol -- %s %d", symbol:get_name(), 1493) )
         end
      end
      
      symbol:set_convModuleParam( param )
   end
   
end

function convFilter:processBuiltin(  )

   local nodeManager = Nodes.NodeManager.new()
   local dummyPos = Parser.Position.new(0, 0)
   
   local function createNodeFromSymbol( classInfo, symbol )
   
      local token = Parser.Token.new(Parser.TokenKind.Symb, symbol:get_name(), dummyPos, false)
      do
         local _switchExp = symbol:get_kind()
         if _switchExp == Ast.SymbolKind.Mtd or _switchExp == Ast.SymbolKind.Fun then
            local argList = {}
            for index, argType in pairs( symbol:get_typeInfo():get_argTypeInfoList() ) do
               local argToken = Parser.Token.new(Parser.TokenKind.Symb, string.format( "arg%d", index), dummyPos, false)
               local dummyScope = Ast.Scope.new(nil, false)
               local argSym = BuiltinArgSymbolInfo.new(dummyScope, argToken.txt, argType, nil, symbol:get_typeInfo())
               
               table.insert( argList, Nodes.DeclArgNode.create( nodeManager, dummyPos, false, {argType}, argToken, argSym ) )
            end
            
            
            if classInfo ~= nil then
               local declFuncInfo = Nodes.DeclFuncInfo.new(Nodes.FuncKind.Mtd, classInfo, nil, token, argList, false, Ast.AccessMode.Pub, nil, symbol:get_typeInfo():get_retTypeInfoList(), false, false)
               return Nodes.DeclMethodNode.create( nodeManager, dummyPos, false, {symbol:get_typeInfo()}, declFuncInfo )
            else
               local declFuncInfo = Nodes.DeclFuncInfo.new(Nodes.FuncKind.Func, nil, nil, token, argList, false, Ast.AccessMode.Pub, nil, symbol:get_typeInfo():get_retTypeInfoList(), false, false)
               return Nodes.DeclFuncNode.create( nodeManager, dummyPos, false, {symbol:get_typeInfo()}, declFuncInfo )
            end
            
         elseif _switchExp == Ast.SymbolKind.Var then
            local varToken = Parser.Token.new(Parser.TokenKind.Symb, symbol:get_name(), dummyPos, false)
            
            return Nodes.DeclVarNode.create( nodeManager, dummyPos, false, {symbol:get_typeInfo()}, Nodes.DeclVarMode.Let, Ast.AccessMode.Pub, true, {Nodes.VarInfo.new(varToken, nil, symbol:get_typeInfo())}, nil, {symbol}, {symbol:get_typeInfo()}, false, nil, nil, {}, nil )
         elseif _switchExp == Ast.SymbolKind.Mbr then
            return nil
         else 
            
               Util.err( string.format( "illegal kind -- %s", Ast.SymbolKind:_getTxt( symbol:get_kind())
               ) )
         end
      end
      
   end
   
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
                  
                  local declClassNode = Nodes.DeclClassNode.create( nodeManager, dummyPos, false, {classInfo}, Ast.AccessMode.Pub, Parser.Token.new(Parser.TokenKind.Symb, classInfo:get_rawTxt(), dummyPos, false), classInfo:get_rawTxt(), nil, fieldList, {}, fieldList, {}, classScope, Nodes.ClassInitBlockInfo.new(nil), {}, {}, {}, {} )
                  
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
                           do
                              local node = createNodeFromSymbol( classInfo, field )
                              if node ~= nil then
                                 table.insert( fieldList, node )
                              end
                           end
                           
                        end
                     end
                  end
                  
               end
               
         end
      end
      
   end
   
   for __index, symbol in pairs( builtin:get_allSymbol() ) do
      if symbol:get_kind() == Ast.SymbolKind.Var or symbol:get_kind() == Ast.SymbolKind.Fun or symbol:get_namespaceTypeInfo():get_kind() == Ast.TypeInfoKind.Root then
         createNodeFromSymbol( nil, symbol )
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
   
   
   self.scopeMgr:setup( self.ast:get_moduleScope(), node:get_nodeManager():getDeclMemberNodeList(  ) )
   
   Ast.pushProcessInfo( node:get_processInfo() )
   
   for pragma, __val in pairs( node:get_luneHelperInfo().pragmaSet ) do
      do
         local _matchExp = pragma
         if _matchExp[1] == LuneControl.Pragma.can_not_conv_code[1] then
            local codeSet = _matchExp[2][1]
         
            if _lune._Set_has(codeSet, LuneControl.Code.C ) then
               self.canConv = false
               break
            end
            
         end
      end
      
   end
   
   
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
   
   
   self.processMode = ProcessMode.Prototype
   for __index, workNode in pairs( nodeManager:getTestBlockNodeList(  ) ) do
      filter( workNode, self, node )
   end
   
   
   local moduleName = self.moduleCtrl:getFullName( node:get_moduleTypeInfo() )
   self:write( string.format( 'static %s lns_moduleInfo_%s = {NULL,NULL,false, NULL, "%s", {', cTypeMod, moduleName, moduleName) )
   for __index, workNode in pairs( nodeManager:getTestBlockNodeList(  ) ) do
      self:write( string.format( "%s__test_%s, ", moduleName, workNode:get_name().txt) )
   end
   
   self:writeln( "NULL } };" )
   self:writeln( string.format( "static %s lns_module_globalStemList;", cTypeAnyPP) )
   self:writeln( string.format( "static %s lns_module_path = NULL;", cTypeAnyPP) )
   
   local function process( onlyPub )
   
      
      
      self.processMode = ProcessMode.Prototype
      
      
      for __index, workNode in pairs( nodeManager:getDeclEnumNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclFormNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclFuncNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclAlgeNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclClassNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclConstrNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclMethodNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getProtoMethodNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      if self.canConv then
         
         for __index, workNode in pairs( nodeManager:getExpToDDDNodeList(  ) ) do
            if not workNode:get_macroArgFlag() then
               if onlyPub then
                  if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                     filter( workNode, self, node )
                  end
                  
               else
                
                  filter( workNode, self, node )
               end
               
            end
            
         end
         
         
         
         for __index, workNode in pairs( nodeManager:getLiteralStringNodeList(  ) ) do
            if not workNode:get_macroArgFlag() then
               if onlyPub then
                  if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                     filter( workNode, self, node )
                  end
                  
               else
                
                  filter( workNode, self, node )
               end
               
            end
            
         end
         
         
         
         for __index, workNode in pairs( nodeManager:getExpCastNodeList(  ) ) do
            if not workNode:get_macroArgFlag() then
               if onlyPub then
                  if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                     filter( workNode, self, node )
                  end
                  
               else
                
                  filter( workNode, self, node )
               end
               
            end
            
         end
         
         
      end
      
      
      self.processMode = ProcessMode.WideScopeVer
      for __index, decl in pairs( nodeManager:getDeclVarNodeList(  ) ) do
         filter( decl, self, node )
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclAlgeNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclClassNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclConstrNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclMethodNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      for __index, workNode in pairs( nodeManager:getDeclFuncNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
      
      self.processMode = ProcessMode.DefClass
      
      for __index, workNode in pairs( nodeManager:getDeclClassNodeList(  ) ) do
         if not workNode:get_macroArgFlag() then
            if onlyPub then
               if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
                  filter( workNode, self, node )
               end
               
            else
             
               filter( workNode, self, node )
            end
            
         end
         
      end
      
      
   end
   
   process( not self.canConv )
   
   if self.canConv then
      self.processMode = ProcessMode.StringFormat
      for __index, litStr in pairs( nodeManager:getLiteralStringNodeList(  ) ) do
         if not litStr:get_macroArgFlag() then
            filter( litStr, self, node )
         end
         
      end
      
      
      self.processMode = ProcessMode.Immediate
      self.processedNodeSet = {}
      
      local function procssLiteralCtor( literalNodeList )
      
         for __index, literalNode in pairs( literalNodeList ) do
            self.processingNode = literalNode
            if not _lune._Set_has(self.processedNodeSet, literalNode ) and not literalNode:get_macroArgFlag() then
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
      
      for __index, castNode in pairs( nodeManager:getExpCastNodeList(  ) ) do
         filter( castNode, self, node )
      end
      
      
      self.processMode = ProcessMode.DefWrap
      for __index, callNode in pairs( nodeManager:getExpCallNodeList(  ) ) do
         filter( callNode, self, node )
      end
      
      
      self.processMode = ProcessMode.InitFuncSym
      self:writeln( string.format( "static void initFuncSym( %s _pEnv, %s pBlock )\n{", cTypeEnvP, cTypeBlockP) )
      self:pushIndent(  )
      
      for __index, declConstrNode in pairs( nodeManager:getDeclConstrNodeList(  ) ) do
         filter( declConstrNode, self, node )
      end
      
      for __index, declMethodNode in pairs( nodeManager:getDeclMethodNodeList(  ) ) do
         filter( declMethodNode, self, node )
      end
      
      for __index, declFuncNode in pairs( nodeManager:getDeclFuncNodeList(  ) ) do
         filter( declFuncNode, self, node )
      end
      
      
      self:popIndent(  )
      self:writeln( "}" )
      
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
      
      for __index, declMethodNode in pairs( nodeManager:getProtoMethodNodeList(  ) ) do
         filter( declMethodNode, self, node )
      end
      
      for __index, declFormNode in pairs( nodeManager:getDeclFormNodeList(  ) ) do
         filter( declFormNode, self, node )
      end
      
      for __index, declFuncNode in pairs( nodeManager:getDeclFuncNodeList(  ) ) do
         self.duringDeclFunc = false
         filter( declFuncNode, self, node )
      end
      
   end
   
   
   self:processInitModule( node )
   
   for __index, testBlock in pairs( nodeManager:getTestBlockNodeList(  ) ) do
      filter( testBlock, self, node )
   end
   
   
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

function convFilter:processBlockPreProcess( scope )

   self:pushIndent(  )
   local anyNum, stemNum, varNum = self.scopeMgr:setupScopeParam( scope )
   self:writeln( string.format( "lns_block_t * %s = lns_enter_block( _pEnv, %d, %d, %d );", getBlockName( scope ), anyNum, stemNum, varNum) )
   self.routineInfoStack:pushDepth(  )
   self.loopInfoStack:pushDepth(  )
end


function convFilter:processBlockPostProcess(  )

   self.loopInfoStack:popDepth(  )
   self.routineInfoStack:popDepth(  )
   self:writeln( "lns_leave_block( _pEnv );" )
   self:popIndent(  )
end


function convFilter:pushRoutine( funcType, blockNode )

   self:processBlockPreProcess( blockNode:get_scope() )
   self.routineInfoStack:newInfo( RoutineInfo.new(funcType) )
end


function convFilter:popRoutine(  )

   self.routineInfoStack:delInfo(  )
   self:processBlockPostProcess(  )
end


function convFilter:processLoopPreProcess( blockNode )

   self:processBlockPreProcess( blockNode:get_scope() )
   self.loopInfoStack:newInfo( DepthInfo.new() )
end


function convFilter:processLoopPostProcess(  )

   self.loopInfoStack:delInfo(  )
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
            self:write( string.format( "%s %s = l_form_closure_var( _pForm, %d )", typeTxt, self.moduleCtrl:getSymbolName( symbol ), _lune.unwrap( scope:get_closureSym2NumMap()[symbol])) )
            
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
         loopFlag = true
      elseif _switchExp == Nodes.BlockKind.Repeat then
         word = ""
         loopFlag = true
      elseif _switchExp == Nodes.BlockKind.For then
         word = ""
         loopFlag = true
      elseif _switchExp == Nodes.BlockKind.Apply then
         word = ""
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
         readyBlock = true
         word = ""
      elseif _switchExp == Nodes.BlockKind.LetUnwrapThenDo then
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
      self:processBlockPreProcess( node:get_scope() )
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
         return string.format( "LNS_STEM_INT( %s )", valTxt)
      elseif _switchExp == Ast.builtinTypeReal then
         return string.format( "LNS_STEM_REAL( %s )", valTxt)
      elseif _switchExp == Ast.builtinTypeBool then
         return string.format( "LNS_STEM_BOOL( %s )", valTxt)
      else 
         
            return "NULL"
      end
   end
   
end

local function getStemTypeId( typeInfo )

   if typeInfo:get_nilable() then
      return "lns_stem_type_none"
   end
   
   do
      local _switchExp = typeInfo
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         return "lns_stem_type_int"
      elseif _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeReal then
         return "lns_stem_type_real"
      elseif _switchExp == Ast.builtinTypeBool then
         return "lns_stem_type_bool"
      else 
         
            return "lns_stem_type_any"
      end
   end
   
end

local function getPrepareClosure( scopeMgr, funcName, argNum, hasDDD, symList )

   local txt
   
   
   txt = string.format( "lns_func2any( _pEnv, (lns_closure_t *)%s, %d, %s, %d", funcName, argNum, tostring( hasDDD), #symList)
   for __index, symbolInfo in pairs( symList ) do
      txt = txt .. ", "
      txt = txt .. scopeMgr:symbol2Any( symbolInfo )
   end
   
   txt = txt .. ")"
   return txt
end

local function getFunc2any( moduleCtrl, scopeMgr, funcType )

   local argList = funcType:get_argTypeInfoList()
   local hasDDD = #argList > 0 and argList[#argList]:get_kind() == Ast.TypeInfoKind.DDD or false
   
   return getPrepareClosure( scopeMgr, moduleCtrl:getFuncName( funcType ), #funcType:get_argTypeInfoList(), hasDDD, (_lune.unwrap( funcType:get_scope()) ):get_closureSymList() )
end

function ScopeMgr:symbol2Any( symbol )

   if symbol:get_kind() == Ast.SymbolKind.Fun then
      return getFunc2any( self.moduleCtrl, self, symbol:get_typeInfo() )
   end
   
   local name = self.moduleCtrl:getSymbolName( symbol )
   do
      local _switchExp = self:getSymbolValKind( symbol )
      if _switchExp == ValKind.Var then
         return name
      else 
         
            Util.err( string.format( "not support -- %s", symbol:get_typeInfo():getTxt(  )) )
      end
   end
   
end


function convFilter:processSym2stem( symbolInfo )

   local valKind = self.scopeMgr:getSymbolValKind( symbolInfo )
   do
      local _switchExp = valKind
      if _switchExp == ValKind.Any then
         self:write( "LNS_STEM_ANY( " )
         if symbolInfo:get_kind() == Ast.SymbolKind.Var then
            self:write( "*" )
         end
         
         self:write( self.moduleCtrl:getSymbolName( symbolInfo ) )
         self:write( ")" )
         return 
      elseif _switchExp == ValKind.Var then
         self:write( self.moduleCtrl:getSymbolName( symbolInfo ) )
         self:write( "->stem" )
         return 
      elseif _switchExp == ValKind.Stem then
         self:write( self.moduleCtrl:getSymbolName( symbolInfo ) )
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
         self:write( "LNS_STEM_INT( " )
         self:write( "" )
         self:write( self.scopeMgr:getAccessPrimValFromSymbol( symbolInfo ) )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeReal then
         self:write( "LNS_STEM_REAL( " )
         self:write( self.scopeMgr:getAccessPrimValFromSymbol( symbolInfo ) )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeBool then
         self:write( "LNS_STEM_BOOL( " )
         self:write( self.scopeMgr:getAccessPrimValFromSymbol( symbolInfo ) )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeStem or _switchExp == Ast.builtinTypeStem_ then
         self:write( self.moduleCtrl:getSymbolName( symbolInfo ) )
      else 
         
            do
               local _switchExp = expType:get_kind()
               if _switchExp == Ast.TypeInfoKind.DDD then
                  self:write( "_pDDD" )
               elseif _switchExp == Ast.TypeInfoKind.Func then
                  do
                     local scope = expType:get_scope()
                     if scope ~= nil then
                        self:write( "LNS_STEM_ANY(" )
                        self:write( getFunc2any( self.moduleCtrl, self.scopeMgr, expType ) )
                        self:write( ")" )
                     else
                        Util.err( "illegal func" )
                     end
                  end
                  
               else 
                  
                     self:write( self.moduleCtrl:getSymbolName( symbolInfo ) )
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
                        local valTxt = string.format( "%s", tostring( Ast.getEnumLiteralVal( valInfo:get_val() )))
                        self:writeln( string.format( "#define %s %s", enumValName, valTxt) )
                     end
                  end
                  
               end
               
            end
            
            do
               local _switchExp = out2HMode
               if _switchExp == Out2HMode.HeaderPub or _switchExp == Out2HMode.SourcePri then
                  self:writeln( string.format( "%s%s %s_get__allList( lns_env_t * _pEnv );", prefix, cTypeAnyP, enumFullName) )
                  self:writeln( string.format( "%s%s %s_get__txt( %s _pEnv, %s val );", prefix, cTypeAnyP, enumFullName, cTypeEnvP, getCType( enumType:get_valTypeInfo() )) )
                  
                  self:writeln( string.format( "%s%s %s( %s _pEnv, %s val );", prefix, cTypeStem, self.moduleCtrl:getEnumFuncName( enumType, "_from" ), cTypeEnvP, getCType( enumType:get_valTypeInfo() )) )
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
         
         self:writeln( string.format( "%s %s_get__allList( lns_env_t * _pEnv )", cTypeAnyP, enumFullName) )
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
         
         
         do
            
            if not Ast.isPubToExternal( enumType:get_accessMode() ) then
               self:write( "static " )
            end
            
            self:writeln( string.format( "%s %s_get__txt( %s _pEnv, %s val ) {", cTypeAnyP, enumFullName, cTypeEnvP, getCType( enumType:get_valTypeInfo() )) )
            self:pushIndent(  )
            
            self:write( string.format( "%s _work =  lns_mtd_Map_get( _pEnv, %s, ", cTypeStem, self.moduleCtrl:getEnumVal2NameMapName( enumType )) )
            local workSym = WorkSymbol.new(_lune.unwrap( self.moduleTypeInfo:get_scope()), Ast.AccessMode.Local, "val", enumType:get_valTypeInfo(), Ast.SymbolKind.Arg, false, SymbolParam.new(getValKind( enumType:get_valTypeInfo() ), 0, getCType( enumType:get_valTypeInfo() )))
            
            self:processSym2stem( workSym )
            self:writeln( ");" )
            self:writeln( string.format( "return _work%s;", accessAny) )
            
            self:popIndent(  )
            self:writeln( "}" )
         end
         
         
         do
            
            if not Ast.isPubToExternal( enumType:get_accessMode() ) then
               self:write( "static " )
            end
            
            self:writeln( string.format( "%s %s( %s _pEnv, %s val ) {", cTypeStem, self.moduleCtrl:getEnumFuncName( enumType, "_from" ), cTypeEnvP, getCType( enumType:get_valTypeInfo() )) )
            self:pushIndent(  )
            
            self:write( string.format( "%s key = ", cTypeStem) )
            local workSym = WorkSymbol.new(_lune.unwrap( self.moduleTypeInfo:get_scope()), Ast.AccessMode.Local, "val", enumType:get_valTypeInfo(), Ast.SymbolKind.Arg, false, SymbolParam.new(getValKind( enumType:get_valTypeInfo() ), 0, getCType( enumType:get_valTypeInfo() )))
            self:processSym2stem( workSym )
            self:writeln( ";" )
            
            self:writeln( string.format( "%s _work = lns_mtd_Map_get( _pEnv, %s, key );", cTypeStem, self.moduleCtrl:getEnumVal2NameMapName( enumType )) )
            
            self:writeln( "if ( _work.type == lns_stem_type_nil ) {" )
            self:pushIndent(  )
            self:writeln( "return lns_global.nilStem;" )
            self:popIndent(  )
            self:writeln( "}" )
            self:writeln( "return key;" )
            
            self:popIndent(  )
            self:writeln( "}" )
         end
         
         
         self:writeln( string.format( "static void init_%s( lns_env_t * _pEnv )", enumFullName) )
         self:writeln( "{" )
         self:pushIndent(  )
         
         local anyVarList = {}
         
         if isStrEnum then
            for index, valName in pairs( node:get_valueNameList() ) do
               local valInfo = _lune.unwrap( enumType:getEnumValInfo( valName.txt ))
               
               local valTxt = string.format( '"%s"', tostring( Ast.getEnumLiteralVal( valInfo:get_val() )))
               local anyVar = self.moduleCtrl:getEnumValCName( enumType, valName.txt )
               table.insert( anyVarList, string.format( "LNS_STEM_ANY( %s )", anyVar) )
               
               self:writeln( string.format( "%s = lns_litStr2any( _pEnv, %s );", anyVar, valTxt) )
            end
            
         else
          
            for index, valName in pairs( node:get_valueNameList() ) do
               local valInfo = _lune.unwrap( enumType:getEnumValInfo( valName.txt ))
               local valTxt = string.format( '%s', tostring( Ast.getEnumLiteralVal( valInfo:get_val() )))
               local anyVar = string.format( "_%s", valName.txt)
               table.insert( anyVarList, anyVar )
               self:write( string.format( "%s %s = ", cTypeStem, anyVar) )
               self:write( getLiteral2Stem( valTxt, enumType:get_valTypeInfo() ) )
               self:writeln( ";" )
            end
            
         end
         
         
         local allListName = string.format( "%s_allList", enumFullName)
         self:write( allListName )
         self:writeln( " = lns_class_List_new( _pEnv );" )
         processAddModuleGlobal( self.stream, string.format( "LNS_STEM_ANY( %s )", allListName) )
         
         for __index, anyVar in pairs( anyVarList ) do
            self:writeln( string.format( "lns_mtd_List_insert( _pEnv, %s_allList, %s );", enumFullName, anyVar) )
         end
         
         
         local val2NameMapName = string.format( "%s_val2NameMap", enumFullName)
         self:write( val2NameMapName )
         self:writeln( " = lns_class_Map_new( _pEnv );" )
         processAddModuleGlobal( self.stream, string.format( "LNS_STEM_ANY( %s )", val2NameMapName) )
         
         for index, anyVar in pairs( anyVarList ) do
            self:writeln( string.format( "lns_mtd_Map_add( _pEnv, %s_val2NameMap, %s, ", enumFullName, anyVar) )
            
            self:writeln( string.format( '  LNS_STEM_ANY( lns_litStr2any( _pEnv, "%s.%s" ) ) );', fullName, node:get_valueNameList()[index].txt) )
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
   
   
   
   local algeTypeName = moduleCtrl:getAlgeCName( node:get_expType() )
   stream:writeln( string.format( "static %s %s_type2NameMap;", cTypeAnyP, algeTypeName) )
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
   
   
   local algeName = moduleCtrl:getAlgeCName( algeType )
   local type2NameMapName = string.format( "%s_type2NameMap", algeName)
   
   if not Ast.isPubToExternal( algeType:get_accessMode() ) then
      stream:write( "static " )
   end
   
   stream:writeln( string.format( "%s %s_get__txt( %s _pEnv, %s pAny ) {", cTypeAnyP, algeName, cTypeEnvP, cTypeAnyP) )
   stream:pushIndent(  )
   stream:writeln( string.format( "return lns_mtd_Map_get( _pEnv, %s, LNS_STEM_INT( pAny->val.alge.type ) )%s;", type2NameMapName, accessAny) )
   stream:popIndent(  )
   stream:writeln( "}" )
   
   stream:writeln( string.format( "static void %s( %s _pEnv ) {", moduleCtrl:getAlgeInitCName( algeType ), cTypeEnvP) )
   
   stream:pushIndent(  )
   
   stream:writeln( type2NameMapName .. " = lns_class_Map_new( _pEnv );" )
   processAddModuleGlobal( stream, string.format( "LNS_STEM_ANY( %s )", type2NameMapName) )
   
   local fullName = moduleCtrl:getFullName( algeType )
   local enumName = moduleCtrl:getAlgeEnumCName( algeType )
   
   for index, valInfo in pairs( valList ) do
      stream:write( string.format( "lns_mtd_Map_add( _pEnv, %s, LNS_STEM_INT( %s_%s ), ", type2NameMapName, enumName, valInfo:get_name()) )
      
      stream:writeln( string.format( 'LNS_STEM_ANY( lns_litStr2any( _pEnv, "%s.%s" ) ) );', fullName, valInfo:get_name()) )
   end
   
   
   for index, valInfo in pairs( valList ) do
      if #valInfo:get_typeList() == 0 then
         local varName = moduleCtrl:getAlgeValCName( algeType, valInfo:get_name() )
         stream:writeln( string.format( "lns_init_alge( &%s, &%s_any, %s_%s );", varName, varName, enumName, valInfo:get_name()) )
      end
      
   end
   
   
   stream:popIndent(  )
   stream:writeln( "}" )
   
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
            gcTxt = string.format( "lns_gc_alge_%s_%s", algeName, valInfo:get_name())
            stream:writeln( string.format( "static void %s( %s _pEnv, void * pVal ) {", gcTxt, cTypeEnvP) )
            stream:pushIndent(  )
            stream:writeln( string.format( "%s *pWorkVal = (%s *)pVal;", valStruct, valStruct) )
            
            for paramIndex, valType in pairs( valInfo:get_typeList() ) do
               if isStemType( valType ) then
                  stream:writeln( string.format( "lns_decre_ref( _pEnv, pWorkVal->_val%d.val.pAny );", paramIndex) )
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
         
         stream:writeln( string.format( "%s pAny = lns_alge_new( _pEnv, %s_%s, sizeof( %s ), %s );", cTypeAnyP, enumName, valInfo:get_name(), valStruct, gcTxt) )
         stream:writeln( string.format( "%s *pVal = pAny->val.alge.pVal;", valStruct) )
         
         for paramIndex, valType in pairs( valInfo:get_typeList() ) do
            if isStemType( valType ) then
               stream:writeln( string.format( "lns_setQ( pVal->_val%d, _val%d );", paramIndex, paramIndex) )
            else
             
               stream:writeln( string.format( "pVal->_val%d = _val%d;", paramIndex, paramIndex) )
            end
            
         end
         
         
         stream:writeln( "return LNS_STEM_ANY( pAny );" )
         
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
      local retType = retTypeList[1]:get_srcTypeInfo()
      do
         local enumType = _lune.__Cast( retType, 3, Ast.EnumTypeInfo )
         if enumType ~= nil then
            retType = enumType:get_valTypeInfo()
         end
      end
      
      do
         local _switchExp = retType
         if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
            return "lns_method_int_t"
         elseif _switchExp == Ast.builtinTypeReal then
            return "lns_method_real_t"
         elseif _switchExp == Ast.builtinTypeBool then
            return "lns_method_bool_t"
         end
      end
      
      if getValKind( retType ) == ValKind.Any then
         return "lns_method_any_t"
      end
      
   end
   
   return "lns_method_t"
end

local function processNewConstrProto( stream, moduleCtrl, node, out2HMode, outputBuiltinFlag )

   local className = moduleCtrl:getClassCName( node:get_expType() )
   
   stream:write( getOut2HeaderPrefix( out2HMode ) )
   stream:write( string.format( "%s %s( %s _pEnv", cTypeAnyP, moduleCtrl:getNewName( node:get_expType() ), cTypeEnvP) )
   
   if not outputBuiltinFlag then
      local scope = _lune.unwrap( node:get_expType():get_scope())
      local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, scopeAccess ))
      for index, argType in pairs( initFuncType:get_argTypeInfoList() ) do
         stream:write( string.format( ", %s arg%d", getCType( argType ), index) )
      end
      
   end
   
   stream:write( ")" )
end

local function processDeclAlgeSub( stream, node )

   stream:write( getCType( node:get_expType() ) )
   if node:get_symbolInfo():get_hasAccessFromClosure() then
      
      stream:write( ' _' )
   elseif node:get_symbolInfo():get_mutable() then
      
      stream:write( ' _' )
   else
    
      stream:write( ' ' )
   end
   
   stream:write( node:get_name(  ).txt )
end

local FuncWrap = {}
FuncWrap._val2NameMap = {}
function FuncWrap:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "FuncWrap.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function FuncWrap._from( val )
   if FuncWrap._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
FuncWrap.__allList = {}
function FuncWrap.get__allList()
   return FuncWrap.__allList
end

FuncWrap.Normal = 0
FuncWrap._val2NameMap[0] = 'Normal'
FuncWrap.__allList[1] = FuncWrap.Normal
FuncWrap.CallWrap = 1
FuncWrap._val2NameMap[1] = 'CallWrap'
FuncWrap.__allList[2] = FuncWrap.CallWrap
FuncWrap.NilWrap = 2
FuncWrap._val2NameMap[2] = 'NilWrap'
FuncWrap.__allList[3] = FuncWrap.NilWrap


local function processMethodDeclTxt( stream, moduleCtrl, wrapKind, methodTypeInfo, argList )

   if methodTypeInfo:get_rawTxt() ~= "__init" and wrapKind == FuncWrap.Normal then
      if not methodTypeInfo:get_staticFlag() or not Ast.isPubToExternal( methodTypeInfo:get_accessMode() ) or not Ast.isPubToExternal( methodTypeInfo:get_parentInfo():get_accessMode() ) then
         stream:write( "static " )
      end
      
   end
   
   
   local name
   
   local objDecl
   
   local retType
   
   do
      local _switchExp = wrapKind
      if _switchExp == FuncWrap.Normal then
         name = moduleCtrl:getMethodCName( methodTypeInfo )
         objDecl = string.format( ", %s pObj", cTypeAnyP)
         retType = getCRetType( methodTypeInfo:get_retTypeInfoList() )
      elseif _switchExp == FuncWrap.CallWrap then
         name = moduleCtrl:getCallMethodCName( methodTypeInfo )
         objDecl = string.format( ", %s pObj", cTypeAnyP)
         retType = getCRetType( methodTypeInfo:get_retTypeInfoList() )
      elseif _switchExp == FuncWrap.NilWrap then
         name = moduleCtrl:getNilMethodCName( methodTypeInfo )
         objDecl = string.format( ", %s obj", cTypeStem)
         if #methodTypeInfo:get_retTypeInfoList() == 0 then
            retType = "void"
         else
          
            retType = cTypeStem
         end
         
      else 
         
            Util.err( string.format( "not support -- %s", FuncWrap:_getTxt( wrapKind)
            ) )
      end
   end
   
   stream:write( string.format( "%s %s( %s _pEnv", retType, name, cTypeEnvP) )
   if methodTypeInfo:get_staticFlag() then
      if isClosure( methodTypeInfo ) then
         stream:write( string.format( ", %s _pForm", cTypeAnyP) )
      end
      
   else
    
      stream:write( objDecl )
   end
   
   
   if methodTypeInfo:get_rawTxt() == "___init" and methodTypeInfo:get_staticFlag() then
      
      stream:write( string.format( ", %s %s", cTypeBlockP, getBlockName( _lune.unwrap( methodTypeInfo:getModule(  ):get_scope()) )) )
   else
    
      if argList ~= nil then
         for index, argNode in pairs( argList ) do
            stream:write( ", " )
            do
               local declArgNode = _lune.__Cast( argNode, 3, Nodes.DeclArgNode )
               if declArgNode ~= nil then
                  processDeclAlgeSub( stream, declArgNode )
               else
                  stream:write( string.format( "%s _pDDD", cTypeStem) )
               end
            end
            
         end
         
      else
         for index, arg in pairs( methodTypeInfo:get_argTypeInfoList() ) do
            stream:write( string.format( ", %s arg%d", getCType( arg ), index) )
         end
         
      end
      
   end
   
   stream:write( ")" )
end

local function processDeclMethodTable( stream, classTypeInfo )

   local function outputField( name, retTypeList )
   
      local methodType = getMethodTypeTxt( retTypeList )
      stream:writeln( string.format( "%s * %s;", methodType, name) )
   end
   
   local nameSet = Ast.getAllMethodName( classTypeInfo, Ast.MethodKind.Object )
   for __index, name in pairs( nameSet:get_list() ) do
      do
         local symbolInfo = _lune.nilacc( classTypeInfo:get_scope(), 'getSymbolInfoField', 'callmtd' , name, true, _lune.unwrap( classTypeInfo:get_scope()), Ast.ScopeAccess.Normal )
         if symbolInfo ~= nil then
            outputField( name, symbolInfo:get_typeInfo():get_retTypeInfoList() )
         end
      end
      
   end
   
end

local function processDeclMemberTable( normal, stream, classTypeInfo )

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
   if normal then
      outputVal( _lune.unwrap( classTypeInfo:get_scope()) )
   else
    
      stream:writeln( "void * pExt;" )
   end
   
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

local function processPrototypeMethod( stream, moduleCtrl, declArgNodeList, funcTypeInfo )

   local function processHeader( out2HMode )
   
      if out2HMode == Out2HMode.HeaderPub then
         stream:write( "extern " )
      end
      
      if not funcTypeInfo:get_staticFlag() then
         
         processMethodDeclTxt( stream, moduleCtrl, FuncWrap.CallWrap, funcTypeInfo, declArgNodeList )
         stream:writeln( ";" )
         processMethodDeclTxt( stream, moduleCtrl, FuncWrap.NilWrap, funcTypeInfo, declArgNodeList )
         stream:writeln( ";" )
      else
       
         
         processMethodDeclTxt( stream, moduleCtrl, FuncWrap.Normal, funcTypeInfo, declArgNodeList )
         stream:writeln( ";" )
      end
      
   end
   
   if funcTypeInfo:get_parentInfo():get_kind() == Ast.TypeInfoKind.Class and funcTypeInfo:get_rawTxt() == "__init" then
      processMethodDeclTxt( stream, moduleCtrl, FuncWrap.Normal, funcTypeInfo, declArgNodeList )
      stream:writeln( ";" )
   else
    
      
      do
         local function processwork( out2HMode )
         
            do
               local _switchExp = out2HMode
               if _switchExp == Out2HMode.SourcePri then
                  if funcTypeInfo:get_parentInfo():get_kind() == Ast.TypeInfoKind.Class then
                     processHeader( out2HMode )
                     processMethodDeclTxt( stream, moduleCtrl, FuncWrap.Normal, funcTypeInfo, declArgNodeList )
                     stream:writeln( ";" )
                  end
                  
               elseif _switchExp == Out2HMode.SourcePub then
                  if funcTypeInfo:get_parentInfo():get_kind() == Ast.TypeInfoKind.Class then
                     if not funcTypeInfo:get_staticFlag() then
                        processMethodDeclTxt( stream, moduleCtrl, FuncWrap.Normal, funcTypeInfo, declArgNodeList )
                        stream:writeln( ";" )
                     end
                     
                  end
                  
               elseif _switchExp == Out2HMode.HeaderPub then
                  processHeader( out2HMode )
               end
            end
            
            
            
         end
         if Ast.isPubToExternal( funcTypeInfo:get_parentInfo():get_accessMode() ) and Ast.isPubToExternal( funcTypeInfo:get_accessMode() ) then
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



local function process2stem( stream, moduleCtrl, scopeMgr, valKind, typeInfo, parent, callback )

   do
      local _switchExp = valKind
      if _switchExp == ValKind.Stem then
         callback(  )
      else 
         
            local expType = typeInfo:get_srcTypeInfo()
            do
               local enumType = _lune.__Cast( expType, 3, Ast.EnumTypeInfo )
               if enumType ~= nil then
                  expType = enumType:get_valTypeInfo()
               end
            end
            
            
            do
               local _switchExp = expType
               if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                  stream:write( "LNS_STEM_INT( " )
                  callback(  )
                  stream:write( ")" )
               elseif _switchExp == Ast.builtinTypeReal then
                  stream:write( "LNS_STEM_REAL( " )
                  callback(  )
                  stream:write( ")" )
               elseif _switchExp == Ast.builtinTypeBool then
                  stream:write( "LNS_STEM_BOOL( " )
                  callback(  )
                  stream:write( ")" )
               else 
                  
                     do
                        local _switchExp = expType:get_kind()
                        if _switchExp == Ast.TypeInfoKind.DDD then
                           stream:write( "_pDDD" )
                        elseif _switchExp == Ast.TypeInfoKind.Func then
                           do
                              local scope = expType:get_scope()
                              if scope ~= nil then
                                 stream:write( "LNS_STEM_ANY(" )
                                 stream:write( getFunc2any( moduleCtrl, scopeMgr, expType ) )
                                 stream:write( ")" )
                              else
                                 Util.err( "illegal func" )
                              end
                           end
                           
                        else 
                           
                              if valKind == ValKind.Var and getValKind( expType ) == ValKind.Stem then
                                 callback(  )
                              else
                               
                                 stream:write( "LNS_STEM_ANY( " )
                                 callback(  )
                                 stream:write( ")" )
                              end
                              
                        end
                     end
                     
               end
            end
            
      end
   end
   
end

local function processDeclCallMethodWrapper( stream, moduleCtrl, scopeMgr, parent, funcTypeInfo, callFlag )

   if funcTypeInfo:get_rawTxt() ~= "__init" and not funcTypeInfo:get_staticFlag() then
      
      processMethodDeclTxt( stream, moduleCtrl, callFlag and FuncWrap.CallWrap or FuncWrap.NilWrap, funcTypeInfo, nil )
      stream:writeln( "{" )
      
      if not callFlag then
         local retVal
         
         do
            local _switchExp = #funcTypeInfo:get_retTypeInfoList()
            if _switchExp == 0 then
               retVal = ""
            elseif _switchExp == 1 then
               retVal = cValNil
            else 
               
                  retVal = cValDDD0
            end
         end
         
         stream:writeln( string.format( "if ( obj.type == lns_stem_type_nil ) { return %s; }", retVal) )
         stream:writeln( string.format( "%s pObj = obj.val.pAny;", cTypeAnyP) )
      end
      
      
      if #funcTypeInfo:get_retTypeInfoList() ~= 0 then
         stream:write( "return " )
      end
      
      
      local retList = funcTypeInfo:get_retTypeInfoList()
      local function process(  )
      
         local classTypeInfo = funcTypeInfo:get_parentInfo()
         local className = moduleCtrl:getClassCName( classTypeInfo )
         stream:write( string.format( "lns_mtd_%s( pObj )->%s( _pEnv, ", className, funcTypeInfo:get_rawTxt()) )
         if classTypeInfo:get_kind() == Ast.TypeInfoKind.IF then
            stream:write( "lns_getImpObj( pObj ) " )
         else
          
            stream:write( "pObj " )
         end
         
         for index, argType in pairs( funcTypeInfo:get_argTypeInfoList() ) do
            stream:write( string.format( ", arg%d", index) )
         end
         
         stream:write( ")" )
      end
      
      if not callFlag and #retList > 0 then
         process2stem( stream, moduleCtrl, scopeMgr, getRetKind( retList ), retList[1], parent, process )
      else
       
         process(  )
      end
      
      stream:writeln( ";" )
      
      stream:writeln( "}" )
   end
   
end

local function getAccessMember( className, obj, member )

   return string.format( "lns_obj_%s( %s )->%s", className, obj, member)
end

local function getAccessMethod( className, obj, method )

   return string.format( "lns_mtd_%s( %s )->%s", className, obj, method)
end

local function processAdvertise( stream, moduleCtrl, scopeMgr, processMode, node )

   
   local declMethodNameSet = {}
   for __index, field in pairs( node:get_fieldList() ) do
      do
         local declMethodNode = _lune.__Cast( field, 3, Nodes.DeclMethodNode )
         if declMethodNode ~= nil then
            do
               local name = declMethodNode:get_declInfo():get_name()
               if name ~= nil then
                  declMethodNameSet[name.txt]= true
               end
            end
            
         end
      end
      
   end
   
   
   for __index, member in pairs( node:get_memberList() ) do
      if member:get_getterMode() ~= Ast.AccessMode.None then
         declMethodNameSet["get_" .. member:get_name().txt]= true
      end
      
      if member:get_setterMode() ~= Ast.AccessMode.None then
         declMethodNameSet["set_" .. member:get_name().txt]= true
      end
      
   end
   
   
   for __index, advInfo in pairs( node:get_advertiseList() ) do
      local member = advInfo:get_member()
      stream:writeln( string.format( "// for advertise %s.%s --->", node:get_name().txt, member:get_name().txt) )
      
      for __index, name in pairs( Ast.getAllMethodName( member:get_expType(), Ast.MethodKind.Object ):get_list() ) do
         if not _lune._Set_has(declMethodNameSet, name ) then
            
            local methodSym = _lune.unwrap( node:get_scope():getSymbolInfoField( name, true, node:get_scope(), Ast.ScopeAccess.Normal ))
            
            local methodType = methodSym:get_typeInfo()
            if methodType:get_accessMode() ~= Ast.AccessMode.Pri then
               do
                  local _switchExp = processMode
                  if _switchExp == ProcessMode.Prototype then
                     processPrototypeMethod( stream, moduleCtrl, nil, methodType )
                  elseif _switchExp == ProcessMode.DefClass then
                     processDeclCallMethodWrapper( stream, moduleCtrl, scopeMgr, node, methodType, true )
                     processDeclCallMethodWrapper( stream, moduleCtrl, scopeMgr, node, methodType, false )
                     
                     processMethodDeclTxt( stream, moduleCtrl, FuncWrap.Normal, methodType, nil )
                     stream:writeln( "{" )
                     local className = moduleCtrl:getClassCName( node:get_expType() )
                     local memberClassName = moduleCtrl:getClassCName( member:get_expType() )
                     stream:pushIndent(  )
                     stream:writeln( string.format( "%s pVal = %s;", cTypeAnyP, getAccessMember( className, "pObj", member:get_name().txt )) )
                     if #methodType:get_retTypeInfoList() ~= 0 then
                        stream:write( "return " )
                     end
                     
                     stream:write( string.format( "%s( _pEnv, pVal", getAccessMethod( memberClassName, "pVal", name )) )
                     for index, argType in pairs( methodType:get_argTypeInfoList() ) do
                        stream:write( string.format( ", arg%d", index) )
                     end
                     
                     stream:writeln( ");" )
                     
                     stream:popIndent(  )
                     stream:writeln( "}" )
                  end
               end
               
            end
            
         end
         
      end
      
      stream:writeln( string.format( "// <-- for advertise %s.%s", node:get_name().txt, member:get_name().txt) )
   end
   
end

local function processDeclClassPrototype( normal, stream, moduleCtrl, node )

   local className = moduleCtrl:getClassCName( node:get_expType() )
   
   stream:writeln( string.format( "static void mtd_%s__del( lns_env_t * _pEnv, %s pObj );", className, cTypeAnyP) )
   if not normal then
      stream:writeln( string.format( "static void mtd_%s__delExt( lns_env_t * _pEnv, %s pObj );", className, cTypeAnyP) )
   end
   
   
   if hasGC( node:get_expType() ) then
      stream:writeln( string.format( "static void mtd_%s__gc( lns_env_t * _pEnv, %s pObj );", className, cTypeAnyP) )
   end
   
   for __index, member in pairs( node:get_memberList() ) do
      local memberName = member:get_name().txt
      if member:get_getterMode() ~= Ast.AccessMode.None then
         local getterType = _lune.unwrap( node:get_scope():getTypeInfoField( string.format( "get_%s", memberName), true, node:get_scope(), scopeAccess ))
         processMethodDeclTxt( stream, moduleCtrl, FuncWrap.Normal, getterType )
         stream:writeln( ";" )
      end
      
      if member:get_setterMode() ~= Ast.AccessMode.None then
         local setterType = _lune.unwrap( node:get_scope():getTypeInfoField( string.format( "set_%s", memberName), true, node:get_scope(), scopeAccess ))
         processMethodDeclTxt( stream, moduleCtrl, FuncWrap.Normal, setterType )
         stream:writeln( ";" )
      end
      
   end
   
end

local function processDefaultCtor( stream, moduleCtrl, scopeMgr, node )

   local className = moduleCtrl:getClassCName( node:get_expType() )
   
   if not node:hasUserInit(  ) then
      
      stream:write( string.format( "static void %s( lns_env_t * _pEnv, %s pAny", moduleCtrl:getCtorName( node:get_expType() ), cTypeAnyP) )
      
      local ctorType = _lune.unwrap( node:get_scope():getTypeInfoField( "__init", true, node:get_scope(), scopeAccess ))
      
      for index, argType in pairs( ctorType:get_argTypeInfoList() ) do
         stream:write( string.format( ", %s _arg%d", getCType( argType ), index) )
      end
      
      stream:writeln( ") {" )
      stream:pushIndent(  )
      
      local memberNum = 0
      for __index, member in pairs( node:get_memberList() ) do
         if not member:get_staticFlag() then
            memberNum = memberNum + 1
         end
         
      end
      
      
      local superArgNum
      
      do
         local baseScope = node:get_scope():get_inherit()
         if baseScope ~= nil then
            
            local superInitType = _lune.unwrap( baseScope:getTypeInfoField( "__init", true, baseScope, scopeAccess ))
            stream:write( string.format( "%s( _pEnv, pAny", moduleCtrl:getCtorName( node:get_expType():get_baseTypeInfo() )) )
            
            if #ctorType:get_argTypeInfoList() >= #superInitType:get_argTypeInfoList() + memberNum then
               superArgNum = #superInitType:get_argTypeInfoList()
               for index, argType in pairs( superInitType:get_argTypeInfoList() ) do
                  stream:write( string.format( ", _arg%d", index) )
               end
               
            else
             
               
               superArgNum = 0
               for index, argType in pairs( superInitType:get_argTypeInfoList() ) do
                  stream:write( string.format( ", %s", cValNil) )
               end
               
            end
            
            stream:writeln( ");" )
         else
            superArgNum = 0
         end
      end
      
      
      stream:writeln( string.format( "%s * pObj = lns_obj_%s( pAny );", className, className) )
      local argIndex = superArgNum
      for index, member in pairs( node:get_memberList() ) do
         if not member:get_staticFlag() then
            argIndex = argIndex + 1
            local valKind = scopeMgr:getSymbolValKind( member:get_symbolInfo() )
            do
               local _switchExp = valKind
               if _switchExp == ValKind.Stem then
                  stream:writeln( string.format( "lns_setQ( pObj->%s, _arg%d );", member:get_name().txt, argIndex) )
               elseif _switchExp == ValKind.Any then
                  stream:writeln( string.format( "lns_setQ_any( &pObj->%s, _arg%d );", member:get_name().txt, argIndex) )
               elseif _switchExp == ValKind.Prim then
                  stream:writeln( string.format( "pObj->%s = _arg%d;", member:get_name().txt, argIndex) )
               else 
                  
                     Util.err( string.format( "no support -- %s:%s:%d", member:get_name().txt, ValKind:_getTxt( valKind)
                     , 3702) )
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
      stream:writeln( string.format( "lns_init_if( &pObj->imp.%s, _pEnv, pAny, &lns_if_%s_imp_%s, %s );", ifName, className, ifName, ifName) )
   end
   
end

function convFilter:processNewInsance( classType, callInit )

   local className = self.moduleCtrl:getClassCName( classType )
   
   self:writeln( string.format( "lns_class_new_( _pEnv, %s, pAny, pObj );", className) )
   
   if callInit and not self.outputBuiltinFlag then
      self:write( string.format( "%s( _pEnv, pAny", self.moduleCtrl:getCtorName( classType )) )
      
      local scope = _lune.unwrap( classType:get_scope())
      if not self.outputBuiltinFlag then
         local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, scopeAccess ))
         for index, argType in pairs( initFuncType:get_argTypeInfoList() ) do
            self:write( string.format( ", arg%d", index) )
         end
         
      end
      
      self:writeln( ");" )
   end
   
   
   if (self.outputBuiltinFlag or not self.canConv ) then
      self:writeln( "pObj->pExt = NULL;" )
   end
   
   self:writeln( "pObj->pImp = &pObj->imp;" )
   self:writeln( "pObj->imp.sentinel.type = lns_value_type_none;" )
   
   processIFObjInit( self.stream, self.moduleCtrl, classType, classType )
end


function convFilter:processMapping( node, classType, out2HMode )

   if not classType:isInheritFrom( Ast.builtinTypeMapping, nil ) then
      return 
   end
   
   
   local classScope = _lune.unwrap( classType:get_scope())
   local toMapMtdSym = _lune.unwrap( classScope:getSymbolInfoChild( "_toMap" ))
   local fromMapMtdSym = _lune.unwrap( classScope:getSymbolInfoChild( "_fromMap" ))
   local className = self.moduleCtrl:getClassCName( classType )
   
   local function processDeclToMap( callFlag )
   
      self:write( string.format( "%s%s ", getOut2HeaderPrefix( out2HMode ), cTypeAnyP) )
      if callFlag then
         self:write( self.moduleCtrl:getCallMethodCName( toMapMtdSym:get_typeInfo() ) )
      else
       
         self:write( self.moduleCtrl:getMethodCName( toMapMtdSym:get_typeInfo() ) )
      end
      
      self:write( string.format( "( %s _pEnv, %s pObj)", cTypeEnvP, getCRetType( toMapMtdSym:get_typeInfo():get_retTypeInfoList() )) )
   end
   
   local function processDeclFromMap( sub )
   
      self:write( string.format( "%s%s ", getOut2HeaderPrefix( out2HMode ), cTypeStem) )
      self:write( self.moduleCtrl:getMethodCName( fromMapMtdSym:get_typeInfo() ) )
      if sub then
         self:write( "Sub" )
      end
      
      self:write( string.format( "( %s _pEnv", cTypeEnvP) )
      if sub then
         self:write( ", const lns_fromVal_info_t * pInfoArray" )
      end
      
      self:write( string.format( ", %s mapStem)", getCRetType( fromMapMtdSym:get_typeInfo():get_retTypeInfoList() )) )
   end
   
   local function processToMapBody(  )
   
      
      processDeclToMap( false )
      self:writeln( "{" )
      self:pushIndent(  )
      
      self:writeln( string.format( "%s pMap = lns_class_Map_new( _pEnv );", cTypeAnyP) )
      for __index, varName in pairs( Ast.getAllNameForKind( classType, Ast.MethodKind.Object, Ast.SymbolKind.Mbr ):get_list() ) do
         
         self:write( "lns_mtd_Map_add( _pEnv, pMap, " )
         self:write( string.format( 'LNS_STEM_ANY( lns_litStr2any( _pEnv, "%s" ) ), ', varName) )
         
         local memberSym = _lune.unwrap( classScope:getSymbolInfoField( varName, true, classScope, Ast.ScopeAccess.Full ))
         local nonNilMemberType = memberSym:get_typeInfo():get_nonnilableType():get_srcTypeInfo()
         
         local valKind = getValKind( memberSym:get_typeInfo() )
         local valTxt = getAccessMember( className, "pObj", varName )
         do
            local _switchExp = valKind
            if _switchExp == ValKind.Prim then
            elseif _switchExp == ValKind.Stem then
               self:writeln( string.format( "lns_toMapFromStem( _pEnv, %s ) );", tostring( valTxt)) )
               valTxt = nil
            elseif _switchExp == ValKind.Any then
               if nonNilMemberType == Ast.builtinTypeString then
               else
                
                  self:writeln( string.format( "lns_toMapFromStem( _pEnv, LNS_STEM_ANY( %s ) ) );", tostring( valTxt)) )
                  valTxt = nil
               end
               
            else 
               
                  Util.err( string.format( "not support -- %s", ValKind:_getTxt( valKind)
                  ) )
            end
         end
         
         if valTxt ~= nil then
            process2stem( self.stream, self.moduleCtrl, self.scopeMgr, getValKind( memberSym:get_typeInfo() ), memberSym:get_typeInfo(), node, function (  )
            
               self:write( valTxt )
            end )
            self:writeln( ");" )
         end
         
      end
      
      
      self:writeln( "return pMap;" )
      
      self:popIndent(  )
      self:writeln( "}" )
      
      processDeclToMap( true )
      self:writeln( "{" )
      self:pushIndent(  )
      
      self:write( "return " )
      self:write( getAccessMethod( className, "pObj", "_toMap" ) )
      self:writeln( "( _pEnv, pObj );" )
      
      self:popIndent(  )
      self:writeln( "}" )
   end
   
   local function processFromMapBody(  )
   
      
      processDeclFromMap( false )
      self:writeln( "{" )
      self:pushIndent(  )
      
      self:writeln( string.format( "return %sSub( _pEnv, NULL, mapStem );", self.moduleCtrl:getMethodCName( fromMapMtdSym:get_typeInfo() )) )
      
      self:popIndent(  )
      self:writeln( "}" )
      
      processDeclFromMap( true )
      self:writeln( "{" )
      self:pushIndent(  )
      
      self:writeln( string.format( "if ( mapStem.type == lns_stem_type_nil ) { return %s; }", cValNil) )
      self:writeln( string.format( "lns_any_t * pMap = mapStem%s;", accessAny) )
      self:writeln( "lns_any_t * pErr = NULL;" )
      
      for __index, varName in pairs( Ast.getAllNameForKind( classType, Ast.MethodKind.Object, Ast.SymbolKind.Mbr ):get_list() ) do
         local memberSym = _lune.unwrap( classScope:getSymbolInfoField( varName, true, classScope, Ast.ScopeAccess.Full ))
         local nonNilMemberType = memberSym:get_typeInfo():get_nonnilableType():get_srcTypeInfo()
         
         local fromMapSym
         
         do
            local _switchExp = nonNilMemberType:get_kind()
            if _switchExp == Ast.TypeInfoKind.List then
               fromMapSym = "lns_fromMapToList"
            elseif _switchExp == Ast.TypeInfoKind.Array then
               fromMapSym = "lns_fromMapToArray"
            elseif _switchExp == Ast.TypeInfoKind.Set then
               fromMapSym = "lns_fromMapToSet"
            elseif _switchExp == Ast.TypeInfoKind.Map then
               fromMapSym = "lns_fromMapToMap"
            else 
               
                  do
                     local memberClassScope = nonNilMemberType:get_scope()
                     if memberClassScope ~= nil then
                        do
                           local symbol = memberClassScope:getSymbolInfoField( "_fromMap", true, memberClassScope, Ast.ScopeAccess.Normal )
                           if symbol ~= nil then
                              fromMapSym = self.moduleCtrl:getMethodCName( symbol:get_typeInfo() )
                           else
                              fromMapSym = nil
                           end
                        end
                        
                     else
                        fromMapSym = nil
                     end
                  end
                  
            end
         end
         
         
         local function process( nilable )
         
            local kind
            
            do
               local _switchExp = nonNilMemberType
               if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                  kind = "lns_stem_type_int"
               elseif _switchExp == Ast.builtinTypeReal then
                  kind = "lns_stem_type_real"
               elseif _switchExp == Ast.builtinTypeBool then
                  kind = "lns_stem_type_bool"
               else 
                  
                     Util.err( string.format( "not support -- %s", memberSym:get_typeInfo():getTxt(  )) )
               end
            end
            
            self:writeln( string.format( "lns_check_err_from_map( pErr, _pEnv, pMap, %s, %s, %s, %s );", nilable and "true" or "false", memberSym:get_name(), kind, getAccessPrimValFromStem( false, memberSym:get_typeInfo(), 0 )) )
         end
         
         self:writeln( string.format( "%s %s;", getCType( memberSym:get_typeInfo() ), memberSym:get_name()) )
         local nilable = memberSym:get_typeInfo():get_nilable()
         if nilable then
            self:writeln( string.format( "%s = lns_global.nilStem;", memberSym:get_name()) )
         end
         
         
         if fromMapSym ~= nil then
            self:write( "lns_check_err_from_map_class" )
            local infoValName
            
            if #memberSym:get_typeInfo():get_itemTypeInfoList() > 0 then
               infoValName = string.format( "&info_0_1_%s_%s", className, memberSym:get_name())
            else
             
               infoValName = "NULL"
            end
            
            
            self:writeln( string.format( "( pErr, _pEnv, pMap, %s, %s, %sSub, %s, %s );", tostring( nilable), memberSym:get_name(), fromMapSym, infoValName, getAccessPrimValFromStem( false, memberSym:get_typeInfo(), 0 )) )
         else
            if nonNilMemberType:equals( Ast.builtinTypeString ) then
               self:writeln( string.format( "lns_check_err_from_map_str( pErr, _pEnv, pMap, %s, %s, %s );", tostring( nilable), memberSym:get_name(), getAccessPrimValFromStem( false, memberSym:get_typeInfo(), 0 )) )
            elseif nonNilMemberType:get_kind() == Ast.TypeInfoKind.Alternate then
               self:writeln( string.format( "lns_check_err_from_map_stem( pErr, _pEnv, pMap, %s, %s );", tostring( nilable), memberSym:get_name()) )
            else
             
               process( nilable )
            end
            
         end
         
      end
      
      
      self:writeln( "if ( pErr != NULL ) {" )
      self:pushIndent(  )
      
      self:write( string.format( "return lns_createMRet( _pEnv, false, 2, %s, ", cValNil) )
      self:writeln( "LNS_STEM_ANY( pErr ) );" )
      
      self:popIndent(  )
      self:writeln( "}" )
      
      self:processNewInsance( classType, false )
      
      for __index, varName in pairs( Ast.getAllNameForKind( classType, Ast.MethodKind.Object, Ast.SymbolKind.Mbr ):get_list() ) do
         local memberSym = _lune.unwrap( classScope:getSymbolInfoField( varName, true, classScope, Ast.ScopeAccess.Full ))
         
         do
            local _switchExp = getValKind( memberSym:get_typeInfo() )
            if _switchExp == ValKind.Stem then
               self:writeln( string.format( "lns_setQ( pObj->%s, %s );", varName, varName) )
            elseif _switchExp == ValKind.Any then
               self:writeln( string.format( "lns_setQ_any( &pObj->%s, %s );", varName, varName) )
            elseif _switchExp == ValKind.Prim then
               self:writeln( string.format( "pObj->%s = %s;", varName, varName) )
            end
         end
         
      end
      
      
      self:writeln( "return lns_createMRet( _pEnv, false, 1, LNS_STEM_ANY( pAny ) );" )
      
      self:popIndent(  )
      self:writeln( "}" )
   end
   
   local function processFromMapInfo( memberSym )
   
      if #memberSym:get_typeInfo():get_nonnilableType():get_itemTypeInfoList() == 0 then
         return 
      end
      
      
      local function processGenType( genType, name, depth, index )
      
         self:write( string.format( "const lns_fromVal_info_t info_%d_%d_%s = { ", depth, index, name) )
         self:write( string.format( "%s, ", tostring( genType:get_nilable())) )
         do
            local _switchExp = genType:get_nonnilableType():get_srcTypeInfo()
            if _switchExp == Ast.builtinTypeStem then
               self:write( "lns_fromMapToStemSub" )
            elseif _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
               self:write( "lns_fromMapToIntSub" )
            elseif _switchExp == Ast.builtinTypeReal then
               self:write( "lns_fromMapToRealSub" )
            elseif _switchExp == Ast.builtinTypeBool then
               self:write( "lns_fromMapToBoolSub" )
            elseif _switchExp == Ast.builtinTypeString then
               self:write( "lns_fromMapToStrSub" )
            else 
               
                  do
                     local _switchExp = genType:get_nonnilableType():get_kind()
                     if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
                        self:write( "lns_fromMapToListSub" )
                     elseif _switchExp == Ast.TypeInfoKind.Set then
                        self:write( "lns_fromMapToSetSub" )
                     elseif _switchExp == Ast.TypeInfoKind.Map then
                        self:write( "lns_fromMapToMapSub" )
                     else 
                        
                           do
                              local memberClassScope = genType:get_scope()
                              if memberClassScope ~= nil then
                                 do
                                    local symbol = memberClassScope:getSymbolInfoField( "_fromMap", true, memberClassScope, Ast.ScopeAccess.Normal )
                                    if symbol ~= nil then
                                       self:write( string.format( "%sSub", self.moduleCtrl:getMethodCName( symbol:get_typeInfo() )) )
                                    end
                                 end
                                 
                              end
                           end
                           
                     end
                  end
                  
            end
         end
         
         if #genType:get_itemTypeInfoList() > 0 then
            self:write( ", { " )
            for subIndex, subGenType in pairs( genType:get_itemTypeInfoList() ) do
               if subIndex > 1 then
                  self:write( ", " )
               end
               
               self:write( string.format( "&info_%d_%d_%s", depth + 1, subIndex, name) )
            end
            
            self:write( "}" )
         end
         
         self:writeln( "};" )
      end
      
      local function process( typeInfo, name, depth, index )
      
         for genIndex, genType in pairs( typeInfo:get_itemTypeInfoList() ) do
            if #genType:get_itemTypeInfoList() > 0 then
               for __index, subGenType in pairs( genType:get_itemTypeInfoList() ) do
                  process( subGenType, name, depth + 2, genIndex )
               end
               
            end
            
            processGenType( genType, name, depth + 1, genIndex )
         end
         
         processGenType( typeInfo, name, depth, index )
      end
      
      process( memberSym:get_typeInfo(), string.format( "%s_%s", className, memberSym:get_name()), 0, 1 )
   end
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         processDeclToMap( true )
         self:writeln( ";" )
         processDeclToMap( false )
         self:writeln( ";" )
         
         for __index, varName in pairs( Ast.getAllNameForKind( classType, Ast.MethodKind.Object, Ast.SymbolKind.Mbr ):get_list() ) do
            local memberSym = _lune.unwrap( classScope:getSymbolInfoField( varName, true, classScope, Ast.ScopeAccess.Full ))
            processFromMapInfo( memberSym )
         end
         
         
         processDeclFromMap( false )
         self:writeln( ";" )
         processDeclFromMap( true )
         self:writeln( ";" )
      elseif _switchExp == ProcessMode.DefClass then
         processToMapBody(  )
         processFromMapBody(  )
      end
   end
   
end


function convFilter:processDeclClassNodePrototype( node )

   local className = self.moduleCtrl:getClassCName( node:get_expType() )
   
   local kind = node:get_expType():get_kind()
   
   local function process( out2HMode )
   
      
      self:writeln( string.format( "typedef struct lns_mtd_%s_t {", className) )
      self:pushIndent(  )
      
      if kind == Ast.TypeInfoKind.Class then
         self:writeln( "lns_del_t * _del;" )
         self:writeln( "lns_gc_t * _gc;" )
      end
      
      
      processDeclMethodTable( self.stream, node:get_expType() )
      
      self:popIndent(  )
      self:writeln( string.format( "} lns_mtd_%s_t;", className) )
      
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
      
      self:writeln( "lns_type_meta_t * pMeta;" )
      do
         local _switchExp = kind
         if _switchExp == Ast.TypeInfoKind.Class then
            self:writeln( string.format( "u_if_imp_%s_t * pImp;", className) )
            self:writeln( string.format( "lns_mtd_%s_t * pMtd;", className) )
            processDeclMemberTable( not self.outputBuiltinFlag and self.canConv, self.stream, node:get_expType() )
            self:writeln( "// interface implements" )
            self:writeln( string.format( "u_if_imp_%s_t imp;", className) )
         elseif _switchExp == Ast.TypeInfoKind.IF then
            self:writeln( string.format( "%s pObj;", cTypeAnyP) )
            self:writeln( string.format( "lns_mtd_%s_t * pMtd;", className) )
         end
      end
      
      
      self:popIndent(  )
      self:writeln( string.format( "} %s;", className) )
      
      do
         local _switchExp = kind
         if _switchExp == Ast.TypeInfoKind.Class then
            self:writeln( string.format( [==[#define lns_mtd_%s( OBJ )                     \
                (((%s*)OBJ->val.classVal)->pMtd )]==], className, className) )
            self:writeln( string.format( "#define lns_obj_%s( OBJ ) ((%s*)OBJ->val.classVal)", className, className) )
            self:writeln( string.format( "#define lns_if_%s( OBJ ) ((%s*)OBJ->val.classVal)->pImp", className, className) )
            
            if not node:get_expType():get_abstractFlag() then
               processNewConstrProto( self.stream, self.moduleCtrl, node, out2HMode, self.outputBuiltinFlag )
               self.stream:writeln( ";" )
            end
            
            
            self:processMapping( node, node:get_expType(), out2HMode )
         elseif _switchExp == Ast.TypeInfoKind.IF then
            self:writeln( string.format( [==[#define lns_mtd_%s( OBJ )                     \
             ((%s*)&OBJ->val.ifVal)->pMtd]==], className, className) )
            if out2HMode == Out2HMode.HeaderPub then
               self:writeln( string.format( 'extern lns_type_meta_t %s;', self.moduleCtrl:getClassMetaName( node:get_expType() )) )
            end
            
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
      processDeclClassPrototype( not self.outputBuiltinFlag and self.canConv, self.stream, self.moduleCtrl, node )
      processAdvertise( self.stream, self.moduleCtrl, self.scopeMgr, self.processMode, node )
      
      if not self.outputBuiltinFlag then
         processDefaultCtor( self.stream, self.moduleCtrl, self.scopeMgr, node )
      end
      
   end
   
end


function convFilter:isManagedAnySymbol( symbol )

   local scope = symbol:get_scope()
   
   if scope:getNamespaceTypeInfo(  ):getModule(  ) ~= self.moduleTypeInfo then
      return false
   end
   
   
   local valKind = self.scopeMgr:getSymbolValKind( symbol )
   
   local varName = self.moduleCtrl:getSymbolName( symbol )
   if valKind == ValKind.Any and (symbol:get_kind() == Ast.SymbolKind.Var or symbol:get_kind() == Ast.SymbolKind.Arg and symbol:get_mutable() or isClassMember( symbol ) ) then
      if varName == "self" then
         return false
      end
      
      return true
   end
   
   return false
end


function convFilter:processDeclClassDef( node )

   local className = self.moduleCtrl:getClassCName( node:get_expType() )
   
   self:writeln( string.format( "static void mtd_%s__del( lns_env_t * _pEnv, %s pObj ) {", className, cTypeAnyP) )
   self:pushIndent(  )
   
   if self.outputBuiltinFlag or not self.canConv then
      self:writeln( string.format( "mtd_%s__delExt( _pEnv, pObj );", className) )
   end
   
   
   if node:get_expType():hasBase(  ) then
      self:writeln( string.format( "mtd_%s__del( _pEnv, pObj );", self.moduleCtrl:getClassCName( node:get_expType():get_baseTypeInfo() )) )
   end
   
   for __index, member in pairs( node:get_memberList() ) do
      if not member:get_staticFlag() then
         local valKind = self.scopeMgr:getSymbolValKind( member:get_symbolInfo() )
         do
            local _switchExp = valKind
            if _switchExp == ValKind.Stem then
               local typeInfo = member:get_symbolInfo():get_typeInfo()
               if typeInfo:get_nilable() and getValKind( typeInfo:get_nonnilableType() ) == ValKind.Prim then
               else
                
                  self:writeln( string.format( "lns_decre_ref_stem( _pEnv, %s );", getAccessMember( className, "pObj", member:get_name().txt )) )
               end
               
            elseif _switchExp == ValKind.Any then
               self:writeln( string.format( "lns_decre_ref( _pEnv, %s );", getAccessMember( className, "pObj", member:get_name().txt )) )
            end
         end
         
      end
      
   end
   
   self:popIndent(  )
   self:writeln( "}" )
   
   if not node:get_expType():get_abstractFlag() then
      
      processNewConstrProto( self.stream, self.moduleCtrl, node, Out2HMode.SourcePub, self.outputBuiltinFlag )
      self:writeln( string.format( "{ // %d", node:get_pos().lineNo) )
      self:pushIndent(  )
      
      self:processNewInsance( node:get_expType(), true )
      
      self:writeln( "return pAny;" )
      self:popIndent(  )
      self:writeln( "}" )
   end
   
   for __index, member in pairs( node:get_memberList() ) do
      local memberName = member:get_name().txt
      if member:get_getterMode() ~= Ast.AccessMode.None then
         local getterType = _lune.unwrap( node:get_scope():getTypeInfoField( string.format( "get_%s", memberName), true, node:get_scope(), scopeAccess ))
         if getterType:get_autoFlag() then
            if getterType:get_staticFlag() then
               processMethodDeclTxt( self.stream, self.moduleCtrl, FuncWrap.CallWrap, getterType )
               self:writeln( "{" )
               self:pushIndent(  )
               if self:isManagedAnySymbol( member:get_symbolInfo() ) then
                  self:writeln( string.format( "return *%s;", self.moduleCtrl:getClassMemberName( member:get_symbolInfo() )) )
               else
                
                  self:writeln( string.format( "return %s;", self.moduleCtrl:getClassMemberName( member:get_symbolInfo() )) )
               end
               
               self:popIndent(  )
               self:writeln( "}" )
            else
             
               processMethodDeclTxt( self.stream, self.moduleCtrl, FuncWrap.Normal, getterType )
               self:writeln( "{" )
               self:pushIndent(  )
               self:writeln( string.format( "return %s;", getAccessMember( className, "pObj", memberName )) )
               self:popIndent(  )
               self:writeln( "}" )
               
               processMethodDeclTxt( self.stream, self.moduleCtrl, FuncWrap.CallWrap, getterType )
               self:writeln( "{" )
               self:pushIndent(  )
               self:writeln( string.format( "return lns_mtd_%s( pObj )->get_%s( _pEnv, pObj );", className, memberName) )
               self:popIndent(  )
               self:writeln( "}" )
            end
            
         end
         
      end
      
      if member:get_setterMode() ~= Ast.AccessMode.None then
         local setterType = _lune.unwrap( node:get_scope():getTypeInfoField( string.format( "set_%s", memberName), true, node:get_scope(), scopeAccess ))
         
         local function process( accessMemberTxt )
         
            self:writeln( "{" )
            self:pushIndent(  )
            local valKind = self.scopeMgr:getSymbolValKind( member:get_symbolInfo() )
            do
               local _switchExp = valKind
               if _switchExp == ValKind.Stem then
                  self:writeln( string.format( 'lns_setq( _pEnv, %s, arg1 );', accessMemberTxt) )
               elseif _switchExp == ValKind.Any then
                  self:writeln( string.format( 'lns_setq_any( _pEnv, &%s, arg1 );', accessMemberTxt) )
               elseif _switchExp == ValKind.Prim then
                  self:writeln( string.format( "%s = arg1;", accessMemberTxt) )
               else 
                  
                     Util.err( string.format( "no support -- %s:%s:%d", member:get_symbolInfo():get_name(), ValKind:_getTxt( valKind)
                     , 4407) )
               end
            end
            
            self:popIndent(  )
            self:writeln( "}" )
         end
         
         if setterType:get_autoFlag() then
            if setterType:get_staticFlag() then
               processMethodDeclTxt( self.stream, self.moduleCtrl, FuncWrap.Normal, setterType )
               local txt
               
               if self:isManagedAnySymbol( member:get_symbolInfo() ) then
                  txt = string.format( "(*%s)", self.moduleCtrl:getClassMemberName( member:get_symbolInfo() ))
               else
                
                  txt = self.moduleCtrl:getClassMemberName( member:get_symbolInfo() )
               end
               
               
               process( txt )
            else
             
               processMethodDeclTxt( self.stream, self.moduleCtrl, FuncWrap.Normal, setterType )
               process( getAccessMember( className, "pObj", memberName ) )
               
               processMethodDeclTxt( self.stream, self.moduleCtrl, FuncWrap.CallWrap, setterType )
               self:writeln( "{" )
               self:pushIndent(  )
               self:writeln( string.format( "lns_mtd_%s( pObj )->set_%s( _pEnv, pObj, arg1 );", className, memberName) )
               self:popIndent(  )
               self:writeln( "}" )
            end
            
         end
         
      end
      
   end
   
   
   processAdvertise( self.stream, self.moduleCtrl, self.scopeMgr, self.processMode, node )
end


local function processInitMethodTable( stream, moduleCtrl, classTypeInfo )

   local function outputField( name, retTypeList )
   
      local methodType = getMethodTypeTxt( retTypeList )
      stream:writeln( string.format( "(%s *)%s,", methodType, name) )
   end
   
   local scope = _lune.unwrap( classTypeInfo:get_scope())
   
   local nameSet = Ast.getAllMethodName( classTypeInfo, Ast.MethodKind.Object )
   
   for __index, name in pairs( nameSet:get_list() ) do
      local symbolInfo = _lune.unwrap( scope:getSymbolInfoField( name, true, scope, Ast.ScopeAccess.Normal ))
      if not symbolInfo:get_typeInfo():get_abstractFlag() then
         outputField( moduleCtrl:getMethodCName( symbolInfo:get_typeInfo() ), symbolInfo:get_typeInfo():get_retTypeInfoList() )
      else
       
         stream:writeln( "NULL," )
      end
      
   end
   
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
                        local impMethodSym = _lune.unwrap( impScope:getSymbolInfoField( symbolInfo:get_name(), true, impScope, scopeAccess ))
                        local impMethodType = impMethodSym:get_typeInfo()
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

   if classType:hasBase(  ) then
      processIFMethodDataInit( stream, moduleCtrl, classType:get_baseTypeInfo(), orgClassType )
   end
   
   if not orgClassType:get_abstractFlag() then
      local className = moduleCtrl:getClassCName( orgClassType )
      for __index, ifType in pairs( classType:get_interfaceList() ) do
         local ifName = moduleCtrl:getClassCName( ifType )
         stream:writeln( string.format( "static lns_mtd_%s_t lns_if_%s_imp_%s = {", ifName, className, ifName) )
         stream:pushIndent(  )
         
         processInitIFMethodTable( stream, moduleCtrl, ifType, orgClassType )
         
         stream:popIndent(  )
         stream:writeln( "};" )
      end
      
   end
   
end

local function processClassMeta( stream, moduleCtrl, classTypeInfo )

   local className = moduleCtrl:getClassCName( classTypeInfo )
   stream:write( string.format( 'lns_type_meta_t %s = { "%s", &%s, {', moduleCtrl:getClassMetaName( classTypeInfo ), className, moduleCtrl:getClassMetaName( classTypeInfo:get_baseTypeInfo() )) )
   for index, ifType in pairs( classTypeInfo:get_interfaceList() ) do
      stream:write( string.format( "&%s, ", moduleCtrl:getClassMetaName( ifType )) )
   end
   
   stream:writeln( "NULL } };" )
end

local function processClassDataInit( stream, moduleCtrl, scopeMgr, classTypeInfo, fieldList )

   local className = moduleCtrl:getClassCName( classTypeInfo )
   
   if not Ast.isPubToExternal( classTypeInfo:get_accessMode() ) then
      stream:write( "static " )
   end
   
   processClassMeta( stream, moduleCtrl, classTypeInfo )
   
   if not classTypeInfo:get_abstractFlag() then
      
      if not Ast.isPubToExternal( classTypeInfo:get_accessMode() ) then
         stream:write( "static " )
      end
      
      stream:writeln( string.format( "lns_mtd_%s_t lns_mtd_%s = {", className, className) )
      stream:pushIndent(  )
      
      stream:writeln( string.format( "mtd_%s__del,", className) )
      
      if hasGC( classTypeInfo ) then
         stream:writeln( string.format( "mtd_%s__gc,", className) )
      else
       
         stream:writeln( "NULL," )
      end
      
      
      processInitMethodTable( stream, moduleCtrl, classTypeInfo )
      stream:popIndent(  )
      stream:writeln( "};" )
   end
   
   
   local function process( out2HMode, symbolInfo )
   
      do
         local _switchExp = out2HMode
         if _switchExp == Out2HMode.HeaderPub or _switchExp == Out2HMode.SourcePub or _switchExp == Out2HMode.SourcePri then
            stream:writeln( string.format( "%s%s %s;", getOut2HeaderPrefix( out2HMode ), scopeMgr:getCTypeForSym( symbolInfo ), moduleCtrl:getClassMemberName( symbolInfo )) )
         end
      end
      
   end
   
   for __index, symbolInfo in pairs( (_lune.unwrap( classTypeInfo:get_scope()) ):get_symbol2SymbolInfoMap() ) do
      if isClassMember( symbolInfo ) then
         
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



function convFilter:processDeclVarC( declFlag, var, init0, manageScope )

   if declFlag then
      local typeTxt = self.scopeMgr:getCTypeForSym( var )
      self:writeln( string.format( "%s %s;", typeTxt, self.moduleCtrl:getSymbolName( var )) )
   end
   
   
   local valKind = self.scopeMgr:getSymbolValKind( var )
   if valKind == ValKind.Prim then
      if init0 then
         self:writeln( string.format( "%s = 0;", self.moduleCtrl:getSymbolName( var )) )
      end
      
      return 
   end
   
   local scope = manageScope
   if  nil == scope then
      local _scope = scope
   
      scope = var:get_scope()
   end
   
   
   local initVal
   
   if not init0 or isStemType( var:get_typeInfo() ) then
      do
         local symbolInfo = _lune.__Cast( var, 3, Ast.SymbolInfo )
         if symbolInfo ~= nil then
            do
               local _switchExp = valKind
               if _switchExp == ValKind.Any then
                  self:write( "lns_set_block_any" )
                  self:writeln( string.format( "( %s, %d, %s );", getBlockName( scope ), getSymbolIndex( symbolInfo ), self.moduleCtrl:getSymbolName( var )) )
               elseif _switchExp == ValKind.Stem then
                  self:write( "lns_set_block_stem" )
                  self:writeln( string.format( "( %s, %d, %s );", getBlockName( scope ), getSymbolIndex( symbolInfo ), self.moduleCtrl:getSymbolName( var )) )
               elseif _switchExp == ValKind.Var then
                  local typeTxt = getStemTypeId( getOrgTypeInfo( var:get_typeInfo() ) )
                  self:writeln( string.format( "lns_set_block_var( %s, %d, %s, %s );", getBlockName( scope ), getSymbolIndex( symbolInfo ), typeTxt, self.moduleCtrl:getSymbolName( var )) )
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
                  self:write( "lns_initVal_stem" )
               elseif _switchExp == ValKind.Any then
                  self:write( "lns_initVal_any" )
               elseif _switchExp == ValKind.Var then
                  self:write( "lns_initVal_var" )
               else 
                  
                     Util.err( string.format( "not support -- %s", ValKind:_getTxt( valKind)
                     ) )
               end
            end
            
            
            self:writeln( string.format( "( %s, %s, %d, %s );", self.moduleCtrl:getSymbolName( var ), getBlockName( scope ), getSymbolIndex( symbolInfo ), initVal) )
         else
            self:writeln( string.format( "%s = %s;", self.moduleCtrl:getSymbolName( var ), initVal) )
         end
      end
      
   end
   
end


function convFilter:process__func__symbol( funcTypeInfo, has__func__Symbol, funcName )

   if not has__func__Symbol then
      return 
   end
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.WideScopeVer then
         local scope = _lune.unwrap( funcTypeInfo:get_scope())
         local symbol = _lune.unwrap( scope:getSymbolInfoChild( "__func__" ))
         self:writeln( string.format( "static %s %s = NULL;", cTypeAnyPP, self.moduleCtrl:getSymbolName( symbol )) )
      elseif _switchExp == ProcessMode.InitFuncSym then
         local scope = _lune.unwrap( funcTypeInfo:get_scope())
         local symbol = _lune.unwrap( scope:getSymbolInfoChild( "__func__" ))
         local symbolParam = self.scopeMgr:getSymbolParam( symbol )
         local symbolName = self.moduleCtrl:getSymbolName( symbol )
         self:writeln( string.format( "lns_set_block_any( pBlock, %d, %s );", symbolParam.index, symbolName) )
         self:writeln( string.format( 'lns_setQ_any( %s, lns_litStr2any( _pEnv, "%s") );', symbolName, funcName) )
      end
   end
   
end


function convFilter:processArgClosure( declInfo )

   for __index, argNode in pairs( declInfo:get_argList() ) do
      do
         local declArg = _lune.__Cast( argNode, 3, Nodes.DeclArgNode )
         if declArg ~= nil then
            local symbolInfo = declArg:get_symbolInfo()
            if symbolInfo:get_hasAccessFromClosure() then
               
               local symbolParam = self.scopeMgr:getSymbolParam( symbolInfo )
               self:writeln( string.format( "%s %s;", symbolParam.typeTxt, self.moduleCtrl:getSymbolName( symbolInfo )) )
               
               self:write( "lns_initVal_var(" )
               self:write( string.format( " %s, %s, %d, ", self.moduleCtrl:getSymbolName( symbolInfo ), getBlockName( symbolInfo:get_scope() ), symbolParam.index) )
               
               local valKind = getValKind( symbolInfo:get_typeInfo() )
               local workSymbol = WorkSymbol.new(symbolInfo:get_scope(), symbolInfo:get_accessMode(), string.format( "_%s", symbolInfo:get_name()), symbolInfo:get_typeInfo(), symbolInfo:get_kind(), symbolInfo:get_staticFlag(), SymbolParam.new(valKind, 0, getCType( symbolInfo:get_typeInfo() )))
               
               self:processSym2stem( workSymbol )
               self:writeln( ");" )
            elseif symbolInfo:get_mutable() then
               
               local symbolParam = self.scopeMgr:getSymbolParam( symbolInfo )
               self:processDeclVarC( true, symbolInfo, false )
               do
                  local _switchExp = getValKind( symbolInfo:get_typeInfo() )
                  if _switchExp == ValKind.Stem then
                     self:writeln( string.format( "lns_setQ( %s, _%s );", symbolInfo:get_name(), symbolInfo:get_name()) )
                  elseif _switchExp == ValKind.Any then
                     self:writeln( string.format( "lns_setQ_any( %s, _%s );", symbolInfo:get_name(), symbolInfo:get_name()) )
                  elseif _switchExp == ValKind.Prim then
                     self:writeln( string.format( "%s = _%s;", symbolInfo:get_name(), symbolInfo:get_name()) )
                  end
               end
               
            end
            
         end
      end
      
   end
   
end


function convFilter:processDeclMethodInfo( declInfo, funcTypeInfo, parent )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         processPrototypeMethod( self.stream, self.moduleCtrl, declInfo:get_argList(), funcTypeInfo )
      elseif _switchExp == ProcessMode.Form then
         local classType = funcTypeInfo:get_parentInfo()
         local className = self.moduleCtrl:getClassCName( classType )
         do
            local body = declInfo:get_body()
            if body ~= nil then
               local methodNodeToken = _lune.unwrap( declInfo:get_name(  ))
               processMethodDeclTxt( self.stream, self.moduleCtrl, FuncWrap.Normal, funcTypeInfo, declInfo:get_argList() )
               self:writeln( "{" )
               
               self:pushIndent(  )
               
               self:pushRoutine( funcTypeInfo, body )
               
               do
                  local declClassNode = declInfo:get_declClassNode()
                  if declClassNode ~= nil then
                     if _lune.nilacc( declInfo:get_name(), "txt" ) == "__init" then
                        
                        for __index, memberSym in pairs( declClassNode:get_uninitMemberList() ) do
                           if declInfo:get_staticFlag() == memberSym:get_staticFlag() then
                              local memberAccess
                              
                              
                              if declInfo:get_staticFlag() then
                                 memberAccess = self.moduleCtrl:getClassMemberName( memberSym )
                              else
                               
                                 memberAccess = getAccessMember( className, "pObj", memberSym:get_name() )
                              end
                              
                              self:writeln( string.format( "%s = %s;", memberAccess, cValNil) )
                           end
                           
                        end
                        
                     end
                     
                  end
               end
               
               
               if funcTypeInfo:get_staticFlag() and funcTypeInfo:get_rawTxt() == "___init" then
                  for __index, symbol in pairs( (_lune.unwrap( classType:get_scope()) ):get_symbol2SymbolInfoMap() ) do
                     if isClassMember( symbol ) then
                        self:processDeclVarC( false, symbol, false, self.moduleTypeInfo:get_scope() )
                     end
                     
                  end
                  
               end
               
               
               local scope = body:get_scope()
               do
                  local selfSymbol = scope:getSymbolInfoChild( "self" )
                  if selfSymbol ~= nil then
                     local symbolParam = self.scopeMgr:getSymbolParam( selfSymbol )
                     if symbolParam.kind == ValKind.Var then
                        self:writeln( string.format( "%s self;", cTypeVarP) )
                        self:writeln( string.format( "lns_initVal_var( self, %s, %d, LNS_STEM_ANY( pObj ) );", getBlockName( scope ), symbolParam.index) )
                     else
                      
                        if not funcTypeInfo:get_staticFlag() then
                           self:writeln( string.format( "%s self = pObj;", cTypeAnyP) )
                        end
                        
                     end
                     
                  end
               end
               
               
               self:processArgClosure( declInfo )
               
               self.duringDeclFunc = true
               filter( body, self, parent )
               self.duringDeclFunc = false
               
               self:popRoutine(  )
               
               self:popIndent(  )
               self:writeln( "}" )
            end
         end
         
         processDeclCallMethodWrapper( self.stream, self.moduleCtrl, self.scopeMgr, parent, funcTypeInfo, true )
         processDeclCallMethodWrapper( self.stream, self.moduleCtrl, self.scopeMgr, parent, funcTypeInfo, false )
      elseif _switchExp == ProcessMode.InitFuncSym or _switchExp == ProcessMode.WideScopeVer then
         self:process__func__symbol( funcTypeInfo, declInfo:get_has__func__Symbol(), self.moduleCtrl:getMethodCName( funcTypeInfo ) )
      end
   end
   
end


function convFilter:processDeclConstr( node, opt )

   self:processDeclMethodInfo( node:get_declInfo(), node:get_expType(), node )
end



function convFilter:processDeclDestr( node, opt )

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


function convFilter:processVal2stem( node, parent )

   process2stem( self.stream, self.moduleCtrl, self.scopeMgr, self:getValKindOfNode( node ), node:get_expType(), parent, function (  )
   
      filter( node, self, parent )
   end )
end


function convFilter:processCallArgList( funcType, expListNode )

   local funcArgTypeList = funcType:get_argTypeInfoList()
   
   local abbrValTxt
   
   if self.moduleCtrl:getBuiltinFuncNameFromType( funcType ) or (funcType:get_kind() == Ast.TypeInfoKind.Method and funcType:get_parentInfo():equals( Ast.builtinTypeString ) ) then
      
      abbrValTxt = cValNone
   else
    
      abbrValTxt = cValNil
   end
   
   
   local function processAbbr( funcArgType )
   
      if funcArgType:get_kind() == Ast.TypeInfoKind.DDD then
         self:write( cValDDD0 )
      else
       
         self:write( abbrValTxt )
      end
      
   end
   
   if expListNode ~= nil then
      local expList = expListNode:get_expList()
      for index, funcArgType in pairs( funcArgTypeList ) do
         self:write( ", " )
         if #expList >= index then
            local expNode = expList[index]
            if expNode:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
               processAbbr( funcArgType )
            else
             
               if funcArgType:get_kind() == Ast.TypeInfoKind.DDD then
                  if expNode:get_kind() == Nodes.NodeKind.get_Abbr() then
                     self:write( cValDDD0 )
                  else
                   
                     filter( expNode, self, expListNode )
                  end
                  
                  
                  return 
               else
                
                  
                  if isStemType( funcArgType ) then
                     self:processVal2stem( expNode, expListNode )
                  else
                   
                     filter( expNode, self, expListNode )
                  end
                  
               end
               
            end
            
         else
          
            processAbbr( funcArgType )
         end
         
      end
      
   else
      for __index, funcArgType in pairs( funcArgTypeList ) do
         self:write( ", " )
         processAbbr( funcArgType )
      end
      
   end
   
end


function convFilter:processExpCallSuper( node, opt )

   local funcType
   
   if node:get_methodType():get_rawTxt() == "__init" then
      self:write( string.format( "%s( _pEnv, pObj", self.moduleCtrl:getCtorName( node:get_superType() )) )
      local superScope = _lune.unwrap( node:get_superType():get_scope())
      funcType = _lune.unwrap( superScope:getTypeInfoChild( "__init" ))
   else
    
      self:write( string.format( "%s( _pEnv, pObj", self.moduleCtrl:getMethodCName( node:get_methodType() )) )
      funcType = node:get_methodType()
   end
   
   
   self:processCallArgList( funcType, node:get_expList() )
   self:writeln( ");" )
end



function convFilter:processDeclMethod( node, opt )

   self:processDeclMethodInfo( node:get_declInfo(), node:get_expType(), node )
end



function convFilter:processProtoMethod( node, opt )

   if node:get_expType():get_abstractFlag() then
      self:processDeclMethodInfo( node:get_declInfo(), node:get_expType(), node )
   end
   
end



function convFilter:processUnwrapSet( node, opt )

end


function convFilter:accessPrimValFromAny( dddFlag, typeInfo, index )

   self:write( getAccessPrimValFromStem( dddFlag, typeInfo, index ) )
end


function convFilter:isStemSym( symbolInfo )

   return self.scopeMgr:getSymbolValKind( symbolInfo ) ~= ValKind.Prim
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
      if _switchExp == ValKind.Var then
         filter( exp, self, parent )
      elseif _switchExp == ValKind.Prim then
         filter( exp, self, parent )
      elseif _switchExp == ValKind.Stem then
         filter( exp, self, parent )
         self:accessPrimValFromAny( #exp:get_expTypeList() > 1, exp:get_expType(), 0 )
      elseif _switchExp == ValKind.Any then
         filter( exp, self, parent )
      else 
         
            Util.err( string.format( "not support -- %d", 5284) )
      end
   end
   
end


function convFilter:processSym2Any( symbol )

   local valKind = self.scopeMgr:getSymbolValKind( symbol )
   local symName
   
   if self:isManagedAnySymbol( symbol ) then
      symName = string.format( "(*%s)", self.moduleCtrl:getSymbolName( symbol ))
   else
    
      symName = self.moduleCtrl:getSymbolName( symbol )
   end
   
   do
      local _switchExp = valKind
      if _switchExp == ValKind.Stem then
         self:write( symName )
         self:write( accessAny )
      elseif _switchExp == ValKind.Any then
         self:write( symName )
      elseif _switchExp == ValKind.Var then
         self:write( symName )
         self:write( string.format( "->stem%s", accessAny) )
      else 
         
            Util.err( string.format( "not suppport -- %s, %d", ValKind:_getTxt( valKind)
            , 5341) )
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
         
            Util.err( string.format( "not suppport -- %d, %s, %s, %d", node:get_pos().lineNo, ValKind:_getTxt( valKind)
            , Nodes.getNodeKindName( node:get_kind() ), 5366) )
      end
   end
   
end




function convFilter:processSetValSingleDirect( parent, node, var, initFlag, expValKind, expValType, index, firstMRet, processVal )

   local valKind = self.scopeMgr:getSymbolValKind( var )
   
   local varName = self.moduleCtrl:getSymbolName( var )
   if self:isManagedAnySymbol( var ) then
      varName = string.format( "(*%s)", varName)
   end
   
   local processPrefix = nil
   do
      local fieldNode = _lune.__Cast( node, 3, Nodes.RefFieldNode )
      if fieldNode ~= nil then
         if _lune.nilacc( fieldNode:get_symbolInfo(), 'get_staticFlag', 'callmtd' ) then
         else
          
            processPrefix = function (  )
            
               local prefixNode = fieldNode:get_prefix()
               local className = self.moduleCtrl:getClassCName( prefixNode:get_expType() )
               self:write( string.format( "lns_obj_%s( ", className) )
               self:processVal2any( prefixNode, fieldNode )
               
               self:write( ")->" )
            end
         end
         
      end
   end
   
   
   if var:get_symbolId() == invalidSymbolId then
      if valKind == expValKind then
         self:write( string.format( "%s = ", varName) )
         processVal(  )
         self:write( ";" )
         return 
      end
      
      Util.err( string.format( "illegal %s %s %s -- %d", var:get_name(), ValKind:_getTxt( valKind)
      , ValKind:_getTxt( expValKind)
      , 5425) )
   end
   
   
   do
      local _switchExp = valKind
      if _switchExp == ValKind.Var then
         
         if initFlag then
            self:write( string.format( "lns_setQ( %s->stem, ", varName) )
         else
          
            self:write( string.format( "lns_setq( _pEnv, %s->stem, ", varName) )
         end
         
         
         process2stem( self.stream, self.moduleCtrl, self.scopeMgr, expValKind, expValType, parent, processVal )
         self:write( " );" )
      elseif _switchExp == ValKind.Stem then
         do
            local _switchExp = expValKind
            if _switchExp == ValKind.Stem or _switchExp == ValKind.Any or _switchExp == ValKind.Prim then
               if initFlag then
                  self:write( "lns_setQ( " )
               else
                
                  self:write( "lns_setq( _pEnv, " )
               end
               
               if processPrefix ~= nil then
                  processPrefix(  )
               end
               
               self:write( string.format( "%s, ", varName) )
               processVal(  )
               self:write( " );" )
            end
         end
         
      elseif _switchExp == ValKind.Any then
         if initFlag then
            self:write( "lns_setQ_any( &" )
         else
          
            self:write( "lns_setq_any( _pEnv, &" )
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
   
   local symName
   
   if self:isManagedAnySymbol( symbol ) then
      symName = string.format( "(*%s)", self.moduleCtrl:getSymbolName( symbol ))
   else
    
      symName = self.moduleCtrl:getSymbolName( symbol )
   end
   
   self:write( symName )
end


local function processToIF( stream, moduleCtrl, expType, process )

   if expType:get_kind() == Ast.TypeInfoKind.IF then
      stream:write( "lns_toIF( _pEnv, " )
      process(  )
      stream:write( string.format( ", &lns_type_meta_%s )", moduleCtrl:getClassCName( expType )) )
      stream:write( accessAny )
   else
    
      process(  )
   end
   
end

local function processGetMRet( stream, moduleCtrl, typeInfo, index )

   local function process(  )
   
      stream:write( string.format( "lns_getMRet( _pEnv, %d )", index) )
      stream:write( getAccessValFromStem( typeInfo ) )
   end
   if typeInfo:get_kind() == Ast.TypeInfoKind.IF then
      processToIF( stream, moduleCtrl, typeInfo, process )
   else
    
      process(  )
   end
   
end

local function needsWrapper( orgFunc, castType )

   if #orgFunc:get_argTypeInfoList() == #castType:get_argTypeInfoList() and #orgFunc:get_retTypeInfoList() == #castType:get_retTypeInfoList() then
      local function check( typeList1, typeList2 )
      
         for index, type1 in pairs( typeList1 ) do
            if not type1:equals( typeList2[index] ) then
               return false
            end
            
         end
         
         return true
      end
      if check( orgFunc:get_argTypeInfoList(), castType:get_argTypeInfoList() ) and check( orgFunc:get_retTypeInfoList(), castType:get_retTypeInfoList() ) then
         
         return false
      end
      
   end
   
   return true
end

function convFilter:processFuncCast2Form( castType, orgFunc )

   if not needsWrapper( orgFunc, castType ) then
      self:write( getFunc2any( self.moduleCtrl, self.scopeMgr, orgFunc ) )
      return 
   end
   
   
   local closureSymList = _lune.unwrapDefault( _lune.nilacc( orgFunc:get_scope(), 'get_closureSymList', 'callmtd' ), {})
   
   local argList = castType:get_argTypeInfoList()
   local hasDDD = #argList > 0 and argList[#argList]:get_kind() == Ast.TypeInfoKind.DDD or false
   
   self:write( getPrepareClosure( self.scopeMgr, self.moduleCtrl:getFuncCastWrapName( orgFunc, castType ), #argList, hasDDD, closureSymList ) )
end


function convFilter:processValForSetOp( parent, dstKind, dstTypeInfo, exp, index, firstMRet )

   local valKind = self:getValKindOfNode( exp )
   
   local function accessVal(  )
   
      if firstMRet then
         processGetMRet( self.stream, self.moduleCtrl, exp:get_expType(), 0 )
      else
       
         if exp:get_expType():get_kind() == Ast.TypeInfoKind.Func then
            self:processFuncCast2Form( dstTypeInfo, exp:get_expType() )
         else
          
            filter( exp, self, parent )
         end
         
      end
      
   end
   
   local function processVal(  )
   
      
      local setValTxt = ""
      if firstMRet then
         accessVal(  )
      elseif not firstMRet and Nodes.hasMultiValNode( exp ) then
         self:write( "lns_fromDDD( " )
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
   
   
   self:processSetValSingleDirect( parent, node, var, initFlag, expValKind, exp:get_expType(), index, firstMRet, function (  )
   
      self:processValForSetOp( parent, self.scopeMgr:getSymbolValKind( var ), var:get_typeInfo(), exp, index, firstMRet )
   end )
end


function convFilter:processSetSymSingle( parent, node, var, initFlag, symbol, toIF )

   local function process(  )
   
      self:processSymForSetOp( parent, self.scopeMgr:getSymbolValKind( var ), var:get_typeInfo(), symbol )
   end
   
   self:processSetValSingleDirect( parent, node, var, initFlag, self.scopeMgr:getSymbolValKind( symbol ), symbol:get_typeInfo(), 1, false, function (  )
   
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
               do
                  local _switchExp = dstType:get_kind()
                  if _switchExp == Ast.TypeInfoKind.Map then
                     self:write( "lns_mtd_Map_add( _pEnv, " )
                     self:processVal2any( refItemNode:get_val(), var )
                     self:write( ", " )
                     do
                        local indexNode = refItemNode:get_index()
                        if indexNode ~= nil then
                           self:processVal2stem( indexNode, var )
                        else
                           self:write( getLiteralStrStem( string.format( '"%s"', _lune.unwrap( refItemNode:get_symbol())) ) )
                        end
                     end
                     
                     self:write( ", " )
                     self:processValForSetOp( parent, ValKind.Stem, dstType:get_itemTypeInfoList()[2], exp, index, firstMRet )
                     self:write( ")" )
                  elseif _switchExp == Ast.TypeInfoKind.List then
                     self:write( "void lns_mtd_List_setAt( _pEnv, " )
                     self:processVal2any( refItemNode:get_val(), var )
                     self:write( ", " )
                     do
                        local indexNode = refItemNode:get_index()
                        if indexNode ~= nil then
                           self:processVal2stem( indexNode, var )
                        else
                           self:write( getLiteralStrStem( string.format( '"%s"', _lune.unwrap( refItemNode:get_symbol())) ) )
                        end
                     end
                     
                     self:write( ", " )
                     self:processValForSetOp( parent, ValKind.Stem, dstType:get_itemTypeInfoList()[2], exp, index, firstMRet )
                     self:write( ")" )
                  else 
                     
                        Util.err( string.format( "not support type -- %s", Ast.TypeInfoKind:_getTxt( dstType:get_kind())
                        ) )
                  end
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

DstInfo.Node = { "Node", {{ func=Nodes.Node._fromMap, nilable=false, child={} },{ func=_lune._toBool, nilable=false, child={} }}}
DstInfo._name2Val["Node"] = DstInfo.Node
DstInfo.Symbol = { "Symbol", {{ func=Ast.LowSymbol._fromMap, nilable=false, child={} },{ func=Nodes.Node._fromMap, nilable=true, child={} },{ func=_lune._toBool, nilable=false, child={} }}}
DstInfo._name2Val["Symbol"] = DstInfo.Symbol


function convFilter:processSetValToDst( parent, dstList, expList, mRetExp )

   local mRetIndex = _lune.nilacc( mRetExp, 'get_index', 'callmtd' )
   
   for index, exp in pairs( expList ) do
      local is1stMRet = index == mRetIndex
      if is1stMRet then
         if mRetExp ~= nil then
            self:write( "lns_setMRet( _pEnv, " )
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
                  local initFlag = _matchExp[2][3]
               
                  self:processSetValSingle( parent, dstNode, symbolInfo, initFlag, exp, accessIndex, is1stMRet and dstIndex == index )
               elseif _matchExp[1] == DstInfo.Node[1] then
                  local dstNode = _matchExp[2][1]
                  local initFlag = _matchExp[2][2]
               
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
               local initFlag = _matchExp[2][3]
            
               self:processSetValSingle( parent, dstNode, symbolInfo, initFlag, exp, accessIndex, is1stMRet )
            elseif _matchExp[1] == DstInfo.Node[1] then
               local dstNode = _matchExp[2][1]
               local initFlag = _matchExp[2][2]
            
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
      
      table.insert( dstList, _lune.newAlge( DstInfo.Symbol, {symbol,node,initFlag}) )
   end
   
   self:processSetValToDst( parent, dstList, expList, mRetExp )
end


function convFilter:processSetValToNode( parent, dstNode, initSymSet, expList, mRetExp )

   local function isInitSym( node )
   
      local symbolList = node:getSymbolInfo(  )
      if #symbolList > 0 then
         return _lune._Set_has(initSymSet, symbolList[1] )
      end
      
      return false
   end
   
   local dstList = {}
   do
      local expListNode = _lune.__Cast( dstNode, 3, Nodes.ExpListNode )
      if expListNode ~= nil then
         for __index, node in pairs( expListNode:get_expList() ) do
            table.insert( dstList, _lune.newAlge( DstInfo.Node, {node,isInitSym( node )}) )
         end
         
      else
         table.insert( dstList, _lune.newAlge( DstInfo.Node, {dstNode,isInitSym( dstNode )}) )
      end
   end
   
   
   self:processSetValToDst( parent, dstList, expList, mRetExp )
end


function convFilter:processDeclVarAndSet( varSymList, expListNode )

   for index, var in pairs( varSymList ) do
      local symbolParam = self.scopeMgr:getSymbolParam( var )
      local typeTxt, valKind = symbolParam.typeTxt, symbolParam.kind
      local declVarFlag
      
      if varSymList[1]:get_scope() ~= self.ast:get_moduleScope() or symbolParam.index == invalidSymbolId then
         declVarFlag = true
      else
       
         declVarFlag = false
      end
      
      
      if valKind ~= ValKind.Prim then
         self:processDeclVarC( declVarFlag, var, valKind ~= ValKind.Var )
      else
       
         if declVarFlag then
            self:writeln( string.format( "%s %s;", typeTxt, self.moduleCtrl:getSymbolName( var )) )
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
      local workSymbol = WorkSymbol.new(varSym:get_scope(), varSym:get_accessMode(), string.format( "_%s", varSym:get_name()), varSym:get_typeInfo():get_nilableTypeInfo(), varSym:get_kind(), varSym:get_staticFlag(), SymbolParam.new(ValKind.Stem, 0, cTypeStem))
      table.insert( workSymList, workSymbol )
   end
   
   self:processDeclVarAndSet( workSymList, expListNode )
   
   self:write( "if ( " )
   for index, workSym in pairs( workSymList ) do
      self:write( string.format( "%s.type != lns_stem_type_nil", self.moduleCtrl:getSymbolName( workSym )) )
      if index ~= #workSymList then
         self:write( " && " )
      end
      
   end
   
   self:writeln( ") {" )
   
   self:processBlockPreProcess( node:get_block():get_scope() )
   
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
         if node:get_mode() == Nodes.DeclVarMode.Let then
            local varSymList = node:get_symbolInfoList()
            if varSymList[1]:get_scope() == self.moduleTypeInfo:get_scope() or self.outputBuiltinFlag then
               local function process( out2HMode, var )
               
                  local typeTxt = self.scopeMgr:getCTypeForSym( var )
                  self:write( getOut2HeaderPrefix( out2HMode ) )
                  self:writeln( string.format( "%s %s;", typeTxt, self.moduleCtrl:getSymbolName( var )) )
               end
               for index, var in pairs( varSymList ) do
                  
                  do
                     local function processwork( out2HMode )
                     
                        process( out2HMode, var )
                        
                        
                     end
                     if Ast.isPubToExternal( var:get_accessMode() ) then
                        self.stream:switchToHeader(  )
                        processwork( Out2HMode.HeaderPub )
                        self.stream:returnToSource(  )
                        
                        processwork( Out2HMode.SourcePub )
                     else
                      
                        processwork( Out2HMode.SourcePri )
                     end
                     
                  end
                  
                  
               end
               
            end
            
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
   
   if node:get_unwrapFlag() then
      
      local unwrapScope = _lune.unwrap( _lune.nilacc( node:get_unwrapBlock(), 'get_scope', 'callmtd' ))
      
      local workSymList = {}
      for index, varSym in pairs( varSymList ) do
         local tmpVarSym = _lune.unwrap( unwrapScope:getSymbolInfoChild( string.format( "_%s", varSym:get_name()) ))
         
         table.insert( workSymList, tmpVarSym )
         if node:get_mode() == Nodes.DeclVarMode.Let and varSym:get_scope() ~= self.moduleTypeInfo:get_scope() then
            self:processDeclVarC( true, varSym, false )
         end
         
      end
      
      
      self:writeln( "{" )
      self:processBlockPreProcess( unwrapScope )
      
      self:processDeclVarAndSet( workSymList, node:get_expList() )
      
      do
         local unwrapBlock = node:get_unwrapBlock()
         if unwrapBlock ~= nil then
            self:writeln( "" )
            self:write( "if ( " )
            local firstFlag = true
            for __index, var in pairs( workSymList ) do
               if self.scopeMgr:getSymbolValKind( var ) == ValKind.Stem then
                  if var:get_typeInfo():get_nilable() then
                     if firstFlag then
                        firstFlag = false
                     else
                      
                        self:write( " || " )
                     end
                     
                     self:write( string.format( "lns_stem_type_nil == %s.type", self.moduleCtrl:getSymbolName( var )) )
                  end
                  
               end
               
            end
            
            self:writeln( " ) {" )
            
            self:pushIndent(  )
            filter( unwrapBlock, self, node )
            self:popIndent(  )
            
            do
               local thenBlock = node:get_thenBlock()
               if thenBlock ~= nil then
                  
                  self:writeln( "}" )
                  self:writeln( "else {" )
                  
                  self:pushIndent(  )
                  
                  for index, var in pairs( varSymList ) do
                     self:processSetSymSingle( node, nil, var, true, workSymList[index], false )
                     self:writeln( "" )
                  end
                  
                  self:popIndent(  )
                  filter( thenBlock, self, node )
                  self:writeln( "}" )
               else
                  self:writeln( "}" )
                  self:writeln( "else {" )
                  self:pushIndent(  )
                  for index, var in pairs( varSymList ) do
                     self:processSetSymSingle( node, nil, var, true, workSymList[index], false )
                     self:writeln( "" )
                  end
                  
                  self:popIndent(  )
                  self:writeln( "}" )
               end
            end
            
            self:processBlockPostProcess(  )
            self:writeln( "}" )
         end
      end
      
   else
    
      self:processDeclVarAndSet( varSymList, node:get_expList() )
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
         local name = self.moduleCtrl:getSymbolName( var )
         self.pubVarName2InfoMap[name] = PubVarInfo.new(node:get_staticFlag(), node:get_accessMode(), node:get_symbolInfoList()[index]:get_mutable(), node:get_typeInfoList()[index])
      end
      
   end
   
end



function convFilter:processWhen( node, opt )

   self:write( "if ( " )
   for index, symPair in pairs( node:get_symPairList() ) do
      self:processSym2stem( symPair:get_src() )
      self:write( ".type != lns_stem_type_nil" )
      if index ~= #node:get_symPairList() then
         self:write( " && " )
      end
      
   end
   
   self:writeln( " ) " )
   self:writeln( "{" )
   
   self:processBlockPreProcess( node:get_block():get_scope() )
   
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
       
         self:writeln( string.format( "%s %s = %s;", dstTypeTxt, self.moduleCtrl:getSymbolName( dstSymbol ), self.moduleCtrl:getSymbolName( srcSymbol )) )
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

   processDeclAlgeSub( self.stream, node )
end



function convFilter:processDeclArgDDD( node, opt )

   self:write( string.format( "%s _pDDD", cTypeStem) )
end



function convFilter:processExpDDD( node, opt )

end



function convFilter:processExpSubDDD( node, opt )

   self:write( string.format( "lns_createSubDDD( _pEnv, %d, LNS_STEM_ANY( _pEnv->pMRet ) )", node:get_remainIndex()) )
end



function convFilter:processFuncPrototype( parent, accessMode, needFormVal, name, retType, argList, termFlag )

   local function process( out2HMode )
   
      self:write( string.format( "%s%s %s( %s _pEnv", getOut2HeaderPrefix( out2HMode ), retType, name, cTypeEnvP ) )
      if needFormVal then
         self:write( string.format( ", %s _pForm", cTypeAnyP ) )
      end
      
      
      for index, arg in pairs( argList ) do
         self:write( ", " )
         filter( arg, self, parent )
      end
      
      self:write( " )" )
      if termFlag then
         self:writeln( ";" )
      end
      
   end
   
   if termFlag then
      
      do
         local function processwork( out2HMode )
         
            if out2HMode ~= Out2HMode.SourcePub then
               process( out2HMode )
            end
            
            
            
         end
         if Ast.isPubToExternal( accessMode ) then
            self.stream:switchToHeader(  )
            processwork( Out2HMode.HeaderPub )
            self.stream:returnToSource(  )
            
            processwork( Out2HMode.SourcePub )
         else
          
            processwork( Out2HMode.SourcePri )
         end
         
      end
      
      
   else
    
      process( Out2HMode.SourcePub )
   end
   
end


function convFilter:processCallUserForm( formName, formType, argNameList )

   local function process( prefix )
   
      self:pushIndent(  )
      self:write( prefix )
      for index, argName in pairs( argNameList ) do
         self:write( ", " )
         self:write( argName )
      end
      
      self:writeln( ");" )
      self:popIndent(  )
   end
   
   self:writeln( string.format( "if lns_isClosure( %s ) {", formName) )
   do
      local _switchExp = getCRetType( formType:get_retTypeInfoList() )
      if _switchExp == "void" then
         process( string.format( "lns_closure( %s )( _pEnv, %s", formName, formName) )
      elseif _switchExp == cTypeAny then
         process( string.format( "return lns_closure_any( %s )( _pEnv, %s", formName, formName) )
      elseif _switchExp == cTypeInt then
         process( string.format( "return lns_closure_int( %s )( _pEnv, %s", formName, formName) )
      elseif _switchExp == cTypeReal then
         process( string.format( "return lns_closure_real( %s )( _pEnv, %s", formName, formName) )
      elseif _switchExp == cTypeBool then
         process( string.format( "return lns_closure_bool( %s )( _pEnv, %s", formName, formName) )
      else 
         
            process( string.format( "return lns_closure( %s )( _pEnv, %s", formName, formName) )
      end
   end
   
   self:writeln( "}" )
   self:writeln( "else {" )
   do
      local _switchExp = getCRetType( formType:get_retTypeInfoList() )
      if _switchExp == "void" then
         process( string.format( "lns_func( %s )( _pEnv", formName) )
      elseif _switchExp == cTypeAnyP then
         process( string.format( "return lns_func_any( %s )( _pEnv", formName) )
      elseif _switchExp == cTypeInt then
         process( string.format( "return lns_func_int( %s )( _pEnv", formName) )
      elseif _switchExp == cTypeReal then
         process( string.format( "return lns_func_real( %s )( _pEnv", formName) )
      elseif _switchExp == cTypeBool then
         process( string.format( "return lns_func_bool( %s )( _pEnv", formName) )
      else 
         
            process( string.format( "return lns_func( %s )( _pEnv", formName) )
      end
   end
   
   self:writeln( "}" )
end


function convFilter:processDeclForm( node, opt )

   local formType = node:get_expType()
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         self:processFuncPrototype( node, formType:get_accessMode(), true, self.moduleCtrl:getCallFormName( formType ), getCRetType( formType:get_retTypeInfoList() ), node:get_argList(), true )
      elseif _switchExp == ProcessMode.Form then
         self:processFuncPrototype( node, formType:get_accessMode(), true, self.moduleCtrl:getCallFormName( formType ), getCRetType( formType:get_retTypeInfoList() ), node:get_argList(), false )
         
         self:writeln( "{" )
         self:pushIndent(  )
         
         local argNameList = {}
         for index, arg in pairs( node:get_argList() ) do
            do
               local workArg = _lune.__Cast( arg, 3, Nodes.DeclArgNode )
               if workArg ~= nil then
                  table.insert( argNameList, workArg:get_name().txt )
               else
                  do
                     local workArg = _lune.__Cast( arg, 3, Nodes.DeclArgDDDNode )
                     if workArg ~= nil then
                        table.insert( argNameList, "_pDDD" )
                     end
                  end
                  
               end
            end
            
         end
         
         self:processCallUserForm( "_pForm", formType, argNameList )
         
         self:popIndent(  )
         self:writeln( "}" )
      end
   end
   
end


function convFilter:processClosureFunc( declInfo )

   local simpleName = declInfo:get_name()
   if  nil == simpleName then
      local _simpleName = simpleName
   
      return 
   end
   
   local scope = _lune.nilacc( declInfo:get_body(), 'get_scope', 'callmtd' )
   if  nil == scope then
      local _scope = scope
   
      return 
   end
   
   local funcSym = scope:getSymbolInfo( simpleName.txt, scope, false, Ast.ScopeAccess.Normal )
   if  nil == funcSym then
      local _funcSym = funcSym
   
      return 
   end
   
   if not funcSym:get_hasAccessFromClosure() then
      return 
   end
   
   
   local symbolParam = self.scopeMgr:getSymbolParam( funcSym )
   local symbolName = self.moduleCtrl:getSymbolName( funcSym )
   self:writeln( string.format( "%s %s;", cTypeVarP, symbolName) )
   self:write( string.format( "lns_initVal_var( %s, %s, %d, ", symbolName, getBlockName( scope:get_parent() ), symbolParam.index) )
   self:writeln( string.format( "LNS_STEM_ANY( %s ) );", getFunc2any( self.moduleCtrl, self.scopeMgr, funcSym:get_typeInfo() )) )
end


function convFilter:processDeclFunc( node, opt )

   local declInfo = node:get_declInfo()
   local name = self.moduleCtrl:getFuncName( node:get_expType() )
   
   local function processFuncPrototype( termFlag )
   
      self:processFuncPrototype( node, declInfo:get_accessMode(), isClosure( node:get_expType() ), name, getCRetType( node:get_expType():get_retTypeInfoList() ), declInfo:get_argList(), termFlag )
   end
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Form then
      elseif _switchExp == ProcessMode.Prototype then
         processFuncPrototype( true )
         return 
      elseif _switchExp == ProcessMode.InitFuncSym or _switchExp == ProcessMode.WideScopeVer then
         self:process__func__symbol( node:get_expType(), declInfo:get_has__func__Symbol(), self.moduleCtrl:getFuncName( node:get_expType() ) )
         return 
      else 
         
            do
               local _switchExp = opt.node:get_kind()
               if _switchExp == Nodes.NodeKind.get_Block() or _switchExp == Nodes.NodeKind.get_ExpMacroExp() then
               else 
                  
                     self:write( getFunc2any( self.moduleCtrl, self.scopeMgr, node:get_expType() ) )
               end
            end
            
            
            return 
      end
   end
   
   if self.duringDeclFunc then
      if opt.node:get_kind() == Nodes.NodeKind.get_Block() then
         self:processClosureFunc( declInfo )
         
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
   
   processFuncPrototype( false )
   self:writeln( "" )
   self:writeln( "{" )
   
   self:pushRoutine( node:get_expType(), body )
   
   self:processArgClosure( declInfo )
   
   local breakKind = Nodes.BreakKind.None
   
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
         self:write( "if ( lns_isCondTrue( " )
         self:processVal2stem( val:get_exp(), node )
         self:write( ") )" )
      elseif val:get_kind() == Nodes.IfKind.ElseIf then
         self:write( "else if ( lns_isCondTrue( " )
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





function convFilter:processEquals( eqFlag, type1, type2, process1, process2 )

   local valKind1 = getValKind( type1 )
   local valKind2 = getValKind( type2 )
   if valKind1 == ValKind.Stem or valKind2 == ValKind.Stem then
      if not eqFlag then
         self:write( "!" )
      end
      
      self:write( "lns_equals( " )
      process1( ValKind.Stem )
      self:write( "," )
      process2( ValKind.Stem )
      self:write( ")" )
   elseif valKind1 == ValKind.Any and valKind2 == ValKind.Any then
      if not eqFlag then
         self:write( "!" )
      end
      
      self:write( "lns_equals_any( " )
      process1( ValKind.Any )
      self:write( "," )
      process2( ValKind.Any )
      self:write( ")" )
   elseif valKind1 == ValKind.Prim and valKind2 == ValKind.Prim then
      process1( ValKind.Prim )
      
      if eqFlag then
         self:write( " == " )
      else
       
         self:write( " != " )
      end
      
      process2( ValKind.Prim )
   else
    
      Util.err( string.format( "illegal kind %s -- %s", ValKind:_getTxt( valKind1)
      , ValKind:_getTxt( valKind2)
      ) )
   end
   
end


function convFilter:processSwitch( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   
   local expType = node:get_exp():get_expType()
   local expSymName = string.format( "_switchExp%d", node:get_id())
   self:write( string.format( "%s %s = ", getCType( expType ), expSymName) )
   filter( node:get_exp(), self, node )
   self:writeln( ";" )
   
   local expValKind = self:getValKindOfNode( node:get_exp() )
   
   local expStemName
   
   if expValKind ~= ValKind.Stem then
      expStemName = expSymName .. "Stem"
      self:write( string.format( "%s %s = ", cTypeStem, expStemName) )
      
      process2stem( self.stream, self.moduleCtrl, self.scopeMgr, expValKind, expType, node, function (  )
      
         self:write( expSymName )
      end )
      
      self:writeln( ";" )
   else
    
      expStemName = expSymName
   end
   
   
   for index, caseInfo in pairs( node:get_caseList() ) do
      if index == 1 then
         self:write( "if ( " )
      else
       
         self:writeln( "}" )
         self:write( "else if ( " )
      end
      
      local expList = caseInfo:get_expList()
      for listIndex, expNode in pairs( expList:get_expList() ) do
         if listIndex ~= 1 then
            self:write( " || " )
         end
         
         
         self:processEquals( true, expType, expNode:get_expType(), function ( valKind )
         
            if valKind == ValKind.Stem then
               self:write( expStemName )
            else
             
               self:write( expSymName )
            end
            
         end, function ( valKind )
         
            do
               local _switchExp = valKind
               if _switchExp == ValKind.Stem then
                  self:processVal2stem( expNode, node )
               elseif _switchExp == ValKind.Any then
                  self:processVal2any( expNode, node )
               elseif _switchExp == ValKind.Prim then
                  self:accessPrimVal( expNode, node )
               end
            end
            
         end )
      end
      
      self:writeln( " ) {" )
      filter( caseInfo:get_block(), self, node )
   end
   
   do
      local _exp = node:get_default()
      if _exp ~= nil then
         self:writeln( "}" )
         self:writeln( "else {" )
         filter( _exp, self, node )
      end
   end
   
   
   self:writeln( "}" )
   
   self:popIndent(  )
   self:writeln( "}" )
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

   self:processLoopPreProcess( node:get_block() )
   
   self:write( "while ( " )
   
   if node:get_exp():get_expType():get_srcTypeInfo() == Ast.builtinTypeBool then
      filter( node:get_exp(), self, node )
   else
    
      self:write( "lns_isCondTrue( " )
      self:processVal2stem( node:get_exp(), node )
      self:write( ")" )
   end
   
   self:writeln( " )" )
   
   self:writeln( "{" )
   self:pushIndent(  )
   self:writeln( "lns_reset_block( _pEnv );" )
   
   filter( node:get_block(), self, node )
   
   self:popIndent(  )
   self:writeln( "}" )
   self:processLoopPostProcess(  )
end



function convFilter:processRepeat( node, opt )

   self:writeln( "{" )
   self:processLoopPreProcess( node:get_block() )
   
   self:writeln( "while ( true ) {" )
   self:pushIndent(  )
   self:writeln( "lns_reset_block( _pEnv );" )
   
   filter( node:get_block(  ), self, node )
   
   self:write( "if ( " )
   if node:get_exp(  ):get_expType():get_srcTypeInfo() == Ast.builtinTypeBool then
      filter( node:get_exp(  ), self, node )
   else
    
      self:write( "lns_isCondTrue(" )
      self:processVal2stem( node:get_exp(), node )
      self:write( ")" )
   end
   
   self:writeln( ") { break; }" )
   
   self:popIndent(  )
   self:writeln( "}" )
   self:processLoopPostProcess(  )
   self:writeln( "}" )
end



function convFilter:processFor( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   self:writeln( string.format( "%s _to;", cTypeInt) )
   self:writeln( string.format( "%s _inc;", cTypeInt) )
   self:writeln( string.format( "%s %s;", cTypeInt, self.moduleCtrl:getSymbolName( node:get_val() )) )
   
   self:processSetValSingle( node, nil, node:get_val(), true, node:get_to(), 0, false )
   self:writeln( "" )
   self:writeln( string.format( "_to = %s;", self.moduleCtrl:getSymbolName( node:get_val() )) )
   do
      local _exp = node:get_delta(  )
      if _exp ~= nil then
         self:processSetValToSym( node, {node:get_val()}, true, {_exp} )
         self:writeln( string.format( "_inc = %s;", self.moduleCtrl:getSymbolName( node:get_val() )) )
      else
         self:writeln( "_inc = 1;" )
      end
   end
   
   self:processSetValToSym( node, {node:get_val()}, true, {node:get_init()} )
   self:writeln( "" )
   
   self:processLoopPreProcess( node:get_block() )
   
   local indexSym = self.moduleCtrl:getSymbolName( node:get_val() )
   self:writeln( string.format( "for (; ; %s += _inc ) {", indexSym) )
   self:pushIndent(  )
   self:writeln( string.format( "if ( ( _inc >= 0 && %s > _to ) || ( _inc < 0 && %s < _to ) ) { break; }", indexSym, indexSym) )
   self:writeln( "lns_reset_block( _pEnv );" )
   filter( node:get_block(), self, node )
   self:popIndent(  )
   self:writeln( "}" )
   
   self:processLoopPostProcess(  )
   
   self:popIndent(  )
   self:writeln( "}" )
end



function convFilter:processCreateDDD( parent, expList )

   self:write( "lns_createDDD" )
   local lastExp = expList[#expList]
   self:write( string.format( "( _pEnv, %s, %d", tostring( Nodes.hasMultiValNode( lastExp )), #expList) )
   
   for __index, exp in pairs( expList ) do
      self:write( ", " )
      self:processVal2stem( exp, parent )
   end
   
   self:write( ")" )
end


function convFilter:processApply( node, opt )

   self:writeln( "{" )
   local varList = node:get_varList()
   
   local scope = node:get_block():get_scope()
   
   local iteExpTypeList = node:get_expList():get_expTypeList()
   local iteFuncType = iteExpTypeList[1]
   
   local dummyId = varList[1]:get_symbolId()
   
   local dummyScope = Ast.Scope.new(self.moduleTypeInfo:get_scope(), false)
   
   local formSym = _lune.unwrap( dummyScope:addLocalVar( false, false, string.format( "_form%d", dummyId), node:get_pos(), iteFuncType, Ast.MutMode.IMut ))
   local paramSym = _lune.unwrap( dummyScope:addLocalVar( false, false, string.format( "_param%d", dummyId), node:get_pos(), iteExpTypeList[2], Ast.MutMode.IMut ))
   local stateSym = _lune.unwrap( dummyScope:addLocalVar( false, false, string.format( "_state%d", dummyId), node:get_pos(), iteExpTypeList[3], Ast.MutMode.IMut ))
   self.scopeMgr:setupScopeParam( dummyScope )
   
   self:processBlockPreProcess( dummyScope )
   local symList = {}
   table.insert( symList, formSym )
   table.insert( symList, paramSym )
   table.insert( symList, stateSym )
   self:processDeclVarAndSet( symList, node:get_expList() )
   
   self:writeln( "{" )
   self:processLoopPreProcess( node:get_block() )
   
   self:writeln( "while ( true ) {" )
   self:pushIndent(  )
   self:writeln( "lns_reset_block( _pEnv );" )
   
   for __index, varSym in pairs( node:get_varList() ) do
      self:processDeclVarC( true, varSym, false )
   end
   
   
   local workSymName = string.format( "_workMret%d", dummyId)
   self:write( string.format( "%s %s = ", cTypeAnyP, workSymName) )
   
   if formSym:get_typeInfo():get_kind() == Ast.TypeInfoKind.Ext then
      self:write( string.format( "lns_lua_callForm( _pEnv, *%s, ", self.moduleCtrl:getSymbolName( formSym )) )
      self:write( "" )
      local expList = {}
      table.insert( expList, self:createRefNodeFromSym( paramSym ) )
      table.insert( expList, self:createRefNodeFromSym( stateSym ) )
      self:processCreateDDD( node, expList )
      self:writeln( string.format( ")%s;", accessAny) )
   else
    
      self:write( string.format( "lns_func( *%s )( _pEnv, ", self.moduleCtrl:getSymbolName( formSym )) )
      if self.scopeMgr:getSymbolValKind( paramSym ) == ValKind.Any then
         self:write( "*" )
      end
      
      self:writeln( string.format( "%s, %s)%s;", self.moduleCtrl:getSymbolName( paramSym ), self.moduleCtrl:getSymbolName( stateSym ), accessAny) )
   end
   
   
   self:writeln( string.format( "if ( lns_fromDDD( %s, 0 ).type == lns_stem_type_nil ) {", workSymName) )
   self:writeln( "   break;" )
   self:writeln( "}" )
   
   local nodeManager = self.dummyNodeManager
   for index, varSym in pairs( node:get_varList() ) do
      local valKind = self.scopeMgr:getSymbolValKind( varSym )
      local varName = self.moduleCtrl:getSymbolName( varSym )
      do
         local _switchExp = valKind
         if _switchExp == ValKind.Stem then
            self:writeln( string.format( 'lns_setq( _pEnv, %s, lns_fromDDD( %s, %d ) );', varName, workSymName, index - 1) )
         elseif _switchExp == ValKind.Any then
            self:writeln( string.format( 'lns_setq_any( _pEnv, %s, lns_fromDDD( %s, %d )%s );', varName, workSymName, index - 1, accessAny) )
         elseif _switchExp == ValKind.Prim then
            self:writeln( string.format( "%s = lns_fromDDD( %s, %d )%s", varName, workSymName, index - 1, getAccessValFromStem( varSym:get_typeInfo() )) )
         else 
            
               Util.err( string.format( "no support -- %s:%s:%d", varSym:get_name(), ValKind:_getTxt( valKind)
               , 7265) )
         end
      end
      
   end
   
   
   filter( node:get_block(), self, node )
   
   self:processSetSymSingle( node, nil, stateSym, false, node:get_varList()[1], false )
   self:writeln( "" )
   
   self:popIndent(  )
   self:writeln( "}" )
   
   self:processLoopPostProcess(  )
   
   self:writeln( "}" )
   self:processBlockPostProcess(  )
   
   self:writeln( "}" )
end



function convFilter:processForeachSetupVal( parent, scope, workTxt, symTxt, symType )

   local symbolInfo = scope:getSymbolInfoChild( symTxt )
   if  nil == symbolInfo then
      local _symbolInfo = symbolInfo
   
      Util.err( string.format( "not found symTxt -- %s", symTxt) )
   end
   
   
   self:processDeclVarC( true, symbolInfo, false )
   
   local srcSymbol = WorkSymbol.new(symbolInfo:get_scope(), symbolInfo:get_accessMode(), workTxt, symbolInfo:get_typeInfo(), symbolInfo:get_kind(), symbolInfo:get_staticFlag(), SymbolParam.new(ValKind.Stem, 0, cTypeStem))
   
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
         local extedType = extType:get_extedType()
         do
            local _switchExp = extedType:get_kind()
            if _switchExp == Ast.TypeInfoKind.List then
               return _lune.newAlge( CollectionKind.ExtList, {extedType})
            elseif _switchExp == Ast.TypeInfoKind.Array then
               return _lune.newAlge( CollectionKind.ExtArray, {extedType})
            elseif _switchExp == Ast.TypeInfoKind.Set then
               return _lune.newAlge( CollectionKind.ExtSet, {extedType})
            elseif _switchExp == Ast.TypeInfoKind.Map then
               return _lune.newAlge( CollectionKind.ExtMap, {extedType})
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
         self:writeln( string.format( "%s _itAny = lns_itList_new( _pEnv, _obj );", cTypeAnyP) )
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
         self:writeln( string.format( "%s _itAny = lns_itSet_new( _pEnv, _obj );", cTypeAnyP) )
         indexSymbol = nil
         self:writeln( string.format( "%s _val;", cTypeStem) )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( string.format( "%s _itAny = lns_itMap_new( _pEnv, _obj );", cTypeAnyP) )
         indexSymbol = nil
         self:writeln( "lns_Map_entry_t _entry;" )
      elseif _switchExp == Ast.TypeInfoKind.Ext then
         indexSymbol = nil
         do
            local _matchExp = collectionKind
            if _matchExp[1] == CollectionKind.ExtMap[1] then
               local extedType = _matchExp[2][1]
            
               self:writeln( string.format( "%s _itAny = lns_lua_itMap_new( _pEnv, _obj );", cTypeAnyP) )
               self:writeln( "lns_Map_entry_t _entry;" )
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
         
         self:writeln( "for ( ; lns_itList_hasNext( _pEnv, _itAny, &_val );" )
         self:writeln( "      lns_itList_inc( _pEnv, _itAny ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Set then
         
         self:writeln( "for ( ; lns_itSet_hasNext( _pEnv, _itAny, &_val );" )
         self:writeln( "      lns_itSet_inc( _pEnv, _itAny ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         
         self:writeln( "for ( ; lns_itMap_hasNext( _pEnv, _itAny, &_entry );" )
         self:writeln( "      lns_itMap_inc( _pEnv, _itAny ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Ext then
         do
            local _matchExp = collectionKind
            if _matchExp[1] == CollectionKind.ExtMap[1] then
               local extedType = _matchExp[2][1]
            
               self:writeln( "while ( lns_lua_itMap_hasNext( _pEnv, _itAny ) )" )
            end
         end
         
      end
   end
   
   self:writeln( "{" )
   self:pushIndent(  )
   self:writeln( "lns_reset_block( _pEnv );" )
   
   if indexSymbol ~= nil then
      if self.scopeMgr:getSymbolValKind( indexSymbol ) ~= ValKind.Prim then
         self:writeln( string.format( "_%s++;", self.moduleCtrl:getSymbolName( indexSymbol )) )
         self:processDeclVarC( true, indexSymbol, true )
         self:processSetValSingleDirect( node, nil, indexSymbol, true, ValKind.Prim, Ast.builtinTypeInt, 0, false, function (  )
         
            self:write( string.format( "_%s", self.moduleCtrl:getSymbolName( indexSymbol )) )
         end )
         self:writeln( "" )
      else
       
         self:writeln( string.format( "%s++;", self.moduleCtrl:getSymbolName( indexSymbol )) )
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
                  local extedType = _matchExp[2][1]
               
                  self:writeln( "lns_lua_itMap_getEntry( _pEnv, _itAny, &_entry );" )
                  self:processMapForeachSetupVal( node, extedType, node:get_block():get_scope(), node:get_key(), node:get_val(), "_entry.key", "_entry.val" )
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
         self:writeln( string.format( "%s _pList = lns_mtd_Map_createKeyList( _pEnv, _obj );", cTypeAnyP) )
         self:writeln( string.format( "lns_mtd_List( _pList )->sort( _pEnv, _pList, %s );", cValNil) )
         self:writeln( string.format( "%s _itAny = lns_itList_new( _pEnv, _pList );", cTypeAnyP) )
         self:writeln( string.format( "%s _val;", cTypeStem) )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( string.format( "%s _pKeyList = lns_mtd_Map_createKeyList( _pEnv, _obj );", cTypeAnyP) )
         self:writeln( string.format( "lns_mtd_List( _pKeyList )->sort( _pEnv, _pKeyList, %s );", cValNil) )
         self:writeln( string.format( "%s _itAny = lns_itList_new( _pEnv, _pKeyList );", cTypeAnyP) )
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
         self:writeln( "for ( ; lns_itList_hasNext( _pEnv, _itAny, &_val );" )
         self:writeln( "      lns_itList_inc( _pEnv, _itAny ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( "for ( ; lns_itList_hasNext( _pEnv, _itAny, &_key );" )
         self:writeln( "      lns_itList_inc( _pEnv, _itAny ) )" )
      end
   end
   
   self:writeln( "{" )
   
   self:writeln( "lns_reset_block( _pEnv );" )
   
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Set then
         self:processPoolForeachSetupVal( node, loopType, node:get_block():get_scope(), node:get_key(), node:get_val() )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:processMapForeachSetupVal( node, loopType, node:get_block():get_scope(), node:get_key(), node:get_val(), "_key", "lns_mtd_Map_get( _pEnv, _obj, _key )" )
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
            self:write( string.format( "lns_unwrap_%sDefault( ", typeTxt) )
            self:processVal2stem( node:get_exp(), node )
            self:write( "," )
            self:accessPrimVal( defVal, node )
            self:write( ")" )
         else
            self:write( string.format( "lns_unwrap_%s( ", typeTxt) )
            self:processVal2stem( node:get_exp(), node )
            self:write( ")" )
         end
      end
      
   end
   
   do
      local _switchExp = getOrgTypeInfo( node:get_expType() )
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         processUnwrap( "int" )
      elseif _switchExp == Ast.builtinTypeReal then
         processUnwrap( "real" )
      elseif _switchExp == Ast.builtinTypeBool then
         processUnwrap( "bool" )
      else 
         
            do
               local _switchExp = self:getValKindOfNode( node )
               if _switchExp == ValKind.Stem then
                  self:write( "lns_unwrap_stem( " )
               elseif _switchExp == ValKind.Any then
                  self:write( "lns_unwrap_any( " )
               else 
                  
                     Util.err( string.format( "no support -- %s: %d", ValKind:_getTxt( self:getValKindOfNode( node ))
                     , 7686) )
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
                  self:write( string.format( ", %s )", cValNone) )
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
   
   
   self:write( "lns_createMRet" )
   
   local lastExp = expList[#expList]
   self:write( string.format( "( _pEnv, %s, %d", tostring( Nodes.hasMultiValNode( lastExp )), #expList) )
   
   for expIndex, exp in pairs( expList ) do
      self:write( ", " )
      self:processVal2stem( exp, parent )
   end
   
   self:write( ")" )
end


local function needMRetWrap( funcArgTypeList, argNodeList )

   local mRetExp = argNodeList:get_mRetExp()
   if  nil == mRetExp then
      local _mRetExp = mRetExp
   
      return false
   end
   
   local argTypeList = argNodeList:get_expTypeList()
   for index, funcArgType in pairs( funcArgTypeList ) do
      local argType = argTypeList[index]
      if funcArgType:get_kind() == Ast.TypeInfoKind.DDD then
         
         return argType:get_kind() ~= Ast.TypeInfoKind.DDD
      end
      
      if argType:get_kind() == Ast.TypeInfoKind.DDD or mRetExp:get_index() == index then
         
         return true
      end
      
      if mRetExp:get_index() == index then
         return true
      end
      
   end
   
   return true
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
MRetInfo.FormFunc = { "FormFunc", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
MRetInfo._name2Val["FormFunc"] = MRetInfo.FormFunc
MRetInfo.Format = { "Format", {{ func=_lune._toStr, nilable=false, child={} },{ func=Nodes.ExpListNode._fromMap, nilable=false, child={} }}}
MRetInfo._name2Val["Format"] = MRetInfo.Format
MRetInfo.Func = { "Func", {{ func=Nodes.Node._fromMap, nilable=false, child={} }}}
MRetInfo._name2Val["Func"] = MRetInfo.Func
MRetInfo.Method = { "Method", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
MRetInfo._name2Val["Method"] = MRetInfo.Method


function convFilter:processCallWithMRet( parent, mRetFuncName, retTypeName, mRetInfo, argList )

   local mRetExp = argList:get_mRetExp()
   if  nil == mRetExp then
      local _mRetExp = mRetExp
   
      return 
   end
   
   
   do
      local _matchExp = mRetInfo
      if _matchExp[1] == MRetInfo.Method[1] then
         local funcType = _matchExp[2][1]
      
         if not needMRetWrap( funcType:get_argTypeInfoList(), argList ) then
            return 
         end
         
      elseif _matchExp[1] == MRetInfo.Form[1] then
      
         if not needMRetWrap( {Ast.builtinTypeDDD}, argList ) then
            return 
         end
         
      elseif _matchExp[1] == MRetInfo.FormFunc[1] then
         local funcType = _matchExp[2][1]
      
         if not needMRetWrap( funcType:get_argTypeInfoList(), argList ) then
            return 
         end
         
      elseif _matchExp[1] == MRetInfo.Func[1] then
         local funcNode = _matchExp[2][1]
      
         if not needMRetWrap( funcNode:get_expType():get_argTypeInfoList(), argList ) then
            return 
         end
         
      elseif _matchExp[1] == MRetInfo.DDD[1] then
         local node = _matchExp[2][1]
      
      elseif _matchExp[1] == MRetInfo.Format[1] then
         local format = _matchExp[2][1]
         local node = _matchExp[2][2]
      
      end
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
            local funcType = _matchExp[2][1]
         
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
                  self:writeln( string.format( "lns_setMRet( _pEnv, pMRet%s );", accessAny) )
               end
               
               if index >= mRetExp:get_index() then
                  do
                     local _switchExp = argNode:get_kind()
                     if _switchExp == Nodes.NodeKind.get_ExpToDDD() then
                        local toDDDNode = _lune.unwrap( _lune.__Cast( argNode, 3, Nodes.ExpToDDDNode ))
                        self:write( string.format( "%s arg%d = ", cTypeStem, index) )
                        self:write( "lns_createDDD" )
                        local expList = toDDDNode:get_expList():get_expList()
                        local lastExp = expList[#expList]
                        self:write( string.format( "( _pEnv, %s, %d", tostring( Nodes.hasMultiValNode( lastExp )), #expList) )
                        for workIndex, exp in pairs( expList ) do
                           
                           self:write( string.format( ", lns_getMRet( _pEnv, %d )", workIndex + index - 2) )
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
                              if index == mRetExp:get_index() then
                                 self:write( "pMRet" )
                              else
                               
                                 self:write( string.format( "lns_createSubDDD( _pEnv, %d, pMRet )", index - mRetExp:get_index()) )
                              end
                              
                           else
                            
                              self:write( string.format( "lns_getMRet( _pEnv, %d )", index - mRetExp:get_index()) )
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
         
         local function processArg2Stem( index, typeInfo )
         
            if #argTypeList >= index then
               do
                  local _switchExp = argTypeList[index]
                  if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                     self:write( "LNS_STEM_INT(" )
                     self:write( string.format( "arg%d )", index) )
                  elseif _switchExp == Ast.builtinTypeReal then
                     self:write( "LNS_STEM_REAL(" )
                     self:write( string.format( "arg%d )", index) )
                  elseif _switchExp == Ast.builtinTypeBool then
                     self:write( "LNS_STEM_BOOL(" )
                     self:write( string.format( "arg%d )", index) )
                  else 
                     
                        if getValKind( argTypeList[index] ) == ValKind.Any then
                           self:write( "LNS_STEM_ANY(" )
                           self:write( string.format( "arg%d )", index) )
                        else
                         
                           self:write( string.format( "arg%d", index) )
                        end
                        
                  end
               end
               
            else
             
               if typeInfo:get_kind() == Ast.TypeInfoKind.DDD then
                  local offset = index - mRetExp:get_index()
                  if offset > 0 then
                     self:write( string.format( "lns_createSubDDD( _pEnv, %d, pMRet )", index - mRetExp:get_index()) )
                  elseif offset == 0 then
                     self:write( "pMRet" )
                  else
                   
                     self:write( cValDDD0 )
                  end
                  
               else
                
                  self:write( cValNone )
               end
               
            end
            
         end
         
         local function processCreateDDD( expList )
         
            self:write( "lns_createDDD" )
            local lastExp = expList[#expList]
            self:write( string.format( "( _pEnv, %s, %d", tostring( Nodes.hasMultiValNode( lastExp )), #expList) )
            
            for index = 1, #expList do
               self:write( ", " )
               processArg2Stem( index, Ast.builtinTypeNone )
            end
            
         end
         
         local funcTypeInfo = nil
         
         do
            local _matchExp = mRetInfo
            if _matchExp[1] == MRetInfo.Method[1] then
               local funcType = _matchExp[2][1]
            
               funcTypeInfo = funcType
               
               processSetArg( true )
               self:write( string.format( "%s( _pEnv, _pObj", self.moduleCtrl:getCallMethodCName( funcType )) )
            elseif _matchExp[1] == MRetInfo.Form[1] then
            
               processSetArg( true )
               self:write( "lns_closure( _pForm )( _pEnv, pForm" )
               
               processCreateDDD( argList:get_expList() )
            elseif _matchExp[1] == MRetInfo.FormFunc[1] then
               local funcType = _matchExp[2][1]
            
               funcTypeInfo = funcType
               
               processSetArg( true )
               self:write( string.format( "%s( _pEnv, _pForm", self.moduleCtrl:getCallFormName( funcType )) )
            elseif _matchExp[1] == MRetInfo.Func[1] then
               local funcNode = _matchExp[2][1]
            
               funcTypeInfo = funcNode:get_expType()
               
               processSetArg( true )
               local wroteFuncFlag = false
               local builtinFunc = TransUnit.getBuiltinFunc(  )
               
               do
                  local cFuncName = self.moduleCtrl:getBuiltinFuncNameFromType( funcNode:get_expType() )
                  if cFuncName ~= nil then
                     wroteFuncFlag = true
                     self:write( cFuncName .. "(" )
                  end
               end
               
               
               if not wroteFuncFlag then
                  filter( funcNode, self, parent )
                  self:write( "(" )
               end
               
               self:write( " _pEnv" )
            elseif _matchExp[1] == MRetInfo.DDD[1] then
               local expListNode = _matchExp[2][1]
            
               processSetArg( true )
               processCreateDDD( expListNode:get_expList() )
            elseif _matchExp[1] == MRetInfo.Format[1] then
               local format = _matchExp[2][1]
               local expListNode = _matchExp[2][2]
            
               processSetArg( true )
               self:write( "mtd_lns_string_format( _pEnv, " )
               self:write( getLiteralStrAny( format ) )
               self:write( ", " )
               processCreateDDD( expListNode:get_expList() )
               self:write( ")" )
            end
         end
         
         
         if funcTypeInfo ~= nil then
            for index, argType in pairs( funcTypeInfo:get_argTypeInfoList() ) do
               if getValKind( argType ) == ValKind.Stem then
                  self:write( ", " )
                  processArg2Stem( index, argType )
               else
                
                  self:write( string.format( ", arg%d", index) )
               end
               
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

   return string.format( "l_call_mret_%d", node:get_id())
end

function convFilter:processExpToDDD( node, opt )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Intermediate or _switchExp == ProcessMode.Prototype then
         do
            local mRetExp = node:get_expList():get_mRetExp()
            if mRetExp ~= nil then
               if mRetExp:get_index() > 0 then
                  
                  self:processCallWithMRet( node, getMRetFuncName( node ), cTypeStem, _lune.newAlge( MRetInfo.DDD, {node:get_expList()}), node:get_expList() )
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
         
         self:write( ")" )
      else
         self:processCreateDDD( node, expList )
      end
   end
   
end


function convFilter:processExpNew( node, opt )

   self:write( string.format( "%s( _pEnv", self.moduleCtrl:getNewName( node:get_symbol():get_expType() )) )
   do
      local _exp = node:get_argList()
      if _exp ~= nil then
         self:processCallArgList( node:get_ctorTypeInfo(), _exp )
      end
   end
   
   self:write( ")" )
end



function convFilter:processCall( funcSym, funcType, setArgFlag, argList )

   if not setArgFlag then
      self:write( "_pEnv" )
      
      do
         local scope = funcType:get_scope()
         if scope ~= nil then
            if #scope:get_closureSymList() > 0 then
               
               self:write( ", " )
               local setFlag = false
               if funcSym ~= nil then
                  if funcSym:get_hasAccessFromClosure() then
                     self:processSym2Any( funcSym )
                     setFlag = true
                  end
                  
               end
               
               if not setFlag then
                  self:write( getPrepareClosure( self.scopeMgr, "NULL", 0, false, scope:get_closureSymList() ) )
               end
               
            end
            
         end
      end
      
   end
   
   
   if funcType:get_kind() == Ast.TypeInfoKind.Func and funcType:get_rawTxt() == "___init" and funcType:get_parentInfo():get_kind() == Ast.TypeInfoKind.Class then
      self:write( string.format( ", %s", getBlockName( self.ast:get_moduleScope() )) )
   else
    
      if argList ~= nil then
         local expList = {}
         for __index, expNode in pairs( argList:get_expList() ) do
            if expNode:get_expType():get_kind() ~= Ast.TypeInfoKind.Abbr then
               table.insert( expList, expNode )
            end
            
         end
         
         
         self:processCallArgList( funcType, argList )
      else
         self:processCallArgList( funcType, nil )
      end
      
   end
   
   self:write( " )" )
end


function convFilter:processDeclClass( node, opt )

   local classType = node:get_expType()
   local className = self.moduleCtrl:getClassCName( classType )
   local classCanonicalName = self.moduleCtrl:getCanonicalName( classType )
   
   self:writeln( string.format( "// decl class %s (%s)-->", classCanonicalName, ProcessMode:_getTxt( self.processMode)
   ) )
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         self:processDeclClassNodePrototype( node )
      elseif _switchExp == ProcessMode.WideScopeVer then
         do
            local _switchExp = classType:get_kind()
            if _switchExp == Ast.TypeInfoKind.Class then
               
               processIFMethodDataInit( self.stream, self.moduleCtrl, classType, classType )
               
               processClassDataInit( self.stream, self.moduleCtrl, self.scopeMgr, classType, node:get_fieldList() )
            elseif _switchExp == Ast.TypeInfoKind.IF then
               processClassMeta( self.stream, self.moduleCtrl, classType )
            end
         end
         
      elseif _switchExp == ProcessMode.DefClass then
         if classType:get_kind() == Ast.TypeInfoKind.Class then
            self:processDeclClassDef( node )
            self:processMapping( node, classType, Ast.isPubToExternal( classType:get_accessMode() ) and Out2HMode.SourcePub or Out2HMode.SourcePri )
         end
         
      elseif _switchExp == ProcessMode.Form or _switchExp == ProcessMode.InitModule then
         
         do
            local initBlockNode = node:get_initBlock():get_func()
            if initBlockNode ~= nil then
               self:write( string.format( "%s( ", self.moduleCtrl:getMethodCName( initBlockNode:get_expType() )) )
               
               self:processCall( nil, initBlockNode:get_expType(), false, nil )
               self:writeln( ";" )
            end
         end
         
      end
   end
   
   
   self:writeln( string.format( "// <--- decl class %s (%s)", classCanonicalName, ProcessMode:_getTxt( self.processMode)
   ) )
end



local function getFormNilWrapper( node )

   return string.format( "l_nil_form_%d", node:get_id())
end

function convFilter:processExpCallDefWrap( node, opt )

   local funcType
   
   if node:get_nilAccess() then
      funcType = node:get_func():get_expType():get_nonnilableType()
   else
    
      return 
   end
   
   
   if funcType:get_kind() ~= Ast.TypeInfoKind.FormFunc then
      return 
   end
   
   
   local retCode
   
   local retType
   
   if getCRetType( funcType:get_retTypeInfoList() ) == "void" then
      retType = "void"
      retCode = ""
   else
    
      retType = cTypeStem
      retCode = cValNil
   end
   
   
   local argNameList = {}
   self:write( string.format( "static %s %s( %s * _pEnv, %s form", retType, getFormNilWrapper( node ), cTypeEnvP, cTypeStem) )
   for index, argType in pairs( funcType:get_argTypeInfoList() ) do
      local name = string.format( "arg%d", index)
      self:write( string.format( ", %s %s", getCType( argType ), name) )
      table.insert( argNameList, name )
   end
   
   self:writeln( ")" )
   self:writeln( "{" )
   self:pushIndent(  )
   
   self:write( string.format( [==[
if ( form.type == lns_stem_type_nil ) {
   return %s;
}
]==], retCode) )
   
   self:processCallUserForm( string.format( "form%s", accessAny), funcType, argNameList )
   
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processExpCall( node, opt )

   local funcType = node:get_func():get_expType():get_nonnilableType()
   
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
                     mRetInfo = _lune.newAlge( MRetInfo.FormFunc, {node:get_func():get_expType()})
                  else 
                     
                        mRetInfo = _lune.newAlge( MRetInfo.Func, {node:get_func()})
                  end
               end
               
               
               self:processCallWithMRet( node, getMRetFuncName( node ), getCRetType( node:get_expTypeList() ), mRetInfo, argList )
            end
         end
         
         return 
      elseif _switchExp == ProcessMode.DefWrap then
         self:processExpCallDefWrap( node, opt )
         return 
      end
   end
   
   
   local classTypeInfo = getBelongClassType( node:get_func() )
   
   local function process(  )
   
      if funcType:get_kind() == Ast.TypeInfoKind.Form then
         self:write( 'lns_call_form( _pEnv, ' )
         self:processVal2any( node:get_func(), node )
         do
            local argList = node:get_argList()
            if argList ~= nil then
               if #argList:get_expList() > 0 then
                  self:processCallArgList( funcType, argList )
               end
               
            else
               self:write( ', lns_global.ddd0' )
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
                     self:write( string.format( "mtd_lns_string_%s( _pEnv, ", fieldNode:get_field().txt) )
                     filter( prefixNode, self, fieldNode )
                  end
                  
               end
            end
            
         end
         
         return true
      end
      
      local funcSym = nil
      if not fieldCall(  ) then
         return 
      end
      
      
      do
         local refNode = _lune.__Cast( node:get_func(), 3, Nodes.ExpRefNode )
         if refNode ~= nil then
            local builtinFunc = TransUnit.getBuiltinFunc(  )
            
            do
               local cFuncTxt = self.moduleCtrl:getBuiltinFuncNameFromType( refNode:get_expType() )
               if cFuncTxt ~= nil then
                  wroteFuncFlag = true
                  self:write( cFuncTxt .. "(" )
               end
            end
            
         end
      end
      
      
      if not wroteFuncFlag then
         local funcSymList = node:get_func():getSymbolInfo(  )
         if #funcSymList > 0 then
            local workFuncSym = funcSymList[1]:getOrg(  )
            funcSym = workFuncSym
            do
               local cFuncName = self.moduleCtrl:getBuiltinFuncNameFromType( workFuncSym:get_typeInfo() )
               if cFuncName ~= nil then
                  wroteFuncFlag = true
                  self:write( cFuncName )
                  self:write( "(" )
               end
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
                     if node:get_nilAccess() then
                        self:write( self.moduleCtrl:getNilMethodCName( funcType ) )
                        self:write( "( _pEnv, " )
                        self:processVal2stem( fieldNode:get_prefix(), fieldNode )
                     else
                      
                        self:write( self.moduleCtrl:getCallMethodCName( funcType ) )
                        self:write( "( _pEnv, " )
                        self:processVal2any( fieldNode:get_prefix(), fieldNode )
                     end
                     
                  end
               end
               
               wroteFuncFlag = true
               setArgFlag = true
            elseif _switchExp == Ast.TypeInfoKind.Func then
               self:write( string.format( "%s( ", self.moduleCtrl:getFuncName( funcType )) )
               wroteFuncFlag = true
            elseif _switchExp == Ast.TypeInfoKind.FormFunc then
               if node:get_nilAccess() then
                  self:write( string.format( "%s( _pEnv, ", getFormNilWrapper( node )) )
               else
                
                  self:write( string.format( "%s( _pEnv, ", self.moduleCtrl:getCallFormName( funcType )) )
               end
               
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
      
      
      self:processCall( funcSym, funcType, setArgFlag, node:get_argList() )
   end
   
   local function call(  )
   
      local isMret = false
      do
         local argList = node:get_argList()
         if argList ~= nil then
            do
               local mRetExp = argList:get_mRetExp()
               if mRetExp ~= nil then
                  if needMRetWrap( node:get_func():get_expType():get_argTypeInfoList(), argList ) then
                     
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
         self:write( "lns_op_not( _pEnv, " )
         self:processVal2stem( node:get_exp(), node )
         self:write( ")" )
      elseif _switchExp == "#" then
         local expType = node:get_exp():get_expType():get_srcTypeInfo()
         if expType:get_kind() == Ast.TypeInfoKind.List then
            self:write( "lns_mtd_List_len( _pEnv, " )
            self:processVal2any( node:get_exp(), node )
            self:write( ")" )
         elseif expType == Ast.builtinTypeString then
            self:processVal2any( node:get_exp(), node )
            self:write( "->val.str.len" )
         else
          
            Util.err( string.format( "not support type -- %s", expType:getTxt(  )) )
         end
         
      else 
         
            Util.err( string.format( "not support op -- %s", op) )
      end
   end
   
end



function convFilter:processExpMultiTo1( node, opt )

   self:write( "lns_fromDDD( " )
   filter( node:get_exp(), self, node )
   self:write( accessAny )
   self:write( ", 0 )" )
   if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.DDD and Ast.isNumberType( node:get_expType():get_srcTypeInfo():get_nonnilableType() ) then
   else
    
      self:write( getAccessValFromStem( node:get_exp():get_expType() ) )
   end
   
end


function convFilter:processStme2Val( dstType, srcStemTxt )

   do
      local _switchExp = dstType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         self:write( "lns_stem2int( " )
         self:write( srcStemTxt )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeReal then
         self:write( "lns_stem2real( " )
         self:write( srcStemTxt )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeBool then
         self:write( "lns_stem2bool( " )
         self:write( srcStemTxt )
         self:write( ")" )
      else 
         
            self:write( srcStemTxt )
            if getValKind( dstType ) == ValKind.Any then
               self:write( accessAny )
            end
            
      end
   end
   
end


function convFilter:processFuncCast( node )

   
   local castType = node:get_castType()
   do
      local _switchExp = castType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
      else 
         
            return 
      end
   end
   
   
   local orgFunc = node:get_exp():get_expType()
   local closureSymList = _lune.unwrapDefault( _lune.nilacc( orgFunc:get_scope(), 'get_closureSymList', 'callmtd' ), {})
   
   do
      local _switchExp = orgFunc:get_nonnilableType():get_kind()
      if _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
      elseif _switchExp == Ast.TypeInfoKind.Stem then
         
         return 
      else 
         
            Util.err( string.format( "illegal kind -- %s, %d", Ast.TypeInfoKind:_getTxt( orgFunc:get_nonnilableType():get_kind())
            , 9148) )
      end
   end
   
   
   if not needsWrapper( orgFunc, castType ) then
      return 
   end
   
   
   self:write( string.format( "static %s %s( %s _pEnv", getCRetType( castType:get_retTypeInfoList() ), self.moduleCtrl:getFuncCastWrapName( orgFunc, castType ), cTypeEnvP) )
   if #closureSymList > 0 then
      self:write( string.format( ", %s _pForm", cTypeAnyP) )
   end
   
   
   for index, argType in pairs( castType:get_argTypeInfoList() ) do
      self:write( ", " )
      self:write( string.format( "%s arg%d", getCType( argType ), index) )
   end
   
   
   self:write( ")" )
   if self.processMode == ProcessMode.Prototype then
      self:writeln( ";" )
      return 
   end
   
   
   self:writeln( string.format( "// %d", node:get_pos().lineNo) )
   self:writeln( "{" )
   self:pushIndent(  )
   
   for index, typeInfo in pairs( orgFunc:get_argTypeInfoList() ) do
      self:write( string.format( "%s var%d = ", getCType( typeInfo ), index) )
      
      local dstType = orgFunc:get_argTypeInfoList()[index]
      local srcType = castType:get_argTypeInfoList()[index]
      
      if index == #castType:get_argTypeInfoList() and (#orgFunc:get_argTypeInfoList() ~= #castType:get_argTypeInfoList() or not dstType:equals( srcType ) ) then
         if srcType:get_kind() == Ast.TypeInfoKind.DDD and dstType:get_kind() ~= Ast.TypeInfoKind.DDD then
            local dddSym = string.format( "arg%d%s", index, accessAny)
            self:processStme2Val( dstType, string.format( "lns_fromDDD( %s, 0 )", dddSym) )
            self:writeln( ";" )
            for subIndex = index + 1, #orgFunc:get_argTypeInfoList() do
               local dstTypeSub = orgFunc:get_argTypeInfoList()[subIndex]
               self:write( string.format( "%s var%d = ", getCType( typeInfo ), subIndex) )
               if dstTypeSub:get_kind() == Ast.TypeInfoKind.DDD then
                  self:write( string.format( "lns_createSubDDD( _pEnv, %d, arg%d )", subIndex - index, index) )
               else
                
                  local dddSymSub = string.format( "arg%d%s", index, accessAny)
                  self:processStme2Val( dstTypeSub, string.format( "lns_fromDDD( %s, %d )", dddSymSub, subIndex - index) )
               end
               
               self:writeln( ";" )
            end
            
            break
         else
          
            self:write( string.format( "arg%d", index) )
         end
         
      else
       
         self:write( string.format( "arg%d", index) )
      end
      
      
      self:writeln( ";" )
   end
   
   
   if #orgFunc:get_retTypeInfoList() > 0 then
      self:write( string.format( "%s ret = ", getCRetType( orgFunc:get_retTypeInfoList() )) )
   end
   
   self:write( string.format( "%s( _pEnv", self.moduleCtrl:getFuncName( orgFunc )) )
   if #closureSymList > 0 then
      self:write( ", _pForm" )
   end
   
   for index, typeInfo in pairs( orgFunc:get_argTypeInfoList() ) do
      self:write( string.format( ", var%s", tostring( index)) )
   end
   
   self:writeln( ");" )
   
   if #castType:get_retTypeInfoList() > 0 then
      self:write( "return " )
      
      if #orgFunc:get_retTypeInfoList() == 0 then
         if #castType:get_retTypeInfoList() > 1 then
            self:write( cValDDD0 )
         else
          
            self:write( cValNil )
         end
         
      elseif #orgFunc:get_retTypeInfoList() == 1 then
         if #castType:get_retTypeInfoList() > 1 then
            if orgFunc:get_retTypeInfoList()[1]:get_kind() == Ast.TypeInfoKind.DDD then
               self:write( "ret" )
            else
             
               self:write( "lns_createMRet( _pEnv, false, 1, " )
               
               process2stem( self.stream, self.moduleCtrl, self.scopeMgr, getValKind( orgFunc:get_retTypeInfoList()[1] ), orgFunc:get_retTypeInfoList()[1], node, function (  )
               
                  self:write( "ret" )
               end )
               self:write( ")" )
            end
            
         else
          
            if getValKind( castType:get_retTypeInfoList()[1] ) == ValKind.Stem then
               process2stem( self.stream, self.moduleCtrl, self.scopeMgr, getValKind( orgFunc:get_retTypeInfoList()[1] ), orgFunc:get_retTypeInfoList()[1], node, function (  )
               
                  self:write( "ret" )
               end )
            else
             
               self:write( "ret" )
            end
            
         end
         
      else
       
         self:write( "ret" )
      end
      
      self:writeln( ";" )
   end
   
   
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processExpCast( node, opt )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype or _switchExp == ProcessMode.Intermediate then
         self:processFuncCast( node )
         return 
      else 
         
      end
   end
   
   
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
                  self:write( string.format( "lns_getIF( _pEnv, &lns_if_%s( ", self.moduleCtrl:getClassCName( expType )) )
                  self:processVal2any( node:get_exp(), node )
                  
                  self:write( string.format( ")->%s )", self.moduleCtrl:getClassCName( castType )) )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.FormFunc then
               self:processFuncCast2Form( castType, expType )
            elseif _switchExp == Ast.TypeInfoKind.Form then
               self:processFuncCast2Form( castType, expType )
            else 
               
                  filter( exp, self, node )
            end
         end
         
      elseif _switchExp == Nodes.CastKind.Force then
         do
            local _switchExp = getValKind( castType )
            if _switchExp == ValKind.Stem then
               self:processVal2stem( exp, node )
            elseif _switchExp == ValKind.Any then
               self:processVal2any( exp, node )
            elseif _switchExp == ValKind.Prim then
               if isStemType( expType ) then
                  do
                     local _switchExp = castType:get_srcTypeInfo()
                     if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                        self:write( "lns_stem2int( " )
                        filter( exp, self, node )
                        self:write( ")" )
                     elseif _switchExp == Ast.builtinTypeReal then
                        self:write( "lns_stem2real( " )
                        filter( exp, self, node )
                        self:write( ")" )
                     elseif _switchExp == Ast.builtinTypeBool then
                        self:write( "lns_stem2bool( " )
                        filter( exp, self, node )
                        self:write( ")" )
                     else 
                        
                           Util.err( string.format( "illegal cast -- %s", castType:getTxt(  )) )
                     end
                  end
                  
               else
                
                  filter( exp, self, node )
               end
               
            end
         end
         
      elseif _switchExp == Nodes.CastKind.Normal then
         local nonNilCastType = castType:get_nonnilableType()
         if nonNilCastType:get_kind() == Ast.TypeInfoKind.Class and not nonNilCastType:equals( Ast.builtinTypeString ) then
            self:write( "lns_castClass( " )
            self:processVal2stem( exp, node )
            self:write( string.format( ", &%s )", self.moduleCtrl:getClassMetaName( nonNilCastType )) )
         elseif nonNilCastType:get_kind() == Ast.TypeInfoKind.IF then
            self:write( "lns_castIf( _pEnv, " )
            self:processVal2stem( exp, node )
            self:write( string.format( ", &%s )", self.moduleCtrl:getClassMetaName( nonNilCastType )) )
         else
          
            if getValKind( nonNilCastType ) == ValKind.Any then
               local kindTxt
               
               local workType
               
               do
                  local enumType = _lune.__Cast( nonNilCastType, 3, Ast.EnumTypeInfo )
                  if enumType ~= nil then
                     workType = enumType:get_valTypeInfo()
                  else
                     workType = nonNilCastType
                  end
               end
               
               do
                  local _switchExp = workType:get_kind()
                  if _switchExp == Ast.TypeInfoKind.List then
                     kindTxt = "lns_value_type_List"
                  elseif _switchExp == Ast.TypeInfoKind.Array then
                     kindTxt = "lns_value_type_Array"
                  elseif _switchExp == Ast.TypeInfoKind.Map then
                     kindTxt = "lns_value_type_Map"
                  elseif _switchExp == Ast.TypeInfoKind.Class then
                     if workType:equals( Ast.builtinTypeString ) then
                        kindTxt = "lns_value_type_str"
                     else
                      
                        Util.err( "not support" )
                     end
                     
                  elseif _switchExp == Ast.TypeInfoKind.IF then
                     Util.err( "not support" )
                  elseif _switchExp == Ast.TypeInfoKind.Func then
                     kindTxt = "lns_value_type_form"
                  elseif _switchExp == Ast.TypeInfoKind.Alge then
                     kindTxt = "lns_value_type_alge"
                  elseif _switchExp == Ast.TypeInfoKind.DDD then
                     kindTxt = "lns_value_type_ddd"
                  elseif _switchExp == Ast.TypeInfoKind.Set then
                     kindTxt = "lns_value_type_Set"
                  elseif _switchExp == Ast.TypeInfoKind.Form then
                     kindTxt = "lns_value_type_form"
                  elseif _switchExp == Ast.TypeInfoKind.FormFunc then
                     kindTxt = "lns_value_type_form"
                  else 
                     
                        Util.err( string.format( "not support -- %s", castType:getTxt(  )) )
                  end
               end
               
               self:write( "lns_castAny( " )
               self:processVal2stem( exp, node )
               self:write( string.format( ", %s )", kindTxt) )
            else
             
               local kindTxt
               
               if nonNilCastType:get_kind() ~= Ast.TypeInfoKind.Stem then
                  do
                     local _switchExp = nonNilCastType:get_srcTypeInfo()
                     if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                        kindTxt = "lns_stem_type_int"
                     elseif _switchExp == Ast.builtinTypeReal then
                        kindTxt = "lns_stem_type_real"
                     elseif _switchExp == Ast.builtinTypeBool then
                        kindTxt = "lns_stem_type_bool"
                     else 
                        
                           Util.err( string.format( "not support -- %s", castType:getTxt(  )) )
                     end
                  end
                  
                  self:write( "lns_castStem( " )
                  self:processVal2stem( exp, node )
                  self:write( string.format( ", %s )", kindTxt) )
               else
                
                  filter( exp, self, node )
               end
               
            end
            
         end
         
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
      self:writeln( "lns_popVal( _pEnv, lns_incStack( _pEnv ) ||" )
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
    
      self:write( "lns_setStackVal( _pEnv, " )
      self:processVal2stem( node:get_exp1(), node )
      self:write( ") " )
   end
   
   self:writeln( opCC )
   if isAndOr( node:get_exp2() ) then
      filter( node:get_exp2(), self, node )
   else
    
      self:write( "lns_setStackVal( _pEnv, " )
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


function convFilter:processConcat( node, parent )

   self:write( "lns_strconcat( _pEnv, " )
   self:processVal2any( node:get_exp1(), node )
   self:write( ", " )
   self:processVal2any( node:get_exp2(), node )
   self:write( ")" )
end


function convFilter:processExpSetVal( node, opt )

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
   
   
   self:processSetValToNode( node, node:get_exp1(), node:get_initSymSet(), expList, mRetExp )
end



function convFilter:processExpOp2( node, opt )

   local opTxt = node:get_op().txt
   
   do
      local _switchExp = opTxt
      if _switchExp == "and" or _switchExp == "or" then
         self:processAndOr( node, opTxt, opt.node )
      elseif _switchExp == ".." then
         self:processConcat( node, opt.node )
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
                     self:processEquals( opTxt == "==", node:get_exp1():get_expType(), node:get_exp2():get_expType(), function ( valKind )
                     
                        do
                           local _switchExp = valKind
                           if _switchExp == ValKind.Stem then
                              self:processVal2stem( node:get_exp1(), node )
                           elseif _switchExp == ValKind.Any then
                              self:processVal2any( node:get_exp1(), node )
                           elseif _switchExp == ValKind.Prim then
                              self:accessPrimVal( node:get_exp1(), node )
                           end
                        end
                        
                     end, function ( valKind )
                     
                        do
                           local _switchExp = valKind
                           if _switchExp == ValKind.Stem then
                              self:processVal2stem( node:get_exp2(), node )
                           elseif _switchExp == ValKind.Any then
                              self:processVal2any( node:get_exp2(), node )
                           elseif _switchExp == ValKind.Prim then
                              self:accessPrimVal( node:get_exp2(), node )
                           end
                        end
                        
                     end )
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
   
   
   if node:get_symbolInfo():get_name() == "super" then
      local funcType = node:get_expType()
      self:write( string.format( "%s.%s", self:getFullName( funcType:get_parentInfo() ), funcType:get_rawTxt()) )
   elseif node:get_symbolInfo():get_name() == "__mod__" then
      self:write( "*lns_module_path" )
   elseif node:get_symbolInfo():get_name() == "..." then
      self:write( "_pDDD" )
   else
    
      do
         local cFuncName = self.moduleCtrl:getBuiltinFuncNameFromType( node:get_expType() )
         if cFuncName ~= nil then
            self:write( cFuncName )
         else
            local symbolInfo = node:get_symbolInfo()
            local valKind = self.scopeMgr:getSymbolValKind( symbolInfo )
            if valKind == ValKind.Var then
               self:write( string.format( "%s->stem", self.moduleCtrl:getSymbolName( symbolInfo )) )
               self:write( getAccessValFromStem( symbolInfo:get_typeInfo() ) )
            else
             
               
               if symbolInfo:get_kind() == Ast.SymbolKind.Fun or symbolInfo:get_typeInfo():get_kind() == Ast.TypeInfoKind.Func then
                  self:write( self.moduleCtrl:getFuncName( symbolInfo:get_typeInfo() ) )
               else
                
                  
                  if self:isManagedAnySymbol( symbolInfo ) then
                     self:write( string.format( "(*%s)", self.moduleCtrl:getSymbolName( symbolInfo )) )
                  else
                   
                     self:write( self.moduleCtrl:getSymbolName( symbolInfo ) )
                  end
                  
               end
               
            end
            
         end
      end
      
   end
   
end



function convFilter:processExpRefItem( node, opt )

   local function process(  )
   
      self:write( "lns_stem_refAt( _pEnv, " )
      self:processVal2stem( node:get_val(), node )
      self:write( ", " )
      do
         local index = node:get_index()
         if index ~= nil then
            self:processVal2stem( index, node )
         else
            self:write( getLiteralStrStem( string.format( '"%s"', _lune.unwrap( node:get_symbol())) ) )
         end
      end
      
      self:write( ")" )
   end
   
   if node:get_nilAccess() then
      processToIF( self.stream, self.moduleCtrl, node:get_expType(), process )
      return 
   end
   
   
   local val = node:get_val()
   local valType = val:get_expType()
   
   local parent = opt.node
   
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
   elseif node:get_isLValue() then
      Util.err( "not support -- L-Value" )
   else
    
      processToIF( self.stream, self.moduleCtrl, node:get_expType(), function (  )
      
         do
            local _switchExp = valType:get_kind()
            if _switchExp == Ast.TypeInfoKind.List then
               self:write( "lns_mtd_List_refAt( _pEnv, " )
               self:processVal2any( val, node )
               self:write( ", " )
               self:accessPrimVal( _lune.unwrap( node:get_index()), node )
               self:write( ")" )
               self:write( getAccessValFromStem( valType:get_itemTypeInfoList()[1] ) )
            elseif _switchExp == Ast.TypeInfoKind.Map then
               self:write( "lns_mtd_Map_get( _pEnv, " )
               self:processVal2any( val, node )
               self:write( ", " )
               do
                  local index = node:get_index()
                  if index ~= nil then
                     self:processVal2stem( index, node )
                  else
                     self:write( getLiteralStrStem( string.format( '"%s"', _lune.unwrap( node:get_symbol())) ) )
                  end
               end
               
               self:write( ")" )
            elseif _switchExp == Ast.TypeInfoKind.Stem then
               process(  )
            else 
               
                  Util.err( string.format( "not support:%s -- %d:%d", Ast.TypeInfoKind:_getTxt( valType:get_kind())
                  , 9849, node:get_pos().lineNo) )
            end
         end
         
      end )
   end
   
end



function convFilter:processRefField( node, opt )

   if node:get_nilAccess() then
      do
         local symbolInfo = node:get_symbolInfo()
         if symbolInfo ~= nil then
            do
               local _switchExp = symbolInfo:get_kind()
               if _switchExp == Ast.SymbolKind.Mbr then
                  local prefixType = getOrgTypeInfo( node:get_prefix():get_expType() )
                  if prefixType:get_kind() == Ast.TypeInfoKind.Class then
                     self:write( "lns_refFieldNil( _pEnv, " )
                     self:processVal2stem( node:get_prefix(), node )
                     self:write( string.format( ", offsetof( %s, %s ), %s )", self.moduleCtrl:getClassCName( prefixType ), symbolInfo:get_name(), getStemTypeId( symbolInfo:get_typeInfo():get_srcTypeInfo() )) )
                  else
                   
                     Util.err( "not support -- " .. prefixType:getTxt(  ) )
                  end
                  
               else 
                  
                     Util.err( "not support -- " .. Ast.SymbolKind:_getTxt( symbolInfo:get_kind())
                      )
               end
            end
            
         else
            Util.err( "not support" )
         end
      end
      
      return 
   end
   
   
   do
      local symbolInfo = node:get_symbolInfo()
      if symbolInfo ~= nil then
         if symbolInfo:get_typeInfo():get_kind() == Ast.TypeInfoKind.Enum then
            if symbolInfo:get_kind() == Ast.SymbolKind.Mbr then
               if symbolInfo:get_namespaceTypeInfo():get_kind() == Ast.TypeInfoKind.Enum then
                  
                  self:write( self.moduleCtrl:getEnumTypeName( symbolInfo:get_typeInfo() ) )
                  self:write( string.format( "__%s", self.moduleCtrl:getSymbolName( symbolInfo )) )
                  return 
               end
               
            else
             
               Util.err( "illegal access" )
            end
            
         end
         
         
         do
            local _switchExp = symbolInfo:get_kind()
            if _switchExp == Ast.SymbolKind.Mbr then
               if node:get_prefix():get_expType():get_kind() == Ast.TypeInfoKind.Class then
                  if symbolInfo:get_staticFlag() then
                     local symbolName = self.moduleCtrl:getClassMemberName( symbolInfo )
                     if self:isManagedAnySymbol( symbolInfo ) then
                        self:write( string.format( "(*%s)", symbolName) )
                     else
                      
                        self:write( string.format( "%s", symbolName) )
                     end
                     
                  else
                   
                     local className = self.moduleCtrl:getClassCName( node:get_prefix():get_expType() )
                     self:write( string.format( "lns_obj_%s( ", className) )
                     self:processVal2any( node:get_prefix(), node )
                     self:write( string.format( ")->%s", node:get_field().txt) )
                  end
                  
               end
               
            elseif _switchExp == Ast.SymbolKind.Var then
               if node:get_prefix():get_expType():get_kind() == Ast.TypeInfoKind.Module then
                  if symbolInfo:get_staticFlag() then
                     local symbolName = self.moduleCtrl:getSymbolName( symbolInfo )
                     if self:isManagedAnySymbol( symbolInfo ) then
                        self:write( string.format( "(*%s)", symbolName) )
                     else
                      
                        self:write( string.format( "%s", symbolName) )
                     end
                     
                  else
                   
                     local className = self.moduleCtrl:getClassCName( node:get_prefix():get_expType() )
                     self:write( string.format( "lns_obj_%s( ", className) )
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
   local prefixType = prefixNode:get_expType():get_nonnilableType()
   local fieldTxt = node:get_field(  ).txt
   
   do
      local _switchExp = prefixType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Enum then
         if node:get_nilAccess() then
            Util.err( string.format( "not support -- %d:%d:%s", 10005, node:get_pos().lineNo, fieldTxt) )
         end
         
         local enumFullName = self.moduleCtrl:getEnumTypeName( prefixType )
         do
            local _switchExp = fieldTxt
            if _switchExp == "_allList" then
               self:write( string.format( "%s_get__allList( _pEnv )", enumFullName) )
            elseif _switchExp == "_txt" then
               self:write( string.format( "%s_get__txt( _pEnv, ", enumFullName) )
               filter( prefixNode, self, node )
               self:write( ")" )
            else 
               
                  Util.err( string.format( "not support -- %d:%d:%s", 10019, node:get_pos().lineNo, fieldTxt) )
            end
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Alge then
         if node:get_nilAccess() then
            Util.err( string.format( "not support -- %d:%d:%s", 10026, node:get_pos().lineNo, fieldTxt) )
         end
         
         local algeName = self.moduleCtrl:getAlgeCName( prefixType )
         do
            local _switchExp = fieldTxt
            if _switchExp == "_txt" then
               self:write( string.format( "%s_get__txt( _pEnv, ", algeName) )
               self:processVal2any( prefixNode, node )
               self:write( ")" )
            else 
               
                  Util.err( string.format( "not support -- %d:%d:%s", 10037, node:get_pos().lineNo, fieldTxt) )
            end
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
         local getterType = _lune.unwrap( _lune.nilacc( prefixType:get_scope(), 'getTypeInfoField', 'callmtd' , string.format( "get_%s", fieldTxt), true, _lune.unwrap( prefixType:get_scope()), scopeAccess ))
         
         local function process(  )
         
            if node:get_nilAccess() then
               local typeInfo = getCType( getterType:get_retTypeInfoList()[1] )
               do
                  local _switchExp = typeInfo
                  if _switchExp == cTypeInt then
                     self:write( "l_nil_mtd_getter_int( _pEnv, " )
                  elseif _switchExp == cTypeReal then
                     self:write( "l_nil_mtd_getter_real( _pEnv, " )
                  elseif _switchExp == cTypeBool then
                     self:write( "l_nil_mtd_getter_bool( _pEnv, " )
                  elseif _switchExp == cTypeAnyP then
                     self:write( "l_nil_mtd_getter_any( _pEnv, " )
                  elseif _switchExp == cTypeStem then
                     self:write( "l_nil_mtd_getter( _pEnv, " )
                  else 
                     
                        Util.err( string.format( "not support -- %d:%d:%s", 10066, node:get_pos().lineNo, fieldTxt) )
                  end
               end
               
               self:processVal2stem( prefixNode, node )
               self:write( ", " )
               self:write( self.moduleCtrl:getMethodCName( getterType ) )
               self:write( ")" )
            else
             
               self:write( string.format( "%s( _pEnv", self.moduleCtrl:getCallMethodCName( getterType )) )
               if not getterType:get_staticFlag() then
                  self:write( ", " )
                  self:processVal2any( prefixNode, node )
               end
               
               self:write( ")" )
               
               processAlterAccessVal( self.stream, getterType:get_retTypeInfoList(), node:get_expTypeList() )
            end
            
         end
         
         if #node:get_expTypeList() == 1 then
            processAlterToActualType( self.stream, self.moduleCtrl, getterType:get_retTypeInfoList()[1], node:get_expType(), process )
         else
          
            process(  )
         end
         
      else 
         
            Util.err( string.format( "not support -- %d:%d:%s", 10099, node:get_pos().lineNo, Ast.TypeInfoKind:_getTxt( prefixType:get_kind())
            ) )
      end
   end
   
end



function convFilter:processReturn( node, opt )

   local retTypeInfoList = self.routineInfoStack:current(  ):get_funcInfo():get_retTypeInfoList()
   
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
                  
                     Util.err( string.format( "no support -- %d", 10157) )
               end
            end
            
         else
          
         end
         
         self:writeln( ";" )
         if needSetRet then
            do
               local _switchExp = retKind
               if _switchExp == ValKind.Stem then
                  if self.routineInfoStack:get_blockDepth() == 1 then
                     self:writeln( "lns_setRet( _pEnv, _ret );" )
                  else
                   
                     self:writeln( string.format( "lns_setRetAtBlock( LNS_BLOCK_AT( _pEnv, %d ), _ret );", self.routineInfoStack:get_blockDepth()) )
                  end
                  
               elseif _switchExp == ValKind.Any then
                  if self.routineInfoStack:get_blockDepth() == 1 then
                     self:writeln( "lns_setRet( _pEnv, LNS_STEM_ANY( _ret ) );" )
                  else
                   
                     self:writeln( string.format( "lns_setRetAtBlock( LNS_BLOCK_AT( _pEnv, %d ), LNS_STEM_ANY( _ret ) );", self.routineInfoStack:get_blockDepth()) )
                  end
                  
               elseif _switchExp == ValKind.Prim then
               else 
                  
                     Util.err( string.format( "no support -- %d", 10190) )
               end
            end
            
         end
         
      else
         blockStart = false
      end
   end
   
   
   if self.routineInfoStack:get_blockDepth() == 1 then
      self:writeln( "lns_leave_block( _pEnv );" )
   else
    
      self:writeln( string.format( "lns_leave_blockMulti( _pEnv, %d );", self.routineInfoStack:get_blockDepth()) )
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

   if not self.enableTest then
      return 
   end
   
   
   local moduleName = self.moduleCtrl:getFullName( self.moduleTypeInfo )
   local function processDecl(  )
   
      self:write( string.format( "void %s__test_%s( %s _pEnv )", moduleName, node:get_name().txt, cTypeEnvP) )
   end
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         processDecl(  )
         self:writeln( ";" )
      else 
         
            processDecl(  )
            self:writeln( "{" )
            self:pushIndent(  )
            
            self:writeln( string.format( 'printf( "%s:\\n" );', node:get_name().txt) )
            filter( node:get_block(), self, node )
            
            self:writeln( "lns_init_lune_base_Testing( _pEnv );" )
            self:writeln( "lune_base_Testing_outputAllResult( _pEnv, lns_io_stdout );" )
            
            self:popIndent(  )
            self:writeln( "}" )
      end
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
         self:write( "lns_imdInt( " )
         filter( exp, self, parent )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeReal then
         self:write( "lns_imdReal( " )
         filter( exp, self, parent )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeBool then
         self:write( "lns_imdBool( " )
         filter( exp, self, parent )
         self:write( ")" )
      elseif _switchExp == Ast.builtinTypeString then
         do
            local strNode = _lune.__Cast( exp, 3, Nodes.LiteralStringNode )
            if strNode ~= nil then
               if not strNode:get_expList() then
                  self:write( string.format( "lns_imdStr( %s )", str2cstr( strNode:get_token().txt )) )
                  return 
               end
               
            end
         end
         
         self:write( "lns_imdAny( " )
         filter( exp, self, parent )
         
         self:write( ")" )
      else 
         
            do
               local _switchExp = valType:get_kind()
               if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Map or _switchExp == Ast.TypeInfoKind.Array or _switchExp == Ast.TypeInfoKind.Class then
                  self:write( "lns_imdAny( " )
                  filter( exp, self, parent )
                  
                  self:write( ")" )
               elseif _switchExp == Ast.TypeInfoKind.DDD then
                  self:write( "lns_imdAny( " )
                  self:processVal2any( exp, parent )
                  self:write( ")" )
               elseif _switchExp == Ast.TypeInfoKind.Alternate or _switchExp == Ast.TypeInfoKind.Stem or _switchExp == Ast.TypeInfoKind.Alge then
                  self:write( "lns_imdStem( " )
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

   return string.format( "lns_list_%X", node:get_id())
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
      
      self:write( string.format( ", %s %s", ctype, self.moduleCtrl:getSymbolName( symbol )) )
   end
   
   self:writeln( ")" )
   self:writeln( "{" )
   
   self:pushIndent(  )
   self:write( string.format( "lns_imd%s( list", collectionType) )
   self:pushIndent(  )
   for __index, exp in pairs( expListNode:get_expList() ) do
      self:write( ", " )
      self:processLiteralVal( exp, node )
   end
   
   self:popIndent(  )
   self:writeln( ");" )
   self:writeln( string.format( "return lns_create%s( _pEnv, list );", collectionType) )
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
            self:write( string.format( ", %s", self.moduleCtrl:getSymbolName( symbol )) )
         end
         
         self:write( ")" )
      else
       
         self:write( "lns_class_List_new( _pEnv )" )
      end
      
   end
   
end



local function getLiteralSetFuncName( node )

   return string.format( "lns_set_%X", node:get_id())
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
            self:write( string.format( ", %s", self.moduleCtrl:getSymbolName( symbol )) )
         end
         
         self:write( ")" )
      else
       
         self:write( "lns_class_Set_new( _pEnv )" )
      end
      
   end
   
end



local function getLiteralMapFuncName( node )

   return string.format( "lns_map_%X", node:get_id())
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
      self:write( string.format( ", %s %s", self.scopeMgr:getCTypeForSym( symbol ), self.moduleCtrl:getSymbolName( symbol )) )
   end
   
   self:writeln( ")" )
   self:writeln( "{" )
   
   self:pushIndent(  )
   self:write( "lns_imdMap( list" )
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
   self:writeln( "return lns_createMap( _pEnv, list );" )
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
            self:write( string.format( ", %s", self.moduleCtrl:getSymbolName( symbol )) )
         end
         
         self:write( ")" )
      else
       
         self:write( "lns_class_Map_new( _pEnv )" )
      end
      
   end
   
end



local function getLiteralArrayFuncName( node )

   return string.format( "lns_array_%X", node:get_id())
end

function convFilter:processLiteralArray( node, opt )

   if self.processMode == ProcessMode.Immediate and self.processingNode == node then
      self:processLiteralListSub( "List", node, node:get_expList(), getLiteralArrayFuncName( node ) )
   else
    
      if node:get_expList() then
         self:write( string.format( "%s( _pEnv", getLiteralArrayFuncName( node )) )
         local symbolSet = self.literalNode2AccessSymbolSet[node]
         if  nil == symbolSet then
            local _symbolSet = symbolSet
         
            return 
         end
         
         for __index, symbol in pairs( symbolSet:get_list() ) do
            self:write( string.format( ", %s", self.moduleCtrl:getSymbolName( symbol )) )
         end
         
         self:write( ")" )
      else
       
         self:write( "lns_class_List_new( _pEnv )" )
      end
      
   end
   
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

   local txt = str2cstr( node:get_token().txt )
   
   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Prototype then
         do
            local expListNode = node:get_expList()
            if expListNode ~= nil then
               do
                  local mRetExp = expListNode:get_mRetExp()
                  if mRetExp ~= nil then
                     self:processCallWithMRet( node, getMRetFuncName( node ), cTypeAnyP, _lune.newAlge( MRetInfo.Format, {txt,expListNode}), expListNode )
                  else
                     self:write( string.format( "static %s lns_litstr_%d( %s _pEnv", cTypeAnyP, node:get_id(), cTypeEnvP) )
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
                     
                     self:processCallWithMRet( node, getMRetFuncName( node ), cTypeAnyP, _lune.newAlge( MRetInfo.Format, {txt,expListNode}), expListNode )
                  else
                     self:write( string.format( "static %s lns_litstr_%d( %s _pEnv", cTypeAnyP, node:get_id(), cTypeEnvP) )
                     for index, exp in pairs( expListNode:get_expList() ) do
                        self:write( string.format( ", %s arg%d", cTypeStem, index) )
                     end
                     
                     self:writeln( string.format( ") // %d", node:get_pos().lineNo) )
                     self:writeln( "{" )
                     self:pushIndent(  )
                     
                     self:write( "return mtd_lns_string_format( _pEnv, " )
                     self:write( getLiteralStrAny( txt ) )
                     self:write( ", " )
                     
                     local expList = expListNode:get_expList()
                     self:write( "lns_createDDD" )
                     local lastExp = expList[#expList]
                     self:write( string.format( "( _pEnv, %s, %d", tostring( Nodes.hasMultiValNode( lastExp )), #expList) )
                     
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
               self:write( string.format( "lns_litstr_%d( _pEnv ", node:get_id()) )
               for __index, exp in pairs( expListNode:get_expList() ) do
                  self:write( ", " )
                  self:processVal2stem( exp, node )
               end
               
               self:write( ")" )
            end
         end
         
      else
         local opList = TransUnit.findForm( txt )
         
         self:write( getLiteralStrAny( txt ) )
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

   if self.loopInfoStack:get_blockDepth() > 1 then
      if self.loopInfoStack:get_blockDepth() == 2 then
         self:writeln( "lns_leave_block( _pEnv );" )
      else
       
         self:writeln( string.format( "lns_leave_blockMulti( _pEnv, %d );", self.loopInfoStack:get_blockDepth() - 1) )
      end
      
   end
   
   self:write( "break;" )
end



function convFilter:processLiteralSymbol( node, opt )

end



function convFilter:processAbbr( node, opt )

   
   Util.err( "illegal" )
end



local function createFilter( enableTest, outputBuiltin, streamName, stream, headerStream, ast )

   return convFilter.new(enableTest, outputBuiltin, streamName, stream, headerStream, ast)
end
_moduleObj.createFilter = createFilter

local function outputBootcode( stream, launchModuleName )

   local srcStream = Util.SimpleSourceOStream.new(stream, nil, stepIndent)
   
   local launchModulePath = launchModuleName:gsub( "%.", "/" )
   local moduleName = launchModuleName:gsub( "%.", "_" )
   srcStream:writeln( string.format( [==[
#include <lunescript.h>
#include <%s.h>
    
void lns_run_module( %s _pEnv ) {
   %s pInfo = lns_init_%s( _pEnv );
   lns_test( _pEnv, pInfo );      
}
]==], launchModulePath, cTypeEnvP, cTypeModP, moduleName) )
end
_moduleObj.outputBootcode = outputBootcode

return _moduleObj
