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
  dump( prefix, depth, node, getTxt( node:get_name(  ) ) )
  for index, field in pairs( node:get_fieldList(  ) ) do
    filter( field, self, prefix .. "  ", depth + 1 )
  end
end

-- none

function dumpFilter:processDeclMember( node, prefix, depth )
  dump( prefix, depth, node, getTxt( node:get_name(  ) ) )
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
    varName = string.format( "%s %s", varName, getTxt( var:get_name(  ) ))
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
  dump( prefix, depth, node, getTxt( node:get_name(  ) ) )
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
  dump( prefix, depth, node, name and getTxt( name ) or "<anonymous>" )
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
  dump( prefix, depth, node, getTxt( node:get_val(  ) ) )
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
  local index = node:get_key(  ) and getTxt( node:get_key(  ) ) or ""
  dump( prefix, depth, node, getTxt( node:get_val(  ) ) .. " " .. index )
  filter( node:get_exp(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_block(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processForsort( node, prefix, depth )
  local index = node:get_key(  ) and getTxt( node:get_key(  ) ) or ""
  dump( prefix, depth, node, getTxt( node:get_val(  ) ) .. " " .. index )
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
  dump( prefix, depth, node, getTxt( node:get_op(  ) ) )
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
  dump( prefix, depth, node, getTxt( node:get_op(  ) ) )
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
  dump( prefix, depth, node, getTxt( node:get_token(  ) ) )
end

-- none

function dumpFilter:processExpRefItem( node, prefix, depth )
  dump( prefix, depth, node, "seq[exp] " .. node:get_expType(  ):getTxt(  ) )
  filter( node:get_val(  ), self, prefix .. "  ", depth + 1 )
  filter( node:get_index(  ), self, prefix .. "  ", depth + 1 )
end

-- none

function dumpFilter:processRefField( node, prefix, depth )
  dump( prefix, depth, node, getTxt( node:get_field(  ) ) )
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
  dump( prefix, depth, node, string.format( "%s(%s)", node:get_num(  ), getTxt( node:get_token(  ) ) ) )
end

-- none

function dumpFilter:processLiteralInt( node, prefix, depth )
  dump( prefix, depth, node, string.format( "%s(%s)", node:get_num(  ), getTxt( node:get_token(  ) ) ) )
end

-- none

function dumpFilter:processLiteralReal( node, prefix, depth )
  dump( prefix, depth, node, string.format( "%s(%s)", node:get_num(  ), getTxt( node:get_token(  ) ) ) )
end

-- none

function dumpFilter:processLiteralString( node, prefix, depth )
  dump( prefix, depth, node, getTxt( node:get_token(  ) ) )
end

-- none

function dumpFilter:processLiteralBool( node, prefix, depth )
  dump( prefix, depth, node, getTxt( node:get_token(  ) ) == "true" )
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
  dump( prefix, depth, node, getTxt( node:get_token(  ) ) )
end

-- none

----- meta -----
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfodumpFilter = {}
  _className2InfoMap.dumpFilter = _classInfodumpFilter
  end
do
  local _classInfoASTInfo = {}
  _className2InfoMap.ASTInfo = _classInfoASTInfo
  end
do
  local _classInfoApplyNode = {}
  _className2InfoMap.ApplyNode = _classInfoApplyNode
  end
do
  local _classInfoBlockNode = {}
  _className2InfoMap.BlockNode = _classInfoBlockNode
  end
do
  local _classInfoBreakNode = {}
  _className2InfoMap.BreakNode = _classInfoBreakNode
  end
do
  local _classInfoCaseInfo = {}
  _className2InfoMap.CaseInfo = _classInfoCaseInfo
  end
do
  local _classInfoDeclArgDDDNode = {}
  _className2InfoMap.DeclArgDDDNode = _classInfoDeclArgDDDNode
  end
do
  local _classInfoDeclArgNode = {}
  _className2InfoMap.DeclArgNode = _classInfoDeclArgNode
  end
do
  local _classInfoDeclClassNode = {}
  _className2InfoMap.DeclClassNode = _classInfoDeclClassNode
  end
do
  local _classInfoDeclConstrNode = {}
  _className2InfoMap.DeclConstrNode = _classInfoDeclConstrNode
  end
do
  local _classInfoDeclFuncInfo = {}
  _className2InfoMap.DeclFuncInfo = _classInfoDeclFuncInfo
  end
do
  local _classInfoDeclFuncNode = {}
  _className2InfoMap.DeclFuncNode = _classInfoDeclFuncNode
  end
do
  local _classInfoDeclMacroInfo = {}
  _className2InfoMap.DeclMacroInfo = _classInfoDeclMacroInfo
  end
do
  local _classInfoDeclMacroNode = {}
  _className2InfoMap.DeclMacroNode = _classInfoDeclMacroNode
  end
do
  local _classInfoDeclMemberNode = {}
  _className2InfoMap.DeclMemberNode = _classInfoDeclMemberNode
  end
do
  local _classInfoDeclMethodNode = {}
  _className2InfoMap.DeclMethodNode = _classInfoDeclMethodNode
  end
do
  local _classInfoDeclVarNode = {}
  _className2InfoMap.DeclVarNode = _classInfoDeclVarNode
  end
do
  local _classInfoExpCallNode = {}
  _className2InfoMap.ExpCallNode = _classInfoExpCallNode
  end
do
  local _classInfoExpCallSuperNode = {}
  _className2InfoMap.ExpCallSuperNode = _classInfoExpCallSuperNode
  end
do
  local _classInfoExpCastNode = {}
  _className2InfoMap.ExpCastNode = _classInfoExpCastNode
  end
do
  local _classInfoExpDDDNode = {}
  _className2InfoMap.ExpDDDNode = _classInfoExpDDDNode
  end
do
  local _classInfoExpListNode = {}
  _className2InfoMap.ExpListNode = _classInfoExpListNode
  end
do
  local _classInfoExpMacroExpNode = {}
  _className2InfoMap.ExpMacroExpNode = _classInfoExpMacroExpNode
  end
do
  local _classInfoExpMacroStatNode = {}
  _className2InfoMap.ExpMacroStatNode = _classInfoExpMacroStatNode
  end
do
  local _classInfoExpNewNode = {}
  _className2InfoMap.ExpNewNode = _classInfoExpNewNode
  end
do
  local _classInfoExpOp1Node = {}
  _className2InfoMap.ExpOp1Node = _classInfoExpOp1Node
  end
do
  local _classInfoExpOp2Node = {}
  _className2InfoMap.ExpOp2Node = _classInfoExpOp2Node
  end
do
  local _classInfoExpParenNode = {}
  _className2InfoMap.ExpParenNode = _classInfoExpParenNode
  end
do
  local _classInfoExpRefItemNode = {}
  _className2InfoMap.ExpRefItemNode = _classInfoExpRefItemNode
  end
do
  local _classInfoExpRefNode = {}
  _className2InfoMap.ExpRefNode = _classInfoExpRefNode
  end
do
  local _classInfoFilter = {}
  _className2InfoMap.Filter = _classInfoFilter
  end
do
  local _classInfoForNode = {}
  _className2InfoMap.ForNode = _classInfoForNode
  end
do
  local _classInfoForeachNode = {}
  _className2InfoMap.ForeachNode = _classInfoForeachNode
  end
do
  local _classInfoForsortNode = {}
  _className2InfoMap.ForsortNode = _classInfoForsortNode
  end
do
  local _classInfoIfNode = {}
  _className2InfoMap.IfNode = _classInfoIfNode
  end
do
  local _classInfoIfStmtInfo = {}
  _className2InfoMap.IfStmtInfo = _classInfoIfStmtInfo
  end
do
  local _classInfoImportNode = {}
  _className2InfoMap.ImportNode = _classInfoImportNode
  end
do
  local _classInfoLiteralArrayNode = {}
  _className2InfoMap.LiteralArrayNode = _classInfoLiteralArrayNode
  end
do
  local _classInfoLiteralBoolNode = {}
  _className2InfoMap.LiteralBoolNode = _classInfoLiteralBoolNode
  end
do
  local _classInfoLiteralCharNode = {}
  _className2InfoMap.LiteralCharNode = _classInfoLiteralCharNode
  end
do
  local _classInfoLiteralIntNode = {}
  _className2InfoMap.LiteralIntNode = _classInfoLiteralIntNode
  end
do
  local _classInfoLiteralListNode = {}
  _className2InfoMap.LiteralListNode = _classInfoLiteralListNode
  end
do
  local _classInfoLiteralMapNode = {}
  _className2InfoMap.LiteralMapNode = _classInfoLiteralMapNode
  end
do
  local _classInfoLiteralNilNode = {}
  _className2InfoMap.LiteralNilNode = _classInfoLiteralNilNode
  end
do
  local _classInfoLiteralRealNode = {}
  _className2InfoMap.LiteralRealNode = _classInfoLiteralRealNode
  end
do
  local _classInfoLiteralStringNode = {}
  _className2InfoMap.LiteralStringNode = _classInfoLiteralStringNode
  end
do
  local _classInfoLiteralSymbolNode = {}
  _className2InfoMap.LiteralSymbolNode = _classInfoLiteralSymbolNode
  end
do
  local _classInfoMacroEval = {}
  _className2InfoMap.MacroEval = _classInfoMacroEval
  end
do
  local _classInfoNamespaceInfo = {}
  _className2InfoMap.NamespaceInfo = _classInfoNamespaceInfo
  _classInfoNamespaceInfo.name = {
    name='name', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 18 }
  _classInfoNamespaceInfo.scope = {
    name='scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3146 }
  _classInfoNamespaceInfo.typeInfo = {
    name='typeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3082 }
  end
do
  local _classInfoNode = {}
  _className2InfoMap.Node = _classInfoNode
  _classInfoNode.filter = {
    name='filter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 26 }
  end
do
  local _classInfoNodePos = {}
  _className2InfoMap.NodePos = _classInfoNodePos
  end
do
  local _classInfoNoneNode = {}
  _className2InfoMap.NoneNode = _classInfoNoneNode
  end
do
  local _classInfoPairItem = {}
  _className2InfoMap.PairItem = _classInfoPairItem
  end
do
  local _classInfoParser = {}
  _className2InfoMap.Parser = _classInfoParser
  end
do
  local _classInfoPosition = {}
  _className2InfoMap.Position = _classInfoPosition
  _classInfoPosition.column = {
    name='column', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoPosition.lineNo = {
    name='lineNo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  end
do
  local _classInfoRefFieldNode = {}
  _className2InfoMap.RefFieldNode = _classInfoRefFieldNode
  end
do
  local _classInfoRefTypeNode = {}
  _className2InfoMap.RefTypeNode = _classInfoRefTypeNode
  end
do
  local _classInfoRepeatNode = {}
  _className2InfoMap.RepeatNode = _classInfoRepeatNode
  end
do
  local _classInfoReturnNode = {}
  _className2InfoMap.ReturnNode = _classInfoReturnNode
  end
do
  local _classInfoRootNode = {}
  _className2InfoMap.RootNode = _classInfoRootNode
  end
do
  local _classInfoScope = {}
  _className2InfoMap.Scope = _classInfoScope
  end
do
  local _classInfoStmtExpNode = {}
  _className2InfoMap.StmtExpNode = _classInfoStmtExpNode
  end
do
  local _classInfoStream = {}
  _className2InfoMap.Stream = _classInfoStream
  end
do
  local _classInfoStreamParser = {}
  _className2InfoMap.StreamParser = _classInfoStreamParser
  _classInfoStreamParser.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3040 }
  end
do
  local _classInfoSwitchNode = {}
  _className2InfoMap.SwitchNode = _classInfoSwitchNode
  end
do
  local _classInfoToken = {}
  _className2InfoMap.Token = _classInfoToken
  _classInfoToken.kind = {
    name='kind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoToken.pos = {
    name='pos', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4004 }
  _classInfoToken.txt = {
    name='txt', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 18 }
  end
do
  local _classInfoTransUnit = {}
  _className2InfoMap.TransUnit = _classInfoTransUnit
  _classInfoTransUnit.ASTInfo = {
    name='ASTInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3962 }
  _classInfoTransUnit.ApplyNode = {
    name='ApplyNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3384 }
  _classInfoTransUnit.BlockNode = {
    name='BlockNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3274 }
  _classInfoTransUnit.BreakNode = {
    name='BreakNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3444 }
  _classInfoTransUnit.CaseInfo = {
    name='CaseInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3320 }
  _classInfoTransUnit.DeclArgDDDNode = {
    name='DeclArgDDDNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3726 }
  _classInfoTransUnit.DeclArgNode = {
    name='DeclArgNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3714 }
  _classInfoTransUnit.DeclClassNode = {
    name='DeclClassNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3734 }
  _classInfoTransUnit.DeclConstrNode = {
    name='DeclConstrNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3672 }
  _classInfoTransUnit.DeclFuncInfo = {
    name='DeclFuncInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3628 }
  _classInfoTransUnit.DeclFuncNode = {
    name='DeclFuncNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3652 }
  _classInfoTransUnit.DeclMacroInfo = {
    name='DeclMacroInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3198 }
  _classInfoTransUnit.DeclMacroNode = {
    name='DeclMacroNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3760 }
  _classInfoTransUnit.DeclMemberNode = {
    name='DeclMemberNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3694 }
  _classInfoTransUnit.DeclMethodNode = {
    name='DeclMethodNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3662 }
  _classInfoTransUnit.DeclVarNode = {
    name='DeclVarNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3606 }
  _classInfoTransUnit.ExpCallNode = {
    name='ExpCallNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3522 }
  _classInfoTransUnit.ExpCallSuperNode = {
    name='ExpCallSuperNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3682 }
  _classInfoTransUnit.ExpCastNode = {
    name='ExpCastNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3488 }
  _classInfoTransUnit.ExpDDDNode = {
    name='ExpDDDNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3534 }
  _classInfoTransUnit.ExpListNode = {
    name='ExpListNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3308 }
  _classInfoTransUnit.ExpMacroExpNode = {
    name='ExpMacroExpNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3554 }
  _classInfoTransUnit.ExpMacroStatNode = {
    name='ExpMacroStatNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3566 }
  _classInfoTransUnit.ExpNewNode = {
    name='ExpNewNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3452 }
  _classInfoTransUnit.ExpOp1Node = {
    name='ExpOp1Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3498 }
  _classInfoTransUnit.ExpOp2Node = {
    name='ExpOp2Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3474 }
  _classInfoTransUnit.ExpParenNode = {
    name='ExpParenNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3544 }
  _classInfoTransUnit.ExpRefItemNode = {
    name='ExpRefItemNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3510 }
  _classInfoTransUnit.ExpRefNode = {
    name='ExpRefNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3464 }
  _classInfoTransUnit.Filter = {
    name='Filter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3170 }
  _classInfoTransUnit.ForNode = {
    name='ForNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3366 }
  _classInfoTransUnit.ForeachNode = {
    name='ForeachNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3400 }
  _classInfoTransUnit.ForsortNode = {
    name='ForsortNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3416 }
  _classInfoTransUnit.IfNode = {
    name='IfNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3296 }
  _classInfoTransUnit.IfStmtInfo = {
    name='IfStmtInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3288 }
  _classInfoTransUnit.ImportNode = {
    name='ImportNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3232 }
  _classInfoTransUnit.LiteralArrayNode = {
    name='LiteralArrayNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3814 }
  _classInfoTransUnit.LiteralBoolNode = {
    name='LiteralBoolNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3870 }
  _classInfoTransUnit.LiteralCharNode = {
    name='LiteralCharNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3778 }
  _classInfoTransUnit.LiteralIntNode = {
    name='LiteralIntNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3790 }
  _classInfoTransUnit.LiteralListNode = {
    name='LiteralListNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3824 }
  _classInfoTransUnit.LiteralMapNode = {
    name='LiteralMapNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3840 }
  _classInfoTransUnit.LiteralNilNode = {
    name='LiteralNilNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3770 }
  _classInfoTransUnit.LiteralRealNode = {
    name='LiteralRealNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3802 }
  _classInfoTransUnit.LiteralStringNode = {
    name='LiteralStringNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3856 }
  _classInfoTransUnit.LiteralSymbolNode = {
    name='LiteralSymbolNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3880 }
  _classInfoTransUnit.MacroEval = {
    name='MacroEval', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3192 }
  _classInfoTransUnit.NamespaceInfo = {
    name='NamespaceInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3190 }
  _classInfoTransUnit.Node = {
    name='Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3174 }
  _classInfoTransUnit.NodePos = {
    name='NodePos', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3172 }
  _classInfoTransUnit.NoneNode = {
    name='NoneNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3224 }
  _classInfoTransUnit.PairItem = {
    name='PairItem', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3834 }
  _classInfoTransUnit.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3970 }
  _classInfoTransUnit.Position = {
    name='Position', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3972 }
  _classInfoTransUnit.RefFieldNode = {
    name='RefFieldNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3588 }
  _classInfoTransUnit.RefTypeNode = {
    name='RefTypeNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3258 }
  _classInfoTransUnit.RepeatNode = {
    name='RepeatNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3354 }
  _classInfoTransUnit.ReturnNode = {
    name='ReturnNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3434 }
  _classInfoTransUnit.RootNode = {
    name='RootNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3242 }
  _classInfoTransUnit.Scope = {
    name='Scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3146 }
  _classInfoTransUnit.StmtExpNode = {
    name='StmtExpNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3578 }
  _classInfoTransUnit.Stream = {
    name='Stream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3974 }
  _classInfoTransUnit.StreamParser = {
    name='StreamParser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3976 }
  _classInfoTransUnit.SwitchNode = {
    name='SwitchNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3326 }
  _classInfoTransUnit.Token = {
    name='Token', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3978 }
  _classInfoTransUnit.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3212 }
  _classInfoTransUnit.TxtStream = {
    name='TxtStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3980 }
  _classInfoTransUnit.TypeInfo = {
    name='TypeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3082 }
  _classInfoTransUnit.TypeInfoKindArray = {
    name='TypeInfoKindArray', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindClass = {
    name='TypeInfoKindClass', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindFunc = {
    name='TypeInfoKindFunc', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindList = {
    name='TypeInfoKindList', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindMacro = {
    name='TypeInfoKindMacro', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindMap = {
    name='TypeInfoKindMap', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindMethod = {
    name='TypeInfoKindMethod', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindModule = {
    name='TypeInfoKindModule', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindNilable = {
    name='TypeInfoKindNilable', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindPrim = {
    name='TypeInfoKindPrim', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.TypeInfoKindRoot = {
    name='TypeInfoKindRoot', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3982 }
  _classInfoTransUnit.VarInfo = {
    name='VarInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3600 }
  _classInfoTransUnit.WhileNode = {
    name='WhileNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3342 }
  _classInfoTransUnit.WrapParser = {
    name='WrapParser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3984 }
  _classInfoTransUnit.base = {
    name='base', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3986 }
  _classInfoTransUnit.getNodeKindName = {
    name='getNodeKindName', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3220 }
  _classInfoTransUnit.isBuiltin = {
    name='isBuiltin', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3086 }
  _classInfoTransUnit.lune = {
    name='lune', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3004 }
  _classInfoTransUnit.memStream = {
    name='memStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3988 }
  _classInfoTransUnit.nodeFilter = {
    name='nodeFilter', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3222 }
  _classInfoTransUnit.nodeKind = {
    name='nodeKind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3218 }
  _classInfoTransUnit.outStream = {
    name='outStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3990 }
  _classInfoTransUnit.rootTypeId = {
    name='rootTypeId', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfoTransUnit.typeInfoKind = {
    name='typeInfoKind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3084 }
  end
do
  local _classInfoTxtStream = {}
  _className2InfoMap.TxtStream = _classInfoTxtStream
  end
do
  local _classInfoTypeInfo = {}
  _className2InfoMap.TypeInfo = _classInfoTypeInfo
  _classInfoTypeInfo.cloneToPublic = {
    name='cloneToPublic', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3098 }
  _classInfoTypeInfo.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3100 }
  _classInfoTypeInfo.createArray = {
    name='createArray', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3106 }
  _classInfoTypeInfo.createBuiltin = {
    name='createBuiltin', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3102 }
  _classInfoTypeInfo.createClass = {
    name='createClass', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3110 }
  _classInfoTypeInfo.createFunc = {
    name='createFunc', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3112 }
  _classInfoTypeInfo.createList = {
    name='createList', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3104 }
  _classInfoTypeInfo.createMap = {
    name='createMap', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3108 }
  end
do
  local _classInfoUtil = {}
  _className2InfoMap.Util = _classInfoUtil
  _classInfoUtil.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3076 }
  _classInfoUtil.base = {
    name='base', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3078 }
  _classInfoUtil.debugLog = {
    name='debugLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3072 }
  _classInfoUtil.errorLog = {
    name='errorLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3070 }
  _classInfoUtil.lune = {
    name='lune', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3080 }
  _classInfoUtil.memStream = {
    name='memStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3062 }
  _classInfoUtil.outStream = {
    name='outStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3058 }
  _classInfoUtil.profile = {
    name='profile', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3074 }
  end
do
  local _classInfoVarInfo = {}
  _className2InfoMap.VarInfo = _classInfoVarInfo
  end
do
  local _classInfoWhileNode = {}
  _className2InfoMap.WhileNode = _classInfoWhileNode
  end
do
  local _classInfoWrapParser = {}
  _className2InfoMap.WrapParser = _classInfoWrapParser
  end
do
  local _classInfobase = {}
  _className2InfoMap.base = _classInfobase
  _classInfobase.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3992 }
  _classInfobase.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 110 }
  end
do
  local _classInfodumpNode = {}
  _className2InfoMap.dumpNode = _classInfodumpNode
  _classInfodumpNode.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3992 }
  _classInfodumpNode.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 110 }
  _classInfodumpNode.dumpFilter = {
    name='dumpFilter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4042 }
  _classInfodumpNode.lune = {
    name='lune', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 106 }
  end
do
  local _classInfolune = {}
  _className2InfoMap.lune = _classInfolune
  _classInfolune.base = {
    name='base', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 108 }
  end
do
  local _classInfomemStream = {}
  _className2InfoMap.memStream = _classInfomemStream
  end
do
  local _classInfooutStream = {}
  _className2InfoMap.outStream = _classInfooutStream
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
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {106, 4042}, }
_typeInfoList[4] = { parentId = 104, typeId = 106, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {108}, }
_typeInfoList[5] = { parentId = 106, typeId = 108, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {110, 3992}, }
_typeInfoList[6] = { parentId = 108, typeId = 110, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3004, 3082, 3084, 3086, 3114, 3118, 3142, 3146, 3162, 3166, 3170, 3172, 3174, 3176, 3178, 3190, 3192, 3194, 3198, 3202, 3208, 3212, 3214, 3218, 3220, 3222, 3224, 3232, 3242, 3250, 3254, 3258, 3274, 3284, 3288, 3296, 3304, 3308, 3316, 3320, 3326, 3336, 3342, 3354, 3366, 3384, 3392, 3400, 3416, 3434, 3444, 3452, 3464, 3474, 3488, 3498, 3510, 3522, 3534, 3544, 3554, 3562, 3566, 3574, 3578, 3588, 3600, 3606, 3616, 3622, 3628, 3634, 3644, 3648, 3652, 3662, 3672, 3682, 3694, 3714, 3726, 3734, 3746, 3750, 3756, 3760, 3770, 3778, 3790, 3802, 3814, 3824, 3834, 3840, 3848, 3852, 3856, 3866, 3870, 3880, 3890, 3892, 3896, 3898, 3902, 3904, 3908, 3910, 3914, 3916, 3920, 3922, 3926, 3928, 3932, 3934, 3938, 3940, 3944, 3946, 3950, 3952, 3956, 3958, 3962, 3970, 3972, 3974, 3976, 3978, 3980, 3982, 3984, 3986, 3988, 3990}, }
_typeInfoList[7] = { parentId = 110, typeId = 3004, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3006}, }
_typeInfoList[8] = { parentId = 3004, typeId = 3006, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3008, 3056}, }
_typeInfoList[9] = { parentId = 3006, typeId = 3008, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3010, 3014, 3020, 3022, 3024, 3030, 3036, 3042, 3044, 3046, 3050, 3052, 3054}, }
_typeInfoList[10] = { parentId = 3008, typeId = 3010, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3012}, }
_typeInfoList[11] = { parentId = 3010, typeId = 3012, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[12] = { parentId = 3008, typeId = 3014, baseId = 3010, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3016, 3018}, }
_typeInfoList[13] = { parentId = 3014, typeId = 3016, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[14] = { parentId = 3014, typeId = 3018, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[15] = { parentId = 3008, typeId = 3020, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[16] = { parentId = 3008, typeId = 3022, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[17] = { parentId = 3008, typeId = 3024, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3026, 3028}, }
_typeInfoList[18] = { parentId = 3024, typeId = 3026, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[19] = { parentId = 3024, typeId = 3028, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[20] = { parentId = 3008, typeId = 3030, baseId = 3024, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3032, 3034}, }
_typeInfoList[21] = { parentId = 3030, typeId = 3032, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[22] = { parentId = 3030, typeId = 3034, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[23] = { parentId = 3008, typeId = 3036, baseId = 3024, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3038, 3040, 3048}, }
_typeInfoList[24] = { parentId = 3036, typeId = 3038, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[25] = { parentId = 3036, typeId = 3040, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3036}, children = {}, }
_typeInfoList[26] = { parentId = 3008, typeId = 3042, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[27] = { parentId = 3008, typeId = 3044, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[28] = { parentId = 3008, typeId = 3046, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[29] = { parentId = 3036, typeId = 3048, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[30] = { parentId = 3008, typeId = 3050, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[31] = { parentId = 3008, typeId = 3052, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[32] = { parentId = 3008, typeId = 3054, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[33] = { parentId = 3006, typeId = 3056, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3058, 3062, 3070, 3072, 3074, 3076, 3078, 3080}, }
_typeInfoList[34] = { parentId = 3056, typeId = 3058, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3060}, }
_typeInfoList[35] = { parentId = 3058, typeId = 3060, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[36] = { parentId = 3056, typeId = 3062, baseId = 3058, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3064, 3066, 3068}, }
_typeInfoList[37] = { parentId = 3062, typeId = 3064, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[38] = { parentId = 3062, typeId = 3066, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[39] = { parentId = 3062, typeId = 3068, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[40] = { parentId = 3056, typeId = 3070, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[41] = { parentId = 3056, typeId = 3072, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[42] = { parentId = 3056, typeId = 3074, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[43] = { parentId = 3056, typeId = 3076, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[44] = { parentId = 3056, typeId = 3078, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[45] = { parentId = 3056, typeId = 3080, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[46] = { parentId = 110, typeId = 3082, baseId = 1, txt = 'TypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3088, 3090, 3092, 3094, 3096, 3098, 3100, 3102, 3104, 3106, 3108, 3110, 3112, 3116, 3120, 3122, 3124, 3126, 3128, 3130, 3132, 3134, 3136, 3138, 3140, 3144}, }
_typeInfoList[47] = { parentId = 110, typeId = 3086, baseId = 1, txt = 'isBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[48] = { parentId = 3082, typeId = 3088, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[49] = { parentId = 3082, typeId = 3090, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[50] = { parentId = 3082, typeId = 3092, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[51] = { parentId = 3082, typeId = 3094, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[52] = { parentId = 3082, typeId = 3096, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[53] = { parentId = 3082, typeId = 3098, baseId = 1, txt = 'cloneToPublic',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[54] = { parentId = 3082, typeId = 3100, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[55] = { parentId = 3082, typeId = 3102, baseId = 1, txt = 'createBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[56] = { parentId = 3082, typeId = 3104, baseId = 1, txt = 'createList',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[57] = { parentId = 3082, typeId = 3106, baseId = 1, txt = 'createArray',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[58] = { parentId = 3082, typeId = 3108, baseId = 1, txt = 'createMap',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[59] = { parentId = 3082, typeId = 3110, baseId = 1, txt = 'createClass',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[60] = { parentId = 3082, typeId = 3112, baseId = 1, txt = 'createFunc',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[61] = { parentId = 110, typeId = 3114, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[62] = { parentId = 3082, typeId = 3116, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3114}, children = {}, }
_typeInfoList[63] = { parentId = 110, typeId = 3118, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[64] = { parentId = 3082, typeId = 3120, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3118}, children = {}, }
_typeInfoList[65] = { parentId = 3082, typeId = 3122, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[66] = { parentId = 3082, typeId = 3124, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[67] = { parentId = 3082, typeId = 3126, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[68] = { parentId = 3082, typeId = 3128, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[69] = { parentId = 3082, typeId = 3130, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[70] = { parentId = 3082, typeId = 3132, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[71] = { parentId = 3082, typeId = 3134, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[72] = { parentId = 3082, typeId = 3136, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[73] = { parentId = 3082, typeId = 3138, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[74] = { parentId = 3082, typeId = 3140, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[75] = { parentId = 110, typeId = 3142, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[76] = { parentId = 3082, typeId = 3144, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3142}, children = {}, }
_typeInfoList[77] = { parentId = 110, typeId = 3146, baseId = 1, txt = 'Scope',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3148, 3150, 3152, 3154, 3156, 3158, 3160, 3164, 3168}, }
_typeInfoList[78] = { parentId = 3146, typeId = 3148, baseId = 1, txt = 'add',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[79] = { parentId = 3146, typeId = 3150, baseId = 1, txt = 'addClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[80] = { parentId = 3146, typeId = 3152, baseId = 1, txt = 'getClassScope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3146}, children = {}, }
_typeInfoList[81] = { parentId = 3146, typeId = 3154, baseId = 1, txt = 'getTypeInfoChild',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[82] = { parentId = 3146, typeId = 3156, baseId = 1, txt = 'getTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[83] = { parentId = 3146, typeId = 3158, baseId = 1, txt = 'getTypeInfoMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[84] = { parentId = 3146, typeId = 3160, baseId = 1, txt = 'get_parent',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3146}, children = {}, }
_typeInfoList[85] = { parentId = 110, typeId = 3162, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 3082}, retTypeId = {}, children = {}, }
_typeInfoList[86] = { parentId = 3146, typeId = 3164, baseId = 1, txt = 'get_symbol2TypeInfoMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3162}, children = {}, }
_typeInfoList[87] = { parentId = 110, typeId = 3166, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 3146}, retTypeId = {}, children = {}, }
_typeInfoList[88] = { parentId = 3146, typeId = 3168, baseId = 1, txt = 'get_className2ScopeMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3166}, children = {}, }
_typeInfoList[89] = { parentId = 110, typeId = 3170, baseId = 1, txt = 'Filter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3226, 3234, 3244, 3260, 3276, 3298, 3310, 3328, 3344, 3356, 3368, 3386, 3402, 3418, 3436, 3446, 3454, 3466, 3476, 3490, 3500, 3512, 3524, 3536, 3546, 3556, 3568, 3580, 3590, 3608, 3654, 3664, 3674, 3684, 3696, 3716, 3728, 3736, 3762, 3772, 3780, 3792, 3804, 3816, 3826, 3842, 3858, 3872, 3882}, }
_typeInfoList[90] = { parentId = 110, typeId = 3172, baseId = 1, txt = 'NodePos',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[91] = { parentId = 110, typeId = 3174, baseId = 1, txt = 'Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3180, 3182, 3184, 3186, 3188}, }
_typeInfoList[92] = { parentId = 110, typeId = 3176, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[93] = { parentId = 110, typeId = 3178, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[94] = { parentId = 3174, typeId = 3180, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3176, 3178}, children = {}, }
_typeInfoList[95] = { parentId = 3174, typeId = 3182, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[96] = { parentId = 3174, typeId = 3184, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[97] = { parentId = 3174, typeId = 3186, baseId = 1, txt = 'get_expType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[98] = { parentId = 3174, typeId = 3188, baseId = 1, txt = 'get_info',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[99] = { parentId = 110, typeId = 3190, baseId = 1, txt = 'NamespaceInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[100] = { parentId = 110, typeId = 3192, baseId = 1, txt = 'MacroEval',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3196}, }
_typeInfoList[101] = { parentId = 110, typeId = 3194, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, retTypeId = {}, children = {}, }
_typeInfoList[102] = { parentId = 3192, typeId = 3196, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3194}, children = {}, }
_typeInfoList[103] = { parentId = 110, typeId = 3198, baseId = 1, txt = 'DeclMacroInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3200, 3204, 3206, 3210}, }
_typeInfoList[104] = { parentId = 3198, typeId = 3200, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[105] = { parentId = 110, typeId = 3202, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3174}, retTypeId = {}, children = {}, }
_typeInfoList[106] = { parentId = 3198, typeId = 3204, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3202}, children = {}, }
_typeInfoList[107] = { parentId = 3198, typeId = 3206, baseId = 1, txt = 'get_ast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[108] = { parentId = 110, typeId = 3208, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3022}, retTypeId = {}, children = {}, }
_typeInfoList[109] = { parentId = 3198, typeId = 3210, baseId = 1, txt = 'get_tokenList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3208}, children = {}, }
_typeInfoList[110] = { parentId = 110, typeId = 3212, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3216, 3968}, }
_typeInfoList[111] = { parentId = 110, typeId = 3214, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[112] = { parentId = 3212, typeId = 3216, baseId = 1, txt = 'get_errMessList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3214}, children = {}, }
_typeInfoList[113] = { parentId = 110, typeId = 3220, baseId = 1, txt = 'getNodeKindName',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[114] = { parentId = 110, typeId = 3222, baseId = 1, txt = 'nodeFilter',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[115] = { parentId = 110, typeId = 3224, baseId = 3174, txt = 'NoneNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3228, 3230}, }
_typeInfoList[116] = { parentId = 3170, typeId = 3226, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[117] = { parentId = 3224, typeId = 3228, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[118] = { parentId = 3224, typeId = 3230, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[119] = { parentId = 110, typeId = 3232, baseId = 3174, txt = 'ImportNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3236, 3238, 3240}, }
_typeInfoList[120] = { parentId = 3170, typeId = 3234, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[121] = { parentId = 3232, typeId = 3236, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[122] = { parentId = 3232, typeId = 3238, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[123] = { parentId = 3232, typeId = 3240, baseId = 1, txt = 'get_modulePath',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[124] = { parentId = 110, typeId = 3242, baseId = 3174, txt = 'RootNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3246, 3248, 3252, 3256}, }
_typeInfoList[125] = { parentId = 3170, typeId = 3244, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[126] = { parentId = 3242, typeId = 3246, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[127] = { parentId = 3242, typeId = 3248, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[128] = { parentId = 110, typeId = 3250, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3174}, retTypeId = {}, children = {}, }
_typeInfoList[129] = { parentId = 3242, typeId = 3252, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3250}, children = {}, }
_typeInfoList[130] = { parentId = 110, typeId = 3254, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {12, 3190}, retTypeId = {}, children = {}, }
_typeInfoList[131] = { parentId = 3242, typeId = 3256, baseId = 1, txt = 'get_typeId2ClassMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3254}, children = {}, }
_typeInfoList[132] = { parentId = 110, typeId = 3258, baseId = 3174, txt = 'RefTypeNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3262, 3264, 3266, 3268, 3270, 3272}, }
_typeInfoList[133] = { parentId = 3170, typeId = 3260, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[134] = { parentId = 3258, typeId = 3262, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[135] = { parentId = 3258, typeId = 3264, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[136] = { parentId = 3258, typeId = 3266, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[137] = { parentId = 3258, typeId = 3268, baseId = 1, txt = 'get_refFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[138] = { parentId = 3258, typeId = 3270, baseId = 1, txt = 'get_mutFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[139] = { parentId = 3258, typeId = 3272, baseId = 1, txt = 'get_array',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[140] = { parentId = 110, typeId = 3274, baseId = 3174, txt = 'BlockNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3278, 3280, 3282, 3286}, }
_typeInfoList[141] = { parentId = 3170, typeId = 3276, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[142] = { parentId = 3274, typeId = 3278, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[143] = { parentId = 3274, typeId = 3280, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[144] = { parentId = 3274, typeId = 3282, baseId = 1, txt = 'get_blockKind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[145] = { parentId = 110, typeId = 3284, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3174}, retTypeId = {}, children = {}, }
_typeInfoList[146] = { parentId = 3274, typeId = 3286, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3284}, children = {}, }
_typeInfoList[147] = { parentId = 110, typeId = 3288, baseId = 1, txt = 'IfStmtInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3290, 3292, 3294}, }
_typeInfoList[148] = { parentId = 3288, typeId = 3290, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[149] = { parentId = 3288, typeId = 3292, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[150] = { parentId = 3288, typeId = 3294, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[151] = { parentId = 110, typeId = 3296, baseId = 3174, txt = 'IfNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3300, 3302, 3306}, }
_typeInfoList[152] = { parentId = 3170, typeId = 3298, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[153] = { parentId = 3296, typeId = 3300, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[154] = { parentId = 3296, typeId = 3302, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[155] = { parentId = 110, typeId = 3304, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3288}, retTypeId = {}, children = {}, }
_typeInfoList[156] = { parentId = 3296, typeId = 3306, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3304}, children = {}, }
_typeInfoList[157] = { parentId = 110, typeId = 3308, baseId = 3174, txt = 'ExpListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3312, 3314, 3318}, }
_typeInfoList[158] = { parentId = 3170, typeId = 3310, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[159] = { parentId = 3308, typeId = 3312, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[160] = { parentId = 3308, typeId = 3314, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[161] = { parentId = 110, typeId = 3316, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3174}, retTypeId = {}, children = {}, }
_typeInfoList[162] = { parentId = 3308, typeId = 3318, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3316}, children = {}, }
_typeInfoList[163] = { parentId = 110, typeId = 3320, baseId = 1, txt = 'CaseInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3322, 3324}, }
_typeInfoList[164] = { parentId = 3320, typeId = 3322, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[165] = { parentId = 3320, typeId = 3324, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[166] = { parentId = 110, typeId = 3326, baseId = 3174, txt = 'SwitchNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3330, 3332, 3334, 3338, 3340}, }
_typeInfoList[167] = { parentId = 3170, typeId = 3328, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[168] = { parentId = 3326, typeId = 3330, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[169] = { parentId = 3326, typeId = 3332, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[170] = { parentId = 3326, typeId = 3334, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[171] = { parentId = 110, typeId = 3336, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3320}, retTypeId = {}, children = {}, }
_typeInfoList[172] = { parentId = 3326, typeId = 3338, baseId = 1, txt = 'get_caseList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3336}, children = {}, }
_typeInfoList[173] = { parentId = 3326, typeId = 3340, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[174] = { parentId = 110, typeId = 3342, baseId = 3174, txt = 'WhileNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3346, 3348, 3350, 3352}, }
_typeInfoList[175] = { parentId = 3170, typeId = 3344, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[176] = { parentId = 3342, typeId = 3346, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[177] = { parentId = 3342, typeId = 3348, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[178] = { parentId = 3342, typeId = 3350, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[179] = { parentId = 3342, typeId = 3352, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[180] = { parentId = 110, typeId = 3354, baseId = 3174, txt = 'RepeatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3358, 3360, 3362, 3364}, }
_typeInfoList[181] = { parentId = 3170, typeId = 3356, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[182] = { parentId = 3354, typeId = 3358, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[183] = { parentId = 3354, typeId = 3360, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[184] = { parentId = 3354, typeId = 3362, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[185] = { parentId = 3354, typeId = 3364, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[186] = { parentId = 110, typeId = 3366, baseId = 3174, txt = 'ForNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3370, 3372, 3374, 3376, 3378, 3380, 3382}, }
_typeInfoList[187] = { parentId = 3170, typeId = 3368, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[188] = { parentId = 3366, typeId = 3370, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[189] = { parentId = 3366, typeId = 3372, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[190] = { parentId = 3366, typeId = 3374, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[191] = { parentId = 3366, typeId = 3376, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[192] = { parentId = 3366, typeId = 3378, baseId = 1, txt = 'get_init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[193] = { parentId = 3366, typeId = 3380, baseId = 1, txt = 'get_to',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[194] = { parentId = 3366, typeId = 3382, baseId = 1, txt = 'get_delta',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[195] = { parentId = 110, typeId = 3384, baseId = 3174, txt = 'ApplyNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3388, 3390, 3394, 3396, 3398}, }
_typeInfoList[196] = { parentId = 3170, typeId = 3386, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[197] = { parentId = 3384, typeId = 3388, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[198] = { parentId = 3384, typeId = 3390, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[199] = { parentId = 110, typeId = 3392, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3022}, retTypeId = {}, children = {}, }
_typeInfoList[200] = { parentId = 3384, typeId = 3394, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3392}, children = {}, }
_typeInfoList[201] = { parentId = 3384, typeId = 3396, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[202] = { parentId = 3384, typeId = 3398, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[203] = { parentId = 110, typeId = 3400, baseId = 3174, txt = 'ForeachNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3404, 3406, 3408, 3410, 3412, 3414}, }
_typeInfoList[204] = { parentId = 3170, typeId = 3402, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[205] = { parentId = 3400, typeId = 3404, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[206] = { parentId = 3400, typeId = 3406, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[207] = { parentId = 3400, typeId = 3408, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[208] = { parentId = 3400, typeId = 3410, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[209] = { parentId = 3400, typeId = 3412, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[210] = { parentId = 3400, typeId = 3414, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[211] = { parentId = 110, typeId = 3416, baseId = 3174, txt = 'ForsortNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3420, 3422, 3424, 3426, 3428, 3430, 3432}, }
_typeInfoList[212] = { parentId = 3170, typeId = 3418, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[213] = { parentId = 3416, typeId = 3420, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[214] = { parentId = 3416, typeId = 3422, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[215] = { parentId = 3416, typeId = 3424, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[216] = { parentId = 3416, typeId = 3426, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[217] = { parentId = 3416, typeId = 3428, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[218] = { parentId = 3416, typeId = 3430, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[219] = { parentId = 3416, typeId = 3432, baseId = 1, txt = 'get_sort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[220] = { parentId = 110, typeId = 3434, baseId = 3174, txt = 'ReturnNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3438, 3440, 3442}, }
_typeInfoList[221] = { parentId = 3170, typeId = 3436, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[222] = { parentId = 3434, typeId = 3438, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[223] = { parentId = 3434, typeId = 3440, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[224] = { parentId = 3434, typeId = 3442, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[225] = { parentId = 110, typeId = 3444, baseId = 3174, txt = 'BreakNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3448, 3450}, }
_typeInfoList[226] = { parentId = 3170, typeId = 3446, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[227] = { parentId = 3444, typeId = 3448, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[228] = { parentId = 3444, typeId = 3450, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[229] = { parentId = 110, typeId = 3452, baseId = 3174, txt = 'ExpNewNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3456, 3458, 3460, 3462}, }
_typeInfoList[230] = { parentId = 3170, typeId = 3454, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[231] = { parentId = 3452, typeId = 3456, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[232] = { parentId = 3452, typeId = 3458, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[233] = { parentId = 3452, typeId = 3460, baseId = 1, txt = 'get_symbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[234] = { parentId = 3452, typeId = 3462, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[235] = { parentId = 110, typeId = 3464, baseId = 3174, txt = 'ExpRefNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3468, 3470, 3472}, }
_typeInfoList[236] = { parentId = 3170, typeId = 3466, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[237] = { parentId = 3464, typeId = 3468, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[238] = { parentId = 3464, typeId = 3470, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[239] = { parentId = 3464, typeId = 3472, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[240] = { parentId = 110, typeId = 3474, baseId = 3174, txt = 'ExpOp2Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3478, 3480, 3482, 3484, 3486}, }
_typeInfoList[241] = { parentId = 3170, typeId = 3476, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[242] = { parentId = 3474, typeId = 3478, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[243] = { parentId = 3474, typeId = 3480, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[244] = { parentId = 3474, typeId = 3482, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[245] = { parentId = 3474, typeId = 3484, baseId = 1, txt = 'get_exp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[246] = { parentId = 3474, typeId = 3486, baseId = 1, txt = 'get_exp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[247] = { parentId = 110, typeId = 3488, baseId = 3174, txt = 'ExpCastNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3492, 3494, 3496}, }
_typeInfoList[248] = { parentId = 3170, typeId = 3490, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[249] = { parentId = 3488, typeId = 3492, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[250] = { parentId = 3488, typeId = 3494, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[251] = { parentId = 3488, typeId = 3496, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[252] = { parentId = 110, typeId = 3498, baseId = 3174, txt = 'ExpOp1Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3502, 3504, 3506, 3508}, }
_typeInfoList[253] = { parentId = 3170, typeId = 3500, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[254] = { parentId = 3498, typeId = 3502, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[255] = { parentId = 3498, typeId = 3504, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[256] = { parentId = 3498, typeId = 3506, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[257] = { parentId = 3498, typeId = 3508, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[258] = { parentId = 110, typeId = 3510, baseId = 3174, txt = 'ExpRefItemNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3514, 3516, 3518, 3520}, }
_typeInfoList[259] = { parentId = 3170, typeId = 3512, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[260] = { parentId = 3510, typeId = 3514, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[261] = { parentId = 3510, typeId = 3516, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[262] = { parentId = 3510, typeId = 3518, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[263] = { parentId = 3510, typeId = 3520, baseId = 1, txt = 'get_index',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[264] = { parentId = 110, typeId = 3522, baseId = 3174, txt = 'ExpCallNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3526, 3528, 3530, 3532}, }
_typeInfoList[265] = { parentId = 3170, typeId = 3524, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[266] = { parentId = 3522, typeId = 3526, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[267] = { parentId = 3522, typeId = 3528, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[268] = { parentId = 3522, typeId = 3530, baseId = 1, txt = 'get_func',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[269] = { parentId = 3522, typeId = 3532, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[270] = { parentId = 110, typeId = 3534, baseId = 3174, txt = 'ExpDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3538, 3540, 3542}, }
_typeInfoList[271] = { parentId = 3170, typeId = 3536, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[272] = { parentId = 3534, typeId = 3538, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[273] = { parentId = 3534, typeId = 3540, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[274] = { parentId = 3534, typeId = 3542, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[275] = { parentId = 110, typeId = 3544, baseId = 3174, txt = 'ExpParenNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3548, 3550, 3552}, }
_typeInfoList[276] = { parentId = 3170, typeId = 3546, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[277] = { parentId = 3544, typeId = 3548, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[278] = { parentId = 3544, typeId = 3550, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[279] = { parentId = 3544, typeId = 3552, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[280] = { parentId = 110, typeId = 3554, baseId = 3174, txt = 'ExpMacroExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3558, 3560, 3564}, }
_typeInfoList[281] = { parentId = 3170, typeId = 3556, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[282] = { parentId = 3554, typeId = 3558, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[283] = { parentId = 3554, typeId = 3560, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[284] = { parentId = 110, typeId = 3562, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3174}, retTypeId = {}, children = {}, }
_typeInfoList[285] = { parentId = 3554, typeId = 3564, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3562}, children = {}, }
_typeInfoList[286] = { parentId = 110, typeId = 3566, baseId = 3174, txt = 'ExpMacroStatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3570, 3572, 3576, 3960}, }
_typeInfoList[287] = { parentId = 3170, typeId = 3568, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[288] = { parentId = 3566, typeId = 3570, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[289] = { parentId = 3566, typeId = 3572, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[290] = { parentId = 110, typeId = 3574, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3174}, retTypeId = {}, children = {}, }
_typeInfoList[291] = { parentId = 3566, typeId = 3576, baseId = 1, txt = 'get_expStrList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3574}, children = {}, }
_typeInfoList[292] = { parentId = 110, typeId = 3578, baseId = 3174, txt = 'StmtExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3582, 3584, 3586}, }
_typeInfoList[293] = { parentId = 3170, typeId = 3580, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[294] = { parentId = 3578, typeId = 3582, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[295] = { parentId = 3578, typeId = 3584, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[296] = { parentId = 3578, typeId = 3586, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[297] = { parentId = 110, typeId = 3588, baseId = 3174, txt = 'RefFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3592, 3594, 3596, 3598, 3954}, }
_typeInfoList[298] = { parentId = 3170, typeId = 3590, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[299] = { parentId = 3588, typeId = 3592, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[300] = { parentId = 3588, typeId = 3594, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[301] = { parentId = 3588, typeId = 3596, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[302] = { parentId = 3588, typeId = 3598, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[303] = { parentId = 110, typeId = 3600, baseId = 1, txt = 'VarInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3602, 3604}, }
_typeInfoList[304] = { parentId = 3600, typeId = 3602, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[305] = { parentId = 3600, typeId = 3604, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[306] = { parentId = 110, typeId = 3606, baseId = 3174, txt = 'DeclVarNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3610, 3612, 3614, 3618, 3620, 3624, 3626}, }
_typeInfoList[307] = { parentId = 3170, typeId = 3608, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[308] = { parentId = 3606, typeId = 3610, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[309] = { parentId = 3606, typeId = 3612, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[310] = { parentId = 3606, typeId = 3614, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[311] = { parentId = 110, typeId = 3616, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3600}, retTypeId = {}, children = {}, }
_typeInfoList[312] = { parentId = 3606, typeId = 3618, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3616}, children = {}, }
_typeInfoList[313] = { parentId = 3606, typeId = 3620, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[314] = { parentId = 110, typeId = 3622, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[315] = { parentId = 3606, typeId = 3624, baseId = 1, txt = 'get_typeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3622}, children = {}, }
_typeInfoList[316] = { parentId = 3606, typeId = 3626, baseId = 1, txt = 'get_unwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3274}, children = {}, }
_typeInfoList[317] = { parentId = 110, typeId = 3628, baseId = 1, txt = 'DeclFuncInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3630, 3632, 3636, 3638, 3640, 3642, 3646, 3650}, }
_typeInfoList[318] = { parentId = 3628, typeId = 3630, baseId = 1, txt = 'get_className',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[319] = { parentId = 3628, typeId = 3632, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[320] = { parentId = 110, typeId = 3634, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3174}, retTypeId = {}, children = {}, }
_typeInfoList[321] = { parentId = 3628, typeId = 3636, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3634}, children = {}, }
_typeInfoList[322] = { parentId = 3628, typeId = 3638, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[323] = { parentId = 3628, typeId = 3640, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[324] = { parentId = 3628, typeId = 3642, baseId = 1, txt = 'get_body',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[325] = { parentId = 110, typeId = 3644, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[326] = { parentId = 3628, typeId = 3646, baseId = 1, txt = 'get_retTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3644}, children = {}, }
_typeInfoList[327] = { parentId = 110, typeId = 3648, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[328] = { parentId = 3628, typeId = 3650, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3648}, children = {}, }
_typeInfoList[329] = { parentId = 110, typeId = 3652, baseId = 3174, txt = 'DeclFuncNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3656, 3658, 3660}, }
_typeInfoList[330] = { parentId = 3170, typeId = 3654, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[331] = { parentId = 3652, typeId = 3656, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[332] = { parentId = 3652, typeId = 3658, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[333] = { parentId = 3652, typeId = 3660, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3628}, children = {}, }
_typeInfoList[334] = { parentId = 110, typeId = 3662, baseId = 3174, txt = 'DeclMethodNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3666, 3668, 3670}, }
_typeInfoList[335] = { parentId = 3170, typeId = 3664, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[336] = { parentId = 3662, typeId = 3666, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[337] = { parentId = 3662, typeId = 3668, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[338] = { parentId = 3662, typeId = 3670, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3628}, children = {}, }
_typeInfoList[339] = { parentId = 110, typeId = 3672, baseId = 3174, txt = 'DeclConstrNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3676, 3678, 3680}, }
_typeInfoList[340] = { parentId = 3170, typeId = 3674, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[341] = { parentId = 3672, typeId = 3676, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[342] = { parentId = 3672, typeId = 3678, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[343] = { parentId = 3672, typeId = 3680, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3628}, children = {}, }
_typeInfoList[344] = { parentId = 110, typeId = 3682, baseId = 3174, txt = 'ExpCallSuperNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3686, 3688, 3690, 3692}, }
_typeInfoList[345] = { parentId = 3170, typeId = 3684, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[346] = { parentId = 3682, typeId = 3686, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[347] = { parentId = 3682, typeId = 3688, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[348] = { parentId = 3682, typeId = 3690, baseId = 1, txt = 'get_superType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[349] = { parentId = 3682, typeId = 3692, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[350] = { parentId = 110, typeId = 3694, baseId = 3174, txt = 'DeclMemberNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3698, 3700, 3702, 3704, 3706, 3708, 3710, 3712}, }
_typeInfoList[351] = { parentId = 3170, typeId = 3696, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[352] = { parentId = 3694, typeId = 3698, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[353] = { parentId = 3694, typeId = 3700, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[354] = { parentId = 3694, typeId = 3702, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[355] = { parentId = 3694, typeId = 3704, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3258}, children = {}, }
_typeInfoList[356] = { parentId = 3694, typeId = 3706, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[357] = { parentId = 3694, typeId = 3708, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[358] = { parentId = 3694, typeId = 3710, baseId = 1, txt = 'get_getterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[359] = { parentId = 3694, typeId = 3712, baseId = 1, txt = 'get_setterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[360] = { parentId = 110, typeId = 3714, baseId = 3174, txt = 'DeclArgNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3718, 3720, 3722, 3724}, }
_typeInfoList[361] = { parentId = 3170, typeId = 3716, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[362] = { parentId = 3714, typeId = 3718, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[363] = { parentId = 3714, typeId = 3720, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[364] = { parentId = 3714, typeId = 3722, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[365] = { parentId = 3714, typeId = 3724, baseId = 1, txt = 'get_argType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3258}, children = {}, }
_typeInfoList[366] = { parentId = 110, typeId = 3726, baseId = 3174, txt = 'DeclArgDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3730, 3732}, }
_typeInfoList[367] = { parentId = 3170, typeId = 3728, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[368] = { parentId = 3726, typeId = 3730, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[369] = { parentId = 3726, typeId = 3732, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[370] = { parentId = 110, typeId = 3734, baseId = 3174, txt = 'DeclClassNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3738, 3740, 3742, 3744, 3748, 3752, 3754, 3758}, }
_typeInfoList[371] = { parentId = 3170, typeId = 3736, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[372] = { parentId = 3734, typeId = 3738, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[373] = { parentId = 3734, typeId = 3740, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[374] = { parentId = 3734, typeId = 3742, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[375] = { parentId = 3734, typeId = 3744, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[376] = { parentId = 110, typeId = 3746, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3174}, retTypeId = {}, children = {}, }
_typeInfoList[377] = { parentId = 3734, typeId = 3748, baseId = 1, txt = 'get_fieldList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3746}, children = {}, }
_typeInfoList[378] = { parentId = 110, typeId = 3750, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3694}, retTypeId = {}, children = {}, }
_typeInfoList[379] = { parentId = 3734, typeId = 3752, baseId = 1, txt = 'get_memberList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3750}, children = {}, }
_typeInfoList[380] = { parentId = 3734, typeId = 3754, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3146}, children = {}, }
_typeInfoList[381] = { parentId = 110, typeId = 3756, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[382] = { parentId = 3734, typeId = 3758, baseId = 1, txt = 'get_outerMethodSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3756}, children = {}, }
_typeInfoList[383] = { parentId = 110, typeId = 3760, baseId = 3174, txt = 'DeclMacroNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3764, 3766, 3768}, }
_typeInfoList[384] = { parentId = 3170, typeId = 3762, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[385] = { parentId = 3760, typeId = 3764, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[386] = { parentId = 3760, typeId = 3766, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[387] = { parentId = 3760, typeId = 3768, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3198}, children = {}, }
_typeInfoList[388] = { parentId = 110, typeId = 3770, baseId = 3174, txt = 'LiteralNilNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3774, 3776, 3894}, }
_typeInfoList[389] = { parentId = 3170, typeId = 3772, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[390] = { parentId = 3770, typeId = 3774, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[391] = { parentId = 3770, typeId = 3776, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[392] = { parentId = 110, typeId = 3778, baseId = 3174, txt = 'LiteralCharNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3782, 3784, 3786, 3788, 3900}, }
_typeInfoList[393] = { parentId = 3170, typeId = 3780, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[394] = { parentId = 3778, typeId = 3782, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[395] = { parentId = 3778, typeId = 3784, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[396] = { parentId = 3778, typeId = 3786, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[397] = { parentId = 3778, typeId = 3788, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[398] = { parentId = 110, typeId = 3790, baseId = 3174, txt = 'LiteralIntNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3794, 3796, 3798, 3800, 3906}, }
_typeInfoList[399] = { parentId = 3170, typeId = 3792, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[400] = { parentId = 3790, typeId = 3794, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[401] = { parentId = 3790, typeId = 3796, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[402] = { parentId = 3790, typeId = 3798, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[403] = { parentId = 3790, typeId = 3800, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[404] = { parentId = 110, typeId = 3802, baseId = 3174, txt = 'LiteralRealNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3806, 3808, 3810, 3812, 3912}, }
_typeInfoList[405] = { parentId = 3170, typeId = 3804, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[406] = { parentId = 3802, typeId = 3806, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[407] = { parentId = 3802, typeId = 3808, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[408] = { parentId = 3802, typeId = 3810, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[409] = { parentId = 3802, typeId = 3812, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[410] = { parentId = 110, typeId = 3814, baseId = 3174, txt = 'LiteralArrayNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3818, 3820, 3822, 3918}, }
_typeInfoList[411] = { parentId = 3170, typeId = 3816, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[412] = { parentId = 3814, typeId = 3818, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[413] = { parentId = 3814, typeId = 3820, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[414] = { parentId = 3814, typeId = 3822, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[415] = { parentId = 110, typeId = 3824, baseId = 3174, txt = 'LiteralListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3828, 3830, 3832, 3924}, }
_typeInfoList[416] = { parentId = 3170, typeId = 3826, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[417] = { parentId = 3824, typeId = 3828, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[418] = { parentId = 3824, typeId = 3830, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[419] = { parentId = 3824, typeId = 3832, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3308}, children = {}, }
_typeInfoList[420] = { parentId = 110, typeId = 3834, baseId = 1, txt = 'PairItem',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3836, 3838}, }
_typeInfoList[421] = { parentId = 3834, typeId = 3836, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[422] = { parentId = 3834, typeId = 3838, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[423] = { parentId = 110, typeId = 3840, baseId = 3174, txt = 'LiteralMapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3844, 3846, 3850, 3854, 3930}, }
_typeInfoList[424] = { parentId = 3170, typeId = 3842, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[425] = { parentId = 3840, typeId = 3844, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[426] = { parentId = 3840, typeId = 3846, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[427] = { parentId = 110, typeId = 3848, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {3174, 3174}, retTypeId = {}, children = {}, }
_typeInfoList[428] = { parentId = 3840, typeId = 3850, baseId = 1, txt = 'get_map',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3848}, children = {}, }
_typeInfoList[429] = { parentId = 110, typeId = 3852, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3834}, retTypeId = {}, children = {}, }
_typeInfoList[430] = { parentId = 3840, typeId = 3854, baseId = 1, txt = 'get_pairList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3852}, children = {}, }
_typeInfoList[431] = { parentId = 110, typeId = 3856, baseId = 3174, txt = 'LiteralStringNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3860, 3862, 3864, 3868, 3936}, }
_typeInfoList[432] = { parentId = 3170, typeId = 3858, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[433] = { parentId = 3856, typeId = 3860, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[434] = { parentId = 3856, typeId = 3862, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[435] = { parentId = 3856, typeId = 3864, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[436] = { parentId = 110, typeId = 3866, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3174}, retTypeId = {}, children = {}, }
_typeInfoList[437] = { parentId = 3856, typeId = 3868, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3866}, children = {}, }
_typeInfoList[438] = { parentId = 110, typeId = 3870, baseId = 3174, txt = 'LiteralBoolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3874, 3876, 3878, 3942}, }
_typeInfoList[439] = { parentId = 3170, typeId = 3872, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[440] = { parentId = 3870, typeId = 3874, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[441] = { parentId = 3870, typeId = 3876, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[442] = { parentId = 3870, typeId = 3878, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[443] = { parentId = 110, typeId = 3880, baseId = 3174, txt = 'LiteralSymbolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3884, 3886, 3888, 3948}, }
_typeInfoList[444] = { parentId = 3170, typeId = 3882, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[445] = { parentId = 3880, typeId = 3884, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[446] = { parentId = 3880, typeId = 3886, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[447] = { parentId = 3880, typeId = 3888, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3022}, children = {}, }
_typeInfoList[448] = { parentId = 110, typeId = 3890, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[449] = { parentId = 110, typeId = 3892, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[450] = { parentId = 3770, typeId = 3894, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3890, 3892}, children = {}, }
_typeInfoList[451] = { parentId = 110, typeId = 3896, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[452] = { parentId = 110, typeId = 3898, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[453] = { parentId = 3778, typeId = 3900, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3896, 3898}, children = {}, }
_typeInfoList[454] = { parentId = 110, typeId = 3902, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[455] = { parentId = 110, typeId = 3904, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[456] = { parentId = 3790, typeId = 3906, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3902, 3904}, children = {}, }
_typeInfoList[457] = { parentId = 110, typeId = 3908, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[458] = { parentId = 110, typeId = 3910, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[459] = { parentId = 3802, typeId = 3912, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3908, 3910}, children = {}, }
_typeInfoList[460] = { parentId = 110, typeId = 3914, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[461] = { parentId = 110, typeId = 3916, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[462] = { parentId = 3814, typeId = 3918, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3914, 3916}, children = {}, }
_typeInfoList[463] = { parentId = 110, typeId = 3920, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[464] = { parentId = 110, typeId = 3922, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[465] = { parentId = 3824, typeId = 3924, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3920, 3922}, children = {}, }
_typeInfoList[466] = { parentId = 110, typeId = 3926, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[467] = { parentId = 110, typeId = 3928, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[468] = { parentId = 3840, typeId = 3930, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3926, 3928}, children = {}, }
_typeInfoList[469] = { parentId = 110, typeId = 3932, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[470] = { parentId = 110, typeId = 3934, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[471] = { parentId = 3856, typeId = 3936, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3932, 3934}, children = {}, }
_typeInfoList[472] = { parentId = 110, typeId = 3938, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[473] = { parentId = 110, typeId = 3940, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[474] = { parentId = 3870, typeId = 3942, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3938, 3940}, children = {}, }
_typeInfoList[475] = { parentId = 110, typeId = 3944, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[476] = { parentId = 110, typeId = 3946, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[477] = { parentId = 3880, typeId = 3948, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3944, 3946}, children = {}, }
_typeInfoList[478] = { parentId = 110, typeId = 3950, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[479] = { parentId = 110, typeId = 3952, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[480] = { parentId = 3588, typeId = 3954, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3950, 3952}, children = {}, }
_typeInfoList[481] = { parentId = 110, typeId = 3956, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[482] = { parentId = 110, typeId = 3958, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3082}, retTypeId = {}, children = {}, }
_typeInfoList[483] = { parentId = 3566, typeId = 3960, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3956, 3958}, children = {}, }
_typeInfoList[484] = { parentId = 110, typeId = 3962, baseId = 1, txt = 'ASTInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3964, 3966}, }
_typeInfoList[485] = { parentId = 3962, typeId = 3964, baseId = 1, txt = 'get_node',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3174}, children = {}, }
_typeInfoList[486] = { parentId = 3962, typeId = 3966, baseId = 1, txt = 'get_moduleTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3082}, children = {}, }
_typeInfoList[487] = { parentId = 3212, typeId = 3968, baseId = 1, txt = 'createAST',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3962}, children = {}, }
_typeInfoList[488] = { parentId = 110, typeId = 3970, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[489] = { parentId = 110, typeId = 3972, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[490] = { parentId = 110, typeId = 3974, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[491] = { parentId = 110, typeId = 3976, baseId = 1, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[492] = { parentId = 110, typeId = 3978, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[493] = { parentId = 110, typeId = 3980, baseId = 1, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[494] = { parentId = 110, typeId = 3982, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[495] = { parentId = 110, typeId = 3984, baseId = 1, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[496] = { parentId = 110, typeId = 3986, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[497] = { parentId = 110, typeId = 3988, baseId = 1, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[498] = { parentId = 110, typeId = 3990, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[499] = { parentId = 108, typeId = 3992, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3994, 3998, 4004, 4006, 4008, 4014, 4020, 4026, 4028, 4030, 4032, 4036, 4038, 4040}, }
_typeInfoList[500] = { parentId = 3992, typeId = 3994, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3996}, }
_typeInfoList[501] = { parentId = 3994, typeId = 3996, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[502] = { parentId = 3992, typeId = 3998, baseId = 3994, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4000, 4002}, }
_typeInfoList[503] = { parentId = 3998, typeId = 4000, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[504] = { parentId = 3998, typeId = 4002, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[505] = { parentId = 3992, typeId = 4004, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[506] = { parentId = 3992, typeId = 4006, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[507] = { parentId = 3992, typeId = 4008, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4010, 4012}, }
_typeInfoList[508] = { parentId = 4008, typeId = 4010, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4006}, children = {}, }
_typeInfoList[509] = { parentId = 4008, typeId = 4012, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[510] = { parentId = 3992, typeId = 4014, baseId = 4008, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4016, 4018}, }
_typeInfoList[511] = { parentId = 4014, typeId = 4016, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4006}, children = {}, }
_typeInfoList[512] = { parentId = 4014, typeId = 4018, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[513] = { parentId = 3992, typeId = 4020, baseId = 4008, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4022, 4024, 4034}, }
_typeInfoList[514] = { parentId = 4020, typeId = 4022, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[515] = { parentId = 4020, typeId = 4024, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {4020}, children = {}, }
_typeInfoList[516] = { parentId = 3992, typeId = 4028, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[517] = { parentId = 3992, typeId = 4030, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[518] = { parentId = 3992, typeId = 4032, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[519] = { parentId = 4020, typeId = 4034, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {4006}, children = {}, }
_typeInfoList[520] = { parentId = 3992, typeId = 4036, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[521] = { parentId = 3992, typeId = 4038, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[522] = { parentId = 3992, typeId = 4040, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[523] = { parentId = 104, typeId = 4042, baseId = 3170, txt = 'dumpFilter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {4050, 4052, 4054, 4056, 4058, 4060, 4062, 4064, 4066, 4068, 4070, 4072, 4074, 4076, 4084, 4086, 4088, 4090, 4092, 4094, 4098, 4102, 4104, 4106, 4108, 4112, 4114, 4116, 4118, 4122, 4124, 4126, 4128, 4130, 4132, 4134, 4136, 4138, 4140, 4142, 4146, 4148, 4150, 4152, 4154, 4156, 4158, 4160, 4162}, }
_typeInfoList[524] = { parentId = 4042, typeId = 4050, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[525] = { parentId = 4042, typeId = 4052, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[526] = { parentId = 4042, typeId = 4054, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[527] = { parentId = 4042, typeId = 4056, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[528] = { parentId = 4042, typeId = 4058, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[529] = { parentId = 4042, typeId = 4060, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[530] = { parentId = 4042, typeId = 4062, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[531] = { parentId = 4042, typeId = 4064, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[532] = { parentId = 4042, typeId = 4066, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[533] = { parentId = 4042, typeId = 4068, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[534] = { parentId = 4042, typeId = 4070, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[535] = { parentId = 4042, typeId = 4072, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[536] = { parentId = 4042, typeId = 4074, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[537] = { parentId = 4042, typeId = 4076, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[538] = { parentId = 4042, typeId = 4084, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[539] = { parentId = 4042, typeId = 4086, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[540] = { parentId = 4042, typeId = 4088, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[541] = { parentId = 4042, typeId = 4090, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[542] = { parentId = 4042, typeId = 4092, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[543] = { parentId = 4042, typeId = 4094, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[544] = { parentId = 4042, typeId = 4098, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[545] = { parentId = 4042, typeId = 4102, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[546] = { parentId = 4042, typeId = 4104, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[547] = { parentId = 4042, typeId = 4106, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[548] = { parentId = 4042, typeId = 4108, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[549] = { parentId = 4042, typeId = 4112, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[550] = { parentId = 4042, typeId = 4114, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[551] = { parentId = 4042, typeId = 4116, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[552] = { parentId = 4042, typeId = 4118, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[553] = { parentId = 4042, typeId = 4122, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[554] = { parentId = 4042, typeId = 4124, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[555] = { parentId = 4042, typeId = 4126, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[556] = { parentId = 4042, typeId = 4128, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[557] = { parentId = 4042, typeId = 4130, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[558] = { parentId = 4042, typeId = 4132, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[559] = { parentId = 4042, typeId = 4134, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[560] = { parentId = 4042, typeId = 4136, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[561] = { parentId = 4042, typeId = 4138, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[562] = { parentId = 4042, typeId = 4140, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[563] = { parentId = 4042, typeId = 4142, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[564] = { parentId = 4042, typeId = 4146, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[565] = { parentId = 4042, typeId = 4148, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[566] = { parentId = 4042, typeId = 4150, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[567] = { parentId = 4042, typeId = 4152, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[568] = { parentId = 4042, typeId = 4154, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[569] = { parentId = 4042, typeId = 4156, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[570] = { parentId = 4042, typeId = 4158, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[571] = { parentId = 4042, typeId = 4160, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[572] = { parentId = 4042, typeId = 4162, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
----- meta -----
return moduleObj
