--lune/base/ConvNodes.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@ConvNodes'
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
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Util = _lune.loadModule( 'lune.base.Util' )

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


local Node = {}



_moduleObj.Node = Node
function Node:processFilter( filter, opt )

end
function Node.setmeta( obj )
  setmetatable( obj, { __index = Node  } )
end
function Node.new( id, pos, kind, hasNilAcc, parent )
   local obj = {}
   Node.setmeta( obj )
   if obj.__init then
      obj:__init( id, pos, kind, hasNilAcc, parent )
   end
   return obj
end
function Node:__init( id, pos, kind, hasNilAcc, parent )

   self.id = id
   self.pos = pos
   self.kind = kind
   self.hasNilAcc = hasNilAcc
   self.parent = parent
end
function Node:get_id()
   return self.id
end
function Node:get_pos()
   return self.pos
end
function Node:get_kind()
   return self.kind
end
function Node:get_hasNilAcc()
   return self.hasNilAcc
end
function Node:get_parent()
   return self.parent
end
function Node:set_parent( parent )
   self.parent = parent
end


local ExpNode = {}
setmetatable( ExpNode, { __index = Node } )
_moduleObj.ExpNode = ExpNode
function ExpNode:visit( visitor, depth )

   return false
end
function ExpNode.setmeta( obj )
  setmetatable( obj, { __index = ExpNode  } )
end
function ExpNode.new( __superarg1, __superarg2, __superarg3, __superarg4, __superarg5 )
   local obj = {}
   ExpNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3, __superarg4, __superarg5 )
   end
   return obj
end
function ExpNode:__init( __superarg1, __superarg2, __superarg3, __superarg4, __superarg5 )

   Node.__init( self, __superarg1, __superarg2, __superarg3, __superarg4, __superarg5 )
end


local DummyExpNode = {}
setmetatable( DummyExpNode, { __index = ExpNode } )
_moduleObj.DummyExpNode = DummyExpNode
function DummyExpNode.new(  )
   local obj = {}
   DummyExpNode.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function DummyExpNode:__init() 
   ExpNode.__init( self,0, Parser.getEofToken(  ).pos, "DummyExpNode", false, nil)
   
end
function DummyExpNode.setmeta( obj )
  setmetatable( obj, { __index = DummyExpNode  } )
end

local dummyExpNode = DummyExpNode.new()

local NilAccNode = {}
setmetatable( NilAccNode, { __index = ExpNode } )
_moduleObj.NilAccNode = NilAccNode
function NilAccNode.new( parent, prefix, acc )
   local obj = {}
   NilAccNode.setmeta( obj )
   if obj.__init then obj:__init( parent, prefix, acc ); end
   return obj
end
function NilAccNode:__init(parent, prefix, acc) 
   ExpNode.__init( self,0, acc:get_pos(), "NicAccNode", false, parent)
   
   self.prefix = prefix
   self.acc = acc
end
function NilAccNode:visit( visitor, depth )

   do
      local _switchExp = visitor( self.prefix, self, "prefix", depth )
      if _switchExp == Nodes.NodeVisitMode.Child then
         if not self.prefix:visit( visitor, depth + 1 ) then
            return false
         end
         
      elseif _switchExp == Nodes.NodeVisitMode.End then
         return false
      end
   end
   
   do
      local _switchExp = visitor( self.acc, self, "acc", depth )
      if _switchExp == Nodes.NodeVisitMode.Child then
         if not self.acc:visit( visitor, depth + 1 ) then
            return false
         end
         
      elseif _switchExp == Nodes.NodeVisitMode.End then
         return false
      end
   end
   
   return true
end
function NilAccNode.setmeta( obj )
  setmetatable( obj, { __index = NilAccNode  } )
end
function NilAccNode:get_prefix()
   return self.prefix
end
function NilAccNode:get_acc()
   return self.acc
end
function NilAccNode:set_acc( acc )
   self.acc = acc
end

local State = {}
_moduleObj.State = State
function State.new( node )
   local obj = {}
   State.setmeta( obj )
   if obj.__init then obj:__init( node ); end
   return obj
end
function State:__init(node) 
   self.topNode = node
   self.nilAccNode = nil
end
function State:setNilAcc( node, parent )

   if not self.nilAccNode then
      if not node:get_hasNilAcc() and parent:get_hasNilAcc() then
         self.nilAccNode = NilAccNode.new(parent, node, self.topNode)
      end
      
   end
   
end
function State.setmeta( obj )
  setmetatable( obj, { __index = State  } )
end
function State:get_nilAccNode()
   return self.nilAccNode
end





local nodeKind2createFromFunc = {}

local function Node_createFromNode( node, parent, state )

   local func = nodeKind2createFromFunc[node:get_kind()]
   if  nil == func then
      local _func = func
   
      Util.err( string.format( "not support -- %s", Nodes.getNodeKindName( node:get_kind() )) )
   end
   
   local convNode = func( node, parent, state )
   
   return convNode
end

local FieldIndex = {}
FieldIndex._val2NameMap = {}
function FieldIndex:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "FieldIndex.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function FieldIndex._from( val )
   if FieldIndex._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
FieldIndex.__allList = {}
function FieldIndex.get__allList()
   return FieldIndex.__allList
end

FieldIndex.Name = 1
FieldIndex._val2NameMap[1] = 'Name'
FieldIndex.__allList[1] = FieldIndex.Name
FieldIndex.Type = 2
FieldIndex._val2NameMap[2] = 'Type'
FieldIndex.__allList[2] = FieldIndex.Type
FieldIndex.Init = 3
FieldIndex._val2NameMap[3] = 'Init'
FieldIndex.__allList[3] = FieldIndex.Init





local ExpListNode = {}



setmetatable( ExpListNode, { __index = ExpNode } )
_moduleObj.ExpListNode = ExpListNode
function ExpListNode:processFilter( filter, opt )

   filter:processExpList( self, opt )
end
function ExpListNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpListNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpListNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.expList = {}
   
   
end
function ExpListNode:visit( visitor, depth )

   do
      local list = self.expList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'expList', depth )
            if _switchExp == Nodes.NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == Nodes.NodeVisitMode.End then
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
function ExpListNode:get_orgNode()
   return self.orgNode
end
function ExpListNode:get_expList()
   return self.expList
end
function ExpListNode:set_expList( expList )
   self.expList = expList
end

