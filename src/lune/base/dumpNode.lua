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

filterObj[TransUnit.nodeKind.ExpMacroExp] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.func.expType:getTxt(  ) )
end

filterObj[TransUnit.nodeKind.DeclMacro] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.expType:getTxt(  ) )
  local nodeInfo = node.info
  TransUnit.nodeFilter( nodeInfo:get_ast(  ), self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.ExpMacroStat] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node:get_expType(  ):getTxt(  ) )
  local list = node:get_info(  )
  for __index, node in pairs( list ) do
    TransUnit.nodeFilter( node, self, prefix .. "  ", depth + 1 )
  end
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
_typeInfoList[1] = { parentId = 1, typeId = 94, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[2] = { parentId = 1, typeId = 158, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[3] = { parentId = 1, typeId = 270, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {272}, }
_typeInfoList[4] = { parentId = 1, typeId = 274, baseId = 270, txt = 'TxtStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {276, 278}, }
_typeInfoList[5] = { parentId = 1, typeId = 280, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[6] = { parentId = 1, typeId = 282, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[7] = { parentId = 1, typeId = 284, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {286, 288}, }
_typeInfoList[8] = { parentId = 1, typeId = 290, baseId = 1, txt = 'WrapParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {292, 294}, }
_typeInfoList[9] = { parentId = 1, typeId = 296, baseId = 284, txt = 'StreamParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {314, 316, 366}, }
_typeInfoList[10] = { parentId = 1, typeId = 342, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[11] = { parentId = 1, typeId = 344, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[12] = { parentId = 1, typeId = 346, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[13] = { parentId = 1, typeId = 374, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[14] = { parentId = 1, typeId = 376, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[15] = { parentId = 1, typeId = 378, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[16] = { parentId = 1, typeId = 380, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {404}, }
_typeInfoList[17] = { parentId = 1, typeId = 382, baseId = 380, txt = 'TxtStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {406, 408}, }
_typeInfoList[18] = { parentId = 1, typeId = 384, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[19] = { parentId = 1, typeId = 386, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[20] = { parentId = 1, typeId = 388, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {410, 412}, }
_typeInfoList[21] = { parentId = 1, typeId = 390, baseId = 1, txt = 'WrapParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {414, 416}, }
_typeInfoList[22] = { parentId = 1, typeId = 392, baseId = 388, txt = 'StreamParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {418, 420, 422}, }
_typeInfoList[23] = { parentId = 1, typeId = 396, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[24] = { parentId = 1, typeId = 398, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[25] = { parentId = 1, typeId = 400, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[26] = { parentId = 1, typeId = 402, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[27] = { parentId = 1, typeId = 424, baseId = 1, txt = 'Util',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[28] = { parentId = 1, typeId = 488, baseId = 1, txt = 'outStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {490}, }
_typeInfoList[29] = { parentId = 1, typeId = 492, baseId = 488, txt = 'memStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {494, 496, 498}, }
_typeInfoList[30] = { parentId = 1, typeId = 500, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[31] = { parentId = 1, typeId = 502, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[32] = { parentId = 1, typeId = 504, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {558}, }
_typeInfoList[33] = { parentId = 1, typeId = 506, baseId = 504, txt = 'TxtStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {560, 562}, }
_typeInfoList[34] = { parentId = 1, typeId = 508, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[35] = { parentId = 1, typeId = 510, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[36] = { parentId = 1, typeId = 512, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {564, 566}, }
_typeInfoList[37] = { parentId = 1, typeId = 514, baseId = 1, txt = 'WrapParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {568, 570}, }
_typeInfoList[38] = { parentId = 1, typeId = 516, baseId = 512, txt = 'StreamParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {572, 574, 576}, }
_typeInfoList[39] = { parentId = 1, typeId = 518, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[40] = { parentId = 1, typeId = 520, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[41] = { parentId = 1, typeId = 522, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[42] = { parentId = 1, typeId = 524, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[43] = { parentId = 1, typeId = 526, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[44] = { parentId = 1, typeId = 528, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[45] = { parentId = 1, typeId = 530, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {578}, }
_typeInfoList[46] = { parentId = 1, typeId = 532, baseId = 530, txt = 'TxtStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {580, 582}, }
_typeInfoList[47] = { parentId = 1, typeId = 534, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[48] = { parentId = 1, typeId = 536, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[49] = { parentId = 1, typeId = 538, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {584, 586}, }
_typeInfoList[50] = { parentId = 1, typeId = 540, baseId = 1, txt = 'WrapParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {588, 590}, }
_typeInfoList[51] = { parentId = 1, typeId = 542, baseId = 538, txt = 'StreamParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {592, 594, 596}, }
_typeInfoList[52] = { parentId = 1, typeId = 544, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[53] = { parentId = 1, typeId = 546, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[54] = { parentId = 1, typeId = 548, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[55] = { parentId = 1, typeId = 550, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[56] = { parentId = 1, typeId = 552, baseId = 1, txt = 'Util',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[57] = { parentId = 1, typeId = 554, baseId = 1, txt = 'outStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {598}, }
_typeInfoList[58] = { parentId = 1, typeId = 556, baseId = 554, txt = 'memStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {600, 602, 604}, }
_typeInfoList[59] = { parentId = 1, typeId = 606, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[60] = { parentId = 1, typeId = 608, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[61] = { parentId = 1, typeId = 610, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {648, 650, 652, 654, 660, 662, 664, 670, 672, 676, 680, 682, 686, 692, 700, 704, 706, 708, 710, 712, 714, 716, 718, 720, 722, 724, 728}, }
_typeInfoList[62] = { parentId = 1, typeId = 624, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[63] = { parentId = 1, typeId = 698, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {610}, retTypeId = {}, children = {}, }
_typeInfoList[64] = { parentId = 1, typeId = 702, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {610}, retTypeId = {}, children = {}, }
_typeInfoList[65] = { parentId = 1, typeId = 726, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {610}, retTypeId = {}, children = {}, }
_typeInfoList[66] = { parentId = 1, typeId = 730, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {746, 748, 750, 752, 754, 756, 758, 762, 766}, }
_typeInfoList[67] = { parentId = 1, typeId = 760, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 610}, retTypeId = {}, children = {}, }
_typeInfoList[68] = { parentId = 1, typeId = 764, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 730}, retTypeId = {}, children = {}, }
_typeInfoList[69] = { parentId = 1, typeId = 770, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[70] = { parentId = 1, typeId = 772, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {776, 778, 782}, }
_typeInfoList[71] = { parentId = 1, typeId = 784, baseId = 772, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[72] = { parentId = 1, typeId = 786, baseId = 772, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[73] = { parentId = 1, typeId = 788, baseId = 772, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[74] = { parentId = 1, typeId = 790, baseId = 772, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[75] = { parentId = 1, typeId = 792, baseId = 772, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[76] = { parentId = 1, typeId = 794, baseId = 772, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[77] = { parentId = 1, typeId = 796, baseId = 772, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[78] = { parentId = 1, typeId = 798, baseId = 772, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[79] = { parentId = 1, typeId = 800, baseId = 772, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[80] = { parentId = 1, typeId = 802, baseId = 772, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[81] = { parentId = 1, typeId = 804, baseId = 772, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[82] = { parentId = 1, typeId = 806, baseId = 772, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[83] = { parentId = 1, typeId = 808, baseId = 772, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[84] = { parentId = 1, typeId = 810, baseId = 772, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[85] = { parentId = 1, typeId = 812, baseId = 772, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[86] = { parentId = 1, typeId = 814, baseId = 772, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[87] = { parentId = 1, typeId = 816, baseId = 772, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[88] = { parentId = 1, typeId = 818, baseId = 772, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[89] = { parentId = 1, typeId = 820, baseId = 772, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[90] = { parentId = 1, typeId = 822, baseId = 772, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[91] = { parentId = 1, typeId = 824, baseId = 772, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[92] = { parentId = 1, typeId = 826, baseId = 772, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[93] = { parentId = 1, typeId = 828, baseId = 772, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[94] = { parentId = 1, typeId = 830, baseId = 772, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[95] = { parentId = 1, typeId = 832, baseId = 772, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[96] = { parentId = 1, typeId = 834, baseId = 772, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[97] = { parentId = 1, typeId = 836, baseId = 772, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[98] = { parentId = 1, typeId = 838, baseId = 772, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[99] = { parentId = 1, typeId = 840, baseId = 772, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[100] = { parentId = 1, typeId = 842, baseId = 772, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[101] = { parentId = 1, typeId = 844, baseId = 772, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[102] = { parentId = 1, typeId = 846, baseId = 772, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[103] = { parentId = 1, typeId = 848, baseId = 772, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[104] = { parentId = 1, typeId = 850, baseId = 772, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[105] = { parentId = 1, typeId = 852, baseId = 772, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[106] = { parentId = 1, typeId = 854, baseId = 772, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[107] = { parentId = 1, typeId = 856, baseId = 772, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[108] = { parentId = 1, typeId = 860, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {942, 1310}, }
_typeInfoList[109] = { parentId = 1, typeId = 940, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[110] = { parentId = 1, typeId = 996, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[111] = { parentId = 1, typeId = 998, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[112] = { parentId = 1, typeId = 1324, baseId = 1, txt = 'DeclMacroInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1330, 1334, 1336, 1340}, }
_typeInfoList[113] = { parentId = 1, typeId = 1332, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {772}, retTypeId = {}, children = {}, }
_typeInfoList[114] = { parentId = 1, typeId = 1338, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {386}, retTypeId = {}, children = {}, }
_typeInfoList[115] = { parentId = 1, typeId = 1388, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1396, 1398, 1402, 1404, 1406, 1408, 1412, 1416}, }
_typeInfoList[116] = { parentId = 1, typeId = 1400, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {772}, retTypeId = {}, children = {}, }
_typeInfoList[117] = { parentId = 1, typeId = 1410, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {610}, retTypeId = {}, children = {}, }
_typeInfoList[118] = { parentId = 1, typeId = 1414, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {610}, retTypeId = {}, children = {}, }
_typeInfoList[119] = { parentId = 1, typeId = 1498, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1502, 1506}, }
_typeInfoList[120] = { parentId = 1, typeId = 1504, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {772}, retTypeId = {}, children = {}, }
_typeInfoList[121] = { parentId = 1, typeId = 1564, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[122] = { parentId = 1, typeId = 1566, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[123] = { parentId = 1, typeId = 1568, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1808}, }
_typeInfoList[124] = { parentId = 1, typeId = 1570, baseId = 1568, txt = 'TxtStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1810, 1812}, }
_typeInfoList[125] = { parentId = 1, typeId = 1572, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[126] = { parentId = 1, typeId = 1574, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[127] = { parentId = 1, typeId = 1576, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1814, 1816}, }
_typeInfoList[128] = { parentId = 1, typeId = 1578, baseId = 1, txt = 'WrapParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1818, 1820}, }
_typeInfoList[129] = { parentId = 1, typeId = 1580, baseId = 1576, txt = 'StreamParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1822, 1824, 1826}, }
_typeInfoList[130] = { parentId = 1, typeId = 1582, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[131] = { parentId = 1, typeId = 1584, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[132] = { parentId = 1, typeId = 1586, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[133] = { parentId = 1, typeId = 1588, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[134] = { parentId = 1, typeId = 1590, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[135] = { parentId = 1, typeId = 1592, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[136] = { parentId = 1, typeId = 1594, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1828}, }
_typeInfoList[137] = { parentId = 1, typeId = 1596, baseId = 1594, txt = 'TxtStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1830, 1832}, }
_typeInfoList[138] = { parentId = 1, typeId = 1598, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[139] = { parentId = 1, typeId = 1600, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[140] = { parentId = 1, typeId = 1602, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1834, 1836}, }
_typeInfoList[141] = { parentId = 1, typeId = 1604, baseId = 1, txt = 'WrapParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1838, 1840}, }
_typeInfoList[142] = { parentId = 1, typeId = 1606, baseId = 1602, txt = 'StreamParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1842, 1844, 1846}, }
_typeInfoList[143] = { parentId = 1, typeId = 1608, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[144] = { parentId = 1, typeId = 1610, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[145] = { parentId = 1, typeId = 1612, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[146] = { parentId = 1, typeId = 1614, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[147] = { parentId = 1, typeId = 1616, baseId = 1, txt = 'Util',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[148] = { parentId = 1, typeId = 1618, baseId = 1, txt = 'outStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1848}, }
_typeInfoList[149] = { parentId = 1, typeId = 1620, baseId = 1618, txt = 'memStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1850, 1852, 1854}, }
_typeInfoList[150] = { parentId = 1, typeId = 1622, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[151] = { parentId = 1, typeId = 1624, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[152] = { parentId = 1, typeId = 1626, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1856}, }
_typeInfoList[153] = { parentId = 1, typeId = 1628, baseId = 1626, txt = 'TxtStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1858, 1860}, }
_typeInfoList[154] = { parentId = 1, typeId = 1630, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[155] = { parentId = 1, typeId = 1632, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[156] = { parentId = 1, typeId = 1634, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1862, 1864}, }
_typeInfoList[157] = { parentId = 1, typeId = 1636, baseId = 1, txt = 'WrapParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1866, 1868}, }
_typeInfoList[158] = { parentId = 1, typeId = 1638, baseId = 1634, txt = 'StreamParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1870, 1872, 1874}, }
_typeInfoList[159] = { parentId = 1, typeId = 1640, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[160] = { parentId = 1, typeId = 1642, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[161] = { parentId = 1, typeId = 1644, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[162] = { parentId = 1, typeId = 1646, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[163] = { parentId = 1, typeId = 1648, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[164] = { parentId = 1, typeId = 1650, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[165] = { parentId = 1, typeId = 1652, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1876}, }
_typeInfoList[166] = { parentId = 1, typeId = 1654, baseId = 1652, txt = 'TxtStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1878, 1880}, }
_typeInfoList[167] = { parentId = 1, typeId = 1656, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[168] = { parentId = 1, typeId = 1658, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[169] = { parentId = 1, typeId = 1660, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1882, 1884}, }
_typeInfoList[170] = { parentId = 1, typeId = 1662, baseId = 1, txt = 'WrapParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1886, 1888}, }
_typeInfoList[171] = { parentId = 1, typeId = 1664, baseId = 1660, txt = 'StreamParser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1890, 1892, 1894}, }
_typeInfoList[172] = { parentId = 1, typeId = 1666, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[173] = { parentId = 1, typeId = 1668, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[174] = { parentId = 1, typeId = 1670, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[175] = { parentId = 1, typeId = 1672, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[176] = { parentId = 1, typeId = 1674, baseId = 1, txt = 'Util',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[177] = { parentId = 1, typeId = 1676, baseId = 1, txt = 'outStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1896}, }
_typeInfoList[178] = { parentId = 1, typeId = 1678, baseId = 1676, txt = 'memStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1898, 1900, 1902}, }
_typeInfoList[179] = { parentId = 1, typeId = 1680, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[180] = { parentId = 1, typeId = 1682, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[181] = { parentId = 1, typeId = 1684, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1904, 1906, 1908, 1910, 1912, 1914, 1916, 1918, 1920, 1922, 1924, 1926, 1928, 1930, 1932, 1934, 1936, 1938, 1940, 1942, 1944, 1946, 1948, 1950, 1952, 1954, 1956}, }
_typeInfoList[182] = { parentId = 1, typeId = 1688, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[183] = { parentId = 1, typeId = 1690, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1684}, retTypeId = {}, children = {}, }
_typeInfoList[184] = { parentId = 1, typeId = 1692, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1684}, retTypeId = {}, children = {}, }
_typeInfoList[185] = { parentId = 1, typeId = 1694, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1684}, retTypeId = {}, children = {}, }
_typeInfoList[186] = { parentId = 1, typeId = 1696, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1958, 1960, 1962, 1964, 1966, 1968, 1970, 1972, 1974}, }
_typeInfoList[187] = { parentId = 1, typeId = 1698, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1684}, retTypeId = {}, children = {}, }
_typeInfoList[188] = { parentId = 1, typeId = 1700, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1696}, retTypeId = {}, children = {}, }
_typeInfoList[189] = { parentId = 1, typeId = 1702, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[190] = { parentId = 1, typeId = 1704, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1976, 1978, 1980}, }
_typeInfoList[191] = { parentId = 1, typeId = 1706, baseId = 1704, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[192] = { parentId = 1, typeId = 1708, baseId = 1704, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[193] = { parentId = 1, typeId = 1710, baseId = 1704, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[194] = { parentId = 1, typeId = 1712, baseId = 1704, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[195] = { parentId = 1, typeId = 1714, baseId = 1704, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[196] = { parentId = 1, typeId = 1716, baseId = 1704, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[197] = { parentId = 1, typeId = 1718, baseId = 1704, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[198] = { parentId = 1, typeId = 1720, baseId = 1704, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[199] = { parentId = 1, typeId = 1722, baseId = 1704, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[200] = { parentId = 1, typeId = 1724, baseId = 1704, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[201] = { parentId = 1, typeId = 1726, baseId = 1704, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[202] = { parentId = 1, typeId = 1728, baseId = 1704, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[203] = { parentId = 1, typeId = 1730, baseId = 1704, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[204] = { parentId = 1, typeId = 1732, baseId = 1704, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[205] = { parentId = 1, typeId = 1734, baseId = 1704, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[206] = { parentId = 1, typeId = 1736, baseId = 1704, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[207] = { parentId = 1, typeId = 1738, baseId = 1704, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[208] = { parentId = 1, typeId = 1740, baseId = 1704, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[209] = { parentId = 1, typeId = 1742, baseId = 1704, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[210] = { parentId = 1, typeId = 1744, baseId = 1704, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[211] = { parentId = 1, typeId = 1746, baseId = 1704, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[212] = { parentId = 1, typeId = 1748, baseId = 1704, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[213] = { parentId = 1, typeId = 1750, baseId = 1704, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[214] = { parentId = 1, typeId = 1752, baseId = 1704, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[215] = { parentId = 1, typeId = 1754, baseId = 1704, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[216] = { parentId = 1, typeId = 1756, baseId = 1704, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[217] = { parentId = 1, typeId = 1758, baseId = 1704, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[218] = { parentId = 1, typeId = 1760, baseId = 1704, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[219] = { parentId = 1, typeId = 1762, baseId = 1704, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[220] = { parentId = 1, typeId = 1764, baseId = 1704, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[221] = { parentId = 1, typeId = 1766, baseId = 1704, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[222] = { parentId = 1, typeId = 1768, baseId = 1704, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[223] = { parentId = 1, typeId = 1770, baseId = 1704, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[224] = { parentId = 1, typeId = 1772, baseId = 1704, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[225] = { parentId = 1, typeId = 1774, baseId = 1704, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[226] = { parentId = 1, typeId = 1776, baseId = 1704, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[227] = { parentId = 1, typeId = 1778, baseId = 1704, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[228] = { parentId = 1, typeId = 1780, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1982, 1984}, }
_typeInfoList[229] = { parentId = 1, typeId = 1782, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[230] = { parentId = 1, typeId = 1786, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[231] = { parentId = 1, typeId = 1788, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[232] = { parentId = 1, typeId = 1790, baseId = 1, txt = 'DeclMacroInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1986, 1988, 1990, 1992}, }
_typeInfoList[233] = { parentId = 1, typeId = 1792, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1704}, retTypeId = {}, children = {}, }
_typeInfoList[234] = { parentId = 1, typeId = 1794, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1600}, retTypeId = {}, children = {}, }
_typeInfoList[235] = { parentId = 1, typeId = 1796, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1994, 1996, 1998, 2000, 2002, 2004, 2006, 2008}, }
_typeInfoList[236] = { parentId = 1, typeId = 1798, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1704}, retTypeId = {}, children = {}, }
_typeInfoList[237] = { parentId = 1, typeId = 1800, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1684}, retTypeId = {}, children = {}, }
_typeInfoList[238] = { parentId = 1, typeId = 1802, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1684}, retTypeId = {}, children = {}, }
_typeInfoList[239] = { parentId = 1, typeId = 1804, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {2010, 2012}, }
_typeInfoList[240] = { parentId = 1, typeId = 1806, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1704}, retTypeId = {}, children = {}, }
_typeInfoList[241] = { parentId = 1, typeId = 2014, baseId = 1, txt = 'filterObj',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[242] = { parentId = 270, typeId = 272, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[243] = { parentId = 274, typeId = 276, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[244] = { parentId = 274, typeId = 278, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[245] = { parentId = 284, typeId = 286, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {282}, children = {}, }
_typeInfoList[246] = { parentId = 284, typeId = 288, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[247] = { parentId = 290, typeId = 292, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {282}, children = {}, }
_typeInfoList[248] = { parentId = 290, typeId = 294, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[249] = { parentId = 296, typeId = 314, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[250] = { parentId = 296, typeId = 316, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {296}, children = {}, }
_typeInfoList[251] = { parentId = 296, typeId = 366, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {282}, children = {}, }
_typeInfoList[252] = { parentId = 380, typeId = 404, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[253] = { parentId = 382, typeId = 406, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[254] = { parentId = 382, typeId = 408, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[255] = { parentId = 388, typeId = 410, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
_typeInfoList[256] = { parentId = 388, typeId = 412, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[257] = { parentId = 390, typeId = 414, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
_typeInfoList[258] = { parentId = 390, typeId = 416, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[259] = { parentId = 392, typeId = 418, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[260] = { parentId = 392, typeId = 420, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {392}, children = {}, }
_typeInfoList[261] = { parentId = 392, typeId = 422, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
_typeInfoList[262] = { parentId = 488, typeId = 490, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[263] = { parentId = 492, typeId = 494, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[264] = { parentId = 492, typeId = 496, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[265] = { parentId = 492, typeId = 498, baseId = 1, txt = 'get_txt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[266] = { parentId = 504, typeId = 558, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[267] = { parentId = 506, typeId = 560, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[268] = { parentId = 506, typeId = 562, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[269] = { parentId = 512, typeId = 564, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {510}, children = {}, }
_typeInfoList[270] = { parentId = 512, typeId = 566, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[271] = { parentId = 514, typeId = 568, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {510}, children = {}, }
_typeInfoList[272] = { parentId = 514, typeId = 570, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[273] = { parentId = 516, typeId = 572, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[274] = { parentId = 516, typeId = 574, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[275] = { parentId = 516, typeId = 576, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {510}, children = {}, }
_typeInfoList[276] = { parentId = 530, typeId = 578, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[277] = { parentId = 532, typeId = 580, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[278] = { parentId = 532, typeId = 582, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[279] = { parentId = 538, typeId = 584, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[280] = { parentId = 538, typeId = 586, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[281] = { parentId = 540, typeId = 588, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[282] = { parentId = 540, typeId = 590, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[283] = { parentId = 542, typeId = 592, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[284] = { parentId = 542, typeId = 594, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {542}, children = {}, }
_typeInfoList[285] = { parentId = 542, typeId = 596, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[286] = { parentId = 554, typeId = 598, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[287] = { parentId = 556, typeId = 600, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[288] = { parentId = 556, typeId = 602, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[289] = { parentId = 556, typeId = 604, baseId = 1, txt = 'get_txt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[290] = { parentId = 610, typeId = 648, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[291] = { parentId = 610, typeId = 650, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[292] = { parentId = 610, typeId = 652, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[293] = { parentId = 610, typeId = 654, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[294] = { parentId = 610, typeId = 660, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[295] = { parentId = 610, typeId = 662, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[296] = { parentId = 610, typeId = 664, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[297] = { parentId = 610, typeId = 670, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[298] = { parentId = 610, typeId = 672, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[299] = { parentId = 610, typeId = 676, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[300] = { parentId = 610, typeId = 680, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[301] = { parentId = 610, typeId = 682, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[302] = { parentId = 610, typeId = 686, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[303] = { parentId = 610, typeId = 692, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[304] = { parentId = 610, typeId = 700, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {698}, children = {}, }
_typeInfoList[305] = { parentId = 610, typeId = 704, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {702}, children = {}, }
_typeInfoList[306] = { parentId = 610, typeId = 706, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[307] = { parentId = 610, typeId = 708, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[308] = { parentId = 610, typeId = 710, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[309] = { parentId = 610, typeId = 712, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[310] = { parentId = 610, typeId = 714, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[311] = { parentId = 610, typeId = 716, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[312] = { parentId = 610, typeId = 718, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[313] = { parentId = 610, typeId = 720, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[314] = { parentId = 610, typeId = 722, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[315] = { parentId = 610, typeId = 724, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[316] = { parentId = 610, typeId = 728, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {726}, children = {}, }
_typeInfoList[317] = { parentId = 730, typeId = 746, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[318] = { parentId = 730, typeId = 748, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[319] = { parentId = 730, typeId = 750, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {730}, children = {}, }
_typeInfoList[320] = { parentId = 730, typeId = 752, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[321] = { parentId = 730, typeId = 754, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[322] = { parentId = 730, typeId = 756, baseId = 1, txt = 'getTypeInfoMethod',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[323] = { parentId = 730, typeId = 758, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {730}, children = {}, }
_typeInfoList[324] = { parentId = 730, typeId = 762, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {760}, children = {}, }
_typeInfoList[325] = { parentId = 730, typeId = 766, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {764}, children = {}, }
_typeInfoList[326] = { parentId = 772, typeId = 776, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[327] = { parentId = 772, typeId = 778, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[328] = { parentId = 772, typeId = 782, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[329] = { parentId = 860, typeId = 942, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {940}, children = {}, }
_typeInfoList[330] = { parentId = 860, typeId = 1310, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[331] = { parentId = 1324, typeId = 1330, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
_typeInfoList[332] = { parentId = 1324, typeId = 1334, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1332}, children = {}, }
_typeInfoList[333] = { parentId = 1324, typeId = 1336, baseId = 1, txt = 'get_ast',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {772}, children = {}, }
_typeInfoList[334] = { parentId = 1324, typeId = 1340, baseId = 1, txt = 'get_tokenList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1338}, children = {}, }
_typeInfoList[335] = { parentId = 1388, typeId = 1396, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
_typeInfoList[336] = { parentId = 1388, typeId = 1398, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
_typeInfoList[337] = { parentId = 1388, typeId = 1402, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1400}, children = {}, }
_typeInfoList[338] = { parentId = 1388, typeId = 1404, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[339] = { parentId = 1388, typeId = 1406, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[340] = { parentId = 1388, typeId = 1408, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {772}, children = {}, }
_typeInfoList[341] = { parentId = 1388, typeId = 1412, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1410}, children = {}, }
_typeInfoList[342] = { parentId = 1388, typeId = 1416, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1414}, children = {}, }
_typeInfoList[343] = { parentId = 1498, typeId = 1502, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
_typeInfoList[344] = { parentId = 1498, typeId = 1506, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1504}, children = {}, }
_typeInfoList[345] = { parentId = 1568, typeId = 1808, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[346] = { parentId = 1570, typeId = 1810, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[347] = { parentId = 1570, typeId = 1812, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[348] = { parentId = 1576, typeId = 1814, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1574}, children = {}, }
_typeInfoList[349] = { parentId = 1576, typeId = 1816, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[350] = { parentId = 1578, typeId = 1818, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1574}, children = {}, }
_typeInfoList[351] = { parentId = 1578, typeId = 1820, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[352] = { parentId = 1580, typeId = 1822, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[353] = { parentId = 1580, typeId = 1824, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1580}, children = {}, }
_typeInfoList[354] = { parentId = 1580, typeId = 1826, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1574}, children = {}, }
_typeInfoList[355] = { parentId = 1594, typeId = 1828, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[356] = { parentId = 1596, typeId = 1830, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[357] = { parentId = 1596, typeId = 1832, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[358] = { parentId = 1602, typeId = 1834, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1600}, children = {}, }
_typeInfoList[359] = { parentId = 1602, typeId = 1836, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[360] = { parentId = 1604, typeId = 1838, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1600}, children = {}, }
_typeInfoList[361] = { parentId = 1604, typeId = 1840, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[362] = { parentId = 1606, typeId = 1842, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[363] = { parentId = 1606, typeId = 1844, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1606}, children = {}, }
_typeInfoList[364] = { parentId = 1606, typeId = 1846, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1600}, children = {}, }
_typeInfoList[365] = { parentId = 1618, typeId = 1848, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[366] = { parentId = 1620, typeId = 1850, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[367] = { parentId = 1620, typeId = 1852, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[368] = { parentId = 1620, typeId = 1854, baseId = 1, txt = 'get_txt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[369] = { parentId = 1626, typeId = 1856, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[370] = { parentId = 1628, typeId = 1858, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[371] = { parentId = 1628, typeId = 1860, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[372] = { parentId = 1634, typeId = 1862, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1632}, children = {}, }
_typeInfoList[373] = { parentId = 1634, typeId = 1864, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[374] = { parentId = 1636, typeId = 1866, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1632}, children = {}, }
_typeInfoList[375] = { parentId = 1636, typeId = 1868, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[376] = { parentId = 1638, typeId = 1870, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[377] = { parentId = 1638, typeId = 1872, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1638}, children = {}, }
_typeInfoList[378] = { parentId = 1638, typeId = 1874, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1632}, children = {}, }
_typeInfoList[379] = { parentId = 1652, typeId = 1876, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[380] = { parentId = 1654, typeId = 1878, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[381] = { parentId = 1654, typeId = 1880, baseId = 1, txt = 'read',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[382] = { parentId = 1660, typeId = 1882, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1658}, children = {}, }
_typeInfoList[383] = { parentId = 1660, typeId = 1884, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[384] = { parentId = 1662, typeId = 1886, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1658}, children = {}, }
_typeInfoList[385] = { parentId = 1662, typeId = 1888, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[386] = { parentId = 1664, typeId = 1890, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[387] = { parentId = 1664, typeId = 1892, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1664}, children = {}, }
_typeInfoList[388] = { parentId = 1664, typeId = 1894, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1658}, children = {}, }
_typeInfoList[389] = { parentId = 1676, typeId = 1896, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[390] = { parentId = 1678, typeId = 1898, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[391] = { parentId = 1678, typeId = 1900, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[392] = { parentId = 1678, typeId = 1902, baseId = 1, txt = 'get_txt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[393] = { parentId = 1684, typeId = 1904, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[394] = { parentId = 1684, typeId = 1906, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[395] = { parentId = 1684, typeId = 1908, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[396] = { parentId = 1684, typeId = 1910, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[397] = { parentId = 1684, typeId = 1912, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[398] = { parentId = 1684, typeId = 1914, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[399] = { parentId = 1684, typeId = 1916, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[400] = { parentId = 1684, typeId = 1918, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[401] = { parentId = 1684, typeId = 1920, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[402] = { parentId = 1684, typeId = 1922, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[403] = { parentId = 1684, typeId = 1924, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[404] = { parentId = 1684, typeId = 1926, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[405] = { parentId = 1684, typeId = 1928, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[406] = { parentId = 1684, typeId = 1930, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[407] = { parentId = 1684, typeId = 1932, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1690}, children = {}, }
_typeInfoList[408] = { parentId = 1684, typeId = 1934, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1692}, children = {}, }
_typeInfoList[409] = { parentId = 1684, typeId = 1936, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[410] = { parentId = 1684, typeId = 1938, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[411] = { parentId = 1684, typeId = 1940, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[412] = { parentId = 1684, typeId = 1942, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[413] = { parentId = 1684, typeId = 1944, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[414] = { parentId = 1684, typeId = 1946, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[415] = { parentId = 1684, typeId = 1948, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[416] = { parentId = 1684, typeId = 1950, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[417] = { parentId = 1684, typeId = 1952, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[418] = { parentId = 1684, typeId = 1954, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[419] = { parentId = 1684, typeId = 1956, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1694}, children = {}, }
_typeInfoList[420] = { parentId = 1696, typeId = 1958, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[421] = { parentId = 1696, typeId = 1960, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[422] = { parentId = 1696, typeId = 1962, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1696}, children = {}, }
_typeInfoList[423] = { parentId = 1696, typeId = 1964, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[424] = { parentId = 1696, typeId = 1966, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[425] = { parentId = 1696, typeId = 1968, baseId = 1, txt = 'getTypeInfoMethod',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[426] = { parentId = 1696, typeId = 1970, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1696}, children = {}, }
_typeInfoList[427] = { parentId = 1696, typeId = 1972, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1698}, children = {}, }
_typeInfoList[428] = { parentId = 1696, typeId = 1974, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1700}, children = {}, }
_typeInfoList[429] = { parentId = 1704, typeId = 1976, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[430] = { parentId = 1704, typeId = 1978, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1684}, children = {}, }
_typeInfoList[431] = { parentId = 1704, typeId = 1980, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[432] = { parentId = 1780, typeId = 1982, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1782}, children = {}, }
_typeInfoList[433] = { parentId = 1780, typeId = 1984, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[434] = { parentId = 1790, typeId = 1986, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1600}, children = {}, }
_typeInfoList[435] = { parentId = 1790, typeId = 1988, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1792}, children = {}, }
_typeInfoList[436] = { parentId = 1790, typeId = 1990, baseId = 1, txt = 'get_ast',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1704}, children = {}, }
_typeInfoList[437] = { parentId = 1790, typeId = 1992, baseId = 1, txt = 'get_tokenList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1794}, children = {}, }
_typeInfoList[438] = { parentId = 1796, typeId = 1994, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1600}, children = {}, }
_typeInfoList[439] = { parentId = 1796, typeId = 1996, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1600}, children = {}, }
_typeInfoList[440] = { parentId = 1796, typeId = 1998, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1798}, children = {}, }
_typeInfoList[441] = { parentId = 1796, typeId = 2000, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[442] = { parentId = 1796, typeId = 2002, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[443] = { parentId = 1796, typeId = 2004, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1704}, children = {}, }
_typeInfoList[444] = { parentId = 1796, typeId = 2006, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1800}, children = {}, }
_typeInfoList[445] = { parentId = 1796, typeId = 2008, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1802}, children = {}, }
_typeInfoList[446] = { parentId = 1804, typeId = 2010, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1600}, children = {}, }
_typeInfoList[447] = { parentId = 1804, typeId = 2012, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1806}, children = {}, }
----- meta -----
return moduleObj
