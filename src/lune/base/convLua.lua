--lune/base/convLua.lns
local _moduleObj = {}
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
do
  end

-- none

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
do
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
function ConvMode:_from( val )
  if self._val2NameMap[ val ] then
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

local convFilter = {}
setmetatable( convFilter, { __index = Filter } )
function convFilter.new( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo )
  local obj = {}
  convFilter.setmeta( obj )
  if obj.__init then obj:__init( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo ); end
return obj
end
function convFilter:__init(streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo) 
  self.macroDepth = 0
  self.streamName = streamName
  self.stream = stream
  self.metaStream = metaStream
  self.outMetaFlag = false
  self.typeInfo2ModuleName = {}
  self.convMode = convMode
  self.inMacro = inMacro
  self.indent = 0
  self.curLineNo = 1
  self.classId2TypeInfo = {}
  self.classId2MemberList = {}
  self.pubVarName2InfoMap = {}
  self.pubFuncName2InfoMap = {}
  self.pubEnumId2EnumTypeInfo = {}
  self.needIndent = false
  self.moduleTypeInfo = moduleTypeInfo
end
function convFilter:getCanonicalName( typeInfo )
  local canonicalName = typeInfo:get_rawTxt()
  
  if self.moduleTypeInfo == typeInfo:getModule(  ) then
    return canonicalName
  end
  local workType = typeInfo:get_parentInfo()
  
  while workType ~= Ast.rootTypeInfo do
    canonicalName = string.format( "%s.%s", workType:get_rawTxt(), canonicalName)
    if self.typeInfo2ModuleName[workType] then
      break
    end
    workType = workType:get_parentInfo()
  end
  return canonicalName
end
function convFilter:close(  )
end
function convFilter:write( txt )
  local stream = self.stream
  
  if self.outMetaFlag then
    stream = self.metaStream
  end
  if self.needIndent then
    stream:write( string.rep( " ", self.indent ) )
    self.needIndent = false
  end
  for cr in string.gmatch( txt, "\n" ) do
    self.curLineNo = self.curLineNo + 1
  end
  stream:write( txt )
end
function convFilter.setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end
do
  end

local function filter( node, filter, parent, baseIndent )
  node:processFilter( filter, parent, baseIndent )
end

local stepIndent = 2

local builtInModuleSet = {}

builtInModuleSet["io"] = true
builtInModuleSet["string"] = true
builtInModuleSet["table"] = true
builtInModuleSet["math"] = true
builtInModuleSet["debug"] = true
function convFilter:setIndent( indent )
  self.indent = indent
end

function convFilter:writeln( txt, baseIndent )
  self:write( txt )
  self:write( "\n" )
  self.needIndent = true
  self.indent = baseIndent
end

function convFilter:processNone( node, parent, baseIndent )
  self:writeln( "-- none", baseIndent )
end

-- none

function convFilter:processImport( node, parent, baseIndent )
  local module = node:get_modulePath(  )
  
  local moduleName = string.gsub( module, ".*%.", "" )
  
  self.typeInfo2ModuleName[node:get_moduleTypeInfo()] = module
  if self.convMode == ConvMode.Exec then
    self:writeln( string.format( "local %s = __luneScript:loadModule( '%s' )", moduleName, module), baseIndent )
  else 
    self:writeln( string.format( "local %s = require( '%s' )", moduleName, module), baseIndent )
  end
end

-- none

