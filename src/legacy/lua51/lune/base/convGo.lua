--lune/base/convGo.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@convGo'
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

ProcessMode.NonClosureFuncDecl = 0
ProcessMode._val2NameMap[0] = 'NonClosureFuncDecl'
ProcessMode.__allList[1] = ProcessMode.NonClosureFuncDecl
ProcessMode.Main = 1
ProcessMode._val2NameMap[1] = 'Main'
ProcessMode.__allList[2] = ProcessMode.Main


local convFilter = {}
setmetatable( convFilter, { __index = Nodes.Filter } )
function convFilter.new( enableTest, streamName, stream, ast )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( enableTest, streamName, stream, ast ); end
   return obj
end
function convFilter:__init(enableTest, streamName, stream, ast) 
   Nodes.Filter.__init( self,true, ast:get_moduleTypeInfo(), ast:get_moduleTypeInfo():get_scope())
   
   
   self.stream = Util.SimpleSourceOStream.new(stream, nil, 4)
   self.processMode = ProcessMode.Main
   self.processModeStack = {}
end
function convFilter:pushProcessMode( mode )

   table.insert( self.processModeStack, self.processMode )
   self.processMode = mode
end
function convFilter:popProcessMode(  )

   self.processMode = self.processModeStack[#self.processModeStack]
   table.remove( self.processModeStack )
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



local function filter( node, filter, parent )

   node:processFilter( filter, Opt.new(parent) )
end

local function isAnyType( typeInfo )

   local work = typeInfo:get_srcTypeInfo()
   return typeInfo:get_nilable() or work == Ast.builtinTypeStem
end

local function isClosure( typeInfo )

   local scope = typeInfo:get_scope()
   if  nil == scope then
      local _scope = scope
   
      return false
   end
   
   return #scope:get_closureSymList() > 0
end

local golanKeywordSet = {["func"] = true, ["select"] = true, ["defer"] = true, ["go"] = true, ["chan"] = true, ["package"] = true, ["const"] = true, ["fallthrough"] = true, ["range"] = true, ["continue"] = true, ["var"] = true}

local SymbolKind = {}
SymbolKind._name2Val = {}
function SymbolKind:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "SymbolKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function SymbolKind._from( val )
   return _lune._AlgeFrom( SymbolKind, val )
end

SymbolKind.Class = { "Class", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
SymbolKind._name2Val["Class"] = SymbolKind.Class
SymbolKind.Func = { "Func", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
SymbolKind._name2Val["Func"] = SymbolKind.Func
SymbolKind.Member = { "Member", {{ func=_lune._toBool, nilable=false, child={} }}}
SymbolKind._name2Val["Member"] = SymbolKind.Member
SymbolKind.Normal = { "Normal"}
SymbolKind._name2Val["Normal"] = SymbolKind.Normal
SymbolKind.PubVar = { "PubVar"}
SymbolKind._name2Val["PubVar"] = SymbolKind.PubVar


local function concatGLSym( name, external )

   return (external and "G" or "l" ) .. name
end
local function isInnerDeclType( typeInfo )

   if typeInfo:get_parentInfo():get_kind() ~= Ast.TypeInfoKind.Module or _lune.nilacc( _lune.nilacc( typeInfo:get_scope(), 'get_parent', 'callmtd' ), 'get_ownerTypeInfo', 'callmtd' ) == nil then
      return true
   end
   
   return false
end

local function concatSymWithType( name, typeInfo )

   return concatGLSym( string.format( "%s_", (typeInfo:getModule(  ):get_rawTxt():gsub( "@", "" ) )) .. name, Ast.isPubToExternal( typeInfo:get_accessMode() ) )
end

function convFilter:getSymbol( kind, name )

   local symbolName
   
   if _lune._Set_has(golanKeywordSet, name ) then
      symbolName = string.format( "%s_", name)
   else
    
      symbolName = name
   end
   
   
   do
      local _matchExp = kind
      if _matchExp[1] == SymbolKind.PubVar[1] then
      
         symbolName = concatGLSym( symbolName, true )
      elseif _matchExp[1] == SymbolKind.Member[1] then
         local external = _matchExp[2][1]
      
         symbolName = concatGLSym( symbolName, external )
      elseif _matchExp[1] == SymbolKind.Func[1] then
         local typeInfo = _matchExp[2][1]
      
         if typeInfo:get_kind() == Ast.TypeInfoKind.Method then
            symbolName = concatGLSym( symbolName, Ast.isPubToExternal( typeInfo:get_accessMode() ) )
         else
          
            if isInnerDeclType( typeInfo ) then
               if not isClosure( typeInfo ) then
                  local parentName = typeInfo:getParentFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager(), true )
                  symbolName = string.format( "%s_%s_%d_", parentName:gsub( "[%.@]", "_" ), symbolName, typeInfo:get_typeId())
               end
               
            end
            
            symbolName = concatSymWithType( symbolName, typeInfo )
         end
         
      elseif _matchExp[1] == SymbolKind.Class[1] then
         local typeInfo = _matchExp[2][1]
      
         local workName
         
         if isInnerDeclType( typeInfo ) then
            workName = string.format( "%s%d", name, typeInfo:get_typeId())
         else
          
            workName = symbolName
         end
         
         symbolName = concatSymWithType( workName, typeInfo )
      elseif _matchExp[1] == SymbolKind.Normal[1] then
      
      end
   end
   
   return symbolName
end


function convFilter:getClassSymbol( typeInfo )

   local orgType = typeInfo:get_srcTypeInfo():get_nonnilableType()
   return self:getSymbol( _lune.newAlge( SymbolKind.Class, {orgType}), orgType:get_rawTxt() )
end


function convFilter:getConstrSymbol( typeInfo )

   return string.format( "Init%s", self:getClassSymbol( typeInfo ))
end


function convFilter:getMethodSymbol( typeInfo )

   return self:getSymbol( _lune.newAlge( SymbolKind.Func, {typeInfo}), typeInfo:get_rawTxt() )
end


function convFilter:outputSymbol( kind, name )

   self:write( self:getSymbol( kind, name ) )
end

local function str2gostr( txt )

   local work = txt
   if string.find( work, '^```' ) then
      work = (string.format( "%q", work:sub( 4, -4 )) ):gsub( "\\\n", "\\n" )
   elseif string.find( work, "^'" ) then
      work = string.format( '"%s"', ((string.format( "%s", work:sub( 2, -2 )) ):gsub( '"', '\\"' ) ))
   end
   
   work = work:gsub( "\\9", "\\t" )
   return work
end

local type2gotypeMap = {[Ast.builtinTypeInt] = "LnsInt", [Ast.builtinTypeReal] = "LnsReal", [Ast.builtinTypeStem] = "LnsAny", [Ast.builtinTypeString] = "string"}

local function getOrgTypeInfo( typeInfo )

   do
      local enumType = _lune.__Cast( typeInfo:get_srcTypeInfo():get_nonnilableType(), 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         return enumType:get_valTypeInfo()
      end
   end
   
   return typeInfo:get_srcTypeInfo():get_nonnilableType()
end

function convFilter:type2gotype( typeInfo )

   if typeInfo:get_kind() == Ast.TypeInfoKind.DDD then
      return "[]LnsAny"
   end
   
   if isAnyType( typeInfo ) then
      return "LnsAny"
   end
   
   local orgType = getOrgTypeInfo( typeInfo )
   do
      local goType = type2gotypeMap[orgType]
      if goType ~= nil then
         return goType
      end
   end
   
   do
      local _switchExp = orgType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List then
         return "*LnsList"
      elseif _switchExp == Ast.TypeInfoKind.Map then
         return "map[LnsAny]LnsAny"
      elseif _switchExp == Ast.TypeInfoKind.Class then
         return "*" .. self:getClassSymbol( typeInfo )
      elseif _switchExp == Ast.TypeInfoKind.IF then
         return self:getClassSymbol( typeInfo )
      end
   end
   
   Util.err( string.format( "not support yet -- %s", typeInfo:getTxt(  )) )
end


local function getExpType( expListNode, index )
   local __func__ = '@lune.@base.@convGo.getExpType'

   local list = expListNode:get_expTypeList()
   if #list >= index then
      return list[index]
   end
   
   Util.err( string.format( "not support yet -- %s, %d: %d", __func__, expListNode:get_pos().lineNo, index) )
end

function convFilter:outputConv( srcType, dstType )

   if not isAnyType( dstType ) then
      self:write( string.format( ".(%s)", self:type2gotype( dstType )) )
   end
   
end





function convFilter:processBlankLine( node, opt )

end



function convFilter:processNone( node, opt )

end



function convFilter:processImport( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processImport'

   Util.err( string.format( "not support -- %s", __func__) )
end

local ExpListKind = {}
ExpListKind._val2NameMap = {}
function ExpListKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ExpListKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ExpListKind._from( val )
   if ExpListKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ExpListKind.__allList = {}
function ExpListKind.get__allList()
   return ExpListKind.__allList
end

ExpListKind.Direct = 0
ExpListKind._val2NameMap[0] = 'Direct'
ExpListKind.__allList[1] = ExpListKind.Direct
ExpListKind.Slice = 1
ExpListKind._val2NameMap[1] = 'Slice'
ExpListKind.__allList[2] = ExpListKind.Slice
ExpListKind.Conv = 2
ExpListKind._val2NameMap[2] = 'Conv'
ExpListKind.__allList[3] = ExpListKind.Conv


local function getExpListKind( dstTypeList, node )

   local lastExp = node:get_expList()[#node:get_expList()]
   local hasAbbr
   
   if lastExp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
      hasAbbr = true
      if #node:get_expList() < 2 then
         return ExpListKind.Direct
      end
      
      lastExp = node:get_expList()[#node:get_expList() - 1]
   else
    
      hasAbbr = false
   end
   
   if _lune.__Cast( lastExp, 3, Nodes.ExpToDDDNode ) then
      local mRetExp = node:get_mRetExp()
      if  nil == mRetExp then
         local _mRetExp = mRetExp
      
         return ExpListKind.Slice
      end
      
      if mRetExp:get_index() == 1 and dstTypeList[mRetExp:get_index()]:get_kind() == Ast.TypeInfoKind.DDD then
         return ExpListKind.Slice
      end
      
      return ExpListKind.Conv
   end
   
   if lastExp:get_expType():get_kind() == Ast.TypeInfoKind.DDD then
      local mRetExp = node:get_mRetExp()
      if  nil == mRetExp then
         local _mRetExp = mRetExp
      
         return ExpListKind.Slice
      end
      
      if mRetExp:get_index() == 1 and dstTypeList[mRetExp:get_index()]:get_kind() == Ast.TypeInfoKind.DDD then
         return ExpListKind.Direct
      end
      
   else
    
      local mRetExp = node:get_mRetExp()
      if  nil == mRetExp then
         local _mRetExp = mRetExp
      
         return ExpListKind.Direct
      end
      
      if not hasAbbr and mRetExp:get_index() == 1 then
         return ExpListKind.Direct
      end
      
   end
   
   return ExpListKind.Conv
end

local function getConvExpName( nodeId, argListNode )

   return string.format( "_convExp%d_%d", nodeId, argListNode:get_id())
end

function convFilter:processConvExp( nodeId, dstTypeList, argListNode )

   local argList = argListNode
   if  nil == argList then
      local _argList = argList
   
      return 
   end
   
   
   if getExpListKind( dstTypeList, argList ) ~= ExpListKind.Conv then
      return 
   end
   
   
   
   self:write( string.format( "func %s(", getConvExpName( nodeId, argList )) )
   local mRetIndex = _lune.nilacc( argList:get_mRetExp(), 'get_index', 'callmtd' )
   if  nil == mRetIndex then
      local _mRetIndex = mRetIndex
   
      return 
   end
   
   for index, argExp in ipairs( argList:get_expList() ) do
      do
         local exp2ddd = _lune.__Cast( argExp, 3, Nodes.ExpToDDDNode )
         if exp2ddd ~= nil then
            for __index, exp in ipairs( exp2ddd:get_expList():get_expList() ) do
               if index ~= 1 then
                  self:write( ", " )
               end
               
               self:write( string.format( "arg%d ", index) )
               self:write( self:type2gotype( exp:get_expType() ) )
            end
            
         else
            if index ~= 1 then
               self:write( ", " )
            end
            
            if mRetIndex == index then
               self:write( string.format( "arg%d []LnsAny", index) )
               break
            else
             
               self:write( string.format( "arg%d ", index) )
               self:write( self:type2gotype( argExp:get_expType() ) )
            end
            
         end
      end
      
   end
   
   self:write( ")" )
   if #dstTypeList >= 2 then
      self:write( ")" )
      for index, argType in ipairs( dstTypeList ) do
         if index ~= 1 then
            self:write( ", " )
         end
         
         self:write( self:type2gotype( argType ) )
      end
      
      self:writeln( ") {" )
   elseif #dstTypeList == 1 then
      self:writeln( string.format( " %s {", self:type2gotype( dstTypeList[1] )) )
   else
    
      self:writeln( "{" )
   end
   
   self:write( "return " )
   
   local restIndex = nil
   for index, dstType in ipairs( dstTypeList ) do
      if index ~= 1 then
         self:write( ", " )
      end
      
      if dstType:get_kind() == Ast.TypeInfoKind.DDD then
         restIndex = index
         break
      end
      
      if index >= mRetIndex then
         self:write( string.format( "arg%d[%d].(%s)", mRetIndex, index - mRetIndex, self:type2gotype( dstType )) )
      else
       
         self:write( string.format( "arg%d", index) )
      end
      
   end
   
   if restIndex ~= nil then
      self:write( "[]LnsAny{ " )
      for index, _5313 in ipairs( argList:get_expList() ) do
         if index >= #dstTypeList then
            self:write( string.format( "arg%d", index) )
         end
         
      end
      
      self:writeln( "}" )
   else
      self:writeln( "" )
   end
   
   
   self:writeln( "}" )
end


function convFilter:processRoot( node, opt )

   Ast.pushProcessInfo( node:get_processInfo() )
   
   self:writeln( "package main" )
   
   for __index, workNode in ipairs( node:get_nodeManager():getExpCallNodeList(  ) ) do
      self:processConvExp( workNode:get_id(), workNode:get_func():get_expType():get_argTypeInfoList(), workNode:get_argList() )
   end
   
   for __index, workNode in ipairs( node:get_nodeManager():getDeclVarNodeList(  ) ) do
      self:processConvExp( workNode:get_id(), workNode:get_typeInfoList(), workNode:get_expList() )
   end
   
   
   self:pushProcessMode( ProcessMode.NonClosureFuncDecl )
   for __index, declFuncNode in ipairs( node:get_nodeManager():getDeclFuncNodeList(  ) ) do
      filter( declFuncNode, self, node )
      self:writeln( "" )
   end
   
   self:popProcessMode(  )
   
   for __index, declFuncNode in ipairs( node:get_nodeManager():getDeclClassNodeList(  ) ) do
      filter( declFuncNode, self, node )
      self:writeln( "" )
   end
   
   
   self:writeln( "func Lns_init() {" )
   self:pushIndent(  )
   
   for __index, child in ipairs( node:get_children() ) do
      
      do
         local _switchExp = child:get_kind()
         if _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_DeclClass() or _switchExp == Nodes.NodeKind.get_DeclMacro() or _switchExp == Nodes.NodeKind.get_TestBlock() then
            
         else 
            
               filter( child, self, node )
               self:writeln( "" )
         end
      end
      
   end
   
   
   self:popIndent(  )
   self:writeln( "}" )
   
   Ast.popProcessInfo(  )
end


function convFilter:processSubfile( node, opt )

   
end


function convFilter:processBlockSub( node, opt )

   if node:get_blockKind() == Nodes.BlockKind.Block then
      self:writeln( "{" )
   end
   
   
   self:pushProcessMode( ProcessMode.Main )
   self:pushIndent(  )
   for __index, child in ipairs( node:get_stmtList() ) do
      
      do
         local _switchExp = child:get_kind()
         if _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_DeclClass() or _switchExp == Nodes.NodeKind.get_DeclMacro() or _switchExp == Nodes.NodeKind.get_TestBlock() then
            
         else 
            
               filter( child, self, node )
               self:writeln( "" )
         end
      end
      
   end
   
   self:popIndent(  )
   self:popProcessMode(  )
   
   if node:get_blockKind() == Nodes.BlockKind.Block then
      self:write( "}" )
   end
   
end


function convFilter:expList2Slice( subList )

   local mRetIndex = _lune.nilacc( subList:get_mRetExp(), 'get_index', 'callmtd' )
   if mRetIndex and mRetIndex == 1 then
      local subExp = subList:get_expList()[1]
      if subExp:get_expType():get_kind() ~= Ast.TypeInfoKind.DDD then
         self:write( "Lns_2DDD(" )
         filter( subExp, self, subList )
         self:write( ")" )
      else
       
         filter( subExp, self, subList )
      end
      
   else
    
      if mRetIndex and mRetIndex ~= 1 then
         self:write( "append( " )
      end
      
      self:write( "[]LnsAny{" )
      for subIndex, subExp in ipairs( subList:get_expList() ) do
         if mRetIndex == subIndex then
            if mRetIndex ~= 1 then
               self:write( "}, " )
            end
            
            if subExp:get_expType():get_kind() ~= Ast.TypeInfoKind.DDD then
               self:write( "Lns_2DDD(" )
               filter( subExp, self, subList )
               self:write( ")" )
            else
             
               filter( subExp, self, subList )
            end
            
            self:write( "..." )
            break
         end
         
         if subIndex ~= 1 then
            self:write( ', ' )
         end
         
         filter( subExp, self, subList )
      end
      
      if mRetIndex and mRetIndex ~= 1 then
         self:write( ")" )
      else
       
         self:write( "}" )
      end
      
   end
   
end

function convFilter:processSetFromExpList( convArgFuncName, dstTypeList, expListNode )

   do
      local _switchExp = getExpListKind( dstTypeList, expListNode )
      if _switchExp == ExpListKind.Conv then
         self:write( string.format( "%s(", convArgFuncName) )
         local mRetIndex = _lune.nilacc( expListNode:get_mRetExp(), 'get_index', 'callmtd' )
         
         for index, exp in ipairs( expListNode:get_expList() ) do
            if index ~= 1 then
               self:write( ', ' )
            end
            
            if mRetIndex == index then
               self:write( "Lns_2DDD(" )
               filter( exp, self, expListNode )
               self:write( ")" )
               break
            end
            
            filter( exp, self, expListNode )
            if _lune.nilacc( expListNode:get_mRetExp(), 'get_index', 'callmtd' ) == index then
               break
            end
            
         end
         
         self:write( ")" )
      elseif _switchExp == ExpListKind.Slice then
         for index, argType in ipairs( dstTypeList ) do
            if index ~= 1 then
               self:write( ', ' )
            end
            
            if #expListNode:get_expList() >= index then
               local argExp = expListNode:get_expList()[index]
               
               do
                  local exp2ddd = _lune.__Cast( argExp, 3, Nodes.ExpToDDDNode )
                  if exp2ddd ~= nil then
                     self:expList2Slice( exp2ddd:get_expList() )
                  else
                     if argExp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
                        if argType:get_kind() == Ast.TypeInfoKind.DDD then
                           self:write( "[]LnsAny{}" )
                        else
                         
                           self:write( "nil" )
                        end
                        
                     else
                      
                        filter( argExp, self, expListNode )
                     end
                     
                  end
               end
               
            else
             
               self:write( "[]LnsAny{}" )
            end
            
         end
         
      elseif _switchExp == ExpListKind.Direct then
         local mRetIndex = _lune.nilacc( expListNode:get_mRetExp(), 'get_index', 'callmtd' )
         for index, funcArgType in ipairs( dstTypeList ) do
            if mRetIndex == index - 1 then
               break
            end
            
            if index ~= 1 then
               self:write( ', ' )
            end
            
            if index == #dstTypeList and funcArgType:get_kind() == Ast.TypeInfoKind.DDD then
               if #expListNode:get_expList() < index or expListNode:get_expList()[index]:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
                  self:write( "[]LnsAny{}" )
               else
                
                  filter( expListNode:get_expList()[index], self, expListNode )
               end
               
            else
             
               if #expListNode:get_expList() < index or expListNode:get_expList()[index]:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
                  self:write( "nil" )
               else
                
                  filter( expListNode:get_expList()[index], self, expListNode )
               end
               
            end
            
         end
         
      end
   end
   
end


function convFilter:processStmtExp( node, opt )

   filter( node:get_exp(), self, node )
end


function convFilter:processDeclEnum( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclEnum'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclAlge( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclAlge'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processNewAlgeVal( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processNewAlgeVal'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclMember( node, opt )

   self:outputSymbol( _lune.newAlge( SymbolKind.Member, {Ast.isPubToExternal( node:get_accessMode() )}), node:get_name().txt )
   self:write( string.format( " %s", self:type2gotype( node:get_refType():get_expType() )) )
end


function convFilter:processExpMacroExp( node, opt )

   for __index, stmt in ipairs( node:get_stmtList() ) do
      filter( stmt, self, node )
      self:writeln( "" )
   end
   
end


function convFilter:processDeclMacro( node, opt )

   
end


function convFilter:processExpMacroStat( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpMacroStat'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:outputDeclFuncArg( funcType )

   for index, argType in ipairs( funcType:get_argTypeInfoList() ) do
      if index ~= 1 then
         self:write( ", " )
      end
      
      self:write( string.format( "arg%d ", index) )
      self:write( self:type2gotype( argType ) )
   end
   
end


function convFilter:processDeclConstr( node, opt )

   local classType = node:get_expType():get_parentInfo()
   local className = self:getClassSymbol( classType )
   self:write( string.format( "func (self *%s) %s(", className, self:getConstrSymbol( classType )) )
   for index, arg in ipairs( node:get_declInfo():get_argList() ) do
      if index ~= 1 then
         self:write( "," )
      end
      
      filter( arg, self, node )
   end
   
   self:writeln( ") {" )
   
   filter( _lune.unwrap( node:get_declInfo():get_body()), self, node )
   
   self:writeln( "}" )
end


function convFilter:processDeclDestr( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclDestr'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpCallSuper( node, opt )

   if node:get_methodType():get_rawTxt() == "__init" then
      self:write( string.format( "self.%s", self:getConstrSymbol( node:get_superType() )) )
   else
    
      self:write( string.format( "self.%s.%s", self:getClassSymbol( node:get_methodType():get_parentInfo() ), self:getMethodSymbol( node:get_methodType() )) )
   end
   
   self:write( "(" )
   do
      local argList = node:get_expList()
      if argList ~= nil then
         self:processSetFromExpList( getConvExpName( node:get_id(), argList ), node:get_methodType():get_argTypeInfoList(), argList )
      end
   end
   
   self:write( ")" )
end


function convFilter:outputRetType( retTypeList )

   do
      local _switchExp = #retTypeList
      if _switchExp == 0 then
         self:write( "" )
      elseif _switchExp == 1 then
         self:write( " " .. self:type2gotype( retTypeList[1] ) )
      else 
         
            self:write( "(" )
            for index, retType in ipairs( retTypeList ) do
               if index ~= 1 then
                  self:write( ", " )
               end
               
               self:write( self:type2gotype( retType ) )
            end
            
            self:write( ")" )
      end
   end
   
end


function convFilter:outputDeclFuncInfo( node, declInfo )

   
   if isClosure( node:get_expType() ) then
      do
         local name = declInfo:get_name()
         if name ~= nil then
            self:outputSymbol( _lune.newAlge( SymbolKind.Func, {node:get_expType()}), name.txt )
            self:write( " := " )
         end
      end
      
      self:write( "func" )
   else
    
      if declInfo:get_kind() == Nodes.FuncKind.Mtd then
         self:write( "func " )
         self:write( "(self *" )
         self:write( self:getClassSymbol( _lune.unwrap( declInfo:get_classTypeInfo()) ) )
         self:write( ") " )
      else
       
         self:write( "func " )
      end
      
      
      do
         local name = declInfo:get_name()
         if name ~= nil then
            self:outputSymbol( _lune.newAlge( SymbolKind.Func, {node:get_expType()}), name.txt )
         end
      end
      
   end
   
   
   self:write( "(" )
   
   for index, arg in ipairs( declInfo:get_argList() ) do
      if index ~= 1 then
         self:write( "," )
      end
      
      filter( arg, self, node )
   end
   
   self:write( ")" )
   
   self:outputRetType( declInfo:get_retTypeInfoList() )
   
   self:writeln( " {" )
   
   do
      local body = declInfo:get_body()
      if body ~= nil then
         filter( body, self, node )
      end
   end
   
   self:write( "}" )
end


function convFilter:processDeclMethod( node, opt )

   self:outputDeclFuncInfo( node, node:get_declInfo() )
end


function convFilter:processProtoMethod( node, opt )

   
end


function convFilter:processUnwrapSet( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processUnwrapSet'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processIfUnwrap( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   for index, varSym in ipairs( node:get_varSymList() ) do
      if index > 1 then
         self:write( ", " )
      end
      
      self:write( "_" .. varSym:get_name() )
   end
   
   self:write( " := " )
   self:processSetFromExpList( getConvExpName( node:get_id(), node:get_expList() ), node:get_expList():get_expTypeList(), node:get_expList() )
   self:writeln( "" )
   self:write( "if " )
   for index, varSym in ipairs( node:get_varSymList() ) do
      if index > 1 then
         self:write( " && " )
      end
      
      self:write( string.format( "_%s != nil", varSym:get_name()) )
   end
   
   self:writeln( " {" )
   self:pushIndent(  )
   for index, varSym in ipairs( node:get_varSymList() ) do
      self:write( string.format( "%s := _%s", varSym:get_name(), varSym:get_name()) )
      self:outputConv( getExpType( node:get_expList(), index ), varSym:get_typeInfo() )
      self:writeln( "" )
   end
   
   self:popIndent(  )
   
   filter( node:get_block(), self, node )
   
   do
      local nilBlock = node:get_nilBlock()
      if nilBlock ~= nil then
         self:writeln( "} else {" )
         
         filter( nilBlock, self, node )
         
         self:writeln( "}" )
      else
         self:writeln( "}" )
      end
   end
   
   
   self:popIndent(  )
   self:write( "}" )
end


function convFilter:processDeclVar( node, opt )

   self:write( "var" )
   
   local prevType = ""
   
   local function declVar(  )
   
      for index, varInfo in ipairs( node:get_varList() ) do
         local goType = self:type2gotype( varInfo:get_actualType() )
         if goType ~= prevType then
            self:write( string.format( " %s", prevType) )
         end
         
         if index ~= 1 then
            self:write( ", " )
         end
         
         self:write( string.format( "%s", varInfo:get_name().txt) )
         prevType = goType
      end
      
      self:write( string.format( " %s", prevType) )
   end
   
   if node:get_unwrapFlag() then
      do
         local expList, unwrapBlock = node:get_expList(), node:get_unwrapBlock()
         if expList ~= nil and unwrapBlock ~= nil then
            declVar(  )
            self:writeln( "" )
            self:writeln( "{" )
            self:pushIndent(  )
            for index, varInfo in ipairs( node:get_varList() ) do
               if index ~= 1 then
                  self:write( ", " )
               end
               
               self:write( string.format( "_%s", varInfo:get_name().txt) )
            end
            
            self:write( " := " )
            self:processSetFromExpList( getConvExpName( node:get_id(), expList ), node:get_typeInfoList(), expList )
            self:writeln( "" )
            self:write( "if " )
            for index, varInfo in ipairs( node:get_varList() ) do
               if index ~= 1 then
                  self:write( " || " )
               end
               
               self:write( string.format( "_%s == nil", varInfo:get_name().txt) )
            end
            
            self:writeln( "{" )
            
            filter( unwrapBlock, self, node )
            
            do
               local thenBlock = node:get_thenBlock()
               if thenBlock ~= nil then
                  self:writeln( "} else {" )
                  self:pushIndent(  )
                  for index, varInfo in ipairs( node:get_varList() ) do
                     self:write( string.format( "%s = _%s", varInfo:get_name().txt, varInfo:get_name().txt) )
                     self:outputConv( getExpType( expList, index ), varInfo:get_actualType() )
                     self:writeln( "" )
                  end
                  
                  self:popIndent(  )
                  filter( thenBlock, self, node )
                  self:writeln( "}" )
               else
                  self:writeln( "}" )
               end
            end
            
            
            self:popIndent(  )
            self:write( "}" )
         end
      end
      
   else
    
      do
         local expList = node:get_expList()
         if expList ~= nil then
            for index, exp in ipairs( expList:get_expList() ) do
               local goType
               
               local varName
               
               if #node:get_varList() >= index then
                  local varInfo = node:get_varList()[index]
                  goType = self:type2gotype( varInfo:get_actualType() )
                  varName = varInfo:get_name().txt
               else
                
                  goType = self:type2gotype( exp:get_expType() )
                  varName = "_"
               end
               
               if goType ~= prevType then
                  self:write( string.format( " %s", prevType) )
               end
               
               if index ~= 1 then
                  self:write( ", " )
               end
               
               self:write( varName )
               prevType = goType
            end
            
            self:write( string.format( " %s", prevType) )
            
            self:write( " = " )
            self:processSetFromExpList( getConvExpName( node:get_id(), expList ), node:get_typeInfoList(), expList )
         else
            declVar(  )
         end
      end
      
   end
   
end


function convFilter:processWhen( node, opt )

   self:write( "if " )
   for index, symPair in ipairs( node:get_symPairList() ) do
      if index > 1 then
         self:write( " && " )
      end
      
      self:write( string.format( "%s != nil", symPair:get_src():get_name()) )
      symPair:get_dst():set_convModuleParam( true )
   end
   
   self:writeln( "{" )
   self:pushIndent(  )
   for __index, symPair in ipairs( node:get_symPairList() ) do
      self:write( string.format( "%s_%d := %s", symPair:get_dst():get_name(), symPair:get_dst():get_symbolId(), symPair:get_src():get_name()) )
      self:outputConv( symPair:get_src():get_typeInfo(), symPair:get_dst():get_typeInfo() )
      self:writeln( "" )
      symPair:get_dst():set_convModuleParam( true )
   end
   
   self:popIndent(  )
   
   filter( node:get_block(), self, node )
   
   do
      local elseBlock = node:get_elseBlock()
      if elseBlock ~= nil then
         self:writeln( "} else {" )
         filter( elseBlock, self, node )
         self:write( "}" )
      else
         self:write( "}" )
      end
   end
   
end


function convFilter:processDeclArg( node, opt )

   self:write( string.format( "%s ", node:get_name().txt) )
   filter( _lune.unwrap( node:get_argType()), self, node )
end


function convFilter:processDeclArgDDD( node, opt )

   self:write( "ddd []LnsAny" )
end


function convFilter:processExpDDD( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpDDD'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpSubDDD( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpSubDDD'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclForm( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclForm'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclFunc( node, opt )

   if (self.processMode == ProcessMode.NonClosureFuncDecl ) == isClosure( node:get_expType() ) then
      return 
   end
   
   
   self:outputDeclFuncInfo( node, node:get_declInfo() )
end


function convFilter:processRefType( node, opt )

   self:write( self:type2gotype( node:get_expType() ) )
end


function convFilter:processIf( node, opt )

   for __index, stmt in ipairs( node:get_stmtList() ) do
      do
         local _switchExp = stmt:get_kind()
         if _switchExp == Nodes.IfKind.If then
            self:write( "if " )
            filter( stmt:get_exp(), self, node )
            self:writeln( "{" )
            filter( stmt:get_block(), self, node )
         elseif _switchExp == Nodes.IfKind.ElseIf then
            self:write( "} else if " )
            filter( stmt:get_exp(), self, node )
            self:writeln( "{" )
            filter( stmt:get_block(), self, node )
         elseif _switchExp == Nodes.IfKind.Else then
            self:writeln( "} else { " )
            filter( stmt:get_block(), self, node )
         end
      end
      
   end
   
   self:write( "}" )
end


function convFilter:processSwitch( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processSwitch'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processMatch( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processMatch'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processWhile( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processWhile'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processRepeat( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processRepeat'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processFor( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   local fromSym = string.format( "_from%d", node:get_id())
   local toSym = string.format( "_to%d", node:get_id())
   local deltaSym = string.format( "_delta%d", node:get_id())
   local workSym = string.format( "_work%d", node:get_id())
   self:write( string.format( "%s := ", fromSym) )
   filter( node:get_init(), self, node )
   self:writeln( "" )
   
   self:write( string.format( "%s := ", toSym) )
   filter( node:get_to(), self, node )
   self:writeln( "" )
   
   do
      local delta = node:get_delta()
      if delta ~= nil then
         self:write( string.format( "%s := ", deltaSym) )
         filter( delta, self, node )
         self:writeln( "" )
      else
         self:writeln( string.format( "var %s LnsInt", deltaSym) )
         self:writeln( string.format( "if %s <= %s { %s = 1 } else { %s = -1 }", fromSym, toSym, deltaSym, deltaSym) )
      end
   end
   
   
   self:writeln( string.format( "for %s := %s; %s <= %s; %s += %s {", workSym, fromSym, workSym, toSym, workSym, deltaSym) )
   
   self:pushIndent(  )
   self:writeln( string.format( "%s := %s", node:get_val():get_name(), workSym) )
   self:popIndent(  )
   
   filter( node:get_block(), self, node )
   
   self:writeln( "}" )
   
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processApply( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processApply'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processForeach( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processForeach'

   
   
   
   self:write( "for " )
   local loopExpType = node:get_exp():get_expType()
   do
      local _switchExp = loopExpType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List then
         do
            local key = node:get_key()
            if key ~= nil then
               if key:get_name() == "_" then
                  self:write( "_, " )
               else
                
                  self:write( string.format( "_%s, ", key:get_name()) )
               end
               
            else
               self:write( "_, " )
            end
         end
         
         
         local valName = node:get_val():get_name()
         local itemType = loopExpType:get_itemTypeInfoList()[1]
         
         if valName ~= "_" then
            self:write( string.format( "_%s", valName) )
         else
          
            self:write( string.format( "%s", valName) )
         end
         
         
         self:write( " := range( " )
         filter( node:get_exp(), self, node )
         self:writeln( ".Items ) {" )
         self:pushIndent(  )
         do
            local key = node:get_key()
            if key ~= nil then
               self:writeln( string.format( "%s := _%s + 1", key:get_name(), key:get_name()) )
            end
         end
         
         
         if valName ~= "_" then
            self:write( string.format( "%s := _%s", valName, valName) )
            self:outputConv( Ast.builtinTypeStem_, itemType )
            self:writeln( "" )
         end
         
         
         self:popIndent(  )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         local keyType = loopExpType:get_itemTypeInfoList()[1]
         do
            local key = node:get_key()
            if key ~= nil then
               
               if key:get_name() ~= "_" then
                  self:write( string.format( "_%s", key:get_name()) )
               else
                
                  self:write( string.format( "%s", key:get_name()) )
               end
               
               
               self:write( ", " )
            else
               self:write( "_, " )
            end
         end
         
         
         local valName = node:get_val():get_name()
         local itemType = loopExpType:get_itemTypeInfoList()[2]
         
         if valName ~= "_" then
            self:write( string.format( "_%s", valName) )
         else
          
            self:write( string.format( "%s", valName) )
         end
         
         
         self:write( " := range( " )
         filter( node:get_exp(), self, node )
         self:writeln( ") {" )
         self:pushIndent(  )
         do
            local key = node:get_key()
            if key ~= nil then
               
               if key:get_name() ~= "_" then
                  self:write( string.format( "%s := _%s", key:get_name(), key:get_name()) )
                  self:outputConv( Ast.builtinTypeStem_, keyType )
                  self:writeln( "" )
               end
               
               
            end
         end
         
         
         if valName ~= "_" then
            self:write( string.format( "%s := _%s", valName, valName) )
            self:outputConv( Ast.builtinTypeStem_, itemType )
            self:writeln( "" )
         end
         
         
         self:popIndent(  )
      else 
         
            Util.err( string.format( "not support -- %s", __func__) )
      end
   end
   
   filter( node:get_block(), self, node )
   self:write( "}" )
end


function convFilter:processForsort( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processForsort'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpUnwrap( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpUnwrap'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpToDDD( node, opt )

   do
      local mRetExp = node:get_expList():get_mRetExp()
      if mRetExp ~= nil then
         filter( node:get_expList(), self, node )
      else
         self:write( "[]LnsAny{ " )
         filter( node:get_expList(), self, node )
         self:write( "}" )
      end
   end
   
end


function convFilter:processExpNew( node, opt )

   local className = self:getClassSymbol( node:get_expType() )
   self:write( string.format( "New%s(", className) )
   do
      local argList = node:get_argList()
      if argList ~= nil then
         local scope = _lune.unwrap( node:get_expType():get_scope())
         local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, Ast.ScopeAccess.Normal ))
         
         self:processSetFromExpList( getConvExpName( node:get_id(), argList ), initFuncType:get_argTypeInfoList(), argList )
      end
   end
   
   self:write( ")" )
end


function convFilter:outputIFMethods( node )

   self:pushIndent(  )
   
   local name2MtdType = {}
   local scope = _lune.unwrap( node:get_expType():get_scope())
   scope:filterTypeInfoField( true, scope, Ast.ScopeAccess.Normal, function ( symbolInfo )
   
      if symbolInfo:get_kind() == Ast.SymbolKind.Mtd and symbolInfo:get_name() ~= "__init" then
         name2MtdType[symbolInfo:get_name()] = symbolInfo:get_typeInfo()
      end
      
      return true
   end )
   do
      local __sorted = {}
      local __map = name2MtdType
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, name in ipairs( __sorted ) do
         local typeInfo = __map[ name ]
         do
            self:write( string.format( "%s(", self:getSymbol( _lune.newAlge( SymbolKind.Func, {typeInfo}), name )) )
            for index, argType in ipairs( typeInfo:get_argTypeInfoList() ) do
               if index ~= 1 then
                  self:write( ", " )
               end
               
               self:write( string.format( "arg%d %s", index, self:type2gotype( argType )) )
            end
            
            self:write( ")" )
            self:outputRetType( typeInfo:get_retTypeInfoList() )
            self:writeln( "" )
         end
      end
   end
   
   
   self:popIndent(  )
end


function convFilter:outputMethodIF( node )

   self:write( "type " )
   self:write( self:getClassSymbol( node:get_expType() ) )
   self:writeln( "Mtd interface {" )
   
   self:outputIFMethods( node )
   
   self:writeln( "}" )
end


function convFilter:outputInterfaceType( node )

   self:writeln( string.format( "type %s interface {", self:getClassSymbol( node:get_expType() )) )
   
   self:pushIndent(  )
   
   self:outputIFMethods( node )
   
   self:popIndent(  )
   
   self:writeln( "}" )
end


function convFilter:outputClassType( node )

   self:write( "type " )
   self:write( self:getClassSymbol( node:get_expType() ) )
   self:writeln( " struct {" )
   
   self:pushIndent(  )
   
   if node:get_expType():hasBase(  ) then
      local superClassType = node:get_expType():get_baseTypeInfo()
      self:writeln( self:getClassSymbol( superClassType ) )
   end
   
   
   for __index, memberNode in ipairs( node:get_memberList() ) do
      filter( memberNode, self, node )
      self:writeln( "" )
   end
   
   self:write( "FP " )
   self:write( self:getClassSymbol( node:get_expType() ) )
   self:writeln( "Mtd" )
   
   self:popIndent(  )
   
   self:writeln( "}" )
end


function convFilter:outputDownCast( node )

   local symbol = self:getClassSymbol( node:get_expType() )
   self:write( "type " )
   self:writeln( string.format( "%sDownCast interface {", symbol) )
   self:pushIndent(  )
   self:write( "To" )
   self:write( self:getClassSymbol( node:get_expType() ) )
   self:write( "() *" )
   self:write( self:getClassSymbol( node:get_expType() ) )
   self:writeln( "" )
   self:popIndent(  )
   self:writeln( "}" )
   
   self:writeln( string.format( "func %sDownCastF( obj LnsAny ) LnsAny {", symbol) )
   self:pushIndent(  )
   self:writeln( string.format( "work, ok := obj.(%sDownCast)", symbol) )
   self:writeln( string.format( "if ok { return work.To%s() }", symbol) )
   self:writeln( "return nil" )
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:outputCastReceiver( node )

   self:write( "func (obj *" )
   self:write( self:getClassSymbol( node:get_expType() ) )
   self:write( ") To" )
   self:write( self:getClassSymbol( node:get_expType() ) )
   self:write( "() *" )
   self:write( self:getClassSymbol( node:get_expType() ) )
   self:writeln( " {" )
   self:pushIndent(  )
   self:writeln( "return obj" )
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:outputConstructor( node )

   local scope = _lune.unwrap( node:get_expType():get_scope())
   local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, Ast.ScopeAccess.Normal ))
   
   local className = self:getClassSymbol( node:get_expType() )
   self:write( string.format( "func New%s(", className) )
   self:outputDeclFuncArg( initFuncType )
   self:writeln( string.format( ") *%s {", className) )
   self:pushIndent(  )
   self:writeln( string.format( "obj := &%s{}", className) )
   self:writeln( "obj.FP = obj" )
   
   do
      local workType = node:get_expType()
      while workType:hasBase(  ) do
         workType = workType:get_baseTypeInfo()
         
         local superName = self:getClassSymbol( workType )
         self:writeln( string.format( "obj.%s.FP = obj", superName) )
      end
      
   end
   
   local ctorName = self:getConstrSymbol( node:get_expType() )
   self:write( string.format( "obj.%s(", ctorName) )
   for index, _5711 in ipairs( initFuncType:get_argTypeInfoList() ) do
      if index ~= 1 then
         self:write( ", " )
      end
      
      self:write( string.format( "arg%d", index) )
   end
   
   self:writeln( ")" )
   self:writeln( "return obj" )
   self:popIndent(  )
   
   self:writeln( "}" )
   
   if not node:hasUserInit(  ) then
      self:write( string.format( "func (self *%s) %s(", className, ctorName) )
      
      self:outputDeclFuncArg( initFuncType )
      self:writeln( ") {" )
      self:pushIndent(  )
      
      local superArgNum
      
      if node:get_expType():hasBase(  ) then
         local superType = node:get_expType():get_baseTypeInfo()
         local baseScope = _lune.unwrap( superType:get_scope())
         local baseInitFuncType = _lune.unwrap( baseScope:getTypeInfoField( "__init", true, baseScope, Ast.ScopeAccess.Normal ))
         superArgNum = #baseInitFuncType:get_argTypeInfoList()
         local superName = self:getClassSymbol( superType )
         self:write( string.format( "self.%s.%s( ", superName, self:getConstrSymbol( superType )) )
         for index = 1, superArgNum do
            if index ~= 1 then
               self:write( "," )
            end
            
            self:write( string.format( "arg%d", index) )
         end
         
         self:writeln( ")" )
      else
       
         superArgNum = 0
      end
      
      for index, _5719 in ipairs( initFuncType:get_argTypeInfoList() ) do
         if superArgNum < index then
            local sIndex = index - superArgNum
            local memberNode = node:get_memberList()[sIndex]
            self:writeln( string.format( "self.%s = arg%d", self:getSymbol( _lune.newAlge( SymbolKind.Member, {Ast.isPubToExternal( memberNode:get_accessMode() )}), memberNode:get_name().txt ), index) )
         end
         
      end
      
      
      self:popIndent(  )
      self:writeln( "}" )
   end
   
end


function convFilter:processDeclClass( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclClass'

   do
      local _switchExp = node:get_expType():get_kind()
      if _switchExp == Ast.TypeInfoKind.Class then
         self:writeln( string.format( "// declaration Class -- %s", node:get_expType():get_rawTxt()) )
         self:outputMethodIF( node )
         self:outputClassType( node )
         self:outputDownCast( node )
         self:outputCastReceiver( node )
         self:outputConstructor( node )
         
         for __index, fieldNode in ipairs( node:get_fieldList() ) do
            do
               local methodNode = _lune.__Cast( fieldNode, 3, Nodes.DeclMemberNode )
               if methodNode ~= nil then
                  
               else
                  filter( fieldNode, self, node )
                  self:writeln( "" )
               end
            end
            
         end
         
      elseif _switchExp == Ast.TypeInfoKind.IF then
         self:outputInterfaceType( node )
      else 
         
            Util.err( string.format( "%s: not support -- %s", __func__, Ast.TypeInfoKind:_getTxt( node:get_expType():get_kind())
            ) )
      end
   end
   
end


function convFilter:processExpCall( node, opt )

   local funcType = node:get_func():get_expType()
   if Ast.isBuiltin( funcType:get_typeId() ) then
      local builtinFuncs = TransUnit.getBuiltinFunc(  )
      do
         local _switchExp = funcType
         if _switchExp == builtinFuncs.lns_print then
            self:write( "Lns_print" )
         else 
            
               filter( node:get_func(), self, node )
         end
      end
      
   else
    
      filter( node:get_func(), self, node )
   end
   
   self:write( "(" )
   
   do
      local argList = node:get_argList()
      if argList ~= nil then
         self:processSetFromExpList( getConvExpName( node:get_id(), argList ), funcType:get_argTypeInfoList(), argList )
      end
   end
   
   
   self:write( ")" )
end


function convFilter:processExpAccessMRet( node, opt )

   
end


function convFilter:processExpList( node, opt )

   for index, exp in ipairs( node:get_expList() ) do
      if index ~= 1 then
         self:write( ", " )
      end
      
      do
         local mRetExp = node:get_mRetExp()
         if mRetExp ~= nil then
            if mRetExp:get_index() == index then
               self:write( "Lns_2DDD(" )
               filter( exp, self, node )
               self:write( ")" )
               break
            end
            
         end
      end
      
      filter( exp, self, node )
   end
   
end


function convFilter:processExpOp1( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpOp1'

   do
      local _switchExp = node:get_op().txt
      if _switchExp == "~" then
         self:write( "^" )
         filter( node:get_exp(), self, node )
      elseif _switchExp == "+" or _switchExp == "-" then
         self:write( node:get_op().txt )
         filter( node:get_exp(), self, node )
      elseif _switchExp == "not" then
         self:write( "Lns_op_not(" )
         filter( node:get_exp(), self, node )
         self:write( ")" )
      else 
         
            Util.err( string.format( "%s: not support -- %s", __func__, node:get_op().txt) )
      end
   end
   
end


function convFilter:processExpMultiTo1( node, opt )

   self:write( "Lns_car(" )
   filter( node:get_exp(), self, node )
   self:write( ")" )
   self:outputConv( Ast.builtinTypeStem_, node:get_expType() )
end


function convFilter:processExpCast( node, opt )

   do
      local _switchExp = node:get_castKind()
      if _switchExp == Nodes.CastKind.Force then
         filter( node:get_exp(), self, node )
      elseif _switchExp == Nodes.CastKind.Implicit then
         do
            local _switchExp = node:get_castType():get_kind()
            if _switchExp == Ast.TypeInfoKind.Class then
               self:write( "&" )
               filter( node:get_exp(), self, node )
               self:write( string.format( ".%s", self:getClassSymbol( node:get_castType() )) )
            elseif _switchExp == Ast.TypeInfoKind.IF then
               filter( node:get_exp(), self, node )
            else 
               
                  filter( node:get_exp(), self, node )
                  if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Class then
                     self:write( ".FP" )
                  end
                  
            end
         end
         
      elseif _switchExp == Nodes.CastKind.Normal then
         local typeName = self:getClassSymbol( node:get_castType() )
         self:write( string.format( "%sDownCastF(", typeName) )
         filter( node:get_exp(), self, node )
         if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Class then
            self:write( ".FP" )
         end
         
         self:write( ")" )
      end
   end
   
end


function convFilter:processExpParen( node, opt )

   if #node:get_exp():get_expTypeList() >= 2 then
      self:write( "Lns_car(" )
      filter( node:get_exp(), self, node )
      self:write( ")" )
      self:outputConv( Ast.builtinTypeStem_, node:get_expType() )
   else
    
      self:write( "(" )
      filter( node:get_exp(), self, node )
      self:write( ")" )
   end
   
end


function convFilter:processExpSetVal( node, opt )

   filter( node:get_exp1(), self, node )
   self:write( " = " )
   filter( node:get_exp2(), self, node )
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
      self:writeln( "Lns_popVal( Lns_incStack() ||" )
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
    
      self:write( "Lns_setStackVal( " )
      filter( node:get_exp1(), self, node )
      self:write( ") " )
   end
   
   self:writeln( opCC )
   if isAndOr( node:get_exp2() ) then
      filter( node:get_exp2(), self, node )
   else
    
      self:write( "Lns_setStackVal( " )
      filter( node:get_exp2(), self, node )
      self:write( ") " )
   end
   
   
   if firstFlag then
      self:write( ")" )
      
      self:outputConv( Ast.builtinTypeStem_, node:get_expType() )
      self:popIndent(  )
   end
   
end


function convFilter:processExpOp2( node, opt )

   local opTxt = node:get_op().txt
   
   do
      local _switchExp = opTxt
      if _switchExp == "and" or _switchExp == "or" then
         self:processAndOr( node, opTxt, opt.node )
      elseif _switchExp == ".." then
         filter( node:get_exp1(), self, node )
         self:write( " + " )
         filter( node:get_exp2(), self, node )
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
                  filter( node:get_exp1(), self, node )
                  self:write( " " .. opTxt .. " " )
                  filter( node:get_exp2(), self, node )
               end
            end
            
      end
   end
   
end


function convFilter:processExpRef( node, opt )

   if node:get_expType():get_kind() == Ast.TypeInfoKind.DDD then
      self:write( "ddd" )
   else
    
      do
         local param = node:get_symbolInfo():get_convModuleParam()
         if param ~= nil then
            self:write( string.format( "%s_%d", node:get_symbolInfo():get_name(), node:get_symbolInfo():get_symbolId()) )
         else
            do
               local _switchExp = node:get_symbolInfo():get_kind()
               if _switchExp == Ast.SymbolKind.Var then
                  if Ast.isPubToExternal( node:get_symbolInfo():get_accessMode() ) then
                     self:write( self:getSymbol( _lune.newAlge( SymbolKind.PubVar), node:get_symbolInfo():get_name() ) )
                  else
                   
                     self:write( self:getSymbol( _lune.newAlge( SymbolKind.Normal), node:get_symbolInfo():get_name() ) )
                  end
                  
               elseif _switchExp == Ast.SymbolKind.Fun then
                  self:write( self:getSymbol( _lune.newAlge( SymbolKind.Func, {node:get_expType()}), node:get_symbolInfo():get_name() ) )
               elseif _switchExp == Ast.SymbolKind.Typ then
                  self:write( self:getClassSymbol( node:get_expType() ) )
               else 
                  
                     self:write( node:get_symbolInfo():get_name() )
               end
            end
            
         end
      end
      
   end
   
end


function convFilter:processExpRefItem( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpRefItem'

   do
      local _switchExp = node:get_val():get_expType():get_kind()
      if _switchExp == Ast.TypeInfoKind.List then
         filter( node:get_val(), self, node )
         self:write( ".GetAt(" )
         filter( _lune.unwrap( node:get_index()), self, node )
         self:write( ")" )
         if not isAnyType( node:get_expType() ) then
            self:write( string.format( ".(%s)", self:type2gotype( node:get_expType() )) )
         end
         
      else 
         
            Util.err( string.format( "not support -- %s", __func__) )
      end
   end
   
end


function convFilter:processRefField( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processRefField'

   filter( node:get_prefix(), self, node )
   
   if node:get_expType():get_kind() == Ast.TypeInfoKind.Method and Ast.isBuiltin( node:get_expType():get_typeId() ) then
      local builtinFuncs = TransUnit.getBuiltinFunc(  )
      do
         local _switchExp = node:get_expType()
         if _switchExp == builtinFuncs.list_insert then
            self:write( ".Insert" )
         elseif _switchExp == builtinFuncs.list_remove then
            self:write( ".Remove" )
         else 
            
               Util.err( string.format( "%s: not support -- %s", __func__, node:get_expType():getTxt(  )) )
         end
      end
      
   else
    
      self:write( "." )
      do
         local symbol = node:get_symbolInfo()
         if symbol ~= nil then
            do
               local _switchExp = node:get_expType():get_kind()
               if _switchExp == Ast.TypeInfoKind.Method then
                  self:write( "FP." )
                  self:write( self:getMethodSymbol( node:get_expType() ) )
               elseif _switchExp == Ast.TypeInfoKind.Class then
                  self:write( "FP." )
                  self:write( self:getClassSymbol( node:get_expType() ) )
               else 
                  
                     self:outputSymbol( _lune.newAlge( SymbolKind.Member, {Ast.isPubToExternal( symbol:get_accessMode() )}), node:get_field().txt )
               end
            end
            
         end
      end
      
   end
   
end


function convFilter:processExpOmitEnum( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpOmitEnum'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processGetField( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processGetField'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processReturn( node, opt )

   self:write( "return " )
   do
      local expList = node:get_expList()
      if expList ~= nil then
         filter( expList, self, node )
      end
   end
   
end


function convFilter:processTestBlock( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processTestBlock'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processProvide( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processProvide'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processAlias( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processAlias'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processBoxing( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processBoxing'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processUnboxing( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processUnboxing'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processLiteralList( node, opt )

   self:write( "NewLnsList(" )
   do
      local expList = node:get_expList()
      if expList ~= nil then
         self:expList2Slice( expList )
      else
         self:write( "[]LnsAny{}" )
      end
   end
   
   self:write( ")" )
end


function convFilter:processLiteralSet( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processLiteralSet'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processLiteralMap( node, opt )

   self:write( "map[LnsAny]LnsAny{" )
   for __index, pair in ipairs( node:get_pairList() ) do
      filter( pair:get_key(), self, node )
      self:write( ":" )
      filter( pair:get_val(), self, node )
      self:write( "," )
   end
   
   self:write( "}" )
end


function convFilter:processLiteralArray( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processLiteralArray'

   Util.err( string.format( "not support -- %s", __func__) )
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
   self:write( string.format( '%s', str2gostr( txt )) )
end


function convFilter:processLiteralBool( node, opt )

   self:write( node:get_token().txt )
end


function convFilter:processLiteralNil( node, opt )

   self:write( "nil" )
end


function convFilter:processBreak( node, opt )

   self:write( "break" )
end


function convFilter:processLiteralSymbol( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processLiteralSymbol'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processAbbr( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processAbbr'

   Util.err( string.format( "not support -- %s", __func__) )
end


local function createFilter( enableTest, streamName, stream, ast )

   return convFilter.new(enableTest, streamName, stream, ast)
end
_moduleObj.createFilter = createFilter

return _moduleObj
