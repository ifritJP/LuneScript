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
function Node.new( id, pos, kind )
   local obj = {}
   Node.setmeta( obj )
   if obj.__init then
      obj:__init( id, pos, kind )
   end
   return obj
end
function Node:__init( id, pos, kind )

   self.id = id
   self.pos = pos
   self.kind = kind
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


local ExpNode = {}
setmetatable( ExpNode, { __index = Node } )
_moduleObj.ExpNode = ExpNode
function ExpNode.setmeta( obj )
  setmetatable( obj, { __index = ExpNode  } )
end
function ExpNode.new( __superarg1, __superarg2, __superarg3 )
   local obj = {}
   ExpNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3 )
   end
   return obj
end
function ExpNode:__init( __superarg1, __superarg2, __superarg3 )

   Node.__init( self, __superarg1, __superarg2, __superarg3 )
end



local nodeKind2createFromFunc = {}

function Node.createFrom( node )

   local func = nodeKind2createFromFunc[node:get_kind()]
   if  nil == func then
      local _func = func
   
      Util.err( string.format( "not support -- %s", Nodes.getNodeKindName( node:get_kind() )) )
   end
   
   return func( node )
end


local function Node_createFromNode( node )

   return Node.createFrom( node )
end




local ExpListNode = {}



setmetatable( ExpListNode, { __index = ExpNode } )
_moduleObj.ExpListNode = ExpListNode
function ExpListNode:processFilter( filter, opt )

   filter:processExpList( self, opt )
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
function ExpListNode.new( __superarg1, __superarg2, __superarg3,orgNode, expList )
   local obj = {}
   ExpListNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )
   end
   return obj
end
function ExpListNode:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.expList = expList
end
function ExpListNode:get_orgNode()
   return self.orgNode
end
function ExpListNode:get_expList()
   return self.expList
end

local function ExpListNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpListNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpListNode') )
   end
   
   local function createexpList(  )
   
      local list = node:get_expList()
      local expList = {}
      for __index, exp in ipairs( list ) do
         table.insert( expList, Node.createFrom( exp ) )
      end
      
      return expList
   end
   
   return ExpListNode.new(node:get_id(), node:get_pos(), 'ExpListNode', node, createexpList(  ))
end
_moduleObj.ExpListNode_createFromNode = ExpListNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpList()] = function ( node )

   return ExpListNode_createFromNode( node )
end




local ExpNewNode = {}



setmetatable( ExpNewNode, { __index = ExpNode } )
_moduleObj.ExpNewNode = ExpNewNode
function ExpNewNode:processFilter( filter, opt )

   filter:processExpNew( self, opt )
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
function ExpNewNode.new( __superarg1, __superarg2, __superarg3,orgNode, argList )
   local obj = {}
   ExpNewNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, argList )
   end
   return obj
end
function ExpNewNode:__init( __superarg1, __superarg2, __superarg3,orgNode, argList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.argList = argList
end
function ExpNewNode:get_orgNode()
   return self.orgNode
end
function ExpNewNode:get_argList()
   return self.argList
end

local function ExpNewNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpNewNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpNewNode') )
   end
   
   local function createargList(  )
   
      local child = node:get_argList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      return ExpListNode_createFromNode( child )
   end
   
   return ExpNewNode.new(node:get_id(), node:get_pos(), 'ExpNewNode', node, createargList(  ))
end
_moduleObj.ExpNewNode_createFromNode = ExpNewNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpNew()] = function ( node )

   return ExpNewNode_createFromNode( node )
end




local ExpUnwrapNode = {}



setmetatable( ExpUnwrapNode, { __index = ExpNode } )
_moduleObj.ExpUnwrapNode = ExpUnwrapNode
function ExpUnwrapNode:processFilter( filter, opt )

   filter:processExpUnwrap( self, opt )
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
function ExpUnwrapNode.new( __superarg1, __superarg2, __superarg3,orgNode, exp, default )
   local obj = {}
   ExpUnwrapNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, exp, default )
   end
   return obj
end
function ExpUnwrapNode:__init( __superarg1, __superarg2, __superarg3,orgNode, exp, default )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.exp = exp
   self.default = default
end
function ExpUnwrapNode:get_orgNode()
   return self.orgNode
end
function ExpUnwrapNode:get_exp()
   return self.exp
end
function ExpUnwrapNode:get_default()
   return self.default
end

local function ExpUnwrapNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpUnwrapNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpUnwrapNode') )
   end
   
   local function createexp(  )
   
      local child = node:get_exp()
      return Node_createFromNode( child )
   end
   local function createdefault(  )
   
      local child = node:get_default()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      return Node_createFromNode( child )
   end
   
   return ExpUnwrapNode.new(node:get_id(), node:get_pos(), 'ExpUnwrapNode', node, createexp(  ), createdefault(  ))
end
_moduleObj.ExpUnwrapNode_createFromNode = ExpUnwrapNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpUnwrap()] = function ( node )

   return ExpUnwrapNode_createFromNode( node )
end




local ExpRefNode = {}



setmetatable( ExpRefNode, { __index = ExpNode } )
_moduleObj.ExpRefNode = ExpRefNode
function ExpRefNode:processFilter( filter, opt )

   filter:processExpRef( self, opt )
end
function ExpRefNode:visit( visitor, depth )

   
   return true
end
function ExpRefNode.setmeta( obj )
  setmetatable( obj, { __index = ExpRefNode  } )
end
function ExpRefNode.new( __superarg1, __superarg2, __superarg3,orgNode )
   local obj = {}
   ExpRefNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode )
   end
   return obj
end
function ExpRefNode:__init( __superarg1, __superarg2, __superarg3,orgNode )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
end
function ExpRefNode:get_orgNode()
   return self.orgNode
end

local function ExpRefNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpRefNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpRefNode') )
   end
   
   
   return ExpRefNode.new(node:get_id(), node:get_pos(), 'ExpRefNode', node)
