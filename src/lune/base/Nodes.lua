--lune/base/Nodes.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Nodes'
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
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )

local SimpleModuleInfoManager = {}
setmetatable( SimpleModuleInfoManager, { ifList = {Ast.ModuleInfoManager,} } )
function SimpleModuleInfoManager.new( moduleInfoManager )
   local obj = {}
   SimpleModuleInfoManager.setmeta( obj )
   if obj.__init then obj:__init( moduleInfoManager ); end
   return obj
end
function SimpleModuleInfoManager:__init(moduleInfoManager) 
   if moduleInfoManager ~= nil then
      self.moduleInfoManager = moduleInfoManager
   else
      self.moduleInfoManager = Ast.DummyModuleInfoManager.get_instance()
   end
   
   self.moduleInfoManagerHist = {}
end
function SimpleModuleInfoManager:push( moduleInfoManager )

   table.insert( self.moduleInfoManagerHist, self.moduleInfoManager )
   self.moduleInfoManager = moduleInfoManager
end
function SimpleModuleInfoManager:pop(  )

   self.moduleInfoManager = self.moduleInfoManagerHist[#self.moduleInfoManagerHist]
   table.remove( self.moduleInfoManagerHist )
end
function SimpleModuleInfoManager.setmeta( obj )
  setmetatable( obj, { __index = SimpleModuleInfoManager  } )
end
function SimpleModuleInfoManager:getModuleInfo( ... )
   return self.moduleInfoManager:getModuleInfo( ... )
end



local Filter = {}
_moduleObj.Filter = Filter
function Filter.new( errorOnDefault, moduleTypeInfo, moduleInfoManager )
   local obj = {}
   Filter.setmeta( obj )
   if obj.__init then obj:__init( errorOnDefault, moduleTypeInfo, moduleInfoManager ); end
   return obj
end
function Filter:__init(errorOnDefault, moduleTypeInfo, moduleInfoManager) 
   self.errorOnDefault = errorOnDefault
   self.moduleInfoManager = SimpleModuleInfoManager.new(moduleInfoManager)
   local function process(  )
   
      if moduleTypeInfo ~= nil then
         return Ast.TypeNameCtrl.new(moduleTypeInfo)
      end
      
      return Ast.defaultTypeNameCtrl
   end
   
   self.typeNameCtrl = process(  )
end
function Filter:get_moduleInfoManager(  )

   return self.moduleInfoManager
end
function Filter.setmeta( obj )
  setmetatable( obj, { __index = Filter  } )
end
function Filter:get_typeNameCtrl()
   return self.typeNameCtrl
end

local BreakKind = {}
_moduleObj.BreakKind = BreakKind
BreakKind._val2NameMap = {}
function BreakKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "BreakKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function BreakKind._from( val )
   if BreakKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
BreakKind.__allList = {}
function BreakKind.get__allList()
   return BreakKind.__allList
end

BreakKind.None = 0
BreakKind._val2NameMap[0] = 'None'
BreakKind.__allList[1] = BreakKind.None
BreakKind.Break = 1
BreakKind._val2NameMap[1] = 'Break'
BreakKind.__allList[2] = BreakKind.Break
BreakKind.Return = 2
BreakKind._val2NameMap[2] = 'Return'
BreakKind.__allList[3] = BreakKind.Return
BreakKind.NeverRet = 3
BreakKind._val2NameMap[3] = 'NeverRet'
BreakKind.__allList[4] = BreakKind.NeverRet


local CheckBreakMode = {}
_moduleObj.CheckBreakMode = CheckBreakMode
CheckBreakMode._val2NameMap = {}
function CheckBreakMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CheckBreakMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function CheckBreakMode._from( val )
   if CheckBreakMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
CheckBreakMode.__allList = {}
function CheckBreakMode.get__allList()
   return CheckBreakMode.__allList
end

CheckBreakMode.Normal = 0
CheckBreakMode._val2NameMap[0] = 'Normal'
CheckBreakMode.__allList[1] = CheckBreakMode.Normal
CheckBreakMode.Return = 1
CheckBreakMode._val2NameMap[1] = 'Return'
CheckBreakMode.__allList[2] = CheckBreakMode.Return
CheckBreakMode.IgnoreFlow = 2
CheckBreakMode._val2NameMap[2] = 'IgnoreFlow'
CheckBreakMode.__allList[3] = CheckBreakMode.IgnoreFlow
CheckBreakMode.IgnoreFlowReturn = 3
CheckBreakMode._val2NameMap[3] = 'IgnoreFlowReturn'
CheckBreakMode.__allList[4] = CheckBreakMode.IgnoreFlowReturn


local Literal = {}
Literal._name2Val = {}
_moduleObj.Literal = Literal
function Literal:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "Literal.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function Literal._from( val )
   return _lune._AlgeFrom( Literal, val )
end

Literal.ARRAY = { "ARRAY", {{ func=_lune._toList, nilable=false, child={ { func = Literal._from, nilable = false, child = {} } } }}}
Literal._name2Val["ARRAY"] = Literal.ARRAY
Literal.Bool = { "Bool", {{ func=_lune._toBool, nilable=false, child={} }}}
Literal._name2Val["Bool"] = Literal.Bool
Literal.Field = { "Field", {{ func=_lune._toList, nilable=false, child={ { func = _lune._toStr, nilable = false, child = {} } } }}}
Literal._name2Val["Field"] = Literal.Field
Literal.Int = { "Int", {{ func=_lune._toInt, nilable=false, child={} }}}
Literal._name2Val["Int"] = Literal.Int
Literal.LIST = { "LIST", {{ func=_lune._toList, nilable=false, child={ { func = Literal._from, nilable = false, child = {} } } }}}
Literal._name2Val["LIST"] = Literal.LIST
Literal.MAP = { "MAP", {{ func=_lune._toMap, nilable=false, child={ { func = Literal._from, nilable = false, child = {} }, 
{ func = Literal._from, nilable = false, child = {} } } }}}
Literal._name2Val["MAP"] = Literal.MAP
Literal.Nil = { "Nil"}
Literal._name2Val["Nil"] = Literal.Nil
Literal.Real = { "Real", {{ func=_lune._toReal, nilable=false, child={} }}}
Literal._name2Val["Real"] = Literal.Real
Literal.SET = { "SET", {{ func=_lune._toList, nilable=false, child={ { func = Literal._from, nilable = false, child = {} } } }}}
Literal._name2Val["SET"] = Literal.SET
Literal.Str = { "Str", {{ func=_lune._toStr, nilable=false, child={} }}}
Literal._name2Val["Str"] = Literal.Str
Literal.Symbol = { "Symbol", {{ func=_lune._toStr, nilable=false, child={} }}}
Literal._name2Val["Symbol"] = Literal.Symbol


local function getLiteralObj( obj )

   do
      local _matchExp = obj
      if _matchExp[1] == Literal.Nil[1] then
      
         return nil
      elseif _matchExp[1] == Literal.Int[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Literal.Real[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Literal.Str[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Literal.Bool[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Literal.Symbol[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Literal.Field[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Literal.LIST[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Literal.ARRAY[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Literal.SET[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == Literal.MAP[1] then
         local val = _matchExp[2][1]
      
         return val
      end
   end
   
end
_moduleObj.getLiteralObj = getLiteralObj



local NodeVisitMode = {}
_moduleObj.NodeVisitMode = NodeVisitMode
NodeVisitMode._val2NameMap = {}
function NodeVisitMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "NodeVisitMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function NodeVisitMode._from( val )
   if NodeVisitMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
NodeVisitMode.__allList = {}
function NodeVisitMode.get__allList()
   return NodeVisitMode.__allList
end

NodeVisitMode.Child = 0
NodeVisitMode._val2NameMap[0] = 'Child'
NodeVisitMode.__allList[1] = NodeVisitMode.Child
NodeVisitMode.Next = 1
NodeVisitMode._val2NameMap[1] = 'Next'
NodeVisitMode.__allList[2] = NodeVisitMode.Next
NodeVisitMode.End = 2
NodeVisitMode._val2NameMap[2] = 'End'
NodeVisitMode.__allList[3] = NodeVisitMode.End



local Node = {}
_moduleObj.Node = Node
function Node:setLValue(  )

   self.isLValue = true
end
function Node.new( id, kind, pos, macroArgFlag, expTypeList )
   local obj = {}
   Node.setmeta( obj )
   if obj.__init then obj:__init( id, kind, pos, macroArgFlag, expTypeList ); end
   return obj
end
function Node:__init(id, kind, pos, macroArgFlag, expTypeList) 
   self.isLValue = false
   self.id = id
   self.kind = kind
   self.pos = pos
   self.expTypeList = expTypeList
   self.commentList = nil
   self.tailComment = nil
   self.macroArgFlag = macroArgFlag
end
function Node:addComment( commentList )

   if #commentList ~= 0 then
      local workList
      
      do
         local _exp = self.commentList
         if _exp ~= nil then
            workList = _exp
         else
            workList = {}
            self.commentList = workList
         end
      end
      
      for __index, comment in ipairs( commentList ) do
         table.insert( workList, comment )
      end
      
   end
   
end
function Node:get_expType(  )

   if #self.expTypeList == 0 then
      return Ast.builtinTypeNone
   end
   
   return self.expTypeList[1]
end
function Node:addTokenList( list, kind, txt )

   table.insert( list, Parser.Token.new(kind, txt, self.pos, false) )
end
function Node:setupLiteralTokenList( list )

   return false
end
function Node:getLiteral(  )

   return nil, nil
end
function Node:processFilter( filter, opt )

end
function Node:canBeLeft(  )

   return false
end
function Node:canBeRight(  )

   return false
end
function Node:canBeStatement(  )

   return false
end
function Node:getBreakKind( checkMode )

   return BreakKind.None
end
function Node:hasNilAccess(  )

   return false
end
function Node.setmeta( obj )
  setmetatable( obj, { __index = Node  } )
end
function Node:get_id()
   return self.id
end
function Node:get_kind()
   return self.kind
end
function Node:get_pos()
   return self.pos
end
function Node:get_expTypeList()
   return self.expTypeList
end
function Node:get_commentList()
   return self.commentList
end
function Node:get_tailComment()
   return self.tailComment
end
function Node:set_tailComment( tailComment )
   self.tailComment = tailComment
end
function Node:get_isLValue()
   return self.isLValue
end
function Node:get_macroArgFlag()
   return self.macroArgFlag
end


local NamespaceInfo = {}
_moduleObj.NamespaceInfo = NamespaceInfo
function NamespaceInfo.setmeta( obj )
  setmetatable( obj, { __index = NamespaceInfo  } )
end
function NamespaceInfo.new( name, scope, typeInfo )
   local obj = {}
   NamespaceInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, scope, typeInfo )
   end
   return obj
end
function NamespaceInfo:__init( name, scope, typeInfo )

   self.name = name
   self.scope = scope
   self.typeInfo = typeInfo
end








local DeclMacroInfo = {}
_moduleObj.DeclMacroInfo = DeclMacroInfo
function DeclMacroInfo.setmeta( obj )
  setmetatable( obj, { __index = DeclMacroInfo  } )
end
function DeclMacroInfo.new( pubFlag, name, argList, stmtBlock, tokenList )
   local obj = {}
   DeclMacroInfo.setmeta( obj )
   if obj.__init then
      obj:__init( pubFlag, name, argList, stmtBlock, tokenList )
   end
   return obj
end
function DeclMacroInfo:__init( pubFlag, name, argList, stmtBlock, tokenList )

   self.pubFlag = pubFlag
   self.name = name
   self.argList = argList
   self.stmtBlock = stmtBlock
   self.tokenList = tokenList
end
function DeclMacroInfo:get_pubFlag()
   return self.pubFlag
end
function DeclMacroInfo:get_name()
   return self.name
end
function DeclMacroInfo:get_argList()
   return self.argList
end
function DeclMacroInfo:get_stmtBlock()
   return self.stmtBlock
end
function DeclMacroInfo:get_tokenList()
   return self.tokenList
end


local nodeKind2NameMap = {}
local nodeKindSeed = 0

local function regKind( name )

   local kind = nodeKindSeed
   nodeKindSeed = nodeKindSeed + 1
   nodeKind2NameMap[kind] = name
   return kind
end

local function getNodeKindName( kind )

   return _lune.unwrap( nodeKind2NameMap[kind])
end
_moduleObj.getNodeKindName = getNodeKindName

function Filter:defaultProcess( node, opt )

   if self.errorOnDefault then
      Util.err( string.format( "not implement yet -- %s", getNodeKindName( node:get_kind() )) )
   end
   
end

local NodeManager = {}
_moduleObj.NodeManager = NodeManager
function NodeManager.new(  )
   local obj = {}
   NodeManager.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function NodeManager:__init() 
   self.idSeed = 0
   self.nodeKind2NodeList = {}
   for kind, _2458 in pairs( nodeKind2NameMap ) do
      if not self.nodeKind2NodeList[kind] then
         self.nodeKind2NodeList[kind] = {}
      end
      
   end
   
end
function NodeManager:nextId(  )

   self.idSeed = self.idSeed + 1
   return self.idSeed
end
function NodeManager:getList( kind )

   local list = self.nodeKind2NodeList[kind]
   if  nil == list then
      local _list = list
   
      return {}
   end
   
   return list
end
function NodeManager:addNode( node )

   local list = self.nodeKind2NodeList[node:get_kind()]
   if  nil == list then
      local _list = list
   
      list = {}
      self.nodeKind2NodeList[node:get_kind()] = list
   end
   
   table.insert( list, node )
end
function NodeManager:delNode( node )

   local list = self.nodeKind2NodeList[node:get_kind()]
   if  nil == list then
      local _list = list
   
      return 
   end
   
   local findIndex = -1
   for index, item in ipairs( list ) do
      if item == node then
         findIndex = index
         break
      end
      
   end
   
   if findIndex ~= -1 then
      table.remove( list, findIndex )
   end
   
end
function NodeManager.setmeta( obj )
  setmetatable( obj, { __index = NodeManager  } )
end


local NodeKind = {}
_moduleObj.NodeKind = NodeKind
function NodeKind.setmeta( obj )
  setmetatable( obj, { __index = NodeKind  } )
end
function NodeKind.new(  )
   local obj = {}
   NodeKind.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function NodeKind:__init(  )

end





function NodeKind.get_None(  )

   return 0
end



regKind( "None" )
function Filter:processNone( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getNoneNodeList(  )

   return self:getList( 0 )
end



local NoneNode = {}
setmetatable( NoneNode, { __index = Node } )
_moduleObj.NoneNode = NoneNode
function NoneNode:processFilter( filter, opt )

   filter:processNone( self, opt )
end
function NoneNode:canBeRight(  )

   return false
end
function NoneNode:canBeLeft(  )

   return false
end
function NoneNode:canBeStatement(  )

   return true
end
function NoneNode.new( id, pos, macroArgFlag, typeList )
   local obj = {}
   NoneNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList ); end
   return obj
end
function NoneNode:__init(id, pos, macroArgFlag, typeList) 
   Node.__init( self,id, 0, pos, macroArgFlag, typeList)
   
   
   
   
end
function NoneNode.create( nodeMan, pos, macroArgFlag, typeList )

   local node = NoneNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList)
   nodeMan:addNode( node )
   return node
end
function NoneNode:visit( visitor, depth )

   
   return true
end
function NoneNode.setmeta( obj )
  setmetatable( obj, { __index = NoneNode  } )
end



function NodeKind.get_BlankLine(  )

   return 1
end



regKind( "BlankLine" )
function Filter:processBlankLine( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getBlankLineNodeList(  )

   return self:getList( 1 )
end



local BlankLineNode = {}
setmetatable( BlankLineNode, { __index = Node } )
_moduleObj.BlankLineNode = BlankLineNode
function BlankLineNode:processFilter( filter, opt )

   filter:processBlankLine( self, opt )
end
function BlankLineNode:canBeRight(  )

   return false
end
function BlankLineNode:canBeLeft(  )

   return false
end
function BlankLineNode:canBeStatement(  )

   return true
end
function BlankLineNode.new( id, pos, macroArgFlag, typeList, lineNum )
   local obj = {}
   BlankLineNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, lineNum ); end
   return obj
end
function BlankLineNode:__init(id, pos, macroArgFlag, typeList, lineNum) 
   Node.__init( self,id, 1, pos, macroArgFlag, typeList)
   
   
   
   self.lineNum = lineNum
   
   
end
function BlankLineNode.create( nodeMan, pos, macroArgFlag, typeList, lineNum )

   local node = BlankLineNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, lineNum)
   nodeMan:addNode( node )
   return node
end
function BlankLineNode:visit( visitor, depth )

   
   return true
end
function BlankLineNode.setmeta( obj )
  setmetatable( obj, { __index = BlankLineNode  } )
end
function BlankLineNode:get_lineNum()
   return self.lineNum
end



function NodeKind.get_Subfile(  )

   return 2
end



regKind( "Subfile" )
function Filter:processSubfile( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getSubfileNodeList(  )

   return self:getList( 2 )
end



local SubfileNode = {}
setmetatable( SubfileNode, { __index = Node } )
_moduleObj.SubfileNode = SubfileNode
function SubfileNode:processFilter( filter, opt )

   filter:processSubfile( self, opt )
end
function SubfileNode:canBeRight(  )

   return false
end
function SubfileNode:canBeLeft(  )

   return false
end
function SubfileNode:canBeStatement(  )

   return true
end
function SubfileNode.new( id, pos, macroArgFlag, typeList, usePath )
   local obj = {}
   SubfileNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, usePath ); end
   return obj
end
function SubfileNode:__init(id, pos, macroArgFlag, typeList, usePath) 
   Node.__init( self,id, 2, pos, macroArgFlag, typeList)
   
   
   
   self.usePath = usePath
   
   
end
function SubfileNode.create( nodeMan, pos, macroArgFlag, typeList, usePath )

   local node = SubfileNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, usePath)
   nodeMan:addNode( node )
   return node
end
function SubfileNode:visit( visitor, depth )

   
   return true
end
function SubfileNode.setmeta( obj )
  setmetatable( obj, { __index = SubfileNode  } )
end
function SubfileNode:get_usePath()
   return self.usePath
end



function NodeKind.get_Import(  )

   return 3
end



regKind( "Import" )
function Filter:processImport( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getImportNodeList(  )

   return self:getList( 3 )
end



local ImportNode = {}
setmetatable( ImportNode, { __index = Node } )
_moduleObj.ImportNode = ImportNode
function ImportNode:processFilter( filter, opt )

   filter:processImport( self, opt )
end
function ImportNode:canBeRight(  )

   return false
end
function ImportNode:canBeLeft(  )

   return false
end
function ImportNode:canBeStatement(  )

   return true
end
function ImportNode.new( id, pos, macroArgFlag, typeList, modulePath, assignName, symbolInfo, moduleTypeInfo )
   local obj = {}
   ImportNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, modulePath, assignName, symbolInfo, moduleTypeInfo ); end
   return obj
end
function ImportNode:__init(id, pos, macroArgFlag, typeList, modulePath, assignName, symbolInfo, moduleTypeInfo) 
   Node.__init( self,id, 3, pos, macroArgFlag, typeList)
   
   
   
   self.modulePath = modulePath
   self.assignName = assignName
   self.symbolInfo = symbolInfo
   self.moduleTypeInfo = moduleTypeInfo
   
   
end
function ImportNode.create( nodeMan, pos, macroArgFlag, typeList, modulePath, assignName, symbolInfo, moduleTypeInfo )

   local node = ImportNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, modulePath, assignName, symbolInfo, moduleTypeInfo)
   nodeMan:addNode( node )
   return node
end
function ImportNode:visit( visitor, depth )

   
   return true
end
function ImportNode.setmeta( obj )
  setmetatable( obj, { __index = ImportNode  } )
end
function ImportNode:get_modulePath()
   return self.modulePath
end
function ImportNode:get_assignName()
   return self.assignName
end
function ImportNode:get_symbolInfo()
   return self.symbolInfo
end
function ImportNode:get_moduleTypeInfo()
   return self.moduleTypeInfo
end





local LuneHelperInfo = {}
_moduleObj.LuneHelperInfo = LuneHelperInfo
function LuneHelperInfo.new(  )
   local obj = {}
   LuneHelperInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function LuneHelperInfo:__init() 
   self.useNilAccess = false
   self.useUnwrapExp = false
   self.hasMappingClassDef = false
   self.useLoad = false
   self.useUnpack = false
   self.useAlge = false
   self.useSet = false
   self.callAnonymous = false
   self.pragmaSet = {}
end
function LuneHelperInfo.setmeta( obj )
  setmetatable( obj, { __index = LuneHelperInfo  } )
end


local ModuleInfo = {}
setmetatable( ModuleInfo, { ifList = {Ast.ModuleInfoIF,} } )
_moduleObj.ModuleInfo = ModuleInfo
function ModuleInfo.new( fullName, assignName, idMap, moduleId )
   local obj = {}
   ModuleInfo.setmeta( obj )
   if obj.__init then obj:__init( fullName, assignName, idMap, moduleId ); end
   return obj
end
function ModuleInfo:__init(fullName, assignName, idMap, moduleId) 
   self.moduleId = moduleId
   self.fullName = fullName
   self.assignName = assignName
   self.localTypeInfo2importIdMap = idMap
   self.importId2localTypeInfoMap = {}
   for typeInfo, importId in pairs( idMap ) do
      self.importId2localTypeInfoMap[importId] = typeInfo
   end
   
end
function ModuleInfo:get_modulePath(  )

   return self.fullName
end
function ModuleInfo:assign( assignName )

   return ModuleInfo.new(self.fullName, assignName, self.localTypeInfo2importIdMap, self.moduleId)
end
function ModuleInfo.setmeta( obj )
  setmetatable( obj, { __index = ModuleInfo  } )
end
function ModuleInfo:get_fullName()
   return self.fullName
end
function ModuleInfo:get_localTypeInfo2importIdMap()
   return self.localTypeInfo2importIdMap
end
function ModuleInfo:get_importId2localTypeInfoMap()
   return self.importId2localTypeInfoMap
end
function ModuleInfo:get_assignName()
   return self.assignName
end
function ModuleInfo:get_moduleId()
   return self.moduleId
end

local MacroValInfo = {}
_moduleObj.MacroValInfo = MacroValInfo
function MacroValInfo.setmeta( obj )
  setmetatable( obj, { __index = MacroValInfo  } )
end
function MacroValInfo.new( val, typeInfo, argNode )
   local obj = {}
   MacroValInfo.setmeta( obj )
   if obj.__init then
      obj:__init( val, typeInfo, argNode )
   end
   return obj
end
function MacroValInfo:__init( val, typeInfo, argNode )

   self.val = val
   self.typeInfo = typeInfo
   self.argNode = argNode
end


local MacroArgInfo = {}
_moduleObj.MacroArgInfo = MacroArgInfo
function MacroArgInfo.setmeta( obj )
  setmetatable( obj, { __index = MacroArgInfo  } )
end
function MacroArgInfo.new( name, typeInfo )
   local obj = {}
   MacroArgInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, typeInfo )
   end
   return obj
end
function MacroArgInfo:__init( name, typeInfo )

   self.name = name
   self.typeInfo = typeInfo
end
function MacroArgInfo:get_name()
   return self.name
end
function MacroArgInfo:get_typeInfo()
   return self.typeInfo
end



local MacroInfo = {}
_moduleObj.MacroInfo = MacroInfo
function MacroInfo.setmeta( obj )
  setmetatable( obj, { __index = MacroInfo  } )
end
function MacroInfo.new( func, symbol2MacroValInfoMap )
   local obj = {}
   MacroInfo.setmeta( obj )
   if obj.__init then
      obj:__init( func, symbol2MacroValInfoMap )
   end
   return obj
end
function MacroInfo:__init( func, symbol2MacroValInfoMap )

   self.func = func
   self.symbol2MacroValInfoMap = symbol2MacroValInfoMap
end



function NodeKind.get_Root(  )

   return 4
end



regKind( "Root" )
function Filter:processRoot( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getRootNodeList(  )

   return self:getList( 4 )
end



local RootNode = {}
setmetatable( RootNode, { __index = Node } )
_moduleObj.RootNode = RootNode
function RootNode:processFilter( filter, opt )

   filter:processRoot( self, opt )
end
function RootNode:canBeRight(  )

   return false
end
function RootNode:canBeLeft(  )

   return false
end
function RootNode:canBeStatement(  )

   return false
end
function RootNode.new( id, pos, macroArgFlag, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap )
   local obj = {}
   RootNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap ); end
   return obj
end
function RootNode:__init(id, pos, macroArgFlag, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap) 
   Node.__init( self,id, 4, pos, macroArgFlag, typeList)
   
   
   
   self.children = children
   self.moduleScope = moduleScope
   self.useModuleMacroSet = useModuleMacroSet
   self.moduleId = moduleId
   self.processInfo = processInfo
   self.moduleTypeInfo = moduleTypeInfo
   self.provideNode = provideNode
   self.luneHelperInfo = luneHelperInfo
   self.nodeManager = nodeManager
   self.importModule2moduleInfo = importModule2moduleInfo
   self.typeId2MacroInfo = typeId2MacroInfo
   self.typeId2ClassMap = typeId2ClassMap
   
   
end
function RootNode.create( nodeMan, pos, macroArgFlag, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap )

   local node = RootNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap)
   nodeMan:addNode( node )
   return node
end
function RootNode:visit( visitor, depth )

   do
      local list = self.children
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'children', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   do
      do
         local child = self.provideNode
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'provideNode', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function RootNode.setmeta( obj )
  setmetatable( obj, { __index = RootNode  } )
end
function RootNode:get_children()
   return self.children
end
function RootNode:get_moduleScope()
   return self.moduleScope
end
function RootNode:get_useModuleMacroSet()
   return self.useModuleMacroSet
end
function RootNode:get_moduleId()
   return self.moduleId
end
function RootNode:get_processInfo()
   return self.processInfo
end
function RootNode:get_moduleTypeInfo()
   return self.moduleTypeInfo
end
function RootNode:get_provideNode()
   return self.provideNode
end
function RootNode:get_luneHelperInfo()
   return self.luneHelperInfo
end
function RootNode:get_nodeManager()
   return self.nodeManager
end
function RootNode:get_importModule2moduleInfo()
   return self.importModule2moduleInfo
end
function RootNode:get_typeId2MacroInfo()
   return self.typeId2MacroInfo
end
function RootNode:get_typeId2ClassMap()
   return self.typeId2ClassMap
end



function RootNode:set_provide( node )

   self.provideNode = node
end



function NodeKind.get_RefType(  )

   return 5
end



regKind( "RefType" )
function Filter:processRefType( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getRefTypeNodeList(  )

   return self:getList( 5 )
end



local RefTypeNode = {}
setmetatable( RefTypeNode, { __index = Node } )
_moduleObj.RefTypeNode = RefTypeNode
function RefTypeNode:processFilter( filter, opt )

   filter:processRefType( self, opt )
end
function RefTypeNode:canBeRight(  )

   return false
end
function RefTypeNode:canBeLeft(  )

   return false
end
function RefTypeNode:canBeStatement(  )

   return false
end
function RefTypeNode.new( id, pos, macroArgFlag, typeList, name, itemNodeList, refFlag, mutFlag, array )
   local obj = {}
   RefTypeNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, name, itemNodeList, refFlag, mutFlag, array ); end
   return obj
end
function RefTypeNode:__init(id, pos, macroArgFlag, typeList, name, itemNodeList, refFlag, mutFlag, array) 
   Node.__init( self,id, 5, pos, macroArgFlag, typeList)
   
   
   
   self.name = name
   self.itemNodeList = itemNodeList
   self.refFlag = refFlag
   self.mutFlag = mutFlag
   self.array = array
   
   
end
function RefTypeNode.create( nodeMan, pos, macroArgFlag, typeList, name, itemNodeList, refFlag, mutFlag, array )

   local node = RefTypeNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, name, itemNodeList, refFlag, mutFlag, array)
   nodeMan:addNode( node )
   return node
end
function RefTypeNode:visit( visitor, depth )

   do
      local child = self.name
      do
         local _switchExp = visitor( child, self, 'name', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local list = self.itemNodeList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'itemNodeList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   
   
   return true
end
function RefTypeNode.setmeta( obj )
  setmetatable( obj, { __index = RefTypeNode  } )
end
function RefTypeNode:get_name()
   return self.name
end
function RefTypeNode:get_itemNodeList()
   return self.itemNodeList
end
function RefTypeNode:get_refFlag()
   return self.refFlag
end
function RefTypeNode:get_mutFlag()
   return self.mutFlag
end
function RefTypeNode:get_array()
   return self.array
end



local BlockKind = {}
_moduleObj.BlockKind = BlockKind
BlockKind._val2NameMap = {}
function BlockKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "BlockKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function BlockKind._from( val )
   if BlockKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
BlockKind.__allList = {}
function BlockKind.get__allList()
   return BlockKind.__allList
end

BlockKind.If = 0
BlockKind._val2NameMap[0] = 'If'
BlockKind.__allList[1] = BlockKind.If
BlockKind.Elseif = 1
BlockKind._val2NameMap[1] = 'Elseif'
BlockKind.__allList[2] = BlockKind.Elseif
BlockKind.Else = 2
BlockKind._val2NameMap[2] = 'Else'
BlockKind.__allList[3] = BlockKind.Else
BlockKind.While = 3
BlockKind._val2NameMap[3] = 'While'
BlockKind.__allList[4] = BlockKind.While
BlockKind.Switch = 4
BlockKind._val2NameMap[4] = 'Switch'
BlockKind.__allList[5] = BlockKind.Switch
BlockKind.Match = 5
BlockKind._val2NameMap[5] = 'Match'
BlockKind.__allList[6] = BlockKind.Match
BlockKind.Repeat = 6
BlockKind._val2NameMap[6] = 'Repeat'
BlockKind.__allList[7] = BlockKind.Repeat
BlockKind.For = 7
BlockKind._val2NameMap[7] = 'For'
BlockKind.__allList[8] = BlockKind.For
BlockKind.Apply = 8
BlockKind._val2NameMap[8] = 'Apply'
BlockKind.__allList[9] = BlockKind.Apply
BlockKind.Foreach = 9
BlockKind._val2NameMap[9] = 'Foreach'
BlockKind.__allList[10] = BlockKind.Foreach
BlockKind.Macro = 14
BlockKind._val2NameMap[14] = 'Macro'
BlockKind.__allList[11] = BlockKind.Macro
BlockKind.Func = 11
BlockKind._val2NameMap[11] = 'Func'
BlockKind.__allList[12] = BlockKind.Func
BlockKind.Default = 12
BlockKind._val2NameMap[12] = 'Default'
BlockKind.__allList[13] = BlockKind.Default
BlockKind.Block = 13
BlockKind._val2NameMap[13] = 'Block'
BlockKind.__allList[14] = BlockKind.Block
BlockKind.Macro = 14
BlockKind._val2NameMap[14] = 'Macro'
BlockKind.__allList[15] = BlockKind.Macro
BlockKind.LetUnwrap = 15
BlockKind._val2NameMap[15] = 'LetUnwrap'
BlockKind.__allList[16] = BlockKind.LetUnwrap
BlockKind.LetUnwrapThenDo = 16
BlockKind._val2NameMap[16] = 'LetUnwrapThenDo'
BlockKind.__allList[17] = BlockKind.LetUnwrapThenDo
BlockKind.IfUnwrap = 17
BlockKind._val2NameMap[17] = 'IfUnwrap'
BlockKind.__allList[18] = BlockKind.IfUnwrap
BlockKind.When = 18
BlockKind._val2NameMap[18] = 'When'
BlockKind.__allList[19] = BlockKind.When
BlockKind.Test = 19
BlockKind._val2NameMap[19] = 'Test'
BlockKind.__allList[20] = BlockKind.Test



function NodeKind.get_Block(  )

   return 6
end



regKind( "Block" )

function NodeManager:getBlockNodeList(  )

   return self:getList( 6 )
end



local BlockNode = {}
setmetatable( BlockNode, { __index = Node } )
_moduleObj.BlockNode = BlockNode
function BlockNode:processFilter( filter, opt )

   filter:processBlock( self, opt )
end
function BlockNode:canBeRight(  )

   return false
end
function BlockNode:canBeLeft(  )

   return false
end
function BlockNode:canBeStatement(  )

   return true
end
function BlockNode.new( id, pos, macroArgFlag, typeList, blockKind, scope, stmtList )
   local obj = {}
   BlockNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, blockKind, scope, stmtList ); end
   return obj
end
function BlockNode:__init(id, pos, macroArgFlag, typeList, blockKind, scope, stmtList) 
   Node.__init( self,id, 6, pos, macroArgFlag, typeList)
   
   
   
   self.blockKind = blockKind
   self.scope = scope
   self.stmtList = stmtList
   
   
end
function BlockNode.create( nodeMan, pos, macroArgFlag, typeList, blockKind, scope, stmtList )

   local node = BlockNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, blockKind, scope, stmtList)
   nodeMan:addNode( node )
   return node
end
function BlockNode:visit( visitor, depth )

   do
      local list = self.stmtList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'stmtList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   
   
   return true
end
function BlockNode.setmeta( obj )
  setmetatable( obj, { __index = BlockNode  } )
end
function BlockNode:get_blockKind()
   return self.blockKind
end
function BlockNode:get_scope()
   return self.scope
end
function BlockNode:get_stmtList()
   return self.stmtList
end



local ScopeKind = {}
_moduleObj.ScopeKind = ScopeKind
ScopeKind._val2NameMap = {}
function ScopeKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ScopeKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ScopeKind._from( val )
   if ScopeKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ScopeKind.__allList = {}
function ScopeKind.get__allList()
   return ScopeKind.__allList
end

ScopeKind.Root = 0
ScopeKind._val2NameMap[0] = 'Root'
ScopeKind.__allList[1] = ScopeKind.Root


function NodeKind.get_Scope(  )

   return 7
end



regKind( "Scope" )
function Filter:processScope( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getScopeNodeList(  )

   return self:getList( 7 )
end



local ScopeNode = {}
setmetatable( ScopeNode, { __index = Node } )
_moduleObj.ScopeNode = ScopeNode
function ScopeNode:processFilter( filter, opt )

   filter:processScope( self, opt )
end
function ScopeNode:canBeRight(  )

   return false
end
function ScopeNode:canBeLeft(  )

   return false
end
function ScopeNode:canBeStatement(  )

   return true
end
function ScopeNode.new( id, pos, macroArgFlag, typeList, scopeKind, scope, symbolList, block )
   local obj = {}
   ScopeNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, scopeKind, scope, symbolList, block ); end
   return obj
end
function ScopeNode:__init(id, pos, macroArgFlag, typeList, scopeKind, scope, symbolList, block) 
   Node.__init( self,id, 7, pos, macroArgFlag, typeList)
   
   
   
   self.scopeKind = scopeKind
   self.scope = scope
   self.symbolList = symbolList
   self.block = block
   
   
end
function ScopeNode.create( nodeMan, pos, macroArgFlag, typeList, scopeKind, scope, symbolList, block )

   local node = ScopeNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, scopeKind, scope, symbolList, block)
   nodeMan:addNode( node )
   return node
