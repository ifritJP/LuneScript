--lune/base/dumpNode.lns
local _moduleObj = {}
local __mod__ = 'lune.base.dumpNode'
if not _lune then
   _lune = {}
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
local dumpFilter = {}
setmetatable( dumpFilter, { __index = Ast.Filter } )
_moduleObj.dumpFilter = dumpFilter
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

   Ast.Filter.__init( self )
end

local localFilter = dumpFilter.new()
local function dump( prefix, depth, node, txt )

   local typeStr = ""
   local expType = node:get_expType(  )
   if expType:equals( Ast.builtinTypeNone ) then
      typeStr = string.format( "(%d:%s:%s)", expType:get_typeId(  ), expType:getTxt(  ), tostring( expType:get_kind(  )))
   end
   
   print( string.format( "%s: %s %s %s", prefix, Ast.getNodeKindName( node:get_kind(  ) ), txt, typeStr) )
end

local function filter( node, filter, prefix, depth )

   node:processFilter( localFilter, prefix, depth )
end

local function getTxt( token )

   return token.txt
end

function dumpFilter:processNone( node, prefix, depth )

   dump( prefix, depth, node, "" )
end


function dumpFilter:processImport( node, prefix, depth )

   dump( prefix, depth, node, node:get_modulePath(  ) )
end


function dumpFilter:processRoot( node, prefix, depth )

   dump( prefix, depth, node, "" )
   for index, child in pairs( node:get_children(  ) ) do
      filter( child, self, prefix .. "  ", depth + 1 )
   end
   
end


function dumpFilter:processSubfile( node, prefix, depth )

   dump( prefix, depth, node, "" )
end

function dumpFilter:processBlock( node, prefix, depth )

   dump( prefix, depth, node, "" )
   for index, statement in pairs( node:get_stmtList(  ) ) do
      filter( statement, self, prefix .. "  ", depth + 1 )
   end
   
end