local function ExpListNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpListNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpListNode') )
   end
   
   local convNode = ExpListNode.new(node:get_id(), node:get_effectivePos(), 'ExpListNode', workNode:hasNilAccess(  ), parent, node)
   local function createexpList(  )
   
      local list = node:get_expList()
      local expList = {}
      for __index, exp in ipairs( list ) do
         local newConv
         
         
         do
            local newState = State.new(convNode)
            local newNode = Node_createFromNode( exp, convNode, newState )
            do
               local nilAccNode = newState:get_nilAccNode()
               if nilAccNode ~= nil then
                  newConv = nilAccNode
               else
                  newConv = newNode
               end
            end
            
         end
         
         
         table.insert( expList, newConv )
      end
      
      return expList
   end
   convNode:set_expList( createexpList(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpListNode_createFromNode = ExpListNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpList()] = function ( node, workParent, state )

   return ExpListNode_createFromNode( node, workParent, state )
end




local ExpNewNode = {}



setmetatable( ExpNewNode, { __index = ExpNode } )
_moduleObj.ExpNewNode = ExpNewNode
function ExpNewNode:processFilter( filter, opt )

   filter:processExpNew( self, opt )
end
function ExpNewNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpNewNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpNewNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.argList = nil
   
   
end
function ExpNewNode:visit( visitor, depth )

   do
      do
         local child = self.argList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'argList', depth )
               if _switchExp == Nodes.NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == Nodes.NodeVisitMode.End then
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
function ExpNewNode:get_orgNode()
   return self.orgNode
end
function ExpNewNode:get_argList()
   return self.argList
end
function ExpNewNode:set_argList( argList )
   self.argList = argList
end

