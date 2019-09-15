--lune/base/Nodes.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Nodes'
local _lune = {}
if _lune1 then
   _lune = _lune1
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

if not table.unpack then
   table.unpack = unpack
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

if not _lune1 then
   _lune1 = _lune
end
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )
local Filter = {}
_moduleObj.Filter = Filter
function Filter.setmeta( obj )
  setmetatable( obj, { __index = Filter  } )
end
function Filter.new(  )
   local obj = {}
   Filter.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Filter:__init(  )

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
   
   Util.errorLog( "unknown literal obj -- " .. Literal:_getTxt( obj)
    )
   return nil
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

   return nil
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
function Node.setmeta( obj )
  setmetatable( obj, { __index = Node  } )
end
function Node.new( id, kind, pos, expTypeList )
   local obj = {}
   Node.setmeta( obj )
   if obj.__init then
      obj:__init( id, kind, pos, expTypeList )
   end
   return obj
end
function Node:__init( id, kind, pos, expTypeList )

   self.id = id
   self.kind = kind
   self.pos = pos
   self.expTypeList = expTypeList
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
local nodeKindSeed = 1
local nodeKind = {}
_moduleObj.nodeKind = nodeKind

local function regKind( name )

   local kind = nodeKindSeed
   nodeKindSeed = nodeKindSeed + 1
   nodeKind2NameMap[kind] = name
   _moduleObj.nodeKind[name] = kind
   return kind
end

local function getNodeKindName( kind )

   return _lune.unwrap( nodeKind2NameMap[kind])
end
_moduleObj.getNodeKindName = getNodeKindName
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
   for __index, kind in pairs( _moduleObj.nodeKind ) do
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

   return _lune.unwrap( _moduleObj.nodeKind['None'])
end


regKind( "None" )
function Filter:processNone( node, opt )

end


function NodeManager:getNoneNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['None']) )
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
function NoneNode.new( id, pos, typeList )
   local obj = {}
   NoneNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList ); end
   return obj
end
function NoneNode:__init(id, pos, typeList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['None']), pos, typeList)
   
   
   
end
function NoneNode.create( nodeMan, pos, typeList )

   local node = NoneNode.new(nodeMan:nextId(  ), pos, typeList)
   nodeMan:addNode( node )
   return node
end
function NoneNode:visit( visitor, depth )

   
   return true
end
function NoneNode.setmeta( obj )
  setmetatable( obj, { __index = NoneNode  } )
end


function NodeKind.get_Subfile(  )

   return _lune.unwrap( _moduleObj.nodeKind['Subfile'])
end


regKind( "Subfile" )
function Filter:processSubfile( node, opt )

end


function NodeManager:getSubfileNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Subfile']) )
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
function SubfileNode.new( id, pos, typeList, usePath )
   local obj = {}
   SubfileNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, usePath ); end
   return obj
end
function SubfileNode:__init(id, pos, typeList, usePath) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Subfile']), pos, typeList)
   
   
   self.usePath = usePath
   
end
function SubfileNode.create( nodeMan, pos, typeList, usePath )

   local node = SubfileNode.new(nodeMan:nextId(  ), pos, typeList, usePath)
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

   return _lune.unwrap( _moduleObj.nodeKind['Import'])
end


regKind( "Import" )
function Filter:processImport( node, opt )

end


function NodeManager:getImportNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Import']) )
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
function ImportNode.new( id, pos, typeList, modulePath, assignName, moduleTypeInfo )
   local obj = {}
   ImportNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, modulePath, assignName, moduleTypeInfo ); end
   return obj
end
function ImportNode:__init(id, pos, typeList, modulePath, assignName, moduleTypeInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Import']), pos, typeList)
   
   
   self.modulePath = modulePath
   self.assignName = assignName
   self.moduleTypeInfo = moduleTypeInfo
   
end
function ImportNode.create( nodeMan, pos, typeList, modulePath, assignName, moduleTypeInfo )

   local node = ImportNode.new(nodeMan:nextId(  ), pos, typeList, modulePath, assignName, moduleTypeInfo)
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
function ImportNode:get_moduleTypeInfo()
   return self.moduleTypeInfo
end



local LuneHelperInfo = {}
_moduleObj.LuneHelperInfo = LuneHelperInfo
function LuneHelperInfo.setmeta( obj )
  setmetatable( obj, { __index = LuneHelperInfo  } )
end
function LuneHelperInfo.new( useNilAccess, useUnwrapExp, hasMappingClassDef, useLoad, useUnpack, useAlge, useSet )
   local obj = {}
   LuneHelperInfo.setmeta( obj )
   if obj.__init then
      obj:__init( useNilAccess, useUnwrapExp, hasMappingClassDef, useLoad, useUnpack, useAlge, useSet )
   end
   return obj
end
function LuneHelperInfo:__init( useNilAccess, useUnwrapExp, hasMappingClassDef, useLoad, useUnpack, useAlge, useSet )

   self.useNilAccess = useNilAccess
   self.useUnwrapExp = useUnwrapExp
   self.hasMappingClassDef = hasMappingClassDef
   self.useLoad = useLoad
   self.useUnpack = useUnpack
   self.useAlge = useAlge
   self.useSet = useSet
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

   return _lune.unwrap( _moduleObj.nodeKind['Root'])
end


regKind( "Root" )
function Filter:processRoot( node, opt )

end


function NodeManager:getRootNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Root']) )
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
function RootNode.new( id, pos, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap )
   local obj = {}
   RootNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap ); end
   return obj
end
function RootNode:__init(id, pos, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Root']), pos, typeList)
   
   
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
function RootNode.create( nodeMan, pos, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap )

   local node = RootNode.new(nodeMan:nextId(  ), pos, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap)
   nodeMan:addNode( node )
   return node
end
function RootNode:visit( visitor, depth )

   do
      local list = self.children
      for __index, child in pairs( list ) do
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

   return _lune.unwrap( _moduleObj.nodeKind['RefType'])
end


regKind( "RefType" )
function Filter:processRefType( node, opt )

end


function NodeManager:getRefTypeNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['RefType']) )
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
function RefTypeNode.new( id, pos, typeList, name, refFlag, mutFlag, array )
   local obj = {}
   RefTypeNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, name, refFlag, mutFlag, array ); end
   return obj
end
function RefTypeNode:__init(id, pos, typeList, name, refFlag, mutFlag, array) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['RefType']), pos, typeList)
   
   
   self.name = name
   self.refFlag = refFlag
   self.mutFlag = mutFlag
   self.array = array
   
end
function RefTypeNode.create( nodeMan, pos, typeList, name, refFlag, mutFlag, array )

   local node = RefTypeNode.new(nodeMan:nextId(  ), pos, typeList, name, refFlag, mutFlag, array)
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
   
   
   return true
end
function RefTypeNode.setmeta( obj )
  setmetatable( obj, { __index = RefTypeNode  } )
end
function RefTypeNode:get_name()
   return self.name
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
BlockKind.IfUnwrap = 16
BlockKind._val2NameMap[16] = 'IfUnwrap'
BlockKind.__allList[17] = BlockKind.IfUnwrap
BlockKind.When = 17
BlockKind._val2NameMap[17] = 'When'
BlockKind.__allList[18] = BlockKind.When

function NodeKind.get_Block(  )

   return _lune.unwrap( _moduleObj.nodeKind['Block'])
end


regKind( "Block" )
function Filter:processBlock( node, opt )

end


function NodeManager:getBlockNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Block']) )
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
function BlockNode.new( id, pos, typeList, blockKind, scope, stmtList )
   local obj = {}
   BlockNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, blockKind, scope, stmtList ); end
   return obj
end
function BlockNode:__init(id, pos, typeList, blockKind, scope, stmtList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Block']), pos, typeList)
   
   
   self.blockKind = blockKind
   self.scope = scope
   self.stmtList = stmtList
   
end
function BlockNode.create( nodeMan, pos, typeList, blockKind, scope, stmtList )

   local node = BlockNode.new(nodeMan:nextId(  ), pos, typeList, blockKind, scope, stmtList)
   nodeMan:addNode( node )
   return node
end
function BlockNode:visit( visitor, depth )

   do
      local list = self.stmtList
      for __index, child in pairs( list ) do
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



function BlockNode:getBreakKind( checkMode )

   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      local kind = BreakKind.None
      for __index, stmt in pairs( self.stmtList ) do
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
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
      end
      
      return kind
   else
    
      if #self.stmtList > 0 then
         local node = self.stmtList[#self.stmtList]
         return node:getBreakKind( checkMode )
      end
      
   end
   
   return BreakKind.None
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

   return _lune.unwrap( _moduleObj.nodeKind['If'])
end


regKind( "If" )
function Filter:processIf( node, opt )

end


function NodeManager:getIfNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['If']) )
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
function IfNode.new( id, pos, typeList, stmtList )
   local obj = {}
   IfNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, stmtList ); end
   return obj
end
function IfNode:__init(id, pos, typeList, stmtList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['If']), pos, typeList)
   
   
   self.stmtList = stmtList
   
end
function IfNode.create( nodeMan, pos, typeList, stmtList )

   local node = IfNode.new(nodeMan:nextId(  ), pos, typeList, stmtList)
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
   for __index, stmtInfo in pairs( self.stmtList ) do
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
                  return BreakKind.None
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpList'])
end


regKind( "ExpList" )
function Filter:processExpList( node, opt )

end


function NodeManager:getExpListNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpList']) )
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
function ExpListNode.new( id, pos, typeList, expList, mRetExp, followOn )
   local obj = {}
   ExpListNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, expList, mRetExp, followOn ); end
   return obj
