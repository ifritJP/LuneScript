--lune/base/dumpNode.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@dumpNode'
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

function _lune.nilacc( val, fieldName, access, ... )
   if not val then
      return nil
   end
   if fieldName then
      local field = val[ fieldName ]
      if not field then
         return nil
      end
      if access == "item" then
         local typeId = type( field )
         if typeId == "table" then
            return field[ ... ]
         elseif typeId == "string" then
            return string.byte( field, ... )
         end
      elseif access == "call" then
         return field( ... )
      elseif access == "callmtd" then
         return field( val, ... )
      end
      return field
   end
   if access == "item" then
      local typeId = type( val )
      if typeId == "table" then
         return val[ ... ]
      elseif typeId == "string" then
         return string.byte( val, ... )
      end
   elseif access == "call" then
      return val( ... )
   elseif access == "list" then
      local list, arg = ...
      if not list then
         return nil
      end
      return val( list, arg )
   end
   error( string.format( "illegal access -- %s", access ) )
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
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )
local Util = _lune.loadModule( 'lune.base.Util' )

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
setmetatable( dumpFilter, { __index = Nodes.Filter } )
function dumpFilter.setmeta( obj )
  setmetatable( obj, { __index = dumpFilter  } )
end
function dumpFilter.new( __superarg1, __superarg2, __superarg3 )
   local obj = {}
   dumpFilter.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3 )
   end
   return obj
end
function dumpFilter:__init( __superarg1, __superarg2, __superarg3 )

   Nodes.Filter.__init( self, __superarg1, __superarg2, __superarg3 )
end


local function createFilter( moduleTypeInfo )

   return dumpFilter.new(true, moduleTypeInfo, moduleTypeInfo:get_scope())
end
_moduleObj.createFilter = createFilter

