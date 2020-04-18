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
   Nodes.Filter.__init( self,ast:get_moduleTypeInfo(), ast:get_moduleTypeInfo():get_scope())
   
   
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

function convFilter:processNone( node, opt )

end



function convFilter:processImport( node, opt )

end



function convFilter:processRoot( node, opt )

   self:writeln( "package main" )
   
   self:writeln( "func Lns_init() {" )
   self:writeln( '    Lns_print( 1, 2.0, 2.5, true, "abc" )' )
   self:writeln( "}" )
end



function convFilter:processSubfile( node, opt )

end


function convFilter:processBlockSub( node, opt )

end



function convFilter:processStmtExp( node, opt )

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

   for __index, stmt in pairs( node:get_stmtList() ) do
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

end


function convFilter:processExpNew( node, opt )

end



function convFilter:processDeclClass( node, opt )

end



function convFilter:processExpCall( node, opt )

end



function convFilter:processExpAccessMRet( node, opt )

end



function convFilter:processExpList( node, opt )

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