end
function ExpListNode:__init(id, pos, typeList, expList, mRetExp, followOn) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpList']), pos, typeList)
   
   
   self.expList = expList
   self.mRetExp = mRetExp
   self.followOn = followOn
   
end
function ExpListNode.create( nodeMan, pos, typeList, expList, mRetExp, followOn )

   local node = ExpListNode.new(nodeMan:nextId(  ), pos, typeList, expList, mRetExp, followOn)
   nodeMan:addNode( node )
   return node
end
function ExpListNode:visit( visitor, depth )

   do
      local list = self.expList
      for __index, child in pairs( list ) do
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

   for __index, expNode in pairs( self:get_expList() ) do
      if not expNode:canBeLeft(  ) then
         return false
      end
      
   end
   
   return true
end

function ExpListNode:canBeRight(  )

   for __index, expNode in pairs( self:get_expList() ) do
      if not expNode:canBeRight(  ) then
         return false
      end
      
   end
   
   return true
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

function NodeKind.get_Switch(  )

   return _lune.unwrap( _moduleObj.nodeKind['Switch'])
end


regKind( "Switch" )
function Filter:processSwitch( node, opt )

end


function NodeManager:getSwitchNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Switch']) )
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
function SwitchNode.new( id, pos, typeList, exp, caseList, default )
   local obj = {}
   SwitchNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, exp, caseList, default ); end
   return obj
end
function SwitchNode:__init(id, pos, typeList, exp, caseList, default) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Switch']), pos, typeList)
   
   
   self.exp = exp
   self.caseList = caseList
   self.default = default
   
end
function SwitchNode.create( nodeMan, pos, typeList, exp, caseList, default )

   local node = SwitchNode.new(nodeMan:nextId(  ), pos, typeList, exp, caseList, default)
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


function SwitchNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   for __index, caseInfo in pairs( self.caseList ) do
      local work = caseInfo:get_block():getBreakKind( checkMode )
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
                  return BreakKind.None
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
                     return BreakKind.None
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


function NodeKind.get_While(  )

   return _lune.unwrap( _moduleObj.nodeKind['While'])
end


regKind( "While" )
function Filter:processWhile( node, opt )

end


function NodeManager:getWhileNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['While']) )
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
function WhileNode.new( id, pos, typeList, exp, block )
   local obj = {}
   WhileNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, exp, block ); end
   return obj
end
function WhileNode:__init(id, pos, typeList, exp, block) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['While']), pos, typeList)
   
   
   self.exp = exp
   self.block = block
   
end
function WhileNode.create( nodeMan, pos, typeList, exp, block )

   local node = WhileNode.new(nodeMan:nextId(  ), pos, typeList, exp, block)
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

   return _lune.unwrap( _moduleObj.nodeKind['Repeat'])
end


regKind( "Repeat" )
function Filter:processRepeat( node, opt )

end


function NodeManager:getRepeatNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Repeat']) )
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
function RepeatNode.new( id, pos, typeList, block, exp )
   local obj = {}
   RepeatNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, block, exp ); end
   return obj
end
function RepeatNode:__init(id, pos, typeList, block, exp) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Repeat']), pos, typeList)
   
   
   self.block = block
   self.exp = exp
   
end
function RepeatNode.create( nodeMan, pos, typeList, block, exp )

   local node = RepeatNode.new(nodeMan:nextId(  ), pos, typeList, block, exp)
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

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_For(  )

   return _lune.unwrap( _moduleObj.nodeKind['For'])
end


regKind( "For" )
function Filter:processFor( node, opt )

end


function NodeManager:getForNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['For']) )
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
function ForNode.new( id, pos, typeList, block, val, init, to, delta )
   local obj = {}
   ForNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, block, val, init, to, delta ); end
   return obj
end
function ForNode:__init(id, pos, typeList, block, val, init, to, delta) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['For']), pos, typeList)
   
   
   self.block = block
   self.val = val
   self.init = init
   self.to = to
   self.delta = delta
   
end
function ForNode.create( nodeMan, pos, typeList, block, val, init, to, delta )

   local node = ForNode.new(nodeMan:nextId(  ), pos, typeList, block, val, init, to, delta)
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

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_Apply(  )

   return _lune.unwrap( _moduleObj.nodeKind['Apply'])
end


regKind( "Apply" )
function Filter:processApply( node, opt )

end


function NodeManager:getApplyNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Apply']) )
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
function ApplyNode.new( id, pos, typeList, varList, exp, block )
   local obj = {}
   ApplyNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, varList, exp, block ); end
   return obj
end
function ApplyNode:__init(id, pos, typeList, varList, exp, block) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Apply']), pos, typeList)
   
   
   self.varList = varList
   self.exp = exp
   self.block = block
   
end
function ApplyNode.create( nodeMan, pos, typeList, varList, exp, block )

   local node = ApplyNode.new(nodeMan:nextId(  ), pos, typeList, varList, exp, block)
   nodeMan:addNode( node )
   return node
end
function ApplyNode:visit( visitor, depth )

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
function ApplyNode.setmeta( obj )
  setmetatable( obj, { __index = ApplyNode  } )
end
function ApplyNode:get_varList()
   return self.varList
end
function ApplyNode:get_exp()
   return self.exp
end
function ApplyNode:get_block()
   return self.block
end


function ApplyNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_Foreach(  )

   return _lune.unwrap( _moduleObj.nodeKind['Foreach'])
end


regKind( "Foreach" )
function Filter:processForeach( node, opt )

end


function NodeManager:getForeachNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Foreach']) )
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
function ForeachNode.new( id, pos, typeList, val, key, exp, block )
   local obj = {}
   ForeachNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, val, key, exp, block ); end
   return obj
end
function ForeachNode:__init(id, pos, typeList, val, key, exp, block) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Foreach']), pos, typeList)
   
   
   self.val = val
   self.key = key
   self.exp = exp
   self.block = block
   
end
function ForeachNode.create( nodeMan, pos, typeList, val, key, exp, block )

   local node = ForeachNode.new(nodeMan:nextId(  ), pos, typeList, val, key, exp, block)
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

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_Forsort(  )

   return _lune.unwrap( _moduleObj.nodeKind['Forsort'])
end


regKind( "Forsort" )
function Filter:processForsort( node, opt )

end


function NodeManager:getForsortNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Forsort']) )
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
function ForsortNode.new( id, pos, typeList, val, key, exp, block, sort )
   local obj = {}
   ForsortNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, val, key, exp, block, sort ); end
   return obj
end
function ForsortNode:__init(id, pos, typeList, val, key, exp, block, sort) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Forsort']), pos, typeList)
   
   
   self.val = val
   self.key = key
   self.exp = exp
   self.block = block
   self.sort = sort
   
end
function ForsortNode.create( nodeMan, pos, typeList, val, key, exp, block, sort )

   local node = ForsortNode.new(nodeMan:nextId(  ), pos, typeList, val, key, exp, block, sort)
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

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_Return(  )

   return _lune.unwrap( _moduleObj.nodeKind['Return'])
end


regKind( "Return" )
function Filter:processReturn( node, opt )

end


function NodeManager:getReturnNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Return']) )
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
function ReturnNode.new( id, pos, typeList, expList )
   local obj = {}
   ReturnNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, expList ); end
   return obj
end
function ReturnNode:__init(id, pos, typeList, expList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Return']), pos, typeList)
   
   
   self.expList = expList
   
end
function ReturnNode.create( nodeMan, pos, typeList, expList )

   local node = ReturnNode.new(nodeMan:nextId(  ), pos, typeList, expList)
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

   return _lune.unwrap( _moduleObj.nodeKind['Break'])
end


regKind( "Break" )
function Filter:processBreak( node, opt )

end


function NodeManager:getBreakNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Break']) )
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
function BreakNode.new( id, pos, typeList )
   local obj = {}
   BreakNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList ); end
   return obj
end
function BreakNode:__init(id, pos, typeList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Break']), pos, typeList)
   
   
   
end
function BreakNode.create( nodeMan, pos, typeList )

   local node = BreakNode.new(nodeMan:nextId(  ), pos, typeList)
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

   return _lune.unwrap( _moduleObj.nodeKind['Provide'])
end


regKind( "Provide" )
function Filter:processProvide( node, opt )

end


function NodeManager:getProvideNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Provide']) )
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
function ProvideNode.new( id, pos, typeList, symbol )
   local obj = {}
   ProvideNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, symbol ); end
   return obj
end
function ProvideNode:__init(id, pos, typeList, symbol) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Provide']), pos, typeList)
   
   
   self.symbol = symbol
   
end
function ProvideNode.create( nodeMan, pos, typeList, symbol )

   local node = ProvideNode.new(nodeMan:nextId(  ), pos, typeList, symbol)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpNew'])
end


regKind( "ExpNew" )
function Filter:processExpNew( node, opt )

end


function NodeManager:getExpNewNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpNew']) )
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
function ExpNewNode.new( id, pos, typeList, symbol, argList )
   local obj = {}
   ExpNewNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, symbol, argList ); end
   return obj
end
function ExpNewNode:__init(id, pos, typeList, symbol, argList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpNew']), pos, typeList)
   
   
   self.symbol = symbol
   self.argList = argList
   
end
function ExpNewNode.create( nodeMan, pos, typeList, symbol, argList )

   local node = ExpNewNode.new(nodeMan:nextId(  ), pos, typeList, symbol, argList)
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
function ExpNewNode:get_argList()
   return self.argList
end


function NodeKind.get_ExpUnwrap(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpUnwrap'])
end


regKind( "ExpUnwrap" )
function Filter:processExpUnwrap( node, opt )

end


function NodeManager:getExpUnwrapNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpUnwrap']) )
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
function ExpUnwrapNode.new( id, pos, typeList, exp, default )
   local obj = {}
   ExpUnwrapNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, exp, default ); end
   return obj
end
function ExpUnwrapNode:__init(id, pos, typeList, exp, default) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpUnwrap']), pos, typeList)
   
   
   self.exp = exp
   self.default = default
   
end
function ExpUnwrapNode.create( nodeMan, pos, typeList, exp, default )

   local node = ExpUnwrapNode.new(nodeMan:nextId(  ), pos, typeList, exp, default)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpRef'])
end


regKind( "ExpRef" )
function Filter:processExpRef( node, opt )

end


function NodeManager:getExpRefNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpRef']) )
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
function ExpRefNode.new( id, pos, typeList, token, symbolInfo )
   local obj = {}
   ExpRefNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, token, symbolInfo ); end
   return obj
