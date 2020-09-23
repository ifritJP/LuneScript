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

local MaxNilAccNum = 3

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

ProcessMode.DeclTopScopeVar = 0
ProcessMode._val2NameMap[0] = 'DeclTopScopeVar'
ProcessMode.__allList[1] = ProcessMode.DeclTopScopeVar
ProcessMode.DeclClass = 1
ProcessMode._val2NameMap[1] = 'DeclClass'
ProcessMode.__allList[2] = ProcessMode.DeclClass
ProcessMode.NonClosureFuncDecl = 2
ProcessMode._val2NameMap[2] = 'NonClosureFuncDecl'
ProcessMode.__allList[3] = ProcessMode.NonClosureFuncDecl
ProcessMode.Main = 3
ProcessMode._val2NameMap[3] = 'Main'
ProcessMode.__allList[4] = ProcessMode.Main


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
   self.moduleTypeInfo = ast:get_moduleTypeInfo()
   self.moduleScope = _lune.unwrap( ast:get_moduleTypeInfo():get_scope())
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
   if typeInfo:get_nilable() or work == Ast.builtinTypeStem then
      return true
   end
   
   do
      local _switchExp = typeInfo:get_kind()
      if _switchExp == Ast.TypeInfoKind.Stem or _switchExp == Ast.TypeInfoKind.Alge or _switchExp == Ast.TypeInfoKind.Alternate then
         return true
      end
   end
   
   return false
end

local function isClosure( typeInfo )

   local scope = typeInfo:get_scope()
   if  nil == scope then
      local _scope = scope
   
      return false
   end
   
   return #scope:get_closureSymList() > 0
end

local golanKeywordSet = {["func"] = true, ["select"] = true, ["defer"] = true, ["go"] = true, ["chan"] = true, ["package"] = true, ["const"] = true, ["fallthrough"] = true, ["range"] = true, ["continue"] = true, ["var"] = true, ["map"] = true}

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

