--lune/base/convLua.lns
local _moduleObj = {}
local function _lune_nilacc( val, fieldName, access, ... )
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
function _lune_unwrap( val )
  if val == nil then
     _luneScript.error( 'unwrap val is nil' )
  end
  return val
end
function _lune_unwrapDefault( val, defval )
  if val == nil then
     return defval
  end
  return val
end

local Ast = require( 'lune.base.Ast' )

local Util = require( 'lune.base.Util' )

local PubVerInfo = {}
function PubVerInfo.new( staticFlag, accessMode, typeInfo )
  local obj = {}
  setmetatable( obj, { __index = PubVerInfo } )
  if obj.__init then
    obj:__init( staticFlag, accessMode, typeInfo )
  end        
  return obj 
 end         
function PubVerInfo:__init( staticFlag, accessMode, typeInfo ) 
            
self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.typeInfo = typeInfo
  end
do
  end

-- none

local PubFuncInfo = {}
function PubFuncInfo.new( accessMode, typeInfo )
  local obj = {}
  setmetatable( obj, { __index = PubFuncInfo } )
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

local convFilter = {}
setmetatable( convFilter, { __index = Filter } )
_moduleObj.convFilter = convFilter
function convFilter.new( streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo )
  local obj = {}
  setmetatable( obj, { __index = convFilter } )
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

  local canonicalName = typeInfo:getTxt(  )
  
  local workType = typeInfo:get_parentInfo()
  
  while workType ~= Ast.rootTypeInfo do
    canonicalName = string.format( "%s.%s", workType:getTxt(  ), canonicalName)
    if self.typeInfo2ModuleName[workType] then
      break
    end
    workType = workType:get_parentInfo()
  end
  return canonicalName
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
builtInModuleSet["_luneScript"] = true
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
  
  self.typeInfo2ModuleName[node:get_moduleTypeInfo()] = moduleName
  if self.convMode == "exe" or self.convMode == "ast" then
    self:writeln( string.format( "local %s = _luneScript.loadModule( '%s' )", moduleName, module), baseIndent )
  else 
    self:writeln( string.format( "local %s = require( '%s' )", moduleName, module), baseIndent )
  end
end

-- none