end
function ExpRefNode:__init(id, pos, typeList, token, symbolInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpRef']), pos, typeList)
   
   
   self.token = token
   self.symbolInfo = symbolInfo
   
end
function ExpRefNode.create( nodeMan, pos, typeList, token, symbolInfo )

   local node = ExpRefNode.new(nodeMan:nextId(  ), pos, typeList, token, symbolInfo)
   nodeMan:addNode( node )
   return node
end
function ExpRefNode:visit( visitor, depth )

   
   return true
end
function ExpRefNode.setmeta( obj )
  setmetatable( obj, { __index = ExpRefNode  } )
end
function ExpRefNode:get_token()
   return self.token
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

function NodeKind.get_ExpOp2(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpOp2'])
end


regKind( "ExpOp2" )
function Filter:processExpOp2( node, opt )

end


function NodeManager:getExpOp2NodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpOp2']) )
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
function ExpOp2Node.new( id, pos, typeList, op, exp1, exp2 )
   local obj = {}
   ExpOp2Node.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, op, exp1, exp2 ); end
   return obj
end
function ExpOp2Node:__init(id, pos, typeList, op, exp1, exp2) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpOp2']), pos, typeList)
   
   
   self.op = op
   self.exp1 = exp1
   self.exp2 = exp2
   
end
function ExpOp2Node.create( nodeMan, pos, typeList, op, exp1, exp2 )

   local node = ExpOp2Node.new(nodeMan:nextId(  ), pos, typeList, op, exp1, exp2)
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


function ExpOp2Node:canBeStatement(  )

   return self:get_op().txt == '='
end

function NodeKind.get_UnwrapSet(  )

   return _lune.unwrap( _moduleObj.nodeKind['UnwrapSet'])
end


regKind( "UnwrapSet" )
function Filter:processUnwrapSet( node, opt )

end


function NodeManager:getUnwrapSetNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['UnwrapSet']) )
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
function UnwrapSetNode.new( id, pos, typeList, dstExpList, srcExpList, unwrapBlock )
   local obj = {}
   UnwrapSetNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, dstExpList, srcExpList, unwrapBlock ); end
   return obj
end
function UnwrapSetNode:__init(id, pos, typeList, dstExpList, srcExpList, unwrapBlock) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['UnwrapSet']), pos, typeList)
   
   
   self.dstExpList = dstExpList
   self.srcExpList = srcExpList
   self.unwrapBlock = unwrapBlock
   
end
function UnwrapSetNode.create( nodeMan, pos, typeList, dstExpList, srcExpList, unwrapBlock )

   local node = UnwrapSetNode.new(nodeMan:nextId(  ), pos, typeList, dstExpList, srcExpList, unwrapBlock)
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

   return _lune.unwrap( _moduleObj.nodeKind['IfUnwrap'])
end


regKind( "IfUnwrap" )
function Filter:processIfUnwrap( node, opt )

end


function NodeManager:getIfUnwrapNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['IfUnwrap']) )
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
function IfUnwrapNode.new( id, pos, typeList, varNameList, expList, block, nilBlock )
   local obj = {}
   IfUnwrapNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, varNameList, expList, block, nilBlock ); end
   return obj
end
function IfUnwrapNode:__init(id, pos, typeList, varNameList, expList, block, nilBlock) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['IfUnwrap']), pos, typeList)
   
   
   self.varNameList = varNameList
   self.expList = expList
   self.block = block
   self.nilBlock = nilBlock
   
end
function IfUnwrapNode.create( nodeMan, pos, typeList, varNameList, expList, block, nilBlock )

   local node = IfUnwrapNode.new(nodeMan:nextId(  ), pos, typeList, varNameList, expList, block, nilBlock)
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
function IfUnwrapNode:get_varNameList()
   return self.varNameList
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
               return BreakKind.None
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
                     return BreakKind.None
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

function NodeKind.get_When(  )

   return _lune.unwrap( _moduleObj.nodeKind['When'])
end


regKind( "When" )
function Filter:processWhen( node, opt )

end


function NodeManager:getWhenNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['When']) )
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
function WhenNode.new( id, pos, typeList, varNameList, expNodeList, block, elseBlock )
   local obj = {}
   WhenNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, varNameList, expNodeList, block, elseBlock ); end
   return obj
end
function WhenNode:__init(id, pos, typeList, varNameList, expNodeList, block, elseBlock) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['When']), pos, typeList)
   
   
   self.varNameList = varNameList
   self.expNodeList = expNodeList
   self.block = block
   self.elseBlock = elseBlock
   
end
function WhenNode.create( nodeMan, pos, typeList, varNameList, expNodeList, block, elseBlock )

   local node = WhenNode.new(nodeMan:nextId(  ), pos, typeList, varNameList, expNodeList, block, elseBlock)
   nodeMan:addNode( node )
   return node
end
function WhenNode:visit( visitor, depth )

   do
      local list = self.expNodeList
      for __index, child in pairs( list ) do
         do
            local _switchExp = visitor( child, self, 'expNodeList', depth )
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
function WhenNode:get_varNameList()
   return self.varNameList
end
function WhenNode:get_expNodeList()
   return self.expNodeList
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
               return BreakKind.None
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
                     return BreakKind.None
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpCast'])
end


regKind( "ExpCast" )
function Filter:processExpCast( node, opt )

end


function NodeManager:getExpCastNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpCast']) )
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
function ExpCastNode.new( id, pos, typeList, exp, castType, castKind )
   local obj = {}
   ExpCastNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, exp, castType, castKind ); end
   return obj
end
function ExpCastNode:__init(id, pos, typeList, exp, castType, castKind) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpCast']), pos, typeList)
   
   
   self.exp = exp
   self.castType = castType
   self.castKind = castKind
   
end
function ExpCastNode.create( nodeMan, pos, typeList, exp, castType, castKind )

   local node = ExpCastNode.new(nodeMan:nextId(  ), pos, typeList, exp, castType, castKind)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpToDDD'])
end


regKind( "ExpToDDD" )
function Filter:processExpToDDD( node, opt )

end


function NodeManager:getExpToDDDNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpToDDD']) )
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
function ExpToDDDNode.new( id, pos, typeList, expList )
   local obj = {}
   ExpToDDDNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, expList ); end
   return obj
end
function ExpToDDDNode:__init(id, pos, typeList, expList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpToDDD']), pos, typeList)
   
   
   self.expList = expList
   
end
function ExpToDDDNode.create( nodeMan, pos, typeList, expList )

   local node = ExpToDDDNode.new(nodeMan:nextId(  ), pos, typeList, expList)
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
MacroMode.Analyze = 2
MacroMode._val2NameMap[2] = 'Analyze'
MacroMode.__allList[3] = MacroMode.Analyze

function NodeKind.get_ExpOp1(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpOp1'])
end


regKind( "ExpOp1" )
function Filter:processExpOp1( node, opt )

end


function NodeManager:getExpOp1NodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpOp1']) )
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
function ExpOp1Node.new( id, pos, typeList, op, macroMode, exp )
   local obj = {}
   ExpOp1Node.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, op, macroMode, exp ); end
   return obj
end
function ExpOp1Node:__init(id, pos, typeList, op, macroMode, exp) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpOp1']), pos, typeList)
   
   
   self.op = op
   self.macroMode = macroMode
   self.exp = exp
   
end
function ExpOp1Node.create( nodeMan, pos, typeList, op, macroMode, exp )

   local node = ExpOp1Node.new(nodeMan:nextId(  ), pos, typeList, op, macroMode, exp)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpRefItem'])
end


regKind( "ExpRefItem" )
function Filter:processExpRefItem( node, opt )

end


function NodeManager:getExpRefItemNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpRefItem']) )
end


local ExpRefItemNode = {}
setmetatable( ExpRefItemNode, { __index = Node } )
_moduleObj.ExpRefItemNode = ExpRefItemNode
function ExpRefItemNode:processFilter( filter, opt )

   filter:processExpRefItem( self, opt )
end
function ExpRefItemNode:canBeRight(  )

   return true
end
function ExpRefItemNode:canBeStatement(  )

   return false
end
function ExpRefItemNode.new( id, pos, typeList, val, nilAccess, symbol, index )
   local obj = {}
   ExpRefItemNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, val, nilAccess, symbol, index ); end
   return obj
end
function ExpRefItemNode:__init(id, pos, typeList, val, nilAccess, symbol, index) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpRefItem']), pos, typeList)
   
   
   self.val = val
   self.nilAccess = nilAccess
   self.symbol = symbol
   self.index = index
   
end
function ExpRefItemNode.create( nodeMan, pos, typeList, val, nilAccess, symbol, index )

   local node = ExpRefItemNode.new(nodeMan:nextId(  ), pos, typeList, val, nilAccess, symbol, index)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpCall'])
end


regKind( "ExpCall" )
function Filter:processExpCall( node, opt )

end


function NodeManager:getExpCallNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpCall']) )
end