end
function ScopeNode:visit( visitor, depth )

   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ScopeNode.setmeta( obj )
  setmetatable( obj, { __index = ScopeNode  } )
end
function ScopeNode:get_scopeKind()
   return self.scopeKind
end
function ScopeNode:get_scope()
   return self.scope
end
function ScopeNode:get_symbolList()
   return self.symbolList
end
function ScopeNode:get_block()
   return self.block
end





local function getBreakKindForStmtList( checkMode, stmtList )

   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      local kind = BreakKind.None
      for __index, stmt in ipairs( stmtList ) do
         if stmt:get_kind() ~= NodeKind.get_BlankLine() then
            local work = stmt:getBreakKind( checkMode )
            
            if checkMode == CheckBreakMode.IgnoreFlowReturn then
               if work == BreakKind.Return then
                  return BreakKind.Return
               end
               
               if work == BreakKind.NeverRet then
                  return BreakKind.NeverRet
               end
               
            else
             
               do
                  local _switchExp = work
                  if _switchExp == BreakKind.None then
                     if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                        if false then
                           return BreakKind.None
                        end
                        
                     end
                     
                  else 
                     
                        if kind == BreakKind.None or kind > work then
                           kind = work
                        end
                        
                  end
               end
               
            end
            
            
         end
         
      end
      
      return kind
   end
   
   
   if #stmtList > 0 then
      for index = #stmtList, 1, -1 do
         local stmt = stmtList[index]
         if stmt:get_kind() ~= NodeKind.get_BlankLine() then
            return stmt:getBreakKind( checkMode )
         end
         
      end
      
   end
   
   return BreakKind.None
end

function BlockNode:getBreakKind( checkMode )

   return getBreakKindForStmtList( checkMode, self.stmtList )
end


local IfKind = {}
_moduleObj.IfKind = IfKind
IfKind._val2NameMap = {}
function IfKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "IfKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function IfKind._from( val )
   if IfKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
IfKind.__allList = {}
function IfKind.get__allList()
   return IfKind.__allList
end

IfKind.If = 0
IfKind._val2NameMap[0] = 'If'
IfKind.__allList[1] = IfKind.If
IfKind.ElseIf = 1
IfKind._val2NameMap[1] = 'ElseIf'
IfKind.__allList[2] = IfKind.ElseIf
IfKind.Else = 2
IfKind._val2NameMap[2] = 'Else'
IfKind.__allList[3] = IfKind.Else


local IfStmtInfo = {}
_moduleObj.IfStmtInfo = IfStmtInfo
function IfStmtInfo.setmeta( obj )
  setmetatable( obj, { __index = IfStmtInfo  } )
end
function IfStmtInfo.new( kind, exp, block )
   local obj = {}
   IfStmtInfo.setmeta( obj )
   if obj.__init then
      obj:__init( kind, exp, block )
   end
   return obj
end
function IfStmtInfo:__init( kind, exp, block )

   self.kind = kind
   self.exp = exp
   self.block = block
end
function IfStmtInfo:get_kind()
   return self.kind
end
function IfStmtInfo:get_exp()
   return self.exp
end
function IfStmtInfo:get_block()
   return self.block
end


function NodeKind.get_If(  )

   return 8
end



regKind( "If" )
function Filter:processIf( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getIfNodeList(  )

   return self:getList( 8 )
end



local IfNode = {}
setmetatable( IfNode, { __index = Node } )
_moduleObj.IfNode = IfNode
function IfNode:processFilter( filter, opt )

   filter:processIf( self, opt )
end
function IfNode:canBeRight(  )

   return false
end
function IfNode:canBeLeft(  )

   return false
end
function IfNode:canBeStatement(  )

   return true
end
function IfNode.new( id, pos, macroArgFlag, typeList, stmtList )
   local obj = {}
   IfNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, stmtList ); end
   return obj
end
function IfNode:__init(id, pos, macroArgFlag, typeList, stmtList) 
   Node.__init( self,id, 8, pos, macroArgFlag, typeList)
   
   
   
   self.stmtList = stmtList
   
   
end
function IfNode.create( nodeMan, pos, macroArgFlag, typeList, stmtList )

   local node = IfNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, stmtList)
   nodeMan:addNode( node )
   return node
end
function IfNode:visit( visitor, depth )

   
   return true
end
function IfNode.setmeta( obj )
  setmetatable( obj, { __index = IfNode  } )
end
function IfNode:get_stmtList()
   return self.stmtList
end



function IfNode:getBreakKind( checkMode )

   local hasElseFlag = false
   local kind = BreakKind.None
   for __index, stmtInfo in ipairs( self.stmtList ) do
      local work = stmtInfo:get_block():getBreakKind( checkMode )
      
      if checkMode == CheckBreakMode.IgnoreFlowReturn then
         if work == BreakKind.Return then
            return BreakKind.Return
         end
         
         if work == BreakKind.NeverRet then
            return BreakKind.NeverRet
         end
         
      else
       
         do
            local _switchExp = work
            if _switchExp == BreakKind.None then
               if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                  if true then
                     return BreakKind.None
                  end
                  
               end
               
            else 
               
                  if kind == BreakKind.None or kind > work then
                     kind = work
                  end
                  
            end
         end
         
      end
      
      
      if stmtInfo:get_kind() == IfKind.Else then
         hasElseFlag = true
      end
      
   end
   
   if hasElseFlag or (checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return ) then
      return kind
   end
   
   
   return BreakKind.None
end


local MRetExp = {}
_moduleObj.MRetExp = MRetExp
function MRetExp.setmeta( obj )
  setmetatable( obj, { __index = MRetExp  } )
end
function MRetExp.new( exp, index )
   local obj = {}
   MRetExp.setmeta( obj )
   if obj.__init then
      obj:__init( exp, index )
   end
   return obj
end
function MRetExp:__init( exp, index )

   self.exp = exp
   self.index = index
end
function MRetExp:get_exp()
   return self.exp
end
function MRetExp:get_index()
   return self.index
end



function NodeKind.get_ExpList(  )

   return 9
end



regKind( "ExpList" )
function Filter:processExpList( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpListNodeList(  )

   return self:getList( 9 )
end



local ExpListNode = {}
setmetatable( ExpListNode, { __index = Node } )
_moduleObj.ExpListNode = ExpListNode
function ExpListNode:processFilter( filter, opt )

   filter:processExpList( self, opt )
end
function ExpListNode:canBeStatement(  )

   return false
end
function ExpListNode.new( id, pos, macroArgFlag, typeList, expList, mRetExp, followOn )
   local obj = {}
   ExpListNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, expList, mRetExp, followOn ); end
   return obj
end
function ExpListNode:__init(id, pos, macroArgFlag, typeList, expList, mRetExp, followOn) 
   Node.__init( self,id, 9, pos, macroArgFlag, typeList)
   
   
   
   self.expList = expList
   self.mRetExp = mRetExp
   self.followOn = followOn
   
   
end
function ExpListNode.create( nodeMan, pos, macroArgFlag, typeList, expList, mRetExp, followOn )

   local node = ExpListNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, expList, mRetExp, followOn)
   nodeMan:addNode( node )
   return node
end
function ExpListNode:visit( visitor, depth )

   do
      local list = self.expList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'expList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   
   
   return true
end
function ExpListNode.setmeta( obj )
  setmetatable( obj, { __index = ExpListNode  } )
end
function ExpListNode:get_expList()
   return self.expList
end
function ExpListNode:get_mRetExp()
   return self.mRetExp
end
function ExpListNode:get_followOn()
   return self.followOn
end


function ExpListNode:canBeLeft(  )

   for __index, expNode in ipairs( self:get_expList() ) do
      if not expNode:canBeLeft(  ) then
         return false
      end
      
   end
   
   return true
end

function ExpListNode:canBeRight(  )

   for __index, expNode in ipairs( self:get_expList() ) do
      if not expNode:canBeRight(  ) then
         return false
      end
      
   end
   
   return true
end

function ExpListNode:setLValue(  )

   for __index, expNode in ipairs( self:get_expList() ) do
      expNode.isLValue = true
   end
   
end


local CaseInfo = {}
_moduleObj.CaseInfo = CaseInfo
function CaseInfo.setmeta( obj )
  setmetatable( obj, { __index = CaseInfo  } )
end
function CaseInfo.new( expList, block )
   local obj = {}
   CaseInfo.setmeta( obj )
   if obj.__init then
      obj:__init( expList, block )
   end
   return obj
end
function CaseInfo:__init( expList, block )

   self.expList = expList
   self.block = block
end
function CaseInfo:get_expList()
   return self.expList
end
function CaseInfo:get_block()
   return self.block
end

local CaseKind = {}
_moduleObj.CaseKind = CaseKind
CaseKind._val2NameMap = {}
function CaseKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CaseKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function CaseKind._from( val )
   if CaseKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
CaseKind.__allList = {}
function CaseKind.get__allList()
   return CaseKind.__allList
end

CaseKind.Lack = 0
CaseKind._val2NameMap[0] = 'Lack'
CaseKind.__allList[1] = CaseKind.Lack
CaseKind.Full = 1
CaseKind._val2NameMap[1] = 'Full'
CaseKind.__allList[2] = CaseKind.Full
CaseKind.MustFull = 2
CaseKind._val2NameMap[2] = 'MustFull'
CaseKind.__allList[3] = CaseKind.MustFull


function NodeKind.get_Switch(  )

   return 10
end



regKind( "Switch" )
function Filter:processSwitch( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getSwitchNodeList(  )

   return self:getList( 10 )
end



local SwitchNode = {}
setmetatable( SwitchNode, { __index = Node } )
_moduleObj.SwitchNode = SwitchNode
function SwitchNode:processFilter( filter, opt )

   filter:processSwitch( self, opt )
end
function SwitchNode:canBeRight(  )

   return false
end
function SwitchNode:canBeLeft(  )

   return false
end
function SwitchNode:canBeStatement(  )

   return true
end
function SwitchNode.new( id, pos, macroArgFlag, typeList, exp, caseList, default, caseKind, failSafeDefault )
   local obj = {}
   SwitchNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp, caseList, default, caseKind, failSafeDefault ); end
   return obj
end
function SwitchNode:__init(id, pos, macroArgFlag, typeList, exp, caseList, default, caseKind, failSafeDefault) 
   Node.__init( self,id, 10, pos, macroArgFlag, typeList)
   
   
   
   self.exp = exp
   self.caseList = caseList
   self.default = default
   self.caseKind = caseKind
   self.failSafeDefault = failSafeDefault
   
   
end
function SwitchNode.create( nodeMan, pos, macroArgFlag, typeList, exp, caseList, default, caseKind, failSafeDefault )

   local node = SwitchNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp, caseList, default, caseKind, failSafeDefault)
   nodeMan:addNode( node )
   return node
end
function SwitchNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.default
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'default', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function SwitchNode.setmeta( obj )
  setmetatable( obj, { __index = SwitchNode  } )
end
function SwitchNode:get_exp()
   return self.exp
end
function SwitchNode:get_caseList()
   return self.caseList
end
function SwitchNode:get_default()
   return self.default
end
function SwitchNode:get_caseKind()
   return self.caseKind
end
function SwitchNode:get_failSafeDefault()
   return self.failSafeDefault
end



function SwitchNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   local fullCase = self.caseKind ~= CaseKind.Lack
   for __index, caseInfo in ipairs( self.caseList ) do
      local work = caseInfo:get_block():getBreakKind( checkMode )
      local goNext = (work == BreakKind.None ) or not fullCase
      
      if checkMode == CheckBreakMode.IgnoreFlowReturn then
         if work == BreakKind.Return then
            return BreakKind.Return
         end
         
         if work == BreakKind.NeverRet then
            return BreakKind.NeverRet
         end
         
      else
       
         do
            local _switchExp = work
            if _switchExp == BreakKind.None then
               if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                  if goNext then
                     return BreakKind.None
                  end
                  
               end
               
            else 
               
                  if kind == BreakKind.None or kind > work then
                     kind = work
                  end
                  
            end
         end
         
      end
      
      
   end
   
   do
      local block = self.default
      if block ~= nil then
         local work = block:getBreakKind( checkMode )
         
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     if true then
                        return BreakKind.None
                     end
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
         return kind
      end
   end
   
   
   if fullCase then
      return kind
   end
   
   return BreakKind.None
end





function NodeKind.get_While(  )

   return 11
end



regKind( "While" )
function Filter:processWhile( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getWhileNodeList(  )

   return self:getList( 11 )
end



local WhileNode = {}
setmetatable( WhileNode, { __index = Node } )
_moduleObj.WhileNode = WhileNode
function WhileNode:processFilter( filter, opt )

   filter:processWhile( self, opt )
end
function WhileNode:canBeRight(  )

   return false
end
function WhileNode:canBeLeft(  )

   return false
end
function WhileNode:canBeStatement(  )

   return true
end
function WhileNode.new( id, pos, macroArgFlag, typeList, exp, block )
   local obj = {}
   WhileNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp, block ); end
   return obj
end
function WhileNode:__init(id, pos, macroArgFlag, typeList, exp, block) 
   Node.__init( self,id, 11, pos, macroArgFlag, typeList)
   
   
   
   self.exp = exp
   self.block = block
   
   
end
function WhileNode.create( nodeMan, pos, macroArgFlag, typeList, exp, block )

   local node = WhileNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp, block)
   nodeMan:addNode( node )
   return node
end
function WhileNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function WhileNode.setmeta( obj )
  setmetatable( obj, { __index = WhileNode  } )
end
function WhileNode:get_exp()
   return self.exp
end
function WhileNode:get_block()
   return self.block
end




function NodeKind.get_Repeat(  )

   return 12
end



regKind( "Repeat" )
function Filter:processRepeat( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getRepeatNodeList(  )

   return self:getList( 12 )
end



local RepeatNode = {}
setmetatable( RepeatNode, { __index = Node } )
_moduleObj.RepeatNode = RepeatNode
function RepeatNode:processFilter( filter, opt )

   filter:processRepeat( self, opt )
end
function RepeatNode:canBeRight(  )

   return false
end
function RepeatNode:canBeLeft(  )

   return false
end
function RepeatNode:canBeStatement(  )

   return true
end
function RepeatNode.new( id, pos, macroArgFlag, typeList, block, exp )
   local obj = {}
   RepeatNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, block, exp ); end
   return obj
end
function RepeatNode:__init(id, pos, macroArgFlag, typeList, block, exp) 
   Node.__init( self,id, 12, pos, macroArgFlag, typeList)
   
   
   
   self.block = block
   self.exp = exp
   
   
end
function RepeatNode.create( nodeMan, pos, macroArgFlag, typeList, block, exp )

   local node = RepeatNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, block, exp)
   nodeMan:addNode( node )
   return node
end
function RepeatNode:visit( visitor, depth )

   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function RepeatNode.setmeta( obj )
  setmetatable( obj, { __index = RepeatNode  } )
end
function RepeatNode:get_block()
   return self.block
end
function RepeatNode:get_exp()
   return self.exp
end




function RepeatNode:getBreakKind( checkMode )

   
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end




function NodeKind.get_For(  )

   return 13
end



regKind( "For" )
function Filter:processFor( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getForNodeList(  )

   return self:getList( 13 )
end



local ForNode = {}
setmetatable( ForNode, { __index = Node } )
_moduleObj.ForNode = ForNode
function ForNode:processFilter( filter, opt )

   filter:processFor( self, opt )
end
function ForNode:canBeRight(  )

   return false
end
function ForNode:canBeLeft(  )

   return false
end
function ForNode:canBeStatement(  )

   return true
end
function ForNode.new( id, pos, macroArgFlag, typeList, block, val, init, to, delta )
   local obj = {}
   ForNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, block, val, init, to, delta ); end
   return obj
end
function ForNode:__init(id, pos, macroArgFlag, typeList, block, val, init, to, delta) 
   Node.__init( self,id, 13, pos, macroArgFlag, typeList)
   
   
   
   self.block = block
   self.val = val
   self.init = init
   self.to = to
   self.delta = delta
   
   
end
function ForNode.create( nodeMan, pos, macroArgFlag, typeList, block, val, init, to, delta )

   local node = ForNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, block, val, init, to, delta)
   nodeMan:addNode( node )
   return node
end
function ForNode:visit( visitor, depth )

   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.init
      do
         local _switchExp = visitor( child, self, 'init', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.to
      do
         local _switchExp = visitor( child, self, 'to', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.delta
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'delta', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function ForNode.setmeta( obj )
  setmetatable( obj, { __index = ForNode  } )
end
function ForNode:get_block()
   return self.block
end
function ForNode:get_val()
   return self.val
end
function ForNode:get_init()
   return self.init
end
function ForNode:get_to()
   return self.to
end
function ForNode:get_delta()
   return self.delta
end




function ForNode:getBreakKind( checkMode )

   
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end




function NodeKind.get_Apply(  )

   return 14
end



regKind( "Apply" )
function Filter:processApply( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getApplyNodeList(  )

   return self:getList( 14 )
end



local ApplyNode = {}
setmetatable( ApplyNode, { __index = Node } )
_moduleObj.ApplyNode = ApplyNode
function ApplyNode:processFilter( filter, opt )

   filter:processApply( self, opt )
end
function ApplyNode:canBeRight(  )

   return false
end
function ApplyNode:canBeLeft(  )

   return false
end
function ApplyNode:canBeStatement(  )

   return true
end
function ApplyNode.new( id, pos, macroArgFlag, typeList, varList, expList, block )
   local obj = {}
   ApplyNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, varList, expList, block ); end
   return obj
end
function ApplyNode:__init(id, pos, macroArgFlag, typeList, varList, expList, block) 
   Node.__init( self,id, 14, pos, macroArgFlag, typeList)
   
   
   
   self.varList = varList
   self.expList = expList
   self.block = block
   
   
end
function ApplyNode.create( nodeMan, pos, macroArgFlag, typeList, varList, expList, block )

   local node = ApplyNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, varList, expList, block)
   nodeMan:addNode( node )
   return node
end
function ApplyNode:visit( visitor, depth )

   do
      local child = self.expList
      do
         local _switchExp = visitor( child, self, 'expList', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ApplyNode.setmeta( obj )
  setmetatable( obj, { __index = ApplyNode  } )
end
function ApplyNode:get_varList()
   return self.varList
end
function ApplyNode:get_expList()
   return self.expList
end
function ApplyNode:get_block()
   return self.block
end




function ApplyNode:getBreakKind( checkMode )

   
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end




function NodeKind.get_Foreach(  )

   return 15
end



regKind( "Foreach" )
function Filter:processForeach( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getForeachNodeList(  )

   return self:getList( 15 )
end



local ForeachNode = {}
setmetatable( ForeachNode, { __index = Node } )
_moduleObj.ForeachNode = ForeachNode
function ForeachNode:processFilter( filter, opt )

   filter:processForeach( self, opt )
end
function ForeachNode:canBeRight(  )

   return false
end
function ForeachNode:canBeLeft(  )

   return false
end
function ForeachNode:canBeStatement(  )

   return true
end
function ForeachNode.new( id, pos, macroArgFlag, typeList, val, key, exp, block )
   local obj = {}
   ForeachNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, val, key, exp, block ); end
   return obj
end
function ForeachNode:__init(id, pos, macroArgFlag, typeList, val, key, exp, block) 
   Node.__init( self,id, 15, pos, macroArgFlag, typeList)
   
   
   
   self.val = val
   self.key = key
   self.exp = exp
   self.block = block
   
   
end
function ForeachNode.create( nodeMan, pos, macroArgFlag, typeList, val, key, exp, block )

   local node = ForeachNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, val, key, exp, block)
   nodeMan:addNode( node )
   return node
end
function ForeachNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ForeachNode.setmeta( obj )
  setmetatable( obj, { __index = ForeachNode  } )
end
function ForeachNode:get_val()
   return self.val
end
function ForeachNode:get_key()
   return self.key
end
function ForeachNode:get_exp()
   return self.exp
end
function ForeachNode:get_block()
   return self.block
end




function ForeachNode:getBreakKind( checkMode )

   
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end




function NodeKind.get_Forsort(  )

   return 16
end



regKind( "Forsort" )
function Filter:processForsort( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getForsortNodeList(  )

   return self:getList( 16 )
end



local ForsortNode = {}
setmetatable( ForsortNode, { __index = Node } )
_moduleObj.ForsortNode = ForsortNode
function ForsortNode:processFilter( filter, opt )

   filter:processForsort( self, opt )
end
function ForsortNode:canBeRight(  )

   return false
end
function ForsortNode:canBeLeft(  )

   return false
end
function ForsortNode:canBeStatement(  )

   return true
end
function ForsortNode.new( id, pos, macroArgFlag, typeList, val, key, exp, block, sort )
   local obj = {}
   ForsortNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, val, key, exp, block, sort ); end
   return obj
end
function ForsortNode:__init(id, pos, macroArgFlag, typeList, val, key, exp, block, sort) 
   Node.__init( self,id, 16, pos, macroArgFlag, typeList)
   
   
   
   self.val = val
   self.key = key
   self.exp = exp
   self.block = block
   self.sort = sort
   
   
end
function ForsortNode.create( nodeMan, pos, macroArgFlag, typeList, val, key, exp, block, sort )

   local node = ForsortNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, val, key, exp, block, sort)
   nodeMan:addNode( node )
   return node