function convFilter:outputMeta( node, baseIndent )

  do
    local _switchExp = self.convMode
    if _switchExp == "lua" or _switchExp == "save" then
      return 
    end
  end
  
  self.outMetaFlag = true
  if self.stream ~= self.metaStream then
    self:writeln( "local _moduleObj = {}", baseIndent )
  end
  self:writeln( "----- meta -----", baseIndent )
  local typeId2TypeInfo = {}
  
  local typeId2UseFlag = {}
  
  local pickupClassMap = {}
  
  local function pickupTypeId( typeInfo, forceFlag, pickupChildFlag )
  
    if typeInfo:get_typeId(  ) == Ast.rootTypeId then
      return 
    end
    if not forceFlag and typeInfo:get_accessMode(  ) ~= "pub" then
      return 
    end
    if typeId2TypeInfo[typeInfo:get_typeId(  )] then
      if pickupChildFlag and not typeInfo:get_nilable() then
        for __index, itemTypeInfo in pairs( typeInfo:get_children(  ) ) do
          if itemTypeInfo:get_accessMode() == "pub" and (itemTypeInfo:get_kind(  ) == Ast.TypeInfoKindClass or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKindIF or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKindFunc or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKindMethod ) then
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
      if typeInfo:get_kind() == Ast.TypeInfoKindClass or typeInfo:get_kind() == Ast.TypeInfoKindIF then
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
          if itemTypeInfo:get_accessMode() == "pub" and (itemTypeInfo:get_kind(  ) == Ast.TypeInfoKindClass or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKindIF or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKindFunc or itemTypeInfo:get_kind(  ) == Ast.TypeInfoKindMethod ) then
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
    if namespaceInfo.typeInfo:get_accessMode(  ) == "pub" and not namespaceInfo.typeInfo:get_externalFlag() then
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
        if classTypeInfo:get_accessMode() == "pub" then
          pickupTypeId( classTypeInfo, true, validChildrenSet[classTypeInfo] == nil and not classTypeInfo:get_externalFlag() )
          pickupClassMap[classTypeId] = nil
          self:writeln( "do", baseIndent + stepIndent )
          self:writeln( string.format( "local _classInfo%d = {}", classTypeId), baseIndent + stepIndent )
          self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId), baseIndent + stepIndent )
          for __index, memberNode in pairs( _lune_unwrap( self.classId2MemberList[classTypeId]) ) do
            if memberNode:get_accessMode() ~= "pri" then
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
          local scope = _lune_unwrap( classTypeInfo:get_scope())
          
          if not Ast.isBuiltin( classTypeId ) then
            local className = classTypeInfo:getTxt(  )
            
            self:writeln( "do", baseIndent + stepIndent )
            self:writeln( string.format( "local _classInfo%s = {}", classTypeId), baseIndent + stepIndent )
            self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId), baseIndent + stepIndent )
            pickupTypeId( classTypeInfo, true, validChildrenSet[classTypeInfo] == nil and not classTypeInfo:get_externalFlag() )
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
                    if symbolInfo:get_accessMode() == "pub" then
                      self:writeln( string.format( "_classInfo%d.%s = {", classTypeId, fieldName), baseIndent + stepIndent )
                      self:writeln( string.format( "  name='%s', staticFlag = %s, ", fieldName, symbolInfo:get_staticFlag()) .. string.format( "accessMode = '%s', typeId = %d }", symbolInfo:get_accessMode(), typeInfo:get_typeId(  )), baseIndent + stepIndent )
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
        self:writeln( string.format( "  name='%s', accessMode = '%s', typeId = %d }", varName, varInfo.accessMode, varInfo.typeInfo:get_typeId()), baseIndent )
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
    if typeId2TypeInfo[typeId] and not Ast.isBuiltin( typeId ) then
      self:write( string.format( "_typeInfoList[%d] = ", listIndex) )
      listIndex = listIndex + 1
      local validChildren = validChildrenSet[typeInfo]
      
      if not validChildren then
        validChildren = typeId2TypeInfo
      end
      typeInfo:serialize( self, validChildren )
    end
  end
  
  for typeId, typeInfo in pairs( self.pubEnumId2EnumTypeInfo ) do
    typeId2TypeInfo[typeId] = typeInfo
  end
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
      end
    end
  end
  
  do
    local _exp = node:get_provideNode()
    if _exp ~= nil then
    
        self:writeln( string.format( "_moduleObj._moduleTypeId = %d", _exp:get_val():get_expType():get_typeId()), baseIndent )
      else
    
        self:writeln( string.format( "_moduleObj._moduleTypeId = %d", self.moduleTypeInfo:get_typeId()), baseIndent )
      end
  end
  
  self:writeln( "----- meta -----", baseIndent )
  if self.stream ~= self.metaStream then
    self:writeln( "return _moduleObj", baseIndent )
  end
  self.outMetaFlag = false
end