function convFilter:outputMeta( node, baseIndent )
  if self.convMode == ConvMode.Convert then
    return 
  end
  self.outMetaFlag = true
  if self.stream ~= self.metaStream then
    self:writeln( "local _moduleObj = {}", baseIndent )
  end
  self:writeln( "----- meta -----", baseIndent )
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
    
    return typeId2TypeInfo[typeId] and not Ast.isBuiltin( typeId ) and (moduleTypeInfo:hasRouteNamespaceFrom( node:get_moduleTypeInfo() ) or typeInfo:get_srcTypeInfo() ~= typeInfo or moduleTypeInfo:equals( Ast.rootTypeInfo ) )
  end
  
  local function pickupTypeId( typeInfo, forceFlag, pickupChildFlag )
    if typeInfo:get_typeId(  ) == Ast.rootTypeId then
      return 
    end
    if not forceFlag and typeInfo:get_accessMode(  ) ~= Ast.AccessMode.Pub then
      return 
    end
    if typeId2TypeInfo[typeInfo:get_typeId(  )] then
      if pickupChildFlag and not typeInfo:get_nilable() then
        for __index, itemTypeInfo in pairs( typeInfo:get_children(  ) ) do
          if itemTypeInfo:get_accessMode() == Ast.AccessMode.Pub and (itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Class or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.IF or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Func or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Method ) then
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
    
    while typeInfo ~= Ast.rootTypeInfo do
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
  self:writeln( "local _typeId2ClassInfoMap = {}", baseIndent )
  self:writeln( "_moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap", baseIndent )
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
          self:writeln( "do", baseIndent + stepIndent )
          self:writeln( string.format( "local _classInfo%d = {}", classTypeId), baseIndent + stepIndent )
          self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId), baseIndent + stepIndent )
          for __index, memberNode in pairs( _lune.unwrap( self.classId2MemberList[classTypeId]) ) do
            if memberNode:get_accessMode() ~= Ast.AccessMode.Pri then
              local memberName = memberNode:get_name().txt
              
              local memberTypeInfo = memberNode:get_expType(  )
              
              self:writeln( string.format( "_classInfo%d.%s = {", classTypeId, memberName), baseIndent + stepIndent )
              self:writeln( string.format( "  name='%s', staticFlag = %s, ", memberName, memberNode:get_staticFlag()) .. string.format( "accessMode = '%s', typeId = %d }", memberNode:get_accessMode(), memberTypeInfo:get_typeId(  )), baseIndent + stepIndent )
              pickupTypeId( memberTypeInfo, true )
            end
          end
          self:writeln( "end", baseIndent )
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
              
              self:writeln( "do", baseIndent + stepIndent )
              self:writeln( string.format( "local _classInfo%s = {}", classTypeId), baseIndent + stepIndent )
              self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId), baseIndent + stepIndent )
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
                        self:writeln( string.format( "_classInfo%d.%s = {", classTypeId, fieldName), baseIndent + stepIndent )
                        self:writeln( string.format( "  name='%s', staticFlag = %s, ", fieldName, symbolInfo:get_staticFlag()) .. string.format( "accessMode = %d, typeId = %d }", symbolInfo:get_accessMode(), typeInfo:get_typeId(  )), baseIndent + stepIndent )
                        pickupTypeId( typeInfo )
                      end
                    end
                  end
                end
              end
              
              self:writeln( "end", baseIndent )
            end
          end
        end
      end
    end
    
  end
  self:writeln( "local _varName2InfoMap = {}", baseIndent )
  self:writeln( "_moduleObj._varName2InfoMap = _varName2InfoMap", baseIndent )
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
        self:writeln( string.format( "_varName2InfoMap.%s = {", varName ), baseIndent )
        self:writeln( string.format( "  name='%s', accessMode = %d, typeId = %d, mutable = %s }", varName, varInfo.accessMode, varInfo.typeInfo:get_typeId(), true), baseIndent )
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
  
  self:writeln( "local _typeInfoList = {}", baseIndent )
  self:writeln( "_moduleObj._typeInfoList = _typeInfoList", baseIndent )
  local listIndex = 1
  
  local wroteTypeIdSet = {}
  
  local function outputTypeInfo( typeInfo )
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
  self:writeln( "local _dependIdMap = {}", baseIndent )
  self:writeln( "_moduleObj._dependIdMap = _dependIdMap", baseIndent )
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
                      self:writeln( string.format( "_dependIdMap[ %d ] = { %d, %d }", typeInfo:get_typeId(), moduleIndex, extId), baseIndent )
                    end
                  end
              end
              
            end
        end
        
      end
    end
  end
  
  self:writeln( "local _dependModuleMap = {}", baseIndent )
  self:writeln( "_moduleObj._dependModuleMap = _dependModuleMap", baseIndent )
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
        self:writeln( string.format( "_dependModuleMap[ '%s' ] = { id = %d, use = %s }", name, _lune.unwrap( importModuleType2Index[moduleTypeInfo]), exportNeedModuleTypeInfo[moduleTypeInfo] or false), baseIndent )
      end
    end
  end
  
  local moduleTypeInfo = self.moduleTypeInfo
  
  do
    local _exp = node:get_provideNode()
    if _exp ~= nil then
    
        moduleTypeInfo = _exp:get_val():get_expType()
      end
  end
  
  self:writeln( string.format( "_moduleObj._moduleTypeId = %d", moduleTypeInfo:get_typeId()), baseIndent )
  self:writeln( string.format( "_moduleObj._moduleMutable = %s", moduleTypeInfo:get_mutable()), baseIndent )
  self:writeln( "----- meta -----", baseIndent )
  if self.stream ~= self.metaStream then
    self:writeln( "return _moduleObj", baseIndent )
  end
  self.outMetaFlag = false
end

function convFilter:processRoot( node, parent, baseIndent )
  self:writeln( string.format( "--%s", self.streamName), baseIndent )
  self:writeln( "local _moduleObj = {}", baseIndent )
  if node:get_luneHelperInfo():get_useNilAccess() or node:get_luneHelperInfo():get_useUnwrapExp() then
    self:writeln( [==[
if not _ENV._lune then
   _lune = {}
end]==], baseIndent )
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
]==], baseIndent )
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
]==], baseIndent )
    end
  end
  local children = node:get_children(  )
  
  for __index, child in pairs( children ) do
    filter( child, self, node, baseIndent )
    self:writeln( "", baseIndent )
  end
  self:outputMeta( node, baseIndent )
  do
    local _exp = node:get_provideNode()
    if _exp ~= nil then
    
        self:write( "return " )
        filter( _exp:get_val(), self, node, baseIndent )
        self:writeln( "", baseIndent )
      else
    
        self:writeln( "return _moduleObj", baseIndent )
      end
  end
  
end

-- none

function convFilter:processSubfile( node, parent, baseIndent )
end