end
function ForsortNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ForsortNode.setmeta( obj )
  setmetatable( obj, { __index = ForsortNode  } )
end
function ForsortNode:get_val()
   return self.val
end
function ForsortNode:get_key()
   return self.key
end
function ForsortNode:get_exp()
   return self.exp
end
function ForsortNode:get_block()
   return self.block
end
function ForsortNode:get_sort()
   return self.sort
end




function ForsortNode:getBreakKind( checkMode )

   
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end




function NodeKind.get_Return(  )

   return 17
end



regKind( "Return" )
function Filter:processReturn( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getReturnNodeList(  )

   return self:getList( 17 )
end



local ReturnNode = {}
setmetatable( ReturnNode, { __index = Node } )
_moduleObj.ReturnNode = ReturnNode
function ReturnNode:processFilter( filter, opt )

   filter:processReturn( self, opt )
end
function ReturnNode:canBeRight(  )

   return false
end
function ReturnNode:canBeLeft(  )

   return false
end
function ReturnNode:canBeStatement(  )

   return true
end
function ReturnNode.new( id, pos, macroArgFlag, typeList, expList )
   local obj = {}
   ReturnNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, expList ); end
   return obj
end
function ReturnNode:__init(id, pos, macroArgFlag, typeList, expList) 
   Node.__init( self,id, 17, pos, macroArgFlag, typeList)
   
   
   
   self.expList = expList
   
   
end
function ReturnNode.create( nodeMan, pos, macroArgFlag, typeList, expList )

   local node = ReturnNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, expList)
   nodeMan:addNode( node )
   return node
end
function ReturnNode:visit( visitor, depth )

   do
      do
         local child = self.expList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'expList', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function ReturnNode.setmeta( obj )
  setmetatable( obj, { __index = ReturnNode  } )
end
function ReturnNode:get_expList()
   return self.expList
end


function ReturnNode:getBreakKind( checkMode )

   return BreakKind.Return
end



function NodeKind.get_Break(  )

   return 18
end



regKind( "Break" )
function Filter:processBreak( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getBreakNodeList(  )

   return self:getList( 18 )
end



local BreakNode = {}
setmetatable( BreakNode, { __index = Node } )
_moduleObj.BreakNode = BreakNode
function BreakNode:processFilter( filter, opt )

   filter:processBreak( self, opt )
end
function BreakNode:canBeRight(  )

   return false
end
function BreakNode:canBeLeft(  )

   return false
end
function BreakNode:canBeStatement(  )

   return true
end
function BreakNode.new( id, pos, macroArgFlag, typeList )
   local obj = {}
   BreakNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList ); end
   return obj
end
function BreakNode:__init(id, pos, macroArgFlag, typeList) 
   Node.__init( self,id, 18, pos, macroArgFlag, typeList)
   
   
   
   
end
function BreakNode.create( nodeMan, pos, macroArgFlag, typeList )

   local node = BreakNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList)
   nodeMan:addNode( node )
   return node
end
function BreakNode:visit( visitor, depth )

   
   return true
end
function BreakNode.setmeta( obj )
  setmetatable( obj, { __index = BreakNode  } )
end



function BreakNode:getBreakKind( checkMode )

   return BreakKind.Break
end



function NodeKind.get_Provide(  )

   return 19
end



regKind( "Provide" )
function Filter:processProvide( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getProvideNodeList(  )

   return self:getList( 19 )
end



local ProvideNode = {}
setmetatable( ProvideNode, { __index = Node } )
_moduleObj.ProvideNode = ProvideNode
function ProvideNode:processFilter( filter, opt )

   filter:processProvide( self, opt )
end
function ProvideNode:canBeRight(  )

   return false
end
function ProvideNode:canBeLeft(  )

   return false
end
function ProvideNode:canBeStatement(  )

   return true
end
function ProvideNode.new( id, pos, macroArgFlag, typeList, symbol )
   local obj = {}
   ProvideNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, symbol ); end
   return obj
end
function ProvideNode:__init(id, pos, macroArgFlag, typeList, symbol) 
   Node.__init( self,id, 19, pos, macroArgFlag, typeList)
   
   
   
   self.symbol = symbol
   
   
end
function ProvideNode.create( nodeMan, pos, macroArgFlag, typeList, symbol )

   local node = ProvideNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, symbol)
   nodeMan:addNode( node )
   return node
end
function ProvideNode:visit( visitor, depth )

   
   return true
end
function ProvideNode.setmeta( obj )
  setmetatable( obj, { __index = ProvideNode  } )
end
function ProvideNode:get_symbol()
   return self.symbol
end




function NodeKind.get_ExpNew(  )

   return 20
end



regKind( "ExpNew" )
function Filter:processExpNew( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpNewNodeList(  )

   return self:getList( 20 )
end



local ExpNewNode = {}
setmetatable( ExpNewNode, { __index = Node } )
_moduleObj.ExpNewNode = ExpNewNode
function ExpNewNode:processFilter( filter, opt )

   filter:processExpNew( self, opt )
end
function ExpNewNode:canBeRight(  )

   return true
end
function ExpNewNode:canBeLeft(  )

   return false
end
function ExpNewNode:canBeStatement(  )

   return true
end
function ExpNewNode.new( id, pos, macroArgFlag, typeList, symbol, ctorTypeInfo, argList )
   local obj = {}
   ExpNewNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, symbol, ctorTypeInfo, argList ); end
   return obj
end
function ExpNewNode:__init(id, pos, macroArgFlag, typeList, symbol, ctorTypeInfo, argList) 
   Node.__init( self,id, 20, pos, macroArgFlag, typeList)
   
   
   
   self.symbol = symbol
   self.ctorTypeInfo = ctorTypeInfo
   self.argList = argList
   
   
end
function ExpNewNode.create( nodeMan, pos, macroArgFlag, typeList, symbol, ctorTypeInfo, argList )

   local node = ExpNewNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, symbol, ctorTypeInfo, argList)
   nodeMan:addNode( node )
   return node
end
function ExpNewNode:visit( visitor, depth )

   do
      local child = self.symbol
      do
         local _switchExp = visitor( child, self, 'symbol', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.argList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'argList', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function ExpNewNode.setmeta( obj )
  setmetatable( obj, { __index = ExpNewNode  } )
end
function ExpNewNode:get_symbol()
   return self.symbol
end
function ExpNewNode:get_ctorTypeInfo()
   return self.ctorTypeInfo
end
function ExpNewNode:get_argList()
   return self.argList
end




function NodeKind.get_ExpUnwrap(  )

   return 21
end



regKind( "ExpUnwrap" )
function Filter:processExpUnwrap( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpUnwrapNodeList(  )

   return self:getList( 21 )
end



local ExpUnwrapNode = {}
setmetatable( ExpUnwrapNode, { __index = Node } )
_moduleObj.ExpUnwrapNode = ExpUnwrapNode
function ExpUnwrapNode:processFilter( filter, opt )

   filter:processExpUnwrap( self, opt )
end
function ExpUnwrapNode:canBeRight(  )

   return true
end
function ExpUnwrapNode:canBeLeft(  )

   return false
end
function ExpUnwrapNode:canBeStatement(  )

   return false
end
function ExpUnwrapNode.new( id, pos, macroArgFlag, typeList, exp, default )
   local obj = {}
   ExpUnwrapNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp, default ); end
   return obj
end
function ExpUnwrapNode:__init(id, pos, macroArgFlag, typeList, exp, default) 
   Node.__init( self,id, 21, pos, macroArgFlag, typeList)
   
   
   
   self.exp = exp
   self.default = default
   
   
end
function ExpUnwrapNode.create( nodeMan, pos, macroArgFlag, typeList, exp, default )

   local node = ExpUnwrapNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp, default)
   nodeMan:addNode( node )
   return node
end
function ExpUnwrapNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.default
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'default', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function ExpUnwrapNode.setmeta( obj )
  setmetatable( obj, { __index = ExpUnwrapNode  } )
end
function ExpUnwrapNode:get_exp()
   return self.exp
end
function ExpUnwrapNode:get_default()
   return self.default
end




function NodeKind.get_ExpRef(  )

   return 22
end



regKind( "ExpRef" )
function Filter:processExpRef( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpRefNodeList(  )

   return self:getList( 22 )
end



local ExpRefNode = {}
setmetatable( ExpRefNode, { __index = Node } )
_moduleObj.ExpRefNode = ExpRefNode
function ExpRefNode:processFilter( filter, opt )

   filter:processExpRef( self, opt )
end
function ExpRefNode:canBeStatement(  )

   return false
end
function ExpRefNode.new( id, pos, macroArgFlag, typeList, symbolInfo )
   local obj = {}
   ExpRefNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, symbolInfo ); end
   return obj
end
function ExpRefNode:__init(id, pos, macroArgFlag, typeList, symbolInfo) 
   Node.__init( self,id, 22, pos, macroArgFlag, typeList)
   
   
   
   self.symbolInfo = symbolInfo
   
   
end
function ExpRefNode.create( nodeMan, pos, macroArgFlag, typeList, symbolInfo )

   local node = ExpRefNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, symbolInfo)
   nodeMan:addNode( node )
   return node
end
function ExpRefNode:visit( visitor, depth )

   
   return true
end
function ExpRefNode.setmeta( obj )
  setmetatable( obj, { __index = ExpRefNode  } )
end
function ExpRefNode:get_symbolInfo()
   return self.symbolInfo
end



function ExpRefNode:canBeLeft(  )

   return self:get_symbolInfo():get_canBeLeft()
end


function ExpRefNode:canBeRight(  )

   return self:get_symbolInfo():get_canBeRight() and self:get_symbolInfo():get_hasValueFlag()
   
end



function NodeKind.get_ExpSetVal(  )

   return 23
end



regKind( "ExpSetVal" )
function Filter:processExpSetVal( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpSetValNodeList(  )

   return self:getList( 23 )
end



local ExpSetValNode = {}
setmetatable( ExpSetValNode, { __index = Node } )
_moduleObj.ExpSetValNode = ExpSetValNode
function ExpSetValNode:processFilter( filter, opt )

   filter:processExpSetVal( self, opt )
end
function ExpSetValNode:canBeRight(  )

   return false
end
function ExpSetValNode:canBeLeft(  )

   return false
end
function ExpSetValNode:canBeStatement(  )

   return true
end
function ExpSetValNode.new( id, pos, macroArgFlag, typeList, exp1, exp2, initSymSet )
   local obj = {}
   ExpSetValNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp1, exp2, initSymSet ); end
   return obj
end
function ExpSetValNode:__init(id, pos, macroArgFlag, typeList, exp1, exp2, initSymSet) 
   Node.__init( self,id, 23, pos, macroArgFlag, typeList)
   
   
   
   self.exp1 = exp1
   self.exp2 = exp2
   self.initSymSet = initSymSet
   
   
end
function ExpSetValNode.create( nodeMan, pos, macroArgFlag, typeList, exp1, exp2, initSymSet )

   local node = ExpSetValNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp1, exp2, initSymSet)
   nodeMan:addNode( node )
   return node
end
function ExpSetValNode:visit( visitor, depth )

   do
      local child = self.exp1
      do
         local _switchExp = visitor( child, self, 'exp1', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.exp2
      do
         local _switchExp = visitor( child, self, 'exp2', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpSetValNode.setmeta( obj )
  setmetatable( obj, { __index = ExpSetValNode  } )
end
function ExpSetValNode:get_exp1()
   return self.exp1
end
function ExpSetValNode:get_exp2()
   return self.exp2
end
function ExpSetValNode:get_initSymSet()
   return self.initSymSet
end




function NodeKind.get_ExpOp2(  )

   return 24
end



regKind( "ExpOp2" )
function Filter:processExpOp2( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpOp2NodeList(  )

   return self:getList( 24 )
end



local ExpOp2Node = {}
setmetatable( ExpOp2Node, { __index = Node } )
_moduleObj.ExpOp2Node = ExpOp2Node
function ExpOp2Node:processFilter( filter, opt )

   filter:processExpOp2( self, opt )
end
function ExpOp2Node:canBeRight(  )

   return true
end
function ExpOp2Node:canBeLeft(  )

   return false
end
function ExpOp2Node:canBeStatement(  )

   return false
end
function ExpOp2Node.new( id, pos, macroArgFlag, typeList, op, exp1, exp2 )
   local obj = {}
   ExpOp2Node.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, op, exp1, exp2 ); end
   return obj
end
function ExpOp2Node:__init(id, pos, macroArgFlag, typeList, op, exp1, exp2) 
   Node.__init( self,id, 24, pos, macroArgFlag, typeList)
   
   
   
   self.op = op
   self.exp1 = exp1
   self.exp2 = exp2
   
   
end
function ExpOp2Node.create( nodeMan, pos, macroArgFlag, typeList, op, exp1, exp2 )

   local node = ExpOp2Node.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, op, exp1, exp2)
   nodeMan:addNode( node )
   return node
end
function ExpOp2Node:visit( visitor, depth )

   do
      local child = self.exp1
      do
         local _switchExp = visitor( child, self, 'exp1', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.exp2
      do
         local _switchExp = visitor( child, self, 'exp2', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpOp2Node.setmeta( obj )
  setmetatable( obj, { __index = ExpOp2Node  } )
end
function ExpOp2Node:get_op()
   return self.op
end
function ExpOp2Node:get_exp1()
   return self.exp1
end
function ExpOp2Node:get_exp2()
   return self.exp2
end




function NodeKind.get_UnwrapSet(  )

   return 25
end



regKind( "UnwrapSet" )
function Filter:processUnwrapSet( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getUnwrapSetNodeList(  )

   return self:getList( 25 )
end



local UnwrapSetNode = {}
setmetatable( UnwrapSetNode, { __index = Node } )
_moduleObj.UnwrapSetNode = UnwrapSetNode
function UnwrapSetNode:processFilter( filter, opt )

   filter:processUnwrapSet( self, opt )
end
function UnwrapSetNode:canBeRight(  )

   return false
end
function UnwrapSetNode:canBeLeft(  )

   return false
end
function UnwrapSetNode:canBeStatement(  )

   return true
end
function UnwrapSetNode.new( id, pos, macroArgFlag, typeList, dstExpList, srcExpList, unwrapBlock )
   local obj = {}
   UnwrapSetNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, dstExpList, srcExpList, unwrapBlock ); end
   return obj
end
function UnwrapSetNode:__init(id, pos, macroArgFlag, typeList, dstExpList, srcExpList, unwrapBlock) 
   Node.__init( self,id, 25, pos, macroArgFlag, typeList)
   
   
   
   self.dstExpList = dstExpList
   self.srcExpList = srcExpList
   self.unwrapBlock = unwrapBlock
   
   
end
function UnwrapSetNode.create( nodeMan, pos, macroArgFlag, typeList, dstExpList, srcExpList, unwrapBlock )

   local node = UnwrapSetNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, dstExpList, srcExpList, unwrapBlock)
   nodeMan:addNode( node )
   return node
end
function UnwrapSetNode:visit( visitor, depth )

   do
      local child = self.dstExpList
      do
         local _switchExp = visitor( child, self, 'dstExpList', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.srcExpList
      do
         local _switchExp = visitor( child, self, 'srcExpList', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.unwrapBlock
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'unwrapBlock', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function UnwrapSetNode.setmeta( obj )
  setmetatable( obj, { __index = UnwrapSetNode  } )
end
function UnwrapSetNode:get_dstExpList()
   return self.dstExpList
end
function UnwrapSetNode:get_srcExpList()
   return self.srcExpList
end
function UnwrapSetNode:get_unwrapBlock()
   return self.unwrapBlock
end




function NodeKind.get_IfUnwrap(  )

   return 26
end



regKind( "IfUnwrap" )
function Filter:processIfUnwrap( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getIfUnwrapNodeList(  )

   return self:getList( 26 )
end



local IfUnwrapNode = {}
setmetatable( IfUnwrapNode, { __index = Node } )
_moduleObj.IfUnwrapNode = IfUnwrapNode
function IfUnwrapNode:processFilter( filter, opt )

   filter:processIfUnwrap( self, opt )
end
function IfUnwrapNode:canBeRight(  )

   return false
end
function IfUnwrapNode:canBeLeft(  )

   return false
end
function IfUnwrapNode:canBeStatement(  )

   return true
end
function IfUnwrapNode.new( id, pos, macroArgFlag, typeList, varSymList, expList, block, nilBlock )
   local obj = {}
   IfUnwrapNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, varSymList, expList, block, nilBlock ); end
   return obj
end
function IfUnwrapNode:__init(id, pos, macroArgFlag, typeList, varSymList, expList, block, nilBlock) 
   Node.__init( self,id, 26, pos, macroArgFlag, typeList)
   
   
   
   self.varSymList = varSymList
   self.expList = expList
   self.block = block
   self.nilBlock = nilBlock
   
   
end
function IfUnwrapNode.create( nodeMan, pos, macroArgFlag, typeList, varSymList, expList, block, nilBlock )

   local node = IfUnwrapNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, varSymList, expList, block, nilBlock)
   nodeMan:addNode( node )
   return node
end
function IfUnwrapNode:visit( visitor, depth )

   do
      local child = self.expList
      do
         local _switchExp = visitor( child, self, 'expList', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.nilBlock
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'nilBlock', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function IfUnwrapNode.setmeta( obj )
  setmetatable( obj, { __index = IfUnwrapNode  } )
end
function IfUnwrapNode:get_varSymList()
   return self.varSymList
end
function IfUnwrapNode:get_expList()
   return self.expList
end
function IfUnwrapNode:get_block()
   return self.block
end
function IfUnwrapNode:get_nilBlock()
   return self.nilBlock
end



function IfUnwrapNode:getBreakKind( checkMode )

   local kind = self.block:getBreakKind( checkMode )
   local work = kind
   
   if checkMode == CheckBreakMode.IgnoreFlowReturn then
      if work == BreakKind.Return then
         return BreakKind.Return
      end
      
      if work == BreakKind.NeverRet then
         return BreakKind.NeverRet
      end
      
   else
    
      do
         local _switchExp = work
         if _switchExp == BreakKind.None then
            if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
               if true then
                  return BreakKind.None
               end
               
            end
            
         else 
            
               if kind == BreakKind.None or kind > work then
                  kind = work
               end
               
         end
      end
      
   end
   
   
   do
      local block = self.nilBlock
      if block ~= nil then
         work = block:getBreakKind( checkMode )
         
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     if true then
                        return BreakKind.None
                     end
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
         return kind
      end
   end
   
   return BreakKind.None
end


local UnwrapSymbolPair = {}
_moduleObj.UnwrapSymbolPair = UnwrapSymbolPair
function UnwrapSymbolPair.setmeta( obj )
  setmetatable( obj, { __index = UnwrapSymbolPair  } )
end
function UnwrapSymbolPair.new( src, dst )
   local obj = {}
   UnwrapSymbolPair.setmeta( obj )
   if obj.__init then
      obj:__init( src, dst )
   end
   return obj
end
function UnwrapSymbolPair:__init( src, dst )

   self.src = src
   self.dst = dst
end
function UnwrapSymbolPair:get_src()
   return self.src
end
function UnwrapSymbolPair:get_dst()
   return self.dst
end



function NodeKind.get_When(  )

   return 27
end



regKind( "When" )
function Filter:processWhen( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getWhenNodeList(  )

   return self:getList( 27 )
end



local WhenNode = {}
setmetatable( WhenNode, { __index = Node } )
_moduleObj.WhenNode = WhenNode
function WhenNode:processFilter( filter, opt )

   filter:processWhen( self, opt )
end
function WhenNode:canBeRight(  )

   return false
end
function WhenNode:canBeLeft(  )

   return false
end
function WhenNode:canBeStatement(  )

   return true
end
function WhenNode.new( id, pos, macroArgFlag, typeList, symPairList, block, elseBlock )
   local obj = {}
   WhenNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, symPairList, block, elseBlock ); end
   return obj
end
function WhenNode:__init(id, pos, macroArgFlag, typeList, symPairList, block, elseBlock) 
   Node.__init( self,id, 27, pos, macroArgFlag, typeList)
   
   
   
   self.symPairList = symPairList
   self.block = block
   self.elseBlock = elseBlock
   
   
end
function WhenNode.create( nodeMan, pos, macroArgFlag, typeList, symPairList, block, elseBlock )

   local node = WhenNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, symPairList, block, elseBlock)
   nodeMan:addNode( node )
   return node
end
function WhenNode:visit( visitor, depth )

   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.elseBlock
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'elseBlock', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function WhenNode.setmeta( obj )
  setmetatable( obj, { __index = WhenNode  } )
end
function WhenNode:get_symPairList()
   return self.symPairList
end
function WhenNode:get_block()
   return self.block
end
function WhenNode:get_elseBlock()
   return self.elseBlock
end



function WhenNode:getBreakKind( checkMode )

   local kind = self.block:getBreakKind( checkMode )
   local work = kind
   
   if checkMode == CheckBreakMode.IgnoreFlowReturn then
      if work == BreakKind.Return then
         return BreakKind.Return
      end
      
      if work == BreakKind.NeverRet then
         return BreakKind.NeverRet
      end
      
   else
    
      do
         local _switchExp = work
         if _switchExp == BreakKind.None then
            if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
               if true then
                  return BreakKind.None
               end
               
            end
            
         else 
            
               if kind == BreakKind.None or kind > work then
                  kind = work
               end
               
         end
      end
      
   end
   
   
   do
      local block = self.elseBlock
      if block ~= nil then
         work = block:getBreakKind( checkMode )
         
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     if true then
                        return BreakKind.None
                     end
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
         return kind
      end
   end
   
   return BreakKind.None
end


local CastKind = {}
_moduleObj.CastKind = CastKind
CastKind._val2NameMap = {}
function CastKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CastKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function CastKind._from( val )
   if CastKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
CastKind.__allList = {}
function CastKind.get__allList()
   return CastKind.__allList
end

CastKind.Normal = 0
CastKind._val2NameMap[0] = 'Normal'
CastKind.__allList[1] = CastKind.Normal
CastKind.Force = 1
CastKind._val2NameMap[1] = 'Force'
CastKind.__allList[2] = CastKind.Force
CastKind.Implicit = 2
CastKind._val2NameMap[2] = 'Implicit'
CastKind.__allList[3] = CastKind.Implicit



function NodeKind.get_ExpCast(  )

   return 28
end



regKind( "ExpCast" )
function Filter:processExpCast( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpCastNodeList(  )

   return self:getList( 28 )
end



local ExpCastNode = {}
setmetatable( ExpCastNode, { __index = Node } )
_moduleObj.ExpCastNode = ExpCastNode
function ExpCastNode:processFilter( filter, opt )

   filter:processExpCast( self, opt )
end
function ExpCastNode:canBeRight(  )

   return true
end
function ExpCastNode:canBeLeft(  )

   return false
end
function ExpCastNode:canBeStatement(  )

   return false
end
function ExpCastNode.new( id, pos, macroArgFlag, typeList, exp, castType, castKind )
   local obj = {}
   ExpCastNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp, castType, castKind ); end
   return obj
end
function ExpCastNode:__init(id, pos, macroArgFlag, typeList, exp, castType, castKind) 
   Node.__init( self,id, 28, pos, macroArgFlag, typeList)
   
   
   
   self.exp = exp
   self.castType = castType
   self.castKind = castKind
   
   
end
function ExpCastNode.create( nodeMan, pos, macroArgFlag, typeList, exp, castType, castKind )

   local node = ExpCastNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp, castType, castKind)
   nodeMan:addNode( node )
   return node
end
function ExpCastNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpCastNode.setmeta( obj )
  setmetatable( obj, { __index = ExpCastNode  } )
end
function ExpCastNode:get_exp()
   return self.exp
end
function ExpCastNode:get_castType()
   return self.castType
end
function ExpCastNode:get_castKind()
   return self.castKind
end



function ExpCastNode:getLiteral(  )

   return self.exp:getLiteral(  )
end

function ExpCastNode:setupLiteralTokenList( list )

   return self.exp:setupLiteralTokenList( list )
end



function NodeKind.get_ExpToDDD(  )

   return 29
end



regKind( "ExpToDDD" )
function Filter:processExpToDDD( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpToDDDNodeList(  )

   return self:getList( 29 )
end



local ExpToDDDNode = {}
setmetatable( ExpToDDDNode, { __index = Node } )
_moduleObj.ExpToDDDNode = ExpToDDDNode
function ExpToDDDNode:processFilter( filter, opt )

   filter:processExpToDDD( self, opt )
end
function ExpToDDDNode:canBeRight(  )

   return true
end
function ExpToDDDNode:canBeLeft(  )

   return false
end
function ExpToDDDNode:canBeStatement(  )

   return false
end
function ExpToDDDNode.new( id, pos, macroArgFlag, typeList, expList )
   local obj = {}
   ExpToDDDNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, expList ); end
   return obj
end
function ExpToDDDNode:__init(id, pos, macroArgFlag, typeList, expList) 
   Node.__init( self,id, 29, pos, macroArgFlag, typeList)
   
   
   
   self.expList = expList
   
   
end
function ExpToDDDNode.create( nodeMan, pos, macroArgFlag, typeList, expList )

   local node = ExpToDDDNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, expList)
   nodeMan:addNode( node )
   return node
end
function ExpToDDDNode:visit( visitor, depth )

   do
      local child = self.expList
      do
         local _switchExp = visitor( child, self, 'expList', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpToDDDNode.setmeta( obj )
  setmetatable( obj, { __index = ExpToDDDNode  } )
end
function ExpToDDDNode:get_expList()
   return self.expList
end




function NodeKind.get_ExpSubDDD(  )

   return 30
end



regKind( "ExpSubDDD" )
function Filter:processExpSubDDD( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpSubDDDNodeList(  )

   return self:getList( 30 )
end



local ExpSubDDDNode = {}
setmetatable( ExpSubDDDNode, { __index = Node } )
_moduleObj.ExpSubDDDNode = ExpSubDDDNode
function ExpSubDDDNode:processFilter( filter, opt )

   filter:processExpSubDDD( self, opt )
end
function ExpSubDDDNode:canBeRight(  )

   return true
end
function ExpSubDDDNode:canBeLeft(  )

   return false
end
function ExpSubDDDNode:canBeStatement(  )

   return false
end
function ExpSubDDDNode.new( id, pos, macroArgFlag, typeList, src, remainIndex )
   local obj = {}
   ExpSubDDDNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, src, remainIndex ); end
   return obj
end
function ExpSubDDDNode:__init(id, pos, macroArgFlag, typeList, src, remainIndex) 
   Node.__init( self,id, 30, pos, macroArgFlag, typeList)
   
   
   
   self.src = src
   self.remainIndex = remainIndex
   
   
end
function ExpSubDDDNode.create( nodeMan, pos, macroArgFlag, typeList, src, remainIndex )

   local node = ExpSubDDDNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, src, remainIndex)
   nodeMan:addNode( node )
   return node
end
function ExpSubDDDNode:visit( visitor, depth )

   do
      local child = self.src
      do
         local _switchExp = visitor( child, self, 'src', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpSubDDDNode.setmeta( obj )
  setmetatable( obj, { __index = ExpSubDDDNode  } )
end
function ExpSubDDDNode:get_src()
   return self.src
end
function ExpSubDDDNode:get_remainIndex()
   return self.remainIndex
end



local MacroMode = {}
_moduleObj.MacroMode = MacroMode
MacroMode._val2NameMap = {}
function MacroMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "MacroMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function MacroMode._from( val )
   if MacroMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
MacroMode.__allList = {}
function MacroMode.get__allList()
   return MacroMode.__allList
end

MacroMode.None = 0
MacroMode._val2NameMap[0] = 'None'
MacroMode.__allList[1] = MacroMode.None
MacroMode.Expand = 1
MacroMode._val2NameMap[1] = 'Expand'
MacroMode.__allList[2] = MacroMode.Expand
MacroMode.AnalyzeArg = 2
MacroMode._val2NameMap[2] = 'AnalyzeArg'
MacroMode.__allList[3] = MacroMode.AnalyzeArg



function NodeKind.get_ExpOp1(  )

   return 31
end



regKind( "ExpOp1" )
function Filter:processExpOp1( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpOp1NodeList(  )

   return self:getList( 31 )
end



local ExpOp1Node = {}
setmetatable( ExpOp1Node, { __index = Node } )
_moduleObj.ExpOp1Node = ExpOp1Node
function ExpOp1Node:processFilter( filter, opt )

   filter:processExpOp1( self, opt )
end
function ExpOp1Node:canBeRight(  )

   return true
end
function ExpOp1Node:canBeLeft(  )

   return false
end
function ExpOp1Node:canBeStatement(  )

   return false
end
function ExpOp1Node.new( id, pos, macroArgFlag, typeList, op, macroMode, exp )
   local obj = {}
   ExpOp1Node.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, op, macroMode, exp ); end
   return obj
end
function ExpOp1Node:__init(id, pos, macroArgFlag, typeList, op, macroMode, exp) 
   Node.__init( self,id, 31, pos, macroArgFlag, typeList)
   
   
   
   self.op = op
   self.macroMode = macroMode
   self.exp = exp
   
   
end
function ExpOp1Node.create( nodeMan, pos, macroArgFlag, typeList, op, macroMode, exp )

   local node = ExpOp1Node.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, op, macroMode, exp)
   nodeMan:addNode( node )
   return node
end
function ExpOp1Node:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpOp1Node.setmeta( obj )
  setmetatable( obj, { __index = ExpOp1Node  } )
end
function ExpOp1Node:get_op()
   return self.op
end
function ExpOp1Node:get_macroMode()
   return self.macroMode
end
function ExpOp1Node:get_exp()
   return self.exp
end




function NodeKind.get_ExpRefItem(  )

   return 32
end



regKind( "ExpRefItem" )
function Filter:processExpRefItem( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpRefItemNodeList(  )

   return self:getList( 32 )
end



local ExpRefItemNode = {}
setmetatable( ExpRefItemNode, { __index = Node } )
_moduleObj.ExpRefItemNode = ExpRefItemNode
function ExpRefItemNode:hasNilAccess(  )

   return self.nilAccess
end
function ExpRefItemNode:processFilter( filter, opt )

   filter:processExpRefItem( self, opt )
end
function ExpRefItemNode:canBeRight(  )

   return true
end
function ExpRefItemNode:canBeStatement(  )

   return false
end
function ExpRefItemNode.new( id, pos, macroArgFlag, typeList, val, nilAccess, symbol, index )
   local obj = {}
   ExpRefItemNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, val, nilAccess, symbol, index ); end
   return obj
end
function ExpRefItemNode:__init(id, pos, macroArgFlag, typeList, val, nilAccess, symbol, index) 
   Node.__init( self,id, 32, pos, macroArgFlag, typeList)
   
   
   
   self.val = val
   self.nilAccess = nilAccess
   self.symbol = symbol
   self.index = index
   
   
end
function ExpRefItemNode.create( nodeMan, pos, macroArgFlag, typeList, val, nilAccess, symbol, index )

   local node = ExpRefItemNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, val, nilAccess, symbol, index)
   nodeMan:addNode( node )
   return node
