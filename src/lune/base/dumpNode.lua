--lune/base/dumpNode.lns
local moduleObj = {}
local TransUnit = require( 'lune.base.TransUnit' )

local filterObj = {}
moduleObj.filterObj = filterObj
function filterObj.new(  )
  local obj = {}
  setmetatable( obj, { __index = filterObj } )
  if obj.__init then
    obj:__init(  )
  end
  return obj
end
function filterObj:__init(  )
            
end

local function dump( prefix, depth, node, txt )
  local typeStr = ""
  if node.expType and node.expType ~= TransUnit.typeInfoKind.None then
    typeStr = string.format( "(%d:%s:%s)", node.expType:get_typeId(  ), node.expType:getTxt(  ), node.expType:get_kind(  ))
  end
  print( string.format( "%s: %s %s %s", prefix, TransUnit.getNodeKindName( node.kind ), txt, typeStr) )
end

filterObj[TransUnit.nodeKind.None] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
end

filterObj[TransUnit.nodeKind.Import] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info )
end

filterObj[TransUnit.nodeKind.Root] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  local children = node.info.children
  for index, child in pairs( children ) do
    TransUnit.nodeFilter( child, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.Block] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  local stmtList = node.info.stmtList
  for index, statement in pairs( stmtList ) do
    TransUnit.nodeFilter( statement, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.StmtExp] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  TransUnit.nodeFilter( node.info, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.DeclClass] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.name.txt )
  local fieldList = node.info.fieldList
  for index, field in pairs( fieldList ) do
    TransUnit.nodeFilter( field, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.DeclMember] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.name.txt )
  TransUnit.nodeFilter( node.info.refType, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.DeclVar] = function ( self, node, prefix, depth )
  local varName = ""
  local varList = node.info.varList
  for index, var in pairs( varList ) do
    if index > 1 then
      varName = varName .. ","
    end
    varName = string.format( "%s %s", varName, var["name"].txt)
  end
  dump( prefix, depth, node, node.info.unwrap and "! " or " " .. varName )
  for index, var in pairs( varList ) do
    if var["refType"] then
      TransUnit.nodeFilter( var["refType"], self, prefix .. "  ", depth + 1 )
    end
  end
  if node.info.expList then
    TransUnit.nodeFilter( node.info.expList, self, prefix .. "  ", depth + 1 )
  end
  if node.info.unwrap then
    TransUnit.nodeFilter( node.info.unwrap, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.DeclArg] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.name.txt )
  TransUnit.nodeFilter( node.info.argType, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.DeclArgDDD] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "..." )
end

filterObj[TransUnit.nodeKind.ExpDDD] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "..." )
end

filterObj[TransUnit.nodeKind.DeclFunc] = function ( self, node, prefix, depth )
  local name = node.info.name
  dump( prefix, depth, node, name and name.txt or "<anonymous>" )
  local argList = node.info.argList
  for index, arg in pairs( argList ) do
    TransUnit.nodeFilter( arg, self, prefix .. "  ", depth + 1 )
  end
  local retTypeList = node.info.retTypeList
  for index, refType in pairs( retTypeList ) do
    TransUnit.nodeFilter( refType, self, prefix .. "  ", depth + 1 )
  end
  TransUnit.nodeFilter( node.info.body, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.DeclMethod] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.name.txt )
  local argList = node.info.argList
  for index, arg in pairs( argList ) do
    TransUnit.nodeFilter( arg, self, prefix .. "  ", depth + 1 )
  end
  local retTypeList = node.info.retTypeList
  for index, refType in pairs( retTypeList ) do
    TransUnit.nodeFilter( refType, self, prefix .. "  ", depth + 1 )
  end
  TransUnit.nodeFilter( node.info.body, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.DeclConstr] = function ( self, node, prefix, depth )
  filterObj[TransUnit.nodeKind.DeclMethod]( self, node, prefix, depth )
end

