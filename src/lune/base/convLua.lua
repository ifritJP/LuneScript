--lune/base/convLua.lns
local _moduleObj = {}
local __mod__ = 'lune.base.convLua'
if not _ENV._lune then
   _lune = {}
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

local Ast = require( 'lune.base.Ast' )
local Util = require( 'lune.base.Util' )
local TransUnit = require( 'lune.base.TransUnit' )
local frontInterface = require( 'lune.base.frontInterface' )
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
    
ConvMode.Exec = 0
ConvMode._val2NameMap[0] = 'Exec'
ConvMode.Convert = 1
ConvMode._val2NameMap[1] = 'Convert'
ConvMode.ConvMeta = 2
ConvMode._val2NameMap[2] = 'ConvMeta'

local hasBitOpFlag = _VERSION:gsub( "^[^%d]+", "" ):find( "^5%.3" ) ~= nil
local BitOpKind = {}
BitOpKind._val2NameMap = {}
function BitOpKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "BitOpKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function BitOpKind._from( val )
   if BitOpKind._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
BitOpKind.And = 0
BitOpKind._val2NameMap[0] = 'And'
BitOpKind.Or = 1
BitOpKind._val2NameMap[1] = 'Or'
BitOpKind.Xor = 2
BitOpKind._val2NameMap[2] = 'Xor'
BitOpKind.LShift = 3
BitOpKind._val2NameMap[3] = 'LShift'
BitOpKind.RShift = 4
BitOpKind._val2NameMap[4] = 'RShift'

local bitBinOpMap = {["&"] = BitOpKind.And, ["|"] = BitOpKind.Or, ["~"] = BitOpKind.Xor, ["|>>"] = BitOpKind.RShift, ["|<<"] = BitOpKind.LShift}
local convFilter = {}
setmetatable( convFilter, { __index = Ast.Filter } )
function convFilter.new( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, moduleSymbolKind )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, moduleSymbolKind ); end
   return obj
end
function convFilter:__init(streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, moduleSymbolKind) 
   self.needModuleObj = true
   self.indentQueue = {0}
   self.moduleSymbolKind = moduleSymbolKind
   self.macroDepth = 0
   self.streamName = streamName
   self.stream = stream
   self.metaStream = metaStream
   self.outMetaFlag = false
   self.typeInfo2ModuleName = {}
   self.convMode = convMode
   self.inMacro = inMacro
   self.curLineNo = 1
   self.classId2TypeInfo = {}
   self.classId2MemberList = {}
   self.pubVarName2InfoMap = {}
   self.pubFuncName2InfoMap = {}
   self.pubEnumId2EnumTypeInfo = {}
   self.needIndent = false
   self.moduleTypeInfo = moduleTypeInfo
