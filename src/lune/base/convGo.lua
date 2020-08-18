--lune/base/convGo.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@convGo'
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
local Ver = _lune.loadModule( 'lune.base.Ver' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Util = _lune.loadModule( 'lune.base.Util' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local LuaMod = _lune.loadModule( 'lune.base.LuaMod' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )

local Opt = {}
_moduleObj.Opt = Opt
function Opt.setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
end
function Opt.new( node )
   local obj = {}
   Opt.setmeta( obj )
   if obj.__init then
      obj:__init( node )
   end
   return obj
end
function Opt:__init( node )

   self.node = node
end


local convFilter = {}
setmetatable( convFilter, { __index = Nodes.Filter } )
function convFilter.new( enableTest, streamName, stream, ast )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( enableTest, streamName, stream, ast ); end
   return obj
end
function convFilter:__init(enableTest, streamName, stream, ast) 
   Nodes.Filter.__init( self,true, ast:get_moduleTypeInfo(), ast:get_moduleTypeInfo():get_scope())
   
   
   self.stream = Util.SimpleSourceOStream.new(stream, nil, 4)
end
function convFilter.setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end
function convFilter:popIndent( ... )
   return self.stream:popIndent( ... )
end

function convFilter:pushIndent( ... )
   return self.stream:pushIndent( ... )
end

function convFilter:returnToSource( ... )
   return self.stream:returnToSource( ... )
end

function convFilter:switchToHeader( ... )
   return self.stream:switchToHeader( ... )
end

function convFilter:write( ... )
   return self.stream:write( ... )
end

function convFilter:writeln( ... )
   return self.stream:writeln( ... )
end



local function filter( node, filter, parent )

   node:processFilter( filter, Opt.new(parent) )
end

local function isAnyType( typeInfo )

   local work = typeInfo:get_srcTypeInfo()
   return typeInfo:get_nilable() or work:get_kind() == Ast.TypeInfoKind.IF or work == Ast.builtinTypeStem
end

local golanKeywordSet = {["func"] = true, ["select"] = true, ["defer"] = true, ["go"] = true, ["chan"] = true, ["package"] = true, ["const"] = true, ["fallthrough"] = true, ["range"] = true, ["continue"] = true, ["var"] = true}

function convFilter:outputSymbol( name )

   if _lune._Set_has(golanKeywordSet, name ) then
      self:write( string.format( "%s_", name) )
   else
    
      self:write( name )
   end
   
end

local function str2gostr( txt )

   local work = txt
   if string.find( work, '^```' ) then
      work = (string.format( "%q", work:sub( 4, -4 )) ):gsub( "\\\n", "\\n" )
   elseif string.find( work, "^'" ) then
      work = string.format( '"%s"', ((string.format( "%s", work:sub( 2, -2 )) ):gsub( '"', '\\"' ) ))
   end
   
   work = work:gsub( "\\9", "\\t" )
   return work
end

local type2gotypeMap = {[Ast.builtinTypeInt] = "LnsInt", [Ast.builtinTypeReal] = "LnsReal", [Ast.builtinTypeStem] = "LnsAny", [Ast.builtinTypeString] = "string"}

local function getOrgTypeInfo( typeInfo )

   do
      local enumType = _lune.__Cast( typeInfo:get_srcTypeInfo():get_nonnilableType(), 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         return enumType:get_valTypeInfo()
      end
   end
   
   return typeInfo:get_srcTypeInfo():get_nonnilableType()
end

local function type2gotype( typeInfo )

   if isAnyType( typeInfo ) then
      return "LnsAny"
   end
   
   local orgType = getOrgTypeInfo( typeInfo )
   do
      local goType = type2gotypeMap[orgType]
      if goType ~= nil then
         return goType
      end
   end
   
   if typeInfo:get_kind() == Ast.TypeInfoKind.DDD then
      return "[]LnsAny"
   end
   
   Util.err( string.format( "not support yet -- %s", typeInfo:getTxt(  )) )
end




function convFilter:processBlankLine( node, opt )

end



function convFilter:processNone( node, opt )

end



function convFilter:processImport( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processImport'

   Util.err( string.format( "not support -- %s", __func__) )
end

local ExpListKind = {}
ExpListKind._val2NameMap = {}
function ExpListKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ExpListKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ExpListKind._from( val )
   if ExpListKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ExpListKind.__allList = {}
function ExpListKind.get__allList()
   return ExpListKind.__allList
end

ExpListKind.Direct = 0
ExpListKind._val2NameMap[0] = 'Direct'
ExpListKind.__allList[1] = ExpListKind.Direct
ExpListKind.Slice = 1
ExpListKind._val2NameMap[1] = 'Slice'
ExpListKind.__allList[2] = ExpListKind.Slice
ExpListKind.Conv = 2
ExpListKind._val2NameMap[2] = 'Conv'
ExpListKind.__allList[3] = ExpListKind.Conv


local function getExpListKind( dstTypeList, node )

   local lastExp = node:get_expList()[#node:get_expList()]
   local hasAbbr
   
   if lastExp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
      hasAbbr = true
      if #node:get_expList() < 2 then
         return ExpListKind.Direct
      end
      
      lastExp = node:get_expList()[#node:get_expList() - 1]
   else
    
      hasAbbr = false
   end
   
   do
      local exp2ddd = _lune.__Cast( lastExp, 3, Nodes.ExpToDDDNode )
      if exp2ddd ~= nil then
         local mRetExp = node:get_mRetExp()
         if  nil == mRetExp then
            local _mRetExp = mRetExp
         
            return ExpListKind.Slice
         end
         
         if mRetExp:get_index() == 1 and dstTypeList[mRetExp:get_index()]:get_kind() == Ast.TypeInfoKind.DDD then
            return ExpListKind.Slice
         end
         
         return ExpListKind.Conv
      end
   end
   
   if lastExp:get_expType():get_kind() == Ast.TypeInfoKind.DDD then
      local mRetExp = node:get_mRetExp()
      if  nil == mRetExp then
         local _mRetExp = mRetExp
      
         return ExpListKind.Slice
      end
      
      if mRetExp:get_index() == 1 and dstTypeList[mRetExp:get_index()]:get_kind() == Ast.TypeInfoKind.DDD then
         return ExpListKind.Direct
      end
      
   else
    
      local mRetExp = node:get_mRetExp()
      if  nil == mRetExp then
         local _mRetExp = mRetExp
      
         return ExpListKind.Direct
      end
      
      if not hasAbbr and mRetExp:get_index() == 1 then
         return ExpListKind.Direct
      end
      
   end
   
   return ExpListKind.Conv
end

function convFilter:processConvArg( nodeId, dstTypeList, argListNode )

   local argList = argListNode
   if  nil == argList then
      local _argList = argList
   
      return 
   end
   
   
   if getExpListKind( dstTypeList, argList ) ~= ExpListKind.Conv then
      return 
   end
   
   
   
   self:write( string.format( "func _convArg%d_%d(", nodeId, argList:get_id()) )
   local mRetIndex = _lune.nilacc( argList:get_mRetExp(), 'get_index', 'callmtd' )
   if  nil == mRetIndex then
      local _mRetIndex = mRetIndex
   
      return 
   end
   
   for index, argExp in ipairs( argList:get_expList() ) do
      do
         local exp2ddd = _lune.__Cast( argExp, 3, Nodes.ExpToDDDNode )
         if exp2ddd ~= nil then
            for __index, exp in ipairs( exp2ddd:get_expList():get_expList() ) do
               if index ~= 1 then
                  self:write( ", " )
               end
               
               self:write( string.format( "arg%d ", index) )
               self:write( type2gotype( exp:get_expType() ) )
            end
            
         else
            if index ~= 1 then
               self:write( ", " )
            end
            
            if mRetIndex == index then
               self:write( string.format( "arg%d []LnsAny", index) )
               break
            else
             
               self:write( string.format( "arg%d ", index) )
               self:write( type2gotype( argExp:get_expType() ) )
            end
            
         end
      end
      
   end
   
   self:write( ") (" )
   for index, argType in ipairs( dstTypeList ) do
      if index ~= 1 then
         self:write( ", " )
      end
      
      self:write( type2gotype( argType ) )
   end
   
   self:writeln( ") {" )
   self:write( "return " )
   
   local restIndex = nil
   for index, dstType in ipairs( dstTypeList ) do
      if index ~= 1 then
         self:write( ", " )
      end
      
      if dstType:get_kind() == Ast.TypeInfoKind.DDD then
         restIndex = index
         break
      end
      
      if index >= mRetIndex then
         self:write( string.format( "arg%d[%d].(%s)", mRetIndex, index - mRetIndex, type2gotype( dstType )) )
      else
       
         self:write( string.format( "arg%d", index) )
      end
      
   end
   
   if restIndex ~= nil then
      self:write( "[]LnsAny{ " )
      for index, argExp in ipairs( argList:get_expList() ) do
         if index >= #dstTypeList then
            self:write( string.format( "arg%d", index) )
         end
         
      end
      
      self:writeln( "}" )
   else
      self:writeln( "" )
   end
   
   
   self:writeln( "}" )
end


function convFilter:processRoot( node, opt )

   self:writeln( "package main" )
   
   for __index, workNode in ipairs( node:get_nodeManager():getExpCallNodeList(  ) ) do
      self:processConvArg( workNode:get_id(), workNode:get_func():get_expType():get_argTypeInfoList(), workNode:get_argList() )
   end
   
   for __index, workNode in ipairs( node:get_nodeManager():getDeclVarNodeList(  ) ) do
      self:processConvArg( workNode:get_id(), workNode:get_typeInfoList(), workNode:get_expList() )
   end
   
   
   for __index, declFuncNode in ipairs( node:get_nodeManager():getDeclFuncNodeList(  ) ) do
      filter( declFuncNode, self, node )
      self:writeln( "" )
   end
   
   
   self:writeln( "func Lns_init() {" )
   self:pushIndent(  )
   
   for __index, child in ipairs( node:get_children() ) do
      
      do
         local _switchExp = child:get_kind()
         if _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_DeclFunc() or _switchExp == Nodes.NodeKind.get_DeclMacro() or _switchExp == Nodes.NodeKind.get_TestBlock() then
            
         else 
            
               filter( child, self, node )
               self:writeln( "" )
         end
      end
      
   end
   
   
   self:popIndent(  )
   self:writeln( "}" )
end


function convFilter:processSubfile( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processSubfile'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processBlockSub( node, opt )

   if node:get_blockKind() == Nodes.BlockKind.Block then
      self:writeln( "{" )
   end
   
   
   self:pushIndent(  )
   for __index, child in ipairs( node:get_stmtList() ) do
      
      do
         local _switchExp = child:get_kind()
         if _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_DeclFunc() or _switchExp == Nodes.NodeKind.get_DeclMacro() or _switchExp == Nodes.NodeKind.get_TestBlock() then
            
         else 
            
               filter( child, self, node )
               self:writeln( "" )
         end
      end
      
   end
   
   self:popIndent(  )
   
   if node:get_blockKind() == Nodes.BlockKind.Block then
      self:write( "}" )
   end
   
end


function convFilter:processStmtExp( node, opt )

   filter( node:get_exp(), self, node )
end


function convFilter:processDeclEnum( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclEnum'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclAlge( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclAlge'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processNewAlgeVal( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processNewAlgeVal'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclMember( node, opt )

   
end


function convFilter:processExpMacroExp( node, opt )

   for __index, stmt in ipairs( node:get_stmtList() ) do
      filter( stmt, self, node )
      self:writeln( "" )
   end
   
end


function convFilter:processDeclMacro( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclMacro'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpMacroStat( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpMacroStat'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclConstr( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclConstr'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclDestr( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclDestr'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpCallSuper( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpCallSuper'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclMethod( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclMethod'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processProtoMethod( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processProtoMethod'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processUnwrapSet( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processUnwrapSet'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processIfUnwrap( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processIfUnwrap'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processSetFromExpList( convArgFuncName, dstTypeList, expListNode )

   do
      local _switchExp = getExpListKind( dstTypeList, expListNode )
      if _switchExp == ExpListKind.Conv then
         self:write( string.format( "%s(", convArgFuncName) )
         local mRetIndex = _lune.nilacc( expListNode:get_mRetExp(), 'get_index', 'callmtd' )
         
         for index, exp in ipairs( expListNode:get_expList() ) do
            if index ~= 1 then
               self:write( ', ' )
            end
            
            if mRetIndex == index then
               self:write( "Lns_2DDD(" )
               filter( exp, self, expListNode )
               self:write( ")" )
               break
            end
            
            filter( exp, self, expListNode )
            if _lune.nilacc( expListNode:get_mRetExp(), 'get_index', 'callmtd' ) == index then
               break
            end
            
         end
         
         self:write( ")" )
      elseif _switchExp == ExpListKind.Slice then
         for index, argType in ipairs( dstTypeList ) do
            if index ~= 1 then
               self:write( ', ' )
            end
            
            if #expListNode:get_expList() >= index then
               local argExp = expListNode:get_expList()[index]
               
               do
                  local exp2ddd = _lune.__Cast( argExp, 3, Nodes.ExpToDDDNode )
                  if exp2ddd ~= nil then
                     local subList = exp2ddd:get_expList()
                     local mRetIndex = _lune.nilacc( subList:get_mRetExp(), 'get_index', 'callmtd' )
                     if mRetIndex and mRetIndex == 1 then
                        local subExp = subList:get_expList()[1]
                        if subExp:get_expType():get_kind() ~= Ast.TypeInfoKind.DDD then
                           self:write( "Lns_2DDD(" )
                           filter( subExp, self, expListNode )
                           self:write( ")" )
                        else
                         
                           filter( subExp, self, expListNode )
                        end
                        
                     else
                      
                        if mRetIndex and mRetIndex ~= 1 then
                           self:write( "append( " )
                        end
                        
                        self:write( "[]LnsAny{" )
                        for subIndex, subExp in ipairs( subList:get_expList() ) do
                           if mRetIndex == subIndex then
                              if mRetIndex ~= 1 then
                                 self:write( "}, " )
                              end
                              
                              if subExp:get_expType():get_kind() ~= Ast.TypeInfoKind.DDD then
                                 self:write( "Lns_2DDD(" )
                                 filter( subExp, self, expListNode )
                                 self:write( ")" )
                              else
                               
                                 filter( subExp, self, expListNode )
                              end
                              
                              self:write( "..." )
                              break
                           end
                           
                           if subIndex ~= 1 then
                              self:write( ', ' )
                           end
                           
                           filter( subExp, self, expListNode )
                        end
                        
                        if mRetIndex and mRetIndex ~= 1 then
                           self:write( ")" )
                        else
                         
                           self:write( "}" )
                        end
                        
                     end
                     
                  else
                     if argExp:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
                        if argType:get_kind() == Ast.TypeInfoKind.DDD then
                           self:write( "[]LnsAny{}" )
                        else
                         
                           self:write( "nil" )
                        end
                        
                     else
                      
                        filter( argExp, self, expListNode )
                     end
                     
                  end
               end
               
            else
             
               self:write( "[]LnsAny{}" )
            end
            
         end
         
      elseif _switchExp == ExpListKind.Direct then
         local mRetIndex = _lune.nilacc( expListNode:get_mRetExp(), 'get_index', 'callmtd' )
         for index, funcArgType in ipairs( dstTypeList ) do
            if mRetIndex == index - 1 then
               break
            end
            
            if index ~= 1 then
               self:write( ', ' )
            end
            
            if index == #dstTypeList and funcArgType:get_kind() == Ast.TypeInfoKind.DDD then
               if #expListNode:get_expList() < index or expListNode:get_expList()[index]:get_expType():get_kind() == Ast.TypeInfoKind.Abbr then
                  self:write( "[]LnsAny{}" )
               else
                
                  filter( expListNode:get_expList()[index], self, expListNode )
               end
               
            else
             
               if #expListNode:get_expList() < index then
                  self:write( "nil" )
               else
                
                  filter( expListNode:get_expList()[index], self, expListNode )
               end
               
            end
            
         end
         
      end
   end
   
end


function convFilter:processDeclVar( node, opt )

   self:write( "var" )
   
   local prevType = ""
   
   do
      local expList = node:get_expList()
      if expList ~= nil then
         for index, exp in ipairs( expList:get_expList() ) do
            local goType
            
            local varName
            
            if #node:get_varList() >= index then
               local varInfo = node:get_varList()[index]
               goType = type2gotype( varInfo:get_actualType() )
               varName = varInfo:get_name().txt
            else
             
               goType = type2gotype( exp:get_expType() )
               varName = "_"
            end
            
            if goType ~= prevType then
               self:write( string.format( " %s", prevType) )
            end
            
            if index ~= 1 then
               self:write( ", " )
            end
            
            self:write( varName )
            prevType = goType
         end
         
         self:write( string.format( " %s", prevType) )
         
         self:write( " = " )
         self:processSetFromExpList( string.format( "_convArg%d_%d", node:get_id(), expList:get_id()), node:get_typeInfoList(), expList )
      else
         for index, varInfo in ipairs( node:get_varList() ) do
            local goType = type2gotype( varInfo:get_actualType() )
            if goType ~= prevType then
               self:write( string.format( " %s", prevType) )
            end
            
            if index ~= 1 then
               self:write( ", " )
            end
            
            self:write( string.format( "%s", varInfo:get_name().txt) )
            prevType = goType
         end
         
         self:write( string.format( " %s", prevType) )
      end
   end
   
end


function convFilter:processWhen( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processWhen'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclArg( node, opt )

   self:write( string.format( "%s ", node:get_name().txt) )
   filter( _lune.unwrap( node:get_argType()), self, node )
end


function convFilter:processDeclArgDDD( node, opt )

   self:write( "ddd []LnsAny" )
end


function convFilter:processExpDDD( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpDDD'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpSubDDD( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpSubDDD'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclForm( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclForm'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclFunc( node, opt )

   self:write( "func " )
   do
      local name = node:get_declInfo():get_name()
      if name ~= nil then
         self:outputSymbol( name.txt )
      end
   end
   
   self:write( "(" )
   
   for index, arg in ipairs( node:get_declInfo():get_argList() ) do
      if index ~= 1 then
         self:write( "," )
      end
      
      filter( arg, self, node )
   end
   
   self:write( ") " )
   
   local retTypeList = node:get_declInfo():get_retTypeInfoList()
   do
      local _switchExp = #retTypeList
      if _switchExp == 0 then
         self:write( "" )
      elseif _switchExp == 1 then
         self:write( type2gotype( retTypeList[1] ) )
      else 
         
            self:write( "(" )
            for index, retType in ipairs( retTypeList ) do
               if index ~= 1 then
                  self:write( ", " )
               end
               
               self:write( type2gotype( retType ) )
            end
            
            self:write( ")" )
      end
   end
   
   self:writeln( " {" )
   
   do
      local body = node:get_declInfo():get_body()
      if body ~= nil then
         filter( body, self, node )
      end
   end
   
   self:write( "}" )
end


function convFilter:processRefType( node, opt )

   if isAnyType( node:get_expType() ) then
      self:write( "LnsAny" )
   else
    
      do
         local goType = type2gotypeMap[getOrgTypeInfo( node:get_expType() )]
         if goType ~= nil then
            self:write( goType )
         else
            filter( node:get_name(), self, node )
         end
      end
      
   end
   
end


function convFilter:processIf( node, opt )

   for __index, stmt in ipairs( node:get_stmtList() ) do
      do
         local _switchExp = stmt:get_kind()
         if _switchExp == Nodes.IfKind.If then
            self:write( "if " )
            filter( stmt:get_exp(), self, node )
            self:writeln( "{" )
            filter( stmt:get_block(), self, node )
         elseif _switchExp == Nodes.IfKind.ElseIf then
            self:write( "} else if " )
            filter( stmt:get_exp(), self, node )
            self:writeln( "{" )
            filter( stmt:get_block(), self, node )
         elseif _switchExp == Nodes.IfKind.Else then
            self:writeln( "} else { " )
            filter( stmt:get_block(), self, node )
         end
      end
      
   end
   
   self:write( "}" )
end


function convFilter:processSwitch( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processSwitch'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processMatch( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processMatch'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processWhile( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processWhile'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processRepeat( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processRepeat'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processFor( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processFor'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processApply( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processApply'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processForeach( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processForeach'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processForsort( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processForsort'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpUnwrap( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpUnwrap'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpToDDD( node, opt )

   do
      local mRetExp = node:get_expList():get_mRetExp()
      if mRetExp ~= nil then
         filter( node:get_expList(), self, node )
      else
         self:write( "[]LnsAny{ " )
         filter( node:get_expList(), self, node )
         self:write( "}" )
      end
   end
   
end


function convFilter:processExpNew( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpNew'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processDeclClass( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processDeclClass'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpCall( node, opt )

   local funcType = node:get_func():get_expType()
   if Ast.isBuiltin( funcType:get_typeId() ) then
      if funcType:get_rawTxt() == "print" then
         self:write( "Lns_print" )
      end
      
   else
    
      self:outputSymbol( funcType:get_rawTxt() )
   end
   
   self:write( "(" )
   
   do
      local argList = node:get_argList()
      if argList ~= nil then
         self:processSetFromExpList( string.format( "_convArg%d_%d( ", node:get_id(), argList:get_id()), funcType:get_argTypeInfoList(), argList )
      end
   end
   
   
   self:write( ")" )
end


function convFilter:processExpAccessMRet( node, opt )

   
end


function convFilter:processExpList( node, opt )

   for index, exp in ipairs( node:get_expList() ) do
      if index ~= 1 then
         self:write( ", " )
      end
      
      do
         local mRetExp = node:get_mRetExp()
         if mRetExp ~= nil then
            if mRetExp:get_index() == index then
               self:write( "Lns_2DDD(" )
               filter( exp, self, node )
               self:write( ")" )
               break
            end
            
         end
      end
      
      filter( exp, self, node )
   end
   
end


function convFilter:processExpOp1( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpOp1'

   do
      local _switchExp = node:get_op().txt
      if _switchExp == "~" then
         self:write( "^" )
         filter( node:get_exp(), self, node )
      elseif _switchExp == "+" or _switchExp == "-" then
         self:write( node:get_op().txt )
         filter( node:get_exp(), self, node )
      elseif _switchExp == "not" then
         self:write( "Lns_op_not(" )
         filter( node:get_exp(), self, node )
         self:write( ")" )
      else 
         
            Util.err( string.format( "%s: not support -- %s", __func__, node:get_op().txt) )
      end
   end
   
end


function convFilter:processExpMultiTo1( node, opt )

   self:write( "Lns_car(" )
   filter( node:get_exp(), self, node )
   self:write( ")" )
   if not isAnyType( node:get_expType() ) then
      self:write( string.format( ".(%s)", type2gotype( node:get_expType() )) )
   end
   
end


function convFilter:processExpCast( node, opt )

   do
      local _switchExp = node:get_castKind()
      if _switchExp == Nodes.CastKind.Force then
         filter( node:get_exp(), self, node )
      elseif _switchExp == Nodes.CastKind.Implicit then
         filter( node:get_exp(), self, node )
      elseif _switchExp == Nodes.CastKind.Normal then
         filter( node:get_exp(), self, node )
      end
   end
   
end


function convFilter:processExpParen( node, opt )

   if #node:get_exp():get_expTypeList() >= 2 then
      self:write( "Lns_car(" )
      filter( node:get_exp(), self, node )
      self:write( ")" )
      if not isAnyType( node:get_expType() ) then
         self:write( string.format( ".(%s)", type2gotype( node:get_expType() )) )
      end
      
   else
    
      self:write( "(" )
      filter( node:get_exp(), self, node )
      self:write( ")" )
   end
   
end


function convFilter:processExpSetVal( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpSetVal'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processAndOr( node, opTxt, parent )

   local function isAndOr( exp )
   
      do
         local parentNode = _lune.__Cast( exp, 3, Nodes.ExpOp2Node )
         if parentNode ~= nil then
            do
               local _switchExp = parentNode:get_op().txt
               if _switchExp == "and" or _switchExp == "or" then
                  return true
               end
            end
            
         end
      end
      
      return false
   end
   
   local firstFlag = not isAndOr( parent )
   if firstFlag then
      self:writeln( "Lns_popVal( Lns_incStack() ||" )
      self:pushIndent(  )
   end
   
   
   local opCC
   
   if opTxt == "and" then
      opCC = "&&"
   else
    
      opCC = "||"
   end
   
   
   if isAndOr( node:get_exp1() ) then
      filter( node:get_exp1(), self, node )
   else
    
      self:write( "Lns_setStackVal( " )
      filter( node:get_exp1(), self, node )
      self:write( ") " )
   end
   
   self:writeln( opCC )
   if isAndOr( node:get_exp2() ) then
      filter( node:get_exp2(), self, node )
   else
    
      self:write( "Lns_setStackVal( " )
      filter( node:get_exp2(), self, node )
      self:write( ") " )
   end
   
   
   if firstFlag then
      self:write( ")" )
      
      if not isAnyType( node:get_expType() ) then
         self:write( string.format( ".(%s)", type2gotype( node:get_expType() )) )
      end
      
      self:popIndent(  )
   end
   
end


function convFilter:processExpOp2( node, opt )

   local opTxt = node:get_op().txt
   
   do
      local _switchExp = opTxt
      if _switchExp == "and" or _switchExp == "or" then
         self:processAndOr( node, opTxt, opt.node )
      elseif _switchExp == ".." then
         filter( node:get_exp1(), self, node )
         self:write( " + " )
         filter( node:get_exp2(), self, node )
      else 
         
            do
               local _exp = Ast.bitBinOpMap[opTxt]
               if _exp ~= nil then
                  do
                     local _switchExp = _exp
                     if _switchExp == Ast.BitOpKind.LShift then
                        opTxt = "<<"
                     elseif _switchExp == Ast.BitOpKind.RShift then
                        opTxt = ">>"
                     end
                  end
                  
                  filter( node:get_exp1(), self, node )
                  self:write( " " .. opTxt .. " " )
                  filter( node:get_exp2(), self, node )
               else
                  filter( node:get_exp1(), self, node )
                  self:write( " " .. opTxt .. " " )
                  filter( node:get_exp2(), self, node )
               end
            end
            
      end
   end
   
end


function convFilter:processExpRef( node, opt )

   if node:get_expType():get_kind() == Ast.TypeInfoKind.DDD then
      self:write( "ddd" )
   else
    
      self:write( node:get_symbolInfo():get_name() )
   end
   
end


function convFilter:processExpRefItem( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpRefItem'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processRefField( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processRefField'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processExpOmitEnum( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processExpOmitEnum'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processGetField( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processGetField'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processReturn( node, opt )

   self:write( "return " )
   do
      local expList = node:get_expList()
      if expList ~= nil then
         filter( expList, self, node )
      end
   end
   
end


function convFilter:processTestBlock( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processTestBlock'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processProvide( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processProvide'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processAlias( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processAlias'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processBoxing( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processBoxing'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processUnboxing( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processUnboxing'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processLiteralList( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processLiteralList'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processLiteralSet( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processLiteralSet'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processLiteralMap( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processLiteralMap'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processLiteralArray( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processLiteralArray'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processLiteralChar( node, opt )

   self:write( string.format( "%d", node:get_num() ) )
end


function convFilter:processLiteralInt( node, opt )

   self:write( node:get_token().txt )
end


function convFilter:processLiteralReal( node, opt )

   self:write( node:get_token().txt )
end


function convFilter:processLiteralString( node, opt )

   local txt = node:get_token().txt
   self:write( string.format( '%s', str2gostr( txt )) )
end


function convFilter:processLiteralBool( node, opt )

   self:write( node:get_token().txt )
end


function convFilter:processLiteralNil( node, opt )

   self:write( "nil" )
end


function convFilter:processBreak( node, opt )

   self:write( "break" )
end


function convFilter:processLiteralSymbol( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processLiteralSymbol'

   Util.err( string.format( "not support -- %s", __func__) )
end


function convFilter:processAbbr( node, opt )
   local __func__ = '@lune.@base.@convGo.convFilter.processAbbr'

   Util.err( string.format( "not support -- %s", __func__) )
end


local function createFilter( enableTest, streamName, stream, ast )

   return convFilter.new(enableTest, streamName, stream, ast)
end
_moduleObj.createFilter = createFilter

return _moduleObj