SymbolKind.Func = { "Func", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
SymbolKind._name2Val["Func"] = SymbolKind.Func
SymbolKind.Member = { "Member", {{ func=_lune._toBool, nilable=false, child={} }}}
SymbolKind._name2Val["Member"] = SymbolKind.Member
SymbolKind.Normal = { "Normal"}
SymbolKind._name2Val["Normal"] = SymbolKind.Normal
SymbolKind.Static = { "Static", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
SymbolKind._name2Val["Static"] = SymbolKind.Static
SymbolKind.Type = { "Type", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
SymbolKind._name2Val["Type"] = SymbolKind.Type
SymbolKind.Var = { "Var", {{ func=Ast.SymbolInfo._fromMap, nilable=false, child={} }}}
SymbolKind._name2Val["Var"] = SymbolKind.Var


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
      symbolName = string.format( "_%s", name)
   else
    
      symbolName = name
   end
   
   
   do
      local _matchExp = kind
      if _matchExp[1] == SymbolKind.Var[1] then
         local symbolInfo = _matchExp[2][1]
      
         if symbolInfo:get_scope() == self.moduleScope then
            symbolName = concatGLSym( string.format( "%s_", (self.moduleTypeInfo:get_rawTxt():gsub( "@", "" ) )) .. symbolName, Ast.isPubToExternal( symbolInfo:get_accessMode() ) )
         end
         
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
         
      elseif _matchExp[1] == SymbolKind.Type[1] then
         local typeInfo = _matchExp[2][1]
      
         local workName
         
         if isInnerDeclType( typeInfo ) then
            workName = string.format( "%s%d", name, typeInfo:get_typeId())
         else
          
            workName = symbolName
         end
         
         symbolName = concatSymWithType( workName, typeInfo )
      elseif _matchExp[1] == SymbolKind.Static[1] then
         local typeInfo = _matchExp[2][1]
      
         local workName = self:getSymbol( _lune.newAlge( SymbolKind.Type, {typeInfo}), typeInfo:get_rawTxt() )
         symbolName = string.format( "%s__%s", workName, name)
      elseif _matchExp[1] == SymbolKind.Normal[1] then
      
      end
   end
   
   return symbolName
end


function convFilter:getTypeSymbol( typeInfo )

   local orgType = typeInfo:get_srcTypeInfo():get_nonnilableType()
   return self:getSymbol( _lune.newAlge( SymbolKind.Type, {orgType}), orgType:get_rawTxt() )
end


function convFilter:getConstrSymbol( typeInfo )

   return string.format( "Init%s", self:getTypeSymbol( typeInfo ))
end


function convFilter:getFuncSymbol( typeInfo )

   if typeInfo:get_kind() == Ast.TypeInfoKind.Method and typeInfo:get_staticFlag() then
      return self:getSymbol( _lune.newAlge( SymbolKind.Static, {typeInfo:get_parentInfo()}), typeInfo:get_rawTxt() )
   end
   
   return self:getSymbol( _lune.newAlge( SymbolKind.Func, {typeInfo}), typeInfo:get_rawTxt() == "" and "_anonymous" or typeInfo:get_rawTxt() )
end


function convFilter:getAlgeSymbol( valInfo )

   return self:getSymbol( _lune.newAlge( SymbolKind.Static, {valInfo:get_algeTpye()}), valInfo:get_name() )
end


function convFilter:getSymbolSym( symbolInfo )

   do
      local _switchExp = symbolInfo:get_kind()
      if _switchExp == Ast.SymbolKind.Fun or _switchExp == Ast.SymbolKind.Mtd then
         return self:getFuncSymbol( symbolInfo:get_typeInfo() )
      elseif _switchExp == Ast.SymbolKind.Var then
         return self:getSymbol( _lune.newAlge( SymbolKind.Var, {symbolInfo}), symbolInfo:get_name() )
      elseif _switchExp == Ast.SymbolKind.Mbr then
         if symbolInfo:get_staticFlag() then
            return self:getSymbol( _lune.newAlge( SymbolKind.Static, {symbolInfo:get_namespaceTypeInfo()}), symbolInfo:get_name() )
         end
         
         return self:getSymbol( _lune.newAlge( SymbolKind.Member, {Ast.isPubToExternal( symbolInfo:get_accessMode() )}), symbolInfo:get_name() )
      elseif _switchExp == Ast.SymbolKind.Arg then
         return symbolInfo:get_name()
      else 
         
            Util.err( string.format( "not support -- %s", Ast.SymbolKind:_getTxt( symbolInfo:get_kind())
            ) )
      end
   end
   
end


function convFilter:getAccessorSym( accessMode, name )

   return self:getSymbol( _lune.newAlge( SymbolKind.Member, {Ast.isPubToExternal( accessMode )}), name )
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
         return "*" .. self:getTypeSymbol( typeInfo )
      elseif _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.FormFunc then
         return self:getTypeSymbol( typeInfo )
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

function convFilter:getAny2Type( typeInfo )

   return string.format( ".(%s)", self:type2gotype( typeInfo ))
end


function convFilter:outputConv( srcType, dstType )

   if not isAnyType( dstType ) then
      self:write( self:getAny2Type( dstType ) )
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
      self:write( "(" )
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
         self:write( string.format( "Lns_getFromMulti( arg%d, %d )", mRetIndex, index - mRetIndex) )
         if not isAnyType( dstType ) then
            self:write( self:getAny2Type( dstType ) )
         end
         
      else
       
         self:write( string.format( "arg%d", index) )
      end
      
   end
   
   if restIndex ~= nil then
      self:write( "[]LnsAny{ " )
      for index, _5355 in ipairs( argList:get_expList() ) do
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


function convFilter:outputNilAccCall( node )

   if not node:hasNilAccess(  ) then
      return 
   end
   
   if #node:get_expTypeList() > MaxNilAccNum then
      local anys = "LnsAny"
      local nils = "nil"
      local lists = "list[0]"
      for count = 2, #node:get_expTypeList() do
         anys = string.format( "%s,LnsAny", anys)
         nils = string.format( "%s,nil", nils)
         lists = string.format( "%s,list[%d]", lists, count - 1)
      end
      
      local name = string.format( "%s_%d", self.moduleTypeInfo:get_rawTxt():gsub( "@", "" ), node:get_id())
      self:write( string.format( [==[
func lns_NilAccCall_%s( call func () (%s) ) bool {
    Lns_NilAccLast( Lns_2DDD( call() ) )
    return true
}
func lns_NilAccFinCall_%s( ret LnsAny ) (%s) {
    if Lns_IsNil( ret ) {
        return %s
    }
    list := ret.([]LnsAny)
    return %s
}
]==], name, anys, name, anys, nils, lists) )
   end
   
end


local function isRetGenerics( node )

   local funcType = node:get_func():get_expType()
   for index, retType in ipairs( funcType:get_retTypeInfoList() ) do
      if retType:get_kind() == Ast.TypeInfoKind.Alternate and not isAnyType( node:get_expTypeList()[index] ) then
         return true
      end
      
   end
   
   return false
end

function convFilter:processGenericsCall( node )

   if not isRetGenerics( node ) or #node:get_expTypeList() < 2 then
      return 
   end
   
   local srcTypeList = node:get_func():get_expType():get_retTypeInfoList()
   local dstTypeList = node:get_expTypeList()
   local srcTxt = string.format( "arg1 %s", self:type2gotype( srcTypeList[1] ))
   local dstTxt = string.format( "%s", self:type2gotype( dstTypeList[1] ))
   
   for index = 2, #srcTypeList do
      srcTxt = string.format( "%s,arg%d %s", srcTxt, index, self:type2gotype( srcTypeList[index] ))
   end
   
   for index = 2, #dstTypeList do
      dstTxt = string.format( "%s,%s", dstTxt, self:type2gotype( dstTypeList[index] ))
   end
   
   self:writeln( string.format( "func lns_convGenerics%d(%s) (%s) {", node:get_id(), srcTxt, dstTxt) )
   self:pushIndent(  )
   self:write( "return " )
   for index, dstType in ipairs( dstTypeList ) do
      if index > 1 then
         self:write( ", " )
      end
      
      if index > #srcTypeList then
         self:write( "nil" )
      else
       
         self:write( string.format( "arg%d", index) )
         local srcType = srcTypeList[index]
         if srcType:get_kind() == Ast.TypeInfoKind.Alternate then
            self:write( self:getAny2Type( dstType ) )
         end
         
      end
      
   end
   
   self:writeln( "" )
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processRoot( node, opt )

   Ast.pushProcessInfo( node:get_processInfo() )
   
   self:writeln( "package main" )
   
   self:pushProcessMode( ProcessMode.DeclTopScopeVar )
   for __index, child in ipairs( node:get_nodeManager():getDeclEnumNodeList(  ) ) do
      filter( child, self, node )
   end
   
   for __index, child in ipairs( node:get_children() ) do
      if child:get_kind() == Nodes.NodeKind.get_DeclVar() then
         filter( child, self, node )
      end
      
   end
   
   self:popProcessMode(  )
   
   for __index, workNode in ipairs( node:get_nodeManager():getDeclAlgeNodeList(  ) ) do
      filter( workNode, self, node )
   end
   
   for __index, workNode in ipairs( node:get_nodeManager():getDeclFormNodeList(  ) ) do
      filter( workNode, self, node )
   end
   
   for __index, workNode in ipairs( node:get_nodeManager():getExpCallNodeList(  ) ) do
      self:processGenericsCall( workNode )
      self:outputNilAccCall( workNode )
   end
   
   
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
   
   self:pushProcessMode( ProcessMode.DeclClass )
   for __index, declNode in ipairs( node:get_nodeManager():getDeclClassNodeList(  ) ) do
      filter( declNode, self, node )
      self:writeln( "" )
   end
   
   self:popProcessMode(  )
   
   self:writeln( "func Lns_init() {" )
   self:pushIndent(  )
   
   for __index, child in ipairs( node:get_children() ) do
      
      do
         local _switchExp = child:get_kind()
         if _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_DeclEnum() or _switchExp == Nodes.NodeKind.get_DeclMacro() or _switchExp == Nodes.NodeKind.get_DeclForm() or _switchExp == Nodes.NodeKind.get_TestBlock() then
            
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
         if _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_DeclClass() or _switchExp == Nodes.NodeKind.get_DeclForm() or _switchExp == Nodes.NodeKind.get_DeclMacro() or _switchExp == Nodes.NodeKind.get_TestBlock() then
            
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


function convFilter:processDeclAlge( node, opt )

   do
      local __sorted = {}
      local __map = node:get_algeType():get_valInfoMap()
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local valInfo = __map[ __key ]
         do
            local algeSym = self:getAlgeSymbol( valInfo )
            self:writeln( string.format( "type %s struct{", algeSym) )
            for index, paramType in ipairs( valInfo:get_typeList() ) do
               self:writeln( string.format( "Val%d %s", index, self:type2gotype( paramType )) )
            end
            
            self:writeln( "}" )
            if #valInfo:get_typeList() == 0 then
               self:writeln( string.format( "var %s_Obj = &%s{}", algeSym, algeSym) )
            end
            
            self:writeln( string.format( "func (self *%s) getTxt() string {", algeSym) )
            self:writeln( string.format( 'return "%s.%s"', node:get_algeType():get_rawTxt(), valInfo:get_name()) )
            self:writeln( "}" )
         end
      end
   end
   
end


function convFilter:processNewAlgeVal( node, opt )

   local algeSym = self:getAlgeSymbol( node:get_valInfo() )
   if #node:get_valInfo():get_typeList() == 0 then
      self:write( string.format( "%s_Obj", algeSym) )
   else
    
      self:write( string.format( "&%s{", algeSym) )
      for index, param in ipairs( node:get_paramList() ) do
         if index > 1 then
            self:write( ", " )
         end
         
         filter( param, self, node )
      end
      
      self:write( "}" )
   end
   
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
   local className = self:getTypeSymbol( classType )
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
    
      self:write( string.format( "self.%s.%s", self:getTypeSymbol( node:get_methodType():get_parentInfo() ), self:getFuncSymbol( node:get_methodType() )) )
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


local FuncInfo = {}
FuncInfo._name2Val = {}
function FuncInfo:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "FuncInfo.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function FuncInfo._from( val )
   return _lune._AlgeFrom( FuncInfo, val )
end

FuncInfo.DeclInfo = { "DeclInfo", {{ func=Nodes.Node._fromMap, nilable=false, child={} },{ func=Nodes.DeclFuncInfo._fromMap, nilable=false, child={} }}}
FuncInfo._name2Val["DeclInfo"] = FuncInfo.DeclInfo
FuncInfo.Type = { "Type", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
FuncInfo._name2Val["Type"] = FuncInfo.Type


function convFilter:outputDeclFunc( funcInfo )

   local typeInfo
   
   local name
   
   do
      local _matchExp = funcInfo
      if _matchExp[1] == FuncInfo.DeclInfo[1] then
         local node = _matchExp[2][1]
         local workDeclInfo = _matchExp[2][2]
      
         typeInfo = node:get_expType()
         if not workDeclInfo:get_name() then
            if self.processMode == ProcessMode.NonClosureFuncDecl then
               name = "_anonymous"
            else
             
               name = nil
            end
            
         else
          
            name = typeInfo:get_rawTxt()
         end
         
      elseif _matchExp[1] == FuncInfo.Type[1] then
         local workTypeInfo = _matchExp[2][1]
      
         typeInfo = workTypeInfo
         name = typeInfo:get_rawTxt()
      end
   end
   
   
   if isClosure( typeInfo ) then
      if name ~= nil then
         self:outputSymbol( _lune.newAlge( SymbolKind.Func, {typeInfo}), name )
         self:write( " := " )
      end
      
      self:write( "func" )
   else
    
      if typeInfo:get_kind() == Ast.TypeInfoKind.Method then
         self:write( "func " )
         self:write( "(self *" )
         self:write( self:getTypeSymbol( typeInfo:get_parentInfo() ) )
         self:write( ") " )
      else
       
         self:write( "func " )
      end
      
      
      if typeInfo:get_kind() ~= Ast.TypeInfoKind.FormFunc then
         if name ~= nil then
            self:outputSymbol( _lune.newAlge( SymbolKind.Func, {typeInfo}), name )
         end
         
      end
      
   end
   
   
   self:write( "(" )
   
   do
      local _matchExp = funcInfo
      if _matchExp[1] == FuncInfo.DeclInfo[1] then
         local node = _matchExp[2][1]
         local declInfo = _matchExp[2][2]
      
         for index, arg in ipairs( declInfo:get_argList() ) do
            if index ~= 1 then
               self:write( "," )
            end
            
            filter( arg, self, node )
         end
         
      elseif _matchExp[1] == FuncInfo.Type[1] then
         local _ = _matchExp[2][1]
      
         for index, argType in ipairs( typeInfo:get_argTypeInfoList() ) do
            if index ~= 1 then
               self:write( "," )
            end
            
            self:write( string.format( "arg%d %s", index, self:type2gotype( argType )) )
         end
         
      end
   end
   
   self:write( ")" )
   
   self:outputRetType( typeInfo:get_retTypeInfoList() )
end


function convFilter:outputDeclFuncInfo( node, declInfo )

   self:outputDeclFunc( _lune.newAlge( FuncInfo.DeclInfo, {node,declInfo}) )
   
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


function convFilter:getEnumGetTxtSym( enumType )

   return string.format( "%s_getTxt", self:getTypeSymbol( enumType ))
end


function convFilter:processDeclEnum( node, opt )

   do
      local _switchExp = self.processMode
      if _switchExp == ProcessMode.DeclTopScopeVar then
         self:writeln( string.format( "// decl enum -- %s ", node:get_enumType():getTxt(  )) )
         do
            local __sorted = {}
            local __map = node:get_enumType():get_name2EnumValInfo()
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, __key in ipairs( __sorted ) do
               local valInfo = __map[ __key ]
               do
                  self:write( string.format( "const %s = ", self:getSymbol( _lune.newAlge( SymbolKind.Static, {node:get_enumType()}), valInfo:get_name() )) )
                  do
                     local _matchExp = valInfo:get_val()
                     if _matchExp[1] == Ast.EnumLiteral.Int[1] then
                        local val = _matchExp[2][1]
                     
                        self:write( string.format( "%d", val) )
                     elseif _matchExp[1] == Ast.EnumLiteral.Real[1] then
                        local val = _matchExp[2][1]
                     
                        self:write( string.format( "%g", val) )
                     elseif _matchExp[1] == Ast.EnumLiteral.Str[1] then
                        local val = _matchExp[2][1]
                     
                        self:write( str2gostr( string.format( '"%s"', val) ) )
                     end
                  end
                  
                  self:writeln( "" )
               end
            end
         end
         
         local listName = string.format( "%sList_", self:getTypeSymbol( node:get_enumType() ))
         self:writeln( string.format( "var %s = NewLnsList( []LnsAny {", listName) )
         do
            local __sorted = {}
            local __map = node:get_enumType():get_name2EnumValInfo()
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, __key in ipairs( __sorted ) do
               local valInfo = __map[ __key ]
               do
                  self:writeln( string.format( "  %s,", self:getSymbol( _lune.newAlge( SymbolKind.Static, {node:get_enumType()}), valInfo:get_name() )) )
               end
            end
         end
         
         self:writeln( "})" )
         
         local scope = _lune.unwrap( node:get_enumType():get_scope())
         self:outputDeclFunc( _lune.newAlge( FuncInfo.Type, {_lune.unwrap( scope:getTypeInfoChild( "get__allList" ))}) )
         self:writeln( "{" )
         self:pushIndent(  )
         self:writeln( string.format( "return %s", listName) )
         self:popIndent(  )
         self:writeln( "}" )
         local mapName = string.format( "%sMap_", self:getTypeSymbol( node:get_enumType() ))
         self:writeln( string.format( "var %s = map[%s]string {", mapName, self:type2gotype( node:get_enumType():get_valTypeInfo() )) )
         do
            local __sorted = {}
            local __map = node:get_enumType():get_name2EnumValInfo()
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, __key in ipairs( __sorted ) do
               local valInfo = __map[ __key ]
               do
                  self:writeln( string.format( '  %s: "%s.%s",', self:getSymbol( _lune.newAlge( SymbolKind.Static, {node:get_enumType()}), valInfo:get_name() ), node:get_expType():get_rawTxt(), valInfo:get_name()) )
               end
            end
         end
         
         self:writeln( "}" )
         
         self:outputDeclFunc( _lune.newAlge( FuncInfo.Type, {_lune.unwrap( scope:getTypeInfoChild( "_from" ))}) )
         self:writeln( "{" )
         self:pushIndent(  )
         self:writeln( string.format( "if _, ok := %s[arg1]; ok { return arg1 }", mapName) )
         self:writeln( "return nil" )
         self:popIndent(  )
         self:writeln( "}" )
         
         self:writeln( "" )
         self:writeln( string.format( "func %s(arg1 %s) string {", self:getEnumGetTxtSym( node:get_enumType() ), self:type2gotype( node:get_enumType():get_valTypeInfo() )) )
         self:pushIndent(  )
         self:writeln( string.format( "return %s[arg1];", mapName) )
         self:popIndent(  )
         self:writeln( "}" )
      else 
         
      end
   end
   
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

   local needType
   
   if node:get_symbolInfoList()[1]:get_scope() ~= self.moduleScope then
      needType = true
      self:write( "var" )
   else
    
      if self.processMode == ProcessMode.DeclTopScopeVar then
         needType = true
      else
       
         needType = false
      end
      
   end
   
   
   local function declVar(  )
   
      local prevType = ""
      for index, symbolInfo in ipairs( node:get_symbolInfoList() ) do
         local goType = self:type2gotype( symbolInfo:get_typeInfo() )
         if needType and goType ~= prevType then
            self:write( string.format( " %s", prevType) )
         end
         
         if index ~= 1 then
            self:write( ", " )
         end
         
         self:write( self:getSymbolSym( symbolInfo ) )
         prevType = goType
      end
      
      if needType then
         self:write( string.format( " %s", prevType) )
      end
      
   end
   
   if self.processMode == ProcessMode.DeclTopScopeVar then
      self:write( "var" )
      declVar(  )
      self:writeln( "" )
      return 
   end
   
   
   if node:get_unwrapFlag() then
      do
         local expList, unwrapBlock = node:get_expList(), node:get_unwrapBlock()
         if expList ~= nil and unwrapBlock ~= nil then
            if node:get_symbolInfoList()[1]:get_scope() ~= self.moduleScope then
               declVar(  )
            end
            
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
                  for index, varInfo in ipairs( node:get_symbolInfoList() ) do
                     self:write( string.format( "%s = _%s", self:getSymbolSym( varInfo ), varInfo:get_name()) )
                     self:outputConv( getExpType( expList, index ), varInfo:get_typeInfo() )
                     self:writeln( "" )
                  end
                  
                  self:popIndent(  )
                  filter( thenBlock, self, node )
                  self:writeln( "}" )
               else
                  self:writeln( "} else {" )
                  self:pushIndent(  )
                  for index, varInfo in ipairs( node:get_symbolInfoList() ) do
                     self:write( string.format( "%s = _%s", self:getSymbolSym( varInfo ), varInfo:get_name()) )
                     self:outputConv( getExpType( expList, index ), varInfo:get_typeInfo() )
                     self:writeln( "" )
                  end
                  
                  self:popIndent(  )
                  self:writeln( "}" )
               end
            end
            
            
            self:popIndent(  )
            self:write( "}" )
         end
      end
      
   else
    
      declVar(  )
      do
         local expList = node:get_expList()
         if expList ~= nil then
            self:write( " = " )
            self:processSetFromExpList( getConvExpName( node:get_id(), expList ), node:get_typeInfoList(), expList )
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
      
      self:write( string.format( "%s != nil", self:getSymbolSym( symPair:get_src() )) )
      symPair:get_dst():set_convModuleParam( true )
   end
   
   self:writeln( "{" )
   self:pushIndent(  )
   for __index, symPair in ipairs( node:get_symPairList() ) do
      self:write( string.format( "%s_%d := %s", symPair:get_dst():get_name(), symPair:get_dst():get_symbolId(), self:getSymbolSym( symPair:get_src() )) )
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

   self:write( string.format( "type %s ", self:getTypeSymbol( node:get_expType() )) )
   self:outputDeclFunc( _lune.newAlge( FuncInfo.Type, {node:get_expType()}) )
   self:writeln( "" )
   
end


function convFilter:processDeclFunc( node, opt )

   if (self.processMode == ProcessMode.NonClosureFuncDecl ) == isClosure( node:get_expType() ) then
      if self.processMode ~= ProcessMode.NonClosureFuncDecl and node:get_expType():get_rawTxt() == "" then
         self:write( self:getFuncSymbol( node:get_expType() ) )
      end
      
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

   local valName = string.format( "_switch%d", node:get_id())
   self:write( string.format( "if %s := ", valName) )
   filter( node:get_exp(), self, node )
   self:write( "; " )
   
   for caseIndex, caseNode in ipairs( node:get_caseList() ) do
      if caseIndex ~= 1 then
         self:write( "} else if " )
      end
      
      for index, exp in ipairs( caseNode:get_expList():get_expList() ) do
         if index ~= 1 then
            self:write( " || " )
         end
         
         self:write( string.format( "%s == ", valName) )
         filter( exp, self, caseNode:get_expList() )
      end
      
      self:writeln( " {" )
      
      filter( caseNode:get_block(), self, node )
   end
   
   
   do
      local defaultBlock = node:get_default()
      if defaultBlock ~= nil then
         self:writeln( "} else {" )
         filter( defaultBlock, self, node )
      end
   end
   
   
   self:writeln( "}" )
end


function convFilter:processMatch( node, opt )

   local val = string.format( "_exp%d", node:get_id())
   self:write( string.format( "switch %s := ", val) )
   filter( node:get_val(), self, node )
   self:writeln( ".(type) {" )
   for __index, caseInfo in ipairs( node:get_caseList() ) do
      self:writeln( string.format( "case *%s:", self:getAlgeSymbol( caseInfo:get_valInfo() )) )
      for index, symbol in ipairs( caseInfo:get_valParamNameList() ) do
         if symbol:get_posForModToRef() then
            self:writeln( string.format( "%s := %s.Val%d", symbol:get_name(), val, index) )
         end
         
      end
      
      filter( caseInfo:get_block(), self, node )
   end
   
   do
      local defaultBlock = node:get_defaultBlock()
      if defaultBlock ~= nil then
         self:writeln( "default:" )
         filter( defaultBlock, self, node )
      end
   end
   
   self:writeln( "}" )
end


function convFilter:processWhile( node, opt )

   self:write( "for " )
   filter( node:get_exp(), self, node )
   self:writeln( " {" )
   
   filter( node:get_block(), self, node )
   
   self:writeln( "}" )
end


function convFilter:processRepeat( node, opt )

   self:writeln( "for {" )
   
   filter( node:get_block(), self, node )
   
   self:pushIndent(  )
   
   self:write( "if " )
   filter( node:get_exp(), self, node )
   self:writeln( "{ break }" )
   
   self:popIndent(  )
   
   self:writeln( "}" )
end


function convFilter:processFor( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   
   local fromSym = string.format( "_from%d", node:get_id())
   local toSym = string.format( "_to%d", node:get_id())
   local deltaSym = string.format( "_delta%d", node:get_id())
   local workSym = string.format( "_work%d", node:get_id())
   
   local valType = string.format( "%s", self:type2gotype( node:get_init():get_expType() ))
   self:write( string.format( "var %s %s = ", fromSym, valType) )
   filter( node:get_init(), self, node )
   self:writeln( "" )
   
   self:write( string.format( "var %s %s = ", toSym, valType) )
   filter( node:get_to(), self, node )
   self:writeln( "" )
   
   do
      local delta = node:get_delta()
      if delta ~= nil then
         self:writeln( string.format( "%s := %s", workSym, fromSym) )
         self:write( string.format( "%s := ", deltaSym) )
         filter( delta, self, node )
         self:writeln( "" )
         
         self:writeln( "for {" )
         
         self:pushIndent(  )
         self:writeln( string.format( "if %s > 0 {", deltaSym) )
         self:writeln( string.format( "   if %s > %s { break }", workSym, toSym) )
         self:writeln( "} else {" )
         self:writeln( string.format( "   if %s < %s { break }", workSym, toSym) )
         self:writeln( "}" )
         self:popIndent(  )
      else
         self:writeln( string.format( "for %s := %s; %s <= %s; %s++ {", workSym, fromSym, workSym, toSym, workSym) )
      end
   end
   
   
   self:pushIndent(  )
   self:writeln( string.format( "%s := %s", node:get_val():get_name(), workSym) )
   self:popIndent(  )
   
   filter( node:get_block(), self, node )
   
   if node:get_delta() then
      self:pushIndent(  )
      self:writeln( string.format( "%s += %s", workSym, deltaSym) )
      self:popIndent(  )
   end
   
   
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

   local className = self:getTypeSymbol( node:get_expType() )
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
   
      if symbolInfo:get_kind() == Ast.SymbolKind.Mtd and symbolInfo:get_name() ~= "__init" and not symbolInfo:get_staticFlag() then
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
   self:write( self:getTypeSymbol( node:get_expType() ) )
   self:writeln( "Mtd interface {" )
   
   self:outputIFMethods( node )
   
   self:writeln( "}" )
end


function convFilter:outputInterfaceType( node )

   self:writeln( string.format( "type %s interface {", self:getTypeSymbol( node:get_expType() )) )
   
   self:pushIndent(  )
   
   self:outputIFMethods( node )
   
   self:popIndent(  )
   
   self:writeln( "}" )
end


function convFilter:outputClassType( node )

   self:write( "type " )
   self:write( self:getTypeSymbol( node:get_expType() ) )
   self:writeln( " struct {" )
   
   self:pushIndent(  )
   
   if node:get_expType():hasBase(  ) then
      local superClassType = node:get_expType():get_baseTypeInfo()
      self:writeln( self:getTypeSymbol( superClassType ) )
   end
   
   
   for __index, memberNode in ipairs( node:get_memberList() ) do
      filter( memberNode, self, node )
      self:writeln( "" )
   end
   
   self:write( "FP " )
   self:write( self:getTypeSymbol( node:get_expType() ) )
   self:writeln( "Mtd" )
   
   self:popIndent(  )
   
   self:writeln( "}" )
end


function convFilter:outputDownCast( node )

   local symbol = self:getTypeSymbol( node:get_expType() )
   self:write( "type " )
   self:writeln( string.format( "%sDownCast interface {", symbol) )
   self:pushIndent(  )
   self:write( "To" )
   self:write( self:getTypeSymbol( node:get_expType() ) )
   self:write( "() *" )
   self:write( self:getTypeSymbol( node:get_expType() ) )
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
   self:write( self:getTypeSymbol( node:get_expType() ) )
   self:write( ") To" )
   self:write( self:getTypeSymbol( node:get_expType() ) )
   self:write( "() *" )
   self:write( self:getTypeSymbol( node:get_expType() ) )
   self:writeln( " {" )
   self:pushIndent(  )
   self:writeln( "return obj" )
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:outputConstructor( node )

   local scope = _lune.unwrap( node:get_expType():get_scope())
   local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, Ast.ScopeAccess.Normal ))
   
   local className = self:getTypeSymbol( node:get_expType() )
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
         
         local superName = self:getTypeSymbol( workType )
         self:writeln( string.format( "obj.%s.FP = obj", superName) )
      end
      
   end
   
   local ctorName = self:getConstrSymbol( node:get_expType() )
   self:write( string.format( "obj.%s(", ctorName) )
   for index, _5834 in ipairs( initFuncType:get_argTypeInfoList() ) do
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
         local superName = self:getTypeSymbol( superType )
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
      
      for index, _5842 in ipairs( initFuncType:get_argTypeInfoList() ) do
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


function convFilter:outputAccessor( node )

   local classType = node:get_expType()
   if classType:get_kind() == Ast.TypeInfoKind.IF then
      return 
   end
   
   
   for __index, memberNode in ipairs( node:get_memberList() ) do
      local memberNameToken = memberNode:get_name(  )
      local memberName = memberNameToken.txt
      local memberSym = memberNode:get_symbolInfo()
      
      if memberNode:get_getterMode(  ) ~= Ast.AccessMode.None then
         local getterName = string.format( "get_%s", memberName)
         local getterSym = _lune.unwrap( _lune.nilacc( classType:get_scope(), 'getSymbolInfoChild', 'callmtd' , getterName ))
         self:outputDeclFunc( _lune.newAlge( FuncInfo.Type, {getterSym:get_typeInfo()}) )
         self:write( string.format( "{ return self.%s ", self:getSymbolSym( memberSym )) )
         self:writeln( "}" )
      end
      
      if memberNode:get_setterMode(  ) ~= Ast.AccessMode.None then
         local setterName = string.format( "set_%s", memberName)
         local setterSym = _lune.unwrap( _lune.nilacc( classType:get_scope(), 'getSymbolInfoChild', 'callmtd' , setterName ))
         self:outputDeclFunc( _lune.newAlge( FuncInfo.Type, {setterSym:get_typeInfo()}) )
         self:write( string.format( "{ self.%s = arg1 ", self:getSymbolSym( memberSym )) )
         self:writeln( "}" )
      end
      
   end
   
end


function convFilter:outputStaticMember( node )

   if self.processMode == ProcessMode.DeclClass then
      for __index, memberNode in ipairs( node:get_memberList() ) do
         if memberNode:get_staticFlag() then
            self:writeln( string.format( "var %s %s", self:getSymbol( _lune.newAlge( SymbolKind.Static, {node:get_expType()}), memberNode:get_name().txt ), self:type2gotype( memberNode:get_expType() )) )
         end
         
      end
      
      do
         local initBlock = node:get_initBlock():get_func()
         if initBlock ~= nil then
            filter( initBlock, self, node )
            self:writeln( "" )
         end
      end
      
   else
    
      do
         local initBlock = node:get_initBlock():get_func()
         if initBlock ~= nil then
            self:writeln( string.format( "%s()", self:getFuncSymbol( initBlock:get_expType() )) )
         end
      end
      
   end
   
end


function convFilter:processDeclClass( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclClass'

   if self.processMode == ProcessMode.DeclClass then
      do
         local _switchExp = node:get_expType():get_kind()
         if _switchExp == Ast.TypeInfoKind.Class then
            self:writeln( string.format( "// declaration Class -- %s", node:get_expType():get_rawTxt()) )
            self:outputStaticMember( node )
            self:outputMethodIF( node )
            self:outputClassType( node )
            self:outputDownCast( node )
            self:outputCastReceiver( node )
            self:outputConstructor( node )
            self:outputAccessor( node )
            
            if #node:get_advertiseList() ~= 0 then
               Util.err( string.format( "%s: not support advertise", __func__) )
            end
            
            
            for __index, ifType in ipairs( node:get_expType():get_interfaceList() ) do
               if ifType == Ast.builtinTypeMapping then
                  Util.err( string.format( "%s: not support mapping.", __func__) )
               end
               
            end
            
            
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
      
   else
    
      self:outputStaticMember( node )
   end
   
end


function convFilter:outputCallPrefix( callId, node, prefixNode, funcSymbol )
   local __func__ = '@lune.@base.@convGo.convFilter.outputCallPrefix'

   local funcType = funcSymbol:get_typeInfo()
   local nilAccName = string.format( "%s_%d", self.moduleTypeInfo:get_rawTxt():gsub( "@", "" ), callId)
   local function processNilAcc( workPrefixNode )
   
      if not node:hasNilAccess(  ) then
         return 
      end
      
      local retNum = #funcType:get_retTypeInfoList()
      do
         local _switchExp = retNum
         if _switchExp == 0 then
            self:write( "Lns_NilAccCall0( func () {" )
         elseif _switchExp == 1 then
            self:write( "Lns_NilAccCall1( func () LnsAny { return " )
         else 
            
               if retNum <= MaxNilAccNum then
                  local anys = "LnsAny"
                  for _5889 = 2, retNum do
                     anys = string.format( "%s,LnsAny", anys)
                  end
                  
                  self:write( string.format( "Lns_NilAccCall%d( func () (%s) { return ", retNum, anys) )
               else
                
                  local args = "LnsAny"
                  for _5891 = 2, retNum do
                     args = string.format( "%s,LnsAny", args)
                  end
                  
                  self:write( string.format( "lns_NilAccCall_%s( func () (%s) { return ", nilAccName, args) )
               end
               
         end
      end
      
      self:write( string.format( "Lns_NilAccPop().(%s)", self:type2gotype( workPrefixNode:get_expType():get_nonnilableType() )) )
   end
   
   local closeParen = false
   if prefixNode ~= nil then
      if node:hasNilAccess(  ) then
         if #funcType:get_retTypeInfoList() >= 2 then
            if #funcType:get_retTypeInfoList() <= MaxNilAccNum then
               self:write( string.format( "Lns_NilAccFinCall%d(", #funcType:get_retTypeInfoList()) )
            else
             
               self:write( string.format( "lns_NilAccFinCall_%s(", nilAccName) )
            end
            
            closeParen = true
         end
         
      end
      
      if not funcType:get_staticFlag() then
         if node:hasNilAccess(  ) and not prefixNode:hasNilAccess(  ) then
            self:write( "Lns_NilAccFin(" )
            self:write( "Lns_NilAccPush(" )
            filter( prefixNode, self, node )
            self:writeln( ") && " )
         else
          
            filter( prefixNode, self, node )
         end
         
      end
      
      
      processNilAcc( prefixNode )
      
      if Ast.isBuiltin( funcType:get_typeId() ) then
         local builtinFuncs = TransUnit.getBuiltinFunc(  )
         
         do
            local _switchExp = funcType
            if _switchExp == builtinFuncs.list_insert then
               self:write( ".Insert" )
            elseif _switchExp == builtinFuncs.list_remove then
               self:write( ".Remove" )
            else 
               
                  Util.err( string.format( "%s: not support -- %s", __func__, funcType:getTxt(  )) )
            end
         end
         
      else
       
         if Ast.isBuiltin( funcType:get_typeId() ) then
            local builtinFuncs = TransUnit.getBuiltinFunc(  )
            do
               local _switchExp = funcType
               if _switchExp == builtinFuncs.list_insert then
                  self:write( ".Insert" )
               elseif _switchExp == builtinFuncs.list_remove then
                  self:write( ".Remove" )
               else 
                  
                     Util.err( string.format( "%s: not support -- %s", __func__, funcType:getTxt(  )) )
               end
            end
            
         else
          
            do
               local _switchExp = funcType:get_kind()
               if _switchExp == Ast.TypeInfoKind.Method then
                  self:write( ".FP." )
                  self:write( self:getFuncSymbol( funcType ) )
               else 
                  
                     self:write( self:getSymbolSym( funcSymbol ) )
               end
            end
            
         end
         
      end
      
   end
   
   
   return closeParen
end


function convFilter:processExpCall( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpCall'

   local funcType = node:get_func():get_expType()
   
   if funcType:get_kind() == Ast.TypeInfoKind.Method and funcType:get_parentInfo():get_kind() == Ast.TypeInfoKind.Enum and funcType:get_rawTxt() == "get__txt" then
      self:write( string.format( "%s(", self:getEnumGetTxtSym( _lune.unwrap( _lune.__Cast( funcType:get_parentInfo(), 3, Ast.EnumTypeInfo )) )) )
      local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
      if  nil == fieldNode then
         local _fieldNode = fieldNode
      
         Util.err( "not support -- __func__" )
      end
      
      
      filter( fieldNode:get_prefix(), self, node )
      self:write( ")" )
      return 
   end
   
   
   local retGenerics = isRetGenerics( node )
   if retGenerics and #funcType:get_retTypeInfoList() ~= 1 then
      self:write( string.format( "lns_convGenerics%d( ", node:get_id()) )
   end
   
   
   local closeParen = false
   do
      local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
      if fieldNode ~= nil then
         closeParen = self:outputCallPrefix( node:get_id(), fieldNode, fieldNode:get_prefix(), _lune.unwrap( fieldNode:get_symbolInfo()) )
      else
         if Ast.isBuiltin( funcType:get_typeId() ) then
            local builtinFuncs = TransUnit.getBuiltinFunc(  )
            
            do
               local _switchExp = funcType
               if _switchExp == builtinFuncs.lns_print then
                  self:write( "Lns_print" )
               else 
                  
                     Util.err( string.format( "%s: not support -- %s", __func__, funcType:getTxt(  )) )
               end
            end
            
         else
          
            filter( node:get_func(), self, node )
         end
         
      end
   end
   
   
   self:write( "(" )
   
   do
      local argList = node:get_argList()
      if argList ~= nil then
         self:processSetFromExpList( getConvExpName( node:get_id(), argList ), funcType:get_argTypeInfoList(), argList )
      end
   end
   
   
   self:write( ")" )
   
   if retGenerics then
      if #funcType:get_retTypeInfoList() == 1 then
         self:write( self:getAny2Type( node:get_expType() ) )
      else
       
         self:write( ")" )
      end
      
   end
   
   
   if node:hasNilAccess(  ) then
      self:write( "})" )
      if opt.node:hasNilAccess(  ) then
         self:writeln( " &&" )
      else
       
         self:write( ")" )
      end
      
      if closeParen then
         self:write( ")" )
      end
      
   end
   
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
      elseif _switchExp == "#" then
         do
            local _switchExp = node:get_exp():get_expType():get_kind()
            if _switchExp == Ast.TypeInfoKind.List then
               filter( node:get_exp(), self, node )
               self:write( ".Len()" )
            else 
               
                  self:write( "len(" )
                  filter( node:get_exp(), self, node )
                  self:write( ")" )
            end
         end
         
      else 
         
            Util.err( string.format( "%s: not support -- %s", __func__, node:get_op().txt) )
      end
   end
   
end


function convFilter:processExpMultiTo1( node, opt )

   self:write( "Lns_car(" )
   filter( node:get_exp(), self, node )
   if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.DDD then
      self:write( "..." )
   end
   
   self:write( ")" )
   self:outputConv( Ast.builtinTypeStem_, node:get_expType() )
end


function convFilter:outputImplicitCast( castType, node )

   do
      local _switchExp = castType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Nilable then
         self:outputImplicitCast( castType:get_nonnilableType(), node )
      elseif _switchExp == Ast.TypeInfoKind.Class then
         if castType == Ast.builtinTypeString then
            filter( node:get_exp(), self, node )
         else
          
            self:write( "&" )
            filter( node:get_exp(), self, node )
            self:write( string.format( ".%s", self:getTypeSymbol( castType )) )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.IF then
         filter( node:get_exp(), self, node )
      elseif _switchExp == Ast.TypeInfoKind.FormFunc then
         self:write( string.format( "%s(", self:getTypeSymbol( castType )) )
         filter( node:get_exp(), self, node )
         self:write( ")" )
      else 
         
            filter( node:get_exp(), self, node )
            if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Class and node:get_exp():get_expType() ~= Ast.builtinTypeString then
               self:write( ".FP" )
            end
            
      end
   end
   
end


function convFilter:processExpCast( node, opt )

   do
      local _switchExp = node:get_castKind()
      if _switchExp == Nodes.CastKind.Force then
         filter( node:get_exp(), self, node )
      elseif _switchExp == Nodes.CastKind.Implicit then
         self:outputImplicitCast( node:get_castType(), node )
      elseif _switchExp == Nodes.CastKind.Normal then
         local typeName = self:getTypeSymbol( node:get_castType() )
         self:write( string.format( "%sDownCastF(", typeName) )
         filter( node:get_exp(), self, node )
         if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Class and node:get_exp():get_expType() ~= Ast.builtinTypeString then
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
                  self:write( self:getSymbolSym( node:get_symbolInfo() ) )
               elseif _switchExp == Ast.SymbolKind.Fun then
                  self:write( self:getSymbol( _lune.newAlge( SymbolKind.Func, {node:get_expType()}), node:get_symbolInfo():get_name() ) )
               elseif _switchExp == Ast.SymbolKind.Typ then
                  self:write( self:getTypeSymbol( node:get_expType() ) )
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
            self:write( self:getAny2Type( node:get_expType() ) )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Map then
         filter( node:get_val(), self, node )
         self:write( "[" )
         do
            local index = node:get_index()
            if index ~= nil then
               filter( index, self, node )
            else
               self:write( str2gostr( _lune.unwrap( node:get_symbol()) ) )
            end
         end
         
         self:write( "]" )
      else 
         
            Util.err( string.format( "not support -- %s", __func__) )
      end
   end
   
end


function convFilter:processRefField( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processRefField'

   if node:get_expType():get_kind() == Ast.TypeInfoKind.Enum then
      self:write( self:getSymbol( _lune.newAlge( SymbolKind.Static, {node:get_expType()}), node:get_field().txt ) )
      return 
   end
   
   do
      local symbol = node:get_symbolInfo()
      if symbol ~= nil then
         if symbol:get_staticFlag() then
            self:write( self:getSymbolSym( symbol ) )
            return 
         end
         
      end
   end
   
   
   local openParenNum
   
   local andOp
   
   if not node:hasNilAccess(  ) then
      openParenNum = 0
      andOp = false
      if not node:get_prefix():hasNilAccess(  ) then
         filter( node:get_prefix(), self, node )
      else
       
         Util.err( string.format( "%s: not support", __func__) )
      end
      
   else
    
      if not node:get_prefix():hasNilAccess(  ) then
         self:write( "Lns_NilAccFin(" )
         self:write( "Lns_NilAccPush(" )
         filter( node:get_prefix(), self, node )
         self:writeln( ") && " )
      else
       
         filter( node:get_prefix(), self, node )
      end
      
      self:write( "Lns_NilAccPush(" )
      if opt.node:hasNilAccess(  ) then
         andOp = true
         openParenNum = 1
      else
       
         andOp = false
         openParenNum = 2
      end
      
      self:write( string.format( "Lns_NilAccPop().(%s)", self:type2gotype( node:get_prefix():get_expType():get_nonnilableType() )) )
   end
   
   
   self:write( "." )
   do
      local symbol = node:get_symbolInfo()
      if symbol ~= nil then
         self:write( self:getSymbolSym( symbol ) )
         local orgSym = symbol:getOrg(  )
         if orgSym:get_kind() == Ast.SymbolKind.Mbr and orgSym:get_typeInfo():get_kind() == Ast.TypeInfoKind.Alternate and not isAnyType( symbol:get_typeInfo() ) then
            self:write( self:getAny2Type( symbol:get_typeInfo() ) )
         end
         
      else
         Util.err( string.format( "not support -- %s", __func__) )
      end
   end
   
   
   for _6006 = 1, openParenNum do
      self:write( ")" )
   end
   
   if andOp then
      self:writeln( "&&" )
   end
   
end


function convFilter:processExpOmitEnum( node, opt )

   self:write( self:getSymbol( _lune.newAlge( SymbolKind.Static, {node:get_expType()}), node:get_valInfo():get_name() ) )
end


function convFilter:processGetField( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processGetField'

   do
      local symbolInfo = node:get_symbolInfo()
      if symbolInfo ~= nil then
         if symbolInfo:get_kind() == Ast.SymbolKind.Mtd and symbolInfo:get_name() == "get__txt" then
            do
               local enumType = _lune.__Cast( symbolInfo:get_namespaceTypeInfo(), 3, Ast.EnumTypeInfo )
               if enumType ~= nil then
                  self:write( string.format( "%s( ", self:getEnumGetTxtSym( enumType )) )
                  filter( node:get_prefix(), self, node )
                  self:write( ")" )
                  return 
               end
            end
            
            if _lune.__Cast( symbolInfo:get_namespaceTypeInfo(), 3, Ast.AlgeTypeInfo ) then
               filter( node:get_prefix(), self, node )
               self:write( ".(LnsAlgeVal).getTxt()" )
               return 
            end
            
         end
         
         
         if symbolInfo:get_staticFlag() then
            self:write( self:getSymbolSym( symbolInfo ) )
            self:write( "()" )
         else
          
            self:outputCallPrefix( node:get_id(), node, node:get_prefix(), symbolInfo )
            self:write( "()" )
            if symbolInfo:get_typeInfo():get_retTypeInfoList()[1]:get_kind() == Ast.TypeInfoKind.Alternate and not isAnyType( node:get_expType() ) then
               self:write( self:getAny2Type( node:get_expType() ) )
            end
            
            
            if node:hasNilAccess(  ) then
               self:write( "} )" )
               if opt.node:hasNilAccess(  ) then
                  self:writeln( " &&" )
               else
                
                  self:write( ")" )
               end
               
            end
            
         end
         
      else
         Util.err( string.format( "not support -- %s", __func__) )
      end
   end
   
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
   do
      local expList = node:get_dddParam()
      if expList ~= nil then
         self:write( string.format( 'Lns_getVM().string_format(%s, ', str2gostr( txt )) )
         self:processSetFromExpList( getConvExpName( node:get_id(), expList ), {Ast.builtinTypeDDD}, expList )
         self:write( ")" )
      else
         self:write( string.format( '%s', str2gostr( txt )) )
      end
   end
   
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