end
function convFilter:get_indent(  )

   if #self.indentQueue > 0 then
      return self.indentQueue[#self.indentQueue]
   end
   
   return 0
end
function convFilter:getFullName( typeInfo )

   local enumName = typeInfo:getFullName( self.typeInfo2ModuleName, true )
   return string.format( "%s", enumName:gsub( "&", "" ))
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
   if self.outMetaFlag then
      stream = self.metaStream
   end
   
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
   
end
function convFilter.setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end

local function filter( node, filter, parent )

   node:processFilter( filter, parent )
end

local stepIndent = 3
local builtInModuleSet = {}
builtInModuleSet["io"] = true
builtInModuleSet["string"] = true
builtInModuleSet["table"] = true
builtInModuleSet["math"] = true
builtInModuleSet["debug"] = true
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

function convFilter:processNone( node, parent )

end


function convFilter:processImport( node, parent )

   local module = node:get_modulePath(  )
   local moduleName = string.gsub( module, ".*%.", "" )
   self.typeInfo2ModuleName[node:get_moduleTypeInfo()] = module
   if self.convMode == ConvMode.Exec then
      self:write( string.format( "local %s = __luneScript:loadModule( '%s' )", moduleName, module) )
   else
    
      self:write( string.format( "local %s = require( '%s' )", moduleName, module) )
   end
   
end


function convFilter:outputMeta( node )

   if self.convMode == ConvMode.Convert then
      return 
   end
   
   self.outMetaFlag = true
   if self.stream ~= self.metaStream then
      self:writeln( "local _moduleObj = {}" )
   end
   
   self:writeln( "----- meta -----" )
   local importModuleType2Index = {}
   local importNameMap = {}
   do
      for typeInfo, moduleInfo in pairs( node:get_importModule2moduleInfo() ) do
         importNameMap[moduleInfo:get_fullName()] = typeInfo
      end
      
      local index = 0
      do
         local __sorted = {}
         local __map = importNameMap
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, importName in ipairs( __sorted ) do
            typeInfo = __map[ importName ]
            do
               index = index + 1
               importModuleType2Index[typeInfo] = index
            end
         end
      end
      
   end
   
   local typeId2TypeInfo = {}
   local typeId2UseFlag = {}
   local pickupClassMap = {}
   local function checkExportTypeInfo( typeInfo )
   
      local moduleTypeInfo = typeInfo:getModule(  )
      local typeId = typeInfo:get_typeId()
      return typeId2TypeInfo[typeId] and not Ast.isBuiltin( typeId ) and (moduleTypeInfo:hasRouteNamespaceFrom( node:get_moduleTypeInfo() ) or typeInfo:get_srcTypeInfo() ~= typeInfo or moduleTypeInfo:equals( Ast.headTypeInfo ) )
   end
   
   local function pickupTypeId( typeInfo, forceFlag, pickupChildFlag )
   
      if typeInfo:get_typeId(  ) == Ast.rootTypeId then
         return 
      end
      
      if not forceFlag and not Ast.isPubToExternal( typeInfo:get_accessMode() ) then
         return 
      end
      
      if typeId2TypeInfo[typeInfo:get_typeId(  )] then
         if pickupChildFlag and not typeInfo:get_nilable() then
            for __index, itemTypeInfo in pairs( typeInfo:get_children(  ) ) do
               if Ast.isPubToExternal( itemTypeInfo:get_accessMode() ) and (itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Class or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.IF or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Func or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Method ) then
                  pickupTypeId( itemTypeInfo, true, true )
               end
               
            end
            
         end
         
         return 
      end
      
      typeId2TypeInfo[typeInfo:get_typeId(  )] = typeInfo
      if typeInfo:get_nilable() then
         pickupTypeId( typeInfo:get_orgTypeInfo(  ), true, false )
      else
       
         if typeInfo:get_kind() == Ast.TypeInfoKind.Class or typeInfo:get_kind() == Ast.TypeInfoKind.IF then
            pickupClassMap[typeInfo:get_typeId()] = typeInfo
         end
         
         local parentInfo = typeInfo:get_parentInfo(  )
         pickupTypeId( parentInfo, true, false )
         local baseInfo = typeInfo:get_baseTypeInfo(  )
         if baseInfo:get_typeId() ~= Ast.rootTypeId then
            pickupTypeId( baseInfo, true, true )
         end
         
         local typeInfoList = typeInfo:get_itemTypeInfoList(  )
         if typeInfoList then
            for __index, itemTypeInfo in pairs( typeInfoList ) do
               pickupTypeId( itemTypeInfo, true, false )
            end
            
         end
         
         typeInfoList = typeInfo:get_argTypeInfoList(  )
         if typeInfoList then
            for __index, itemTypeInfo in pairs( typeInfoList ) do
               pickupTypeId( itemTypeInfo, true, false )
            end
            
         end
         
         typeInfoList = typeInfo:get_retTypeInfoList(  )
         if typeInfoList then
            for __index, itemTypeInfo in pairs( typeInfoList ) do
               pickupTypeId( itemTypeInfo, true, true )
            end
            
         end
         
         if pickupChildFlag then
            for __index, itemTypeInfo in pairs( typeInfo:get_children(  ) ) do
               if itemTypeInfo:get_accessMode() == Ast.AccessMode.Pub and (itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Class or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.IF or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Func or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Method ) then
                  pickupTypeId( itemTypeInfo, true, true )
               end
               
            end
            
         end
         
         pickupTypeId( typeInfo:get_nilableTypeInfo(  ), true, false )
      end
      
   end
   
   local classId2TypeInfo = {}
   local validChildrenSet = {}
   do
      local typeInfo = self.moduleTypeInfo
      while typeInfo ~= Ast.headTypeInfo do
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
   
   classId2TypeInfo = self.classId2TypeInfo
   self:writeln( "local _typeId2ClassInfoMap = {}" )
   self:writeln( "_moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap" )
   do
      local __sorted = {}
      local __map = classId2TypeInfo
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, classTypeId in ipairs( __sorted ) do
         classTypeInfo = __map[ classTypeId ]
         do
            if classTypeInfo:get_accessMode() == Ast.AccessMode.Pub then
               pickupTypeId( classTypeInfo, true, validChildrenSet[classTypeInfo] == nil and not classTypeInfo:get_externalFlag() )
               pickupClassMap[classTypeId] = nil
               self:writeln( "do" )
               self:pushIndent(  )
               self:writeln( string.format( "local _classInfo%d = {}", classTypeId) )
               self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId) )
               for __index, memberNode in pairs( _lune.unwrap( self.classId2MemberList[classTypeId]) ) do
                  if memberNode:get_accessMode() ~= Ast.AccessMode.Pri then
                     local memberName = memberNode:get_name().txt
                     local memberTypeInfo = memberNode:get_expType(  )
                     self:writeln( string.format( "_classInfo%d.%s = {", classTypeId, memberName) )
                     self:writeln( string.format( "  name='%s', staticFlag = %s, ", memberName, memberNode:get_staticFlag()) .. string.format( "accessMode = '%s', typeId = %d }", memberNode:get_accessMode(), memberTypeInfo:get_typeId(  )) )
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
            workClassMap[classTypeId] = classTypeInfo
            hasWorkClassFlag = true
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
            classTypeInfo = __map[ classTypeId ]
            do
               local scope = _lune.unwrap( classTypeInfo:get_scope())
               if not Ast.isBuiltin( classTypeId ) then
                  pickupTypeId( classTypeInfo, true, validChildrenSet[classTypeInfo] == nil and not classTypeInfo:get_externalFlag() )
                  if checkExportTypeInfo( classTypeInfo ) then
                     local className = classTypeInfo:getTxt(  )
                     self:writeln( "do" )
                     self:pushIndent(  )
                     self:writeln( string.format( "local _classInfo%s = {}", classTypeId) )
                     self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId) )
                     do
                        local __sorted = {}
                        local __map = scope:get_symbol2TypeInfoMap()
                        for __key in pairs( __map ) do
                           table.insert( __sorted, __key )
                        end
                        table.sort( __sorted )
                        for __index, fieldName in ipairs( __sorted ) do
                           symbolInfo = __map[ fieldName ]
                           do
                              local typeInfo = symbolInfo:get_typeInfo()
                              if symbolInfo:get_kind() == Ast.SymbolKind.Mbr or symbolInfo:get_kind() == Ast.SymbolKind.Var then
                                 if symbolInfo:get_accessMode() == Ast.AccessMode.Pub then
                                    self:writeln( string.format( "_classInfo%d.%s = {", classTypeId, fieldName) )
                                    self:writeln( string.format( "  name='%s', staticFlag = %s, ", fieldName, symbolInfo:get_staticFlag()) .. string.format( "accessMode = %d, typeId = %d }", symbolInfo:get_accessMode(), typeInfo:get_typeId(  )) )
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
   
   self:writeln( "local _varName2InfoMap = {}" )
   self:writeln( "_moduleObj._varName2InfoMap = _varName2InfoMap" )
   do
      local __sorted = {}
      local __map = self.pubVarName2InfoMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, varName in ipairs( __sorted ) do
         varInfo = __map[ varName ]
         do
            self:writeln( string.format( "_varName2InfoMap.%s = {", varName ) )
            self:writeln( string.format( "  name='%s', accessMode = %d, typeId = %d, mutable = %s }", varName, varInfo.accessMode, varInfo.typeInfo:get_typeId(), true) )
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
      for __index, funcName in ipairs( __sorted ) do
         funcInfo = __map[ funcName ]
         do
            pickupTypeId( funcInfo.typeInfo, true )
         end
      end
   end
   
   self:writeln( "local _typeInfoList = {}" )
   self:writeln( "_moduleObj._typeInfoList = _typeInfoList" )
   local listIndex = 1
   local wroteTypeIdSet = {}
   local function outputTypeInfo( typeInfo )
   
      do
         local _switchExp = typeInfo:get_kind()
         if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
            do
               local _switchExp = typeInfo:get_accessMode()
               if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Pro or _switchExp == Ast.AccessMode.Global then
               else 
                  
                     Util.errorLog( string.format( "skip: %s %s", self:getFullName( typeInfo )) )
                     return 
               end
            end
            
         end
      end
      
      local typeId = typeInfo:get_typeId(  )
      if wroteTypeIdSet[typeId] then
         return 
      end
      
      wroteTypeIdSet[typeId] = true
      if checkExportTypeInfo( typeInfo ) then
         self:write( string.format( "_typeInfoList[%d] = ", listIndex) )
         listIndex = listIndex + 1
         local validChildren = validChildrenSet[typeInfo]
         if not validChildren then
            validChildren = typeId2TypeInfo
         end
         
         typeInfo:serialize( self, validChildren )
      else
       
      end
      
   end
   
   for typeId, typeInfo in pairs( self.pubEnumId2EnumTypeInfo ) do
      typeId2TypeInfo[typeId] = typeInfo
   end
   
   self:writeln( "local _dependIdMap = {}" )
   self:writeln( "_moduleObj._dependIdMap = _dependIdMap" )
   local exportNeedModuleTypeInfo = {}
   do
      local __sorted = {}
      local __map = typeId2TypeInfo
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, typeId in ipairs( __sorted ) do
         typeInfo = __map[ typeId ]
         do
            outputTypeInfo( typeInfo )
            local moduleTypeInfo = typeInfo:getModule(  )
            exportNeedModuleTypeInfo[moduleTypeInfo] = true
            do
               local moduleIndex = importModuleType2Index[moduleTypeInfo]
               if moduleIndex ~= nil then
                  local moduleInfo = _lune.unwrap( node:get_importModule2moduleInfo()[moduleTypeInfo])
                  do
                     local extId = moduleInfo:get_localTypeInfo2importIdMap()[typeInfo]
                     if extId ~= nil then
                        local valid = false
                        if typeInfo:get_srcTypeInfo() ~= typeInfo then
                           valid = true
                        else
                         
                           local orgTypeInfo = typeInfo
                           if typeInfo:get_nilable() then
                              orgTypeInfo = typeInfo:get_orgTypeInfo()
                           end
                           
                           do
                              local _switchExp = orgTypeInfo:get_kind()
                              if _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.Map or _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array or _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Module then
                                 valid = true
                              end
                           end
                           
                        end
                        
                        if valid then
                           self:writeln( string.format( "_dependIdMap[ %d ] = { %d, %d } -- %s", typeInfo:get_typeId(), moduleIndex, extId, typeInfo:getTxt(  )) )
                        end
                        
                     end
                  end
                  
               end
            end
            
         end
      end
   end
   
   self:writeln( "local _dependModuleMap = {}" )
   self:writeln( "_moduleObj._dependModuleMap = _dependModuleMap" )
   do
      local __sorted = {}
      local __map = importNameMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, name in ipairs( __sorted ) do
         moduleTypeInfo = __map[ name ]
         do
            self:writeln( string.format( "_dependModuleMap[ '%s' ] = { id = %d, use = %s }", name, _lune.unwrap( importModuleType2Index[moduleTypeInfo]), exportNeedModuleTypeInfo[moduleTypeInfo] or false) )
         end
      end
   end
   
   local moduleTypeInfo = self.moduleTypeInfo
   local moduleSymbolKind = Ast.SymbolKind.Typ
   do
      local _exp = node:get_provideNode()
      if _exp ~= nil then
         moduleTypeInfo = _exp:get_symbol():get_typeInfo()
         moduleSymbolKind = _exp:get_symbol():get_kind()
      end
   end
   
   self:writeln( string.format( "_moduleObj._moduleTypeId = %d", moduleTypeInfo:get_typeId()) )
   self:writeln( string.format( "_moduleObj._moduleSymbolKind = %d", moduleSymbolKind) )
   self:writeln( string.format( "_moduleObj._moduleMutable = %s", moduleTypeInfo:get_mutable()) )
   self:writeln( "----- meta -----" )
   if self.stream ~= self.metaStream then
      self:writeln( "return _moduleObj" )
   end
   
   self.outMetaFlag = false
