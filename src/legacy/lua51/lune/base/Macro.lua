--lune/base/Macro.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Macro'
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

function _lune.loadstring51( txt, env )
   local func = loadstring( txt )
   if func and env then
      setfenv( func, env )
   end
   return func
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
local Util = _lune.loadModule( 'lune.base.Util' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Formatter = _lune.loadModule( 'lune.base.Formatter' )
local function loadCode( code )

   local loaded, mess = _lune.loadstring51( code )
   if loaded ~= nil then
      do
         local obj = loaded(  )
         if obj ~= nil then
            return obj
         end
      end
      
      error( "failed to load" )
   end
   
   error( string.format( "%s -- %s", tostring( mess), code) )
end


local toList = loadCode( "return function( ... ) return { ... } end" )
_moduleObj.toList = toList


local toListEmpty = loadCode( "return function() return {} end" )
_moduleObj.toListEmpty = toListEmpty


local toLuaval = loadCode( "return function( val ) return val end" )
_moduleObj.toLuaval = toLuaval


local MacroMetaArgInfo = {}
setmetatable( MacroMetaArgInfo, { ifList = {Mapping,} } )
function MacroMetaArgInfo.setmeta( obj )
  setmetatable( obj, { __index = MacroMetaArgInfo  } )
end
function MacroMetaArgInfo.new( name, typeId )
   local obj = {}
   MacroMetaArgInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, typeId )
   end
   return obj
end
function MacroMetaArgInfo:__init( name, typeId )

   self.name = name
   self.typeId = typeId
end
function MacroMetaArgInfo:_toMap()
  return self
end
function MacroMetaArgInfo._fromMap( val )
  local obj, mes = MacroMetaArgInfo._fromMapSub( {}, val )
  if obj then
     MacroMetaArgInfo.setmeta( obj )
  end
  return obj, mes
end
function MacroMetaArgInfo._fromStem( val )
  return MacroMetaArgInfo._fromMap( val )
end

function MacroMetaArgInfo._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "name", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "typeId", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local MacroMetaInfo = {}
setmetatable( MacroMetaInfo, { ifList = {Mapping,} } )
function MacroMetaInfo.setmeta( obj )
  setmetatable( obj, { __index = MacroMetaInfo  } )
end
function MacroMetaInfo.new( name, argList, symList, stmtBlock, tokenList )
   local obj = {}
   MacroMetaInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, argList, symList, stmtBlock, tokenList )
   end
   return obj
end
function MacroMetaInfo:__init( name, argList, symList, stmtBlock, tokenList )

   self.name = name
   self.argList = argList
   self.symList = symList
   self.stmtBlock = stmtBlock
   self.tokenList = tokenList
end
function MacroMetaInfo:_toMap()
  return self
end
function MacroMetaInfo._fromMap( val )
  local obj, mes = MacroMetaInfo._fromMapSub( {}, val )
  if obj then
     MacroMetaInfo.setmeta( obj )
  end
  return obj, mes
end
function MacroMetaInfo._fromStem( val )
  return MacroMetaInfo._fromMap( val )
end

