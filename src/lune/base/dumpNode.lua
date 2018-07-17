--lune/base/dumpNode.lns
local moduleObj = {}
local TransUnit = require( 'lune.base.TransUnit' )

local Parser = require( 'lune.base.Parser' )

local dumpFilter = {}
setmetatable( dumpFilter, { __index = Filter } )
moduleObj.dumpFilter = dumpFilter
            function dumpFilter.new(  )
            local obj = {}
            setmetatable( obj, { __index = dumpFilter } )
            if obj.__init then
            obj:__init(  )
            end       
            return obj
            end         
            function dumpFilter:__init(  )
            
end

local function dump( prefix, depth, node, txt )
  local typeStr = ""
  local expType = node:get_expType(  )
  if expType and expType ~= TransUnit.typeInfoKind.None then
    typeStr = string.format( "(%d:%s:%s)", expType:get_typeId(  ), expType:getTxt(  ), expType:get_kind(  ))
  end
  print( string.format( "%s: %s %s %s", prefix, TransUnit.getNodeKindName( node:get_kind(  ) ), txt, typeStr) )
end

local function filter( node, filter, prefix, depth )
  node:processFilter( dumpFilter, prefix, depth )
end

local function getTxt( token )
  return token.txt
end

function dumpFilter:processNone( node, prefix, depth )
  dump( prefix, depth, node, "" )
end

-- none

function dumpFilter:processImport( node, prefix, depth )
  dump( prefix, depth, node, node:get_modulePath(  ) )
end

-- none