end

function convFilter:processRoot( node, parent )

   self:writeln( string.format( "--%s", self.streamName) )
   self.needModuleObj = node:get_provideNode() == nil
   if self.needModuleObj then
      self:writeln( "local _moduleObj = {}" )
   end
   
   self:writeln( string.format( "local __mod__ = '%s'", node:get_moduleTypeInfo():getFullName( {} )) )
   if node:get_luneHelperInfo():get_useNilAccess() or node:get_luneHelperInfo():get_useUnwrapExp() or node:get_luneHelperInfo():get_hasMappingClassDef() then
      self:writeln( [==[
if not _ENV._lune then
   _lune = {}
end]==] )
      if node:get_luneHelperInfo():get_useNilAccess() then
         self:writeln( [==[
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
]==] )
      end
      
      if node:get_luneHelperInfo():get_useUnwrapExp() then
         self:writeln( [==[
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
]==] )
      end
      
   end
   
   if node:get_luneHelperInfo():get_hasMappingClassDef() then
      self:writeln( [==[      
function _lune._toStem( val )
   return val
end
function _lune._toInt( val )
   if type( val ) == "number" then
      return math.floor( val )
   end
   return nil
end
function _lune._toReal( val )
   if type( val ) == "number" then
      return val
   end
   return nil
end
function _lune._toBool( val )
   if type( val ) == "boolean" then
      return val
   end
   return nil
end
function _lune._toStr( val )
   if type( val ) == "string" then
      return val
   end
   return nil
end
function _lune._toList( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toValInfo = toValInfoList[ 1 ]
      for index, mem in ipairs( val ) do
         local memval, mess = toValInfo.func( mem, toValInfo.child )
         if memval == nil and not toValInfo.nilable then
            if mess then
              return nil, string.format( "%d.%s", index, mess )
            end
            return nil, index
         end
         tbl[ index ] = memval
      end
      return tbl
   end
   return nil
end
function _lune._toMap( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toKeyInfo = toValInfoList[ 1 ]
      local toValInfo = toValInfoList[ 2 ]
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal, valSub = toValInfo.func( mem, toValInfo.child )
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
function _lune._fromMap( obj, map, memInfoList )
   if type( map ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( map[ memInfo.name ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s.%s", memInfo.name, key) or memInfo.name
      end
      obj[ memInfo.name ] = val
   end
   return true
end
]==] )
   end
   
   local children = node:get_children(  )
   for __index, child in pairs( children ) do
      filter( child, self, node )
      self:writeln( "" )
   end
   
   self:outputMeta( node )
   do
      local _exp = node:get_provideNode()
      if _exp ~= nil then
         self:write( "return " )
         self:write( _exp:get_symbol():get_name() )
         self:writeln( "" )
      else
         self:writeln( "return _moduleObj" )
      end
   end
   
end


function convFilter:processSubfile( node, parent )

end

function convFilter:processBlock( node, parent )

   local word = ""
   do
      local _switchExp = node:get_blockKind(  )
      if _switchExp == Ast.BlockKind.If or _switchExp == Ast.BlockKind.Elseif then
         word = "then"
      elseif _switchExp == Ast.BlockKind.Else then
         word = ""
      elseif _switchExp == Ast.BlockKind.While then
         word = "do"
      elseif _switchExp == Ast.BlockKind.Repeat then
         word = ""
      elseif _switchExp == Ast.BlockKind.For then
         word = "do"
      elseif _switchExp == Ast.BlockKind.Apply then
         word = "do"
      elseif _switchExp == Ast.BlockKind.Foreach then
         word = "do"
      elseif _switchExp == Ast.BlockKind.Macro then
         word = ""
      elseif _switchExp == Ast.BlockKind.Func then
         word = ""
      elseif _switchExp == Ast.BlockKind.Default then
         word = ""
      elseif _switchExp == Ast.BlockKind.Block then
         word = "do"
      elseif _switchExp == Ast.BlockKind.Macro then
         word = ""
      elseif _switchExp == Ast.BlockKind.LetUnwrap then
         word = ""
      elseif _switchExp == Ast.BlockKind.IfUnwrap then
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
   if node:get_blockKind(  ) == Ast.BlockKind.Block then
      self:writeln( "end" )
   end
   
end


function convFilter:processStmtExp( node, parent )

   filter( node:get_exp(  ), self, node )
end


function convFilter:processDeclEnum( node, parent )

   local access = node:get_accessMode() == Ast.AccessMode.Global and "" or "local "
   local enumFullName = node:get_name().txt
   local typeInfo = node:get_expType()
   local parentInfo = typeInfo:get_parentInfo()
   local isTopNS = true
   if parentInfo ~= Ast.headTypeInfo and parentInfo:get_kind() == Ast.TypeInfoKind.Class then
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
   for __index, valName in pairs( node:get_valueNameList() ) do
      local valInfo = _lune.unwrap( typeInfo:getEnumValInfo( valName.txt ))
      local valTxt = string.format( "%s", valInfo:get_val())
      if typeInfo:get_valTypeInfo():equals( Ast.builtinTypeString ) then
         valTxt = string.format( "'%s'", valInfo:get_val())
      end
      
      self:writeln( string.format( "%s.%s = %s", enumFullName, valName.txt, valTxt) )
      self:writeln( string.format( "%s._val2NameMap[%s] = '%s'", enumFullName, valTxt, valName.txt) )
   end
   
end

function convFilter:getDestrClass( classTypeInfo )

   local typeInfo = classTypeInfo
   while not typeInfo:equals( Ast.headTypeInfo ) do
      local scope = _lune.unwrap( typeInfo:get_scope())
      do
         local _exp = scope:getTypeInfoChild( "__free" )
         if _exp ~= nil then
            return typeInfo
         end
      end
      
      typeInfo = typeInfo:get_baseTypeInfo()
   end
   
   return nil
end

function convFilter:processDeclClass( node, parent )

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
         self:write( string.format( "local %s = require( %s )", className, _exp.txt ) )
         do
            local _switchExp = node:get_accessMode()
            if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Pro then
               if self.needModuleObj then
                  self:writeln( "" )
                  self:write( string.format( "_moduleObj.%s = %s", className, className) )
               end
               
            end
         end
         
         return 
      end
   end
   
   self:writeln( string.format( "local %s = {}", className ) )
   local baseInfo = node:get_expType(  ):get_baseTypeInfo(  )
   if baseInfo:get_typeId(  ) ~= Ast.rootTypeId then
      self:writeln( string.format( "setmetatable( %s, { __index = %s } )", className, self:getFullName( _lune.unwrap( baseInfo) )) )
   end
   
   if nodeInfo:get_accessMode(  ) == Ast.AccessMode.Pub then
      if self.needModuleObj then
         self:writeln( string.format( "_moduleObj.%s = %s", className, className ) )
      end
      
   end
   
   for __index, declNode in pairs( node:get_declStmtList() ) do
      filter( declNode, self, node )
   end
   
   local hasConstrFlag = false
   local hasDestrFlag = false
   local memberList = {}
   local fieldList = nodeInfo:get_fieldList(  )
   local outerMethodSet = nodeInfo:get_outerMethodSet(  )
   local methodNameSet = {}
   for __index, field in pairs( fieldList ) do
      local ignoreFlag = false
      if field:get_kind() == Ast.nodeKind['DeclConstr'] then
         hasConstrFlag = true
         methodNameSet["__init"] = true
      end
      
      if field:get_kind() == Ast.nodeKind['DeclDestr'] then
         hasDestrFlag = true
         methodNameSet["__free"] = true
      end
      
      if field:get_kind() == Ast.nodeKind['DeclMember'] then
         local declMemberNode = field
         if not declMemberNode:get_staticFlag() then
            table.insert( memberList, declMemberNode )
         end
         
      end
      
      if field:get_kind() == Ast.nodeKind['DeclMethod'] then
         local methodNode = field
         local declInfo = methodNode:get_declInfo(  )
         local methodNameToken = _lune.unwrap( declInfo:get_name(  ))
         if outerMethodSet[methodNameToken.txt] then
            ignoreFlag = true
         end
         
         methodNameSet[methodNameToken.txt] = true
      end
      
      if (not ignoreFlag ) then
         filter( field, self, node )
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
      methodNameSet["__init"] = true
      local argTxt = ""
      for index, member in pairs( memberList ) do
         if index > 1 then
            argTxt = argTxt .. ", "
         end
         
         argTxt = argTxt .. member:get_name().txt
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
      for __index, member in pairs( memberList ) do
         local memberName = member:get_name().txt
         self:writeln( string.format( "self.%s = %s", memberName, memberName ) )
      end
      
      self:popIndent(  )
      self:writeln( 'end' )
   end
   
   local scope = nodeInfo:get_scope(  )
   for __index, memberNode in pairs( nodeInfo:get_memberList() ) do
      local memberNameToken = memberNode:get_name(  )
      local memberName = memberNameToken.txt
      local getterName = "get_" .. memberName
      local autoFlag = not methodNameSet[getterName]
      local prefix = memberNode:get_staticFlag() and className or "self"
      if memberNode:get_getterMode(  ) ~= Ast.AccessMode.None and autoFlag then
         self:writeln( string.format( [==[
function %s:%s()       
   return %s.%s         
end]==], className, getterName, prefix, memberName) )
         methodNameSet[getterName] = true
      end
      
      local setterName = "set_" .. memberName
      autoFlag = not methodNameSet[setterName]
      if memberNode:get_setterMode(  ) ~= Ast.AccessMode.None and autoFlag then
         self:writeln( string.format( [==[
function %s:%s( %s )   
   %s.%s = %s              
end]==], className, setterName, memberName, prefix, memberName, memberName) )
         methodNameSet[setterName] = true
      end
      
   end
   
   for __index, advertiseInfo in pairs( node:get_advertiseList() ) do
      local memberName = advertiseInfo:get_member():get_name().txt
      local memberType = advertiseInfo:get_member():get_expType()
      for __index, child in pairs( memberType:get_children() ) do
         if child:get_kind() == Ast.TypeInfoKind.Method and child:get_accessMode() ~= Ast.AccessMode.Pri and not child:get_staticFlag() then
            local childName = advertiseInfo:get_prefix() .. child:getTxt(  )
            if not methodNameSet[childName] then
               self:writeln( string.format( [==[
function %s:%s( ... )
   return self.%s:%s( ... )
end       
]==], className, childName, memberName, childName) )
            end
            
         end
         
      end
      
   end
   
   if #nodeInfo:get_initStmtList() > 0 then
      self:writeln( "do" )
      self:pushIndent(  )
      for __index, initStmt in pairs( nodeInfo:get_initStmtList() ) do
         filter( initStmt, self, node )
         self:writeln( "" )
      end
      
      self:popIndent(  )
      self:writeln( "end" )
   end
   
   if classTypeInfo:isInheritFrom( TransUnit.typeInfoMappingIF ) then
      self:writeln( string.format( [==[
function %s:_toMap()
  return self
end
function %s._fromMap( val )
  local obj, mes = %s._fromMapSub( {}, val )
  if obj then
     %s.setmeta( obj )
  end
  return obj, mes
end
function %s._fromStem( val )
  return %s._fromMap( val )
end
]==], classTypeInfo:getTxt(  ), classTypeInfo:getTxt(  ), classTypeInfo:getTxt(  ), classTypeInfo:getTxt(  ), classTypeInfo:getTxt(  ), classTypeInfo:getTxt(  )) )
      self:writeln( string.format( 'function %s._fromMapSub( obj, val )', classTypeInfo:getTxt(  )) )
      if classTypeInfo:get_baseTypeInfo() ~= Ast.headTypeInfo then
         self:writeln( string.format( [==[
   local result, mes = %s._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end
]==], self:getFullName( classTypeInfo:get_baseTypeInfo() )) )
      end
      
      self:writeln( '   local memInfo = {}' )
      local function getMapInfo( typeInfo )
      
         local orgTypeInfo = typeInfo:get_srcTypeInfo()
         if typeInfo:get_nilable() then
            orgTypeInfo = typeInfo:get_orgTypeInfo()
         end
         
         local child = "{}"
         local funcTxt = ""
         do
            local _switchExp = orgTypeInfo:get_kind()
            if _switchExp == Ast.TypeInfoKind.Stem then
               funcTxt = '_lune._toStem'
            elseif _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
               if not orgTypeInfo:equals( Ast.builtinTypeString ) then
                  funcTxt = string.format( '%s._fromMap', self:getFullName( orgTypeInfo ))
               else
                
                  funcTxt = '_lune._toStr'
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Enum then
               funcTxt = string.format( '%s._from', self:getFullName( orgTypeInfo ))
            elseif _switchExp == Ast.TypeInfoKind.Prim then
               do
                  local _switchExp = orgTypeInfo
                  if _switchExp == Ast.builtinTypeInt then
                     funcTxt = '_lune._toInt'
                  elseif _switchExp == Ast.builtinTypeReal then
                     funcTxt = '_lune._toReal'
                  elseif _switchExp == Ast.builtinTypeBool then
                     funcTxt = '_lune._toBool'
                  else 
                     
                        Util.err( string.format( "unknown type -- %s", orgTypeInfo:getTxt(  )) )
                  end
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Map then
               funcTxt = '_lune._toMap'
               local itemList = orgTypeInfo:get_itemTypeInfoList()
               local keyFuncTxt, keyNilable, keyChild = getMapInfo( itemList[1] )
               local valFuncTxt, valNilable, valChild = getMapInfo( itemList[2] )
               child = string.format( "{ { func = %s, nilable = %s, child = %s }, \n", keyFuncTxt, keyNilable, keyChild) .. string.format( "{ func = %s, nilable = %s, child = %s } }", valFuncTxt, valNilable, valChild)
            elseif _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
               funcTxt = '_lune._toList'
               local itemList = orgTypeInfo:get_itemTypeInfoList()
               local valFuncTxt, valNilable, valChild = getMapInfo( itemList[1] )
               child = string.format( "{ { func = %s, nilable = %s, child = %s } }", valFuncTxt, valNilable, valChild)
            end
         end
         
         return funcTxt, typeInfo:get_nilable(), child
      end
      
      for __index, memberNode in pairs( node:get_memberList() ) do
         local funcTxt, nilable, child = getMapInfo( memberNode:get_expType() )
         self:writeln( string.format( '   table.insert( memInfo, { name = "%s", func = %s, nilable = %s, child = %s } )', memberNode:get_name().txt, funcTxt, nilable, child) )
      end
      
      self:writeln( string.format( [==[
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end]==], classTypeInfo:getTxt(  )) )
   end
   
end


function convFilter:processDeclMember( node, parent )

end


function convFilter:processExpMacroExp( node, parent )

   local stmtList = node:get_stmtList(  )
   if stmtList then
      for __index, stmt in pairs( stmtList ) do
         filter( stmt, self, node )
         self:writeln( "" )
      end
      
   end
   
end


function convFilter:processDeclMacro( node, parent )

   if self.inMacro then
      local nodeInfo = node:get_declInfo(  )
      local name = nodeInfo:get_name(  )
      self:write( string.format( "local function %s(", name.txt) )
      self:writeln( "__macroArgs )" )
      self:pushIndent(  )
      for index, arg in pairs( nodeInfo:get_argList(  ) ) do
         local argName = arg:get_name().txt
         self:writeln( string.format( "local %s = __macroArgs.%s", argName, argName) )
      end
      
      self:writeln( "local macroVar = {}" )
      self:writeln( "macroVar._names = {}" )
      self.macroDepth = self.macroDepth + 1
      do
         local _exp = nodeInfo:get_ast(  )
         if _exp ~= nil then
            filter( _exp, self, node )
         end
      end
      
      self.macroDepth = self.macroDepth - 1
      self:writeln( "" )
      self:writeln( "return macroVar" )
      self:popIndent(  )
      self:writeln( "end" )
      self:writeln( string.format( "return %s", name.txt) )
   end
   
end


function convFilter:processExpMacroStat( node, parent )

   for index, token in pairs( node:get_expStrList(  ) ) do
      if index ~= 1 then
         self:write( '..' )
      end
      
      filter( token, self, node )
   end
   
end


function convFilter:processExpNew( node, parent )

   filter( node:get_symbol(  ), self, node )
   self:write( ".new(" )
   do
      local _exp = node:get_argList(  )
      if _exp ~= nil then
         filter( _exp, self, node )
      end
   end
   
   self:write( ")" )
end


function convFilter:process__func__symbol( has__func__Symbol, classType, funcName )

   if has__func__Symbol then
      local nameSpace = ""
      if classType ~= Ast.headTypeInfo then
         nameSpace = self:getFullName( classType )
      end
      
      if funcName == "" then
         funcName = "<anonymous>"
      end
      
      self:writeln( string.format( "local __func__ = '%s.%s'", nameSpace, funcName) )
   end
   
end

function convFilter:processDeclConstr( node, parent )

   local declInfo = node:get_declInfo(  )
   local classTypeInfo = _lune.unwrap( declInfo:get_classTypeInfo())
   local className = self:getFullName( classTypeInfo )
   self:write( string.format( "function %s.new( ", className ) )
   local argTxt = ""
   local argList = declInfo:get_argList(  )
   for index, arg in pairs( argList ) do
      if index > 1 then
         self:write( ", " )
         argTxt = argTxt .. ", "
      end
      
      filter( arg, self, node )
      if arg:get_kind(  ) == Ast.nodeKind['DeclArg'] then
         argTxt = argTxt .. (arg ):get_name().txt
      else
       
         local name = _lune.unwrap( node:get_declInfo(  ):get_name())
         Util.err( string.format( "not support ... in macro -- %s", name.txt) )
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
   self:write( string.format( "function %s:__init(%s) ", className, argTxt ) )
   do
      local _exp = declInfo:get_body()
      if _exp ~= nil then
         self:process__func__symbol( declInfo:get_has__func__Symbol(), node:get_expType():get_parentInfo(), "__init" )
         filter( _exp, self, node )
      end
   end
   
   self:writeln( "end" )
end


function convFilter:processDeclDestr( node, parent )

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

function convFilter:processExpCallSuper( node, parent )

   local typeInfo = node:get_superType(  )
   self:write( string.format( "%s.__init( self ", self:getFullName( typeInfo )) )
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         self:write( "," )
         filter( _exp, self, node )
      end
   end
   
   self:writeln( ")" )
end


function convFilter:processDeclMethod( node, parent )

   local declInfo = node:get_declInfo(  )
   local delimit = ":"
   if declInfo:get_staticFlag(  ) then
      delimit = "."
   end
   
   local methodNodeToken = _lune.unwrap( declInfo:get_name(  ))
   local methodName = methodNodeToken.txt
   local classTypeInfo = _lune.unwrap( declInfo:get_classTypeInfo())
   self:write( string.format( "function %s%s%s( ", self:getFullName( classTypeInfo ), delimit, methodName) )
   local argList = declInfo:get_argList(  )
   for index, arg in pairs( argList ) do
      if index > 1 then
         self:write( ", " )
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


function convFilter:processUnwrapSet( node, parent )

   local dstExpList = node:get_dstExpList()
   filter( dstExpList, self, node )
   self:write( " = " )
   filter( node:get_srcExpList(), self, node )
   self:writeln( "" )
   self:write( "if " )
   for index, expNode in pairs( dstExpList:get_expList() ) do
      if index > 1 then
         self:write( " or " )
      end
      
      self:write( "nil == " )
      filter( expNode, self, node )
   end
   
   self:writeln( " then" )
   self:pushIndent(  )
   for index, expNode in pairs( dstExpList:get_expList() ) do
      self:write( string.format( "local _exp%d = ", index) )
      filter( expNode, self, node )
      self:writeln( "" )
   end
   
   if node:get_unwrapBlock() then
      filter( _lune.unwrap( node:get_unwrapBlock()), self, node )
   end
   
   self:popIndent(  )
   self:writeln( "end" )
end

function convFilter:processIfUnwrap( node, parent )

   self:writeln( "do" )
   self:pushIndent(  )
   self:write( "local " )
   for index, varName in pairs( node:get_varNameList() ) do
      self:write( varName )
      if index ~= #node:get_varNameList() then
         self:write( ", " )
      end
      
   end
   
   self:write( " = " )
   for index, expNode in pairs( node:get_expNodeList() ) do
      filter( expNode, self, node )
      if index ~= #node:get_expNodeList() then
         self:write( ", " )
      end
      
   end
   
   self:writeln( "" )
   self:write( "if " )
   for index, varName in pairs( node:get_varNameList() ) do
      self:write( string.format( "%s ~= nil", varName) )
      if index ~= #node:get_varNameList() then
         self:write( " and " )
      end
      
   end
   
   self:write( " then" )
   filter( node:get_block(), self, node )
   if node:get_nilBlock() then
      self:write( "else" )
      filter( _lune.unwrap( node:get_nilBlock()), self, node )
   end
   
   self:writeln( "end" )
   self:popIndent(  )
   self:writeln( "end" )
end

function convFilter:processDeclVar( node, parent )

   if node:get_syncBlock() then
      self:writeln( "do" )
      self:pushIndent(  )
      for __index, varInfo in pairs( node:get_syncVarList() ) do
         self:writeln( string.format( "local _sync_%s", varInfo:get_name().txt) )
      end
      
      self:writeln( "do" )
      self:pushIndent(  )
   end
   
   if node:get_mode() ~= Ast.DeclVarMode.Unwrap and node:get_accessMode(  ) ~= Ast.AccessMode.Global then
      self:write( "local " )
   end
   
   local varList = node:get_varList(  )
   for index, var in pairs( varList ) do
      if index > 1 then
         self:write( ", " )
      end
      
      self:write( var:get_name().txt )
   end
   
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         self:write( " = " )
         filter( _exp, self, node )
      else
         self:writeln( "" )
      end
   end
   
   do
      local _exp = node:get_unwrapBlock()
      if _exp ~= nil then
         self:writeln( "" )
         self:write( "if " )
         for index, var in pairs( varList ) do
            if index > 1 then
               self:write( " or " )
            end
            
            self:write( " nil == " .. var:get_name().txt )
         end
         
         self:writeln( " then" )
         self:pushIndent(  )
         for index, var in pairs( varList ) do
            self:writeln( string.format( "local _%s = %s", var:get_name().txt, var:get_name().txt) )
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
         for __index, varInfo in pairs( node:get_syncVarList() ) do
            self:writeln( string.format( "_sync_%s = %s", varInfo:get_name().txt, varInfo:get_name().txt) )
         end
         
         self:popIndent(  )
         self:writeln( "end" )
         for __index, varInfo in pairs( node:get_syncVarList() ) do
            self:writeln( string.format( "%s = _sync_%s", varInfo:get_name().txt, varInfo:get_name().txt) )
         end
         
         self:popIndent(  )
         self:writeln( "end" )
      end
   end
   
   if node:get_accessMode(  ) == Ast.AccessMode.Pub then
      self:writeln( "" )
      for index, var in pairs( varList ) do
         local name = var:get_name().txt
         if self.needModuleObj then
            self:writeln( string.format( "_moduleObj.%s = %s", name, name) )
         end
         
         self.pubVarName2InfoMap[name] = PubVerInfo.new(node:get_staticFlag(), node:get_accessMode(), node:get_symbolInfoList()[index]:get_mutable(), node:get_typeInfoList()[index])
      end
      
   end
   
   if self.macroDepth > 0 then
      self:writeln( "" )
      for index, var in pairs( varList ) do
         local varName = var:get_name().txt
         self:writeln( string.format( "table.insert( macroVar._names, '%s' )", varName) )
         self:writeln( string.format( "macroVar.%s = %s", varName, varName) )
      end
      
   end
   
end


function convFilter:processDeclArg( node, parent )

   self:write( string.format( "%s", node:get_name(  ).txt ) )
end


function convFilter:processDeclArgDDD( node, parent )

   self:write( "..." )
end


function convFilter:processExpDDD( node, parent )

   self:write( "..." )
end


function convFilter:processDeclFunc( node, parent )

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
   
   self:write( string.format( "%sfunction %s( ", letTxt, name ) )
   local argList = declInfo:get_argList(  )
   for index, arg in pairs( argList ) do
      if index > 1 then
         self:write( ", " )
      end
      
      filter( arg, self, node )
   end
   
   self:writeln( " )" )
   do
      local _exp = declInfo:get_body()
      if _exp ~= nil then
         self:process__func__symbol( declInfo:get_has__func__Symbol(), node:get_expType():get_parentInfo(), name )
         filter( _exp, self, node )
      end
   end
   
   self:writeln( "end" )
   local expType = node:get_expType(  )
   if expType:get_accessMode(  ) == Ast.AccessMode.Pub then
      if self.needModuleObj then
         self:write( string.format( "_moduleObj.%s = %s", name, name) )
      end
      
      self.pubFuncName2InfoMap[name] = PubFuncInfo.new(declInfo:get_accessMode(  ), node:get_expType(  ))
   end
   
end


function convFilter:processRefType( node, parent )

   self:write( (node:get_refFlag(  ) and "&" or "" ) .. (node:get_mutFlag(  ) and "mut " or "" ) )
   filter( node:get_name(  ), self, node )
   if node:get_array(  ) == "array" then
      self:write( "[@]" )
   elseif node:get_array(  ) == "list" then
      self:write( "[]" )
   end
   
end


function convFilter:processIf( node, parent )

   local valList = node:get_stmtList(  )
   for index, val in pairs( valList ) do
      if index == 1 then
         self:write( "if " )
         filter( val:get_exp(), self, node )
      elseif val:get_kind() == "elseif" then
         self:write( "elseif " )
         filter( val:get_exp(), self, node )
      else
       
         self:writeln( "else" )
      end
      
      self:write( " " )
      filter( val:get_block(), self, node )
   end
   
   self:writeln( "end" )
end


function convFilter:processSwitch( node, parent )

   self:writeln( "do" )
   self:pushIndent(  )
   self:write( "local _switchExp = " )
   filter( node:get_exp(  ), self, node )
   self:writeln( "" )
   for index, caseInfo in pairs( node:get_caseList(  ) ) do
      if index == 1 then
         self:write( "if " )
      else
       
         self:write( "elseif " )
      end
      
      local expList = caseInfo:get_expList(  )
      for listIndex, expNode in pairs( expList:get_expList(  ) ) do
         if listIndex ~= 1 then
            self:write( " or " )
         end
         
         self:write( "_switchExp == " )
         filter( expNode, self, node )
      end
      
      self:write( " then" )
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
   self:popIndent(  )
   self:writeln( "end" )
end


function convFilter:processWhile( node, parent )

   self:write( "while " )
   filter( node:get_exp(  ), self, node )
   self:write( " " )
   filter( node:get_block(  ), self, node )
   self:writeln( "end" )
end


function convFilter:processRepeat( node, parent )

   self:write( "repeat " )
   filter( node:get_block(  ), self, node )
   self:write( "until " )
   filter( node:get_exp(  ), self, node )
end


function convFilter:processFor( node, parent )

   self:write( string.format( "for %s = ", node:get_val(  ).txt ) )
   filter( node:get_init(  ), self, node )
   self:write( ", " )
   filter( node:get_to(  ), self, node )
   do
      local _exp = node:get_delta(  )
      if _exp ~= nil then
         self:write( ", " )
         filter( _exp, self, node )
      end
   end
   
   self:write( " " )
   filter( node:get_block(  ), self, node )
   self:writeln( "end" )
end


function convFilter:processApply( node, parent )

   self:write( "for " )
   local varList = node:get_varList(  )
   for index, var in pairs( varList ) do
      if index > 1 then
         self:write( ", " )
      end
      
      self:write( var.txt )
   end
   
   self:write( " in " )
   filter( node:get_exp(), self, node )
   self:write( " " )
   filter( node:get_block(), self, node )
   self:writeln( "end" )
end


function convFilter:processForeach( node, parent )

   self:write( "for " )
   do
      local _exp = node:get_key()
      if _exp ~= nil then
         self:write( _exp.txt )
      else
         self:write( "__index" )
      end
   end
   
   self:write( ", " )
   self:write( node:get_val().txt )
   self:write( " in pairs( " )
   filter( node:get_exp(), self, node )
   self:write( " ) " )
   filter( node:get_block(), self, node )
   self:writeln( "end" )
end


function convFilter:processForsort( node, parent )

   self:writeln( "do" )
   self:pushIndent(  )
   self:writeln( "local __sorted = {}" )
   self:write( "local __map = " )
   filter( node:get_exp(), self, node )
   self:writeln( "" )
   self:writeln( "for __key in pairs( __map ) do" )
   self:pushIndent(  )
   self:writeln( "table.insert( __sorted, __key )" )
   self:popIndent(  )
   self:writeln( "end" )
   self:writeln( "table.sort( __sorted )" )
   self:write( "for __index, " )
   local key = "__key"
   do
      local _exp = node:get_key()
      if _exp ~= nil then
         key = _exp.txt
      end
   end
   
   self:write( key )
   self:writeln( " in ipairs( __sorted ) do" )
   self:pushIndent(  )
   self:writeln( string.format( "%s = __map[ %s ]", node:get_val().txt, key ) )
   filter( node:get_block(), self, node )
   self:writeln( "end" )
   self:popIndent(  )
   self:writeln( "end" )
   self:popIndent(  )
   self:writeln( "end" )
end


function convFilter:processExpUnwrap( node, parent )

   do
      local _exp = node:get_default()
      if _exp ~= nil then
         self:write( '_lune.unwrapDefault( ' )
         filter( node:get_exp(), self, node )
         self:write( ', ' )
         filter( _exp, self, node )
         self:write( ')' )
      else
         self:write( '_lune.unwrap( ' )
         filter( node:get_exp(), self, node )
         self:write( ')' )
      end
   end
   
end

function convFilter:processExpCall( node, parent )

   local wroteFuncFlag = false
   local setArgFlag = false
   
   if node:get_func():get_kind() == Ast.nodeKind['RefField'] then
      local fieldNode = node:get_func()
      local prefixNode = fieldNode:get_prefix()
      local prefixType = prefixNode:get_expType()
      if node:get_nilAccess() then
         wroteFuncFlag = true
         setArgFlag = true
         if prefixType:get_kind() == Ast.TypeInfoKind.List then
            self:write( string.format( "_lune.nilacc( table.%s, nil, 'list', ", fieldNode:get_field().txt) )
            filter( prefixNode, self, fieldNode )
         else
          
            self:write( "_lune.nilacc( " )
            filter( prefixNode, self, fieldNode )
            self:write( string.format( ", '%s', 'callmtd' ", fieldNode:get_field().txt) )
         end
         
      else
       
         if prefixType:get_kind() == Ast.TypeInfoKind.List then
            setArgFlag = true
            wroteFuncFlag = true
            self:write( string.format( "table.%s( ", fieldNode:get_field().txt) )
            filter( prefixNode, self, fieldNode )
         elseif prefixType:get_kind() == Ast.TypeInfoKind.Enum then
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
   
   do
      local _exp = node:get_argList()
      if _exp ~= nil then
         if wroteFuncFlag and setArgFlag then
            if #_exp:get_expList() > 0 then
               self:write( ", " )
            end
            
         end
         
         filter( _exp, self, node )
      end
   end
   
   self:write( " )" )
end


function convFilter:processExpList( node, parent )

   local expList = node:get_expList(  )
   for index, exp in pairs( expList ) do
      if index > 1 then
         self:write( ", " )
      end
      
      filter( exp, self, node )
   end
   
end


function convFilter:processExpOp1( node, parent )

   local op = node:get_op().txt
   if op == ",,," then
      filter( node:get_exp(), self, node )
   elseif op == ",,,," then
      if node:get_macroMode() == Ast.MacroMode.Expand then
         filter( node:get_exp(), self, node )
      else
       
         self:write( "_luneSym2Str( " )
         filter( node:get_exp(), self, node )
         self:write( " )" )
      end
      
   elseif op == ",," then
      self:write( "_luneGetLocal( " )
      filter( node:get_exp(), self, node )
      self:write( " )" )
   elseif op == "~" then
      if hasBitOpFlag then
         self:write( op )
         filter( node:get_exp(), self, node )
      else
       
         self:write( "bit32.bnot( " )
         filter( node:get_exp(), self, node )
         self:write( " )" )
      end
      
   else
    
      if op == "not" then
         op = op .. " "
      end
      
      self:write( op )
      filter( node:get_exp(), self, node )
   end
   
end


function convFilter:processExpCast( node, parent )

   if node:get_expType():equals( Ast.builtinTypeInt ) then
      self:write( "math.floor(" )
      filter( node:get_exp(), self, node )
      self:write( ")" )
   else
    
      filter( node:get_exp(), self, node )
   end
   
end


function convFilter:processExpParen( node, parent )

   self:write( "(" )
   filter( node:get_exp(), self, node )
   self:write( " )" )
end


function convFilter:processExpOp2( node, parent )

   local intCast = false
   if node:get_expType():equals( Ast.builtinTypeInt ) and node:get_op().txt == "/" then
      intCast = true
      self:write( "math.floor(" )
   end
   
   local opTxt = node:get_op().txt
   do
      local _exp = bitBinOpMap[opTxt]
      if _exp ~= nil then
         if hasBitOpFlag then
            do
               local _switchExp = _exp
               if _switchExp == BitOpKind.LShift then
                  opTxt = "<<"
               elseif _switchExp == BitOpKind.RShift then
                  opTxt = ">>"
               end
            end
            
            filter( node:get_exp1(), self, node )
            self:write( " " .. opTxt .. " " )
            filter( node:get_exp2(), self, node )
         else
          
            local binfunc = ""
            local exp2Mod = ""
            do
               local _switchExp = _exp
               if _switchExp == BitOpKind.And then
                  binfunc = "band"
               elseif _switchExp == BitOpKind.Or then
                  binfunc = "bor"
               elseif _switchExp == BitOpKind.Xor then
                  binfunc = "bxor"
               elseif _switchExp == BitOpKind.LShift then
                  binfunc = "lshift"
               elseif _switchExp == BitOpKind.RShift then
                  binfunc = "lshift"
                  exp2Mod = "-"
               else 
                  
                     Util.err( string.format( "illegal op -- %s", opTxt) )
               end
            end
            
            self:write( string.format( "bit32.%s(", binfunc) )
            filter( node:get_exp1(), self, node )
            self:write( ", " )
            self:write( exp2Mod )
            filter( node:get_exp2(), self, node )
            self:write( " )" )
         end
         
      else
         filter( node:get_exp1(), self, node )
         self:write( " " .. opTxt .. " " )
         filter( node:get_exp2(), self, node )
      end
   end
   
   if intCast then
      self:write( ")" )
   end
   
end


function convFilter:processExpRef( node, parent )

   if node:get_symbolInfo():get_accessMode() == Ast.AccessMode.Pub and node:get_symbolInfo():get_kind() == Ast.SymbolKind.Var then
      if self.needModuleObj then
         self:write( "_moduleObj." )
      end
      
   end
   
   self:write( node:get_token().txt )
end


function convFilter:processExpRefItem( node, parent )

   if node:get_nilAccess() then
      self:write( "_lune.nilacc( " )
      filter( node:get_val(), self, node )
      self:write( ", nil, 'item', " )
      do
         local _exp = node:get_index()
         if _exp ~= nil then
            filter( _exp, self, node )
         else
            self:write( string.format( "'%s'", _lune.unwrap( node:get_symbol())) )
         end
      end
      
      self:write( ")" )
   else
    
      if node:get_val():get_expType():equals( Ast.builtinTypeString ) then
         self:write( "string.byte( " )
         filter( node:get_val(), self, node )
         self:write( ", " )
         do
            local _exp = node:get_index()
            if _exp ~= nil then
               filter( _exp, self, node )
            else
               error( "index is nil" )
            end
         end
         
         self:write( " )" )
      else
       
         filter( node:get_val(), self, node )
         self:write( "[" )
         do
            local _exp = node:get_index()
            if _exp ~= nil then
               filter( _exp, self, node )
            else
               self:write( string.format( "'%s'", _lune.unwrap( node:get_symbol())) )
            end
         end
         
         self:write( "]" )
      end
      
   end
   
end


function convFilter:processRefField( node, parent )

   local prefix = node:get_prefix(  )
   if node:get_nilAccess() then
      self:write( '_lune.nilacc( ' )
      filter( prefix, self, node )
      self:write( string.format( ', "%s" )', node:get_field().txt) )
   else
    
      filter( prefix, self, node )
      local delimit = "."
      if parent:get_kind() == Ast.nodeKind['ExpCall'] then
         if node:get_expType(  ):get_kind(  ) == Ast.TypeInfoKind.Method then
            delimit = ":"
         else
          
            delimit = "."
         end
         
      end
      
      local fieldToken = node:get_field(  )
      self:write( delimit .. fieldToken.txt )
   end
   
end


function convFilter:processExpOmitEnum( node, parent )

   self:write( string.format( "%s.%s", self:getFullName( node:get_expType() ), node:get_valToken().txt) )
end


function convFilter:processGetField( node, parent )

   local prefixNode = node:get_prefix(  )
   local prefixType = prefixNode:get_expType()
   local fieldTxt = node:get_field(  ).txt
   if fieldTxt == "_txt" and prefixType:get_kind() == Ast.TypeInfoKind.Enum then
      self:write( string.format( "%s:_getTxt( ", self:getFullName( prefixType )) )
      filter( prefixNode, self, node )
      self:writeln( ")" )
   else
    
      filter( prefixNode, self, node )
      local delimit = "."
      if node:get_getterTypeInfo(  ):get_kind(  ) == Ast.TypeInfoKind.Method then
         delimit = ":"
      else
       
         delimit = "."
      end
      
      if node:get_getterTypeInfo(  ) then
         fieldTxt = string.format( "get_%s()", fieldTxt)
      end
      
      self:write( delimit .. fieldTxt )
   end
   
end


function convFilter:processReturn( node, parent )

   self:write( "return " )
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         filter( _exp, self, node )
      end
   end
   
end


function convFilter:processProvide( node, parent )

end

function convFilter:processLiteralList( node, parent )

   self:write( "{" )
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         filter( _exp, self, node )
      end
   end
   
   self:write( "}" )
end


function convFilter:processLiteralMap( node, parent )

   self:write( "{" )
   local pairList = node:get_pairList()
   for index, pair in pairs( pairList ) do
      if index > 1 then
         self:write( ", " )
      end
      
      self:write( "[" )
      filter( pair:get_key(), self, node )
      self:write( "] = " )
      filter( pair:get_val(), self, node )
   end
   
   self:write( "}" )
end


function convFilter:processLiteralArray( node, parent )

   self:write( "{" )
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         filter( _exp, self, node )
      end
   end
   
   self:write( "}" )
end


function convFilter:processLiteralChar( node, parent )

   self:write( string.format( "%g", node:get_num() ) )
end


function convFilter:processLiteralInt( node, parent )

   self:write( node:get_token().txt )
end


function convFilter:processLiteralReal( node, parent )

   self:write( node:get_token().txt )
end


function convFilter:processLiteralString( node, parent )

   local txt = node:get_token(  ).txt
   if string.find( txt, '^```' ) then
      txt = '[==[' .. txt:sub( 4, -4 ) .. ']==]'
   end
   
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
   
end


function convFilter:processLiteralBool( node, parent )

   self:write( node:get_token().txt )
end


function convFilter:processLiteralNil( node, parent )

   self:write( "nil" )
end


function convFilter:processBreak( node, parent )

   self:write( "break" )
end


function convFilter:processLiteralSymbol( node, parent )

   self:write( string.format( '%s', node:get_token().txt) )
end


local function createFilter( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, moduleSymbolKind )

   return convFilter.new(streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, moduleSymbolKind)
end
_moduleObj.createFilter = createFilter
local MacroEvalImp = {}
setmetatable( MacroEvalImp, { __index = Ast.MacroEval } )
_moduleObj.MacroEvalImp = MacroEvalImp
function MacroEvalImp:eval( node )

   local oStream = Util.memStream.new()
   local conv = convFilter.new("macro", oStream, oStream, ConvMode.Exec, true, Ast.headTypeInfo, Ast.SymbolKind.Typ)
   conv:processDeclMacro( node, node )
   local newEnv = {}
   for key, val in pairs( _ENV ) do
      newEnv[key] = val
   end
   
   newEnv["_lnsLoad"] = function ( name, txt )
   
      local metaInfo, val = frontInterface.loadFromLnsTxt( name, txt, false )
      return val
   end
   
   local chunk, err = load( oStream:get_txt(  ), "", "bt", newEnv )
   if err then
      Util.err( err )
   end
   
   do
      local _exp = chunk
      if _exp ~= nil then
         local mod = _exp(  )
         if not mod then
            Util.err( "macro load error" )
         end
         
         return (_lune.unwrap( mod) )
      end
   end
   
   Util.err( "failed to load -- " .. node:get_declInfo():get_name() )
end
function MacroEvalImp.setmeta( obj )
  setmetatable( obj, { __index = MacroEvalImp  } )
end
function MacroEvalImp.new( mode )
   local obj = {}
   MacroEvalImp.setmeta( obj )
   if obj.__init then
      obj:__init( mode )
   end        
   return obj 
end         
function MacroEvalImp:__init( mode ) 

   self.mode = mode
end

return _moduleObj
