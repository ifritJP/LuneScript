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
  local expType = node:get_expType(  )
  if expType and expType ~= TransUnit.typeInfoKind.None then
    typeStr = string.format( "(%d:%s:%s)", expType:get_typeId(  ), expType:getTxt(  ), expType:get_kind(  ))
  end
  print( string.format( "%s: %s %s %s", prefix, TransUnit.getNodeKindName( node:get_kind(  ) ), txt, typeStr) )
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
  TransUnit.nodeFilter( node.exp, self, prefix .. "  ", depth + 1 )
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
  dump( prefix, depth, node, "" )
  local stmtList = node:get_stmtList(  )
  if stmtList then
    for __index, stmt in pairs( stmtList ) do
      TransUnit.nodeFilter( stmt, self, prefix .. "  ", depth + 1 )
    end
  end
end

filterObj[TransUnit.nodeKind.DeclMacro] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node:get_expType(  ):getTxt(  ) )
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
  local varList = node.varList
  for index, var in pairs( varList ) do
    if index > 1 then
      varName = varName .. ","
    end
    varName = string.format( "%s %s", varName, var["name"].txt)
  end
  dump( prefix, depth, node, node.unwrap and "! " or " " .. varName )
  for index, var in pairs( varList ) do
    if var["refType"] then
      TransUnit.nodeFilter( var["refType"], self, prefix .. "  ", depth + 1 )
    end
  end
  if node.expList then
    TransUnit.nodeFilter( node.expList, self, prefix .. "  ", depth + 1 )
  end
  if node.unwrap then
    TransUnit.nodeFilter( node.unwrap, self, prefix .. "  ", depth + 1 )
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
  local name = node.declInfo.name
  dump( prefix, depth, node, name and name.txt or "<anonymous>" )
  local argList = node.declInfo.argList
  for index, arg in pairs( argList ) do
    TransUnit.nodeFilter( arg, self, prefix .. "  ", depth + 1 )
  end
  local retTypeList = node.declInfo.retTypeList
  for index, refType in pairs( retTypeList ) do
    TransUnit.nodeFilter( refType, self, prefix .. "  ", depth + 1 )
  end
  if node.declInfo.body then
    TransUnit.nodeFilter( node.declInfo.body, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.DeclMethod] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.declInfo.name.txt )
  local argList = node.declInfo.argList
  for index, arg in pairs( argList ) do
    TransUnit.nodeFilter( arg, self, prefix .. "  ", depth + 1 )
  end
  local retTypeList = node.declInfo.retTypeList
  for index, refType in pairs( retTypeList ) do
    TransUnit.nodeFilter( refType, self, prefix .. "  ", depth + 1 )
  end
  TransUnit.nodeFilter( node.declInfo.body, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.DeclConstr] = function ( self, node, prefix, depth )
  filterObj[TransUnit.nodeKind.DeclMethod]( self, node, prefix, depth )
end

filterObj[TransUnit.nodeKind.ExpCallSuper] = function ( self, node, prefix, depth )
  local typeInfo = node.info.superType
  dump( prefix, depth, node, typeInfo:getTxt(  ) )
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
  TransUnit.nodeFilter( node.func, self, prefix .. "  ", depth + 1 )
  if node.argList then
    TransUnit.nodeFilter( node.argList, self, prefix .. "  ", depth + 1 )
  end
end

filterObj[TransUnit.nodeKind.ExpList] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  local expList = node.expList
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
  TransUnit.nodeFilter( node.exp, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.ExpOp2] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.op.txt )
  TransUnit.nodeFilter( node.exp1, self, prefix .. "  ", depth + 1 )
  TransUnit.nodeFilter( node.exp2, self, prefix .. "  ", depth + 1 )
end

filterObj[TransUnit.nodeKind.ExpNew] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, "" )
  TransUnit.nodeFilter( node.symbol, self, prefix .. "  ", depth + 1 )
  if node.argList then
    TransUnit.nodeFilter( node.argList, self, prefix .. "  ", depth + 1 )
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
  if node.expList then
    TransUnit.nodeFilter( node.expList, self, prefix .. "  ", depth + 1 )
  end
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

filterObj[TransUnit.nodeKind.LiteralSymbol] = function ( self, node, prefix, depth )
  dump( prefix, depth, node, node.info.txt )
end

