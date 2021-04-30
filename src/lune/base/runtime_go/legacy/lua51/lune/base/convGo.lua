--lune/base/convGo.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@convGo'
local _lune = {}
if _lune3 then
   _lune = _lune3
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

if not _lune3 then
   _lune3 = _lune
end
local Ver = _lune.loadModule( 'lune.base.Ver' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Util = _lune.loadModule( 'lune.base.Util' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )
local Types = _lune.loadModule( 'lune.base.Types' )
local LnsOpt = _lune.loadModule( 'lune.base.Option' )

local MaxNilAccNum = 3

local Opt = {}
_moduleObj.Opt = Opt
function Opt.new( parent )
   local obj = {}
   Opt.setmeta( obj )
   if obj.__init then obj:__init( parent ); end
   return obj
end
function Opt:__init(parent) 
   self.parent = parent
end
function Opt.setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
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


local Env = {}
function Env.new(  )
   local obj = {}
   Env.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Env:__init() 
   self.stack = {}
end
function Env:push( node )

   table.insert( self.stack, node )
end
function Env:pop(  )

   table.remove( self.stack )
end
function Env:getLuavm( threading )

   if #self.stack == 0 then
      if threading then
         return "self.LnsEnv.LuaVM"
      else
       
         return "Lns_getVM()"
      end
      
   else
    
      return "Lns_getVMSync()"
   end
   
end
function Env:getEnv( threading )

   if #self.stack == 0 then
      if threading then
         return "self.LnsEnv"
      else
       
         return "Lns_GetEnv()"
      end
      
   else
    
      return "Lns_GetEnvSync()"
   end
   
end
function Env.setmeta( obj )
  setmetatable( obj, { __index = Env  } )
end


local function isMain( funcType )

   if funcType:get_kind() == Ast.TypeInfoKind.Func and funcType:get_rawTxt() == "__main" and funcType:get_accessMode() == Ast.AccessMode.Pub then
      return true
   end
   
   return false
end

local Option = {}
_moduleObj.Option = Option
function Option.setmeta( obj )
  setmetatable( obj, { __index = Option  } )
end
function Option.new( packageName, appName, mainModule )
   local obj = {}
   Option.setmeta( obj )
   if obj.__init then
      obj:__init( packageName, appName, mainModule )
   end
   return obj
end
function Option:__init( packageName, appName, mainModule )

   self.packageName = packageName
   self.appName = appName
   self.mainModule = mainModule
end
function Option:get_packageName()
   return self.packageName
end
function Option:get_appName()
   return self.appName
end
function Option:get_mainModule()
   return self.mainModule
end


local convFilter = {}
setmetatable( convFilter, { __index = Nodes.Filter } )
function convFilter.new( enableTest, streamName, stream, ast, option )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( enableTest, streamName, stream, ast, option ); end
   return obj
end
function convFilter:__init(enableTest, streamName, stream, ast, option) 
   Nodes.Filter.__init( self,true, ast:get_moduleTypeInfo(), ast:get_moduleTypeInfo():get_scope())
   
   
   self.moduleType2SymbolMap = {}
   self.env = Env.new()
   self.option = option
   self.processInfo = ast:get_processInfo()
   self.stream = Util.SimpleSourceOStream.new(stream, nil, 4)
   self.processMode = ProcessMode.Main
   self.processModeStack = {}
   self.moduleTypeInfo = ast:get_moduleTypeInfo()
   self.moduleScope = _lune.unwrap( ast:get_moduleTypeInfo():get_scope())
   self.builtin2runtime = {}
   self.type2gotypeMap = {}
   self.nodeManager = Nodes.NodeManager.new()
   self.enableTest = enableTest
   self.module2PackSym = {}
   
   local modDir = self.moduleTypeInfo:getParentFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager(), false )
   self.modDir = modDir:gsub( "@", "" ):gsub( "%.$", "" )
   
   self.noneNode = Nodes.NoneNode.create( self.nodeManager, Parser.noneToken.pos, false, {Ast.builtinTypeNone} )
end
function convFilter:getVM( threading, typeInfo )

   local txt = self.builtin2runtime[typeInfo]
   if  nil == txt then
      local _txt = txt
   
      return nil
   end
   
   return (txt:gsub( "GETVM", self.env:getLuavm( threading ) ) )
end
function convFilter:pushProcessMode( mode )

   table.insert( self.processModeStack, self.processMode )
   self.processMode = mode
end
function convFilter:popProcessMode(  )

   self.processMode = self.processModeStack[#self.processModeStack]
   table.remove( self.processModeStack )
end
function convFilter:isPubType( typeInfo )

   if Ast.isPubToExternal( typeInfo:get_accessMode() ) then
      if typeInfo:get_kind() == Ast.TypeInfoKind.Func then
         do
            local _switchExp = typeInfo:get_parentInfo():get_kind()
            if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Enum then
               return self:isPubType( typeInfo:get_parentInfo() )
            end
         end
         
      end
      
      return true
   end
   
   return typeInfo:getModule(  ) ~= self.moduleTypeInfo
end
function convFilter:isPubSym( symbol )

   if Ast.isPubToExternal( symbol:get_accessMode() ) then
      return true
   end
   
   return symbol:getModule(  ) ~= self.moduleTypeInfo
end
function convFilter:isExtType( typeInfo )

   return typeInfo:getModule(  ) ~= self.moduleTypeInfo
end
function convFilter:isExtSymbol( symbol )

   return symbol:getModule(  ) ~= self.moduleTypeInfo
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



local ignoreNodeInInnerBlockSet = {[Nodes.NodeKind.get_DeclAlge()] = true, [Nodes.NodeKind.get_DeclEnum()] = true, [Nodes.NodeKind.get_DeclMethod()] = true, [Nodes.NodeKind.get_DeclForm()] = true, [Nodes.NodeKind.get_DeclMacro()] = true, [Nodes.NodeKind.get_TestCase()] = true}

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
      elseif _switchExp == Ast.TypeInfoKind.Ext then
         if typeInfo:get_extedType():get_kind() == Ast.TypeInfoKind.Stem then
            return true
         end
         
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

local golanKeywordSet = {["func"] = true, ["select"] = true, ["defer"] = true, ["go"] = true, ["chan"] = true, ["package"] = true, ["const"] = true, ["fallthrough"] = true, ["range"] = true, ["continue"] = true, ["var"] = true, ["map"] = true, ["default"] = true}

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
SymbolKind.Func = { "Func", {{}}}
SymbolKind._name2Val["Func"] = SymbolKind.Func
SymbolKind.Member = { "Member", {{}}}
SymbolKind._name2Val["Member"] = SymbolKind.Member
SymbolKind.Normal = { "Normal"}
SymbolKind._name2Val["Normal"] = SymbolKind.Normal
SymbolKind.Static = { "Static", {{}}}
SymbolKind._name2Val["Static"] = SymbolKind.Static
SymbolKind.Type = { "Type", {{},{}}}
SymbolKind._name2Val["Type"] = SymbolKind.Type
SymbolKind.Var = { "Var", {{}}}
SymbolKind._name2Val["Var"] = SymbolKind.Var


local function concatGLSym( name, external )

   if external then
      local frontChar = string.byte( name, 1 )
      local front
      
      if frontChar >= 97 and frontChar <= 122 then
         front = string.format( "%c", 65 + frontChar - 97)
      elseif frontChar >= 65 and frontChar <= 90 then
         front = name:sub( 1, 1 )
      else
       
         front = string.format( "G%c", frontChar)
      end
      
      return front .. name:sub( 2 )
   end
   
   return name
   
end

local function outputGoMain( appName, mod, testing, path, opt )

   
   local lune_path = "main.go"
   if path ~= nil then
      if path ~= "" then
         lune_path = path
      end
      
   end
   
   local fileObj = io.open( lune_path, "w" )
   if  nil == fileObj then
      local _fileObj = fileObj
   
      return string.format( "failed to open -- %s", lune_path)
   end
   
   
   local base_mainCode = [==[
package main

import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
//IMPORT_MAIN:

//IMPORT:
////TEST:import . "lns/lune/base"

func main() {
    Lns_InitModOnce($opt)
    //TEST:Lns_Testing_init()
    Lns_init()
    //TEST:Testing_run( "" )
    //TEST:Testing_outputAllResult(Lns_io_stdout)
}
]==]
   
   local mainMod = mod:gsub( ".*%.", "" )
   local code = base_mainCode
   if mod ~= mainMod then
      local importPath = mod:gsub( "%.[^%.]+$", "" ):gsub( "%.", "/" )
      code = code:gsub( "//IMPORT_MAIN:", string.format( 'import . "%s/%s"', appName, importPath) )
   end
   
   
   if testing then
      code = code:gsub( "Lns_init", string.format( "Lns_%s_init", mainMod) )
      code = code:gsub( "//TEST:", "" )
      code = code:gsub( 'run%( "" %)', string.format( 'run( "lune.base.%s" )', mainMod) )
      code = code:gsub( '//IMPORT:', 'import . "github.com/ifritJP/LuneScript/src/lune/base"' )
   else
    
      code = code:gsub( "Lns_init%(%)", string.format( "Lns_RunMain( %s___main )", concatGLSym( mainMod, true )) )
   end
   
   do
      local _switchExp = opt:get_int2strMode()
      if _switchExp == LnsOpt.Int2strMode.Int2strModeNeed0 then
         code = code:gsub( "$opt", "LnsRuntimeOpt{ Int2strModeNeed0 }" )
      elseif _switchExp == LnsOpt.Int2strMode.Int2strModeUnneed0 then
         code = code:gsub( "$opt", "LnsRuntimeOpt{ Int2strModeUnneed0 }" )
      elseif _switchExp == LnsOpt.Int2strMode.Int2strModeDepend then
         code = code:gsub( "$opt", "LnsRuntimeOpt{ Int2strModeDepend }" )
      end
   end
   
   
   fileObj:write( code )
   
   return nil
end
_moduleObj.outputGoMain = outputGoMain

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


function convFilter:getModuleName( typeInfo, assign )

   if typeInfo == Ast.headTypeInfo then
      return ""
   end
   
   local moduleType = typeInfo:getModule(  )
   if assign then
      do
         local symbolInfo = self.moduleType2SymbolMap[moduleType]
         if symbolInfo ~= nil then
            return symbolInfo:get_name()
         end
      end
      
   end
   
   return (moduleType:get_rawTxt():gsub( "@", "" ) )
end


function convFilter:concatSymWithType( name, typeInfo )

   local modName = self:getModuleName( typeInfo:getModule(  ), false )
   local typeName
   
   if modName == "" then
      typeName = name
   else
    
      typeName = string.format( "%s_%s", modName, name)
   end
   
   return concatGLSym( typeName, self:isPubType( typeInfo ) )
end


function convFilter:isSamePackageExtModule( typeInfo )

   local extModuleType = _lune.__Cast( typeInfo:get_srcTypeInfo():get_nonnilableType():get_aliasSrc(), 3, Ast.NormalTypeInfo )
   if  nil == extModuleType then
      local _extModuleType = extModuleType
   
      Util.err( string.format( "illegal type -- %s", typeInfo:getTxt(  )) )
   end
   
   local requireParent = extModuleType:get_requirePath():gsub( "[^%.]+$", "" )
   local moduleParent = self:getFull( self.moduleTypeInfo, false ):gsub( "[^@%.]+$", "" ):gsub( "@", "" )
   return requireParent == moduleParent
end


function convFilter:isSameModDir( moduleTypeInfo )

   if moduleTypeInfo:get_parentInfo() == self.moduleTypeInfo:get_parentInfo() then
      return true
   end
   
   return false
end


function convFilter:isSamePackage( typeInfo )

   if typeInfo:get_kind() == Ast.TypeInfoKind.ExtModule then
      return self:isSamePackageExtModule( typeInfo )
   end
   
   return self:isSameModDir( typeInfo )
end


function convFilter:needPrefix( typeInfo )

   if Ast.isBuiltin( typeInfo:get_typeId() ) then
      return false
   end
   
   return not self:isSamePackage( typeInfo )
end


function convFilter:getSymbol( kind, name )
   local __func__ = '@lune.@base.@convGo.convFilter.getSymbol'

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
         if symbolInfo:getModule(  ) ~= self.moduleTypeInfo and symbolInfo:get_staticFlag() then
            symbolName = string.format( "%s_%s", self:getModuleName( symbolInfo:getModule(  ), false ), symbolInfo:get_name())
         elseif name == "__mod__" then
            symbolName = string.format( "%s__mod__", modName)
         elseif symbolInfo:get_scope() == self.moduleScope then
            symbolName = concatGLSym( string.format( "%s_", modName) .. symbolName, Ast.isPubToExternal( symbolInfo:get_accessMode() ) )
         elseif not symbolInfo:get_posForModToRef() then
            if symbolName ~= "__func__" then
               symbolName = "_"
            end
            
         end
         
         if self:needPrefix( symbolInfo:getModule(  ) ) then
            symbolName = string.format( "%s.%s", self:getModuleName( symbolInfo:getModule(  ), true ), symbolName)
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
          
            local prefix = nil
            do
               local _switchExp = typeInfo:get_parentInfo():get_kind()
               if _switchExp == Ast.TypeInfoKind.Module or _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Method then
                  if isInnerDeclType( typeInfo ) then
                     if not isClosure( typeInfo ) then
                        local parentName = typeInfo:getParentFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager(), true )
                        symbolName = string.format( "%s_%s_%d_", parentName:gsub( "[%.@]", "_" ), symbolName, typeInfo:get_typeId())
                     end
                     
                  else
                   
                     if not Ast.isBuiltin( typeInfo:get_typeId() ) and not self:isPubType( typeInfo ) then
                        symbolName = string.format( "%s_%d_", symbolName, typeInfo:get_typeId())
                     end
                     
                     symbolName = self:concatSymWithType( symbolName, typeInfo )
                  end
                  
               elseif _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.Class then
                  local parentInfo = typeInfo:get_parentInfo()
                  symbolName = string.format( "%s_%s", self:getSymbol( _lune.newAlge( SymbolKind.Type, {parentInfo,false}), parentInfo:get_rawTxt() ), symbolName)
                  if not self:isPubType( typeInfo ) then
                     symbolName = string.format( "%s_%d_", symbolName, typeInfo:get_typeId())
                  end
                  
               elseif _switchExp == Ast.TypeInfoKind.ExtModule then
                  symbolName = concatGLSym( symbolName, true )
                  if not self:isSamePackageExtModule( typeInfo:get_parentInfo() ) then
                     prefix = typeInfo:get_parentInfo():get_rawTxt()
                  end
                  
               else 
                  
                     Util.err( string.format( "%s: not support -- %s:%s", __func__, typeInfo:getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager(), true ), Ast.TypeInfoKind:_getTxt( typeInfo:get_parentInfo():get_kind())
                     ) )
               end
            end
            
            if not prefix then
               if self:needPrefix( typeInfo:getModule(  ) ) then
                  prefix = self:getModuleName( typeInfo, true )
               end
               
            end
            
            if prefix ~= nil then
               symbolName = string.format( "%s.%s", prefix, symbolName)
            end
            
         end
         
      elseif _matchExp[1] == SymbolKind.Type[1] then
         local typeInfo = _matchExp[2][1]
         local needPrefix = _matchExp[2][2]
      
         if typeInfo:get_kind() == Ast.TypeInfoKind.FormFunc then
            return self:getSymbol( _lune.newAlge( SymbolKind.Func, {typeInfo}), symbolName )
         end
         
         local workName
         
         if isInnerDeclType( typeInfo ) and not Ast.isBuiltin( typeInfo:get_typeId() ) then
            workName = string.format( "%s%d", name, typeInfo:get_typeId())
         else
          
            workName = symbolName
         end
         
         symbolName = self:concatSymWithType( workName, typeInfo )
         if needPrefix and self:needPrefix( typeInfo:getModule(  ) ) then
            symbolName = string.format( "%s.%s", self:getModuleName( typeInfo, true ), symbolName)
         end
         
      elseif _matchExp[1] == SymbolKind.Static[1] then
         local typeInfo = _matchExp[2][1]
      
         local workName = self:getSymbol( _lune.newAlge( SymbolKind.Type, {typeInfo,true}), typeInfo:get_rawTxt() )
         symbolName = string.format( "%s__%s", workName, name)
      elseif _matchExp[1] == SymbolKind.Normal[1] then
      
      end
   end
   
   
   return symbolName