local ExpCallNode = {}
setmetatable( ExpCallNode, { __index = Node } )
_moduleObj.ExpCallNode = ExpCallNode
function ExpCallNode:processFilter( filter, opt )

   filter:processExpCall( self, opt )
end
function ExpCallNode:canBeLeft(  )

   return false
end
function ExpCallNode:canBeStatement(  )

   return true
end
function ExpCallNode.new( id, pos, typeList, func, errorFunc, nilAccess, argList )
   local obj = {}
   ExpCallNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, func, errorFunc, nilAccess, argList ); end
   return obj
end
function ExpCallNode:__init(id, pos, typeList, func, errorFunc, nilAccess, argList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpCall']), pos, typeList)
   
   
   self.func = func
   self.errorFunc = errorFunc
   self.nilAccess = nilAccess
   self.argList = argList
   
end
function ExpCallNode.create( nodeMan, pos, typeList, func, errorFunc, nilAccess, argList )

   local node = ExpCallNode.new(nodeMan:nextId(  ), pos, typeList, func, errorFunc, nilAccess, argList)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpAccessMRet'])
end


regKind( "ExpAccessMRet" )
function Filter:processExpAccessMRet( node, opt )

end


function NodeManager:getExpAccessMRetNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpAccessMRet']) )
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
function ExpAccessMRetNode.new( id, pos, typeList, mRet, index )
   local obj = {}
   ExpAccessMRetNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, mRet, index ); end
   return obj
end
function ExpAccessMRetNode:__init(id, pos, typeList, mRet, index) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpAccessMRet']), pos, typeList)
   
   
   self.mRet = mRet
   self.index = index
   
end
function ExpAccessMRetNode.create( nodeMan, pos, typeList, mRet, index )

   local node = ExpAccessMRetNode.new(nodeMan:nextId(  ), pos, typeList, mRet, index)
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


function NodeKind.get_ExpDDD(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpDDD'])
end


regKind( "ExpDDD" )
function Filter:processExpDDD( node, opt )

end


function NodeManager:getExpDDDNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpDDD']) )
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
function ExpDDDNode.new( id, pos, typeList, token )
   local obj = {}
   ExpDDDNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, token ); end
   return obj
end
function ExpDDDNode:__init(id, pos, typeList, token) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpDDD']), pos, typeList)
   
   
   self.token = token
   
end
function ExpDDDNode.create( nodeMan, pos, typeList, token )

   local node = ExpDDDNode.new(nodeMan:nextId(  ), pos, typeList, token)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpParen'])
end


regKind( "ExpParen" )
function Filter:processExpParen( node, opt )

end


function NodeManager:getExpParenNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpParen']) )
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
function ExpParenNode.new( id, pos, typeList, exp )
   local obj = {}
   ExpParenNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, exp ); end
   return obj
end
function ExpParenNode:__init(id, pos, typeList, exp) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpParen']), pos, typeList)
   
   
   self.exp = exp
   
end
function ExpParenNode.create( nodeMan, pos, typeList, exp )

   local node = ExpParenNode.new(nodeMan:nextId(  ), pos, typeList, exp)
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


function NodeKind.get_ExpMacroExp(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpMacroExp'])
end


regKind( "ExpMacroExp" )
function Filter:processExpMacroExp( node, opt )

end


function NodeManager:getExpMacroExpNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpMacroExp']) )
end


local ExpMacroExpNode = {}
setmetatable( ExpMacroExpNode, { __index = Node } )
_moduleObj.ExpMacroExpNode = ExpMacroExpNode
function ExpMacroExpNode:processFilter( filter, opt )

   filter:processExpMacroExp( self, opt )
end
function ExpMacroExpNode:canBeRight(  )

   return false
end
function ExpMacroExpNode:canBeLeft(  )

   return false
end
function ExpMacroExpNode:canBeStatement(  )

   return true
end
function ExpMacroExpNode.new( id, pos, typeList, stmtList )
   local obj = {}
   ExpMacroExpNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, stmtList ); end
   return obj
end
function ExpMacroExpNode:__init(id, pos, typeList, stmtList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpMacroExp']), pos, typeList)
   
   
   self.stmtList = stmtList
   
end
function ExpMacroExpNode.create( nodeMan, pos, typeList, stmtList )

   local node = ExpMacroExpNode.new(nodeMan:nextId(  ), pos, typeList, stmtList)
   nodeMan:addNode( node )
   return node
end
function ExpMacroExpNode:visit( visitor, depth )

   do
      local list = self.stmtList
      for __index, child in pairs( list ) do
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


function ExpMacroExpNode:getBreakKind( checkMode )

   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      local kind = BreakKind.None
      for __index, stmt in pairs( self.stmtList ) do
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
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
      end
      
      return kind
   else
    
      if #self.stmtList > 0 then
         return self.stmtList[#self.stmtList]:getBreakKind( checkMode )
      end
      
   end
   
   return BreakKind.None
end

function NodeKind.get_ExpMacroStat(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpMacroStat'])
end


regKind( "ExpMacroStat" )
function Filter:processExpMacroStat( node, opt )

end


function NodeManager:getExpMacroStatNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpMacroStat']) )
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
function ExpMacroStatNode.new( id, pos, typeList, expStrList )
   local obj = {}
   ExpMacroStatNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, expStrList ); end
   return obj
end
function ExpMacroStatNode:__init(id, pos, typeList, expStrList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpMacroStat']), pos, typeList)
   
   
   self.expStrList = expStrList
   
end
function ExpMacroStatNode.create( nodeMan, pos, typeList, expStrList )

   local node = ExpMacroStatNode.new(nodeMan:nextId(  ), pos, typeList, expStrList)
   nodeMan:addNode( node )
   return node
end
function ExpMacroStatNode:visit( visitor, depth )

   do
      local list = self.expStrList
      for __index, child in pairs( list ) do
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


function NodeKind.get_StmtExp(  )

   return _lune.unwrap( _moduleObj.nodeKind['StmtExp'])
end


regKind( "StmtExp" )
function Filter:processStmtExp( node, opt )

end


function NodeManager:getStmtExpNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['StmtExp']) )
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
function StmtExpNode.new( id, pos, typeList, exp )
   local obj = {}
   StmtExpNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, exp ); end
   return obj
end
function StmtExpNode:__init(id, pos, typeList, exp) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['StmtExp']), pos, typeList)
   
   
   self.exp = exp
   
end
function StmtExpNode.create( nodeMan, pos, typeList, exp )

   local node = StmtExpNode.new(nodeMan:nextId(  ), pos, typeList, exp)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpMacroStatList'])
end


regKind( "ExpMacroStatList" )
function Filter:processExpMacroStatList( node, opt )

end


function NodeManager:getExpMacroStatListNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpMacroStatList']) )
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
function ExpMacroStatListNode.new( id, pos, typeList, exp )
   local obj = {}
   ExpMacroStatListNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, exp ); end
   return obj
end
function ExpMacroStatListNode:__init(id, pos, typeList, exp) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpMacroStatList']), pos, typeList)
   
   
   self.exp = exp
   
end
function ExpMacroStatListNode.create( nodeMan, pos, typeList, exp )

   local node = ExpMacroStatListNode.new(nodeMan:nextId(  ), pos, typeList, exp)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpOmitEnum'])
end


regKind( "ExpOmitEnum" )
function Filter:processExpOmitEnum( node, opt )

end


function NodeManager:getExpOmitEnumNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpOmitEnum']) )
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
function ExpOmitEnumNode.new( id, pos, typeList, valToken, valInfo, enumTypeInfo )
   local obj = {}
   ExpOmitEnumNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, valToken, valInfo, enumTypeInfo ); end
   return obj
end
function ExpOmitEnumNode:__init(id, pos, typeList, valToken, valInfo, enumTypeInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpOmitEnum']), pos, typeList)
   
   
   self.valToken = valToken
   self.valInfo = valInfo
   self.enumTypeInfo = enumTypeInfo
   
end
function ExpOmitEnumNode.create( nodeMan, pos, typeList, valToken, valInfo, enumTypeInfo )

   local node = ExpOmitEnumNode.new(nodeMan:nextId(  ), pos, typeList, valToken, valInfo, enumTypeInfo)
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

   return _lune.unwrap( _moduleObj.nodeKind['RefField'])
end


regKind( "RefField" )
function Filter:processRefField( node, opt )

end


function NodeManager:getRefFieldNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['RefField']) )
end


local RefFieldNode = {}
setmetatable( RefFieldNode, { __index = Node } )
_moduleObj.RefFieldNode = RefFieldNode
function RefFieldNode:processFilter( filter, opt )

   filter:processRefField( self, opt )
end
function RefFieldNode:canBeStatement(  )

   return false
end
function RefFieldNode.new( id, pos, typeList, field, symbolInfo, nilAccess, prefix )
   local obj = {}
   RefFieldNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, field, symbolInfo, nilAccess, prefix ); end
   return obj
end
function RefFieldNode:__init(id, pos, typeList, field, symbolInfo, nilAccess, prefix) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['RefField']), pos, typeList)
   
   
   self.field = field
   self.symbolInfo = symbolInfo
   self.nilAccess = nilAccess
   self.prefix = prefix
   
end
function RefFieldNode.create( nodeMan, pos, typeList, field, symbolInfo, nilAccess, prefix )

   local node = RefFieldNode.new(nodeMan:nextId(  ), pos, typeList, field, symbolInfo, nilAccess, prefix)
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

   return _lune.unwrap( _moduleObj.nodeKind['GetField'])
end


regKind( "GetField" )
function Filter:processGetField( node, opt )

end


function NodeManager:getGetFieldNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['GetField']) )
end


