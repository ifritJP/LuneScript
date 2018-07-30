--lune/base/TransUnit.lns
local moduleObj = {}
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
   end
   error( string.format( "illegal access -- %s", access ) )
end




-- none

local Parser = require( 'lune.base.Parser' )

local Util = require( 'lune.base.Util' )

local Ast = require( 'lune.base.Ast' )

local TransUnit = {}
moduleObj.TransUnit = TransUnit
function TransUnit.new( macroEval )
  local obj = {}
  setmetatable( obj, { __index = TransUnit } )
  if obj.__init then obj:__init( macroEval ); end
return obj
end
function TransUnit:__init(macroEval) 
  self.validMutControl = true
  self.moduleName = ""
  self.parser = Parser.DummyParser.new()
  self.subfileList = {}
  self.pushbackList = {}
  self.usedTokenList = {}
  self.scope = Ast.rootScope
  self.typeId2ClassMap = {}
  self.typeInfo2ClassNode = {}
  self.currentToken = nil
  self.errMessList = {}
  self.macroEval = macroEval
  self.typeId2MacroInfo = {}
  self.macroMode = "none"
  self.symbol2ValueMapForMacro = {}
end
function TransUnit:addErrMess( pos, mess )
  table.insert( self.errMessList, string.format( "%s:%d:%d: %s", self.parser:getStreamName(  ), pos.lineNo, pos.column, mess) )
end
function TransUnit:pushScope( classFlag, inheritList )
  self.scope = Ast.Scope.new(self.scope, classFlag, inheritList or {})
  return self.scope
end
function TransUnit:popScope(  )
  self.scope = self.scope:get_parent(  )
end
function TransUnit:getCurrentClass(  )
  local typeInfo = Ast.rootTypeInfo
  
  local scope = self.scope
  
  repeat 
    do
      local _exp = scope:get_ownerTypeInfo()
      if _exp then
      
          if _exp:get_kind() == Ast.TypeInfoKindClass then
            return _exp
          end
        end
    end
    
    scope = scope:get_parent()
  until scope == Ast.rootScope
  return typeInfo
end
function TransUnit:getCurrentNamespaceTypeInfo(  )
  local typeInfo = Ast.rootTypeInfo
  
  local scope = self.scope
  
  repeat 
    do
      local _exp = scope:get_ownerTypeInfo()
      if _exp then
      
          return _exp
        end
    end
    
    scope = scope:get_parent()
  until scope == Ast.rootScope
  return typeInfo
end
function TransUnit:pushClass( baseInfo, externalFlag, name, accessMode, defNamespace )
  local typeInfo = Ast.rootTypeInfo
  
  do
    local _exp = self.scope:getTypeInfoChild( name )
    if _exp then
    
        typeInfo = _exp
        self.scope = typeInfo:get_scope() or _luneScript.error( 'unwrap val is nil' )
      else
    
        local parentInfo = self:getCurrentNamespaceTypeInfo(  )
        
        local inheritList = {}
        
        do
          local _exp = baseInfo
          if _exp then
          
              inheritList = {_exp:get_scope(  ) or _luneScript.error( 'unwrap val is nil' )}
            end
        end
        
        local scope = self:pushScope( true, inheritList )
        
        typeInfo = Ast.NormalTypeInfo.createClass( scope, baseInfo, parentInfo, externalFlag, accessMode, name )
        local parentScope = scope:get_parent(  )
        
        parentScope:addClass( name, typeInfo, scope, false )
      end
  end
  
  local namespace = defNamespace
  
      if  not namespace then
        local _namespace = namespace
        
        namespace = Ast.NamespaceInfo.new(name, self.scope, typeInfo)
      end
    
  self.typeId2ClassMap[typeInfo:get_typeId(  )] = namespace
  return typeInfo
end
function TransUnit:popClass(  )
  self:popScope(  )
end
-- none
-- none
-- none
-- none
-- none
-- none
-- none
-- none
-- none
function TransUnit:get_errMessList()
  return self.errMessList
end
do
  end

local opLevelBase = 0

local op2levelMap = {}

local op1levelMap = {}

local function regOpLevel( opnum, opList )
  opLevelBase = opLevelBase + 1
  if opnum == 1 then
    for __index, op in pairs( opList ) do
      op1levelMap[op] = opLevelBase
    end
  else 
    for __index, op in pairs( opList ) do
      op2levelMap[op] = opLevelBase
    end
  end
end

regOpLevel( 2, {"="} )
regOpLevel( 2, {"or"} )
regOpLevel( 2, {"and"} )
regOpLevel( 2, {"<", ">", "<=", ">=", "~=", "=="} )
regOpLevel( 2, {"|"} )
regOpLevel( 2, {"~"} )
regOpLevel( 2, {"&"} )
regOpLevel( 2, {"<<", ">>"} )
regOpLevel( 2, {".."} )
regOpLevel( 2, {"+", "-"} )
regOpLevel( 2, {"*", "/", "//", "%"} )
regOpLevel( 1, {"`", ",,", ",,,", ",,,,"} )
regOpLevel( 1, {"not", "#", "-", "~"} )
regOpLevel( 1, {"^"} )
local quotedChar2Code = {}

quotedChar2Code['a'] = 7
quotedChar2Code['b'] = 8
quotedChar2Code['t'] = 9
quotedChar2Code['n'] = 10
quotedChar2Code['v'] = 11
quotedChar2Code['f'] = 12
quotedChar2Code['r'] = 13
quotedChar2Code['\\'] = 92
quotedChar2Code['"'] = 34
quotedChar2Code["'"] = 39
local _TypeInfo = {}
function _TypeInfo.new( baseId, itemTypeId, argTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children, accessMode )
  local obj = {}
  setmetatable( obj, { __index = _TypeInfo } )
  if obj.__init then
    obj:__init( baseId, itemTypeId, argTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children, accessMode )
  end        
  return obj 
 end         
function _TypeInfo:__init( baseId, itemTypeId, argTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children, accessMode ) 
            
self.baseId = baseId
  self.itemTypeId = itemTypeId
  self.argTypeId = argTypeId
  self.retTypeId = retTypeId
  self.parentId = parentId
  self.typeId = typeId
  self.txt = txt
  self.kind = kind
  self.staticFlag = staticFlag
  self.nilable = nilable
  self.orgTypeId = orgTypeId
  self.children = children
  self.accessMode = accessMode
  end
do
  end

local _ModuleInfo = {}
function _ModuleInfo.new( _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap )
  local obj = {}
  setmetatable( obj, { __index = _ModuleInfo } )
  if obj.__init then
    obj:__init( _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap )
  end        
  return obj 
 end         
function _ModuleInfo:__init( _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap ) 
            
self._typeId2ClassInfoMap = _typeId2ClassInfoMap
  self._typeInfoList = _typeInfoList
  self._varName2InfoMap = _varName2InfoMap
  self._funcName2InfoMap = _funcName2InfoMap
  end
do
  end

local typeInfoListInsert = Ast.typeInfoRoot

moduleObj.typeInfoListInsert = typeInfoListInsert

local typeInfoListRemove = Ast.typeInfoRoot

moduleObj.typeInfoListRemove = typeInfoListRemove