function convFilter:processBlock( node, parent, baseIndent )
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
  
  self:writeln( word, baseIndent + stepIndent )
  local stmtList = node:get_stmtList(  )
  
  for __index, statement in pairs( stmtList ) do
    filter( statement, self, node, baseIndent + stepIndent )
    self:writeln( "", baseIndent + stepIndent )
  end
  self:setIndent( baseIndent )
  if node:get_blockKind(  ) == Ast.BlockKind.Block then
    self:writeln( "end", baseIndent )
  end
end

-- none

function convFilter:processStmtExp( node, parent, baseIndent )
  filter( node:get_exp(  ), self, node, baseIndent )
end

-- none

function convFilter:processDeclEnum( node, parent, baseIndent )
  local access = node:get_accessMode() == Ast.AccessMode.Global and "" or "local "
  
  self:writeln( string.format( "%s%s = {}", access, node:get_name().txt), baseIndent )
  if node:get_accessMode() == Ast.AccessMode.Pub then
    self:writeln( string.format( "_moduleObj.%s = %s", node:get_name().txt, node:get_name().txt), baseIndent )
  end
  local typeInfo = node:get_expType()
  
  if typeInfo:get_accessMode() ~= Ast.AccessMode.Pri then
    self.pubEnumId2EnumTypeInfo[typeInfo:get_typeId()] = typeInfo
  end
  self:writeln( string.format( "%s._val2NameMap = {}", node:get_name().txt), baseIndent )
  self:writeln( string.format( [==[function %s:_getTxt( val )
  local name = self._val2NameMap[ val ]
  if name then
    return string.format( "%s.%%s", name )
  end
  return string.format( "illegal val -- %%s", val )
end 
function %s:_from( val )
  if self._val2NameMap[ val ] then
    return val
  end
  return nil
end 
    ]==], node:get_name().txt, self:getCanonicalName( typeInfo ), node:get_name().txt), baseIndent )
  for __index, valName in pairs( node:get_valueNameList() ) do
    local valInfo = _lune.unwrap( typeInfo:getEnumValInfo( valName.txt ))
    
    local valTxt = string.format( "%s", valInfo:get_val())
    
    if typeInfo:get_valTypeInfo():equals( Ast.builtinTypeString ) then
      valTxt = string.format( "'%s'", valInfo:get_val())
    end
    self:writeln( string.format( "%s.%s = %s", node:get_name().txt, valName.txt, valTxt), baseIndent )
    self:writeln( string.format( "%s._val2NameMap[%s] = '%s'", node:get_name().txt, valTxt, valName.txt), baseIndent )
  end
end

function convFilter:getDestrClass( classTypeInfo )
  local typeInfo = classTypeInfo
  
  while not typeInfo:equals( Ast.rootTypeInfo ) do
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

function convFilter:processDeclClass( node, parent, baseIndent )
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
    
        self:writeln( string.format( "local %s = require( %s )", className, _exp.txt ), baseIndent )
        if node:get_accessMode() ~= Ast.AccessMode.Pri then
          self:writeln( string.format( "_moduleObj.%s = %s", className, className), baseIndent )
        end
        return 
      end
  end
  
  self:writeln( string.format( "local %s = {}", className ), baseIndent )
  local baseInfo = node:get_expType(  ):get_baseTypeInfo(  )
  
  if baseInfo:get_typeId(  ) ~= Ast.rootTypeId then
    self:writeln( string.format( "setmetatable( %s, { __index = %s } )", className, (_lune.unwrap( baseInfo) ):getTxt(  )), baseIndent )
  end
  if nodeInfo:get_accessMode(  ) == Ast.AccessMode.Pub then
    self:writeln( string.format( "_moduleObj.%s = %s", className, className ), baseIndent )
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
      filter( field, self, node, baseIndent )
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
end]==], className, className, destTxt), baseIndent )
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
]==], className, argTxt, className, argTxt, className, argTxt), baseIndent )
    for __index, member in pairs( memberList ) do
      local memberName = member:get_name().txt
      
      self:writeln( string.format( "self.%s = %s", memberName, memberName ), baseIndent + stepIndent )
    end
    self:writeln( 'end', baseIndent )
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
end]==], className, getterName, prefix, memberName), baseIndent )
      methodNameSet[getterName] = true
    end
    local setterName = "set_" .. memberName
    
    autoFlag = not methodNameSet[setterName]
    if memberNode:get_setterMode(  ) ~= Ast.AccessMode.None and autoFlag then
      self:writeln( string.format( [==[
function %s:%s( %s )   
  %s.%s = %s              
end]==], className, setterName, memberName, prefix, memberName, memberName), baseIndent )
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
]==], className, childName, memberName, childName), baseIndent )
        end
      end
    end
  end
  self:writeln( "do", baseIndent + stepIndent )
  for __index, initStmt in pairs( nodeInfo:get_initStmtList() ) do
    filter( initStmt, self, node, baseIndent + stepIndent )
    self:writeln( "", baseIndent + stepIndent )
  end
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processDeclMember( node, parent, baseIndent )
end

-- none

function convFilter:processExpMacroExp( node, parent, baseIndent )
  local stmtList = node:get_stmtList(  )
  
  if stmtList then
    for __index, stmt in pairs( stmtList ) do
      filter( stmt, self, node, baseIndent )
      self:writeln( "", baseIndent )
    end
  end
end

-- none