local GetFieldNode = {}
setmetatable( GetFieldNode, { __index = Node } )
_moduleObj.GetFieldNode = GetFieldNode
function GetFieldNode:processFilter( filter, opt )

   filter:processGetField( self, opt )
end
function GetFieldNode:canBeRight(  )

   return true
end
function GetFieldNode:canBeStatement(  )

   return false
end
function GetFieldNode.new( id, pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo )
   local obj = {}
   GetFieldNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo ); end
   return obj
end
function GetFieldNode:__init(id, pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['GetField']), pos, typeList)
   
   
   self.field = field
   self.symbolInfo = symbolInfo
   self.nilAccess = nilAccess
   self.prefix = prefix
   self.getterTypeInfo = getterTypeInfo
   
end
function GetFieldNode.create( nodeMan, pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo )

   local node = GetFieldNode.new(nodeMan:nextId(  ), pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo)
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

   return _lune.unwrap( _moduleObj.nodeKind['Alias'])
end


regKind( "Alias" )
function Filter:processAlias( node, opt )

end


function NodeManager:getAliasNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Alias']) )
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
function AliasNode.new( id, pos, typeList, newName, srcNode, typeInfo )
   local obj = {}
   AliasNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, newName, srcNode, typeInfo ); end
   return obj
end
function AliasNode:__init(id, pos, typeList, newName, srcNode, typeInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Alias']), pos, typeList)
   
   
   self.newName = newName
   self.srcNode = srcNode
   self.typeInfo = typeInfo
   
end
function AliasNode.create( nodeMan, pos, typeList, newName, srcNode, typeInfo )

   local node = AliasNode.new(nodeMan:nextId(  ), pos, typeList, newName, srcNode, typeInfo)
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

   return _lune.unwrap( _moduleObj.nodeKind['DeclVar'])
end


regKind( "DeclVar" )
function Filter:processDeclVar( node, opt )

end


function NodeManager:getDeclVarNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclVar']) )
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
function DeclVarNode.new( id, pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock )
   local obj = {}
   DeclVarNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock ); end
   return obj
end
function DeclVarNode:__init(id, pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclVar']), pos, typeList)
   
   
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
function DeclVarNode.create( nodeMan, pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock )

   local node = DeclVarNode.new(nodeMan:nextId(  ), pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock)
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
                     return BreakKind.None
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
                           return BreakKind.None
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
                                 return BreakKind.None
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

   return _lune.unwrap( _moduleObj.nodeKind['DeclForm'])
end


regKind( "DeclForm" )
function Filter:processDeclForm( node, opt )

end


function NodeManager:getDeclFormNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclForm']) )
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
function DeclFormNode.new( id, pos, typeList, argList )
   local obj = {}
   DeclFormNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, argList ); end
   return obj
end
function DeclFormNode:__init(id, pos, typeList, argList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclForm']), pos, typeList)
   
   
   self.argList = argList
   
end
function DeclFormNode.create( nodeMan, pos, typeList, argList )

   local node = DeclFormNode.new(nodeMan:nextId(  ), pos, typeList, argList)
   nodeMan:addNode( node )
   return node
end
function DeclFormNode:visit( visitor, depth )

   do
      local list = self.argList
      for __index, child in pairs( list ) do
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


local DeclFuncInfo = {}
_moduleObj.DeclFuncInfo = DeclFuncInfo
function DeclFuncInfo.setmeta( obj )
  setmetatable( obj, { __index = DeclFuncInfo  } )
end
function DeclFuncInfo.new( classTypeInfo, name, argList, staticFlag, accessMode, body, retTypeInfoList, has__func__Symbol )
   local obj = {}
   DeclFuncInfo.setmeta( obj )
   if obj.__init then
      obj:__init( classTypeInfo, name, argList, staticFlag, accessMode, body, retTypeInfoList, has__func__Symbol )
   end
   return obj
end
function DeclFuncInfo:__init( classTypeInfo, name, argList, staticFlag, accessMode, body, retTypeInfoList, has__func__Symbol )

   self.classTypeInfo = classTypeInfo
   self.name = name
   self.argList = argList
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.body = body
   self.retTypeInfoList = retTypeInfoList
   self.has__func__Symbol = has__func__Symbol
end
function DeclFuncInfo:get_classTypeInfo()
   return self.classTypeInfo
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

function NodeKind.get_DeclFunc(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclFunc'])
end


regKind( "DeclFunc" )
function Filter:processDeclFunc( node, opt )

end


function NodeManager:getDeclFuncNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclFunc']) )
end


local DeclFuncNode = {}
setmetatable( DeclFuncNode, { __index = Node } )
_moduleObj.DeclFuncNode = DeclFuncNode
function DeclFuncNode:processFilter( filter, opt )

   filter:processDeclFunc( self, opt )
end
function DeclFuncNode:canBeRight(  )

   return true
end
function DeclFuncNode:canBeLeft(  )

   return false
end
function DeclFuncNode:canBeStatement(  )

   return true
end
function DeclFuncNode.new( id, pos, typeList, declInfo )
   local obj = {}
   DeclFuncNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, declInfo ); end
   return obj
end
function DeclFuncNode:__init(id, pos, typeList, declInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclFunc']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclFuncNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclFuncNode.new(nodeMan:nextId(  ), pos, typeList, declInfo)
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


function NodeKind.get_DeclMethod(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclMethod'])
end


regKind( "DeclMethod" )
function Filter:processDeclMethod( node, opt )

end


function NodeManager:getDeclMethodNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclMethod']) )
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
function DeclMethodNode.new( id, pos, typeList, declInfo )
   local obj = {}
   DeclMethodNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, declInfo ); end
   return obj
end
function DeclMethodNode:__init(id, pos, typeList, declInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclMethod']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclMethodNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclMethodNode.new(nodeMan:nextId(  ), pos, typeList, declInfo)
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


function NodeKind.get_DeclConstr(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclConstr'])
end


regKind( "DeclConstr" )
function Filter:processDeclConstr( node, opt )

end


function NodeManager:getDeclConstrNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclConstr']) )
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
function DeclConstrNode.new( id, pos, typeList, declInfo )
   local obj = {}
   DeclConstrNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, declInfo ); end
   return obj
end
function DeclConstrNode:__init(id, pos, typeList, declInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclConstr']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclConstrNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclConstrNode.new(nodeMan:nextId(  ), pos, typeList, declInfo)
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

   return _lune.unwrap( _moduleObj.nodeKind['DeclDestr'])
end


regKind( "DeclDestr" )
function Filter:processDeclDestr( node, opt )

end


function NodeManager:getDeclDestrNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclDestr']) )
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
function DeclDestrNode.new( id, pos, typeList, declInfo )
   local obj = {}
   DeclDestrNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, declInfo ); end
   return obj
end
function DeclDestrNode:__init(id, pos, typeList, declInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclDestr']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclDestrNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclDestrNode.new(nodeMan:nextId(  ), pos, typeList, declInfo)
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

   return _lune.unwrap( _moduleObj.nodeKind['ExpCallSuper'])
end


regKind( "ExpCallSuper" )
function Filter:processExpCallSuper( node, opt )

end


function NodeManager:getExpCallSuperNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpCallSuper']) )
end


local ExpCallSuperNode = {}
setmetatable( ExpCallSuperNode, { __index = Node } )
_moduleObj.ExpCallSuperNode = ExpCallSuperNode
function ExpCallSuperNode:processFilter( filter, opt )

   filter:processExpCallSuper( self, opt )
end
function ExpCallSuperNode:canBeRight(  )

   return false
end
function ExpCallSuperNode:canBeLeft(  )

   return false
end
function ExpCallSuperNode:canBeStatement(  )

   return true
end
function ExpCallSuperNode.new( id, pos, typeList, superType, methodType, expList )
   local obj = {}
   ExpCallSuperNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, superType, methodType, expList ); end
   return obj
end
function ExpCallSuperNode:__init(id, pos, typeList, superType, methodType, expList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['ExpCallSuper']), pos, typeList)
   
   
   self.superType = superType
   self.methodType = methodType
   self.expList = expList
   
end
function ExpCallSuperNode.create( nodeMan, pos, typeList, superType, methodType, expList )

   local node = ExpCallSuperNode.new(nodeMan:nextId(  ), pos, typeList, superType, methodType, expList)
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


function NodeKind.get_DeclMember(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclMember'])
end


regKind( "DeclMember" )
function Filter:processDeclMember( node, opt )

end


function NodeManager:getDeclMemberNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclMember']) )
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
function DeclMemberNode.new( id, pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode )
   local obj = {}
   DeclMemberNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode ); end
   return obj
end
function DeclMemberNode:__init(id, pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclMember']), pos, typeList)
   
   
   self.name = name
   self.refType = refType
   self.symbolInfo = symbolInfo
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.getterMutable = getterMutable
   self.getterMode = getterMode
   self.getterRetType = getterRetType
   self.setterMode = setterMode
   
end
function DeclMemberNode.create( nodeMan, pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode )

   local node = DeclMemberNode.new(nodeMan:nextId(  ), pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode)
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

   return _lune.unwrap( _moduleObj.nodeKind['DeclArg'])
end


regKind( "DeclArg" )
function Filter:processDeclArg( node, opt )

end


function NodeManager:getDeclArgNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclArg']) )
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
function DeclArgNode.new( id, pos, typeList, name, argType )
   local obj = {}
   DeclArgNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, name, argType ); end
   return obj
end
function DeclArgNode:__init(id, pos, typeList, name, argType) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclArg']), pos, typeList)
   
   
   self.name = name
   self.argType = argType
   
end
function DeclArgNode.create( nodeMan, pos, typeList, name, argType )

   local node = DeclArgNode.new(nodeMan:nextId(  ), pos, typeList, name, argType)
   nodeMan:addNode( node )
   return node
end
function DeclArgNode:visit( visitor, depth )

   do
      local child = self.argType
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
   
   
   return true
end
function DeclArgNode.setmeta( obj )
  setmetatable( obj, { __index = DeclArgNode  } )
end
function DeclArgNode:get_name()
   return self.name
end
function DeclArgNode:get_argType()
   return self.argType
end


function NodeKind.get_DeclArgDDD(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclArgDDD'])
end


regKind( "DeclArgDDD" )
function Filter:processDeclArgDDD( node, opt )

end


function NodeManager:getDeclArgDDDNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclArgDDD']) )
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
function DeclArgDDDNode.new( id, pos, typeList )
   local obj = {}
   DeclArgDDDNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList ); end
   return obj
end
function DeclArgDDDNode:__init(id, pos, typeList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclArgDDD']), pos, typeList)
   
   
   
end
function DeclArgDDDNode.create( nodeMan, pos, typeList )

   local node = DeclArgDDDNode.new(nodeMan:nextId(  ), pos, typeList)
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
function AdvertiseInfo.new( member, prefix )
   local obj = {}
   AdvertiseInfo.setmeta( obj )
   if obj.__init then
      obj:__init( member, prefix )
   end
   return obj
end
function AdvertiseInfo:__init( member, prefix )

   self.member = member
   self.prefix = prefix
end
function AdvertiseInfo:get_member()
   return self.member
end
function AdvertiseInfo:get_prefix()
   return self.prefix
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

   return _lune.unwrap( _moduleObj.nodeKind['DeclClass'])
end


regKind( "DeclClass" )
function Filter:processDeclClass( node, opt )

end


function NodeManager:getDeclClassNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclClass']) )
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
function DeclClassNode.new( id, pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initBlock, advertiseList, trustList, outerMethodSet )
   local obj = {}
   DeclClassNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initBlock, advertiseList, trustList, outerMethodSet ); end
   return obj
end
function DeclClassNode:__init(id, pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initBlock, advertiseList, trustList, outerMethodSet) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclClass']), pos, typeList)
   
   
   self.accessMode = accessMode
   self.name = name
   self.gluePrefix = gluePrefix
   self.declStmtList = declStmtList
   self.fieldList = fieldList
   self.moduleName = moduleName
   self.memberList = memberList
   self.scope = scope
   self.initBlock = initBlock
   self.advertiseList = advertiseList
   self.trustList = trustList
   self.outerMethodSet = outerMethodSet
   
end
function DeclClassNode.create( nodeMan, pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initBlock, advertiseList, trustList, outerMethodSet )

   local node = DeclClassNode.new(nodeMan:nextId(  ), pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initBlock, advertiseList, trustList, outerMethodSet)
   nodeMan:addNode( node )
   return node
end
function DeclClassNode:visit( visitor, depth )

   do
      local list = self.declStmtList
      for __index, child in pairs( list ) do
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
      for __index, child in pairs( list ) do
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
      for __index, child in pairs( list ) do
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
function DeclClassNode:get_declStmtList()
   return self.declStmtList
end
function DeclClassNode:get_fieldList()
   return self.fieldList
end
function DeclClassNode:get_moduleName()
   return self.moduleName
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
function DeclClassNode:get_outerMethodSet()
   return self.outerMethodSet
end


function DeclClassNode:hasUserInit(  )

   local scope = _lune.unwrap( self:get_expType():get_scope())
   local initFuncType = _lune.unwrap( scope:getTypeInfoField( "__init", true, scope ))
   return not initFuncType:get_autoFlag()
end

function NodeKind.get_DeclEnum(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclEnum'])
end


regKind( "DeclEnum" )
function Filter:processDeclEnum( node, opt )

end


function NodeManager:getDeclEnumNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclEnum']) )
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
function DeclEnumNode.new( id, pos, typeList, accessMode, name, valueNameList, scope )
   local obj = {}
   DeclEnumNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, accessMode, name, valueNameList, scope ); end
   return obj
end
function DeclEnumNode:__init(id, pos, typeList, accessMode, name, valueNameList, scope) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclEnum']), pos, typeList)
   
   
   self.accessMode = accessMode
   self.name = name
   self.valueNameList = valueNameList
   self.scope = scope
   
end
function DeclEnumNode.create( nodeMan, pos, typeList, accessMode, name, valueNameList, scope )

   local node = DeclEnumNode.new(nodeMan:nextId(  ), pos, typeList, accessMode, name, valueNameList, scope)
   nodeMan:addNode( node )
   return node
end
function DeclEnumNode:visit( visitor, depth )

   
   return true
end
function DeclEnumNode.setmeta( obj )
  setmetatable( obj, { __index = DeclEnumNode  } )
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

   return _lune.unwrap( _moduleObj.nodeKind['DeclAlge'])
end


regKind( "DeclAlge" )
function Filter:processDeclAlge( node, opt )

end


function NodeManager:getDeclAlgeNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclAlge']) )
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
function DeclAlgeNode.new( id, pos, typeList, accessMode, algeType, scope )
   local obj = {}
   DeclAlgeNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, accessMode, algeType, scope ); end
   return obj
end
function DeclAlgeNode:__init(id, pos, typeList, accessMode, algeType, scope) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclAlge']), pos, typeList)
   
   
   self.accessMode = accessMode
   self.algeType = algeType
   self.scope = scope
   
end
function DeclAlgeNode.create( nodeMan, pos, typeList, accessMode, algeType, scope )

   local node = DeclAlgeNode.new(nodeMan:nextId(  ), pos, typeList, accessMode, algeType, scope)
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

   return _lune.unwrap( _moduleObj.nodeKind['NewAlgeVal'])
end


regKind( "NewAlgeVal" )
function Filter:processNewAlgeVal( node, opt )

end


function NodeManager:getNewAlgeValNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['NewAlgeVal']) )
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
function NewAlgeValNode.new( id, pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList )
   local obj = {}
   NewAlgeValNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList ); end
   return obj