end


function convFilter:getTypeSymbol( typeInfo )

   local orgType = typeInfo:get_srcTypeInfo():get_nonnilableType():get_aliasSrc()
   return self:getSymbol( _lune.newAlge( SymbolKind.Type, {orgType,false}), orgType:get_rawTxt() )
end

function convFilter:getTypeSymbolWithPrefix( typeInfo )

   local orgType = typeInfo:get_srcTypeInfo():get_nonnilableType():get_aliasSrc()
   return self:getSymbol( _lune.newAlge( SymbolKind.Type, {orgType,true}), orgType:get_rawTxt() )
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
      elseif _switchExp == Ast.SymbolKind.Typ then
         if _lune.__Cast( symbolInfo:get_typeInfo(), 3, Ast.AliasTypeInfo ) then
            return self:getSymbol( _lune.newAlge( SymbolKind.Var, {symbolInfo}), symbolInfo:get_name() )
         end
         
         return self:getTypeSymbol( symbolInfo:get_typeInfo() )
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


function convFilter:getConv2formName( node )

   return string.format( "conv2Form%d", node:get_id())
end


function convFilter:getConvGenericsName( node )

   return string.format( "lns_convGenerics%d", node:get_id())
end


function convFilter:getModuleSym( moduleTypeInfo, addDot )

   do
      local packSym = self.module2PackSym[moduleTypeInfo]
      if packSym ~= nil then
         if addDot then
            return string.format( "%s.", packSym)
         end
         
         return packSym
      end
   end
   
   return ""
end

local function str2gostr( txt )

   local work = txt
   if string.find( work, '^```' ) then
      work = work:sub( 4, -4 ):gsub( "^\n", "" )
      work = Parser.quoteStr( work )
   elseif string.find( work, "^'" ) then
      work = Parser.convFromRawToStr( work ):gsub( '"', '\\"' )
      work = string.format( '"%s"', work)
   end
   
   
   return work
end

