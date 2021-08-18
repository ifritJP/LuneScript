--lune/base/convLua.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@convLua'
local _lune = {}
if _lune6 then
   _lune = _lune6
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

function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
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

if not _lune6 then
   _lune6 = _lune
end


local Ver = _lune.loadModule( 'lune.base.Ver' )
local Str = _lune.loadModule( 'lune.base.Str' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Util = _lune.loadModule( 'lune.base.Util' )
local AstInfo = _lune.loadModule( 'lune.base.AstInfo' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local LuaMod = _lune.loadModule( 'lune.base.LuaMod' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Types = _lune.loadModule( 'lune.base.Types' )
local Log = _lune.loadModule( 'lune.base.Log' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )
local LnsOpt = _lune.loadModule( 'lune.base.Option' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Builtin = _lune.loadModule( 'lune.base.Builtin' )

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

ConvMode.Convert = 0
ConvMode._val2NameMap[0] = 'Convert'
ConvMode.__allList[1] = ConvMode.Convert
ConvMode.ConvMeta = 1
ConvMode._val2NameMap[1] = 'ConvMeta'
ConvMode.__allList[2] = ConvMode.ConvMeta


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

local function getSymbolTxt( symbolInfo )

   if symbolInfo:get_name() == "_" then
      do
         local annonymous = _lune.__Cast( symbolInfo, 3, Ast.AnonymousSymbolInfo )
         if annonymous ~= nil then
            return string.format( "_%d", annonymous:get_anonymousId())
         else
            Util.err( string.format( "can't cast to AnonymousSymbolInfo. -- %s:%s", _lune.nilacc( symbolInfo:get_pos(), "lineNo" ), _lune.nilacc( symbolInfo:get_pos(), "column" )) )
         end
      end
      
   end
   
   return symbolInfo:get_name()
end

local Option = {}
_moduleObj.Option = Option
function Option.setmeta( obj )
  setmetatable( obj, { __index = Option  } )
end
function Option.new( mainModule )
   local obj = {}
   Option.setmeta( obj )
   if obj.__init then
      obj:__init( mainModule )
   end
   return obj
end
function Option:__init( mainModule )

   self.mainModule = mainModule
end
function Option:get_mainModule()
   return self.mainModule
end


local ConvFilter = {}
setmetatable( ConvFilter, { __index = Nodes.Filter,ifList = {oStream,} } )
function ConvFilter.new( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, processInfo, moduleSymbolKind, builtinFunc, useLuneRuntime, targetLuaVer, enableTest, useIpairs, option )
   local obj = {}
   ConvFilter.setmeta( obj )
   if obj.__init then obj:__init( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, processInfo, moduleSymbolKind, builtinFunc, useLuneRuntime, targetLuaVer, enableTest, useIpairs, option ); end
   return obj
end
function ConvFilter:__init(streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, processInfo, moduleSymbolKind, builtinFunc, useLuneRuntime, targetLuaVer, enableTest, useIpairs, option) 
   Nodes.Filter.__init( self,true, moduleTypeInfo, moduleTypeInfo:get_scope())
   
   
   if stream == metaStream then
      Util.err( "streamName == stream" )
   end
   
   
   self.option = option
   self.builtinFunc = builtinFunc
   self.moduleType2SymbolMap = {}
   self.processInfo = processInfo
   self.enableTest = enableTest
   self.macroVarSymSet = {}
   self.needModuleObj = true
   self.indentQueue = {0}
   self.moduleSymbolKind = moduleSymbolKind
   self.macroDepth = 0
   self.streamName = streamName
   self.stream = stream
   self.metaStream = metaStream
   self.outMetaFlag = false
   self.convMode = convMode
   self.inMacro = inMacro
   self.curLineNo = 1
   self.classId2TypeInfo = {}
   self.classId2MemberList = {}
   self.pubVarName2InfoMap = {}
   self.pubFuncName2InfoMap = {}
   self.pubEnumId2EnumTypeInfo = {}
   self.pubAlgeId2AlgeTypeInfo = {}
   self.needIndent = false
   self.moduleTypeInfo = moduleTypeInfo
   self.useLuneRuntime = useLuneRuntime
   self.targetLuaVer = targetLuaVer
   self.useIpairs = useIpairs
   
   self.builtinSym2code = {[builtinFunc.__lns_runmode_Sync_sym] = string.format( "%d", 0), [builtinFunc.__lns_runmode_Queue_sym] = string.format( "%d", 1), [builtinFunc.__lns_runmode_Skip_sym] = string.format( "%d", 2), [builtinFunc.__lns_capability_async_sym] = "false"}
end
function ConvFilter:get_indent(  )

   if #self.indentQueue > 0 then
      return self.indentQueue[#self.indentQueue]
   end
   
   return 0
end
function ConvFilter:getCanonicalName( typeInfo, localFlag )

   local enumName = typeInfo:getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager(), localFlag )
   local moduleType = typeInfo:get_genSrcTypeInfo():get_srcTypeInfo():getModule(  )
   local canonical = (enumName:gsub( "&", "" ) )
   do
      local assignSym = self.moduleType2SymbolMap[moduleType]
      if assignSym ~= nil then
         if assignSym:get_isLazyLoad() then
            local index = _lune.unwrap( canonical:find( ".", 1, true ))
            return string.format( "%s().%s", canonical:sub( 1, index - 1 ), canonical:sub( index + 1 ))
         end
         
      end
   end
   
   return canonical
end
function ConvFilter:getFullName( typeInfo )

   return self:getCanonicalName( typeInfo, true )
end
function ConvFilter:close(  )

end
function ConvFilter:flush(  )

end
function ConvFilter:writeRaw( txt )

   local stream = self.stream
   if self.outMetaFlag then
      stream = self.metaStream
   end
   
   if self.needIndent then
      stream:write( string.rep( " ", self:get_indent() ) )
      self.needIndent = false
   end
   
   stream:write( txt )
end
function ConvFilter:write( txt )

   local stream = self.stream
   if self.outMetaFlag then
      stream = self.metaStream
   end
   
   for __index, line in ipairs( Str.getLineList( txt ) ) do
      if self.needIndent then
         stream:write( string.rep( " ", self:get_indent() ) )
         self.needIndent = false
      end
      
      
      if #line > 0 and string.byte( line, #line ) == 10 then
         self.curLineNo = self.curLineNo + 1
      end
      
      stream:write( line )
   end
   
   return self, nil
end
function ConvFilter.setmeta( obj )
  setmetatable( obj, { __index = ConvFilter  } )
end


local function filter( node, filter, parent )

   node:processFilter( filter, Opt.new(parent) )
end

local stepIndent = 3



function ConvFilter:processBlankLine( node, opt )

end


function ConvFilter:processDeclForm( node, opt )

end


function ConvFilter:processProtoMethod( node, opt )

end


function ConvFilter:pushIndent( newIndent )

   local indent = _lune.unwrapDefault( newIndent, self:get_indent() + stepIndent)
   table.insert( self.indentQueue, indent )
end


function ConvFilter:popIndent(  )

   if #self.indentQueue == 0 then
      Util.err( "self.indentQueue == 0" )
   end
   
   table.remove( self.indentQueue )
end


function ConvFilter:writeln( txt )

   self:write( txt )
   self:writeRaw( "\n" )
   self.needIndent = true
end


function ConvFilter:processNone( node, opt )

end


function ConvFilter:processShebang( node, opt )

end



function ConvFilter:processImport( node, opt )

   local info = node:get_info()
   local modulePath = info:get_modulePath()
   local modSym = modulePath:gsub( ".*%.", "" )
   modSym = info:get_assignName()
   self:writeRaw( string.format( "local %s = _lune.", modSym) )
   do
      local _switchExp = info:get_lazy()
      if _switchExp == Nodes.LazyLoad.Off then
         self:writeRaw( "loadModule" )
      elseif _switchExp == Nodes.LazyLoad.On or _switchExp == Nodes.LazyLoad.Auto then
         self:writeRaw( "_lazyImport" )
      end
   end
   
   self:writeRaw( string.format( "( '%s' )", modulePath) )
end



local ExportIdKind = {}
ExportIdKind._val2NameMap = {}
function ExportIdKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ExportIdKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ExportIdKind._from( val )
   if ExportIdKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ExportIdKind.__allList = {}
function ExportIdKind.get__allList()
   return ExportIdKind.__allList
end

ExportIdKind.Discarded = 0
ExportIdKind._val2NameMap[0] = 'Discarded'
ExportIdKind.__allList[1] = ExportIdKind.Discarded
ExportIdKind.Depend = 1
ExportIdKind._val2NameMap[1] = 'Depend'
ExportIdKind.__allList[2] = ExportIdKind.Depend
ExportIdKind.Normal = 2
ExportIdKind._val2NameMap[2] = 'Normal'
ExportIdKind.__allList[3] = ExportIdKind.Normal

function ConvFilter:outputMeta( node )

   if self.convMode == ConvMode.Convert then
      return 
   end
   
   
   self.outMetaFlag = true
   
   self:writeln( "local _moduleObj = {}" )
   
   self:writeln( "----- meta -----" )
   
   self:writeln( string.format( "_moduleObj.__version = '%s'", Ver.version) )
   self:writeln( string.format( "_moduleObj.__formatVersion = '%s'", Ver.metaVersion) )
   self:writeln( string.format( "_moduleObj.__buildId = %q", node:get_moduleId():getNextModuleId(  ):get_idStr()) )
   self:writeln( string.format( "_moduleObj.__enableTest = %s", self.enableTest) )
   self:writeln( string.format( "_moduleObj.__hasTest = %s", #node:get_nodeManager():getTestCaseNodeList(  ) ~= 0) )
   
   self:writeRaw( "_moduleObj.__lazyModuleList = {" )
   do
      local firstFlag = true
      for __index, declClass in ipairs( node:get_nodeManager():getDeclClassNodeList(  ) ) do
         if declClass:get_lazyLoad() ~= Nodes.LazyLoad.Off and Ast.isPubToExternal( declClass:get_accessMode() ) then
            if not firstFlag then
               self:writeRaw( "," )
            else
             
               firstFlag = false
            end
            
            self:writeRaw( string.format( "%d", declClass:get_expType():get_typeId().id) )
         end
         
      end
      
   end
   
   self:writeln( "}" )
   
   local importModuleType2Index = {}
   local importProcessInfo2Index = {}
   
   importProcessInfo2Index[Ast.getRootProcessInfoRo(  )] = frontInterface.getRootDependModId(  )
   importProcessInfo2Index[self.processInfo] = 0
   importProcessInfo2Index[self.processInfo:get_orgInfo()] = 0
   
   local importNameMap = {}
   do
      
      for __index, exportInfo in pairs( node:get_importModule2moduleInfo() ) do
         
         importNameMap[exportInfo:get_fullName()] = exportInfo
      end
      
      
      local index = 0
      do
         local __sorted = {}
         local __map = importNameMap
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, __key in ipairs( __sorted ) do
            local exportInfo = __map[ __key ]
            do
               index = index + 1
               importModuleType2Index[exportInfo:get_moduleTypeInfo()] = index
               importProcessInfo2Index[exportInfo:get_processInfo()] = index
               importProcessInfo2Index[exportInfo:get_processInfo():get_orgInfo()] = index
            end
         end
      end
      
   end
   
   
   local typeId2TypeInfo = {}
   do
      local work = node:get_moduleTypeInfo()
      while Ast.TypeInfo.hasParent( work ) do
         typeId2TypeInfo[work:get_typeId()] = work
         work = work:get_parentInfo()
      end
      
   end
   
   
   local pickupClassMap = {}
   
   local function checkExportTypeInfo( typeInfo )
   
      local moduleTypeInfo = typeInfo:getModule(  )
      local typeId = typeInfo:get_typeId()
      return typeId2TypeInfo[typeId] and not Ast.isBuiltin( typeId.id ) and (moduleTypeInfo:hasRouteNamespaceFrom( node:get_moduleTypeInfo() ) or typeInfo:get_srcTypeInfo() ~= typeInfo or moduleTypeInfo:equals( self.processInfo, Ast.headTypeInfo ) )
   end
   
   local function isDependOnExt( typeInfo )
   
      if Ast.isExtId( typeInfo ) then
         return true
      end
      
      return self.moduleTypeInfo:get_processInfo() ~= typeInfo:get_processInfo()
   end
   
   local function pickupTypeId( typeInfo, forceFlag, pickupChildFlag )
   
      
      if typeInfo:get_typeId() == Ast.rootTypeIdInfo then
         return 
      end
      
      
      if not forceFlag and not Ast.isPubToExternal( typeInfo:get_accessMode() ) then
         return 
      end
      
      
      if typeId2TypeInfo[typeInfo:get_typeId(  )] then
         
         if isDependOnExt( typeInfo ) then
            
            return 
         end
         
         
         if pickupChildFlag and not typeInfo:get_nilable() then
            for __index, itemTypeInfo in ipairs( typeInfo:get_children() ) do
               if Ast.isPubToExternal( itemTypeInfo:get_accessMode() ) then
                  do
                     local _switchExp = itemTypeInfo:get_kind()
                     if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc or _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Method then
                        pickupTypeId( itemTypeInfo, true, true )
                     end
                  end
                  
               end
               
            end
            
         end
         
         return 
      end
      
      
      typeId2TypeInfo[typeInfo:get_typeId(  )] = typeInfo
      if typeInfo:isModule(  ) then
         return 
      end
      
      if Ast.isBuiltin( typeInfo:get_srcTypeInfo():get_typeId().id ) then
         return 
      end
      
      
      if typeInfo:get_kind() == Ast.TypeInfoKind.Ext then
         pickupTypeId( typeInfo:get_extedType(), true, true )
      end
      
      
      if typeInfo ~= typeInfo:get_srcTypeInfo() then
         pickupTypeId( typeInfo:get_srcTypeInfo(), true, false )
      elseif typeInfo:get_nilable() then
         pickupTypeId( typeInfo:get_nonnilableType(), true, false )
      else
       
         if isDependOnExt( typeInfo ) then
            
            return 
         end
         
         
         if typeInfo:get_kind() == Ast.TypeInfoKind.Class or typeInfo:get_kind() == Ast.TypeInfoKind.IF then
            pickupClassMap[typeInfo:get_typeId()] = typeInfo
         end
         
         
         if not typeInfo:get_externalFlag() then
            do
               local _switchExp = typeInfo:get_kind()
               if _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc or _switchExp == Ast.TypeInfoKind.Alge or _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.Map or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array or _switchExp == Ast.TypeInfoKind.Alternate or _switchExp == Ast.TypeInfoKind.Box then
                  pickupTypeId( typeInfo:get_nilableTypeInfo(), true, false )
               end
            end
            
         end
         
         
         local parentInfo = typeInfo:get_parentInfo()
         pickupTypeId( parentInfo, true, false )
         
         pickupTypeId( typeInfo:get_genSrcTypeInfo(), true, false )
         
         local baseInfo = typeInfo:get_baseTypeInfo()
         if baseInfo:get_typeId() ~= Ast.rootTypeIdInfo then
            pickupTypeId( baseInfo, true, true )
         end
         
         for __index, itemTypeInfo in ipairs( typeInfo:get_interfaceList() ) do
            pickupTypeId( itemTypeInfo, true, true )
         end
         
         
         for __index, itemTypeInfo in ipairs( typeInfo:get_itemTypeInfoList() ) do
            pickupTypeId( itemTypeInfo, true, false )
         end
         
         for __index, itemTypeInfo in ipairs( typeInfo:get_argTypeInfoList() ) do
            pickupTypeId( itemTypeInfo, true, false )
         end
         
         for __index, itemTypeInfo in ipairs( typeInfo:get_retTypeInfoList() ) do
            pickupTypeId( itemTypeInfo, true, true )
         end
         
         if pickupChildFlag then
            for __index, itemTypeInfo in ipairs( typeInfo:get_children() ) do
               if itemTypeInfo:get_accessMode() == Ast.AccessMode.Pub then
                  do
                     local _switchExp = itemTypeInfo:get_kind()
                     if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc or _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Method then
                        pickupTypeId( itemTypeInfo, true, true )
                     end
                  end
                  
               end
               
            end
            
         end
         
      end
      
   end
   
   local validChildrenSet = {}
   
   do
      local typeInfo = self.moduleTypeInfo
      while Ast.TypeInfo.hasParent( typeInfo ) do
         validChildrenSet[typeInfo:get_parentInfo()] = {[typeInfo:get_typeId()] = typeInfo}
         typeInfo = typeInfo:get_parentInfo()
      end
      
      pickupTypeId( self.moduleTypeInfo, true )
   end
   
   
   local typeId2ClassMap = node:get_typeId2ClassMap(  )
   for __index, namespaceInfo in pairs( typeId2ClassMap ) do
      if namespaceInfo.typeInfo:get_accessMode(  ) == Ast.AccessMode.Pub and not namespaceInfo.typeInfo:get_externalFlag() then
         pickupClassMap[namespaceInfo.typeInfo:get_typeId()] = namespaceInfo.typeInfo
      end
      
   end
   
   
   local classId2TypeInfo = {}
   for idInfo, classTypeInfo in pairs( self.classId2TypeInfo ) do
      classId2TypeInfo[idInfo.id] = classTypeInfo
   end
   
   
   local serializeInfo = Ast.SerializeInfo.new(importProcessInfo2Index, {})
   
   self:writeln( "local __typeId2ClassInfoMap = {}" )
   self:writeln( "_moduleObj.__typeId2ClassInfoMap = __typeId2ClassInfoMap" )
   
   do
      local __sorted = {}
      local __map = classId2TypeInfo
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, classTypeId in ipairs( __sorted ) do
         local classTypeInfo = __map[ classTypeId ]
         do
            if classTypeInfo:get_accessMode() == Ast.AccessMode.Pub then
               pickupTypeId( classTypeInfo, true, validChildrenSet[classTypeInfo] == nil and not classTypeInfo:get_externalFlag() )
               pickupClassMap[classTypeInfo:get_typeId()] = nil
               
               self:writeln( "do" )
               self:pushIndent(  )
               self:writeln( string.format( "local __classInfo%d = {}", classTypeId) )
               self:writeln( string.format( "__typeId2ClassInfoMap[ %d ] = __classInfo%d", classTypeId, classTypeId) )
               
               for __index, memberNode in ipairs( _lune.unwrap( self.classId2MemberList[classTypeInfo:get_typeId()]) ) do
                  if Ast.isPubToExternal( memberNode:get_accessMode() ) then
                     local memberName = memberNode:get_name().txt
                     local memberTypeInfo = memberNode:get_expType(  )
                     self:writeln( string.format( "__classInfo%d.%s = {", classTypeId, memberName) )
                     self:writeln( string.format( "  name='%s', staticFlag = %s, mutMode = %d,", memberName, memberNode:get_staticFlag(), memberNode:get_symbolInfo():get_mutMode()) .. string.format( "accessMode = %d, typeId = %s }", memberNode:get_accessMode(), serializeInfo:serializeId( memberTypeInfo:get_typeId() )) )
                     
                     pickupTypeId( memberTypeInfo, true )
                  end
                  
               end
               
               
               self:popIndent(  )
               self:writeln( "end" )
            end
            
         end
      end
   end
   
   
   local pickupedClassMap = {}
   while true do
      local workClassMap = {}
      local hasWorkClassFlag = false
      for classTypeId, classTypeInfo in pairs( pickupClassMap ) do
         if not pickupedClassMap[classTypeId] then
            pickupedClassMap[classTypeId] = classTypeInfo
            if checkExportTypeInfo( classTypeInfo ) then
               workClassMap[classTypeId.id] = classTypeInfo
               hasWorkClassFlag = true
            end
            
         end
         
      end
      
      
      if not hasWorkClassFlag then
         break
      end
      
      do
         local __sorted = {}
         local __map = workClassMap
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, classTypeId in ipairs( __sorted ) do
            local classTypeInfo = __map[ classTypeId ]
            do
               if not Ast.isBuiltin( classTypeId ) then
                  local scope = classTypeInfo:get_scope()
                  if  nil == scope then
                     local _scope = scope
                  
                     Util.err( string.format( "%s.scope is nil", classTypeInfo:getTxt(  )) )
                  end
                  
                  
                  pickupTypeId( classTypeInfo, true, validChildrenSet[classTypeInfo] == nil and not classTypeInfo:get_externalFlag() )
                  
                  if checkExportTypeInfo( classTypeInfo ) then
                     
                     self:writeln( "do" )
                     self:pushIndent(  )
                     self:writeln( string.format( "local __classInfo%d = {}", classTypeId) )
                     self:writeln( string.format( "__typeId2ClassInfoMap[ %d ] = __classInfo%d", classTypeId, classTypeId) )
                     
                     do
                        local __sorted = {}
                        local __map = scope:get_symbol2SymbolInfoMap()
                        for __key in pairs( __map ) do
                           table.insert( __sorted, __key )
                        end
                        table.sort( __sorted )
                        for __index, fieldName in ipairs( __sorted ) do
                           local symbolInfo = __map[ fieldName ]
                           do
                              local typeInfo = symbolInfo:get_typeInfo()
                              if symbolInfo:get_kind() == Ast.SymbolKind.Mbr or symbolInfo:get_kind() == Ast.SymbolKind.Var then
                                 if symbolInfo:get_accessMode() == Ast.AccessMode.Pub then
                                    self:writeln( string.format( "__classInfo%d.%s = {", classTypeId, fieldName) )
                                    self:writeln( string.format( "  name='%s', staticFlag = %s, ", fieldName, symbolInfo:get_staticFlag()) .. string.format( "accessMode = %d, typeId = %d }", symbolInfo:get_accessMode(), typeInfo:get_typeId().id) )
                                    pickupTypeId( typeInfo )
                                 end
                                 
                              end
                              
                           end
                        end
                     end
                     
                     
                     self:popIndent(  )
                     self:writeln( "end" )
                  end
                  
               end
               
            end
         end
      end
      
   end
   
   
   self:writeln( "local __macroName2InfoMap = {}" )
   self:writeln( "_moduleObj.__macroName2InfoMap = __macroName2InfoMap" )
   for __index, macroDeclNode in ipairs( node:get_nodeManager():getDeclMacroNodeList(  ) ) do
      local declInfo = macroDeclNode:get_declInfo()
      if declInfo:get_pubFlag() then
         local macroInfo = _lune.unwrap( node:get_typeId2MacroInfo()[macroDeclNode:get_expType():get_typeId()])
         
         local macroTypeInfo = macroDeclNode:get_expType()
         pickupTypeId( macroTypeInfo, true )
         
         self:writeln( "do" )
         self:pushIndent(  )
         self:writeln( "local info = {}" )
         self:writeln( string.format( "__macroName2InfoMap[ %d ] = info", macroTypeInfo:get_typeId().id) )
         local pos = macroDeclNode:get_pos():get_orgPos()
         self:writeln( string.format( "info.pos = {%d,%d}", pos.lineNo, pos.column) )
         self:writeln( string.format( "info.name = %q", declInfo:get_name().txt) )
         self:writeRaw( "info.argList = {" )
         for index, argNode in ipairs( declInfo:get_argList() ) do
            if index ~= 1 then
               self:writeRaw( "," )
            end
            
            self:writeRaw( string.format( "{name=%q,typeId=%d}", argNode:get_name().txt, argNode:get_expType():get_typeId().id) )
         end
         
         self:writeln( "}" )
         self:writeRaw( "info.symList = {" )
         local firstFlag = true
         do
            local __sorted = {}
            local __map = macroInfo:get_symbol2MacroValInfoMap()
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, name in ipairs( __sorted ) do
               local symInfo = __map[ name ]
               do
                  if firstFlag then
                     firstFlag = false
                  else
                   
                     self:writeRaw( "," )
                  end
                  
                  self:writeRaw( string.format( "{name=%q,typeId=%d}", name, symInfo.typeInfo:get_typeId().id) )
                  pickupTypeId( symInfo.typeInfo, true )
               end
            end
         end
         
         self:writeln( "}" )
         
         do
            local stmtBlock = declInfo:get_stmtBlock()
            if stmtBlock ~= nil then
               local memStream = Util.memStream.new()
               
               local workFilter = ConvFilter.new(declInfo:get_name().txt, memStream, Util.NullOStream.new(), ConvMode.Convert, false, Ast.headTypeInfo, self.processInfo, Ast.SymbolKind.Typ, self.builtinFunc, self.useLuneRuntime, self.targetLuaVer, self.enableTest, self.useIpairs, self.option)
               
               workFilter.macroDepth = workFilter.macroDepth + 1
               workFilter:processBlock( stmtBlock, Opt.new(node) )
               workFilter.macroDepth = workFilter.macroDepth - 1
               
               memStream:close(  )
               self:writeln( string.format( 'info.stmtBlock = %s', Parser.quoteStr( memStream:get_txt() )) )
            end
         end
         
         
         self:writeln( 'info.tokenList = {' )
         local prevLineNo = -1
         for index, token in ipairs( declInfo:get_tokenList() ) do
            if index > 1 then
               self:writeRaw( "," )
            end
            
            if prevLineNo ~= -1 and prevLineNo ~= token.pos.lineNo then
               self:writeRaw( string.format( "{%d,%s},", Parser.TokenKind.Dlmt, Parser.quoteStr( "\n" )) )
            end
            
            prevLineNo = token.pos.lineNo
            self:writeRaw( string.format( "{%d,%s}", token.kind, Parser.quoteStr( token.txt )) )
         end
         
         self:writeln( '}' )
         
         self:popIndent(  )
         self:writeln( "end" )
      end
      
   end
   
   
   self:writeln( "local __varName2InfoMap = {}" )
   self:writeln( "_moduleObj.__varName2InfoMap = __varName2InfoMap" )
   
   do
      local __sorted = {}
      local __map = self.pubVarName2InfoMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, varName in ipairs( __sorted ) do
         local varInfo = __map[ varName ]
         do
            self:writeln( string.format( "__varName2InfoMap.%s = {", varName ) )
            self:writeln( string.format( "  name='%s', accessMode = %d, typeId = %s, mutable = %s }", varName, varInfo.accessMode, serializeInfo:serializeId( varInfo.typeInfo:get_typeId() ), true) )
            pickupTypeId( varInfo.typeInfo, true )
         end
      end
   end
   
   
   do
      local __sorted = {}
      local __map = self.pubFuncName2InfoMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local funcInfo = __map[ __key ]
         do
            pickupTypeId( funcInfo.typeInfo, true )
         end
      end
   end
   
   
   for __index, aliasNode in ipairs( node:get_nodeManager():getAliasNodeList(  ) ) do
      pickupTypeId( aliasNode:get_expType(), false )
   end
   
   
   self:writeln( "local __typeInfoList = {}" )
   self:writeln( "_moduleObj.__typeInfoList = __typeInfoList" )
   
   local listIndex = 1
   
   local function outputDepend( typeInfo, moduleTypeInfo )
   
      local typeId = typeInfo:get_typeId()
      if self.processInfo:get_orgInfo() == typeId:get_processInfo() then
         return ExportIdKind.Normal
      end
      
      
      if typeId:isSwichingId(  ) then
         return ExportIdKind.Discarded
      end
      
      if Ast.isExtId( typeInfo ) then
         
         return ExportIdKind.Normal
      end
      
      
      return ExportIdKind.Discarded
   end
   
   local wroteTypeIdSet = {}
   
   local function outputTypeInfo( typeInfo )
   
      if Ast.isBuiltin( typeInfo:get_typeId().id ) then
         return 
      end
      
      if not Ast.TypeInfo.hasParent( typeInfo ) then
         return 
      end
      
      
      do
         local _switchExp = typeInfo:get_kind()
         if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
            do
               local _switchExp = typeInfo:get_accessMode()
               if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Pro or _switchExp == Ast.AccessMode.Global then
               else 
                  
                     Util.errorLog( string.format( "skip: %s %s", typeInfo:get_accessMode(), self:getFullName( typeInfo )) )
                     return 
               end
            end
            
         end
      end
      
      local typeId = typeInfo:get_typeId(  )
      if _lune._Set_has(wroteTypeIdSet, typeId ) then
         return 
      end
      
      wroteTypeIdSet[typeId]= true
      
      self:writeRaw( string.format( "__typeInfoList[%d] = ", listIndex) )
      listIndex = listIndex + 1
      
      local validChildren = validChildrenSet[typeInfo]
      if  nil == validChildren then
         local _validChildren = validChildren
      
         validChildren = typeId2TypeInfo
      end
      
      typeInfo:serialize( self, Ast.SerializeInfo.new(importProcessInfo2Index, validChildren) )
   end
   
   for typeId, typeInfo in pairs( self.pubEnumId2EnumTypeInfo ) do
      typeId2TypeInfo[typeId] = typeInfo
   end
   
   for typeId, typeInfo in pairs( self.pubAlgeId2AlgeTypeInfo ) do
      typeId2TypeInfo[typeId] = typeInfo
      do
         local __sorted = {}
         local __map = typeInfo:get_valInfoMap()
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, __key in ipairs( __sorted ) do
            local valInfo = __map[ __key ]
            do
               for __index, valType in ipairs( valInfo:get_typeList() ) do
                  pickupTypeId( valType, true )
               end
               
            end
         end
      end
      
   end
   
   for __index, workNode in ipairs( node:get_nodeManager():getAliasNodeList(  ) ) do
      if Ast.isPubToExternal( workNode:get_expType():get_accessMode() ) then
         local aliasType = _lune.unwrap( _lune.__Cast( workNode:get_expType(), 3, Ast.AliasTypeInfo ))
         local aliasSrcType = aliasType:get_aliasSrc()
         typeId2TypeInfo[aliasType:get_typeId()] = aliasType
         typeId2TypeInfo[aliasSrcType:get_typeId()] = aliasSrcType
      end
      
   end
   
   
   self:writeln( "local __dependIdMap = {}" )
   self:writeln( "_moduleObj.__dependIdMap = __dependIdMap" )
   local exportNeedModuleTypeInfo = {}
   
   do
      local module2TypeList = {}
      for __index, typeInfo in pairs( typeId2TypeInfo ) do
         local modIndex = _lune.unwrap( importProcessInfo2Index[typeInfo:get_typeId():get_processInfo()])
         local map = module2TypeList[modIndex]
         if  nil == map then
            local _map = map
         
            map = {}
            module2TypeList[modIndex] = map
         end
         
         map[typeInfo:get_typeId().id] = typeInfo
      end
      
      do
         local __sorted = {}
         local __map = module2TypeList
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, __key in ipairs( __sorted ) do
            local map = __map[ __key ]
            do
               do
                  local __sorted = {}
                  local __map = map
                  for __key in pairs( __map ) do
                     table.insert( __sorted, __key )
                  end
                  table.sort( __sorted )
                  for __index, __key in ipairs( __sorted ) do
                     local typeInfo = __map[ __key ]
                     do
                        local moduleTypeInfo = typeInfo:getModule(  )
                        exportNeedModuleTypeInfo[moduleTypeInfo]= true
                        local ret = outputDepend( typeInfo, moduleTypeInfo )
                        
                        if ret == ExportIdKind.Normal then
                           outputTypeInfo( typeInfo )
                        end
                        
                     end
                  end
               end
               
            end
         end
      end
      
   end
   
   
   for moduleTypeInfo, __val in pairs( node:get_useModuleMacroSet() ) do
      exportNeedModuleTypeInfo[moduleTypeInfo]= true
   end
   
   
   self:writeln( "local __dependModuleMap = {}" )
   self:writeln( "_moduleObj.__dependModuleMap = __dependModuleMap" )
   do
      local __sorted = {}
      local __map = importNameMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, name in ipairs( __sorted ) do
         local exportInfo = __map[ name ]
         do
            local moduleTypeInfo = exportInfo:get_moduleTypeInfo()
            self:writeln( string.format( "__dependModuleMap[ '%s' ] = { typeId = %d, use = %s, buildId = %q }", name, _lune.unwrap( importModuleType2Index[moduleTypeInfo]), _lune._Set_has(exportNeedModuleTypeInfo, moduleTypeInfo ), (_lune.unwrap( node:get_importModule2moduleInfo()[moduleTypeInfo]) ):get_moduleId():get_idStr()) )
         end
      end
   end
   
   
   self:writeRaw( "_moduleObj.__subModuleMap = {" )
   do
      local firstFlag = true
      for __index, subfileNode in ipairs( node:get_nodeManager():getSubfileNodeList(  ) ) do
         do
            local usePath = subfileNode:get_usePath()
            if usePath ~= nil then
               if firstFlag then
                  firstFlag = false
               else
                
                  self:writeRaw( "," )
               end
               
               self:writeRaw( string.format( "%q", usePath) )
            end
         end
         
      end
      
   end
   
   self:writeln( "}" )
   
   self:writeRaw( "_moduleObj.__moduleHierarchy = {" )
   local workType = node:get_moduleTypeInfo()
   while Ast.TypeInfo.hasParent( workType ) do
      self:writeRaw( string.format( "%d,", workType:get_typeId().id) )
      workType = workType:get_parentInfo()
   end
   
   self:writeln( "}" )
   
   local moduleTypeInfo = self.moduleTypeInfo
   local moduleSymbolKind = Ast.SymbolKind.Typ
   do
      local _exp = node:get_provideNode()
      if _exp ~= nil then
         moduleTypeInfo = _exp:get_symbol():get_typeInfo()
         moduleSymbolKind = _exp:get_symbol():get_kind()
      end
   end
   
   self:writeln( string.format( "_moduleObj.__moduleTypeId = %d", moduleTypeInfo:get_typeId().id) )
   self:writeln( string.format( "_moduleObj.__moduleSymbolKind = %d", moduleSymbolKind) )
   self:writeln( string.format( "_moduleObj.__moduleMutable = %s", Ast.TypeInfo.isMut( moduleTypeInfo )) )
   
   self:writeln( "----- meta -----" )
   
   self:writeln( "return _moduleObj" )
   
   self.outMetaFlag = false
end


function ConvFilter:outputMainDirect(  )

   local code = string.format( [==[
do
   local loaded, mess = _lune.%s( [=[
if _lune and _lune._shebang then
  return nil
else
  return arg
end
]=] )
   if loaded ~= nil then
      local args = loaded(  )
      do
         local obj = (args )
         if obj ~= nil then
            local work = obj
            local argList = {""}
            do
               local _exp = work[0]
               if _exp ~= nil then
                  argList[1] = _exp
               end
            end
            for key, val in pairs( work ) do
               if key > 0 then
                  table.insert( argList, val )
               end
            end
            __main( argList )
         else
            -- print( "via lnsc" )
         end
      end
   else
      error( mess )
   end
end
]==], self.targetLuaVer:get_loadStrFuncName())
   self:writeln( code )
end


function ConvFilter:processRoot( node, opt )

   self:writeln( string.format( "--%s", self.streamName) )
   
   self.needModuleObj = node:get_provideNode() == nil
   
   if self.needModuleObj then
      self:writeln( "local _moduleObj = {}" )
   end
   
   if self.enableTest then
      self:writeln( "_moduleObj.__enableTest = true" )
   end
   
   
   self:writeln( string.format( "local __mod__ = '%s'", node:get_moduleTypeInfo():getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager() )) )
   
   local luneSymbol = string.format( "_lune%d", Ver.luaModVersion)
   
   do
      local runtime = self.useLuneRuntime
      if runtime ~= nil then
         self:writeln( string.format( 'local _lune = require( "%s" )', runtime) )
      else
         self:writeln( "local _lune = {}" )
         self:writeln( string.format( [==[
if %s then
   _lune = %s
end]==], luneSymbol, luneSymbol) )
         
         if node:get_luneHelperInfo().useAlge then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.Alge ) )
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.AlgeMapping ) )
         end
         
         if node:get_luneHelperInfo().useSet then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.SetOp ) )
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.SetMapping ) )
         end
         
         if node:get_luneHelperInfo().useUnpack and not self.targetLuaVer:get_hasTableUnpack() then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.Unpack ) )
         end
         
         if node:get_luneHelperInfo().useLoad or self.option:get_mainModule() == self:getCanonicalName( self.moduleTypeInfo, false ):gsub( "@", "" ) then
            self:writeln( self.targetLuaVer:getLoadCode(  ) )
         end
         
         if node:get_luneHelperInfo().useNilAccess then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.NilAcc ) )
         end
         
         if node:get_luneHelperInfo().useUnwrapExp then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.Unwrap ) )
         end
         
         if node:get_luneHelperInfo().hasMappingClassDef then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.Mapping ) )
         end
         
         if #node:get_nodeManager():getImportNodeList(  ) ~= 0 then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.LoadModule ) )
         end
         
         if #node:get_nodeManager():getExpCastNodeList(  ) ~= 0 then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.InstanceOf ) )
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.Cast ) )
         end
         
         if node:get_luneHelperInfo().useLazyLoad then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.LazyLoad ) )
         end
         
         if node:get_luneHelperInfo().useLazyRequire then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.LazyRequire ) )
         end
         
         if node:get_luneHelperInfo().useRun then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.Run ) )
         end
         
         if node:get_luneHelperInfo().useStrReplace then
            self:writeln( LuaMod.getCode( LuaMod.CodeKind.StrReplace ) )
         end
         
      end
   end
   
   self:writeln( string.format( [==[
if not %s then
   %s = _lune
end]==], luneSymbol, luneSymbol) )
   
   for __index, importNode in ipairs( node:get_nodeManager():getImportNodeList(  ) ) do
      local info = importNode:get_info()
      if info:get_lazy() ~= Nodes.LazyLoad.Off then
         self.moduleType2SymbolMap[info:get_moduleTypeInfo()] = info:get_symbolInfo()
      end
      
   end
   
   
   if _lune.nilacc( node:get_moduleScope():getSymbolInfoChild( "_" ), 'get_posForLatestMod', 'callmtd' ) then
      self:writeln( "local _" )
   end
   
   
   local children = node:get_children()
   
   for __index, child in ipairs( children ) do
      filter( child, self, node )
      self:writeln( "" )
   end
   
   
   if self.option:get_mainModule() == self:getCanonicalName( self.moduleTypeInfo, false ):gsub( "@", "" ) then
      self:outputMainDirect(  )
   end
   
   
   do
      local _exp = node:get_provideNode()
      if _exp ~= nil then
         self:writeRaw( "return " )
         self:writeRaw( _exp:get_symbol():get_name() )
         
         self:writeln( "" )
      else
         self:writeln( "return _moduleObj" )
      end
   end
   
   
   self:outputMeta( node )
