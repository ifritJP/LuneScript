--lune/base/dumpNode.lns
local _moduleObj = {}
local __mod__ = 'lune.base.dumpNode'
if not _lune then
   _lune = {}
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

local Ast = _lune.loadModule( 'lune.base.Ast' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Opt = {}
_moduleObj.Opt = Opt
function Opt.new( prefix, depth )
   local obj = {}
   Opt.setmeta( obj )
   if obj.__init then obj:__init( prefix, depth ); end
   return obj
end
function Opt:__init(prefix, depth) 
   self.prefix = prefix
   self.depth = depth
   self.next = nil
end
function Opt:get(  )

   return self.prefix, self.depth
end
function Opt:nextOpt(  )

   do
      local _exp = self.next
      if _exp ~= nil then
         return _exp
      end
   end
   
   local opt = Opt.new(self.prefix .. "  ", self.depth + 1)
   self.next = opt
   return opt
end
function Opt.setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
end

local dumpFilter = {}
setmetatable( dumpFilter, { __index = Ast.Filter } )
function dumpFilter.setmeta( obj )
  setmetatable( obj, { __index = dumpFilter  } )
end
function dumpFilter.new(  )
   local obj = {}
   dumpFilter.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function dumpFilter:__init(  ) 

   Ast.Filter.__init( self )end

local function createFilter(  )

   return dumpFilter.new()
end
_moduleObj.createFilter = createFilter
local localFilter = dumpFilter.new()
local function dump( prefix, depth, node, txt )

   local typeStr = ""
   local expType = node:get_expType(  )
   if expType:equals( Ast.builtinTypeNone ) then
      typeStr = string.format( "(%d:%s:%s)", expType:get_typeId(  ), expType:getTxt(  ), tostring( expType:get_kind(  )))
   end
   
   print( string.format( "%s: %s %s %s", prefix, Ast.getNodeKindName( node:get_kind(  ) ), txt, typeStr) )
end

local function filter( node, filter, opt )

   node:processFilter( localFilter, opt )
end

local function getTxt( token )

   return token.txt
end

function dumpFilter:processNone( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
end


function dumpFilter:processImport( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_modulePath(  ) )
end


function dumpFilter:processRoot( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   for index, child in pairs( node:get_children(  ) ) do
      filter( child, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processSubfile( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
end

function dumpFilter:processBlock( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   for index, statement in pairs( node:get_stmtList(  ) ) do
      filter( statement, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processStmtExp( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processDeclEnum( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_name().txt )
   local enumTypeInfo = node:get_expType()
   for __index, name in pairs( node:get_valueNameList() ) do
      local valInfo = _lune.unwrap( enumTypeInfo:getEnumValInfo( name.txt ))
      print( string.format( "%s  %s: %s", prefix, name.txt, tostring( valInfo:get_val())) )
   end
   
end

function dumpFilter:processDeclAlge( node, opt )

   local prefix, depth = opt:get(  )
   local algeTypeInfo = node:get_algeType()
   dump( prefix, depth, node, algeTypeInfo:get_rawTxt() )
   do
      local __sorted = {}
      local __map = algeTypeInfo:get_valInfoMap()
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local valInfo = __map[ __key ]
         do
            print( string.format( "%s  %s: %s", prefix, algeTypeInfo:get_rawTxt(), valInfo:get_name()) )
         end
      end
   end
   
end

function dumpFilter:processNewAlgeVal( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_name().txt )
   for __index, exp in pairs( node:get_paramList() ) do
      filter( exp, self, opt:nextOpt(  ) )
   end
   
end

function dumpFilter:processDeclClass( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_name(  ).txt )
   for index, field in pairs( node:get_fieldList(  ) ) do
      filter( field, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processDeclMember( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_name(  ).txt )
   filter( node:get_refType(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpMacroExp( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   local stmtList = node:get_stmtList(  )
   for __index, stmt in pairs( stmtList ) do
      filter( stmt, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processDeclMacro( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_expType(  ):getTxt(  ) )
end


function dumpFilter:processExpMacroStat( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_expType(  ):getTxt(  ) )
   for __index, node in pairs( node:get_expStrList(  ) ) do
      filter( node, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processUnwrapSet( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_dstExpList(), self, opt:nextOpt(  ) )
   filter( node:get_srcExpList(), self, opt:nextOpt(  ) )
   if node:get_unwrapBlock() then
      filter( _lune.unwrap( node:get_unwrapBlock()), self, opt:nextOpt(  ) )
   end
   
end

function dumpFilter:processIfUnwrap( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   for index, expNode in pairs( node:get_expNodeList() ) do
      filter( expNode, self, opt:nextOpt(  ) )
   end
   
   filter( node:get_block(), self, opt:nextOpt(  ) )
   if node:get_nilBlock() then
      filter( _lune.unwrap( node:get_nilBlock()), self, opt:nextOpt(  ) )
   end
   
end

function dumpFilter:processWhen( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   for index, expNode in pairs( node:get_expNodeList() ) do
      filter( expNode, self, opt:nextOpt(  ) )
   end
   
   filter( node:get_block(), self, opt:nextOpt(  ) )
   do
      local _exp = node:get_elseBlock()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end

function dumpFilter:processDeclVar( node, opt )

   local prefix, depth = opt:get(  )
   local varName = ""
   for index, var in pairs( node:get_varList(  ) ) do
      if index > 1 then
         varName = varName .. ","
      end
      
      varName = string.format( "%s %s", varName, var:get_name(  ).txt)
   end
   
   if node:get_unwrapBlock() then
      varName = "!" .. varName
   end
   
   varName = string.format( "%s %s", tostring( node:get_mode()), varName)
   dump( prefix, depth, node, varName )
   for index, var in pairs( node:get_varList(  ) ) do
      do
         local _exp = var:get_refType()
         if _exp ~= nil then
            filter( _exp, self, opt:nextOpt(  ) )
         end
      end
      
   end
   
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
   do
      local _exp = node:get_unwrapBlock()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
   do
      local _exp = node:get_thenBlock()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
   for __index, var in pairs( node:get_syncVarList() ) do
      do
         local _exp = var:get_refType()
         if _exp ~= nil then
            filter( _exp, self, opt:nextOpt(  ) )
         end
      end
      
   end
   
   do
      local _exp = node:get_syncBlock()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processDeclArg( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_name(  ).txt )
   filter( node:get_argType(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processDeclArgDDD( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "..." )
end


function dumpFilter:processExpDDD( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "..." )
end


function dumpFilter:processDeclFuncInfo( node, declInfo, opt )

   local prefix, depth = opt:get(  )
   local name = "<anonymous>"
   do
      local _exp = declInfo:get_name(  )
      if _exp ~= nil then
         name = _exp.txt
      end
   end
   
   if node:get_expType():get_mutable() then
      name = name .. " mut"
   end
   
   dump( prefix, depth, node, name )
   local argList = declInfo:get_argList(  )
   for index, arg in pairs( argList ) do
      filter( arg, self, opt:nextOpt(  ) )
   end
   
   do
      local _exp = declInfo:get_body(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end

function dumpFilter:processDeclFunc( node, opt )

   self:processDeclFuncInfo( node, node:get_declInfo(  ), opt )
end


function dumpFilter:processDeclMethod( node, opt )

   self:processDeclFuncInfo( node, node:get_declInfo(  ), opt )
end


function dumpFilter:processDeclConstr( node, opt )

   self:processDeclFuncInfo( node, node:get_declInfo(  ), opt )
end


function dumpFilter:processDeclDestr( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
end


function dumpFilter:processExpCallSuper( node, opt )

   local prefix, depth = opt:get(  )
   local typeInfo = node:get_superType(  )
   dump( prefix, depth, node, typeInfo:getTxt(  ) )
end


function dumpFilter:processRefType( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, (node:get_refFlag(  ) and "&" or "" ) .. (node:get_mutFlag(  ) and "mut " or "" ) )
   filter( node:get_name(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processIf( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   local stmtList = node:get_stmtList(  )
   for index, stmt in pairs( stmtList ) do
      if stmt:get_exp():get_kind() ~= Ast.nodeKind['None'] then
         filter( stmt:get_exp(), self, opt:nextOpt(  ) )
      end
      
      filter( stmt:get_block(), self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processSwitch( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
   local caseList = node:get_caseList()
   for __index, caseInfo in pairs( caseList ) do
      filter( caseInfo:get_expList(), self, opt:nextOpt(  ) )
      filter( caseInfo:get_block(), self, opt:nextOpt(  ) )
   end
   
   do
      local _exp = node:get_default(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processMatch( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_val(), self, opt:nextOpt(  ) )
   local caseList = node:get_caseList()
   for __index, caseInfo in pairs( caseList ) do
      filter( caseInfo:get_block(), self, Opt.new(prefix .. "  " .. caseInfo:get_valInfo():get_name(), depth + 1) )
   end
   
   do
      local _exp = node:get_defaultBlock()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processWhile( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
   filter( node:get_block(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processRepeat( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_block(  ), self, opt:nextOpt(  ) )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processFor( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_val(  ).txt )
   filter( node:get_init(  ), self, opt:nextOpt(  ) )
   filter( node:get_to(  ), self, opt:nextOpt(  ) )
   do
      local _exp = node:get_delta(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
   filter( node:get_block(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processApply( node, opt )

   local prefix, depth = opt:get(  )
   local varNames = ""
   local varList = node:get_varList(  )
   for index, var in pairs( varList ) do
      varNames = varNames .. var.txt .. " "
   end
   
   dump( prefix, depth, node, varNames )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
   filter( node:get_block(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processForeach( node, opt )

   local prefix, depth = opt:get(  )
   local index = ""
   do
      local _exp = node:get_key()
      if _exp ~= nil then
         index = _exp.txt
      end
   end
   
   do
      local _exp = node:get_val()
      if _exp ~= nil then
         dump( prefix, depth, node, _exp.txt .. " " .. index )
      end
   end
   
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
   filter( node:get_block(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processForsort( node, opt )

   local prefix, depth = opt:get(  )
   local index = ""
   do
      local _exp = node:get_key(  )
      if _exp ~= nil then
         index = _exp.txt
      end
   end
   
   dump( prefix, depth, node, node:get_val(  ).txt .. " " .. index )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
   filter( node:get_block(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpUnwrap( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_exp(), self, opt:nextOpt(  ) )
   do
      local _exp = node:get_default()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end

function dumpFilter:processExpCall( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_func(  ), self, opt:nextOpt(  ) )
   do
      local _exp = node:get_argList(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processExpList( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   local expList = node:get_expList(  )
   for index, exp in pairs( expList ) do
      filter( exp, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processExpOp1( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_op(  ).txt )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpCast( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpParen( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "()" )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpOp2( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_op(  ).txt )
   filter( node:get_exp1(  ), self, opt:nextOpt(  ) )
   filter( node:get_exp2(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpNew( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_symbol(  ), self, opt:nextOpt(  ) )
   do
      local _exp = node:get_argList(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processExpRef( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_token(  ).txt )
end


function dumpFilter:processExpRefItem( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "seq[exp] " .. node:get_expType(  ):getTxt(  ) )
   filter( node:get_val(), self, opt:nextOpt(  ) )
   do
      local _exp = node:get_index()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processRefField( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_field(  ).txt )
   filter( node:get_prefix(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpOmitEnum( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%s.%s", node:get_expType():getTxt(  ), node:get_valToken().txt) )
end


function dumpFilter:processGetField( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "get_" .. node:get_field(  ).txt )
   filter( node:get_prefix(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processReturn( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processProvide( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_symbol():get_name() )
end

function dumpFilter:processAlias( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%s = %s", node:get_newName(), node:get_typeInfo():getTxt(  )) )
end

function dumpFilter:processBoxing( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_src(), self, opt:nextOpt(  ) )
end

function dumpFilter:processUnboxing( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_src(), self, opt:nextOpt(  ) )
end

function dumpFilter:processLiteralList( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "[]" )
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processLiteralSet( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "(@)" )
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processLiteralMap( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "{}" )
   local pairList = node:get_pairList(  )
   for __index, pair in pairs( pairList ) do
      filter( pair:get_key(  ), self, opt:nextOpt(  ) )
      filter( pair:get_val(  ), self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processLiteralArray( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "[@]" )
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processLiteralChar( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%s(%s)", tostring( node:get_num(  )), node:get_token(  ).txt ) )
end


function dumpFilter:processLiteralInt( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%s(%s)", tostring( node:get_num(  )), node:get_token(  ).txt ) )
end


function dumpFilter:processLiteralReal( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%s(%s)", tostring( node:get_num(  )), node:get_token(  ).txt ) )
end


function dumpFilter:processLiteralString( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_token(  ).txt )
   for __index, param in pairs( node:get_argList() ) do
      filter( param, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processLiteralBool( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_token(  ).txt == "true" and "true" or "false" )
end


function dumpFilter:processLiteralNil( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
end


function dumpFilter:processBreak( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
end


function dumpFilter:processLiteralSymbol( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_token(  ).txt )
end


return _moduleObj