end
_moduleObj.ExpRefNode_createFromNode = ExpRefNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpRef()] = function ( node )

   return ExpRefNode_createFromNode( node )
end




local ExpOp2Node = {}



setmetatable( ExpOp2Node, { __index = ExpNode } )
_moduleObj.ExpOp2Node = ExpOp2Node
function ExpOp2Node:processFilter( filter, opt )

   filter:processExpOp2( self, opt )
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
function ExpOp2Node.new( __superarg1, __superarg2, __superarg3,orgNode, exp1, exp2 )
   local obj = {}
   ExpOp2Node.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, exp1, exp2 )
   end
   return obj
end
function ExpOp2Node:__init( __superarg1, __superarg2, __superarg3,orgNode, exp1, exp2 )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.exp1 = exp1
   self.exp2 = exp2
end
function ExpOp2Node:get_orgNode()
   return self.orgNode
end
function ExpOp2Node:get_exp1()
   return self.exp1
end
function ExpOp2Node:get_exp2()
   return self.exp2
end

local function ExpOp2Node_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpOp2Node )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpOp2Node') )
   end
   
   local function createexp1(  )
   
      local child = node:get_exp1()
      return Node_createFromNode( child )
   end
   local function createexp2(  )
   
      local child = node:get_exp2()
      return Node_createFromNode( child )
   end
   
   return ExpOp2Node.new(node:get_id(), node:get_pos(), 'ExpOp2Node', node, createexp1(  ), createexp2(  ))
end
_moduleObj.ExpOp2Node_createFromNode = ExpOp2Node_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpOp2()] = function ( node )

   return ExpOp2Node_createFromNode( node )
end




local ExpCastNode = {}



setmetatable( ExpCastNode, { __index = ExpNode } )
_moduleObj.ExpCastNode = ExpCastNode
function ExpCastNode:processFilter( filter, opt )

   filter:processExpCast( self, opt )
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
function ExpCastNode.new( __superarg1, __superarg2, __superarg3,orgNode, exp )
   local obj = {}
   ExpCastNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )
   end
   return obj
end
function ExpCastNode:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.exp = exp
end
function ExpCastNode:get_orgNode()
   return self.orgNode
end
function ExpCastNode:get_exp()
   return self.exp
end

local function ExpCastNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpCastNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpCastNode') )
   end
   
   local function createexp(  )
   
      local child = node:get_exp()
      return Node_createFromNode( child )
   end
   
   return ExpCastNode.new(node:get_id(), node:get_pos(), 'ExpCastNode', node, createexp(  ))
end
_moduleObj.ExpCastNode_createFromNode = ExpCastNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpCast()] = function ( node )

   return ExpCastNode_createFromNode( node )
end




local ExpToDDDNode = {}



setmetatable( ExpToDDDNode, { __index = ExpNode } )
_moduleObj.ExpToDDDNode = ExpToDDDNode
function ExpToDDDNode:processFilter( filter, opt )

   filter:processExpToDDD( self, opt )
end
function ExpToDDDNode:visit( visitor, depth )

   do
      local child = self.expList
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
   
   
   
   return true
end
function ExpToDDDNode.setmeta( obj )
  setmetatable( obj, { __index = ExpToDDDNode  } )
end
function ExpToDDDNode.new( __superarg1, __superarg2, __superarg3,orgNode, expList )
   local obj = {}
   ExpToDDDNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )
   end
   return obj
end
function ExpToDDDNode:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.expList = expList
end
function ExpToDDDNode:get_orgNode()
   return self.orgNode
end
function ExpToDDDNode:get_expList()
   return self.expList
end

local function ExpToDDDNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpToDDDNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpToDDDNode') )
   end
   
   local function createexpList(  )
   
      local child = node:get_expList()
      return ExpListNode_createFromNode( child )
   end
   
   return ExpToDDDNode.new(node:get_id(), node:get_pos(), 'ExpToDDDNode', node, createexpList(  ))
end
_moduleObj.ExpToDDDNode_createFromNode = ExpToDDDNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpToDDD()] = function ( node )

   return ExpToDDDNode_createFromNode( node )
end




local ExpSubDDDNode = {}



setmetatable( ExpSubDDDNode, { __index = ExpNode } )
_moduleObj.ExpSubDDDNode = ExpSubDDDNode
function ExpSubDDDNode:processFilter( filter, opt )

   filter:processExpSubDDD( self, opt )
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
function ExpSubDDDNode.new( __superarg1, __superarg2, __superarg3,orgNode, src )
   local obj = {}
   ExpSubDDDNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, src )
   end
   return obj
end
function ExpSubDDDNode:__init( __superarg1, __superarg2, __superarg3,orgNode, src )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.src = src
end
function ExpSubDDDNode:get_orgNode()
   return self.orgNode
end
function ExpSubDDDNode:get_src()
   return self.src
end

local function ExpSubDDDNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpSubDDDNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpSubDDDNode') )
   end
   
   local function createsrc(  )
   
      local child = node:get_src()
      return Node_createFromNode( child )
   end
   
   return ExpSubDDDNode.new(node:get_id(), node:get_pos(), 'ExpSubDDDNode', node, createsrc(  ))
end
_moduleObj.ExpSubDDDNode_createFromNode = ExpSubDDDNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpSubDDD()] = function ( node )

   return ExpSubDDDNode_createFromNode( node )
end




local ExpOp1Node = {}



setmetatable( ExpOp1Node, { __index = ExpNode } )
_moduleObj.ExpOp1Node = ExpOp1Node
function ExpOp1Node:processFilter( filter, opt )

   filter:processExpOp1( self, opt )
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
function ExpOp1Node.new( __superarg1, __superarg2, __superarg3,orgNode, exp )
   local obj = {}
   ExpOp1Node.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )
   end
   return obj