function dumpFilter:processStmtExp( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processDeclEnum( node, prefix, depth )

   dump( prefix, depth, node, node:get_name().txt )
   local enumTypeInfo = node:get_expType()
   for __index, name in pairs( node:get_valueNameList() ) do
      local valInfo = _lune.unwrap( enumTypeInfo:getEnumValInfo( name.txt ))
      print( string.format( "%s  %s: %s", prefix, name.txt, tostring( valInfo:get_val())) )
   end
   
end

function dumpFilter:processDeclAlge( node, prefix, depth )

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

function dumpFilter:processNewAlgeVal( node, prefix, depth )

   dump( prefix, depth, node, node:get_name().txt )
   for __index, exp in pairs( node:get_paramList() ) do
      filter( exp, self, prefix .. "  ", depth + 1 )
   end
   
end

function dumpFilter:processDeclClass( node, prefix, depth )

   dump( prefix, depth, node, node:get_name(  ).txt )
   for index, field in pairs( node:get_fieldList(  ) ) do
      filter( field, self, prefix .. "  ", depth + 1 )
   end
   
end


function dumpFilter:processDeclMember( node, prefix, depth )

   dump( prefix, depth, node, node:get_name(  ).txt )
   filter( node:get_refType(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processExpMacroExp( node, prefix, depth )

   dump( prefix, depth, node, "" )
   local stmtList = node:get_stmtList(  )
   for __index, stmt in pairs( stmtList ) do
      filter( stmt, self, prefix .. "  ", depth + 1 )
   end
   
end


function dumpFilter:processDeclMacro( node, prefix, depth )

   dump( prefix, depth, node, node:get_expType(  ):getTxt(  ) )
end


function dumpFilter:processExpMacroStat( node, prefix, depth )

   dump( prefix, depth, node, node:get_expType(  ):getTxt(  ) )
   for __index, node in pairs( node:get_expStrList(  ) ) do
      filter( node, self, prefix .. "  ", depth + 1 )
   end
   
end


function dumpFilter:processUnwrapSet( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_dstExpList(), self, prefix .. "  ", depth + 1 )
   filter( node:get_srcExpList(), self, prefix .. "  ", depth + 1 )
   if node:get_unwrapBlock() then
      filter( _lune.unwrap( node:get_unwrapBlock()), self, prefix .. "  ", depth + 1 )
   end
   
end

function dumpFilter:processIfUnwrap( node, prefix, depth )

   dump( prefix, depth, node, "" )
   for index, expNode in pairs( node:get_expNodeList() ) do
      filter( expNode, self, prefix .. "  ", depth + 1 )
   end
   
   filter( node:get_block(), self, prefix .. "  ", depth + 1 )
   if node:get_nilBlock() then
      filter( _lune.unwrap( node:get_nilBlock()), self, prefix .. "  ", depth + 1 )
   end
   
end

function dumpFilter:processWhen( node, prefix, depth )

   dump( prefix, depth, node, "" )
   for index, expNode in pairs( node:get_expNodeList() ) do
      filter( expNode, self, prefix .. "  ", depth + 1 )
   end
   
   filter( node:get_block(), self, prefix .. "  ", depth + 1 )
   do
      local _exp = node:get_elseBlock()
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end

function dumpFilter:processDeclVar( node, prefix, depth )

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
            filter( _exp, self, prefix .. "  ", depth + 1 )
         end
      end
      
   end
   
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
   do
      local _exp = node:get_unwrapBlock()
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
   do
      local _exp = node:get_thenBlock()
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
   for __index, var in pairs( node:get_syncVarList() ) do
      do
         local _exp = var:get_refType()
         if _exp ~= nil then
            filter( _exp, self, prefix .. "  ", depth + 1 )
         end
      end
      
   end
   
   do
      local _exp = node:get_syncBlock()
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end


function dumpFilter:processDeclArg( node, prefix, depth )

   dump( prefix, depth, node, node:get_name(  ).txt )
   filter( node:get_argType(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processDeclArgDDD( node, prefix, depth )

   dump( prefix, depth, node, "..." )
end


function dumpFilter:processExpDDD( node, prefix, depth )

   dump( prefix, depth, node, "..." )
end


function dumpFilter:processDeclFuncInfo( node, declInfo, prefix, depth )

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
      filter( arg, self, prefix .. "  ", depth + 1 )
   end
   
   do
      local _exp = declInfo:get_body(  )
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end

function dumpFilter:processDeclFunc( node, prefix, depth )

   self:processDeclFuncInfo( node, node:get_declInfo(  ), prefix, depth )
end


function dumpFilter:processDeclMethod( node, prefix, depth )

   self:processDeclFuncInfo( node, node:get_declInfo(  ), prefix, depth )
end


function dumpFilter:processDeclConstr( node, prefix, depth )

   self:processDeclFuncInfo( node, node:get_declInfo(  ), prefix, depth )
end


function dumpFilter:processDeclDestr( node, prefix, depth )

   dump( prefix, depth, node, "" )
end


function dumpFilter:processExpCallSuper( node, prefix, depth )

   local typeInfo = node:get_superType(  )
   dump( prefix, depth, node, typeInfo:getTxt(  ) )
end


function dumpFilter:processRefType( node, prefix, depth )

   dump( prefix, depth, node, (node:get_refFlag(  ) and "&" or "" ) .. (node:get_mutFlag(  ) and "mut " or "" ) )
   filter( node:get_name(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processIf( node, prefix, depth )

   dump( prefix, depth, node, "" )
   local stmtList = node:get_stmtList(  )
   for index, stmt in pairs( stmtList ) do
      if stmt:get_exp():get_kind() ~= Ast.nodeKind['None'] then
         filter( stmt:get_exp(), self, prefix .. "  ", depth + 1 )
      end
      
      filter( stmt:get_block(), self, prefix .. "  ", depth + 1 )
   end
   
end


function dumpFilter:processSwitch( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
   local caseList = node:get_caseList()
   for __index, caseInfo in pairs( caseList ) do
      filter( caseInfo:get_expList(), self, prefix .. "  ", depth + 1 )
      filter( caseInfo:get_block(), self, prefix .. "  ", depth + 1 )
   end
   
   do
      local _exp = node:get_default(  )
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end


function dumpFilter:processMatch( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_val(), self, prefix .. "  ", depth + 1 )
   local caseList = node:get_caseList()
   for __index, caseInfo in pairs( caseList ) do
      filter( caseInfo:get_block(), self, prefix .. "  " .. caseInfo:get_valInfo():get_name(), depth + 1 )
   end
   
   do
      local _exp = node:get_defaultBlock()
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end


function dumpFilter:processWhile( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
   filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processRepeat( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processFor( node, prefix, depth )

   dump( prefix, depth, node, node:get_val(  ).txt )
   filter( node:get_init(  ), self, prefix .. "  ", depth + 1 )
   filter( node:get_to(  ), self, prefix .. "  ", depth + 1 )
   do
      local _exp = node:get_delta(  )
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
   filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processApply( node, prefix, depth )

   local varNames = ""
   local varList = node:get_varList(  )
   for index, var in pairs( varList ) do
      varNames = varNames .. var.txt .. " "
   end
   
   dump( prefix, depth, node, varNames )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
   filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processForeach( node, prefix, depth )

   local index = ""
   do
      local _exp = node:get_key(  )
      if _exp ~= nil then
         index = _exp.txt
      end
   end
   
   dump( prefix, depth, node, node:get_val(  ).txt .. " " .. index )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
   filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processForsort( node, prefix, depth )

   local index = ""
   do
      local _exp = node:get_key(  )
      if _exp ~= nil then
         index = _exp.txt
      end
   end
   
   dump( prefix, depth, node, node:get_val(  ).txt .. " " .. index )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
   filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processExpUnwrap( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_exp(), self, prefix .. "  ", depth + 1 )
   do
      local _exp = node:get_default()
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end

function dumpFilter:processExpCall( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_func(  ), self, prefix .. "  ", depth + 1 )
   do
      local _exp = node:get_argList(  )
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end


function dumpFilter:processExpList( node, prefix, depth )

   dump( prefix, depth, node, "" )
   local expList = node:get_expList(  )
   for index, exp in pairs( expList ) do
      filter( exp, self, prefix .. "  ", depth + 1 )
   end
   
end


function dumpFilter:processExpOp1( node, prefix, depth )

   dump( prefix, depth, node, node:get_op(  ).txt )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processExpCast( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processExpParen( node, prefix, depth )

   dump( prefix, depth, node, "()" )
   filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processExpOp2( node, prefix, depth )

   dump( prefix, depth, node, node:get_op(  ).txt )
   filter( node:get_exp1(  ), self, prefix .. "  ", depth + 1 )
   filter( node:get_exp2(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processExpNew( node, prefix, depth )

   dump( prefix, depth, node, "" )
   filter( node:get_symbol(  ), self, prefix .. "  ", depth + 1 )
   do
      local _exp = node:get_argList(  )
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end


function dumpFilter:processExpRef( node, prefix, depth )

   dump( prefix, depth, node, node:get_token(  ).txt )
end


function dumpFilter:processExpRefItem( node, prefix, depth )

   dump( prefix, depth, node, "seq[exp] " .. node:get_expType(  ):getTxt(  ) )
   filter( node:get_val(), self, prefix .. "  ", depth + 1 )
   do
      local _exp = node:get_index()
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end


function dumpFilter:processRefField( node, prefix, depth )

   dump( prefix, depth, node, node:get_field(  ).txt )
   filter( node:get_prefix(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processExpOmitEnum( node, prefix, depth )

   dump( prefix, depth, node, string.format( "%s.%s", node:get_expType():getTxt(  ), node:get_valToken().txt) )
end


function dumpFilter:processGetField( node, prefix, depth )

   dump( prefix, depth, node, "get_" .. node:get_field(  ).txt )
   filter( node:get_prefix(  ), self, prefix .. "  ", depth + 1 )
end


function dumpFilter:processReturn( node, prefix, depth )

   dump( prefix, depth, node, "" )
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end


function dumpFilter:processProvide( node, prefix, depth )

   dump( prefix, depth, node, node:get_symbol():get_name() )
end

function dumpFilter:processLiteralList( node, prefix, depth )

   dump( prefix, depth, node, "[]" )
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end


function dumpFilter:processLiteralMap( node, prefix, depth )

   dump( prefix, depth, node, "{}" )
   local pairList = node:get_pairList(  )
   for __index, pair in pairs( pairList ) do
      filter( pair:get_key(  ), self, prefix .. "  ", depth + 1 )
      filter( pair:get_val(  ), self, prefix .. "  ", depth + 1 )
   end
   
end


function dumpFilter:processLiteralArray( node, prefix, depth )

   dump( prefix, depth, node, "[@]" )
   do
      local _exp = node:get_expList(  )
      if _exp ~= nil then
         filter( _exp, self, prefix .. "  ", depth + 1 )
      end
   end
   
end


function dumpFilter:processLiteralChar( node, prefix, depth )

   dump( prefix, depth, node, string.format( "%s(%s)", tostring( node:get_num(  )), node:get_token(  ).txt ) )
end


function dumpFilter:processLiteralInt( node, prefix, depth )

   dump( prefix, depth, node, string.format( "%s(%s)", tostring( node:get_num(  )), node:get_token(  ).txt ) )
end


function dumpFilter:processLiteralReal( node, prefix, depth )

   dump( prefix, depth, node, string.format( "%s(%s)", tostring( node:get_num(  )), node:get_token(  ).txt ) )
end


function dumpFilter:processLiteralString( node, prefix, depth )

   dump( prefix, depth, node, node:get_token(  ).txt )
end


function dumpFilter:processLiteralBool( node, prefix, depth )

   dump( prefix, depth, node, node:get_token(  ).txt == "true" and "true" or "false" )
end


function dumpFilter:processLiteralNil( node, prefix, depth )

   dump( prefix, depth, node, "" )
end


function dumpFilter:processBreak( node, prefix, depth )

   dump( prefix, depth, node, "" )
end


function dumpFilter:processLiteralSymbol( node, prefix, depth )

   dump( prefix, depth, node, node:get_token(  ).txt )
end


return _moduleObj