local function dump( prefix, depth, node, txt )

   local typeStr = ""
   local expType = node:get_expType(  )
   if expType:equals( Ast.builtinTypeNone ) then
      typeStr = string.format( "(%d:%s:%s)", expType:get_typeId(  ), expType:getTxt(  ), tostring( expType:get_kind(  )))
   end
   
   
   local comment
   
   do
      local commentList = node:get_commentList()
      if commentList ~= nil then
         comment = string.format( "comment:%d,%d", #commentList, node:get_tailComment() and 1 or 0)
      else
         if node:get_tailComment() then
            comment = "comment:0,1"
         else
          
            comment = ""
         end
         
      end
   end
   
   
   print( string.format( "%s: %s %s %s %s", prefix, Nodes.getNodeKindName( node:get_kind(  ) ), txt, typeStr, comment) )
end

local function filter( node, filter, opt )

   node:processFilter( filter, opt )
end

local function getTxt( token )

   return token.txt
end

function dumpFilter:processNone( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
end


function dumpFilter:processLuneControl( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, LuneControl.Pragma:_getTxt( node:get_pragma())
    )
end


function dumpFilter:processBlankLine( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%d", node:get_lineNum()) )
end


function dumpFilter:processLuneKind( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_exp(), self, opt:nextOpt(  ) )
end


function dumpFilter:processImport( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_modulePath(  ) )
end


function dumpFilter:processRoot( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   for __index, child in ipairs( node:get_children(  ) ) do
      filter( child, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processSubfile( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
end

function dumpFilter:processBlockSub( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   for __index, statement in ipairs( node:get_stmtList(  ) ) do
      filter( statement, self, opt:nextOpt(  ) )
   end
   
end



function dumpFilter:processScope( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, Nodes.ScopeKind:_getTxt( node:get_scopeKind())
    )
   filter( node:get_block(), self, opt:nextOpt(  ) )
end


function dumpFilter:processStmtExp( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
end



function dumpFilter:processDeclEnum( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_name().txt )
   local enumTypeInfo = _lune.unwrap( (_lune.__Cast( node:get_expType(), 3, Ast.EnumTypeInfo ) ))
   for __index, name in ipairs( node:get_valueNameList() ) do
      local valInfo = _lune.unwrap( enumTypeInfo:getEnumValInfo( name.txt ))
      print( string.format( "%s  %s: %s", prefix, name.txt, Ast.EnumLiteral:_getTxt( valInfo:get_val())
      ) )
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
   for __index, exp in ipairs( node:get_paramList() ) do
      filter( exp, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processDeclClass( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_name(  ).txt )
   for __index, field in ipairs( node:get_fieldList(  ) ) do
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
   for __index, stmt in ipairs( stmtList ) do
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
   for __index, expStr in ipairs( node:get_expStrList(  ) ) do
      filter( expStr, self, opt:nextOpt(  ) )
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
   for __index, expNode in ipairs( node:get_expList():get_expList() ) do
      filter( expNode, self, opt:nextOpt(  ) )
   end
   
   filter( node:get_block(), self, opt:nextOpt(  ) )
   if node:get_nilBlock() then
      filter( _lune.unwrap( node:get_nilBlock()), self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processWhen( node, opt )

   local prefix, depth = opt:get(  )
   
   local symTxt = ""
   for __index, symPair in ipairs( node:get_symPairList() ) do
      symTxt = string.format( "%s %s", symTxt, tostring( symPair:get_src()))
   end
   
   
   dump( prefix, depth, node, symTxt )
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
   for index, var in ipairs( node:get_varList(  ) ) do
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
   for __index, var in ipairs( node:get_varList(  ) ) do
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
   
   do
      local _exp = node:get_syncBlock()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt(  ) )
      end
   end
   
end


function dumpFilter:processDeclArg( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%s:%s", node:get_name(  ).txt, node:get_expType():getTxt(  )) )
   
end


function dumpFilter:processDeclArgDDD( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "..." )
end


function dumpFilter:processExpDDD( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "..." )
end


function dumpFilter:processExpSubDDD( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "... (%d)", node:get_remainIndex()) )
end



function dumpFilter:processDeclForm( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_expType():getTxt( self:get_typeNameCtrl() ) )
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
   
   if Ast.TypeInfo.isMut( node:get_expType() ) then
      name = name .. " mut"
   end
   
   do
      local _exp = _lune.nilacc( node:get_expType():get_scope(), 'get_closureSymList', 'callmtd' )
      if _exp ~= nil then
         if #_exp > 0 then
            name = name .. " closure"
         end
         
      end
   end
   
   dump( prefix, depth, node, name )
   local argList = declInfo:get_argList(  )
   for __index, arg in ipairs( argList ) do
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


function dumpFilter:processProtoMethod( node, opt )

   self:processDeclFuncInfo( node, node:get_declInfo(  ), opt )
end


function dumpFilter:processDeclConstr( node, opt )

   self:processDeclFuncInfo( node, node:get_declInfo(  ), opt )
end



function dumpFilter:processDeclDestr( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
end



function dumpFilter:processExpCallSuperCtor( node, opt )

   local prefix, depth = opt:get(  )
   local typeInfo = node:get_superType(  )
   dump( prefix, depth, node, typeInfo:getTxt(  ) )
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
   for __index, stmt in ipairs( stmtList ) do
      if stmt:get_exp():get_kind() ~= Nodes.nodeKindEnum.None then
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
   for __index, caseInfo in ipairs( caseList ) do
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
   for __index, caseInfo in ipairs( caseList ) do
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
   dump( prefix, depth, node, node:get_val(  ):get_name() )
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
   for __index, var in ipairs( varList ) do
      varNames = varNames .. var:get_name() .. " "
   end
   
   dump( prefix, depth, node, varNames )
   filter( node:get_expList(), self, opt:nextOpt(  ) )
   filter( node:get_block(), self, opt:nextOpt(  ) )
end


function dumpFilter:processForeach( node, opt )

   local prefix, depth = opt:get(  )
   local index = ""
   do
      local _exp = node:get_key()
      if _exp ~= nil then
         index = _exp:get_name()
      end
   end
   
   dump( prefix, depth, node, node:get_val():get_name() .. " " .. index )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
   filter( node:get_block(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processForsort( node, opt )

   local prefix, depth = opt:get(  )
   local index = ""
   do
      local _exp = node:get_key(  )
      if _exp ~= nil then
         index = _exp:get_name()
      end
   end
   
   dump( prefix, depth, node, node:get_val(  ):get_name() .. " " .. index )
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


local function getTypeListTxt( typeList )

   local txt = ""
   for index, typeInfo in ipairs( typeList ) do
      if index > 1 then
         txt = txt .. ", "
      end
      
      txt = txt .. typeInfo:getTxt(  )
   end
   
   return txt
end

function dumpFilter:processExpCall( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "[nilacc:%s] -> %s", tostring( node:get_nilAccess()), getTypeListTxt( node:get_expTypeList() )) )
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
   local mess
   
   do
      local mRetExp = node:get_mRetExp()
      if mRetExp ~= nil then
         mess = string.format( "hasMRetExp (%d)", mRetExp:get_index())
      else
         mess = "noMRetExp"
      end
   end
   
   dump( prefix, depth, node, mess )
   local expList = node:get_expList(  )
   for __index, exp in ipairs( expList ) do
      filter( exp, self, opt:nextOpt(  ) )
   end
   
end


function dumpFilter:processExpMRet( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_mRet(), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpAccessMRet( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%d", node:get_index()) )
end


function dumpFilter:processExpOp1( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_op(  ).txt )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
end



function dumpFilter:processExpToDDD( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_expList(), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpMultiTo1( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "" )
   filter( node:get_exp(), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpCast( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%s -> %s", node:get_exp():get_expType():getTxt( self:get_typeNameCtrl() ), node:get_castType():getTxt( self:get_typeNameCtrl() )) )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpParen( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "()" )
   filter( node:get_exp(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpSetVal( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "= %s", node:get_expType():getTxt( self:get_typeNameCtrl() )) )
   filter( node:get_exp1(  ), self, opt:nextOpt(  ) )
   filter( node:get_exp2(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpSetItem( node, opt )

   local prefix, depth = opt:get(  )
   local indexSym = ""
   local indexNode = nil
   do
      local _matchExp = node:get_index()
      if _matchExp[1] == Nodes.IndexVal.NodeIdx[1] then
         local index = _matchExp[2][1]
      
         indexNode = index
      elseif _matchExp[1] == Nodes.IndexVal.SymIdx[1] then
         local index = _matchExp[2][1]
      
         indexSym = index
      end
   end
   
   dump( prefix, depth, node, indexSym )
   filter( node:get_val(), self, opt:nextOpt(  ) )
   if indexNode ~= nil then
      filter( indexNode, self, opt:nextOpt(  ) )
   end
   
   filter( node:get_exp2(), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpOp2( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%s -> %s", node:get_op(  ).txt, node:get_expType():getTxt( self:get_typeNameCtrl() )) )
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
   dump( prefix, depth, node, string.format( "%s:%s", node:get_symbolInfo():get_name(), node:get_expType():getTxt(  )) )
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
   dump( prefix, depth, node, string.format( "%s:%s:%s", node:get_field().txt, _lune.nilacc( node:get_symbolInfo(), 'get_mutable', 'callmtd' ) and "mut" or "imut", node:get_expType():getTxt(  )) )
   filter( node:get_prefix(  ), self, opt:nextOpt(  ) )
end


function dumpFilter:processExpOmitEnum( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "%s.%s", node:get_expType():getTxt(  ), node:get_valToken().txt) )
end



function dumpFilter:processGetField( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, string.format( "get_%s:%s", node:get_field(  ).txt, node:get_expType():getTxt(  )) )
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


function dumpFilter:processTestBlock( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, node:get_name().txt )
   filter( node:get_block(), self, opt:nextOpt(  ) )
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
   for __index, pair in ipairs( pairList ) do
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
   do
      local expList = node:get_orgParam()
      if expList ~= nil then
         for __index, param in ipairs( expList:get_expList() ) do
            filter( param, self, opt:nextOpt(  ) )
         end
         
      end
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



function dumpFilter:processAbbr( node, opt )

   local prefix, depth = opt:get(  )
   dump( prefix, depth, node, "##" )
end



return _moduleObj