end
function ExpOp1Node:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.exp = exp
end
function ExpOp1Node:get_orgNode()
   return self.orgNode
end
function ExpOp1Node:get_exp()
   return self.exp
end

local function ExpOp1Node_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpOp1Node )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpOp1Node') )
   end
   
   local function createexp(  )
   
      local child = node:get_exp()
      return Node_createFromNode( child )
   end
   
   return ExpOp1Node.new(node:get_id(), node:get_pos(), 'ExpOp1Node', node, createexp(  ))
end
_moduleObj.ExpOp1Node_createFromNode = ExpOp1Node_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpOp1()] = function ( node )

   return ExpOp1Node_createFromNode( node )
end




local ExpRefItemNode = {}



setmetatable( ExpRefItemNode, { __index = ExpNode } )
_moduleObj.ExpRefItemNode = ExpRefItemNode
function ExpRefItemNode:processFilter( filter, opt )

   filter:processExpRefItem( self, opt )
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
function ExpRefItemNode.new( __superarg1, __superarg2, __superarg3,orgNode, val, index )
   local obj = {}
   ExpRefItemNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, val, index )
   end
   return obj
end
function ExpRefItemNode:__init( __superarg1, __superarg2, __superarg3,orgNode, val, index )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.val = val
   self.index = index
end
function ExpRefItemNode:get_orgNode()
   return self.orgNode
end
function ExpRefItemNode:get_val()
   return self.val
end
function ExpRefItemNode:get_index()
   return self.index
end

local function ExpRefItemNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpRefItemNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpRefItemNode') )
   end
   
   local function createval(  )
   
      local child = node:get_val()
      return Node_createFromNode( child )
   end
   local function createindex(  )
   
      local child = node:get_index()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      return Node_createFromNode( child )
   end
   
   return ExpRefItemNode.new(node:get_id(), node:get_pos(), 'ExpRefItemNode', node, createval(  ), createindex(  ))
end
_moduleObj.ExpRefItemNode_createFromNode = ExpRefItemNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpRefItem()] = function ( node )

   return ExpRefItemNode_createFromNode( node )
end




local ExpCallNode = {}



setmetatable( ExpCallNode, { __index = ExpNode } )
_moduleObj.ExpCallNode = ExpCallNode
function ExpCallNode:processFilter( filter, opt )

   filter:processExpCall( self, opt )
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
function ExpCallNode.new( __superarg1, __superarg2, __superarg3,orgNode, func, argList )
   local obj = {}
   ExpCallNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, func, argList )
   end
   return obj
end
function ExpCallNode:__init( __superarg1, __superarg2, __superarg3,orgNode, func, argList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.func = func
   self.argList = argList
end
function ExpCallNode:get_orgNode()
   return self.orgNode
end
function ExpCallNode:get_func()
   return self.func
end
function ExpCallNode:get_argList()
   return self.argList
end

local function ExpCallNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpCallNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpCallNode') )
   end
   
   local function createfunc(  )
   
      local child = node:get_func()
      return Node_createFromNode( child )
   end
   local function createargList(  )
   
      local child = node:get_argList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      return ExpListNode_createFromNode( child )
   end
   
   return ExpCallNode.new(node:get_id(), node:get_pos(), 'ExpCallNode', node, createfunc(  ), createargList(  ))
end
_moduleObj.ExpCallNode_createFromNode = ExpCallNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpCall()] = function ( node )

   return ExpCallNode_createFromNode( node )
end




local ExpMRetNode = {}



setmetatable( ExpMRetNode, { __index = ExpNode } )
_moduleObj.ExpMRetNode = ExpMRetNode
function ExpMRetNode:processFilter( filter, opt )

   filter:processExpMRet( self, opt )
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
function ExpMRetNode.new( __superarg1, __superarg2, __superarg3,orgNode, mRet )
   local obj = {}
   ExpMRetNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, mRet )
   end
   return obj
end
function ExpMRetNode:__init( __superarg1, __superarg2, __superarg3,orgNode, mRet )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.mRet = mRet
end
function ExpMRetNode:get_orgNode()
   return self.orgNode
end
function ExpMRetNode:get_mRet()
   return self.mRet
end

local function ExpMRetNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpMRetNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpMRetNode') )
   end
   
   local function createmRet(  )
   
      local child = node:get_mRet()
      return Node_createFromNode( child )
   end
   
   return ExpMRetNode.new(node:get_id(), node:get_pos(), 'ExpMRetNode', node, createmRet(  ))
end
_moduleObj.ExpMRetNode_createFromNode = ExpMRetNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpMRet()] = function ( node )

   return ExpMRetNode_createFromNode( node )
end




local ExpAccessMRetNode = {}



setmetatable( ExpAccessMRetNode, { __index = ExpNode } )
_moduleObj.ExpAccessMRetNode = ExpAccessMRetNode
function ExpAccessMRetNode:processFilter( filter, opt )

   filter:processExpAccessMRet( self, opt )
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
function ExpAccessMRetNode.new( __superarg1, __superarg2, __superarg3,orgNode, mRet )
   local obj = {}
   ExpAccessMRetNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, mRet )
   end
   return obj
end
function ExpAccessMRetNode:__init( __superarg1, __superarg2, __superarg3,orgNode, mRet )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.mRet = mRet
end
function ExpAccessMRetNode:get_orgNode()
   return self.orgNode
end
function ExpAccessMRetNode:get_mRet()
   return self.mRet
end

local function ExpAccessMRetNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpAccessMRetNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpAccessMRetNode') )
   end
   
   local function createmRet(  )
   
      local child = node:get_mRet()
      return Node_createFromNode( child )
   end
   
   return ExpAccessMRetNode.new(node:get_id(), node:get_pos(), 'ExpAccessMRetNode', node, createmRet(  ))