local function getOrgTypeInfo( typeInfo )

   do
      local enumType = _lune.__Cast( typeInfo:get_srcTypeInfo():get_nonnilableType():get_aliasSrc(), 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         return enumType:get_valTypeInfo()
      end
   end
   
   return typeInfo:get_srcTypeInfo():get_nonnilableType()
end

function convFilter:type2gotypeOrg( typeInfo, addClassAster )

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
         if orgType:get_extedType():get_kind() == Ast.TypeInfoKind.Stem then
            return "LnsAny"
         end
         
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
         if typeInfo:get_genSrcTypeInfo() == TransUnit.getBuiltinFunc(  ).__pipe_ then
            return "*Lns__pipe"
         end
         
         local symbol = self:getTypeSymbolWithPrefix( typeInfo )
         if addClassAster then
            return "*" .. symbol
         end
         
         return symbol
      elseif _switchExp == Ast.TypeInfoKind.IF then
         return self:getTypeSymbolWithPrefix( typeInfo )
      elseif _switchExp == Ast.TypeInfoKind.Alternate then
         return self:type2gotypeOrg( typeInfo:get_baseTypeInfo(), addClassAster )
      end
   end
   
   Util.err( string.format( "not support yet -- %s", typeInfo:getTxt(  )) )
end


function convFilter:type2gotype( typeInfo )

   return self:type2gotypeOrg( typeInfo, true )
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

function convFilter:outputStem2Type( dstType )

   if dstType:get_kind() == Ast.TypeInfoKind.Alternate and dstType:hasBase(  ) then
      self:write( string.format( ".(%s)", self:type2gotype( dstType )) )
   elseif dstType:get_kind() == Ast.TypeInfoKind.Class and dstType ~= Ast.builtinTypeString then
      self:write( string.format( ".(%sDownCast).To%s()", self:getTypeSymbolWithPrefix( dstType ), self:getTypeSymbol( dstType )) )
   else
    
      self:outputAny2Type( dstType )
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
   
   
   if not self:isSameModDir( node:get_moduleTypeInfo() ) and not Ast.isBuiltin( node:get_moduleTypeInfo():get_typeId() ) then
      self:write( string.format( "%s.", self:getModuleName( node:get_moduleTypeInfo(), true )) )
   end
   
   self:writeln( string.format( "Lns_%s_init()", self:getModuleName( node:get_moduleTypeInfo(), false )) )
end


function convFilter:needConvFormFunc( node )

   local castType = node:get_castType():get_extedType():get_nonnilableType()
   if castType:get_kind() ~= Ast.TypeInfoKind.FormFunc then
      return false
   end
   
   
   local funcType = node:get_exp():get_expType():get_extedType()
   if #castType:get_argTypeInfoList() ~= #funcType:get_argTypeInfoList() or #castType:get_retTypeInfoList() ~= #funcType:get_retTypeInfoList() then
      return true
   end
   
   for index, argType in ipairs( castType:get_argTypeInfoList() ) do
      if not argType:equals( self.processInfo, funcType:get_argTypeInfoList()[index] ) then
         return true
      end
      
   end
   
   for index, retType in ipairs( castType:get_retTypeInfoList() ) do
      if not retType:equals( self.processInfo, funcType:get_retTypeInfoList()[index] ) then
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
               self:write( string.format( "%sDownCastF(", self:getTypeSymbolWithPrefix( castType )) )
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
         if Ast.isClass( node:get_expType() ) then
            self:write( ".FP" )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.FormFunc then
         local expType = node:get_expType()
         if self:needConvFormFunc( parent ) then
            self:write( string.format( "%s(", self:getConv2formName( parent )) )
            filter( node, self, parent )
            self:write( ")" )
         else
          
            self:write( string.format( "%s(", self:getTypeSymbol( castType )) )
            filter( node, self, parent )
            self:write( ")" )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Alternate then
         if not castType:hasBase(  ) then
            if Ast.isClass( node:get_expType():get_nonnilableType() ) then
               self:write( string.format( "%s2Stem(", self:getTypeSymbolWithPrefix( node:get_expType():get_nonnilableType() )) )
               filter( node, self, parent )
               self:write( ")" )
            else
             
               filter( node, self, parent )
            end
            
         else
          
            self:outputImplicitCast( castType:get_baseTypeInfo(), node, parent )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Form then
         self:write( string.format( "%s(", self:getConv2formName( parent )) )
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
         
            if node:get_expType():get_nilable() and Ast.isClass( node:get_expType():get_nonnilableType() ) then
               self:write( string.format( "%s2Stem(", self:getTypeSymbolWithPrefix( node:get_expType():get_nonnilableType() )) )
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

   if #dstTypeList < #node:get_expList() then
      if dstTypeList[#dstTypeList]:get_kind() ~= Ast.TypeInfoKind.DDD then
         return ExpListKind.Conv
      end
      
   end
   
   
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

function convFilter:getConvExpName( nodeId, argListNode )

   return string.format( "%s_convExp%d", self.moduleTypeInfo:get_rawTxt():gsub( "@", "" ), nodeId)
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
   
   self:writeln( string.format( "// for %d", argList:get_pos().lineNo) )
   self:write( string.format( "func %s(", self:getConvExpName( nodeId, argList )) )
   
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
   
   self:write( ") " )
   
   local function getRetType( retType, index )
   
      if retType == Ast.builtinTypeEmpty then
         return argList:getExpTypeNoDDDAt( index )
      end
      
      return retType
   end
   
   local retTypeList = {}
   for index, dstType in ipairs( dstTypeList ) do
      table.insert( retTypeList, getRetType( dstType, index ) )
   end
   
   
   if #retTypeList >= 2 then
      self:write( "(" )
      for index, retType in ipairs( retTypeList ) do
         if index ~= 1 then
            self:write( ", " )
         end
         
         self:write( self:type2gotype( getRetType( retType, index ) ) )
      end
      
      self:writeln( ") {" )
   elseif #retTypeList == 1 then
      self:write( self:type2gotype( getRetType( retTypeList[1], 1 ) ) )
      self:writeln( " {" )
   else
    
      self:writeln( " {" )
   end
   
   self:write( "    return " )
   
   if mRetIndex ~= nil then
      local restIndex = nil
      for index, retType in ipairs( retTypeList ) do
         if index ~= 1 then
            self:write( ", " )
         end
         
         if retType:get_kind() == Ast.TypeInfoKind.DDD then
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
               self:outputAny2Type( retType )
            end
            
         else
          
            self:write( string.format( "arg%d", index) )
         end
         
      end
      
      if restIndex ~= nil then
         self:write( "Lns_2DDD( " )
         for index, _6385 in ipairs( expList ) do
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
      for index, _6393 in ipairs( retTypeList ) do
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
func lns_NilAccCall_%s( env *LnsEnv, call func () (%s) ) bool {
    return env.NilAccPush( Lns_2DDD( call() ) )
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
   
   self:writeln( string.format( "func %s(%s) (%s) {", self:getConvGenericsName( node ), srcTxt, dstTxt) )
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

FuncInfo.DeclInfo = { "DeclInfo", {{},{}}}
FuncInfo._name2Val["DeclInfo"] = FuncInfo.DeclInfo
FuncInfo.Type = { "Type", {{}}}
FuncInfo._name2Val["Type"] = FuncInfo.Type
FuncInfo.WithClass = { "WithClass", {{},{}}}
FuncInfo._name2Val["WithClass"] = FuncInfo.WithClass


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


local FuncConv = {}
function FuncConv.new( retList )
   local obj = {}
   FuncConv.setmeta( obj )
   if obj.__init then obj:__init( retList ); end
   return obj
end
function FuncConv:__init(retList) 
   self.argList = {}
   self.retList = retList
end
function FuncConv.setmeta( obj )
  setmetatable( obj, { __index = FuncConv  } )
end
function FuncConv:get_argList()
   return self.argList
end
function FuncConv:get_retList()
   return self.retList
end


function convFilter:outputDeclFunc( funcInfo )

   local typeInfo
   
   local name
   
   local prefixType
   
   local extFlag
   
   do
      local _matchExp = funcInfo
      if _matchExp[1] == FuncInfo.DeclInfo[1] then
         local node = _matchExp[2][1]
         local workDeclInfo = _matchExp[2][2]
      
         extFlag = false
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
      
         extFlag = workTypeInfo:get_kind() == Ast.TypeInfoKind.Ext
         typeInfo = workTypeInfo
         prefixType = typeInfo:get_parentInfo()
         name = typeInfo:get_rawTxt()
      elseif _matchExp[1] == FuncInfo.WithClass[1] then
         local classType = _matchExp[2][1]
         local methodType = _matchExp[2][2]
      
         extFlag = false
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
      
      
      if typeInfo:get_extedType():get_kind() ~= Ast.TypeInfoKind.FormFunc then
         if name ~= nil then
            self:outputSymbol( _lune.newAlge( SymbolKind.Func, {typeInfo}), name )
         end
         
      end
      
   end
   
   
   self:write( "(" )
   
   local workType = typeInfo:getOverridingType(  )
   if  nil == workType then
      local _workType = workType
   
      workType = typeInfo
   end
   
   
   local retTypeList
   
   if extFlag then
      retTypeList = _lune.unwrap( Ast.convToExtTypeList( self.processInfo, workType:get_retTypeInfoList() ))
   else
    
      retTypeList = workType:get_retTypeInfoList()
   end
   
   
   local funcConv = FuncConv.new(retTypeList)
   
   do
      local _matchExp = funcInfo
      if _matchExp[1] == FuncInfo.DeclInfo[1] then
         local node = _matchExp[2][1]
         local declInfo = _matchExp[2][2]
      
         for index, arg in ipairs( declInfo:get_argList() ) do
            if index ~= 1 then
               self:write( "," )
            end
            
            local argType = workType:get_argTypeInfoList()[index]
            if argType:get_nonnilableType():get_kind() == Ast.TypeInfoKind.Alternate then
               do
                  local argNode = _lune.__Cast( arg, 3, Nodes.DeclArgNode )
                  if argNode ~= nil then
                     local argName = self:getSymbolSym( argNode:get_symbolInfo() )
                     self:write( string.format( "_%s ", argName) )
                     self:write( self:type2gotype( argType ) )
                     table.insert( funcConv:get_argList(), argNode:get_symbolInfo() )
                  else
                     filter( arg, self, node )
                  end
               end
               
            else
             
               filter( arg, self, node )
            end
            
         end
         
      elseif _matchExp[1] == FuncInfo.Type[1] then
         local _ = _matchExp[2][1]
      
         for index, argType in ipairs( workType:get_argTypeInfoList() ) do
            if index ~= 1 then
               self:write( "," )
            end
            
            self:write( string.format( "arg%d %s", index, self:type2gotype( argType )) )
         end
         
      elseif _matchExp[1] == FuncInfo.WithClass[1] then
         local _ = _matchExp[2][1]
         local _ = _matchExp[2][2]
      
         for index, argType in ipairs( workType:get_argTypeInfoList() ) do
            if index ~= 1 then
               self:write( "," )
            end
            
            self:write( string.format( "arg%d %s", index, self:type2gotype( argType )) )
         end
         
      end
   end
   
   self:write( ")" )
   
   self:outputRetType( retTypeList )
   
   return funcConv
end


function convFilter:outputConvToForm( node )

   local castType = node:get_castType():get_nonnilableType():get_extedType()
   if castType:get_kind() ~= Ast.TypeInfoKind.Form then
      return 
   end
   
   
   local funcType = node:get_exp():get_expType():get_extedType()
   if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Ext and funcType:get_srcTypeInfo():get_kind() == Ast.TypeInfoKind.Form then
      self:writeln( string.format( [==[      
func %s( luaform LnsAny ) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        return %s.RunLoadedfunc( luaform.(*Lns_luaValue), argList )
    }
}]==], self:getConv2formName( node ), self.env:getLuavm( node:get_exp():isThreading(  ) )) )
      return 
   end
   
   
   self:writeln( string.format( "// for %d: %s", node:get_pos().lineNo, Nodes.getNodeKindName( node:get_kind() )) )
   self:write( string.format( "func %s( src func (", self:getConv2formName( node )) )
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
      if symbolInfo:get_scope() == self.moduleScope and node:get_mode() == Nodes.DeclVarMode.Let then
         self:writeln( string.format( "var %s %s", self:getSymbolSym( symbolInfo ), self:type2gotype( symbolInfo:get_typeInfo() )) )
      end
      
   end
   
end


function convFilter:outputConvExt( funcNode )

   do
      local fieldNode = _lune.__Cast( funcNode, 3, Nodes.RefFieldNode )
      if fieldNode ~= nil then
         if fieldNode:get_prefix():get_expType():get_nonnilableType():get_kind() ~= Ast.TypeInfoKind.Ext then
            return 
         end
         
      else
         return 
      end
   end
   
   self:write( string.format( "func Lns_callExt%d( args []LnsAny ) (", funcNode:get_id()) )
   for index, retType in ipairs( funcNode:get_expType():get_retTypeInfoList() ) do
      if index > 1 then
         self:write( "," )
      end
      
      self:write( self:type2gotype( retType ) )
   end
   
   self:writeln( ") {" )
   self:write( "    return " )
   for index, _6585 in ipairs( funcNode:get_expType():get_retTypeInfoList() ) do
      if index > 1 then
         self:write( "," )
      end
      
      self:write( string.format( "Lns_getFromMulti( args, %d )", index - 1) )
   end
   
   self:writeln( "" )
   self:writeln( "}" )
end


function convFilter:outputModule( moduleTypeInfo, addDot )

   self:write( self:getModuleSym( moduleTypeInfo, addDot ) )
end


local function getModulePrefix( node )

   if node:get_assigned() then
      return string.format( "%s.", node:get_assignName())
   end
   
   local mod = node:get_moduleTypeInfo():get_parentInfo():get_rawTxt():gsub( "@", "" )
   if mod == "" then
      return "main."
   end
   
   return string.format( "%s.", mod)
end

function convFilter:outputModuleImport( node )

   if node:get_expType():get_kind() ~= Ast.TypeInfoKind.ExtModule or (node:get_lang() ~= Types.Lang.Go and node:get_lang() ~= Types.Lang.Same ) or self:isSamePackageExtModule( node:get_expType() ) then
      return 
   end
   
   
   local normalType = _lune.unwrap( _lune.__Cast( node:get_expType(), 3, Ast.NormalTypeInfo ))
   self:write( string.format( "import %s ", node:get_expType():get_rawTxt()) )
   
   local mod = normalType:get_requirePath():gsub( "%.[^%.]+$", "" )
   if mod:find( "^go/" ) then
      local workMod = mod:gsub( "^go/", "" ):gsub( "%.", "/" ):gsub( ":", "." )
      self:writeln( string.format( '"%s"', workMod) )
   else
    
      local path = normalType:get_requirePath():gsub( "%.", "/" ):gsub( ":", "." )
      self:writeln( string.format( '"%s/%s"', self.option:get_appName(), path) )
   end
   
end


function convFilter:outputImport( node )

   if self:isSameModDir( node:get_moduleTypeInfo() ) or Ast.isBuiltin( node:get_moduleTypeInfo():get_typeId() ) then
      return 
   end
   
   
   self:write( "import " )
   local packSym = node:get_assignName()
   if node:get_modulePath():find( "^go/" ) then
      local workMod = node:get_modulePath():gsub( "^go/", "" ):gsub( "%.", "/" ):gsub( ":", "." )
      self:writeln( string.format( '%s "%s"', packSym, (workMod:gsub( "/[^/]+$", "" ) )) )
   else
    
      local modulePath, count = node:get_modulePath():gsub( "([^%.]+)%.[^%.]+$", "/%1" )
      if count == 0 then
         self:writeln( string.format( '%s "%s"', packSym, self.option.appName) )
      else
       
         local modDir = modulePath:gsub( "/", "" )
         self:writeln( string.format( '%s "%s/%s"', packSym, self.option.appName, (modDir:gsub( "%.", "/" ) )) )
      end
      
   end
   
   
   self.module2PackSym[node:get_moduleTypeInfo()] = packSym
end


function convFilter:processRoot( node, opt )

   for __index, importNode in ipairs( node:get_nodeManager():getImportNodeList(  ) ) do
      self.moduleType2SymbolMap[importNode:get_moduleTypeInfo()] = importNode:get_symbolInfo()
   end
   
   
   for pragma, __val in pairs( node:get_luneHelperInfo().pragmaSet ) do
      do
         local _matchExp = pragma
         if _matchExp[1] == LuneControl.Pragma.limit_conv_code[1] then
            local codeSet = _matchExp[2][1]
         
            if not _lune._Set_has(codeSet, LuneControl.Code.Go ) then
               self:writeln( "// This code is transcompiled by LuneScript." )
               self:writeln( string.format( "package %s", self.option.packageName) )
               return 
            end
            
         end
      end
      
   end
   
   
   local function isUsingInTest( aNode )
   
      for __index, testBlock in ipairs( node:get_nodeManager():getTestBlockNodeList(  ) ) do
         if testBlock:isInnerPos( aNode:get_pos() ) then
            return true
         end
         
      end
      
      return false
   end
   
   
   
   
   
   local builtinFuncs = TransUnit.getBuiltinFunc(  )
   
   
   local builtin2runtime = {[builtinFuncs.str_gsub] = 'GETVM.String_gsub', [builtinFuncs.string_gsub] = 'GETVM.String_gsub', [builtinFuncs.str_find] = 'GETVM.String_find', [builtinFuncs.string_find] = 'GETVM.String_find', [builtinFuncs.str_byte] = 'GETVM.String_byte', [builtinFuncs.string_byte] = 'GETVM.String_byte', [builtinFuncs.str_format] = 'GETVM.String_format', [builtinFuncs.string_format] = 'GETVM.String_format', [builtinFuncs.str_rep] = 'GETVM.String_rep', [builtinFuncs.string_rep] = 'GETVM.String_rep', [builtinFuncs.str_gmatch] = 'GETVM.String_gmatch', [builtinFuncs.string_gmatch] = 'GETVM.String_gmatch', [builtinFuncs.str_sub] = 'GETVM.String_sub', [builtinFuncs.string_sub] = 'GETVM.String_sub', [builtinFuncs.str_lower] = 'GETVM.String_lower', [builtinFuncs.string_lower] = 'GETVM.String_lower', [builtinFuncs.str_upper] = 'GETVM.String_upper', [builtinFuncs.string_upper] = 'GETVM.String_upper', [builtinFuncs.str_reverse] = 'GETVM.String_reverse', [builtinFuncs.string_reverse] = 'GETVM.String_reverse', [Ast.builtinTypeNone] = ""}
   
   
   builtin2runtime[builtinFuncs.lns_error] = "panic"
   builtin2runtime[builtinFuncs.lns_print] = "Lns_print"
   builtin2runtime[builtinFuncs.lns_type] = "Lns_type"
   builtin2runtime[builtinFuncs.lns_require] = "Lns_require"
   builtin2runtime[builtinFuncs.lns_tonumber] = "Lns_tonumber"
   builtin2runtime[builtinFuncs.lns__load] = "GETVM.Load"
   builtin2runtime[builtinFuncs.lns_loadfile] = "GETVM.Loadfile"
   builtin2runtime[builtinFuncs.lns_expandLuavalMap] = "GETVM.ExpandLuavalMap"
   
   builtin2runtime[builtinFuncs.string_dump] = "GETVM.String_dump"
   
   builtin2runtime[builtinFuncs.io_open] = "Lns_io_open"
   builtin2runtime[builtinFuncs.io_popen] = "GETVM.IO_popen"
   builtin2runtime[builtinFuncs.package_searchpath] = "GETVM.Package_searchpath"
   builtin2runtime[builtinFuncs.os_clock] = "GETVM.OS_clock"
   builtin2runtime[builtinFuncs.os_exit] = "GETVM.OS_exit"
   builtin2runtime[builtinFuncs.os_remove] = "GETVM.OS_remove"
   builtin2runtime[builtinFuncs.os_date] = "GETVM.OS_date"
   builtin2runtime[builtinFuncs.os_time] = "GETVM.OS_time"
   builtin2runtime[builtinFuncs.os_difftime] = "GETVM.OS_difftime"
   builtin2runtime[builtinFuncs.os_rename] = "GETVM.OS_rename"
   builtin2runtime[builtinFuncs.math_random] = "GETVM.Math_random"
   builtin2runtime[builtinFuncs.math_randomseed] = "GETVM.Math_randomseed"
   
   self.builtin2runtime = builtin2runtime
   
   self.type2gotypeMap = {[Ast.builtinTypeInt] = "LnsInt", [Ast.builtinTypeReal] = "LnsReal", [Ast.builtinTypeStem] = "LnsAny", [Ast.builtinTypeString] = "string", [Ast.builtinTypeBool] = "bool", [builtinFuncs.ostream_] = "Lns_oStream", [builtinFuncs.istream_] = "Lns_iStream", [builtinFuncs.luastream_] = "Lns_luaStream"}
   
   self:writeln( "// This code is transcompiled by LuneScript." )
   self:writeln( string.format( "package %s", self.option.packageName) )
   self:writeln( 'import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"' )
   
   for __index, workNode in ipairs( node:get_nodeManager():getImportNodeList(  ) ) do
      self:outputImport( workNode )
   end
   
   for __index, workNode in ipairs( node:get_nodeManager():getDeclClassNodeList(  ) ) do
      self:outputModuleImport( workNode )
   end
   
   
   local initModVar = string.format( "init_%s", self:getModuleName( node:get_moduleTypeInfo(), false ))
   self:writeln( string.format( "var %s bool", initModVar) )
   
   self:pushProcessMode( ProcessMode.DeclTopScopeVar )
   local modSym = _lune.unwrap( self.moduleScope:getSymbolInfoChild( "__mod__" ))
   self:writeln( string.format( "var %s string", self:getSymbolSym( modSym )) )
   
   
   do
      local function procNode( workNode )
      
         filter( workNode, self, node )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getDeclEnumNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   for __index, workNode in ipairs( node:get_nodeManager():getDeclVarNodeList(  ) ) do
      self:outputTopScopeVar( workNode )
   end
   
   self:popProcessMode(  )
   
   
   do
      local function procNode( workNode )
      
         filter( workNode, self, node )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getDeclAlgeNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   do
      local function procNode( workNode )
      
         filter( workNode, self, node )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getDeclFormNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   do
      local function procNode( workNode )
      
         self:processGenericsCall( workNode )
         self:outputNilAccCall( workNode )
         self:outputConvExt( workNode:get_func() )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getExpCallNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   do
      local function procNode( workNode )
      
         self:outputConvToForm( workNode )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getExpCastNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   do
      local function procNode( workNode )
      
         filter( workNode, self, node )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getTestCaseNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   do
      local function procNode( workNode )
      
         local symTypeList = {}
         for _6763 = 1, #workNode:get_varSymList() do
            table.insert( symTypeList, Ast.builtinTypeStem_ )
         end
         
         self:processConvExp( workNode:get_id(), symTypeList, workNode:get_expList() )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getIfUnwrapNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   do
      local function procNode( workNode )
      
         self:processConvExp( workNode:get_id(), workNode:get_exp1():get_expTypeList(), workNode:get_exp2() )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getExpSetValNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   
   do
      local function procNode( workNode )
      
         self:processConvExp( workNode:get_id(), workNode:get_func():get_expType():get_argTypeInfoList(), workNode:get_argList() )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getExpCallNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   do
      local function procNode( workNode )
      
         local typeList = {}
         do
            local expList = workNode:get_expList()
            if expList ~= nil then
               for index, symbolInfo in ipairs( workNode:get_symbolInfoList() ) do
                  if workNode:get_mode() == Nodes.DeclVarMode.Let or workNode:get_mode() == Nodes.DeclVarMode.Unwrap then
                     if workNode:get_unwrapFlag() then
                        table.insert( typeList, expList:getExpTypeNoDDDAt( index ) )
                     else
                      
                        table.insert( typeList, symbolInfo:get_typeInfo() )
                     end
                     
                  end
                  
               end
               
               self:processConvExp( workNode:get_id(), typeList, expList )
            end
         end
         
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getDeclVarNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   do
      local function procNode( workNode )
      
         self:writeln( string.format( "type %s = %s", self:getSymbolSym( workNode:get_newSymbol() ), self:getTypeSymbol( workNode:get_typeInfo():get_aliasSrc() )) )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getAliasNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   
   self:pushProcessMode( ProcessMode.NonClosureFuncDecl )
   
   do
      local function procNode( workNode )
      
         filter( workNode, self, node )
         self:writeln( "" )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getDeclFuncNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   self:popProcessMode(  )
   
   self:pushProcessMode( ProcessMode.DeclClass )
   
   do
      local function procNode( workNode )
      
         filter( workNode, self, node )
         self:writeln( "" )
      end
      
      
      for __index, tmpNode in ipairs( node:get_nodeManager():getDeclClassNodeList(  ) ) do
         if self.enableTest or not isUsingInTest( tmpNode ) then
            procNode( tmpNode )
         end
         
      end
      
   end
   
   
   self:popProcessMode(  )
   
   self:writeln( string.format( "func Lns_%s_init() {", self:getModuleName( node:get_moduleTypeInfo(), false )) )
   self:pushIndent(  )
   
   self:writeln( string.format( "if %s { return }", initModVar) )
   self:writeln( string.format( "%s = true", initModVar) )
   
   self:writeln( string.format( '%s = "%s"', self:getSymbolSym( modSym ), node:get_moduleTypeInfo():getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager() )) )
   
   self:writeln( "Lns_InitMod()" )
   
   local modulePath = node:get_moduleTypeInfo():getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager() ):gsub( "@", "" )
   local moduleName = self:getModuleName( node:get_moduleTypeInfo(), false )
   
   if self.enableTest then
      for __index, workNode in ipairs( node:get_nodeManager():getTestCaseNodeList(  ) ) do
         self:writeln( string.format( 'Testing_registerTestcase( "%s", "%s", lns_test_%s_%s )', modulePath, workNode:get_name().txt, moduleName, workNode:get_name().txt) )
      end
      
   end
   
   
   for __index, child in ipairs( node:get_children() ) do
      if not _lune._Set_has(ignoreNodeInInnerBlockSet, child:get_kind() ) then
         filter( child, self, node )
      end
      
   end
   
   
   self:popIndent(  )
   self:writeln( "}" )
   
   if self.option.mainModule == self:getCanonicalName( self.moduleTypeInfo, false ):gsub( "@", "" ) then
      local hasMain = false
      for __index, workNode in ipairs( node:get_nodeManager():getDeclFuncNodeList(  ) ) do
         if isMain( workNode:get_expType() ) then
            hasMain = true
            break
         end
         
      end
      
      if not hasMain then
         local moduleSym = self:getModuleName( self.moduleTypeInfo, false )
         self:writeln( string.format( "func %s___main(args *LnsList) LnsInt {", concatGLSym( moduleSym, true )) )
         self:writeln( string.format( "Lns_%s_init()", moduleSym) )
         self:writeln( "return 0" )
         self:writeln( "}" )
      end
      
   end
   
   
   self:writeln( "func init() {" )
   self:pushIndent(  )
   self:writeln( string.format( "%s = false", initModVar) )
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processSubfile( node, opt )

   
end


