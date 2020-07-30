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

local type2gotypeMap = {[Ast.builtinTypeInt] = "LnsInt", [Ast.builtinTypeReal] = "LnsReal", [Ast.builtinTypeStem] = "LnsSfem", [Ast.builtinTypeString] = "string"}

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

   local orgType = getOrgTypeInfo( typeInfo )
   do
      local goType = type2gotypeMap[orgType]
      if goType ~= nil then
         return goType
      end
   end
   
   Util.err( string.format( "not support yet -- %s", typeInfo:getTxt(  )) )
end

function convFilter:processNone( node, opt )

end



function convFilter:processImport( node, opt )

end



function convFilter:processRoot( node, opt )

   self:writeln( "package main" )
   
   for __index, declFuncNode in ipairs( node:get_nodeManager():getDeclFuncNodeList(  ) ) do
      filter( declFuncNode, self, node )
      self:writeln( "" )
   end
   
   
   self:writeln( "func Lns_init() {" )
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
   
   self:writeln( "}" )
end



function convFilter:processSubfile( node, opt )

end


function convFilter:processBlockSub( node, opt )

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
end



function convFilter:processStmtExp( node, opt )

   filter( node:get_exp(), self, node )
end



function convFilter:processDeclEnum( node, opt )

end


function convFilter:processDeclAlge( node, opt )

end


function convFilter:processNewAlgeVal( node, opt )

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

end



function convFilter:processExpMacroStat( node, opt )

end



function convFilter:processDeclConstr( node, opt )

end



function convFilter:processDeclDestr( node, opt )

end


function convFilter:processExpCallSuper( node, opt )

end



function convFilter:processDeclMethod( node, opt )

end



function convFilter:processProtoMethod( node, opt )

end



function convFilter:processUnwrapSet( node, opt )

end


function convFilter:processIfUnwrap( node, opt )

end


function convFilter:processDeclVar( node, opt )

end



function convFilter:processWhen( node, opt )

end


function convFilter:processDeclArg( node, opt )

end



function convFilter:processDeclArgDDD( node, opt )

end



function convFilter:processExpDDD( node, opt )

end



function convFilter:processExpSubDDD( node, opt )

end



function convFilter:processDeclForm( node, opt )

end


function convFilter:processDeclFunc( node, opt )

   self:write( "func " )
   do
      local name = node:get_declInfo():get_name()
      if name ~= nil then
         self:write( name.txt )
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

end



function convFilter:processIf( node, opt )

end



function convFilter:processSwitch( node, opt )

end



function convFilter:processMatch( node, opt )

end



function convFilter:processWhile( node, opt )

end



function convFilter:processRepeat( node, opt )

end



function convFilter:processFor( node, opt )

end



function convFilter:processApply( node, opt )

end



function convFilter:processForeach( node, opt )

end



function convFilter:processForsort( node, opt )

end



function convFilter:processExpUnwrap( node, opt )

end


function convFilter:processExpToDDD( node, opt )

   filter( node:get_expList(), self, node )
end


function convFilter:processExpNew( node, opt )

end



function convFilter:processDeclClass( node, opt )

end



function convFilter:processExpCall( node, opt )

   local funcType = node:get_func():get_expType()
   if Ast.isBuiltin( funcType:get_typeId() ) then
      if funcType:get_rawTxt() == "print" then
         self:write( "Lns_print" )
      end
      
   else
    
      self:write( funcType:get_rawTxt() )
   end
   
   self:write( "(" )
   
   do
      local argList = node:get_argList()
      if argList ~= nil then
         for index, exp in ipairs( argList:get_expList() ) do
            if index ~= 1 then
               self:write( ', ' )
            end
            
            filter( exp, self, node )
         end
         
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
      
      filter( exp, self, node )
      do
         local mRetExp = node:get_mRetExp()
         if mRetExp ~= nil then
            if mRetExp:get_index() == index then
               break
            end
            
         end
      end
      
   end
   
end



function convFilter:processExpOp1( node, opt )

end



function convFilter:processExpMultiTo1( node, opt )

end


function convFilter:processExpCast( node, opt )

end



function convFilter:processExpParen( node, opt )

end



function convFilter:processExpSetVal( node, opt )

end



function convFilter:processExpOp2( node, opt )

   local opTxt = node:get_op().txt
   
   do
      local _switchExp = opTxt
      if _switchExp == "and" or _switchExp == "or" then
         Util.err( "not support yet" )
      elseif _switchExp == ".." then
         Util.err( "not support yet" )
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

end



function convFilter:processExpRefItem( node, opt )

end



function convFilter:processRefField( node, opt )

end



function convFilter:processExpOmitEnum( node, opt )

end



function convFilter:processGetField( node, opt )

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

end


function convFilter:processProvide( node, opt )

end


function convFilter:processAlias( node, opt )

end


function convFilter:processBoxing( node, opt )

end


function convFilter:processUnboxing( node, opt )

end


function convFilter:processLiteralList( node, opt )

end



function convFilter:processLiteralSet( node, opt )

end



function convFilter:processLiteralMap( node, opt )

end



function convFilter:processLiteralArray( node, opt )

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

end



function convFilter:processLiteralNil( node, opt )

end



function convFilter:processBreak( node, opt )

end



function convFilter:processLiteralSymbol( node, opt )

end



function convFilter:processAbbr( node, opt )

   
   Util.err( "illegal" )
end



local function createFilter( enableTest, streamName, stream, ast )

   return convFilter.new(enableTest, streamName, stream, ast)
end
_moduleObj.createFilter = createFilter

return _moduleObj