end



function ConvFilter:processSubfile( node, opt )

end


function ConvFilter:processRequest( node, opt )

   filter( node:get_exp(), self, node )
end



function ConvFilter:processAsyncLock( node, opt )

   filter( node:get_block(), self, node )
end


function ConvFilter:processBlockSub( node, opt )

   local word = ""
   do
      local _switchExp = node:get_blockKind(  )
      if _switchExp == Nodes.BlockKind.If or _switchExp == Nodes.BlockKind.Elseif then
         word = "then"
      elseif _switchExp == Nodes.BlockKind.Else then
         word = ""
      elseif _switchExp == Nodes.BlockKind.While then
         word = "do"
      elseif _switchExp == Nodes.BlockKind.Repeat then
         word = ""
      elseif _switchExp == Nodes.BlockKind.For then
         word = "do"
      elseif _switchExp == Nodes.BlockKind.Apply then
         word = "do"
      elseif _switchExp == Nodes.BlockKind.Foreach then
         word = "do"
      elseif _switchExp == Nodes.BlockKind.Macro then
         word = ""
      elseif _switchExp == Nodes.BlockKind.Func then
         word = ""
      elseif _switchExp == Nodes.BlockKind.Default then
         word = ""
      elseif _switchExp == Nodes.BlockKind.Block or _switchExp == Nodes.BlockKind.AsyncLock then
         word = "do"
      elseif _switchExp == Nodes.BlockKind.LetUnwrap then
         word = ""
      elseif _switchExp == Nodes.BlockKind.LetUnwrapThenDo then
         word = ""
      elseif _switchExp == Nodes.BlockKind.IfUnwrap then
         word = ""
      end
   end
   
   self:writeln( word )
   self:pushIndent(  )
   
   if _lune.nilacc( node:get_scope():getSymbolInfoChild( "_" ), 'get_posForLatestMod', 'callmtd' ) then
      self:writeln( "local _" )
   end
   
   
   local stmtList = node:get_stmtList(  )
   for __index, statement in ipairs( stmtList ) do
      filter( statement, self, node )
      self:writeln( "" )
   end
   
   
   self:popIndent(  )
   do
      local _switchExp = node:get_blockKind(  )
      if _switchExp == Nodes.BlockKind.Block or _switchExp == Nodes.BlockKind.AsyncLock then
         self:writeln( "end" )
      end
   end
   