function convFilter:processEnv( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   self:writeln( "Lns_LockEnvSync()" )
   self:popIndent(  )
   
   self.env:push( node )
   
   filter( node:get_block(), self, node )
   
   self.env:pop(  )
   
   self:pushIndent(  )
   self:writeln( "Lns_UnlockEnvSync()" )
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processBlockSub( node, opt )

   do
      local _switchExp = node:get_blockKind()
      if _switchExp == Nodes.BlockKind.Block then
         self:writeln( "{" )
      end
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
   
   do
      local _switchExp = node:get_blockKind()
      if _switchExp == Nodes.BlockKind.Block then
         self:writeln( "}" )
      end
   end
   
end


function convFilter:expList2Slice( subList, toStem )

   local function processExp( exp )
   
      if toStem and Ast.isClass( exp:get_expType() ) then
         self:write( string.format( "%s2Stem(", self:getTypeSymbolWithPrefix( exp:get_expType():get_nonnilableType() )) )
         filter( Nodes.getCastUnwraped( exp ), self, subList )
         self:write( ")" )
      else
       
         filter( exp, self, subList )
      end
      
   end
   
   local mRetIndex = _lune.nilacc( subList:get_mRetExp(), 'get_index', 'callmtd' )
   if mRetIndex and mRetIndex == 1 then
      local subExp = subList:get_expList()[1]
      if subExp:get_expType():get_kind() ~= Ast.TypeInfoKind.DDD then
         self:write( "Lns_2DDD(" )
         processExp( subExp )
         self:write( ")" )
      else
       
         processExp( subExp )
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
               processExp( subExp )
               self:write( ")" )
            else
             
               processExp( subExp )
            end
            
            self:write( "..." )
            break
         end
         
         if subIndex ~= 1 then
            self:write( ', ' )
         end
         
         processExp( subExp )
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
                     self:expList2Slice( exp2ddd:get_expList(), false )
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
         for index, dstType in ipairs( dstTypeList ) do
            if mRetIndex == index - 1 then
               break
            end
            
            if index ~= 1 then
               self:write( ', ' )
            end
            
            local exp
            
            if #expListNode:get_expList() < index then
               exp = self.noneNode
            else
             
               exp = expListNode:get_expList()[index]
            end
            
            if index == #dstTypeList and dstType:get_kind() == Ast.TypeInfoKind.DDD then
               if #expListNode:get_expList() < index or exp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
                  self:write( "[]LnsAny{}" )
               else
                
                  filter( exp, self, expListNode )
               end
               
            else
             
               if #expListNode:get_expList() < index or exp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
                  self:write( "nil" )
               elseif expListNode:get_expTypeList()[index]:get_kind() == Ast.TypeInfoKind.DDD then
                  self:write( "Lns_car(" )
                  filter( exp, self, expListNode )
                  self:write( ")" )
               else
                
                  filter( exp, self, expListNode )
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

   self:writeln( string.format( "// decl alge -- %s", node:get_expType():getTxt(  )) )
   self:writeln( string.format( "type %s = LnsAny", self:getTypeSymbol( node:get_algeType() )) )
   
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
         self:processSetFromExpList( self:getConvExpName( node:get_id(), argList ), node:get_methodType():get_argTypeInfoList(), argList )
      end
   end
   
   self:writeln( ")" )
end


function convFilter:processExpCallSuper( node, opt )

   self:write( string.format( "self.%s.%s(", self:getTypeSymbol( node:get_methodType():get_parentInfo() ), self:getFuncSymbol( node:get_methodType() )) )
   do
      local argList = node:get_expList()
      if argList ~= nil then
         self:processSetFromExpList( self:getConvExpName( node:get_id(), argList ), node:get_methodType():get_argTypeInfoList(), argList )
      end
   end
   
   self:write( ")" )
end


function convFilter:outputDummyReturn( retTypeInfoList )

   self:writeln( "// insert a dummy" )
   self:write( "    return " )
   for index, retType in ipairs( retTypeInfoList ) do
      if index > 1 then
         self:write( "," )
      end
      
      if isAnyType( retType ) then
         self:write( "nil" )
      else
       
         do
            local _switchExp = getOrgTypeInfo( retType )
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


function convFilter:outputDeclFuncInfo( node, declInfo )

   local funcType = node:get_expType()
   if funcType:get_abstractFlag() then
      return 
   end
   
   if declInfo:get_name() and not isClosure( funcType ) then
      self:writeln( string.format( "// %d: decl %s", node:get_pos().lineNo, self:getCanonicalName( funcType, false )) )
   end
   
   local convFunc = self:outputDeclFunc( _lune.newAlge( FuncInfo.DeclInfo, {node,declInfo}) )
   
   self:writeln( " {" )
   
   self:pushIndent(  )
   
   if isMain( funcType ) then
      self:writeln( string.format( "Lns_%s_init()", self:getModuleName( self.moduleTypeInfo, false )) )
   end
   
   
   if declInfo:get_has__func__Symbol() then
      local nameSpace = self:getCanonicalName( funcType:get_parentInfo(), false )
      local funcName
      
      do
         local name = declInfo:get_name()
         if name ~= nil then
            funcName = name.txt
         else
            funcName = "<anonymous>"
         end
      end
      
      local funcSym_ = _lune.unwrap( _lune.nilacc( funcType:get_scope(), 'getSymbolInfoChild', 'callmtd' , "__func__" ))
      self:writeln( string.format( '%s := "%s.%s"', self:getSymbolSym( funcSym_ ), nameSpace, funcName) )
   end
   
   
   for __index, convArg in ipairs( convFunc:get_argList() ) do
      if convArg:get_posForModToRef() then
         self:write( string.format( "%s := _%s", self:getSymbolSym( convArg ), self:getSymbolSym( convArg )) )
         self:outputAny2Type( convArg:get_typeInfo() )
         self:writeln( "" )
      end
      
   end
   
   
   self:popIndent(  )
   
   do
      local body = declInfo:get_body()
      if body ~= nil then
         filter( body, self, node )
         
         local retTypeInfoList = funcType:get_retTypeInfoList()
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
               self:outputDummyReturn( funcType:get_retTypeInfoList() )
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
         self:writeln( string.format( "type %s = %s", self:getTypeSymbol( node:get_enumType() ), self:type2gotype( node:get_enumType() )) )
         
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
         for __index, valName in ipairs( node:get_valueNameList() ) do
            self:writeln( string.format( "  %s,", self:getSymbol( _lune.newAlge( SymbolKind.Static, {node:get_enumType()}), valName.txt )) )
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
       
         self:write( "_" .. self:getSymbolSym( varSym ) )
      end
      
      table.insert( tempTypeList, Ast.builtinTypeStem_ )
   end
   
   if getExpListKind( tempTypeList, node:get_expList() ) == ExpListKind.Direct then
      for _7199 = #node:get_varSymList() + 1, #node:get_expList():get_expTypeList() do
         self:write( ", _" )
      end
      
   end
   
   self:write( " := " )
   self:processSetFromExpList( self:getConvExpName( node:get_id(), node:get_expList() ), tempTypeList, node:get_expList() )
   self:writeln( "" )
   self:write( "if " )
   local hasSym = false
   for __index, varSym in ipairs( node:get_varSymList() ) do
      if varSym:get_name() ~= "_" then
         if hasSym then
            self:write( " && " )
         end
         
         self:write( string.format( "_%s != nil", self:getSymbolSym( varSym )) )
         hasSym = true
      end
      
   end
   
   self:writeln( " {" )
   self:pushIndent(  )
   for index, varSym in ipairs( node:get_varSymList() ) do
      if varSym:get_name() ~= "_" then
         if varSym:hasAccess(  ) then
            self:write( string.format( "%s := _%s", self:getSymbolSym( varSym ), self:getSymbolSym( varSym )) )
            if node:get_expList():getExpTypeNoDDDAt( index ):get_nilable() then
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
            end
            
            
            local tmpVarTypeList = {}
            for index, _7237 in ipairs( node:get_symbolInfoList() ) do
               table.insert( tmpVarTypeList, expList:getExpTypeNoDDDAt( index ) )
            end
            
            
            if getExpListKind( tmpVarTypeList, expList ) == ExpListKind.Direct then
               for _7241 = #tmpVarTypeList + 1, #expList:get_expTypeList() do
                  self:write( ", _" )
               end
               
            end
            
            self:write( " := " )
            self:processSetFromExpList( self:getConvExpName( node:get_id(), expList ), tmpVarTypeList, expList )
            self:writeln( "" )
            self:write( "if " )
            local hasCond = false
            for index, varInfo in ipairs( node:get_varList() ) do
               if expList:getExpTypeNoDDDAt( index ):get_nilable() then
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
                     if expList:getExpTypeNoDDDAt( index ):get_nilable() then
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
                     if expList:getExpTypeNoDDDAt( index ):get_nilable() then
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
            local varTypeList = {}
            for index, symbolInfo in ipairs( node:get_symbolInfoList() ) do
               table.insert( varTypeList, symbolInfo:get_typeInfo() )
               if index > 1 then
                  self:write( "," )
               end
               
               if symbolInfo:get_posForModToRef() or Ast.isPubToExternal( symbolInfo:get_accessMode() ) then
                  self:write( string.format( "%s", self:getSymbolSym( symbolInfo )) )
               else
                
                  self:write( "_" )
               end
               
            end
            
            
            if getExpListKind( varTypeList, expList ) == ExpListKind.Direct then
               for _7271 = #varTypeList + 1, #expList:get_expTypeList() do
                  self:write( ", _" )
               end
               
            end
            
            self:write( " = " )
            self:processSetFromExpList( self:getConvExpName( node:get_id(), expList ), varTypeList, expList )
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
               self:write( string.format( "%s = _%s", self:getSymbolSym( varSym ), varSym:get_name()) )
               if expList:getExpTypeNoDDDAt( index ):get_nilable() then
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
         
         if getExpListKind( typeList, expList ) == ExpListKind.Direct then
            for _7300 = #node:get_symbolInfoList() + 1, #expList:get_expTypeList() do
               self:write( ",_" )
            end
            
         end
         
         self:write( " = " )
         self:processSetFromExpList( self:getConvExpName( node:get_id(), expList ), typeList, expList )
         self:writeln( "" )
         self:write( "if " )
         local hasCond = false
         for index, varSym in ipairs( node:get_symbolInfoList() ) do
            if expList:getExpTypeNoDDDAt( index ):get_nilable() then
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
         for _7431, symbol in ipairs( caseInfo:get_valParamNameList() ) do
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
   if node:get_val():get_posForModToRef() then
      self:writeln( string.format( "%s := %s", node:get_val():get_name(), workSym) )
   end
   
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
      for index, _7495 in ipairs( node:get_varList() ) do
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


function convFilter:outputForeachLua( node, sortFlag, exp, val, key, block )
   local __func__ = '@lune.@base.@convGo.convFilter.outputForeachLua'

   do
      local _switchExp = exp:get_expType():get_extedType():get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Map then
         self:writeln( "{" )
         self:pushIndent(  )
         local tmpExp = string.format( "_exp%d", node:get_id())
         local tmpKey = string.format( "_key%d", node:get_id())
         local tmpVal = string.format( "_val%d", node:get_id())
         local tmpIndex = string.format( "_index%d", node:get_id())
         local sorted = string.format( "_sorted%d", node:get_id())
         self:write( string.format( "%s := ", tmpExp) )
         filter( exp, self, node )
         self:writeln( "" )
         
         local tmpValName
         
         if val:hasAccess(  ) then
            tmpValName = tmpVal
         else
          
            tmpValName = "_"
         end
         
         if sortFlag then
            self:writeln( string.format( "%s := %s.SortMapKeyList( %s )", sorted, self.env:getLuavm( node:isThreading(  ) ), tmpExp) )
            self:writeln( string.format( "%s, %s := %s.Get1stFromMap()", tmpIndex, tmpKey, sorted) )
            self:writeln( string.format( "for %s != nil {", tmpIndex) )
            self:pushIndent(  )
         else
          
            self:writeln( string.format( "%s, %s := %s.Get1stFromMap()", tmpKey, tmpValName, tmpExp) )
            self:writeln( string.format( "for %s != nil {", tmpKey) )
            self:pushIndent(  )
         end
         
         do
            local keySym = key
            if keySym ~= nil then
               if keySym:hasAccess(  ) then
                  self:write( string.format( "%s := %s", self:getSymbolSym( keySym ), tmpKey) )
                  self:outputAny2Type( keySym:get_typeInfo() )
                  self:writeln( "" )
               end
               
            end
         end
         
         if val:hasAccess(  ) then
            if sortFlag then
               self:write( string.format( "%s := %s.GetAt( %s )", self:getSymbolSym( val ), tmpExp, tmpKey) )
            else
             
               self:write( string.format( "%s := %s", self:getSymbolSym( val ), tmpVal) )
            end
            
            self:outputAny2Type( val:get_typeInfo() )
            self:writeln( "" )
         end
         
         self:popIndent(  )
         
         filter( block, self, node )
         
         self:pushIndent(  )
         if not sortFlag then
            self:write( string.format( "%s, %s = ", tmpKey, tmpValName) )
            self:writeln( string.format( "%s.NextFromMap( %s )", tmpExp, tmpKey) )
         else
          
            self:write( string.format( "%s, %s = ", tmpIndex, tmpKey) )
            self:writeln( string.format( "%s.NextFromMap( %s )", sorted, tmpIndex) )
         end
         
         
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

   if node:get_exp():get_expType():get_srcTypeInfo():get_kind() == Ast.TypeInfoKind.Ext then
      self:outputForeachLua( node, false, node:get_exp(), node:get_val(), node:get_key(), node:get_block() )
      return 
   end
   
   
   
   
   
   local hasAccessLoopSym = _lune.nilacc( node:get_key(), 'get_posForModToRef', 'callmtd' ) or node:get_val():get_posForModToRef()
   
   self:write( "for " )
   local loopExpType = node:get_exp():get_expType()
   do
      local _switchExp = loopExpType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         local valName = self:getSymbolSym( node:get_val() )
         local itemType = loopExpType:get_itemTypeInfoList()[1]
         if hasAccessLoopSym then
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
            
            
            
            if valName ~= "_" then
               self:write( string.format( "_%s", valName) )
            else
             
               self:write( string.format( "%s", valName) )
            end
            
            
            self:write( " := " )
         end
         
         self:write( "range( " )
         filter( node:get_exp(), self, node )
         self:writeln( ".Items ) {" )
         self:pushIndent(  )
         do
            local key = node:get_key()
            if key ~= nil then
               if key:get_posForModToRef() then
                  self:writeln( string.format( "%s := _%s + 1", self:getSymbolSym( key ), key:get_name()) )
               end
               
            end
         end
         
         
         if valName ~= "_" then
            self:write( string.format( "%s := _%s", valName, valName) )
            self:outputStem2Type( itemType )
            self:writeln( "" )
         end
         
         
         self:popIndent(  )
      elseif _switchExp == Ast.TypeInfoKind.Map then
         local valName = self:getSymbolSym( node:get_val() )
         local itemType = loopExpType:get_itemTypeInfoList()[2]
         local keyType = loopExpType:get_itemTypeInfoList()[1]
         if hasAccessLoopSym then
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
            
            
            
            if valName ~= "_" then
               self:write( string.format( "_%s", valName) )
            else
             
               self:write( string.format( "%s", valName) )
            end
            
            
            self:write( " := " )
         end
         
         self:write( "range( " )
         filter( node:get_exp(), self, node )
         self:writeln( ".Items ) {" )
         self:pushIndent(  )
         do
            local key = node:get_key()
            if key ~= nil then
               
               if key:get_name() ~= "_" then
                  self:write( string.format( "%s := _%s", key:get_name(), key:get_name()) )
                  self:outputStem2Type( keyType )
                  self:writeln( "" )
               end
               
               
            end
         end
         
         
         if valName ~= "_" then
            self:write( string.format( "%s := _%s", valName, valName) )
            self:outputStem2Type( itemType )
            self:writeln( "" )
         end
         
         
         self:popIndent(  )
      elseif _switchExp == Ast.TypeInfoKind.Set then
         local valType = loopExpType:get_itemTypeInfoList()[1]
         local valName = self:getSymbolSym( node:get_val() )
         if hasAccessLoopSym then
            
            if valName ~= "_" then
               self:write( string.format( "_%s", valName) )
            else
             
               self:write( string.format( "%s", valName) )
            end
            
            
            self:write( " := " )
         end
         
         self:write( "range( " )
         filter( node:get_exp(), self, node )
         self:writeln( ".Items ) {" )
         self:pushIndent(  )
         
         if valName ~= "_" then
            self:write( string.format( "%s := _%s", valName, valName) )
            self:outputStem2Type( valType )
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

   if node:get_exp():get_expType():get_srcTypeInfo():get_kind() == Ast.TypeInfoKind.Ext then
      self:outputForeachLua( node, true, node:get_exp(), node:get_val(), node:get_key(), node:get_block() )
      return 
   end
   
   
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
         self:outputStem2Type( valSym:get_typeInfo() )
         self:writeln( "" )
      end
      
   end
   
   if keySym ~= nil then
      if keySym:get_posForModToRef() then
         self:write( string.format( "%s := _%s", key, key) )
         self:outputStem2Type( keySym:get_typeInfo() )
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
   if self:isSamePackage( node:get_expType():getModule(  ) ) then
      self:write( string.format( "New%s(", className) )
   else
    
      do
         local refTypeNode = _lune.__Cast( node:get_symbol(), 3, Nodes.RefTypeNode )
         if refTypeNode ~= nil then
            do
               local refNode = _lune.__Cast( refTypeNode:get_name(), 3, Nodes.RefFieldNode )
               if refNode ~= nil then
                  filter( refNode:get_prefix(), self, node )
                  self:write( "." )
               end
            end
            
         end
      end
      
      self:write( string.format( "New%s(", className) )
   end
   
   do
      local argList = node:get_argList()
      if argList ~= nil then
         local scope = _lune.unwrap( node:get_expType():get_scope())
         local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, Ast.ScopeAccess.Normal ))
         
         self:processSetFromExpList( self:getConvExpName( node:get_id(), argList ), initFuncType:get_argTypeInfoList(), argList )
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
      self:writeln( self:getTypeSymbolWithPrefix( superClassType ) )
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

   self:writeln( string.format( "func %s2Stem( obj LnsAny ) LnsAny {", self:getTypeSymbolWithPrefix( node:get_expType() )) )
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
   self:write( symbol )
   self:write( "() *" )
   self:write( symbol )
   self:writeln( "" )
   self:popIndent(  )
   self:writeln( "}" )
   
   self:writeln( string.format( "func %sDownCastF( multi ...LnsAny ) LnsAny {", symbol) )
   self:pushIndent(  )
   self:writeln( "if len( multi ) == 0 { return nil }" )
   self:writeln( "obj := multi[ 0 ]" )
   self:writeln( "if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }" )
   self:writeln( string.format( "work, ok := obj.(%sDownCast)", self:getTypeSymbolWithPrefix( node:get_expType() )) )
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
      for index, _7735 in ipairs( initFuncType:get_argTypeInfoList() ) do
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
         
         
         if retType:get_srcTypeInfo() ~= memberSym:get_typeInfo():get_srcTypeInfo() then
            if retType:get_kind() == Ast.TypeInfoKind.IF then
               if Ast.isClass( memberSym:get_typeInfo():get_srcTypeInfo() ) then
                  self:write( string.format( "%s%s.FP", prefix, self:getSymbolSym( memberSym )) )
               end
               
            elseif Ast.isClass( retType ) then
               self:write( string.format( "&%s%s.%s", prefix, self:getSymbolSym( memberSym ), self:getTypeSymbol( retType )) )
            elseif retType:get_kind() == Ast.TypeInfoKind.Alternate and retType:hasBase(  ) then
               self:write( string.format( "%s%s.FP", prefix, self:getSymbolSym( memberSym )) )
               self:outputStem2Type( retType )
            else
             
               if isAnyType( retType ) and Ast.isClass( memberSym:get_typeInfo() ) then
                  self:write( string.format( "%s%s.FP", prefix, self:getSymbolSym( memberSym )) )
               else
                
                  Util.err( "not support" )
               end
               
            end
            
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

   local workTypeInfo = getOrgTypeInfo( typeInfo )
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
         self:write( string.format( "%sSub, %s,", self:getFromStemName( workTypeInfo ), tostring( typeInfo:get_nilable())) )
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