function convFilter:processDeclMacro( node, parent, baseIndent )
  if self.inMacro then
    local nodeInfo = node:get_declInfo(  )
    
    local name = nodeInfo:get_name(  )
    
    self:write( string.format( "local function %s(", name.txt) )
    self:writeln( "__macroArgs )", baseIndent )
    for index, arg in pairs( nodeInfo:get_argList(  ) ) do
      local argName = arg:get_name().txt
      
      self:writeln( string.format( "local %s = __macroArgs.%s", argName, argName), baseIndent )
    end
    self:writeln( "local macroVar = {}", baseIndent )
    self:writeln( "macroVar._names = {}", baseIndent )
    self.macroDepth = self.macroDepth + 1
    do
      local _exp = nodeInfo:get_ast(  )
      if _exp ~= nil then
      
          filter( _exp, self, node, baseIndent )
        end
    end
    
    self.macroDepth = self.macroDepth - 1
    self:writeln( "", baseIndent )
    self:writeln( "return macroVar", baseIndent )
    self:writeln( "end", baseIndent )
    self:writeln( string.format( "return %s", name.txt), baseIndent )
  end
end

-- none

function convFilter:processExpMacroStat( node, parent, baseIndent )
  for index, token in pairs( node:get_expStrList(  ) ) do
    if index ~= 1 then
      self:write( '..' )
    end
    filter( token, self, node, baseIndent )
  end
end

-- none

function convFilter:processExpNew( node, parent, baseIndent )
  filter( node:get_symbol(  ), self, node, baseIndent )
  self:write( ".new(" )
  do
    local _exp = node:get_argList(  )
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( ")" )
end

-- none

function convFilter:processDeclConstr( node, parent, baseIndent )
  local declInfo = node:get_declInfo(  )
  
  local classNameToken = _lune.unwrap( declInfo:get_className(  ))
  
  local className = classNameToken.txt
  
  self:write( string.format( "function %s.new( ", className ) )
  local argTxt = ""
  
  local argList = declInfo:get_argList(  )
  
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
      argTxt = argTxt .. ", "
    end
    filter( arg, self, node, baseIndent )
    if arg:get_kind(  ) == Ast.nodeKind['DeclArg'] then
      argTxt = argTxt .. (arg ):get_name().txt
    else 
      local name = _lune.unwrap( node:get_declInfo(  ):get_name())
      
      Util.err( string.format( "not support ... in macro -- %s", name.txt) )
    end
  end
  self:writeln( " )", baseIndent + stepIndent )
  self:writeln( "local obj = {}", baseIndent + stepIndent )
  self:writeln( string.format( "%s.setmeta( obj )", className), baseIndent + stepIndent )
  self:writeln( string.format( "if obj.__init then obj:__init( %s ); end", argTxt ), baseIndent )
  self:writeln( "return obj", baseIndent )
  self:writeln( "end", baseIndent )
  self:write( string.format( "function %s:__init(%s) ", className, argTxt ) )
  do
    local _exp = declInfo:get_body()
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processDeclDestr( node, parent, baseIndent )
  self:writeln( string.format( "function %s.__free( self )", _lune.nilacc( node:get_declInfo():get_className(), "txt" )), baseIndent + stepIndent )
  filter( _lune.unwrap( node:get_declInfo():get_body()), self, node, baseIndent + stepIndent )
  local classTypeInfo = node:get_expType():get_parentInfo()
  
  do
    local _exp = self:getDestrClass( classTypeInfo:get_baseTypeInfo() )
    if _exp ~= nil then
    
        self:writeln( string.format( "%s.__free( self )", _exp:getTxt(  )), baseIndent )
      end
  end
  
  self:writeln( "end", baseIndent )
end

function convFilter:processExpCallSuper( node, parent, baseIndent )
  local typeInfo = node:get_superType(  )
  
  self:write( string.format( "%s.__init( self, ", typeInfo:getTxt(  )) )
  do
    local _exp = node:get_expList()
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:writeln( ")", baseIndent )
end

-- none

function convFilter:processDeclMethod( node, parent, baseIndent )
  local declInfo = node:get_declInfo(  )
  
  local delimit = ":"
  
  if declInfo:get_staticFlag(  ) then
    delimit = "."
  end
  local methodNodeToken = _lune.unwrap( declInfo:get_name(  ))
  
  local methodName = methodNodeToken.txt
  
  local classNameToken = _lune.unwrap( declInfo:get_className(  ))
  
  self:write( string.format( "function %s%s%s( ", classNameToken.txt, delimit, methodName) )
  local argList = declInfo:get_argList(  )
  
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
    end
    filter( arg, self, node, baseIndent )
  end
  self:write( " )" )
  do
    local _exp = declInfo:get_body()
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      else
    
        self:writeln( "", baseIndent )
      end
  end
  
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processUnwrapSet( node, parent, baseIndent )
  local dstExpList = node:get_dstExpList()
  
  filter( dstExpList, self, node, baseIndent )
  self:write( " = " )
  filter( node:get_srcExpList(), self, node, baseIndent )
  self:writeln( "", baseIndent )
  self:write( "if " )
  for index, expNode in pairs( dstExpList:get_expList() ) do
    if index > 1 then
      self:write( " or " )
    end
    self:write( "nil == " )
    filter( expNode, self, node, baseIndent )
  end
  self:writeln( " then", baseIndent + stepIndent )
  for index, expNode in pairs( dstExpList:get_expList() ) do
    self:write( string.format( "local _exp%d = ", index) )
    filter( expNode, self, node, baseIndent + stepIndent )
    self:writeln( "", baseIndent + stepIndent )
  end
  if node:get_unwrapBlock() then
    filter( _lune.unwrap( node:get_unwrapBlock()), self, node, baseIndent + stepIndent )
  end
  self:writeln( "end", baseIndent )