end



function ConvFilter:processLoadRuntime(  )

   do
      local _exp = self.useLuneRuntime
      if _exp ~= nil then
         self:writeln( string.format( 'local _lune = require( "%s" )', _exp) )
      else
         self:writeln( string.format( 'local _lune = require( "%s" )', LnsOpt.getRuntimeModule(  )) )
      end
   end
   
end


function ConvFilter:processScope( node, opt )

   if node:get_scopeKind() == Nodes.ScopeKind.Root then
      self:processLoadRuntime(  )
   end
   
   filter( node:get_block(), self, node )
end


function ConvFilter:processStmtExp( node, opt )

   filter( node:get_exp(  ), self, node )
end



function ConvFilter:processDeclEnum( node, opt )

   local access = node:get_accessMode() == Ast.AccessMode.Global and "" or "local "
   local enumFullName = node:get_name().txt
   local typeInfo = _lune.unwrap( _lune.__Cast( node:get_expType():get_aliasSrc(), 3, Ast.EnumTypeInfo ))
   local parentInfo = typeInfo:get_parentInfo()
   local isTopNS = true
   if Ast.TypeInfo.hasParent( typeInfo ) and parentInfo:get_kind() == Ast.TypeInfoKind.Class then
      enumFullName = string.format( "%s.%s", self:getFullName( parentInfo ), enumFullName)
      access = ""
      isTopNS = false
   end
   
   
   self:writeln( string.format( "%s%s = {}", access, enumFullName) )
   if isTopNS and node:get_accessMode() == Ast.AccessMode.Pub then
      if self.needModuleObj then
         self:writeln( string.format( "_moduleObj.%s = %s", enumFullName, enumFullName) )
      end
      
   end
   
   
   if typeInfo:get_accessMode() == Ast.AccessMode.Pub then
      self.pubEnumId2EnumTypeInfo[typeInfo:get_typeId()] = typeInfo
   end
   
   
   self:writeln( string.format( "%s._val2NameMap = {}", enumFullName) )
   self:writeln( string.format( [==[function %s:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "%s.%%s", name )
   end
   return string.format( "illegal val -- %%s", val )
end
function %s._from( val )
   if %s._val2NameMap[ val ] then
      return val
   end
   return nil
end
    ]==], enumFullName, enumFullName, enumFullName, enumFullName) )
   self:writeln( string.format( [==[
%s.__allList = {}
function %s.get__allList()
   return %s.__allList
end
]==], enumFullName, enumFullName, enumFullName) )
   
   for index, valName in ipairs( node:get_valueNameList() ) do
      local valInfo = _lune.unwrap( typeInfo:getEnumValInfo( valName.txt ))
      local valTxt = string.format( "%s", Ast.getEnumLiteralVal( valInfo:get_val() ))
      if typeInfo:get_valTypeInfo():equals( self.processInfo, Ast.builtinTypeString ) then
         valTxt = string.format( "'%s'", Ast.getEnumLiteralVal( valInfo:get_val() ))
      end
      
      self:writeln( string.format( "%s.%s = %s", enumFullName, valName.txt, valTxt) )
      
      self:writeln( string.format( "%s._val2NameMap[%s] = '%s'", enumFullName, valTxt, valName.txt) )
      self:writeln( string.format( "%s.__allList[%d] = %s.%s", enumFullName, index, enumFullName, valName.txt) )
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