end
_moduleObj.ExpAccessMRetNode_createFromNode = ExpAccessMRetNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpAccessMRet()] = function ( node )

   return ExpAccessMRetNode_createFromNode( node )
end




local ExpMultiTo1Node = {}



setmetatable( ExpMultiTo1Node, { __index = ExpNode } )
_moduleObj.ExpMultiTo1Node = ExpMultiTo1Node
function ExpMultiTo1Node:processFilter( filter, opt )

   filter:processExpMultiTo1( self, opt )
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
function ExpMultiTo1Node.new( __superarg1, __superarg2, __superarg3,orgNode, exp )
   local obj = {}
   ExpMultiTo1Node.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )
   end
   return obj
end
function ExpMultiTo1Node:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.exp = exp
end
function ExpMultiTo1Node:get_orgNode()
   return self.orgNode
end
function ExpMultiTo1Node:get_exp()
   return self.exp
end

local function ExpMultiTo1Node_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpMultiTo1Node )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpMultiTo1Node') )
   end
   
   local function createexp(  )
   
      local child = node:get_exp()
      return Node_createFromNode( child )
   end
   
   return ExpMultiTo1Node.new(node:get_id(), node:get_pos(), 'ExpMultiTo1Node', node, createexp(  ))
end
_moduleObj.ExpMultiTo1Node_createFromNode = ExpMultiTo1Node_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpMultiTo1()] = function ( node )

   return ExpMultiTo1Node_createFromNode( node )
end




local ExpParenNode = {}



setmetatable( ExpParenNode, { __index = ExpNode } )
_moduleObj.ExpParenNode = ExpParenNode
function ExpParenNode:processFilter( filter, opt )

   filter:processExpParen( self, opt )
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
function ExpParenNode.new( __superarg1, __superarg2, __superarg3,orgNode, exp )
   local obj = {}
   ExpParenNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )
   end
   return obj
end
function ExpParenNode:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.exp = exp
end
function ExpParenNode:get_orgNode()
   return self.orgNode
end
function ExpParenNode:get_exp()
   return self.exp
end

local function ExpParenNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpParenNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpParenNode') )
   end
   
   local function createexp(  )
   
      local child = node:get_exp()
      return Node_createFromNode( child )
   end
   
   return ExpParenNode.new(node:get_id(), node:get_pos(), 'ExpParenNode', node, createexp(  ))
end
_moduleObj.ExpParenNode_createFromNode = ExpParenNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpParen()] = function ( node )

   return ExpParenNode_createFromNode( node )
end




local ExpOmitEnumNode = {}



setmetatable( ExpOmitEnumNode, { __index = ExpNode } )
_moduleObj.ExpOmitEnumNode = ExpOmitEnumNode
function ExpOmitEnumNode:processFilter( filter, opt )

   filter:processExpOmitEnum( self, opt )
end
function ExpOmitEnumNode:visit( visitor, depth )

   
   return true
end
function ExpOmitEnumNode.setmeta( obj )
  setmetatable( obj, { __index = ExpOmitEnumNode  } )
end
function ExpOmitEnumNode.new( __superarg1, __superarg2, __superarg3,orgNode )
   local obj = {}
   ExpOmitEnumNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode )
   end
   return obj
end
function ExpOmitEnumNode:__init( __superarg1, __superarg2, __superarg3,orgNode )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
end
function ExpOmitEnumNode:get_orgNode()
   return self.orgNode
end

local function ExpOmitEnumNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.ExpOmitEnumNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'ExpOmitEnumNode') )
   end
   
   
   return ExpOmitEnumNode.new(node:get_id(), node:get_pos(), 'ExpOmitEnumNode', node)
end
_moduleObj.ExpOmitEnumNode_createFromNode = ExpOmitEnumNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_ExpOmitEnum()] = function ( node )

   return ExpOmitEnumNode_createFromNode( node )
end




local RefFieldNode = {}



setmetatable( RefFieldNode, { __index = ExpNode } )
_moduleObj.RefFieldNode = RefFieldNode
function RefFieldNode:processFilter( filter, opt )

   filter:processRefField( self, opt )
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
function RefFieldNode.new( __superarg1, __superarg2, __superarg3,orgNode, prefix )
   local obj = {}
   RefFieldNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, prefix )
   end
   return obj
end
function RefFieldNode:__init( __superarg1, __superarg2, __superarg3,orgNode, prefix )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.prefix = prefix
end
function RefFieldNode:get_orgNode()
   return self.orgNode
end
function RefFieldNode:get_prefix()
   return self.prefix
end

local function RefFieldNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.RefFieldNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'RefFieldNode') )
   end
   
   local function createprefix(  )
   
      local child = node:get_prefix()
      return Node_createFromNode( child )
   end
   
   return RefFieldNode.new(node:get_id(), node:get_pos(), 'RefFieldNode', node, createprefix(  ))
end
_moduleObj.RefFieldNode_createFromNode = RefFieldNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_RefField()] = function ( node )

   return RefFieldNode_createFromNode( node )
end




local GetFieldNode = {}



setmetatable( GetFieldNode, { __index = ExpNode } )
_moduleObj.GetFieldNode = GetFieldNode
function GetFieldNode:processFilter( filter, opt )

   filter:processGetField( self, opt )
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
function GetFieldNode.new( __superarg1, __superarg2, __superarg3,orgNode, prefix )
   local obj = {}
   GetFieldNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, prefix )
   end
   return obj
end
function GetFieldNode:__init( __superarg1, __superarg2, __superarg3,orgNode, prefix )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.prefix = prefix
end
function GetFieldNode:get_orgNode()
   return self.orgNode
end
function GetFieldNode:get_prefix()
   return self.prefix
end

local function GetFieldNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.GetFieldNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'GetFieldNode') )
   end
   
   local function createprefix(  )
   
      local child = node:get_prefix()
      return Node_createFromNode( child )
   end
   
   return GetFieldNode.new(node:get_id(), node:get_pos(), 'GetFieldNode', node, createprefix(  ))