filterObj[TransUnit.nodeKind.RefType] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, (node.info.refFlag and "&" or "" ) .. (node.info.mutFlag and "mut " or "" ) )
  TransUnit.nodeFilter( node.info.name, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.If] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  local valList = node.info
  for index, val in pairs( valList ) do
    if val['exp'] then
      TransUnit.nodeFilter( val['exp'], self, prefix .. "  ", depth + 1 )
    end
    TransUnit.nodeFilter( val['block'], self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.Switch] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  TransUnit.nodeFilter( node.info.exp, self, prefix .. "  ", depth + 1 )
  local caseList = node.info.caseList
  for __index, caseInfo in pairs( caseList ) do
    TransUnit.nodeFilter( caseInfo['expList'], self, prefix .. "  ", depth + 1 )
    TransUnit.nodeFilter( caseInfo['block'], self, prefix .. "  ", depth + 1 )
  end
  TransUnit.nodeFilter( node.info.default, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.While] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  TransUnit.nodeFilter( node.info.exp, self, prefix .. "  ", depth + 1 )
  TransUnit.nodeFilter( node.info.block, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.Repeat] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  TransUnit.nodeFilter( node.info.block, self, prefix .. "  ", depth + 1 )
  TransUnit.nodeFilter( node.info.exp, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.For] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.val.txt )
  TransUnit.nodeFilter( node.info.init, self, prefix .. "  ", depth + 1 )
  TransUnit.nodeFilter( node.info.to, self, prefix .. "  ", depth + 1 )
  if node.info.delta then
    TransUnit.nodeFilter( node.info.delta, self, prefix .. "  ", depth + 1 )
  end
  TransUnit.nodeFilter( node.info.block, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.Apply] = function ( self, node, prefix, depth )
  local varNames = ""
  local varList = node.info.varList
  for index, var in pairs( varList ) do
    varNames = varNames .. var['txt'] .. " "
  end
  dump( prefix, depth, node, varNames )
  TransUnit.nodeFilter( node.info.exp, self, prefix .. "  ", depth + 1 )
  TransUnit.nodeFilter( node.info.block, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.Foreach] = function ( self, node, prefix, depth )
  local index = node.info.key and node.info.key.txt or ""
  dump( prefix, depth, node, node.info.val.txt .. " " .. index )
  TransUnit.nodeFilter( node.info.exp, self, prefix .. "  ", depth + 1 )
  TransUnit.nodeFilter( node.info.block, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.Forsort] = function ( self, node, prefix, depth )
  local index = node.info.key and node.info.key.txt or ""
  dump( prefix, depth, node, node.info.val.txt .. " " .. index )
  TransUnit.nodeFilter( node.info.exp, self, prefix .. "  ", depth + 1 )
  TransUnit.nodeFilter( node.info.block, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.ExpCall] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  TransUnit.nodeFilter( node.info.func, self, prefix .. "  ", depth + 1 )
  if node.info.argList then
    TransUnit.nodeFilter( node.info.argList, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.ExpList] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  local expList = node.info
  for index, exp in pairs( expList ) do
    TransUnit.nodeFilter( exp, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.ExpOp1] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.op.txt )
  TransUnit.nodeFilter( node.info.exp, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.ExpCast] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  TransUnit.nodeFilter( node.info, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.ExpParen] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "()" )
  TransUnit.nodeFilter( node.info, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.ExpOp2] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.op.txt )
  node.info.exp1:filter( self, prefix .. "  ", depth + 1 )
  node.info.exp2:filter( self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.ExpNew] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  TransUnit.nodeFilter( node.info.symbol, self, prefix .. "  ", depth + 1 )
  if node.info.argList then
    TransUnit.nodeFilter( node.info.argList, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.ExpRef] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.txt )
end

filterObj[TransUnit.nodeKind.ExpRefItem] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "seq[exp] " .. node:get_expType(  ):getTxt(  ) )
  TransUnit.nodeFilter( node:get_info(  ).val, self, prefix .. "  ", depth + 1 )
  TransUnit.nodeFilter( node:get_info(  ).index, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.RefField] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.field.txt )
  TransUnit.nodeFilter( node.info.prefix, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.Return] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  TransUnit.nodeFilter( node.info, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.LiteralList] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "[]" )
  if node.info then
    TransUnit.nodeFilter( node.info, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.LiteralMap] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  local pairList = node.info.pairList
  for __index, pair in pairs( pairList ) do
    TransUnit.nodeFilter( pair['key'], self, prefix .. "  ", depth + 1 )
    TransUnit.nodeFilter( pair['val'], self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.LiteralArray] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "[@]" )
  if node.info then
    TransUnit.nodeFilter( node.info, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.LiteralChar] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, string.format( "%s(%s)", node.info.num, node.info.token.txt ) )