function ConvFilter:getMapInfo( typeInfo )

   local nonnilableType = typeInfo:get_srcTypeInfo()
   if typeInfo:get_nilable() then
      nonnilableType = typeInfo:get_nonnilableType()
   end
   
   
   local child = "{}"
   local funcTxt = ""
   do
      local _switchExp = nonnilableType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Stem then
         funcTxt = '_lune._toStem'
      elseif _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
         if not nonnilableType:equals( self.processInfo, Ast.builtinTypeString ) then
            if Ast.NormalTypeInfo.isAvailableMapping( self.processInfo, nonnilableType, {} ) then
               funcTxt = string.format( '%s._fromMap', self:getFullName( nonnilableType ))
               if isGenericType( nonnilableType ) then
                  local memStream = Util.memStream.new()
                  self:outputAlter2MapFunc( memStream, nonnilableType:createAlt2typeMap( false ) )
                  child = memStream:get_txt()
               end
               
            else
             
               funcTxt = "nil"
            end
            
         else
          
            funcTxt = '_lune._toStr'
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.Alge then
         funcTxt = string.format( '%s._from', self:getFullName( nonnilableType ))
      elseif _switchExp == Ast.TypeInfoKind.Prim then
         do
            local _switchExp = nonnilableType
            if _switchExp == Ast.builtinTypeInt then
               funcTxt = '_lune._toInt'
            elseif _switchExp == Ast.builtinTypeReal then
               funcTxt = '_lune._toReal'
            elseif _switchExp == Ast.builtinTypeBool then
               funcTxt = '_lune._toBool'
            else 
               
                  Util.err( string.format( "unknown type -- %s", nonnilableType:getTxt(  )) )
            end
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Map then
         funcTxt = '_lune._toMap'
         local itemList = nonnilableType:get_itemTypeInfoList()
         local keyFuncTxt, keyNilable, keyChild = self:getMapInfo( itemList[1] )
         local valFuncTxt, valNilable, valChild = self:getMapInfo( itemList[2] )
         
         child = string.format( "{ { func = %s, nilable = %s, child = %s }, \n", keyFuncTxt, keyNilable, keyChild) .. string.format( "{ func = %s, nilable = %s, child = %s } }", valFuncTxt, valNilable, valChild)
      elseif _switchExp == Ast.TypeInfoKind.Set then
         funcTxt = '_lune._toSet'
         local itemList = nonnilableType:get_itemTypeInfoList()
         local valFuncTxt, valNilable, valChild = self:getMapInfo( itemList[1] )
         
         child = string.format( "{ func = %s, nilable = %s, child = %s }", valFuncTxt, valNilable, valChild)
      elseif _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         funcTxt = '_lune._toList'
         local itemList = nonnilableType:get_itemTypeInfoList()
         local valFuncTxt, valNilable, valChild = self:getMapInfo( itemList[1] )
         
         child = string.format( "{ { func = %s, nilable = %s, child = %s } }", valFuncTxt, valNilable, valChild)
      elseif _switchExp == Ast.TypeInfoKind.Alternate then
         local prefix = string.format( "obj.__alt2mapFunc.%s", nonnilableType:get_rawTxt())
         funcTxt = string.format( "%s.func", prefix)
         child = string.format( "%s.child", prefix)
      end
   end
   
   return funcTxt, typeInfo:get_nilable(), child
end


function ConvFilter:processDeclAlge( node, opt )

   local access = node:get_accessMode() == Ast.AccessMode.Global and "" or "local "
   local algeFullName = node:get_algeType():get_rawTxt()
   local typeInfo = _lune.unwrap( _lune.__Cast( node:get_expType(), 3, Ast.AlgeTypeInfo ))
   local parentInfo = typeInfo:get_parentInfo()
   local isTopNS = true
   if Ast.TypeInfo.hasParent( typeInfo ) and parentInfo:get_kind() == Ast.TypeInfoKind.Class then
      algeFullName = string.format( "%s.%s", self:getFullName( parentInfo ), algeFullName)
      access = ""
      isTopNS = false
   end
   
   
   self:writeln( string.format( "%s%s = {}", access, algeFullName) )
   self:writeln( string.format( "%s._name2Val = {}", algeFullName) )
   if isTopNS and node:get_accessMode() == Ast.AccessMode.Pub then
      if self.needModuleObj then
         self:writeln( string.format( "_moduleObj.%s = %s", algeFullName, algeFullName) )
      end
      
   end
   
   
   if typeInfo:get_accessMode() == Ast.AccessMode.Pub then
      self.pubAlgeId2AlgeTypeInfo[typeInfo:get_typeId()] = typeInfo
   end
   
   
   self:writeln( string.format( [==[function %s:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "%s.%%s", name )
   end
   return string.format( "illegal val -- %%s", val )
end
]==], algeFullName, algeFullName) )
   self:writeln( string.format( [==[
function %s._from( val )
   return _lune._AlgeFrom( %s, val )
end
]==], algeFullName, algeFullName) )
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
            self:writeRaw( string.format( '%s.%s = { "%s"', algeFullName, valInfo:get_name(), valInfo:get_name()) )
            
            if #valInfo:get_typeList() > 0 then
               self:writeRaw( ", {" )
               for index, paramType in ipairs( valInfo:get_typeList() ) do
                  if index > 1 then
                     self:writeRaw( "," )
                  end
                  
                  if Ast.NormalTypeInfo.isAvailableMapping( self.processInfo, node:get_algeType(), {} ) then
                     local funcTxt, nilable, child = self:getMapInfo( paramType )
                     self:writeRaw( string.format( "{ func=%s, nilable=%s, child=%s }", funcTxt, nilable, child) )
                  else
                   
                     self:writeRaw( "{}" )
                  end
                  
               end
               
               self:writeRaw( "}" )
            end
            
            self:writeln( "}" )
            self:writeln( string.format( '%s._name2Val["%s"] = %s.%s', algeFullName, valInfo:get_name(), algeFullName, valInfo:get_name()) )
         end
      end
   end
   
end


function ConvFilter:processNewAlgeVal( node, opt )

   local valInfo = node:get_valInfo()
   self:writeRaw( string.format( '_lune.newAlge( %s.%s', self:getFullName( node:get_algeTypeInfo() ), valInfo:get_name()) )
   if #valInfo:get_typeList() > 0 then
      self:writeRaw( ", {" )
      for index, exp in ipairs( node:get_paramList() ) do
         if index > 1 then
            self:writeRaw( "," )
         end
         
         filter( exp, self, node )
      end
      
      self:writeRaw( "}" )
   end
   
   self:writeRaw( ")" )
end


function ConvFilter:getDestrClass( classTypeInfo )

   local typeInfo = classTypeInfo
   while typeInfo ~= Ast.headTypeInfo do
      local scope = _lune.unwrap( typeInfo:get_scope())
      if scope:getTypeInfoChild( "__free" ) then
         return typeInfo
      end
      
      typeInfo = typeInfo:get_baseTypeInfo()
   end
   
   return nil
end


function ConvFilter:outputAlter2MapFunc( stream, alt2Map )

   stream:write( "{" )
   
   for altType, assinType in pairs( alt2Map ) do
      if altType:get_kind() == Ast.TypeInfoKind.Alternate then
         if assinType:get_kind() == Ast.TypeInfoKind.Alternate then
            stream:write( string.format( "%s = obj.__alt2mapFunc.%s,", altType:get_rawTxt(), assinType:get_rawTxt()) )
         else
          
            local funcTxt, nilable, child = self:getMapInfo( assinType )
            stream:write( string.format( "%s = { func=%s, nilable=%s, child=%s },", altType:get_rawTxt(), funcTxt, nilable, child) )
         end
         
      end
      
   end
   
   
   stream:write( "}" )
end


function ConvFilter:processProtoClass( node, opt )

   self:writeRaw( string.format( "local %s = {}", node:get_name().txt) )
end


function ConvFilter:processDeclClass( node, opt )

   local nodeInfo = node
   local classNameToken = nodeInfo:get_name(  )
   local className = classNameToken.txt
   local classTypeInfo = node:get_expType(  )
   local classTypeId = classTypeInfo:get_typeId()
   
   if nodeInfo:get_accessMode(  ) == Ast.AccessMode.Pub then
      self.classId2TypeInfo[classTypeId] = classTypeInfo
   end
   
   self.classId2MemberList[classTypeId] = nodeInfo:get_memberList(  )
   
   do
      local _exp = node:get_moduleName()
      if _exp ~= nil then
         self:writeRaw( string.format( "local %s = ", className) )
         if node:get_lazyLoad() == Nodes.LazyLoad.Off then
            self:writeRaw( "require" )
         else
          
            self:writeRaw( "_lune._lazyRequire" )
         end
         
         self:writeRaw( string.format( "( %s )", _exp.txt) )
         do
            local _switchExp = node:get_accessMode()
            if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Pro then
               if self.needModuleObj then
                  self:writeln( "" )
                  self:writeRaw( string.format( "_moduleObj.%s = %s", className, className) )
               end
               
            end
         end
         
         return 
      end
   end
   
   if not node:get_hasPrototype() then
      self:writeln( string.format( "local %s = {}", className ) )
   end
   
   
   local ifTxt = ""
   if #classTypeInfo:get_interfaceList() > 0 then
      for __index, ifType in ipairs( classTypeInfo:get_interfaceList() ) do
         ifTxt = ifTxt .. self:getFullName( ifType ) .. ","
      end
      
      ifTxt = string.format( "ifList = {%s}", ifTxt)
   end
   
   
   local baseInfo = classTypeInfo:get_baseTypeInfo(  )
   local baseTxt = ""
   if baseInfo:get_typeId(  ) ~= Ast.rootTypeIdInfo then
      baseTxt = string.format( "__index = %s", self:getFullName( baseInfo ))
   end
   
   
   if #ifTxt > 0 or #baseTxt > 0 then
      local metaTxt = baseTxt
      if #baseTxt > 0 and #ifTxt > 0 then
         metaTxt = string.format( "%s,%s", baseTxt, ifTxt)
      elseif #ifTxt > 0 then
         metaTxt = ifTxt
      end
      
      self:writeln( string.format( "setmetatable( %s, { %s } )", className, metaTxt) )
   end
   
   
   if nodeInfo:get_accessMode(  ) == Ast.AccessMode.Pub then
      if self.needModuleObj then
         self:writeln( string.format( "_moduleObj.%s = %s", className, className ) )
      end
      
   end
   
   
   for __index, declNode in ipairs( node:get_declStmtList() ) do
      filter( declNode, self, node )
   end
   
   
   local hasConstrFlag = false
   local memberList = {}
   local fieldList = nodeInfo:get_fieldList(  )
   local outerMethodSet = nodeInfo:get_outerMethodSet(  )
   local methodNameSet = {}
   
   if classTypeInfo:get_kind() ~= Ast.TypeInfoKind.IF then
      for __index, field in ipairs( fieldList ) do
         local ignoreFlag = false
         if field:get_kind() == Nodes.NodeKind.get_DeclConstr() then
            hasConstrFlag = true
            methodNameSet["__init"]= true
         end
         
         if field:get_kind() == Nodes.NodeKind.get_DeclDestr() then
            methodNameSet["__free"]= true
         end
         
         do
            local declMemberNode = _lune.__Cast( field, 3, Nodes.DeclMemberNode )
            if declMemberNode ~= nil then
               if not declMemberNode:get_staticFlag() then
                  table.insert( memberList, declMemberNode )
               end
               
            end
         end
         
         do
            local methodNode = _lune.__Cast( field, 3, Nodes.DeclMethodNode )
            if methodNode ~= nil then
               local declInfo = methodNode:get_declInfo(  )
               local methodNameToken = _lune.unwrap( declInfo:get_name(  ))
               if _lune._Set_has(outerMethodSet, methodNameToken.txt ) then
                  ignoreFlag = true
               end
               
               methodNameSet[methodNameToken.txt]= true
            end
         end
         
         
         if (not ignoreFlag ) then
            filter( field, self, node )
         end
         
      end
      
   end
   
   
   local destTxt = ""
   do
      local _exp = self:getDestrClass( node:get_expType(  ) )
      if _exp ~= nil then
         destTxt = string.format( ", __gc = %s.__free", _exp:getTxt(  ))
      end
   end
   
   
   self:writeln( string.format( [==[
function %s.setmeta( obj )
  setmetatable( obj, { __index = %s %s } )
end]==], className, className, destTxt) )
   
   if not hasConstrFlag then
      
      methodNameSet["__init"]= true
      
      local oldFlag = node:get_hasOldCtor()
      
      local superArgTxt = ""
      local thisArgTxt = ""
      
      if not oldFlag and baseInfo ~= Ast.headTypeInfo then
         do
            local superInit = (_lune.unwrap( baseInfo:get_scope()) ):getSymbolInfoChild( "__init" )
            if superInit ~= nil then
               for index, _1 in ipairs( superInit:get_typeInfo():get_argTypeInfoList() ) do
                  if #superArgTxt > 0 then
                     superArgTxt = superArgTxt .. ", "
                  end
                  
                  superArgTxt = string.format( "%s__superarg%d", superArgTxt, index)
               end
               
            end
         end
         
      end
      
      
      for __index, member in ipairs( memberList ) do
         if #thisArgTxt > 0 then
            thisArgTxt = thisArgTxt .. ", "
         end
         
         thisArgTxt = thisArgTxt .. member:get_name().txt
      end
      
      
      local argTxt = superArgTxt
      if thisArgTxt ~= "" then
         if #argTxt > 0 then
            argTxt = argTxt .. ","
         end
         
         argTxt = argTxt .. thisArgTxt
      end
      
      
      self:writeln( string.format( [==[
function %s.new( %s )
   local obj = {}
   %s.setmeta( obj )
   if obj.__init then
      obj:__init( %s )
   end
   return obj
end
function %s:__init( %s )
]==], className, argTxt, className, argTxt, className, argTxt) )
      self:pushIndent(  )
      
      if baseInfo ~= Ast.headTypeInfo then
         
         if (_lune.unwrap( baseInfo:get_scope()) ):getSymbolInfoChild( "__init" ) then
            self:writeRaw( string.format( "%s.__init( self", self:getFullName( baseInfo )) )
            if #superArgTxt > 0 then
               self:writeln( string.format( ", %s )", superArgTxt) )
            else
             
               self:writeln( ")" )
            end
            
         end
         
      end
      
      
      for __index, member in ipairs( memberList ) do
         local memberName = member:get_name().txt
         self:writeln( string.format( "self.%s = %s", memberName, memberName ) )
      end
      
      self:popIndent(  )
      self:writeln( 'end' )
   end
   
   
   for __index, memberNode in ipairs( nodeInfo:get_memberList() ) do
      local memberNameToken = memberNode:get_name(  )
      local memberName = memberNameToken.txt
      local getterName = "get_" .. memberName
      
      local autoFlag = not _lune._Set_has(methodNameSet, getterName )
      local prefix
      
      local delimit
      
      if memberNode:get_staticFlag() then
         prefix = className
         delimit = "."
      else
       
         prefix = "self"
         delimit = ":"
      end
      
      if memberNode:get_getterMode(  ) ~= Ast.AccessMode.None and autoFlag then
         self:writeln( string.format( [==[
function %s%s%s()
   return %s.%s
end]==], className, delimit, getterName, prefix, memberName) )
         methodNameSet[getterName]= true
      end
      
      local setterName = "set_" .. memberName
      
      autoFlag = not _lune._Set_has(methodNameSet, setterName )
      if memberNode:get_setterMode(  ) ~= Ast.AccessMode.None and autoFlag then
         self:writeln( string.format( [==[
function %s%s%s( %s )
   %s.%s = %s
end]==], className, delimit, setterName, memberName, prefix, memberName, memberName) )
         methodNameSet[setterName]= true
      end
      
   end
   
   
   for __index, advertiseInfo in ipairs( node:get_advertiseList() ) do
      local memberName = advertiseInfo:get_member():get_name().txt
      local memberType = advertiseInfo:get_member():get_expType()
      for __index, mtdName in ipairs( Ast.getAllMethodName( memberType, Ast.MethodKind.Object ):get_list() ) do
         local mbrScope = _lune.unwrap( memberType:get_scope())
         local child = _lune.unwrap( mbrScope:getTypeInfoField( mtdName, true, mbrScope, Ast.ScopeAccess.Normal ))
         if child:get_accessMode() ~= Ast.AccessMode.Pri then
            local childName = advertiseInfo:get_prefix() .. child:getTxt(  )
            if not _lune._Set_has(methodNameSet, childName ) then
               self:writeln( string.format( [==[
function %s:%s( ... )
   return self.%s:%s( ... )
end
]==], className, childName, memberName, childName) )
            end
            
         end
         
      end
      
   end
   
   
   do
      local initBlock = _lune.nilacc( _lune.nilacc( nodeInfo:get_initBlock():get_func(), 'get_declInfo', 'callmtd' ), 'get_body', 'callmtd' )
      if initBlock ~= nil then
         if #initBlock:get_stmtList() > 0 then
            self:writeln( "do" )
            self:pushIndent(  )
            for __index, initStmt in ipairs( initBlock:get_stmtList() ) do
               filter( initStmt, self, node )
               self:writeln( "" )
            end
            
            self:popIndent(  )
            self:writeln( "end" )
         end
         
      end
   end
   
   
   if classTypeInfo:isInheritFrom( self.processInfo, Ast.builtinTypeMapping, nil ) then
      local declArgTxt = "val"
      local argTxt = "{}, val"
      if isGenericType( classTypeInfo ) then
         declArgTxt = "val, __alt2mapFunc"
         argTxt = "{ __alt2mapFunc = __alt2mapFunc }, val"
      end
      
      
      self:writeln( string.format( [==[
function %s:_toMap()
  return self
end
function %s._fromMap( %s )
  local obj, mes = %s._fromMapSub( %s )
  if obj then
     %s.setmeta( obj )
  end
  return obj, mes
end
function %s._fromStem( %s )
  return %s._fromMap( %s )
end
]==], className, className, declArgTxt, className, argTxt, className, className, declArgTxt, className, declArgTxt) )
      
      self:writeln( string.format( 'function %s._fromMapSub( obj, val )', className) )
      
      if classTypeInfo:get_baseTypeInfo() ~= Ast.headTypeInfo then
         self:writeln( string.format( [==[
   local result, mes = %s._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end
]==], self:getFullName( classTypeInfo:get_baseTypeInfo() )) )
      end
      
      
      self:writeln( '   local memInfo = {}' )
      
      for __index, memberNode in ipairs( node:get_memberList() ) do
         local funcTxt, nilable, child = self:getMapInfo( memberNode:get_expType() )
         self:writeln( string.format( '   table.insert( memInfo, { name = "%s", func = %s, nilable = %s, child = %s } )', memberNode:get_name().txt, funcTxt, nilable, child) )
      end
      
      
      self:writeln( [==[
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end]==] )
   end
   