local function ExpNewNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpNewNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpNewNode') )
   end
   
   local convNode = ExpNewNode.new(node:get_id(), node:get_effectivePos(), 'ExpNewNode', workNode:hasNilAccess(  ), parent, node)
   local function createargList(  )
   
      local child = node:get_argList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      local paramNode = ExpListNode_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_argList( createargList(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpNewNode_createFromNode = ExpNewNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpNew()] = function ( node, workParent, state )

   return ExpNewNode_createFromNode( node, workParent, state )
end




local ExpUnwrapNode = {}



setmetatable( ExpUnwrapNode, { __index = ExpNode } )
_moduleObj.ExpUnwrapNode = ExpUnwrapNode
function ExpUnwrapNode:processFilter( filter, opt )

   filter:processExpUnwrap( self, opt )
end
function ExpUnwrapNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpUnwrapNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpUnwrapNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.exp = dummyExpNode
   self.default = nil
   
   
end
function ExpUnwrapNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
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
               if _switchExp == Nodes.NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == Nodes.NodeVisitMode.End then
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
function ExpUnwrapNode:get_orgNode()
   return self.orgNode
end
function ExpUnwrapNode:get_exp()
   return self.exp
end
function ExpUnwrapNode:set_exp( exp )
   self.exp = exp
end
function ExpUnwrapNode:get_default()
   return self.default
end
function ExpUnwrapNode:set_default( default )
   self.default = default
end

local function ExpUnwrapNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpUnwrapNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpUnwrapNode') )
   end
   
   local convNode = ExpUnwrapNode.new(node:get_id(), node:get_effectivePos(), 'ExpUnwrapNode', workNode:hasNilAccess(  ), parent, node)
   local function createexp(  )
   
      local child = node:get_exp()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_exp( createexp(  ) )
   local function createdefault(  )
   
      local child = node:get_default()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_default( createdefault(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpUnwrapNode_createFromNode = ExpUnwrapNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpUnwrap()] = function ( node, workParent, state )

   return ExpUnwrapNode_createFromNode( node, workParent, state )
end




local ExpRefNode = {}



setmetatable( ExpRefNode, { __index = ExpNode } )
_moduleObj.ExpRefNode = ExpRefNode
function ExpRefNode:processFilter( filter, opt )

   filter:processExpRef( self, opt )
end
function ExpRefNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpRefNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpRefNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   
   
end
function ExpRefNode:visit( visitor, depth )

   
   return true
end
function ExpRefNode.setmeta( obj )
  setmetatable( obj, { __index = ExpRefNode  } )
end
function ExpRefNode:get_orgNode()
   return self.orgNode
end

local function ExpRefNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpRefNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpRefNode') )
   end
   
   local convNode = ExpRefNode.new(node:get_id(), node:get_effectivePos(), 'ExpRefNode', workNode:hasNilAccess(  ), parent, node)
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpRefNode_createFromNode = ExpRefNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpRef()] = function ( node, workParent, state )

   return ExpRefNode_createFromNode( node, workParent, state )
end




local ExpOp2Node = {}



setmetatable( ExpOp2Node, { __index = ExpNode } )
_moduleObj.ExpOp2Node = ExpOp2Node
function ExpOp2Node:processFilter( filter, opt )

   filter:processExpOp2( self, opt )
end
function ExpOp2Node.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpOp2Node.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpOp2Node:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.exp1 = dummyExpNode
   self.exp2 = dummyExpNode
   
   
end
function ExpOp2Node:visit( visitor, depth )

   do
      local child = self.exp1
      do
         local _switchExp = visitor( child, self, 'exp1', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.exp2
      do
         local _switchExp = visitor( child, self, 'exp2', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpOp2Node.setmeta( obj )
  setmetatable( obj, { __index = ExpOp2Node  } )
end
function ExpOp2Node:get_orgNode()
   return self.orgNode
end
function ExpOp2Node:get_exp1()
   return self.exp1
end
function ExpOp2Node:set_exp1( exp1 )
   self.exp1 = exp1
end
function ExpOp2Node:get_exp2()
   return self.exp2
end
function ExpOp2Node:set_exp2( exp2 )
   self.exp2 = exp2
end

local function ExpOp2Node_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpOp2Node )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpOp2Node') )
   end
   
   local convNode = ExpOp2Node.new(node:get_id(), node:get_effectivePos(), 'ExpOp2Node', workNode:hasNilAccess(  ), parent, node)
   local function createexp1(  )
   
      local child = node:get_exp1()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_exp1( createexp1(  ) )
   local function createexp2(  )
   
      local child = node:get_exp2()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_exp2( createexp2(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpOp2Node_createFromNode = ExpOp2Node_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpOp2()] = function ( node, workParent, state )

   return ExpOp2Node_createFromNode( node, workParent, state )
end




local ExpCastNode = {}



setmetatable( ExpCastNode, { __index = ExpNode } )
_moduleObj.ExpCastNode = ExpCastNode
function ExpCastNode:processFilter( filter, opt )

   filter:processExpCast( self, opt )
end
function ExpCastNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpCastNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpCastNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.exp = dummyExpNode
   
   
end
function ExpCastNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpCastNode.setmeta( obj )
  setmetatable( obj, { __index = ExpCastNode  } )
end
function ExpCastNode:get_orgNode()
   return self.orgNode
end
function ExpCastNode:get_exp()
   return self.exp
end
function ExpCastNode:set_exp( exp )
   self.exp = exp
end

local function ExpCastNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpCastNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpCastNode') )
   end
   
   local convNode = ExpCastNode.new(node:get_id(), node:get_effectivePos(), 'ExpCastNode', workNode:hasNilAccess(  ), parent, node)
   local function createexp(  )
   
      local child = node:get_exp()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_exp( createexp(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpCastNode_createFromNode = ExpCastNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpCast()] = function ( node, workParent, state )

   return ExpCastNode_createFromNode( node, workParent, state )
end



local ExpToDDDNode = {}
setmetatable( ExpToDDDNode, { __index = ExpNode } )
_moduleObj.ExpToDDDNode = ExpToDDDNode
function ExpToDDDNode.setmeta( obj )
  setmetatable( obj, { __index = ExpToDDDNode  } )
end
function ExpToDDDNode.new( __superarg1, __superarg2, __superarg3, __superarg4, __superarg5,orgNode, expList )
   local obj = {}
   ExpToDDDNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3, __superarg4, __superarg5,orgNode, expList )
   end
   return obj
end
function ExpToDDDNode:__init( __superarg1, __superarg2, __superarg3, __superarg4, __superarg5,orgNode, expList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3, __superarg4, __superarg5 )
   self.orgNode = orgNode
   self.expList = expList
end
function ExpToDDDNode:get_orgNode()
   return self.orgNode
end
function ExpToDDDNode:get_expList()
   return self.expList
end
function ExpToDDDNode:set_expList( expList )
   self.expList = expList
end


local function ExpToDDDNode_createFromNode( workNode, parent, state )
   local __func__ = '@lune.@base.@ConvNodes.ExpToDDDNode_createFromNode'

   local node = _lune.__Cast( workNode, 3, Nodes.ExpToDDDNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s -- %s", tostring( workNode:get_kind()), __func__) )
   end
   
   local expList = ExpListNode_createFromNode( node:get_expList(), dummyExpNode, state )
   
   local convNode = ExpToDDDNode.new(node:get_id(), node:get_effectivePos(), "ExpToDDDNode", false, parent, node, expList)
   expList:set_parent( convNode )
   
   state:setNilAcc( convNode, parent )
   
   return convNode
end
_moduleObj.ExpToDDDNode_createFromNode = ExpToDDDNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpToDDD()] = function ( node, parent, state )

   return ExpToDDDNode_createFromNode( node, parent, state )
end


local ExpSubDDDNode = {}



setmetatable( ExpSubDDDNode, { __index = ExpNode } )
_moduleObj.ExpSubDDDNode = ExpSubDDDNode
function ExpSubDDDNode:processFilter( filter, opt )

   filter:processExpSubDDD( self, opt )
end
function ExpSubDDDNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpSubDDDNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpSubDDDNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.src = dummyExpNode
   
   
end
function ExpSubDDDNode:visit( visitor, depth )

   do
      local child = self.src
      do
         local _switchExp = visitor( child, self, 'src', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpSubDDDNode.setmeta( obj )
  setmetatable( obj, { __index = ExpSubDDDNode  } )
end
function ExpSubDDDNode:get_orgNode()
   return self.orgNode
end
function ExpSubDDDNode:get_src()
   return self.src
end
function ExpSubDDDNode:set_src( src )
   self.src = src
end

local function ExpSubDDDNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpSubDDDNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpSubDDDNode') )
   end
   
   local convNode = ExpSubDDDNode.new(node:get_id(), node:get_effectivePos(), 'ExpSubDDDNode', workNode:hasNilAccess(  ), parent, node)
   local function createsrc(  )
   
      local child = node:get_src()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_src( createsrc(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpSubDDDNode_createFromNode = ExpSubDDDNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpSubDDD()] = function ( node, workParent, state )

   return ExpSubDDDNode_createFromNode( node, workParent, state )
end




local ExpOp1Node = {}



setmetatable( ExpOp1Node, { __index = ExpNode } )
_moduleObj.ExpOp1Node = ExpOp1Node
function ExpOp1Node:processFilter( filter, opt )

   filter:processExpOp1( self, opt )
end
function ExpOp1Node.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpOp1Node.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpOp1Node:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.exp = dummyExpNode
   
   
end
function ExpOp1Node:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpOp1Node.setmeta( obj )
  setmetatable( obj, { __index = ExpOp1Node  } )
end
function ExpOp1Node:get_orgNode()
   return self.orgNode
end
function ExpOp1Node:get_exp()
   return self.exp
end
function ExpOp1Node:set_exp( exp )
   self.exp = exp
end

local function ExpOp1Node_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpOp1Node )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpOp1Node') )
   end
   
   local convNode = ExpOp1Node.new(node:get_id(), node:get_effectivePos(), 'ExpOp1Node', workNode:hasNilAccess(  ), parent, node)
   local function createexp(  )
   
      local child = node:get_exp()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_exp( createexp(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpOp1Node_createFromNode = ExpOp1Node_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpOp1()] = function ( node, workParent, state )

   return ExpOp1Node_createFromNode( node, workParent, state )
end




local ExpRefItemNode = {}



setmetatable( ExpRefItemNode, { __index = ExpNode } )
_moduleObj.ExpRefItemNode = ExpRefItemNode
function ExpRefItemNode:processFilter( filter, opt )

   filter:processExpRefItem( self, opt )
end
function ExpRefItemNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpRefItemNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpRefItemNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.val = dummyExpNode
   self.index = nil
   
   
end
function ExpRefItemNode:visit( visitor, depth )

   do
      local child = self.val
      do
         local _switchExp = visitor( child, self, 'val', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
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
               if _switchExp == Nodes.NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == Nodes.NodeVisitMode.End then
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
function ExpRefItemNode:get_orgNode()
   return self.orgNode
end
function ExpRefItemNode:get_val()
   return self.val
end
function ExpRefItemNode:set_val( val )
   self.val = val
end
function ExpRefItemNode:get_index()
   return self.index
end
function ExpRefItemNode:set_index( index )
   self.index = index
end

local function ExpRefItemNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpRefItemNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpRefItemNode') )
   end
   
   local convNode = ExpRefItemNode.new(node:get_id(), node:get_effectivePos(), 'ExpRefItemNode', workNode:hasNilAccess(  ), parent, node)
   local function createval(  )
   
      local child = node:get_val()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_val( createval(  ) )
   local function createindex(  )
   
      local child = node:get_index()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_index( createindex(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpRefItemNode_createFromNode = ExpRefItemNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpRefItem()] = function ( node, workParent, state )

   return ExpRefItemNode_createFromNode( node, workParent, state )
end




local ExpCallNode = {}



setmetatable( ExpCallNode, { __index = ExpNode } )
_moduleObj.ExpCallNode = ExpCallNode
function ExpCallNode:processFilter( filter, opt )

   filter:processExpCall( self, opt )
end
function ExpCallNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpCallNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpCallNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.func = dummyExpNode
   self.argList = nil
   
   
end
function ExpCallNode:visit( visitor, depth )

   do
      local child = self.func
      do
         local _switchExp = visitor( child, self, 'func', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
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
               if _switchExp == Nodes.NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == Nodes.NodeVisitMode.End then
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
function ExpCallNode:get_orgNode()
   return self.orgNode
end
function ExpCallNode:get_func()
   return self.func
end
function ExpCallNode:set_func( func )
   self.func = func
end
function ExpCallNode:get_argList()
   return self.argList
end
function ExpCallNode:set_argList( argList )
   self.argList = argList
end

local function ExpCallNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpCallNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpCallNode') )
   end
   
   local convNode = ExpCallNode.new(node:get_id(), node:get_effectivePos(), 'ExpCallNode', workNode:hasNilAccess(  ), parent, node)
   local function createfunc(  )
   
      local child = node:get_func()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_func( createfunc(  ) )
   local function createargList(  )
   
      local child = node:get_argList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      local paramNode = ExpListNode_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_argList( createargList(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpCallNode_createFromNode = ExpCallNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpCall()] = function ( node, workParent, state )

   return ExpCallNode_createFromNode( node, workParent, state )
end




local ExpMRetNode = {}



setmetatable( ExpMRetNode, { __index = ExpNode } )
_moduleObj.ExpMRetNode = ExpMRetNode
function ExpMRetNode:processFilter( filter, opt )

   filter:processExpMRet( self, opt )
end
function ExpMRetNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpMRetNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpMRetNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.mRet = dummyExpNode
   
   
end
function ExpMRetNode:visit( visitor, depth )

   do
      local child = self.mRet
      do
         local _switchExp = visitor( child, self, 'mRet', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpMRetNode.setmeta( obj )
  setmetatable( obj, { __index = ExpMRetNode  } )
end
function ExpMRetNode:get_orgNode()
   return self.orgNode
end
function ExpMRetNode:get_mRet()
   return self.mRet
end
function ExpMRetNode:set_mRet( mRet )
   self.mRet = mRet
end

local function ExpMRetNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpMRetNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpMRetNode') )
   end
   
   local convNode = ExpMRetNode.new(node:get_id(), node:get_effectivePos(), 'ExpMRetNode', workNode:hasNilAccess(  ), parent, node)
   local function createmRet(  )
   
      local child = node:get_mRet()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_mRet( createmRet(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpMRetNode_createFromNode = ExpMRetNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpMRet()] = function ( node, workParent, state )

   return ExpMRetNode_createFromNode( node, workParent, state )
end




local ExpAccessMRetNode = {}



setmetatable( ExpAccessMRetNode, { __index = ExpNode } )
_moduleObj.ExpAccessMRetNode = ExpAccessMRetNode
function ExpAccessMRetNode:processFilter( filter, opt )

   filter:processExpAccessMRet( self, opt )
end
function ExpAccessMRetNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpAccessMRetNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpAccessMRetNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.mRet = dummyExpNode
   
   
end
function ExpAccessMRetNode:visit( visitor, depth )

   do
      local child = self.mRet
      do
         local _switchExp = visitor( child, self, 'mRet', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpAccessMRetNode.setmeta( obj )
  setmetatable( obj, { __index = ExpAccessMRetNode  } )
end
function ExpAccessMRetNode:get_orgNode()
   return self.orgNode
end
function ExpAccessMRetNode:get_mRet()
   return self.mRet
end
function ExpAccessMRetNode:set_mRet( mRet )
   self.mRet = mRet
end

local function ExpAccessMRetNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpAccessMRetNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpAccessMRetNode') )
   end
   
   local convNode = ExpAccessMRetNode.new(node:get_id(), node:get_effectivePos(), 'ExpAccessMRetNode', workNode:hasNilAccess(  ), parent, node)
   local function createmRet(  )
   
      local child = node:get_mRet()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_mRet( createmRet(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpAccessMRetNode_createFromNode = ExpAccessMRetNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpAccessMRet()] = function ( node, workParent, state )

   return ExpAccessMRetNode_createFromNode( node, workParent, state )
end




local ExpMultiTo1Node = {}



setmetatable( ExpMultiTo1Node, { __index = ExpNode } )
_moduleObj.ExpMultiTo1Node = ExpMultiTo1Node
function ExpMultiTo1Node:processFilter( filter, opt )

   filter:processExpMultiTo1( self, opt )
end
function ExpMultiTo1Node.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpMultiTo1Node.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpMultiTo1Node:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.exp = dummyExpNode
   
   
end
function ExpMultiTo1Node:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpMultiTo1Node.setmeta( obj )
  setmetatable( obj, { __index = ExpMultiTo1Node  } )
end
function ExpMultiTo1Node:get_orgNode()
   return self.orgNode
end
function ExpMultiTo1Node:get_exp()
   return self.exp
end
function ExpMultiTo1Node:set_exp( exp )
   self.exp = exp
end

local function ExpMultiTo1Node_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpMultiTo1Node )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpMultiTo1Node') )
   end
   
   local convNode = ExpMultiTo1Node.new(node:get_id(), node:get_effectivePos(), 'ExpMultiTo1Node', workNode:hasNilAccess(  ), parent, node)
   local function createexp(  )
   
      local child = node:get_exp()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_exp( createexp(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpMultiTo1Node_createFromNode = ExpMultiTo1Node_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpMultiTo1()] = function ( node, workParent, state )

   return ExpMultiTo1Node_createFromNode( node, workParent, state )
end




local ExpParenNode = {}



setmetatable( ExpParenNode, { __index = ExpNode } )
_moduleObj.ExpParenNode = ExpParenNode
function ExpParenNode:processFilter( filter, opt )

   filter:processExpParen( self, opt )
end
function ExpParenNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpParenNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpParenNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.exp = dummyExpNode
   
   
end
function ExpParenNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function ExpParenNode.setmeta( obj )
  setmetatable( obj, { __index = ExpParenNode  } )
end
function ExpParenNode:get_orgNode()
   return self.orgNode
end
function ExpParenNode:get_exp()
   return self.exp
end
function ExpParenNode:set_exp( exp )
   self.exp = exp
end

local function ExpParenNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpParenNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpParenNode') )
   end
   
   local convNode = ExpParenNode.new(node:get_id(), node:get_effectivePos(), 'ExpParenNode', workNode:hasNilAccess(  ), parent, node)
   local function createexp(  )
   
      local child = node:get_exp()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_exp( createexp(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpParenNode_createFromNode = ExpParenNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpParen()] = function ( node, workParent, state )

   return ExpParenNode_createFromNode( node, workParent, state )
end




local ExpOmitEnumNode = {}



setmetatable( ExpOmitEnumNode, { __index = ExpNode } )
_moduleObj.ExpOmitEnumNode = ExpOmitEnumNode
function ExpOmitEnumNode:processFilter( filter, opt )

   filter:processExpOmitEnum( self, opt )
end
function ExpOmitEnumNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   ExpOmitEnumNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function ExpOmitEnumNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   
   
end
function ExpOmitEnumNode:visit( visitor, depth )

   
   return true
end
function ExpOmitEnumNode.setmeta( obj )
  setmetatable( obj, { __index = ExpOmitEnumNode  } )
end
function ExpOmitEnumNode:get_orgNode()
   return self.orgNode
end

local function ExpOmitEnumNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpOmitEnumNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpOmitEnumNode') )
   end
   
   local convNode = ExpOmitEnumNode.new(node:get_id(), node:get_effectivePos(), 'ExpOmitEnumNode', workNode:hasNilAccess(  ), parent, node)
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.ExpOmitEnumNode_createFromNode = ExpOmitEnumNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpOmitEnum()] = function ( node, workParent, state )

   return ExpOmitEnumNode_createFromNode( node, workParent, state )
end




local RefFieldNode = {}



setmetatable( RefFieldNode, { __index = ExpNode } )
_moduleObj.RefFieldNode = RefFieldNode
function RefFieldNode:processFilter( filter, opt )

   filter:processRefField( self, opt )
end
function RefFieldNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   RefFieldNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function RefFieldNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.prefix = dummyExpNode
   
   
end
function RefFieldNode:visit( visitor, depth )

   do
      local child = self.prefix
      do
         local _switchExp = visitor( child, self, 'prefix', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function RefFieldNode.setmeta( obj )
  setmetatable( obj, { __index = RefFieldNode  } )
end
function RefFieldNode:get_orgNode()
   return self.orgNode
end
function RefFieldNode:get_prefix()
   return self.prefix
end
function RefFieldNode:set_prefix( prefix )
   self.prefix = prefix
end

local function RefFieldNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.RefFieldNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'RefFieldNode') )
   end
   
   local convNode = RefFieldNode.new(node:get_id(), node:get_effectivePos(), 'RefFieldNode', workNode:hasNilAccess(  ), parent, node)
   local function createprefix(  )
   
      local child = node:get_prefix()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_prefix( createprefix(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.RefFieldNode_createFromNode = RefFieldNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_RefField()] = function ( node, workParent, state )

   return RefFieldNode_createFromNode( node, workParent, state )
end




local GetFieldNode = {}



setmetatable( GetFieldNode, { __index = ExpNode } )
_moduleObj.GetFieldNode = GetFieldNode
function GetFieldNode:processFilter( filter, opt )

   filter:processGetField( self, opt )
end
function GetFieldNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   GetFieldNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function GetFieldNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.prefix = dummyExpNode
   
   
end
function GetFieldNode:visit( visitor, depth )

   do
      local child = self.prefix
      do
         local _switchExp = visitor( child, self, 'prefix', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function GetFieldNode.setmeta( obj )
  setmetatable( obj, { __index = GetFieldNode  } )
end
function GetFieldNode:get_orgNode()
   return self.orgNode
end
function GetFieldNode:get_prefix()
   return self.prefix
end
function GetFieldNode:set_prefix( prefix )
   self.prefix = prefix
end

local function GetFieldNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.GetFieldNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'GetFieldNode') )
   end
   
   local convNode = GetFieldNode.new(node:get_id(), node:get_effectivePos(), 'GetFieldNode', workNode:hasNilAccess(  ), parent, node)
   local function createprefix(  )
   
      local child = node:get_prefix()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_prefix( createprefix(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.GetFieldNode_createFromNode = GetFieldNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_GetField()] = function ( node, workParent, state )

   return GetFieldNode_createFromNode( node, workParent, state )
end




local DeclFuncNode = {}



setmetatable( DeclFuncNode, { __index = ExpNode } )
_moduleObj.DeclFuncNode = DeclFuncNode
function DeclFuncNode:processFilter( filter, opt )

   filter:processDeclFunc( self, opt )
end
function DeclFuncNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   DeclFuncNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function DeclFuncNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   
   
end
function DeclFuncNode:visit( visitor, depth )

   
   return true
end
function DeclFuncNode.setmeta( obj )
  setmetatable( obj, { __index = DeclFuncNode  } )
end
function DeclFuncNode:get_orgNode()
   return self.orgNode
end

local function DeclFuncNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.DeclFuncNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'DeclFuncNode') )
   end
   
   local convNode = DeclFuncNode.new(node:get_id(), node:get_effectivePos(), 'DeclFuncNode', workNode:hasNilAccess(  ), parent, node)
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.DeclFuncNode_createFromNode = DeclFuncNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_DeclFunc()] = function ( node, workParent, state )

   return DeclFuncNode_createFromNode( node, workParent, state )
end




local NewAlgeValNode = {}



setmetatable( NewAlgeValNode, { __index = ExpNode } )
_moduleObj.NewAlgeValNode = NewAlgeValNode
function NewAlgeValNode:processFilter( filter, opt )

   filter:processNewAlgeVal( self, opt )
end
function NewAlgeValNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   NewAlgeValNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function NewAlgeValNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.paramList = {}
   
   
end
function NewAlgeValNode:visit( visitor, depth )

   do
      local list = self.paramList
      for __index, child in ipairs( list ) do
         do
            local _switchExp = visitor( child, self, 'paramList', depth )
            if _switchExp == Nodes.NodeVisitMode.Child then
               if not child:visit( visitor, depth + 1 ) then
                  return false
               end
               
            elseif _switchExp == Nodes.NodeVisitMode.End then
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
function NewAlgeValNode:get_orgNode()
   return self.orgNode
end
function NewAlgeValNode:get_paramList()
   return self.paramList
end
function NewAlgeValNode:set_paramList( paramList )
   self.paramList = paramList
end

local function NewAlgeValNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.NewAlgeValNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'NewAlgeValNode') )
   end
   
   local convNode = NewAlgeValNode.new(node:get_id(), node:get_effectivePos(), 'NewAlgeValNode', workNode:hasNilAccess(  ), parent, node)
   local function createparamList(  )
   
      local list = node:get_paramList()
      local expList = {}
      for __index, exp in ipairs( list ) do
         local newConv
         
         
         do
            local newState = State.new(convNode)
            local newNode = Node_createFromNode( exp, convNode, newState )
            do
               local nilAccNode = newState:get_nilAccNode()
               if nilAccNode ~= nil then
                  newConv = nilAccNode
               else
                  newConv = newNode
               end
            end
            
         end
         
         
         table.insert( expList, newConv )
      end
      
      return expList
   end
   convNode:set_paramList( createparamList(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.NewAlgeValNode_createFromNode = NewAlgeValNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_NewAlgeVal()] = function ( node, workParent, state )

   return NewAlgeValNode_createFromNode( node, workParent, state )
end




local LuneKindNode = {}



setmetatable( LuneKindNode, { __index = ExpNode } )
_moduleObj.LuneKindNode = LuneKindNode
function LuneKindNode:processFilter( filter, opt )

   filter:processLuneKind( self, opt )
end
function LuneKindNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LuneKindNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LuneKindNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.exp = dummyExpNode
   
   
end
function LuneKindNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function LuneKindNode.setmeta( obj )
  setmetatable( obj, { __index = LuneKindNode  } )
end
function LuneKindNode:get_orgNode()
   return self.orgNode
end
function LuneKindNode:get_exp()
   return self.exp
end
function LuneKindNode:set_exp( exp )
   self.exp = exp
end

local function LuneKindNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LuneKindNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LuneKindNode') )
   end
   
   local convNode = LuneKindNode.new(node:get_id(), node:get_effectivePos(), 'LuneKindNode', workNode:hasNilAccess(  ), parent, node)
   local function createexp(  )
   
      local child = node:get_exp()
      local paramNode = Node_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_exp( createexp(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LuneKindNode_createFromNode = LuneKindNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LuneKind()] = function ( node, workParent, state )

   return LuneKindNode_createFromNode( node, workParent, state )
end




local LiteralNilNode = {}



setmetatable( LiteralNilNode, { __index = ExpNode } )
_moduleObj.LiteralNilNode = LiteralNilNode
function LiteralNilNode:processFilter( filter, opt )

   filter:processLiteralNil( self, opt )
end
function LiteralNilNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralNilNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralNilNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   
   
end
function LiteralNilNode:visit( visitor, depth )

   
   return true
end
function LiteralNilNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralNilNode  } )
end
function LiteralNilNode:get_orgNode()
   return self.orgNode
end

local function LiteralNilNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralNilNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralNilNode') )
   end
   
   local convNode = LiteralNilNode.new(node:get_id(), node:get_effectivePos(), 'LiteralNilNode', workNode:hasNilAccess(  ), parent, node)
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralNilNode_createFromNode = LiteralNilNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralNil()] = function ( node, workParent, state )

   return LiteralNilNode_createFromNode( node, workParent, state )
end




local LiteralCharNode = {}



setmetatable( LiteralCharNode, { __index = ExpNode } )
_moduleObj.LiteralCharNode = LiteralCharNode
function LiteralCharNode:processFilter( filter, opt )

   filter:processLiteralChar( self, opt )
end
function LiteralCharNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralCharNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralCharNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   
   
end
function LiteralCharNode:visit( visitor, depth )

   
   return true
end
function LiteralCharNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralCharNode  } )
end
function LiteralCharNode:get_orgNode()
   return self.orgNode
end

local function LiteralCharNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralCharNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralCharNode') )
   end
   
   local convNode = LiteralCharNode.new(node:get_id(), node:get_effectivePos(), 'LiteralCharNode', workNode:hasNilAccess(  ), parent, node)
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralCharNode_createFromNode = LiteralCharNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralChar()] = function ( node, workParent, state )

   return LiteralCharNode_createFromNode( node, workParent, state )
end




local LiteralIntNode = {}



setmetatable( LiteralIntNode, { __index = ExpNode } )
_moduleObj.LiteralIntNode = LiteralIntNode
function LiteralIntNode:processFilter( filter, opt )

   filter:processLiteralInt( self, opt )
end
function LiteralIntNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralIntNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralIntNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   
   
end
function LiteralIntNode:visit( visitor, depth )

   
   return true
end
function LiteralIntNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralIntNode  } )
end
function LiteralIntNode:get_orgNode()
   return self.orgNode
end

local function LiteralIntNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralIntNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralIntNode') )
   end
   
   local convNode = LiteralIntNode.new(node:get_id(), node:get_effectivePos(), 'LiteralIntNode', workNode:hasNilAccess(  ), parent, node)
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralIntNode_createFromNode = LiteralIntNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralInt()] = function ( node, workParent, state )

   return LiteralIntNode_createFromNode( node, workParent, state )
end




local LiteralRealNode = {}



setmetatable( LiteralRealNode, { __index = ExpNode } )
_moduleObj.LiteralRealNode = LiteralRealNode
function LiteralRealNode:processFilter( filter, opt )

   filter:processLiteralReal( self, opt )
end
function LiteralRealNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralRealNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralRealNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   
   
end
function LiteralRealNode:visit( visitor, depth )

   
   return true
end
function LiteralRealNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralRealNode  } )
end
function LiteralRealNode:get_orgNode()
   return self.orgNode