end

function convFilter:processIfUnwrap( node, parent, baseIndent )
  self:writeln( "do", baseIndent + stepIndent )
  self:write( "local " )
  for index, varName in pairs( node:get_varNameList() ) do
    self:write( varName )
    if index ~= #node:get_varNameList() then
      self:write( ", " )
    end
  end
  self:write( " = " )
  for index, expNode in pairs( node:get_expNodeList() ) do
    filter( expNode, self, node, baseIndent + stepIndent )
    if index ~= #node:get_expNodeList() then
      self:write( ", " )
    end
  end
  self:writeln( "", baseIndent + stepIndent )
  self:write( "if " )
  for index, varName in pairs( node:get_varNameList() ) do
    self:write( string.format( "%s ~= nil", varName) )
    if index ~= #node:get_varNameList() then
      self:write( " and " )
    end
  end
  self:writeln( " then", baseIndent + stepIndent )
  filter( node:get_block(), self, node, baseIndent + stepIndent * 2 )
  if node:get_nilBlock() then
    self:writeln( "else", baseIndent + stepIndent )
    filter( _lune.unwrap( node:get_nilBlock()), self, node, baseIndent + stepIndent * 2 )
  end
  self:writeln( "end", baseIndent )
  self:writeln( "end", baseIndent )
end

function convFilter:processDeclVar( node, parent, baseIndent )
  if node:get_syncBlock() then
    self:writeln( "do", baseIndent + stepIndent )
    for __index, varInfo in pairs( node:get_syncVarList() ) do
      self:writeln( string.format( "local _sync_%s", varInfo:get_name().txt), baseIndent + stepIndent )
    end
    self:writeln( "do", baseIndent + stepIndent * 2 )
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
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:writeln( "", baseIndent )
  do
    local _exp = node:get_unwrapBlock()
    if _exp ~= nil then
    
        self:writeln( "", baseIndent + stepIndent * 2 )
        self:write( "if " )
        for index, var in pairs( varList ) do
          if index > 1 then
            self:write( " or " )
          end
          self:write( " nil == " .. var:get_name().txt )
        end
        self:writeln( " then", baseIndent + stepIndent * 3 )
        for index, var in pairs( varList ) do
          self:writeln( string.format( "local _%s = %s", var:get_name().txt, var:get_name().txt), baseIndent + stepIndent * 3 )
        end
        filter( _exp, self, node, baseIndent + stepIndent * 2 )
        do
          local thenBlock = node:get_thenBlock()
          if thenBlock ~= nil then
          
              self:writeln( "else", baseIndent + stepIndent * 3 )
              filter( thenBlock, self, node, baseIndent + stepIndent * 3 )
            end
        end
        
        -- none
        
        self:writeln( "end", baseIndent + stepIndent * 1 )
      end
  end
  
  do
    local _exp = node:get_syncBlock()
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent + stepIndent * 1 )
        for __index, varInfo in pairs( node:get_syncVarList() ) do
          self:writeln( string.format( "_sync_%s = %s", varInfo:get_name().txt, varInfo:get_name().txt), baseIndent + stepIndent )
        end
        self:writeln( "end", baseIndent + stepIndent )
        for __index, varInfo in pairs( node:get_syncVarList() ) do
          self:writeln( string.format( "%s = _sync_%s", varInfo:get_name().txt, varInfo:get_name().txt), baseIndent + stepIndent )
        end
        self:writeln( "end", baseIndent )
      end
  end
  
  if node:get_accessMode(  ) == Ast.AccessMode.Pub then
    self:writeln( "", baseIndent )
    for index, var in pairs( varList ) do
      local name = var:get_name().txt
      
      self:writeln( string.format( "_moduleObj.%s = %s", name, name), baseIndent )
      self.pubVarName2InfoMap[name] = PubVerInfo.new(node:get_staticFlag(), node:get_accessMode(), node:get_symbolInfoList()[index]:get_mutable(), node:get_typeInfoList()[index])
    end
  end
  if self.macroDepth > 0 then
    self:writeln( "", baseIndent )
    for index, var in pairs( varList ) do
      local varName = var:get_name().txt
      
      self:writeln( string.format( "table.insert( macroVar._names, '%s' )", varName), baseIndent )
      self:writeln( string.format( "macroVar.%s = %s", varName, varName), baseIndent )
    end
  end
end

-- none

function convFilter:processDeclArg( node, parent, baseIndent )
  self:write( string.format( "%s", node:get_name(  ).txt ) )
end

-- none

function convFilter:processDeclArgDDD( node, parent, baseIndent )
  self:write( "..." )
end

-- none

function convFilter:processExpDDD( node, parent, baseIndent )
  self:write( "..." )