end
function NewAlgeValNode:__init(id, pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['NewAlgeVal']), pos, typeList)
   
   
   self.name = name
   self.prefix = prefix
   self.algeTypeInfo = algeTypeInfo
   self.valInfo = valInfo
   self.paramList = paramList
   
end
function NewAlgeValNode.create( nodeMan, pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList )

   local node = NewAlgeValNode.new(nodeMan:nextId(  ), pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList)
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
      for __index, child in pairs( list ) do
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

   return _lune.unwrap( _moduleObj.nodeKind['LuneControl'])
end


regKind( "LuneControl" )
function Filter:processLuneControl( node, opt )

end


function NodeManager:getLuneControlNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LuneControl']) )
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
function LuneControlNode.new( id, pos, typeList, pragma )
   local obj = {}
   LuneControlNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, pragma ); end
   return obj
end
function LuneControlNode:__init(id, pos, typeList, pragma) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LuneControl']), pos, typeList)
   
   
   self.pragma = pragma
   
end
function LuneControlNode.create( nodeMan, pos, typeList, pragma )

   local node = LuneControlNode.new(nodeMan:nextId(  ), pos, typeList, pragma)
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

   return _lune.unwrap( _moduleObj.nodeKind['Match'])
end


regKind( "Match" )
function Filter:processMatch( node, opt )

end


function NodeManager:getMatchNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Match']) )
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
function MatchNode.new( id, pos, typeList, val, algeTypeInfo, caseList, defaultBlock )
   local obj = {}
   MatchNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, val, algeTypeInfo, caseList, defaultBlock ); end
   return obj
end
function MatchNode:__init(id, pos, typeList, val, algeTypeInfo, caseList, defaultBlock) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Match']), pos, typeList)
   
   
   self.val = val
   self.algeTypeInfo = algeTypeInfo
   self.caseList = caseList
   self.defaultBlock = defaultBlock
   
end
function MatchNode.create( nodeMan, pos, typeList, val, algeTypeInfo, caseList, defaultBlock )

   local node = MatchNode.new(nodeMan:nextId(  ), pos, typeList, val, algeTypeInfo, caseList, defaultBlock)
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


function NodeKind.get_LuneKind(  )

   return _lune.unwrap( _moduleObj.nodeKind['LuneKind'])
end


regKind( "LuneKind" )
function Filter:processLuneKind( node, opt )

end


function NodeManager:getLuneKindNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LuneKind']) )
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
function LuneKindNode.new( id, pos, typeList, exp )
   local obj = {}
   LuneKindNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, exp ); end
   return obj
end
function LuneKindNode:__init(id, pos, typeList, exp) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LuneKind']), pos, typeList)
   
   
   self.exp = exp
   
end
function LuneKindNode.create( nodeMan, pos, typeList, exp )

   local node = LuneKindNode.new(nodeMan:nextId(  ), pos, typeList, exp)
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

   return _lune.unwrap( _moduleObj.nodeKind['DeclMacro'])
end


regKind( "DeclMacro" )
function Filter:processDeclMacro( node, opt )

end


function NodeManager:getDeclMacroNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclMacro']) )
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
function DeclMacroNode.new( id, pos, typeList, declInfo )
   local obj = {}
   DeclMacroNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, declInfo ); end
   return obj
end
function DeclMacroNode:__init(id, pos, typeList, declInfo) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['DeclMacro']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclMacroNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclMacroNode.new(nodeMan:nextId(  ), pos, typeList, declInfo)
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

function NodeKind.get_Abbr(  )

   return _lune.unwrap( _moduleObj.nodeKind['Abbr'])
end


regKind( "Abbr" )
function Filter:processAbbr( node, opt )

end


function NodeManager:getAbbrNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Abbr']) )
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
function AbbrNode.new( id, pos, typeList )
   local obj = {}
   AbbrNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList ); end
   return obj