end
_moduleObj.GetFieldNode_createFromNode = GetFieldNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_GetField()] = function ( node )

   return GetFieldNode_createFromNode( node )
end




local DeclFuncNode = {}



setmetatable( DeclFuncNode, { __index = ExpNode } )
_moduleObj.DeclFuncNode = DeclFuncNode
function DeclFuncNode:processFilter( filter, opt )

   filter:processDeclFunc( self, opt )
end
function DeclFuncNode:visit( visitor, depth )

   
   return true
end
function DeclFuncNode.setmeta( obj )
  setmetatable( obj, { __index = DeclFuncNode  } )
end
function DeclFuncNode.new( __superarg1, __superarg2, __superarg3,orgNode )
   local obj = {}
   DeclFuncNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode )
   end
   return obj
end
function DeclFuncNode:__init( __superarg1, __superarg2, __superarg3,orgNode )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
end
function DeclFuncNode:get_orgNode()
   return self.orgNode
end

local function DeclFuncNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.DeclFuncNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'DeclFuncNode') )
   end
   
   
   return DeclFuncNode.new(node:get_id(), node:get_pos(), 'DeclFuncNode', node)
end
_moduleObj.DeclFuncNode_createFromNode = DeclFuncNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_DeclFunc()] = function ( node )

   return DeclFuncNode_createFromNode( node )
end




local NewAlgeValNode = {}



setmetatable( NewAlgeValNode, { __index = ExpNode } )
_moduleObj.NewAlgeValNode = NewAlgeValNode
function NewAlgeValNode:processFilter( filter, opt )

   filter:processNewAlgeVal( self, opt )
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
function NewAlgeValNode.new( __superarg1, __superarg2, __superarg3,orgNode, paramList )
   local obj = {}
   NewAlgeValNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, paramList )
   end
   return obj
end
function NewAlgeValNode:__init( __superarg1, __superarg2, __superarg3,orgNode, paramList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.paramList = paramList
end
function NewAlgeValNode:get_orgNode()
   return self.orgNode
end
function NewAlgeValNode:get_paramList()
   return self.paramList
end

local function NewAlgeValNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.NewAlgeValNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'NewAlgeValNode') )
   end
   
   local function createparamList(  )
   
      local list = node:get_paramList()
      local expList = {}
      for __index, exp in ipairs( list ) do
         table.insert( expList, Node.createFrom( exp ) )
      end
      
      return expList
   end
   
   return NewAlgeValNode.new(node:get_id(), node:get_pos(), 'NewAlgeValNode', node, createparamList(  ))
end
_moduleObj.NewAlgeValNode_createFromNode = NewAlgeValNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_NewAlgeVal()] = function ( node )

   return NewAlgeValNode_createFromNode( node )
end




local LuneKindNode = {}



setmetatable( LuneKindNode, { __index = ExpNode } )
_moduleObj.LuneKindNode = LuneKindNode
function LuneKindNode:processFilter( filter, opt )

   filter:processLuneKind( self, opt )
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
function LuneKindNode.new( __superarg1, __superarg2, __superarg3,orgNode, exp )
   local obj = {}
   LuneKindNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )
   end
   return obj
end
function LuneKindNode:__init( __superarg1, __superarg2, __superarg3,orgNode, exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.exp = exp
end
function LuneKindNode:get_orgNode()
   return self.orgNode
end
function LuneKindNode:get_exp()
   return self.exp
end

local function LuneKindNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LuneKindNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LuneKindNode') )
   end
   
   local function createexp(  )
   
      local child = node:get_exp()
      return Node_createFromNode( child )
   end
   
   return LuneKindNode.new(node:get_id(), node:get_pos(), 'LuneKindNode', node, createexp(  ))
end
_moduleObj.LuneKindNode_createFromNode = LuneKindNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LuneKind()] = function ( node )

   return LuneKindNode_createFromNode( node )
end




local LiteralCharNode = {}



setmetatable( LiteralCharNode, { __index = ExpNode } )
_moduleObj.LiteralCharNode = LiteralCharNode
function LiteralCharNode:processFilter( filter, opt )

   filter:processLiteralChar( self, opt )
end
function LiteralCharNode:visit( visitor, depth )

   
   return true
end
function LiteralCharNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralCharNode  } )
end
function LiteralCharNode.new( __superarg1, __superarg2, __superarg3,orgNode )
   local obj = {}
   LiteralCharNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode )
   end
   return obj
end
function LiteralCharNode:__init( __superarg1, __superarg2, __superarg3,orgNode )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
end
function LiteralCharNode:get_orgNode()
   return self.orgNode
end

local function LiteralCharNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralCharNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralCharNode') )
   end
   
   
   return LiteralCharNode.new(node:get_id(), node:get_pos(), 'LiteralCharNode', node)
end
_moduleObj.LiteralCharNode_createFromNode = LiteralCharNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralChar()] = function ( node )

   return LiteralCharNode_createFromNode( node )
end




local LiteralIntNode = {}



setmetatable( LiteralIntNode, { __index = ExpNode } )
_moduleObj.LiteralIntNode = LiteralIntNode
function LiteralIntNode:processFilter( filter, opt )

   filter:processLiteralInt( self, opt )
end
function LiteralIntNode:visit( visitor, depth )

   
   return true
end
function LiteralIntNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralIntNode  } )
end
function LiteralIntNode.new( __superarg1, __superarg2, __superarg3,orgNode )
   local obj = {}
   LiteralIntNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode )
   end
   return obj
end
function LiteralIntNode:__init( __superarg1, __superarg2, __superarg3,orgNode )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
end
function LiteralIntNode:get_orgNode()
   return self.orgNode
end