end

local function LiteralRealNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralRealNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralRealNode') )
   end
   
   local convNode = LiteralRealNode.new(node:get_id(), node:get_effectivePos(), 'LiteralRealNode', workNode:hasNilAccess(  ), parent, node)
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralRealNode_createFromNode = LiteralRealNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralReal()] = function ( node, workParent, state )

   return LiteralRealNode_createFromNode( node, workParent, state )
end




local LiteralArrayNode = {}



setmetatable( LiteralArrayNode, { __index = ExpNode } )
_moduleObj.LiteralArrayNode = LiteralArrayNode
function LiteralArrayNode:processFilter( filter, opt )

   filter:processLiteralArray( self, opt )
end
function LiteralArrayNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralArrayNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralArrayNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.expList = nil
   
   
end
function LiteralArrayNode:visit( visitor, depth )

   do
      do
         local child = self.expList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'expList', depth )
               if _switchExp == Nodes.NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == Nodes.NodeVisitMode.End then
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
function LiteralArrayNode:get_orgNode()
   return self.orgNode
end
function LiteralArrayNode:get_expList()
   return self.expList
end
function LiteralArrayNode:set_expList( expList )
   self.expList = expList
end

local function LiteralArrayNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralArrayNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralArrayNode') )
   end
   
   local convNode = LiteralArrayNode.new(node:get_id(), node:get_effectivePos(), 'LiteralArrayNode', workNode:hasNilAccess(  ), parent, node)
   local function createexpList(  )
   
      local child = node:get_expList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      local paramNode = ExpListNode_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_expList( createexpList(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralArrayNode_createFromNode = LiteralArrayNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralArray()] = function ( node, workParent, state )

   return LiteralArrayNode_createFromNode( node, workParent, state )
end




local LiteralListNode = {}



setmetatable( LiteralListNode, { __index = ExpNode } )
_moduleObj.LiteralListNode = LiteralListNode
function LiteralListNode:processFilter( filter, opt )

   filter:processLiteralList( self, opt )
end
function LiteralListNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralListNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralListNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.expList = nil
   
   
end
function LiteralListNode:visit( visitor, depth )

   do
      do
         local child = self.expList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'expList', depth )
               if _switchExp == Nodes.NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == Nodes.NodeVisitMode.End then
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
function LiteralListNode:get_orgNode()
   return self.orgNode