end



function ConvFilter:processDeclMember( node, opt )

end



function ConvFilter:processExpMacroExp( node, opt )

   for __index, stmt in ipairs( node:get_stmtList() ) do
      filter( stmt, self, node )
      self:writeln( "" )
   end
   
end





function ConvFilter:outputDeclMacro( name, argNameList, callback )

   self:writeRaw( string.format( "local function %s(", name) )
   
   self:writeln( "__macroArgs )" )
   self:pushIndent(  )
   
   self:writeln( string.format( 'local _lune = require( "%s" )', LnsOpt.getRuntimeModule(  )) )
   
   self:writeln( "local __var = __macroArgs.__var" )
   for __index, argName in ipairs( argNameList ) do
      self:writeln( string.format( "local %s = __macroArgs.%s", argName, argName) )
   end
   
   
   self:writeln( "local macroVar = {}" )
   self:writeln( "macroVar.__names = {}" )
   
   self:writeln( [==[
local function __expStatList( list )
  local ret = ""
  for index, txt in ipairs( list ) do
    ret = string.format( "%s %s ", ret, txt )
  end
  return ret
end
]==] )
   
   self.macroDepth = self.macroDepth + 1
   
   callback(  )
   
   self.macroDepth = self.macroDepth - 1
   
   self:writeln( "" )
   
   self:writeln( "macroVar.__var = __var" )
   self:writeln( "return macroVar" )
   self:popIndent(  )
   self:writeln( "end" )
   self:writeln( string.format( "return %s", name) )
end


function ConvFilter:processExpMacroStatList( node, opt )

   self:writeRaw( "__expStatList(" )
   filter( node:get_exp(), self, node )
   self:writeRaw( ")" )
end


function ConvFilter:processDeclMacro( node, opt )

   if self.inMacro then
      local macroInfo = node:get_declInfo(  )
      local argNameList = {}
      for __index, arg in ipairs( macroInfo:get_argList() ) do
         table.insert( argNameList, arg:get_name().txt )
      end
      
      self:outputDeclMacro( macroInfo:get_name().txt, argNameList, function (  )
      
         do
            local stmtBlock = macroInfo:get_stmtBlock()
            if stmtBlock ~= nil then
               filter( stmtBlock, self, node )
            end
         end
         
      end )
   end
   
end



function ConvFilter:processExpMacroStat( node, opt )

   if #node:get_expStrList() == 0 then
      self:writeRaw( "''" )
   else
    
      for index, token in ipairs( node:get_expStrList() ) do
         if index ~= 1 then
            self:writeRaw( '..' )
         end
         
         
         filter( token, self, node )
      end
      
   end
   
end



function ConvFilter:processExpNew( node, opt )

   filter( node:get_symbol(  ), self, node )
   self:writeRaw( ".new(" )
   do
      local _exp = node:get_argList(  )
      if _exp ~= nil then
         filter( _exp, self, node )
      end
   end
   
   self:writeRaw( ")" )
end



function ConvFilter:process__func__symbol( has__func__Symbol, parentType, funcName )

   if has__func__Symbol then
      local nameSpace = self:getCanonicalName( parentType, false )
      if funcName == "" then
         funcName = "<anonymous>"
      end
      
      self:pushIndent(  )
      self:writeln( string.format( "local __func__ = '%s.%s'", nameSpace, funcName) )
      self:popIndent(  )
   end
   
end


function ConvFilter:processDeclConstr( node, opt )

   local declInfo = node:get_declInfo(  )
   local classTypeInfo = _lune.unwrap( declInfo:get_classTypeInfo())
   local className = self:getFullName( classTypeInfo )
   self:writeRaw( string.format( "function %s.new( ", className ) )
   
   local argTxt = ""
   
   self:writeRaw( argTxt )
   local argList = declInfo:get_argList(  )
   for __index, arg in ipairs( argList ) do
      if #argTxt > 0 then
         self:writeRaw( ", " )
         argTxt = argTxt .. ", "
      end
      
      filter( arg, self, node )
      
      do
         local _exp = _lune.__Cast( arg, 3, Nodes.DeclArgNode )
         if _exp ~= nil then
            argTxt = argTxt .. _exp:get_name().txt
         else
            local name = _lune.unwrap( node:get_declInfo(  ):get_name())
            Util.err( string.format( "not support ... in macro -- %s", name.txt) )
         end
      end
      
   end
   
   self:writeln( " )" )
   self:pushIndent(  )
   self:writeln( "local obj = {}" )
   self:writeln( string.format( "%s.setmeta( obj )", className) )
   self:writeln( string.format( "if obj.__init then obj:__init( %s ); end", argTxt ) )
   self:writeln( "return obj" )
   self:popIndent(  )
   self:writeln( "end" )
   
   self:writeRaw( string.format( "function %s:__init(%s) ", className, argTxt ) )
   do
      local _exp = declInfo:get_body()
      if _exp ~= nil then
         self:process__func__symbol( declInfo:get_has__func__Symbol(), node:get_expType():get_parentInfo(), "__init" )
         
         filter( _exp, self, node )
      end
   end
   
   self:writeln( "end" )
end



function ConvFilter:processDeclDestr( node, opt )

   self:writeln( string.format( "function %s.__free( self )", _lune.nilacc( node:get_declInfo():get_classTypeInfo(), 'getTxt', 'callmtd'  )) )
   
   self:process__func__symbol( node:get_declInfo():get_has__func__Symbol(), node:get_expType():get_parentInfo(), "__free" )
   
   filter( _lune.unwrap( node:get_declInfo():get_body()), self, node )
   
   local classTypeInfo = node:get_expType():get_parentInfo()
   do
      local _exp = self:getDestrClass( classTypeInfo:get_baseTypeInfo() )
      if _exp ~= nil then
         self:writeln( string.format( "%s.__free( self )", _exp:getTxt(  )) )
      end
   end
   
   
   self:writeln( "end" )
end


function ConvFilter:processExpCallSuperCtor( node, opt )

   local typeInfo = node:get_superType()
   
   self:writeRaw( string.format( "%s.%s( self", self:getFullName( typeInfo ), node:get_methodType():get_rawTxt()) )
   
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         self:writeRaw( "," )
         filter( _exp, self, node )
      end
   end
   
   self:writeln( ")" )
end



function ConvFilter:processExpCallSuper( node, opt )

   local typeInfo = node:get_superType()
   
   self:writeRaw( string.format( "%s.%s( self", self:getFullName( typeInfo ), node:get_methodType():get_rawTxt()) )
   
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         self:writeRaw( "," )
         filter( _exp, self, node )
      end
   end
   
   self:writeRaw( ")" )
end



function ConvFilter:processDeclMethod( node, opt )

   local declInfo = node:get_declInfo(  )
   local classTypeInfo = _lune.unwrap( declInfo:get_classTypeInfo())
   
   local delimit = ":"
   if declInfo:get_staticFlag(  ) then
      delimit = "."
   end
   
   local methodNodeToken = _lune.unwrap( declInfo:get_name(  ))
   local methodName = methodNodeToken.txt
   self:writeRaw( string.format( "function %s%s%s( ", classTypeInfo:get_rawTxt(), delimit, methodName) )
   
   local argList = declInfo:get_argList(  )
   for index, arg in ipairs( argList ) do
      if index > 1 then
         self:writeRaw( ", " )
      end
      
      filter( arg, self, node )
   end
   
   self:writeln( " )" )
   do
      local _exp = declInfo:get_body()
      if _exp ~= nil then
         self:process__func__symbol( declInfo:get_has__func__Symbol(), node:get_expType():get_parentInfo(), methodName )
         filter( _exp, self, node )
      end
   end
   
   self:writeln( "end" )
end



function ConvFilter:processUnwrapSet( node, opt )

   local dstExpList = node:get_dstExpList()
   filter( dstExpList, self, node )
   self:writeRaw( " = " )
   filter( node:get_srcExpList(), self, node )
   self:writeln( "" )
   
   self:writeRaw( "if " )
   for index, expNode in ipairs( dstExpList:get_expList() ) do
      if index > 1 then
         self:writeRaw( " or " )
      end
      
      self:writeRaw( "nil == " )
      filter( expNode, self, node )
   end
   
   self:writeln( " then" )
   self:pushIndent(  )
   
   for index, expNode in ipairs( dstExpList:get_expList() ) do
      self:writeRaw( string.format( "local _exp%d = ", index) )
      filter( expNode, self, node )
      self:writeln( "" )
   end
   
   
   if node:get_unwrapBlock() then
      filter( _lune.unwrap( node:get_unwrapBlock()), self, node )
   end
   
   self:popIndent(  )
   self:writeln( "end" )
end


function ConvFilter:processExpListSub( parent, expList, mRetExp )

   local mRetIndex = _lune.nilacc( mRetExp, 'get_index', 'callmtd' )
   
   for index, exp in ipairs( expList ) do
      if exp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
         break
      end
      
      
      do
         local castNode = _lune.__Cast( exp, 3, Nodes.ExpCastNode )
         if castNode ~= nil then
            if castNode:get_castKind() == Nodes.CastKind.Implicit then
               if castNode:get_exp():get_kind() == Nodes.NodeKind.get_ExpAccessMRet() then
                  break
               end
               
            end
            
         end
      end
      
      if index > 1 then
         self:writeRaw( ", " )
      end
      
      filter( exp, self, parent )
      if index == mRetIndex then
         break
      end
      
   end
   