function convFilter:outputAsyncItem( node )

   local classType = node:get_expType()
   local classScope = _lune.unwrap( classType:get_scope())
   local createSym = _lune.unwrap( classScope:getSymbolInfoChild( "_createPipe" ))
   self:outputDeclFunc( _lune.newAlge( FuncInfo.WithClass, {classType,createSym:get_typeInfo()}) )
   self:writeln( "{" )
   self:writeln( "   return NewLnspipe( arg1 )" )
   self:writeln( "}" )
end


function convFilter:outputMapping( node )

   local classType = node:get_expType()
   local className = self:getTypeSymbol( classType )
   self:writeln( string.format( "func (self *%s) ToMapSetup( obj *LnsMap ) *LnsMap {", className) )
   self:pushIndent(  )
   for __index, memberNode in ipairs( node:get_memberList() ) do
      if not memberNode:get_staticFlag() then
         self:writeln( string.format( 'obj.Items["%s"] = Lns_ToCollection( self.%s )', memberNode:get_symbolInfo():get_name(), self:getSymbolSym( memberNode:get_symbolInfo() )) )
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
   if not classType:get_abstractFlag() then
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
      self:writeln( '   return false, nil, "no map -- " + Lns_ToString(obj)' )
      self:writeln( "} else {" )
      self:writeln( '   objMap = work' )
      self:writeln( "}" )
      self:outputNewSetup( "newObj", classType )
      self:writeln( string.format( "return %sMain( newObj, objMap, paramList )", fromStemName) )
      self:popIndent(  )
      self:writeln( "}" )
   end
   
   
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
                  self:write( string.format( 'paramList[%d].Func( objMap.Items["%s"], %s, paramList[%d].Child', index - 1, memberNode:get_symbolInfo():get_name(), tostring( memberNode:get_expType():get_nilable()), index - 1) )
               end
               
            end
            
         else
          
            self:write( string.format( '%sSub( objMap.Items["%s"], %s, ', self:getFromStemName( memberNode:get_expType() ), memberNode:get_symbolInfo():get_name(), tostring( memberNode:get_expType():get_nilable())) )
            self:outputConvItemTypeList( memberNode:get_expType():get_itemTypeInfoList(), nil )
         end
         
         self:writeln( "); !ok {" )
         self:writeln( string.format( '   return false,nil,"%s:" + mess.(string)', memberNode:get_symbolInfo():get_name()) )
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