function TransUnit:registBuiltInScope(  )
  local builtInInfo = {{[""] = {["type"] = {["arg"] = {"stem!"}, ["ret"] = {"str"}}, ["error"] = {["arg"] = {"str"}, ["ret"] = {}}, ["print"] = {["arg"] = {"..."}, ["ret"] = {}}, ["tonumber"] = {["arg"] = {"str"}, ["ret"] = {"real"}}, ["load"] = {["arg"] = {"str"}, ["ret"] = {"form!", "str"}}, ["require"] = {["arg"] = {"str"}, ["ret"] = {"stem!"}}, ["_fcall"] = {["arg"] = {"form", "..."}, ["ret"] = {""}}}}, {["ioStream"] = {["read"] = {["methodFlag"] = {}, ["arg"] = {"stem!"}, ["ret"] = {"str!"}}, ["close"] = {["methodFlag"] = {}, ["arg"] = {}, ["ret"] = {}}}}, {["io"] = {["open"] = {["arg"] = {"str", "str!"}, ["ret"] = {"ioStream!"}}, ["popen"] = {["arg"] = {"str"}, ["ret"] = {"ioStream!"}}}}, {["os"] = {["clock"] = {["arg"] = {}, ["ret"] = {"int"}}, ["exit"] = {["arg"] = {"int!"}, ["ret"] = {}}}}, {["string"] = {["find"] = {["arg"] = {"str", "str", "int!", "bool!"}, ["ret"] = {"int", "int"}}, ["byte"] = {["arg"] = {"str", "int"}, ["ret"] = {"int"}}, ["format"] = {["arg"] = {"str", "..."}, ["ret"] = {"str"}}, ["rep"] = {["arg"] = {"str", "int"}, ["ret"] = {"str"}}, ["gmatch"] = {["arg"] = {"str", "str"}, ["ret"] = {"stem!"}}, ["gsub"] = {["arg"] = {"str", "str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["arg"] = {"str", "int", "int!"}, ["ret"] = {"str"}}}}, {["str"] = {["find"] = {["methodFlag"] = {}, ["arg"] = {"str", "int!", "bool!"}, ["ret"] = {"int", "int"}}, ["byte"] = {["methodFlag"] = {}, ["arg"] = {"int"}, ["ret"] = {"int"}}, ["format"] = {["methodFlag"] = {}, ["arg"] = {"..."}, ["ret"] = {"str"}}, ["rep"] = {["methodFlag"] = {}, ["arg"] = {"int"}, ["ret"] = {"str"}}, ["gmatch"] = {["methodFlag"] = {}, ["arg"] = {"str"}, ["ret"] = {"stem!"}}, ["gsub"] = {["methodFlag"] = {}, ["arg"] = {"str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["methodFlag"] = {}, ["arg"] = {"int", "int!"}, ["ret"] = {"str"}}}}, {["table"] = {["unpack"] = {["arg"] = {"stem"}, ["ret"] = {"..."}}}}, {["List"] = {["insert"] = {["methodFlag"] = {}, ["arg"] = {"stem!"}, ["ret"] = {""}}, ["remove"] = {["methodFlag"] = {}, ["arg"] = {"int!"}, ["ret"] = {""}}}}, {["debug"] = {["getinfo"] = {["arg"] = {"int"}, ["ret"] = {"stem"}}}}, {["_luneScript"] = {["loadModule"] = {["arg"] = {"str"}, ["ret"] = {"stem"}}, ["searchModule"] = {["arg"] = {"str"}, ["ret"] = {"str!"}}}}}
  
  local function getTypeInfo( typeName )
    if typeName:find( "!$" ) then
      local orgTypeName = typeName:gsub( "!$", "" )
      
      local typeInfo = Ast.rootScope:getTypeInfo( orgTypeName, Ast.rootScope, false ) or _luneScript.error( 'unwrap val is nil' )
      
      return typeInfo:get_nilableTypeInfo()
    end
    return Ast.rootScope:getTypeInfo( typeName, Ast.rootScope, false ) or _luneScript.error( 'unwrap val is nil' )
  end
  
  local builtinModuleName2Scope = {}
  
  for __index, builtinClassInfo in pairs( builtInInfo ) do
    for name, name2FuncInfo in pairs( builtinClassInfo ) do
      local parentInfo = Ast.typeInfoRoot
      
      if name ~= "" then
        parentInfo = self:pushClass( nil, true, name, "pub" )
        Ast.builtInTypeIdSet[parentInfo:get_typeId(  )] = true
      end
      if not parentInfo then
        error( "parentInfo is nil" )
      end
      if not builtinModuleName2Scope[name] then
        if name ~= "" and getTypeInfo( name ) then
          builtinModuleName2Scope[name] = self.scope
        end
        do
          local __sorted = {}
          local __map = name2FuncInfo
          for __key in pairs( __map ) do
            table.insert( __sorted, __key )
          end
          table.sort( __sorted )
          for __index, funcName in ipairs( __sorted ) do
            info = __map[ funcName ]
            do
              local argTypeList = {}
              
              for __index, argType in pairs( info["arg"] or _luneScript.error( 'unwrap val is nil' ) ) do
                table.insert( argTypeList, getTypeInfo( argType ) )
              end
              local retTypeList = {}
              
              for __index, retType in pairs( info["ret"] or _luneScript.error( 'unwrap val is nil' ) ) do
                local retTypeInfo = getTypeInfo( retType )
                
                table.insert( retTypeList, retTypeInfo )
              end
              local methodFlag = info["methodFlag"]
              
              self:pushScope( false )
              local typeInfo = Ast.NormalTypeInfo.createFunc( true, self.scope, methodFlag and Ast.TypeInfoKindMethod or Ast.TypeInfoKindFunc, parentInfo, false, true, not methodFlag, "pub", funcName, argTypeList, retTypeList )
              
              self:popScope(  )
              Ast.builtInTypeIdSet[typeInfo:get_typeId(  )] = true
              self.scope:add( funcName, typeInfo, "pub", false )
              if methodFlag then
                do
                  local _switchExp = (name )
                  if _switchExp == "List" then
                    do
                      local _switchExp = (funcName )
                      if _switchExp == "insert" then
                        typeInfoListInsert = typeInfo
                      elseif _switchExp == "remove" then
                        typeInfoListRemove = typeInfo
                      end
                    end
                    
                  end
                end
                
              end
            end
          end
        end
        
      end
      if name ~= "" then
        self:popClass(  )
      end
    end
  end
end

function TransUnit:error( mess )
  local pos = Parser.Position.new(0, 0)
  
  local txt = ""
  
  do
    local _exp = self.currentToken
    if _exp then
    
        pos = _exp.pos
        txt = _exp.txt
      end
  end
  
  error( string.format( "error:%s:%d:%d:(%s) %s", self.parser:getStreamName(  ), pos.lineNo, pos.column, txt, mess ) )
end

function TransUnit:createNoneNode( pos )
  return Ast.NoneNode.new(pos, {Ast.builtinTypeNone})
end

function TransUnit:pushbackToken( token )
  table.insert( self.pushbackList, token )
  self.currentToken = self.usedTokenList[#self.usedTokenList]
end

local function expandVal( tokenList, val, pos )
  do
    local _exp = val
    if _exp then
    
        do
          local _switchExp = type( _exp )
          if _switchExp == "number" then
            local num = string.format( "%g", _exp)
            
            local kind = Parser.kind.Int
            
            if string.find( num, ".", 1, true ) then
              kind = Parser.kind.Real
            end
            table.insert( tokenList, Parser.Token.new(kind, num, pos) )
          elseif _switchExp == "string" then
            table.insert( tokenList, Parser.Token.new(Parser.kind.Str, string.format( '[[%s]]', _exp), pos) )
          elseif _switchExp == "table" then
            table.insert( tokenList, Parser.Token.new(Parser.kind.Dlmt, string.format( "{", _exp), pos) )
            for key, item in pairs( _exp ) do
              expandVal( tokenList, item, pos )
              table.insert( tokenList, Parser.Token.new(Parser.kind.Dlmt, string.format( ",", _exp), pos) )
            end
            table.insert( tokenList, Parser.Token.new(Parser.kind.Dlmt, string.format( "}", _exp), pos) )
          end
        end
        
      end
  end
  
end

function TransUnit:newPushback( tokenList )
  for index = #tokenList, 1, -1 do
    table.insert( self.pushbackList, tokenList[index] )
  end
  self.currentToken = self.usedTokenList[#self.usedTokenList]
end

function TransUnit:getTokenNoErr(  )
  if #self.pushbackList > 0 then
    table.insert( self.usedTokenList, self.currentToken )
    self.currentToken = self.pushbackList[#self.pushbackList]
    table.remove( self.pushbackList )
    return self.currentToken
  end
  local commentList = {}
  
  local token = nil
  
  while true do
    token = self.parser:getToken(  )
    do
      local _exp = token
      if _exp then
      
          if _exp.kind ~= Parser.kind.Cmnt then
            break
          end
          table.insert( commentList, _exp )
        else
      
          break
        end
    end
    
  end
  do
    local _exp = token
    if _exp then
    
        if self.macroMode == "expand" then
          local tokenTxt = _exp.txt
          
          if tokenTxt == ',,' or tokenTxt == ',,,,' then
            local nextToken = self:getTokenNoErr(  ) or _luneScript.error( 'unwrap val is nil' )
            
            local macroVal = self.symbol2ValueMapForMacro[nextToken.txt]
            
                if  not macroVal then
                  local _macroVal = macroVal
                  
                  self:error( string.format( "unknown macro val %s", nextToken.txt) )
                end
              
            if tokenTxt == ',,' then
              if macroVal.typeInfo == Ast.builtinTypeSymbol then
                local txtList = (macroVal.val or _luneScript.error( 'unwrap val is nil' ) )
                
                for index = #txtList, 1, -1 do
                  nextToken = Parser.Token.new(nextToken.kind, txtList[index], nextToken.pos)
                  self:pushbackToken( nextToken )
                end
              elseif macroVal.typeInfo == Ast.builtinTypeStat then
                self:pushbackStr( string.format( "macroVal %s", nextToken.txt), (macroVal.val or _luneScript.error( 'unwrap val is nil' ) ) )
              elseif macroVal.typeInfo:get_kind(  ) == Ast.TypeInfoKindArray or macroVal.typeInfo:get_kind(  ) == Ast.TypeInfoKindList then
                local strList = (macroVal.val or _luneScript.error( 'unwrap val is nil' ) )
                
                if strList then
                  for index = #strList, 1, -1 do
                    self:pushbackStr( string.format( "macroVal %s[%d]", nextToken.txt, index), strList[index] )
                  end
                else 
                  self:error( string.format( "macro val is nil %s", nextToken.txt) )
                end
              else 
                local tokenList = {}
                
                expandVal( tokenList, macroVal.val, nextToken.pos )
                self:newPushback( tokenList )
              end
            elseif tokenTxt == ',,,,' then
              if macroVal.typeInfo == Ast.builtinTypeSymbol then
                local txtList = (macroVal.val or _luneScript.error( 'unwrap val is nil' ) )
                
                local newToken = ""
                
                for __index, txt in pairs( txtList ) do
                  newToken = string.format( "%s%s", newToken, txt)
                end
                nextToken = Parser.Token.new(Parser.kind.Str, string.format( "'%s'", newToken), nextToken.pos)
                self:pushbackToken( nextToken )
              elseif macroVal.typeInfo == Ast.builtinTypeStat then
                nextToken = Parser.Token.new(Parser.kind.Str, string.format( "'%s'", macroVal.val or _luneScript.error( 'unwrap val is nil' )), nextToken.pos)
                self:pushbackToken( nextToken )
              else 
                self:error( string.format( "not support this symbol -- %s%s", tokenTxt, nextToken.txt) )
              end
            end
            nextToken = self:getTokenNoErr(  ) or _luneScript.error( 'unwrap val is nil' )
            token = nextToken
          end
        end
      end
  end
  
  do
    local _exp = token
    if _exp then
    
        _exp:set_commentList( commentList )
      end
  end
  
  table.insert( self.usedTokenList, self.currentToken )
  self.currentToken = token
  return token
end

function TransUnit:getToken(  )
  local token = self:getTokenNoErr(  )
  
      if  not token then
        local _token = token
        
        return Parser.getEofToken(  )
      end
    
  self.currentToken = token
  return token
end

function TransUnit:pushback(  )
  table.insert( self.pushbackList, self.currentToken )
  self.currentToken = self.usedTokenList[#self.usedTokenList]
  table.remove( self.usedTokenList )
end

function TransUnit:pushbackStr( name, statement )
  local memStream = Parser.TxtStream.new(statement)
  
  local parser = Parser.StreamParser.new(memStream, name, false)
  
  local list = {}
  
  while true do
    local token = parser:getToken(  )
    
    if not token then
      break
    end
    table.insert( list, token )
  end
  for index = #list, 1, -1 do
    self:pushbackToken( list[index] )
  end
end

function TransUnit:checkSymbol( token )
  if token.kind ~= Parser.kind.Symb and token.kind ~= Parser.kind.Kywd and token.kind ~= Parser.kind.Type then
    self:error( "illegal symbol" )
  end
  return token
end

function TransUnit:getSymbolToken(  )
  return self:checkSymbol( self:getToken(  ) )
end

function TransUnit:checkToken( token, txt )
  if not token or token.txt ~= txt then
    self:error( string.format( "not found -- %s", txt ) )
  end
  return token
end

function TransUnit:checkNextToken( txt )
  return self:checkToken( self:getToken(  ), txt )
end

function TransUnit:getContinueToken(  )
  local prevToken = self.currentToken
  
      if  not prevToken then
        local _prevToken = prevToken
        
        return self:getToken(  ), false
      end
    
  local token = self:getToken(  )
  
  if prevToken.pos.lineNo ~= token.pos.lineNo or prevToken.pos.column + #prevToken.txt ~= token.pos.column then
    return token, false
  end
  return token, true
end

function TransUnit:analyzeStatementList( stmtList, termTxt )
  while true do
    local statement = self:analyzeStatement( termTxt )
    
    do
      local _exp = statement
      if _exp then
      
          table.insert( stmtList, _exp )
        else
      
          break
        end
    end
    
  end
end

function TransUnit:analyzeStatementListSubfile( stmtList )
  local statement = self:analyzeStatement(  )
  
  do
    local _exp = statement
    if _exp then
    
        if _exp:get_kind() ~= Ast.nodeKindSubfile then
          self:error( "subfile must have 'subfile' declaration at top." )
        end
      end
  end
  
  self:analyzeStatementList( stmtList )
end

function TransUnit:analyzeLuneControl( firstToken )
  local nextToken = self:getToken(  )
  
  do
    local _switchExp = (nextToken.txt )
    if _switchExp == "disable_mut_control" then
      self.validMutControl = false
    else 
      self:addErrMess( nextToken.pos, string.format( "unknown option -- %s", nextToken.txt) )
    end
  end
  
  self:checkNextToken( ";" )
end

function TransUnit:analyzeBlock( blockKind, scope )
  local token = self:checkNextToken( "{" )
  
  if not scope then
    self:pushScope( false )
  end
  local stmtList = {}
  
  self:analyzeStatementList( stmtList, "}" )
  self:checkNextToken( "}" )
  if not scope then
    self:popScope(  )
  end
  local node = Ast.BlockNode.new(token.pos, {Ast.builtinTypeNone}, blockKind, stmtList)
  
  return node
end

function TransUnit:analyzeImport( token )
  if self.moduleScope ~= self.scope then
    self:error( "'import' must call at top scope." )
  end
  self.scope = Ast.rootScope
  local moduleToken = self:getToken(  )
  
  local modulePath = moduleToken.txt
  
  local nextToken = moduleToken
  
  local nameList = {moduleToken.txt}
  
  while true do
    nextToken = self:getToken(  )
    if nextToken.txt == "." then
      nextToken = self:getToken(  )
      moduleToken = nextToken
      modulePath = string.format( "%s.%s", modulePath, moduleToken.txt)
      table.insert( nameList, moduleToken.txt )
    else 
      break
    end
  end
  local typeId2TypeInfo = {}
  
  typeId2TypeInfo[Ast.rootTypeId] = Ast.typeInfoRoot
  local moduleTypeInfo = Ast.rootTypeInfo
  
  for __index, moduleName in pairs( nameList ) do
    moduleTypeInfo = self:pushClass( nil, true, moduleName, "pub" )
  end
  for __index, moduleName in pairs( nameList ) do
    self:popClass(  )
  end
  local moduleInfo = _luneScript.loadModule( modulePath )
  
  for __index, symbolInfo in pairs( Ast.sym2builtInTypeMap ) do
    typeId2TypeInfo[symbolInfo:get_typeInfo():get_typeId(  )] = symbolInfo:get_typeInfo()
  end
  local typeId2Scope = {}
  
  typeId2Scope[Ast.rootTypeId] = self.scope
  local newId2OldIdMap = {}
  
  local function registTypeInfo( atomInfo )
    local newTypeInfo = nil
    
    if not Ast.builtInTypeIdSet[atomInfo.typeId] then
      if atomInfo.nilable then
        local orgTypeInfo = typeId2TypeInfo[atomInfo.orgTypeId] or _luneScript.error( 'unwrap val is nil' )
        
        newTypeInfo = orgTypeInfo:get_nilableTypeInfo(  )
        typeId2TypeInfo[atomInfo.typeId] = newTypeInfo or _luneScript.error( 'unwrap val is nil' )
      else 
        local itemTypeInfo = {}
        
        for __index, typeId in pairs( atomInfo.itemTypeId ) do
          table.insert( itemTypeInfo, typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' ) )
        end
        local argTypeInfo = {}
        
        for __index, typeId in pairs( atomInfo.argTypeId ) do
          if not typeId2TypeInfo[typeId] then
            Util.errorLog( string.format( "not found -- %s,%d", atomInfo.txt, #atomInfo.argTypeId) )
          end
          table.insert( argTypeInfo, typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' ) )
        end
        local retTypeInfo = {}
        
        for __index, typeId in pairs( atomInfo.retTypeId ) do
          table.insert( retTypeInfo, typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' ) )
        end
        local parentInfo = Ast.typeInfoRoot
        
        if atomInfo.parentId ~= Ast.rootTypeId then
          local workTypeInfo = typeId2TypeInfo[atomInfo.parentId]
          
              if  not workTypeInfo then
                local _workTypeInfo = workTypeInfo
                
                error( string.format( "not found parentInfo %s %s", atomInfo.parentId, atomInfo.txt) )
              end
            
          parentInfo = workTypeInfo
        end
        local baseInfo = typeId2TypeInfo[atomInfo.baseId] or _luneScript.error( 'unwrap val is nil' )
        
        local parentScope = typeId2Scope[atomInfo.parentId]
        
            if  not parentScope then
              local _parentScope = parentScope
              
              self:error( string.format( "not found parentScope %s %s", atomInfo.parentId, atomInfo.txt) )
            end
          
        if atomInfo.txt ~= "" then
          newTypeInfo = parentScope:getTypeInfoChild( atomInfo.txt )
        end
        if newTypeInfo and atomInfo.kind == Ast.TypeInfoKindClass then
          local newTypeInfo = newTypeInfo
          
              if  not newTypeInfo then
                local _newTypeInfo = newTypeInfo
                
              else
                
                  typeId2Scope[atomInfo.typeId] = newTypeInfo:get_scope()
                  if not newTypeInfo:get_scope() then
                    error( string.format( "not found scope %s %s %s %s %s", parentScope, atomInfo.parentId, atomInfo.typeId, atomInfo.txt, newTypeInfo:getTxt(  )) )
                  end
                  typeId2TypeInfo[atomInfo.typeId] = newTypeInfo
                end
            
        else 
          if atomInfo.kind == Ast.TypeInfoKindClass then
            local baseScope = typeId2Scope[atomInfo.baseId] or _luneScript.error( 'unwrap val is nil' )
            
            local scope = Ast.Scope.new(parentScope, true, baseScope and {baseScope} or {})
            
            local workTypeInfo = Ast.NormalTypeInfo.createClass( scope, baseInfo, parentInfo, true, "pub", atomInfo.txt )
            
            newTypeInfo = workTypeInfo
            typeId2Scope[atomInfo.typeId] = scope
            typeId2TypeInfo[atomInfo.typeId] = workTypeInfo
            parentScope:addClass( atomInfo.txt, workTypeInfo, scope, false )
          else 
            local scope = nil
            
            if atomInfo.kind == Ast.TypeInfoKindFunc or atomInfo.kind == Ast.TypeInfoKindMethod then
              scope = Ast.Scope.new(parentScope, false, {})
            end
            local workTypeInfo = Ast.NormalTypeInfo.create( scope, baseInfo, parentInfo, atomInfo.staticFlag, atomInfo.kind, atomInfo.txt, itemTypeInfo, argTypeInfo, retTypeInfo )
            
            newTypeInfo = workTypeInfo
            typeId2TypeInfo[atomInfo.typeId] = workTypeInfo
            if atomInfo.kind == Ast.TypeInfoKindFunc or atomInfo.kind == Ast.TypeInfoKindMethod then
              (typeId2Scope[atomInfo.parentId] or _luneScript.error( 'unwrap val is nil' ) ):add( atomInfo.txt, workTypeInfo, atomInfo.accessMode, false )
              typeId2Scope[atomInfo.typeId] = scope
            end
          end
        end
      end
    else 
      newTypeInfo = Ast.rootScope:getTypeInfo( atomInfo.txt, Ast.rootScope, false )
      typeId2TypeInfo[atomInfo.typeId] = newTypeInfo or _luneScript.error( 'unwrap val is nil' )
    end
    return newTypeInfo or _luneScript.error( 'unwrap val is nil' )
  end
  
  for __index, atomInfo in pairs( moduleInfo._typeInfoList ) do
    registTypeInfo( atomInfo )
  end
  for __index, atomInfo in pairs( moduleInfo._typeInfoList ) do
    if atomInfo.children and #atomInfo.children > 0 then
      local scope = typeId2Scope[atomInfo.typeId] or _luneScript.error( 'unwrap val is nil' )
      
      for __index, childId in pairs( atomInfo.children ) do
        local typeInfo = typeId2TypeInfo[childId] or _luneScript.error( 'unwrap val is nil' )
        
        scope:add( typeInfo:getTxt(  ), typeInfo, typeInfo:get_accessMode(), typeInfo:get_mutable() )
      end
    end
  end
  for typeId, typeInfo in pairs( typeId2TypeInfo ) do
    newId2OldIdMap[typeInfo:get_typeId(  )] = typeId
  end
  local function registMember( classTypeId )
    local classTypeInfo = typeId2TypeInfo[classTypeId] or _luneScript.error( 'unwrap val is nil' )
    
    self:pushClass( nil, true, classTypeInfo:getTxt(  ), "pub" )
    do
      local _exp = moduleInfo._typeId2ClassInfoMap[classTypeId]
      if _exp then
      
          local classInfo = _exp
          
          for fieldName, fieldInfo in pairs( classInfo ) do
            local typeId = fieldInfo.typeId
            
            local fieldTypeInfo = typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' )
            
            self.scope:add( fieldName, fieldTypeInfo, (fieldInfo.accessMode or _luneScript.error( 'unwrap val is nil' ) ), (fieldInfo.mutable or false ) )
          end
          for __index, child in pairs( classTypeInfo:get_children(  ) ) do
            if child:get_kind(  ) == Ast.TypeInfoKindClass then
              local oldId = newId2OldIdMap[child:get_typeId(  )]
              
              if oldId then
                registMember( oldId or _luneScript.error( 'unwrap val is nil' ) )
              end
            end
          end
        else
      
          self:error( string.format( "not found class -- %d, %s", classTypeId, classTypeInfo:getTxt(  )) )
        end
    end
    
    self:popClass(  )
  end
  
  for __index, atomInfo in pairs( moduleInfo._typeInfoList ) do
    if atomInfo.parentId == Ast.rootTypeId and atomInfo.kind == Ast.TypeInfoKindClass then
      registMember( atomInfo.typeId )
    end
  end
  for __index, moduleName in pairs( nameList ) do
    self:pushClass( nil, true, moduleName, "pub" )
  end
  for varName, varInfo in pairs( moduleInfo._varName2InfoMap ) do
    self.scope:add( varName, typeId2TypeInfo[varInfo.typeId] or _luneScript.error( 'unwrap val is nil' ), "pub", (varInfo.mutable or false ) )
  end
  for __index, moduleName in pairs( nameList ) do
    self:popClass(  )
  end
  self.scope = self.moduleScope
  self.scope:add( moduleToken.txt, moduleTypeInfo, "local", false )
  self:checkToken( nextToken, ";" )
  if self.moduleScope ~= self.scope then
    self:error( "illegal top scope." )
  end
  return Ast.ImportNode.new(token.pos, {Ast.builtinTypeNone}, modulePath)
end

function TransUnit:analyzeSubfile( token )
  if self.scope ~= self.moduleScope then
    self:error( "'module' must be top scope." )
  end
  local mode = self:getToken(  )
  
  local moduleName = ""
  
  while true do
    local nextToken = self:getToken(  )
    
    if nextToken.txt == ";" then
      break
    end
    if moduleName == "" then
      moduleName = nextToken.txt
    else 
      moduleName = string.format( "%s%s", moduleName, nextToken.txt)
    end
  end
  if moduleName == "" then
    self:addErrMess( token.pos, "illegal subfile" )
  else 
    if mode.txt == "use" then
      do
        local _exp = _luneScript.searchModule( moduleName )
        if _exp then
        
            table.insert( self.subfileList, _exp )
          else
        
            self:addErrMess( token.pos, string.format( "not found subfile -- %s", moduleName) )
          end
      end
      
    elseif mode.txt == "owner" then
      if self.moduleName ~= moduleName then
        self:addErrMess( token.pos, string.format( "illegal owner module -- %s, %s", moduleName, self.moduleName) )
      end
    else 
      self:addErrMess( mode.pos, string.format( "illegal module mode -- %s", mode.txt) )
    end
  end
  return Ast.SubfileNode.new(token.pos, {Ast.builtinTypeNone})
end

function TransUnit:analyzeIfUnwrap( firstToken )
  local exp = self:analyzeExp(  )
  
  local scope = self:pushScope(  )
  
  local expType = exp:get_expType()
  
  if not expType:get_nilable() then
    self:addErrMess( exp:get_pos(), "this is not nilable" )
    scope:add( "_exp", expType, "local", false )
  else 
    scope:add( "_exp", expType:get_orgTypeInfo(), "local", false )
  end
  local block = self:analyzeBlock( "if!", scope )
  
  self:popScope(  )
  local elseBlock = nil
  
  local nextToken = self:getToken(  )
  
  if nextToken.txt == "else" then
    elseBlock = self:analyzeBlock( "if!" )
  else 
    self:pushback(  )
  end
  return Ast.IfUnwrapNode.new(firstToken.pos, {Ast.builtinTypeNone}, exp, block, elseBlock)
end

function TransUnit:analyzeIf( token )
  local nextToken, continueFlag = self:getContinueToken(  )
  
  if continueFlag and nextToken.txt == "!" then
    return self:analyzeIfUnwrap( token )
  end
  self:pushback(  )
  local list = {}
  
  table.insert( list, Ast.IfStmtInfo.new("if", self:analyzeExp(  ), self:analyzeBlock( "if" )) )
  nextToken = self:getToken(  )
  if nextToken.txt == "elseif" then
    while nextToken.txt == "elseif" do
      table.insert( list, Ast.IfStmtInfo.new("elseif", self:analyzeExp(  ), self:analyzeBlock( "elseif" )) )
      nextToken = self:getToken(  )
    end
  end
  if nextToken.txt == "else" then
    table.insert( list, Ast.IfStmtInfo.new("else", Ast.NoneNode.new(nextToken.pos, {Ast.builtinTypeNone}), self:analyzeBlock( "else" )) )
  else 
    self:pushback(  )
  end
  return Ast.IfNode.new(token.pos, {Ast.builtinTypeNone}, list)
end

function TransUnit:analyzeSwitch( firstToken )
  local exp = self:analyzeExp(  )
  
  self:checkNextToken( "{" )
  local caseList = {}
  
  local nextToken = self:getToken(  )
  
  while (nextToken.txt == "case" ) do
    self:checkToken( nextToken, "case" )
    local condexpList = self:analyzeExpList(  )
    
    local condBock = self:analyzeBlock( "switch" )
    
    table.insert( caseList, Ast.CaseInfo.new(condexpList, condBock) )
    nextToken = self:getToken(  )
  end
  local defaultBlock = nil
  
  if nextToken.txt == "default" then
    defaultBlock = self:analyzeBlock( "default" )
  else 
    self:pushback(  )
  end
  self:checkNextToken( "}" )
  return Ast.SwitchNode.new(firstToken.pos, {Ast.builtinTypeNone}, exp, caseList, defaultBlock)
end

function TransUnit:analyzeWhile( token )
  return Ast.WhileNode.new(token.pos, {Ast.builtinTypeNone}, self:analyzeExp(  ), self:analyzeBlock( "while" ))
end

function TransUnit:analyzeRepeat( token )
  local scope = self:pushScope( false )
  
  local node = Ast.RepeatNode.new(token.pos, {Ast.builtinTypeNone}, self:analyzeBlock( "repeat", scope ), self:analyzeExp(  ))
  
  self:popScope(  )
  self:checkNextToken( ";" )
  return node
end

function TransUnit:analyzeFor( firstToken )
  local scope = self:pushScope( false )
  
  local val = self:getToken(  )
  
  if val.kind ~= Parser.kind.Symb then
    self:error( "not symbol" )
  end
  self:checkNextToken( "=" )
  local exp1 = self:analyzeExp(  )
  
  if exp1:get_expType() ~= Ast.builtinTypeInt then
    self:addErrMess( exp1:get_pos(), string.format( "exp1 is not int -- %s", exp1:get_expType():getTxt(  )) )
  end
  self.scope:add( val.txt, exp1:get_expType(), "local", false )
  self:checkNextToken( "," )
  local exp2 = self:analyzeExp(  )
  
  if exp2:get_expType() ~= Ast.builtinTypeInt then
    self:addErrMess( exp2:get_pos(), string.format( "exp2 is not int -- %s", exp2:get_expType():getTxt(  )) )
  end
  local token = self:getToken(  )
  
  local exp3 = nil
  
  if token.txt == "," then
    exp3 = self:analyzeExp(  )
    do
      local _exp = exp3
      if _exp then
      
          if _exp:get_expType() ~= Ast.builtinTypeInt then
            self:addErrMess( _exp:get_pos(), string.format( "exp is not int -- %s", _exp:get_expType():getTxt(  )) )
          end
        end
    end
    
  else 
    self:pushback(  )
  end
  local node = Ast.ForNode.new(firstToken.pos, {Ast.builtinTypeNone}, self:analyzeBlock( "for", scope ), val, exp1, exp2, exp3)
  
  self:popScope(  )
  return node
end

function TransUnit:analyzeApply( token )
  local scope = self:pushScope(  )
  
  local varList = {}
  
  local nextToken = Parser.getEofToken(  )
  
  repeat 
    local var = self:getSymbolToken(  )
    
    if var.kind ~= Parser.kind.Symb then
      self:error( "illegal symbol" )
    end
    table.insert( varList, var )
    nextToken = self:getToken(  )
    scope:add( var.txt, Ast.builtinTypeStem, "local", false )
  until nextToken.txt ~= ","
  self:checkToken( nextToken, "of" )
  local exp = self:analyzeExp(  )
  
  if exp:get_kind() ~= Ast.nodeKindExpCall then
    self:error( "not call" )
  end
  local block = self:analyzeBlock( "apply", scope )
  
  self:popScope(  )
  return Ast.ApplyNode.new(token.pos, {Ast.builtinTypeNone}, varList, exp, block)
end

function TransUnit:analyzeForeach( token, sortFlag )
  local scope = self:pushScope(  )
  
  local valSymbol = Parser.getEofToken(  )
  
  local keySymbol = nil
  
  local nextToken = Parser.getEofToken(  )
  
  for index = 1, 2 do
    local sym = self:getToken(  )
    
    if sym.kind ~= Parser.kind.Symb then
      self:error( "illegal symbol" )
    end
    if index == 1 then
      valSymbol = sym
    else 
      keySymbol = sym
    end
    nextToken = self:getToken(  )
    if nextToken.txt ~= "," then
      break
    end
  end
  self:checkToken( nextToken, "in" )
  local exp = self:analyzeExp(  )
  
  if not exp:get_expType() then
    self:error( string.format( "unknown type of exp -- %d:%d", token.pos.lineNo, token.pos.column) )
  else 
    local itemTypeInfoList = exp:get_expType():get_itemTypeInfoList(  )
    
    if exp:get_expType():get_kind(  ) == Ast.TypeInfoKindMap then
      self.scope:add( valSymbol.txt, itemTypeInfoList[2], "local", false )
      do
        local _exp = keySymbol
        if _exp then
        
            self.scope:add( _exp.txt, itemTypeInfoList[1], "local", false )
          end
      end
      
    elseif exp:get_expType():get_kind(  ) == Ast.TypeInfoKindList or exp:get_expType():get_kind(  ) == Ast.TypeInfoKindArray then
      self.scope:add( valSymbol.txt, itemTypeInfoList[1], "local", false )
      do
        local _exp = keySymbol
        if _exp then
        
            self.scope:add( _exp.txt, Ast.builtinTypeInt, "local", false )
          else
        
            self.scope:add( "__index", Ast.builtinTypeInt, "local", false )
          end
      end
      
    else 
      self:error( string.format( "unknown kind type of exp for foreach-- %d:%d", exp:get_pos().lineNo, exp:get_pos().column) )
    end
  end
  local block = self:analyzeBlock( "foreach", scope )
  
  self:popScope(  )
  if sortFlag then
    return Ast.ForsortNode.new(token.pos, {Ast.builtinTypeNone}, valSymbol, keySymbol, exp, block, sortFlag)
  else 
    return Ast.ForeachNode.new(token.pos, {Ast.builtinTypeNone}, valSymbol, keySymbol, exp, block, sortFlag)
  end
end

function TransUnit:analyzeRefType( accessMode )
  local firstToken = self:getToken(  )
  
  local token = firstToken
  
  local refFlag = false
  
  if token.txt == "&" then
    refFlag = true
    token = self:getToken(  )
  end
  local mutFlag = false
  
  if token.txt == "mut" then
    mutFlag = true
    token = self:getToken(  )
  end
  local typeInfo = Ast.builtinTypeStem_
  
  self:checkSymbol( token )
  local name = self:analyzeExpSymbol( firstToken, token, "symbol", nil, true )
  
  typeInfo = name:get_expType()
  local continueToken, continueFlag = self:getContinueToken(  )
  
  token = continueToken
  if continueFlag and token.txt == "!" then
    typeInfo = typeInfo:get_nilableTypeInfo(  ) or _luneScript.error( 'unwrap val is nil' )
    token = self:getToken(  )
  end
  local arrayMode = "no"
  
  while true do
    if token.txt == '[' or token.txt == '[@' then
      if token.txt == '[' then
        arrayMode = "list"
        typeInfo = Ast.NormalTypeInfo.createList( accessMode, self:getCurrentClass(  ), {typeInfo} )
      else 
        arrayMode = "array"
        typeInfo = Ast.NormalTypeInfo.createArray( accessMode, self:getCurrentClass(  ), {typeInfo} )
      end
      token = self:getToken(  )
      if token.txt ~= ']' then
        self:pushback(  )
        self:checkNextToken( ']' )
      end
    elseif token.txt == "<" then
      local genericList = {}
      
      local nextToken = Parser.getEofToken(  )
      
      repeat 
        local typeExp = self:analyzeRefType( accessMode )
        
        table.insert( genericList, typeExp:get_expType() )
        nextToken = self:getToken(  )
      until nextToken.txt ~= ","
      self:checkToken( nextToken, '>' )
      if typeInfo:get_kind() == Ast.TypeInfoKindMap then
        typeInfo = Ast.NormalTypeInfo.createMap( accessMode, self:getCurrentClass(  ), genericList[1] or Ast.builtinTypeStem, genericList[2] or Ast.builtinTypeStem )
      else 
        self:error( string.format( "not support generic: %s", typeInfo:getTxt(  ) ) )
      end
    else 
      self:pushback(  )
      break
    end
    token = self:getToken(  )
  end
  if token.txt == "!" then
    typeInfo = typeInfo:get_nilableTypeInfo(  ) or _luneScript.error( 'unwrap val is nil' )
    token = self:getToken(  )
  end
  return Ast.RefTypeNode.new(firstToken.pos, {typeInfo}, name, refFlag, mutFlag, arrayMode)
end

function TransUnit:analyzeDeclArgList( accessMode, argList )
  local token = Parser.noneToken
  
  repeat 
    local argName = self:getToken(  )
    
    if argName.txt == ")" then
      token = argName
      break
    elseif argName.txt == "..." then
      table.insert( argList, Ast.DeclArgDDDNode.new(argName.pos, {Ast.builtinTypeDDD}) )
    else 
      argName = self:checkSymbol( argName )
      self:checkNextToken( ":" )
      local refType = self:analyzeRefType( accessMode )
      
      local arg = Ast.DeclArgNode.new(argName.pos, refType:get_expTypeList(), argName, refType)
      
      self.scope:add( argName.txt, refType:get_expType(), "local", false )
      table.insert( argList, arg )
    end
    token = self:getToken(  )
  until token.txt ~= ","
  self:checkToken( token, ")" )
  return token
end

local ASTInfo = {}
moduleObj.ASTInfo = ASTInfo
function ASTInfo.new( node, moduleTypeInfo )
  local obj = {}
  setmetatable( obj, { __index = ASTInfo } )
  if obj.__init then
    obj:__init( node, moduleTypeInfo )
  end        
  return obj 
 end         
function ASTInfo:__init( node, moduleTypeInfo ) 
            
self.node = node
  self.moduleTypeInfo = moduleTypeInfo
  end
function ASTInfo:get_node()
  return self.node
end
function ASTInfo:get_moduleTypeInfo()
  return self.moduleTypeInfo
end
do
  end

function TransUnit:createAST( parser, macroFlag, module )
  self.moduleName = module
  self:registBuiltInScope(  )
  local moduleTypeInfo = Ast.typeInfoRoot
  
  if module then
    for txt in string.gmatch( module, '[^%.]+' ) do
      moduleTypeInfo = self:pushClass( nil, false, txt, "pub" ) or _luneScript.error( 'unwrap val is nil' )
    end
  end
  self.moduleScope = self.scope
  self.parser = parser
  local ast = nil
  
  if macroFlag then
    ast = self:analyzeBlock( "macro" )
  else 
    local children = {}
    
    ast = Ast.RootNode.new(Parser.Position.new(0, 0), {Ast.builtinTypeNone}, children, self.typeId2ClassMap)
    self:analyzeStatementList( children )
    local token = self:getTokenNoErr(  )
    
    do
      local _exp = token
      if _exp then
      
          error( string.format( "unknown:%d:%d:(%s) %s", _exp.pos.lineNo, _exp.pos.column, Parser.getKindTxt( _exp.kind ), _exp.txt) )
        end
    end
    
    for __index, file in pairs( self.subfileList ) do
      if self.scope ~= self.moduleScope then
        self:error( "scope does not close" )
      end
      do
        local _exp = Parser.StreamParser.create( file, false )
        if _exp then
        
            self.parser = _exp
            self:analyzeStatementListSubfile( children )
            token = self:getTokenNoErr(  )
            do
              local _exp = token
              if _exp then
              
                  error( string.format( "unknown:%d:%d:(%s) %s", _exp.pos.lineNo, _exp.pos.column, Parser.getKindTxt( _exp.kind ), _exp.txt) )
                end
            end
            
          else
        
            self:error( string.format( "open error -- %s", file) )
          end
      end
      
    end
  end
  if module then
    for txt in string.gmatch( module, '[^%.]+' ) do
      self:popClass(  )
    end
  end
  if #self.errMessList > 0 then
    for __index, mess in pairs( self.errMessList ) do
      Util.errorLog( "error:" .. mess )
    end
    error( "has error" )
  end
  return ASTInfo.new(ast or _luneScript.error( 'unwrap val is nil' ), moduleTypeInfo)
end

function TransUnit:analyzeDeclMacro( accessMode, firstToken )
  local nameToken = self:getToken(  )
  
  self:checkNextToken( "(" )
  local scope = self:pushScope(  )
  
  local argList = {}
  
  local nextToken = self:analyzeDeclArgList( accessMode, argList )
  
  local argTypeList = {}
  
  for __index, argNode in pairs( argList ) do
    table.insert( argTypeList, argNode:get_expType() )
  end
  self:checkNextToken( "{" )
  nextToken = self:getToken(  )
  local ast = nil
  
  if nextToken.txt == "{" then
    local parser = Parser.WrapParser.new(self.parser, string.format( "decl macro %s", nameToken.txt))
    
    self.macroScope = scope
    local bakParser = self.parser
    
    self.parser = parser
    local stmtList = {}
    
    self:analyzeStatementList( stmtList, "}" )
    self:checkNextToken( "}" )
    self.parser = bakParser
    self.macroScope = nil
    ast = Ast.BlockNode.new(firstToken.pos, {Ast.builtinTypeNone}, "macro", stmtList)
  else 
    self:pushback(  )
  end
  self:popScope(  )
  local tokenList = {}
  
  local braceCount = 0
  
  while true do
    nextToken = self:getToken(  )
    if nextToken.txt == "{" then
      braceCount = braceCount + 1
    elseif nextToken.txt == "}" then
      if braceCount == 0 then
        break
      end
      braceCount = braceCount - 1
    end
    table.insert( tokenList, nextToken )
  end
  local typeInfo = Ast.NormalTypeInfo.createFunc( false, scope, Ast.TypeInfoKindMacro, self:getCurrentNamespaceTypeInfo(  ), false, false, false, accessMode, nameToken.txt, argTypeList )
  
  self.scope:add( nameToken.txt, typeInfo, "local", false )
  local declMacroInfo = Ast.DeclMacroInfo.new(nameToken, argList, ast, tokenList)
  
  local node = Ast.DeclMacroNode.new(firstToken.pos, {typeInfo}, declMacroInfo)
  
  local macroObj = self.macroEval:eval( node )
  
  self.typeId2MacroInfo[typeInfo:get_typeId(  )] = Ast.MacroInfo.new(macroObj, declMacroInfo, self.symbol2ValueMapForMacro)
  self.symbol2ValueMapForMacro = {}
  return node
end

function TransUnit:analyzeDeclProto( accessMode, firstToken )
  local nextToken = self:getToken(  )
  
  if nextToken.txt == "class" then
    local name = self:getToken(  )
    
    nextToken = self:getToken(  )
    local baseRef = nil
    
    if nextToken.txt == "extend" then
      baseRef = self:analyzeRefType( accessMode )
      nextToken = self:getToken(  )
    end
    self:checkToken( nextToken, ";" )
    local typeInfo = nil
    
    do
      local _exp = baseRef
      if _exp then
      
          typeInfo = _exp:get_expType(  )
        end
    end
    
    self:pushClass( typeInfo, false, name.txt, accessMode )
    self:popClass(  )
  else 
    self:error( "illegal proto" )
  end
  return self:createNoneNode( firstToken.pos )
end

function TransUnit:analyzeDecl( accessMode, staticFlag, firstToken, token )
  if not staticFlag then
    if token.txt == "static" then
      staticFlag = true
      token = self:getToken(  )
    end
  end
  local overrideFlag = false
  
  if token.txt == "override" then
    overrideFlag = true
    token = self:getToken(  )
  end
  if token.txt == "let" then
    return self:analyzeDeclVar( "let", accessMode, staticFlag, firstToken )
  elseif token.txt == "fn" then
    return self:analyzeDeclFunc( overrideFlag, accessMode, staticFlag, nil, firstToken, nil )
  elseif token.txt == "class" then
    return self:analyzeDeclClass( accessMode, firstToken, true )
  elseif token.txt == "module" then
    return self:analyzeDeclClass( accessMode, firstToken, false )
  elseif token.txt == "proto" then
    return self:analyzeDeclProto( accessMode, firstToken )
  elseif token.txt == "macro" then
    return self:analyzeDeclMacro( accessMode, firstToken )
  end
  return nil
end

function TransUnit:analyzeDeclMember( accessMode, staticFlag, firstToken )
  local varName = self:getSymbolToken(  )
  
  local token = self:getToken(  )
  
  local refType = self:analyzeRefType( accessMode )
  
  token = self:getToken(  )
  local getterMode = "none"
  
  local setterMode = "none"
  
  if token.txt == "{" then
    local nextToken = self:getToken(  )
    
    if nextToken.txt == "pub" or nextToken.txt == "pri" then
      getterMode = nextToken.txt
      nextToken = self:getToken(  )
      if nextToken.txt == "," then
        nextToken = self:getToken(  )
        if nextToken.txt == "pub" or nextToken.txt == "pri" then
          setterMode = nextToken.txt
          nextToken = self:getToken(  )
        end
      end
    end
    self:checkToken( nextToken, "}" )
    token = self:getToken(  )
  end
  self:checkToken( token, ";" )
  self.scope:add( varName.txt, refType:get_expType(), accessMode, false )
  return Ast.DeclMemberNode.new(firstToken.pos, refType:get_expTypeList(), varName, refType, staticFlag, accessMode, getterMode, setterMode)
end

function TransUnit:analyzeDeclMethod( overrideFlag, accessMode, staticFlag, className, firstToken, name )
  local node = self:analyzeDeclFunc( overrideFlag, accessMode, staticFlag, className, name, name )
  
  return node
end

function TransUnit:analyzeDeclClass( classAccessMode, firstToken, classFlag )
  local name = self:getSymbolToken(  )
  
  local moduleName = nil
  
  if not classFlag then
    self:checkNextToken( "require" )
    moduleName = self:getToken(  )
  end
  local nextToken = self:getToken(  )
  
  local baseRef = nil
  
  if nextToken.txt == "extend" then
    baseRef = self:analyzeRefType( classAccessMode )
    nextToken = self:getToken(  )
  end
  self:checkToken( nextToken, "{" )
  local baseTypeInfo = nil
  
  do
    local _exp = baseRef
    if _exp then
    
        baseTypeInfo = _exp:get_expType(  )
        do
          local _exp = _exp:get_expType(  ):get_scope()
          if _exp then
          
              do
                local _exp = _exp:getTypeInfoChild( "__init" )
                if _exp then
                
                    if _exp:get_accessMode() == "pri" then
                      self:addErrMess( firstToken.pos, "The access mode of '__init' is 'pri'." )
                    end
                  end
              end
              
            end
        end
        
      end
  end
  
  local classTypeInfo = self:pushClass( baseTypeInfo, false, name.txt, classAccessMode )
  
  self.scope:add( "self", classTypeInfo, "pri", false )
  local fieldList = {}
  
  local memberList = {}
  
  local methodName2Node = {}
  
  local initStmtList = {}
  
  local node = Ast.DeclClassNode.new(firstToken.pos, {classTypeInfo}, classAccessMode, name, fieldList, moduleName, memberList, self.scope, initStmtList, {})
  
  self.typeInfo2ClassNode[classTypeInfo] = node
  while true do
    local token = self:getToken(  )
    
    if token.txt == "}" then
      break
    end
    local accessMode = "pri"
    
    if token.txt == "pub" or token.txt == "pro" or token.txt == "pri" or token.txt == "global" then
      accessMode = token.txt
      token = self:getToken(  )
    end
    local staticFlag = false
    
    if token.txt == "static" then
      staticFlag = true
      token = self:getToken(  )
    end
    local overrideFlag = false
    
    if token.txt == "override" then
      overrideFlag = true
      token = self:getToken(  )
    end
    if token.txt == "let" then
      local memberNode = self:analyzeDeclMember( accessMode, staticFlag, token )
      
      table.insert( fieldList, memberNode )
      table.insert( memberList, memberNode )
    elseif token.txt == "fn" then
      local nameToken = self:getToken(  )
      
      local methodNode = self:analyzeDeclMethod( overrideFlag, accessMode, staticFlag, name, token, nameToken )
      
      table.insert( fieldList, methodNode )
    elseif token.txt == "__init" then
      self:checkNextToken( "{" )
      self:analyzeStatementList( initStmtList, "}" )
      self:checkNextToken( "}" )
    elseif token.txt == ";" then
    else 
      self:error( "illegal field" )
    end
  end
  local parentInfo = classTypeInfo
  
  local memberTypeList = {}
  
  for __index, memberNode in pairs( memberList ) do
    local memberType = memberNode:get_expType()
    
    table.insert( memberTypeList, memberType )
    if memberNode:get_expType():get_accessMode() ~= "pub" then
      memberType = Ast.NormalTypeInfo.cloneToPublic( memberType )
    end
    local memberName = memberNode:get_name()
    
    local getterName = "get_" .. memberName.txt
    
    local accessMode = memberNode:get_getterMode()
    
    if accessMode ~= "none" and not self.scope:getTypeInfoChild( getterName ) then
      local retTypeInfo = Ast.NormalTypeInfo.createFunc( false, self:pushScope( false ), Ast.TypeInfoKindMethod, parentInfo, true, false, false, accessMode, getterName, {}, {memberType} )
      
      self:popScope(  )
      self.scope:add( getterName, retTypeInfo, accessMode, false )
    end
    local setterName = "set_" .. memberName.txt
    
    accessMode = memberNode:get_setterMode()
    if memberNode:get_setterMode() ~= "none" and not self.scope:getTypeInfoChild( setterName ) then
      self.scope:add( setterName, Ast.NormalTypeInfo.createFunc( false, self:pushScope( false ), Ast.TypeInfoKindMethod, parentInfo, true, false, false, accessMode, setterName, {memberType}, nil ), accessMode, false )
      self:popScope(  )
    end
  end
  if not self.scope:getTypeInfoChild( "__init" ) then
    local initTypeInfo = Ast.NormalTypeInfo.createFunc( false, self:pushScope( false ), Ast.TypeInfoKindMethod, parentInfo, true, false, false, "pub", "__init", memberTypeList, {} )
    
    self:popScope(  )
    self.scope:add( "__init", initTypeInfo, "pub", false )
  end
  self:popClass(  )
  return node
end

function TransUnit:addMethod( className, methodNode, name )
  local classTypeInfo = self.scope:getTypeInfo( className, self.scope, false )
  
  local classNodeInfo = self.typeInfo2ClassNode[classTypeInfo] or _luneScript.error( 'unwrap val is nil' )
  
  classNodeInfo:get_outerMethodSet()[name] = true
  table.insert( classNodeInfo:get_fieldList(), methodNode )
end

function TransUnit:analyzeDeclFunc( overrideFlag, accessMode, staticFlag, classNameToken, firstToken, name )
  local token = self:getToken(  )
  
  do
    local _exp = name
    if _exp then
    
        name = self:checkSymbol( _exp )
      else
    
        if token.txt ~= "(" then
          name = self:checkSymbol( token )
          token = self:getToken(  )
        end
      end
  end
  
  local needPopFlag = false
  
  if token.txt == "." then
    needPopFlag = true
    classNameToken = name
    self:pushClass( nil, false, (name or _luneScript.error( 'unwrap val is nil' ) ).txt, "pub" )
    name = self:getSymbolToken(  )
    token = self:getToken(  )
  end
  local kind = Ast.nodeKindDeclConstr
  
  local typeKind = Ast.TypeInfoKindFunc
  
  if classNameToken then
    if not staticFlag then
      typeKind = Ast.TypeInfoKindMethod
    end
    if (name or _luneScript.error( 'unwrap val is nil' ) ).txt == "__init" then
      kind = Ast.nodeKindDeclConstr
    else 
      kind = Ast.nodeKindDeclMethod
    end
  else 
    kind = Ast.nodeKindDeclFunc
    if not staticFlag then
      staticFlag = true
    end
  end
  local funcName = ""
  
  do
    local _exp = name
    if _exp then
    
        funcName = _exp.txt
      end
  end
  
  if overrideFlag then
    if not name then
      self:error( "name is nil" )
    end
    -- none
    
    local overrideType = self.scope:getTypeInfoField( funcName, false, self.scope )
    
        if  not overrideType then
          local _overrideType = overrideType
          
          self:error( "not found override -- " .. funcName )
        end
      
    if overrideType:get_accessMode(  ) ~= accessMode then
      self:error( string.format( "missmatch override accessMode -- %s,%s,%s", funcName, overrideType:get_accessMode(  ), accessMode) )
    end
    if overrideType:get_staticFlag(  ) ~= staticFlag then
      self:error( "missmatch override staticFlag -- " .. funcName )
    end
    if overrideType:get_kind(  ) ~= Ast.TypeInfoKindMethod then
      self:error( string.format( "missmatch override kind -- %s, %d", funcName, overrideType:get_kind(  )) )
    end
  else 
    do
      local _exp = name
      if _exp then
      
          if _exp.txt ~= "__init" and self.scope:getTypeInfoField( _exp.txt, false, self.scope ) then
            self:error( "missmatch override --" .. funcName )
          end
        end
    end
    
  end
  self:checkToken( token, "(" )
  local scope = self:pushScope(  )
  
  local argList = {}
  
  token = self:analyzeDeclArgList( accessMode, argList )
  local argTypeList = {}
  
  for __index, argNode in pairs( argList ) do
    table.insert( argTypeList, argNode:get_expType() )
  end
  self:checkToken( token, ")" )
  token = self:getToken(  )
  local retTypeInfoList = {}
  
  if token.txt == ":" then
    repeat 
      local refType = self:analyzeRefType( accessMode )
      
      table.insert( retTypeInfoList, refType:get_expType() )
      token = self:getToken(  )
    until token.txt ~= ","
  end
  local typeInfo = Ast.NormalTypeInfo.createFunc( false, scope, typeKind, self:getCurrentNamespaceTypeInfo(  ), false, false, staticFlag, accessMode, funcName, argTypeList, retTypeInfoList )
  
  do
    local _exp = name
    if _exp then
    
        scope:get_parent(  ):add( _exp.txt, typeInfo, accessMode, false )
      end
  end
  
  local node = self:createNoneNode( firstToken.pos )
  
  if token.txt == ";" then
    node = self:createNoneNode( firstToken.pos )
  else 
    self:pushback(  )
    local body = self:analyzeBlock( "func", scope )
    
    local info = Ast.DeclFuncInfo.new(classNameToken, name, argList, staticFlag, accessMode, body, retTypeInfoList)
    
    do
      local _switchExp = (kind )
      if _switchExp == Ast.nodeKindDeclConstr then
        node = Ast.DeclConstrNode.new(firstToken.pos, {typeInfo}, info)
      elseif _switchExp == Ast.nodeKindDeclMethod then
        node = Ast.DeclMethodNode.new(firstToken.pos, {typeInfo}, info)
      elseif _switchExp == Ast.nodeKindDeclFunc then
        node = Ast.DeclFuncNode.new(firstToken.pos, {typeInfo}, info)
      else 
        self:error( string.format( "illegal kind -- %d", kind) )
      end
    end
    
  end
  self:popScope(  )
  if needPopFlag then
    self:addMethod( (classNameToken or _luneScript.error( 'unwrap val is nil' ) ).txt, node, funcName )
    self:popClass(  )
  end
  return node
end

function TransUnit:analyzeDeclVar( mode, accessMode, staticFlag, firstToken )
  local unwrapFlag = false
  
  local token, continueFlag = self:getContinueToken(  )
  
  if continueFlag and token.txt == "!" then
    unwrapFlag = true
  else 
    self:pushback(  )
    if mode ~= "let" then
      Util.errorLog( "need '!'" )
    end
  end
  local typeInfoList = {}
  
  local varNameList = {}
  
  local varTypeList = {}
  
  repeat 
    local varName = self:getSymbolToken(  )
    
    token = self:getToken(  )
    local typeInfo = Ast.builtinTypeNone
    
    if token.txt == ":" then
      local refType = self:analyzeRefType( accessMode )
      
      table.insert( varTypeList, refType )
      typeInfo = refType:get_expType()
      token = self:getToken(  )
    else 
      table.insert( varTypeList, nil )
    end
    table.insert( varNameList, varName )
    table.insert( typeInfoList, typeInfo )
  until token.txt ~= ","
  local expList = nil
  
  if token.txt == "=" then
    expList = self:analyzeExpList(  )
    if not expList then
      self:error( "expList is nil" )
    end
  end
  local orgExpTypeList = {}
  
  do
    local _exp = expList
    if _exp then
    
        local expTypeList = {}
        
        for index, exp in pairs( _exp:get_expList() ) do
          if index == #_exp:get_expList() then
            if exp:get_expType() == Ast.builtinTypeDDD then
              for subIndex = index, #varNameList do
                local argType = typeInfoList[subIndex]
                
                if argType ~= Ast.builtinTypeNone and not argType:isSettableFrom( Ast.builtinTypeStem_ ) then
                  self:addErrMess( firstToken.pos, string.format( "unmatch value type %s(%d) <- %s(%d)", argType:getTxt(  ), argType:get_typeId(), Ast.builtinTypeStem_:getTxt(  ), Ast.builtinTypeStem_:get_typeId()) )
                end
                if unwrapFlag then
                  table.insert( expTypeList, Ast.builtinTypeStem )
                else 
                  table.insert( expTypeList, Ast.builtinTypeStem_ )
                end
                table.insert( orgExpTypeList, Ast.builtinTypeStem_ )
              end
            else 
              for __index, typeInfo in pairs( exp:get_expTypeList() ) do
                table.insert( orgExpTypeList, typeInfo )
                if unwrapFlag and typeInfo:get_nilable() then
                  typeInfo = typeInfo:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
                end
                table.insert( expTypeList, typeInfo )
                local argType = typeInfoList[index]
                
                if not (unwrapFlag and typeInfo == Ast.builtinTypeNil ) and argType ~= Ast.builtinTypeNone and not argType:isSettableFrom( typeInfo ) then
                  self:addErrMess( firstToken.pos, string.format( "unmatch value type %s <- %s", argType:getTxt(  ), typeInfo:getTxt(  )) )
                end
              end
            end
          else 
            local expTypeInfo = Ast.builtinTypeStem_
            
            if exp:get_expType() == Ast.builtinTypeDDD then
              expTypeInfo = Ast.builtinTypeStem_
            else 
              expTypeInfo = exp:get_expType()
            end
            table.insert( orgExpTypeList, expTypeInfo )
            if unwrapFlag and expTypeInfo:get_nilable() then
              expTypeInfo = expTypeInfo:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
            end
            local argType = typeInfoList[index]
            
            if argType ~= Ast.builtinTypeNone and not argType:isSettableFrom( expTypeInfo ) then
              self:addErrMess( firstToken.pos, string.format( "unmatch value type %s <- %s", argType:getTxt(  ), expTypeInfo:getTxt(  )) )
            end
            table.insert( expTypeList, expTypeInfo )
          end
        end
        for index, typeInfo in pairs( expTypeList ) do
          if not typeInfoList[index] or typeInfoList[index] == Ast.builtinTypeNone then
            typeInfoList[index] = typeInfo
          end
        end
      end
  end
  
  if self.macroScope == self.scope then
    for index, varName in pairs( varNameList ) do
      local typeInfo = typeInfoList[index]
      
      self.symbol2ValueMapForMacro[varName.txt] = Ast.MacroValInfo.new(nil, typeInfo)
    end
  end
  local syncScope = self.scope
  
  if mode == "sync" then
    syncScope = self:pushScope(  )
  end
  local varList = {}
  
  local sameSymbolList = {}
  
  local currentClass = self:getCurrentClass(  )
  
  for index, varName in pairs( varNameList ) do
    local typeInfo = typeInfoList[index]
    
    local varInfo = Ast.VarInfo.new(varName, varTypeList[index], typeInfo)
    
    table.insert( varList, varInfo )
    if not varTypeList[index] and typeInfo == Ast.builtinTypeNil then
      self:addErrMess( varName.pos, string.format( 'need type -- %s', varName.txt) )
    end
    local sameTypeInfo = self.scope:getTypeInfo( varName.txt, self.scope, true )
    
    do
      local _exp = sameTypeInfo
      if _exp then
      
          table.insert( sameSymbolList, varInfo )
        end
    end
    
    if mode == "let" or mode == "sync" then
      if mode == "let" then
        if self.scope:getTypeInfo( varName.txt, self.scope, true ) then
          self:addErrMess( varName.pos, string.format( "shadowing variable -- %s", varName.txt) )
        end
      end
      self.scope:add( varName.txt, typeInfo, "local", false )
    end
  end
  local unwrapBlock = nil
  
  local thenBlock = nil
  
  if unwrapFlag then
    local scope = self:pushScope(  )
    
    for index, varName in pairs( varNameList ) do
      self.scope:add( "_" .. varName.txt, orgExpTypeList[index], "local", false )
    end
    unwrapBlock = self:analyzeBlock( "let!", scope )
    self:popScope(  )
    token = self:getToken(  )
    if token.txt == "then" then
      thenBlock = self:analyzeBlock( "let!", scope )
    else 
      self:pushback(  )
    end
  end
  local syncBlock = nil
  
  if mode == "sync" then
    local nextToken = self:getToken(  )
    
    if nextToken.txt == "do" then
      syncBlock = self:analyzeBlock( "let!", syncScope )
    else 
      self:pushback(  )
    end
    self:popScope(  )
  end
  self:checkNextToken( ";" )
  local node = Ast.DeclVarNode.new(firstToken.pos, {Ast.builtinTypeNone}, mode, accessMode, staticFlag, varList, expList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, sameSymbolList, syncBlock)
  
  return node
end

function TransUnit:analyzeExpList( skipOp2Flag )
  local expList = {}
  
  local pos = nil
  
  local expTypeList = {}
  
  repeat 
    local exp = self:analyzeExp( skipOp2Flag, 0 )
    
    if not pos then
      pos = exp:get_pos()
    end
    table.insert( expList, exp )
    table.insert( expTypeList, exp:get_expType() )
    local token = self:getToken(  )
    
  until token.txt ~= ","
  self:pushback(  )
  return Ast.ExpListNode.new(pos or Parser.Position.new(0, 0), expTypeList, expList)
end

function TransUnit:analyzeListConst( token )
  local nextToken = self:getToken(  )
  
  local expList = nil
  
  local itemTypeInfo = Ast.builtinTypeNone
  
  if nextToken.txt ~= "]" then
    self:pushback(  )
    expList = self:analyzeExpList(  )
    self:checkNextToken( "]" )
    local nodeList = (expList or _luneScript.error( 'unwrap val is nil' ) ):get_expList()
    
    for __index, exp in pairs( nodeList ) do
      local expType = exp:get_expType()
      
      if itemTypeInfo == Ast.builtinTypeNone then
        itemTypeInfo = expType
      elseif not itemTypeInfo:isSettableFrom( expType ) then
        if expType == Ast.builtinTypeNil then
          itemTypeInfo = itemTypeInfo:get_nilableTypeInfo() or _luneScript.error( 'unwrap val is nil' )
        elseif expType:get_nilable() then
          itemTypeInfo = Ast.builtinTypeStem_
        else 
          itemTypeInfo = Ast.builtinTypeStem
        end
      end
    end
  end
  local kind = Ast.nodeKindLiteralArray
  
  local typeInfoList = {Ast.builtinTypeNone}
  
  if token.txt == '[' then
    kind = Ast.nodeKindLiteralList
    typeInfoList = {Ast.NormalTypeInfo.createList( "local", self:getCurrentClass(  ), {itemTypeInfo} )}
    return Ast.LiteralListNode.new(token.pos, typeInfoList, expList)
  else 
    typeInfoList = {Ast.NormalTypeInfo.createArray( "local", self:getCurrentClass(  ), {itemTypeInfo} )}
    return Ast.LiteralArrayNode.new(token.pos, typeInfoList, expList)
  end
end

function TransUnit:analyzeMapConst( token )
  local nextToken = self:getToken(  )
  
  local map = {}
  
  local pairList = {}
  
  local keyTypeInfo = Ast.builtinTypeNone
  
  local valTypeInfo = Ast.builtinTypeNone
  
  local function getMapKeyValType( pos, keyFlag, typeInfo, expType )
    if expType:get_nilable() then
      if keyFlag then
        self:addErrMess( pos, string.format( "map key can't set a nilable -- %s", expType:getTxt(  )) )
      end
      if expType == Ast.builtinTypeNil then
        return typeInfo
      end
      expType = expType:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
    end
    if not typeInfo:isSettableFrom( expType ) then
      if typeInfo ~= Ast.builtinTypeNone then
        typeInfo = Ast.builtinTypeStem
      else 
        typeInfo = expType
      end
    end
    return typeInfo
  end
  
  while true do
    if nextToken.txt == "}" then
      break
    end
    self:pushback(  )
    local key = self:analyzeExp(  )
    
    keyTypeInfo = getMapKeyValType( key:get_pos(), true, keyTypeInfo, key:get_expType() )
    self:checkNextToken( ":" )
    local val = self:analyzeExp(  )
    
    valTypeInfo = getMapKeyValType( val:get_pos(), false, valTypeInfo, val:get_expType() )
    table.insert( pairList, Ast.PairItem.new(key, val) )
    map[key] = val
    nextToken = self:getToken(  )
    if nextToken.txt ~= "," then
      break
    end
    nextToken = self:getToken(  )
  end
  local typeInfo = Ast.NormalTypeInfo.createMap( "local", self:getCurrentClass(  ), keyTypeInfo, valTypeInfo )
  
  self:checkToken( nextToken, "}" )
  return Ast.LiteralMapNode.new(token.pos, {typeInfo}, map, pairList)
end

function TransUnit:analyzeExpRefItem( token, exp, nilAccess )
  local indexExp = self:analyzeExp(  )
  
  self:checkNextToken( "]" )
  local expType = exp:get_expType()
  
  if nilAccess then
    if not expType:get_nilable() then
      nilAccess = false
    else 
      expType = expType:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
    end
  end
  local typeInfo = Ast.builtinTypeStem_
  
  if expType:get_kind() == Ast.TypeInfoKindMap then
    typeInfo = expType:get_itemTypeInfoList(  )[2]
    if typeInfo ~= Ast.builtinTypeStem_ and not typeInfo:get_nilable() then
      typeInfo = typeInfo:get_nilableTypeInfo()
    end
  elseif expType:get_kind() == Ast.TypeInfoKindArray or expType:get_kind() == Ast.TypeInfoKindList then
    typeInfo = expType:get_itemTypeInfoList(  )[1]
  elseif expType == Ast.builtinTypeString then
    typeInfo = Ast.builtinTypeInt
  elseif expType == Ast.builtinTypeStem then
    typeInfo = Ast.builtinTypeStem
  else 
    self:addErrMess( exp:get_pos(), string.format( "could not access with []. -- %s", expType:getTxt(  )) )
  end
  if not typeInfo then
    self:error( "illegal type" )
  end
  return Ast.ExpRefItemNode.new(token.pos, {typeInfo}, exp, nilAccess, indexExp)
end

function TransUnit:checkMatchValType( pos, funcTypeInfo, expList, genericTypeList )
  local argTypeList = funcTypeInfo:get_argTypeInfoList()
  
  do
    local _switchExp = funcTypeInfo
    if _switchExp == typeInfoListInsert then
      argTypeList = genericTypeList
    elseif _switchExp == typeInfoListRemove then
    end
  end
  
  do
    local _exp = expList
    if _exp then
    
        local expNodeList = _exp:get_expList()
        
        for index, expNode in pairs( expNodeList ) do
          if #argTypeList < index then
            self:addErrMess( pos, string.format( "%s: over argument. expect:%d, actual:%d", funcTypeInfo:getTxt(  ), #argTypeList, #expNodeList) )
            break
          end
          local argType = argTypeList[index]
          
          local expType = expNode:get_expType()
          
          if #argTypeList == index then
            if argType ~= Ast.builtinTypeDDD then
              if not argType:isSettableFrom( expType ) then
                self:addErrMess( expNode:get_pos(), string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), index, argType:getTxt(  ), expType:getTxt(  )) )
              end
            end
            break
          elseif #expNodeList == index then
            if expType == Ast.builtinTypeDDD then
              for argIndex = index, #argTypeList do
                local workArgType = argTypeList[argIndex]
                
                if not workArgType:isSettableFrom( Ast.builtinTypeStem_ ) then
                  self:addErrMess( expNode:get_pos(), string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), argIndex, workArgType:getTxt(  ), expType:getTxt(  )) )
                end
              end
            else 
              for argIndex = index, #argTypeList do
                local argTypeInfo = argTypeList[argIndex]
                
                if not argTypeInfo:isSettableFrom( expType ) then
                  self:addErrMess( expNode:get_pos(), string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), argIndex, argTypeInfo:getTxt(  ), expType:getTxt(  )) )
                end
                expType = Ast.builtinTypeNil
              end
            end
            break
          end
          if not argType:isSettableFrom( expType ) then
            self:addErrMess( expNode:get_pos(), string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), index, argType:getTxt(  ), expType:getTxt(  )) )
          end
        end
      else
    
        if #argTypeList ~= 0 then
        end
      end
  end
  
end

local MacroPaser = {}
setmetatable( MacroPaser, { __index = Parser } )
function MacroPaser.new( tokenList, name )
  local obj = {}
  setmetatable( obj, { __index = MacroPaser } )
  if obj.__init then obj:__init( tokenList, name ); end
return obj
end
function MacroPaser:__init(tokenList, name) 
  self.pos = 1
  self.tokenList = tokenList
  self.name = name
end
function MacroPaser:getToken(  )
  if #self.tokenList < self.pos then
    return nil
  end
  local token = self.tokenList[self.pos]
  
  self.pos = self.pos + 1
  return token
end
function MacroPaser:getStreamName(  )
  return self.name
end
do
  end

function TransUnit:evalMacro( firstToken, macroTypeInfo, expList )
  do
    local _exp = expList
    if _exp then
    
        if _exp:get_expList(  ) then
          for __index, exp in pairs( _exp:get_expList(  ) ) do
            local kind = exp:get_kind()
            
            if kind ~= Ast.nodeKindLiteralNil and kind ~= Ast.nodeKindLiteralChar and kind ~= Ast.nodeKindLiteralInt and kind ~= Ast.nodeKindLiteralReal and kind ~= Ast.nodeKindLiteralArray and kind ~= Ast.nodeKindLiteralList and kind ~= Ast.nodeKindLiteralMap and kind ~= Ast.nodeKindLiteralString and kind ~= Ast.nodeKindLiteralBool and kind ~= Ast.nodeKindLiteralSymbol and kind ~= Ast.nodeKindRefField and kind ~= Ast.nodeKindExpMacroStat then
              self:error( "Macro arguments must be literal value." )
            end
          end
        end
      end
  end
  
  local macroInfo = self.typeId2MacroInfo[macroTypeInfo:get_typeId(  )] or _luneScript.error( 'unwrap val is nil' )
  
  local argVal = {}
  
  do
    local _exp = expList
    if _exp then
    
        for __index, argNode in pairs( _exp:get_expList(  ) ) do
          local valList, typeInfoList = argNode:getLiteral(  )
          
          local val = valList[1]
          
          local typeInfo = typeInfoList[1]
          
          table.insert( argVal, val )
        end
      end
  end
  
  local func = macroInfo.func
  
  local macroVars = func( table.unpack( argVal ) )
  
  for __index, name in pairs( (macroVars._names or _luneScript.error( 'unwrap val is nil' ) ) ) do
    local valInfo = macroInfo.symbol2MacroValInfoMap[name] or _luneScript.error( 'unwrap val is nil' )
    
    local typeInfo = valInfo and valInfo.typeInfo or Ast.builtinTypeStem_
    
    local val = macroVars[name]
    
    if typeInfo == Ast.builtinTypeSymbol then
      val = {val}
    end
    self.symbol2ValueMapForMacro[name] = Ast.MacroValInfo.new(val, typeInfo)
  end
  local argList = macroInfo.declInfo:get_argList(  )
  
  if argList then
    for index, arg in pairs( argList ) do
      if arg:get_kind(  ) == Ast.nodeKindDeclArg then
        local argInfo = arg
        
        local argType = argInfo:get_argType()
        
        local argName = argInfo:get_name().txt
        
        self.symbol2ValueMapForMacro[argName] = Ast.MacroValInfo.new(argVal[index], argType:get_expType())
      else 
        self:error( "not support ... in macro" )
      end
    end
  end
  local parser = MacroPaser.new(macroInfo.declInfo:get_tokenList(), string.format( "macro %s", macroTypeInfo:getTxt(  )))
  
  local bakParser = self.parser
  
  self.parser = parser
  self.macroMode = "expand"
  local stmtList = {}
  
  self:analyzeStatementList( stmtList, "}" )
  self.macroMode = "none"
  self.parser = bakParser
  return Ast.ExpMacroExpNode.new(firstToken.pos, {Ast.builtinTypeNone}, stmtList)
end

function TransUnit:analyzeExpCont( firstToken, exp, skipFlag )
  local nextToken = self:getToken(  )
  
  if not skipFlag then
    repeat 
      local matchFlag = false
      
      if nextToken.txt == "[" or nextToken.txt == "$[" then
        matchFlag = true
        exp = self:analyzeExpRefItem( nextToken, exp, nextToken.txt == "$[" )
        nextToken = self:getToken(  )
      end
      if nextToken.txt == "(" then
        local macroFlag = false
        
        local funcTypeInfo = exp:get_expType()
        
        if funcTypeInfo:get_kind(  ) == Ast.TypeInfoKindMacro then
          macroFlag = true
          self.symbol2ValueMapForMacro = {}
          self.macroMode = "analyze"
        end
        matchFlag = true
        local work = self:getToken(  )
        
        local expList = nil
        
        if work.txt ~= ")" then
          self:pushback(  )
          expList = self:analyzeExpList(  )
          self:checkNextToken( ")" )
        end
        local genericTypeList = funcTypeInfo:get_itemTypeInfoList()
        
        if funcTypeInfo:get_kind() == Ast.TypeInfoKindMethod and exp:get_kind() == Ast.nodeKindRefField then
          local refField = exp
          
          local classType = refField:get_prefix():get_expType()
          
          genericTypeList = classType:get_itemTypeInfoList()
        end
        self:checkMatchValType( exp:get_pos(), funcTypeInfo, expList, genericTypeList )
        if macroFlag then
          self.macroMode = "none"
          exp = self:evalMacro( firstToken, funcTypeInfo, expList )
        else 
          do
            local _switchExp = (exp:get_expType():get_kind() )
            if _switchExp == Ast.TypeInfoKindMethod or _switchExp == Ast.TypeInfoKindFunc then
            else 
              self:error( string.format( "can't call the type -- %s", exp:get_expType():getTxt(  )) )
            end
          end
          
          exp = Ast.ExpCallNode.new(firstToken.pos, funcTypeInfo:get_retTypeInfoList(  ), exp, expList)
        end
        nextToken = self:getToken(  )
      end
    until not matchFlag
  end
  do
    local _switchExp = nextToken.txt
    if _switchExp == "." then
      return self:analyzeExpSymbol( firstToken, self:getToken(  ), "field", exp, skipFlag )
    elseif _switchExp == "$." then
      return self:analyzeExpSymbol( firstToken, self:getToken(  ), "field_nil", exp, skipFlag )
    elseif _switchExp == ".$" then
      return self:analyzeExpSymbol( firstToken, self:getToken(  ), "get", exp, skipFlag )
    elseif _switchExp == "$.$" then
      return self:analyzeExpSymbol( firstToken, self:getToken(  ), "get_nil", exp, skipFlag )
    end
  end
  
  self:pushback(  )
  return exp
end

function TransUnit:analyzeExpSymbol( firstToken, token, mode, prefixExp, skipFlag )
  local exp = nil
  
  if mode == "field" or mode == "get" or mode == "field_nil" or mode == "get_nil" then
    local prefixExp = prefixExp
    
        if  not prefixExp then
          local _prefixExp = prefixExp
          
          self:error( "prefix is nil" )
        else
          
            local accessNil = false
            
            if mode == "field_nil" or mode == "get_nil" then
              accessNil = true
            end
            if self.macroMode == "analyze" then
              exp = Ast.RefFieldNode.new(firstToken.pos, {Ast.builtinTypeSymbol}, token, accessNil, prefixExp or _luneScript.error( 'unwrap val is nil' ))
            else 
              local typeInfo = Ast.builtinTypeStem_
              
              local prefixExpType = prefixExp:get_expType()
              
              if accessNil then
                if prefixExpType:get_nilable() then
                  prefixExpType = prefixExpType:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
                else 
                  accessNil = false
                end
              end
              local getterTypeInfo = nil
              
              if prefixExpType:get_kind(  ) == Ast.TypeInfoKindClass or prefixExpType:get_kind(  ) == Ast.TypeInfoKindList then
                if prefixExpType:get_kind(  ) == Ast.TypeInfoKindList then
                  prefixExpType = Ast.builtinTypeList
                end
                local className = prefixExpType:getTxt(  )
                
                local classScope = prefixExpType:get_scope()
                
                    if  not classScope then
                      local _classScope = classScope
                      
                      self:error( string.format( "not found field: %s, %s, %s", classScope, className, prefixExpType) )
                    end
                  
                typeInfo = nil
                if mode == "get" or mode == "get_nil" then
                  typeInfo = classScope:getTypeInfo( string.format( "get_%s", token.txt), self.scope, false )
                  do
                    local _exp = typeInfo
                    if _exp then
                    
                        if (_exp:get_kind(  ) == Ast.TypeInfoKindMethod ) then
                          local retTypeList = _exp:get_retTypeInfoList(  )
                          
                          getterTypeInfo = _exp
                          typeInfo = retTypeList[1]
                        end
                      end
                  end
                  
                end
                if not getterTypeInfo then
                  typeInfo = classScope:getTypeInfoField( token.txt, true, self.scope )
                end
                if not typeInfo then
                  for name, val in pairs( classScope:get_symbol2TypeInfoMap() ) do
                    Util.errorLog( string.format( "debug: %s, %s", name, val) )
                  end
                  self:error( string.format( "not found field typeInfo: %s.%s", className, token.txt ) )
                end
              elseif prefixExpType:get_kind(  ) == Ast.TypeInfoKindMap then
                local work = prefixExpType:get_itemTypeInfoList()[1]
                
                if work ~= Ast.builtinTypeString then
                  self:addErrMess( token.pos, string.format( "map key type is not str. (%s)", work:getTxt(  )) )
                end
                typeInfo = prefixExpType:get_itemTypeInfoList()[2]
                do
                  local _exp = typeInfo
                  if _exp then
                  
                      if not _exp:get_nilable() then
                        typeInfo = _exp:get_nilableTypeInfo()
                      end
                    end
                end
                
              elseif prefixExpType == Ast.builtinTypeStem then
              else 
                self:error( string.format( "illegal type -- %s, %d", prefixExpType:getTxt(  ), prefixExpType:get_kind(  )) )
              end
              do
                local _exp = getterTypeInfo
                if _exp then
                
                    exp = Ast.GetFieldNode.new(firstToken.pos, {typeInfo or _luneScript.error( 'unwrap val is nil' )}, token, accessNil, prefixExp, _exp)
                  else
                
                    exp = Ast.RefFieldNode.new(firstToken.pos, {typeInfo or _luneScript.error( 'unwrap val is nil' )}, token, accessNil, prefixExp)
                  end
              end
              
            end
          end
      
  elseif mode == "symbol" then
    if self.macroMode == "analyze" then
      exp = Ast.LiteralSymbolNode.new(firstToken.pos, {Ast.builtinTypeSymbol}, token)
    else 
      local symbolInfo = self.scope:getSymbolTypeInfo( token.txt, self.scope, self.moduleScope )
      
          if  not symbolInfo then
            local _symbolInfo = symbolInfo
            
            self:error( "not found type -- " .. token.txt )
          end
        
      local typeInfo = symbolInfo:get_typeInfo()
      
      if typeInfo == Ast.builtinTypeSymbol then
        skipFlag = true
      end
      exp = Ast.ExpRefNode.new(firstToken.pos, {typeInfo}, token, symbolInfo)
    end
  elseif mode == "fn" then
    exp = self:analyzeDeclFunc( false, "local", false, nil, token, nil )
  else 
    self:error( "illegal mode", mode )
  end
  return self:analyzeExpCont( firstToken, exp or _luneScript.error( 'unwrap val is nil' ), skipFlag )
end

function TransUnit:analyzeExpOp2( firstToken, exp, prevOpLevel )
  while true do
    local nextToken = self:getToken(  )
    
    local opTxt = nextToken.txt
    
    if opTxt == "@" then
      local castType = self:analyzeRefType( "local" )
      
      if exp:get_expType():get_nilable() and not castType:get_expType():get_nilable() then
        self:addErrMess( firstToken.pos, string.format( "can't cast from nilable to not nilable  -- %s->%s", exp:get_expType():getTxt(  ), castType:get_expType():getTxt(  )) )
      end
      exp = Ast.ExpCastNode.new(firstToken.pos, castType:get_expTypeList(), exp)
    elseif nextToken.kind == Parser.kind.Ope then
      if Parser.isOp2( opTxt ) then
        local opLevel = op2levelMap[opTxt]
        
            if  not opLevel then
              local _opLevel = opLevel
              
              error( string.format( "unknown op -- %s %s", opTxt, prevOpLevel ) )
            end
          
        if prevOpLevel and opLevel <= prevOpLevel then
          self:pushback(  )
          return exp
        end
        local exp2 = self:analyzeExp( false, opLevel )
        
        local info = {["op"] = nextToken, ["exp1"] = exp, ["exp2"] = exp2}
        
        if not exp:get_expType() or not exp2:get_expType() then
          self:error( string.format( "illegal exp or exp2 %s, %s, %s , %s,%d:%d", exp:get_expType(), exp2:get_expType(), nextToken.txt, self.parser:getStreamName(  ), nextToken.pos.lineNo, nextToken.pos.column) )
        end
        local retType = Ast.builtinTypeNone
        
        local exp1Type = exp:get_expType()
        
        local exp2Type = exp2:get_expType()
        
        if not exp1Type then
          self:error( string.format( "expType is nil %s:%d:%d", self.parser:getStreamName(  ), firstToken.pos.lineNo, firstToken.pos.column) )
        end
        do
          local _switchExp = opTxt
          if _switchExp == "or" then
            if exp1Type:equals( exp2Type ) then
              retType = exp1Type
            elseif exp1Type:get_kind() == exp2Type:get_kind() then
              retType = exp1Type
            elseif exp2Type == Ast.builtinTypeNil then
              retType = exp1Type
            elseif exp1Type == Ast.builtinTypeNil then
              retType = exp2Type
            else 
              retType = Ast.builtinTypeStem_
            end
          elseif _switchExp == "and" then
            retType = exp2Type
          elseif _switchExp == "<" or _switchExp == ">" or _switchExp == "<=" or _switchExp == ">=" then
            if (exp1Type ~= Ast.builtinTypeInt and exp1Type ~= Ast.builtinTypeReal ) or (exp2Type ~= Ast.builtinTypeInt and exp2Type ~= Ast.builtinTypeReal ) then
              self:addErrMess( nextToken.pos, string.format( "no int type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            retType = Ast.builtinTypeBool
          elseif _switchExp == "~=" or _switchExp == "==" then
            if (not exp1Type:isSettableFrom( exp2Type ) and not exp2Type:isSettableFrom( exp1Type ) ) then
              self:addErrMess( nextToken.pos, string.format( "not compatible type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            retType = Ast.builtinTypeBool
          elseif _switchExp == "^" or _switchExp == "|" or _switchExp == "~" or _switchExp == "&" or _switchExp == "<<" or _switchExp == ">>" then
            if exp1Type ~= Ast.builtinTypeInt or exp2Type ~= Ast.builtinTypeInt then
              self:addErrMess( nextToken.pos, string.format( "no int type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            retType = Ast.builtinTypeInt
          elseif _switchExp == ".." then
            if exp1Type ~= Ast.builtinTypeString or exp1Type ~= Ast.builtinTypeString then
              self:addErrMess( nextToken.pos, string.format( "no string type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            retType = Ast.builtinTypeString
          elseif _switchExp == "+" or _switchExp == "-" or _switchExp == "*" or _switchExp == "/" or _switchExp == "//" or _switchExp == "%" then
            if (exp1Type ~= Ast.builtinTypeReal and exp1Type ~= Ast.builtinTypeInt ) or (exp2Type ~= Ast.builtinTypeReal and exp2Type ~= Ast.builtinTypeInt ) then
              self:addErrMess( nextToken.pos, string.format( "no numeric type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            if exp1Type == Ast.builtinTypeReal or exp2Type == Ast.builtinTypeReal then
              retType = Ast.builtinTypeReal
            else 
              retType = Ast.builtinTypeInt
            end
          elseif _switchExp == "=" then
            if not exp1Type:isSettableFrom( exp2Type ) then
              self:addErrMess( nextToken.pos, string.format( "unmatch type %s and %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
          else 
            self:error( "unknown op " .. opTxt )
          end
        end
        
        exp = Ast.ExpOp2Node.new(firstToken.pos, {retType}, nextToken, exp, exp2)
      else 
        self:error( "illegal op" )
      end
    else 
      self:pushback(  )
      return exp
    end
  end
  return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
end

function TransUnit:analyzeExpMacroStat( firstToken )
  local expStrList = {}
  
  self:checkNextToken( "{" )
  local braceCount = 0
  
  local prevToken = firstToken
  
  while true do
    local token = self:getToken(  )
    
    if token.txt == ",," or token.txt == ",,," or token.txt == ",,,," then
      local exp = self:analyzeExp( true, op1levelMap[token.txt] or _luneScript.error( 'unwrap val is nil' ) )
      
      local nextToken = self:getToken(  )
      
      if nextToken.txt ~= "$" then
        self:pushback(  )
      end
      local format = token.txt == ",,," and "'%s '" or '"\'%s\'"'
      
      if token.txt == ",," and exp:get_kind() == Ast.nodeKindExpRef then
        local refToken = (exp ):get_token(  )
        
        local macroInfo = self.symbol2ValueMapForMacro[refToken.txt]
        
        if macroInfo then
          if (macroInfo or _luneScript.error( 'unwrap val is nil' ) ).typeInfo == Ast.builtinTypeSymbol then
            format = "'%s '"
          end
        else 
          if exp:get_expType() == Ast.builtinTypeInt or exp:get_expType() == Ast.builtinTypeReal then
            format = "'%s' "
          end
        end
      end
      local newToken = Parser.Token.new(Parser.kind.Str, format, token.pos)
      
      local literalStr = Ast.LiteralStringNode.new(token.pos, {Ast.builtinTypeString}, newToken, {exp})
      
      table.insert( expStrList, literalStr )
    else 
      if token.txt == "{" then
        braceCount = braceCount + 1
      elseif token.txt == "}" then
        if braceCount == 0 then
          break
        end
        braceCount = braceCount - 1
      end
      local format = "' %s '"
      
      if prevToken == firstToken or (prevToken.pos.lineNo == token.pos.lineNo and prevToken.pos.column + #prevToken.txt == token.pos.column ) then
        format = "'%s'"
      end
      local newToken = Parser.Token.new(token.kind, string.format( format, token.txt ), token.pos)
      
      local literalStr = Ast.LiteralStringNode.new(token.pos, {Ast.builtinTypeString}, newToken, {})
      
      table.insert( expStrList, literalStr )
    end
    prevToken = token
  end
  return Ast.ExpMacroStatNode.new(firstToken.pos, {Ast.builtinTypeStat}, expStrList)
end

function TransUnit:analyzeSuper( firstToken )
  self:checkNextToken( "(" )
  local expList = self:analyzeExpList(  )
  
  self:checkNextToken( ")" )
  self:checkNextToken( ";" )
  local classType = self:getCurrentClass(  )
  
  local superType = classType:get_baseTypeInfo(  )
  
  return Ast.ExpCallSuperNode.new(firstToken.pos, {Ast.builtinTypeNone}, superType, expList)
end

function TransUnit:analyzeUnwrap( firstToken )
  local nextToken, continueFlag = self:getContinueToken(  )
  
  if not continueFlag or nextToken.txt ~= "!" then
    self:pushback(  )
    self:pushbackToken( firstToken )
    local exp = self:analyzeExp(  )
    
    self:checkNextToken( ";" )
    return Ast.StmtExpNode.new(nextToken.pos, {Ast.builtinTypeNone}, exp)
  end
  self:pushback(  )
  return self:analyzeDeclVar( "unwrap", "local", false, firstToken )
end

function TransUnit:analyzeExpUnwrap( firstToken )
  local expNode = self:analyzeExp(  )
  
  local nextToken = self:getToken(  )
  
  local insNode = nil
  
  if nextToken.txt == "default" then
    insNode = self:analyzeExp(  )
  else 
    self:pushback(  )
  end
  local unwrapType = Ast.builtinTypeStem_
  
  local expType = expNode:get_expType()
  
  if not expType:get_nilable() then
    unwrapType = expType
  else 
    unwrapType = expType:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
    do
      local _exp = insNode
      if _exp then
      
          local insType = _exp:get_expType()
          
          unwrapType = insType
          if not unwrapType:isSettableFrom( insType ) then
            self:addErrMess( _exp:get_pos(), string.format( "unmatch type: %s <- %s", unwrapType:getTxt(  ), insType:getTxt(  )) )
          end
        end
    end
    
  end
  return Ast.ExpUnwrapNode.new(firstToken.pos, {unwrapType}, expNode, insNode)
end

function TransUnit:analyzeExp( skipOp2Flag, prevOpLevel )
  local firstToken = self:getToken(  )
  
  local token = firstToken
  
  local exp = Ast.NoneNode.new(firstToken.pos, {Ast.builtinTypeNone})
  
  if token.kind == Parser.kind.Dlmt then
    if token.txt == "..." then
      return Ast.ExpDDDNode.new(firstToken.pos, {Ast.builtinTypeNone}, token)
    end
    if token.txt == '[' or token.txt == '[@' then
      exp = self:analyzeListConst( token )
    end
    if token.txt == '{' then
      exp = self:analyzeMapConst( token )
    end
    if token.txt == "(" then
      exp = self:analyzeExp(  )
      self:checkNextToken( ")" )
      exp = Ast.ExpParenNode.new(firstToken.pos, exp:get_expTypeList(), exp)
      exp = self:analyzeExpCont( firstToken, exp, false )
    end
  end
  if token.txt == "new" then
    exp = self:analyzeRefType( "local" )
    self:checkNextToken( "(" )
    local nextToken = self:getToken(  )
    
    local argList = nil
    
    if nextToken.txt ~= ")" then
      self:pushback(  )
      argList = self:analyzeExpList(  )
      self:checkNextToken( ")" )
    end
    local classTypeInfo = exp:get_expType()
    
    local classScope = classTypeInfo:get_scope(  )
    
    local initTypeInfo = (classScope or _luneScript.error( 'unwrap val is nil' ) ):getTypeInfoChild( "__init" )
    
        if  not initTypeInfo then
          local _initTypeInfo = initTypeInfo
          
          self:error( "not found __init" )
        end
      
    if initTypeInfo:get_accessMode() == "pub" or (initTypeInfo:get_accessMode() == "pro" and self.scope:getClassTypeInfo(  ):isInheritFrom( classTypeInfo ) ) or (self.scope:getClassTypeInfo(  ) == classTypeInfo ) then
    else 
      self:addErrMess( token.pos, string.format( "can't access to __init of %s", classTypeInfo:getTxt(  )) )
    end
    self:checkMatchValType( exp:get_pos(), initTypeInfo, argList, exp:get_expType():get_itemTypeInfoList() )
    exp = Ast.ExpNewNode.new(firstToken.pos, exp:get_expTypeList(), exp, argList)
    exp = self:analyzeExpCont( firstToken, exp, false )
  end
  if token.kind == Parser.kind.Ope and Parser.isOp1( token.txt ) then
    if token.txt == "`" then
      exp = self:analyzeExpMacroStat( token )
    else 
      exp = self:analyzeExp( true, op1levelMap[token.txt] or _luneScript.error( 'unwrap val is nil' ) )
      local typeInfo = Ast.builtinTypeNone
      
      local macroExpFlag = false
      
      do
        local _switchExp = (token.txt )
        if _switchExp == "-" then
          if exp:get_expType() ~= Ast.builtinTypeInt and exp:get_expType() ~= Ast.builtinTypeReal then
            self:addErrMess( token.pos, string.format( 'unmatch type for "-" -- %s', exp:get_expType():getTxt(  )) )
          end
          typeInfo = exp:get_expType()
        elseif _switchExp == "#" then
          if exp:get_expType():get_kind() ~= Ast.TypeInfoKindList and exp:get_expType():get_kind() ~= Ast.TypeInfoKindArray and exp:get_expType():get_kind() ~= Ast.TypeInfoKindMap and exp:get_expType() ~= Ast.builtinTypeString then
            self:addErrMess( token.pos, string.format( 'unmatch type for "#" -- %s', exp:get_expType():getTxt(  )) )
          end
          typeInfo = Ast.builtinTypeInt
        elseif _switchExp == "not" then
          typeInfo = Ast.builtinTypeBool
        elseif _switchExp == ",," then
          macroExpFlag = true
        elseif _switchExp == ",,," then
          macroExpFlag = true
          if exp:get_expType() ~= Ast.builtinTypeString then
            self:error( "unmatch ,,, type, need string type" )
          end
          typeInfo = Ast.builtinTypeSymbol
        elseif _switchExp == ",,,," then
          macroExpFlag = true
          if exp:get_expType() ~= Ast.builtinTypeSymbol then
            self:error( "unmatch ,,, type, need symbol type" )
          end
          typeInfo = Ast.builtinTypeString
        elseif _switchExp == "`" then
          typeInfo = Ast.builtinTypeNone
        elseif _switchExp == "not" then
          typeInfo = Ast.builtinTypeBool
        else 
          self:error( "unknown op1" )
        end
      end
      
      if macroExpFlag then
        local nextToken = self:getToken(  )
        
        if nextToken.txt ~= "$" then
          self:pushback(  )
        end
      end
      exp = Ast.ExpOp1Node.new(firstToken.pos, {typeInfo}, token, self.macroMode, exp)
      return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
    end
  end
  if token.kind == Parser.kind.Int then
    exp = Ast.LiteralIntNode.new(firstToken.pos, {Ast.builtinTypeInt}, token, math.floor(tonumber( token.txt )))
  elseif token.kind == Parser.kind.Real then
    exp = Ast.LiteralRealNode.new(firstToken.pos, {Ast.builtinTypeReal}, token, tonumber( token.txt ))
  elseif token.kind == Parser.kind.Char then
    local num = 0
    
    if #(token.txt ) == 1 then
      num = token.txt:byte( 1 )
    else 
      num = quotedChar2Code[token.txt:sub( 2, 2 )] or _luneScript.error( 'unwrap val is nil' )
    end
    exp = Ast.LiteralCharNode.new(firstToken.pos, {Ast.builtinTypeChar}, token, num)
  elseif token.kind == Parser.kind.Str then
    local nextToken = self:getToken(  )
    
    local formatArgList = {}
    
    if nextToken.txt == "(" then
      repeat 
        local arg = self:analyzeExp(  )
        
        table.insert( formatArgList, arg )
        nextToken = self:getToken(  )
      until nextToken.txt ~= ","
      self:checkToken( nextToken, ")" )
      nextToken = self:getToken(  )
    end
    exp = Ast.LiteralStringNode.new(firstToken.pos, {Ast.builtinTypeString}, token, formatArgList)
    token = nextToken
    if token.txt == "[" or token.txt == "$[" then
      exp = self:analyzeExpRefItem( token, exp, token.txt == "$[" )
    else 
      self:pushback(  )
    end
  elseif token.kind == Parser.kind.Kywd and token.txt == "fn" then
    exp = self:analyzeExpSymbol( firstToken, token, "fn", nil, false )
  elseif token.kind == Parser.kind.Kywd and token.txt == "unwrap" then
    exp = self:analyzeExpUnwrap( token )
  elseif token.kind == Parser.kind.Symb then
    exp = self:analyzeExpSymbol( firstToken, token, "symbol", nil, false )
  elseif token.kind == Parser.kind.Type then
    local symbolTypeInfo = Ast.sym2builtInTypeMap[token.txt]
    
        if  not symbolTypeInfo then
          local _symbolTypeInfo = symbolTypeInfo
          
          self:error( string.format( "unknown type -- %s", token.txt) )
        end
      
    exp = Ast.ExpRefNode.new(firstToken.pos, {Ast.builtinTypeNone}, token, symbolTypeInfo)
  elseif token.kind == Parser.kind.Kywd and (token.txt == "true" or token.txt == "false" ) then
    exp = Ast.LiteralBoolNode.new(firstToken.pos, {Ast.builtinTypeBool}, token)
  elseif token.kind == Parser.kind.Kywd and (token.txt == "nil" or token.txt == "null" ) then
    exp = Ast.LiteralNilNode.new(firstToken.pos, {Ast.builtinTypeNil})
  end
  if not exp then
    self:error( "illegal exp" )
  end
  if skipOp2Flag then
    return exp
  end
  return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
end

function TransUnit:analyzeStatement( termTxt )
  local token = self:getTokenNoErr(  )
  
      if  not token then
        local _token = token
        
        return nil
      end
    
  local statement = self:analyzeDecl( "local", false, token, token )
  
  if not statement then
    if token.txt == termTxt then
      self:pushback(  )
      return nil
    elseif token.txt == "pub" or token.txt == "pro" or token.txt == "pri" or token.txt == "global" or token.txt == "static" then
      local accessMode = (token.txt ~= "static" ) and token.txt or "pri"
      
      local staticFlag = (token.txt == "static" )
      
      local nextToken = token
      
      if token.txt ~= "static" then
        nextToken = self:getToken(  )
      end
      statement = self:analyzeDecl( accessMode, staticFlag, token, nextToken )
    elseif token.txt == "{" then
      self:pushback(  )
      statement = self:analyzeBlock( "{" )
    elseif token.txt == "super" then
      statement = self:analyzeSuper( token )
    elseif token.txt == "if" then
      statement = self:analyzeIf( token )
    elseif token.txt == "switch" then
      statement = self:analyzeSwitch( token )
    elseif token.txt == "while" then
      statement = self:analyzeWhile( token )
    elseif token.txt == "repeat" then
      statement = self:analyzeRepeat( token )
    elseif token.txt == "for" then
      statement = self:analyzeFor( token )
    elseif token.txt == "apply" then
      statement = self:analyzeApply( token )
    elseif token.txt == "foreach" then
      statement = self:analyzeForeach( token, false )
    elseif token.txt == "forsort" then
      statement = self:analyzeForeach( token, true )
    elseif token.txt == "return" then
      local nextToken = self:getToken(  )
      
      local expList = nil
      
      if nextToken.txt ~= ";" then
        self:pushback(  )
        expList = self:analyzeExpList(  )
        self:checkNextToken( ";" )
      end
      local funcTypeInfo = self:getCurrentNamespaceTypeInfo(  )
      
      if not funcTypeInfo then
        self:addErrMess( token.pos, "'return' could not exist here" )
      else 
        do
          local _exp = expList
          if _exp then
          
              local retTypeList = funcTypeInfo:get_retTypeInfoList()
              
              local expNodeList = _exp:get_expList()
              
              if #retTypeList == 0 and #expNodeList > 0 then
                self:addErrMess( token.pos, "this function can't return value." )
              end
              for index, retType in pairs( retTypeList ) do
                local expNode = expNodeList[index]
                
                local expType = expNode:get_expType()
                
                if not retType:isSettableFrom( expType ) then
                  self:addErrMess( token.pos, string.format( "return type of arg(%d) is not compatible -- %s(%d) and %s(%d)", index, retType:getTxt(  ), retType:get_typeId(  ), expType:getTxt(  ), expType:get_typeId(  )) )
                end
                if index == #retTypeList then
                  if #retTypeList < #expNodeList and retType ~= Ast.builtinTypeDDD then
                    self:addErrMess( token.pos, "over return value" )
                  end
                end
              end
            end
        end
        
      end
      statement = Ast.ReturnNode.new(token.pos, {Ast.builtinTypeNone}, expList)
    elseif token.txt == "break" then
      self:checkNextToken( ";" )
      statement = Ast.BreakNode.new(token.pos, {Ast.builtinTypeNone})
    elseif token.txt == "unwrap" then
      statement = self:analyzeUnwrap( token )
    elseif token.txt == "sync" then
      statement = self:analyzeDeclVar( "sync", "local", false, token )
    elseif token.txt == "import" then
      statement = self:analyzeImport( token )
    elseif token.txt == "subfile" then
      statement = self:analyzeSubfile( token )
    elseif token.txt == "luneConstrol" then
      self:analyzeLuneControl( token )
      statement = self:createNoneNode( token.pos )
    elseif token.txt == ";" then
      statement = self:createNoneNode( token.pos )
    elseif token.txt == ",," or token.txt == ",,," or token.txt == ",,,," then
      self:error( string.format( "illegal macro op -- %s", token.txt) )
    else 
      self:pushback(  )
      local exp = self:analyzeExp(  )
      
      self:checkNextToken( ";" )
      statement = Ast.StmtExpNode.new(exp:get_pos(), {Ast.builtinTypeNone}, exp)
    end
  end
  return statement
end

return moduleObj