local function LiteralIntNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralIntNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralIntNode') )
   end
   
   
   return LiteralIntNode.new(node:get_id(), node:get_pos(), 'LiteralIntNode', node)
end
_moduleObj.LiteralIntNode_createFromNode = LiteralIntNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralInt()] = function ( node )

   return LiteralIntNode_createFromNode( node )
end




local LiteralRealNode = {}



setmetatable( LiteralRealNode, { __index = ExpNode } )
_moduleObj.LiteralRealNode = LiteralRealNode
function LiteralRealNode:processFilter( filter, opt )

   filter:processLiteralReal( self, opt )
end
function LiteralRealNode:visit( visitor, depth )

   
   return true
end
function LiteralRealNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralRealNode  } )
end
function LiteralRealNode.new( __superarg1, __superarg2, __superarg3,orgNode )
   local obj = {}
   LiteralRealNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode )
   end
   return obj
end
function LiteralRealNode:__init( __superarg1, __superarg2, __superarg3,orgNode )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
end
function LiteralRealNode:get_orgNode()
   return self.orgNode
end

local function LiteralRealNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralRealNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralRealNode') )
   end
   
   
   return LiteralRealNode.new(node:get_id(), node:get_pos(), 'LiteralRealNode', node)
end
_moduleObj.LiteralRealNode_createFromNode = LiteralRealNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralReal()] = function ( node )

   return LiteralRealNode_createFromNode( node )
end




local LiteralArrayNode = {}



setmetatable( LiteralArrayNode, { __index = ExpNode } )
_moduleObj.LiteralArrayNode = LiteralArrayNode
function LiteralArrayNode:processFilter( filter, opt )

   filter:processLiteralArray( self, opt )
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
function LiteralArrayNode.new( __superarg1, __superarg2, __superarg3,orgNode, expList )
   local obj = {}
   LiteralArrayNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )
   end
   return obj
end
function LiteralArrayNode:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.expList = expList
end
function LiteralArrayNode:get_orgNode()
   return self.orgNode
end
function LiteralArrayNode:get_expList()
   return self.expList
end

local function LiteralArrayNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralArrayNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralArrayNode') )
   end
   
   local function createexpList(  )
   
      local child = node:get_expList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      return ExpListNode_createFromNode( child )
   end
   
   return LiteralArrayNode.new(node:get_id(), node:get_pos(), 'LiteralArrayNode', node, createexpList(  ))
end
_moduleObj.LiteralArrayNode_createFromNode = LiteralArrayNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralArray()] = function ( node )

   return LiteralArrayNode_createFromNode( node )
end




local LiteralListNode = {}



setmetatable( LiteralListNode, { __index = ExpNode } )
_moduleObj.LiteralListNode = LiteralListNode
function LiteralListNode:processFilter( filter, opt )

   filter:processLiteralList( self, opt )
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
function LiteralListNode.new( __superarg1, __superarg2, __superarg3,orgNode, expList )
   local obj = {}
   LiteralListNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )
   end
   return obj
end
function LiteralListNode:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.expList = expList
end
function LiteralListNode:get_orgNode()
   return self.orgNode
end
function LiteralListNode:get_expList()
   return self.expList
end

local function LiteralListNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralListNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralListNode') )
   end
   
   local function createexpList(  )
   
      local child = node:get_expList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      return ExpListNode_createFromNode( child )
   end
   
   return LiteralListNode.new(node:get_id(), node:get_pos(), 'LiteralListNode', node, createexpList(  ))
end
_moduleObj.LiteralListNode_createFromNode = LiteralListNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralList()] = function ( node )

   return LiteralListNode_createFromNode( node )
end




local LiteralSetNode = {}



setmetatable( LiteralSetNode, { __index = ExpNode } )
_moduleObj.LiteralSetNode = LiteralSetNode
function LiteralSetNode:processFilter( filter, opt )

   filter:processLiteralSet( self, opt )
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
function LiteralSetNode.new( __superarg1, __superarg2, __superarg3,orgNode, expList )
   local obj = {}
   LiteralSetNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )
   end
   return obj
end
function LiteralSetNode:__init( __superarg1, __superarg2, __superarg3,orgNode, expList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.expList = expList
end
function LiteralSetNode:get_orgNode()
   return self.orgNode
end
function LiteralSetNode:get_expList()
   return self.expList
end

local function LiteralSetNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralSetNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralSetNode') )
   end
   
   local function createexpList(  )
   
      local child = node:get_expList()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      return ExpListNode_createFromNode( child )
   end
   
   return LiteralSetNode.new(node:get_id(), node:get_pos(), 'LiteralSetNode', node, createexpList(  ))
end
_moduleObj.LiteralSetNode_createFromNode = LiteralSetNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralSet()] = function ( node )

   return LiteralSetNode_createFromNode( node )
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
function LiteralMapNode:visit( visitor, depth )

   
   return true
end
function LiteralMapNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralMapNode  } )
end
function LiteralMapNode.new( __superarg1, __superarg2, __superarg3,orgNode, pairList )
   local obj = {}
   LiteralMapNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, pairList )
   end
   return obj
end
function LiteralMapNode:__init( __superarg1, __superarg2, __superarg3,orgNode, pairList )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.pairList = pairList
end
function LiteralMapNode:get_orgNode()
   return self.orgNode
end
function LiteralMapNode:get_pairList()
   return self.pairList
end



local function LiteralMapNode_createFromNode( workNode )
   local __func__ = '@lune.@base.@ConvNodes.LiteralMapNode_createFromNode'

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralMapNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s -- %s", tostring( workNode:get_kind()), __func__) )
   end
   
   local pairList = {}
   for __index, item in ipairs( node:get_pairList() ) do
      local key = Node_createFromNode( item:get_key() )
      local val = Node_createFromNode( item:get_val() )
      table.insert( pairList, PairItem.new(key, val) )
   end
   
   return LiteralMapNode.new(node:get_id(), node:get_pos(), "LiteralMapNode", node, pairList)
