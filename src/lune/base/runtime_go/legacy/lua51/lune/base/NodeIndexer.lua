--lune/base/NodeIndexer.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@NodeIndexer'
local _lune = {}
if _lune4 then
   _lune = _lune4
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

if not _lune4 then
   _lune4 = _lune
end


local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Ast = _lune.loadModule( 'lune.base.Ast' )

local NamespaceInfo = {}
function NamespaceInfo.new( parent, nsType )
   local obj = {}
   NamespaceInfo.setmeta( obj )
   if obj.__init then obj:__init( parent, nsType ); end
   return obj
end
function NamespaceInfo:__init(parent, nsType) 
   self.parent = parent
   self.childCount = 1
   local subId
   
   local depth
   
   if parent ~= nil then
      subId = parent.childCount
      depth = parent.depth + 1
      parent.childCount = parent.childCount + 1
   else
      subId = 0
      depth = 0
   end
   
   self.subId = subId
   self.depth = depth
   self.nsType = nsType
   
   self.idProvMap = {}
end
function NamespaceInfo:getNewId( nodeKind )

   local idProv = self.idProvMap[nodeKind]
   if  nil == idProv then
      local _idProv = idProv
   
      local work = Ast.IdProvider.new(0, 100000)
      self.idProvMap[nodeKind] = work
      idProv = work
   end
   
   return idProv:getNewId(  )
end
function NamespaceInfo:getIdTxt(  )

   do
      local parent = self.parent
      if parent ~= nil then
         if self.depth < 4 then
            return string.format( "%s_%s", parent:getIdTxt(  ), self:get_nsType():get_rawTxt())
         else
          
            return string.format( "%s%s", parent:getIdTxt(  ), string.format( "_%d", self.subId))
         end
         
      else
         return ""
      end
   end
   
end
function NamespaceInfo.setmeta( obj )
  setmetatable( obj, { __index = NamespaceInfo  } )
end
function NamespaceInfo:get_subId()
   return self.subId
end
function NamespaceInfo:get_childCount()
   return self.childCount
end
function NamespaceInfo:get_nsType()
   return self.nsType
end

local declNameSpaceNodeKindSet = {[Nodes.NodeKind.get_DeclClass()] = true, [Nodes.NodeKind.get_DeclConstr()] = true, [Nodes.NodeKind.get_DeclFunc()] = true, [Nodes.NodeKind.get_ProtoMethod()] = true, [Nodes.NodeKind.get_DeclMethod()] = true, [Nodes.NodeKind.get_DeclEnum()] = true, [Nodes.NodeKind.get_DeclAlge()] = true, [Nodes.NodeKind.get_DeclMacro()] = true}

local Index = {}
_moduleObj.Index = Index
function Index.new( nsInfo, index )
   local obj = {}
   Index.setmeta( obj )
   if obj.__init then obj:__init( nsInfo, index ); end
   return obj
end
function Index:__init(nsInfo, index) 
   self.nsInfo = nsInfo
   self.index = index
end
function Index:getIdTxt(  )

   return string.format( "%s_%d", self.nsInfo:getIdTxt(  ), self.index)
end
function Index.setmeta( obj )
  setmetatable( obj, { __index = Index  } )
end
function Index:get_nsInfo()
   return self.nsInfo
end
function Index:get_index()
   return self.index
end

local Indexer = {}
_moduleObj.Indexer = Indexer
function Indexer.new( processInfo )
   local obj = {}
   Indexer.setmeta( obj )
   if obj.__init then obj:__init( processInfo ); end
   return obj
end
function Indexer:__init(processInfo) 
   self.nsType2nsInfo = {}
   self.processInfo = processInfo
   self.nsInfo = NamespaceInfo.new(nil, Ast.headTypeInfo)
   self.node2Index = {}
   self.targetNodeKindSet = {}
end
function Indexer:getIndex( node )

   return _lune.unwrap( self.node2Index[node])
end
function Indexer:visit( targetNode, dept )

   targetNode:visit( function ( node, parent, relation, depth )
   
      if _lune._Set_has(self.targetNodeKindSet, node:get_kind() ) then
         self.node2Index[node] = Index.new(self.nsInfo, self.nsInfo:getNewId( node:get_kind() ))
      end
      
      
      if _lune._Set_has(declNameSpaceNodeKindSet, node:get_kind() ) then
         local bakNsInfo = self.nsInfo
         local parentNsInfo = self.nsInfo
         
         do
            local declMethodNode = _lune.__Cast( node, 3, Nodes.DeclMethodNode )
            if declMethodNode ~= nil then
               if declMethodNode:get_declInfo():get_outsizeOfClass() then
                  parentNsInfo = _lune.unwrap( self.nsType2nsInfo[node:get_expType():get_parentInfo()])
               end
               
            end
         end
         
         
         self.nsInfo = NamespaceInfo.new(parentNsInfo, node:get_expType())
         self.nsType2nsInfo[node:get_expType()] = self.nsInfo
         
         self:visit( node, depth + 1 )
         self.nsInfo = bakNsInfo
         
         return Nodes.NodeVisitMode.Next
      end
      
      
      return Nodes.NodeVisitMode.Child
   end, 0, {} )
end
function Indexer:start( targetNode, targetNodeKindSet )

   self.targetNodeKindSet = targetNodeKindSet
   self:visit( targetNode, 0 )
end
function Indexer:dump(  )

   local list = {}
   for node, _1 in pairs( self.node2Index ) do
      table.insert( list, node )
   end
   
   local function comp( node1, node2 )
   
      return node1:get_id() < node2:get_id()
   end
   table.sort( list, comp )
   
   for __index, node in ipairs( list ) do
      local index = _lune.unwrap( self.node2Index[node])
      print( string.format( "%d:%d:", node:get_pos().lineNo, node:get_pos().column), index:getIdTxt(  ), Nodes.getNodeKindName( node:get_kind() ) )
   end
   
end
function Indexer.setmeta( obj )
  setmetatable( obj, { __index = Indexer  } )
end
function Indexer:get_node2Index()
   return self.node2Index
end


return _moduleObj