function convFilter:outputDummyAbstractMethod( classType, methodType )

   self:outputDeclFunc( _lune.newAlge( FuncInfo.WithClass, {classType,methodType}) )
   self:writeln( "{" )
   self:outputDummyReturn( methodType:get_retTypeInfoList() )
   self:writeln( "}" )
end


function convFilter:outputDummyAbstractMethodOfClass( classType )

   local name2typeMap = {}
   
   local scope = _lune.unwrap( classType:get_scope())
   scope:filterSymbolInfoIfField( scope, Ast.ScopeAccess.Normal, function ( symbolInfo )
   
      if symbolInfo:get_kind() == Ast.SymbolKind.Mtd and symbolInfo:get_typeInfo():get_abstractFlag() then
         do
            local methodType = scope:getTypeInfoChild( symbolInfo:get_name() )
            if methodType ~= nil then
               if methodType:get_abstractFlag() then
                  name2typeMap[symbolInfo:get_name()] = symbolInfo:get_typeInfo()
               end
               
            else
               name2typeMap[symbolInfo:get_name()] = symbolInfo:get_typeInfo()
            end
         end
         
      end
      
      return true
   end )
   
   do
      local __sorted = {}
      local __map = name2typeMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local methodType = __map[ __key ]
         do
            self:outputDummyAbstractMethod( classType, methodType )
         end
      end
   end
   
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
                  self:outputDeclFunc( _lune.newAlge( FuncInfo.WithClass, {node:get_expType(),funcType}) )
                  self:writeln( " {" )
                  if #funcType:get_retTypeInfoList() > 0 then
                     self:write( "    return " )
                  end
                  
                  self:write( string.format( "self.%s. ", self:getSymbolSym( adv:get_member():get_symbolInfo() )) )
                  if adv:get_member():get_expType():get_kind() == Ast.TypeInfoKind.Class then
                     self:write( "FP." )
                  end
                  
                  self:write( string.format( "%s( ", self:getSymbolSym( symbol )) )
                  for index, _7938 in ipairs( funcType:get_argTypeInfoList() ) do
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


