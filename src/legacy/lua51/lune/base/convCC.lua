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
local cTypeStem = "lune_stem_t"
local cTypeStemP = "lune_stem_t *"
local cTypeEnvP = "lune_env_t *"
local SymbolParam = {}
function SymbolParam.setmeta( obj )
  setmetatable( obj, { __index = SymbolParam  } )
end
function SymbolParam.new( index )
   local obj = {}
   SymbolParam.setmeta( obj )
   if obj.__init then
      obj:__init( index )
   end
   return obj
end
function SymbolParam:__init( index )

   self.index = index
end

local PubVerInfo = {}
function PubVerInfo.setmeta( obj )
  setmetatable( obj, { __index = PubVerInfo  } )
end
function PubVerInfo.new( staticFlag, accessMode, mutable, typeInfo )
   local obj = {}
   PubVerInfo.setmeta( obj )
   if obj.__init then
      obj:__init( staticFlag, accessMode, mutable, typeInfo )
   end
   return obj
end
function PubVerInfo:__init( staticFlag, accessMode, mutable, typeInfo )

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

local function isStemType( valType )

   local expType = valType:get_srcTypeInfo()
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         return false
      elseif _switchExp == Ast.builtinTypeReal then
         return false
      else 
         
            return true
      end
   end
   
end

local function isStemRet( retTypeList )

   do
      local _switchExp = #retTypeList
      if _switchExp == 0 then
         return false
      elseif _switchExp == 1 then
         return isStemType( retTypeList[1] )
      end
   end
   
   return true
end

local function getCType( valType, varFlag )

   local expType = valType:get_srcTypeInfo()
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         return cTypeInt
      elseif _switchExp == Ast.builtinTypeReal then
         return "lune_real_t"
      else 
         
            if varFlag then
               return "lune_var_t *"
            end
            
            return cTypeStemP
      end
   end
   
end

local function getCRetType( retTypeList )

   do
      local _switchExp = #retTypeList
      if _switchExp == 0 then
         return "void"
      elseif _switchExp == 1 then
         return getCType( retTypeList[1], false )
      end
   end
   
   return cTypeStemP
end

local function getCTypeForSym( symbol )

   local typeTxt
   
   if symbol:get_isSetFromClosuer() then
      typeTxt = "lune_var_t *"
   else
    
      typeTxt = getCType( symbol:get_typeInfo(), symbol:get_kind() == Ast.SymbolKind.Var )
   end
   
   return typeTxt, typeTxt == "lune_var_t *"
end

local function isStemSym( symbolInfo )

   if symbolInfo:get_isSetFromClosuer() then
      return true
   end
   
   local typeTxt, isStem = getCTypeForSym( symbolInfo )
   return isStem
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

ProcessMode.Prototype = 0
ProcessMode._val2NameMap[0] = 'Prototype'
ProcessMode.__allList[1] = ProcessMode.Prototype
ProcessMode.WideScopeVer = 1
ProcessMode._val2NameMap[1] = 'WideScopeVer'
ProcessMode.__allList[2] = ProcessMode.WideScopeVer
ProcessMode.InitModule = 2
ProcessMode._val2NameMap[2] = 'InitModule'
ProcessMode.__allList[3] = ProcessMode.InitModule
ProcessMode.Form = 3
ProcessMode._val2NameMap[3] = 'Form'
ProcessMode.__allList[4] = ProcessMode.Form
ProcessMode.Immediate = 4
ProcessMode._val2NameMap[4] = 'Immediate'
ProcessMode.__allList[5] = ProcessMode.Immediate
ProcessMode.Main = 5
ProcessMode._val2NameMap[5] = 'Main'
ProcessMode.__allList[6] = ProcessMode.Main

local convFilter = {}
setmetatable( convFilter, { __index = Nodes.Filter,ifList = {oStream,} } )
function convFilter:pushRoutine( funcType )

   self.currentRoutineInfo = RoutineInfo.new(funcType)
   table.insert( self.routineInfoQueue, self.currentRoutineInfo )