end
function LiteralListNode:get_expList()
   return self.expList
end
function LiteralListNode:set_expList( expList )
   self.expList = expList
end

local function LiteralListNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralListNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralListNode') )
   end
   
   local convNode = LiteralListNode.new(node:get_id(), node:get_effectivePos(), 'LiteralListNode', workNode:hasNilAccess(  ), parent, node)
   local function createexpList(  )
   
      local child = node:get_expList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      local paramNode = ExpListNode_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_expList( createexpList(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralListNode_createFromNode = LiteralListNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralList()] = function ( node, workParent, state )

   return LiteralListNode_createFromNode( node, workParent, state )
end




local LiteralSetNode = {}



setmetatable( LiteralSetNode, { __index = ExpNode } )
_moduleObj.LiteralSetNode = LiteralSetNode
function LiteralSetNode:processFilter( filter, opt )

   filter:processLiteralSet( self, opt )
end
function LiteralSetNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralSetNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralSetNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.expList = nil
   
   
end
function LiteralSetNode:visit( visitor, depth )

   do
      do
         local child = self.expList
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'expList', depth )
               if _switchExp == Nodes.NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == Nodes.NodeVisitMode.End then
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
function LiteralSetNode:get_orgNode()
   return self.orgNode
end
function LiteralSetNode:get_expList()
   return self.expList