end
function ExpRefItemNode:visit( visitor, depth )

   do
      local child = self.val
      do
         local _switchExp = visitor( child, self, 'val', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.index
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'index', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function ExpRefItemNode.setmeta( obj )
  setmetatable( obj, { __index = ExpRefItemNode  } )
end
function ExpRefItemNode:get_val()
   return self.val
end
function ExpRefItemNode:get_nilAccess()
   return self.nilAccess
end
function ExpRefItemNode:get_symbol()
   return self.symbol
end
function ExpRefItemNode:get_index()
   return self.index
end


function ExpRefItemNode:canBeLeft(  )

   if self.val:get_expType() == Ast.builtinTypeStem then
      return false
   end
   
   return Ast.TypeInfo.isMut( self:get_val():get_expType() ) and not self.nilAccess
end



function NodeKind.get_ExpCall(  )

   return 33
end



regKind( "ExpCall" )
function Filter:processExpCall( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpCallNodeList(  )

   return self:getList( 33 )
end



local ExpCallNode = {}
setmetatable( ExpCallNode, { __index = Node } )
_moduleObj.ExpCallNode = ExpCallNode
function ExpCallNode:hasNilAccess(  )

   return self.nilAccess
end
function ExpCallNode:processFilter( filter, opt )

   filter:processExpCall( self, opt )
end
function ExpCallNode:canBeLeft(  )

   return false
end
function ExpCallNode:canBeStatement(  )

   return true
end
function ExpCallNode.new( id, pos, macroArgFlag, typeList, func, errorFunc, nilAccess, argList )
   local obj = {}
   ExpCallNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, func, errorFunc, nilAccess, argList ); end
   return obj
end
function ExpCallNode:__init(id, pos, macroArgFlag, typeList, func, errorFunc, nilAccess, argList) 
   Node.__init( self,id, 33, pos, macroArgFlag, typeList)
   
   
   
   self.func = func
   self.errorFunc = errorFunc
   self.nilAccess = nilAccess
   self.argList = argList
   
   
end
function ExpCallNode.create( nodeMan, pos, macroArgFlag, typeList, func, errorFunc, nilAccess, argList )

   local node = ExpCallNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, func, errorFunc, nilAccess, argList)
   nodeMan:addNode( node )
   return node
end
function ExpCallNode:visit( visitor, depth )

   do
      local child = self.func
      do
         local _switchExp = visitor( child, self, 'func', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.argList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'argList', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function ExpCallNode.setmeta( obj )
  setmetatable( obj, { __index = ExpCallNode  } )
end
function ExpCallNode:get_func()
   return self.func
end
function ExpCallNode:get_errorFunc()
   return self.errorFunc
end
function ExpCallNode:get_nilAccess()
   return self.nilAccess
end
function ExpCallNode:get_argList()
   return self.argList
end


function ExpCallNode:canBeRight(  )

   local expType = self:get_expType()
   if expType:equals( Ast.builtinTypeNone ) or expType:equals( Ast.builtinTypeNeverRet ) then
      return false
   end
   
   return true
end


function ExpCallNode:getBreakKind( checkMode )

   if self.errorFunc then
      return BreakKind.NeverRet
   end
   
   return BreakKind.None
end



function NodeKind.get_ExpAccessMRet(  )

   return 34
end



regKind( "ExpAccessMRet" )
function Filter:processExpAccessMRet( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpAccessMRetNodeList(  )

   return self:getList( 34 )
end



local ExpAccessMRetNode = {}
setmetatable( ExpAccessMRetNode, { __index = Node } )
_moduleObj.ExpAccessMRetNode = ExpAccessMRetNode
function ExpAccessMRetNode:processFilter( filter, opt )

   filter:processExpAccessMRet( self, opt )
end
function ExpAccessMRetNode:canBeRight(  )

   return true
end
function ExpAccessMRetNode:canBeLeft(  )

   return false
end
function ExpAccessMRetNode:canBeStatement(  )

   return false
end
function ExpAccessMRetNode.new( id, pos, macroArgFlag, typeList, mRet, index )
   local obj = {}
   ExpAccessMRetNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, mRet, index ); end
   return obj
end
function ExpAccessMRetNode:__init(id, pos, macroArgFlag, typeList, mRet, index) 
   Node.__init( self,id, 34, pos, macroArgFlag, typeList)
   
   
   
   self.mRet = mRet
   self.index = index
   
   
end
function ExpAccessMRetNode.create( nodeMan, pos, macroArgFlag, typeList, mRet, index )

   local node = ExpAccessMRetNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, mRet, index)
   nodeMan:addNode( node )
   return node
end
function ExpAccessMRetNode:visit( visitor, depth )

   do
      local child = self.mRet
      do
         local _switchExp = visitor( child, self, 'mRet', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpAccessMRetNode.setmeta( obj )
  setmetatable( obj, { __index = ExpAccessMRetNode  } )
end
function ExpAccessMRetNode:get_mRet()
   return self.mRet
end
function ExpAccessMRetNode:get_index()
   return self.index
end




function NodeKind.get_ExpMultiTo1(  )

   return 35
end



regKind( "ExpMultiTo1" )
function Filter:processExpMultiTo1( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpMultiTo1NodeList(  )

   return self:getList( 35 )
end



local ExpMultiTo1Node = {}
setmetatable( ExpMultiTo1Node, { __index = Node } )
_moduleObj.ExpMultiTo1Node = ExpMultiTo1Node
function ExpMultiTo1Node:processFilter( filter, opt )

   filter:processExpMultiTo1( self, opt )
end
function ExpMultiTo1Node:canBeRight(  )

   return true
end
function ExpMultiTo1Node:canBeLeft(  )

   return false
end
function ExpMultiTo1Node:canBeStatement(  )

   return false
end
function ExpMultiTo1Node.new( id, pos, macroArgFlag, typeList, exp )
   local obj = {}
   ExpMultiTo1Node.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp ); end
   return obj
end
function ExpMultiTo1Node:__init(id, pos, macroArgFlag, typeList, exp) 
   Node.__init( self,id, 35, pos, macroArgFlag, typeList)
   
   
   
   self.exp = exp
   
   
end
function ExpMultiTo1Node.create( nodeMan, pos, macroArgFlag, typeList, exp )

   local node = ExpMultiTo1Node.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp)
   nodeMan:addNode( node )
   return node
end
function ExpMultiTo1Node:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpMultiTo1Node.setmeta( obj )
  setmetatable( obj, { __index = ExpMultiTo1Node  } )
end
function ExpMultiTo1Node:get_exp()
   return self.exp
end




function NodeKind.get_ExpDDD(  )

   return 36
end



regKind( "ExpDDD" )
function Filter:processExpDDD( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpDDDNodeList(  )

   return self:getList( 36 )
end



local ExpDDDNode = {}
setmetatable( ExpDDDNode, { __index = Node } )
_moduleObj.ExpDDDNode = ExpDDDNode
function ExpDDDNode:processFilter( filter, opt )

   filter:processExpDDD( self, opt )
end
function ExpDDDNode:canBeRight(  )

   return true
end
function ExpDDDNode:canBeLeft(  )

   return false
end
function ExpDDDNode:canBeStatement(  )

   return false
end
function ExpDDDNode.new( id, pos, macroArgFlag, typeList, token )
   local obj = {}
   ExpDDDNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, token ); end
   return obj
end
function ExpDDDNode:__init(id, pos, macroArgFlag, typeList, token) 
   Node.__init( self,id, 36, pos, macroArgFlag, typeList)
   
   
   
   self.token = token
   
   
end
function ExpDDDNode.create( nodeMan, pos, macroArgFlag, typeList, token )

   local node = ExpDDDNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, token)
   nodeMan:addNode( node )
   return node
end
function ExpDDDNode:visit( visitor, depth )

   
   return true
end
function ExpDDDNode.setmeta( obj )
  setmetatable( obj, { __index = ExpDDDNode  } )
end
function ExpDDDNode:get_token()
   return self.token
end




function NodeKind.get_ExpParen(  )

   return 37
end



regKind( "ExpParen" )
function Filter:processExpParen( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpParenNodeList(  )

   return self:getList( 37 )
end



local ExpParenNode = {}
setmetatable( ExpParenNode, { __index = Node } )
_moduleObj.ExpParenNode = ExpParenNode
function ExpParenNode:processFilter( filter, opt )

   filter:processExpParen( self, opt )
end
function ExpParenNode:canBeRight(  )

   return true
end
function ExpParenNode:canBeLeft(  )

   return false
end
function ExpParenNode:canBeStatement(  )

   return false
end
function ExpParenNode.new( id, pos, macroArgFlag, typeList, exp )
   local obj = {}
   ExpParenNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp ); end
   return obj
end
function ExpParenNode:__init(id, pos, macroArgFlag, typeList, exp) 
   Node.__init( self,id, 37, pos, macroArgFlag, typeList)
   
   
   
   self.exp = exp
   
   
end
function ExpParenNode.create( nodeMan, pos, macroArgFlag, typeList, exp )

   local node = ExpParenNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp)
   nodeMan:addNode( node )
   return node
end
function ExpParenNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpParenNode.setmeta( obj )
  setmetatable( obj, { __index = ExpParenNode  } )
end
function ExpParenNode:get_exp()
   return self.exp
end


function ExpParenNode:getSymbolInfo(  )

   return self.exp:getSymbolInfo(  )
end



function NodeKind.get_ExpMacroExp(  )

   return 38
end



regKind( "ExpMacroExp" )
function Filter:processExpMacroExp( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpMacroExpNodeList(  )

   return self:getList( 38 )
end



local ExpMacroExpNode = {}
setmetatable( ExpMacroExpNode, { __index = Node } )
_moduleObj.ExpMacroExpNode = ExpMacroExpNode
function ExpMacroExpNode:processFilter( filter, opt )

   filter:processExpMacroExp( self, opt )
end
function ExpMacroExpNode:canBeLeft(  )

   return false
end
function ExpMacroExpNode:canBeStatement(  )

   return true
end
function ExpMacroExpNode.new( id, pos, macroArgFlag, typeList, stmtList )
   local obj = {}
   ExpMacroExpNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, stmtList ); end
   return obj
end
function ExpMacroExpNode:__init(id, pos, macroArgFlag, typeList, stmtList) 
   Node.__init( self,id, 38, pos, macroArgFlag, typeList)
   
   
   
   self.stmtList = stmtList
   
   
end
function ExpMacroExpNode.create( nodeMan, pos, macroArgFlag, typeList, stmtList )

   local node = ExpMacroExpNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, stmtList)
   nodeMan:addNode( node )
   return node
end
function ExpMacroExpNode:visit( visitor, depth )

   do
      local list = self.stmtList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'stmtList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   
   
   return true
end
function ExpMacroExpNode.setmeta( obj )
  setmetatable( obj, { __index = ExpMacroExpNode  } )
end
function ExpMacroExpNode:get_stmtList()
   return self.stmtList
end



function ExpMacroExpNode:canBeRight(  )

   return self:get_expType() ~= Ast.builtinTypeNone
end


function ExpMacroExpNode:getBreakKind( checkMode )

   return getBreakKindForStmtList( checkMode, self.stmtList )
end


local MacroStatKind = {}
_moduleObj.MacroStatKind = MacroStatKind
MacroStatKind._val2NameMap = {}
function MacroStatKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "MacroStatKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function MacroStatKind._from( val )
   if MacroStatKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
MacroStatKind.__allList = {}
function MacroStatKind.get__allList()
   return MacroStatKind.__allList
end

MacroStatKind.Stat = 0
MacroStatKind._val2NameMap[0] = 'Stat'
MacroStatKind.__allList[1] = MacroStatKind.Stat
MacroStatKind.Exp = 1
MacroStatKind._val2NameMap[1] = 'Exp'
MacroStatKind.__allList[2] = MacroStatKind.Exp



function NodeKind.get_ExpMacroStat(  )

   return 39
end



regKind( "ExpMacroStat" )
function Filter:processExpMacroStat( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpMacroStatNodeList(  )

   return self:getList( 39 )
end



local ExpMacroStatNode = {}
setmetatable( ExpMacroStatNode, { __index = Node } )
_moduleObj.ExpMacroStatNode = ExpMacroStatNode
function ExpMacroStatNode:processFilter( filter, opt )

   filter:processExpMacroStat( self, opt )
end
function ExpMacroStatNode:canBeRight(  )

   return true
end
function ExpMacroStatNode:canBeLeft(  )

   return false
end
function ExpMacroStatNode:canBeStatement(  )

   return false
end
function ExpMacroStatNode.new( id, pos, macroArgFlag, typeList, expStrList )
   local obj = {}
   ExpMacroStatNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, expStrList ); end
   return obj
end
function ExpMacroStatNode:__init(id, pos, macroArgFlag, typeList, expStrList) 
   Node.__init( self,id, 39, pos, macroArgFlag, typeList)
   
   
   
   self.expStrList = expStrList
   
   
end
function ExpMacroStatNode.create( nodeMan, pos, macroArgFlag, typeList, expStrList )

   local node = ExpMacroStatNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, expStrList)
   nodeMan:addNode( node )
   return node
end
function ExpMacroStatNode:visit( visitor, depth )

   do
      local list = self.expStrList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'expStrList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   
   
   return true
end
function ExpMacroStatNode.setmeta( obj )
  setmetatable( obj, { __index = ExpMacroStatNode  } )
end
function ExpMacroStatNode:get_expStrList()
   return self.expStrList
end




function NodeKind.get_ExpMacroArgExp(  )

   return 40
end



regKind( "ExpMacroArgExp" )
function Filter:processExpMacroArgExp( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpMacroArgExpNodeList(  )

   return self:getList( 40 )
end



local ExpMacroArgExpNode = {}
setmetatable( ExpMacroArgExpNode, { __index = Node } )
_moduleObj.ExpMacroArgExpNode = ExpMacroArgExpNode
function ExpMacroArgExpNode:processFilter( filter, opt )

   filter:processExpMacroArgExp( self, opt )
end
function ExpMacroArgExpNode:canBeRight(  )

   return true
end
function ExpMacroArgExpNode:canBeLeft(  )

   return false
end
function ExpMacroArgExpNode:canBeStatement(  )

   return false
end
function ExpMacroArgExpNode.new( id, pos, macroArgFlag, typeList, codeTxt )
   local obj = {}
   ExpMacroArgExpNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, codeTxt ); end
   return obj
end
function ExpMacroArgExpNode:__init(id, pos, macroArgFlag, typeList, codeTxt) 
   Node.__init( self,id, 40, pos, macroArgFlag, typeList)
   
   
   
   self.codeTxt = codeTxt
   
   
end
function ExpMacroArgExpNode.create( nodeMan, pos, macroArgFlag, typeList, codeTxt )

   local node = ExpMacroArgExpNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, codeTxt)
   nodeMan:addNode( node )
   return node
end
function ExpMacroArgExpNode:visit( visitor, depth )

   
   return true
end
function ExpMacroArgExpNode.setmeta( obj )
  setmetatable( obj, { __index = ExpMacroArgExpNode  } )
end
function ExpMacroArgExpNode:get_codeTxt()
   return self.codeTxt
end




function NodeKind.get_StmtExp(  )

   return 41
end



regKind( "StmtExp" )
function Filter:processStmtExp( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getStmtExpNodeList(  )

   return self:getList( 41 )
end



local StmtExpNode = {}
setmetatable( StmtExpNode, { __index = Node } )
_moduleObj.StmtExpNode = StmtExpNode
function StmtExpNode:processFilter( filter, opt )

   filter:processStmtExp( self, opt )
end
function StmtExpNode:canBeRight(  )

   return true
end
function StmtExpNode:canBeLeft(  )

   return false
end
function StmtExpNode.new( id, pos, macroArgFlag, typeList, exp )
   local obj = {}
   StmtExpNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp ); end
   return obj
end
function StmtExpNode:__init(id, pos, macroArgFlag, typeList, exp) 
   Node.__init( self,id, 41, pos, macroArgFlag, typeList)
   
   
   
   self.exp = exp
   
   
end
function StmtExpNode.create( nodeMan, pos, macroArgFlag, typeList, exp )

   local node = StmtExpNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp)
   nodeMan:addNode( node )
   return node
end
function StmtExpNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function StmtExpNode.setmeta( obj )
  setmetatable( obj, { __index = StmtExpNode  } )
end
function StmtExpNode:get_exp()
   return self.exp
end


function StmtExpNode:canBeStatement(  )

   return self:get_exp():canBeStatement(  )
end


function StmtExpNode:getBreakKind( checkMode )

   return self:get_exp():getBreakKind( checkMode )
end



function NodeKind.get_ExpMacroStatList(  )

   return 42
end



regKind( "ExpMacroStatList" )
function Filter:processExpMacroStatList( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpMacroStatListNodeList(  )

   return self:getList( 42 )
end



local ExpMacroStatListNode = {}
setmetatable( ExpMacroStatListNode, { __index = Node } )
_moduleObj.ExpMacroStatListNode = ExpMacroStatListNode
function ExpMacroStatListNode:processFilter( filter, opt )

   filter:processExpMacroStatList( self, opt )
end
function ExpMacroStatListNode:canBeRight(  )

   return true
end
function ExpMacroStatListNode:canBeLeft(  )

   return false
end
function ExpMacroStatListNode:canBeStatement(  )

   return false
end
function ExpMacroStatListNode.new( id, pos, macroArgFlag, typeList, exp )
   local obj = {}
   ExpMacroStatListNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp ); end
   return obj
end
function ExpMacroStatListNode:__init(id, pos, macroArgFlag, typeList, exp) 
   Node.__init( self,id, 42, pos, macroArgFlag, typeList)
   
   
   
   self.exp = exp
   
   
end
function ExpMacroStatListNode.create( nodeMan, pos, macroArgFlag, typeList, exp )

   local node = ExpMacroStatListNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp)
   nodeMan:addNode( node )
   return node
end
function ExpMacroStatListNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpMacroStatListNode.setmeta( obj )
  setmetatable( obj, { __index = ExpMacroStatListNode  } )
end
function ExpMacroStatListNode:get_exp()
   return self.exp
end




function NodeKind.get_ExpOmitEnum(  )

   return 43
end



regKind( "ExpOmitEnum" )
function Filter:processExpOmitEnum( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpOmitEnumNodeList(  )

   return self:getList( 43 )
end



local ExpOmitEnumNode = {}
setmetatable( ExpOmitEnumNode, { __index = Node } )
_moduleObj.ExpOmitEnumNode = ExpOmitEnumNode
function ExpOmitEnumNode:processFilter( filter, opt )

   filter:processExpOmitEnum( self, opt )
end
function ExpOmitEnumNode:canBeRight(  )

   return true
end
function ExpOmitEnumNode:canBeLeft(  )

   return true
end
function ExpOmitEnumNode:canBeStatement(  )

   return false
end
function ExpOmitEnumNode.new( id, pos, macroArgFlag, typeList, valToken, valInfo, enumTypeInfo )
   local obj = {}
   ExpOmitEnumNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, valToken, valInfo, enumTypeInfo ); end
   return obj
end
function ExpOmitEnumNode:__init(id, pos, macroArgFlag, typeList, valToken, valInfo, enumTypeInfo) 
   Node.__init( self,id, 43, pos, macroArgFlag, typeList)
   
   
   
   self.valToken = valToken
   self.valInfo = valInfo
   self.enumTypeInfo = enumTypeInfo
   
   
end
function ExpOmitEnumNode.create( nodeMan, pos, macroArgFlag, typeList, valToken, valInfo, enumTypeInfo )

   local node = ExpOmitEnumNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, valToken, valInfo, enumTypeInfo)
   nodeMan:addNode( node )
   return node
end
function ExpOmitEnumNode:visit( visitor, depth )

   
   return true
end
function ExpOmitEnumNode.setmeta( obj )
  setmetatable( obj, { __index = ExpOmitEnumNode  } )
end
function ExpOmitEnumNode:get_valToken()
   return self.valToken
end
function ExpOmitEnumNode:get_valInfo()
   return self.valInfo
end
function ExpOmitEnumNode:get_enumTypeInfo()
   return self.enumTypeInfo
end




function NodeKind.get_RefField(  )

   return 44
end



regKind( "RefField" )
function Filter:processRefField( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getRefFieldNodeList(  )

   return self:getList( 44 )
end



local RefFieldNode = {}
setmetatable( RefFieldNode, { __index = Node } )
_moduleObj.RefFieldNode = RefFieldNode
function RefFieldNode:hasNilAccess(  )

   return self.nilAccess
end
function RefFieldNode:processFilter( filter, opt )

   filter:processRefField( self, opt )
end
function RefFieldNode:canBeStatement(  )

   return false
end
function RefFieldNode.new( id, pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix )
   local obj = {}
   RefFieldNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix ); end
   return obj
end
function RefFieldNode:__init(id, pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix) 
   Node.__init( self,id, 44, pos, macroArgFlag, typeList)
   
   
   
   self.field = field
   self.symbolInfo = symbolInfo
   self.nilAccess = nilAccess
   self.prefix = prefix
   
   
end
function RefFieldNode.create( nodeMan, pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix )

   local node = RefFieldNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix)
   nodeMan:addNode( node )
   return node
end
function RefFieldNode:visit( visitor, depth )

   do
      local child = self.prefix
      do
         local _switchExp = visitor( child, self, 'prefix', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function RefFieldNode.setmeta( obj )
  setmetatable( obj, { __index = RefFieldNode  } )
end
function RefFieldNode:get_field()
   return self.field
end
function RefFieldNode:get_symbolInfo()
   return self.symbolInfo
end
function RefFieldNode:get_nilAccess()
   return self.nilAccess
end
function RefFieldNode:get_prefix()
   return self.prefix
end


function RefFieldNode:canBeLeft(  )

   do
      local _exp = self:get_symbolInfo()
      if _exp ~= nil then
         return _exp:get_canBeLeft()
      end
   end
   
   return false
end

function RefFieldNode:canBeRight(  )

   do
      local _exp = self:get_symbolInfo()
      if _exp ~= nil then
         return _exp:get_canBeRight()
      end
   end
   
   return true
end



function NodeKind.get_GetField(  )

   return 45
end



regKind( "GetField" )
function Filter:processGetField( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getGetFieldNodeList(  )

   return self:getList( 45 )
end



local GetFieldNode = {}
setmetatable( GetFieldNode, { __index = Node } )
_moduleObj.GetFieldNode = GetFieldNode
function GetFieldNode:hasNilAccess(  )

   return self.nilAccess
end
function GetFieldNode:processFilter( filter, opt )

   filter:processGetField( self, opt )
end
function GetFieldNode:canBeRight(  )

   return true
end
function GetFieldNode:canBeStatement(  )

   return false
end
function GetFieldNode.new( id, pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo )
   local obj = {}
   GetFieldNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo ); end
   return obj
end
function GetFieldNode:__init(id, pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo) 
   Node.__init( self,id, 45, pos, macroArgFlag, typeList)
   
   
   
   self.field = field
   self.symbolInfo = symbolInfo
   self.nilAccess = nilAccess
   self.prefix = prefix
   self.getterTypeInfo = getterTypeInfo
   
   
end
function GetFieldNode.create( nodeMan, pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo )

   local node = GetFieldNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo)
   nodeMan:addNode( node )
   return node
