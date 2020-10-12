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
   self.builtin2runtime = {}
   self.type2gotypeMap = {}
   self.nodeManager = Nodes.NodeManager.new()
   self.enableTest = enableTest
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



local ignoreNodeInInnerBlockSet = {[Nodes.NodeKind.get_DeclAlge()] = true, [Nodes.NodeKind.get_DeclEnum()] = true, [Nodes.NodeKind.get_DeclMethod()] = true, [Nodes.NodeKind.get_DeclForm()] = true, [Nodes.NodeKind.get_DeclMacro()] = true, [Nodes.NodeKind.get_TestBlock()] = true}

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
      if _switchExp == Ast.TypeInfoKind.Stem or _switchExp == Ast.TypeInfoKind.Alge then
         return true
      elseif _switchExp == Ast.TypeInfoKind.Alternate then
         if typeInfo:hasBase(  ) then
            return isAnyType( typeInfo:get_baseTypeInfo() )
         end
         
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

SymbolKind.Arg = { "Arg"}
SymbolKind._name2Val["Arg"] = SymbolKind.Arg
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

   if external then
      return name:sub( 1, 1 ):upper(  ) .. name:sub( 2 )
   end
   
   return name
   
end
local function isInnerDeclType( typeInfo )

   if typeInfo:get_kind() == Ast.TypeInfoKind.FormFunc then
      return typeInfo:get_parentInfo():get_kind() ~= Ast.TypeInfoKind.Module
   end
   
   
   if typeInfo:get_parentInfo():get_kind() ~= Ast.TypeInfoKind.Module or _lune.nilacc( _lune.nilacc( typeInfo:get_scope(), 'get_parent', 'callmtd' ), 'get_ownerTypeInfo', 'callmtd' ) == nil then
      return true
   end
   
   return false
end

function convFilter:getCanonicalName( typeInfo, localFlag )

   local enumName = typeInfo:getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager(), localFlag )
   return string.format( "%s", (enumName:gsub( "&", "" ) ))
end


local function getModuleName( typeInfo )

   return (typeInfo:getModule(  ):get_rawTxt():gsub( "@", "" ) )
end

local function concatSymWithType( name, typeInfo )

   return concatGLSym( string.format( "%s_", getModuleName( typeInfo:getModule(  ) )) .. name, Ast.isPubToExternal( typeInfo:get_accessMode() ) )
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
      if _matchExp[1] == SymbolKind.Arg[1] then
      
         return symbolName
      elseif _matchExp[1] == SymbolKind.Var[1] then
         local symbolInfo = _matchExp[2][1]
      
         local modName = self.moduleTypeInfo:get_rawTxt():gsub( "@", "" )
         if name == "__mod__" then
            symbolName = string.format( "%s__mod__", modName)
         elseif symbolInfo:get_scope() == self.moduleScope then
            symbolName = concatGLSym( string.format( "%s_", modName) .. symbolName, Ast.isPubToExternal( symbolInfo:get_accessMode() ) )
         elseif not symbolInfo:get_posForModToRef() then
            if symbolName ~= "__func__" then
               symbolName = "_"
            end
            
         end
         
      elseif _matchExp[1] == SymbolKind.Member[1] then
         local external = _matchExp[2][1]
      
         symbolName = concatGLSym( symbolName, external )
      elseif _matchExp[1] == SymbolKind.Func[1] then
         local typeInfo = _matchExp[2][1]
      
         if typeInfo:get_kind() == Ast.TypeInfoKind.Method then
            do
               local _switchExp = symbolName
               if _switchExp == "_toMap" then
                  return "ToMap"
               else 
                  
                     symbolName = concatGLSym( symbolName, Ast.isPubToExternal( typeInfo:get_accessMode() ) )
               end
            end
            
         else
          
            if isInnerDeclType( typeInfo ) then
               if not isClosure( typeInfo ) then
                  local parentName = typeInfo:getParentFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager(), true )
                  symbolName = string.format( "%s_%s_%d_", parentName:gsub( "[%.@]", "_" ), symbolName, typeInfo:get_typeId())
               end
               
            else
             
               if not Ast.isBuiltin( typeInfo:get_typeId() ) and not Ast.isPubToExternal( typeInfo:get_accessMode() ) then
                  symbolName = string.format( "%s_%d_", symbolName, typeInfo:get_typeId())
               end
               
            end
            
            symbolName = concatSymWithType( symbolName, typeInfo )
         end
         
      elseif _matchExp[1] == SymbolKind.Type[1] then
         local typeInfo = _matchExp[2][1]
      
         if typeInfo:get_kind() == Ast.TypeInfoKind.FormFunc then
            return self:getSymbol( _lune.newAlge( SymbolKind.Func, {typeInfo}), symbolName )
         end
         
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
      elseif _switchExp == Ast.SymbolKind.Arg then
         return self:getSymbol( _lune.newAlge( SymbolKind.Arg), symbolInfo:get_name() )
      elseif _switchExp == Ast.SymbolKind.Var then
         return self:getSymbol( _lune.newAlge( SymbolKind.Var, {symbolInfo}), symbolInfo:get_name() )
      elseif _switchExp == Ast.SymbolKind.Mbr then
         if symbolInfo:get_staticFlag() then
            return self:getSymbol( _lune.newAlge( SymbolKind.Static, {symbolInfo:get_namespaceTypeInfo()}), symbolInfo:get_name() )
         end
         
         return self:getSymbol( _lune.newAlge( SymbolKind.Member, {Ast.isPubToExternal( symbolInfo:get_accessMode() )}), symbolInfo:get_name() )
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
      work = Parser.convFromRawToStr( work ):gsub( '"', '\\"' )
      work = string.format( '"%s"', work)
   end
   
   
   work = work:gsub( "\\\n", "\\n" )
   work = work:gsub( "\\9", "\\t" )
   return work
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

function convFilter:type2gotype( typeInfo )

   if typeInfo:get_kind() == Ast.TypeInfoKind.DDD then
      return "[]LnsAny"
   end
   
   if isAnyType( typeInfo ) then
      return "LnsAny"
   end
   
   local orgType = getOrgTypeInfo( typeInfo )
   do
      local goType = self.type2gotypeMap[orgType]
      if goType ~= nil then
         return goType
      end
   end
   
   
   do
      local _switchExp = orgType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Ext then
         return "*Lns_luaValue"
      elseif _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         return "*LnsList"
      elseif _switchExp == Ast.TypeInfoKind.Set then
         return "*LnsSet"
      elseif _switchExp == Ast.TypeInfoKind.Map then
         return "*LnsMap"
      elseif _switchExp == Ast.TypeInfoKind.Form then
         return "LnsForm"
      elseif _switchExp == Ast.TypeInfoKind.FormFunc then
         return self:getFuncSymbol( typeInfo )
      elseif _switchExp == Ast.TypeInfoKind.Class then
         return "*" .. self:getTypeSymbol( typeInfo )
      elseif _switchExp == Ast.TypeInfoKind.IF then
         return self:getTypeSymbol( typeInfo )
      elseif _switchExp == Ast.TypeInfoKind.Alternate then
         return self:type2gotype( typeInfo:get_baseTypeInfo() )
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
function convFilter:outputAny2Type( dstType )

   if not isAnyType( dstType ) and dstType:get_kind() ~= Ast.TypeInfoKind.Alternate then
      self:write( string.format( ".(%s)", self:type2gotype( dstType )) )
   end
   
end





function convFilter:processBlankLine( node, opt )

end



function convFilter:processNone( node, opt )

end



function convFilter:processImport( node, opt )

   if node:get_modulePath() == "lune.base.Depend" then
      self:writeln( "Lns_LuaVer_init()" )
   end
   
   self:writeln( string.format( "Lns_%s_init()", getModuleName( node:get_moduleTypeInfo() )) )
end


local function needConvFormFunc( node )

   local castType = node:get_castType():get_extedType()
   if castType:get_kind() ~= Ast.TypeInfoKind.FormFunc then
      return false
   end
   
   
   local funcType = node:get_exp():get_expType():get_extedType()
   if #castType:get_argTypeInfoList() ~= #funcType:get_argTypeInfoList() or #castType:get_retTypeInfoList() ~= #funcType:get_retTypeInfoList() then
      return true
   end
   
   for index, argType in ipairs( castType:get_argTypeInfoList() ) do
      if argType ~= funcType:get_argTypeInfoList()[index] then
         return true
      end
      
   end
   
   for index, retType in ipairs( castType:get_retTypeInfoList() ) do
      if retType ~= funcType:get_retTypeInfoList()[index] then
         return true
      end
      
   end
   
   return false
end

local function needConvCast( dstType, srcType )

   do
      local _switchExp = dstType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Nilable then
         return needConvCast( dstType:get_nonnilableType(), srcType )
      elseif _switchExp == Ast.TypeInfoKind.Class then
         if dstType == Ast.builtinTypeString or srcType:get_genSrcTypeInfo():get_srcTypeInfo():get_nonnilableType() == dstType:get_genSrcTypeInfo():get_srcTypeInfo():get_nonnilableType() then
            return false
         else
          
            return true
         end
         
      elseif _switchExp == Ast.TypeInfoKind.IF then
         return false
      elseif _switchExp == Ast.TypeInfoKind.FormFunc then
         return true
      elseif _switchExp == Ast.TypeInfoKind.Alternate then
         if not dstType:hasBase(  ) then
            return false
         else
          
            return needConvCast( dstType:get_baseTypeInfo(), srcType )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Form then
         return true
      elseif _switchExp == Ast.TypeInfoKind.Prim then
         if not dstType:get_nilable() then
            do
               local _switchExp = dstType
               if _switchExp == Ast.builtinTypeInt then
                  return true
               elseif _switchExp == Ast.builtinTypeReal then
                  return true
               else 
                  
                     return false
               end
            end
            
         else
          
            return false
         end
         
      else 
         
            if srcType:get_kind() == Ast.TypeInfoKind.Class and srcType ~= Ast.builtinTypeString then
               return true
            else
             
               return false
            end
            
      end
   end
   