----- meta -----
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfofilterObj = {}
  _className2InfoMap.filterObj = _classInfofilterObj
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
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {106, 3132}, }
_typeInfoList[4] = { parentId = 104, typeId = 106, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {108}, }
_typeInfoList[5] = { parentId = 106, typeId = 108, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {110}, }
_typeInfoList[6] = { parentId = 108, typeId = 110, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2458, 2526, 2528, 2530, 2558, 2562, 2586, 2590, 2606, 2610, 2614, 2616, 2626, 2628, 2632, 2636, 2642, 2646, 2648, 2652, 2654, 2656, 2658, 2662, 2668, 2672, 2676, 2688, 2694, 2698, 2706, 2710, 2714, 2718, 2722, 2728, 2734, 2740, 2748, 2756, 2770, 2774, 2782, 2794, 2808, 2814, 2818, 2826, 2832, 2842, 2848, 2856, 2864, 2872, 2878, 2884, 2888, 2892, 2896, 2900, 2906, 2914, 2920, 2926, 2932, 2938, 2944, 2954, 2958, 2962, 2968, 2974, 2980, 2988, 3004, 3012, 3016, 3024, 3028, 3034, 3038, 3044, 3048, 3056, 3064, 3072, 3078, 3084, 3090, 3094, 3100, 3106, 3110, 3116, 3122, 3124}, }
_typeInfoList[7] = { parentId = 110, typeId = 2458, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2460}, }
_typeInfoList[8] = { parentId = 2458, typeId = 2460, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2462, 2506}, }
_typeInfoList[9] = { parentId = 2460, typeId = 2462, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2464, 2468, 2474, 2476, 2478, 2484, 2490, 2496, 2498, 2500, 2504}, }
_typeInfoList[10] = { parentId = 2462, typeId = 2464, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2466}, }
_typeInfoList[11] = { parentId = 2464, typeId = 2466, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[12] = { parentId = 2462, typeId = 2468, baseId = 2464, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2470, 2472}, }
_typeInfoList[13] = { parentId = 2468, typeId = 2470, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[14] = { parentId = 2468, typeId = 2472, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[15] = { parentId = 2462, typeId = 2474, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[16] = { parentId = 2462, typeId = 2476, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[17] = { parentId = 2462, typeId = 2478, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2480, 2482}, }
_typeInfoList[18] = { parentId = 2478, typeId = 2480, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[19] = { parentId = 2478, typeId = 2482, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[20] = { parentId = 2462, typeId = 2484, baseId = 2478, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2486, 2488}, }
_typeInfoList[21] = { parentId = 2484, typeId = 2486, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[22] = { parentId = 2484, typeId = 2488, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[23] = { parentId = 2462, typeId = 2490, baseId = 2478, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2492, 2494, 2502}, }
_typeInfoList[24] = { parentId = 2490, typeId = 2492, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[25] = { parentId = 2490, typeId = 2494, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {2490}, children = {}, }
_typeInfoList[26] = { parentId = 2462, typeId = 2496, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[27] = { parentId = 2462, typeId = 2498, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[28] = { parentId = 2462, typeId = 2500, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[29] = { parentId = 2490, typeId = 2502, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[30] = { parentId = 2462, typeId = 2504, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[31] = { parentId = 2460, typeId = 2506, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2508, 2512, 2520, 2522, 2524}, }
_typeInfoList[32] = { parentId = 2506, typeId = 2508, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2510}, }
_typeInfoList[33] = { parentId = 2508, typeId = 2510, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[34] = { parentId = 2506, typeId = 2512, baseId = 2508, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2514, 2516, 2518}, }
_typeInfoList[35] = { parentId = 2512, typeId = 2514, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[36] = { parentId = 2512, typeId = 2516, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[37] = { parentId = 2512, typeId = 2518, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[38] = { parentId = 2506, typeId = 2520, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[39] = { parentId = 2506, typeId = 2522, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[40] = { parentId = 2506, typeId = 2524, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[41] = { parentId = 110, typeId = 2526, baseId = 1, txt = 'TypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2532, 2534, 2536, 2538, 2540, 2542, 2544, 2546, 2548, 2550, 2552, 2554, 2556, 2560, 2564, 2566, 2568, 2570, 2572, 2574, 2576, 2578, 2580, 2582, 2584, 2588}, }
_typeInfoList[42] = { parentId = 110, typeId = 2530, baseId = 1, txt = 'isBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[43] = { parentId = 2526, typeId = 2532, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[44] = { parentId = 2526, typeId = 2534, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[45] = { parentId = 2526, typeId = 2536, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[46] = { parentId = 2526, typeId = 2538, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[47] = { parentId = 2526, typeId = 2540, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[48] = { parentId = 2526, typeId = 2542, baseId = 1, txt = 'cloneToPublic',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[49] = { parentId = 2526, typeId = 2544, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[50] = { parentId = 2526, typeId = 2546, baseId = 1, txt = 'createBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[51] = { parentId = 2526, typeId = 2548, baseId = 1, txt = 'createList',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[52] = { parentId = 2526, typeId = 2550, baseId = 1, txt = 'createArray',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[53] = { parentId = 2526, typeId = 2552, baseId = 1, txt = 'createMap',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[54] = { parentId = 2526, typeId = 2554, baseId = 1, txt = 'createClass',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[55] = { parentId = 2526, typeId = 2556, baseId = 1, txt = 'createFunc',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[56] = { parentId = 110, typeId = 2558, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {2526}, retTypeId = {}, children = {}, }
_typeInfoList[57] = { parentId = 2526, typeId = 2560, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2558}, children = {}, }
_typeInfoList[58] = { parentId = 110, typeId = 2562, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {2526}, retTypeId = {}, children = {}, }
_typeInfoList[59] = { parentId = 2526, typeId = 2564, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2562}, children = {}, }
_typeInfoList[60] = { parentId = 2526, typeId = 2566, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[61] = { parentId = 2526, typeId = 2568, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[62] = { parentId = 2526, typeId = 2570, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[63] = { parentId = 2526, typeId = 2572, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[64] = { parentId = 2526, typeId = 2574, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[65] = { parentId = 2526, typeId = 2576, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[66] = { parentId = 2526, typeId = 2578, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[67] = { parentId = 2526, typeId = 2580, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[68] = { parentId = 2526, typeId = 2582, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[69] = { parentId = 2526, typeId = 2584, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[70] = { parentId = 110, typeId = 2586, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2526}, retTypeId = {}, children = {}, }
_typeInfoList[71] = { parentId = 2526, typeId = 2588, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2586}, children = {}, }
_typeInfoList[72] = { parentId = 110, typeId = 2590, baseId = 1, txt = 'Scope',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2592, 2594, 2596, 2598, 2600, 2602, 2604, 2608, 2612}, }
_typeInfoList[73] = { parentId = 2590, typeId = 2592, baseId = 1, txt = 'add',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[74] = { parentId = 2590, typeId = 2594, baseId = 1, txt = 'addClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[75] = { parentId = 2590, typeId = 2596, baseId = 1, txt = 'getClassScope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2590}, children = {}, }
_typeInfoList[76] = { parentId = 2590, typeId = 2598, baseId = 1, txt = 'getTypeInfoChild',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[77] = { parentId = 2590, typeId = 2600, baseId = 1, txt = 'getTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[78] = { parentId = 2590, typeId = 2602, baseId = 1, txt = 'getTypeInfoMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[79] = { parentId = 2590, typeId = 2604, baseId = 1, txt = 'get_parent',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2590}, children = {}, }
_typeInfoList[80] = { parentId = 110, typeId = 2606, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 2526}, retTypeId = {}, children = {}, }
_typeInfoList[81] = { parentId = 2590, typeId = 2608, baseId = 1, txt = 'get_symbol2TypeInfoMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2606}, children = {}, }
_typeInfoList[82] = { parentId = 110, typeId = 2610, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 2590}, retTypeId = {}, children = {}, }
_typeInfoList[83] = { parentId = 2590, typeId = 2612, baseId = 1, txt = 'get_className2ScopeMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2610}, children = {}, }
_typeInfoList[84] = { parentId = 110, typeId = 2614, baseId = 1, txt = 'NodePos',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[85] = { parentId = 110, typeId = 2616, baseId = 1, txt = 'Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2618, 2620, 2622, 2624}, }
_typeInfoList[86] = { parentId = 2616, typeId = 2618, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[87] = { parentId = 2616, typeId = 2620, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[88] = { parentId = 2616, typeId = 2622, baseId = 1, txt = 'get_expType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[89] = { parentId = 2616, typeId = 2624, baseId = 1, txt = 'get_info',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[90] = { parentId = 110, typeId = 2626, baseId = 1, txt = 'MacroEval',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2630}, }
_typeInfoList[91] = { parentId = 110, typeId = 2628, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, retTypeId = {}, children = {}, }
_typeInfoList[92] = { parentId = 2626, typeId = 2630, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2628}, children = {}, }
_typeInfoList[93] = { parentId = 110, typeId = 2632, baseId = 1, txt = 'DeclMacroInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2634, 2638, 2640, 2644}, }
_typeInfoList[94] = { parentId = 2632, typeId = 2634, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[95] = { parentId = 110, typeId = 2636, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2616}, retTypeId = {}, children = {}, }
_typeInfoList[96] = { parentId = 2632, typeId = 2638, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2636}, children = {}, }
_typeInfoList[97] = { parentId = 2632, typeId = 2640, baseId = 1, txt = 'get_ast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[98] = { parentId = 110, typeId = 2642, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2476}, retTypeId = {}, children = {}, }
_typeInfoList[99] = { parentId = 2632, typeId = 2644, baseId = 1, txt = 'get_tokenList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2642}, children = {}, }
_typeInfoList[100] = { parentId = 110, typeId = 2646, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2650, 3130}, }
_typeInfoList[101] = { parentId = 110, typeId = 2648, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[102] = { parentId = 2646, typeId = 2650, baseId = 1, txt = 'get_errMessList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2648}, children = {}, }
_typeInfoList[103] = { parentId = 110, typeId = 2654, baseId = 1, txt = 'getNodeKindName',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[104] = { parentId = 110, typeId = 2656, baseId = 1, txt = 'nodeFilter',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[105] = { parentId = 110, typeId = 2658, baseId = 2616, txt = 'NoneNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2660}, }
_typeInfoList[106] = { parentId = 2658, typeId = 2660, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[107] = { parentId = 110, typeId = 2662, baseId = 2616, txt = 'ImportNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2664, 2666}, }
_typeInfoList[108] = { parentId = 2662, typeId = 2664, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[109] = { parentId = 2662, typeId = 2666, baseId = 1, txt = 'get_modulePath',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[110] = { parentId = 110, typeId = 2668, baseId = 2616, txt = 'RootNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2670, 2674}, }
_typeInfoList[111] = { parentId = 2668, typeId = 2670, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[112] = { parentId = 110, typeId = 2672, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2616}, retTypeId = {}, children = {}, }
_typeInfoList[113] = { parentId = 2668, typeId = 2674, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2672}, children = {}, }
_typeInfoList[114] = { parentId = 110, typeId = 2676, baseId = 2616, txt = 'RefTypeNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2678, 2680, 2682, 2684, 2686}, }
_typeInfoList[115] = { parentId = 2676, typeId = 2678, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[116] = { parentId = 2676, typeId = 2680, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[117] = { parentId = 2676, typeId = 2682, baseId = 1, txt = 'get_refFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[118] = { parentId = 2676, typeId = 2684, baseId = 1, txt = 'get_mutFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[119] = { parentId = 2676, typeId = 2686, baseId = 1, txt = 'get_array',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[120] = { parentId = 110, typeId = 2688, baseId = 2616, txt = 'BlockNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2690, 2692, 2696}, }
_typeInfoList[121] = { parentId = 2688, typeId = 2690, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[122] = { parentId = 2688, typeId = 2692, baseId = 1, txt = 'get_blockKind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[123] = { parentId = 110, typeId = 2694, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2616}, retTypeId = {}, children = {}, }
_typeInfoList[124] = { parentId = 2688, typeId = 2696, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2694}, children = {}, }
_typeInfoList[125] = { parentId = 110, typeId = 2698, baseId = 1, txt = 'IfStmtInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2700, 2702, 2704}, }
_typeInfoList[126] = { parentId = 2698, typeId = 2700, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[127] = { parentId = 2698, typeId = 2702, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[128] = { parentId = 2698, typeId = 2704, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[129] = { parentId = 110, typeId = 2706, baseId = 2616, txt = 'IfNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2708, 2712}, }
_typeInfoList[130] = { parentId = 2706, typeId = 2708, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[131] = { parentId = 110, typeId = 2710, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2698}, retTypeId = {}, children = {}, }
_typeInfoList[132] = { parentId = 2706, typeId = 2712, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2710}, children = {}, }
_typeInfoList[133] = { parentId = 110, typeId = 2714, baseId = 2616, txt = 'ExpListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2716, 2720}, }
_typeInfoList[134] = { parentId = 2714, typeId = 2716, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[135] = { parentId = 110, typeId = 2718, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2616}, retTypeId = {}, children = {}, }
_typeInfoList[136] = { parentId = 2714, typeId = 2720, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2718}, children = {}, }
_typeInfoList[137] = { parentId = 110, typeId = 2722, baseId = 1, txt = 'CaseInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2724, 2726}, }
_typeInfoList[138] = { parentId = 2722, typeId = 2724, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2714}, children = {}, }
_typeInfoList[139] = { parentId = 2722, typeId = 2726, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[140] = { parentId = 110, typeId = 2728, baseId = 2616, txt = 'SwitchNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2730, 2732, 2736, 2738}, }
_typeInfoList[141] = { parentId = 2728, typeId = 2730, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[142] = { parentId = 2728, typeId = 2732, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[143] = { parentId = 110, typeId = 2734, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2722}, retTypeId = {}, children = {}, }
_typeInfoList[144] = { parentId = 2728, typeId = 2736, baseId = 1, txt = 'get_caseList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2734}, children = {}, }
_typeInfoList[145] = { parentId = 2728, typeId = 2738, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[146] = { parentId = 110, typeId = 2740, baseId = 2616, txt = 'WhileNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2742, 2744, 2746}, }
_typeInfoList[147] = { parentId = 2740, typeId = 2742, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[148] = { parentId = 2740, typeId = 2744, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[149] = { parentId = 2740, typeId = 2746, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[150] = { parentId = 110, typeId = 2748, baseId = 2616, txt = 'RepeatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2750, 2752, 2754}, }
_typeInfoList[151] = { parentId = 2748, typeId = 2750, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[152] = { parentId = 2748, typeId = 2752, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[153] = { parentId = 2748, typeId = 2754, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[154] = { parentId = 110, typeId = 2756, baseId = 2616, txt = 'ForNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2758, 2760, 2762, 2764, 2766, 2768}, }
_typeInfoList[155] = { parentId = 2756, typeId = 2758, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[156] = { parentId = 2756, typeId = 2760, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[157] = { parentId = 2756, typeId = 2762, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[158] = { parentId = 2756, typeId = 2764, baseId = 1, txt = 'get_init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[159] = { parentId = 2756, typeId = 2766, baseId = 1, txt = 'get_to',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[160] = { parentId = 2756, typeId = 2768, baseId = 1, txt = 'get_delta',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[161] = { parentId = 110, typeId = 2770, baseId = 2616, txt = 'ApplyNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2772, 2776, 2778, 2780}, }
_typeInfoList[162] = { parentId = 2770, typeId = 2772, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[163] = { parentId = 110, typeId = 2774, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2476}, retTypeId = {}, children = {}, }
_typeInfoList[164] = { parentId = 2770, typeId = 2776, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2774}, children = {}, }
_typeInfoList[165] = { parentId = 2770, typeId = 2778, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[166] = { parentId = 2770, typeId = 2780, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[167] = { parentId = 110, typeId = 2782, baseId = 2616, txt = 'ForeachNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2784, 2786, 2788, 2790, 2792}, }
_typeInfoList[168] = { parentId = 2782, typeId = 2784, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[169] = { parentId = 2782, typeId = 2786, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[170] = { parentId = 2782, typeId = 2788, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[171] = { parentId = 2782, typeId = 2790, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[172] = { parentId = 2782, typeId = 2792, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[173] = { parentId = 110, typeId = 2794, baseId = 2616, txt = 'ForsortNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2796, 2798, 2800, 2802, 2804, 2806}, }
_typeInfoList[174] = { parentId = 2794, typeId = 2796, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[175] = { parentId = 2794, typeId = 2798, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[176] = { parentId = 2794, typeId = 2800, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[177] = { parentId = 2794, typeId = 2802, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[178] = { parentId = 2794, typeId = 2804, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[179] = { parentId = 2794, typeId = 2806, baseId = 1, txt = 'get_sort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[180] = { parentId = 110, typeId = 2808, baseId = 2616, txt = 'ReturnNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2810, 2812}, }
_typeInfoList[181] = { parentId = 2808, typeId = 2810, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[182] = { parentId = 2808, typeId = 2812, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2714}, children = {}, }
_typeInfoList[183] = { parentId = 110, typeId = 2814, baseId = 2616, txt = 'BreakNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2816}, }
_typeInfoList[184] = { parentId = 2814, typeId = 2816, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[185] = { parentId = 110, typeId = 2818, baseId = 2616, txt = 'ExpNewNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2820, 2822, 2824}, }
_typeInfoList[186] = { parentId = 2818, typeId = 2820, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[187] = { parentId = 2818, typeId = 2822, baseId = 1, txt = 'get_symbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[188] = { parentId = 2818, typeId = 2824, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2714}, children = {}, }
_typeInfoList[189] = { parentId = 110, typeId = 2826, baseId = 2616, txt = 'ExpRefNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2828, 2830}, }
_typeInfoList[190] = { parentId = 2826, typeId = 2828, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[191] = { parentId = 2826, typeId = 2830, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[192] = { parentId = 110, typeId = 2832, baseId = 2616, txt = 'ExpOp2Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2834, 2836, 2838, 2840}, }
_typeInfoList[193] = { parentId = 2832, typeId = 2834, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[194] = { parentId = 2832, typeId = 2836, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[195] = { parentId = 2832, typeId = 2838, baseId = 1, txt = 'get_exp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[196] = { parentId = 2832, typeId = 2840, baseId = 1, txt = 'get_exp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[197] = { parentId = 110, typeId = 2842, baseId = 2616, txt = 'ExpCastNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2844, 2846}, }
_typeInfoList[198] = { parentId = 2842, typeId = 2844, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[199] = { parentId = 2842, typeId = 2846, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[200] = { parentId = 110, typeId = 2848, baseId = 2616, txt = 'ExpOp1Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2850, 2852, 2854}, }
_typeInfoList[201] = { parentId = 2848, typeId = 2850, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[202] = { parentId = 2848, typeId = 2852, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[203] = { parentId = 2848, typeId = 2854, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[204] = { parentId = 110, typeId = 2856, baseId = 2616, txt = 'ExpRefItemNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2858, 2860, 2862}, }
_typeInfoList[205] = { parentId = 2856, typeId = 2858, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[206] = { parentId = 2856, typeId = 2860, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[207] = { parentId = 2856, typeId = 2862, baseId = 1, txt = 'get_index',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[208] = { parentId = 110, typeId = 2864, baseId = 2616, txt = 'ExpCallNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2866, 2868, 2870}, }
_typeInfoList[209] = { parentId = 2864, typeId = 2866, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[210] = { parentId = 2864, typeId = 2868, baseId = 1, txt = 'get_func',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[211] = { parentId = 2864, typeId = 2870, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2714}, children = {}, }
_typeInfoList[212] = { parentId = 110, typeId = 2872, baseId = 2616, txt = 'ExpDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2874, 2876}, }
_typeInfoList[213] = { parentId = 2872, typeId = 2874, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[214] = { parentId = 2872, typeId = 2876, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[215] = { parentId = 110, typeId = 2878, baseId = 2616, txt = 'ExpParenNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2880, 2882}, }
_typeInfoList[216] = { parentId = 2878, typeId = 2880, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[217] = { parentId = 2878, typeId = 2882, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[218] = { parentId = 110, typeId = 2884, baseId = 2616, txt = 'ExpMacroExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2886, 2890}, }
_typeInfoList[219] = { parentId = 2884, typeId = 2886, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[220] = { parentId = 110, typeId = 2888, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2616}, retTypeId = {}, children = {}, }
_typeInfoList[221] = { parentId = 2884, typeId = 2890, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2888}, children = {}, }
_typeInfoList[222] = { parentId = 110, typeId = 2892, baseId = 2616, txt = 'ExpMacroStatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2894, 2898}, }
_typeInfoList[223] = { parentId = 2892, typeId = 2894, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[224] = { parentId = 110, typeId = 2896, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2616}, retTypeId = {}, children = {}, }
_typeInfoList[225] = { parentId = 2892, typeId = 2898, baseId = 1, txt = 'get_expStrList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2896}, children = {}, }
_typeInfoList[226] = { parentId = 110, typeId = 2900, baseId = 2616, txt = 'StmtExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2902, 2904}, }
_typeInfoList[227] = { parentId = 2900, typeId = 2902, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[228] = { parentId = 2900, typeId = 2904, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[229] = { parentId = 110, typeId = 2906, baseId = 2616, txt = 'RefFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2908, 2910, 2912}, }
_typeInfoList[230] = { parentId = 2906, typeId = 2908, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[231] = { parentId = 2906, typeId = 2910, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[232] = { parentId = 2906, typeId = 2912, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[233] = { parentId = 110, typeId = 2914, baseId = 1, txt = 'VarInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2916, 2918}, }
_typeInfoList[234] = { parentId = 2914, typeId = 2916, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[235] = { parentId = 2914, typeId = 2918, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[236] = { parentId = 110, typeId = 2920, baseId = 2616, txt = 'DeclVarNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2922, 2924, 2928, 2930, 2934, 2936}, }
_typeInfoList[237] = { parentId = 2920, typeId = 2922, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[238] = { parentId = 2920, typeId = 2924, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[239] = { parentId = 110, typeId = 2926, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2914}, retTypeId = {}, children = {}, }
_typeInfoList[240] = { parentId = 2920, typeId = 2928, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2926}, children = {}, }
_typeInfoList[241] = { parentId = 2920, typeId = 2930, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2714}, children = {}, }
_typeInfoList[242] = { parentId = 110, typeId = 2932, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2526}, retTypeId = {}, children = {}, }
_typeInfoList[243] = { parentId = 2920, typeId = 2934, baseId = 1, txt = 'get_typeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2932}, children = {}, }
_typeInfoList[244] = { parentId = 2920, typeId = 2936, baseId = 1, txt = 'get_unwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2688}, children = {}, }
_typeInfoList[245] = { parentId = 110, typeId = 2938, baseId = 1, txt = 'DeclFuncInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2940, 2942, 2946, 2948, 2950, 2952, 2956, 2960}, }
_typeInfoList[246] = { parentId = 2938, typeId = 2940, baseId = 1, txt = 'get_className',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[247] = { parentId = 2938, typeId = 2942, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[248] = { parentId = 110, typeId = 2944, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2616}, retTypeId = {}, children = {}, }
_typeInfoList[249] = { parentId = 2938, typeId = 2946, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2944}, children = {}, }
_typeInfoList[250] = { parentId = 2938, typeId = 2948, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[251] = { parentId = 2938, typeId = 2950, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[252] = { parentId = 2938, typeId = 2952, baseId = 1, txt = 'get_body',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[253] = { parentId = 110, typeId = 2954, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2526}, retTypeId = {}, children = {}, }
_typeInfoList[254] = { parentId = 2938, typeId = 2956, baseId = 1, txt = 'get_retTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2954}, children = {}, }
_typeInfoList[255] = { parentId = 110, typeId = 2958, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2526}, retTypeId = {}, children = {}, }
_typeInfoList[256] = { parentId = 2938, typeId = 2960, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2958}, children = {}, }
_typeInfoList[257] = { parentId = 110, typeId = 2962, baseId = 2616, txt = 'DeclFuncNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2964, 2966}, }
_typeInfoList[258] = { parentId = 2962, typeId = 2964, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[259] = { parentId = 2962, typeId = 2966, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2938}, children = {}, }
_typeInfoList[260] = { parentId = 110, typeId = 2968, baseId = 2616, txt = 'DeclMethodNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2970, 2972}, }
_typeInfoList[261] = { parentId = 2968, typeId = 2970, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[262] = { parentId = 2968, typeId = 2972, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2938}, children = {}, }
_typeInfoList[263] = { parentId = 110, typeId = 2974, baseId = 2616, txt = 'DeclConstrNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2976, 2978}, }
_typeInfoList[264] = { parentId = 2974, typeId = 2976, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[265] = { parentId = 2974, typeId = 2978, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2938}, children = {}, }
_typeInfoList[266] = { parentId = 110, typeId = 2980, baseId = 2616, txt = 'ExpCallSuperNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2982, 2984, 2986}, }
_typeInfoList[267] = { parentId = 2980, typeId = 2982, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[268] = { parentId = 2980, typeId = 2984, baseId = 1, txt = 'get_superType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[269] = { parentId = 2980, typeId = 2986, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2714}, children = {}, }
_typeInfoList[270] = { parentId = 110, typeId = 2988, baseId = 2616, txt = 'DeclMemberNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2990, 2992, 2994, 2996, 2998, 3000, 3002}, }
_typeInfoList[271] = { parentId = 2988, typeId = 2990, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[272] = { parentId = 2988, typeId = 2992, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[273] = { parentId = 2988, typeId = 2994, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2676}, children = {}, }
_typeInfoList[274] = { parentId = 2988, typeId = 2996, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[275] = { parentId = 2988, typeId = 2998, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[276] = { parentId = 2988, typeId = 3000, baseId = 1, txt = 'get_getterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[277] = { parentId = 2988, typeId = 3002, baseId = 1, txt = 'get_setterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[278] = { parentId = 110, typeId = 3004, baseId = 2616, txt = 'DeclArgNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3006, 3008, 3010}, }
_typeInfoList[279] = { parentId = 3004, typeId = 3006, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[280] = { parentId = 3004, typeId = 3008, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[281] = { parentId = 3004, typeId = 3010, baseId = 1, txt = 'get_argType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2676}, children = {}, }
_typeInfoList[282] = { parentId = 110, typeId = 3012, baseId = 2616, txt = 'DeclArgDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3014}, }
_typeInfoList[283] = { parentId = 3012, typeId = 3014, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[284] = { parentId = 110, typeId = 3016, baseId = 2616, txt = 'DeclClassNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3018, 3020, 3022, 3026, 3030, 3032, 3036}, }
_typeInfoList[285] = { parentId = 3016, typeId = 3018, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[286] = { parentId = 3016, typeId = 3020, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[287] = { parentId = 3016, typeId = 3022, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[288] = { parentId = 110, typeId = 3024, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2616}, retTypeId = {}, children = {}, }
_typeInfoList[289] = { parentId = 3016, typeId = 3026, baseId = 1, txt = 'get_fieldList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3024}, children = {}, }
_typeInfoList[290] = { parentId = 110, typeId = 3028, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2988}, retTypeId = {}, children = {}, }
_typeInfoList[291] = { parentId = 3016, typeId = 3030, baseId = 1, txt = 'get_memberList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3028}, children = {}, }
_typeInfoList[292] = { parentId = 3016, typeId = 3032, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2590}, children = {}, }
_typeInfoList[293] = { parentId = 110, typeId = 3034, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[294] = { parentId = 3016, typeId = 3036, baseId = 1, txt = 'get_outerMethodSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3034}, children = {}, }
_typeInfoList[295] = { parentId = 110, typeId = 3038, baseId = 2616, txt = 'DeclMacroNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3040, 3042}, }
_typeInfoList[296] = { parentId = 3038, typeId = 3040, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[297] = { parentId = 3038, typeId = 3042, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2632}, children = {}, }
_typeInfoList[298] = { parentId = 110, typeId = 3044, baseId = 2616, txt = 'LiteralNilNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3046}, }
_typeInfoList[299] = { parentId = 3044, typeId = 3046, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[300] = { parentId = 110, typeId = 3048, baseId = 2616, txt = 'LiteralCharNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3050, 3052, 3054}, }
_typeInfoList[301] = { parentId = 3048, typeId = 3050, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[302] = { parentId = 3048, typeId = 3052, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[303] = { parentId = 3048, typeId = 3054, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[304] = { parentId = 110, typeId = 3056, baseId = 2616, txt = 'LiteralIntNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3058, 3060, 3062}, }
_typeInfoList[305] = { parentId = 3056, typeId = 3058, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[306] = { parentId = 3056, typeId = 3060, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[307] = { parentId = 3056, typeId = 3062, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[308] = { parentId = 110, typeId = 3064, baseId = 2616, txt = 'LiteralRealNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3066, 3068, 3070}, }
_typeInfoList[309] = { parentId = 3064, typeId = 3066, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[310] = { parentId = 3064, typeId = 3068, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[311] = { parentId = 3064, typeId = 3070, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[312] = { parentId = 110, typeId = 3072, baseId = 2616, txt = 'LiteralArrayNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3074, 3076}, }
_typeInfoList[313] = { parentId = 3072, typeId = 3074, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[314] = { parentId = 3072, typeId = 3076, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2714}, children = {}, }
_typeInfoList[315] = { parentId = 110, typeId = 3078, baseId = 2616, txt = 'LiteralListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3080, 3082}, }
_typeInfoList[316] = { parentId = 3078, typeId = 3080, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[317] = { parentId = 3078, typeId = 3082, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2714}, children = {}, }
_typeInfoList[318] = { parentId = 110, typeId = 3084, baseId = 1, txt = 'PairList',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3086, 3088}, }
_typeInfoList[319] = { parentId = 3084, typeId = 3086, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[320] = { parentId = 3084, typeId = 3088, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[321] = { parentId = 110, typeId = 3090, baseId = 2616, txt = 'LiteralMapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3092, 3096, 3098}, }
_typeInfoList[322] = { parentId = 3090, typeId = 3092, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[323] = { parentId = 110, typeId = 3094, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {2616, 2616}, retTypeId = {}, children = {}, }
_typeInfoList[324] = { parentId = 3090, typeId = 3096, baseId = 1, txt = 'get_map',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3094}, children = {}, }
_typeInfoList[325] = { parentId = 3090, typeId = 3098, baseId = 1, txt = 'get_pairList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3084}, children = {}, }
_typeInfoList[326] = { parentId = 110, typeId = 3100, baseId = 2616, txt = 'LiteralStringNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3102, 3104, 3108}, }
_typeInfoList[327] = { parentId = 3100, typeId = 3102, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[328] = { parentId = 3100, typeId = 3104, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[329] = { parentId = 110, typeId = 3106, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2616}, retTypeId = {}, children = {}, }
_typeInfoList[330] = { parentId = 3100, typeId = 3108, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3106}, children = {}, }
_typeInfoList[331] = { parentId = 110, typeId = 3110, baseId = 2616, txt = 'LiteralBoolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3112, 3114}, }
_typeInfoList[332] = { parentId = 3110, typeId = 3112, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[333] = { parentId = 3110, typeId = 3114, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[334] = { parentId = 110, typeId = 3116, baseId = 2616, txt = 'LiteralSymbolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3118, 3120}, }
_typeInfoList[335] = { parentId = 3116, typeId = 3118, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[336] = { parentId = 3116, typeId = 3120, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2476}, children = {}, }
_typeInfoList[337] = { parentId = 110, typeId = 3122, baseId = 1, txt = 'getLiteralValue',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {6, 2526}, children = {}, }
_typeInfoList[338] = { parentId = 110, typeId = 3124, baseId = 1, txt = 'ASTInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {3126, 3128}, }
_typeInfoList[339] = { parentId = 3124, typeId = 3126, baseId = 1, txt = 'get_node',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2616}, children = {}, }
_typeInfoList[340] = { parentId = 3124, typeId = 3128, baseId = 1, txt = 'get_moduleTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2526}, children = {}, }
_typeInfoList[341] = { parentId = 2646, typeId = 3130, baseId = 1, txt = 'createAST',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {3124}, children = {}, }
_typeInfoList[342] = { parentId = 104, typeId = 3132, baseId = 1, txt = 'filterObj',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
----- meta -----
return moduleObj
