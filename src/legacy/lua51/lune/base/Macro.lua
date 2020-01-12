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
         for index, item in pairs( list ) do
            newList[index] = getLiteralMacroVal( item )
         end
         
         return newList
      elseif _matchExp[1] == Nodes.Literal.ARRAY[1] then
         local list = _matchExp[2][1]
      
         local newList = {}
         for index, item in pairs( list ) do
            newList[index] = getLiteralMacroVal( item )
         end
         
         return newList
      elseif _matchExp[1] == Nodes.Literal.SET[1] then
         local list = _matchExp[2][1]
      
         local newSet = {}
         for __index, item in pairs( list ) do
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
   
   Util.err( "unknown literal obj -- " .. Nodes.Literal:_getTxt( obj)
    )
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
   self.macroMode = Nodes.MacroMode.None
   self.macroCallLineNo = 0
   self.macroModeStack = {self.macroMode}
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
function MacroCtrl:get_macroMode()
   return self.macroMode
end
function MacroCtrl:get_tokenExpanding()
   return self.tokenExpanding
end
function MacroCtrl:get_macroCallLineNo()
   return self.macroCallLineNo
end


function MacroCtrl:evalMacroOp( streamName, firstToken, macroTypeInfo, expList )

   self.useModuleMacroSet[macroTypeInfo:getModule(  )]= true
   
   if expList ~= nil then
      for __index, exp in pairs( expList:get_expList() ) do
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
   local macroArgValMap = {}
   local macroArgNodeList = macroInfo:getArgList(  )
   local macroArgName2ArgNode = {}
   if expList ~= nil then
      for index, argNode in pairs( expList:get_expList() ) do
         local literal, mess = argNode:getLiteral(  )
         if literal ~= nil then
            do
               local val = getLiteralMacroVal( literal )
               if val ~= nil then
                  argValMap[index] = val
                  local declArgNode = macroArgNodeList[index]
                  
                  macroArgValMap[declArgNode:get_name()] = val
                  macroArgName2ArgNode[declArgNode:get_name()] = argNode
               end
            end
            
         else
            local errmess = string.format( "not support node -- %s:%s", Nodes.getNodeKindName( argNode:get_kind() ), tostring( mess))
            return nil, errmess
         end
         
      end
      
   end
   
   local func = macroInfo.func
   local macroVars = func( macroArgValMap )
   
   for __index, name in pairs( (_lune.unwrap( macroVars['__names']) ) ) do
      local valInfo = _lune.unwrap( macroInfo.symbol2MacroValInfoMap[name])
      local typeInfo = valInfo.typeInfo
      local val = macroVars[name]
      if typeInfo:equals( Ast.builtinTypeSymbol ) then
         val = {val}
      end
      
      self.symbol2ValueMapForMacro[name] = Nodes.MacroValInfo.new(val, typeInfo, macroArgName2ArgNode[name])
   end
   
   
   for index, arg in pairs( macroInfo:getArgList(  ) ) do
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


function MacroCtrl:importMacro( macroInfoStem, macroTypeInfo, typeId2TypeInfo )

   local macroInfo, err = MacroMetaInfo._fromStem( macroInfoStem )
   if macroInfo ~= nil then
      local argList = {}
      local argNameList = {}
      local symbol2MacroValInfoMap = {}
      for __index, argInfo in pairs( macroInfo.argList ) do
         local argTypeInfo = _lune.unwrap( typeId2TypeInfo[argInfo.typeId])
         table.insert( argList, Nodes.MacroArgInfo.new(argInfo.name, argTypeInfo) )
         table.insert( argNameList, argInfo.name )
      end
      
      for __index, symInfo in pairs( macroInfo.symList ) do
         local symTypeInfo = _lune.unwrap( typeId2TypeInfo[symInfo.typeId])
         symbol2MacroValInfoMap[symInfo.name] = Nodes.MacroValInfo.new(nil, symTypeInfo, nil)
      end
      
      
      local tokenList = {}
      local lineNo = 0
      local column = 1
      for __index, tokenInfo in pairs( macroInfo.tokenList ) do
         local txt = tokenInfo[2]
         if txt == "\n" then
            lineNo = lineNo + 1
            column = 1
         else
          
            table.insert( tokenList, Parser.Token.new(_lune.unwrap( Parser.TokenKind._from( math.floor(tokenInfo[1]) )), txt, Parser.Position.new(lineNo, column), false) )
            column = column + #txt + 1
         end
         
      end
      
      
      self.typeId2MacroInfo[macroTypeInfo:get_typeId()] = ExtMacroInfo.new(macroInfo.name, self.macroEval:evalFromCode( macroInfo.name, argNameList, macroInfo.stmtBlock ), symbol2MacroValInfoMap, argList, tokenList)
   else
      Util.errorLog( string.format( "macro load fail -- %s", _lune.unwrapDefault( err, "")) )
   end
   
end


function MacroCtrl:regist( node )

   local macroObj = self.macroEval:eval( node )
   
   self.typeId2MacroInfo[node:get_expType():get_typeId(  )] = Nodes.DefMacroInfo.new(macroObj, node:get_declInfo(), self.symbol2ValueMapForMacro)
   self.symbol2ValueMapForMacro = {}
end