end

-- none

function convFilter:processDeclFunc( node, parent, baseIndent )
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
    filter( arg, self, node, baseIndent )
  end
  self:write( " )" )
  do
    local _exp = declInfo:get_body()
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      else
    
        self:writeln( "", baseIndent )
      end
  end
  
  self:writeln( "end", baseIndent )
  local expType = node:get_expType(  )
  
  if expType:get_accessMode(  ) == Ast.AccessMode.Pub then
    self:write( string.format( "_moduleObj.%s = %s", name, name) )
    self.pubFuncName2InfoMap[name] = PubFuncInfo.new(declInfo:get_accessMode(  ), node:get_expType(  ))
  end
end

-- none

function convFilter:processRefType( node, parent, baseIndent )
  self:write( (node:get_refFlag(  ) and "&" or "" ) .. (node:get_mutFlag(  ) and "mut " or "" ) )
  filter( node:get_name(  ), self, node, baseIndent )
  if node:get_array(  ) == "array" then
    self:write( "[@]" )
  elseif node:get_array(  ) == "list" then
    self:write( "[]" )
  end
end

-- none

function convFilter:processIf( node, parent, baseIndent )
  local valList = node:get_stmtList(  )
  
  for index, val in pairs( valList ) do
    if index == 1 then
      self:write( "if " )
      filter( val:get_exp(), self, node, baseIndent )
    elseif val:get_kind() == "elseif" then
      self:write( "elseif " )
      filter( val:get_exp(), self, node, baseIndent )
    else 
      self:write( "else" )
    end
    self:write( " " )
    filter( val:get_block(), self, node, baseIndent )
  end
  self:write( "end" )
end

-- none

function convFilter:processSwitch( node, parent, baseIndent )
  self:writeln( "do", baseIndent + 2 )
  self:write( "local _switchExp = " )
  filter( node:get_exp(  ), self, node, baseIndent + 2 )
  self:writeln( "", baseIndent + 2 )
  for index, caseInfo in pairs( node:get_caseList(  ) ) do
    if index == 1 then
      self:write( "if " )
    else 
      self:write( "elseif " )
    end
    local expList = caseInfo:get_expList(  )
    
    for index, expNode in pairs( expList:get_expList(  ) ) do
      if index ~= 1 then
        self:write( " or " )
      end
      self:write( "_switchExp == " )
      filter( expNode, self, node, baseIndent + 2 )
    end
    self:write( " then" )
    filter( caseInfo:get_block(), self, node, baseIndent + 2 )
  end
  do
    local _exp = node:get_default(  )
    if _exp ~= nil then
    
        self:write( "else " )
        filter( _exp, self, node, baseIndent + 2 )
      end
  end
  
  self:writeln( "end", baseIndent )
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processWhile( node, parent, baseIndent )
  self:write( "while " )
  filter( node:get_exp(  ), self, node, baseIndent )
  self:write( " " )
  filter( node:get_block(  ), self, node, baseIndent )
  self:write( "end" )
end

-- none

function convFilter:processRepeat( node, parent, baseIndent )
  self:write( "repeat " )
  filter( node:get_block(  ), self, node, baseIndent )
  self:write( "until " )
  filter( node:get_exp(  ), self, node, baseIndent )
end

-- none

function convFilter:processFor( node, parent, baseIndent )
  self:write( string.format( "for %s = ", node:get_val(  ).txt ) )
  filter( node:get_init(  ), self, node, baseIndent )
  self:write( ", " )
  filter( node:get_to(  ), self, node, baseIndent )
  do
    local _exp = node:get_delta(  )
    if _exp ~= nil then
    
        self:write( ", " )
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( " " )
  filter( node:get_block(  ), self, node, baseIndent )
  self:write( "end" )
end

-- none