end
_moduleObj.LiteralMapNode_createFromNode = LiteralMapNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralMap()] = function ( node )

   return Node_createFromNode( node )
end


local LiteralStringNode = {}



setmetatable( LiteralStringNode, { __index = ExpNode } )
_moduleObj.LiteralStringNode = LiteralStringNode
function LiteralStringNode:processFilter( filter, opt )

   filter:processLiteralString( self, opt )
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
function LiteralStringNode.new( __superarg1, __superarg2, __superarg3,orgNode, dddParam )
   local obj = {}
   LiteralStringNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode, dddParam )
   end
   return obj
end
function LiteralStringNode:__init( __superarg1, __superarg2, __superarg3,orgNode, dddParam )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
   self.dddParam = dddParam
end
function LiteralStringNode:get_orgNode()
   return self.orgNode
end
function LiteralStringNode:get_dddParam()
   return self.dddParam
end

local function LiteralStringNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralStringNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralStringNode') )
   end
   
   local function createdddParam(  )
   
      local child = node:get_dddParam()
      if  nil == child then
         local _child = child
      
         return nil
      end
      
      return ExpListNode_createFromNode( child )
   end
   
   return LiteralStringNode.new(node:get_id(), node:get_pos(), 'LiteralStringNode', node, createdddParam(  ))
end
_moduleObj.LiteralStringNode_createFromNode = LiteralStringNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralString()] = function ( node )

   return LiteralStringNode_createFromNode( node )
end




local LiteralBoolNode = {}



setmetatable( LiteralBoolNode, { __index = ExpNode } )
_moduleObj.LiteralBoolNode = LiteralBoolNode
function LiteralBoolNode:processFilter( filter, opt )

   filter:processLiteralBool( self, opt )
end
function LiteralBoolNode:visit( visitor, depth )

   
   return true
end
function LiteralBoolNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralBoolNode  } )
end
function LiteralBoolNode.new( __superarg1, __superarg2, __superarg3,orgNode )
   local obj = {}
   LiteralBoolNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode )
   end
   return obj
end
function LiteralBoolNode:__init( __superarg1, __superarg2, __superarg3,orgNode )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
end
function LiteralBoolNode:get_orgNode()
   return self.orgNode
end

local function LiteralBoolNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralBoolNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralBoolNode') )
   end
   
   
   return LiteralBoolNode.new(node:get_id(), node:get_pos(), 'LiteralBoolNode', node)
end
_moduleObj.LiteralBoolNode_createFromNode = LiteralBoolNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralBool()] = function ( node )

   return LiteralBoolNode_createFromNode( node )
end




local LiteralSymbolNode = {}



setmetatable( LiteralSymbolNode, { __index = ExpNode } )
_moduleObj.LiteralSymbolNode = LiteralSymbolNode
function LiteralSymbolNode:processFilter( filter, opt )

   filter:processLiteralSymbol( self, opt )
end
function LiteralSymbolNode:visit( visitor, depth )

   
   return true
end
function LiteralSymbolNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralSymbolNode  } )
end
function LiteralSymbolNode.new( __superarg1, __superarg2, __superarg3,orgNode )
   local obj = {}
   LiteralSymbolNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,orgNode )
   end
   return obj
end
function LiteralSymbolNode:__init( __superarg1, __superarg2, __superarg3,orgNode )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.orgNode = orgNode
end
function LiteralSymbolNode:get_orgNode()
   return self.orgNode
end

local function LiteralSymbolNode_createFromNode( workNode )

   local node = _lune.__Cast( workNode, 3, Nodes.LiteralSymbolNode )
   if  nil == node then
      local _node = node
   
      Util.err( string.format( "illegal node -- %s, %s", Nodes.getNodeKindName( workNode:get_kind() ), 'LiteralSymbolNode') )
   end
   
   
   return LiteralSymbolNode.new(node:get_id(), node:get_pos(), 'LiteralSymbolNode', node)
end
_moduleObj.LiteralSymbolNode_createFromNode = LiteralSymbolNode_createFromNode
nodeKind2createFromFunc[Nodes.NodeKind.get_LiteralSymbol()] = function ( node )

   return LiteralSymbolNode_createFromNode( node )
end




local NilAccPushNode = {}



setmetatable( NilAccPushNode, { __index = ExpNode } )
_moduleObj.NilAccPushNode = NilAccPushNode
function NilAccPushNode:processFilter( filter, opt )

   filter:processNilAccPush( self, opt )
end
function NilAccPushNode:visit( visitor, depth )

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
function NilAccPushNode.setmeta( obj )
  setmetatable( obj, { __index = NilAccPushNode  } )
end
function NilAccPushNode.new( __superarg1, __superarg2, __superarg3,exp )
   local obj = {}
   NilAccPushNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,exp )
   end
   return obj
end
function NilAccPushNode:__init( __superarg1, __superarg2, __superarg3,exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.exp = exp
end
function NilAccPushNode:get_exp()
   return self.exp
end




local NilAccFinNode = {}



setmetatable( NilAccFinNode, { __index = ExpNode } )
_moduleObj.NilAccFinNode = NilAccFinNode
function NilAccFinNode:processFilter( filter, opt )

   filter:processNilAccFin( self, opt )
end
function NilAccFinNode:visit( visitor, depth )

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
function NilAccFinNode.setmeta( obj )
  setmetatable( obj, { __index = NilAccFinNode  } )
end
function NilAccFinNode.new( __superarg1, __superarg2, __superarg3,exp )
   local obj = {}
   NilAccFinNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,exp )
   end
   return obj