local function expandVal( tokenList, val, pos )

   if val ~= nil then
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
            table.insert( tokenList, Parser.Token.new(Parser.TokenKind.Str, string.format( '[[%s]]', tostring( val)), pos, false) )
         else 
            
               return string.format( "not support ,, List -- %s", type( val ))
         end
      end
      
   end
   
   return nil
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
         
         if macroVal.typeInfo:equals( Ast.builtinTypeSymbol ) then
            local txtList = (_lune.unwrap( macroVal.val) )
            for index = #txtList, 1, -1 do
               nextToken = Parser.Token.new(nextToken.kind, txtList[index], nextToken.pos, false)
               parser:pushbackToken( nextToken )
            end
            
         elseif macroVal.typeInfo:equals( Ast.builtinTypeStat ) or macroVal.typeInfo:equals( Ast.builtinTypeExp ) then
            parser:pushbackStr( string.format( "macroVal %s", nextToken.txt), (_lune.unwrap( macroVal.val) ) )
         elseif macroVal.typeInfo:get_kind() == Ast.TypeInfoKind.Array or macroVal.typeInfo:get_kind(  ) == Ast.TypeInfoKind.List then
            if macroVal.typeInfo:get_itemTypeInfoList()[1]:equals( Ast.builtinTypeStat ) then
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
         if macroVal.typeInfo:equals( Ast.builtinTypeString ) then
            nextToken = Parser.Token.new(nextToken.kind, (_lune.unwrap( macroVal.val) ), nextToken.pos, false)
            parser:pushbackToken( nextToken )
         else
          
            parser:error( string.format( "',,,' does not support this type -- %s", macroVal.typeInfo:getTxt(  )) )
         end
         
      elseif tokenTxt == ',,,,' then
         if macroVal.typeInfo:equals( Ast.builtinTypeSymbol ) then
            local txtList = (_lune.unwrap( macroVal.val) )
            local newToken = ""
            for __index, txt in pairs( txtList ) do
               newToken = string.format( "%s%s", newToken, txt)
            end
            
            nextToken = Parser.Token.new(Parser.TokenKind.Str, string.format( "'%s'", newToken), nextToken.pos, false)
            parser:pushbackToken( nextToken )
         elseif macroVal.typeInfo:equals( Ast.builtinTypeStat ) or macroVal.typeInfo:equals( Ast.builtinTypeExp ) then
            local txt = (_lune.unwrap( macroVal.val) )
            local rawTxt
            
            if txt:find( "^```" ) then
               
               rawTxt = string.format( "%q", txt)
            else
             
               rawTxt = string.format( "%q", txt)
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
               if valType:equals( Ast.builtinTypeSymbol ) or valType:equals( Ast.builtinTypeExp ) or valType:equals( Ast.builtinTypeStat ) then
                  format = "' %s '"
               elseif valType:get_kind() == Ast.TypeInfoKind.List and valType:get_itemTypeInfoList()[1]:equals( Ast.builtinTypeStat ) then
                  format = "' %s '"
                  exp = Nodes.ExpMacroStatListNode.create( nodeManager, prefixToken.pos, self.macroMode == Nodes.MacroMode.AnalyzeArg, {Ast.builtinTypeString}, exp )
               elseif Ast.builtinTypeString:equals( valType ) then
               else
                
                  table.insert( errMessList, ErrorMess.new(_lune.unwrap( symbolInfo:get_pos()), string.format( "not support ,, -- %s", valType:getTxt(  ))) )
               end
               
            else
               if exp:get_expType():equals( Ast.builtinTypeInt ) or exp:get_expType():equals( Ast.builtinTypeReal ) then
                  format = "'%s' "
               elseif exp:get_expType():equals( Ast.builtinTypeStat ) or exp:get_expType():equals( Ast.builtinTypeExp ) then
                  format = "' %s '"
               end
               
            end
            
         end
      end
      
   end
   
   local newToken = Parser.Token.new(Parser.TokenKind.Str, format, prefixToken.pos, prefixToken.consecutive)
   local literalStr = Nodes.LiteralStringNode.create( nodeManager, prefixToken.pos, self.macroMode == Nodes.MacroMode.AnalyzeArg, {Ast.builtinTypeString}, newToken, Nodes.ExpListNode.create( nodeManager, exp:get_pos(), self.macroMode == Nodes.MacroMode.AnalyzeArg, exp:get_expTypeList(), {exp}, nil, false ) )
   return literalStr
end


function MacroCtrl:registVar( symbolList )

   for __index, symbolInfo in pairs( symbolList ) do
      self.symbol2ValueMapForMacro[symbolInfo:get_name()] = Nodes.MacroValInfo.new(nil, symbolInfo:get_typeInfo(), nil)
   end
   
end


function MacroCtrl:startDecl(  )

   self.symbol2ValueMapForMacro = {}
end




function MacroCtrl:finishMacroMode(  )

   table.remove( self.macroModeStack )
   self.macroMode = self.macroModeStack[#self.macroModeStack]
end


function MacroCtrl:startExpandMode( lineNo, callback )

   self.macroMode = Nodes.MacroMode.Expand
   self.macroCallLineNo = lineNo
   table.insert( self.macroModeStack, self.macroMode )
   
   callback(  )
   
   self:finishMacroMode(  )
end


function MacroCtrl:startAnalyzeArgMode(  )

   self.macroMode = Nodes.MacroMode.AnalyzeArg
   table.insert( self.macroModeStack, self.macroMode )
end


function MacroCtrl:switchMacroMode(  )

   
   self.macroMode = self.macroModeStack[#self.macroModeStack - 1]
   table.insert( self.macroModeStack, self.macroMode )
end


function MacroCtrl:restoreMacroMode(  )

   
   table.remove( self.macroModeStack )
   self.macroMode = self.macroModeStack[#self.macroModeStack]
end


function MacroCtrl:isInAnalyzeArgMode(  )

   if #self.macroModeStack == 0 then
      return false
   end
   
   for __index, mode in pairs( self.macroModeStack ) do
      if mode == Nodes.MacroMode.AnalyzeArg then
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