end
function GetFieldNode:visit( visitor, depth )

   do
      local child = self.prefix
      do
         local _switchExp = visitor( child, self, 'prefix', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function GetFieldNode.setmeta( obj )
  setmetatable( obj, { __index = GetFieldNode  } )
end
function GetFieldNode:get_field()
   return self.field
end
function GetFieldNode:get_symbolInfo()
   return self.symbolInfo
end
function GetFieldNode:get_nilAccess()
   return self.nilAccess
end
function GetFieldNode:get_prefix()
   return self.prefix
end
function GetFieldNode:get_getterTypeInfo()
   return self.getterTypeInfo
end


function GetFieldNode:canBeLeft(  )

   do
      local _exp = self:get_symbolInfo()
      if _exp ~= nil then
         return _exp:get_canBeLeft()
      end
   end
   
   return false
end



function NodeKind.get_Alias(  )

   return 46
end



regKind( "Alias" )
function Filter:processAlias( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getAliasNodeList(  )

   return self:getList( 46 )
end



local AliasNode = {}
setmetatable( AliasNode, { __index = Node } )
_moduleObj.AliasNode = AliasNode
function AliasNode:processFilter( filter, opt )

   filter:processAlias( self, opt )
end
function AliasNode:canBeRight(  )

   return false
end
function AliasNode:canBeLeft(  )

   return false
end
function AliasNode:canBeStatement(  )

   return true
end
function AliasNode.new( id, pos, macroArgFlag, typeList, newName, srcNode, typeInfo )
   local obj = {}
   AliasNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, newName, srcNode, typeInfo ); end
   return obj
end
function AliasNode:__init(id, pos, macroArgFlag, typeList, newName, srcNode, typeInfo) 
   Node.__init( self,id, 46, pos, macroArgFlag, typeList)
   
   
   
   self.newName = newName
   self.srcNode = srcNode
   self.typeInfo = typeInfo
   
   
end
function AliasNode.create( nodeMan, pos, macroArgFlag, typeList, newName, srcNode, typeInfo )

   local node = AliasNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, newName, srcNode, typeInfo)
   nodeMan:addNode( node )
   return node
end
function AliasNode:visit( visitor, depth )

   do
      local child = self.srcNode
      do
         local _switchExp = visitor( child, self, 'srcNode', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function AliasNode.setmeta( obj )
  setmetatable( obj, { __index = AliasNode  } )
end
function AliasNode:get_newName()
   return self.newName
end
function AliasNode:get_srcNode()
   return self.srcNode
end
function AliasNode:get_typeInfo()
   return self.typeInfo
end



local VarInfo = {}
_moduleObj.VarInfo = VarInfo
function VarInfo.setmeta( obj )
  setmetatable( obj, { __index = VarInfo  } )
end
function VarInfo.new( name, refType, actualType )
   local obj = {}
   VarInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, refType, actualType )
   end
   return obj
end
function VarInfo:__init( name, refType, actualType )

   self.name = name
   self.refType = refType
   self.actualType = actualType
end
function VarInfo:get_name()
   return self.name
end
function VarInfo:get_refType()
   return self.refType
end
function VarInfo:get_actualType()
   return self.actualType
end


local DeclVarMode = {}
_moduleObj.DeclVarMode = DeclVarMode
DeclVarMode._val2NameMap = {}
function DeclVarMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "DeclVarMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function DeclVarMode._from( val )
   if DeclVarMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
DeclVarMode.__allList = {}
function DeclVarMode.get__allList()
   return DeclVarMode.__allList
end

DeclVarMode.Let = 0
DeclVarMode._val2NameMap[0] = 'Let'
DeclVarMode.__allList[1] = DeclVarMode.Let
DeclVarMode.Sync = 1
DeclVarMode._val2NameMap[1] = 'Sync'
DeclVarMode.__allList[2] = DeclVarMode.Sync
DeclVarMode.Unwrap = 2
DeclVarMode._val2NameMap[2] = 'Unwrap'
DeclVarMode.__allList[3] = DeclVarMode.Unwrap



function NodeKind.get_DeclVar(  )

   return 47
end



regKind( "DeclVar" )
function Filter:processDeclVar( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclVarNodeList(  )

   return self:getList( 47 )
end



local DeclVarNode = {}
setmetatable( DeclVarNode, { __index = Node } )
_moduleObj.DeclVarNode = DeclVarNode
function DeclVarNode:processFilter( filter, opt )

   filter:processDeclVar( self, opt )
end
function DeclVarNode:canBeRight(  )

   return false
end
function DeclVarNode:canBeLeft(  )

   return false
end
function DeclVarNode:canBeStatement(  )

   return true
end
function DeclVarNode.new( id, pos, macroArgFlag, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock )
   local obj = {}
   DeclVarNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock ); end
   return obj
end
function DeclVarNode:__init(id, pos, macroArgFlag, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock) 
   Node.__init( self,id, 47, pos, macroArgFlag, typeList)
   
   
   
   self.mode = mode
   self.accessMode = accessMode
   self.staticFlag = staticFlag
   self.varList = varList
   self.expList = expList
   self.symbolInfoList = symbolInfoList
   self.typeInfoList = typeInfoList
   self.unwrapFlag = unwrapFlag
   self.unwrapBlock = unwrapBlock
   self.thenBlock = thenBlock
   self.syncVarList = syncVarList
   self.syncBlock = syncBlock
   
   
end
function DeclVarNode.create( nodeMan, pos, macroArgFlag, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock )

   local node = DeclVarNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock)
   nodeMan:addNode( node )
   return node
end
function DeclVarNode:visit( visitor, depth )

   do
      do
         local child = self.expList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'expList', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   do
      do
         local child = self.unwrapBlock
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'unwrapBlock', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   do
      do
         local child = self.thenBlock
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'thenBlock', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   do
      do
         local child = self.syncBlock
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'syncBlock', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function DeclVarNode.setmeta( obj )
  setmetatable( obj, { __index = DeclVarNode  } )
end
function DeclVarNode:get_mode()
   return self.mode
end
function DeclVarNode:get_accessMode()
   return self.accessMode
end
function DeclVarNode:get_staticFlag()
   return self.staticFlag
end
function DeclVarNode:get_varList()
   return self.varList
end
function DeclVarNode:get_expList()
   return self.expList
end
function DeclVarNode:get_symbolInfoList()
   return self.symbolInfoList
end
function DeclVarNode:get_typeInfoList()
   return self.typeInfoList
end
function DeclVarNode:get_unwrapFlag()
   return self.unwrapFlag
end
function DeclVarNode:get_unwrapBlock()
   return self.unwrapBlock
end
function DeclVarNode:get_thenBlock()
   return self.thenBlock
end
function DeclVarNode:get_syncVarList()
   return self.syncVarList
end
function DeclVarNode:get_syncBlock()
   return self.syncBlock
end



function DeclVarNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   local work = BreakKind.None
   do
      local block = self.unwrapBlock
      if block ~= nil then
         work = block:getBreakKind( checkMode )
         
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     if true then
                        return BreakKind.None
                     end
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
         do
            local thenBlock = self.thenBlock
            if thenBlock ~= nil then
               work = thenBlock:getBreakKind( checkMode )
               
               if checkMode == CheckBreakMode.IgnoreFlowReturn then
                  if work == BreakKind.Return then
                     return BreakKind.Return
                  end
                  
                  if work == BreakKind.NeverRet then
                     return BreakKind.NeverRet
                  end
                  
               else
                
                  do
                     local _switchExp = work
                     if _switchExp == BreakKind.None then
                        if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                           if true then
                              return BreakKind.None
                           end
                           
                        end
                        
                     else 
                        
                           if kind == BreakKind.None or kind > work then
                              kind = work
                           end
                           
                     end
                  end
                  
               end
               
               
               do
                  local syncBlock = self.syncBlock
                  if syncBlock ~= nil then
                     work = syncBlock:getBreakKind( checkMode )
                     
                     if checkMode == CheckBreakMode.IgnoreFlowReturn then
                        if work == BreakKind.Return then
                           return BreakKind.Return
                        end
                        
                        if work == BreakKind.NeverRet then
                           return BreakKind.NeverRet
                        end
                        
                     else
                      
                        do
                           local _switchExp = work
                           if _switchExp == BreakKind.None then
                              if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                                 if true then
                                    return BreakKind.None
                                 end
                                 
                              end
                              
                           else 
                              
                                 if kind == BreakKind.None or kind > work then
                                    kind = work
                                 end
                                 
                           end
                        end
                        
                     end
                     
                     
                  end
               end
               
               return kind
            end
         end
         
         if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
            return kind
         end
         
      end
   end
   
   return BreakKind.None
end



function NodeKind.get_DeclForm(  )

   return 48
end



regKind( "DeclForm" )
function Filter:processDeclForm( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclFormNodeList(  )

   return self:getList( 48 )
end



local DeclFormNode = {}
setmetatable( DeclFormNode, { __index = Node } )
_moduleObj.DeclFormNode = DeclFormNode
function DeclFormNode:processFilter( filter, opt )

   filter:processDeclForm( self, opt )
end
function DeclFormNode:canBeRight(  )

   return false
end
function DeclFormNode:canBeLeft(  )

   return false
end
function DeclFormNode:canBeStatement(  )

   return true
end
function DeclFormNode.new( id, pos, macroArgFlag, typeList, argList )
   local obj = {}
   DeclFormNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, argList ); end
   return obj
end
function DeclFormNode:__init(id, pos, macroArgFlag, typeList, argList) 
   Node.__init( self,id, 48, pos, macroArgFlag, typeList)
   
   
   
   self.argList = argList
   
   
end
function DeclFormNode.create( nodeMan, pos, macroArgFlag, typeList, argList )

   local node = DeclFormNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, argList)
   nodeMan:addNode( node )
   return node
end
function DeclFormNode:visit( visitor, depth )

   do
      local list = self.argList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'argList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   
   
   return true
end
function DeclFormNode.setmeta( obj )
  setmetatable( obj, { __index = DeclFormNode  } )
end
function DeclFormNode:get_argList()
   return self.argList
end



local FuncKind = {}
_moduleObj.FuncKind = FuncKind
FuncKind._val2NameMap = {}
function FuncKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "FuncKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function FuncKind._from( val )
   if FuncKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
FuncKind.__allList = {}
function FuncKind.get__allList()
   return FuncKind.__allList
end

FuncKind.Func = 0
FuncKind._val2NameMap[0] = 'Func'
FuncKind.__allList[1] = FuncKind.Func
FuncKind.Mtd = 1
FuncKind._val2NameMap[1] = 'Mtd'
FuncKind.__allList[2] = FuncKind.Mtd
FuncKind.Ctor = 2
FuncKind._val2NameMap[2] = 'Ctor'
FuncKind.__allList[3] = FuncKind.Ctor
FuncKind.Dstr = 3
FuncKind._val2NameMap[3] = 'Dstr'
FuncKind.__allList[4] = FuncKind.Dstr
FuncKind.InitBlock = 4
FuncKind._val2NameMap[4] = 'InitBlock'
FuncKind.__allList[5] = FuncKind.InitBlock


local DeclFuncInfo = {}
_moduleObj.DeclFuncInfo = DeclFuncInfo
function DeclFuncInfo.setmeta( obj )
  setmetatable( obj, { __index = DeclFuncInfo  } )
end
function DeclFuncInfo.new( kind, classTypeInfo, declClassNode, name, argList, staticFlag, accessMode, body, retTypeInfoList, has__func__Symbol, overrideFlag )
   local obj = {}
   DeclFuncInfo.setmeta( obj )
   if obj.__init then
      obj:__init( kind, classTypeInfo, declClassNode, name, argList, staticFlag, accessMode, body, retTypeInfoList, has__func__Symbol, overrideFlag )
   end
   return obj
end
function DeclFuncInfo:__init( kind, classTypeInfo, declClassNode, name, argList, staticFlag, accessMode, body, retTypeInfoList, has__func__Symbol, overrideFlag )

   self.kind = kind
   self.classTypeInfo = classTypeInfo
   self.declClassNode = declClassNode
   self.name = name
   self.argList = argList
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.body = body
   self.retTypeInfoList = retTypeInfoList
   self.has__func__Symbol = has__func__Symbol
   self.overrideFlag = overrideFlag
end
function DeclFuncInfo:get_kind()
   return self.kind
end
function DeclFuncInfo:get_classTypeInfo()
   return self.classTypeInfo
end
function DeclFuncInfo:get_declClassNode()
   return self.declClassNode
end
function DeclFuncInfo:get_name()
   return self.name
end
function DeclFuncInfo:get_argList()
   return self.argList
end
function DeclFuncInfo:get_staticFlag()
   return self.staticFlag
end
function DeclFuncInfo:get_accessMode()
   return self.accessMode
end
function DeclFuncInfo:get_body()
   return self.body
end
function DeclFuncInfo:get_retTypeInfoList()
   return self.retTypeInfoList
end
function DeclFuncInfo:get_has__func__Symbol()
   return self.has__func__Symbol
end
function DeclFuncInfo:get_overrideFlag()
   return self.overrideFlag
end


function DeclFuncInfo.createFrom( info, name )

   return DeclFuncInfo.new(info:get_kind(), info.classTypeInfo, info.declClassNode, name, info.argList, info.staticFlag, info.accessMode, info.body, info.retTypeInfoList, info.has__func__Symbol, info.overrideFlag)
end



function NodeKind.get_DeclFunc(  )

   return 49
end



regKind( "DeclFunc" )
function Filter:processDeclFunc( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclFuncNodeList(  )

   return self:getList( 49 )
end



local DeclFuncNode = {}
setmetatable( DeclFuncNode, { __index = Node } )
_moduleObj.DeclFuncNode = DeclFuncNode
function DeclFuncNode:processFilter( filter, opt )

   filter:processDeclFunc( self, opt )
end
function DeclFuncNode:canBeLeft(  )

   return false
end
function DeclFuncNode.new( id, pos, macroArgFlag, typeList, declInfo )
   local obj = {}
   DeclFuncNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, declInfo ); end
   return obj
end
function DeclFuncNode:__init(id, pos, macroArgFlag, typeList, declInfo) 
   Node.__init( self,id, 49, pos, macroArgFlag, typeList)
   
   
   
   self.declInfo = declInfo
   
   
end
function DeclFuncNode.create( nodeMan, pos, macroArgFlag, typeList, declInfo )

   local node = DeclFuncNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclFuncNode:visit( visitor, depth )

   
   return true
end
function DeclFuncNode.setmeta( obj )
  setmetatable( obj, { __index = DeclFuncNode  } )
end
function DeclFuncNode:get_declInfo()
   return self.declInfo
end


function DeclFuncNode:canBeRight(  )

   return self.declInfo:get_name() == nil
end

function DeclFuncNode:canBeStatement(  )

   return not (self.declInfo:get_name() == nil )
end



function NodeKind.get_DeclMethod(  )

   return 50
end



regKind( "DeclMethod" )
function Filter:processDeclMethod( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclMethodNodeList(  )

   return self:getList( 50 )
end



local DeclMethodNode = {}
setmetatable( DeclMethodNode, { __index = Node } )
_moduleObj.DeclMethodNode = DeclMethodNode
function DeclMethodNode:processFilter( filter, opt )

   filter:processDeclMethod( self, opt )
end
function DeclMethodNode:canBeRight(  )

   return false
end
function DeclMethodNode:canBeLeft(  )

   return false
end
function DeclMethodNode:canBeStatement(  )

   return true
end
function DeclMethodNode.new( id, pos, macroArgFlag, typeList, declInfo )
   local obj = {}
   DeclMethodNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, declInfo ); end
   return obj
end
function DeclMethodNode:__init(id, pos, macroArgFlag, typeList, declInfo) 
   Node.__init( self,id, 50, pos, macroArgFlag, typeList)
   
   
   
   self.declInfo = declInfo
   
   
end
function DeclMethodNode.create( nodeMan, pos, macroArgFlag, typeList, declInfo )

   local node = DeclMethodNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclMethodNode:visit( visitor, depth )

   
   return true
end
function DeclMethodNode.setmeta( obj )
  setmetatable( obj, { __index = DeclMethodNode  } )
end
function DeclMethodNode:get_declInfo()
   return self.declInfo
end




function NodeKind.get_ProtoMethod(  )

   return 51
end



regKind( "ProtoMethod" )
function Filter:processProtoMethod( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getProtoMethodNodeList(  )

   return self:getList( 51 )
end



local ProtoMethodNode = {}
setmetatable( ProtoMethodNode, { __index = Node } )
_moduleObj.ProtoMethodNode = ProtoMethodNode
function ProtoMethodNode:processFilter( filter, opt )

   filter:processProtoMethod( self, opt )
end
function ProtoMethodNode:canBeRight(  )

   return false
end
function ProtoMethodNode:canBeLeft(  )

   return false
end
function ProtoMethodNode:canBeStatement(  )

   return true
end
function ProtoMethodNode.new( id, pos, macroArgFlag, typeList, declInfo )
   local obj = {}
   ProtoMethodNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, declInfo ); end
   return obj
end
function ProtoMethodNode:__init(id, pos, macroArgFlag, typeList, declInfo) 
   Node.__init( self,id, 51, pos, macroArgFlag, typeList)
   
   
   
   self.declInfo = declInfo
   
   
end
function ProtoMethodNode.create( nodeMan, pos, macroArgFlag, typeList, declInfo )

   local node = ProtoMethodNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function ProtoMethodNode:visit( visitor, depth )

   
   return true
end
function ProtoMethodNode.setmeta( obj )
  setmetatable( obj, { __index = ProtoMethodNode  } )
end
function ProtoMethodNode:get_declInfo()
   return self.declInfo
end




function NodeKind.get_DeclConstr(  )

   return 52
end



regKind( "DeclConstr" )
function Filter:processDeclConstr( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclConstrNodeList(  )

   return self:getList( 52 )
end



local DeclConstrNode = {}
setmetatable( DeclConstrNode, { __index = Node } )
_moduleObj.DeclConstrNode = DeclConstrNode
function DeclConstrNode:processFilter( filter, opt )

   filter:processDeclConstr( self, opt )
end
function DeclConstrNode:canBeRight(  )

   return false
end
function DeclConstrNode:canBeLeft(  )

   return false
end
function DeclConstrNode:canBeStatement(  )

   return true
end
function DeclConstrNode.new( id, pos, macroArgFlag, typeList, declInfo )
   local obj = {}
   DeclConstrNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, declInfo ); end
   return obj
end
function DeclConstrNode:__init(id, pos, macroArgFlag, typeList, declInfo) 
   Node.__init( self,id, 52, pos, macroArgFlag, typeList)
   
   
   
   self.declInfo = declInfo
   
   
end
function DeclConstrNode.create( nodeMan, pos, macroArgFlag, typeList, declInfo )

   local node = DeclConstrNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclConstrNode:visit( visitor, depth )

   
   return true
end
function DeclConstrNode.setmeta( obj )
  setmetatable( obj, { __index = DeclConstrNode  } )
end
function DeclConstrNode:get_declInfo()
   return self.declInfo
end




function NodeKind.get_DeclDestr(  )

   return 53
end



regKind( "DeclDestr" )
function Filter:processDeclDestr( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclDestrNodeList(  )

   return self:getList( 53 )
end



local DeclDestrNode = {}
setmetatable( DeclDestrNode, { __index = Node } )
_moduleObj.DeclDestrNode = DeclDestrNode
function DeclDestrNode:processFilter( filter, opt )

   filter:processDeclDestr( self, opt )
end
function DeclDestrNode:canBeRight(  )

   return false
end
function DeclDestrNode:canBeLeft(  )

   return false
end
function DeclDestrNode:canBeStatement(  )

   return true
end
function DeclDestrNode.new( id, pos, macroArgFlag, typeList, declInfo )
   local obj = {}
   DeclDestrNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, declInfo ); end
   return obj
end
function DeclDestrNode:__init(id, pos, macroArgFlag, typeList, declInfo) 
   Node.__init( self,id, 53, pos, macroArgFlag, typeList)
   
   
   
   self.declInfo = declInfo
   
   
end
function DeclDestrNode.create( nodeMan, pos, macroArgFlag, typeList, declInfo )

   local node = DeclDestrNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclDestrNode:visit( visitor, depth )

   
   return true
end
function DeclDestrNode.setmeta( obj )
  setmetatable( obj, { __index = DeclDestrNode  } )
end
function DeclDestrNode:get_declInfo()
   return self.declInfo
end




function NodeKind.get_ExpCallSuper(  )

   return 54
end



regKind( "ExpCallSuper" )
function Filter:processExpCallSuper( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getExpCallSuperNodeList(  )

   return self:getList( 54 )
end



local ExpCallSuperNode = {}
setmetatable( ExpCallSuperNode, { __index = Node } )
_moduleObj.ExpCallSuperNode = ExpCallSuperNode
function ExpCallSuperNode:processFilter( filter, opt )

   filter:processExpCallSuper( self, opt )
end
function ExpCallSuperNode:canBeLeft(  )

   return false
end
function ExpCallSuperNode:canBeStatement(  )

   return true
end
function ExpCallSuperNode.new( id, pos, macroArgFlag, typeList, superType, methodType, expList )
   local obj = {}
   ExpCallSuperNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, superType, methodType, expList ); end
   return obj
end
function ExpCallSuperNode:__init(id, pos, macroArgFlag, typeList, superType, methodType, expList) 
   Node.__init( self,id, 54, pos, macroArgFlag, typeList)
   
   
   
   self.superType = superType
   self.methodType = methodType
   self.expList = expList
   
   
end
function ExpCallSuperNode.create( nodeMan, pos, macroArgFlag, typeList, superType, methodType, expList )

   local node = ExpCallSuperNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, superType, methodType, expList)
   nodeMan:addNode( node )
   return node
end
function ExpCallSuperNode:visit( visitor, depth )

   do
      do
         local child = self.expList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'expList', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function ExpCallSuperNode.setmeta( obj )
  setmetatable( obj, { __index = ExpCallSuperNode  } )
end
function ExpCallSuperNode:get_superType()
   return self.superType
end
function ExpCallSuperNode:get_methodType()
   return self.methodType
end
function ExpCallSuperNode:get_expList()
   return self.expList
end



function ExpCallSuperNode:canBeRight(  )

   return self:get_expType() ~= Ast.builtinTypeNone
end



function NodeKind.get_DeclMember(  )

   return 55
end



regKind( "DeclMember" )
function Filter:processDeclMember( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclMemberNodeList(  )

   return self:getList( 55 )
end



local DeclMemberNode = {}
setmetatable( DeclMemberNode, { __index = Node } )
_moduleObj.DeclMemberNode = DeclMemberNode
function DeclMemberNode:processFilter( filter, opt )

   filter:processDeclMember( self, opt )
end
function DeclMemberNode:canBeRight(  )

   return false
end
function DeclMemberNode:canBeLeft(  )

   return false
end
function DeclMemberNode:canBeStatement(  )

   return true
end
function DeclMemberNode.new( id, pos, macroArgFlag, typeList, name, refType, symbolInfo, classType, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode )
   local obj = {}
   DeclMemberNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, name, refType, symbolInfo, classType, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode ); end
   return obj
end
function DeclMemberNode:__init(id, pos, macroArgFlag, typeList, name, refType, symbolInfo, classType, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode) 
   Node.__init( self,id, 55, pos, macroArgFlag, typeList)
   
   
   
   self.name = name
   self.refType = refType
   self.symbolInfo = symbolInfo
   self.classType = classType
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.getterMutable = getterMutable
   self.getterMode = getterMode
   self.getterRetType = getterRetType
   self.setterMode = setterMode
   
   
end
function DeclMemberNode.create( nodeMan, pos, macroArgFlag, typeList, name, refType, symbolInfo, classType, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode )

   local node = DeclMemberNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, name, refType, symbolInfo, classType, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode)
   nodeMan:addNode( node )
   return node
end
function DeclMemberNode:visit( visitor, depth )

   do
      local child = self.refType
      do
         local _switchExp = visitor( child, self, 'refType', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function DeclMemberNode.setmeta( obj )
  setmetatable( obj, { __index = DeclMemberNode  } )
end
function DeclMemberNode:get_name()
   return self.name
end
function DeclMemberNode:get_refType()
   return self.refType
end
function DeclMemberNode:get_symbolInfo()
   return self.symbolInfo
end
function DeclMemberNode:get_classType()
   return self.classType
end
function DeclMemberNode:get_staticFlag()
   return self.staticFlag
end
function DeclMemberNode:get_accessMode()
   return self.accessMode
end
function DeclMemberNode:get_getterMutable()
   return self.getterMutable
end
function DeclMemberNode:get_getterMode()
   return self.getterMode
end
function DeclMemberNode:get_getterRetType()
   return self.getterRetType
end
function DeclMemberNode:get_setterMode()
   return self.setterMode
end




function NodeKind.get_DeclArg(  )

   return 56
end



regKind( "DeclArg" )
function Filter:processDeclArg( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclArgNodeList(  )

   return self:getList( 56 )
end



local DeclArgNode = {}
setmetatable( DeclArgNode, { __index = Node } )
_moduleObj.DeclArgNode = DeclArgNode
function DeclArgNode:processFilter( filter, opt )

   filter:processDeclArg( self, opt )
end
function DeclArgNode:canBeRight(  )

   return false
end
function DeclArgNode:canBeLeft(  )

   return false
end
function DeclArgNode:canBeStatement(  )

   return false
end
function DeclArgNode.new( id, pos, macroArgFlag, typeList, name, symbolInfo, argType )
   local obj = {}
   DeclArgNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, name, symbolInfo, argType ); end
   return obj
end
function DeclArgNode:__init(id, pos, macroArgFlag, typeList, name, symbolInfo, argType) 
   Node.__init( self,id, 56, pos, macroArgFlag, typeList)
   
   
   
   self.name = name
   self.symbolInfo = symbolInfo
   self.argType = argType
   
   
end
function DeclArgNode.create( nodeMan, pos, macroArgFlag, typeList, name, symbolInfo, argType )

   local node = DeclArgNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, name, symbolInfo, argType)
   nodeMan:addNode( node )
   return node
end
function DeclArgNode:visit( visitor, depth )

   do
      do
         local child = self.argType
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'argType', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function DeclArgNode.setmeta( obj )
  setmetatable( obj, { __index = DeclArgNode  } )
end
function DeclArgNode:get_name()
   return self.name
end
function DeclArgNode:get_symbolInfo()
   return self.symbolInfo
end
function DeclArgNode:get_argType()
   return self.argType
end




function NodeKind.get_DeclArgDDD(  )

   return 57
end



regKind( "DeclArgDDD" )
function Filter:processDeclArgDDD( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclArgDDDNodeList(  )

   return self:getList( 57 )
end



local DeclArgDDDNode = {}
setmetatable( DeclArgDDDNode, { __index = Node } )
_moduleObj.DeclArgDDDNode = DeclArgDDDNode
function DeclArgDDDNode:processFilter( filter, opt )

   filter:processDeclArgDDD( self, opt )
end
function DeclArgDDDNode:canBeRight(  )

   return false
end
function DeclArgDDDNode:canBeLeft(  )

   return false
end
function DeclArgDDDNode:canBeStatement(  )

   return false
end
function DeclArgDDDNode.new( id, pos, macroArgFlag, typeList )
   local obj = {}
   DeclArgDDDNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList ); end
   return obj
end
function DeclArgDDDNode:__init(id, pos, macroArgFlag, typeList) 
   Node.__init( self,id, 57, pos, macroArgFlag, typeList)
   
   
   
   
end
function DeclArgDDDNode.create( nodeMan, pos, macroArgFlag, typeList )

   local node = DeclArgDDDNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList)
   nodeMan:addNode( node )
   return node
end
function DeclArgDDDNode:visit( visitor, depth )

   
   return true