function dumpFilter:processRoot( node, prefix, depth )
  dump( prefix, depth, node, "" )
  for index, child in pairs( node:get_children(  ) ) do
    filter( child, self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processBlock( node, prefix, depth )
  dump( prefix, depth, node, "" )
  for index, statement in pairs( node:get_stmtList(  ) ) do
    filter( statement, self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processStmtExp( node, prefix, depth )
  dump( prefix, depth, node, "" )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processDeclClass( node, prefix, depth )
  dump( prefix, depth, node, node:get_name(  ).txt )
  for index, field in pairs( node:get_fieldList(  ) ) do
    filter( field, self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processDeclMember( node, prefix, depth )
  dump( prefix, depth, node, node:get_name(  ).txt )
  filter( node:get_refType(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processExpMacroExp( node, prefix, depth )
  dump( prefix, depth, node, "" )
  local stmtList = node:get_stmtList(  )
  if stmtList then
    for __index, stmt in pairs( stmtList ) do
      filter( stmt, self, prefix .. "  ", depth + 1 )
    end
  end
end

-- none

function dumpFilter:processDeclMacro( node, prefix, depth )
  dump( prefix, depth, node, node:get_expType(  ):getTxt(  ) )
end

-- none

function dumpFilter:processExpMacroStat( node, prefix, depth )
  dump( prefix, depth, node, node:get_expType(  ):getTxt(  ) )
  for __index, node in pairs( node:get_expStrList(  ) ) do
    filter( node, self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processDeclVar( node, prefix, depth )
  local varName = ""
  for index, var in pairs( node:get_varList(  ) ) do
    if index > 1 then
      varName = varName .. ","
    end
    varName = string.format( "%s %s", varName, var:get_name(  ).txt)
  end
  dump( prefix, depth, node, node:get_unwrap(  ) and "! " or " " .. varName )
  for index, var in pairs( node:get_varList(  ) ) do
    if var["refType"] then
      filter( var["refType"], self, prefix .. "  ", depth + 1 )
    end
  end
  if node:get_expList(  ) then
    filter( node:get_expList(  ), self, prefix .. "  ", depth + 1 )
  end
  if node:get_unwrap(  ) then
    filter( node:get_unwrap(  ), self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processDeclArg( node, prefix, depth )
  dump( prefix, depth, node, node:get_name(  ).txt )
  filter( node:get_argType(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processDeclArgDDD( node, prefix, depth )
  dump( prefix, depth, node, "..." )
end

-- none

function dumpFilter:processExpDDD( node, prefix, depth )
  dump( prefix, depth, node, "..." )
end

-- none

function dumpFilter:processDeclFuncInfo( node, declInfo, prefix, depth )
  local name = declInfo:get_name(  )
  dump( prefix, depth, node, name and name.txt or "<anonymous>" )
  local argList = declInfo:get_argList(  )
  for index, arg in pairs( argList ) do
    filter( arg, self, prefix .. "  ", depth + 1 )
  end
  local retTypeList = declInfo:get_retTypeList(  )
  for index, refType in pairs( retTypeList ) do
    filter( refType, self, prefix .. "  ", depth + 1 )
  end
  if declInfo:get_body(  ) then
    filter( declInfo:get_body(  ), self, prefix .. "  ", depth + 1 )
  end
end

function dumpFilter:processDeclFunc( node, prefix, depth )
  self:processDeclFuncInfo( node, node:get_declInfo(  ), prefix, depth )
end

-- none

function dumpFilter:processDeclMethod( node, prefix, depth )
  self:processDeclFuncInfo( node, node:get_declInfo(  ), prefix, depth )
end

-- none

function dumpFilter:processDeclConstr( node, prefix, depth )
  self:processDeclFuncInfo( node, node:get_declInfo(  ), prefix, depth )
end

-- none

function dumpFilter:processExpCallSuper( node, prefix, depth )
  local typeInfo = node:get_superType(  )
  dump( prefix, depth, node, typeInfo:getTxt(  ) )
end

-- none

function dumpFilter:processRefType( node, prefix, depth )
  dump( prefix, depth, node, (node:get_refFlag(  ) and "&" or "" ) .. (node:get_mutFlag(  ) and "mut " or "" ) )
  filter( node:get_name(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processIf( node, prefix, depth )
  dump( prefix, depth, node, "" )
  local stmtList = node:get_stmtList(  )
  for index, stmt in pairs( stmtList ) do
    if stmt['exp'] then
      filter( stmt['exp'], self, prefix .. "  ", depth + 1 )
    end
    filter( stmt['block'], self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processSwitch( node, prefix, depth )
  dump( prefix, depth, node, "" )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
  local caseList = node:get_caseList(  )
  for __index, caseInfo in pairs( caseList ) do
    filter( caseInfo['expList'], self, prefix .. "  ", depth + 1 )
    filter( caseInfo['block'], self, prefix .. "  ", depth + 1 )
  end
  filter( node:get_default(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processWhile( node, prefix, depth )
  dump( prefix, depth, node, "" )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processRepeat( node, prefix, depth )
  dump( prefix, depth, node, "" )
  filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processFor( node, prefix, depth )
  dump( prefix, depth, node, node:get_val(  ).txt )
  filter( node:get_init(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_to(  ), self, prefix .. "  ", depth + 1 )
  if node:get_delta(  ) then
    filter( node:get_delta(  ), self, prefix .. "  ", depth + 1 )
  end
  filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processApply( node, prefix, depth )
  local varNames = ""
  local varList = node:get_varList(  )
  for index, var in pairs( varList ) do
    varNames = varNames .. var['txt'] .. " "
  end
  dump( prefix, depth, node, varNames )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processForeach( node, prefix, depth )
  local index = node:get_key(  ) and node:get_key(  ).txt or ""
  dump( prefix, depth, node, node:get_val(  ).txt .. " " .. index )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processForsort( node, prefix, depth )
  local index = node:get_key(  ) and node:get_key(  ).txt or ""
  dump( prefix, depth, node, node:get_val(  ).txt .. " " .. index )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processExpCall( node, prefix, depth )
  dump( prefix, depth, node, "" )
  filter( node:get_func(  ), self, prefix .. "  ", depth + 1 )
  if node:get_argList(  ) then
    filter( node:get_argList(  ), self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processExpList( node, prefix, depth )
  dump( prefix, depth, node, "" )
  local expList = node:get_expList(  )
  for index, exp in pairs( expList ) do
    filter( exp, self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processExpOp1( node, prefix, depth )
  dump( prefix, depth, node, node:get_op(  ).txt )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processExpCast( node, prefix, depth )
  dump( prefix, depth, node, "" )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processExpParen( node, prefix, depth )
  dump( prefix, depth, node, "()" )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processExpOp2( node, prefix, depth )
  dump( prefix, depth, node, node:get_op(  ).txt )
  filter( node:get_exp1(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_exp2(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processExpNew( node, prefix, depth )
  dump( prefix, depth, node, "" )
  filter( node:get_symbol(  ), self, prefix .. "  ", depth + 1 )
  if node:get_argList(  ) then
    filter( node:get_argList(  ), self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processExpRef( node, prefix, depth )
  dump( prefix, depth, node, node:get_token(  ).txt )
end

-- none

function dumpFilter:processExpRefItem( node, prefix, depth )
  dump( prefix, depth, node, "seq[exp] " .. node:get_expType(  ):getTxt(  ) )
  filter( node:get_val(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_index(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processRefField( node, prefix, depth )
  dump( prefix, depth, node, node:get_field(  ).txt )
  filter( node:get_prefix(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processGetField( node, prefix, depth )
  dump( prefix, depth, node, (node:get_getterTypeInfo(  ) and "get_" or "" ) .. node:get_field(  ).txt )
  filter( node:get_prefix(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processReturn( node, prefix, depth )
  dump( prefix, depth, node, "" )
  if node:get_expList(  ) then
    filter( node:get_expList(  ), self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processLiteralList( node, prefix, depth )
  dump( prefix, depth, node, "[]" )
  if node:get_expList(  ) then
    filter( node:get_expList(  ), self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processLiteralMap( node, prefix, depth )
  dump( prefix, depth, node, "{}" )
  local pairList = node:get_pairList(  )
  for __index, pair in pairs( pairList ) do
    filter( pair:get_key(  ), self, prefix .. "  ", depth + 1 )
    filter( pair:get_val(  ), self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processLiteralArray( node, prefix, depth )
  dump( prefix, depth, node, "[@]" )
  if node:get_expList(  ) then
    filter( node:get_expList(  ), self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processLiteralChar( node, prefix, depth )
  dump( prefix, depth, node, string.format( "%s(%s)", node:get_num(  ), node:get_token(  ).txt ) )
end

-- none

function dumpFilter:processLiteralInt( node, prefix, depth )
  dump( prefix, depth, node, string.format( "%s(%s)", node:get_num(  ), node:get_token(  ).txt ) )
end

-- none

function dumpFilter:processLiteralReal( node, prefix, depth )
  dump( prefix, depth, node, string.format( "%s(%s)", node:get_num(  ), node:get_token(  ).txt ) )
end

-- none

function dumpFilter:processLiteralString( node, prefix, depth )
  dump( prefix, depth, node, node:get_token(  ).txt )
end

-- none

function dumpFilter:processLiteralBool( node, prefix, depth )
  dump( prefix, depth, node, node:get_token(  ).txt == "true" )
end

-- none

function dumpFilter:processLiteralNil( node, prefix, depth )
  dump( prefix, depth, node, "" )
end

-- none

function dumpFilter:processBreak( node, prefix, depth )
  dump( prefix, depth, node, "" )
end

-- none

function dumpFilter:processLiteralSymbol( node, prefix, depth )
  dump( prefix, depth, node, node:get_token(  ).txt )
end

-- none

----- meta -----
local _typeId2ClassInfoMap = {}
moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfo4092 = {}
  _className2InfoMap.dumpFilter = _classInfo4092
  _typeId2ClassInfoMap[ 4092 ] = _classInfo4092
  end
do
  local _classInfo4032 = {}
  _className2InfoMap.ASTInfo = _classInfo4032
  _typeId2ClassInfoMap[ 4032 ] = _classInfo4032
  end
do
  local _classInfo3434 = {}
  _className2InfoMap.ApplyNode = _classInfo3434
  _typeId2ClassInfoMap[ 3434 ] = _classInfo3434
  end
do
  local _classInfo3324 = {}
  _className2InfoMap.BlockNode = _classInfo3324
  _typeId2ClassInfoMap[ 3324 ] = _classInfo3324
  end
do
  local _classInfo3494 = {}
  _className2InfoMap.BreakNode = _classInfo3494
  _typeId2ClassInfoMap[ 3494 ] = _classInfo3494
  end
do
  local _classInfo3370 = {}
  _className2InfoMap.CaseInfo = _classInfo3370
  _typeId2ClassInfoMap[ 3370 ] = _classInfo3370
  end
do
  local _classInfo3794 = {}
  _className2InfoMap.DeclArgDDDNode = _classInfo3794
  _typeId2ClassInfoMap[ 3794 ] = _classInfo3794
  end
do
  local _classInfo3782 = {}
  _className2InfoMap.DeclArgNode = _classInfo3782
  _typeId2ClassInfoMap[ 3782 ] = _classInfo3782
  end
do
  local _classInfo3140 = {}
  _className2InfoMap.DeclClassNode = _classInfo3140
  _typeId2ClassInfoMap[ 3140 ] = _classInfo3140
  end
do
  local _classInfo3740 = {}
  _className2InfoMap.DeclConstrNode = _classInfo3740
  _typeId2ClassInfoMap[ 3740 ] = _classInfo3740
  end
do
  local _classInfo3696 = {}
  _className2InfoMap.DeclFuncInfo = _classInfo3696
  _typeId2ClassInfoMap[ 3696 ] = _classInfo3696
  end
do
  local _classInfo3720 = {}
  _className2InfoMap.DeclFuncNode = _classInfo3720
  _typeId2ClassInfoMap[ 3720 ] = _classInfo3720
  end
do
  local _classInfo3250 = {}
  _className2InfoMap.DeclMacroInfo = _classInfo3250
  _typeId2ClassInfoMap[ 3250 ] = _classInfo3250
  end
do
  local _classInfo3826 = {}
  _className2InfoMap.DeclMacroNode = _classInfo3826
  _typeId2ClassInfoMap[ 3826 ] = _classInfo3826
  end
do
  local _classInfo3762 = {}
  _className2InfoMap.DeclMemberNode = _classInfo3762
  _typeId2ClassInfoMap[ 3762 ] = _classInfo3762
  end
do
  local _classInfo3730 = {}
  _className2InfoMap.DeclMethodNode = _classInfo3730
  _typeId2ClassInfoMap[ 3730 ] = _classInfo3730
  end
do
  local _classInfo3672 = {}
  _className2InfoMap.DeclVarNode = _classInfo3672
  _typeId2ClassInfoMap[ 3672 ] = _classInfo3672
  end
do
  local _classInfo3572 = {}
  _className2InfoMap.ExpCallNode = _classInfo3572
  _typeId2ClassInfoMap[ 3572 ] = _classInfo3572
  end
do
  local _classInfo3750 = {}
  _className2InfoMap.ExpCallSuperNode = _classInfo3750
  _typeId2ClassInfoMap[ 3750 ] = _classInfo3750
  end
do
  local _classInfo3538 = {}
  _className2InfoMap.ExpCastNode = _classInfo3538
  _typeId2ClassInfoMap[ 3538 ] = _classInfo3538
  end
do
  local _classInfo3584 = {}
  _className2InfoMap.ExpDDDNode = _classInfo3584
  _typeId2ClassInfoMap[ 3584 ] = _classInfo3584
  end
do
  local _classInfo3358 = {}
  _className2InfoMap.ExpListNode = _classInfo3358
  _typeId2ClassInfoMap[ 3358 ] = _classInfo3358
  end
do
  local _classInfo3604 = {}
  _className2InfoMap.ExpMacroExpNode = _classInfo3604
  _typeId2ClassInfoMap[ 3604 ] = _classInfo3604
  end
do
  local _classInfo3616 = {}
  _className2InfoMap.ExpMacroStatNode = _classInfo3616
  _typeId2ClassInfoMap[ 3616 ] = _classInfo3616
  end
do
  local _classInfo3502 = {}
  _className2InfoMap.ExpNewNode = _classInfo3502
  _typeId2ClassInfoMap[ 3502 ] = _classInfo3502
  end
do
  local _classInfo3548 = {}
  _className2InfoMap.ExpOp1Node = _classInfo3548
  _typeId2ClassInfoMap[ 3548 ] = _classInfo3548
  end
do
  local _classInfo3524 = {}
  _className2InfoMap.ExpOp2Node = _classInfo3524
  _typeId2ClassInfoMap[ 3524 ] = _classInfo3524
  end
do
  local _classInfo3594 = {}
  _className2InfoMap.ExpParenNode = _classInfo3594
  _typeId2ClassInfoMap[ 3594 ] = _classInfo3594
  end
do
  local _classInfo3560 = {}
  _className2InfoMap.ExpRefItemNode = _classInfo3560
  _typeId2ClassInfoMap[ 3560 ] = _classInfo3560
  end
do
  local _classInfo3514 = {}
  _className2InfoMap.ExpRefNode = _classInfo3514
  _typeId2ClassInfoMap[ 3514 ] = _classInfo3514
  end
do
  local _classInfo3226 = {}
  _className2InfoMap.Filter = _classInfo3226
  _typeId2ClassInfoMap[ 3226 ] = _classInfo3226
  end
do
  local _classInfo3416 = {}
  _className2InfoMap.ForNode = _classInfo3416
  _typeId2ClassInfoMap[ 3416 ] = _classInfo3416
  end
do
  local _classInfo3450 = {}
  _className2InfoMap.ForeachNode = _classInfo3450
  _typeId2ClassInfoMap[ 3450 ] = _classInfo3450
  end
do
  local _classInfo3466 = {}
  _className2InfoMap.ForsortNode = _classInfo3466
  _typeId2ClassInfoMap[ 3466 ] = _classInfo3466
  end
do
  local _classInfo3650 = {}
  _className2InfoMap.GetFieldNode = _classInfo3650
  _typeId2ClassInfoMap[ 3650 ] = _classInfo3650
  end
do
  local _classInfo3346 = {}
  _className2InfoMap.IfNode = _classInfo3346
  _typeId2ClassInfoMap[ 3346 ] = _classInfo3346
  end
do
  local _classInfo3338 = {}
  _className2InfoMap.IfStmtInfo = _classInfo3338
  _typeId2ClassInfoMap[ 3338 ] = _classInfo3338
  end
do
  local _classInfo3282 = {}
  _className2InfoMap.ImportNode = _classInfo3282
  _typeId2ClassInfoMap[ 3282 ] = _classInfo3282
  end
do
  local _classInfo3880 = {}
  _className2InfoMap.LiteralArrayNode = _classInfo3880
  _typeId2ClassInfoMap[ 3880 ] = _classInfo3880
  end
do
  local _classInfo3936 = {}
  _className2InfoMap.LiteralBoolNode = _classInfo3936
  _typeId2ClassInfoMap[ 3936 ] = _classInfo3936
  end
do
  local _classInfo3844 = {}
  _className2InfoMap.LiteralCharNode = _classInfo3844
  _typeId2ClassInfoMap[ 3844 ] = _classInfo3844
  end
do
  local _classInfo3856 = {}
  _className2InfoMap.LiteralIntNode = _classInfo3856
  _typeId2ClassInfoMap[ 3856 ] = _classInfo3856
  end
do
  local _classInfo3890 = {}
  _className2InfoMap.LiteralListNode = _classInfo3890
  _typeId2ClassInfoMap[ 3890 ] = _classInfo3890
  end
do
  local _classInfo3906 = {}
  _className2InfoMap.LiteralMapNode = _classInfo3906
  _typeId2ClassInfoMap[ 3906 ] = _classInfo3906
  end
do
  local _classInfo3836 = {}
  _className2InfoMap.LiteralNilNode = _classInfo3836
  _typeId2ClassInfoMap[ 3836 ] = _classInfo3836
  end
do
  local _classInfo3868 = {}
  _className2InfoMap.LiteralRealNode = _classInfo3868
  _typeId2ClassInfoMap[ 3868 ] = _classInfo3868
  end
do
  local _classInfo3922 = {}
  _className2InfoMap.LiteralStringNode = _classInfo3922
  _typeId2ClassInfoMap[ 3922 ] = _classInfo3922
  end
do
  local _classInfo3946 = {}
  _className2InfoMap.LiteralSymbolNode = _classInfo3946
  _typeId2ClassInfoMap[ 3946 ] = _classInfo3946
  end
do
  local _classInfo3248 = {}
  _className2InfoMap.MacroEval = _classInfo3248
  _typeId2ClassInfoMap[ 3248 ] = _classInfo3248
  end
do
  local _classInfo3246 = {}
  _className2InfoMap.NamespaceInfo = _classInfo3246
  _typeId2ClassInfoMap[ 3246 ] = _classInfo3246
  _classInfo3246.name = {
    name='name', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 18 }
  _classInfo3246.scope = {
    name='scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3202 }
  _classInfo3246.typeInfo = {
    name='typeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3132 }
  end
do
  local _classInfo3138 = {}
  _className2InfoMap.Node = _classInfo3138
  _typeId2ClassInfoMap[ 3138 ] = _classInfo3138
  end
do
  local _classInfo3228 = {}
  _className2InfoMap.NodePos = _classInfo3228
  _typeId2ClassInfoMap[ 3228 ] = _classInfo3228
  end
do
  local _classInfo3274 = {}
  _className2InfoMap.NoneNode = _classInfo3274
  _typeId2ClassInfoMap[ 3274 ] = _classInfo3274
  end
do
  local _classInfo3900 = {}
  _className2InfoMap.PairItem = _classInfo3900
  _typeId2ClassInfoMap[ 3900 ] = _classInfo3900
  end
do
  local _classInfo3062 = {}
  _className2InfoMap.Parser = _classInfo3062
  _typeId2ClassInfoMap[ 3062 ] = _classInfo3062
  _classInfo3062.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3084 }
  _classInfo3062.Position = {
    name='Position', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3074 }
  _classInfo3062.Stream = {
    name='Stream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3064 }
  _classInfo3062.StreamParser = {
    name='StreamParser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3096 }
  _classInfo3062.Token = {
    name='Token', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3076 }
  _classInfo3062.TxtStream = {
    name='TxtStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3068 }
  _classInfo3062.WrapParser = {
    name='WrapParser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3090 }
  _classInfo3062.getEofToken = {
    name='getEofToken', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3110 }
  _classInfo3062.getKindTxt = {
    name='getKindTxt', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3102 }
  _classInfo3062.isOp1 = {
    name='isOp1', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3106 }
  _classInfo3062.isOp2 = {
    name='isOp2', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3104 }
  end
do
  local _classInfo3074 = {}
  _className2InfoMap.Position = _classInfo3074
  _typeId2ClassInfoMap[ 3074 ] = _classInfo3074
  _classInfo3074.column = {
    name='column', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo3074.lineNo = {
    name='lineNo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  end
do
  local _classInfo3638 = {}
  _className2InfoMap.RefFieldNode = _classInfo3638
  _typeId2ClassInfoMap[ 3638 ] = _classInfo3638
  end
do
  local _classInfo3308 = {}
  _className2InfoMap.RefTypeNode = _classInfo3308
  _typeId2ClassInfoMap[ 3308 ] = _classInfo3308
  end
do
  local _classInfo3404 = {}
  _className2InfoMap.RepeatNode = _classInfo3404
  _typeId2ClassInfoMap[ 3404 ] = _classInfo3404
  end
do
  local _classInfo3484 = {}
  _className2InfoMap.ReturnNode = _classInfo3484
  _typeId2ClassInfoMap[ 3484 ] = _classInfo3484
  end
do
  local _classInfo3292 = {}
  _className2InfoMap.RootNode = _classInfo3292
  _typeId2ClassInfoMap[ 3292 ] = _classInfo3292
  end
do
  local _classInfo3202 = {}
  _className2InfoMap.Scope = _classInfo3202
  _typeId2ClassInfoMap[ 3202 ] = _classInfo3202
  end
do
  local _classInfo3628 = {}
  _className2InfoMap.StmtExpNode = _classInfo3628
  _typeId2ClassInfoMap[ 3628 ] = _classInfo3628
  end
do
  local _classInfo3064 = {}
  _className2InfoMap.Stream = _classInfo3064
  _typeId2ClassInfoMap[ 3064 ] = _classInfo3064
  end
do
  local _classInfo4074 = {}
  _className2InfoMap.StreamParser = _classInfo4074
  _typeId2ClassInfoMap[ 4074 ] = _classInfo4074
  _classInfo4074.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 4078 }
  end
do
  local _classInfo3376 = {}
  _className2InfoMap.SwitchNode = _classInfo3376
  _typeId2ClassInfoMap[ 3376 ] = _classInfo3376
  end
do
  local _classInfo4054 = {}
  _className2InfoMap.Token = _classInfo4054
  _typeId2ClassInfoMap[ 4054 ] = _classInfo4054
  _classInfo4054.kind = {
    name='kind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo4054.pos = {
    name='pos', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4052 }
  _classInfo4054.txt = {
    name='txt', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 18 }
  end
do
  local _classInfo110 = {}
  _className2InfoMap.TransUnit = _classInfo110
  _typeId2ClassInfoMap[ 110 ] = _classInfo110
  _classInfo110.ASTInfo = {
    name='ASTInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4032 }
  _classInfo110.ApplyNode = {
    name='ApplyNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3434 }
  _classInfo110.BlockNode = {
    name='BlockNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3324 }
  _classInfo110.BreakNode = {
    name='BreakNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3494 }
  _classInfo110.CaseInfo = {
    name='CaseInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3370 }
  _classInfo110.DeclArgDDDNode = {
    name='DeclArgDDDNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3794 }
  _classInfo110.DeclArgNode = {
    name='DeclArgNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3782 }
  _classInfo110.DeclClassNode = {
    name='DeclClassNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3140 }
  _classInfo110.DeclConstrNode = {
    name='DeclConstrNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3740 }
  _classInfo110.DeclFuncInfo = {
    name='DeclFuncInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3696 }
  _classInfo110.DeclFuncNode = {
    name='DeclFuncNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3720 }
  _classInfo110.DeclMacroInfo = {
    name='DeclMacroInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3250 }
  _classInfo110.DeclMacroNode = {
    name='DeclMacroNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3826 }
  _classInfo110.DeclMemberNode = {
    name='DeclMemberNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3762 }
  _classInfo110.DeclMethodNode = {
    name='DeclMethodNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3730 }
  _classInfo110.DeclVarNode = {
    name='DeclVarNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3672 }
  _classInfo110.ExpCallNode = {
    name='ExpCallNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3572 }
  _classInfo110.ExpCallSuperNode = {
    name='ExpCallSuperNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3750 }
  _classInfo110.ExpCastNode = {
    name='ExpCastNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3538 }
  _classInfo110.ExpDDDNode = {
    name='ExpDDDNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3584 }
  _classInfo110.ExpListNode = {
    name='ExpListNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3358 }
  _classInfo110.ExpMacroExpNode = {
    name='ExpMacroExpNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3604 }
  _classInfo110.ExpMacroStatNode = {
    name='ExpMacroStatNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3616 }
  _classInfo110.ExpNewNode = {
    name='ExpNewNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3502 }
  _classInfo110.ExpOp1Node = {
    name='ExpOp1Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3548 }
  _classInfo110.ExpOp2Node = {
    name='ExpOp2Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3524 }
  _classInfo110.ExpParenNode = {
    name='ExpParenNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3594 }
  _classInfo110.ExpRefItemNode = {
    name='ExpRefItemNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3560 }
  _classInfo110.ExpRefNode = {
    name='ExpRefNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3514 }
  _classInfo110.Filter = {
    name='Filter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3226 }
  _classInfo110.ForNode = {
    name='ForNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3416 }
  _classInfo110.ForeachNode = {
    name='ForeachNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3450 }
  _classInfo110.ForsortNode = {
    name='ForsortNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3466 }
  _classInfo110.GetFieldNode = {
    name='GetFieldNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3650 }
  _classInfo110.IfNode = {
    name='IfNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3346 }
  _classInfo110.IfStmtInfo = {
    name='IfStmtInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3338 }
  _classInfo110.ImportNode = {
    name='ImportNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3282 }
  _classInfo110.LiteralArrayNode = {
    name='LiteralArrayNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3880 }
  _classInfo110.LiteralBoolNode = {
    name='LiteralBoolNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3936 }
  _classInfo110.LiteralCharNode = {
    name='LiteralCharNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3844 }
  _classInfo110.LiteralIntNode = {
    name='LiteralIntNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3856 }
  _classInfo110.LiteralListNode = {
    name='LiteralListNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3890 }
  _classInfo110.LiteralMapNode = {
    name='LiteralMapNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3906 }
  _classInfo110.LiteralNilNode = {
    name='LiteralNilNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3836 }
  _classInfo110.LiteralRealNode = {
    name='LiteralRealNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3868 }
  _classInfo110.LiteralStringNode = {
    name='LiteralStringNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3922 }
  _classInfo110.LiteralSymbolNode = {
    name='LiteralSymbolNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3946 }
  _classInfo110.MacroEval = {
    name='MacroEval', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3248 }
  _classInfo110.NamespaceInfo = {
    name='NamespaceInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3246 }
  _classInfo110.Node = {
    name='Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3138 }
  _classInfo110.NodePos = {
    name='NodePos', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3228 }
  _classInfo110.NoneNode = {
    name='NoneNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3274 }
  _classInfo110.PairItem = {
    name='PairItem', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3900 }
  _classInfo110.RefFieldNode = {
    name='RefFieldNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3638 }
  _classInfo110.RefTypeNode = {
    name='RefTypeNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3308 }
  _classInfo110.RepeatNode = {
    name='RepeatNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3404 }
  _classInfo110.ReturnNode = {
    name='ReturnNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3484 }
  _classInfo110.RootNode = {
    name='RootNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3292 }
  _classInfo110.Scope = {
    name='Scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3202 }
  _classInfo110.StmtExpNode = {
    name='StmtExpNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3628 }
  _classInfo110.SwitchNode = {
    name='SwitchNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3376 }
  _classInfo110.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3264 }
  _classInfo110.TypeInfo = {
    name='TypeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3132 }
  _classInfo110.TypeInfoKindArray = {
    name='TypeInfoKindArray', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindClass = {
    name='TypeInfoKindClass', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindFunc = {
    name='TypeInfoKindFunc', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindList = {
    name='TypeInfoKindList', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindMacro = {
    name='TypeInfoKindMacro', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindMap = {
    name='TypeInfoKindMap', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindMethod = {
    name='TypeInfoKindMethod', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindModule = {
    name='TypeInfoKindModule', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindNilable = {
    name='TypeInfoKindNilable', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindPrim = {
    name='TypeInfoKindPrim', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.TypeInfoKindRoot = {
    name='TypeInfoKindRoot', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.VarInfo = {
    name='VarInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3664 }
  _classInfo110.WhileNode = {
    name='WhileNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3392 }
  _classInfo110.getNodeKindName = {
    name='getNodeKindName', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3272 }
  _classInfo110.isBuiltin = {
    name='isBuiltin', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3136 }
  _classInfo110.lune = {
    name='lune', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3058 }
  _classInfo110.nodeKind = {
    name='nodeKind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3270 }
  _classInfo110.rootTypeId = {
    name='rootTypeId', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo110.typeInfoKind = {
    name='typeInfoKind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3134 }
  end
do
  local _classInfo3068 = {}
  _className2InfoMap.TxtStream = _classInfo3068
  _typeId2ClassInfoMap[ 3068 ] = _classInfo3068
  end
do
  local _classInfo3132 = {}
  _className2InfoMap.TypeInfo = _classInfo3132
  _typeId2ClassInfoMap[ 3132 ] = _classInfo3132
  _classInfo3132.cloneToPublic = {
    name='cloneToPublic', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3152 }
  _classInfo3132.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3154 }
  _classInfo3132.createArray = {
    name='createArray', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3160 }
  _classInfo3132.createBuiltin = {
    name='createBuiltin', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3156 }
  _classInfo3132.createClass = {
    name='createClass', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3164 }
  _classInfo3132.createFunc = {
    name='createFunc', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3166 }
  _classInfo3132.createList = {
    name='createList', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3158 }
  _classInfo3132.createMap = {
    name='createMap', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3162 }
  end
do
  local _classInfo3112 = {}
  _className2InfoMap.Util = _classInfo3112
  _typeId2ClassInfoMap[ 3112 ] = _classInfo3112
  _classInfo3112.debugLog = {
    name='debugLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3128 }
  _classInfo3112.errorLog = {
    name='errorLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3126 }
  _classInfo3112.memStream = {
    name='memStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3118 }
  _classInfo3112.outStream = {
    name='outStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3114 }
  _classInfo3112.profile = {
    name='profile', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3130 }
  end
do
  local _classInfo3664 = {}
  _className2InfoMap.VarInfo = _classInfo3664
  _typeId2ClassInfoMap[ 3664 ] = _classInfo3664
  end
do
  local _classInfo3392 = {}
  _className2InfoMap.WhileNode = _classInfo3392
  _typeId2ClassInfoMap[ 3392 ] = _classInfo3392
  end
do
  local _classInfo4068 = {}
  _className2InfoMap.WrapParser = _classInfo4068
  _typeId2ClassInfoMap[ 4068 ] = _classInfo4068
  end
do
  local _classInfo3060 = {}
  _className2InfoMap.base = _classInfo3060
  _typeId2ClassInfoMap[ 3060 ] = _classInfo3060
  _classInfo3060.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3062 }
  _classInfo3060.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3112 }
  end
do
  local _classInfo104 = {}
  _className2InfoMap.dumpNode = _classInfo104
  _typeId2ClassInfoMap[ 104 ] = _classInfo104
  _classInfo104.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4040 }
  _classInfo104.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 110 }
  _classInfo104.dumpFilter = {
    name='dumpFilter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4092 }
  _classInfo104.lune = {
    name='lune', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 106 }
  end
do
  local _classInfo3058 = {}
  _className2InfoMap.lune = _classInfo3058
  _typeId2ClassInfoMap[ 3058 ] = _classInfo3058
  _classInfo3058.base = {
    name='base', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3060 }
  end
do
  local _classInfo3118 = {}
  _className2InfoMap.memStream = _classInfo3118
  _typeId2ClassInfoMap[ 3118 ] = _classInfo3118
  end
do
  local _classInfo3114 = {}
  _className2InfoMap.outStream = _classInfo3114
  _typeId2ClassInfoMap[ 3114 ] = _classInfo3114
  end
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 100, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {102}, }
_typeInfoList[2] = { parentId = 100, typeId = 102, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {104}, }
_typeInfoList[3] = { parentId = 102, typeId = 104, baseId = 1, txt = 'dumpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {106, 4092}, }
_typeInfoList[4] = { parentId = 104, typeId = 106, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {108}, }
_typeInfoList[5] = { parentId = 106, typeId = 108, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {110, 4040}, }
_typeInfoList[6] = { parentId = 108, typeId = 110, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3058, 3132, 3136, 3138, 3140, 3202, 3226, 3228, 3246, 3248, 3250, 3264, 3272, 3274, 3282, 3292, 3308, 3324, 3338, 3346, 3358, 3370, 3376, 3392, 3404, 3416, 3434, 3450, 3466, 3484, 3494, 3502, 3514, 3524, 3538, 3548, 3560, 3572, 3584, 3594, 3604, 3616, 3628, 3638, 3650, 3664, 3672, 3696, 3720, 3730, 3740, 3750, 3762, 3782, 3794, 3826, 3836, 3844, 3856, 3868, 3880, 3890, 3900, 3906, 3922, 3936, 3946, 4032}, }
_typeInfoList[7] = { parentId = 110, typeId = 3058, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3060}, }
_typeInfoList[8] = { parentId = 3058, typeId = 3060, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3062, 3112}, }
_typeInfoList[9] = { parentId = 3060, typeId = 3062, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3064, 3068, 3074, 3076, 3084, 3090, 3096, 3102, 3104, 3106, 3110}, }
_typeInfoList[10] = { parentId = 3062, typeId = 3064, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3066}, }
_typeInfoList[11] = { parentId = 3064, typeId = 3066, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[12] = { parentId = 3062, typeId = 3068, baseId = 3064, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3070, 3072}, }
_typeInfoList[13] = { parentId = 3068, typeId = 3070, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[14] = { parentId = 3068, typeId = 3072, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[15] = { parentId = 3062, typeId = 3074, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[16] = { parentId = 3062, typeId = 3076, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3078, 3082}, }
_typeInfoList[17] = { parentId = 3076, typeId = 3078, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[18] = { parentId = 1, typeId = 3080, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3076}, retTypeId = {}, children = {}, }
_typeInfoList[19] = { parentId = 3076, typeId = 3082, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3080}, children = {}, }
_typeInfoList[20] = { parentId = 3062, typeId = 3084, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3086, 3088}, }
_typeInfoList[21] = { parentId = 3084, typeId = 3086, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[22] = { parentId = 3084, typeId = 3088, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[23] = { parentId = 3062, typeId = 3090, baseId = 3084, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3092, 3094}, }
_typeInfoList[24] = { parentId = 3090, typeId = 3092, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[25] = { parentId = 3090, typeId = 3094, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[26] = { parentId = 3062, typeId = 3096, baseId = 3084, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3098, 3100, 3108}, }
_typeInfoList[27] = { parentId = 3096, typeId = 3098, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[28] = { parentId = 3096, typeId = 3100, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3096}, children = {}, }
_typeInfoList[29] = { parentId = 3062, typeId = 3102, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[30] = { parentId = 3062, typeId = 3104, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[31] = { parentId = 3062, typeId = 3106, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[32] = { parentId = 3096, typeId = 3108, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[33] = { parentId = 3062, typeId = 3110, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[34] = { parentId = 3060, typeId = 3112, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3114, 3118, 3126, 3128, 3130}, }
_typeInfoList[35] = { parentId = 3112, typeId = 3114, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3116}, }
_typeInfoList[36] = { parentId = 3114, typeId = 3116, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[37] = { parentId = 3112, typeId = 3118, baseId = 3114, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3120, 3122, 3124}, }
_typeInfoList[38] = { parentId = 3118, typeId = 3120, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[39] = { parentId = 3118, typeId = 3122, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[40] = { parentId = 3118, typeId = 3124, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[41] = { parentId = 3112, typeId = 3126, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[42] = { parentId = 3112, typeId = 3128, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[43] = { parentId = 3112, typeId = 3130, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[44] = { parentId = 110, typeId = 3132, baseId = 1, txt = 'TypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3142, 3144, 3146, 3148, 3150, 3152, 3154, 3156, 3158, 3160, 3162, 3164, 3166, 3170, 3174, 3176, 3178, 3180, 3182, 3184, 3186, 3188, 3190, 3192, 3194, 3198, 3200}, }
_typeInfoList[45] = { parentId = 110, typeId = 3136, baseId = 1, txt = 'isBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[46] = { parentId = 110, typeId = 3138, baseId = 1, txt = 'Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3230, 3236, 3238, 3240, 3244}, }
_typeInfoList[47] = { parentId = 110, typeId = 3140, baseId = 3138, txt = 'DeclClassNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3804, 3806, 3808, 3810, 3814, 3818, 3820, 3824}, }
_typeInfoList[48] = { parentId = 3132, typeId = 3142, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[49] = { parentId = 3132, typeId = 3144, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[50] = { parentId = 3132, typeId = 3146, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[51] = { parentId = 3132, typeId = 3148, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[52] = { parentId = 3132, typeId = 3150, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[53] = { parentId = 3132, typeId = 3152, baseId = 1, txt = 'cloneToPublic',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[54] = { parentId = 3132, typeId = 3154, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[55] = { parentId = 3132, typeId = 3156, baseId = 1, txt = 'createBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[56] = { parentId = 3132, typeId = 3158, baseId = 1, txt = 'createList',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[57] = { parentId = 3132, typeId = 3160, baseId = 1, txt = 'createArray',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[58] = { parentId = 3132, typeId = 3162, baseId = 1, txt = 'createMap',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[59] = { parentId = 3132, typeId = 3164, baseId = 1, txt = 'createClass',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[60] = { parentId = 3132, typeId = 3166, baseId = 1, txt = 'createFunc',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[61] = { parentId = 1, typeId = 3168, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[62] = { parentId = 3132, typeId = 3170, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3168}, children = {}, }
_typeInfoList[63] = { parentId = 1, typeId = 3172, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[64] = { parentId = 3132, typeId = 3174, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3172}, children = {}, }
_typeInfoList[65] = { parentId = 3132, typeId = 3176, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[66] = { parentId = 3132, typeId = 3178, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[67] = { parentId = 3132, typeId = 3180, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[68] = { parentId = 3132, typeId = 3182, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[69] = { parentId = 3132, typeId = 3184, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[70] = { parentId = 3132, typeId = 3186, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[71] = { parentId = 3132, typeId = 3188, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[72] = { parentId = 3132, typeId = 3190, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[73] = { parentId = 3132, typeId = 3192, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[74] = { parentId = 3132, typeId = 3194, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[75] = { parentId = 1, typeId = 3196, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[76] = { parentId = 3132, typeId = 3198, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3196}, children = {}, }
_typeInfoList[77] = { parentId = 3132, typeId = 3200, baseId = 1, txt = 'isSettableFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[78] = { parentId = 110, typeId = 3202, baseId = 1, txt = 'Scope',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3204, 3206, 3208, 3210, 3212, 3214, 3216, 3220, 3224}, }
_typeInfoList[79] = { parentId = 3202, typeId = 3204, baseId = 1, txt = 'add',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[80] = { parentId = 3202, typeId = 3206, baseId = 1, txt = 'addClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[81] = { parentId = 3202, typeId = 3208, baseId = 1, txt = 'getClassScope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3202}, children = {}, }
_typeInfoList[82] = { parentId = 3202, typeId = 3210, baseId = 1, txt = 'getTypeInfoChild',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[83] = { parentId = 3202, typeId = 3212, baseId = 1, txt = 'getTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[84] = { parentId = 3202, typeId = 3214, baseId = 1, txt = 'getTypeInfoMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[85] = { parentId = 3202, typeId = 3216, baseId = 1, txt = 'get_parent',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3202}, children = {}, }
_typeInfoList[86] = { parentId = 1, typeId = 3218, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 3132}, retTypeId = {}, children = {}, }
_typeInfoList[87] = { parentId = 3202, typeId = 3220, baseId = 1, txt = 'get_symbol2TypeInfoMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3218}, children = {}, }
_typeInfoList[88] = { parentId = 1, typeId = 3222, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 3202}, retTypeId = {}, children = {}, }
_typeInfoList[89] = { parentId = 3202, typeId = 3224, baseId = 1, txt = 'get_className2ScopeMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3222}, children = {}, }
_typeInfoList[90] = { parentId = 110, typeId = 3226, baseId = 1, txt = 'Filter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3276, 3284, 3294, 3310, 3326, 3348, 3360, 3378, 3394, 3406, 3418, 3436, 3452, 3468, 3486, 3496, 3504, 3516, 3526, 3540, 3550, 3562, 3574, 3586, 3596, 3606, 3618, 3630, 3640, 3652, 3674, 3722, 3732, 3742, 3752, 3764, 3784, 3796, 3802, 3828, 3838, 3846, 3858, 3870, 3882, 3892, 3908, 3924, 3938, 3948}, }
_typeInfoList[91] = { parentId = 110, typeId = 3228, baseId = 1, txt = 'NodePos',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[92] = { parentId = 3138, typeId = 3230, baseId = 1, txt = 'get_expType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[93] = { parentId = 1, typeId = 3232, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[94] = { parentId = 1, typeId = 3234, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[95] = { parentId = 3138, typeId = 3236, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3232, 3234}, children = {}, }
_typeInfoList[96] = { parentId = 3138, typeId = 3238, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[97] = { parentId = 3138, typeId = 3240, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[98] = { parentId = 1, typeId = 3242, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[99] = { parentId = 3138, typeId = 3244, baseId = 1, txt = 'get_expTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3242}, children = {}, }
_typeInfoList[100] = { parentId = 110, typeId = 3246, baseId = 1, txt = 'NamespaceInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[101] = { parentId = 110, typeId = 3248, baseId = 1, txt = 'MacroEval',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4030}, }
_typeInfoList[102] = { parentId = 110, typeId = 3250, baseId = 1, txt = 'DeclMacroInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3252, 3256, 3258, 3262}, }
_typeInfoList[103] = { parentId = 3250, typeId = 3252, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[104] = { parentId = 1, typeId = 3254, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3138}, retTypeId = {}, children = {}, }
_typeInfoList[105] = { parentId = 3250, typeId = 3256, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3254}, children = {}, }
_typeInfoList[106] = { parentId = 3250, typeId = 3258, baseId = 1, txt = 'get_ast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[107] = { parentId = 1, typeId = 3260, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3076}, retTypeId = {}, children = {}, }
_typeInfoList[108] = { parentId = 3250, typeId = 3262, baseId = 1, txt = 'get_tokenList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3260}, children = {}, }
_typeInfoList[109] = { parentId = 110, typeId = 3264, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3268, 4038}, }
_typeInfoList[110] = { parentId = 1, typeId = 3266, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[111] = { parentId = 3264, typeId = 3268, baseId = 1, txt = 'get_errMessList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3266}, children = {}, }
_typeInfoList[112] = { parentId = 110, typeId = 3272, baseId = 1, txt = 'getNodeKindName',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[113] = { parentId = 110, typeId = 3274, baseId = 3138, txt = 'NoneNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3278, 3280}, }
_typeInfoList[114] = { parentId = 3226, typeId = 3276, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[115] = { parentId = 3274, typeId = 3278, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[116] = { parentId = 3274, typeId = 3280, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[117] = { parentId = 110, typeId = 3282, baseId = 3138, txt = 'ImportNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3286, 3288, 3290}, }
_typeInfoList[118] = { parentId = 3226, typeId = 3284, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[119] = { parentId = 3282, typeId = 3286, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[120] = { parentId = 3282, typeId = 3288, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[121] = { parentId = 3282, typeId = 3290, baseId = 1, txt = 'get_modulePath',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[122] = { parentId = 110, typeId = 3292, baseId = 3138, txt = 'RootNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3296, 3298, 3302, 3306}, }
_typeInfoList[123] = { parentId = 3226, typeId = 3294, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[124] = { parentId = 3292, typeId = 3296, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[125] = { parentId = 3292, typeId = 3298, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[126] = { parentId = 1, typeId = 3300, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3138}, retTypeId = {}, children = {}, }
_typeInfoList[127] = { parentId = 3292, typeId = 3302, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3300}, children = {}, }
_typeInfoList[128] = { parentId = 1, typeId = 3304, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {12, 3246}, retTypeId = {}, children = {}, }
_typeInfoList[129] = { parentId = 3292, typeId = 3306, baseId = 1, txt = 'get_typeId2ClassMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3304}, children = {}, }
_typeInfoList[130] = { parentId = 110, typeId = 3308, baseId = 3138, txt = 'RefTypeNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3312, 3314, 3316, 3318, 3320, 3322}, }
_typeInfoList[131] = { parentId = 3226, typeId = 3310, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[132] = { parentId = 3308, typeId = 3312, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[133] = { parentId = 3308, typeId = 3314, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[134] = { parentId = 3308, typeId = 3316, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[135] = { parentId = 3308, typeId = 3318, baseId = 1, txt = 'get_refFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[136] = { parentId = 3308, typeId = 3320, baseId = 1, txt = 'get_mutFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[137] = { parentId = 3308, typeId = 3322, baseId = 1, txt = 'get_array',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[138] = { parentId = 110, typeId = 3324, baseId = 3138, txt = 'BlockNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3328, 3330, 3332, 3336}, }
_typeInfoList[139] = { parentId = 3226, typeId = 3326, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[140] = { parentId = 3324, typeId = 3328, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[141] = { parentId = 3324, typeId = 3330, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[142] = { parentId = 3324, typeId = 3332, baseId = 1, txt = 'get_blockKind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[143] = { parentId = 1, typeId = 3334, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3138}, retTypeId = {}, children = {}, }
_typeInfoList[144] = { parentId = 3324, typeId = 3336, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3334}, children = {}, }
_typeInfoList[145] = { parentId = 110, typeId = 3338, baseId = 1, txt = 'IfStmtInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3340, 3342, 3344}, }
_typeInfoList[146] = { parentId = 3338, typeId = 3340, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[147] = { parentId = 3338, typeId = 3342, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[148] = { parentId = 3338, typeId = 3344, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[149] = { parentId = 110, typeId = 3346, baseId = 3138, txt = 'IfNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3350, 3352, 3356}, }
_typeInfoList[150] = { parentId = 3226, typeId = 3348, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[151] = { parentId = 3346, typeId = 3350, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[152] = { parentId = 3346, typeId = 3352, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[153] = { parentId = 1, typeId = 3354, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3338}, retTypeId = {}, children = {}, }
_typeInfoList[154] = { parentId = 3346, typeId = 3356, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[155] = { parentId = 110, typeId = 3358, baseId = 3138, txt = 'ExpListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3362, 3364, 3368}, }
_typeInfoList[156] = { parentId = 3226, typeId = 3360, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[157] = { parentId = 3358, typeId = 3362, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[158] = { parentId = 3358, typeId = 3364, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[159] = { parentId = 1, typeId = 3366, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3138}, retTypeId = {}, children = {}, }
_typeInfoList[160] = { parentId = 3358, typeId = 3368, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[161] = { parentId = 110, typeId = 3370, baseId = 1, txt = 'CaseInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3372, 3374}, }
_typeInfoList[162] = { parentId = 3370, typeId = 3372, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3358}, children = {}, }
_typeInfoList[163] = { parentId = 3370, typeId = 3374, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[164] = { parentId = 110, typeId = 3376, baseId = 3138, txt = 'SwitchNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3380, 3382, 3384, 3388, 3390}, }
_typeInfoList[165] = { parentId = 3226, typeId = 3378, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[166] = { parentId = 3376, typeId = 3380, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[167] = { parentId = 3376, typeId = 3382, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[168] = { parentId = 3376, typeId = 3384, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[169] = { parentId = 1, typeId = 3386, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3370}, retTypeId = {}, children = {}, }
_typeInfoList[170] = { parentId = 3376, typeId = 3388, baseId = 1, txt = 'get_caseList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3386}, children = {}, }
_typeInfoList[171] = { parentId = 3376, typeId = 3390, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[172] = { parentId = 110, typeId = 3392, baseId = 3138, txt = 'WhileNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3396, 3398, 3400, 3402}, }
_typeInfoList[173] = { parentId = 3226, typeId = 3394, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[174] = { parentId = 3392, typeId = 3396, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[175] = { parentId = 3392, typeId = 3398, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[176] = { parentId = 3392, typeId = 3400, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[177] = { parentId = 3392, typeId = 3402, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[178] = { parentId = 110, typeId = 3404, baseId = 3138, txt = 'RepeatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3408, 3410, 3412, 3414}, }
_typeInfoList[179] = { parentId = 3226, typeId = 3406, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[180] = { parentId = 3404, typeId = 3408, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[181] = { parentId = 3404, typeId = 3410, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[182] = { parentId = 3404, typeId = 3412, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[183] = { parentId = 3404, typeId = 3414, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[184] = { parentId = 110, typeId = 3416, baseId = 3138, txt = 'ForNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3420, 3422, 3424, 3426, 3428, 3430, 3432}, }
_typeInfoList[185] = { parentId = 3226, typeId = 3418, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[186] = { parentId = 3416, typeId = 3420, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[187] = { parentId = 3416, typeId = 3422, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[188] = { parentId = 3416, typeId = 3424, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[189] = { parentId = 3416, typeId = 3426, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[190] = { parentId = 3416, typeId = 3428, baseId = 1, txt = 'get_init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[191] = { parentId = 3416, typeId = 3430, baseId = 1, txt = 'get_to',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[192] = { parentId = 3416, typeId = 3432, baseId = 1, txt = 'get_delta',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[193] = { parentId = 110, typeId = 3434, baseId = 3138, txt = 'ApplyNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3438, 3440, 3444, 3446, 3448}, }
_typeInfoList[194] = { parentId = 3226, typeId = 3436, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[195] = { parentId = 3434, typeId = 3438, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[196] = { parentId = 3434, typeId = 3440, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[197] = { parentId = 1, typeId = 3442, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3076}, retTypeId = {}, children = {}, }
_typeInfoList[198] = { parentId = 3434, typeId = 3444, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3442}, children = {}, }
_typeInfoList[199] = { parentId = 3434, typeId = 3446, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[200] = { parentId = 3434, typeId = 3448, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[201] = { parentId = 110, typeId = 3450, baseId = 3138, txt = 'ForeachNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3454, 3456, 3458, 3460, 3462, 3464}, }
_typeInfoList[202] = { parentId = 3226, typeId = 3452, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[203] = { parentId = 3450, typeId = 3454, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[204] = { parentId = 3450, typeId = 3456, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[205] = { parentId = 3450, typeId = 3458, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[206] = { parentId = 3450, typeId = 3460, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[207] = { parentId = 3450, typeId = 3462, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[208] = { parentId = 3450, typeId = 3464, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[209] = { parentId = 110, typeId = 3466, baseId = 3138, txt = 'ForsortNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3470, 3472, 3474, 3476, 3478, 3480, 3482}, }
_typeInfoList[210] = { parentId = 3226, typeId = 3468, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[211] = { parentId = 3466, typeId = 3470, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[212] = { parentId = 3466, typeId = 3472, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[213] = { parentId = 3466, typeId = 3474, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[214] = { parentId = 3466, typeId = 3476, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[215] = { parentId = 3466, typeId = 3478, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[216] = { parentId = 3466, typeId = 3480, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[217] = { parentId = 3466, typeId = 3482, baseId = 1, txt = 'get_sort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[218] = { parentId = 110, typeId = 3484, baseId = 3138, txt = 'ReturnNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3488, 3490, 3492}, }
_typeInfoList[219] = { parentId = 3226, typeId = 3486, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[220] = { parentId = 3484, typeId = 3488, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[221] = { parentId = 3484, typeId = 3490, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[222] = { parentId = 3484, typeId = 3492, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3358}, children = {}, }
_typeInfoList[223] = { parentId = 110, typeId = 3494, baseId = 3138, txt = 'BreakNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3498, 3500}, }
_typeInfoList[224] = { parentId = 3226, typeId = 3496, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[225] = { parentId = 3494, typeId = 3498, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[226] = { parentId = 3494, typeId = 3500, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[227] = { parentId = 110, typeId = 3502, baseId = 3138, txt = 'ExpNewNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3506, 3508, 3510, 3512}, }
_typeInfoList[228] = { parentId = 3226, typeId = 3504, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[229] = { parentId = 3502, typeId = 3506, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[230] = { parentId = 3502, typeId = 3508, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[231] = { parentId = 3502, typeId = 3510, baseId = 1, txt = 'get_symbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[232] = { parentId = 3502, typeId = 3512, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3358}, children = {}, }
_typeInfoList[233] = { parentId = 110, typeId = 3514, baseId = 3138, txt = 'ExpRefNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3518, 3520, 3522}, }
_typeInfoList[234] = { parentId = 3226, typeId = 3516, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[235] = { parentId = 3514, typeId = 3518, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[236] = { parentId = 3514, typeId = 3520, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[237] = { parentId = 3514, typeId = 3522, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[238] = { parentId = 110, typeId = 3524, baseId = 3138, txt = 'ExpOp2Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3528, 3530, 3532, 3534, 3536}, }
_typeInfoList[239] = { parentId = 3226, typeId = 3526, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[240] = { parentId = 3524, typeId = 3528, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[241] = { parentId = 3524, typeId = 3530, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[242] = { parentId = 3524, typeId = 3532, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[243] = { parentId = 3524, typeId = 3534, baseId = 1, txt = 'get_exp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[244] = { parentId = 3524, typeId = 3536, baseId = 1, txt = 'get_exp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[245] = { parentId = 110, typeId = 3538, baseId = 3138, txt = 'ExpCastNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3542, 3544, 3546}, }
_typeInfoList[246] = { parentId = 3226, typeId = 3540, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[247] = { parentId = 3538, typeId = 3542, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[248] = { parentId = 3538, typeId = 3544, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[249] = { parentId = 3538, typeId = 3546, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[250] = { parentId = 110, typeId = 3548, baseId = 3138, txt = 'ExpOp1Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3552, 3554, 3556, 3558}, }
_typeInfoList[251] = { parentId = 3226, typeId = 3550, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[252] = { parentId = 3548, typeId = 3552, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[253] = { parentId = 3548, typeId = 3554, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[254] = { parentId = 3548, typeId = 3556, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[255] = { parentId = 3548, typeId = 3558, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[256] = { parentId = 110, typeId = 3560, baseId = 3138, txt = 'ExpRefItemNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3564, 3566, 3568, 3570}, }
_typeInfoList[257] = { parentId = 3226, typeId = 3562, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[258] = { parentId = 3560, typeId = 3564, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[259] = { parentId = 3560, typeId = 3566, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[260] = { parentId = 3560, typeId = 3568, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[261] = { parentId = 3560, typeId = 3570, baseId = 1, txt = 'get_index',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[262] = { parentId = 110, typeId = 3572, baseId = 3138, txt = 'ExpCallNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3576, 3578, 3580, 3582}, }
_typeInfoList[263] = { parentId = 3226, typeId = 3574, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[264] = { parentId = 3572, typeId = 3576, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[265] = { parentId = 3572, typeId = 3578, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[266] = { parentId = 3572, typeId = 3580, baseId = 1, txt = 'get_func',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[267] = { parentId = 3572, typeId = 3582, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3358}, children = {}, }
_typeInfoList[268] = { parentId = 110, typeId = 3584, baseId = 3138, txt = 'ExpDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3588, 3590, 3592}, }
_typeInfoList[269] = { parentId = 3226, typeId = 3586, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[270] = { parentId = 3584, typeId = 3588, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[271] = { parentId = 3584, typeId = 3590, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[272] = { parentId = 3584, typeId = 3592, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[273] = { parentId = 110, typeId = 3594, baseId = 3138, txt = 'ExpParenNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3598, 3600, 3602}, }
_typeInfoList[274] = { parentId = 3226, typeId = 3596, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[275] = { parentId = 3594, typeId = 3598, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[276] = { parentId = 3594, typeId = 3600, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[277] = { parentId = 3594, typeId = 3602, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[278] = { parentId = 110, typeId = 3604, baseId = 3138, txt = 'ExpMacroExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3608, 3610, 3614}, }
_typeInfoList[279] = { parentId = 3226, typeId = 3606, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[280] = { parentId = 3604, typeId = 3608, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[281] = { parentId = 3604, typeId = 3610, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[282] = { parentId = 1, typeId = 3612, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3138}, retTypeId = {}, children = {}, }
_typeInfoList[283] = { parentId = 3604, typeId = 3614, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3612}, children = {}, }
_typeInfoList[284] = { parentId = 110, typeId = 3616, baseId = 3138, txt = 'ExpMacroStatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3620, 3622, 3626, 4026}, }
_typeInfoList[285] = { parentId = 3226, typeId = 3618, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[286] = { parentId = 3616, typeId = 3620, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[287] = { parentId = 3616, typeId = 3622, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[288] = { parentId = 1, typeId = 3624, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3138}, retTypeId = {}, children = {}, }
_typeInfoList[289] = { parentId = 3616, typeId = 3626, baseId = 1, txt = 'get_expStrList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3624}, children = {}, }
_typeInfoList[290] = { parentId = 110, typeId = 3628, baseId = 3138, txt = 'StmtExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3632, 3634, 3636}, }
_typeInfoList[291] = { parentId = 3226, typeId = 3630, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[292] = { parentId = 3628, typeId = 3632, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[293] = { parentId = 3628, typeId = 3634, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[294] = { parentId = 3628, typeId = 3636, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[295] = { parentId = 110, typeId = 3638, baseId = 3138, txt = 'RefFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3642, 3644, 3646, 3648, 4020}, }
_typeInfoList[296] = { parentId = 3226, typeId = 3640, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[297] = { parentId = 3638, typeId = 3642, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[298] = { parentId = 3638, typeId = 3644, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[299] = { parentId = 3638, typeId = 3646, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[300] = { parentId = 3638, typeId = 3648, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[301] = { parentId = 110, typeId = 3650, baseId = 3138, txt = 'GetFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3654, 3656, 3658, 3660, 3662}, }
_typeInfoList[302] = { parentId = 3226, typeId = 3652, baseId = 1, txt = 'processGetField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[303] = { parentId = 3650, typeId = 3654, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[304] = { parentId = 3650, typeId = 3656, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[305] = { parentId = 3650, typeId = 3658, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[306] = { parentId = 3650, typeId = 3660, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[307] = { parentId = 3650, typeId = 3662, baseId = 1, txt = 'get_getterTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[308] = { parentId = 110, typeId = 3664, baseId = 1, txt = 'VarInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3666, 3668, 3670}, }
_typeInfoList[309] = { parentId = 3664, typeId = 3666, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[310] = { parentId = 3664, typeId = 3668, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[311] = { parentId = 3664, typeId = 3670, baseId = 1, txt = 'get_actualType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[312] = { parentId = 110, typeId = 3672, baseId = 3138, txt = 'DeclVarNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3676, 3678, 3680, 3682, 3686, 3688, 3692, 3694}, }
_typeInfoList[313] = { parentId = 3226, typeId = 3674, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[314] = { parentId = 3672, typeId = 3676, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[315] = { parentId = 3672, typeId = 3678, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[316] = { parentId = 3672, typeId = 3680, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[317] = { parentId = 3672, typeId = 3682, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[318] = { parentId = 1, typeId = 3684, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3664}, retTypeId = {}, children = {}, }
_typeInfoList[319] = { parentId = 3672, typeId = 3686, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3684}, children = {}, }
_typeInfoList[320] = { parentId = 3672, typeId = 3688, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3358}, children = {}, }
_typeInfoList[321] = { parentId = 1, typeId = 3690, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[322] = { parentId = 3672, typeId = 3692, baseId = 1, txt = 'get_typeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3690}, children = {}, }
_typeInfoList[323] = { parentId = 3672, typeId = 3694, baseId = 1, txt = 'get_unwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3324}, children = {}, }
_typeInfoList[324] = { parentId = 110, typeId = 3696, baseId = 1, txt = 'DeclFuncInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3698, 3700, 3704, 3706, 3708, 3710, 3714, 3718}, }
_typeInfoList[325] = { parentId = 3696, typeId = 3698, baseId = 1, txt = 'get_className',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[326] = { parentId = 3696, typeId = 3700, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[327] = { parentId = 1, typeId = 3702, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3138}, retTypeId = {}, children = {}, }
_typeInfoList[328] = { parentId = 3696, typeId = 3704, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3702}, children = {}, }
_typeInfoList[329] = { parentId = 3696, typeId = 3706, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[330] = { parentId = 3696, typeId = 3708, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[331] = { parentId = 3696, typeId = 3710, baseId = 1, txt = 'get_body',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[332] = { parentId = 1, typeId = 3712, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[333] = { parentId = 3696, typeId = 3714, baseId = 1, txt = 'get_retTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3712}, children = {}, }
_typeInfoList[334] = { parentId = 1, typeId = 3716, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[335] = { parentId = 3696, typeId = 3718, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3716}, children = {}, }
_typeInfoList[336] = { parentId = 110, typeId = 3720, baseId = 3138, txt = 'DeclFuncNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3724, 3726, 3728}, }
_typeInfoList[337] = { parentId = 3226, typeId = 3722, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[338] = { parentId = 3720, typeId = 3724, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[339] = { parentId = 3720, typeId = 3726, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[340] = { parentId = 3720, typeId = 3728, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3696}, children = {}, }
_typeInfoList[341] = { parentId = 110, typeId = 3730, baseId = 3138, txt = 'DeclMethodNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3734, 3736, 3738}, }
_typeInfoList[342] = { parentId = 3226, typeId = 3732, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[343] = { parentId = 3730, typeId = 3734, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[344] = { parentId = 3730, typeId = 3736, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[345] = { parentId = 3730, typeId = 3738, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3696}, children = {}, }
_typeInfoList[346] = { parentId = 110, typeId = 3740, baseId = 3138, txt = 'DeclConstrNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3744, 3746, 3748}, }
_typeInfoList[347] = { parentId = 3226, typeId = 3742, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[348] = { parentId = 3740, typeId = 3744, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[349] = { parentId = 3740, typeId = 3746, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[350] = { parentId = 3740, typeId = 3748, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3696}, children = {}, }
_typeInfoList[351] = { parentId = 110, typeId = 3750, baseId = 3138, txt = 'ExpCallSuperNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3754, 3756, 3758, 3760}, }
_typeInfoList[352] = { parentId = 3226, typeId = 3752, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[353] = { parentId = 3750, typeId = 3754, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[354] = { parentId = 3750, typeId = 3756, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[355] = { parentId = 3750, typeId = 3758, baseId = 1, txt = 'get_superType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[356] = { parentId = 3750, typeId = 3760, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3358}, children = {}, }
_typeInfoList[357] = { parentId = 110, typeId = 3762, baseId = 3138, txt = 'DeclMemberNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3766, 3768, 3770, 3772, 3774, 3776, 3778, 3780}, }
_typeInfoList[358] = { parentId = 3226, typeId = 3764, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[359] = { parentId = 3762, typeId = 3766, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[360] = { parentId = 3762, typeId = 3768, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[361] = { parentId = 3762, typeId = 3770, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[362] = { parentId = 3762, typeId = 3772, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[363] = { parentId = 3762, typeId = 3774, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[364] = { parentId = 3762, typeId = 3776, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[365] = { parentId = 3762, typeId = 3778, baseId = 1, txt = 'get_getterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[366] = { parentId = 3762, typeId = 3780, baseId = 1, txt = 'get_setterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[367] = { parentId = 110, typeId = 3782, baseId = 3138, txt = 'DeclArgNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3786, 3788, 3790, 3792}, }
_typeInfoList[368] = { parentId = 3226, typeId = 3784, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[369] = { parentId = 3782, typeId = 3786, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[370] = { parentId = 3782, typeId = 3788, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[371] = { parentId = 3782, typeId = 3790, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[372] = { parentId = 3782, typeId = 3792, baseId = 1, txt = 'get_argType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[373] = { parentId = 110, typeId = 3794, baseId = 3138, txt = 'DeclArgDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3798, 3800}, }
_typeInfoList[374] = { parentId = 3226, typeId = 3796, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[375] = { parentId = 3794, typeId = 3798, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[376] = { parentId = 3794, typeId = 3800, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[377] = { parentId = 3226, typeId = 3802, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[378] = { parentId = 3140, typeId = 3804, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[379] = { parentId = 3140, typeId = 3806, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[380] = { parentId = 3140, typeId = 3808, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[381] = { parentId = 3140, typeId = 3810, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[382] = { parentId = 1, typeId = 3812, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3138}, retTypeId = {}, children = {}, }
_typeInfoList[383] = { parentId = 3140, typeId = 3814, baseId = 1, txt = 'get_fieldList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3812}, children = {}, }
_typeInfoList[384] = { parentId = 1, typeId = 3816, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3762}, retTypeId = {}, children = {}, }
_typeInfoList[385] = { parentId = 3140, typeId = 3818, baseId = 1, txt = 'get_memberList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3816}, children = {}, }
_typeInfoList[386] = { parentId = 3140, typeId = 3820, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3202}, children = {}, }
_typeInfoList[387] = { parentId = 1, typeId = 3822, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[388] = { parentId = 3140, typeId = 3824, baseId = 1, txt = 'get_outerMethodSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3822}, children = {}, }
_typeInfoList[389] = { parentId = 110, typeId = 3826, baseId = 3138, txt = 'DeclMacroNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3830, 3832, 3834}, }
_typeInfoList[390] = { parentId = 3226, typeId = 3828, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[391] = { parentId = 3826, typeId = 3830, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[392] = { parentId = 3826, typeId = 3832, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[393] = { parentId = 3826, typeId = 3834, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3250}, children = {}, }
_typeInfoList[394] = { parentId = 110, typeId = 3836, baseId = 3138, txt = 'LiteralNilNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3840, 3842, 3960}, }
_typeInfoList[395] = { parentId = 3226, typeId = 3838, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[396] = { parentId = 3836, typeId = 3840, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[397] = { parentId = 3836, typeId = 3842, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[398] = { parentId = 110, typeId = 3844, baseId = 3138, txt = 'LiteralCharNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3848, 3850, 3852, 3854, 3966}, }
_typeInfoList[399] = { parentId = 3226, typeId = 3846, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[400] = { parentId = 3844, typeId = 3848, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[401] = { parentId = 3844, typeId = 3850, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[402] = { parentId = 3844, typeId = 3852, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[403] = { parentId = 3844, typeId = 3854, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[404] = { parentId = 110, typeId = 3856, baseId = 3138, txt = 'LiteralIntNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3860, 3862, 3864, 3866, 3972}, }
_typeInfoList[405] = { parentId = 3226, typeId = 3858, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[406] = { parentId = 3856, typeId = 3860, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[407] = { parentId = 3856, typeId = 3862, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[408] = { parentId = 3856, typeId = 3864, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[409] = { parentId = 3856, typeId = 3866, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[410] = { parentId = 110, typeId = 3868, baseId = 3138, txt = 'LiteralRealNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3872, 3874, 3876, 3878, 3978}, }
_typeInfoList[411] = { parentId = 3226, typeId = 3870, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[412] = { parentId = 3868, typeId = 3872, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[413] = { parentId = 3868, typeId = 3874, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[414] = { parentId = 3868, typeId = 3876, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[415] = { parentId = 3868, typeId = 3878, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[416] = { parentId = 110, typeId = 3880, baseId = 3138, txt = 'LiteralArrayNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3884, 3886, 3888, 3984}, }
_typeInfoList[417] = { parentId = 3226, typeId = 3882, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[418] = { parentId = 3880, typeId = 3884, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[419] = { parentId = 3880, typeId = 3886, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[420] = { parentId = 3880, typeId = 3888, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3358}, children = {}, }
_typeInfoList[421] = { parentId = 110, typeId = 3890, baseId = 3138, txt = 'LiteralListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3894, 3896, 3898, 3990}, }
_typeInfoList[422] = { parentId = 3226, typeId = 3892, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[423] = { parentId = 3890, typeId = 3894, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[424] = { parentId = 3890, typeId = 3896, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[425] = { parentId = 3890, typeId = 3898, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3358}, children = {}, }
_typeInfoList[426] = { parentId = 110, typeId = 3900, baseId = 1, txt = 'PairItem',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3902, 3904}, }
_typeInfoList[427] = { parentId = 3900, typeId = 3902, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[428] = { parentId = 3900, typeId = 3904, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[429] = { parentId = 110, typeId = 3906, baseId = 3138, txt = 'LiteralMapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3910, 3912, 3916, 3920, 3996}, }
_typeInfoList[430] = { parentId = 3226, typeId = 3908, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[431] = { parentId = 3906, typeId = 3910, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[432] = { parentId = 3906, typeId = 3912, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[433] = { parentId = 1, typeId = 3914, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {3138, 3138}, retTypeId = {}, children = {}, }
_typeInfoList[434] = { parentId = 3906, typeId = 3916, baseId = 1, txt = 'get_map',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3914}, children = {}, }
_typeInfoList[435] = { parentId = 1, typeId = 3918, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3900}, retTypeId = {}, children = {}, }
_typeInfoList[436] = { parentId = 3906, typeId = 3920, baseId = 1, txt = 'get_pairList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3918}, children = {}, }
_typeInfoList[437] = { parentId = 110, typeId = 3922, baseId = 3138, txt = 'LiteralStringNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3926, 3928, 3930, 3934, 4002}, }
_typeInfoList[438] = { parentId = 3226, typeId = 3924, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[439] = { parentId = 3922, typeId = 3926, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[440] = { parentId = 3922, typeId = 3928, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[441] = { parentId = 3922, typeId = 3930, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[442] = { parentId = 1, typeId = 3932, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3138}, retTypeId = {}, children = {}, }
_typeInfoList[443] = { parentId = 3922, typeId = 3934, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3932}, children = {}, }
_typeInfoList[444] = { parentId = 110, typeId = 3936, baseId = 3138, txt = 'LiteralBoolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3940, 3942, 3944, 4008}, }
_typeInfoList[445] = { parentId = 3226, typeId = 3938, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[446] = { parentId = 3936, typeId = 3940, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[447] = { parentId = 3936, typeId = 3942, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[448] = { parentId = 3936, typeId = 3944, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[449] = { parentId = 110, typeId = 3946, baseId = 3138, txt = 'LiteralSymbolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3950, 3952, 3954, 4014}, }
_typeInfoList[450] = { parentId = 3226, typeId = 3948, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[451] = { parentId = 3946, typeId = 3950, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[452] = { parentId = 3946, typeId = 3952, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[453] = { parentId = 3946, typeId = 3954, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3076}, children = {}, }
_typeInfoList[454] = { parentId = 1, typeId = 3956, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[455] = { parentId = 1, typeId = 3958, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[456] = { parentId = 3836, typeId = 3960, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3956, 3958}, children = {}, }
_typeInfoList[457] = { parentId = 1, typeId = 3962, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[458] = { parentId = 1, typeId = 3964, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[459] = { parentId = 3844, typeId = 3966, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3962, 3964}, children = {}, }
_typeInfoList[460] = { parentId = 1, typeId = 3968, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[461] = { parentId = 1, typeId = 3970, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[462] = { parentId = 3856, typeId = 3972, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3968, 3970}, children = {}, }
_typeInfoList[463] = { parentId = 1, typeId = 3974, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[464] = { parentId = 1, typeId = 3976, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[465] = { parentId = 3868, typeId = 3978, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3974, 3976}, children = {}, }
_typeInfoList[466] = { parentId = 1, typeId = 3980, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[467] = { parentId = 1, typeId = 3982, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[468] = { parentId = 3880, typeId = 3984, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3980, 3982}, children = {}, }
_typeInfoList[469] = { parentId = 1, typeId = 3986, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[470] = { parentId = 1, typeId = 3988, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[471] = { parentId = 3890, typeId = 3990, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3986, 3988}, children = {}, }
_typeInfoList[472] = { parentId = 1, typeId = 3992, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[473] = { parentId = 1, typeId = 3994, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[474] = { parentId = 3906, typeId = 3996, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3992, 3994}, children = {}, }
_typeInfoList[475] = { parentId = 1, typeId = 3998, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[476] = { parentId = 1, typeId = 4000, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[477] = { parentId = 3922, typeId = 4002, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3998, 4000}, children = {}, }
_typeInfoList[478] = { parentId = 1, typeId = 4004, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[479] = { parentId = 1, typeId = 4006, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[480] = { parentId = 3936, typeId = 4008, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4004, 4006}, children = {}, }
_typeInfoList[481] = { parentId = 1, typeId = 4010, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[482] = { parentId = 1, typeId = 4012, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[483] = { parentId = 3946, typeId = 4014, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4010, 4012}, children = {}, }
_typeInfoList[484] = { parentId = 1, typeId = 4016, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[485] = { parentId = 1, typeId = 4018, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[486] = { parentId = 3638, typeId = 4020, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4016, 4018}, children = {}, }
_typeInfoList[487] = { parentId = 1, typeId = 4022, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[488] = { parentId = 1, typeId = 4024, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3132}, retTypeId = {}, children = {}, }
_typeInfoList[489] = { parentId = 3616, typeId = 4026, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4022, 4024}, children = {}, }
_typeInfoList[490] = { parentId = 1, typeId = 4028, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, retTypeId = {}, children = {}, }
_typeInfoList[491] = { parentId = 3248, typeId = 4030, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4028}, children = {}, }
_typeInfoList[492] = { parentId = 110, typeId = 4032, baseId = 1, txt = 'ASTInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4034, 4036}, }
_typeInfoList[493] = { parentId = 4032, typeId = 4034, baseId = 1, txt = 'get_node',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3138}, children = {}, }
_typeInfoList[494] = { parentId = 4032, typeId = 4036, baseId = 1, txt = 'get_moduleTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3132}, children = {}, }
_typeInfoList[495] = { parentId = 3264, typeId = 4038, baseId = 1, txt = 'createAST',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4032}, children = {}, }
_typeInfoList[496] = { parentId = 108, typeId = 4040, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4042, 4046, 4052, 4054, 4062, 4068, 4074, 4082, 4084, 4086, 4090}, }
_typeInfoList[497] = { parentId = 4040, typeId = 4042, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4044}, }
_typeInfoList[498] = { parentId = 4042, typeId = 4044, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[499] = { parentId = 4040, typeId = 4046, baseId = 4042, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4048, 4050}, }
_typeInfoList[500] = { parentId = 4046, typeId = 4048, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[501] = { parentId = 4046, typeId = 4050, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[502] = { parentId = 4040, typeId = 4052, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[503] = { parentId = 4040, typeId = 4054, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4056, 4060}, }
_typeInfoList[504] = { parentId = 4054, typeId = 4056, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[505] = { parentId = 1, typeId = 4058, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4054}, retTypeId = {}, children = {}, }
_typeInfoList[506] = { parentId = 4054, typeId = 4060, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4058}, children = {}, }
_typeInfoList[507] = { parentId = 4040, typeId = 4062, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4064, 4066}, }
_typeInfoList[508] = { parentId = 4062, typeId = 4064, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4054}, children = {}, }
_typeInfoList[509] = { parentId = 4062, typeId = 4066, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[510] = { parentId = 4040, typeId = 4068, baseId = 4062, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4070, 4072}, }
_typeInfoList[511] = { parentId = 4068, typeId = 4070, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4054}, children = {}, }
_typeInfoList[512] = { parentId = 4068, typeId = 4072, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[513] = { parentId = 4040, typeId = 4074, baseId = 4062, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4076, 4078, 4088}, }
_typeInfoList[514] = { parentId = 4074, typeId = 4076, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[515] = { parentId = 4074, typeId = 4078, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {4074}, children = {}, }
_typeInfoList[516] = { parentId = 4040, typeId = 4082, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[517] = { parentId = 4040, typeId = 4084, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[518] = { parentId = 4040, typeId = 4086, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[519] = { parentId = 4074, typeId = 4088, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4054}, children = {}, }
_typeInfoList[520] = { parentId = 4040, typeId = 4090, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[521] = { parentId = 104, typeId = 4092, baseId = 3226, txt = 'dumpFilter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4100, 4102, 4104, 4106, 4108, 4110, 4112, 4114, 4116, 4118, 4120, 4122, 4124, 4126, 4134, 4136, 4138, 4140, 4142, 4144, 4148, 4152, 4154, 4156, 4158, 4162, 4164, 4166, 4168, 4172, 4174, 4176, 4178, 4180, 4182, 4184, 4186, 4188, 4190, 4192, 4194, 4198, 4200, 4202, 4204, 4206, 4208, 4210, 4212, 4214}, }
_typeInfoList[522] = { parentId = 4092, typeId = 4100, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[523] = { parentId = 4092, typeId = 4102, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[524] = { parentId = 4092, typeId = 4104, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[525] = { parentId = 4092, typeId = 4106, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[526] = { parentId = 4092, typeId = 4108, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[527] = { parentId = 4092, typeId = 4110, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[528] = { parentId = 4092, typeId = 4112, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[529] = { parentId = 4092, typeId = 4114, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[530] = { parentId = 4092, typeId = 4116, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[531] = { parentId = 4092, typeId = 4118, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[532] = { parentId = 4092, typeId = 4120, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[533] = { parentId = 4092, typeId = 4122, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[534] = { parentId = 4092, typeId = 4124, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[535] = { parentId = 4092, typeId = 4126, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[536] = { parentId = 4092, typeId = 4134, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[537] = { parentId = 4092, typeId = 4136, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[538] = { parentId = 4092, typeId = 4138, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[539] = { parentId = 4092, typeId = 4140, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[540] = { parentId = 4092, typeId = 4142, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[541] = { parentId = 4092, typeId = 4144, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[542] = { parentId = 4092, typeId = 4148, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[543] = { parentId = 4092, typeId = 4152, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[544] = { parentId = 4092, typeId = 4154, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[545] = { parentId = 4092, typeId = 4156, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[546] = { parentId = 4092, typeId = 4158, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[547] = { parentId = 4092, typeId = 4162, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[548] = { parentId = 4092, typeId = 4164, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[549] = { parentId = 4092, typeId = 4166, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[550] = { parentId = 4092, typeId = 4168, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[551] = { parentId = 4092, typeId = 4172, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[552] = { parentId = 4092, typeId = 4174, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[553] = { parentId = 4092, typeId = 4176, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[554] = { parentId = 4092, typeId = 4178, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[555] = { parentId = 4092, typeId = 4180, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[556] = { parentId = 4092, typeId = 4182, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[557] = { parentId = 4092, typeId = 4184, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[558] = { parentId = 4092, typeId = 4186, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[559] = { parentId = 4092, typeId = 4188, baseId = 1, txt = 'processGetField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[560] = { parentId = 4092, typeId = 4190, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[561] = { parentId = 4092, typeId = 4192, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[562] = { parentId = 4092, typeId = 4194, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[563] = { parentId = 4092, typeId = 4198, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[564] = { parentId = 4092, typeId = 4200, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[565] = { parentId = 4092, typeId = 4202, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[566] = { parentId = 4092, typeId = 4204, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[567] = { parentId = 4092, typeId = 4206, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[568] = { parentId = 4092, typeId = 4208, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[569] = { parentId = 4092, typeId = 4210, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[570] = { parentId = 4092, typeId = 4212, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[571] = { parentId = 4092, typeId = 4214, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
----- meta -----
return moduleObj