function MacroMetaInfo._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "name", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "argList", func = _lune._toList, nilable = false, child = { { func = MacroMetaArgInfo._fromMap, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "symList", func = _lune._toList, nilable = false, child = { { func = MacroMetaArgInfo._fromMap, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "stmtBlock", func = _lune._toStr, nilable = true, child = {} } )
   table.insert( memInfo, { name = "tokenList", func = _lune._toList, nilable = false, child = { { func = _lune._toList, nilable = false, child = { { func = _lune._toStem, nilable = false, child = {} } } } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local MacroPaser = {}
setmetatable( MacroPaser, { __index = Parser.Parser } )
function MacroPaser.new( tokenList, name )
   local obj = {}
   MacroPaser.setmeta( obj )
   if obj.__init then obj:__init( tokenList, name ); end
   return obj
end
function MacroPaser:__init(tokenList, name) 
   Parser.Parser.__init( self)
   
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
function MacroPaser.setmeta( obj )
  setmetatable( obj, { __index = MacroPaser  } )
end


local function getLiteralMacroVal( obj )

   do
      local _matchExp = obj
      if _matchExp[1] == Nodes.Literal.Nil[1] then
      
         return nil
      elseif _matchExp[1] == Nodes.Literal.Int[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Nodes.Literal.Real[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Nodes.Literal.Str[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Nodes.Literal.Bool[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Nodes.Literal.Symbol[1] then
         local val = _matchExp[2][1]
      
         return {val}
      elseif _matchExp[1] == Nodes.Literal.Field[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Nodes.Literal.LIST[1] then
         local list = _matchExp[2][1]
      
         local newList = {}
         for index, item in ipairs( list ) do
            newList[index] = getLiteralMacroVal( item )
         end
         
         return newList
      elseif _matchExp[1] == Nodes.Literal.ARRAY[1] then
         local list = _matchExp[2][1]
      
         local newList = {}
         for index, item in ipairs( list ) do
            newList[index] = getLiteralMacroVal( item )
         end
         
         return newList
      elseif _matchExp[1] == Nodes.Literal.SET[1] then
         local list = _matchExp[2][1]
      
         local newSet = {}
         for __index, item in ipairs( list ) do
            do
               local _exp = getLiteralMacroVal( item )
               if _exp ~= nil then
                  newSet[_exp]= true
               end
            end
            
         end
         
         return newSet
      elseif _matchExp[1] == Nodes.Literal.MAP[1] then
         local map = _matchExp[2][1]
      
         local newMap = {}
         for key, val in pairs( map ) do
            local keyObj = getLiteralMacroVal( key )
            local valObj = getLiteralMacroVal( val )
            if keyObj ~= nil and valObj ~= nil then
               newMap[keyObj] = valObj
            end
            
         end
         
         return newMap
      end
   end
   
end

local ExtMacroInfo = {}
setmetatable( ExtMacroInfo, { __index = Nodes.MacroInfo } )
function ExtMacroInfo:getArgList(  )

   return self.argList
end
function ExtMacroInfo:getTokenList(  )

   return self.tokenList
end
function ExtMacroInfo.new( name, func, symbol2MacroValInfoMap, argList, tokenList )
   local obj = {}
   ExtMacroInfo.setmeta( obj )
   if obj.__init then obj:__init( name, func, symbol2MacroValInfoMap, argList, tokenList ); end
   return obj
end
function ExtMacroInfo:__init(name, func, symbol2MacroValInfoMap, argList, tokenList) 
   Nodes.MacroInfo.__init( self,func, symbol2MacroValInfoMap)
   
   self.name = name
   self.argList = argList
   self.tokenList = tokenList
end
function ExtMacroInfo.setmeta( obj )
  setmetatable( obj, { __index = ExtMacroInfo  } )
end
function ExtMacroInfo:get_name()
   return self.name
end


local MacroAnalyzeInfo = {}
_moduleObj.MacroAnalyzeInfo = MacroAnalyzeInfo
function MacroAnalyzeInfo.new( typeInfo, mode )
   local obj = {}
   MacroAnalyzeInfo.setmeta( obj )
   if obj.__init then obj:__init( typeInfo, mode ); end
   return obj
end
function MacroAnalyzeInfo:__init(typeInfo, mode) 
   self.typeInfo = typeInfo
   self.mode = mode
   self.argIndex = 1
end
function MacroAnalyzeInfo:equalsArgTypeList( argTypeList )

   return self.typeInfo:get_argTypeInfoList() == argTypeList
end
function MacroAnalyzeInfo:getCurArgType(  )

   if #self.typeInfo:get_argTypeInfoList() < self.argIndex then
      return Ast.builtinTypeNone
   end
   
   return self.typeInfo:get_argTypeInfoList()[self.argIndex]
end
function MacroAnalyzeInfo:nextArg(  )

   self.argIndex = self.argIndex + 1
end
function MacroAnalyzeInfo:isAnalyzingSymArg(  )

   return self.mode == Nodes.MacroMode.AnalyzeArg and self:getCurArgType(  ) == Ast.builtinTypeSymbol
   
end
function MacroAnalyzeInfo:isAnalyzingExpArg(  )

   return self.mode == Nodes.MacroMode.AnalyzeArg and self:getCurArgType(  ) == Ast.builtinTypeExp
end
function MacroAnalyzeInfo.setmeta( obj )
  setmetatable( obj, { __index = MacroAnalyzeInfo  } )
end
function MacroAnalyzeInfo:get_typeInfo()
   return self.typeInfo
end
function MacroAnalyzeInfo:get_mode()
   return self.mode
end
function MacroAnalyzeInfo:get_argIndex()
   return self.argIndex
end


local MacroCtrl = {}
_moduleObj.MacroCtrl = MacroCtrl
function MacroCtrl.new( macroEval )
   local obj = {}
   MacroCtrl.setmeta( obj )
   if obj.__init then obj:__init( macroEval ); end
   return obj
end
function MacroCtrl:__init(macroEval) 
   self.tokenExpanding = false
   self.useModuleMacroSet = {}
   self.typeId2MacroInfo = {}
   self.symbol2ValueMapForMacro = {}
   self.macroEval = macroEval
   self.analyzeInfo = MacroAnalyzeInfo.new(Ast.builtinTypeNone, Nodes.MacroMode.None)
   self.macroCallLineNo = 0
   self.macroAnalyzeInfoStack = {self.analyzeInfo}
   self.macroLocalVarMap = _moduleObj.toListEmpty(  )
end
function MacroCtrl.setmeta( obj )
  setmetatable( obj, { __index = MacroCtrl  } )
end
function MacroCtrl:get_useModuleMacroSet()
   return self.useModuleMacroSet
end
function MacroCtrl:get_typeId2MacroInfo()
   return self.typeId2MacroInfo
end
function MacroCtrl:get_analyzeInfo()
   return self.analyzeInfo
end
function MacroCtrl:get_tokenExpanding()
   return self.tokenExpanding
end
function MacroCtrl:get_macroCallLineNo()
   return self.macroCallLineNo
end


local function equalsType( typeInfo, builtinType )

   return typeInfo:get_srcTypeInfo() == builtinType
end
function MacroCtrl:evalMacroOp( streamName, firstToken, macroTypeInfo, expList )

   self.useModuleMacroSet[macroTypeInfo:getModule(  )]= true
   
   if expList ~= nil then
      for __index, exp in ipairs( expList:get_expList() ) do
         local kind = exp:get_kind()
         do
            local _switchExp = kind
            if _switchExp == Nodes.NodeKind.get_LiteralNil() or _switchExp == Nodes.NodeKind.get_LiteralChar() or _switchExp == Nodes.NodeKind.get_LiteralInt() or _switchExp == Nodes.NodeKind.get_LiteralReal() or _switchExp == Nodes.NodeKind.get_LiteralArray() or _switchExp == Nodes.NodeKind.get_LiteralList() or _switchExp == Nodes.NodeKind.get_LiteralMap() or _switchExp == Nodes.NodeKind.get_LiteralString() or _switchExp == Nodes.NodeKind.get_LiteralBool() or _switchExp == Nodes.NodeKind.get_LiteralSymbol() or _switchExp == Nodes.NodeKind.get_RefField() or _switchExp == Nodes.NodeKind.get_ExpMacroStat() or _switchExp == Nodes.NodeKind.get_ExpMacroArgExp() or _switchExp == Nodes.NodeKind.get_ExpOmitEnum() or _switchExp == Nodes.NodeKind.get_ExpCast() or _switchExp == Nodes.NodeKind.get_ExpOp2() then
            else 
               
                  local mess = string.format( "Macro arguments must be literal value. -- %d:%d:%s", exp:get_pos().lineNo, exp:get_pos().column, Nodes.getNodeKindName( kind ))
                  return nil, mess
            end
         end
         
      end
      
   end
   
   
   local macroInfo = _lune.unwrap( self.typeId2MacroInfo[macroTypeInfo:get_typeId(  )])
   local argValMap = {}
   local macroArgValMap = {["__var"] = self.macroLocalVarMap}
   local macroArgNodeList = macroInfo:getArgList(  )
   local macroArgName2ArgNode = {}
   if expList ~= nil then
      for index, argNode in ipairs( expList:get_expList() ) do
         local literal, mess = argNode:getLiteral(  )
         if literal ~= nil then
            do
               local val = getLiteralMacroVal( literal )
               if val ~= nil then
                  argValMap[index] = _moduleObj.toLuaval( val )
                  local declArgNode = macroArgNodeList[index]
                  
                  if argNode:get_expType() == Ast.builtinTypeSymbol then
                     macroArgValMap[declArgNode:get_name()] = val[1]
                  else
                   
                     macroArgValMap[declArgNode:get_name()] = val
                  end
                  
                  macroArgName2ArgNode[declArgNode:get_name()] = argNode
               end
            end
            
         else
            local errmess = string.format( "not support node at arg(%d) -- %s:%s", index, Nodes.getNodeKindName( argNode:get_kind() ), tostring( mess))
            return nil, errmess
         end
         
      end
      
   end
   
   local func = macroInfo.func
   local macroVars = func( macroArgValMap )
   self.macroLocalVarMap = _lune.unwrap( macroVars['__var'])
   for __index, name in pairs( (_lune.unwrap( macroVars['__names']) ) ) do
      local valInfo = _lune.unwrap( macroInfo.symbol2MacroValInfoMap[name])
      local typeInfo = valInfo.typeInfo
      local valList
      
      do
         local val = macroVars[name]
         if val ~= nil then
            if equalsType( typeInfo, Ast.builtinTypeSymbol ) then
               valList = _moduleObj.toList( val )
            else
             
               valList = val
            end
            
         else
            valList = _moduleObj.toListEmpty(  )
         end
      end
      
      self.symbol2ValueMapForMacro[name] = Nodes.MacroValInfo.new(valList, typeInfo, macroArgName2ArgNode[name])
   end
   
   for index, arg in ipairs( macroInfo:getArgList(  ) ) do
      if arg:get_typeInfo():get_kind() ~= Ast.TypeInfoKind.DDD then
         local argType = arg:get_typeInfo()
         local argName = arg:get_name()
         self.symbol2ValueMapForMacro[argName] = Nodes.MacroValInfo.new(argValMap[index], argType, macroArgName2ArgNode[argName])
      else
       
         return nil, "not support ... in macro"
      end
      
   end
   
   return MacroPaser.new(macroInfo:getTokenList(  ), string.format( "%s:%d:%d: (macro %s)", streamName, firstToken.pos.lineNo, firstToken.pos.column, macroTypeInfo:getTxt(  ))), nil
end


function MacroCtrl:importMacro( processInfo, macroInfoStem, macroTypeInfo, typeId2TypeInfo )

   local macroInfo, err = MacroMetaInfo._fromStem( macroInfoStem )
   if macroInfo ~= nil then
      local argList = {}
      local argNameList = {}
      local symbol2MacroValInfoMap = {}
      for __index, argInfo in ipairs( macroInfo.argList ) do
         local argTypeInfo = _lune.unwrap( typeId2TypeInfo[argInfo.typeId])
         table.insert( argList, Nodes.MacroArgInfo.new(argInfo.name, argTypeInfo) )
         table.insert( argNameList, argInfo.name )
         
      end
      
      for __index, symInfo in ipairs( macroInfo.symList ) do
         local symTypeInfo = _lune.unwrap( typeId2TypeInfo[symInfo.typeId])
         symbol2MacroValInfoMap[symInfo.name] = Nodes.MacroValInfo.new(nil, symTypeInfo, nil)
      end
      
      
      local tokenList = {}
      local lineNo = 0
      local column = 1
      for __index, tokenInfo in ipairs( macroInfo.tokenList ) do
         local txt = tokenInfo[2]
         if txt == "\n" then
            lineNo = lineNo + 1
            column = 1
         else
          
            local pos = Parser.Position.new(lineNo, column, string.format( "macro:%s", macroInfo.name))
            table.insert( tokenList, Parser.Token.new(_lune.unwrap( Parser.TokenKind._from( math.floor(tokenInfo[1]) )), txt, pos, false) )
            column = column + #txt + 1
         end
         
      end
      
      
      self.typeId2MacroInfo[macroTypeInfo:get_typeId()] = ExtMacroInfo.new(macroInfo.name, self.macroEval:evalFromCode( processInfo, macroInfo.name, argNameList, macroInfo.stmtBlock ), symbol2MacroValInfoMap, argList, tokenList)
   else
      Util.errorLog( string.format( "macro load fail -- %s: %s ", macroTypeInfo:getTxt(  ), _lune.unwrapDefault( err, "")) )
   end
   
end


function MacroCtrl:regist( processInfo, node, macroScope )

   local macroObj = self.macroEval:eval( processInfo, node )
   local remap = {}
   for name, macroValInfo in pairs( self.symbol2ValueMapForMacro ) do
      if equalsType( macroValInfo.typeInfo, Ast.builtinTypeEmpty ) then
         remap[name] = Nodes.MacroValInfo.new(macroValInfo.val, _lune.unwrap( macroScope:getTypeInfoChild( name )), macroValInfo.argNode)
      else
       
         remap[name] = macroValInfo
      end
      
   end
   
   self.symbol2ValueMapForMacro = remap
   self.typeId2MacroInfo[node:get_expType():get_typeId(  )] = Nodes.DefMacroInfo.new(macroObj, node:get_declInfo(), self.symbol2ValueMapForMacro)
   
   self.symbol2ValueMapForMacro = {}
end


local function expandVal( tokenList, workval, pos )

   if workval ~= nil then
      local val = workval
      do
         local _switchExp = type( val )
         if _switchExp == "boolean" then
            local token = string.format( "%s", tostring( val))
            local kind = Parser.TokenKind.Kywd
            table.insert( tokenList, Parser.Token.new(kind, token, pos, false) )
         elseif _switchExp == "number" then
            local num = string.format( "%g", val * 1.0)
            local kind = Parser.TokenKind.Int
            if string.find( num, ".", 1, true ) then
               kind = Parser.TokenKind.Real
            end
            
            table.insert( tokenList, Parser.Token.new(kind, num, pos, false) )
         elseif _switchExp == "string" then
            
            table.insert( tokenList, Parser.Token.new(Parser.TokenKind.Str, Parser.quoteStr( val ), pos, false) )
            
         else 
            
               return string.format( "not support ,, List -- %s", type( val ))
         end
      end
      
   end
   
   return nil
end
local function pushbackTxt( pushbackParser, txtList, streamName, pos )

   local tokenList = {}
   for __index, txt in ipairs( txtList ) do
      local stream = Parser.TxtStream.new(txt)
      local parser = Parser.StreamParser.new(stream, string.format( "macro symbol -- %s", streamName))
      local workParser = Parser.DefaultPushbackParser.new(parser)
      while true do
         local worktoken = workParser:getTokenNoErr(  )
         if worktoken.kind == Parser.TokenKind.Eof then
            break
         end
         
         table.insert( tokenList, Parser.Token.new(worktoken.kind, worktoken.txt, pos, false) )
      end
      
   end
   
   for index = #tokenList, 1, -1 do
      pushbackParser:pushbackToken( tokenList[index] )
   end
   
end

function MacroCtrl:expandMacroVal( typeNameCtrl, scope, parser, token )

   if self.tokenExpanding then
      return token
   end
   
   
   local function getToken(  )
   
      self.tokenExpanding = true
      local work = parser:getTokenNoErr(  )
      self.tokenExpanding = false
      return work
   end
   
   local tokenTxt = token.txt
   if tokenTxt == ',,' or tokenTxt == ',,,' or tokenTxt == ',,,,' then
      local nextToken = getToken(  )
      
      local macroVal = self.symbol2ValueMapForMacro[nextToken.txt]
      if  nil == macroVal then
         local _macroVal = macroVal
      
         parser:error( string.format( "unknown macro val -- %s", nextToken.txt) )
      end
      
      
      if tokenTxt == ',,' then
         if equalsType( macroVal.typeInfo, Ast.builtinTypeSymbol ) then
            local txtList = {}
            for __index, txt in pairs( (_lune.unwrap( macroVal.val) ) ) do
               table.insert( txtList, txt )
            end
            
            pushbackTxt( parser, txtList, nextToken.txt, nextToken.pos )
         elseif equalsType( macroVal.typeInfo, Ast.builtinTypeStat ) or equalsType( macroVal.typeInfo, Ast.builtinTypeExp ) or equalsType( macroVal.typeInfo, Ast.builtinTypeMultiExp ) then
            parser:pushbackStr( string.format( "macroVal %s", nextToken.txt), (_lune.unwrap( macroVal.val) ) )
         elseif macroVal.typeInfo:get_kind() == Ast.TypeInfoKind.Array or macroVal.typeInfo:get_kind(  ) == Ast.TypeInfoKind.List then
            if equalsType( macroVal.typeInfo:get_itemTypeInfoList()[1], Ast.builtinTypeStat ) then
               local strList = (_lune.unwrap( macroVal.val) )
               for index = #strList, 1, -1 do
                  parser:pushbackStr( string.format( "macroVal %s[%d]", nextToken.txt, index), strList[index] )
               end
               
            else
             
               local tokenList = {}
               
               do
                  local argNode = macroVal.argNode
                  if argNode ~= nil then
                     if not argNode:setupLiteralTokenList( tokenList ) then
                        parser:error( string.format( "illegal macro val ,, -- %s", nextToken.txt) )
                     end
                     
                  else
                     parser:error( string.format( "not support ,, -- %s", nextToken.txt) )
                  end
               end
               
               
               parser:newPushback( tokenList )
            end
            
         elseif macroVal.typeInfo:get_kind(  ) == Ast.TypeInfoKind.Enum then
            local enumTypeInfo = _lune.unwrap( _lune.__Cast( macroVal.typeInfo, 3, Ast.EnumTypeInfo ))
            local fullname = macroVal.typeInfo:getFullName( typeNameCtrl, scope, true )
            local nameList = {}
            for name in fullname:gmatch( "[^%.]+" ) do
               table.insert( nameList, name )
            end
            
            local enumValInfo = _lune.unwrap( enumTypeInfo:get_val2EnumValInfo()[_lune.unwrap( macroVal.val)])
            nextToken = Parser.Token.new(Parser.TokenKind.Symb, enumValInfo:get_name(), nextToken.pos, false)
            parser:pushbackToken( nextToken )
            for index = #nameList, 1, -1 do
               nextToken = Parser.Token.new(Parser.TokenKind.Dlmt, ".", nextToken.pos, false)
               parser:pushbackToken( nextToken )
               nextToken = Parser.Token.new(Parser.TokenKind.Symb, nameList[index], nextToken.pos, false)
               parser:pushbackToken( nextToken )
            end
            
         else
          
            local tokenList = {}
            
            do
               local argNode = macroVal.argNode
               if argNode ~= nil then
                  if not argNode:setupLiteralTokenList( tokenList ) then
                     parser:error( string.format( "illegal macro val ,, -- %s", nextToken.txt) )
                  end
                  
               else
                  expandVal( tokenList, macroVal.val, nextToken.pos )
               end
            end
            
            
            parser:newPushback( tokenList )
         end
         
      elseif tokenTxt == ',,,' then
         if equalsType( macroVal.typeInfo, Ast.builtinTypeString ) then
            pushbackTxt( parser, {(_lune.unwrap( macroVal.val) )}, nextToken.txt, nextToken.pos )
         else
          
            parser:error( string.format( "',,,' does not support this type -- %s", macroVal.typeInfo:getTxt(  )) )
         end
         
      elseif tokenTxt == ',,,,' then
         if equalsType( macroVal.typeInfo, Ast.builtinTypeSymbol ) then
            local txtList = (_lune.unwrap( macroVal.val) )
            local newToken = ""
            for __index, txt in pairs( txtList ) do
               newToken = string.format( "%s%s", newToken, txt)
            end
            
            nextToken = Parser.Token.new(Parser.TokenKind.Str, string.format( "'%s'", newToken), nextToken.pos, false)
            parser:pushbackToken( nextToken )
         elseif equalsType( macroVal.typeInfo, Ast.builtinTypeStat ) or equalsType( macroVal.typeInfo, Ast.builtinTypeExp ) or equalsType( macroVal.typeInfo, Ast.builtinTypeMultiExp ) then
            local txt = (_lune.unwrap( macroVal.val) )
            local rawTxt
            
            if txt:find( "^```" ) then
               rawTxt = Parser.quoteStr( txt )
            else
             
               rawTxt = Parser.quoteStr( txt )
            end
            
            nextToken = Parser.Token.new(Parser.TokenKind.Str, rawTxt, nextToken.pos, false)
            parser:pushbackToken( nextToken )
         else
          
            parser:error( string.format( "not support this symbol -- %s%s", tokenTxt, nextToken.txt) )
         end
         
      end
      
      nextToken = getToken(  )
      
      token = nextToken
   end
   
   
   self.tokenExpanding = false
   
   return token
end


local ErrorMess = {}
_moduleObj.ErrorMess = ErrorMess
function ErrorMess.setmeta( obj )
  setmetatable( obj, { __index = ErrorMess  } )
end
function ErrorMess.new( pos, mess )
   local obj = {}
   ErrorMess.setmeta( obj )
   if obj.__init then
      obj:__init( pos, mess )
   end
   return obj
end
function ErrorMess:__init( pos, mess )

   self.pos = pos
   self.mess = mess
end


function MacroCtrl:expandSymbol( parser, prefixToken, exp, nodeManager, errMessList )

   local nextToken = parser:getTokenNoErr(  )
   if nextToken.txt ~= "~~" then
      parser:pushbackToken( nextToken )
   end
   
   
   local format = prefixToken.txt == ",,," and "' %s '" or '"\'%s\'"'
   
   if prefixToken.txt == ",," then
      do
         local refNode = _lune.__Cast( exp, 3, Nodes.ExpRefNode )
         if refNode ~= nil then
            local symbolInfo = refNode:get_symbolInfo()
            local macroInfo = self.symbol2ValueMapForMacro[symbolInfo:get_name()]
            if macroInfo ~= nil then
               local valType = macroInfo.typeInfo
               if equalsType( valType, Ast.builtinTypeSymbol ) or equalsType( valType, Ast.builtinTypeExp ) or equalsType( valType, Ast.builtinTypeMultiExp ) or equalsType( valType, Ast.builtinTypeStat ) then
                  format = "' %s '"
               elseif valType:get_kind() == Ast.TypeInfoKind.List and equalsType( valType:get_itemTypeInfoList()[1], Ast.builtinTypeStat ) then
                  format = "' %s '"
                  exp = Nodes.ExpMacroStatListNode.create( nodeManager, prefixToken.pos, self.analyzeInfo:get_mode() == Nodes.MacroMode.AnalyzeArg, {Ast.builtinTypeString}, exp )
               elseif equalsType( Ast.builtinTypeString, valType ) then
                  
               elseif equalsType( valType, Ast.builtinTypeInt ) or equalsType( valType, Ast.builtinTypeReal ) then
                  format = "' %s' "
               else
                
                  table.insert( errMessList, ErrorMess.new(_lune.unwrap( symbolInfo:get_pos()), string.format( "not support ,, -- %s", valType:getTxt(  ))) )
               end
               
            else
               if equalsType( exp:get_expType(), Ast.builtinTypeInt ) or equalsType( exp:get_expType(), Ast.builtinTypeReal ) then
                  format = "' %s' "
               elseif equalsType( exp:get_expType(), Ast.builtinTypeSymbol ) or equalsType( exp:get_expType(), Ast.builtinTypeExp ) or equalsType( exp:get_expType(), Ast.builtinTypeMultiExp ) or equalsType( exp:get_expType(), Ast.builtinTypeStat ) then
                  format = "' %s '"
               end
               
            end
            
         end
      end
      
   end
   
   local newToken = Parser.Token.new(Parser.TokenKind.Str, format, prefixToken.pos, prefixToken.consecutive)
   
   local expListNode = Nodes.ExpListNode.create( nodeManager, exp:get_pos(), self.analyzeInfo:get_mode() == Nodes.MacroMode.AnalyzeArg, exp:get_expTypeList(), {exp}, nil, false )
   local dddNode = Nodes.ExpToDDDNode.create( nodeManager, exp:get_pos(), self.analyzeInfo:get_mode() == Nodes.MacroMode.AnalyzeArg, exp:get_expTypeList(), expListNode )
   
   local literalStr = Nodes.LiteralStringNode.create( nodeManager, prefixToken.pos, self.analyzeInfo:get_mode() == Nodes.MacroMode.AnalyzeArg, {Ast.builtinTypeString}, newToken, expListNode, Nodes.ExpListNode.create( nodeManager, exp:get_pos(), self.analyzeInfo:get_mode() == Nodes.MacroMode.AnalyzeArg, exp:get_expTypeList(), {dddNode}, nil, false ) )
   return literalStr
end

function MacroCtrl:registVar( symbolList )

   for __index, symbolInfo in ipairs( symbolList ) do
      local info = Nodes.MacroValInfo.new(nil, symbolInfo:get_typeInfo(), nil)
      self.symbol2ValueMapForMacro[symbolInfo:get_name()] = info
   end
   
end


function MacroCtrl:startDecl(  )

   self.symbol2ValueMapForMacro = {}
end




function MacroCtrl:finishMacroMode(  )

   table.remove( self.macroAnalyzeInfoStack )
   self.analyzeInfo = self.macroAnalyzeInfoStack[#self.macroAnalyzeInfoStack]
   
end


function MacroCtrl:startExpandMode( lineNo, typeInfo, callback )

   self.analyzeInfo = MacroAnalyzeInfo.new(typeInfo, Nodes.MacroMode.Expand)
   self.macroCallLineNo = lineNo
   table.insert( self.macroAnalyzeInfoStack, self.analyzeInfo )
   
   callback(  )
   
   self:finishMacroMode(  )
end


function MacroCtrl:startAnalyzeArgMode( macroFuncType )

   self.analyzeInfo = MacroAnalyzeInfo.new(macroFuncType, Nodes.MacroMode.AnalyzeArg)
   table.insert( self.macroAnalyzeInfoStack, self.analyzeInfo )
   
end


function MacroCtrl:switchMacroMode(  )

   self.analyzeInfo = self.macroAnalyzeInfoStack[#self.macroAnalyzeInfoStack - 1]
   table.insert( self.macroAnalyzeInfoStack, self.analyzeInfo )
end


function MacroCtrl:restoreMacroMode(  )

   table.remove( self.macroAnalyzeInfoStack )
   self.analyzeInfo = self.macroAnalyzeInfoStack[#self.macroAnalyzeInfoStack]
end

function MacroCtrl:isInAnalyzeArgMode(  )

   if #self.macroAnalyzeInfoStack == 0 then
      return false
   end
   
   for __index, info in ipairs( self.macroAnalyzeInfoStack ) do
      if info:get_mode() == Nodes.MacroMode.AnalyzeArg then
         return true
      end
      
   end
   
   return false
end


local function nodeToCodeTxt( node, moduleTypeInfo )

   local memStream = Util.memStream.new()
   local formatter = Formatter.createFilter( moduleTypeInfo, memStream )
   
   node:processFilter( formatter, Formatter.Opt.new(node) )
   
   return memStream:get_txt()
end
_moduleObj.nodeToCodeTxt = nodeToCodeTxt



return _moduleObj