end
function LiteralSetNode:set_expList( expList )
   self.expList = expList
end

local function LiteralSetNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralSetNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralSetNode') )
   end
   
   local convNode = LiteralSetNode.new(node:get_id(), node:get_effectivePos(), 'LiteralSetNode', workNode:hasNilAccess(  ), parent, node)
   local function createexpList(  )
   
      local child = node:get_expList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      local paramNode = ExpListNode_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_expList( createexpList(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralSetNode_createFromNode = LiteralSetNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralSet()] = function ( node, workParent, state )

   return LiteralSetNode_createFromNode( node, workParent, state )
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


local LiteralMapNode = {}



setmetatable( LiteralMapNode, { __index = ExpNode } )
_moduleObj.LiteralMapNode = LiteralMapNode
function LiteralMapNode:processFilter( filter, opt )

   filter:processLiteralMap( self, opt )
end
function LiteralMapNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralMapNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralMapNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.pairList = {}
   
   
end
function LiteralMapNode:visit( visitor, depth )

   
   return true
end
function LiteralMapNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralMapNode  } )
end
function LiteralMapNode:get_orgNode()
   return self.orgNode
end
function LiteralMapNode:get_pairList()
   return self.pairList
end
function LiteralMapNode:set_pairList( pairList )
   self.pairList = pairList