end
function DeclArgDDDNode.setmeta( obj )
  setmetatable( obj, { __index = DeclArgDDDNode  } )
end



local AdvertiseInfo = {}
_moduleObj.AdvertiseInfo = AdvertiseInfo
function AdvertiseInfo.setmeta( obj )
  setmetatable( obj, { __index = AdvertiseInfo  } )
end
function AdvertiseInfo.new( member, prefix, pos )
   local obj = {}
   AdvertiseInfo.setmeta( obj )
   if obj.__init then
      obj:__init( member, prefix, pos )
   end
   return obj
end
function AdvertiseInfo:__init( member, prefix, pos )

   self.member = member
   self.prefix = prefix
   self.pos = pos
end
function AdvertiseInfo:get_member()
   return self.member
end
function AdvertiseInfo:get_prefix()
   return self.prefix
end
function AdvertiseInfo:get_pos()
   return self.pos
end



function NodeKind.get_DeclAdvertise(  )

   return 58
end



regKind( "DeclAdvertise" )
function Filter:processDeclAdvertise( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclAdvertiseNodeList(  )

   return self:getList( 58 )
end



local DeclAdvertiseNode = {}
setmetatable( DeclAdvertiseNode, { __index = Node } )
_moduleObj.DeclAdvertiseNode = DeclAdvertiseNode
function DeclAdvertiseNode:processFilter( filter, opt )

   filter:processDeclAdvertise( self, opt )
end
function DeclAdvertiseNode:canBeRight(  )

   return false
end
function DeclAdvertiseNode:canBeLeft(  )

   return false
end
function DeclAdvertiseNode:canBeStatement(  )

   return false
end
function DeclAdvertiseNode.new( id, pos, macroArgFlag, typeList, advInfo )
   local obj = {}
   DeclAdvertiseNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, advInfo ); end
   return obj
end
function DeclAdvertiseNode:__init(id, pos, macroArgFlag, typeList, advInfo) 
   Node.__init( self,id, 58, pos, macroArgFlag, typeList)
   
   
   
   self.advInfo = advInfo
   
   
end
function DeclAdvertiseNode.create( nodeMan, pos, macroArgFlag, typeList, advInfo )

   local node = DeclAdvertiseNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, advInfo)
   nodeMan:addNode( node )
   return node
end
function DeclAdvertiseNode:visit( visitor, depth )

   
   return true
end
function DeclAdvertiseNode.setmeta( obj )
  setmetatable( obj, { __index = DeclAdvertiseNode  } )
end
function DeclAdvertiseNode:get_advInfo()
   return self.advInfo
end



local ClassInitBlockInfo = {}
_moduleObj.ClassInitBlockInfo = ClassInitBlockInfo
function ClassInitBlockInfo.setmeta( obj )
  setmetatable( obj, { __index = ClassInitBlockInfo  } )
end
function ClassInitBlockInfo.new( func )
   local obj = {}
   ClassInitBlockInfo.setmeta( obj )
   if obj.__init then
      obj:__init( func )
   end
   return obj
end
function ClassInitBlockInfo:__init( func )

   self.func = func
end
function ClassInitBlockInfo:get_func()
   return self.func
end
function ClassInitBlockInfo:set_func( func )
   self.func = func
end


function NodeKind.get_DeclClass(  )

   return 59
end



regKind( "DeclClass" )
function Filter:processDeclClass( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclClassNodeList(  )

   return self:getList( 59 )
end



local DeclClassNode = {}
setmetatable( DeclClassNode, { __index = Node } )
_moduleObj.DeclClassNode = DeclClassNode
function DeclClassNode:processFilter( filter, opt )

   filter:processDeclClass( self, opt )
end
function DeclClassNode:canBeRight(  )

   return false
end
function DeclClassNode:canBeLeft(  )

   return false
end
function DeclClassNode:canBeStatement(  )

   return true
end
function DeclClassNode.new( id, pos, macroArgFlag, typeList, accessMode, name, gluePrefix, moduleName, allStmtList, declStmtList, fieldList, memberList, scope, initBlock, advertiseList, trustList, uninitMemberList, outerMethodSet )
   local obj = {}
   DeclClassNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, accessMode, name, gluePrefix, moduleName, allStmtList, declStmtList, fieldList, memberList, scope, initBlock, advertiseList, trustList, uninitMemberList, outerMethodSet ); end
   return obj
end
function DeclClassNode:__init(id, pos, macroArgFlag, typeList, accessMode, name, gluePrefix, moduleName, allStmtList, declStmtList, fieldList, memberList, scope, initBlock, advertiseList, trustList, uninitMemberList, outerMethodSet) 
   Node.__init( self,id, 59, pos, macroArgFlag, typeList)
   
   
   
   self.accessMode = accessMode
   self.name = name
   self.gluePrefix = gluePrefix
   self.moduleName = moduleName
   self.allStmtList = allStmtList
   self.declStmtList = declStmtList
   self.fieldList = fieldList
   self.memberList = memberList
   self.scope = scope
   self.initBlock = initBlock
   self.advertiseList = advertiseList
   self.trustList = trustList
   self.uninitMemberList = uninitMemberList
   self.outerMethodSet = outerMethodSet
   
   
end
function DeclClassNode.create( nodeMan, pos, macroArgFlag, typeList, accessMode, name, gluePrefix, moduleName, allStmtList, declStmtList, fieldList, memberList, scope, initBlock, advertiseList, trustList, uninitMemberList, outerMethodSet )

   local node = DeclClassNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, accessMode, name, gluePrefix, moduleName, allStmtList, declStmtList, fieldList, memberList, scope, initBlock, advertiseList, trustList, uninitMemberList, outerMethodSet)
   nodeMan:addNode( node )
   return node
end
function DeclClassNode:visit( visitor, depth )

   do
      local list = self.allStmtList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'allStmtList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   do
      local list = self.declStmtList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'declStmtList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   do
      local list = self.fieldList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'fieldList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   do
      local list = self.memberList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'memberList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   
   
   return true
end
function DeclClassNode.setmeta( obj )
  setmetatable( obj, { __index = DeclClassNode  } )
end
function DeclClassNode:get_accessMode()
   return self.accessMode
end
function DeclClassNode:get_name()
   return self.name
end
function DeclClassNode:get_gluePrefix()
   return self.gluePrefix
end
function DeclClassNode:get_moduleName()
   return self.moduleName
end
function DeclClassNode:get_allStmtList()
   return self.allStmtList
end
function DeclClassNode:get_declStmtList()
   return self.declStmtList
end
function DeclClassNode:get_fieldList()
   return self.fieldList
end
function DeclClassNode:get_memberList()
   return self.memberList
end
function DeclClassNode:get_scope()
   return self.scope
end
function DeclClassNode:get_initBlock()
   return self.initBlock
end
function DeclClassNode:get_advertiseList()
   return self.advertiseList
end
function DeclClassNode:get_trustList()
   return self.trustList
end
function DeclClassNode:get_uninitMemberList()
   return self.uninitMemberList
end
function DeclClassNode:get_outerMethodSet()
   return self.outerMethodSet
end



function DeclClassNode:hasUserInit(  )

   local scope = _lune.unwrap( self:get_expType():get_scope())
   local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope, Ast.ScopeAccess.Normal ))
   return not initFuncType:get_autoFlag()
end



function NodeKind.get_DeclEnum(  )

   return 60
end



regKind( "DeclEnum" )
function Filter:processDeclEnum( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclEnumNodeList(  )

   return self:getList( 60 )
end



local DeclEnumNode = {}
setmetatable( DeclEnumNode, { __index = Node } )
_moduleObj.DeclEnumNode = DeclEnumNode
function DeclEnumNode:processFilter( filter, opt )

   filter:processDeclEnum( self, opt )
end
function DeclEnumNode:canBeRight(  )

   return false
end
function DeclEnumNode:canBeLeft(  )

   return false
end
function DeclEnumNode:canBeStatement(  )

   return true
end
function DeclEnumNode.new( id, pos, macroArgFlag, typeList, enumType, accessMode, name, valueNameList, scope )
   local obj = {}
   DeclEnumNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, enumType, accessMode, name, valueNameList, scope ); end
   return obj
end
function DeclEnumNode:__init(id, pos, macroArgFlag, typeList, enumType, accessMode, name, valueNameList, scope) 
   Node.__init( self,id, 60, pos, macroArgFlag, typeList)
   
   
   
   self.enumType = enumType
   self.accessMode = accessMode
   self.name = name
   self.valueNameList = valueNameList
   self.scope = scope
   
   
end
function DeclEnumNode.create( nodeMan, pos, macroArgFlag, typeList, enumType, accessMode, name, valueNameList, scope )

   local node = DeclEnumNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, enumType, accessMode, name, valueNameList, scope)
   nodeMan:addNode( node )
   return node
end
function DeclEnumNode:visit( visitor, depth )

   
   return true
end
function DeclEnumNode.setmeta( obj )
  setmetatable( obj, { __index = DeclEnumNode  } )
end
function DeclEnumNode:get_enumType()
   return self.enumType
end
function DeclEnumNode:get_accessMode()
   return self.accessMode
end
function DeclEnumNode:get_name()
   return self.name
end
function DeclEnumNode:get_valueNameList()
   return self.valueNameList
end
function DeclEnumNode:get_scope()
   return self.scope
end




function NodeKind.get_DeclAlge(  )

   return 61
end



regKind( "DeclAlge" )
function Filter:processDeclAlge( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclAlgeNodeList(  )

   return self:getList( 61 )
end



local DeclAlgeNode = {}
setmetatable( DeclAlgeNode, { __index = Node } )
_moduleObj.DeclAlgeNode = DeclAlgeNode
function DeclAlgeNode:processFilter( filter, opt )

   filter:processDeclAlge( self, opt )
end
function DeclAlgeNode:canBeRight(  )

   return false
end
function DeclAlgeNode:canBeLeft(  )

   return false
end
function DeclAlgeNode:canBeStatement(  )

   return true
end
function DeclAlgeNode.new( id, pos, macroArgFlag, typeList, accessMode, algeType, scope )
   local obj = {}
   DeclAlgeNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, accessMode, algeType, scope ); end
   return obj
end
function DeclAlgeNode:__init(id, pos, macroArgFlag, typeList, accessMode, algeType, scope) 
   Node.__init( self,id, 61, pos, macroArgFlag, typeList)
   
   
   
   self.accessMode = accessMode
   self.algeType = algeType
   self.scope = scope
   
   
end
function DeclAlgeNode.create( nodeMan, pos, macroArgFlag, typeList, accessMode, algeType, scope )

   local node = DeclAlgeNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, accessMode, algeType, scope)
   nodeMan:addNode( node )
   return node
end
function DeclAlgeNode:visit( visitor, depth )

   
   return true
end
function DeclAlgeNode.setmeta( obj )
  setmetatable( obj, { __index = DeclAlgeNode  } )
end
function DeclAlgeNode:get_accessMode()
   return self.accessMode
end
function DeclAlgeNode:get_algeType()
   return self.algeType
end
function DeclAlgeNode:get_scope()
   return self.scope
end




function NodeKind.get_NewAlgeVal(  )

   return 62
end



regKind( "NewAlgeVal" )
function Filter:processNewAlgeVal( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getNewAlgeValNodeList(  )

   return self:getList( 62 )
end



local NewAlgeValNode = {}
setmetatable( NewAlgeValNode, { __index = Node } )
_moduleObj.NewAlgeValNode = NewAlgeValNode
function NewAlgeValNode:processFilter( filter, opt )

   filter:processNewAlgeVal( self, opt )
end
function NewAlgeValNode:canBeRight(  )

   return true
end
function NewAlgeValNode:canBeLeft(  )

   return false
end
function NewAlgeValNode:canBeStatement(  )

   return false
end
function NewAlgeValNode.new( id, pos, macroArgFlag, typeList, name, prefix, algeTypeInfo, valInfo, paramList )
   local obj = {}
   NewAlgeValNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, name, prefix, algeTypeInfo, valInfo, paramList ); end
   return obj
end
function NewAlgeValNode:__init(id, pos, macroArgFlag, typeList, name, prefix, algeTypeInfo, valInfo, paramList) 
   Node.__init( self,id, 62, pos, macroArgFlag, typeList)
   
   
   
   self.name = name
   self.prefix = prefix
   self.algeTypeInfo = algeTypeInfo
   self.valInfo = valInfo
   self.paramList = paramList
   
   
end
function NewAlgeValNode.create( nodeMan, pos, macroArgFlag, typeList, name, prefix, algeTypeInfo, valInfo, paramList )

   local node = NewAlgeValNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, name, prefix, algeTypeInfo, valInfo, paramList)
   nodeMan:addNode( node )
   return node
end
function NewAlgeValNode:visit( visitor, depth )

   do
      do
         local child = self.prefix
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'prefix', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   do
      local list = self.paramList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'paramList', depth )
            if _switchExp == NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == NodeVisitMode.End then
               return false
            end
         end
         
         
      end
      
      
   end
   
   
   
   return true
end
function NewAlgeValNode.setmeta( obj )
  setmetatable( obj, { __index = NewAlgeValNode  } )
end
function NewAlgeValNode:get_name()
   return self.name
end
function NewAlgeValNode:get_prefix()
   return self.prefix
end
function NewAlgeValNode:get_algeTypeInfo()
   return self.algeTypeInfo
end
function NewAlgeValNode:get_valInfo()
   return self.valInfo
end
function NewAlgeValNode:get_paramList()
   return self.paramList
end




function NodeKind.get_LuneControl(  )

   return 63
end



regKind( "LuneControl" )
function Filter:processLuneControl( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLuneControlNodeList(  )

   return self:getList( 63 )
end



local LuneControlNode = {}
setmetatable( LuneControlNode, { __index = Node } )
_moduleObj.LuneControlNode = LuneControlNode
function LuneControlNode:processFilter( filter, opt )

   filter:processLuneControl( self, opt )
end
function LuneControlNode:canBeRight(  )

   return false
end
function LuneControlNode:canBeLeft(  )

   return false
end
function LuneControlNode:canBeStatement(  )

   return true
end
function LuneControlNode.new( id, pos, macroArgFlag, typeList, pragma )
   local obj = {}
   LuneControlNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, pragma ); end
   return obj
end
function LuneControlNode:__init(id, pos, macroArgFlag, typeList, pragma) 
   Node.__init( self,id, 63, pos, macroArgFlag, typeList)
   
   
   
   self.pragma = pragma
   
   
end
function LuneControlNode.create( nodeMan, pos, macroArgFlag, typeList, pragma )

   local node = LuneControlNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, pragma)
   nodeMan:addNode( node )
   return node
end
function LuneControlNode:visit( visitor, depth )

   
   return true
end
function LuneControlNode.setmeta( obj )
  setmetatable( obj, { __index = LuneControlNode  } )
end
function LuneControlNode:get_pragma()
   return self.pragma
end



local MatchCase = {}
_moduleObj.MatchCase = MatchCase
function MatchCase.setmeta( obj )
  setmetatable( obj, { __index = MatchCase  } )
end
function MatchCase.new( valInfo, valParamNameList, block )
   local obj = {}
   MatchCase.setmeta( obj )
   if obj.__init then
      obj:__init( valInfo, valParamNameList, block )
   end
   return obj
end
function MatchCase:__init( valInfo, valParamNameList, block )

   self.valInfo = valInfo
   self.valParamNameList = valParamNameList
   self.block = block
end
function MatchCase:get_valInfo()
   return self.valInfo
end
function MatchCase:get_valParamNameList()
   return self.valParamNameList
end
function MatchCase:get_block()
   return self.block
end



function NodeKind.get_Match(  )

   return 64
end



regKind( "Match" )
function Filter:processMatch( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getMatchNodeList(  )

   return self:getList( 64 )
end



local MatchNode = {}
setmetatable( MatchNode, { __index = Node } )
_moduleObj.MatchNode = MatchNode
function MatchNode:processFilter( filter, opt )

   filter:processMatch( self, opt )
end
function MatchNode:canBeRight(  )

   return false
end
function MatchNode:canBeLeft(  )

   return false
end
function MatchNode:canBeStatement(  )

   return true
end
function MatchNode.new( id, pos, macroArgFlag, typeList, val, algeTypeInfo, caseList, defaultBlock, caseKind, failSafeDefault )
   local obj = {}
   MatchNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, val, algeTypeInfo, caseList, defaultBlock, caseKind, failSafeDefault ); end
   return obj
end
function MatchNode:__init(id, pos, macroArgFlag, typeList, val, algeTypeInfo, caseList, defaultBlock, caseKind, failSafeDefault) 
   Node.__init( self,id, 64, pos, macroArgFlag, typeList)
   
   
   
   self.val = val
   self.algeTypeInfo = algeTypeInfo
   self.caseList = caseList
   self.defaultBlock = defaultBlock
   self.caseKind = caseKind
   self.failSafeDefault = failSafeDefault
   
   
end
function MatchNode.create( nodeMan, pos, macroArgFlag, typeList, val, algeTypeInfo, caseList, defaultBlock, caseKind, failSafeDefault )

   local node = MatchNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, val, algeTypeInfo, caseList, defaultBlock, caseKind, failSafeDefault)
   nodeMan:addNode( node )
   return node
end
function MatchNode:visit( visitor, depth )

   do
      local child = self.val
      do
         local _switchExp = visitor( child, self, 'val', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      do
         local child = self.defaultBlock
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'defaultBlock', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function MatchNode.setmeta( obj )
  setmetatable( obj, { __index = MatchNode  } )
end
function MatchNode:get_val()
   return self.val
end
function MatchNode:get_algeTypeInfo()
   return self.algeTypeInfo
end
function MatchNode:get_caseList()
   return self.caseList
end
function MatchNode:get_defaultBlock()
   return self.defaultBlock
end
function MatchNode:get_caseKind()
   return self.caseKind
end
function MatchNode:get_failSafeDefault()
   return self.failSafeDefault
end


function MatchNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   local fullCase = self.caseKind ~= CaseKind.Lack
   for __index, caseInfo in ipairs( self.caseList ) do
      local work = caseInfo:get_block():getBreakKind( checkMode )
      local goNext = (work == BreakKind.None ) or not fullCase
      
      if checkMode == CheckBreakMode.IgnoreFlowReturn then
         if work == BreakKind.Return then
            return BreakKind.Return
         end
         
         if work == BreakKind.NeverRet then
            return BreakKind.NeverRet
         end
         
      else
       
         do
            local _switchExp = work
            if _switchExp == BreakKind.None then
               if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                  if goNext then
                     return BreakKind.None
                  end
                  
               end
               
            else 
               
                  if kind == BreakKind.None or kind > work then
                     kind = work
                  end
                  
            end
         end
         
      end
      
      
   end
   
   do
      local block = self.defaultBlock
      if block ~= nil then
         local work = block:getBreakKind( checkMode )
         
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     if true then
                        return BreakKind.None
                     end
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
         return kind
      end
   end
   
   
   if fullCase then
      return kind
   end
   
   return BreakKind.None
end



function NodeKind.get_LuneKind(  )

   return 65
end



regKind( "LuneKind" )
function Filter:processLuneKind( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLuneKindNodeList(  )

   return self:getList( 65 )
end



local LuneKindNode = {}
setmetatable( LuneKindNode, { __index = Node } )
_moduleObj.LuneKindNode = LuneKindNode
function LuneKindNode:processFilter( filter, opt )

   filter:processLuneKind( self, opt )
end
function LuneKindNode:canBeRight(  )

   return true
end
function LuneKindNode:canBeLeft(  )

   return false
end
function LuneKindNode:canBeStatement(  )

   return false
end
function LuneKindNode.new( id, pos, macroArgFlag, typeList, exp )
   local obj = {}
   LuneKindNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, exp ); end
   return obj
end
function LuneKindNode:__init(id, pos, macroArgFlag, typeList, exp) 
   Node.__init( self,id, 65, pos, macroArgFlag, typeList)
   
   
   
   self.exp = exp
   
   
end
function LuneKindNode.create( nodeMan, pos, macroArgFlag, typeList, exp )

   local node = LuneKindNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, exp)
   nodeMan:addNode( node )
   return node
end
function LuneKindNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function LuneKindNode.setmeta( obj )
  setmetatable( obj, { __index = LuneKindNode  } )
end
function LuneKindNode:get_exp()
   return self.exp
end




function NodeKind.get_DeclMacro(  )

   return 66
end



regKind( "DeclMacro" )
function Filter:processDeclMacro( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getDeclMacroNodeList(  )

   return self:getList( 66 )
end



local DeclMacroNode = {}
setmetatable( DeclMacroNode, { __index = Node } )
_moduleObj.DeclMacroNode = DeclMacroNode
function DeclMacroNode:processFilter( filter, opt )

   filter:processDeclMacro( self, opt )
end
function DeclMacroNode:canBeRight(  )

   return false
end
function DeclMacroNode:canBeLeft(  )

   return false
end
function DeclMacroNode:canBeStatement(  )

   return true
end
function DeclMacroNode.new( id, pos, macroArgFlag, typeList, declInfo )
   local obj = {}
   DeclMacroNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, declInfo ); end
   return obj
end
function DeclMacroNode:__init(id, pos, macroArgFlag, typeList, declInfo) 
   Node.__init( self,id, 66, pos, macroArgFlag, typeList)
   
   
   
   self.declInfo = declInfo
   
   
end
function DeclMacroNode.create( nodeMan, pos, macroArgFlag, typeList, declInfo )

   local node = DeclMacroNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclMacroNode:visit( visitor, depth )

   
   return true
end
function DeclMacroNode.setmeta( obj )
  setmetatable( obj, { __index = DeclMacroNode  } )
end
function DeclMacroNode:get_declInfo()
   return self.declInfo
end



local MacroEval = {}
_moduleObj.MacroEval = MacroEval
function MacroEval.setmeta( obj )
  setmetatable( obj, { __index = MacroEval  } )
end
function MacroEval.new(  )
   local obj = {}
   MacroEval.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function MacroEval:__init(  )

end



function NodeKind.get_TestBlock(  )

   return 67
end



regKind( "TestBlock" )
function Filter:processTestBlock( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getTestBlockNodeList(  )

   return self:getList( 67 )
end



local TestBlockNode = {}
setmetatable( TestBlockNode, { __index = Node } )
_moduleObj.TestBlockNode = TestBlockNode
function TestBlockNode:processFilter( filter, opt )

   filter:processTestBlock( self, opt )
end
function TestBlockNode:canBeRight(  )

   return false
end
function TestBlockNode:canBeLeft(  )

   return false
end
function TestBlockNode:canBeStatement(  )

   return true
end
function TestBlockNode.new( id, pos, macroArgFlag, typeList, name, block )
   local obj = {}
   TestBlockNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, name, block ); end
   return obj
end
function TestBlockNode:__init(id, pos, macroArgFlag, typeList, name, block) 
   Node.__init( self,id, 67, pos, macroArgFlag, typeList)
   
   
   
   self.name = name
   self.block = block
   
   
end
function TestBlockNode.create( nodeMan, pos, macroArgFlag, typeList, name, block )

   local node = TestBlockNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, name, block)
   nodeMan:addNode( node )
   return node
end
function TestBlockNode:visit( visitor, depth )

   do
      local child = self.block
      do
         local _switchExp = visitor( child, self, 'block', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function TestBlockNode.setmeta( obj )
  setmetatable( obj, { __index = TestBlockNode  } )
end
function TestBlockNode:get_name()
   return self.name
end
function TestBlockNode:get_block()
   return self.block
end




function NodeKind.get_Abbr(  )

   return 68
end



regKind( "Abbr" )
function Filter:processAbbr( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getAbbrNodeList(  )

   return self:getList( 68 )
end



local AbbrNode = {}
setmetatable( AbbrNode, { __index = Node } )
_moduleObj.AbbrNode = AbbrNode
function AbbrNode:processFilter( filter, opt )

   filter:processAbbr( self, opt )
end
function AbbrNode:canBeRight(  )

   return true
end
function AbbrNode:canBeLeft(  )

   return false
end
function AbbrNode:canBeStatement(  )

   return false
end
function AbbrNode.new( id, pos, macroArgFlag, typeList )
   local obj = {}
   AbbrNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList ); end
   return obj
end
function AbbrNode:__init(id, pos, macroArgFlag, typeList) 
   Node.__init( self,id, 68, pos, macroArgFlag, typeList)
   
   
   
   
end
function AbbrNode.create( nodeMan, pos, macroArgFlag, typeList )

   local node = AbbrNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList)
   nodeMan:addNode( node )
   return node
end
function AbbrNode:visit( visitor, depth )

   
   return true
end
function AbbrNode.setmeta( obj )
  setmetatable( obj, { __index = AbbrNode  } )
end




function NodeKind.get_Boxing(  )

   return 69
end



regKind( "Boxing" )
function Filter:processBoxing( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getBoxingNodeList(  )

   return self:getList( 69 )
end



local BoxingNode = {}
setmetatable( BoxingNode, { __index = Node } )
_moduleObj.BoxingNode = BoxingNode
function BoxingNode:processFilter( filter, opt )

   filter:processBoxing( self, opt )
end
function BoxingNode:canBeRight(  )

   return true
end
function BoxingNode:canBeLeft(  )

   return false
end
function BoxingNode:canBeStatement(  )

   return false
end
function BoxingNode.new( id, pos, macroArgFlag, typeList, src )
   local obj = {}
   BoxingNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, src ); end
   return obj
end
function BoxingNode:__init(id, pos, macroArgFlag, typeList, src) 
   Node.__init( self,id, 69, pos, macroArgFlag, typeList)
   
   
   
   self.src = src
   
   
end
function BoxingNode.create( nodeMan, pos, macroArgFlag, typeList, src )

   local node = BoxingNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, src)
   nodeMan:addNode( node )
   return node
end
function BoxingNode:visit( visitor, depth )

   do
      local child = self.src
      do
         local _switchExp = visitor( child, self, 'src', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function BoxingNode.setmeta( obj )
  setmetatable( obj, { __index = BoxingNode  } )
end
function BoxingNode:get_src()
   return self.src
end




function NodeKind.get_Unboxing(  )

   return 70
end



regKind( "Unboxing" )
function Filter:processUnboxing( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getUnboxingNodeList(  )

   return self:getList( 70 )
end



local UnboxingNode = {}
setmetatable( UnboxingNode, { __index = Node } )
_moduleObj.UnboxingNode = UnboxingNode
function UnboxingNode:processFilter( filter, opt )

   filter:processUnboxing( self, opt )
end
function UnboxingNode:canBeRight(  )

   return true
end
function UnboxingNode:canBeLeft(  )

   return false
end
function UnboxingNode:canBeStatement(  )

   return false
end
function UnboxingNode.new( id, pos, macroArgFlag, typeList, src )
   local obj = {}
   UnboxingNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, src ); end
   return obj
end
function UnboxingNode:__init(id, pos, macroArgFlag, typeList, src) 
   Node.__init( self,id, 70, pos, macroArgFlag, typeList)
   
   
   
   self.src = src
   
   
end
function UnboxingNode.create( nodeMan, pos, macroArgFlag, typeList, src )

   local node = UnboxingNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, src)
   nodeMan:addNode( node )
   return node
end
function UnboxingNode:visit( visitor, depth )

   do
      local child = self.src
      do
         local _switchExp = visitor( child, self, 'src', depth )
         if _switchExp == NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function UnboxingNode.setmeta( obj )
  setmetatable( obj, { __index = UnboxingNode  } )
end
function UnboxingNode:get_src()
   return self.src
end




function NodeKind.get_LiteralNil(  )

   return 71
end