end

filterObj[TransUnit.nodeKind.LiteralInt] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, string.format( "%s(%s)", node.info.num, node.info.token.txt ) )
end

filterObj[TransUnit.nodeKind.LiteralReal] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, string.format( "%s(%s)", node.info.num, node.info.token.txt ) )
end

filterObj[TransUnit.nodeKind.LiteralString] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.token.txt )
end

filterObj[TransUnit.nodeKind.LiteralBool] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.txt == "true" )
end

filterObj[TransUnit.nodeKind.LiteralNil] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
end

filterObj[TransUnit.nodeKind.Break] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
end

----- meta -----
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
local _classInfofilterObj = {}
_className2InfoMap.filterObj = _classInfofilterObj
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 88, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[2] = { parentId = 1, typeId = 150, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[3] = { parentId = 1, typeId = 260, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[4] = { parentId = 1, typeId = 262, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[5] = { parentId = 1, typeId = 264, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {282, 284, 340}, }
_typeInfoList[6] = { parentId = 1, typeId = 286, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[7] = { parentId = 1, typeId = 312, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[8] = { parentId = 1, typeId = 314, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[9] = { parentId = 1, typeId = 316, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[10] = { parentId = 1, typeId = 348, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[11] = { parentId = 1, typeId = 350, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[12] = { parentId = 1, typeId = 352, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[13] = { parentId = 1, typeId = 354, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[14] = { parentId = 1, typeId = 356, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[15] = { parentId = 1, typeId = 358, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {372, 374, 376}, }
_typeInfoList[16] = { parentId = 1, typeId = 360, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[17] = { parentId = 1, typeId = 364, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[18] = { parentId = 1, typeId = 366, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[19] = { parentId = 1, typeId = 368, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[20] = { parentId = 1, typeId = 370, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[21] = { parentId = 1, typeId = 378, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[22] = { parentId = 1, typeId = 380, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[23] = { parentId = 1, typeId = 392, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[24] = { parentId = 1, typeId = 398, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {418, 420, 422, 424, 430, 432, 434, 440, 442, 446, 450, 452, 456, 462, 470, 474, 476, 478, 480, 482, 484, 486, 488, 490, 492, 494, 498}, }
_typeInfoList[25] = { parentId = 1, typeId = 468, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[26] = { parentId = 1, typeId = 472, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[27] = { parentId = 1, typeId = 496, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[28] = { parentId = 1, typeId = 500, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {516, 518, 520, 522, 524, 526, 530, 534}, }
_typeInfoList[29] = { parentId = 1, typeId = 528, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 398}, retTypeId = {}, children = {}, }
_typeInfoList[30] = { parentId = 1, typeId = 532, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 500}, retTypeId = {}, children = {}, }
_typeInfoList[31] = { parentId = 1, typeId = 538, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[32] = { parentId = 1, typeId = 540, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {544, 546, 550}, }
_typeInfoList[33] = { parentId = 1, typeId = 552, baseId = 540, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[34] = { parentId = 1, typeId = 554, baseId = 540, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[35] = { parentId = 1, typeId = 556, baseId = 540, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[36] = { parentId = 1, typeId = 558, baseId = 540, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[37] = { parentId = 1, typeId = 560, baseId = 540, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[38] = { parentId = 1, typeId = 562, baseId = 540, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[39] = { parentId = 1, typeId = 564, baseId = 540, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[40] = { parentId = 1, typeId = 566, baseId = 540, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[41] = { parentId = 1, typeId = 568, baseId = 540, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[42] = { parentId = 1, typeId = 570, baseId = 540, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[43] = { parentId = 1, typeId = 572, baseId = 540, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[44] = { parentId = 1, typeId = 574, baseId = 540, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[45] = { parentId = 1, typeId = 576, baseId = 540, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[46] = { parentId = 1, typeId = 578, baseId = 540, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[47] = { parentId = 1, typeId = 580, baseId = 540, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[48] = { parentId = 1, typeId = 582, baseId = 540, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[49] = { parentId = 1, typeId = 584, baseId = 540, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[50] = { parentId = 1, typeId = 586, baseId = 540, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[51] = { parentId = 1, typeId = 588, baseId = 540, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[52] = { parentId = 1, typeId = 590, baseId = 540, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[53] = { parentId = 1, typeId = 592, baseId = 540, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[54] = { parentId = 1, typeId = 594, baseId = 540, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[55] = { parentId = 1, typeId = 596, baseId = 540, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[56] = { parentId = 1, typeId = 598, baseId = 540, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[57] = { parentId = 1, typeId = 600, baseId = 540, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[58] = { parentId = 1, typeId = 602, baseId = 540, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[59] = { parentId = 1, typeId = 604, baseId = 540, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[60] = { parentId = 1, typeId = 606, baseId = 540, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[61] = { parentId = 1, typeId = 608, baseId = 540, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[62] = { parentId = 1, typeId = 610, baseId = 540, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[63] = { parentId = 1, typeId = 612, baseId = 540, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[64] = { parentId = 1, typeId = 614, baseId = 540, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[65] = { parentId = 1, typeId = 616, baseId = 540, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[66] = { parentId = 1, typeId = 618, baseId = 540, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[67] = { parentId = 1, typeId = 620, baseId = 540, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[68] = { parentId = 1, typeId = 622, baseId = 540, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[69] = { parentId = 1, typeId = 624, baseId = 540, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[70] = { parentId = 1, typeId = 628, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {708, 1250}, }
_typeInfoList[71] = { parentId = 1, typeId = 706, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[72] = { parentId = 1, typeId = 760, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[73] = { parentId = 1, typeId = 762, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[74] = { parentId = 1, typeId = 1094, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1102, 1104, 1108, 1110, 1112, 1114, 1118, 1122}, }
_typeInfoList[75] = { parentId = 1, typeId = 1106, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {540}, retTypeId = {}, children = {}, }
_typeInfoList[76] = { parentId = 1, typeId = 1116, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[77] = { parentId = 1, typeId = 1120, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[78] = { parentId = 1, typeId = 1208, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1212, 1216}, }
_typeInfoList[79] = { parentId = 1, typeId = 1214, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {540}, retTypeId = {}, children = {}, }
_typeInfoList[80] = { parentId = 1, typeId = 1272, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[81] = { parentId = 1, typeId = 1274, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[82] = { parentId = 1, typeId = 1276, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[83] = { parentId = 1, typeId = 1278, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[84] = { parentId = 1, typeId = 1280, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1434, 1436, 1438}, }
_typeInfoList[85] = { parentId = 1, typeId = 1282, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[86] = { parentId = 1, typeId = 1284, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[87] = { parentId = 1, typeId = 1286, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[88] = { parentId = 1, typeId = 1288, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[89] = { parentId = 1, typeId = 1290, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[90] = { parentId = 1, typeId = 1292, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[91] = { parentId = 1, typeId = 1294, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[92] = { parentId = 1, typeId = 1296, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[93] = { parentId = 1, typeId = 1298, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[94] = { parentId = 1, typeId = 1300, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1440, 1442, 1444}, }
_typeInfoList[95] = { parentId = 1, typeId = 1302, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[96] = { parentId = 1, typeId = 1304, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[97] = { parentId = 1, typeId = 1306, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[98] = { parentId = 1, typeId = 1308, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[99] = { parentId = 1, typeId = 1310, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[100] = { parentId = 1, typeId = 1312, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[101] = { parentId = 1, typeId = 1314, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[102] = { parentId = 1, typeId = 1318, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[103] = { parentId = 1, typeId = 1320, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1446, 1448, 1450, 1452, 1454, 1456, 1458, 1460, 1462, 1464, 1466, 1468, 1470, 1472, 1474, 1476, 1478, 1480, 1482, 1484, 1486, 1488, 1490, 1492, 1494, 1496, 1498}, }
_typeInfoList[104] = { parentId = 1, typeId = 1322, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[105] = { parentId = 1, typeId = 1324, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[106] = { parentId = 1, typeId = 1326, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[107] = { parentId = 1, typeId = 1328, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1500, 1502, 1504, 1506, 1508, 1510, 1512, 1514}, }
_typeInfoList[108] = { parentId = 1, typeId = 1330, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1320}, retTypeId = {}, children = {}, }
_typeInfoList[109] = { parentId = 1, typeId = 1332, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1328}, retTypeId = {}, children = {}, }
_typeInfoList[110] = { parentId = 1, typeId = 1334, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[111] = { parentId = 1, typeId = 1336, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1516, 1518, 1520}, }
_typeInfoList[112] = { parentId = 1, typeId = 1338, baseId = 1336, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[113] = { parentId = 1, typeId = 1340, baseId = 1336, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[114] = { parentId = 1, typeId = 1342, baseId = 1336, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[115] = { parentId = 1, typeId = 1344, baseId = 1336, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[116] = { parentId = 1, typeId = 1346, baseId = 1336, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[117] = { parentId = 1, typeId = 1348, baseId = 1336, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[118] = { parentId = 1, typeId = 1350, baseId = 1336, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[119] = { parentId = 1, typeId = 1352, baseId = 1336, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[120] = { parentId = 1, typeId = 1354, baseId = 1336, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[121] = { parentId = 1, typeId = 1356, baseId = 1336, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[122] = { parentId = 1, typeId = 1358, baseId = 1336, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[123] = { parentId = 1, typeId = 1360, baseId = 1336, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[124] = { parentId = 1, typeId = 1362, baseId = 1336, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[125] = { parentId = 1, typeId = 1364, baseId = 1336, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[126] = { parentId = 1, typeId = 1366, baseId = 1336, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[127] = { parentId = 1, typeId = 1368, baseId = 1336, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[128] = { parentId = 1, typeId = 1370, baseId = 1336, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[129] = { parentId = 1, typeId = 1372, baseId = 1336, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[130] = { parentId = 1, typeId = 1374, baseId = 1336, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[131] = { parentId = 1, typeId = 1376, baseId = 1336, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[132] = { parentId = 1, typeId = 1378, baseId = 1336, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[133] = { parentId = 1, typeId = 1380, baseId = 1336, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[134] = { parentId = 1, typeId = 1382, baseId = 1336, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[135] = { parentId = 1, typeId = 1384, baseId = 1336, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[136] = { parentId = 1, typeId = 1386, baseId = 1336, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[137] = { parentId = 1, typeId = 1388, baseId = 1336, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[138] = { parentId = 1, typeId = 1390, baseId = 1336, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[139] = { parentId = 1, typeId = 1392, baseId = 1336, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[140] = { parentId = 1, typeId = 1394, baseId = 1336, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[141] = { parentId = 1, typeId = 1396, baseId = 1336, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[142] = { parentId = 1, typeId = 1398, baseId = 1336, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[143] = { parentId = 1, typeId = 1400, baseId = 1336, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[144] = { parentId = 1, typeId = 1402, baseId = 1336, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[145] = { parentId = 1, typeId = 1404, baseId = 1336, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[146] = { parentId = 1, typeId = 1406, baseId = 1336, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[147] = { parentId = 1, typeId = 1408, baseId = 1336, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[148] = { parentId = 1, typeId = 1410, baseId = 1336, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[149] = { parentId = 1, typeId = 1412, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1522, 1524}, }
_typeInfoList[150] = { parentId = 1, typeId = 1414, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[151] = { parentId = 1, typeId = 1418, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[152] = { parentId = 1, typeId = 1420, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[153] = { parentId = 1, typeId = 1422, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1526, 1528, 1530, 1532, 1534, 1536, 1538, 1540}, }
_typeInfoList[154] = { parentId = 1, typeId = 1424, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1336}, retTypeId = {}, children = {}, }
_typeInfoList[155] = { parentId = 1, typeId = 1426, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[156] = { parentId = 1, typeId = 1428, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[157] = { parentId = 1, typeId = 1430, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1542, 1544}, }
_typeInfoList[158] = { parentId = 1, typeId = 1432, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1336}, retTypeId = {}, children = {}, }
_typeInfoList[159] = { parentId = 1, typeId = 1546, baseId = 1, txt = 'filterObj',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[160] = { parentId = 264, typeId = 282, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[161] = { parentId = 264, typeId = 284, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {264}, children = {}, }
_typeInfoList[162] = { parentId = 264, typeId = 340, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[163] = { parentId = 358, typeId = 372, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[164] = { parentId = 358, typeId = 374, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {358}, children = {}, }
_typeInfoList[165] = { parentId = 358, typeId = 376, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[166] = { parentId = 398, typeId = 418, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[167] = { parentId = 398, typeId = 420, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[168] = { parentId = 398, typeId = 422, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[169] = { parentId = 398, typeId = 424, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[170] = { parentId = 398, typeId = 430, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[171] = { parentId = 398, typeId = 432, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[172] = { parentId = 398, typeId = 434, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[173] = { parentId = 398, typeId = 440, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[174] = { parentId = 398, typeId = 442, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[175] = { parentId = 398, typeId = 446, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[176] = { parentId = 398, typeId = 450, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[177] = { parentId = 398, typeId = 452, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[178] = { parentId = 398, typeId = 456, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[179] = { parentId = 398, typeId = 462, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[180] = { parentId = 398, typeId = 470, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {468}, children = {}, }
_typeInfoList[181] = { parentId = 398, typeId = 474, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {472}, children = {}, }
_typeInfoList[182] = { parentId = 398, typeId = 476, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[183] = { parentId = 398, typeId = 478, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[184] = { parentId = 398, typeId = 480, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[185] = { parentId = 398, typeId = 482, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[186] = { parentId = 398, typeId = 484, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[187] = { parentId = 398, typeId = 486, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[188] = { parentId = 398, typeId = 488, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[189] = { parentId = 398, typeId = 490, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[190] = { parentId = 398, typeId = 492, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[191] = { parentId = 398, typeId = 494, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[192] = { parentId = 398, typeId = 498, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {496}, children = {}, }
_typeInfoList[193] = { parentId = 500, typeId = 516, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[194] = { parentId = 500, typeId = 518, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[195] = { parentId = 500, typeId = 520, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {500}, children = {}, }
_typeInfoList[196] = { parentId = 500, typeId = 522, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[197] = { parentId = 500, typeId = 524, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[198] = { parentId = 500, typeId = 526, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {500}, children = {}, }
_typeInfoList[199] = { parentId = 500, typeId = 530, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {528}, children = {}, }
_typeInfoList[200] = { parentId = 500, typeId = 534, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {532}, children = {}, }
_typeInfoList[201] = { parentId = 540, typeId = 544, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[202] = { parentId = 540, typeId = 546, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[203] = { parentId = 540, typeId = 550, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[204] = { parentId = 628, typeId = 708, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {706}, children = {}, }
_typeInfoList[205] = { parentId = 628, typeId = 1250, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[206] = { parentId = 1094, typeId = 1102, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {356}, children = {}, }
_typeInfoList[207] = { parentId = 1094, typeId = 1104, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {356}, children = {}, }
_typeInfoList[208] = { parentId = 1094, typeId = 1108, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1106}, children = {}, }
_typeInfoList[209] = { parentId = 1094, typeId = 1110, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[210] = { parentId = 1094, typeId = 1112, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[211] = { parentId = 1094, typeId = 1114, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {540}, children = {}, }
_typeInfoList[212] = { parentId = 1094, typeId = 1118, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1116}, children = {}, }
_typeInfoList[213] = { parentId = 1094, typeId = 1122, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1120}, children = {}, }
_typeInfoList[214] = { parentId = 1208, typeId = 1212, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {356}, children = {}, }
_typeInfoList[215] = { parentId = 1208, typeId = 1216, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1214}, children = {}, }
_typeInfoList[216] = { parentId = 1280, typeId = 1434, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[217] = { parentId = 1280, typeId = 1436, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1280}, children = {}, }
_typeInfoList[218] = { parentId = 1280, typeId = 1438, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[219] = { parentId = 1300, typeId = 1440, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[220] = { parentId = 1300, typeId = 1442, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1300}, children = {}, }
_typeInfoList[221] = { parentId = 1300, typeId = 1444, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[222] = { parentId = 1320, typeId = 1446, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[223] = { parentId = 1320, typeId = 1448, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[224] = { parentId = 1320, typeId = 1450, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[225] = { parentId = 1320, typeId = 1452, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[226] = { parentId = 1320, typeId = 1454, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[227] = { parentId = 1320, typeId = 1456, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[228] = { parentId = 1320, typeId = 1458, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[229] = { parentId = 1320, typeId = 1460, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[230] = { parentId = 1320, typeId = 1462, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[231] = { parentId = 1320, typeId = 1464, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[232] = { parentId = 1320, typeId = 1466, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[233] = { parentId = 1320, typeId = 1468, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[234] = { parentId = 1320, typeId = 1470, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[235] = { parentId = 1320, typeId = 1472, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[236] = { parentId = 1320, typeId = 1474, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1322}, children = {}, }
_typeInfoList[237] = { parentId = 1320, typeId = 1476, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1324}, children = {}, }
_typeInfoList[238] = { parentId = 1320, typeId = 1478, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[239] = { parentId = 1320, typeId = 1480, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[240] = { parentId = 1320, typeId = 1482, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[241] = { parentId = 1320, typeId = 1484, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[242] = { parentId = 1320, typeId = 1486, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[243] = { parentId = 1320, typeId = 1488, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[244] = { parentId = 1320, typeId = 1490, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[245] = { parentId = 1320, typeId = 1492, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[246] = { parentId = 1320, typeId = 1494, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[247] = { parentId = 1320, typeId = 1496, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[248] = { parentId = 1320, typeId = 1498, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[249] = { parentId = 1328, typeId = 1500, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[250] = { parentId = 1328, typeId = 1502, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[251] = { parentId = 1328, typeId = 1504, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1328}, children = {}, }
_typeInfoList[252] = { parentId = 1328, typeId = 1506, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[253] = { parentId = 1328, typeId = 1508, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[254] = { parentId = 1328, typeId = 1510, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1328}, children = {}, }
_typeInfoList[255] = { parentId = 1328, typeId = 1512, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1330}, children = {}, }
_typeInfoList[256] = { parentId = 1328, typeId = 1514, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1332}, children = {}, }
_typeInfoList[257] = { parentId = 1336, typeId = 1516, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[258] = { parentId = 1336, typeId = 1518, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[259] = { parentId = 1336, typeId = 1520, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[260] = { parentId = 1412, typeId = 1522, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1414}, children = {}, }
_typeInfoList[261] = { parentId = 1412, typeId = 1524, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[262] = { parentId = 1422, typeId = 1526, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1298}, children = {}, }
_typeInfoList[263] = { parentId = 1422, typeId = 1528, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1298}, children = {}, }
_typeInfoList[264] = { parentId = 1422, typeId = 1530, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1424}, children = {}, }
_typeInfoList[265] = { parentId = 1422, typeId = 1532, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[266] = { parentId = 1422, typeId = 1534, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[267] = { parentId = 1422, typeId = 1536, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1336}, children = {}, }
_typeInfoList[268] = { parentId = 1422, typeId = 1538, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1426}, children = {}, }
_typeInfoList[269] = { parentId = 1422, typeId = 1540, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1428}, children = {}, }
_typeInfoList[270] = { parentId = 1430, typeId = 1542, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1298}, children = {}, }
_typeInfoList[271] = { parentId = 1430, typeId = 1544, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1432}, children = {}, }
----- meta -----
return moduleObj