end
function NilAccFinNode:__init( __superarg1, __superarg2, __superarg3,exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.exp = exp
end
function NilAccFinNode:get_exp()
   return self.exp
end




local NilAccFinCallNode = {}



setmetatable( NilAccFinCallNode, { __index = ExpNode } )
_moduleObj.NilAccFinCallNode = NilAccFinCallNode
function NilAccFinCallNode:processFilter( filter, opt )

   filter:processNilAccFinCall( self, opt )
end
function NilAccFinCallNode:visit( visitor, depth )

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
function NilAccFinCallNode.setmeta( obj )
  setmetatable( obj, { __index = NilAccFinCallNode  } )
end
function NilAccFinCallNode.new( __superarg1, __superarg2, __superarg3,exp )
   local obj = {}
   NilAccFinCallNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,exp )
   end
   return obj
end
function NilAccFinCallNode:__init( __superarg1, __superarg2, __superarg3,exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.exp = exp
end
function NilAccFinCallNode:get_exp()
   return self.exp
end




local NilAccLastNode = {}



setmetatable( NilAccLastNode, { __index = ExpNode } )
_moduleObj.NilAccLastNode = NilAccLastNode
function NilAccLastNode:processFilter( filter, opt )

   filter:processNilAccLast( self, opt )
end
function NilAccLastNode:visit( visitor, depth )

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
function NilAccLastNode.setmeta( obj )
  setmetatable( obj, { __index = NilAccLastNode  } )
end
function NilAccLastNode.new( __superarg1, __superarg2, __superarg3,exp )
   local obj = {}
   NilAccLastNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,exp )
   end
   return obj
end
function NilAccLastNode:__init( __superarg1, __superarg2, __superarg3,exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.exp = exp
end
function NilAccLastNode:get_exp()
   return self.exp
end




local CallExtNode = {}



setmetatable( CallExtNode, { __index = ExpNode } )
_moduleObj.CallExtNode = CallExtNode
function CallExtNode:processFilter( filter, opt )

   filter:processCallExt( self, opt )
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
function CallExtNode.new( __superarg1, __superarg2, __superarg3,exp )
   local obj = {}
   CallExtNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,exp )
   end
   return obj
end
function CallExtNode:__init( __superarg1, __superarg2, __superarg3,exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.exp = exp
end
function CallExtNode:get_exp()
   return self.exp
end




local CondPopValNode = {}



setmetatable( CondPopValNode, { __index = ExpNode } )
_moduleObj.CondPopValNode = CondPopValNode
function CondPopValNode:processFilter( filter, opt )

   filter:processCondPopVal( self, opt )
end
function CondPopValNode:visit( visitor, depth )

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
function CondPopValNode.setmeta( obj )
  setmetatable( obj, { __index = CondPopValNode  } )
end
function CondPopValNode.new( __superarg1, __superarg2, __superarg3,exp )
   local obj = {}
   CondPopValNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,exp )
   end
   return obj
end
function CondPopValNode:__init( __superarg1, __superarg2, __superarg3,exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.exp = exp
end
function CondPopValNode:get_exp()
   return self.exp
end




local CondSetValNode = {}



setmetatable( CondSetValNode, { __index = ExpNode } )
_moduleObj.CondSetValNode = CondSetValNode
function CondSetValNode:processFilter( filter, opt )

   filter:processCondSetVal( self, opt )
end
function CondSetValNode:visit( visitor, depth )

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
function CondSetValNode.setmeta( obj )
  setmetatable( obj, { __index = CondSetValNode  } )
end
function CondSetValNode.new( __superarg1, __superarg2, __superarg3,exp )
   local obj = {}
   CondSetValNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,exp )
   end
   return obj
end
function CondSetValNode:__init( __superarg1, __superarg2, __superarg3,exp )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.exp = exp
end
function CondSetValNode:get_exp()
   return self.exp
end




local GetAtNode = {}



setmetatable( GetAtNode, { __index = ExpNode } )
_moduleObj.GetAtNode = GetAtNode
function GetAtNode:processFilter( filter, opt )

   filter:processGetAt( self, opt )
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
function GetAtNode.new( __superarg1, __superarg2, __superarg3,exp, index )
   local obj = {}
   GetAtNode.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3,exp, index )
   end
   return obj
end
function GetAtNode:__init( __superarg1, __superarg2, __superarg3,exp, index )

   ExpNode.__init( self, __superarg1, __superarg2, __superarg3 )
   self.exp = exp
   self.index = index
end
function GetAtNode:get_exp()
   return self.exp
end
function GetAtNode:get_index()
   return self.index
end




local function convertNodes( targetNode )

   targetNode:visit( function ( node, parent, relation, depth )
   
      print( string.format( "%s: %s: %d:%d", string.rep( " ", depth * 3 ) .. relation, Nodes.getNodeKindName( node:get_kind() ), node:get_pos().lineNo, node:get_pos().column) )
      do
         local createFunc = nodeKind2createFromFunc[node:get_kind()]
         if createFunc ~= nil then
            local workNode = createFunc( node )
            print( string.format( "%s: conved %s: %d:%d", string.rep( " ", depth * 3 ) .. relation, workNode:get_kind(), workNode:get_pos().lineNo, workNode:get_pos().column) )
            workNode:visit( function ( convNode, convParent, convRelation, convDepth )
            
               print( string.format( "%s: conved %s: %d:%d", string.rep( " ", convDepth * 3 ) .. convRelation, convNode:get_kind(), convNode:get_pos().lineNo, convNode:get_pos().column) )
               return Nodes.NodeVisitMode.Child
            end, depth )
            if node:get_kind() == Nodes.NodeKind.get_DeclFunc() then
               return Nodes.NodeVisitMode.Child
            end
            
            return Nodes.NodeVisitMode.Next
         else
            return Nodes.NodeVisitMode.Child
         end
      end
      
   end, 0 )
end
_moduleObj.convertNodes = convertNodes





return _moduleObj