end
function convFilter:outputImplicitCast( castType, node, parent )

   do
      local _switchExp = castType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Nilable then
         self:outputImplicitCast( castType:get_nonnilableType(), node, parent )
      elseif _switchExp == Ast.TypeInfoKind.Class then
         if castType == Ast.builtinTypeString or node:get_expType():get_kind() == Ast.TypeInfoKind.Alternate or node:get_expType():get_genSrcTypeInfo():get_srcTypeInfo():get_nonnilableType() == castType:get_genSrcTypeInfo():get_srcTypeInfo():get_nonnilableType() then
            filter( node, self, parent )
         else
          
            if isAnyType( node:get_expType() ) then
               self:write( string.format( "%sDownCastF(", self:getTypeSymbol( castType )) )
               filter( node, self, parent )
               self:write( ")" )
            else
             
               self:write( "&" )
               filter( node, self, parent )
               self:write( string.format( ".%s", self:getTypeSymbol( castType )) )
            end
            
         end
         
      elseif _switchExp == Ast.TypeInfoKind.IF then
         filter( node, self, parent )
      elseif _switchExp == Ast.TypeInfoKind.FormFunc then
         local expType = node:get_expType()
         if needConvFormFunc( parent ) then
            self:write( string.format( "conv2Form%d(", parent:get_id()) )
            filter( node, self, parent )
            self:write( ")" )
         else
          
            self:write( string.format( "%s(", self:getTypeSymbol( castType )) )
            filter( node, self, parent )
            self:write( ")" )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Alternate then
         if not castType:hasBase(  ) then
            filter( node, self, parent )
         else
          
            self:outputImplicitCast( castType:get_baseTypeInfo(), node, parent )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Form then
         self:write( string.format( "conv2Form%d(", parent:get_id()) )
         filter( node, self, parent )
         self:write( ")" )
      elseif _switchExp == Ast.TypeInfoKind.Prim then
         if not node:get_expType():get_nilable() then
            do
               local _switchExp = castType
               if _switchExp == Ast.builtinTypeInt then
                  self:write( "LnsInt(" )
                  filter( node, self, parent )
                  self:write( ")" )
               elseif _switchExp == Ast.builtinTypeReal then
                  self:write( "LnsReal(" )
                  filter( node, self, parent )
                  self:write( ")" )
               else 
                  
                     filter( node, self, parent )
               end
            end
            
         else
          
            filter( node, self, parent )
         end
         
      else 
         
            if node:get_expType():get_nilable() and node:get_expType():get_nonnilableType():get_kind() == Ast.TypeInfoKind.Class then
               self:write( string.format( "%s2Stem(", self:getTypeSymbol( node:get_expType():get_nonnilableType() )) )
               filter( node, self, parent )
               self:write( ")" )
            else
             
               filter( node, self, parent )
               if node:get_expType():get_kind() == Ast.TypeInfoKind.Class and node:get_expType() ~= Ast.builtinTypeString then
                  self:write( ".FP" )
               end
               
            end
            
      end
   end
   
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

   if #dstTypeList > 1 and node:get_mRetExp() then
      for __index, exp in ipairs( node:get_expList() ) do
         do
            local castNode = _lune.__Cast( exp, 3, Nodes.ExpCastNode )
            if castNode ~= nil then
               if needConvCast( castNode:get_castType(), castNode:get_exp():get_expType() ) then
                  return ExpListKind.Conv
               end
               
            end
         end
         
      end
      
   end
   
   
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
   
   
   
   local expList = argList:get_expList()
   local mRetIndex = _lune.nilacc( argList:get_mRetExp(), 'get_index', 'callmtd' )
   
   if not mRetIndex then
      if expList[#expList]:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
      else
       
         return 
      end
      
   end
   
   
   
   local workList = {}
   for __index, exp in ipairs( expList ) do
      local workExp = Nodes.getUnwraped( exp )
      if workExp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
         break
      end
      
      table.insert( workList, workExp )
   end
   
   expList = workList
   
   self:write( string.format( "func %s(", getConvExpName( nodeId, argList )) )
   
   for index, argExp in ipairs( expList ) do
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
               do
                  local castNode = _lune.__Cast( argExp, 3, Nodes.ExpCastNode )
                  if castNode ~= nil then
                     self:write( self:type2gotype( castNode:get_castType() ) )
                  else
                     self:write( self:type2gotype( argExp:get_expType() ) )
                  end
               end
               
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
   
   if mRetIndex ~= nil then
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
            local wrote = false
            if index <= #expList then
               local exp = expList[index]
               do
                  local castNode = _lune.__Cast( exp, 3, Nodes.ExpCastNode )
                  if castNode ~= nil then
                     local statNode = Nodes.ConvStatNode.create( self.nodeManager, exp:get_pos(), false, {exp:get_expType()}, string.format( "Lns_getFromMulti( arg%d, 0 )", index) )
                     self:outputImplicitCast( castNode:get_castType(), statNode, castNode )
                     wrote = true
                  end
               end
               
            end
            
            if not wrote then
               self:write( string.format( "Lns_getFromMulti( arg%d, %d )", mRetIndex, index - mRetIndex) )
               self:outputAny2Type( dstType )
            end
            
         else
          
            self:write( string.format( "arg%d", index) )
         end
         
      end
      
      if restIndex ~= nil then
         self:write( "Lns_2DDD( " )
         for index, _5568 in ipairs( expList ) do
            if index >= restIndex then
               if index < #expList then
                  self:write( string.format( "arg%d", index) )
               else
                
                  self:write( string.format( "arg%d[%d:]", mRetIndex, index - mRetIndex) )
                  break
               end
               
            end
            
         end
         
         self:writeln( ")" )
      else
         self:writeln( "" )
      end
      
   else
      for index, _5570 in ipairs( dstTypeList ) do
         if index ~= 1 then
            self:write( ", " )
         end
         
         if index <= #expList then
            self:write( string.format( "arg%d", index) )
         else
          
            self:write( "nil" )
         end
         
      end
      
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
    return Lns_NilAccPush( Lns_2DDD( call() ) )
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
            self:outputAny2Type( dstType )
         end
         
      end
      
   end
   
   self:writeln( "" )
   self:popIndent(  )
   self:writeln( "}" )
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

FuncInfo.Advertise = { "Advertise", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} },{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
FuncInfo._name2Val["Advertise"] = FuncInfo.Advertise
FuncInfo.DeclInfo = { "DeclInfo", {{ func=Nodes.Node._fromMap, nilable=false, child={} },{ func=Nodes.DeclFuncInfo._fromMap, nilable=false, child={} }}}
FuncInfo._name2Val["DeclInfo"] = FuncInfo.DeclInfo
FuncInfo.Type = { "Type", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
FuncInfo._name2Val["Type"] = FuncInfo.Type


function convFilter:outputRetType( retTypeList )

   do
      local _switchExp = #retTypeList
      if _switchExp == 0 then
         self:write( "" )
      elseif _switchExp == 1 then
         if retTypeList[1] ~= Ast.builtinTypeNeverRet then
            self:write( " " .. self:type2gotype( retTypeList[1] ) )
         end
         
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


function convFilter:outputDeclFunc( funcInfo )

   local typeInfo
   
   local name
   
   local prefixType
   
   do
      local _matchExp = funcInfo
      if _matchExp[1] == FuncInfo.DeclInfo[1] then
         local node = _matchExp[2][1]
         local workDeclInfo = _matchExp[2][2]
      
         typeInfo = node:get_expType()
         prefixType = typeInfo:get_parentInfo()
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
         prefixType = typeInfo:get_parentInfo()
         name = typeInfo:get_rawTxt()
      elseif _matchExp[1] == FuncInfo.Advertise[1] then
         local classType = _matchExp[2][1]
         local methodType = _matchExp[2][2]
      
         typeInfo = methodType
         prefixType = classType
         name = typeInfo:get_rawTxt()
      end
   end
   
   
   if isClosure( typeInfo ) then
      self:write( "func" )
   else
    
      if typeInfo:get_kind() == Ast.TypeInfoKind.Method then
         self:write( "func " )
         self:write( "(self *" )
         self:write( self:getTypeSymbol( prefixType ) )
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
         
      elseif _matchExp[1] == FuncInfo.Advertise[1] then
         local _ = _matchExp[2][1]
         local _ = _matchExp[2][2]
      
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


function convFilter:outputConvToFormFunc( node )

   if not needConvFormFunc( node ) then
      return 
   end
   
   local castType = node:get_castType():get_extedType()
   
   local funcType = node:get_exp():get_expType():get_extedType()
   self:writeln( string.format( "// for %d: %s", node:get_pos().lineNo, Nodes.getNodeKindName( node:get_kind() )) )
   self:write( string.format( "func conv2Form%d( src func (", node:get_id()) )
   for index, argType in ipairs( funcType:get_argTypeInfoList() ) do
      if index > 1 then
         self:write( ", " )
      end
      
      self:write( string.format( "arg%d %s", index, self:type2gotype( argType )) )
   end
   
   self:write( ")" )
   self:outputRetType( funcType:get_retTypeInfoList() )
   self:write( ")" )
   self:outputDeclFunc( _lune.newAlge( FuncInfo.Type, {castType}) )
   self:writeln( " {" )
   
   self:pushIndent(  )
   self:write( "return " )
   self:outputDeclFunc( _lune.newAlge( FuncInfo.Type, {castType}) )
   self:writeln( " {" )
   
   self:pushIndent(  )
   if #funcType:get_argTypeInfoList() > 0 and #funcType:get_retTypeInfoList() > 0 then
      self:write( "return Lns_2DDD(" )
   end
   
   self:write( "src(" )
   if #castType:get_argTypeInfoList() > #funcType:get_argTypeInfoList() then
      for index, argType in ipairs( funcType:get_argTypeInfoList() ) do
         if index > 1 then
            self:write( ", " )
         end
         
         if index == #funcType:get_argTypeInfoList() then
            if argType:get_kind() == Ast.TypeInfoKind.DDD then
               self:write( string.format( "Lns_2DDD( arg%d", index) )
               for subIndex = index + 1, #castType:get_argTypeInfoList() do
                  self:write( string.format( ", arg%d", subIndex) )
               end
               
               self:write( ")" )
            else
             
               self:write( string.format( "arg%d", index) )
            end
            
         else
          
            self:write( string.format( "arg%d", index) )
         end
         
      end
      
   else
    
      for index, argType in ipairs( funcType:get_argTypeInfoList() ) do
         if index > 1 then
            self:write( ", " )
         end
         
         if index >= #castType:get_argTypeInfoList() then
            local subIndex = index - #castType:get_argTypeInfoList()
            if argType:get_kind() == Ast.TypeInfoKind.DDD then
               self:write( string.format( "arg%d[%d:]", #castType:get_argTypeInfoList(), subIndex) )
            else
             
               self:write( string.format( "Lns_getFromMulti( arg%d, %d )", #castType:get_argTypeInfoList(), subIndex) )
            end
            
         else
          
            self:write( string.format( "arg%d", index) )
         end
         
      end
      
   end
   
   self:write( ")" )
   if #funcType:get_retTypeInfoList() > 0 then
      if #funcType:get_argTypeInfoList() > 0 then
         self:writeln( ")" )
      end
      
   end
   
   self:writeln( "" )
   self:popIndent(  )
   self:writeln( "}" )
   
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:outputConvToForm( node )

   local castType = node:get_castType():get_extedType()
   if castType:get_kind() ~= Ast.TypeInfoKind.Form then
      return 
   end
   
   
   local funcType = node:get_exp():get_expType():get_extedType()
   if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Ext and funcType:get_srcTypeInfo():get_kind() == Ast.TypeInfoKind.Form then
      self:writeln( string.format( [==[      
func conv2Form%d( luaform LnsAny ) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        return Lns_getVM().RunLoadedfunc( luaform.(*Lns_luaValue), argList )
    }
}]==], node:get_id()) )
      return 
   end
   
   
   self:writeln( string.format( "// for %d: %s", node:get_pos().lineNo, Nodes.getNodeKindName( node:get_kind() )) )
   self:write( string.format( "func conv2Form%d( src func (", node:get_id()) )
   for index, argType in ipairs( funcType:get_argTypeInfoList() ) do
      if index > 1 then
         self:write( ", " )
      end
      
      self:write( string.format( "arg%d %s", index, self:type2gotype( argType )) )
   end
   
   self:write( ")" )
   self:outputRetType( funcType:get_retTypeInfoList() )
   self:writeln( ") LnsForm {" )
   
   self:pushIndent(  )
   self:writeln( "return func (argList []LnsAny) []LnsAny {" )
   
   self:pushIndent(  )
   if #funcType:get_retTypeInfoList() > 0 then
      self:write( "return " )
      if #funcType:get_argTypeInfoList() > 0 then
         self:write( "Lns_2DDD(" )
      end
      
   end
   
   self:write( "src(" )
   for index, argType in ipairs( funcType:get_argTypeInfoList() ) do
      if index > 1 then
         self:write( ", " )
      end
      
      if argType:get_kind() == Ast.TypeInfoKind.DDD then
         self:write( string.format( "argList[ %d: ]", index - 1) )
         break
      end
      
      
      self:write( string.format( "Lns_getFromMulti( argList, %d )", index - 1) )
   end
   
   self:write( ")" )
   if #funcType:get_retTypeInfoList() > 0 then
      if #funcType:get_argTypeInfoList() > 0 then
         self:writeln( ")" )
      else
       
         self:writeln( "" )
      end
      
   else
    
      self:writeln( "" )
      self:writeln( "return []LnsAny{}" )
   end
   
   self:popIndent(  )
   self:writeln( "}" )
   
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processConvStat( node, opt )

   self:write( node:get_txt() )
end


function convFilter:outputTopScopeVar( node )

   for __index, symbolInfo in ipairs( node:get_symbolInfoList() ) do
      if Ast.isPubToExternal( symbolInfo:get_accessMode() ) or symbolInfo:get_posForModToRef() then
         self:writeln( string.format( "var %s %s", self:getSymbolSym( symbolInfo ), self:type2gotype( symbolInfo:get_typeInfo() )) )
      end
      
   end
   
end


function convFilter:processRoot( node, opt )

   local builtinFuncs = TransUnit.getBuiltinFunc(  )
   
   
   local builtin2runtime = {[builtinFuncs.str_gsub] = 'Lns_getVM().String_gsub', [builtinFuncs.string_gsub] = 'Lns_getVM().String_gsub', [builtinFuncs.str_find] = 'Lns_getVM().String_find', [builtinFuncs.string_find] = 'Lns_getVM().String_find', [builtinFuncs.str_byte] = 'Lns_getVM().String_byte', [builtinFuncs.string_byte] = 'Lns_getVM().String_byte', [builtinFuncs.str_format] = 'Lns_getVM().String_format', [builtinFuncs.string_format] = 'Lns_getVM().String_format', [builtinFuncs.str_rep] = 'Lns_getVM().String_rep', [builtinFuncs.string_rep] = 'Lns_getVM().String_rep', [builtinFuncs.str_gmatch] = 'Lns_getVM().String_gmatch', [builtinFuncs.string_gmatch] = 'Lns_getVM().String_gmatch', [builtinFuncs.str_sub] = 'Lns_getVM().String_sub', [builtinFuncs.string_sub] = 'Lns_getVM().String_sub', [builtinFuncs.str_lower] = 'Lns_getVM().String_lower', [builtinFuncs.string_lower] = 'Lns_getVM().String_lower', [builtinFuncs.str_upper] = 'Lns_getVM().String_upper', [builtinFuncs.string_upper] = 'Lns_getVM().String_upper', [builtinFuncs.str_reverse] = 'Lns_getVM().String_reverse', [builtinFuncs.string_reverse] = 'Lns_getVM().String_reverse', [Ast.builtinTypeNone] = ""}
   
   
   builtin2runtime[builtinFuncs.lns_error] = "panic"
   builtin2runtime[builtinFuncs.lns_print] = "Lns_print"
   builtin2runtime[builtinFuncs.lns_type] = "Lns_type"
   builtin2runtime[builtinFuncs.lns_require] = "Lns_require"
   builtin2runtime[builtinFuncs.lns_tonumber] = "Lns_tonumber"
   builtin2runtime[builtinFuncs.lns__load] = "Lns_getVM().Load"
   
   builtin2runtime[builtinFuncs.string_dump] = "Lns_getVM().String_dump"
   
   builtin2runtime[builtinFuncs.io_open] = "Lns_getVM().IO_open"
   builtin2runtime[builtinFuncs.io_popen] = "Lns_getVM().IO_popen"
   builtin2runtime[builtinFuncs.package_searchpath] = "Lns_getVM().Package_searchpath"
   builtin2runtime[builtinFuncs.os_clock] = "Lns_getVM().OS_clock"
   builtin2runtime[builtinFuncs.os_exit] = "Lns_getVM().OS_exit"
   builtin2runtime[builtinFuncs.os_remove] = "Lns_getVM().OS_remove"
   builtin2runtime[builtinFuncs.os_date] = "Lns_getVM().OS_date"
   builtin2runtime[builtinFuncs.os_time] = "Lns_getVM().OS_time"
   builtin2runtime[builtinFuncs.os_difftime] = "Lns_getVM().OS_difftime"
   builtin2runtime[builtinFuncs.os_rename] = "Lns_getVM().OS_rename"
   builtin2runtime[builtinFuncs.math_random] = "Lns_getVM().Math_random"
   builtin2runtime[builtinFuncs.math_randomseed] = "Lns_getVM().Math_randomseed"
   
   self.builtin2runtime = builtin2runtime
   
   self.type2gotypeMap = {[Ast.builtinTypeInt] = "LnsInt", [Ast.builtinTypeReal] = "LnsReal", [Ast.builtinTypeStem] = "LnsAny", [Ast.builtinTypeString] = "string", [Ast.builtinTypeBool] = "bool", [builtinFuncs.ostream_] = "Lns_oStream", [builtinFuncs.istream_] = "Lns_iStream"}
   
   Ast.pushProcessInfo( node:get_processInfo() )
   
   self:writeln( "// This code is transcompiled by LuneScript." )
   
   self:writeln( "package lns" )
   
   self:writeln( 'import . "lns/lune/base/runtime_go"' )
   
   local initModVar = string.format( "init_%s", getModuleName( node:get_moduleTypeInfo() ))
   self:writeln( string.format( "var %s bool", initModVar) )
   
   self:pushProcessMode( ProcessMode.DeclTopScopeVar )
   local modSym = _lune.unwrap( self.moduleScope:getSymbolInfoChild( "__mod__" ))
   self:writeln( string.format( "var %s string", self:getSymbolSym( modSym )) )
   
   for __index, child in ipairs( node:get_nodeManager():getDeclEnumNodeList(  ) ) do
      filter( child, self, node )
   end
   
   for __index, child in ipairs( node:get_children() ) do
      do
         local declVarNode = _lune.__Cast( child, 3, Nodes.DeclVarNode )
         if declVarNode ~= nil then
            self:outputTopScopeVar( declVarNode )
         end
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
   
   for __index, workNode in ipairs( node:get_nodeManager():getExpCastNodeList(  ) ) do
      self:outputConvToForm( workNode )
      self:outputConvToFormFunc( workNode )
   end
   
   for __index, workNode in ipairs( node:get_nodeManager():getTestBlockNodeList(  ) ) do
      filter( workNode, self, node )
   end
   
   
   for __index, workNode in ipairs( node:get_nodeManager():getIfUnwrapNodeList(  ) ) do
      local symTypeList = {}
      for _5702 = 1, #workNode:get_varSymList() do
         table.insert( symTypeList, Ast.builtinTypeStem_ )
      end
      
      self:processConvExp( workNode:get_id(), symTypeList, workNode:get_expList() )
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
   
   self:writeln( string.format( "func Lns_%s_init() {", getModuleName( node:get_moduleTypeInfo() )) )
   self:pushIndent(  )
   
   self:writeln( string.format( "if %s { return }", initModVar) )
   self:writeln( string.format( "%s = true", initModVar) )
   
   self:writeln( string.format( '%s = "%s"', self:getSymbolSym( modSym ), node:get_moduleTypeInfo():getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager() )) )
   
   self:writeln( "Lns_InitMod()" )
   
   local modulePath = node:get_moduleTypeInfo():getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager() ):gsub( "@", "" )
   local moduleName = getModuleName( node:get_moduleTypeInfo() )
   
   for __index, workNode in ipairs( node:get_nodeManager():getTestBlockNodeList(  ) ) do
      self:writeln( string.format( 'Testing_registerTestcase( "%s", "%s", lns_test_%s_%s )', modulePath, workNode:get_name().txt, moduleName, workNode:get_name().txt) )
   end
   
   
   for __index, child in ipairs( node:get_children() ) do
      if not _lune._Set_has(ignoreNodeInInnerBlockSet, child:get_kind() ) then
         filter( child, self, node )
      end
      
   end
   
   
   self:popIndent(  )
   self:writeln( "}" )
   
   self:writeln( "func init() {" )
   self:pushIndent(  )
   self:writeln( string.format( "%s = false", initModVar) )
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
      if not _lune._Set_has(ignoreNodeInInnerBlockSet, child:get_kind() ) then
         filter( child, self, node )
      end
      
   end
   
   self:popIndent(  )
   self:popProcessMode(  )
   
   if node:get_blockKind() == Nodes.BlockKind.Block then
      self:writeln( "}" )
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
            if exp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
               break
            end
            
            if index ~= 1 then
               self:write( ', ' )
            end
            
            if mRetIndex == index then
               self:write( "Lns_2DDD(" )
               filter( Nodes.getCastUnwraped( exp ), self, expListNode )
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
   self:writeln( "" )
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
            
            self:writeln( string.format( "func (self *%s) GetTxt() string {", algeSym) )
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
   self:writeln( "" )
end


function convFilter:processExpMacroExp( node, opt )

   for __index, stmt in ipairs( node:get_stmtList() ) do
      if not _lune._Set_has(ignoreNodeInInnerBlockSet, stmt:get_kind() ) then
         filter( stmt, self, node )
      end
      
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
   self:writeln( string.format( "// %d: %s", node:get_pos().lineNo, Nodes.getNodeKindName( node:get_kind() )) )
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


function convFilter:processExpCallSuperCtor( node, opt )

   self:write( string.format( "self.%s(", self:getConstrSymbol( node:get_superType() )) )
   do
      local argList = node:get_expList()
      if argList ~= nil then
         self:processSetFromExpList( getConvExpName( node:get_id(), argList ), node:get_methodType():get_argTypeInfoList(), argList )
      end
   end
   
   self:writeln( ")" )
end


function convFilter:processExpCallSuper( node, opt )

   self:write( string.format( "self.%s.%s(", self:getTypeSymbol( node:get_methodType():get_parentInfo() ), self:getFuncSymbol( node:get_methodType() )) )
   do
      local argList = node:get_expList()
      if argList ~= nil then
         self:processSetFromExpList( getConvExpName( node:get_id(), argList ), node:get_methodType():get_argTypeInfoList(), argList )
      end
   end
   
   self:write( ")" )
end


function convFilter:outputDeclFuncInfo( node, declInfo )

   if node:get_expType():get_abstractFlag() then
      return 
   end
   
   if declInfo:get_name() and not isClosure( node:get_expType() ) then
      self:writeln( string.format( "// %d: decl %s", node:get_pos().lineNo, self:getCanonicalName( node:get_expType(), false )) )
   end
   
   self:outputDeclFunc( _lune.newAlge( FuncInfo.DeclInfo, {node,declInfo}) )
   
   self:writeln( " {" )
   
   self:pushIndent(  )
   
   if declInfo:get_has__func__Symbol() then
      local nameSpace = self:getCanonicalName( node:get_expType():get_parentInfo(), false )
      local funcName
      
      do
         local name = declInfo:get_name()
         if name ~= nil then
            funcName = name.txt
         else
            funcName = "<anonymous>"
         end
      end
      
      local funcSym_ = _lune.unwrap( _lune.nilacc( node:get_expType():get_scope(), 'getSymbolInfoChild', 'callmtd' , "__func__" ))
      self:writeln( string.format( '%s := "%s.%s"', self:getSymbolSym( funcSym_ ), nameSpace, funcName) )
   end
   
   
   self:popIndent(  )
   
   do
      local body = declInfo:get_body()
      if body ~= nil then
         filter( body, self, node )
         
         local retTypeInfoList = node:get_expType():get_retTypeInfoList()
         if #retTypeInfoList > 0 and retTypeInfoList[#retTypeInfoList] ~= Ast.builtinTypeNeverRet then
            local needReturn = false
            for index = #body:get_stmtList(), 1, -1 do
               local statment = body:get_stmtList()[index]
               do
                  local _switchExp = statment:get_kind()
                  if _switchExp == Nodes.NodeKind.get_BlankLine() then
                  elseif _switchExp == Nodes.NodeKind.get_Return() then
                     break
                  else 
                     
                        needReturn = true
                        break
                  end
               end
               
            end
            
            if needReturn then
               self:write( "    return " )
               for index, retType in ipairs( node:get_expType():get_retTypeInfoList() ) do
                  if index > 1 then
                     self:write( "," )
                  end
                  
                  if isAnyType( retType ) then
                     self:write( "nil" )
                  else
                   
                     do
                        local _switchExp = retType
                        if _switchExp == Ast.builtinTypeInt then
                           self:write( "0" )
                        elseif _switchExp == Ast.builtinTypeReal then
                           self:write( "0.0" )
                        elseif _switchExp == Ast.builtinTypeBool then
                           self:write( "false" )
                        elseif _switchExp == Ast.builtinTypeString then
                           self:write( '""' )
                        else 
                           
                              self:write( "nil" )
                        end
                     end
                     
                  end
                  
               end
               
               self:writeln( "" )
            end
            
         end
         
      end
   end
   
   self:write( "}" )
   if declInfo:get_name() then
      self:writeln( "" )
   end
   
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
   
   local tempTypeList = {}
   for index, varSym in ipairs( node:get_varSymList() ) do
      if index > 1 then
         self:write( ", " )
      end
      
      if varSym:get_name() == "_" then
         self:write( "_" )
      else
       
         self:write( "_" .. varSym:get_name() )
      end
      
      table.insert( tempTypeList, Ast.builtinTypeStem_ )
      if index == #node:get_varSymList() then
         for _5896 = index + 1, #node:get_expList():get_expTypeList() do
            self:write( ", _" )
         end
         
      end
      
   end
   
   self:write( " := " )
   self:processSetFromExpList( getConvExpName( node:get_id(), node:get_expList() ), tempTypeList, node:get_expList() )
   self:writeln( "" )
   self:write( "if " )
   local hasSym = false
   for __index, varSym in ipairs( node:get_varSymList() ) do
      if varSym:get_name() ~= "_" then
         if hasSym then
            self:write( " && " )
         end
         
         self:write( string.format( "_%s != nil", varSym:get_name()) )
         hasSym = true
      end
      
   end
   
   self:writeln( " {" )
   self:pushIndent(  )
   for index, varSym in ipairs( node:get_varSymList() ) do
      if varSym:get_name() ~= "_" then
         if varSym:hasAccess(  ) then
            self:write( string.format( "%s := _%s", varSym:get_name(), varSym:get_name()) )
            if node:get_expList():getExpTypeAt( index ):get_nilable() then
               self:outputAny2Type( varSym:get_typeInfo() )
            end
            
            self:writeln( "" )
         end
         
      end
      
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
   self:writeln( "" )
end


function convFilter:outputLetVar( node )

   local function declVar(  )
   
      if node:get_symbolInfoList()[1]:get_scope() == self.moduleScope then
         return 
      end
      
      for __index, symbolInfo in ipairs( node:get_symbolInfoList() ) do
         if symbolInfo:get_posForModToRef() then
            self:writeln( string.format( "var %s %s", self:getSymbolSym( symbolInfo ), self:type2gotype( symbolInfo:get_typeInfo() )) )
         end
         
      end
      
   end
   
   if node:get_unwrapFlag() then
      do
         local expList, unwrapBlock = node:get_expList(), node:get_unwrapBlock()
         if expList ~= nil and  unwrapBlock ~= nil then
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
               if index == #node:get_varList() then
                  do
                     local expListNode = node:get_expList()
                     if expListNode ~= nil then
                        for _5914 = index + 1, #expListNode:get_expTypeList() do
                           self:write( ", _" )
                        end
                        
                     end
                  end
                  
               end
               
            end
            
            self:write( " := " )
            self:processSetFromExpList( getConvExpName( node:get_id(), expList ), node:get_typeInfoList(), expList )
            self:writeln( "" )
            self:write( "if " )
            local hasCond = false
            for index, varInfo in ipairs( node:get_varList() ) do
               if expList:getExpTypeAt( index ):get_nilable() then
                  if hasCond then
                     self:write( " || " )
                  end
                  
                  self:write( string.format( "_%s == nil", varInfo:get_name().txt) )
                  hasCond = true
               end
               
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
                     if expList:getExpTypeAt( index ):get_nilable() then
                        self:outputAny2Type( varInfo:get_typeInfo() )
                     end
                     
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
                     if expList:getExpTypeAt( index ):get_nilable() then
                        self:outputAny2Type( varInfo:get_typeInfo() )
                     end
                     
                     self:writeln( "" )
                  end
                  
                  self:popIndent(  )
                  self:writeln( "}" )
               end
            end
            
            
            self:popIndent(  )
            self:writeln( "}" )
         end
      end
      
   else
    
      declVar(  )
      
      do
         local expList = node:get_expList()
         if expList ~= nil then
            for index, symbolInfo in ipairs( node:get_symbolInfoList() ) do
               if index > 1 then
                  self:write( "," )
               end
               
               if symbolInfo:get_posForModToRef() or Ast.isPubToExternal( symbolInfo:get_accessMode() ) then
                  self:write( string.format( "%s", self:getSymbolSym( symbolInfo )) )
               else
                
                  self:write( "_" )
               end
               
            end
            
            if #expList:get_expTypeList() > #node:get_symbolInfoList() then
               for _5926 = #node:get_symbolInfoList() + 1, #expList:get_expTypeList() do
                  self:write( ", _" )
               end
               
            end
            
            
            self:write( " = " )
            self:processSetFromExpList( getConvExpName( node:get_id(), expList ), node:get_typeInfoList(), expList )
            self:writeln( "" )
         end
      end
      
   end
   
end


function convFilter:processDeclVar( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclVar'

   do
      local _switchExp = node:get_mode()
      if _switchExp == Nodes.DeclVarMode.Let then
         self:outputLetVar( node )
      elseif _switchExp == Nodes.DeclVarMode.Unwrap then
         self:writeln( "{" )
         self:pushIndent(  )
         for __index, varSym in ipairs( node:get_symbolInfoList() ) do
            self:writeln( string.format( "var _%s LnsAny", varSym:get_name()) )
         end
         
         local expList = node:get_expList()
         if  nil == expList then
            local _expList = expList
         
            Util.err( "illegal" )
         end
         
         
         local function setVals(  )
         
            for index, varSym in ipairs( node:get_symbolInfoList() ) do
               self:write( string.format( "%s = _%s", varSym:get_name(), varSym:get_name()) )
               if expList:getExpTypeAt( index ):get_nilable() then
                  self:outputAny2Type( varSym:get_typeInfo() )
               end
               
               self:writeln( "" )
            end
            
         end
         
         local typeList = {}
         for index, varSym in ipairs( node:get_symbolInfoList() ) do
            table.insert( typeList, varSym:get_typeInfo() )
            if index > 1 then
               self:write( "," )
            end
            
            self:write( string.format( "_%s", varSym:get_name()) )
         end
         
         for _5943 = #node:get_symbolInfoList() + 1, #expList:get_expTypeList() do
            self:write( ",_" )
         end
         
         self:write( " = " )
         self:processSetFromExpList( getConvExpName( node:get_id(), expList ), typeList, expList )
         self:writeln( "" )
         self:write( "if " )
         local hasCond = false
         for index, varSym in ipairs( node:get_symbolInfoList() ) do
            if expList:getExpTypeAt( index ):get_nilable() then
               if hasCond then
                  self:write( " || " )
               end
               
               self:write( string.format( "_%s == nil", varSym:get_name()) )
               hasCond = true
            end
            
         end
         
         self:writeln( " {" )
         filter( _lune.unwrap( node:get_unwrapBlock()), self, node )
         do
            local thenBlock = node:get_thenBlock()
            if thenBlock ~= nil then
               self:writeln( "} else {" )
               self:pushIndent(  )
               setVals(  )
               self:popIndent(  )
               filter( thenBlock, self, node )
               self:writeln( "}" )
            else
               self:writeln( "}" )
               setVals(  )
            end
         end
         
         self:popIndent(  )
         
         self:writeln( "}" )
      else 
         
            Util.err( string.format( "not support -- %s", __func__) )
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
      if symPair:get_dst():get_posForModToRef() then
         self:write( string.format( "%s_%d := %s", symPair:get_dst():get_name(), symPair:get_dst():get_symbolId(), self:getSymbolSym( symPair:get_src() )) )
         self:outputAny2Type( symPair:get_dst():get_typeInfo() )
         self:writeln( "" )
      end
      
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
   
   self:writeln( "" )
end


function convFilter:processDeclArg( node, opt )

   self:write( string.format( "%s ", self:getSymbolSym( node:get_symbolInfo() )) )
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

   self:write( string.format( "type %s ", self:getFuncSymbol( node:get_expType() )) )
   self:outputDeclFunc( _lune.newAlge( FuncInfo.Type, {node:get_expType()}) )
   self:writeln( "" )
   
end


function convFilter:processDeclFunc( node, opt )

   do
      local funcSym = node:get_declInfo():get_symbol()
      if funcSym ~= nil then
         if not funcSym:get_posForModToRef() and not Ast.isPubToExternal( funcSym:get_accessMode() ) then
            return 
         end
         
      end
   end
   
   local funcType = node:get_expType()
   if (self.processMode == ProcessMode.NonClosureFuncDecl ) == isClosure( node:get_expType() ) then
      if self.processMode ~= ProcessMode.NonClosureFuncDecl and not node:get_declInfo():get_symbol() then
         self:write( self:getFuncSymbol( funcType ) )
      end
      
      return 
   end
   
   if isClosure( funcType ) then
      do
         local funcSym = node:get_declInfo():get_symbol()
         if funcSym ~= nil then
            self:write( "var " )
            self:outputSymbol( _lune.newAlge( SymbolKind.Func, {funcType}), funcSym:get_name() )
            self:write( " " )
            self:outputDeclFunc( _lune.newAlge( FuncInfo.DeclInfo, {node,node:get_declInfo()}) )
            self:writeln( "" )
            self:outputSymbol( _lune.newAlge( SymbolKind.Func, {funcType}), funcSym:get_name() )
            self:write( " = " )
         end
      end
      
   end
   
   self:outputDeclFuncInfo( node, node:get_declInfo() )
end


function convFilter:processRefType( node, opt )

   self:write( self:type2gotype( node:get_expType() ) )
end


function convFilter:processCond( node, parent )

   if node:get_expType() ~= Ast.builtinTypeBool then
      self:write( "Lns_isCondTrue( " )
      filter( node, self, parent )
      self:write( ")" )
   else
    
      filter( node, self, parent )
   end
   
end


function convFilter:processIf( node, opt )

   for __index, stmt in ipairs( node:get_stmtList() ) do
      do
         local _switchExp = stmt:get_kind()
         if _switchExp == Nodes.IfKind.If then
            self:write( "if " )
            self:processCond( stmt:get_exp(), node )
            self:writeln( "{" )
            filter( stmt:get_block(), self, node )
         elseif _switchExp == Nodes.IfKind.ElseIf then
            self:write( "} else if " )
            self:processCond( stmt:get_exp(), node )
            self:writeln( "{" )
            filter( stmt:get_block(), self, node )
         elseif _switchExp == Nodes.IfKind.Else then
            self:writeln( "} else { " )
            filter( stmt:get_block(), self, node )
         end
      end
      
   end
   
   self:writeln( "}" )
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

   local function hasAccessing(  )
   
      for __index, caseInfo in ipairs( node:get_caseList() ) do
         for _6037, symbol in ipairs( caseInfo:get_valParamNameList() ) do
            if symbol:get_posForModToRef() then
               return true
            end
            
         end
         
      end
      
      return false
   end
   local val
   
   if hasAccessing(  ) then
      val = string.format( "_exp%d", node:get_id())
      self:write( string.format( "switch %s := ", val) )
   else
    
      val = ""
      self:write( "switch " )
   end
   
   filter( node:get_val(), self, node )
   self:writeln( ".(type) {" )
   for __index, caseInfo in ipairs( node:get_caseList() ) do
      self:writeln( string.format( "case *%s:", self:getAlgeSymbol( caseInfo:get_valInfo() )) )
      for index, symbol in ipairs( caseInfo:get_valParamNameList() ) do
         if symbol:get_posForModToRef() then
            self:writeln( string.format( "%s := %s.Val%d", self:getSymbolSym( symbol ), val, index) )
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
   if not node:get_infinit() then
      self:processCond( node:get_exp(), node )
   end
   
   self:writeln( " {" )
   
   filter( node:get_block(), self, node )
   
   self:writeln( "}" )
end


function convFilter:processRepeat( node, opt )

   self:writeln( "for {" )
   
   filter( node:get_block(), self, node )
   
   self:pushIndent(  )
   
   self:write( "if " )
   self:processCond( node:get_exp(), node )
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

   self:writeln( "{" )
   self:pushIndent(  )
   
   local formSym = string.format( "_form%d", node:get_id())
   local paramSym = string.format( "_param%d", node:get_id())
   local prevSym = string.format( "_prev%d", node:get_id())
   
   if node:get_expList():get_mRetExp() then
      self:write( string.format( "%s, %s, %s := ", formSym, paramSym, prevSym) )
      filter( node:get_expList(), self, node )
      self:writeln( "" )
   else
    
      self:write( string.format( "%s, %s, %s := ", formSym, paramSym, prevSym) )
      filter( node:get_expList():get_expList()[1], self, node:get_expList() )
      self:write( "," )
      filter( node:get_expList():get_expList()[2], self, node:get_expList() )
      self:write( ", LnsAny(" )
      filter( node:get_expList():get_expList()[3], self, node:get_expList() )
      self:writeln( ")" )
   end
   
   
   self:writeln( "for {" )
   self:pushIndent(  )
   
   local setTxt = prevSym
   for index = 2, #node:get_varList() do
      local symInfo = node:get_varList()[index]
      self:writeln( string.format( "var %s %s", self:getSymbolSym( symInfo ), self:type2gotype( symInfo:get_typeInfo() )) )
      setTxt = string.format( "%s, %s", setTxt, self:getSymbolSym( symInfo ))
   end
   
   if node:get_expList():get_expType():get_kind() == Ast.TypeInfoKind.Ext then
      local workSym = string.format( "_work%d", node:get_id())
      self:writeln( string.format( "%s := %s.(*Lns_luaValue).Call( Lns_2DDD( %s, %s ) )", workSym, formSym, paramSym, prevSym) )
      self:write( string.format( "%s = ", setTxt) )
      for index, _6080 in ipairs( node:get_varList() ) do
         if index > 1 then
            self:write( "," )
         end
         
         self:write( string.format( "Lns_getFromMulti(%s,%d)", workSym, index - 1) )
      end
      
      self:writeln( "" )
   else
    
      self:writeln( string.format( "%s = %s( %s, %s )", setTxt, formSym, paramSym, prevSym) )
   end
   
   self:writeln( string.format( "if Lns_IsNil( %s ) { break }", prevSym) )
   local topSymInfo = node:get_varList()[1]
   if topSymInfo:get_name() ~= "_" then
      self:writeln( string.format( "%s := %s.(%s)", self:getSymbolSym( topSymInfo ), prevSym, self:type2gotype( topSymInfo:get_typeInfo() )) )
   end
   
   self:popIndent(  )
   
   filter( node:get_block(), self, node )
   
   self:writeln( "}" )
   
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:outputForeachLua( node, extType )
   local __func__ = '@lune.@base.@convGo.convFilter.outputForeachLua'

   do
      local _switchExp = extType:get_extedType():get_kind()
      if _switchExp == Ast.TypeInfoKind.List then
         self:writeln( "{" )
         self:pushIndent(  )
         self:write( "_index, _val := " )
         filter( node:get_exp(), self, node )
         self:writeln( ".Get1stFromMap()" )
         self:writeln( "for _index != nil {" )
         self:pushIndent(  )
         do
            local keySym = node:get_key()
            if keySym ~= nil then
               self:write( string.format( "%s := _index", self:getSymbolSym( keySym )) )
               self:outputAny2Type( keySym:get_typeInfo() )
               self:writeln( "" )
            end
         end
         
         self:write( string.format( "%s := _val", self:getSymbolSym( node:get_val() )) )
         self:outputAny2Type( node:get_val():get_typeInfo() )
         self:writeln( "" )
         self:popIndent(  )
         
         filter( node:get_block(), self, node )
         
         self:pushIndent(  )
         self:write( "_index, _val = " )
         filter( node:get_exp(), self, node )
         self:writeln( ".NextFromMap( _index )" )
         
         self:popIndent(  )
         
         self:writeln( "}" )
         self:popIndent(  )
         self:writeln( "}" )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( "{" )
         self:pushIndent(  )
         self:write( "_key, _val := " )
         filter( node:get_exp(), self, node )
         self:writeln( ".Get1stFromMap()" )
         self:writeln( "for _key != nil {" )
         self:pushIndent(  )
         do
            local keySym = node:get_key()
            if keySym ~= nil then
               self:write( string.format( "%s := _key", self:getSymbolSym( keySym )) )
               self:outputAny2Type( keySym:get_typeInfo() )
               self:writeln( "" )
            end
         end
         
         self:write( string.format( "%s := _val", self:getSymbolSym( node:get_val() )) )
         self:outputAny2Type( node:get_val():get_typeInfo() )
         self:writeln( "" )
         self:popIndent(  )
         
         filter( node:get_block(), self, node )
         
         self:pushIndent(  )
         self:write( "_key, _val = " )
         filter( node:get_exp(), self, node )
         self:writeln( ".NextFromMap( _key )" )
         
         self:popIndent(  )
         
         self:writeln( "}" )
         self:popIndent(  )
         self:writeln( "}" )
      else 
         
            Util.err( string.format( "not support -- %s", __func__) )
      end
   end
   
end


function convFilter:processForeach( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processForeach'

   
   do
      local extType = _lune.__Cast( node:get_exp():get_expType():get_srcTypeInfo(), 3, Ast.ExtTypeInfo )
      if extType ~= nil then
         self:outputForeachLua( node, _lune.unwrap( _lune.__Cast( extType:get_srcTypeInfo(), 3, Ast.ExtTypeInfo )) )
         return 
      end
   end
   
   
   
   
   
   self:write( "for " )
   local loopExpType = node:get_exp():get_expType()
   do
      local _switchExp = loopExpType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
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
         
         
         local valName = self:getSymbolSym( node:get_val() )
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
               self:writeln( string.format( "%s := _%s + 1", self:getSymbolSym( key ), key:get_name()) )
            end
         end
         
         
         if valName ~= "_" then
            self:write( string.format( "%s := _%s", valName, valName) )
            self:outputAny2Type( itemType )
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
         
         
         local valName = self:getSymbolSym( node:get_val() )
         local itemType = loopExpType:get_itemTypeInfoList()[2]
         
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
               
               if key:get_name() ~= "_" then
                  self:write( string.format( "%s := _%s", key:get_name(), key:get_name()) )
                  self:outputAny2Type( keyType )
                  self:writeln( "" )
               end
               
               
            end
         end
         
         
         if valName ~= "_" then
            self:write( string.format( "%s := _%s", valName, valName) )
            self:outputAny2Type( itemType )
            self:writeln( "" )
         end
         
         
         self:popIndent(  )
      elseif _switchExp == Ast.TypeInfoKind.Set then
         local valType = loopExpType:get_itemTypeInfoList()[1]
         local valName = self:getSymbolSym( node:get_val() )
         
         if valName ~= "_" then
            self:write( string.format( "_%s", valName) )
         else
          
            self:write( string.format( "%s", valName) )
         end
         
         
         self:write( " := range( " )
         filter( node:get_exp(), self, node )
         self:writeln( ".Items ) {" )
         self:pushIndent(  )
         
         if valName ~= "_" then
            self:write( string.format( "%s := _%s", valName, valName) )
            self:outputAny2Type( valType )
            self:writeln( "" )
         end
         
         
         self:popIndent(  )
      else 
         
            Util.err( string.format( "not support -- %s", __func__) )
      end
   end
   
   filter( node:get_block(), self, node )
   self:write( "}" )
   self:writeln( "" )
end


local type2LnsItemKindMap = {[Ast.builtinTypeInt] = "LnsItemKindInt", [Ast.builtinTypeReal] = "LnsItemKindReal", [Ast.builtinTypeString] = "LnsItemKindStr"}

local function getLnsItemKind( typeInfo )

   do
      local kind = type2LnsItemKindMap[typeInfo]
      if kind ~= nil then
         return kind
      end
   end
   
   return "LnsItemKindStem"
end

function convFilter:processForsort( node, opt )

   local keySym
   
   local valSym
   
   local keyTypeInfo = node:get_exp():get_expType():get_itemTypeInfoList()[1]
   if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Set then
      keySym = node:get_val()
      valSym = node:get_key()
   else
    
      keySym = node:get_key()
      valSym = node:get_val()
   end
   
   
   self:writeln( "{" )
   self:pushIndent(  )
   local collSym = string.format( "__collection%d", node:get_id())
   self:write( string.format( "%s := ", collSym) )
   filter( node:get_exp(), self, node )
   self:writeln( "" )
   local sortSym = string.format( "__sorted%d", node:get_id())
   self:write( string.format( "%s := %s.", sortSym, collSym) )
   do
      local _switchExp = keyTypeInfo
      if _switchExp == Ast.builtinTypeInt then
         self:writeln( "CreateKeyListInt()" )
      elseif _switchExp == Ast.builtinTypeReal then
         self:writeln( "CreateKeyListReal()" )
      elseif _switchExp == Ast.builtinTypeString then
         self:writeln( "CreateKeyListStr()" )
      else 
         
            self:writeln( "CreateKeyListStem()" )
      end
   end
   
   self:writeln( string.format( "%s.Sort( %s, nil )", sortSym, getLnsItemKind( keyTypeInfo )) )
   
   self:write( "for _, " )
   local key = string.format( "__key%d", node:get_id())
   if keySym ~= nil then
      key = string.format( "%s", self:getSymbolSym( keySym ))
   end
   
   self:write( string.format( "_%s", key) )
   self:writeln( string.format( " := range( %s.Items ) {", sortSym) )
   self:pushIndent(  )
   if valSym ~= nil then
      if valSym:get_posForModToRef() then
         self:write( string.format( "%s := %s.Items[ _%s ]", self:getSymbolSym( valSym ), collSym, key) )
         self:outputAny2Type( valSym:get_typeInfo() )
         self:writeln( "" )
      end
      
   end
   
   if keySym ~= nil then
      if keySym:get_posForModToRef() then
         self:write( string.format( "%s := _%s", key, key) )
         self:outputAny2Type( keySym:get_typeInfo() )
         self:writeln( "" )
      end
      
   end
   
   self:popIndent(  )
   
   filter( node:get_block(), self, node )
   self:writeln( "}" )
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processExpUnwrap( node, opt )

   do
      local def = node:get_default()
      if def ~= nil then
         self:write( "Lns_unwrapDefault( " )
         filter( node:get_exp(), self, node )
         self:write( ", " )
         filter( def, self, node )
      else
         self:write( "Lns_unwrap( " )
         filter( node:get_exp(), self, node )
      end
   end
   
   self:write( ")" )
   self:outputAny2Type( node:get_expType() )
end


function convFilter:processExpToDDD( node, opt )

   if node:get_expList():get_mRetExp() then
      filter( node:get_expList(), self, node )
   else
    
      self:write( "[]LnsAny{ " )
      filter( node:get_expList(), self, node )
      self:write( "}" )
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
   
   local typeName = self:type2gotype( node:get_expType() )
   self:writeln( string.format( "func Lns_cast2%s( obj LnsAny ) LnsAny {", typeName) )
   self:writeln( string.format( "    if _, ok := obj.(%s); ok { ", typeName) )
   self:writeln( "        return obj" )
   self:writeln( "    }" )
   self:writeln( "    return nil" )
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
   end
   
   self:write( "FP " )
   self:write( self:getTypeSymbol( node:get_expType() ) )
   self:writeln( "Mtd" )
   
   self:popIndent(  )
   
   self:writeln( "}" )
end


function convFilter:outputToStem( node )

   self:writeln( string.format( "func %s2Stem( obj LnsAny ) LnsAny {", self:getTypeSymbol( node:get_expType() )) )
   self:pushIndent(  )
   self:writeln( "if obj == nil {" )
   self:writeln( "    return nil" )
   self:writeln( "}" )
   self:writeln( string.format( "return obj.(%s).FP", self:type2gotype( node:get_expType() )) )
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


function convFilter:outputNewSetup( objName, classType )

   local className = self:getTypeSymbol( classType )
   self:writeln( string.format( "%s := &%s{}", objName, className) )
   self:writeln( string.format( "%s.FP = %s", objName, objName) )
   
   do
      local workType = classType
      while workType:hasBase(  ) do
         workType = workType:get_baseTypeInfo()
         
         local superName = self:getTypeSymbol( workType )
         self:writeln( string.format( "%s.%s.FP = %s", objName, superName, objName) )
      end
      
   end
   
end


function convFilter:outputConstructor( node )

   local scope = _lune.unwrap( node:get_expType():get_scope())
   local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, Ast.ScopeAccess.Normal ))
   
   local className = self:getTypeSymbol( node:get_expType() )
   local ctorName = self:getConstrSymbol( node:get_expType() )
   
   if not node:get_expType():get_abstractFlag() then
      self:write( string.format( "func New%s(", className) )
      self:outputDeclFuncArg( initFuncType )
      self:writeln( string.format( ") *%s {", className) )
      self:pushIndent(  )
      self:outputNewSetup( "obj", node:get_expType() )
      self:write( string.format( "obj.%s(", ctorName) )
      for index, _6214 in ipairs( initFuncType:get_argTypeInfoList() ) do
         if index ~= 1 then
            self:write( ", " )
         end
         
         self:write( string.format( "arg%d", index) )
      end
      
      self:writeln( ")" )
      self:writeln( "return obj" )
      self:popIndent(  )
      
      self:writeln( "}" )
   end
   
   
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
         local superName = self:getTypeSymbol( superType )
         self:write( string.format( "self.%s.%s( ", superName, self:getConstrSymbol( superType )) )
         for index = 1, #baseInitFuncType:get_argTypeInfoList() do
            if index ~= 1 then
               self:write( "," )
            end
            
            if node:get_hasOldCtor() then
               self:write( "nil" )
            else
             
               self:write( string.format( "arg%d", index) )
            end
            
         end
         
         self:writeln( ")" )
         if node:get_hasOldCtor() then
            superArgNum = 0
         else
          
            superArgNum = #baseInitFuncType:get_argTypeInfoList()
         end
         
      else
       
         superArgNum = 0
      end
      
      local argIndex = superArgNum + 1
      for __index, memberNode in ipairs( node:get_memberList() ) do
         if not memberNode:get_staticFlag() then
            self:writeln( string.format( "self.%s = arg%d", self:getSymbol( _lune.newAlge( SymbolKind.Member, {Ast.isPubToExternal( memberNode:get_accessMode() )}), memberNode:get_name().txt ), argIndex) )
            argIndex = argIndex + 1
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
         
         local retType = getterSym:get_typeInfo():get_retTypeInfoList()[1]:get_srcTypeInfo()
         self:write( "{ return " )
         
         local prefix
         
         if memberSym:get_staticFlag() then
            prefix = ""
         else
          
            prefix = "self."
         end
         
         
         if retType ~= memberSym:get_typeInfo():get_srcTypeInfo() then
            self:write( string.format( "&%s%s.%s", prefix, self:getSymbolSym( memberSym ), self:getTypeSymbol( retType )) )
         else
          
            self:write( string.format( "%s%s", prefix, self:getSymbolSym( memberSym )) )
         end
         
         self:writeln( " }" )
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


local type2FromStemNameMap = {[Ast.builtinTypeInt] = "Lns_ToInt", [Ast.builtinTypeReal] = "Lns_ToReal", [Ast.builtinTypeBool] = "Lns_ToBool", [Ast.builtinTypeString] = "Lns_ToStr", [Ast.builtinTypeStem] = "Lns_ToStem"}
function convFilter:getFromStemName( typeInfo )
   local __func__ = '@lune.@base.@convGo.convFilter.getFromStemName'

   local workTypeInfo = typeInfo:get_srcTypeInfo():get_nonnilableType()
   do
      local name = type2FromStemNameMap[workTypeInfo]
      if name ~= nil then
         return name
      end
   end
   
   do
      local _switchExp = workTypeInfo:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         return "Lns_ToList"
      elseif _switchExp == Ast.TypeInfoKind.Set then
         return "Lns_ToSet"
      elseif _switchExp == Ast.TypeInfoKind.Map then
         return "Lns_ToLnsMap"
      elseif _switchExp == Ast.TypeInfoKind.Class then
         return string.format( "%s_FromMap", self:getTypeSymbol( workTypeInfo ))
      else 
         
            Util.err( string.format( "%s: not support -- %s", __func__, Ast.TypeInfoKind:_getTxt( workTypeInfo:get_kind())
            ) )
      end
   end
   
end




function convFilter:outputConvItemType( typeInfo, alt2type )

   local workTypeInfo = typeInfo:get_srcTypeInfo():get_nonnilableType()
   if typeInfo:get_srcTypeInfo():get_nonnilableType():get_kind() == Ast.TypeInfoKind.Alternate then
      if alt2type ~= nil then
         do
            local alt = alt2type[workTypeInfo]
            if alt ~= nil then
               workTypeInfo = alt
            end
         end
         
      end
      
   end
   
   
   do
      local altType = _lune.__Cast( workTypeInfo, 3, Ast.AlternateTypeInfo )
      if altType ~= nil then
         self:write( string.format( 'paramList[%d]', altType:get_altIndex() - 1) )
      else
         self:writeln( "Lns_ToObjParam{" )
         self:pushIndent(  )
         self:write( string.format( "%sSub, %s,", self:getFromStemName( workTypeInfo ), typeInfo:get_nilable()) )
         self:outputConvItemTypeList( workTypeInfo:get_itemTypeInfoList(), alt2type )
         self:popIndent(  )
         self:write( "}" )
      end
   end
   
end


function convFilter:outputConvItemTypeList( itemTypeInfoList, alt2type )

   if #itemTypeInfoList > 0 then
      self:write( "[]Lns_ToObjParam{" )
      self:pushIndent(  )
      for index, itemType in ipairs( itemTypeInfoList ) do
         if index > 1 then
            self:write( "," )
         end
         
         
         self:outputConvItemType( itemType, alt2type )
      end
      
      self:popIndent(  )
      self:write( "}" )
   else
    
      self:write( "nil" )
   end
   
end


function convFilter:outputAlter2MapFunc( alt2Map )
   local __func__ = '@lune.@base.@convGo.convFilter.outputAlter2MapFunc'

   self:write( "{" )
   
   for altType, assinType in pairs( alt2Map ) do
      if altType:get_kind() == Ast.TypeInfoKind.Alternate then
         if assinType:get_kind() == Ast.TypeInfoKind.Alternate then
            Util.err( string.format( "not support: %s", __func__) )
         else
          
            self:outputConvItemType( assinType, alt2Map )
         end
         
      end
      
   end
   
   
   self:write( "}" )
end


function convFilter:outputMapping( node )

   local classType = node:get_expType()
   local className = self:getTypeSymbol( classType )
   self:writeln( string.format( "func (self *%s) ToMapSetup( obj *LnsMap ) *LnsMap {", className) )
   self:pushIndent(  )
   for __index, memberNode in ipairs( node:get_memberList() ) do
      if not memberNode:get_staticFlag() then
         self:writeln( string.format( 'obj.Items["%s"] = Lns_ToCollection( self.%s )', self:getSymbolSym( memberNode:get_symbolInfo() ), self:getSymbolSym( memberNode:get_symbolInfo() )) )
      end
      
   end
   
   if classType:hasBase(  ) then
      self:writeln( string.format( "return self.%s.ToMapSetup( obj )", self:getTypeSymbol( classType:get_baseTypeInfo() )) )
   else
    
      self:writeln( "return obj" )
   end
   
   self:popIndent(  )
   self:writeln( "}" )
   self:writeln( string.format( "func (self *%s) ToMap() *LnsMap {", className) )
   self:writeln( "    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )" )
   self:writeln( "}" )
   
   local fromStemName = self:getFromStemName( classType )
   
   local classScope = _lune.unwrap( classType:get_scope())
   do
      local fromMapSym = _lune.unwrap( classScope:getSymbolInfoChild( "_fromMap" ))
      self:writeln( string.format( "func %s(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){", self:getSymbolSym( fromMapSym )) )
      self:writeln( string.format( "   return %s( arg1, paramList )", fromStemName) )
      self:writeln( "}" )
   end
   
   do
      local fromStemSym = _lune.unwrap( classScope:getSymbolInfoChild( "_fromStem" ))
      self:writeln( string.format( "func %s(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){", self:getSymbolSym( fromStemSym )) )
      self:writeln( string.format( "   return %s( arg1, paramList )", fromStemName) )
      self:writeln( "}" )
   end
   
   
   self:writeln( string.format( "func %s( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {", fromStemName) )
   self:pushIndent(  )
   self:writeln( string.format( "_,conv,mess := %sSub(obj,false, paramList);", fromStemName) )
   self:writeln( "return conv,mess" )
   self:popIndent(  )
   self:writeln( "}" )
   
   self:writeln( string.format( "func %sSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {", fromStemName) )
   self:pushIndent(  )
   self:writeln( "var objMap *LnsMap" )
   self:writeln( "if work, ok := obj.(*LnsMap); !ok {" )
   self:writeln( '   return false, nil, "no map"' )
   self:writeln( "} else {" )
   self:writeln( '   objMap = work' )
   self:writeln( "}" )
   self:outputNewSetup( "newObj", classType )
   self:writeln( string.format( "return %sMain( newObj, objMap, paramList )", fromStemName) )
   self:popIndent(  )
   self:writeln( "}" )
   
   self:writeln( string.format( "func %sMain( newObj %s, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {", fromStemName, self:type2gotype( classType )) )
   self:pushIndent(  )
   
   if classType:hasBase(  ) then
      self:writeln( string.format( "if ok,_,mess := %sMain( &newObj.%s, objMap, paramList ); !ok {", self:getFromStemName( classType:get_baseTypeInfo() ), self:getTypeSymbol( classType:get_baseTypeInfo() )) )
      self:writeln( "    return false,nil,mess" )
      self:writeln( "}" )
   end
   
   
   for __index, memberNode in ipairs( node:get_memberList() ) do
      if not memberNode:get_staticFlag() then
         local memberName = self:getSymbolSym( memberNode:get_symbolInfo() )
         self:write( "if ok,conv,mess := " )
         if memberNode:get_expType():get_nonnilableType():get_kind() == Ast.TypeInfoKind.Alternate then
            for index, itemType in ipairs( classType:get_itemTypeInfoList() ) do
               if itemType == memberNode:get_expType():get_srcTypeInfo() then
                  self:write( string.format( 'paramList[%d].Func( objMap.Items["%s"], %s, paramList[%d].Child', index - 1, memberName, memberNode:get_expType():get_nilable(), index - 1) )
               end
               
            end
            
         else
          
            self:write( string.format( '%sSub( objMap.Items["%s"], %s, ', self:getFromStemName( memberNode:get_expType() ), memberName, memberNode:get_expType():get_nilable()) )
            self:outputConvItemTypeList( memberNode:get_expType():get_itemTypeInfoList(), nil )
         end
         
         self:writeln( "); !ok {" )
         self:writeln( string.format( '   return false,nil,"%s:" + mess.(string)', memberName) )
         self:writeln( "} else {" )
         self:write( string.format( "   newObj.%s = conv", memberName) )
         self:outputAny2Type( memberNode:get_expType() )
         self:writeln( "" )
         self:writeln( "}" )
      end
      
   end
   
   self:writeln( "return true, newObj, nil" )
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:outputAdvertise( node )
   local __func__ = '@lune.@base.@convGo.convFilter.outputAdvertise'

   local methodNameSet = node:createMethodNameSetWithoutAdv(  )
   for __index, adv in ipairs( node:get_advertiseList() ) do
      if adv:get_prefix() ~= "" then
         Util.err( string.format( "%s: not support advertise with prefix", __func__) )
      end
      
      do
         local scope = adv:get_member():get_expType():get_scope()
         if scope ~= nil then
            scope:filterTypeInfoField( true, scope, Ast.ScopeAccess.Normal, function ( symbol )
            
               if symbol:get_kind() == Ast.SymbolKind.Mtd and symbol:get_name() ~= "__init" and not _lune._Set_has(methodNameSet, symbol:get_name() ) and not symbol:get_staticFlag() then
                  local funcType = symbol:get_typeInfo()
                  self:outputDeclFunc( _lune.newAlge( FuncInfo.Advertise, {node:get_expType(),funcType}) )
                  self:writeln( " {" )
                  if #funcType:get_retTypeInfoList() > 0 then
                     self:write( "    return " )
                  end
                  
                  self:write( string.format( "self.%s.FP.%s( ", self:getSymbolSym( adv:get_member():get_symbolInfo() ), self:getSymbolSym( symbol )) )
                  for index, _6304 in ipairs( funcType:get_argTypeInfoList() ) do
                     if index > 1 then
                        self:write( "," )
                     end
                     
                     self:write( string.format( "arg%d", index) )
                  end
                  
                  self:writeln( ")" )
                  self:writeln( "}" )
               end
               
               return true
            end )
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
            self:outputToStem( node )
            self:outputDownCast( node )
            self:outputCastReceiver( node )
            self:outputConstructor( node )
            self:outputAccessor( node )
            self:outputAdvertise( node )
            
            if node:get_expType():isInheritFrom( Ast.builtinTypeMapping, nil ) then
               self:outputMapping( node )
            end
            
            
            for __index, fieldNode in ipairs( node:get_fieldList() ) do
               if _lune.__Cast( fieldNode, 3, Nodes.DeclMemberNode ) then
                  
               else
                
                  filter( fieldNode, self, node )
                  self:writeln( "" )
               end
               
            end
            
         elseif _switchExp == Ast.TypeInfoKind.IF then
            self:outputInterfaceType( node )
         else 
            
               Util.err( string.format( "%s: not support -- %s:%d", __func__, Ast.TypeInfoKind:_getTxt( node:get_expType():get_kind())
               , node:get_pos().lineNo) )
         end
      end
      
   else
    
      self:outputStaticMember( node )
   end
   
end


local CallKind = {}
CallKind._name2Val = {}
function CallKind:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "CallKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function CallKind._from( val )
   return _lune._AlgeFrom( CallKind, val )
end

CallKind.BuiltinCall = { "BuiltinCall", {{ func=_lune._toStr, nilable=false, child={} },{ func=_lune._toStr, nilable=false, child={} }}}
CallKind._name2Val["BuiltinCall"] = CallKind.BuiltinCall
CallKind.FormCall = { "FormCall"}
CallKind._name2Val["FormCall"] = CallKind.FormCall
CallKind.LuaCall = { "LuaCall"}
CallKind._name2Val["LuaCall"] = CallKind.LuaCall
CallKind.Normal = { "Normal"}
CallKind._name2Val["Normal"] = CallKind.Normal
CallKind.RunLoaded = { "RunLoaded"}
CallKind._name2Val["RunLoaded"] = CallKind.RunLoaded
CallKind.RuntimeCall = { "RuntimeCall", {{ func=Nodes.Node._fromMap, nilable=false, child={} }}}
CallKind._name2Val["RuntimeCall"] = CallKind.RuntimeCall
CallKind.SortCall = { "SortCall", {{ func=Ast.TypeInfo._fromMap, nilable=false, child={} }}}
CallKind._name2Val["SortCall"] = CallKind.SortCall

function convFilter:outputCallPrefix( callId, node, prefixNode, funcSymbol )

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
                  for _6336 = 2, retNum do
                     anys = string.format( "%s,LnsAny", anys)
                  end
                  
                  self:write( string.format( "Lns_NilAccCall%d( func () (%s) { return ", retNum, anys) )
               else
                
                  local args = "LnsAny"
                  for _6338 = 2, retNum do
                     args = string.format( "%s,LnsAny", args)
                  end
                  
                  self:write( string.format( "lns_NilAccCall_%s( func () (%s) { return ", nilAccName, args) )
               end
               
         end
      end
      
      self:write( string.format( "Lns_NilAccPop().(%s)", self:type2gotype( workPrefixNode:get_expType():get_nonnilableType() )) )
   end
   
   local callKind = _lune.newAlge( CallKind.Normal)
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
      
      
      local prefixType = prefixNode:get_expType():get_nonnilableType()
      
      if prefixType == Ast.builtinTypeString then
         if node:hasNilAccess(  ) then
            Util.err( "not support nilAccName" )
         end
         
         do
            local runtime = self.builtin2runtime[funcType]
            if runtime ~= nil then
               callKind = _lune.newAlge( CallKind.RuntimeCall, {prefixNode})
               self:write( runtime )
            end
         end
         
      else
       
         if not funcType:get_staticFlag() then
            if node:hasNilAccess(  ) then
               if not prefixNode:hasNilAccess(  ) then
                  self:write( "Lns_NilAccFin(" )
                  self:write( "Lns_NilAccPush(" )
                  filter( prefixNode, self, node )
                  self:writeln( ") && " )
               else
                
                  filter( prefixNode, self, node )
                  do
                     local _switchExp = prefixNode:get_kind()
                     if _switchExp == Nodes.NodeKind.get_RefField() then
                     else 
                        
                           self:writeln( "&&" )
                     end
                  end
                  
               end
               
            else
             
               filter( prefixNode, self, node )
            end
            
         end
         
         
         processNilAcc( prefixNode )
         
         if prefixType:get_kind() == Ast.TypeInfoKind.Ext then
            self:write( string.format( '.CallMethod( "%s", Lns_2DDD', funcSymbol:get_name()) )
            callKind = _lune.newAlge( CallKind.LuaCall)
         else
          
            local prefixKind
            
            if prefixType:get_kind() == Ast.TypeInfoKind.Alternate and prefixType:hasBase(  ) then
               prefixKind = prefixType:get_baseTypeInfo():get_kind()
            else
             
               prefixKind = prefixType:get_kind()
            end
            
            
            if Ast.isBuiltin( funcType:get_typeId() ) then
               do
                  local _switchExp = prefixKind
                  if _switchExp == Ast.TypeInfoKind.Class then
                     self:write( string.format( ".FP.%s", self:getSymbolSym( funcSymbol )) )
                  else 
                     
                        local builtinFuncs = TransUnit.getBuiltinFunc(  )
                        do
                           local runtime = self.builtin2runtime[funcType]
                           if runtime ~= nil then
                              self:write( runtime )
                           else
                              do
                                 local _switchExp = funcType
                                 if _switchExp == builtinFuncs.list_sort or _switchExp == builtinFuncs.array_sort then
                                    callKind = _lune.newAlge( CallKind.SortCall, {prefixType:get_itemTypeInfoList()[1]})
                                 end
                              end
                              
                              self:write( string.format( ".%s", self:getSymbolSym( funcSymbol )) )
                           end
                        end
                        
                  end
               end
               
            else
             
               do
                  local _switchExp = funcType:get_kind()
                  if _switchExp == Ast.TypeInfoKind.Method then
                     do
                        local _switchExp = prefixKind
                        if _switchExp == Ast.TypeInfoKind.Class then
                           self:write( string.format( ".FP.%s", self:getSymbolSym( funcSymbol )) )
                        else 
                           
                              self:write( string.format( ".%s", self:getSymbolSym( funcSymbol )) )
                        end
                     end
                     
                  else 
                     
                        self:write( self:getSymbolSym( funcSymbol ) )
                  end
               end
               
            end
            
         end
         
      end
      
   end
   
   
   return closeParen, callKind
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
   local callKind = _lune.newAlge( CallKind.Normal)
   do
      local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
      if fieldNode ~= nil then
         closeParen, callKind = self:outputCallPrefix( node:get_id(), fieldNode, fieldNode:get_prefix(), _lune.unwrap( fieldNode:get_symbolInfo()) )
      else
         if Ast.isBuiltin( funcType:get_typeId() ) then
            local builtinFuncs = TransUnit.getBuiltinFunc(  )
            
            do
               local runtime = self.builtin2runtime[funcType]
               if runtime ~= nil then
                  self:write( runtime )
               else
                  do
                     local _switchExp = funcType:get_srcTypeInfo()
                     if _switchExp == Ast.builtinTypeForm then
                        filter( node:get_func(), self, node )
                        callKind = _lune.newAlge( CallKind.FormCall)
                     else 
                        
                           Util.err( string.format( "%s: not support -- %s:%d", __func__, funcType:getTxt(  ), node:get_pos().lineNo) )
                     end
                  end
                  
               end
            end
            
         else
          
            if funcType:get_kind() == Ast.TypeInfoKind.Ext then
               self:write( "Lns_getVM().RunLoadedfunc" )
               callKind = _lune.newAlge( CallKind.RunLoaded)
            else
             
               filter( node:get_func(), self, node )
            end
            
         end
         
      end
   end
   
   
   self:write( "(" )
   
   local closeTxt = nil
   do
      local _matchExp = callKind
      if _matchExp[1] == CallKind.RuntimeCall[1] then
         local prefixNode = _matchExp[2][1]
      
         filter( prefixNode, self, node:get_func() )
         if node:get_argList() then
            self:write( "," )
         end
         
      elseif _matchExp[1] == CallKind.FormCall[1] then
      
         self:write( "Lns_2DDD(" )
      elseif _matchExp[1] == CallKind.RunLoaded[1] then
      
         filter( node:get_func(), self, node )
         self:write( "," )
         if not node:get_argList() then
            self:write( "[]LnsAny{}" )
         end
         
      elseif _matchExp[1] == CallKind.SortCall[1] then
         local typeInfo = _matchExp[2][1]
      
         self:write( string.format( "%s, ", getLnsItemKind( typeInfo )) )
      elseif _matchExp[1] == CallKind.BuiltinCall[1] then
         local packName = _matchExp[2][1]
         local funcname = _matchExp[2][2]
      
         closeTxt = "}"
         self:write( string.format( '"%s", "%s"', packName, funcname) )
         if node:get_argList() then
            self:write( ", []LnsAny{" )
         end
         
      elseif _matchExp[1] == CallKind.LuaCall[1] then
      
         closeTxt = ")"
      end
   end
   
   
   do
      local argList = node:get_argList()
      if argList ~= nil then
         self:processSetFromExpList( getConvExpName( node:get_id(), argList ), funcType:get_argTypeInfoList(), argList )
      end
   end
   
   
   if funcType:get_kind() == Ast.TypeInfoKind.Func and funcType:get_staticFlag() and funcType:get_parentInfo():isInheritFrom( Ast.builtinTypeMapping ) and (funcType:get_rawTxt() == "_fromMap" or funcType:get_rawTxt() == "_fromStem" ) then
      local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
      if  nil == fieldNode then
         local _fieldNode = fieldNode
      
         Util.err( "not support -- __func__" )
      end
      
      self:write( "," )
      
      self:outputConvItemTypeList( funcType:get_parentInfo():get_itemTypeInfoList(), fieldNode:get_prefix():get_expType():createAlt2typeMap( false ) )
   end
   
   
   if closeTxt ~= nil then
      self:write( closeTxt )
   end
   
   self:write( ")" )
   if callKind == _lune.newAlge( CallKind.LuaCall) then
      if #funcType:get_retTypeInfoList() == 1 then
         self:write( "[0]" )
         self:outputAny2Type( funcType:get_retTypeInfoList()[1] )
      elseif #funcType:get_retTypeInfoList() > 1 then
         Util.err( string.format( "%s: not support", __func__) )
      end
      
   end
   
   
   if retGenerics then
      if #funcType:get_retTypeInfoList() == 1 then
         self:outputAny2Type( node:get_expType() )
      else
       
         self:write( ")" )
      end
      
   end
   
   
   if node:hasNilAccess(  ) then
      self:write( "})" )
      self:write( string.format( "/* %d:%d */", node:get_pos().lineNo, node:get_pos().column) )
      
      if opt.node:hasNilAccess(  ) then
         
      else
       
         self:write( ")" )
      end
      
      if closeParen then
         self:write( ")" )
      end
      
   end
   
   
   if callKind == _lune.newAlge( CallKind.FormCall) then
      self:write( ")" )
   end
   
end


function convFilter:processExpMRet( node, opt )

   filter( node:get_mRet(), self, node )
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
               if index == 1 or exp:get_expType():get_kind() == Ast.TypeInfoKind.DDD then
                  filter( exp, self, node )
               else
                
                  self:write( "Lns_2DDD(" )
                  filter( exp, self, node )
                  self:write( ")" )
               end
               
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
   self:outputAny2Type( node:get_expType() )
end


function convFilter:processExpCast( node, opt )

   do
      local _switchExp = node:get_castKind()
      if _switchExp == Nodes.CastKind.Force then
         if isAnyType( node:get_exp():get_expType() ) then
            do
               local _switchExp = node:get_castType()
               if _switchExp == Ast.builtinTypeInt then
                  self:write( "Lns_forceCastInt(" )
                  filter( node:get_exp(), self, node )
                  self:write( ")" )
               elseif _switchExp == Ast.builtinTypeReal then
                  self:write( "Lns_forceCastReal(" )
                  filter( node:get_exp(), self, node )
                  self:write( ")" )
               else 
                  
                     filter( node:get_exp(), self, node )
                     self:outputAny2Type( node:get_castType() )
               end
            end
            
         else
          
            self:write( string.format( "%s(", self:type2gotype( node:get_castType() )) )
            filter( node:get_exp(), self, node )
            self:write( ")" )
         end
         
      elseif _switchExp == Nodes.CastKind.Implicit then
         if #node:get_exp():get_expTypeList() > 1 then
            filter( node:get_exp(), self, node )
         else
          
            self:outputImplicitCast( node:get_castType(), node:get_exp(), node )
         end
         
      elseif _switchExp == Nodes.CastKind.Normal then
         local typeName = self:getTypeSymbol( node:get_castType() )
         local castType = node:get_castType():get_nonnilableType()
         if castType:get_kind() == Ast.TypeInfoKind.Class and castType ~= Ast.builtinTypeString then
            self:write( string.format( "%sDownCastF(", typeName) )
            filter( node:get_exp(), self, node )
            if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Class and node:get_exp():get_expType() ~= Ast.builtinTypeString then
               self:write( ".FP" )
            end
            
            self:write( ")" )
         else
          
            self:write( string.format( "Lns_cast2%s( ", self:type2gotype( castType )) )
            filter( node:get_exp(), self, node )
            self:write( ")" )
         end
         
      end
   end
   
end


function convFilter:processExpParen( node, opt )

   if #node:get_exp():get_expTypeList() >= 2 then
      self:write( "Lns_car(" )
      filter( node:get_exp(), self, node )
      self:write( ")" )
      self:outputAny2Type( node:get_expType() )
   else
    
      self:write( "(" )
      filter( node:get_exp(), self, node )
      self:write( ")" )
   end
   
end


function convFilter:processExpSetVal( node, opt )

   do
      local refItemNode = _lune.__Cast( node:get_exp1(), 3, Nodes.ExpRefItemNode )
      if refItemNode ~= nil then
         if refItemNode:get_val():get_expType():get_kind() == Ast.TypeInfoKind.Map then
            filter( refItemNode:get_val(), self, refItemNode )
            self:write( ".Set(" )
            do
               local indexNode = refItemNode:get_index()
               if indexNode ~= nil then
                  filter( indexNode, self, refItemNode )
               else
                  self:write( string.format( '"%s"', _lune.unwrap( refItemNode:get_symbol())) )
               end
            end
            
            self:write( "," )
            filter( node:get_exp2(), self, node )
            self:write( ")" )
            return 
         end
         
      end
   end
   
   filter( node:get_exp1(), self, node )
   for _6426 = #node:get_exp1():get_expTypeList() + 1, #node:get_exp2():get_expTypeList() do
      self:write( ",_" )
   end
   
   self:write( " = " )
   if node:get_exp2():get_expType():get_kind() == Ast.TypeInfoKind.DDD then
      self:write( "Lns_getFromMulti( " )
      filter( node:get_exp2(), self, node )
      self:write( ", 0 )" )
   else
    
      filter( node:get_exp2(), self, node )
   end
   
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
      
      self:outputAny2Type( node:get_expType() )
      self:popIndent(  )
   end
   
end


local op2map = {[".."] = "+", ["~="] = "!="}

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
                  do
                     local op = op2map[opTxt]
                     if op ~= nil then
                        self:write( string.format( " %s ", op) )
                     else
                        self:write( string.format( " %s ", opTxt) )
                     end
                  end
                  
                  filter( node:get_exp2(), self, node )
               end
            end
            
      end
   end
   
end


function convFilter:processExpRef( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpRef'

   if node:get_expType():get_kind() == Ast.TypeInfoKind.DDD then
      self:write( "ddd" )
   else
    
      if node:get_symbolInfo():get_convModuleParam() then
         self:write( string.format( "%s_%d", node:get_symbolInfo():get_name(), node:get_symbolInfo():get_symbolId()) )
      else
       
         do
            local _switchExp = node:get_symbolInfo():get_kind()
            if _switchExp == Ast.SymbolKind.Var or _switchExp == Ast.SymbolKind.Arg then
               self:write( self:getSymbolSym( node:get_symbolInfo() ) )
            elseif _switchExp == Ast.SymbolKind.Fun then
               if Ast.isBuiltin( node:get_expType():get_typeId() ) then
                  local builtinFunc = TransUnit.getBuiltinFunc(  )
                  do
                     local _switchExp = node:get_expType()
                     if _switchExp == builtinFunc.lns_print then
                        self:write( "Lns_print" )
                     else 
                        
                           Util.err( string.format( "%s: not support -- %s", __func__, node:get_symbolInfo():get_name()) )
                     end
                  end
                  
               else
                
                  self:write( self:getSymbol( _lune.newAlge( SymbolKind.Func, {node:get_expType()}), node:get_symbolInfo():get_name() ) )
               end
               
            elseif _switchExp == Ast.SymbolKind.Typ then
               self:write( self:getTypeSymbol( node:get_expType() ) )
            else 
               
                  self:write( node:get_symbolInfo():get_name() )
            end
         end
         
      end
      
   end
   
end


function convFilter:processExpRefItem( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpRefItem'

   local prefixType = node:get_val():get_expType():get_nonnilableType()
   do
      local _switchExp = prefixType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         if node:get_nilAccess() then
            self:write( "Lns_NilAccFin( Lns_NilAccPush( " )
            filter( node:get_val(), self, node )
            self:write( ") && " )
            self:write( "Lns_NilAccLast( Lns_NilAccPop().(*LnsList)" )
            self:write( ".GetAt(" )
            filter( _lune.unwrap( node:get_index()), self, node )
            self:write( ")" )
            self:outputAny2Type( node:get_expType() )
            self:write( "))" )
         else
          
            filter( node:get_val(), self, node )
            self:write( ".GetAt(" )
            filter( _lune.unwrap( node:get_index()), self, node )
            self:write( ")" )
            self:outputAny2Type( node:get_expType() )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Map or _switchExp == Ast.TypeInfoKind.Stem then
         if node:get_nilAccess() then
            self:write( "Lns_NilAccFin( Lns_NilAccPush( " )
            filter( node:get_val(), self, node )
            self:write( ") && " )
            self:write( "Lns_NilAccLast( Lns_NilAccPop().(*LnsMap)" )
         else
          
            filter( node:get_val(), self, node )
            if prefixType:get_kind() == Ast.TypeInfoKind.Stem then
               self:write( ".(*LnsMap)" )
            end
            
         end
         
         self:write( ".Items[" )
         do
            local index = node:get_index()
            if index ~= nil then
               filter( index, self, node )
            else
               self:write( string.format( '"%s"', str2gostr( _lune.unwrap( node:get_symbol()) )) )
            end
         end
         
         self:write( "]" )
         if node:get_nilAccess() then
            self:write( "))" )
         end
         
      else 
         
            if prefixType == Ast.builtinTypeString then
               self:write( "LnsInt(" )
               filter( node:get_val(), self, node )
               self:write( "[" )
               filter( _lune.unwrap( node:get_index()), self, node )
               self:write( "-1])" )
            else
             
               Util.err( string.format( "%s: not support -- %d, %s", __func__, node:get_pos().lineNo, Ast.TypeInfoKind:_getTxt( prefixType:get_kind())
               ) )
            end
            
      end
   end
   
end


function convFilter:processRefField( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processRefField'

   if node:get_prefix():get_expType():get_kind() == Ast.TypeInfoKind.Enum and node:get_expType():get_kind() == Ast.TypeInfoKind.Enum then
      self:write( self:getSymbol( _lune.newAlge( SymbolKind.Static, {node:get_expType()}), node:get_field().txt ) )
      return 
   end
   
   do
      local symbol = node:get_symbolInfo()
      if symbol ~= nil then
         local orgSym = symbol:getOrg(  )
         local builtinFuncs = TransUnit.getBuiltinFunc(  )
         if _lune._Set_has(builtinFuncs:get_allSymbolSet(), orgSym ) then
            self:write( string.format( "Lns_%s_%s", node:get_prefix():get_expType():get_rawTxt():gsub( "@", "" ), orgSym:get_name()) )
            return 
         end
         
         
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
         do
            local _switchExp = Nodes.getUnwraped( node:get_prefix() ):get_kind()
            if _switchExp == Nodes.NodeKind.get_ExpCall() or _switchExp == Nodes.NodeKind.get_GetField() then
               self:writeln( "&&" )
            end
         end
         
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
         if orgSym:get_kind() == Ast.SymbolKind.Mbr and orgSym:get_typeInfo():get_kind() == Ast.TypeInfoKind.Alternate and isAnyType( orgSym:get_typeInfo() ) then
            self:outputAny2Type( symbol:get_typeInfo() )
         end
         
      else
         Util.err( string.format( "not support -- %s", __func__) )
      end
   end
   
   
   for _6477 = 1, openParenNum do
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
               self:write( ".(LnsAlgeVal).GetTxt()" )
               return 
            end
            
         end
         
         
         if symbolInfo:get_staticFlag() then
            self:write( self:getSymbolSym( symbolInfo ) )
            self:write( "()" )
         else
          
            self:outputCallPrefix( node:get_id(), node, node:get_prefix(), symbolInfo )
            self:write( "()" )
            local retType = symbolInfo:get_typeInfo():get_retTypeInfoList()[1]
            if retType:get_kind() == Ast.TypeInfoKind.Alternate and not retType:hasBase(  ) then
               self:outputAny2Type( node:get_expType() )
            end
            
            
            if node:hasNilAccess(  ) then
               self:write( "})" )
               do
                  local _switchExp = opt.node:get_kind()
                  if _switchExp == Nodes.NodeKind.get_GetField() or _switchExp == Nodes.NodeKind.get_RefField() or _switchExp == Nodes.NodeKind.get_ExpCall() then
                  else 
                     
                        self:write( ")" )
                  end
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
   
   self:writeln( "" )
end


function convFilter:processTestBlock( node, opt )

   if not self.enableTest then
      return 
   end
   
   self:writeln( string.format( "func lns_test_%s_%s( %s *Testing_Ctrl ) {", getModuleName( self.moduleTypeInfo ), node:get_name().txt, node:get_ctrlName()) )
   
   filter( node:get_block(), self, node )
   
   self:writeln( "}" )
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

   self:write( "NewLnsSet(" )
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


function convFilter:processLiteralMap( node, opt )

   local hasNilable = false
   self:write( "NewLnsMap( map[LnsAny]LnsAny{" )
   for __index, pair in ipairs( node:get_pairList() ) do
      if pair:get_key():get_kind() == Nodes.NodeKind.get_LiteralNil(  ) or pair:get_val():get_kind() == Nodes.NodeKind.get_LiteralNil(  ) then
         
      else
       
         if pair:get_key():get_expType():get_kind() == Ast.TypeInfoKind.Nilable or pair:get_val():get_expType():get_kind() == Ast.TypeInfoKind.Nilable then
            hasNilable = true
         end
         
         filter( pair:get_key(), self, node )
         self:write( ":" )
         filter( pair:get_val(), self, node )
         self:write( "," )
      end
      
   end
   
   self:write( "})" )
   if hasNilable then
      self:write( ".Correct()" )
   end
   
end


function convFilter:processLiteralArray( node, opt )

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
         self:write( string.format( 'Lns_getVM().String_format(%s, ', str2gostr( txt )) )
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
   self:writeln( "" )
end


function convFilter:processLiteralSymbol( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processLiteralSymbol'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processLuneControl( node, opt )

   do
      local _matchExp = node:get_pragma()
      if _matchExp[1] == LuneControl.Pragma.default__init[1] then
      
      elseif _matchExp[1] == LuneControl.Pragma.default__init_old[1] then
      
      elseif _matchExp[1] == LuneControl.Pragma.disable_mut_control[1] then
      
      elseif _matchExp[1] == LuneControl.Pragma.ignore_symbol_[1] then
      
      elseif _matchExp[1] == LuneControl.Pragma.load__lune_module[1] then
      
      elseif _matchExp[1] == LuneControl.Pragma.can_not_conv_code[1] then
         local _ = _matchExp[2][1]
      
      end
   end
   
end


function convFilter:processAbbr( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processAbbr'

   Util.err( string.format( "not support -- %s:%d", __func__, node:get_pos().lineNo) )
end


local function createFilter( enableTest, streamName, stream, ast )

   return convFilter.new(enableTest, streamName, stream, ast)
end
_moduleObj.createFilter = createFilter

return _moduleObj