end



local function LiteralMapNode_createFromNode( workNode, parent, state )
   local __func__ = '@lune.@base.@ConvNodes.LiteralMapNode_createFromNode'

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralMapNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s -- %s", tostring( workNode:get_kind()), __func__) )
   end
   
   local pairList = {}
   local convNode = LiteralMapNode.new(node:get_id(), node:get_effectivePos(), "LiteralMapNode", false, parent, node)
   for __index, item in ipairs( node:get_pairList() ) do
      local key
      
      
      do
         local newState = State.new(convNode)
         local newNode = Node_createFromNode( item:get_key(), convNode, newState )
         do
            local nilAccNode = newState:get_nilAccNode()
            if nilAccNode ~= nil then
               key = nilAccNode
            else
               key = newNode
            end
         end
         
      end
      
      
      local val
      
      
      do
         local newState = State.new(convNode)
         local newNode = Node_createFromNode( item:get_val(), convNode, newState )
         do
            local nilAccNode = newState:get_nilAccNode()
            if nilAccNode ~= nil then
               val = nilAccNode
            else
               val = newNode
            end
         end
         
      end
      
      
      table.insert( pairList, PairItem.new(key, val) )
   end
   
   convNode:set_pairList( pairList )
   
   state:setNilAcc( convNode, parent )
   
   return convNode
end
_moduleObj.LiteralMapNode_createFromNode = LiteralMapNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralMap()] = function ( node, parent, state )

   return LiteralMapNode_createFromNode( node, parent, state )
end


local LiteralStringNode = {}



setmetatable( LiteralStringNode, { __index = ExpNode } )
_moduleObj.LiteralStringNode = LiteralStringNode
function LiteralStringNode:processFilter( filter, opt )

   filter:processLiteralString( self, opt )
end
function LiteralStringNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralStringNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralStringNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   self.dddParam = nil
   
   
end
function LiteralStringNode:visit( visitor, depth )

   do
      do
         local child = self.dddParam
         if child ~= nil then
            do
               local _switchExp = visitor( child, self, 'dddParam', depth )
               if _switchExp == Nodes.NodeVisitMode.Child then
                  if not child:visit( visitor, depth + 1 ) then
                     return false
                  end
                  
               elseif _switchExp == Nodes.NodeVisitMode.End then
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
function LiteralStringNode:get_orgNode()
   return self.orgNode
end
function LiteralStringNode:get_dddParam()
   return self.dddParam
end
function LiteralStringNode:set_dddParam( dddParam )
   self.dddParam = dddParam
end

local function LiteralStringNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralStringNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralStringNode') )
   end
   
   local convNode = LiteralStringNode.new(node:get_id(), node:get_effectivePos(), 'LiteralStringNode', workNode:hasNilAccess(  ), parent, node)
   local function createdddParam(  )
   
      local child = node:get_dddParam()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      local paramNode = ExpListNode_createFromNode( child, convNode, state )
      return paramNode
   end
   convNode:set_dddParam( createdddParam(  ) )
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralStringNode_createFromNode = LiteralStringNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralString()] = function ( node, workParent, state )

   return LiteralStringNode_createFromNode( node, workParent, state )
end




local LiteralBoolNode = {}



setmetatable( LiteralBoolNode, { __index = ExpNode } )
_moduleObj.LiteralBoolNode = LiteralBoolNode
function LiteralBoolNode:processFilter( filter, opt )

   filter:processLiteralBool( self, opt )