end
function AbbrNode:__init(id, pos, typeList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Abbr']), pos, typeList)
   
   
   
end
function AbbrNode.create( nodeMan, pos, typeList )

   local node = AbbrNode.new(nodeMan:nextId(  ), pos, typeList)
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

   return _lune.unwrap( _moduleObj.nodeKind['Boxing'])
end


regKind( "Boxing" )
function Filter:processBoxing( node, opt )

end


function NodeManager:getBoxingNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Boxing']) )
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
function BoxingNode.new( id, pos, typeList, src )
   local obj = {}
   BoxingNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, src ); end
   return obj
end
function BoxingNode:__init(id, pos, typeList, src) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Boxing']), pos, typeList)
   
   
   self.src = src
   
end
function BoxingNode.create( nodeMan, pos, typeList, src )

   local node = BoxingNode.new(nodeMan:nextId(  ), pos, typeList, src)
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

   return _lune.unwrap( _moduleObj.nodeKind['Unboxing'])
end


regKind( "Unboxing" )
function Filter:processUnboxing( node, opt )

end


function NodeManager:getUnboxingNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Unboxing']) )
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
function UnboxingNode.new( id, pos, typeList, src )
   local obj = {}
   UnboxingNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, src ); end
   return obj
end
function UnboxingNode:__init(id, pos, typeList, src) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['Unboxing']), pos, typeList)
   
   
   self.src = src
   
end
function UnboxingNode.create( nodeMan, pos, typeList, src )

   local node = UnboxingNode.new(nodeMan:nextId(  ), pos, typeList, src)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralNil'])
end


regKind( "LiteralNil" )
function Filter:processLiteralNil( node, opt )

end


function NodeManager:getLiteralNilNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralNil']) )
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
function LiteralNilNode.new( id, pos, typeList )
   local obj = {}
   LiteralNilNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList ); end
   return obj
end
function LiteralNilNode:__init(id, pos, typeList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralNil']), pos, typeList)
   
   
   
end
function LiteralNilNode.create( nodeMan, pos, typeList )

   local node = LiteralNilNode.new(nodeMan:nextId(  ), pos, typeList)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralChar'])
end


regKind( "LiteralChar" )
function Filter:processLiteralChar( node, opt )

end


function NodeManager:getLiteralCharNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralChar']) )
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
function LiteralCharNode.new( id, pos, typeList, token, num )
   local obj = {}
   LiteralCharNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, token, num ); end
   return obj
end
function LiteralCharNode:__init(id, pos, typeList, token, num) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralChar']), pos, typeList)
   
   
   self.token = token
   self.num = num
   
end
function LiteralCharNode.create( nodeMan, pos, typeList, token, num )

   local node = LiteralCharNode.new(nodeMan:nextId(  ), pos, typeList, token, num)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralInt'])
end


regKind( "LiteralInt" )
function Filter:processLiteralInt( node, opt )

end


function NodeManager:getLiteralIntNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralInt']) )
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
function LiteralIntNode.new( id, pos, typeList, token, num )
   local obj = {}
   LiteralIntNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, token, num ); end
   return obj
end
function LiteralIntNode:__init(id, pos, typeList, token, num) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralInt']), pos, typeList)
   
   
   self.token = token
   self.num = num
   
end
function LiteralIntNode.create( nodeMan, pos, typeList, token, num )

   local node = LiteralIntNode.new(nodeMan:nextId(  ), pos, typeList, token, num)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralReal'])
end


regKind( "LiteralReal" )
function Filter:processLiteralReal( node, opt )

end


function NodeManager:getLiteralRealNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralReal']) )
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
function LiteralRealNode.new( id, pos, typeList, token, num )
   local obj = {}
   LiteralRealNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, token, num ); end
   return obj
end
function LiteralRealNode:__init(id, pos, typeList, token, num) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralReal']), pos, typeList)
   
   
   self.token = token
   self.num = num
   
end
function LiteralRealNode.create( nodeMan, pos, typeList, token, num )

   local node = LiteralRealNode.new(nodeMan:nextId(  ), pos, typeList, token, num)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralArray'])
end


regKind( "LiteralArray" )
function Filter:processLiteralArray( node, opt )

end


function NodeManager:getLiteralArrayNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralArray']) )
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
function LiteralArrayNode.new( id, pos, typeList, expList )
   local obj = {}
   LiteralArrayNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, expList ); end
   return obj
end
function LiteralArrayNode:__init(id, pos, typeList, expList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralArray']), pos, typeList)
   
   
   self.expList = expList
   
end
function LiteralArrayNode.create( nodeMan, pos, typeList, expList )

   local node = LiteralArrayNode.new(nodeMan:nextId(  ), pos, typeList, expList)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralList'])
end


regKind( "LiteralList" )
function Filter:processLiteralList( node, opt )

end


function NodeManager:getLiteralListNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralList']) )
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
function LiteralListNode.new( id, pos, typeList, expList )
   local obj = {}
   LiteralListNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, expList ); end
   return obj
end
function LiteralListNode:__init(id, pos, typeList, expList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralList']), pos, typeList)
   
   
   self.expList = expList
   
end
function LiteralListNode.create( nodeMan, pos, typeList, expList )

   local node = LiteralListNode.new(nodeMan:nextId(  ), pos, typeList, expList)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralSet'])
end


regKind( "LiteralSet" )
function Filter:processLiteralSet( node, opt )

end


function NodeManager:getLiteralSetNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralSet']) )
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
function LiteralSetNode.new( id, pos, typeList, expList )
   local obj = {}
   LiteralSetNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, expList ); end
   return obj
end
function LiteralSetNode:__init(id, pos, typeList, expList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralSet']), pos, typeList)
   
   
   self.expList = expList
   
end
function LiteralSetNode.create( nodeMan, pos, typeList, expList )

   local node = LiteralSetNode.new(nodeMan:nextId(  ), pos, typeList, expList)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralMap'])
end


regKind( "LiteralMap" )
function Filter:processLiteralMap( node, opt )

end


function NodeManager:getLiteralMapNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralMap']) )
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
function LiteralMapNode.new( id, pos, typeList, map, pairList )
   local obj = {}
   LiteralMapNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, map, pairList ); end
   return obj
end
function LiteralMapNode:__init(id, pos, typeList, map, pairList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralMap']), pos, typeList)
   
   
   self.map = map
   self.pairList = pairList
   
end
function LiteralMapNode.create( nodeMan, pos, typeList, map, pairList )

   local node = LiteralMapNode.new(nodeMan:nextId(  ), pos, typeList, map, pairList)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralString'])
end


regKind( "LiteralString" )
function Filter:processLiteralString( node, opt )

end


function NodeManager:getLiteralStringNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralString']) )
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
function LiteralStringNode.new( id, pos, typeList, token, argList )
   local obj = {}
   LiteralStringNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, token, argList ); end
   return obj
end
function LiteralStringNode:__init(id, pos, typeList, token, argList) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralString']), pos, typeList)
   
   
   self.token = token
   self.argList = argList
   
end
function LiteralStringNode.create( nodeMan, pos, typeList, token, argList )

   local node = LiteralStringNode.new(nodeMan:nextId(  ), pos, typeList, token, argList)
   nodeMan:addNode( node )
   return node
end
function LiteralStringNode:visit( visitor, depth )

   do
      local list = self.argList
      for __index, child in pairs( list ) do
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
function LiteralStringNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralStringNode  } )
end
function LiteralStringNode:get_token()
   return self.token
end
function LiteralStringNode:get_argList()
   return self.argList
end


function NodeKind.get_LiteralBool(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralBool'])
end


regKind( "LiteralBool" )
function Filter:processLiteralBool( node, opt )

end


function NodeManager:getLiteralBoolNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralBool']) )
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
function LiteralBoolNode.new( id, pos, typeList, token )
   local obj = {}
   LiteralBoolNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, token ); end
   return obj
end
function LiteralBoolNode:__init(id, pos, typeList, token) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralBool']), pos, typeList)
   
   
   self.token = token
   
end
function LiteralBoolNode.create( nodeMan, pos, typeList, token )

   local node = LiteralBoolNode.new(nodeMan:nextId(  ), pos, typeList, token)
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

   return _lune.unwrap( _moduleObj.nodeKind['LiteralSymbol'])
end


regKind( "LiteralSymbol" )
function Filter:processLiteralSymbol( node, opt )

end


function NodeManager:getLiteralSymbolNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralSymbol']) )
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
function LiteralSymbolNode.new( id, pos, typeList, token )
   local obj = {}
   LiteralSymbolNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, typeList, token ); end
   return obj
end
function LiteralSymbolNode:__init(id, pos, typeList, token) 
   Node.__init( self,id, _lune.unwrap( _moduleObj.nodeKind['LiteralSymbol']), pos, typeList)
   
   
   self.token = token
   
end
function LiteralSymbolNode.create( nodeMan, pos, typeList, token )

   local node = LiteralSymbolNode.new(nodeMan:nextId(  ), pos, typeList, token)
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
                  do
                     local _exp = refFieldNode:get_symbolInfo()
                     if _exp ~= nil then
                        return {_exp}
                     end
                  end
                  
               end
            end
            
         elseif _switchExp == NodeKind.get_GetField() then
            do
               local getFieldNode = _lune.__Cast( node, 3, GetFieldNode )
               if getFieldNode ~= nil then
                  do
                     local _exp = getFieldNode:get_symbolInfo()
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
                  for index, expNode in pairs( expListNode:get_expList() ) do
                     if index == #expListNode:get_expList() then
                        for __index, symbolInfo in pairs( processExpNode( expNode ) ) do
                           table.insert( list, symbolInfo )
                        end
                        
                     else
                      
                        for __index, symbolInfo in pairs( processExpNode( expNode ) ) do
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
      for __index, stmt in pairs( self.block:get_stmtList() ) do
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
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
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
      for __index, stmt in pairs( self.block:get_stmtList() ) do
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
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
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

   return _lune.newAlge( Literal.Nil)