function convFilter:processRoot( node, parent, baseIndent )

  self:writeln( string.format( "--%s", self.streamName), baseIndent )
  self:writeln( "local _moduleObj = {}", baseIndent )
  self:writeln( [==[
local function _lune_nilacc( val, fieldName, access, ... )
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
function _lune_unwrap( val )
  if val == nil then
     _luneScript.error( 'unwrap val is nil' )
  end
  return val
end
function _lune_unwrapDefault( val, defval )
  if val == nil then
     return defval
  end
  return val
end
]==], baseIndent )
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
  
  if node:get_blockKind(  ) == "if" or node:get_blockKind(  ) == "elseif" then
    word = "then"
  elseif node:get_blockKind(  ) == "else" then
    word = ""
  elseif node:get_blockKind(  ) == "while" then
    word = "do"
  elseif node:get_blockKind(  ) == "repeat" then
    word = ""
  elseif node:get_blockKind(  ) == "for" then
    word = "do"
  elseif node:get_blockKind(  ) == "apply" then
    word = "do"
  elseif node:get_blockKind(  ) == "foreach" then
    word = "do"
  elseif node:get_blockKind(  ) == "macro" then
    word = ""
  elseif node:get_blockKind(  ) == "func" then
    word = ""
  elseif node:get_blockKind(  ) == "default" then
    word = ""
  elseif node:get_blockKind(  ) == "{" then
    word = "do"
  elseif node:get_blockKind(  ) == "macro" then
    word = ""
  elseif node:get_blockKind(  ) == "let!" then
    word = ""
  elseif node:get_blockKind(  ) == "if!" then
    word = ""
  end
  self:writeln( word, baseIndent + stepIndent )
  local stmtList = node:get_stmtList(  )
  
  for __index, statement in pairs( stmtList ) do
    filter( statement, self, node, baseIndent + stepIndent )
    self:writeln( "", baseIndent + stepIndent )
  end
  self:setIndent( baseIndent )
  if node:get_blockKind(  ) == "{" then
    self:writeln( "end", baseIndent )
  end
end

-- none

function convFilter:processStmtExp( node, parent, baseIndent )

  filter( node:get_exp(  ), self, node, baseIndent )
end

-- none

function convFilter:processDeclEnum( node, parent, baseIndent )

  local access = node:get_accessMode() == "global" and "" or "local "
  
  self:writeln( string.format( "%s%s = {}", access, node:get_name().txt), baseIndent )
  if node:get_accessMode() == "pub" then
    self:writeln( string.format( "_moduleObj.%s = %s", node:get_name().txt, node:get_name().txt), baseIndent )
  end
  local typeInfo = node:get_expType()
  
  if typeInfo:get_accessMode() ~= "pri" then
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
function %s:_fromVal( val )
   if self._val2NameMap[ val ] then
      return val
   end
   return nil
end
]==], node:get_name().txt, self:getCanonicalName( typeInfo ), node:get_name().txt), baseIndent )
  for __index, valName in pairs( node:get_valueNameList() ) do
    local valInfo = _lune_unwrap( typeInfo:getEnumValInfo( valName.txt ))
    
    local valTxt = string.format( "%s", valInfo:get_val())
    
    if typeInfo:get_valTypeInfo() == Ast.builtinTypeString then
      valTxt = string.format( "'%s'", valInfo:get_val())
      Util.errorLog( string.format( "hoge: %s", valTxt) )
    end
    self:writeln( string.format( "%s.%s = %s", node:get_name().txt, valName.txt, valTxt), baseIndent )
    self:writeln( string.format( "%s._val2NameMap[%s] = '%s'", node:get_name().txt, valTxt, valName.txt), baseIndent )
  end
end