regKind( "LiteralNil" )
function Filter:processLiteralNil( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralNilNodeList(  )

   return self:getList( 71 )
end



local LiteralNilNode = {}
setmetatable( LiteralNilNode, { __index = Node } )
_moduleObj.LiteralNilNode = LiteralNilNode
function LiteralNilNode:processFilter( filter, opt )

   filter:processLiteralNil( self, opt )
end
function LiteralNilNode:canBeRight(  )

   return true
end
function LiteralNilNode:canBeLeft(  )

   return false
end
function LiteralNilNode:canBeStatement(  )

   return false
end
function LiteralNilNode.new( id, pos, macroArgFlag, typeList )
   local obj = {}
   LiteralNilNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList ); end
   return obj
end
function LiteralNilNode:__init(id, pos, macroArgFlag, typeList) 
   Node.__init( self,id, 71, pos, macroArgFlag, typeList)
   
   
   
   
end
function LiteralNilNode.create( nodeMan, pos, macroArgFlag, typeList )

   local node = LiteralNilNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList)
   nodeMan:addNode( node )
   return node
end
function LiteralNilNode:visit( visitor, depth )

   
   return true
end
function LiteralNilNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralNilNode  } )
end




function NodeKind.get_LiteralChar(  )

   return 72
end



regKind( "LiteralChar" )
function Filter:processLiteralChar( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralCharNodeList(  )

   return self:getList( 72 )
end



local LiteralCharNode = {}
setmetatable( LiteralCharNode, { __index = Node } )
_moduleObj.LiteralCharNode = LiteralCharNode
function LiteralCharNode:processFilter( filter, opt )

   filter:processLiteralChar( self, opt )
end
function LiteralCharNode:canBeRight(  )

   return true
end
function LiteralCharNode:canBeLeft(  )

   return false
end
function LiteralCharNode:canBeStatement(  )

   return false
end
function LiteralCharNode.new( id, pos, macroArgFlag, typeList, token, num )
   local obj = {}
   LiteralCharNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, token, num ); end
   return obj
end
function LiteralCharNode:__init(id, pos, macroArgFlag, typeList, token, num) 
   Node.__init( self,id, 72, pos, macroArgFlag, typeList)
   
   
   
   self.token = token
   self.num = num
   
   
end
function LiteralCharNode.create( nodeMan, pos, macroArgFlag, typeList, token, num )

   local node = LiteralCharNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, token, num)
   nodeMan:addNode( node )
   return node
end
function LiteralCharNode:visit( visitor, depth )

   
   return true
end
function LiteralCharNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralCharNode  } )
end
function LiteralCharNode:get_token()
   return self.token
end
function LiteralCharNode:get_num()
   return self.num
end




function NodeKind.get_LiteralInt(  )

   return 73
end



regKind( "LiteralInt" )
function Filter:processLiteralInt( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralIntNodeList(  )

   return self:getList( 73 )
end



local LiteralIntNode = {}
setmetatable( LiteralIntNode, { __index = Node } )
_moduleObj.LiteralIntNode = LiteralIntNode
function LiteralIntNode:processFilter( filter, opt )

   filter:processLiteralInt( self, opt )
end
function LiteralIntNode:canBeRight(  )

   return true
end
function LiteralIntNode:canBeLeft(  )

   return false
end
function LiteralIntNode:canBeStatement(  )

   return false
end
function LiteralIntNode.new( id, pos, macroArgFlag, typeList, token, num )
   local obj = {}
   LiteralIntNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, token, num ); end
   return obj
end
function LiteralIntNode:__init(id, pos, macroArgFlag, typeList, token, num) 
   Node.__init( self,id, 73, pos, macroArgFlag, typeList)
   
   
   
   self.token = token
   self.num = num
   
   
end
function LiteralIntNode.create( nodeMan, pos, macroArgFlag, typeList, token, num )

   local node = LiteralIntNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, token, num)
   nodeMan:addNode( node )
   return node
end
function LiteralIntNode:visit( visitor, depth )

   
   return true
end
function LiteralIntNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralIntNode  } )
end
function LiteralIntNode:get_token()
   return self.token
end
function LiteralIntNode:get_num()
   return self.num
end




function NodeKind.get_LiteralReal(  )

   return 74
end



regKind( "LiteralReal" )
function Filter:processLiteralReal( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralRealNodeList(  )

   return self:getList( 74 )
end



local LiteralRealNode = {}
setmetatable( LiteralRealNode, { __index = Node } )
_moduleObj.LiteralRealNode = LiteralRealNode
function LiteralRealNode:processFilter( filter, opt )

   filter:processLiteralReal( self, opt )
end
function LiteralRealNode:canBeRight(  )

   return true
end
function LiteralRealNode:canBeLeft(  )

   return false
end
function LiteralRealNode:canBeStatement(  )

   return false
end
function LiteralRealNode.new( id, pos, macroArgFlag, typeList, token, num )
   local obj = {}
   LiteralRealNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, token, num ); end
   return obj
end
function LiteralRealNode:__init(id, pos, macroArgFlag, typeList, token, num) 
   Node.__init( self,id, 74, pos, macroArgFlag, typeList)
   
   
   
   self.token = token
   self.num = num
   
   
end
function LiteralRealNode.create( nodeMan, pos, macroArgFlag, typeList, token, num )

   local node = LiteralRealNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, token, num)
   nodeMan:addNode( node )
   return node
end
function LiteralRealNode:visit( visitor, depth )

   
   return true
end
function LiteralRealNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralRealNode  } )
end
function LiteralRealNode:get_token()
   return self.token
end
function LiteralRealNode:get_num()
   return self.num
end




function NodeKind.get_LiteralArray(  )

   return 75
end



regKind( "LiteralArray" )
function Filter:processLiteralArray( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralArrayNodeList(  )

   return self:getList( 75 )
end



local LiteralArrayNode = {}
setmetatable( LiteralArrayNode, { __index = Node } )
_moduleObj.LiteralArrayNode = LiteralArrayNode
function LiteralArrayNode:processFilter( filter, opt )

   filter:processLiteralArray( self, opt )
end
function LiteralArrayNode:canBeRight(  )

   return true
end
function LiteralArrayNode:canBeLeft(  )

   return false
end
function LiteralArrayNode:canBeStatement(  )

   return false
end
function LiteralArrayNode.new( id, pos, macroArgFlag, typeList, expList )
   local obj = {}
   LiteralArrayNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, expList ); end
   return obj
end
function LiteralArrayNode:__init(id, pos, macroArgFlag, typeList, expList) 
   Node.__init( self,id, 75, pos, macroArgFlag, typeList)
   
   
   
   self.expList = expList
   
   
end
function LiteralArrayNode.create( nodeMan, pos, macroArgFlag, typeList, expList )

   local node = LiteralArrayNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, expList)
   nodeMan:addNode( node )
   return node
end
function LiteralArrayNode:visit( visitor, depth )

   do
      do
         local child = self.expList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'expList', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function LiteralArrayNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralArrayNode  } )
end
function LiteralArrayNode:get_expList()
   return self.expList
end




function NodeKind.get_LiteralList(  )

   return 76
end



regKind( "LiteralList" )
function Filter:processLiteralList( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralListNodeList(  )

   return self:getList( 76 )
end



local LiteralListNode = {}
setmetatable( LiteralListNode, { __index = Node } )
_moduleObj.LiteralListNode = LiteralListNode
function LiteralListNode:processFilter( filter, opt )

   filter:processLiteralList( self, opt )
end
function LiteralListNode:canBeRight(  )

   return true
end
function LiteralListNode:canBeLeft(  )

   return false
end
function LiteralListNode:canBeStatement(  )

   return false
end
function LiteralListNode.new( id, pos, macroArgFlag, typeList, expList )
   local obj = {}
   LiteralListNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, expList ); end
   return obj
end
function LiteralListNode:__init(id, pos, macroArgFlag, typeList, expList) 
   Node.__init( self,id, 76, pos, macroArgFlag, typeList)
   
   
   
   self.expList = expList
   
   
end
function LiteralListNode.create( nodeMan, pos, macroArgFlag, typeList, expList )

   local node = LiteralListNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, expList)
   nodeMan:addNode( node )
   return node
end
function LiteralListNode:visit( visitor, depth )

   do
      do
         local child = self.expList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'expList', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function LiteralListNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralListNode  } )
end
function LiteralListNode:get_expList()
   return self.expList
end




function NodeKind.get_LiteralSet(  )

   return 77
end



regKind( "LiteralSet" )
function Filter:processLiteralSet( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralSetNodeList(  )

   return self:getList( 77 )
end



local LiteralSetNode = {}
setmetatable( LiteralSetNode, { __index = Node } )
_moduleObj.LiteralSetNode = LiteralSetNode
function LiteralSetNode:processFilter( filter, opt )

   filter:processLiteralSet( self, opt )
end
function LiteralSetNode:canBeRight(  )

   return true
end
function LiteralSetNode:canBeLeft(  )

   return false
end
function LiteralSetNode:canBeStatement(  )

   return false
end
function LiteralSetNode.new( id, pos, macroArgFlag, typeList, expList )
   local obj = {}
   LiteralSetNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, expList ); end
   return obj
end
function LiteralSetNode:__init(id, pos, macroArgFlag, typeList, expList) 
   Node.__init( self,id, 77, pos, macroArgFlag, typeList)
   
   
   
   self.expList = expList
   
   
end
function LiteralSetNode.create( nodeMan, pos, macroArgFlag, typeList, expList )

   local node = LiteralSetNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, expList)
   nodeMan:addNode( node )
   return node
end
function LiteralSetNode:visit( visitor, depth )

   do
      do
         local child = self.expList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'expList', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function LiteralSetNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralSetNode  } )
end
function LiteralSetNode:get_expList()
   return self.expList
end



local PairItem = {}
_moduleObj.PairItem = PairItem
function PairItem.setmeta( obj )
  setmetatable( obj, { __index = PairItem  } )
end
function PairItem.new( key, val )
   local obj = {}
   PairItem.setmeta( obj )
   if obj.__init then
      obj:__init( key, val )
   end
   return obj
end
function PairItem:__init( key, val )

   self.key = key
   self.val = val
end
function PairItem:get_key()
   return self.key
end
function PairItem:get_val()
   return self.val
end


function NodeKind.get_LiteralMap(  )

   return 78
end



regKind( "LiteralMap" )
function Filter:processLiteralMap( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralMapNodeList(  )

   return self:getList( 78 )
end



local LiteralMapNode = {}
setmetatable( LiteralMapNode, { __index = Node } )
_moduleObj.LiteralMapNode = LiteralMapNode
function LiteralMapNode:processFilter( filter, opt )

   filter:processLiteralMap( self, opt )
end
function LiteralMapNode:canBeRight(  )

   return true
end
function LiteralMapNode:canBeLeft(  )

   return false
end
function LiteralMapNode:canBeStatement(  )

   return false
end
function LiteralMapNode.new( id, pos, macroArgFlag, typeList, map, pairList )
   local obj = {}
   LiteralMapNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, map, pairList ); end
   return obj
end
function LiteralMapNode:__init(id, pos, macroArgFlag, typeList, map, pairList) 
   Node.__init( self,id, 78, pos, macroArgFlag, typeList)
   
   
   
   self.map = map
   self.pairList = pairList
   
   
end
function LiteralMapNode.create( nodeMan, pos, macroArgFlag, typeList, map, pairList )

   local node = LiteralMapNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, map, pairList)
   nodeMan:addNode( node )
   return node
end
function LiteralMapNode:visit( visitor, depth )

   do
      local map = self.map
      for key, val in pairs( map ) do
         do
            local child = key
            do
               local _switchExp = visitor( child, self, 'map', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
         
         do
            local child = val
            do
               local _switchExp = visitor( child, self, 'map', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
         
      end
      
      
   end
   
   
   
   return true
end
function LiteralMapNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralMapNode  } )
end
function LiteralMapNode:get_map()
   return self.map
end
function LiteralMapNode:get_pairList()
   return self.pairList
end




function NodeKind.get_LiteralString(  )

   return 79
end



regKind( "LiteralString" )
function Filter:processLiteralString( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralStringNodeList(  )

   return self:getList( 79 )
end



local LiteralStringNode = {}
setmetatable( LiteralStringNode, { __index = Node } )
_moduleObj.LiteralStringNode = LiteralStringNode
function LiteralStringNode:processFilter( filter, opt )

   filter:processLiteralString( self, opt )
end
function LiteralStringNode:canBeRight(  )

   return true
end
function LiteralStringNode:canBeLeft(  )

   return false
end
function LiteralStringNode:canBeStatement(  )

   return false
end
function LiteralStringNode.new( id, pos, macroArgFlag, typeList, token, orgParam, dddParam )
   local obj = {}
   LiteralStringNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, token, orgParam, dddParam ); end
   return obj
end
function LiteralStringNode:__init(id, pos, macroArgFlag, typeList, token, orgParam, dddParam) 
   Node.__init( self,id, 79, pos, macroArgFlag, typeList)
   
   
   
   self.token = token
   self.orgParam = orgParam
   self.dddParam = dddParam
   
   
end
function LiteralStringNode.create( nodeMan, pos, macroArgFlag, typeList, token, orgParam, dddParam )

   local node = LiteralStringNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, token, orgParam, dddParam)
   nodeMan:addNode( node )
   return node
end
function LiteralStringNode:visit( visitor, depth )

   do
      do
         local child = self.orgParam
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'orgParam', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   do
      do
         local child = self.dddParam
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'dddParam', depth )
               if _switchExp == NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == NodeVisitMode.End then
                  return false
               end
            end
            
            
         end
      end
      
   end
   
   
   
   return true
end
function LiteralStringNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralStringNode  } )
end
function LiteralStringNode:get_token()
   return self.token
end
function LiteralStringNode:get_orgParam()
   return self.orgParam
end
function LiteralStringNode:get_dddParam()
   return self.dddParam
end




function NodeKind.get_LiteralBool(  )

   return 80
end



regKind( "LiteralBool" )
function Filter:processLiteralBool( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralBoolNodeList(  )

   return self:getList( 80 )
end



local LiteralBoolNode = {}
setmetatable( LiteralBoolNode, { __index = Node } )
_moduleObj.LiteralBoolNode = LiteralBoolNode
function LiteralBoolNode:processFilter( filter, opt )

   filter:processLiteralBool( self, opt )
end
function LiteralBoolNode:canBeRight(  )

   return true
end
function LiteralBoolNode:canBeLeft(  )

   return false
end
function LiteralBoolNode:canBeStatement(  )

   return false
end
function LiteralBoolNode.new( id, pos, macroArgFlag, typeList, token )
   local obj = {}
   LiteralBoolNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, token ); end
   return obj
end
function LiteralBoolNode:__init(id, pos, macroArgFlag, typeList, token) 
   Node.__init( self,id, 80, pos, macroArgFlag, typeList)
   
   
   
   self.token = token
   
   
end
function LiteralBoolNode.create( nodeMan, pos, macroArgFlag, typeList, token )

   local node = LiteralBoolNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, token)
   nodeMan:addNode( node )
   return node
end
function LiteralBoolNode:visit( visitor, depth )

   
   return true
end
function LiteralBoolNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralBoolNode  } )
end
function LiteralBoolNode:get_token()
   return self.token
end




function NodeKind.get_LiteralSymbol(  )

   return 81
end



regKind( "LiteralSymbol" )
function Filter:processLiteralSymbol( node, opt )

   self:defaultProcess( node, opt )
end


function NodeManager:getLiteralSymbolNodeList(  )

   return self:getList( 81 )
end



local LiteralSymbolNode = {}
setmetatable( LiteralSymbolNode, { __index = Node } )
_moduleObj.LiteralSymbolNode = LiteralSymbolNode
function LiteralSymbolNode:processFilter( filter, opt )

   filter:processLiteralSymbol( self, opt )
end
function LiteralSymbolNode:canBeRight(  )

   return true
end
function LiteralSymbolNode:canBeLeft(  )

   return false
end
function LiteralSymbolNode:canBeStatement(  )

   return false
end
function LiteralSymbolNode.new( id, pos, macroArgFlag, typeList, token )
   local obj = {}
   LiteralSymbolNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, macroArgFlag, typeList, token ); end
   return obj
end
function LiteralSymbolNode:__init(id, pos, macroArgFlag, typeList, token) 
   Node.__init( self,id, 81, pos, macroArgFlag, typeList)
   
   
   
   self.token = token
   
   
end
function LiteralSymbolNode.create( nodeMan, pos, macroArgFlag, typeList, token )

   local node = LiteralSymbolNode.new(nodeMan:nextId(  ), pos, macroArgFlag, typeList, token)
   nodeMan:addNode( node )
   return node
end
function LiteralSymbolNode:visit( visitor, depth )

   
   return true
end
function LiteralSymbolNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralSymbolNode  } )
end
function LiteralSymbolNode:get_token()
   return self.token
end



function Node:getSymbolInfo(  )

   local function processExpNode( node )
   
      do
         local _switchExp = (node:get_kind() )
         if _switchExp == NodeKind.get_ExpRef() then
            return {(_lune.unwrap( (_lune.__Cast( node, 3, ExpRefNode ) )) ):get_symbolInfo()}
         elseif _switchExp == NodeKind.get_RefField() then
            do
               local refFieldNode = _lune.__Cast( node, 3, RefFieldNode )
               if refFieldNode ~= nil then
                  if refFieldNode:get_nilAccess() then
                     return {}
                  end
                  
                  do
                     local _exp = refFieldNode:get_symbolInfo()
                     if _exp ~= nil then
                        return {_exp}
                     end
                  end
                  
               end
            end
            
         elseif _switchExp == NodeKind.get_ExpList() then
            do
               local expListNode = _lune.__Cast( node, 3, ExpListNode )
               if expListNode ~= nil then
                  local list = {}
                  for index, expNode in ipairs( expListNode:get_expList() ) do
                     if index == #expListNode:get_expList() then
                        for __index, symbolInfo in ipairs( processExpNode( expNode ) ) do
                           table.insert( list, symbolInfo )
                        end
                        
                     else
                      
                        for __index, symbolInfo in ipairs( processExpNode( expNode ) ) do
                           table.insert( list, symbolInfo )
                           break
                        end
                        
                     end
                     
                  end
                  
                  return list
               end
            end
            
         elseif _switchExp == NodeKind.get_RefType() then
            do
               local refTypeNode = _lune.__Cast( node, 3, RefTypeNode )
               if refTypeNode ~= nil then
                  return refTypeNode:get_name():getSymbolInfo(  )
               end
            end
            
         end
      end
      
      return {}
   end
   return processExpNode( self )
end


function WhileNode:getBreakKind( checkMode )

   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      local kind = BreakKind.None
      for __index, stmt in ipairs( self.block:get_stmtList() ) do
         if stmt:get_kind() ~= NodeKind.get_BlankLine() then
            local work = stmt:getBreakKind( checkMode )
            
            if checkMode == CheckBreakMode.IgnoreFlowReturn then
               if work == BreakKind.Return then
                  return BreakKind.Return
               end
               
               if work == BreakKind.NeverRet then
                  return BreakKind.NeverRet
               end
               
            else
             
               do
                  local _switchExp = work
                  if _switchExp == BreakKind.None then
                     if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                        if false then
                           return BreakKind.None
                        end
                        
                     end
                     
                  else 
                     
                        if kind == BreakKind.None or kind > work then
                           kind = work
                        end
                        
                  end
               end
               
            end
            
            
         end
         
      end
      
      if kind == BreakKind.Break then
         return BreakKind.None
      end
      
      return kind
   else
    
      if self.exp:get_expType():get_nilable() then
         return BreakKind.None
      end
      
      if self.exp:get_expType():equals( Ast.builtinTypeBool ) then
         do
            local boolNode = _lune.__Cast( self.exp, 3, LiteralBoolNode )
            if boolNode ~= nil then
               if boolNode:get_token().txt == "false" then
                  return BreakKind.None
               end
               
            else
               return BreakKind.None
            end
         end
         
      end
      
      local mode = CheckBreakMode.IgnoreFlow
      local kind = BreakKind.None
      for __index, stmt in ipairs( self.block:get_stmtList() ) do
         if stmt:get_kind() ~= NodeKind.get_BlankLine() then
            local work = stmt:getBreakKind( mode )
            
            if mode == CheckBreakMode.IgnoreFlowReturn then
               if work == BreakKind.Return then
                  return BreakKind.Return
               end
               
               if work == BreakKind.NeverRet then
                  return BreakKind.NeverRet
               end
               
            else
             
               do
                  local _switchExp = work
                  if _switchExp == BreakKind.None then
                     if mode == CheckBreakMode.Normal or mode == CheckBreakMode.Return then
                        if false then
                           return BreakKind.None
                        end
                        
                     end
                     
                  else 
                     
                        if kind == BreakKind.None or kind > work then
                           kind = work
                        end
                        
                  end
               end
               
            end
            
            
         end
         
      end
      
      if kind == BreakKind.Break then
         return BreakKind.None
      end
      
      if kind == BreakKind.Return then
         return BreakKind.Return
      end
      
      return BreakKind.NeverRet
   end
   
end


function LiteralNilNode:getLiteral(  )

   return _lune.newAlge( Literal.Nil), nil
end

function LiteralNilNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Symb, "nil" )
   return true
end


function LiteralCharNode:getLiteral(  )

   return _lune.newAlge( Literal.Int, {self.num}), nil
end

function LiteralCharNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Char, string.format( "%d", self.num) )
   return true
end


function LiteralIntNode:getLiteral(  )

   return _lune.newAlge( Literal.Int, {self.num}), nil
end

function LiteralIntNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Int, string.format( "%d", self.num) )
   return true
end


function LiteralRealNode:getLiteral(  )

   return _lune.newAlge( Literal.Real, {self.num}), nil
end

function LiteralRealNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Real, string.format( "%g", self.num) )
   return true
end


function LiteralArrayNode:getLiteral(  )

   local literalList = {}
   do
      local _exp = self.expList
      if _exp ~= nil then
         for __index, node in ipairs( _exp:get_expList(  ) ) do
            local literal, mess = node:getLiteral(  )
            if literal ~= nil then
               table.insert( literalList, literal )
            else
               return nil, mess
            end
            
         end
         
      end
   end
   
   return _lune.newAlge( Literal.ARRAY, {literalList}), nil
end


function LiteralArrayNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Dlmt, "[@" )
   do
      local _exp = self.expList
      if _exp ~= nil then
         for index, node in ipairs( _exp:get_expList(  ) ) do
            if index > 1 then
               self:addTokenList( list, Parser.TokenKind.Dlmt, "," )
            end
            
            if not node:setupLiteralTokenList( list ) then
               return false
            end
            
         end
         
      end
   end
   
   self:addTokenList( list, Parser.TokenKind.Dlmt, "]" )
   return true
end


function LiteralListNode:getLiteral(  )

   local literalList = {}
   do
      local _exp = self.expList
      if _exp ~= nil then
         for __index, node in ipairs( _exp:get_expList(  ) ) do
            local literal, mess = node:getLiteral(  )
            if literal ~= nil then
               table.insert( literalList, literal )
            else
               return nil, mess
            end
            
         end
         
      end
   end
   
   return _lune.newAlge( Literal.LIST, {literalList}), nil
end


function LiteralListNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Dlmt, "[" )
   do
      local _exp = self.expList
      if _exp ~= nil then
         for index, node in ipairs( _exp:get_expList(  ) ) do
            if index > 1 then
               self:addTokenList( list, Parser.TokenKind.Dlmt, "," )
            end
            
            if not node:setupLiteralTokenList( list ) then
               return false
            end
            
         end
         
      end
   end
   
   self:addTokenList( list, Parser.TokenKind.Dlmt, "]" )
   return true
end


function LiteralSetNode:getLiteral(  )

   local literalList = {}
   do
      local _exp = self.expList
      if _exp ~= nil then
         for __index, node in ipairs( _exp:get_expList(  ) ) do
            local literal, mess = node:getLiteral(  )
            if literal ~= nil then
               table.insert( literalList, literal )
            else
               return nil, mess
            end
            
         end
         
      end
   end
   
   return _lune.newAlge( Literal.SET, {literalList}), nil
end


function LiteralSetNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Dlmt, "(@" )
   do
      local _exp = self.expList
      if _exp ~= nil then
         for index, node in ipairs( _exp:get_expList(  ) ) do
            if index > 1 then
               self:addTokenList( list, Parser.TokenKind.Dlmt, "," )
            end
            
            if not node:setupLiteralTokenList( list ) then
               return false
            end
            
         end
         
      end
   end
   
   self:addTokenList( list, Parser.TokenKind.Dlmt, ")" )
   return true
end


function LiteralMapNode:getLiteral(  )

   local litMap = {}
   for key, val in pairs( self.map ) do
      local keyLiteral, keyMess = key:getLiteral(  )
      local valLiteral, valMess = val:getLiteral(  )
      if keyLiteral ~= nil and valLiteral ~= nil then
         litMap[keyLiteral] = valLiteral
      else
         if not keyLiteral then
            return nil, keyMess
         end
         
         if not valLiteral then
            return nil, valMess
         end
         
      end
      
   end
   
   return _lune.newAlge( Literal.MAP, {litMap}), nil
end


function LiteralMapNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Dlmt, "{" )
   
   local lit2valNode = {}
   for key, _8142 in pairs( self.map ) do
      local literal = key:getLiteral(  )
      if literal ~= nil then
         do
            local _matchExp = literal
            if _matchExp[1] == Literal.Int[1] then
               local param = _matchExp[2][1]
            
               lit2valNode[param] = key
            elseif _matchExp[1] == Literal.Str[1] then
               local param = _matchExp[2][1]
            
               lit2valNode[param] = key
            elseif _matchExp[1] == Literal.Real[1] then
               local param = _matchExp[2][1]
            
               lit2valNode[param] = key
            else 
               
                  return false
            end
         end
         
      end
      
   end
   
   
   do
      local __sorted = {}
      local __map = lit2valNode
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, _8150 in ipairs( __sorted ) do
         local key = __map[ _8150 ]
         do
            if not key:setupLiteralTokenList( list ) then
               return false
            end
            
            self:addTokenList( list, Parser.TokenKind.Dlmt, ":" )
            if not _lune.nilacc( self.map[key], 'setupLiteralTokenList', 'callmtd' , list ) then
               return false
            end
            
            self:addTokenList( list, Parser.TokenKind.Dlmt, "," )
         end
      end
   end
   
   self:addTokenList( list, Parser.TokenKind.Dlmt, "}" )
   return true
end


function LiteralStringNode:getLiteral(  )

   local txt = self.token.txt
   if string.find( txt, '^```' ) then
      txt = txt:sub( 4, -4 )
   else
    
      txt = txt:sub( 2, -2 )
   end
   
   
   do
      local param = self:get_orgParam()
      if param ~= nil then
         local argList = param:get_expList()
         
         local paramList = {}
         for __index, argNode in ipairs( argList ) do
            local arg, mess = argNode:getLiteral(  )
            if arg ~= nil then
               paramList[#paramList + 1] = getLiteralObj( arg )
            else
               return nil, mess
            end
            
         end
         
         txt = string.format( txt, table.unpack( paramList ) )
      end
   end
   
   return _lune.newAlge( Literal.Str, {txt}), nil
end


function LiteralStringNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Str, self.token.txt )
   do
      local param = self:get_orgParam()
      if param ~= nil then
         self:addTokenList( list, Parser.TokenKind.Dlmt, "(" )
         for index, argNode in ipairs( param:get_expList() ) do
            if index > 1 then
               self:addTokenList( list, Parser.TokenKind.Dlmt, "," )
            end
            
            if not argNode:setupLiteralTokenList( list ) then
               return false
            end
            
         end
         
         self:addTokenList( list, Parser.TokenKind.Dlmt, ")" )
      end
   end
   
   return true
end


