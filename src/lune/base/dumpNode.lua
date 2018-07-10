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
local _funcName2InfoMap = {}
moduleObj._funcName2InfoMap = _funcName2InfoMap
moduleObj._typeInfoList = {}
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 84, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 142, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 248, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 250, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 252, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {270, 272, 328}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 274, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 300, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 302, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 304, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 336, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 338, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 340, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 342, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 344, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 346, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {360, 362, 364}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 348, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 352, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 354, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 356, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 358, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 366, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 368, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 380, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 386, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {406, 408, 410, 412, 418, 420, 422, 428, 430, 434, 438, 440, 444, 450, 458, 462, 464, 466, 468, 470, 472, 474, 476, 478, 480, 482, 486}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 456, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 460, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 484, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 488, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {504, 506, 508, 510, 512, 514, 518, 522}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 516, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 520, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 488}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 526, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 528, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {532, 534, 538}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 540, baseId = 528, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 542, baseId = 528, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 544, baseId = 528, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 546, baseId = 528, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 548, baseId = 528, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 550, baseId = 528, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 552, baseId = 528, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 554, baseId = 528, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 556, baseId = 528, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 558, baseId = 528, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 560, baseId = 528, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 562, baseId = 528, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 564, baseId = 528, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 566, baseId = 528, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 568, baseId = 528, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 570, baseId = 528, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 572, baseId = 528, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 574, baseId = 528, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 576, baseId = 528, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 578, baseId = 528, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 580, baseId = 528, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 582, baseId = 528, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 584, baseId = 528, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 586, baseId = 528, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 588, baseId = 528, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 590, baseId = 528, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 592, baseId = 528, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 594, baseId = 528, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 596, baseId = 528, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 598, baseId = 528, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 600, baseId = 528, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 602, baseId = 528, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 604, baseId = 528, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 606, baseId = 528, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 608, baseId = 528, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 610, baseId = 528, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 612, baseId = 528, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 616, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {696, 1232}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 694, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 748, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 750, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1076, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1084, 1086, 1090, 1092, 1094, 1096, 1100, 1104}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1088, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {528}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1098, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1102, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1190, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1194, 1198}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1196, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {528}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1254, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1256, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1258, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1260, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1262, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1416, 1418, 1420}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1264, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1266, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1268, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1270, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1272, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1274, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1276, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1278, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1280, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1282, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1422, 1424, 1426}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1284, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1286, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1288, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1290, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1292, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1294, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1296, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1300, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1302, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1428, 1430, 1432, 1434, 1436, 1438, 1440, 1442, 1444, 1446, 1448, 1450, 1452, 1454, 1456, 1458, 1460, 1462, 1464, 1466, 1468, 1470, 1472, 1474, 1476, 1478, 1480}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1304, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1306, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1308, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1310, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1482, 1484, 1486, 1488, 1490, 1492, 1494, 1496}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1312, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1314, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1310}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1316, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1318, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1498, 1500, 1502}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1320, baseId = 1318, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1322, baseId = 1318, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1324, baseId = 1318, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1326, baseId = 1318, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1328, baseId = 1318, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1330, baseId = 1318, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1332, baseId = 1318, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1334, baseId = 1318, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1336, baseId = 1318, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1338, baseId = 1318, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1340, baseId = 1318, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1342, baseId = 1318, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1344, baseId = 1318, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1346, baseId = 1318, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1348, baseId = 1318, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1350, baseId = 1318, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1352, baseId = 1318, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1354, baseId = 1318, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1356, baseId = 1318, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1358, baseId = 1318, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1360, baseId = 1318, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1362, baseId = 1318, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1364, baseId = 1318, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1366, baseId = 1318, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1368, baseId = 1318, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1370, baseId = 1318, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1372, baseId = 1318, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1374, baseId = 1318, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1376, baseId = 1318, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1378, baseId = 1318, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1380, baseId = 1318, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1382, baseId = 1318, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1384, baseId = 1318, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1386, baseId = 1318, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1388, baseId = 1318, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1390, baseId = 1318, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1392, baseId = 1318, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1394, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1504, 1506}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1396, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1400, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1402, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1404, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1508, 1510, 1512, 1514, 1516, 1518, 1520, 1522}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1406, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1318}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1408, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1410, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1412, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1524, 1526}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1414, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1318}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1528, baseId = 1, txt = 'filterObj',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 252, typeId = 270, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 252, typeId = 272, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {252}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 252, typeId = 328, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 346, typeId = 360, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 346, typeId = 362, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {346}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 346, typeId = 364, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 406, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 408, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 410, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 412, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 418, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 420, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 422, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 428, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 430, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 434, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 438, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 440, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 444, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 450, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 458, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {456}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 462, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {460}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 464, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 466, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 468, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 470, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 472, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 474, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 476, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 478, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 480, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 482, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 486, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {484}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 504, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 506, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 508, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {488}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 510, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 512, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 514, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {488}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 518, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {516}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 522, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {520}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 528, typeId = 532, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 528, typeId = 534, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 528, typeId = 538, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 616, typeId = 696, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {694}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 616, typeId = 1232, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1084, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {344}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1086, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {344}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1090, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1088}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1092, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1094, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1096, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {528}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1100, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1098}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1104, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1102}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1190, typeId = 1194, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {344}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1190, typeId = 1198, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1196}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1262, typeId = 1416, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1262, typeId = 1418, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1262}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1262, typeId = 1420, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1282, typeId = 1422, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1282, typeId = 1424, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1282}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1282, typeId = 1426, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1428, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1430, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1432, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1434, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1436, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1438, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1440, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1442, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1444, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1446, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1448, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1450, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1452, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1454, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1456, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1304}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1458, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1306}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1460, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1462, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1464, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1466, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1468, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1470, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1472, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1474, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1476, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1478, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1480, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1308}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1482, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1484, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1486, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1310}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1488, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1490, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1492, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1310}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1494, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1312}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1496, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1314}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1318, typeId = 1498, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1318, typeId = 1500, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1318, typeId = 1502, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1394, typeId = 1504, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1396}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1394, typeId = 1506, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1508, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1280}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1510, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1280}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1512, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1406}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1514, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1516, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1518, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1318}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1520, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1408}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1522, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1410}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1412, typeId = 1524, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1280}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1412, typeId = 1526, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1414}, children = {}, }
)
----- meta -----
return moduleObj