end


function ConvFilter:processExpMRet( node, opt )

   filter( node:get_mRet(), self, node )
end


function ConvFilter:processIfUnwrap( node, opt )

   self:writeln( "do" )
   self:pushIndent(  )
   self:writeRaw( "local " )
   for index, varSym in ipairs( node:get_varSymList() ) do
      self:writeRaw( getSymbolTxt( varSym ) )
      if index ~= #node:get_varSymList() then
         self:writeRaw( ", " )
      end
      
   end
   
   self:writeRaw( " = " )
   
   self:processExpListSub( node, node:get_expList():get_expList(), node:get_expList():get_mRetExp() )
   self:writeln( "" )
   
   self:writeRaw( "if " )
   local hasSym = false
   for __index, varSym in ipairs( node:get_varSymList() ) do
      if varSym:get_name() ~= "_" then
         if hasSym then
            self:writeRaw( " and  " )
         end
         
         self:writeRaw( string.format( "%s ~= nil", getSymbolTxt( varSym )) )
         hasSym = true
      end
      
   end
   
   self:writeRaw( " then" )
   
   filter( node:get_block(), self, node )
   
   do
      local _exp = node:get_nilBlock()
      if _exp ~= nil then
         self:writeRaw( "else" )
         filter( _exp, self, node )
      end
   end
   
   self:writeln( "end" )
   self:popIndent(  )
   self:writeln( "end" )
end


function ConvFilter:processWhen( node, opt )

   self:writeRaw( "if " )
   for index, symPair in ipairs( node:get_symPairList() ) do
      self:writeRaw( string.format( "%s ~= nil", symPair:get_src():get_name()) )
      if index ~= #node:get_symPairList() then
         self:writeRaw( " and " )
      end
      
   end
   
   self:writeRaw( " then" )
   
   filter( node:get_block(), self, node )
   
   do
      local _exp = node:get_elseBlock()
      if _exp ~= nil then
         self:writeRaw( "else" )
         filter( _exp, self, node )
      end
   end
   
   self:writeln( "end" )
end


function ConvFilter:processDeclVar( node, opt )

   if node:get_syncBlock() then
      self:writeln( "do" )
      self:pushIndent(  )
      for __index, varInfo in ipairs( node:get_symbolInfoList() ) do
         self:writeln( string.format( "local _sync_%s", getSymbolTxt( varInfo )) )
      end
      
      self:writeln( "do" )
      self:pushIndent(  )
   end
   
   
   if node:get_mode() ~= Nodes.DeclVarMode.Unwrap and node:get_accessMode(  ) ~= Ast.AccessMode.Global then
      self:writeRaw( "local " )
   end
   
   
   local varList = node:get_symbolInfoList()
   local varNameList = {}
   for index, var in ipairs( varList ) do
      if index > 1 then
         self:writeRaw( ", " )
      end
      
      local name = getSymbolTxt( var )
      self:writeRaw( name )
      table.insert( varNameList, name )
   end
   
   
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         self:writeRaw( " = " )
         filter( _exp, self, node )
      else
         self:writeln( "" )
      end
   end
   
   
   do
      local _exp = node:get_unwrapBlock()
      if _exp ~= nil then
         self:writeln( "" )
         self:writeRaw( "if " )
         for index, varName in ipairs( varNameList ) do
            if index > 1 then
               self:writeRaw( " or " )
            end
            
            self:writeRaw( " nil == " .. varName )
         end
         
         self:writeln( " then" )
         self:pushIndent(  )
         
         for __index, varName in ipairs( varNameList ) do
            self:writeln( string.format( "local _%s = %s", varName, varName) )
         end
         
         self:popIndent(  )
         
         filter( _exp, self, node )
         
         do
            local thenBlock = node:get_thenBlock()
            if thenBlock ~= nil then
               self:writeln( "else" )
               self:pushIndent(  )
               filter( thenBlock, self, node )
               self:popIndent(  )
            end
         end
         
         
         
         self:writeln( "end" )
      end
   end
   
   
   do
      local _exp = node:get_syncBlock()
      if _exp ~= nil then
         filter( _exp, self, node )
         
         local syncVarNameList = {}
         
         for __index, varInfo in ipairs( node:get_symbolInfoList() ) do
            local name = getSymbolTxt( varInfo )
            table.insert( syncVarNameList, name )
            self:writeln( string.format( "_sync_%s = %s", name, name) )
         end
         
         self:popIndent(  )
         self:writeln( "end" )
         
         for __index, name in ipairs( syncVarNameList ) do
            self:writeln( string.format( "%s = _sync_%s", name, name) )
         end
         
         self:popIndent(  )
         self:writeln( "end" )
      end
   end
   
   
   do
      local _switchExp = node:get_accessMode(  )
      if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
         self:writeln( "" )
         for index, varName in ipairs( varNameList ) do
            local name = varName
            if self.needModuleObj then
               self:writeln( string.format( "_moduleObj.%s = %s", name, name) )
            end
            
            self.pubVarName2InfoMap[name] = PubVerInfo.new(node:get_staticFlag(), node:get_accessMode(), node:get_symbolInfoList()[index]:get_mutable(), node:get_typeInfoList()[index])
         end
         
      end
   end
   
   
   if self.macroDepth > 0 then
      self:writeln( "" )
      for __index, symbolInfo in ipairs( node:get_symbolInfoList() ) do
         local varName = getSymbolTxt( symbolInfo )
         self:writeln( string.format( "table.insert( macroVar.__names, '%s' )", varName) )
         self:writeln( string.format( "macroVar.%s = %s", varName, varName) )
         self.macroVarSymSet[symbolInfo]= true
      end
      
   end
   
end



function ConvFilter:processDeclArg( node, opt )

   self:writeRaw( string.format( "%s", node:get_name(  ).txt ) )
end



function ConvFilter:processDeclArgDDD( node, opt )

   self:writeRaw( "..." )
end



function ConvFilter:processDeclFunc( node, opt )

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
      letTxt = "local "
   end
   
   self:writeRaw( string.format( "%sfunction %s( ", letTxt, name ) )
   
   local argList = declInfo:get_argList(  )
   self:processExpListSub( node, argList, nil )
   
   self:writeln( " )" )
   
   do
      local _exp = declInfo:get_body()
      if _exp ~= nil then
         self:process__func__symbol( declInfo:get_has__func__Symbol(), node:get_expType():get_parentInfo(), name )
         filter( _exp, self, node )
      end
   end
   
   self:writeRaw( "end" )
   
   local expType = node:get_expType(  )
   do
      local _switchExp = expType:get_accessMode(  )
      if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
         if self.needModuleObj then
            self:writeln( "" )
            self:writeRaw( string.format( "_moduleObj.%s = %s", name, name) )
         end
         
         
         self.pubFuncName2InfoMap[name] = PubFuncInfo.new(declInfo:get_accessMode(  ), node:get_expType(  ))
      end
   end
   
end



function ConvFilter:processRefType( node, opt )

   do
      local _switchExp = node:get_mutMode()
      if _switchExp == Ast.MutMode.IMut then
         self:writeRaw( "&" )
      elseif _switchExp == Ast.MutMode.AllMut then
         self:writeRaw( "+" )
      end
   end
   
   filter( node:get_name(  ), self, node )
   if node:get_array(  ) == "array" then
      self:writeRaw( "[@]" )
   elseif node:get_array(  ) == "list" then
      self:writeRaw( "[]" )
   end
   
end



function ConvFilter:processIf( node, opt )

   local valList = node:get_stmtList(  )
   for index, val in ipairs( valList ) do
      if index == 1 then
         self:writeRaw( "if " )
         filter( val:get_exp(), self, node )
      elseif val:get_kind() == Nodes.IfKind.ElseIf then
         self:writeRaw( "elseif " )
         filter( val:get_exp(), self, node )
      else
       
         self:writeln( "else" )
      end
      
      self:writeRaw( " " )
      filter( val:get_block(), self, node )
   end
   
   self:writeln( "end" )
end



function ConvFilter:processSwitch( node, opt )

   self:writeln( "do" )
   self:pushIndent(  )
   self:writeRaw( "local _switchExp = " )
   filter( node:get_exp(  ), self, node )
   self:writeln( "" )
   
   if #node:get_caseList() > 0 then
      for index, caseInfo in ipairs( node:get_caseList() ) do
         if index == 1 then
            self:writeRaw( "if " )
         else
          
            self:writeRaw( "elseif " )
         end
         
         local expList = caseInfo:get_expList(  )
         for listIndex, expNode in ipairs( expList:get_expList(  ) ) do
            if listIndex ~= 1 then
               self:writeRaw( " or " )
            end
            
            
            self:writeRaw( "_switchExp == " )
            filter( expNode, self, node )
         end
         
         self:writeRaw( " then" )
         filter( caseInfo:get_block(), self, node )
      end
      
      do
         local _exp = node:get_default(  )
         if _exp ~= nil then
            self:writeln( "else " )
            self:pushIndent(  )
            filter( _exp, self, node )
            self:popIndent(  )
         end
      end
      
      self:writeln( "end" )
   end
   
   self:popIndent(  )
   
   self:writeln( "end" )
end



function ConvFilter:processMatch( node, opt )

   self:writeln( "do" )
   self:pushIndent(  )
   self:writeRaw( "local _matchExp = " )
   filter( node:get_val(), self, node )
   self:writeln( "" )
   
   if #node:get_caseList() > 0 then
      local fullName = self:getFullName( node:get_algeTypeInfo() )
      for index, caseInfo in ipairs( node:get_caseList() ) do
         if index == 1 then
            self:writeRaw( "if " )
         else
          
            self:writeRaw( "elseif " )
         end
         
         self:writeln( string.format( "_matchExp[1] == %s.%s[1] then", fullName, caseInfo:get_valInfo():get_name()) )
         for paramNum, paramSym in ipairs( caseInfo:get_valParamNameList() ) do
            self:writeln( string.format( "   local %s = _matchExp[2][%d]", paramSym:get_name(), paramNum) )
         end
         
         filter( caseInfo:get_block(), self, node )
      end
      
      do
         local _exp = node:get_defaultBlock()
         if _exp ~= nil then
            self:writeln( "else " )
            self:pushIndent(  )
            filter( _exp, self, node )
            self:popIndent(  )
         end
      end
      
      self:writeln( "end" )
   end
   
   self:popIndent(  )
   
   self:writeln( "end" )
end



function ConvFilter:processWhile( node, opt )

   self:writeRaw( "while " )
   
   filter( node:get_exp(  ), self, node )
   self:writeRaw( " " )
   filter( node:get_block(  ), self, node )
   self:writeln( "end" )
end



function ConvFilter:processRepeat( node, opt )

   self:writeRaw( "repeat " )
   filter( node:get_block(  ), self, node )
   self:writeRaw( "until " )
   filter( node:get_exp(  ), self, node )
end



function ConvFilter:processFor( node, opt )

   self:writeRaw( string.format( "for %s = ", getSymbolTxt( node:get_val() )) )
   filter( node:get_init(  ), self, node )
   self:writeRaw( ", " )
   filter( node:get_to(  ), self, node )
   do
      local _exp = node:get_delta(  )
      if _exp ~= nil then
         self:writeRaw( ", " )
         filter( _exp, self, node )
      end
   end
   
   self:writeRaw( " " )
   filter( node:get_block(  ), self, node )
   self:writeln( "end" )
end



function ConvFilter:processApply( node, opt )

   self:writeRaw( "for " )
   local varList = node:get_varList()
   for index, var in ipairs( varList ) do
      if index > 1 then
         self:writeRaw( ", " )
      end
      
      self:writeRaw( getSymbolTxt( var ) )
   end
   
   self:writeRaw( " in " )
   filter( node:get_expList(), self, node )
   self:writeRaw( " " )
   filter( node:get_block(), self, node )
   self:writeln( "end" )
end



function ConvFilter:processForeach( node, opt )

   local keySym
   
   local valSym
   
   if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Set then
      keySym = node:get_val()
      valSym = node:get_key()
   else
    
      keySym = node:get_key()
      valSym = node:get_val()
   end
   
   
   self:writeRaw( "for " )
   if keySym ~= nil then
      self:writeRaw( getSymbolTxt( keySym ) )
   else
      self:writeRaw( "__index" )
   end
   
   self:writeRaw( ", " )
   if valSym ~= nil then
      self:writeRaw( getSymbolTxt( valSym ) )
   else
      self:writeRaw( "__val" )
   end
   
   
   if self.useIpairs and node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.List then
      if node:get_exp():get_expType():get_itemTypeInfoList()[1]:get_nilable() then
         
         self:writeRaw( " in pairs( " )
      else
       
         
         self:writeRaw( " in ipairs( " )
      end
      
   else
    
      self:writeRaw( " in pairs( " )
   end
   
   filter( node:get_exp(), self, node )
   self:writeRaw( " ) " )
   filter( node:get_block(), self, node )
   self:writeln( "end" )
end



function ConvFilter:processForsort( node, opt )

   local keySym
   
   local valSym
   
   if node:get_exp():get_expType():get_kind() == Ast.TypeInfoKind.Set then
      keySym = node:get_val()
      valSym = node:get_key()
   else
    
      keySym = node:get_key()
      valSym = node:get_val()
   end
   
   
   self:writeln( "do" )
   self:pushIndent(  )
   self:writeln( "local __sorted = {}" )
   self:writeRaw( "local __map = " )
   filter( node:get_exp(), self, node )
   self:writeln( "" )
   self:writeln( "for __key in pairs( __map ) do" )
   self:pushIndent(  )
   self:writeln( "table.insert( __sorted, __key )" )
   self:popIndent(  )
   self:writeln( "end" )
   
   self:writeln( "table.sort( __sorted )" )
   
   self:writeRaw( "for __index, " )
   local key = "__key"
   if keySym ~= nil then
      key = getSymbolTxt( keySym )
   end
   
   self:writeRaw( key )
   self:writeln( " in ipairs( __sorted ) do" )
   self:pushIndent(  )
   if valSym ~= nil then
      self:writeln( string.format( "local %s = __map[ %s ]", getSymbolTxt( valSym ), key ) )
   end
   
   filter( node:get_block(), self, node )
   
   self:writeln( "end" )
   self:popIndent(  )
   self:writeln( "end" )
   self:popIndent(  )
   self:writeln( "end" )
