--lune/base/convCC.lns
local _moduleObj = {}
local __mod__ = 'lune.base.convCC'
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

local Ver = _lune.loadModule( 'lune.base.Ver' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Util = _lune.loadModule( 'lune.base.Util' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local LuaMod = _lune.loadModule( 'lune.base.LuaMod' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
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
ProcessMode.Main = 4
ProcessMode._val2NameMap[4] = 'Main'
ProcessMode.__allList[5] = ProcessMode.Main

local convFilter = {}
setmetatable( convFilter, { __index = Nodes.Filter,ifList = {oStream,} } )
function convFilter:pushRoutine( funcType )

   self.currentRoutineInfo = RoutineInfo.new(funcType)
   table.insert( self.routineInfoQueue, self.currentRoutineInfo )
end
function convFilter:popRoutine(  )

   self.currentRoutineInfo = self.routineInfoQueue[#self.routineInfoQueue]
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
   self.typeInfo2ModuleName = {}
   self.curLineNo = 1
   self.classId2TypeInfo = {}
   self.classId2MemberList = {}
   self.pubVarName2InfoMap = {}
   self.pubFuncName2InfoMap = {}
   self.pubEnumId2EnumTypeInfo = {}
   self.pubAlgeId2AlgeTypeInfo = {}
   self.needIndent = false
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


function convFilter:processRoot( node, opt )

   Ast.pushProcessInfo( node:get_processInfo() )
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
   
   self.processMode = ProcessMode.Form
   for __index, declFuncNode in pairs( node:get_nodeManager():getDeclFuncNodeList(  ) ) do
      filter( declFuncNode, self, node )
   end
   
   self.processMode = ProcessMode.InitModule
   self:writeln( [==[void lune_init_test( lune_env_t * _pEnv )
{
]==] )
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

   local word = ""
   do
      local _switchExp = node:get_blockKind(  )
      if _switchExp == Nodes.BlockKind.If or _switchExp == Nodes.BlockKind.Elseif then
         word = "{"
      elseif _switchExp == Nodes.BlockKind.Else then
         word = ""
      elseif _switchExp == Nodes.BlockKind.While then
         word = "{"
      elseif _switchExp == Nodes.BlockKind.Repeat then
         word = ""
      elseif _switchExp == Nodes.BlockKind.For then
         word = "{"
      elseif _switchExp == Nodes.BlockKind.Apply then
         word = "{"
      elseif _switchExp == Nodes.BlockKind.Foreach then
         word = "{"
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
   local stmtList = node:get_stmtList(  )
   for __index, statement in pairs( stmtList ) do
      filter( statement, self, node )
      self:writeln( "" )
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

local function getCType( valType )

   local expType = valType:get_srcTypeInfo()
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         return "lune_int_t"
      elseif _switchExp == Ast.builtinTypeReal then
         return "lune_real_t"
      elseif _switchExp == Ast.builtinTypeBool then
         return "lune_bool_t"
      else 
         
            return "lune_stem_t *"
      end
   end
   
end

local function getCTypeForSym( symbol )

   local typeTxt
   
   if symbol:get_isSetFromClosuer() then
      typeTxt = "lune_stem_t *"
   else
    
      typeTxt = getCType( symbol:get_typeInfo() )
   end
   
   return typeTxt, typeTxt == "lune_stem_t *"
end

function convFilter:accessPrimValFromStem( dddFlag, typeInfo, index )

   if dddFlag then
      self:write( string.format( "->val.ddd.pStemList[ %d ]", index) )
   end
   
   local expType = typeInfo:get_srcTypeInfo()
   do
      local _switchExp = expType
      if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
         self:write( "->val.intVal" )
      elseif _switchExp == Ast.builtinTypeReal then
         self:write( "->val.realVal" )
      elseif _switchExp == Ast.builtinTypeBool then
         self:write( "->val.boolVal" )
      end
   end
   
end

local function isStemSym( symbolInfo )

   if symbolInfo:get_isSetFromClosuer() then
      return true
   end
   
   local typeTxt, isStem = getCTypeForSym( symbolInfo )
   return isStem
end

local function isStemVal( node )

   do
      local _switchExp = node:get_kind()
      if _switchExp == Nodes.NodeKind.get_ExpCall() then
         return true
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
         
      end
   end
   
   return false
end

local function hasMultiVal( exp )

   return exp:get_expType():get_kind() == Ast.TypeInfoKind.DDD or #exp:get_expTypeList() > 1
end

function convFilter:processSetValToSym( parent, varSymList, expList )

   local function initVar(  )
   
      local function setVal( var, exp, index )
      
         local work, isStem = getCTypeForSym( var )
         self:write( string.format( "%s", var:get_name()) )
         if isStem then
            self:accessPrimValFromStem( false, var:get_typeInfo(), 0 )
         end
         
         self:write( "=" )
         if hasMultiVal( exp ) then
            self:write( string.format( "lune_fromDDD( _work, %d )", index) )
         else
          
            self:write( "_work" )
         end
         
         if not isStem and isStemVal( exp ) then
            self:accessPrimValFromStem( false, var:get_typeInfo(), 0 )
         end
         
         self:writeln( ";" )
      end
      
      if expList ~= nil then
         for index = 1, #expList do
            if index > #varSymList then
               return index
            end
            
            self:writeln( "{" )
            local exp = expList[index]
            if hasMultiVal( exp ) then
               self:write( "lune_stem_t *" )
            else
             
               self:write( getCType( exp:get_expType() ) )
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
   
   for index = initVar(  ), #varSymList do
      self:write( varSymList[index]:get_name() )
      self:writeln( " = _pEnv->pNilStem;" )
   end
   
end

function convFilter:processDeclVar( node, opt )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.WideScopeVer then
         local varSymList = node:get_symbolInfoList()
         for index, var in pairs( varSymList ) do
            local typeTxt, isStem = getCTypeForSym( var )
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
   if varSymList[1]:get_scope() ~= self.ast:get_moduleScope() then
      for index, var in pairs( varSymList ) do
         local typeTxt, isStem = getCTypeForSym( var )
         self:write( string.format( "%s %s", typeTxt, var:get_name()) )
         if isStem then
            do
               local _switchExp = var:get_typeInfo()
               if _switchExp == Ast.builtinTypeBool then
                  self:write( " = lune_bool2stem( _pEnv, true )" )
               elseif _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
                  self:write( " = lune_int2stem( _pEnv, 0 )" )
               elseif _switchExp == Ast.builtinTypeReal then
                  self:write( " = lune_bool2stem( _pEnv, 0.0 )" )
               else 
                  
                     self:write( " = NULL" )
               end
            end
            
         end
         
         self:writeln( ";" )
      end
      
   end
   
   self:processSetValToSym( node, varSymList, node:get_expList(  ):get_expList() )
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

   self:write( getCType( node:get_expType() ) .. " " )
   self:write( node:get_name(  ).txt )
end


function convFilter:processDeclArgDDD( node, opt )

   self:write( "lune_stem_t * _pDDD" )
end


function convFilter:processExpDDD( node, opt )

end


function convFilter:processDeclFunc( node, opt )

   if self.processMode ~= ProcessMode.Form then
      return 
   end
   
   local declInfo = node:get_declInfo(  )
   local nameToken = declInfo:get_name(  )
   local name = ""
   do
      local _exp = nameToken
      if _exp ~= nil then
         name = _exp.txt
      end
   end
   
   local letTxt = ""
   if declInfo:get_accessMode(  ) ~= Ast.AccessMode.Global and #name ~= 0 then
      letTxt = "static "
   end
   
   self:write( string.format( "%slune_stem_t * %s( lune_env_t * _pEnv, lune_stem_t * _pForm", letTxt, name ) )
   local argList = declInfo:get_argList(  )
   for index, arg in pairs( argList ) do
      self:write( ", " )
      filter( arg, self, node )
   end
   
   self:writeln( " )" )
   self:writeln( "{" )
   local localVerNum = 0
   self:writeln( string.format( "lune_enter_block( _pEnv, %d );", localVerNum) )
   local scope = _lune.unwrap( node:get_expType():get_scope())
   do
      local __sorted = {}
      local __map = scope:get_clojureSymMap()
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local symbol = __map[ __key ]
         do
            self:writeln( string.format( "lune_stem_t * %s = lune_form_closure( _pForm, %d );", symbol:get_name(), _lune.unwrap( scope:get_clojureSym2NumMap()[symbol])) )
         end
      end
   end
   
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
         
            self:writeln( "lune_leave_block( _pEnv );" )
            self:writeln( "return _pEnv->pNoneStem;" )
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

end


function convFilter:processSwitch( node, opt )

end


function convFilter:processMatch( node, opt )

end


function convFilter:processWhile( node, opt )

end


function convFilter:processRepeat( node, opt )

end


function convFilter:processFor( node, opt )

end


function convFilter:processApply( node, opt )

end


function convFilter:processForeach( node, opt )

end


function convFilter:processForsort( node, opt )

end


function convFilter:processExpUnwrap( node, opt )

end

function convFilter:processVal2Stem( node, parent )

   if isStemVal( node ) then
      filter( node, self, parent )
   else
    
      do
         local _switchExp = node:get_expType():get_srcTypeInfo()
         if _switchExp == Ast.builtinTypeBool then
            self:write( "lune_bool2stem( _pEnv, " )
            filter( node, self, parent )
            self:write( ")" )
         elseif _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeChar then
            self:write( "lune_int2stem( _pEnv, " )
            filter( node, self, parent )
            self:write( ")" )
         elseif _switchExp == Ast.builtinTypeReal then
            self:write( "lune_real2stem( _pEnv, " )
            filter( node, self, parent )
            self:write( ")" )
         else 
            
               filter( node, self, parent )
         end
      end
      
   end
   
end

function convFilter:processCreateDDD( ddd, expList, num )

   if ddd then
      self:write( "lune_createDDD" )
   else
    
      self:write( "lune_createMRet" )
   end
   
   local lastExp = expList[#expList]
   self:write( string.format( "( _pEnv, %s, %d", tostring( hasMultiVal( lastExp )), num) )
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
               self:processCreateDDD( true, expList, #expList - index + 1 )
               for expIndex = index, #expList do
                  self:write( ", " )
                  self:processVal2Stem( expList[expIndex], parent )
               end
               
               self:write( ")" )
               return 
            else
             
               filter( expList[index], self, parent )
            end
            
         else
          
            self:write( "_pEnv->pNilStem" )
         end
         
      end
      
   end
   
end

function convFilter:processExpCall( node, opt )

   local wroteFuncFlag = false
   local setArgFlag = false
   local function setArg(  )
   
      self:write( "_pEnv" )
   end
   
   local function fieldCall(  )
   
      local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
      if  nil == fieldNode then
         local _fieldNode = fieldNode
      
         return true
      end
      
      local prefixNode = fieldNode:get_prefix()
      local function processSet(  )
      
         setArgFlag = true
         wroteFuncFlag = true
         do
            local _switchExp = fieldNode:get_field().txt
            if _switchExp == "add" or _switchExp == "del" then
               filter( prefixNode, self, fieldNode )
               self:write( "[" )
               do
                  local argList = node:get_argList()
                  if argList ~= nil then
                     filter( argList, self, fieldNode )
                  end
               end
               
               self:write( "]" )
               do
                  local _switchExp = fieldNode:get_field().txt
                  if _switchExp == "add" then
                     self:write( "= true" )
                  elseif _switchExp == "del" then
                     self:write( "= nil" )
                  end
               end
               
               return false
            end
         end
         
         self:write( string.format( "_lune._Set_%s(", fieldNode:get_field().txt) )
         filter( prefixNode, self, fieldNode )
         return true
      end
      
      local prefixType = prefixNode:get_expType()
      local function processEnumAlge(  )
      
         wroteFuncFlag = true
         local fieldExpType = fieldNode:get_expType()
         local canonicalName = self:getCanonicalName( prefixType )
         local methodName = fieldNode:get_field().txt
         local delimit = ":"
         if methodName == "get__txt" then
            methodName = "_getTxt"
         end
         
         if fieldExpType:get_kind() == Ast.TypeInfoKind.Func then
            delimit = "."
         end
         
         self:write( string.format( "%s%s%s( ", canonicalName, delimit, methodName) )
         if fieldExpType:get_staticFlag() then
            setArgFlag = false
         else
          
            filter( prefixNode, self, fieldNode )
            setArgFlag = true
         end
         
      end
      
      if node:get_nilAccess() then
         wroteFuncFlag = true
         setArgFlag = true
         do
            local _switchExp = prefixType:get_kind()
            if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
               self:write( string.format( "_lune.nilacc( table.%s, nil, 'list', ", fieldNode:get_field().txt) )
               filter( prefixNode, self, fieldNode )
            else 
               
                  self:write( "_lune.nilacc( " )
                  filter( prefixNode, self, fieldNode )
                  self:write( string.format( ", '%s', 'callmtd' ", fieldNode:get_field().txt) )
            end
         end
         
      else
       
         do
            local _switchExp = prefixType:get_kind()
            if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
               setArgFlag = true
               wroteFuncFlag = true
               self:write( string.format( "table.%s( ", fieldNode:get_field().txt) )
               filter( prefixNode, self, fieldNode )
            elseif _switchExp == Ast.TypeInfoKind.Set then
               if not processSet(  ) then
                  return false
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.Alge then
               processEnumAlge(  )
            elseif _switchExp == Ast.TypeInfoKind.Box then
               filter( prefixNode, self, fieldNode )
               self:write( "[1]" )
               return false
            elseif _switchExp == Ast.TypeInfoKind.Class then
               if prefixType:isInheritFrom( Ast.builtinTypeMapping, nil ) and isGenericType( prefixType ) and (fieldNode:get_field().txt == "_fromMap" or fieldNode:get_field().txt == "_fromStem" ) then
                  wroteFuncFlag = true
                  setArgFlag = true
                  filter( node:get_func(), self, node )
                  self:write( "( " )
                  do
                     local argList = node:get_argList()
                     if argList ~= nil then
                        filter( argList, self, node )
                        self:write( ", " )
                     end
                  end
                  
                  self:outputAlter2MapFunc( self, prefixType:createAlt2typeMap( false ) )
                  self:write( ")" )
                  return false
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
         if refNode:get_token().txt == "super" then
            wroteFuncFlag = true
            setArgFlag = true
            local funcType = refNode:get_expType()
            self:write( string.format( "%s.%s( self ", self:getFullName( funcType:get_parentInfo() ), funcType:get_rawTxt()) )
         elseif refNode:get_expType() == builtinFunc.lune_print then
            wroteFuncFlag = true
            self:write( "lune_print(" )
         end
         
      end
   end
   
   if not wroteFuncFlag then
      if node:get_nilAccess() then
         self:write( "_lune.nilacc( " )
         filter( node:get_func(), self, node )
         self:write( ", nil, 'call'" )
         wroteFuncFlag = true
      else
       
         filter( node:get_func(), self, node )
         self:write( "( " )
      end
      
   end
   
   if not setArgFlag then
      self:write( "_pEnv, " )
      do
         local scope = node:get_func():get_expType():get_scope()
         if scope ~= nil then
            if #scope:get_clojureSymList() > 0 then
               self:write( string.format( "lune_func2stem( _pEnv, (lune_func_t *)NULL, %d", #scope:get_clojureSymList()) )
               for __index, symbolInfo in pairs( scope:get_clojureSymList() ) do
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
   elseif op == "not" then
      self:write( "!" )
   else
    
      Util.err( string.format( "not support op -- %s", op) )
   end
   
   filter( node:get_exp(), self, node )
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


function convFilter:accessPrimVal( exp, parent )

   if not isStemVal( exp ) then
      filter( exp, self, parent )
   else
    
      filter( exp, self, parent )
      self:accessPrimValFromStem( #exp:get_expTypeList() > 1, exp:get_expType(), 0 )
   end
   
end

function convFilter:processExpOp2( node, opt )

   local intCast = false
   if node:get_expType():equals( Ast.builtinTypeInt ) and node:get_op().txt == "/" then
      intCast = true
      self:write( "math.floor(" )
   end
   
   local opTxt = node:get_op().txt
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
         self:accessPrimVal( node:get_exp1(), node )
         self:write( " " .. opTxt .. " " )
         self:accessPrimVal( node:get_exp2(), node )
      end
   end
   
   if intCast then
      self:write( ")" )
   end
   
end


function convFilter:processExpRef( node, opt )

   if node:get_token().txt == "super" then
      local funcType = node:get_expType()
      self:write( string.format( "%s.%s", self:getFullName( funcType:get_parentInfo() ), funcType:get_rawTxt()) )
   elseif node:get_expType():equals( TransUnit.getBuiltinFunc(  ).lune__load ) then
   else
    
      local symbolInfo = node:get_symbolInfo()
      if symbolInfo:get_isSetFromClosuer() then
         self:write( string.format( "%s->val.closu", symbolInfo:get_name()) )
      else
       
         if symbolInfo:get_accessMode() == Ast.AccessMode.Pub and symbolInfo:get_kind() == Ast.SymbolKind.Var then
            if self.needModuleObj then
               self:write( "_moduleObj." )
            end
            
         end
         
         self:write( node:get_token().txt )
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
         self:writeln( "{" )
         self:write( "lune_stem_t * _pRet = " )
         if #self.currentRoutineInfo:get_funcInfo():get_retTypeInfoList() >= 2 then
            self:processCreateDDD( false, expList, #expList )
            for __index, expNode in pairs( expList ) do
               self:write( ", " )
               self:processVal2Stem( expNode, node )
            end
            
            self:write( ")" )
         elseif #self.currentRoutineInfo:get_funcInfo():get_retTypeInfoList() == 1 then
            self:processVal2Stem( expList[1], node )
         else
          
            self:write( "_pEnv->pNilStem" )
         end
         
         self:writeln( ";" )
         self:writeln( "lune_setRet( _pEnv, _pRet );" )
      end
   end
   
   if self.currentRoutineInfo:get_blockDepth() == 1 then
      self:writeln( "lune_leave_block( _pEnv );" )
   else
    
      self:writeln( string.format( "lune_leave_blockMulti( _pEnv, %d );", self.currentRoutineInfo:get_blockDepth()) )
   end
   
   self:writeln( "return _pRet;" )
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

function convFilter:processLiteralList( node, opt )

end


function convFilter:processLiteralSet( node, opt )

end


function convFilter:processLiteralMap( node, opt )

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

   self:write( node:get_token().txt )
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