function convFilter:processApply( node, parent, baseIndent )
  self:write( "for " )
  local varList = node:get_varList(  )
  
  for index, var in pairs( varList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( var.txt )
  end
  self:write( " in " )
  filter( node:get_exp(), self, node, baseIndent )
  self:write( " " )
  filter( node:get_block(), self, node, baseIndent )
  self:write( "end" )
end

-- none

function convFilter:processForeach( node, parent, baseIndent )
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
  filter( node:get_exp(), self, node, baseIndent )
  self:write( " ) " )
  filter( node:get_block(), self, node, baseIndent )
  self:write( "end" )
end

-- none

function convFilter:processForsort( node, parent, baseIndent )
  self:writeln( "do", baseIndent + stepIndent )
  self:writeln( "local __sorted = {}", baseIndent + stepIndent )
  self:write( "local __map = " )
  filter( node:get_exp(), self, node, baseIndent + stepIndent )
  self:writeln( "", baseIndent + stepIndent )
  self:writeln( "for __key in pairs( __map ) do", baseIndent + stepIndent * 2 )
  self:writeln( "table.insert( __sorted, __key )", baseIndent + stepIndent )
  self:writeln( "end", baseIndent + stepIndent )
  self:writeln( "table.sort( __sorted )", baseIndent + stepIndent )
  self:write( "for __index, " )
  local key = "__key"
  
  do
    local _exp = node:get_key()
    if _exp ~= nil then
    
        key = _exp.txt
      end
  end
  
  self:write( key )
  self:writeln( " in ipairs( __sorted ) do", baseIndent + stepIndent * 2 )
  self:writeln( string.format( "%s = __map[ %s ]", node:get_val().txt, key ), baseIndent + stepIndent * 2 )
  filter( node:get_block(), self, node, baseIndent + stepIndent * 2 )
  self:writeln( "end", baseIndent + stepIndent )
  self:writeln( "end", baseIndent )
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processExpUnwrap( node, parent, baseIndent )
  do
    local _exp = node:get_default()
    if _exp ~= nil then
    
        self:write( '_lune.unwrapDefault( ' )
        filter( node:get_exp(), self, node, baseIndent )
        self:write( ', ' )
        filter( _exp, self, node, baseIndent )
        self:write( ')' )
      else
    
        self:write( '_lune.unwrap( ' )
        filter( node:get_exp(), self, node, baseIndent )
        self:write( ')' )
      end
  end
  
end

function convFilter:processExpCall( node, parent, baseIndent )
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
        filter( prefixNode, self, fieldNode, baseIndent )
      else 
        self:write( "_lune.nilacc( " )
        filter( prefixNode, self, fieldNode, baseIndent )
        self:write( string.format( ", '%s', 'callmtd' ", fieldNode:get_field().txt) )
      end
    else 
      if prefixType:get_kind() == Ast.TypeInfoKind.List then
        setArgFlag = true
        wroteFuncFlag = true
        self:write( string.format( "table.%s( ", fieldNode:get_field().txt) )
        filter( prefixNode, self, fieldNode, baseIndent )
      elseif prefixType:get_kind() == Ast.TypeInfoKind.Enum then
        wroteFuncFlag = true
        local fieldExpType = fieldNode:get_expType()
        
        local canonicalName = self:getCanonicalName( prefixType )
        
        local methodName = fieldNode:get_field().txt
        
        if methodName == "get__txt" then
          methodName = "_getTxt"
        end
        self:write( string.format( "%s:%s( ", canonicalName, methodName) )
        if fieldExpType:get_staticFlag() then
          setArgFlag = false
        else 
          filter( prefixNode, self, fieldNode, baseIndent )
          setArgFlag = true
        end
      end
    end
    
  end
  if not wroteFuncFlag then
    if node:get_nilAccess() then
      self:write( "_lune.nilacc( " )
      filter( node:get_func(), self, node, baseIndent )
      self:write( ", nil, 'call'" )
      wroteFuncFlag = true
    else 
      filter( node:get_func(), self, node, baseIndent )
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
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( " )" )
end

-- none

function convFilter:processExpList( node, parent, baseIndent )
  local expList = node:get_expList(  )
  
  for index, exp in pairs( expList ) do
    if index > 1 then
      self:write( ", " )
    end
    filter( exp, self, node, baseIndent )
  end
end

-- none

function convFilter:processExpOp1( node, parent, baseIndent )
  local op = node:get_op().txt
  
  if op == ",,," then
    filter( node:get_exp(), self, node, baseIndent )
  elseif op == ",,,," then
    if node:get_macroMode() == Ast.MacroMode.Expand then
      filter( node:get_exp(), self, node, baseIndent )
    else 
      self:write( "_luneSym2Str( " )
      filter( node:get_exp(), self, node, baseIndent )
      self:write( " )" )
    end
  elseif op == ",," then
    self:write( "_luneGetLocal( " )
    filter( node:get_exp(), self, node, baseIndent )
    self:write( " )" )
  else 
    if op == "not" then
      op = op .. " "
    end
    self:write( op )
    filter( node:get_exp(), self, node, baseIndent )
  end
end

-- none

function convFilter:processExpCast( node, parent, baseIndent )
  if node:get_expType():equals( Ast.builtinTypeInt ) then
    self:write( "math.floor(" )
    filter( node:get_exp(), self, node, baseIndent )
    self:write( ")" )
  else 
    filter( node:get_exp(), self, node, baseIndent )
  end
end

-- none

function convFilter:processExpParen( node, parent, baseIndent )
  self:write( "(" )
  filter( node:get_exp(), self, node, baseIndent )
  self:write( " )" )
end

-- none

function convFilter:processExpOp2( node, parent, baseIndent )
  local intCast = false
  
  if node:get_expType():equals( Ast.builtinTypeInt ) and node:get_op().txt == "/" then
    intCast = true
    self:write( "math.floor(" )
  end
  filter( node:get_exp1(), self, node, baseIndent )
  self:write( " " .. node:get_op().txt .. " " )
  filter( node:get_exp2(), self, node, baseIndent )
  if intCast then
    self:write( ")" )
  end
end

-- none

function convFilter:processExpRef( node, parent, baseIndent )
  if node:get_symbolInfo():get_accessMode() == Ast.AccessMode.Pub and node:get_symbolInfo():get_kind() == Ast.SymbolKind.Var then
    self:write( "_moduleObj." )
  end
  self:write( node:get_token().txt )
end

-- none

function convFilter:processExpRefItem( node, parent, baseIndent )
  if node:get_nilAccess() then
    self:write( "_lune.nilacc( " )
    filter( node:get_val(), self, node, baseIndent )
    self:write( ", nil, 'item', " )
    do
      local _exp = node:get_index()
      if _exp ~= nil then
      
          filter( _exp, self, node, baseIndent )
        else
      
          self:write( string.format( "'%s'", _lune.unwrap( node:get_symbol())) )
        end
    end
    
    self:write( ")" )
  else 
    if node:get_val():get_expType():equals( Ast.builtinTypeString ) then
      self:write( "string.byte( " )
      filter( node:get_val(), self, node, baseIndent )
      self:write( ", " )
      do
        local _exp = node:get_index()
        if _exp ~= nil then
        
            filter( _exp, self, node, baseIndent )
          else
        
            error( "index is nil" )
          end
      end
      
      self:write( " )" )
    else 
      filter( node:get_val(), self, node, baseIndent )
      self:write( "[" )
      do
        local _exp = node:get_index()
        if _exp ~= nil then
        
            filter( _exp, self, node, baseIndent )
          else
        
            self:write( string.format( "'%s'", _lune.unwrap( node:get_symbol())) )
          end
      end
      
      self:write( "]" )
    end
  end
end

-- none

function convFilter:processRefField( node, parent, baseIndent )
  local prefix = node:get_prefix(  )
  
  if node:get_nilAccess() then
    self:write( '_lune.nilacc( ' )
    filter( prefix, self, node, baseIndent )
    self:write( string.format( ', "%s" )', node:get_field().txt) )
  else 
    filter( prefix, self, node, baseIndent )
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

-- none

function convFilter:processExpOmitEnum( node, parent, baseIndent )
  local moduleName = self.typeInfo2ModuleName[node:get_expType():getModule(  )]
  
      if  nil == moduleName then
        local _moduleName = moduleName
        
        moduleName = ""
      else
        
          moduleName = moduleName:gsub( ".*%.", "" )
          moduleName = moduleName .. "."
        end
    
  self:write( string.format( "%s%s.%s", moduleName, node:get_expType():getTxt(  ), node:get_valToken().txt) )
end

-- none

function convFilter:processGetField( node, parent, baseIndent )
  local prefixNode = node:get_prefix(  )
  
  local prefixType = prefixNode:get_expType()
  
  local fieldTxt = node:get_field(  ).txt
  
  if fieldTxt == "_txt" and prefixType:get_kind() == Ast.TypeInfoKind.Enum then
    local canonicalName = self:getCanonicalName( prefixType )
    
    self:write( string.format( "%s:_getTxt( ", canonicalName) )
    filter( prefixNode, self, node, baseIndent )
    self:writeln( ")", baseIndent )
  else 
    filter( prefixNode, self, node, baseIndent )
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

-- none

function convFilter:processReturn( node, parent, baseIndent )
  self:write( "return " )
  do
    local _exp = node:get_expList()
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
end

-- none

function convFilter:processProvide( node, parent, baseIndent )
end

function convFilter:processLiteralList( node, parent, baseIndent )
  self:write( "{" )
  do
    local _exp = node:get_expList()
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( "}" )
end

-- none

function convFilter:processLiteralMap( node, parent, baseIndent )
  self:write( "{" )
  local pairList = node:get_pairList()
  
  for index, pair in pairs( pairList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( "[" )
    filter( pair:get_key(), self, node, baseIndent )
    self:write( "] = " )
    filter( pair:get_val(), self, node, baseIndent )
  end
  self:write( "}" )
end

-- none

function convFilter:processLiteralArray( node, parent, baseIndent )
  self:write( "{" )
  do
    local _exp = node:get_expList()
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( "}" )
end

-- none

function convFilter:processLiteralChar( node, parent, baseIndent )
  self:write( string.format( "%g", node:get_num() ) )
end

-- none

function convFilter:processLiteralInt( node, parent, baseIndent )
  self:write( node:get_token().txt )
end

-- none

function convFilter:processLiteralReal( node, parent, baseIndent )
  self:write( node:get_token().txt )
end

-- none

function convFilter:processLiteralString( node, parent, baseIndent )
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
      filter( val, self, node, baseIndent )
    end
    self:write( ")" )
  else 
    self:write( txt )
  end
end

-- none

function convFilter:processLiteralBool( node, parent, baseIndent )
  self:write( node:get_token().txt )
end

-- none

function convFilter:processLiteralNil( node, parent, baseIndent )
  self:write( "nil" )
end

-- none

function convFilter:processBreak( node, parent, baseIndent )
  self:write( "break" )
end

-- none

function convFilter:processLiteralSymbol( node, parent, baseIndent )
  self:write( string.format( '%s', node:get_token().txt) )
end

-- none

local function createFilter( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo )
  return convFilter.new(streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo)
end
_moduleObj.createFilter = createFilter
local MacroEvalImp = {}
setmetatable( MacroEvalImp, { __index = MacroEval } )
_moduleObj.MacroEvalImp = MacroEvalImp
function MacroEvalImp:eval( node )
  local oStream = Util.memStream.new()
  
  local conv = convFilter.new("macro", oStream, oStream, ConvMode.Exec, true, Ast.rootTypeInfo)
  
  conv:processDeclMacro( node, node, 0 )
  local chunk, err = load( oStream:get_txt(  ) )
  
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
do
  end

return _moduleObj