end



function ConvFilter:processExpUnwrap( node, opt )

   do
      local _exp = node:get_default()
      if _exp ~= nil then
         self:writeRaw( '_lune.unwrapDefault( ' )
         filter( node:get_exp(), self, node )
         self:writeRaw( ', ' )
         filter( _exp, self, node )
         self:writeRaw( ')' )
      else
         self:writeRaw( '_lune.unwrap( ' )
         filter( node:get_exp(), self, node )
         self:writeRaw( ')' )
      end
   end
   
end


function ConvFilter:processExpCall( node, opt )

   local wroteFuncFlag = false
   local setArgFlag = false
   
   local function fieldCall(  )
   
      local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
      if  nil == fieldNode then
         local _fieldNode = fieldNode
      
         return true
      end
      
      
      do
         local _switchExp = node:get_func():get_expType()
         if _switchExp == self.builtinFunc.__lns_runtime_log or _switchExp == self.builtinFunc.__lns_runtime_enableLog or _switchExp == self.builtinFunc.__lns_runtime_dumpLog or _switchExp == self.builtinFunc.__processor_end then
            return false
         elseif _switchExp == self.builtinFunc.list___new or _switchExp == self.builtinFunc.__lns_sync_createProcesser then
            self:writeRaw( "{}" )
            return false
         elseif _switchExp == self.builtinFunc.__lns_sync_createFlag then
            self:writeRaw( "nil" )
            return false
         end
      end
      
      
      local prefixNode = fieldNode:get_prefix()
      
      if prefixNode:get_expType():isInheritFrom( self.processInfo, Ast.builtinTypeAsyncItem, nil ) and node:get_func():get_expType():get_rawTxt() == "_createPipe" then
         self:writeRaw( "nil" )
         return false
      end
      
      
      local fieldTxt
      
      if node:get_func():get_expType():get_nonnilableType() == self.builtinFunc.str_replace then
         wroteFuncFlag = true
         setArgFlag = true
         self:writeRaw( "_lune.replace( " )
         filter( prefixNode, self, fieldNode )
         return true
      else
       
         fieldTxt = fieldNode:get_field().txt
      end
      
      
      local function processSet(  )
      
         setArgFlag = true
         wroteFuncFlag = true
         
         do
            local _switchExp = fieldTxt
            if _switchExp == "add" or _switchExp == "del" then
               filter( prefixNode, self, fieldNode )
               self:writeRaw( "[" )
               do
                  local argList = node:get_argList()
                  if argList ~= nil then
                     filter( argList, self, fieldNode )
                  end
               end
               
               self:writeRaw( "]" )
               do
                  local _switchExp = fieldTxt
                  if _switchExp == "add" then
                     self:writeRaw( "= true" )
                  elseif _switchExp == "del" then
                     self:writeRaw( "= nil" )
                  end
               end
               
               return false
            end
         end
         
         
         self:writeRaw( string.format( "_lune._Set_%s(", fieldTxt) )
         filter( prefixNode, self, fieldNode )
         return true
      end
      
      local prefixType = prefixNode:get_expType()
      
      local function processEnumAlge(  )
      
         wroteFuncFlag = true
         local fieldExpType = fieldNode:get_expType()
         local canonicalName = self:getFullName( prefixType )
         local methodName = fieldTxt
         local delimit = ":"
         if methodName == "get__txt" then
            methodName = "_getTxt"
         end
         
         if fieldExpType:get_kind() == Ast.TypeInfoKind.Func then
            delimit = "."
         end
         
         self:writeRaw( string.format( "%s%s%s( ", canonicalName, delimit, methodName) )
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
               self:writeRaw( string.format( "_lune.nilacc( table.%s, nil, 'list', ", fieldTxt) )
               filter( prefixNode, self, fieldNode )
            else 
               
                  self:writeRaw( "_lune.nilacc( " )
                  filter( prefixNode, self, fieldNode )
                  self:writeRaw( string.format( ", '%s', 'callmtd' ", fieldTxt) )
            end
         end
         
      else
       
         do
            local _switchExp = prefixType:get_kind()
            if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
               setArgFlag = true
               wroteFuncFlag = true
               self:writeRaw( string.format( "table.%s( ", fieldTxt) )
               filter( prefixNode, self, fieldNode )
            elseif _switchExp == Ast.TypeInfoKind.Set then
               if not processSet(  ) then
                  return false
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.Alge then
               processEnumAlge(  )
            elseif _switchExp == Ast.TypeInfoKind.Box then
               filter( prefixNode, self, fieldNode )
               self:writeRaw( "[1]" )
               return false
            elseif _switchExp == Ast.TypeInfoKind.Class then
               if prefixType:isInheritFrom( self.processInfo, Ast.builtinTypeMapping, nil ) and isGenericType( prefixType ) and (fieldTxt == "_fromMap" or fieldTxt == "_fromStem" ) then
                  wroteFuncFlag = true
                  setArgFlag = true
                  filter( node:get_func(), self, node )
                  self:writeRaw( "( " )
                  do
                     local argList = node:get_argList()
                     if argList ~= nil then
                        filter( argList, self, node )
                        self:writeRaw( ", " )
                     end
                  end
                  
                  self:outputAlter2MapFunc( self, prefixType:createAlt2typeMap( false ) )
                  self:writeRaw( ")" )
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
         if refNode:get_symbolInfo():get_name() == "super" then
            wroteFuncFlag = true
            setArgFlag = true
            local funcType = refNode:get_expType()
            self:writeRaw( string.format( "%s.%s( self ", self:getFullName( funcType:get_parentInfo() ), funcType:get_rawTxt()) )
         else
          
            do
               local _switchExp = refNode:get_expType()
               if _switchExp == self.builtinFunc.lns_expandLuavalMap then
                  
                  wroteFuncFlag = true
                  self:writeRaw( "(" )
               elseif _switchExp == self.builtinFunc.lns___run then
                  self:writeRaw( "_lune._run(" )
                  wroteFuncFlag = true
               elseif _switchExp == self.builtinFunc.lns___join then
                  return 
               end
            end
            
         end
         
      end
   end
   
   
   if not wroteFuncFlag then
      if node:get_nilAccess() then
         self:writeRaw( "_lune.nilacc( " )
         filter( node:get_func(), self, node )
         self:writeRaw( ", nil, 'call'" )
         wroteFuncFlag = true
         setArgFlag = true
      else
       
         filter( node:get_func(), self, node )
         self:writeRaw( "( " )
      end
      
   end
   
   
   local convStrFlag = false
   
   if not self.targetLuaVer:get_canFormStem2Str() and self.builtinFunc:isStrFormFunc( node:get_func():get_expType() ) then
      convStrFlag = true
   end
   
   
   do
      local argList = node:get_argList()
      if argList ~= nil then
         local expList = {}
         for __index, expNode in ipairs( argList:get_expList() ) do
            if expNode:get_expType():get_kind() ~= Ast.TypeInfoKind.Abbr then
               do
                  local toDDD = _lune.__Cast( expNode, 3, Nodes.ExpToDDDNode )
                  if toDDD ~= nil then
                     for __index, appNode in ipairs( toDDD:get_expList():get_expList() ) do
                        table.insert( expList, appNode )
                     end
                     
                  else
                     table.insert( expList, expNode )
                  end
               end
               
            end
            
         end
         
         
         if wroteFuncFlag and setArgFlag then
            if #expList > 0 then
               self:writeRaw( ", " )
            end
            
         end
         
         if convStrFlag then
            local opList = {}
            if #expList > 0 then
               local literal = expList[1]:getLiteral(  )
               if literal ~= nil then
                  do
                     local _matchExp = literal
                     if _matchExp[1] == Nodes.Literal.Str[1] then
                        local txt = _matchExp[2][1]
                     
                        opList = TransUnit.findForm( txt )
                     end
                  end
                  
               end
               
            end
            
            for index, argNode in ipairs( expList ) do
               local filtered = false
               if index > 1 then
                  self:writeRaw( ", " )
                  
                  if index - 1 <= #opList then
                     local formType = TransUnit.isMatchStringFormatType( opList[index - 1], argNode:get_expType(), self.targetLuaVer )
                     if formType == TransUnit.FormType.NeedConv then
                        self:writeRaw( "tostring( " )
                        filter( argNode, self, node )
                        self:writeRaw( ")" )
                        filtered = true
                     end
                     
                  end
                  
               end
               
               if not filtered then
                  filter( argNode, self, node )
               end
               
            end
            
         else
          
            filter( argList, self, node )
         end
         
      end
   end
   
   self:writeRaw( " )" )
end



function ConvFilter:processExpList( node, opt )

   local expList = node:get_expList()
   
   self:processExpListSub( node, expList, node:get_mRetExp() )
end



function ConvFilter:processExpOp1( node, opt )

   local op = node:get_op().txt
   if op == ",,," then
      filter( node:get_exp(), self, node )
   elseif op == ",,,," then
      if node:get_macroMode() ~= Nodes.MacroMode.None then
         filter( node:get_exp(), self, node )
      else
       
         self:writeRaw( "__luneSym2Str( " )
         filter( node:get_exp(), self, node )
         self:writeRaw( " )" )
      end
      
   elseif op == ",," then
      do
         local _switchExp = node:get_exp():get_expType()
         if _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeReal or _switchExp == Ast.builtinTypeBool then
            filter( node:get_exp(), self, node )
         else 
            
               self:writeRaw( "__luneGetLocal( " )
               filter( node:get_exp(), self, node )
               self:writeRaw( " )" )
         end
      end
      
   elseif op == "~" then
      if self.targetLuaVer:get_hasBitOp() == LuaVer.BitOp.HasOp then
         self:writeRaw( op )
         filter( node:get_exp(), self, node )
      else
       
         self:writeRaw( "bit32.bnot( " )
         filter( node:get_exp(), self, node )
         self:writeRaw( " )" )
      end
      
   else
    
      if op == "not" then
         op = op .. " "
      end
      
      self:writeRaw( op )
      filter( node:get_exp(), self, node )
   end
   
end



function ConvFilter:processExpToDDD( node, opt )

   self:processExpListSub( node, node:get_expList():get_expList(), node:get_expList():get_mRetExp() )
end


function ConvFilter:processExpMultiTo1( node, opt )

   filter( node:get_exp(), self, node )
end


function ConvFilter:processExpCast( node, opt )

   do
      local _switchExp = node:get_castKind()
      if _switchExp == Nodes.CastKind.Force then
         if node:get_expType():equals( self.processInfo, Ast.builtinTypeInt ) then
            self:writeRaw( "math.floor(" )
            filter( node:get_exp(), self, node )
            self:writeRaw( ")" )
         elseif node:get_expType():equals( self.processInfo, Ast.builtinTypeReal ) then
            filter( node:get_exp(), self, node )
            self:writeRaw( " * 1.0" )
         else
          
            filter( node:get_exp(), self, node )
         end
         
      elseif _switchExp == Nodes.CastKind.Normal then
         self:writeRaw( "_lune.__Cast( " )
         filter( node:get_exp(), self, node )
         local castKind
         
         local classObj = "nil"
         do
            local _switchExp = node:get_expType():get_nonnilableType()
            if _switchExp == Ast.builtinTypeInt then
               castKind = LuaMod.CastKind.Int
            elseif _switchExp == Ast.builtinTypeReal then
               castKind = LuaMod.CastKind.Real
            elseif _switchExp == Ast.builtinTypeString then
               castKind = LuaMod.CastKind.Str
            else 
               
                  castKind = LuaMod.CastKind.Class
                  classObj = self:getFullName( node:get_expType():get_nonnilableType() )
            end
         end
         
         self:writeRaw( string.format( ", %d, %s )", castKind, classObj) )
      elseif _switchExp == Nodes.CastKind.Implicit then
         filter( node:get_exp(), self, node )
      end
   end
   
end



function ConvFilter:processExpParen( node, opt )

   self:writeRaw( "(" )
   filter( node:get_exp(), self, node )
   self:writeRaw( " )" )
end



function ConvFilter:processExpSetVal( node, opt )

   filter( node:get_exp1(), self, node )
   self:writeRaw( " = " )
   filter( node:get_exp2(), self, node )
end



function ConvFilter:processExpSetItem( node, opt )

   filter( node:get_val(), self, node )
   self:writeRaw( "[" )
   do
      local _matchExp = node:get_index()
      if _matchExp[1] == Nodes.IndexVal.NodeIdx[1] then
         local index = _matchExp[2][1]
      
         filter( index, self, node )
      elseif _matchExp[1] == Nodes.IndexVal.SymIdx[1] then
         local index = _matchExp[2][1]
      
         self:writeRaw( string.format( "'%s'", index) )
      end
   end
   
   self:writeRaw( "]" )
   self:writeRaw( " = " )
   filter( node:get_exp2(), self, node )
end


function ConvFilter:processExpOp2( node, opt )

   local intCast = false
   if node:get_expType():equals( self.processInfo, Ast.builtinTypeInt ) and node:get_op().txt == "/" then
      intCast = true
      self:writeRaw( "math.floor(" )
   end
   
   
   local opTxt = node:get_op().txt
   
   do
      local _exp = Ast.bitBinOpMap[opTxt]
      if _exp ~= nil then
         
         if self.targetLuaVer:get_hasBitOp() == LuaVer.BitOp.HasOp then
            
            do
               local _switchExp = _exp
               if _switchExp == Ast.BitOpKind.LShift then
                  opTxt = "<<"
               elseif _switchExp == Ast.BitOpKind.RShift then
                  opTxt = ">>"
               end
            end
            
            filter( node:get_exp1(), self, node )
            self:writeRaw( " " .. opTxt .. " " )
            filter( node:get_exp2(), self, node )
         else
          
            
            local binfunc = ""
            local exp2Mod = ""
            do
               local _switchExp = _exp
               if _switchExp == Ast.BitOpKind.And then
                  binfunc = "band"
               elseif _switchExp == Ast.BitOpKind.Or then
                  binfunc = "bor"
               elseif _switchExp == Ast.BitOpKind.Xor then
                  binfunc = "bxor"
               elseif _switchExp == Ast.BitOpKind.LShift then
                  binfunc = "lshift"
               elseif _switchExp == Ast.BitOpKind.RShift then
                  binfunc = "lshift"
                  exp2Mod = "-"
               end
            end
            
            self:writeRaw( string.format( "bit32.%s(", binfunc) )
            filter( node:get_exp1(), self, node )
            self:writeRaw( ", " )
            self:writeRaw( exp2Mod )
            filter( node:get_exp2(), self, node )
            self:writeRaw( " )" )
         end
         
      else
         filter( node:get_exp1(), self, node )
         self:writeRaw( " " .. opTxt .. " " )
         filter( node:get_exp2(), self, node )
      end
   end
   
   
   if intCast then
      self:writeRaw( ")" )
   end
   