function convFilter:processDeclClass( node, parent, baseIndent )

  local nodeInfo = node
  
  local classNameToken = nodeInfo:get_name(  )
  
  local className = classNameToken.txt
  
  local classTypeInfo = node:get_expType(  )
  
  local classTypeId = classTypeInfo:get_typeId()
  
  if nodeInfo:get_accessMode(  ) == "pub" then
    self.classId2TypeInfo[classTypeId] = classTypeInfo
  end
  self.classId2MemberList[classTypeId] = nodeInfo:get_memberList(  )
  do
    local _exp = node:get_moduleName()
    if _exp ~= nil then
    
        self:writeln( string.format( "local %s = require( %s )", className, _exp.txt ), baseIndent )
        if node:get_accessMode() ~= "pri" then
          self:writeln( string.format( "_moduleObj.%s = %s", className, className), baseIndent )
        end
        return 
      end
  end
  
  self:writeln( string.format( "local %s = {}", className ), baseIndent )
  local baseInfo = node:get_expType(  ):get_baseTypeInfo(  )
  
  if baseInfo:get_typeId(  ) ~= Ast.rootTypeId then
    self:writeln( string.format( "setmetatable( %s, { __index = %s } )", className, (_lune_unwrap( baseInfo) ):getTxt(  )), baseIndent )
  end
  if nodeInfo:get_accessMode(  ) == "pub" then
    self:writeln( string.format( "_moduleObj.%s = %s", className, className ), baseIndent )
  end
  local hasConstrFlag = false
  
  local memberList = {}
  
  local fieldList = nodeInfo:get_fieldList(  )
  
  local outerMethodSet = nodeInfo:get_outerMethodSet(  )
  
  local methodNameSet = {}
  
  for __index, field in pairs( fieldList ) do
    local ignoreFlag = false
    
    if field:get_kind() == Ast.nodeKind.DeclConstr then
      hasConstrFlag = true
      methodNameSet["__init"] = true
    end
    if field:get_kind() == Ast.nodeKind.DeclMember then
      local declMemberNode = field
      
      if not declMemberNode:get_staticFlag() then
        table.insert( memberList, declMemberNode )
      end
    end
    if field:get_kind() == Ast.nodeKind.DeclMethod then
      local methodNode = field
      
      local declInfo = methodNode:get_declInfo(  )
      
      local methodNameToken = _lune_unwrap( declInfo:get_name(  ))
      
      if outerMethodSet[methodNameToken.txt] then
        ignoreFlag = true
      end
      methodNameSet[methodNameToken.txt] = true
    end
    if (not ignoreFlag ) then
      filter( field, self, node, baseIndent )
    end
  end
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
  setmetatable( obj, { __index = %s } )
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
    
    if memberNode:get_getterMode(  ) ~= "none" and autoFlag then
      self:writeln( string.format( [==[
function %s:%s()
  return %s.%s
end]==], className, getterName, prefix, memberName), baseIndent )
      methodNameSet[getterName] = true
    end
    local setterName = "set_" .. memberName
    
    autoFlag = not methodNameSet[setterName]
    if memberNode:get_setterMode(  ) ~= "none" and autoFlag then
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
      if child:get_kind() == Ast.TypeInfoKindMethod and child:get_accessMode() ~= "pri" and not child:get_staticFlag() then
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
  
  local classNameToken = _lune_unwrap( declInfo:get_className(  ))
  
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
    if arg:get_kind(  ) == Ast.nodeKind.DeclArg then
      argTxt = argTxt .. (arg ):get_name().txt
    else 
      local name = _lune_unwrap( node:get_declInfo(  ):get_name())
      
      Util.err( string.format( "not support ... in macro -- %s", name.txt) )
    end
  end
  self:writeln( " )", baseIndent + stepIndent )
  self:writeln( "local obj = {}", baseIndent + stepIndent )
  self:writeln( string.format( "setmetatable( obj, { __index = %s } )", className ), baseIndent + stepIndent )
  self:writeln( string.format( "if obj.__init then obj:__init( %s ); end", argTxt ), baseIndent )
  self:writeln( "return obj", baseIndent )
  self:writeln( "end", baseIndent )
  self:write( string.format( "function %s:__init(%s) ", className, argTxt ) )
  do
    local _exp = declInfo:get_body(  )
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processExpCallSuper( node, parent, baseIndent )

  local typeInfo = node:get_superType(  )
  
  self:write( string.format( "%s.__init( self, ", typeInfo:getTxt(  )) )
  if node:get_expList(  ) then
    filter( node:get_expList(  ), self, node, baseIndent )
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
  local methodNodeToken = _lune_unwrap( declInfo:get_name(  ))
  
  local methodName = methodNodeToken.txt
  
  local classNameToken = _lune_unwrap( declInfo:get_className(  ))
  
  self:write( string.format( "function %s%s%s( ", classNameToken.txt, delimit, methodName) )
  local argList = declInfo:get_argList(  )
  
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
    end
    filter( arg, self, node, baseIndent )
  end
  self:writeln( " )", baseIndent )
  do
    local _exp = declInfo:get_body(  )
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
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
    filter( _lune_unwrap( node:get_unwrapBlock()), self, node, baseIndent + stepIndent )
  end
  self:writeln( "end", baseIndent )