function convFilter:processProtoClass( node, opt )

end


function convFilter:processDeclClass( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclClass'

   if node:isModule(  ) then
      return 
   end
   
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
            self:outputDummyAbstractMethodOfClass( node:get_expType() )
            self:outputAdvertise( node )
            
            if node:get_expType():isInheritFrom( self.processInfo, Ast.builtinTypeMapping, nil ) then
               self:outputMapping( node )
            end
            
            
            if node:get_expType():isInheritFrom( self.processInfo, Ast.builtinTypeAsyncItem, nil ) then
               self:outputAsyncItem( node )
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

CallKind.BuiltinCall = { "BuiltinCall", {{},{}}}
CallKind._name2Val["BuiltinCall"] = CallKind.BuiltinCall
CallKind.FormCall = { "FormCall"}
CallKind._name2Val["FormCall"] = CallKind.FormCall
CallKind.LuaCall = { "LuaCall"}
CallKind._name2Val["LuaCall"] = CallKind.LuaCall
CallKind.Normal = { "Normal"}
CallKind._name2Val["Normal"] = CallKind.Normal
CallKind.RunLoaded = { "RunLoaded"}
CallKind._name2Val["RunLoaded"] = CallKind.RunLoaded
CallKind.RuntimeCall = { "RuntimeCall", {{}}}
CallKind._name2Val["RuntimeCall"] = CallKind.RuntimeCall
CallKind.SortCall = { "SortCall", {{}}}
CallKind._name2Val["SortCall"] = CallKind.SortCall

function convFilter:outputCallPrefix( threading, callId, node, prefixNode, funcSymbol )

   local funcType = funcSymbol:get_typeInfo()
   local nilAccName = string.format( "%s_%d", self.moduleTypeInfo:get_rawTxt():gsub( "@", "" ), callId)
   
   local callKind = _lune.newAlge( CallKind.Normal)
   
   local extCallFlag
   
   if prefixNode ~= nil then
      extCallFlag = prefixNode:get_expType():get_nonnilableType():get_kind() == Ast.TypeInfoKind.Ext
   else
      extCallFlag = false
   end
   
   if extCallFlag then
      callKind = _lune.newAlge( CallKind.LuaCall)
   end
   
   
   local getEnvTxt = self.env:getEnv( threading )
   local function processNilAcc( workPrefixNode )
   
      if not node:hasNilAccess(  ) then
         return 
      end
      
      local retNum = #funcType:get_retTypeInfoList()
      do
         local _switchExp = retNum
         if _switchExp == 0 then
            self:write( string.format( "Lns_NilAccCall0( %s, func () {", getEnvTxt) )
         elseif _switchExp == 1 then
            
            self:write( string.format( "Lns_NilAccCall1( %s, func () LnsAny { return ", getEnvTxt) )
         else 
            
               if retNum <= MaxNilAccNum then
                  local anys = "LnsAny"
                  for _8006 = 2, retNum do
                     anys = string.format( "%s,LnsAny", anys)
                  end
                  
                  self:write( string.format( "Lns_NilAccCall%d( %s, func () (%s) { return ", retNum, getEnvTxt, anys) )
               else
                
                  local args = "LnsAny"
                  for _8010 = 2, retNum do
                     args = string.format( "%s,LnsAny", args)
                  end
                  
                  self:write( string.format( "lns_NilAccCall_%s( %s, func () (%s) { return ", nilAccName, getEnvTxt, args) )
               end
               
         end
      end
      
      if extCallFlag then
         if #funcType:get_retTypeInfoList() > 1 then
            self:write( string.format( "Lns_callExt%d( ", node:get_id()) )
         end
         
      end
      
      
      self:write( string.format( "%s.NilAccPop().(%s)", getEnvTxt, self:type2gotype( workPrefixNode:get_expType():get_nonnilableType() )) )
   end
   
   local closeParen = false
   if prefixNode ~= nil then
      if node:hasNilAccess(  ) then
         if #funcType:get_retTypeInfoList() >= 2 then
            if #funcType:get_retTypeInfoList() <= MaxNilAccNum then
               self:write( string.format( "Lns_NilAccFinCall%d( ", #funcType:get_retTypeInfoList()) )
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
            local runtime = self:getVM( threading, funcType )
            if runtime ~= nil then
               callKind = _lune.newAlge( CallKind.RuntimeCall, {prefixNode})
               self:write( runtime )
            end
         end
         
      else
       
         if not funcType:get_staticFlag() or funcSymbol:get_kind() == Ast.SymbolKind.Mbr then
            if node:hasNilAccess(  ) then
               if not prefixNode:hasNilAccess(  ) then
                  self:write( string.format( "%s.NilAccFin(", getEnvTxt) )
                  self:write( string.format( "%s.NilAccPush(", getEnvTxt) )
                  filter( prefixNode, self, node )
                  self:writeln( ") && " )
               else
                
                  filter( prefixNode, self, node )
                  self:writeln( "&&" )
               end
               
            else
             
               if extCallFlag then
                  if #funcType:get_retTypeInfoList() > 1 then
                     self:write( string.format( "Lns_callExt%d( ", node:get_id()) )
                  end
                  
               end
               
               filter( prefixNode, self, node )
            end
            
         else
          
            
         end
         
         
         processNilAcc( prefixNode )
         if prefixType:get_kind() == Ast.TypeInfoKind.Ext then
            self:write( string.format( '.CallMethod( "%s", Lns_2DDD', funcSymbol:get_name()) )
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
                           local runtime = self:getVM( threading, funcType )
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
                     
                        if funcSymbol:get_kind() == Ast.SymbolKind.Mbr then
                           self:write( "." )
                        end
                        
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

   local funcType = node:get_func():get_expType():get_nonnilableType()
   
   if funcType:get_kind() == Ast.TypeInfoKind.Method and funcType:get_parentInfo():get_kind() == Ast.TypeInfoKind.Enum and funcType:get_rawTxt() == "get__txt" then
      self:write( string.format( "%s(", self:getEnumGetTxtSym( _lune.unwrap( _lune.__Cast( funcType:get_parentInfo():get_aliasSrc(), 3, Ast.EnumTypeInfo )) )) )
      local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
      if  nil == fieldNode then
         local _fieldNode = fieldNode
      
         Util.err( "not support -- __func__" )
      end
      
      
      filter( fieldNode:get_prefix(), self, node )
      self:write( ")" )
      return 
   end
   
   
   local retGenerics
   
   if opt.parent:get_kind() == Nodes.NodeKind.get_StmtExp() then
      retGenerics = false
   else
    
      retGenerics = isRetGenerics( node )
      if retGenerics and #funcType:get_retTypeInfoList() ~= 1 then
         self:write( string.format( "%s(", self:getConvGenericsName( node )) )
      end
      
   end
   
   
   local closeParen = false
   local callKind = _lune.newAlge( CallKind.Normal)
   local withPrefix
   
   do
      local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
      if fieldNode ~= nil then
         withPrefix = true
         closeParen, callKind = self:outputCallPrefix( node:isThreading(  ), node:get_id(), fieldNode, fieldNode:get_prefix(), _lune.unwrap( fieldNode:get_symbolInfo()) )
      else
         withPrefix = false
         if Ast.isBuiltin( funcType:get_typeId() ) then
            local builtinFuncs = TransUnit.getBuiltinFunc(  )
            
            do
               local runtime = self:getVM( node:isThreading(  ), funcType )
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
               self:write( string.format( "%s.RunLoadedfunc", self.env:getLuavm( node:isThreading(  ) )) )
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
         else
          
            self:write( "Lns_2DDD(" )
            closeTxt = ")"
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
         self:processSetFromExpList( self:getConvExpName( node:get_id(), argList ), funcType:get_argTypeInfoList(), argList )
      end
   end
   
   
   if funcType:get_kind() == Ast.TypeInfoKind.Func and funcType:get_staticFlag() and funcType:get_parentInfo():isInheritFrom( self.processInfo, Ast.builtinTypeMapping ) and (funcType:get_rawTxt() == "_fromMap" or funcType:get_rawTxt() == "_fromStem" ) then
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
   if callKind == _lune.newAlge( CallKind.LuaCall) or callKind == _lune.newAlge( CallKind.RunLoaded) then
      if #funcType:get_retTypeInfoList() == 1 then
         if opt.parent:get_kind() ~= Nodes.NodeKind.get_StmtExp() then
            self:write( "[0]" )
            local retTypeList = _lune.unwrap( Ast.convToExtTypeList( self.processInfo, funcType:get_retTypeInfoList() ))
            self:outputAny2Type( retTypeList[1] )
         end
         
      elseif #funcType:get_retTypeInfoList() > 1 then
         self:write( ")" )
      end
      
   end
   
   
   if retGenerics then
      if #funcType:get_retTypeInfoList() == 1 then
         local retType = funcType:get_retTypeInfoList()[1]
         if isAnyType( retType ) then
            if not isAnyType( node:get_expType() ) then
               self:outputAny2Type( node:get_expType() )
            end
            
         elseif retType:get_kind() == Ast.TypeInfoKind.Alternate then
            if retType:get_srcTypeInfo() ~= node:get_expType():get_srcTypeInfo() then
               self:write( ".FP" )
               self:outputStem2Type( node:get_expType() )
            end
            
         end
         
      else
       
         self:write( ")" )
      end
      
   end
   
   
   if withPrefix and node:hasNilAccess(  ) then
      self:write( "})" )
      self:write( string.format( "/* %d:%d */", node:get_pos().lineNo, node:get_pos().column) )
      
      if opt.parent:hasNilAccess(  ) then
         
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
            elseif _switchExp == Ast.TypeInfoKind.Ext then
               do
                  local _switchExp = node:get_exp():get_expType():get_extedType():get_kind()
                  if _switchExp == Ast.TypeInfoKind.List then
                     filter( node:get_exp(), self, node )
                     self:write( ".Len()" )
                  else 
                     
                        Util.err( string.format( "%s: not support -- %s", __func__, node:get_exp():get_expType():getTxt(  )) )
                  end
               end
               
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
          
            self:write( string.format( "(%s)(", self:type2gotype( node:get_castType() )) )
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
         local typeName = self:getTypeSymbolWithPrefix( node:get_castType() )
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

   filter( node:get_exp1(), self, node )
   if getExpListKind( node:get_exp1():get_expTypeList(), node:get_exp2() ) == ExpListKind.Direct then
      for _8220 = #node:get_exp1():get_expTypeList() + 1, #node:get_exp2():get_expTypeList() do
         self:write( ",_" )
      end
      
   end
   
   self:write( " = " )
   self:processSetFromExpList( self:getConvExpName( node:get_id(), node:get_exp2() ), node:get_exp1():get_expTypeList(), node:get_exp2() )
   self:writeln( "" )
end


function convFilter:processExpSetItem( node, opt )

   filter( node:get_val(), self, node )
   self:write( ".Set(" )
   do
      local _matchExp = node:get_index()
      if _matchExp[1] == Nodes.IndexVal.NodeIdx[1] then
         local index = _matchExp[2][1]
      
         filter( index, self, node )
      elseif _matchExp[1] == Nodes.IndexVal.SymIdx[1] then
         local index = _matchExp[2][1]
      
         self:write( string.format( '"%s"', index) )
      end
   end
   
   self:write( "," )
   filter( node:get_exp2(), self, node )
   self:write( ")" )
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
   
   local getEnvTxt = self.env:getEnv( node:isThreading(  ) )
   local firstFlag = not isAndOr( parent )
   if firstFlag then
      self:writeln( string.format( "%s.PopVal( %s.IncStack() ||", getEnvTxt, getEnvTxt) )
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
    
      self:write( string.format( "%s.SetStackVal( ", getEnvTxt) )
      filter( node:get_exp1(), self, node )
      self:write( ") " )
   end
   
   self:writeln( opCC )
   if isAndOr( node:get_exp2() ) then
      filter( node:get_exp2(), self, node )
   else
    
      self:write( string.format( "%s.SetStackVal( ", getEnvTxt) )
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
         self:processAndOr( node, opTxt, opt.parent )
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
               if node:get_expType():get_kind() == Ast.TypeInfoKind.Module then
                  self:write( node:get_symbolInfo():get_name() )
               else
                
                  self:write( self:getTypeSymbol( node:get_expType() ) )
                  
               end
               
            else 
               
                  self:write( node:get_symbolInfo():get_name() )
            end
         end
         
      end
      
   end
   
end


function convFilter:processExpRefItem( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpRefItem'

   local getEnvTxt = self.env:getEnv( node:isThreading(  ) )
   local prefixType = node:get_val():get_expType():get_nonnilableType()
   do
      local _switchExp = prefixType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Ext then
         if node:get_nilAccess() then
            self:write( string.format( "%s.NilAccFin( %s.NilAccPush( ", getEnvTxt, getEnvTxt) )
            filter( node:get_val(), self, node )
            self:write( ") && " )
            self:write( string.format( "%s.NilAccPush( %s.NilAccPop().(*Lns_luaValue)", getEnvTxt, getEnvTxt) )
         else
          
            filter( node:get_val(), self, node )
            if prefixType:get_extedType():get_kind() == Ast.TypeInfoKind.Stem then
               self:write( ".(*Lns_luaValue)" )
            end
            
         end
         
         self:write( ".GetAt(" )
         do
            local index = node:get_index()
            if index ~= nil then
               filter( index, self, node )
            else
               self:write( string.format( '"%s"', str2gostr( _lune.unwrap( node:get_symbol()) )) )
            end
         end
         
         self:write( ")" )
         self:outputStem2Type( node:get_expType() )
      elseif _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         if node:get_nilAccess() then
            if not node:get_val():hasNilAccess(  ) then
               self:writeln( string.format( "%s.NilAccFin( %s.NilAccPush(", getEnvTxt, getEnvTxt) )
               filter( node:get_val(), self, node )
               self:writeln( ") && " )
            else
             
               filter( node:get_val(), self, node )
               self:writeln( "&&" )
            end
            
            self:write( string.format( "%s.NilAccPush( %s.NilAccPop().(*LnsList)", getEnvTxt, getEnvTxt) )
            self:write( ".GetAt(" )
            filter( _lune.unwrap( node:get_index()), self, node )
            self:write( ")" )
            self:outputStem2Type( node:get_expType() )
            self:write( "))" )
         else
          
            filter( node:get_val(), self, node )
            self:write( ".GetAt(" )
            filter( _lune.unwrap( node:get_index()), self, node )
            self:write( ")" )
            self:outputStem2Type( node:get_expType() )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Map then
         if node:get_nilAccess() then
            if not node:get_val():hasNilAccess(  ) then
               self:writeln( string.format( "%s.NilAccFin( %s.NilAccPush(", getEnvTxt, getEnvTxt) )
               filter( node:get_val(), self, node )
               self:writeln( ") && " )
            else
             
               filter( node:get_val(), self, node )
               self:writeln( "&&" )
            end
            
            self:write( string.format( "%s.NilAccPush( %s.NilAccPop().(*LnsMap)", getEnvTxt, getEnvTxt) )
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
         self:outputStem2Type( node:get_expType() )
         if node:get_nilAccess() then
            self:write( "))" )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Stem then
         self:write( "Lns_FromStemGetAt(" )
         filter( node:get_val(), self, node )
         self:write( "," )
         do
            local index = node:get_index()
            if index ~= nil then
               filter( index, self, node )
            else
               self:write( string.format( '"%s"', str2gostr( _lune.unwrap( node:get_symbol()) )) )
            end
         end
         
         self:write( string.format( ", %s )", tostring( node:get_nilAccess())) )
         self:outputStem2Type( node:get_expType() )
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
   
   
   local getEnvTxt = self.env:getEnv( node:isThreading(  ) )
   
   local openParenNum
   
   if not node:hasNilAccess(  ) then
      openParenNum = 0
      if not node:get_prefix():hasNilAccess(  ) then
         filter( node:get_prefix(), self, node )
      else
       
         Util.err( string.format( "%s: not support", __func__) )
      end
      
   else
    
      if not node:get_prefix():hasNilAccess(  ) then
         self:write( string.format( "%s.NilAccFin(", getEnvTxt) )
         self:write( string.format( "%s.NilAccPush(", getEnvTxt) )
         filter( node:get_prefix(), self, node )
         self:writeln( ") && " )
      else
       
         filter( node:get_prefix(), self, node )
         self:writeln( "&&" )
      end
      
      self:write( string.format( "%s.NilAccPush(", getEnvTxt) )
      if opt.parent:hasNilAccess(  ) then
         openParenNum = 1
      else
       
         openParenNum = 2
      end
      
      self:write( string.format( "%s.NilAccPop().(%s)", getEnvTxt, self:type2gotype( node:get_prefix():get_expType():get_nonnilableType() )) )
   end
   
   
   self:write( "." )
   do
      local symbol = node:get_symbolInfo()
      if symbol ~= nil then
         if node:get_prefix():get_expType():get_kind() == Ast.TypeInfoKind.Ext then
            self:write( string.format( 'GetAt( "%s" )', symbol:get_name()) )
            self:outputAny2Type( node:get_expType() )
         else
          
            self:write( self:getSymbolSym( symbol ) )
            local orgSym = symbol:getOrg(  )
            if orgSym:get_kind() == Ast.SymbolKind.Mbr and orgSym:get_typeInfo():get_kind() == Ast.TypeInfoKind.Alternate and isAnyType( orgSym:get_typeInfo() ) then
               self:outputAny2Type( node:get_expType() )
            end
            
         end
         
      else
         Util.err( string.format( "not support -- %s", __func__) )
      end
   end
   
   
   for _8369 = 1, openParenNum do
      self:write( ")" )
   end
   
end


function convFilter:processExpOmitEnum( node, opt )

   self:write( self:getSymbol( _lune.newAlge( SymbolKind.Static, {node:get_expType():get_aliasSrc()}), node:get_valInfo():get_name() ) )
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
          
            local closeParen = self:outputCallPrefix( node:isThreading(  ), node:get_id(), node, node:get_prefix(), symbolInfo )
            self:write( "()" )
            local retType = symbolInfo:get_typeInfo():get_retTypeInfoList()[1]
            if retType:get_kind() == Ast.TypeInfoKind.Alternate and not retType:hasBase(  ) then
               self:outputAny2Type( node:get_expType() )
            end
            
            
            if node:hasNilAccess(  ) then
               self:write( "})" )
               if opt.parent:hasNilAccess(  ) then
                  
               else
                
                  self:write( ")" )
               end
               
            end
            
            if closeParen then
               self:write( ")" )
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


function convFilter:processTestCase( node, opt )

   if not self.enableTest then
      return 
   end
   
   self:writeln( string.format( "func lns_test_%s_%s( %s *Testing_Ctrl ) {", self:getModuleName( self.moduleTypeInfo, false ), node:get_name().txt, node:get_ctrlName()) )
   
   filter( node:get_block(), self, node )
   
   self:writeln( "}" )
end


function convFilter:processTestBlock( node, opt )

   local stmtList = node:get_stmtList()
   for __index, statement in ipairs( stmtList ) do
      if not _lune._Set_has(ignoreNodeInInnerBlockSet, statement:get_kind() ) then
         filter( statement, self, node )
      end
      
   end
   
end



function convFilter:processProvide( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processProvide'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processAlias( node, opt )

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
         self:expList2Slice( expList, true )
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
         self:expList2Slice( expList, true )
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
         self:expList2Slice( expList, true )
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
         self:write( string.format( '%s.String_format(%s, ', self.env:getLuavm( node:isThreading(  ) ), str2gostr( txt )) )
         self:processSetFromExpList( self:getConvExpName( node:get_id(), expList ), {Ast.builtinTypeDDD}, expList )
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
      
      elseif _matchExp[1] == LuneControl.Pragma.limit_conv_code[1] then
         local _ = _matchExp[2][1]
      
      elseif _matchExp[1] == LuneControl.Pragma.use_async[1] then
      
      elseif _matchExp[1] == LuneControl.Pragma.run_async_pipe[1] then
      
         self:writeln( "go self.LoopMain()" )
      end
   end
   
end


function convFilter:processAbbr( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processAbbr'

   Util.err( string.format( "not support -- %s:%d", __func__, node:get_pos().lineNo) )
end


local function createFilter( enableTest, streamName, stream, ast, option )

   return convFilter.new(enableTest, streamName, stream, ast, option)
end
_moduleObj.createFilter = createFilter





return _moduleObj