end
function LiteralBoolNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralBoolNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralBoolNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   
   
end
function LiteralBoolNode:visit( visitor, depth )

   
   return true
end
function LiteralBoolNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralBoolNode  } )
end
function LiteralBoolNode:get_orgNode()
   return self.orgNode
end

local function LiteralBoolNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralBoolNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralBoolNode') )
   end
   
   local convNode = LiteralBoolNode.new(node:get_id(), node:get_effectivePos(), 'LiteralBoolNode', workNode:hasNilAccess(  ), parent, node)
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralBoolNode_createFromNode = LiteralBoolNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralBool()] = function ( node, workParent, state )

   return LiteralBoolNode_createFromNode( node, workParent, state )
end




local LiteralSymbolNode = {}



setmetatable( LiteralSymbolNode, { __index = ExpNode } )
_moduleObj.LiteralSymbolNode = LiteralSymbolNode
function LiteralSymbolNode:processFilter( filter, opt )

   filter:processLiteralSymbol( self, opt )
end
function LiteralSymbolNode.new( id, pos, kind, hasNilAcc, parent, orgNode )
   local obj = {}
   LiteralSymbolNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent, orgNode ); end
   return obj
end
function LiteralSymbolNode:__init(id, pos, kind, hasNilAcc, parent, orgNode) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.orgNode = orgNode
   
   
end
function LiteralSymbolNode:visit( visitor, depth )

   
   return true
end
function LiteralSymbolNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralSymbolNode  } )
end
function LiteralSymbolNode:get_orgNode()
   return self.orgNode
end

local function LiteralSymbolNode_createFromNode( workNode, parent, state )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralSymbolNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralSymbolNode') )
   end
   
   local convNode = LiteralSymbolNode.new(node:get_id(), node:get_effectivePos(), 'LiteralSymbolNode', workNode:hasNilAccess(  ), parent, node)
   
   state:setNilAcc( convNode, parent )
   return convNode
end
_moduleObj.LiteralSymbolNode_createFromNode = LiteralSymbolNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralSymbol()] = function ( node, workParent, state )

   return LiteralSymbolNode_createFromNode( node, workParent, state )
end





local NoneNilAccNode = {}



setmetatable( NoneNilAccNode, { __index = ExpNode } )
_moduleObj.NoneNilAccNode = NoneNilAccNode
function NoneNilAccNode:processFilter( filter, opt )

   filter:processNoneNilAcc( self, opt )
end
function NoneNilAccNode.new( id, pos, kind, hasNilAcc, parent )
   local obj = {}
   NoneNilAccNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent ); end
   return obj
end
function NoneNilAccNode:__init(id, pos, kind, hasNilAcc, parent) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   
end
function NoneNilAccNode:visit( visitor, depth )

   
   return true
end
function NoneNilAccNode.setmeta( obj )
  setmetatable( obj, { __index = NoneNilAccNode  } )
end




local CallExtNode = {}



setmetatable( CallExtNode, { __index = ExpNode } )
_moduleObj.CallExtNode = CallExtNode
function CallExtNode:processFilter( filter, opt )

   filter:processCallExt( self, opt )
end
function CallExtNode.new( id, pos, kind, hasNilAcc, parent )
   local obj = {}
   CallExtNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent ); end
   return obj
end
function CallExtNode:__init(id, pos, kind, hasNilAcc, parent) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.exp = dummyExpNode
   
   
end
function CallExtNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function CallExtNode.setmeta( obj )
  setmetatable( obj, { __index = CallExtNode  } )
end
function CallExtNode:get_exp()
   return self.exp
end
function CallExtNode:set_exp( exp )
   self.exp = exp
end




local CondStackValNode = {}



setmetatable( CondStackValNode, { __index = ExpNode } )
_moduleObj.CondStackValNode = CondStackValNode
function CondStackValNode:processFilter( filter, opt )

   filter:processCondStackVal( self, opt )
end
function CondStackValNode.new( id, pos, kind, hasNilAcc, parent )
   local obj = {}
   CondStackValNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent ); end
   return obj
end
function CondStackValNode:__init(id, pos, kind, hasNilAcc, parent) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.exp = dummyExpNode
   
   
end
function CondStackValNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function CondStackValNode.setmeta( obj )
  setmetatable( obj, { __index = CondStackValNode  } )
end
function CondStackValNode:get_exp()
   return self.exp
end
function CondStackValNode:set_exp( exp )
   self.exp = exp
end




local GetAtNode = {}



setmetatable( GetAtNode, { __index = ExpNode } )
_moduleObj.GetAtNode = GetAtNode
function GetAtNode:processFilter( filter, opt )

   filter:processGetAt( self, opt )
end
function GetAtNode.new( id, pos, kind, hasNilAcc, parent )
   local obj = {}
   GetAtNode.setmeta( obj )
   if obj.__init then obj:__init( id, pos, kind, hasNilAcc, parent ); end
   return obj
end
function GetAtNode:__init(id, pos, kind, hasNilAcc, parent) 
   ExpNode.__init( self,id, pos, kind, hasNilAcc, parent)
   
   self.exp = dummyExpNode
   self.index = dummyExpNode
   
   
end
function GetAtNode:visit( visitor, depth )

   do
      local child = self.exp
      do
         local _switchExp = visitor( child, self, 'exp', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   do
      local child = self.index
      do
         local _switchExp = visitor( child, self, 'index', depth )
         if _switchExp == Nodes.NodeVisitMode.Child then
            if not child:visit( visitor, depth + 1 ) then
               return false
            end
            
         elseif _switchExp == Nodes.NodeVisitMode.End then
            return false
         end
      end
      
      
   end
   
   
   
   return true
end
function GetAtNode.setmeta( obj )
  setmetatable( obj, { __index = GetAtNode  } )
end
function GetAtNode:get_exp()
   return self.exp
end
function GetAtNode:set_exp( exp )
   self.exp = exp
end
function GetAtNode:get_index()
   return self.index
end
function GetAtNode:set_index( index )
   self.index = index
end




local function convertNodes( targetNode )

   do
      local createFunc = nodeKind2createFromFunc[targetNode:get_kind()]
      if createFunc ~= nil then
         local state = State.new(dummyExpNode)
         local workNode = createFunc( targetNode, dummyExpNode, state )
         do
            local nilAccNode = state:get_nilAccNode()
            if nilAccNode ~= nil then
               if nilAccNode:get_acc() == dummyExpNode then
                  nilAccNode:set_acc( workNode )
               end
               
               return nilAccNode
            end
         end
         
         return workNode
      end
   end
   
   return nil
end
_moduleObj.convertNodes = convertNodes





return _moduleObj