end

function convFilter:processIfUnwrap( node, parent, baseIndent )

  self:writeln( "do", baseIndent + stepIndent )
  self:write( "local _exp = " )
  filter( node:get_exp(), self, node, baseIndent + stepIndent )
  self:writeln( "", baseIndent + stepIndent )
  self:writeln( "if _exp ~= nil then", baseIndent + stepIndent )
  filter( node:get_block(), self, node, baseIndent + stepIndent * 2 )
  if node:get_nilBlock() then
    self:writeln( "else", baseIndent + stepIndent )
    filter( _lune_unwrap( node:get_nilBlock()), self, node, baseIndent + stepIndent * 2 )
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
  if node:get_mode() ~= "unwrap" and node:get_accessMode(  ) ~= "global" then
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
        local thenBlock = node:get_thenBlock()
        
            if  nil == thenBlock then
              local _thenBlock = thenBlock
              
            else
              
                self:writeln( "else", baseIndent + stepIndent * 3 )
                filter( thenBlock, self, node, baseIndent + stepIndent * 3 )
              end
          
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
  
  if node:get_accessMode(  ) == "pub" then
    self:writeln( "", baseIndent )
    for index, var in pairs( varList ) do
      local name = var:get_name().txt
      
      self:writeln( string.format( "_moduleObj.%s = %s", name, name), baseIndent )
      self.pubVarName2InfoMap[name] = PubVerInfo.new(node:get_staticFlag(  ), node:get_accessMode(  ), node:get_typeInfoList(  )[index])
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
  
  if declInfo:get_accessMode(  ) ~= "global" and #name ~= 0 then
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
  self:writeln( " )", baseIndent )
  do
    local _exp = declInfo:get_body(  )
    if _exp ~= nil then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:writeln( "end", baseIndent )
  local expType = node:get_expType(  )
  
  if expType:get_accessMode(  ) == "pub" then
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
    
        self:write( '_lune_unwrapDefault( ' )
        filter( node:get_exp(), self, node, baseIndent )
        self:write( ', ' )
        filter( _exp, self, node, baseIndent )
        self:write( ')' )
      else
    
        self:write( '_lune_unwrap( ' )
        filter( node:get_exp(), self, node, baseIndent )
        self:write( ')' )
      end
  end
  
end