end
function convFilter:popRoutine(  )

   self.currentRoutineInfo = self.routineInfoQueue[#self.routineInfoQueue - 1]
   table.remove( self.routineInfoQueue )
end
function convFilter.new( streamName, stream, ast )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( streamName, stream, ast ); end
   return obj
end
function convFilter:__init(streamName, stream, ast) 
   Nodes.Filter.__init( self)
   
   self.processingNode = nil
   self.processedNodeSet = {}
   self.accessSymbolSet = Util.OrderedSet.new()
   self.literalNode2AccessSymbolSet = {}
   self.duringDeclFunc = false
   self.processMode = ProcessMode.Prototype
   self.routineInfoQueue = {}
   self.currentRoutineInfo = RoutineInfo.new(Ast.builtinTypeNone)
   self.moduleTypeInfo = ast:get_moduleTypeInfo()
   self.moduleSymbolKind = ast:get_moduleSymbolKind()
   self.ast = _lune.unwrap( _lune.__Cast( ast:get_node(), 3, Nodes.RootNode ))
   self.needModuleObj = true
   self.indentQueue = {0}
   self.macroDepth = 0
   self.streamName = streamName
   self.stream = stream
   self.streamQueue = {}
   self.typeInfo2ModuleName = {}
   self.curLineNo = 1
   self.classId2TypeInfo = {}
   self.classId2MemberList = {}
   self.pubVarName2InfoMap = {}
   self.pubFuncName2InfoMap = {}
   self.pubEnumId2EnumTypeInfo = {}
   self.pubAlgeId2AlgeTypeInfo = {}
   self.needIndent = false
   self:pushRoutine( ast:get_moduleTypeInfo() )
end
function convFilter:pushStream(  )

   table.insert( self.streamQueue, self.stream )
   local stream = Util.memStream.new()
   self.stream = stream
   return stream
end
function convFilter:popStream(  )

   if #self.streamQueue == 0 then
      Util.err( "streamQueue is empty." )
   end
   
   self.stream = self.streamQueue[#self.streamQueue]
   table.remove( self.streamQueue )
end
function convFilter:get_indent(  )

   if #self.indentQueue > 0 then
      return self.indentQueue[#self.indentQueue]
   end
   
   return 0
end
function convFilter:getFullName( typeInfo )

   local enumName = typeInfo:getFullName( self.typeInfo2ModuleName, true )
   return string.format( "%s", (enumName:gsub( "&", "" ) ))
end
function convFilter:getCanonicalName( typeInfo )

   return self:getFullName( typeInfo )
end
function convFilter:close(  )

end
function convFilter:flush(  )

end
function convFilter:writeRaw( txt )

   local stream = self.stream
   if self.needIndent then
      stream:write( string.rep( " ", self:get_indent() ) )
      self.needIndent = false
   end
   
   for cr in string.gmatch( txt, "\n" ) do
      self.curLineNo = self.curLineNo + 1
   end
   
   stream:write( txt )
end
function convFilter:write( txt )

   while true do
      do
         local index = string.find( txt, "\n" )
         if index ~= nil then
            self:writeRaw( txt:sub( 1, index ) )
            txt = txt:sub( index + 1 )
         else
            break
         end
      end
      
   end
   
   if #txt > 0 then
      self:writeRaw( txt )
   end
   
   return self, nil
end
function convFilter.setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end

local function filter( node, filter, parent )

   node:processFilter( filter, Opt.new(parent) )
end

local stepIndent = 3
function convFilter:pushIndent( newIndent )

   local indent = _lune.unwrapDefault( newIndent, self:get_indent() + stepIndent)
   table.insert( self.indentQueue, indent )
end

function convFilter:popIndent(  )

   if #self.indentQueue == 0 then
      Util.err( "self.indentQueue == 0" )
   end
   
   table.remove( self.indentQueue )
end

function convFilter:writeln( txt )

   self:write( txt )
   self:write( "\n" )
   self.needIndent = true
end

function convFilter:processNone( node, opt )

end


function convFilter:processImport( node, opt )

   local module = node:get_modulePath(  )
   local moduleName = module:gsub( ".*%.", "" )
   moduleName = node:get_assignName()
   self.typeInfo2ModuleName[node:get_moduleTypeInfo()] = ModuleInfo.new(moduleName, module)
   self:write( string.format( "local %s = _lune.loadModule( '%s' )", moduleName, module) )
end


local function setupScopeParam( scope )

   local stemNum = 0
   for __index, symbol in pairs( scope:get_symbol2SymbolInfoMap() ) do
      if symbol:get_name() ~= "__func__" then
         if symbol:get_kind() == Ast.SymbolKind.Var and isStemSym( symbol ) then
            local param = SymbolParam.new(stemNum)
            symbol:set_convModuleParam( param )
            stemNum = stemNum + 1
         end
         
      end
      
   end
   
   return stemNum
end

local function getSymbolIndex( symbol )

   local param = symbol:get_convModuleParam()
   if  nil == param then
      local _param = param
   
      return 0
   end
   
   return (param ).index
end

function convFilter:processRoot( node, opt )

   Ast.pushProcessInfo( node:get_processInfo() )
   node:visit( function ( target, parent, releation, depth )
   
      print( string.rep( "  ", depth ) .. Nodes.getNodeKindName( target:get_kind() ), releation )
      return Nodes.NodeVisitMode.Child
   end
   , 0 )
   self:writeln( string.format( "// %s", self.streamName) )
   self:writeln( "#include <lunescript.h>" )
   local children = node:get_children(  )
   self.processMode = ProcessMode.Prototype
   for __index, declFuncNode in pairs( node:get_nodeManager():getDeclFuncNodeList(  ) ) do
      filter( declFuncNode, self, node )
   end
   
   self.processMode = ProcessMode.WideScopeVer
   for __index, child in pairs( children ) do
      if child:get_kind() == Nodes.NodeKind.get_DeclVar() then
         filter( child, self, node )
      end
      
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
   
   procssLiteralCtor( node:get_nodeManager():getLiteralListNodeList(  ) )
   procssLiteralCtor( node:get_nodeManager():getLiteralArrayNodeList(  ) )
   procssLiteralCtor( node:get_nodeManager():getLiteralSetNodeList(  ) )
   procssLiteralCtor( node:get_nodeManager():getLiteralMapNodeList(  ) )
   self.processingNode = nil
   self.processMode = ProcessMode.Form
   for __index, declFuncNode in pairs( node:get_nodeManager():getDeclFuncNodeList(  ) ) do
      self.duringDeclFunc = false
      filter( declFuncNode, self, node )
   end
   
   self.processMode = ProcessMode.InitModule
   self:writeln( string.format( [==[void lune_init_test( %s _pEnv )
{
]==], cTypeEnvP) )
   local stemNum = setupScopeParam( self.ast:get_moduleScope() )
   self:writeln( string.format( "lune_block_t * pBlock_%X = lune_enter_module( %d );", self.ast:get_moduleScope():get_scopeId(), stemNum) )
   for __index, child in pairs( children ) do
      do
         local _switchExp = child:get_kind()
         if _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_DeclClass() or _switchExp == Nodes.NodeKind.get_DeclFunc() or _switchExp == Nodes.NodeKind.get_DeclEnum() or _switchExp == Nodes.NodeKind.get_DeclMacro() then
         else 
            
               filter( child, self, node )
               self:writeln( "" )
         end
      end
      
   end
   
   self:writeln( "}" )
   Ast.popProcessInfo(  )
end


function convFilter:processSubfile( node, opt )

end

function convFilter:processBlock( node, opt )

   local stemNum = setupScopeParam( node:get_scope() )
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
            self:writeln( string.format( "lune_var_t * %s = lune_form_closure( _pForm, %d );", symbol:get_name(), _lune.unwrap( scope:get_closureSym2NumMap()[symbol])) )
         end
      end
   end
   
   local loopFlag = false
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
         word = ""
      end
   end
   
   self:writeln( word )
   self:pushIndent(  )
   if not loopFlag then
      self:writeln( string.format( "lune_block_t * pBlock_%X = lune_enter_block( _pEnv, %d );", node:get_scope():get_scopeId(), stemNum) )
   else
    
      self:writeln( "lune_reset_block( _pEnv );" )
   end
   
   local stmtList = node:get_stmtList(  )
   for __index, statement in pairs( stmtList ) do
      filter( statement, self, node )
      self:writeln( "" )
   end
   
   if not loopFlag then
      self:writeln( "lune_leave_block( _pEnv );" )
   end
   
   self:popIndent(  )
   if node:get_blockKind(  ) == Nodes.BlockKind.Block then
      self:writeln( "}" )
   end
   
end


function convFilter:processStmtExp( node, opt )

   filter( node:get_exp(  ), self, node )
   self:write( ";" )
end


function convFilter:processDeclEnum( node, opt )

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

function convFilter:processDeclAlge( node, opt )

end

function convFilter:processNewAlgeVal( node, opt )

end

function convFilter:outputAlter2MapFunc( stream, alt2Map )

end

function convFilter:processDeclClass( node, opt )

end


function convFilter:processDeclMember( node, opt )

end


function convFilter:processExpMacroExp( node, opt )

end



function convFilter:outputDeclMacro( name, argNameList, callback )

end

function convFilter:processDeclMacro( node, opt )

end


function convFilter:processExpMacroStat( node, opt )

end


function convFilter:processExpNew( node, opt )

end


function convFilter:process__func__symbol( has__func__Symbol, classType, funcName )

end

function convFilter:processDeclConstr( node, opt )

end


function convFilter:processDeclDestr( node, opt )

end

function convFilter:processExpCallSuper( node, opt )

end


function convFilter:processDeclMethod( node, opt )

end


function convFilter:processUnwrapSet( node, opt )

end

function convFilter:processIfUnwrap( node, opt )

end

function convFilter:processWhen( node, opt )

end

local function getAccessPrimValFromStem( dddFlag, typeInfo, index )

   local txt = ""
   if dddFlag then
      txt = string.format( "->val.ddd.pStemList[ %d ]", index)
   end
   
   local expType = typeInfo:get_srcTypeInfo()
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         txt = txt .. "->val.intVal"
      elseif _switchExp == Ast.builtinTypeReal then
         txt = txt .. "->val.realVal"
      elseif _switchExp == Ast.builtinTypeBool then
         txt = txt .. "->val.boolVal"
      end
   end
   
   return txt
end

function convFilter:accessPrimValFromStem( dddFlag, typeInfo, index )

   self:write( getAccessPrimValFromStem( dddFlag, typeInfo, index ) )
end

local function isStemVal( node )

   do
      local _switchExp = node:get_kind()
      if _switchExp == Nodes.NodeKind.get_ExpCall() then
         do
            local callNode = _lune.__Cast( node, 3, Nodes.ExpCallNode )
            if callNode ~= nil then
               return isStemRet( callNode:get_func():get_expType():get_retTypeInfoList() )
            end
         end
         
         Util.err( "" )
      elseif _switchExp == Nodes.NodeKind.get_ExpParen() then
         do
            local _exp = _lune.__Cast( node, 3, Nodes.ExpParenNode )
            if _exp ~= nil then
               return isStemVal( _exp:get_exp() )
            end
         end
         
      elseif _switchExp == Nodes.NodeKind.get_ExpRef() then
         do
            local expRefNode = _lune.__Cast( node, 3, Nodes.ExpRefNode )
            if expRefNode ~= nil then
               return isStemSym( expRefNode:get_symbolInfo() )
            end
         end
         
      elseif _switchExp == Nodes.NodeKind.get_ExpOp2() then
         do
            local op2Node = _lune.__Cast( node, 3, Nodes.ExpOp2Node )
            if op2Node ~= nil then
               do
                  local _switchExp = op2Node:get_op().txt
                  if _switchExp == "and" or _switchExp == "or" then
                     return true
                  end
               end
               
            end
         end
         
         return isStemType( node:get_expType() )
      elseif _switchExp == Nodes.NodeKind.get_ExpOp1() then
         do
            local op1Node = _lune.__Cast( node, 3, Nodes.ExpOp1Node )
            if op1Node ~= nil then
               do
                  local _switchExp = op1Node:get_op().txt
                  if _switchExp == "not" then
                     return true
                  end
               end
               
            end
         end
         
         return isStemType( node:get_expType() )
      end
   end
   
   return false
end

function convFilter:accessPrimVal( exp, parent )

   if not isStemVal( exp ) then
      filter( exp, self, parent )
   else
    
      filter( exp, self, parent )
      self:accessPrimValFromStem( #exp:get_expTypeList() > 1, exp:get_expType(), 0 )
   end
   
end

local function hasMultiVal( exp )

   return exp:get_expType():get_kind() == Ast.TypeInfoKind.DDD or #exp:get_expTypeList() > 1
end

function convFilter:processSetValToSym( parent, varSymList, expList )

   local function initVar(  )
   
      local function setVal( var, exp, index )
      
         local work, isStem = getCTypeForSym( var )
         local setValTxt = ""
         if hasMultiVal( exp ) then
            setValTxt = string.format( "lune_fromDDD( _work, %d )", index)
         else
          
            setValTxt = "_work"
         end
         
         if not isStem and isStemVal( exp ) then
            setValTxt = setValTxt .. getAccessPrimValFromStem( false, var:get_typeInfo(), 0 )
         end
         
         if isStem then
            if isStemType( var:get_typeInfo() ) then
               self:writeln( string.format( "lune_setq( _pEnv, &%s->pStem, %s );", var:get_name(), setValTxt) )
            else
             
               self:write( string.format( "%s->pStem", var:get_name()) )
               self:accessPrimValFromStem( false, var:get_typeInfo(), 0 )
               self:writeln( string.format( " = %s;", setValTxt) )
            end
            
         else
          
            self:writeln( string.format( "%s = %s;", var:get_name(), setValTxt) )
         end
         
      end
      
      if expList ~= nil then
         for index = 1, #expList do
            if index > #varSymList then
               return index
            end
            
            self:writeln( "{" )
            local exp = expList[index]
            if hasMultiVal( exp ) then
               self:write( cTypeStemP )
            else
             
               if isStemVal( exp ) then
                  self:write( cTypeStemP )
               else
                
                  self:write( getCType( exp:get_expType(), false ) )
               end
               
            end
            
            self:write( " _work = " )
            filter( exp, self, parent )
            self:writeln( ";" )
            if index == #expList then
               for varIndex = index, #varSymList do
                  if varIndex > #varSymList then
                     self:writeln( "}" )
                     return varIndex
                  end
                  
                  setVal( varSymList[varIndex], exp, varIndex - index )
               end
               
               self:writeln( "}" )
               return #varSymList + 1
            else
             
               setVal( varSymList[index], exp, 0 )
            end
            
            self:writeln( "}" )
         end
         
         return #expList + 1
      end
      
      return 1
   end
   
   initVar(  )
end

function convFilter:processDeclVar( node, opt )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.WideScopeVer then
         local varSymList = node:get_symbolInfoList()
         for index, var in pairs( varSymList ) do
            local typeTxt, isStem = getCTypeForSym( var )
            do
               local _switchExp = var:get_accessMode()
               if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
               else 
                  
                     self:write( "static " )
               end
            end
            
            self:writeln( string.format( "%s %s;", typeTxt, var:get_name()) )
         end
         
         return 
      elseif _switchExp == ProcessMode.Main or _switchExp == ProcessMode.InitModule or _switchExp == ProcessMode.Form then
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
   
   if node:get_mode() ~= Nodes.DeclVarMode.Unwrap and node:get_accessMode(  ) ~= Ast.AccessMode.Global then
   end
   
   local varSymList = node:get_symbolInfoList()
   for index, var in pairs( varSymList ) do
      local typeTxt, isStem = getCTypeForSym( var )
      if isStem then
         if varSymList[1]:get_scope() ~= self.ast:get_moduleScope() then
            self:writeln( string.format( "%s %s; ", typeTxt, var:get_name()) )
         end
         
         local function getDefInitVal( typeInfo )
         
            do
               local _switchExp = var:get_typeInfo():get_srcTypeInfo()
               if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                  return "lune_int2stem( _pEnv, 0 )"
               elseif _switchExp == Ast.builtinTypeReal then
                  return "lune_real2stem( _pEnv, 0.0 )"
               else 
                  
                     return "NULL"
               end
            end
            
         end
         
         local initVal = getDefInitVal( var:get_typeInfo() )
         if initVal == "NULL" then
            self:writeln( string.format( "lune_set_block_stem( pBlock_%X, %d, %s );", var:get_scope():get_scopeId(), getSymbolIndex( var ), var:get_name()) )
         else
          
            self:writeln( string.format( "lune_initVal( %s, pBlock_%X, %d, %s );", var:get_name(), var:get_scope():get_scopeId(), getSymbolIndex( var ), initVal) )
         end
         
      else
       
         if varSymList[1]:get_scope() ~= self.ast:get_moduleScope() then
            self:writeln( string.format( "%s %s;", typeTxt, var:get_name()) )
         end
         
      end
      
   end
   
   self:processSetValToSym( node, varSymList, _lune.nilacc( node:get_expList(  ), 'get_expList', 'callmtd' ) )
   do
      local _exp = node:get_unwrapBlock()
      if _exp ~= nil then
         self:writeln( "" )
         self:write( "if " )
         for index, var in pairs( varSymList ) do
            if index > 1 then
               self:write( " || " )
            end
            
            self:write( " _pEnv->pNilStem == " .. var:get_name() )
         end
         
         self:writeln( " {" )
         self:pushIndent(  )
         for index, var in pairs( varSymList ) do
            self:writeln( string.format( "local _%s = %s", var:get_name(), var:get_name()) )
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
         local name = var:get_name()
         self.pubVarName2InfoMap[name] = PubVerInfo.new(node:get_staticFlag(), node:get_accessMode(), node:get_symbolInfoList()[index]:get_mutable(), node:get_typeInfoList()[index])
      end
      
   end
   
end


function convFilter:processDeclArg( node, opt )

   self:write( getCType( node:get_expType(), false ) .. " " )
   self:write( node:get_name(  ).txt )
end


function convFilter:processDeclArgDDD( node, opt )

   self:write( string.format( "%s _pDDD", cTypeStemP) )
end


function convFilter:processExpDDD( node, opt )

end


local function getFuncName( typeInfo )

   if typeInfo:get_rawTxt() == "" then
      return string.format( "_anonymous_%d", typeInfo:get_typeId())
   end
   
   do
      local _switchExp = typeInfo:get_accessMode()
      if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
         return typeInfo:get_rawTxt()
      end
   end
   
   return string.format( "lune_f_%d_%s", typeInfo:get_typeId(), typeInfo:get_rawTxt())
end

function convFilter:processVal2Stem( node, parent )

   if isStemVal( node ) then
      filter( node, self, parent )
   else
    
      local expType = node:get_expType()
      do
         local _switchExp = expType:get_srcTypeInfo()
         if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
            self:write( "lune_int2stem( _pEnv, " )
            filter( node, self, parent )
            self:write( ")" )
         elseif _switchExp == Ast.builtinTypeReal then
            self:write( "lune_real2stem( _pEnv, " )
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
                           local argList = expType:get_argTypeInfoList()
                           local hasDDD = #argList > 0 and argList[#argList]:get_kind() == Ast.TypeInfoKind.DDD or false
                           self:write( "lune_func2stem( _pEnv, (lune_func_t *)" )
                           self:write( getFuncName( expType ) )
                           self:write( string.format( ", %d, %s, %d", #node:get_expType():get_argTypeInfoList(), tostring( hasDDD), #scope:get_closureSymList()) )
                           for __index, symbolInfo in pairs( scope:get_closureSymList() ) do
                              self:write( string.format( ", %s", symbolInfo:get_name()) )
                           end
                           
                           self:write( ")" )
                        else
                           Util.err( "illegal func" )
                        end
                     end
                     
                  else 
                     
                        filter( node, self, parent )
                  end
               end
               
         end
      end
      
   end
   
end

function convFilter:processDeclFunc( node, opt )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.Form then
      elseif _switchExp == ProcessMode.Prototype then
         return 
      else 
         
            if opt.node:get_kind() ~= Nodes.NodeKind.get_Block() then
               self:processVal2Stem( node, opt.node )
            end
            
            return 
      end
   end
   
   if self.duringDeclFunc then
      if opt.node:get_kind() == Nodes.NodeKind.get_Block() then
         return 
      end
      
      self:write( getFuncName( node:get_expType() ) )
      return 
   end
   
   self.duringDeclFunc = true
   local declInfo = node:get_declInfo(  )
   local name = getFuncName( node:get_expType() )
   local letTxt = ""
   if declInfo:get_accessMode(  ) ~= Ast.AccessMode.Global and #name ~= 0 then
      letTxt = "static "
   end
   
   self:write( string.format( "%s%s %s( %s _pEnv, %s _pForm", letTxt, getCRetType( node:get_expType():get_retTypeInfoList() ), name, cTypeEnvP, cTypeStemP ) )
   local argList = declInfo:get_argList(  )
   for index, arg in pairs( argList ) do
      self:write( ", " )
      filter( arg, self, node )
   end
   
   self:writeln( " )" )
   self:writeln( "{" )
   local breakKind = Nodes.BreakKind.None
   do
      local body = declInfo:get_body()
      if body ~= nil then
         self:process__func__symbol( declInfo:get_has__func__Symbol(), node:get_expType():get_parentInfo(), name )
         self:pushRoutine( node:get_expType() )
         filter( body, self, node )
         self:popRoutine(  )
         breakKind = body:getBreakKind( Nodes.CheckBreakMode.Normal )
      end
   end
   
   do
      local _switchExp = breakKind
      if _switchExp == Nodes.BreakKind.Return or _switchExp == Nodes.BreakKind.NeverRet then
      else 
         
      end
   end
   
   self:writeln( "}" )
   local expType = node:get_expType(  )
   if expType:get_accessMode(  ) == Ast.AccessMode.Pub then
      if self.needModuleObj then
         self:write( string.format( "_moduleObj.%s = %s", name, name) )
      end
      
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
         self:processVal2Stem( val:get_exp(), node )
         self:write( ") )" )
      elseif val:get_kind() == Nodes.IfKind.ElseIf then
         self:write( "else if ( lune_isCondTrue( " )
         self:processVal2Stem( val:get_exp(), node )
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

end


function convFilter:processWhile( node, opt )

end


function convFilter:processRepeat( node, opt )

end


function convFilter:processLoopPreProcess( blockNode )

   local stemNum = setupScopeParam( blockNode:get_scope() )
   self:writeln( string.format( "lune_block_t * pBlock_%X = lune_enter_block( _pEnv, %d );", blockNode:get_scope():get_scopeId(), stemNum) )
end

function convFilter:processLoopPostProcess(  )

   self:writeln( "lune_leave_block( _pEnv );" )
end

function convFilter:processFor( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   self:writeln( string.format( "%s _to;", cTypeInt) )
   self:writeln( string.format( "%s _inc;", cTypeInt) )
   self:writeln( string.format( "%s %s;", cTypeInt, node:get_val():get_name()) )
   self:processSetValToSym( node, {node:get_val()}, {node:get_to(  )} )
   self:writeln( string.format( "_to = %s;", node:get_val():get_name()) )
   do
      local _exp = node:get_delta(  )
      if _exp ~= nil then
         self:processSetValToSym( node, {node:get_val()}, {_exp} )
         self:writeln( string.format( "_inc = %s;", node:get_val():get_name()) )
      else
         self:writeln( "_inc = 1;" )
      end
   end
   
   self:processSetValToSym( node, {node:get_val()}, {node:get_init(  )} )
   self:processLoopPreProcess( node:get_block() )
   self:writeln( string.format( "for (; %s <= _to; %s += _inc )", node:get_val():get_name(), node:get_val():get_name()) )
   filter( node:get_block(), self, node )
   self:writeln( "}" )
   self:processLoopPostProcess(  )
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processApply( node, opt )

end


function convFilter:processForeach( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   self:write( string.format( "%s _obj = ", cTypeStemP) )
   filter( node:get_exp(), self, node )
   self:writeln( ";" )
   local validIndexFlag
   
   local loopType = node:get_exp():get_expType()
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         self:writeln( string.format( "%s _itStem = lune_itList_new( _pEnv, _obj );", cTypeStemP) )
         do
            local _exp = node:get_key()
            if _exp ~= nil then
               self:writeln( string.format( "int %s = 0;", _exp.txt) )
               validIndexFlag = true
            end
         end
         
         self:writeln( string.format( "%s _val;", cTypeStemP) )
      elseif _switchExp == Ast.TypeInfoKind.Set then
         self:writeln( string.format( "%s _itStem = lune_itSet_new( _pEnv, _obj );", cTypeStemP) )
         validIndexFlag = false
         self:writeln( string.format( "%s _val;", cTypeStemP) )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( string.format( "%s _itStem = lune_itMap_new( _pEnv, _obj );", cTypeStemP) )
         validIndexFlag = false
         self:writeln( "lune_Map_entry_t _entry;" )
      else 
         
            Util.err( string.format( "illegal kind -- %s", Ast.TypeInfoKind:_getTxt( loopType:get_kind())
            ) )
      end
   end
   
   self:processLoopPreProcess( node:get_block() )
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         self:writeln( "for ( ; lune_itList_hasNext( _pEnv, _itStem, &_val );" )
         self:writeln( "      lune_itList_inc( _pEnv, _itStem ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Set then
         self:writeln( "for ( ; lune_itSet_hasNext( _pEnv, _itStem, &_val );" )
         self:writeln( "      lune_itSet_inc( _pEnv, _itStem ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( "for ( ; lune_itMap_hasNext( _pEnv, _itStem, &_entry );" )
         self:writeln( "      lune_itMap_inc( _pEnv, _itStem ) )" )
      end
   end
   
   self:writeln( "{" )
   if validIndexFlag then
      do
         local _exp = node:get_key()
         if _exp ~= nil then
            self:writeln( string.format( "   %s++;", _exp.txt) )
         end
      end
      
   end
   
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Array then
         local valType = loopType:get_itemTypeInfoList()[1]
         local valSymTxt
         
         if loopType:get_kind() == Ast.TypeInfoKind.Set then
            do
               local _exp = node:get_key()
               if _exp ~= nil then
                  valSymTxt = _exp.txt
               else
                  valSymTxt = "__val"
               end
            end
            
         else
          
            do
               local _exp = node:get_val()
               if _exp ~= nil then
                  valSymTxt = _exp.txt
               else
                  valSymTxt = "__val"
               end
            end
            
         end
         
         self:writeln( string.format( "   %s %s = _val%s;", getCType( valType, false ), valSymTxt, getAccessPrimValFromStem( false, valType, 0 )) )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         do
            local _exp = node:get_key()
            if _exp ~= nil then
               local keyType = loopType:get_itemTypeInfoList()[1]
               self:writeln( string.format( "   %s %s = _entry.pKey%s;", getCType( keyType, false ), _exp.txt, getAccessPrimValFromStem( false, keyType, 0 )) )
            end
         end
         
         do
            local _exp = node:get_val()
            if _exp ~= nil then
               local valType = loopType:get_itemTypeInfoList()[2]
               self:writeln( string.format( "   %s %s = _entry.pVal%s;", getCType( valType, false ), _exp.txt, getAccessPrimValFromStem( false, valType, 0 )) )
            end
         end
         
      else 
         
      end
   end
   
   filter( node:get_block(), self, node )
   self:writeln( "}" )
   self:processLoopPostProcess(  )
   self:writeln( "}" )
   self:popIndent(  )
end


function convFilter:processForsort( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   self:write( string.format( "%s _obj = ", cTypeStemP) )
   filter( node:get_exp(), self, node )
   self:writeln( ";" )
   local loopType = node:get_exp():get_expType()
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Set then
         self:writeln( "lune_stem_t * _pList = lune_mtd_Map_createKeyList( _pEnv, _obj );" )
         self:writeln( "lune_mtd_List( _pList )->sort( _pEnv, _pList, _pEnv->pNilStem );" )
         self:writeln( string.format( "%s _itStem = lune_itList_new( _pEnv, _pList );", cTypeStemP) )
         self:writeln( string.format( "%s _val;", cTypeStemP) )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( "lune_stem_t * _pKeyList = lune_mtd_Map_createKeyList( _pEnv, _obj );" )
         self:writeln( "lune_mtd_List( _pKeyList )->sort( _pEnv, _pKeyList, _pEnv->pNilStem );" )
         self:writeln( string.format( "%s _itStem = lune_itList_new( _pEnv, _pKeyList );", cTypeStemP) )
         self:writeln( string.format( "%s _key;", cTypeStemP) )
      else 
         
            Util.err( string.format( "illegal kind -- %s", Ast.TypeInfoKind:_getTxt( loopType:get_kind())
            ) )
      end
   end
   
   self:processLoopPreProcess( node:get_block() )
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Set then
         self:writeln( "for ( ; lune_itList_hasNext( _pEnv, _itStem, &_val );" )
         self:writeln( "      lune_itList_inc( _pEnv, _itStem ) )" )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( "for ( ; lune_itList_hasNext( _pEnv, _itStem, &_key );" )
         self:writeln( "      lune_itList_inc( _pEnv, _itStem ) )" )
      end
   end
   
   self:writeln( "{" )
   do
      local _switchExp = loopType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Set then
         local valType = loopType:get_itemTypeInfoList()[1]
         local valSymTxt
         
         do
            local _exp = node:get_key()
            if _exp ~= nil then
               valSymTxt = _exp.txt
            else
               valSymTxt = "__val"
            end
         end
         
         self:writeln( string.format( "   %s %s = _val%s;", getCType( valType, false ), valSymTxt, getAccessPrimValFromStem( false, valType, 0 )) )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         do
            local _exp = node:get_key()
            if _exp ~= nil then
               local keyType = loopType:get_itemTypeInfoList()[1]
               self:writeln( string.format( "   %s %s = _key%s;", getCType( keyType, false ), _exp.txt, getAccessPrimValFromStem( false, keyType, 0 )) )
            end
         end
         
         local valType = loopType:get_itemTypeInfoList()[2]
         self:writeln( string.format( "   %s %s = lune_mtd_Map( _obj )->get( _pEnv, _obj, _key )%s;", getCType( valType, false ), node:get_val().txt, getAccessPrimValFromStem( false, valType, 0 )) )
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
            self:processVal2Stem( node:get_exp(), node )
            self:write( "," )
            self:accessPrimVal( defVal, node )
            self:write( ")" )
         else
            self:write( string.format( "lune_unwrap_%s( ", typeTxt) )
            self:processVal2Stem( node:get_exp(), node )
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
         
            self:write( "lune_unwrap_stem( " )
            self:processVal2Stem( node:get_exp(), node )
            do
               local defVal = node:get_default()
               if defVal ~= nil then
                  self:write( "," )
                  self:processVal2Stem( defVal, node )
                  self:write( ")" )
               else
                  self:write( ", NULL )" )
               end
            end
            
      end
   end
   
end

function convFilter:processCreateDDD( ddd, expList, index, parent )

   if expList[index]:get_expType():get_kind() == Ast.TypeInfoKind.DDD and #expList == index then
      self:write( "_pDDD" )
      return 
   end
   
   if ddd then
      self:write( "lune_createDDD" )
   else
    
      self:write( "lune_createMRet" )
   end
   
   local lastExp = expList[#expList]
   self:write( string.format( "( _pEnv, %s, %d", tostring( hasMultiVal( lastExp )), #expList - index + 1) )
   for expIndex = index, #expList do
      self:write( ", " )
      self:processVal2Stem( expList[expIndex], parent )
   end
   
   self:write( ")" )
end

function convFilter:processCallArgList( funcType, expListNode, parent )

   if expListNode ~= nil then
      local expList = expListNode:get_expList()
      for index, funcArgType in pairs( funcType:get_argTypeInfoList() ) do
         if index ~= 1 then
            self:write( ", " )
         end
         
         if #expList >= index then
            if funcArgType:get_kind() == Ast.TypeInfoKind.DDD then
               self:processCreateDDD( true, expList, index, parent )
               return 
            else
             
               if isStemType( funcArgType ) then
                  self:processVal2Stem( expList[index], parent )
               else
                
                  filter( expList[index], self, parent )
               end
               
            end
            
         else
          
            self:write( "_pEnv->pNilStem" )
         end
         
      end
      
   end
   
end

function convFilter:processExpCall( node, opt )

   if node:get_func():get_expType():get_kind() == Ast.TypeInfoKind.Form then
      self:write( 'lune_call_form( _pEnv, ' )
      filter( node:get_func(), self, node )
      do
         local argList = node:get_argList()
         if argList ~= nil then
            self:write( string.format( ', %d', #argList:get_expList()) )
            if #argList:get_expList() > 0 then
               self:write( ', ' )
               self:processCallArgList( node:get_func():get_expType(), node:get_argList(), node )
            end
            
         else
            self:write( ', 0' )
         end
      end
      
      self:write( ' )' )
      return 
   end
   
   local wroteFuncFlag = false
   local setArgFlag = false
   do
      local refNode = _lune.__Cast( node:get_func(), 3, Nodes.ExpRefNode )
      if refNode ~= nil then
         local builtinFunc = TransUnit.getBuiltinFunc(  )
         if refNode:get_expType() == builtinFunc.lune_print then
            wroteFuncFlag = true
            self:write( "lune_print(" )
         end
         
      end
   end
   
   if not wroteFuncFlag then
      filter( node:get_func(), self, node )
      if node:get_func():get_expType():get_kind() == Ast.TypeInfoKind.Form then
         self:write( "->val.form.pFunc" )
      end
      
      self:write( "( " )
   end
   
   if not setArgFlag then
      self:write( "_pEnv, " )
      do
         local scope = node:get_func():get_expType():get_scope()
         if scope ~= nil then
            if #scope:get_closureSymList() > 0 then
               self:write( string.format( "lune_func2stem( _pEnv, (lune_func_t *)NULL, 0, false, %d", #scope:get_closureSymList()) )
               for __index, symbolInfo in pairs( scope:get_closureSymList() ) do
                  self:write( string.format( ", %s", symbolInfo:get_name()) )
               end
               
               self:write( ")" )
            else
             
               self:write( "NULL" )
            end
            
         else
            self:write( "NULL" )
         end
      end
      
   end
   
   do
      local argList = node:get_argList()
      if argList ~= nil then
         local expList = {}
         for __index, expNode in pairs( argList:get_expList() ) do
            if expNode:get_expType():get_kind() ~= Ast.TypeInfoKind.Abbr then
               table.insert( expList, expNode )
            end
            
         end
         
         if #expList > 0 then
            self:write( ", " )
            self:processCallArgList( node:get_func():get_expType(), node:get_argList(), node )
         end
         
      end
   end
   
   self:write( " )" )
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
   if op == "~" then
      self:write( op )
      self:accessPrimVal( node:get_exp(), node )
   elseif op == "not" then
      self:write( "lune_op_not( _pEnv, " )
      self:processVal2Stem( node:get_exp(), node )
      self:write( ")" )
   else
    
      Util.err( string.format( "not support op -- %s", op) )
   end
   
end


function convFilter:processExpCast( node, opt )

end


function convFilter:processExpParen( node, opt )

   self:write( "(" )
   filter( node:get_exp(), self, node )
   if #node:get_exp():get_expTypeList() > 1 then
      self:write( "->val.ddd.pStemList[ 0 ]" )
   end
   
   self:write( " )" )
end


function convFilter:processWrapForm2Func( funcType )

   self:write( string.format( "static %s _wrap_%s_%d( %s _pEnv, %s _pForm, ", cTypeStemP, funcType:get_rawTxt(), funcType:get_typeId(), cTypeEnvP, cTypeStemP) )
   for index, argType in pairs( funcType:get_argTypeInfoList() ) do
      self:write( string.format( ", %s arg%d", getCType( argType, false ), index) )
   end
   
   self:writeln( ")" )
   self:writeln( "{" )
   self:writeln( 'return %s( _pEnv, pForm' )
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
    
      self:writeln( "lune_setStackVal( _pEnv, " )
      self:processVal2Stem( node:get_exp1(), node )
      self:write( ") " )
   end
   
   if isAndOr( node:get_exp2() ) then
      filter( node:get_exp2(), self, node )
   else
    
      self:writeln( opCC )
      self:writeln( "lune_setStackVal( _pEnv, " )
      self:processVal2Stem( node:get_exp2(), node )
      self:write( ") " )
   end
   
   if firstFlag then
      self:write( ")" )
      self:popIndent(  )
   end
   
end

function convFilter:processExpOp2( node, opt )

   local opTxt = node:get_op().txt
   do
      local _switchExp = opTxt
      if _switchExp == "=" then
         local symbolList = node:get_exp1():getSymbolInfo(  )
         local expList
         
         do
            local expListNode = _lune.__Cast( node:get_exp2(), 3, Nodes.ExpListNode )
            if expListNode ~= nil then
               expList = expListNode:get_expList()
            else
               expList = {node:get_exp2()}
            end
         end
         
         self:processSetValToSym( node, symbolList, expList )
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
                  
                  self:accessPrimVal( node:get_exp1(), node )
                  self:write( " " .. opTxt .. " " )
                  self:accessPrimVal( node:get_exp2(), node )
               else
                  if _lune._Set_has(Ast.compOpSet, opTxt ) then
                     self:write( "lune_bool2stem( _pEnv, " )
                     self:accessPrimVal( node:get_exp1(), node )
                     self:write( " " .. opTxt .. " " )
                     self:accessPrimVal( node:get_exp2(), node )
                     self:write( ")" )
                  else
                   
                     self:accessPrimVal( node:get_exp1(), node )
                     self:write( " " .. opTxt .. " " )
                     self:accessPrimVal( node:get_exp2(), node )
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
   elseif node:get_expType():equals( TransUnit.getBuiltinFunc(  ).lune__load ) then
   else
    
      local symbolInfo = node:get_symbolInfo()
      if isStemSym( symbolInfo ) then
         self:write( string.format( "%s->pStem", symbolInfo:get_name()) )
      else
       
         if symbolInfo:get_accessMode() == Ast.AccessMode.Pub and symbolInfo:get_kind() == Ast.SymbolKind.Var then
            if self.needModuleObj then
               self:write( "_moduleObj." )
            end
            
         end
         
         if symbolInfo:get_kind() == Ast.SymbolKind.Fun then
            self:write( getFuncName( symbolInfo:get_typeInfo() ) )
         else
          
            self:write( node:get_token().txt )
         end
         
      end
      
   end
   
end


function convFilter:processExpRefItem( node, opt )

end


function convFilter:processRefField( node, opt )

end


function convFilter:processExpOmitEnum( node, opt )

end


function convFilter:processGetField( node, opt )

end


function convFilter:processReturn( node, opt )

   do
      local expListNode = node:get_expList()
      if expListNode ~= nil then
         local expList = expListNode:get_expList()
         local retTypeInfoList = self.currentRoutineInfo:get_funcInfo():get_retTypeInfoList()
         local isStem = isStemRet( retTypeInfoList )
         local needSetRet = true
         self:writeln( "{" )
         self:write( string.format( "%s _ret = ", getCRetType( retTypeInfoList )) )
         if #retTypeInfoList >= 2 then
            self:processCreateDDD( false, expList, 1, node )
         elseif #retTypeInfoList == 1 then
            if isStem then
               self:processVal2Stem( expList[1], node )
               if Ast.builtinTypeBool:equals( retTypeInfoList[1] ) then
                  needSetRet = false
               end
               
            else
             
               filter( expList[1], self, node )
            end
            
         else
          
         end
         
         self:writeln( ";" )
         if isStem and needSetRet then
            self:writeln( "lune_setRet( _pEnv, _ret );" )
         end
         
      end
   end
   
   if self.currentRoutineInfo:get_blockDepth() == 1 then
      self:writeln( "lune_leave_block( _pEnv );" )
   else
    
      self:writeln( string.format( "lune_leave_blockMulti( _pEnv, %d );", self.currentRoutineInfo:get_blockDepth()) )
   end
   
   self:writeln( "return _ret;" )
   self:writeln( "}" )
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

   local symbolList = exp:getSymbolInfo(  )
   if #symbolList > 0 then
      local work, varFlag = getCTypeForSym( symbolList[1] )
      if varFlag then
         self:write( "lune_imdStem( " )
         filter( exp, self, parent )
         self:write( ")" )
         return 
      end
      
   end
   
   local valType = exp:get_expType():get_srcTypeInfo()
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
               if #strNode:get_argList() == 0 then
                  self:write( string.format( "lune_imdStr( %s )", strNode:get_token().txt) )
                  return 
               end
               
            end
         end
         
         self:write( "lune_imdStem( " )
         filter( exp, self, parent )
         self:write( ")" )
      else 
         
            do
               local _switchExp = valType:get_kind()
               if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Map or _switchExp == Ast.TypeInfoKind.Array then
                  self:write( "lune_imdStem( " )
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
   self:write( string.format( "static %s %s( %s _pEnv", cTypeStemP, literalFuncName, cTypeEnvP) )
   for __index, symbol in pairs( self.accessSymbolSet:get_list() ) do
      self:write( string.format( ", %s %s", getCTypeForSym( symbol ), symbol:get_name()) )
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
   self:writeln( string.format( "return lune_create%s( _pEnv, &list );", collectionType) )
   self:popIndent(  )
   self:writeln( "}" )
   self.literalNode2AccessSymbolSet[node] = self.accessSymbolSet:clone(  )
end

function convFilter:processLiteralList( node, opt )

   if self.processMode == ProcessMode.Immediate and self.processingNode == node then
      self:processLiteralListSub( "List", node, node:get_expList(), getLiteralListFuncName( node ) )
   else
    
      self:write( string.format( "%s( _pEnv", getLiteralListFuncName( node )) )
      local symbolSet = self.literalNode2AccessSymbolSet[node]
      if  nil == symbolSet then
         local _symbolSet = symbolSet
      
         return 
      end
      
      for __index, symbol in pairs( symbolSet:get_list() ) do
         self:write( string.format( ", %s", symbol:get_name()) )
      end
      
      self:write( ")" )
   end
   
end


local function getLiteralSetFuncName( node )

   return string.format( "lune_set_%X", node:get_id())
end

function convFilter:processLiteralSet( node, opt )

   if self.processMode == ProcessMode.Immediate and self.processingNode == node then
      self:processLiteralListSub( "Set", node, node:get_expList(), getLiteralSetFuncName( node ) )
   else
    
      self:write( string.format( "%s( _pEnv", getLiteralSetFuncName( node )) )
      local symbolSet = self.literalNode2AccessSymbolSet[node]
      if  nil == symbolSet then
         local _symbolSet = symbolSet
      
         return 
      end
      
      for __index, symbol in pairs( symbolSet:get_list() ) do
         self:write( string.format( ", %s", symbol:get_name()) )
      end
      
      self:write( ")" )
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
   self:write( string.format( "static %s %s( %s _pEnv", cTypeStemP, getLiteralMapFuncName( node ), cTypeEnvP) )
   for __index, symbol in pairs( self.accessSymbolSet:get_list() ) do
      self:write( string.format( ", %s %s", getCTypeForSym( symbol ), symbol:get_name()) )
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
   self:writeln( "return lune_createMap( _pEnv, &list );" )
   self:popIndent(  )
   self:writeln( "}" )
   self.literalNode2AccessSymbolSet[node] = self.accessSymbolSet:clone(  )
end

function convFilter:processLiteralMap( node, opt )

   if self.processMode == ProcessMode.Immediate and self.processingNode == node then
      self:processLiteralMapSub( node )
   else
    
      self:write( string.format( "%s( _pEnv", getLiteralMapFuncName( node )) )
      local symbolSet = self.literalNode2AccessSymbolSet[node]
      if  nil == symbolSet then
         local _symbolSet = symbolSet
      
         return 
      end
      
      for __index, symbol in pairs( symbolSet:get_list() ) do
         self:write( string.format( ", %s", symbol:get_name()) )
      end
      
      self:write( ")" )
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

   local txt = node:get_token(  ).txt
   if string.find( txt, '^```' ) then
      txt = '[==[' .. txt:sub( 4, -4 ) .. ']==]'
   end
   
   local opList = TransUnit.findForm( txt )
   self:write( "lune_litStr2stem( _pEnv, " )
   local argList = node:get_argList(  )
   if #argList > 0 then
      self:write( string.format( 'string.format( %s, ', txt ) )
      for index, val in pairs( argList ) do
         if index > 1 then
            self:write( ", " )
         end
         
         filter( val, self, node )
      end
      
      self:write( ")" )
   else
    
      self:write( txt )
   end
   
   self:write( ")" )
end


function convFilter:processLiteralBool( node, opt )

   if node:get_token().txt == "true" then
      self:write( "_pEnv->pTrueStem" )
   else
    
      self:write( "_pEnv->pFalseStem" )
   end
   
end


function convFilter:processLiteralNil( node, opt )

   self:write( "_pEnv->pNilStem" )
end


function convFilter:processBreak( node, opt )

   self:write( "break" )
end


function convFilter:processLiteralSymbol( node, opt )

end


local function createFilter( streamName, stream, ast )

   return convFilter.new(streamName, stream, ast)
end
_moduleObj.createFilter = createFilter
return _moduleObj