end



function ConvFilter:processExpRef( node, opt )

   do
      local _switchExp = node:get_symbolInfo():get_name()
      if _switchExp == "super" then
         local funcType = node:get_expType()
         self:writeRaw( string.format( "%s.%s", self:getFullName( funcType:get_parentInfo() ), funcType:get_rawTxt()) )
      else 
         
            local builtinFunc = self.builtinFunc
            if node:get_expType():equals( self.processInfo, builtinFunc.lns__load ) then
               self:writeRaw( "_lune." .. self.targetLuaVer:get_loadStrFuncName() )
            else
             
               if _lune._Set_has(self.macroVarSymSet, node:get_symbolInfo():getOrg(  ) ) then
                  self:writeRaw( "macroVar." )
               else
                
                  if node:get_symbolInfo():get_accessMode() == Ast.AccessMode.Pub and node:get_symbolInfo():get_kind() == Ast.SymbolKind.Var then
                     if self.needModuleObj then
                        self:writeRaw( "_moduleObj." )
                     end
                     
                  end
                  
               end
               
               self:writeRaw( node:get_symbolInfo():get_name() )
               if node:get_symbolInfo():get_isLazyLoad() then
                  self:writeRaw( "()" )
               end
               
            end
            
      end
   end
   
end



function ConvFilter:processExpRefItem( node, opt )

   if node:get_nilAccess() then
      self:writeRaw( "_lune.nilacc( " )
      filter( node:get_val(), self, node )
      self:writeRaw( ", nil, 'item', " )
      do
         local _exp = node:get_index()
         if _exp ~= nil then
            filter( _exp, self, node )
         else
            self:writeRaw( string.format( "'%s'", _lune.unwrap( node:get_symbol())) )
         end
      end
      
      self:writeRaw( ")" )
   else
    
      if node:get_val():get_expType():equals( self.processInfo, Ast.builtinTypeString ) then
         self:writeRaw( "string.byte( " )
         filter( node:get_val(), self, node )
         self:writeRaw( ", " )
         do
            local _exp = node:get_index()
            if _exp ~= nil then
               filter( _exp, self, node )
            else
               error( "index is nil" )
            end
         end
         
         self:writeRaw( " )" )
      else
       
         filter( node:get_val(), self, node )
         self:writeRaw( "[" )
         do
            local _exp = node:get_index()
            if _exp ~= nil then
               filter( _exp, self, node )
            else
               self:writeRaw( string.format( "'%s'", _lune.unwrap( node:get_symbol())) )
            end
         end
         
         self:writeRaw( "]" )
      end
      
   end
   
end



function ConvFilter:processRefField( node, opt )

   do
      local symbol = node:get_symbolInfo()
      if symbol ~= nil then
         do
            local code = self.builtinSym2code[symbol:getOrg(  )]
            if code ~= nil then
               self:writeRaw( code )
               return 
            end
         end
         
      end
   end
   
   
   local parent = opt.node
   local prefix = node:get_prefix(  )
   
   if node:get_nilAccess() then
      self:writeRaw( '_lune.nilacc( ' )
      filter( prefix, self, node )
      self:writeRaw( string.format( ', "%s" )', node:get_field().txt) )
   else
    
      filter( prefix, self, node )
      local delimit = "."
      if parent:get_kind() == Nodes.NodeKind.get_ExpCall() then
         if node:get_expType(  ):get_kind(  ) == Ast.TypeInfoKind.Method then
            delimit = ":"
         else
          
            delimit = "."
         end
         
      end
      
      local fieldToken = node:get_field(  )
      self:writeRaw( delimit .. fieldToken.txt )
      do
         local symbolInfo = node:get_symbolInfo()
         if symbolInfo ~= nil then
            if symbolInfo:get_isLazyLoad() then
               self:writeRaw( "()" )
            end
            
         end
      end
      
   end
   
end



function ConvFilter:processExpOmitEnum( node, opt )

   do
      local aliasType = node:get_aliasType()
      if aliasType ~= nil then
         self:writeRaw( self:getFullName( aliasType ) )
      else
         self:writeRaw( self:getFullName( node:get_expType() ) )
      end
   end
   
   self:writeRaw( string.format( ".%s", node:get_valToken().txt) )
end



function ConvFilter:processGetField( node, opt )

   local prefixNode = node:get_prefix(  )
   local prefixType = prefixNode:get_expType()
   local fieldTxt = node:get_field(  ).txt
   
   if fieldTxt == "_txt" and (prefixType:get_kind() == Ast.TypeInfoKind.Enum or prefixType:get_kind() == Ast.TypeInfoKind.Alge ) then
      self:writeRaw( string.format( "%s:_getTxt( ", self:getFullName( prefixType )) )
      filter( prefixNode, self, node )
      self:writeln( ")" )
   else
    
      
      if node:get_nilAccess() then
         fieldTxt = string.format( "get_%s", fieldTxt)
         self:writeRaw( "_lune.nilacc( " )
         filter( prefixNode, self, node )
         self:writeRaw( string.format( ", '%s', 'callmtd' )", fieldTxt) )
      else
       
         fieldTxt = string.format( "get_%s()", fieldTxt)
         filter( prefixNode, self, node )
         local delimit = "."
         if node:get_getterTypeInfo(  ):get_kind(  ) == Ast.TypeInfoKind.Method then
            delimit = ":"
         else
          
            delimit = "."
         end
         
         
         self:writeRaw( delimit .. fieldTxt )
      end
      
   end
   
end



function ConvFilter:processReturn( node, opt )

   self:writeRaw( "return " )
   
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         filter( _exp, self, node )
      end
   end
   
end



function ConvFilter:processLuneKind( node, opt )

   do
      local workNode = _lune.__Cast( node:get_exp(), 3, Nodes.ExpCastNode )
      if workNode ~= nil then
         if workNode:get_castKind() == Nodes.CastKind.Implicit then
            self:writeRaw( string.format( "%d", workNode:get_exp():get_expType():get_kind()) )
         end
         
      else
         self:writeRaw( string.format( "%d", node:get_exp():get_expType():get_kind()) )
      end
   end
   
end


function ConvFilter:processTestCase( node, opt )

   if self.enableTest then
      self:writeln( "do" )
      self:pushIndent(  )
      filter( node:get_impNode(), self, node )
      self:writeln( "" )
      self:writeln( string.format( "local function testcase( %s ) ", node:get_ctrlName()) )
      
      filter( node:get_block(), self, node )
      
      self:writeln( "end" )
      self:writeln( string.format( '__t.registerTestcase( "%s", "%s", testcase )', self.moduleTypeInfo:getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager() ):gsub( "@", "" ), node:get_name().txt) )
      self:popIndent(  )
      self:writeln( "end" )
   end
   
end


function ConvFilter:processTestBlock( node, opt )

   if self.enableTest then
      for __index, statement in ipairs( node:get_stmtList() ) do
         filter( statement, self, node )
         self:writeln( "" )
      end
      
   end
   
end



function ConvFilter:processProvide( node, opt )

end


function ConvFilter:processAlias( node, opt )

   self:writeRaw( string.format( "local %s = ", node:get_newSymbol():get_name()) )
   filter( node:get_srcNode(), self, node )
   if Ast.isPubToExternal( node:get_expType():get_accessMode() ) then
      self:writeRaw( string.format( "\n_moduleObj.%s = %s", node:get_newSymbol():get_name(), node:get_newSymbol():get_name()) )
   end
   
end


function ConvFilter:processBoxing( node, opt )

   self:writeRaw( "{" )
   
   filter( node:get_src(), self, node )
   
   self:writeRaw( "}" )
end


function ConvFilter:processUnboxing( node, opt )

   filter( node:get_src(), self, node )
   self:writeRaw( "[1]" )
end


function ConvFilter:processLiteralList( node, opt )

   self:writeRaw( "{" )
   
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         filter( _exp, self, node )
      end
   end
   
   
   self:writeRaw( "}" )
end



function ConvFilter:processLiteralSet( node, opt )

   self:writeRaw( "{" )
   do
      local expListNode = node:get_expList()
      if expListNode ~= nil then
         for index, expNode in ipairs( expListNode:get_expList() ) do
            if index > 1 then
               self:writeRaw( ", " )
            end
            
            self:writeRaw( "[" )
            filter( expNode, self, node )
            self:writeRaw( "] = true" )
         end
         
      end
   end
   
   
   self:writeRaw( "}" )
end



function ConvFilter:processLiteralMap( node, opt )

   self:writeRaw( "{" )
   local pairList = node:get_pairList()
   for index, pair in ipairs( pairList ) do
      if index > 1 then
         self:writeRaw( ", " )
      end
      
      self:writeRaw( "[" )
      filter( pair:get_key(), self, node )
      self:writeRaw( "] = " )
      filter( pair:get_val(), self, node )
   end
   
   
   self:writeRaw( "}" )
end



function ConvFilter:processLiteralArray( node, opt )

   self:writeRaw( "{" )
   
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         filter( _exp, self, node )
      end
   end
   
   
   self:writeRaw( "}" )
end



function ConvFilter:processLiteralChar( node, opt )

   self:writeRaw( string.format( "%d", node:get_num() ) )
end



function ConvFilter:processLiteralInt( node, opt )

   
   self:writeRaw( node:get_token().txt )
end



function ConvFilter:processLiteralReal( node, opt )

   
   self:writeRaw( node:get_token().txt )
end



function ConvFilter:processLiteralString( node, opt )

   local txt = node:get_token(  ).txt
   if string.find( txt, '^```' ) then
      txt = '[==[' .. txt:sub( 4, -4 ) .. ']==]'
   end
   
   local opList = TransUnit.findForm( txt )
   
   do
      local expList = node:get_orgParam()
      if expList ~= nil then
         local mRetIndex = _lune.nilacc( expList:get_mRetExp(), 'get_index', 'callmtd' )
         
         self:writeRaw( string.format( 'string.format( %s, ', txt ) )
         for index, val in ipairs( expList:get_expList() ) do
            if index > 1 then
               self:writeRaw( ", " )
            end
            
            
            local matchFlag = TransUnit.FormType.Match
            if index <= #opList then
               matchFlag = TransUnit.isMatchStringFormatType( opList[index], val:get_expType(), self.targetLuaVer )
            end
            
            if matchFlag == TransUnit.FormType.NeedConv then
               self:writeRaw( "tostring( " )
               filter( val, self, node )
               self:writeRaw( ")" )
            else
             
               filter( val, self, node )
            end
            
            if index == mRetIndex then
               break
            end
            
         end
         
         self:writeRaw( ")" )
      else
         self:writeRaw( txt )
      end
   end
   
end



function ConvFilter:processLiteralBool( node, opt )

   self:writeRaw( node:get_token().txt )
end



function ConvFilter:processLiteralNil( node, opt )

   self:writeRaw( "nil" )
end



function ConvFilter:processBreak( node, opt )

   self:writeRaw( "break" )
end



function ConvFilter:processLiteralSymbol( node, opt )

   self:writeRaw( string.format( '%s', node:get_token().txt) )
end



function ConvFilter:processLuneControl( node, opt )

   do
      local _matchExp = node:get_pragma()
      if _matchExp[1] == LuneControl.Pragma.load__lune_module[1] then
      
         self:processLoadRuntime(  )
      end
   end
   
end


local FilterInfo = {}
_moduleObj.FilterInfo = FilterInfo
function FilterInfo:outputLuaAndMeta( node )

   
   node:processFilter( self.filter, Opt.new(node) )
end
function FilterInfo:outputLua( node )

   node:processFilter( self.filter, Opt.new(node) )
end
function FilterInfo:outputMeta( node )

end
function FilterInfo.setmeta( obj )
  setmetatable( obj, { __index = FilterInfo  } )
end
function FilterInfo.new( filter )
   local obj = {}
   FilterInfo.setmeta( obj )
   if obj.__init then
      obj:__init( filter )
   end
   return obj
end
function FilterInfo:__init( filter )

   self.filter = filter
end
function FilterInfo:get_filter()
   return self.filter
end


local function createFilter( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, processInfo, moduleSymbolKind, builtinFunc, useLuneRuntime, targetLuaVer, enableTest, useIpairs, option )

   local convFilter = ConvFilter.new(streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, processInfo, moduleSymbolKind, builtinFunc, useLuneRuntime, targetLuaVer, enableTest, useIpairs, option)
   return FilterInfo.new(convFilter)
end
_moduleObj.createFilter = createFilter

local MacroEvalImp = {}
setmetatable( MacroEvalImp, { __index = Nodes.MacroEval } )
_moduleObj.MacroEvalImp = MacroEvalImp
function MacroEvalImp:evalFromCodeToLuaCode( processInfo, name, argNameList, code )

   local stream = Util.memStream.new()
   local conv = ConvFilter.new("macro", stream, Util.NullOStream.new(), ConvMode.ConvMeta, true, Ast.headTypeInfo, processInfo, Ast.SymbolKind.Typ, self.builtinFunc, nil, LuaVer.getCurVer(  ), false, true, Option.new(""))
   
   conv:outputDeclMacro( name, argNameList, function (  )
   
      if code ~= nil then
         conv:write( code )
      end
      
   end )
   
   return stream:get_txt()
end
function MacroEvalImp:evalToLuaCode( processInfo, node )

   local stream = Util.memStream.new()
   local conv = ConvFilter.new("macro", stream, Util.NullOStream.new(), ConvMode.ConvMeta, true, Ast.headTypeInfo, processInfo, Ast.SymbolKind.Typ, self.builtinFunc, nil, LuaVer.getCurVer(  ), false, true, Option.new(""))
   
   conv:processDeclMacro( node, Opt.new(node) )
   
   return stream:get_txt()
end
function MacroEvalImp.setmeta( obj )
  setmetatable( obj, { __index = MacroEvalImp  } )
end
function MacroEvalImp.new( builtinFunc )
   local obj = {}
   MacroEvalImp.setmeta( obj )
   if obj.__init then
      obj:__init( builtinFunc )
   end
   return obj
end
function MacroEvalImp:__init( builtinFunc )

   Nodes.MacroEval.__init( self)
   self.builtinFunc = builtinFunc
end






return _moduleObj