function LiteralBoolNode:getLiteral(  )

   return _lune.newAlge( Literal.Bool, {self.token.txt == "true"}), nil
end


function LiteralBoolNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Symb, self.token.txt )
   return true
end


function LiteralSymbolNode:getLiteral(  )

   return _lune.newAlge( Literal.Symbol, {self.token.txt}), nil
end


function LiteralSymbolNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Symb, self.token.txt )
   return true
end


local function enumLiteral2Literal( obj )

   do
      local _matchExp = obj
      if _matchExp[1] == Ast.EnumLiteral.Int[1] then
         local val = _matchExp[2][1]
      
         return _lune.newAlge( Literal.Int, {val}), nil
      elseif _matchExp[1] == Ast.EnumLiteral.Real[1] then
         local val = _matchExp[2][1]
      
         return _lune.newAlge( Literal.Real, {val}), nil
      elseif _matchExp[1] == Ast.EnumLiteral.Str[1] then
         local val = _matchExp[2][1]
      
         return _lune.newAlge( Literal.Str, {val}), nil
      end
   end
   
end

function RefFieldNode:getLiteral(  )

   local typeInfo = self:get_expType()
   do
      local enumTypeInfo = _lune.__Cast( typeInfo, 3, Ast.EnumTypeInfo )
      if enumTypeInfo ~= nil then
         if _lune.__Cast( self.prefix:get_expType(), 3, Ast.EnumTypeInfo ) then
            local enumval = _lune.unwrap( enumTypeInfo:getEnumValInfo( self.field.txt ))
            return enumLiteral2Literal( enumval:get_val() )
         end
         
      end
   end
   
   
   local tokenList = {}
   local literal, mess = self.prefix:getLiteral(  )
   if literal ~= nil then
      do
         local _matchExp = literal
         if _matchExp[1] == Literal.Symbol[1] then
            local symbol = _matchExp[2][1]
         
            table.insert( tokenList, symbol )
         elseif _matchExp[1] == Literal.Field[1] then
            local symList = _matchExp[2][1]
         
            for __index, symbol in ipairs( symList ) do
               table.insert( tokenList, symbol )
            end
            
         else 
            
               return nil, string.format( "not support -- %s", Literal:_getTxt( literal)
               )
         end
      end
      
      if self.nilAccess then
         table.insert( tokenList, "$." )
      else
       
         table.insert( tokenList, "." )
      end
      
      table.insert( tokenList, self.field.txt )
      return _lune.newAlge( Literal.Field, {tokenList}), nil
   end
   
   return nil, mess
end


function ExpMacroStatNode:getLiteral(  )

   local txt = ""
   for __index, token in ipairs( self.expStrList ) do
      local literal = token:getLiteral(  )
      if literal ~= nil then
         do
            local _matchExp = literal
            if _matchExp[1] == Literal.Str[1] then
               local work = _matchExp[2][1]
            
               txt = string.format( "%s%s", txt, work)
            end
         end
         
      else
         return nil, string.format( "illegal literal -- %s", getNodeKindName( token:get_kind() ))
      end
      
   end
   
   return _lune.newAlge( Literal.Str, {txt}), nil
end


function ExpMacroArgExpNode:getLiteral(  )

   return _lune.newAlge( Literal.Str, {self:get_codeTxt()}), nil
end


function ExpRefNode:getLiteral(  )

   local typeInfo = self.symbolInfo:get_typeInfo()
   do
      local enumTypeInfo = _lune.__Cast( typeInfo, 3, Ast.EnumTypeInfo )
      if enumTypeInfo ~= nil then
         if self.symbolInfo:get_kind() == Ast.SymbolKind.Mbr and self.symbolInfo:get_namespaceTypeInfo():get_kind() == Ast.TypeInfoKind.Enum then
            local enumval = _lune.unwrap( enumTypeInfo:getEnumValInfo( self.symbolInfo:get_name() ))
            return enumLiteral2Literal( enumval:get_val() )
         end
         
      end
   end
   
   
   return nil, string.format( "unsupport refnode -- %s", typeInfo:getTxt(  ))
end


function ExpOmitEnumNode:getLiteral(  )

   local enumval = self.valInfo
   return enumLiteral2Literal( enumval:get_val() )
end


function ExpOmitEnumNode:setupLiteralTokenList( list )

   local enumval = self.valInfo
   self:addTokenList( list, Parser.TokenKind.Dlmt, "." )
   
   self:addTokenList( list, Parser.TokenKind.Symb, (Ast.EnumLiteral:_getTxt( enumval:get_val())
   :gsub( ".*%.", "" ) ) )
   return true
end


function ExpOp2Node:getValType( node )

   local literal = node:getLiteral(  )
   if  nil == literal then
      local _literal = literal
   
      return false, 0, 0.0, "", Ast.headTypeInfo
   end
   
   
   local intVal, realVal, strVal = 0, 0.0, ""
   local retTypeInfo = Ast.builtinTypeNone
   
   local function getEnum( txt, typeInfo )
   
      do
         local enumTypeInfo = _lune.__Cast( typeInfo, 3, Ast.EnumTypeInfo )
         if enumTypeInfo ~= nil then
            local valInfo = _lune.unwrap( enumTypeInfo:getEnumValInfo( txt ))
            do
               local _matchExp = valInfo:get_val()
               if _matchExp[1] == Ast.EnumLiteral.Int[1] then
                  local val = _matchExp[2][1]
               
                  intVal = val
                  realVal = val * 1.0
               elseif _matchExp[1] == Ast.EnumLiteral.Real[1] then
                  local val = _matchExp[2][1]
               
                  realVal = val
               elseif _matchExp[1] == Ast.EnumLiteral.Str[1] then
                  local val = _matchExp[2][1]
               
                  strVal = val
               end
            end
            
            retTypeInfo = enumTypeInfo:get_valTypeInfo()
         end
      end
      
   end
   
   do
      local _matchExp = literal
      if _matchExp[1] == Literal.Int[1] then
         local val = _matchExp[2][1]
      
         intVal = val
         realVal = val * 1.0
         retTypeInfo = Ast.builtinTypeInt
      elseif _matchExp[1] == Literal.Real[1] then
         local val = _matchExp[2][1]
      
         realVal = val
         intVal = math.floor(val)
         retTypeInfo = Ast.builtinTypeReal
      elseif _matchExp[1] == Literal.Str[1] then
         local val = _matchExp[2][1]
      
         strVal = val
         retTypeInfo = Ast.builtinTypeString
      else 
         
            return false, 0, 0.0, "", Ast.headTypeInfo
      end
   end
   
   return true, intVal, realVal, strVal, retTypeInfo
end


function ExpOp2Node:setupLiteralTokenList( list )

   local literal = self:getLiteral(  )
   if literal ~= nil then
      do
         local _matchExp = literal
         if _matchExp[1] == Literal.Int[1] then
            local val = _matchExp[2][1]
         
            self:addTokenList( list, Parser.TokenKind.Int, string.format( "%d", val) )
         elseif _matchExp[1] == Literal.Real[1] then
            local val = _matchExp[2][1]
         
            self:addTokenList( list, Parser.TokenKind.Real, string.format( "%g", val) )
         elseif _matchExp[1] == Literal.Str[1] then
            local val = _matchExp[2][1]
         
            self:addTokenList( list, Parser.TokenKind.Str, string.format( "%q", val) )
         else 
            
               return false
         end
      end
      
      return true
   else
      return false
   end
   
end


function ExpOp2Node:getLiteral(  )

   
   local ret1, int1, real1, str1, type1 = self:getValType( self:get_exp1() )
   local ret2, int2, real2, str2, type2 = self:getValType( self:get_exp2() )
   
   if not ret1 then
      return nil, string.format( "not support literal -- %s", getNodeKindName( self:get_exp1():get_kind() ))
   end
   
   if not ret2 then
      return nil, string.format( "not support literal -- %s", getNodeKindName( self:get_exp2():get_kind() ))
   end
   
   
   if (type1 == Ast.builtinTypeInt or type1 == Ast.builtinTypeReal ) and (type2 == Ast.builtinTypeInt or type2 == Ast.builtinTypeReal ) then
      local retType = Ast.builtinTypeInt
      if type1 == Ast.builtinTypeReal or type2 == Ast.builtinTypeReal then
         retType = Ast.builtinTypeReal
      end
      
      
      do
         local _switchExp = (self.op.txt )
         if _switchExp == "+" then
            if retType == Ast.builtinTypeInt then
               return _lune.newAlge( Literal.Int, {int1 + int2}), nil
            end
            
            return _lune.newAlge( Literal.Real, {real1 + real2}), nil
         elseif _switchExp == "-" then
            if retType == Ast.builtinTypeInt then
               return _lune.newAlge( Literal.Int, {int1 - int2}), nil
            end
            
            return _lune.newAlge( Literal.Real, {real1 - real2}), nil
         elseif _switchExp == "*" then
            if retType == Ast.builtinTypeInt then
               return _lune.newAlge( Literal.Int, {int1 * int2}), nil
            end
            
            return _lune.newAlge( Literal.Real, {real1 * real2}), nil
         elseif _switchExp == "/" then
            if retType == Ast.builtinTypeInt then
               return _lune.newAlge( Literal.Int, {math.floor(int1 / int2)}), nil
            end
            
            return _lune.newAlge( Literal.Real, {real1 / real2}), nil
         end
      end
      
   elseif type1 == Ast.builtinTypeString and type2 == Ast.builtinTypeString then
      if self.op.txt == ".." then
         return _lune.newAlge( Literal.Str, {str1 .. str2}), nil
      end
      
   end
   
   
   local mess = string.format( "not support literal operation -- %s %s %s", type1:getTxt(  ), self.op.txt, type2:getTxt(  ))
   return nil, mess
end


local DefMacroInfo = {}
setmetatable( DefMacroInfo, { __index = MacroInfo } )
_moduleObj.DefMacroInfo = DefMacroInfo
function DefMacroInfo:get_name(  )

   return self.declInfo:get_name().txt
end
function DefMacroInfo:getArgList(  )

   return self.argList
end
function DefMacroInfo:getTokenList(  )

   return self.declInfo:get_tokenList()
end
function DefMacroInfo.new( func, declInfo, symbol2MacroValInfoMap )
   local obj = {}
   DefMacroInfo.setmeta( obj )
   if obj.__init then obj:__init( func, declInfo, symbol2MacroValInfoMap ); end
   return obj
end
function DefMacroInfo:__init(func, declInfo, symbol2MacroValInfoMap) 
   MacroInfo.__init( self,func, symbol2MacroValInfoMap)
   
   self.declInfo = declInfo
   self.argList = {}
   for __index, argNode in ipairs( declInfo:get_argList() ) do
      if argNode:get_kind(  ) == NodeKind.get_DeclArg() then
         local argType = argNode:get_expType()
         local argName = argNode:get_name().txt
         table.insert( self.argList, MacroArgInfo.new(argName, argType) )
      end
      
   end
   
end
function DefMacroInfo.setmeta( obj )
  setmetatable( obj, { __index = DefMacroInfo  } )
end


function NodeManager:MultiTo1( node )

   local expType = node:get_expType()
   if #node:get_expTypeList() > 1 then
      return ExpMultiTo1Node.create( self, node:get_pos(), node:get_macroArgFlag(), {expType}, node )
   elseif expType:get_kind() == Ast.TypeInfoKind.DDD then
      return ExpMultiTo1Node.create( self, node:get_pos(), node:get_macroArgFlag(), expType:get_itemTypeInfoList(), node )
   end
   
   return node
end


function Filter:processBlockSub( node, opt )

end


function Filter:processBlock( node, opt )

   self.moduleInfoManager:push( node:get_scope() )
   
   self:processBlockSub( node, opt )
   
   self.moduleInfoManager:pop(  )
end

local function hasMultiValNode( node )

   return #node:get_expTypeList() > 1 or node:get_expType():get_kind() == Ast.TypeInfoKind.DDD
end
_moduleObj.hasMultiValNode = hasMultiValNode



local nodeKindEnum = {}
_moduleObj.nodeKindEnum = nodeKindEnum
nodeKindEnum._val2NameMap = {}
function nodeKindEnum:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "nodeKindEnum.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function nodeKindEnum._from( val )
   if nodeKindEnum._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
nodeKindEnum.__allList = {}
function nodeKindEnum.get__allList()
   return nodeKindEnum.__allList
end

nodeKindEnum.None = 0
nodeKindEnum._val2NameMap[0] = 'None'
nodeKindEnum.__allList[1] = nodeKindEnum.None
nodeKindEnum.BlankLine = 1
nodeKindEnum._val2NameMap[1] = 'BlankLine'
nodeKindEnum.__allList[2] = nodeKindEnum.BlankLine
nodeKindEnum.Subfile = 2
nodeKindEnum._val2NameMap[2] = 'Subfile'
nodeKindEnum.__allList[3] = nodeKindEnum.Subfile
nodeKindEnum.Import = 3
nodeKindEnum._val2NameMap[3] = 'Import'
nodeKindEnum.__allList[4] = nodeKindEnum.Import
nodeKindEnum.Root = 4
nodeKindEnum._val2NameMap[4] = 'Root'
nodeKindEnum.__allList[5] = nodeKindEnum.Root
nodeKindEnum.RefType = 5
nodeKindEnum._val2NameMap[5] = 'RefType'
nodeKindEnum.__allList[6] = nodeKindEnum.RefType
nodeKindEnum.Block = 6
nodeKindEnum._val2NameMap[6] = 'Block'
nodeKindEnum.__allList[7] = nodeKindEnum.Block
nodeKindEnum.Scope = 7
nodeKindEnum._val2NameMap[7] = 'Scope'
nodeKindEnum.__allList[8] = nodeKindEnum.Scope
nodeKindEnum.If = 8
nodeKindEnum._val2NameMap[8] = 'If'
nodeKindEnum.__allList[9] = nodeKindEnum.If
nodeKindEnum.ExpList = 9
nodeKindEnum._val2NameMap[9] = 'ExpList'
nodeKindEnum.__allList[10] = nodeKindEnum.ExpList
nodeKindEnum.Switch = 10
nodeKindEnum._val2NameMap[10] = 'Switch'
nodeKindEnum.__allList[11] = nodeKindEnum.Switch
nodeKindEnum.While = 11
nodeKindEnum._val2NameMap[11] = 'While'
nodeKindEnum.__allList[12] = nodeKindEnum.While
nodeKindEnum.Repeat = 12
nodeKindEnum._val2NameMap[12] = 'Repeat'
nodeKindEnum.__allList[13] = nodeKindEnum.Repeat
nodeKindEnum.For = 13
nodeKindEnum._val2NameMap[13] = 'For'
nodeKindEnum.__allList[14] = nodeKindEnum.For
nodeKindEnum.Apply = 14
nodeKindEnum._val2NameMap[14] = 'Apply'
nodeKindEnum.__allList[15] = nodeKindEnum.Apply
nodeKindEnum.Foreach = 15
nodeKindEnum._val2NameMap[15] = 'Foreach'
nodeKindEnum.__allList[16] = nodeKindEnum.Foreach
nodeKindEnum.Forsort = 16
nodeKindEnum._val2NameMap[16] = 'Forsort'
nodeKindEnum.__allList[17] = nodeKindEnum.Forsort
nodeKindEnum.Return = 17
nodeKindEnum._val2NameMap[17] = 'Return'
nodeKindEnum.__allList[18] = nodeKindEnum.Return
nodeKindEnum.Break = 18
nodeKindEnum._val2NameMap[18] = 'Break'
nodeKindEnum.__allList[19] = nodeKindEnum.Break
nodeKindEnum.Provide = 19
nodeKindEnum._val2NameMap[19] = 'Provide'
nodeKindEnum.__allList[20] = nodeKindEnum.Provide
nodeKindEnum.ExpNew = 20
nodeKindEnum._val2NameMap[20] = 'ExpNew'
nodeKindEnum.__allList[21] = nodeKindEnum.ExpNew
nodeKindEnum.ExpUnwrap = 21
nodeKindEnum._val2NameMap[21] = 'ExpUnwrap'
nodeKindEnum.__allList[22] = nodeKindEnum.ExpUnwrap
nodeKindEnum.ExpRef = 22
nodeKindEnum._val2NameMap[22] = 'ExpRef'
nodeKindEnum.__allList[23] = nodeKindEnum.ExpRef
nodeKindEnum.ExpSetVal = 23
nodeKindEnum._val2NameMap[23] = 'ExpSetVal'
nodeKindEnum.__allList[24] = nodeKindEnum.ExpSetVal
nodeKindEnum.ExpOp2 = 24
nodeKindEnum._val2NameMap[24] = 'ExpOp2'
nodeKindEnum.__allList[25] = nodeKindEnum.ExpOp2
nodeKindEnum.UnwrapSet = 25
nodeKindEnum._val2NameMap[25] = 'UnwrapSet'
nodeKindEnum.__allList[26] = nodeKindEnum.UnwrapSet
nodeKindEnum.IfUnwrap = 26
nodeKindEnum._val2NameMap[26] = 'IfUnwrap'
nodeKindEnum.__allList[27] = nodeKindEnum.IfUnwrap
nodeKindEnum.When = 27
nodeKindEnum._val2NameMap[27] = 'When'
nodeKindEnum.__allList[28] = nodeKindEnum.When
nodeKindEnum.ExpCast = 28
nodeKindEnum._val2NameMap[28] = 'ExpCast'
nodeKindEnum.__allList[29] = nodeKindEnum.ExpCast
nodeKindEnum.ExpToDDD = 29
nodeKindEnum._val2NameMap[29] = 'ExpToDDD'
nodeKindEnum.__allList[30] = nodeKindEnum.ExpToDDD
nodeKindEnum.ExpSubDDD = 30
nodeKindEnum._val2NameMap[30] = 'ExpSubDDD'
nodeKindEnum.__allList[31] = nodeKindEnum.ExpSubDDD
nodeKindEnum.ExpOp1 = 31
nodeKindEnum._val2NameMap[31] = 'ExpOp1'
nodeKindEnum.__allList[32] = nodeKindEnum.ExpOp1
nodeKindEnum.ExpRefItem = 32
nodeKindEnum._val2NameMap[32] = 'ExpRefItem'
nodeKindEnum.__allList[33] = nodeKindEnum.ExpRefItem
nodeKindEnum.ExpCall = 33
nodeKindEnum._val2NameMap[33] = 'ExpCall'
nodeKindEnum.__allList[34] = nodeKindEnum.ExpCall
nodeKindEnum.ExpAccessMRet = 34
nodeKindEnum._val2NameMap[34] = 'ExpAccessMRet'
nodeKindEnum.__allList[35] = nodeKindEnum.ExpAccessMRet
nodeKindEnum.ExpMultiTo1 = 35
nodeKindEnum._val2NameMap[35] = 'ExpMultiTo1'
nodeKindEnum.__allList[36] = nodeKindEnum.ExpMultiTo1
nodeKindEnum.ExpDDD = 36
nodeKindEnum._val2NameMap[36] = 'ExpDDD'
nodeKindEnum.__allList[37] = nodeKindEnum.ExpDDD
nodeKindEnum.ExpParen = 37
nodeKindEnum._val2NameMap[37] = 'ExpParen'
nodeKindEnum.__allList[38] = nodeKindEnum.ExpParen
nodeKindEnum.ExpMacroExp = 38
nodeKindEnum._val2NameMap[38] = 'ExpMacroExp'
nodeKindEnum.__allList[39] = nodeKindEnum.ExpMacroExp
nodeKindEnum.ExpMacroStat = 39
nodeKindEnum._val2NameMap[39] = 'ExpMacroStat'
nodeKindEnum.__allList[40] = nodeKindEnum.ExpMacroStat
nodeKindEnum.ExpMacroArgExp = 40
nodeKindEnum._val2NameMap[40] = 'ExpMacroArgExp'
nodeKindEnum.__allList[41] = nodeKindEnum.ExpMacroArgExp
nodeKindEnum.StmtExp = 41
nodeKindEnum._val2NameMap[41] = 'StmtExp'
nodeKindEnum.__allList[42] = nodeKindEnum.StmtExp
nodeKindEnum.ExpMacroStatList = 42
nodeKindEnum._val2NameMap[42] = 'ExpMacroStatList'
nodeKindEnum.__allList[43] = nodeKindEnum.ExpMacroStatList
nodeKindEnum.ExpOmitEnum = 43
nodeKindEnum._val2NameMap[43] = 'ExpOmitEnum'
nodeKindEnum.__allList[44] = nodeKindEnum.ExpOmitEnum
nodeKindEnum.RefField = 44
nodeKindEnum._val2NameMap[44] = 'RefField'
nodeKindEnum.__allList[45] = nodeKindEnum.RefField
nodeKindEnum.GetField = 45
nodeKindEnum._val2NameMap[45] = 'GetField'
nodeKindEnum.__allList[46] = nodeKindEnum.GetField
nodeKindEnum.Alias = 46
nodeKindEnum._val2NameMap[46] = 'Alias'
nodeKindEnum.__allList[47] = nodeKindEnum.Alias
nodeKindEnum.DeclVar = 47
nodeKindEnum._val2NameMap[47] = 'DeclVar'
nodeKindEnum.__allList[48] = nodeKindEnum.DeclVar
nodeKindEnum.DeclForm = 48
nodeKindEnum._val2NameMap[48] = 'DeclForm'
nodeKindEnum.__allList[49] = nodeKindEnum.DeclForm
nodeKindEnum.DeclFunc = 49
nodeKindEnum._val2NameMap[49] = 'DeclFunc'
nodeKindEnum.__allList[50] = nodeKindEnum.DeclFunc
nodeKindEnum.DeclMethod = 50
nodeKindEnum._val2NameMap[50] = 'DeclMethod'
nodeKindEnum.__allList[51] = nodeKindEnum.DeclMethod
nodeKindEnum.ProtoMethod = 51
nodeKindEnum._val2NameMap[51] = 'ProtoMethod'
nodeKindEnum.__allList[52] = nodeKindEnum.ProtoMethod
nodeKindEnum.DeclConstr = 52
nodeKindEnum._val2NameMap[52] = 'DeclConstr'
nodeKindEnum.__allList[53] = nodeKindEnum.DeclConstr
nodeKindEnum.DeclDestr = 53
nodeKindEnum._val2NameMap[53] = 'DeclDestr'
nodeKindEnum.__allList[54] = nodeKindEnum.DeclDestr
nodeKindEnum.ExpCallSuper = 54
nodeKindEnum._val2NameMap[54] = 'ExpCallSuper'
nodeKindEnum.__allList[55] = nodeKindEnum.ExpCallSuper
nodeKindEnum.DeclMember = 55
nodeKindEnum._val2NameMap[55] = 'DeclMember'
nodeKindEnum.__allList[56] = nodeKindEnum.DeclMember
nodeKindEnum.DeclArg = 56
nodeKindEnum._val2NameMap[56] = 'DeclArg'
nodeKindEnum.__allList[57] = nodeKindEnum.DeclArg
nodeKindEnum.DeclArgDDD = 57
nodeKindEnum._val2NameMap[57] = 'DeclArgDDD'
nodeKindEnum.__allList[58] = nodeKindEnum.DeclArgDDD
nodeKindEnum.DeclAdvertise = 58
nodeKindEnum._val2NameMap[58] = 'DeclAdvertise'
nodeKindEnum.__allList[59] = nodeKindEnum.DeclAdvertise
nodeKindEnum.DeclClass = 59
nodeKindEnum._val2NameMap[59] = 'DeclClass'
nodeKindEnum.__allList[60] = nodeKindEnum.DeclClass
nodeKindEnum.DeclEnum = 60
nodeKindEnum._val2NameMap[60] = 'DeclEnum'
nodeKindEnum.__allList[61] = nodeKindEnum.DeclEnum
nodeKindEnum.DeclAlge = 61
nodeKindEnum._val2NameMap[61] = 'DeclAlge'
nodeKindEnum.__allList[62] = nodeKindEnum.DeclAlge
nodeKindEnum.NewAlgeVal = 62
nodeKindEnum._val2NameMap[62] = 'NewAlgeVal'
nodeKindEnum.__allList[63] = nodeKindEnum.NewAlgeVal
nodeKindEnum.LuneControl = 63
nodeKindEnum._val2NameMap[63] = 'LuneControl'
nodeKindEnum.__allList[64] = nodeKindEnum.LuneControl
nodeKindEnum.Match = 64
nodeKindEnum._val2NameMap[64] = 'Match'
nodeKindEnum.__allList[65] = nodeKindEnum.Match
nodeKindEnum.LuneKind = 65
nodeKindEnum._val2NameMap[65] = 'LuneKind'
nodeKindEnum.__allList[66] = nodeKindEnum.LuneKind
nodeKindEnum.DeclMacro = 66
nodeKindEnum._val2NameMap[66] = 'DeclMacro'
nodeKindEnum.__allList[67] = nodeKindEnum.DeclMacro
nodeKindEnum.TestBlock = 67
nodeKindEnum._val2NameMap[67] = 'TestBlock'
nodeKindEnum.__allList[68] = nodeKindEnum.TestBlock
nodeKindEnum.Abbr = 68
nodeKindEnum._val2NameMap[68] = 'Abbr'
nodeKindEnum.__allList[69] = nodeKindEnum.Abbr
nodeKindEnum.Boxing = 69
nodeKindEnum._val2NameMap[69] = 'Boxing'
nodeKindEnum.__allList[70] = nodeKindEnum.Boxing
nodeKindEnum.Unboxing = 70
nodeKindEnum._val2NameMap[70] = 'Unboxing'
nodeKindEnum.__allList[71] = nodeKindEnum.Unboxing
nodeKindEnum.LiteralNil = 71
nodeKindEnum._val2NameMap[71] = 'LiteralNil'
nodeKindEnum.__allList[72] = nodeKindEnum.LiteralNil
nodeKindEnum.LiteralChar = 72
nodeKindEnum._val2NameMap[72] = 'LiteralChar'
nodeKindEnum.__allList[73] = nodeKindEnum.LiteralChar
nodeKindEnum.LiteralInt = 73
nodeKindEnum._val2NameMap[73] = 'LiteralInt'
nodeKindEnum.__allList[74] = nodeKindEnum.LiteralInt
nodeKindEnum.LiteralReal = 74
nodeKindEnum._val2NameMap[74] = 'LiteralReal'
nodeKindEnum.__allList[75] = nodeKindEnum.LiteralReal
nodeKindEnum.LiteralArray = 75
nodeKindEnum._val2NameMap[75] = 'LiteralArray'
nodeKindEnum.__allList[76] = nodeKindEnum.LiteralArray
nodeKindEnum.LiteralList = 76
nodeKindEnum._val2NameMap[76] = 'LiteralList'
nodeKindEnum.__allList[77] = nodeKindEnum.LiteralList
nodeKindEnum.LiteralSet = 77
nodeKindEnum._val2NameMap[77] = 'LiteralSet'
nodeKindEnum.__allList[78] = nodeKindEnum.LiteralSet
nodeKindEnum.LiteralMap = 78
nodeKindEnum._val2NameMap[78] = 'LiteralMap'
nodeKindEnum.__allList[79] = nodeKindEnum.LiteralMap
nodeKindEnum.LiteralString = 79
nodeKindEnum._val2NameMap[79] = 'LiteralString'
nodeKindEnum.__allList[80] = nodeKindEnum.LiteralString
nodeKindEnum.LiteralBool = 80
nodeKindEnum._val2NameMap[80] = 'LiteralBool'
nodeKindEnum.__allList[81] = nodeKindEnum.LiteralBool
nodeKindEnum.LiteralSymbol = 81
nodeKindEnum._val2NameMap[81] = 'LiteralSymbol'
nodeKindEnum.__allList[82] = nodeKindEnum.LiteralSymbol



return _moduleObj