end

function LiteralNilNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Symb, "nil" )
   return true
end

function LiteralCharNode:getLiteral(  )

   return _lune.newAlge( Literal.Int, {self.num})
end

function LiteralCharNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Char, string.format( "%d", self.num) )
   return true
end

function LiteralIntNode:getLiteral(  )

   return _lune.newAlge( Literal.Int, {self.num})
end

function LiteralIntNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Int, string.format( "%d", self.num) )
   return true
end

function LiteralRealNode:getLiteral(  )

   return _lune.newAlge( Literal.Real, {self.num})
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
         for __index, node in pairs( _exp:get_expList(  ) ) do
            local literal = node:getLiteral(  )
            if literal ~= nil then
               table.insert( literalList, literal )
            else
               return nil
            end
            
         end
         
      end
   end
   
   return _lune.newAlge( Literal.ARRAY, {literalList})
end

function LiteralArrayNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Dlmt, "[@" )
   do
      local _exp = self.expList
      if _exp ~= nil then
         for index, node in pairs( _exp:get_expList(  ) ) do
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
         for __index, node in pairs( _exp:get_expList(  ) ) do
            local literal = node:getLiteral(  )
            if literal ~= nil then
               table.insert( literalList, literal )
            else
               return nil
            end
            
         end
         
      end
   end
   
   return _lune.newAlge( Literal.LIST, {literalList})
end

function LiteralListNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Dlmt, "[" )
   do
      local _exp = self.expList
      if _exp ~= nil then
         for index, node in pairs( _exp:get_expList(  ) ) do
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
         for __index, node in pairs( _exp:get_expList(  ) ) do
            local literal = node:getLiteral(  )
            if literal ~= nil then
               table.insert( literalList, literal )
            else
               return nil
            end
            
         end
         
      end
   end
   
   return _lune.newAlge( Literal.SET, {literalList})
end

function LiteralSetNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Dlmt, "(@" )
   do
      local _exp = self.expList
      if _exp ~= nil then
         for index, node in pairs( _exp:get_expList(  ) ) do
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
      local keyLiteral = key:getLiteral(  )
      local valLiteral = val:getLiteral(  )
      if keyLiteral ~= nil and valLiteral ~= nil then
         litMap[keyLiteral] = valLiteral
      end
      
   end
   
   return _lune.newAlge( Literal.MAP, {litMap})
end

function LiteralMapNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Dlmt, "{" )
   local lit2valNode = {}
   for key, val in pairs( self.map ) do
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
            do
               return false
            end
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
      for __index, literal in ipairs( __sorted ) do
         local key = __map[ literal ]
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
   
   local argList = self:get_argList()
   if #argList > 0 then
      local paramList = {}
      for __index, argNode in pairs( argList ) do
         local arg = argNode:getLiteral(  )
         if  nil == arg then
            local _arg = arg
         
            return nil
         end
         
         paramList[#paramList + 1] = getLiteralObj( arg )
      end
      
      txt = string.format( txt, table.unpack( paramList ) )
   end
   
   return _lune.newAlge( Literal.Str, {txt})
end

function LiteralStringNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Str, self.token.txt )
   if #self:get_argList() > 0 then
      self:addTokenList( list, Parser.TokenKind.Dlmt, self.token.txt )
      for index, argNode in pairs( self:get_argList() ) do
         if index > 1 then
            self:addTokenList( list, Parser.TokenKind.Dlmt, "," )
         end
         
         if not argNode:setupLiteralTokenList( list ) then
            return false
         end
         
      end
      
   end
   
   return true
end

function LiteralBoolNode:getLiteral(  )

   return _lune.newAlge( Literal.Bool, {self.token.txt == "true"})
end

function LiteralBoolNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Symb, self.token.txt )
   return true
end

function LiteralSymbolNode:getLiteral(  )

   return _lune.newAlge( Literal.Symbol, {self.token.txt})
end

function LiteralSymbolNode:setupLiteralTokenList( list )

   self:addTokenList( list, Parser.TokenKind.Symb, self.token.txt )
   return true
end

function RefFieldNode:getLiteral(  )

   local tokenList = {}
   do
      local _matchExp = _lune.unwrap( self.prefix:getLiteral(  ))
      if _matchExp[1] == Literal.Symbol[1] then
         local symbol = _matchExp[2][1]
      
         table.insert( tokenList, symbol )
      elseif _matchExp[1] == Literal.Field[1] then
         local symList = _matchExp[2][1]
      
         for __index, symbol in pairs( symList ) do
            table.insert( tokenList, symbol )
         end
         
      else 
      do
         Util.errorLog( "not support" )
      end
      end
   end
   
   if self.nilAccess then
      table.insert( tokenList, "$." )
   else
    
      table.insert( tokenList, "." )
   end
   
   table.insert( tokenList, self.field.txt )
   return _lune.newAlge( Literal.Field, {tokenList})
end

function ExpMacroStatNode:getLiteral(  )

   local txt = ""
   for __index, token in pairs( self.expStrList ) do
      local literal = token:getLiteral(  )
      if  nil == literal then
         local _literal = literal
      
         Util.err( "illegal literal" )
      end
      
      do
         local _matchExp = literal
         if _matchExp[1] == Literal.Str[1] then
            local work = _matchExp[2][1]
         
            txt = string.format( "%s%s", txt, work)
         end
      end
      
   end
   
   return _lune.newAlge( Literal.Str, {txt})
end

local function enumLiiteral2Literal( obj )

   do
      local _matchExp = obj
      if _matchExp[1] == Ast.EnumLiteral.Int[1] then
         local val = _matchExp[2][1]
      
         return _lune.newAlge( Literal.Int, {val})
      elseif _matchExp[1] == Ast.EnumLiteral.Real[1] then
         local val = _matchExp[2][1]
      
         return _lune.newAlge( Literal.Real, {val})
      elseif _matchExp[1] == Ast.EnumLiteral.Str[1] then
         local val = _matchExp[2][1]
      
         return _lune.newAlge( Literal.Str, {val})
      end
   end
   
   Util.err( "illegal enum" .. Ast.EnumLiteral:_getTxt( obj)
    )
end

function ExpRefNode:getLiteral(  )

   local typeInfo = self.symbolInfo:get_typeInfo()
   do
      local enumTypeInfo = _lune.__Cast( typeInfo, 3, Ast.EnumTypeInfo )
      if enumTypeInfo ~= nil then
         local enumval = _lune.unwrap( enumTypeInfo:getEnumValInfo( self.symbolInfo:get_name() ))
         local val = enumLiiteral2Literal( enumval:get_val() )
         return val
      end
   end
   
   return nil
end

function ExpOmitEnumNode:getLiteral(  )

   local enumTypeInfo = self.enumTypeInfo
   local enumval = self.valInfo
   local val = enumLiiteral2Literal( enumval:get_val() )
   return val
end

function ExpOmitEnumNode:setupLiteralTokenList( list )

   local enumTypeInfo = self.enumTypeInfo
   local enumval = self.valInfo
   self:addTokenList( list, Parser.TokenKind.Dlmt, "." )
   self:addTokenList( list, Parser.TokenKind.Symb, (Ast.EnumLiteral:_getTxt( enumval:get_val())
   :gsub( ".*%.", "" ) ) )
   return true
end

function ExpOp2Node:getLiteral(  )

   local function getValType( node )
   
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
                     realVal = val
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
            realVal = val
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
         do
            return false, 0, 0.0, "", Ast.headTypeInfo
         end
         end
      end
      
      return true, intVal, realVal, strVal, retTypeInfo
   end
   
   local ret1, int1, real1, str1, type1 = getValType( self:get_exp1() )
   local ret2, int2, real2, str2, type2 = getValType( self:get_exp2() )
   if not ret1 or not ret2 then
      return nil
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
               return _lune.newAlge( Literal.Int, {int1 + int2})
            end
            
            return _lune.newAlge( Literal.Real, {real1 + real2})
         elseif _switchExp == "-" then
            if retType == Ast.builtinTypeInt then
               return _lune.newAlge( Literal.Int, {int1 - int2})
            end
            
            return _lune.newAlge( Literal.Real, {real1 - real2})
         elseif _switchExp == "*" then
            if retType == Ast.builtinTypeInt then
               return _lune.newAlge( Literal.Int, {int1 * int2})
            end
            
            return _lune.newAlge( Literal.Real, {real1 * real2})
         elseif _switchExp == "/" then
            if retType == Ast.builtinTypeInt then
               return _lune.newAlge( Literal.Int, {math.floor(int1 / int2)})
            end
            
            return _lune.newAlge( Literal.Real, {real1 / real2})
         end
      end
      
   elseif type1 == Ast.builtinTypeString and type2 == Ast.builtinTypeString then
      if self.op.txt == ".." then
         return _lune.newAlge( Literal.Str, {str1 .. str2})
      end
      
   end
   
   return nil
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
   for __index, argNode in pairs( declInfo:get_argList() ) do
      if argNode:get_kind(  ) == NodeKind.get_DeclArg() then
         local argType = argNode:get_argType():get_expType()
         local argName = argNode:get_name().txt
         table.insert( self.argList, MacroArgInfo.new(argName, argType) )
      end
      
   end
   
end
function DefMacroInfo.setmeta( obj )
  setmetatable( obj, { __index = DefMacroInfo  } )
end

return _moduleObj