function convFilter:processExpCall( node, parent, baseIndent )

  local wroteFuncFlag = false
  
  
  if node:get_func():get_kind() == Ast.nodeKind.RefField then
    local fieldNode = node:get_func()
    
    local prefixNode = fieldNode:get_prefix()
    
    local prefixType = prefixNode:get_expType()
    
    if node:get_nilAccess() then
      wroteFuncFlag = true
      if prefixType:get_kind() == Ast.TypeInfoKindList then
        self:write( string.format( "_lune_nilacc( table.%s, nil, 'list', ", fieldNode:get_field().txt) )
        filter( prefixNode, self, fieldNode, baseIndent )
      else 
        self:write( "_lune_nilacc( " )
        filter( prefixNode, self, fieldNode, baseIndent )
        self:write( string.format( ", '%s', 'callmtd' ", fieldNode:get_field().txt) )
      end
    else 
      if prefixType:get_kind() == Ast.TypeInfoKindList then
        wroteFuncFlag = true
        self:write( string.format( "table.%s( ", fieldNode:get_field().txt) )
        filter( prefixNode, self, fieldNode, baseIndent )
      elseif prefixType:get_kind() == Ast.TypeInfoKindEnum then
        wroteFuncFlag = true
        local canonicalName = self:getCanonicalName( prefixType )
        
        self:write( string.format( "%s:_getTxt( ", canonicalName) )
        filter( prefixNode, self, fieldNode, baseIndent )
      end
    end
    
  end
  if not wroteFuncFlag then
    if node:get_nilAccess() then
      self:write( "_lune_nilacc( " )
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
    
        if wroteFuncFlag then
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
    if node:get_macroMode() == "expand" then
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

  if node:get_expType() == Ast.builtinTypeInt then
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
  
  if node:get_expType() == Ast.builtinTypeInt and node:get_op().txt == "/" then
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

  if node:get_symbolInfo():get_accessMode() == "pub" and node:get_symbolInfo():get_kind() == Ast.SymbolKind.Var then
    self:write( "_moduleObj." )
  end
  self:write( node:get_token().txt )
end

-- none

function convFilter:processExpRefItem( node, parent, baseIndent )

  if node:get_nilAccess() then
    self:write( "_lune_nilacc( " )
    filter( node:get_val(), self, node, baseIndent )
    self:write( ", nil, 'item', " )
    filter( node:get_index(), self, node, baseIndent )
    self:write( ")" )
  else 
    if node:get_val():get_expType() == Ast.builtinTypeString then
      self:write( "string.byte( " )
      filter( node:get_val(), self, node, baseIndent )
      self:write( ", " )
      filter( node:get_index(), self, node, baseIndent )
      self:write( " )" )
    else 
      filter( node:get_val(), self, node, baseIndent )
      self:write( "[" )
      filter( node:get_index(), self, node, baseIndent )
      self:write( "]" )
    end
  end
end

-- none

function convFilter:processRefField( node, parent, baseIndent )

  local prefix = node:get_prefix(  )
  
  if node:get_nilAccess() then
    self:write( '_lune_nilacc( ' )
    filter( prefix, self, node, baseIndent )
    self:write( string.format( ', "%s" )', node:get_field().txt) )
  else 
    filter( prefix, self, node, baseIndent )
    local delimit = "."
    
    if parent:get_kind() == Ast.nodeKind.ExpCall then
      if node:get_expType(  ):get_kind(  ) == Ast.TypeInfoKindMethod then
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

function convFilter:processGetField( node, parent, baseIndent )

  local prefixNode = node:get_prefix(  )
  
  local prefixType = prefixNode:get_expType()
  
  local fieldTxt = node:get_field(  ).txt
  
  if fieldTxt == "_txt" and prefixType:get_kind() == Ast.TypeInfoKindEnum then
    local canonicalName = self:getCanonicalName( prefixType )
    
    self:write( string.format( "%s:_getTxt( ", canonicalName) )
    filter( prefixNode, self, node, baseIndent )
    self:writeln( ")", baseIndent )
  else 
    filter( prefixNode, self, node, baseIndent )
    local delimit = "."
    
    if node:get_getterTypeInfo(  ):get_kind(  ) == Ast.TypeInfoKindMethod then
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

local MacroEvalImp = {}
setmetatable( MacroEvalImp, { __index = MacroEval } )
_moduleObj.MacroEvalImp = MacroEvalImp
function MacroEvalImp:eval( node )

  local oStream = Util.memStream.new()
  
  local conv = convFilter.new("macro", oStream, oStream, "exe", true, Ast.rootTypeInfo)
  
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
        return (_lune_unwrap( mod) )
      end
  end
  
  Util.err( "failed to load -- " .. node:get_declInfo():get_name() )
end
function MacroEvalImp.new( mode )
  local obj = {}
  setmetatable( obj, { __index = MacroEvalImp } )
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
