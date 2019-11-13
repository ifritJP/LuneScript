--lune/base/Formatter.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Formatter'
local _lune = {}
if _lune1 then
   _lune = _lune1
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

if not _lune1 then
   _lune1 = _lune
end

local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )

local Opt = {}
_moduleObj.Opt = Opt
function Opt.new( parent )
   local obj = {}
   Opt.setmeta( obj )
   if obj.__init then obj:__init( parent ); end
   return obj
end
function Opt:__init(parent) 
   self.parent = parent
end
function Opt:nextOpt( parent )

   return Opt.new(parent)
end
function Opt.setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
end
function Opt:get_parent()
   return self.parent
end


local FormatterFilter = {}
setmetatable( FormatterFilter, { __index = Nodes.Filter } )
function FormatterFilter.new( moduleTypeInfo, moduleInfoManager, stream )
   local obj = {}
   FormatterFilter.setmeta( obj )
   if obj.__init then obj:__init( moduleTypeInfo, moduleInfoManager, stream ); end
   return obj
end
function FormatterFilter:__init(moduleTypeInfo, moduleInfoManager, stream) 
   Nodes.Filter.__init( self,moduleTypeInfo, moduleInfoManager)
   
   self.stream = Util.SimpleSourceOStream.new(stream, nil, 3)
end
function FormatterFilter.setmeta( obj )
  setmetatable( obj, { __index = FormatterFilter  } )
end
function FormatterFilter:write( ... )
   return self.stream:write( ... )
end

function FormatterFilter:writeln( ... )
   return self.stream:writeln( ... )
end

function FormatterFilter:pushIndent( ... )
   return self.stream:pushIndent( ... )
end

function FormatterFilter:popIndent( ... )
   return self.stream:popIndent( ... )
end

function FormatterFilter:switchToHeader( ... )
   return self.stream:switchToHeader( ... )
end

function FormatterFilter:returnToSource( ... )
   return self.stream:returnToSource( ... )
end



local function createFilter( moduleTypeInfo, stream )

   return FormatterFilter.new(moduleTypeInfo, moduleTypeInfo:get_scope(), stream)
end
_moduleObj.createFilter = createFilter

local function filter( node, filter, opt )

   node:processFilter( filter, opt )
end

local function getTxt( token )

   return token.txt
end

function FormatterFilter:processNone( node, opt )

end



function FormatterFilter:processBlankLine( node, opt )

   for index = 1, node:get_lineNum() do
      self:writeln( "" )
   end
   
end



function FormatterFilter:processImport( node, opt )

   self:write( string.format( "import %s", node:get_modulePath()) )
   if not node:get_modulePath():find( "%." .. node:get_assignName() .. "$" ) then
      self:write( string.format( " as %s", node:get_assignName()) )
   end
   
   self:writeln( ";" )
end


function FormatterFilter:processRoot( node, opt )

   
   for index, child in pairs( node:get_children() ) do
      filter( child, self, opt:nextOpt( node ) )
   end
   
end


function FormatterFilter:processSubfile( node, opt )

end

function FormatterFilter:processBlockSub( node, opt )

   self:writeln( "{" )
   self:pushIndent(  )
   
   for index, statement in pairs( node:get_stmtList() ) do
      filter( statement, self, opt:nextOpt( node ) )
   end
   
   
   self:popIndent(  )
   
   do
      local _switchExp = node:get_blockKind()
      if _switchExp == Nodes.BlockKind.LetUnwrap then
         self:write( "}" )
      else 
         
            self:writeln( "}" )
      end
   end
   
end


function FormatterFilter:processStmtExp( node, opt )

   filter( node:get_exp(), self, opt:nextOpt( node ) )
   self:writeln( ";" )
end



function FormatterFilter:processDeclEnum( node, opt )

   
   local enumTypeInfo = _lune.unwrap( (_lune.__Cast( node:get_expType(), 3, Ast.EnumTypeInfo ) ))
   for __index, name in pairs( node:get_valueNameList() ) do
      local valInfo = _lune.unwrap( enumTypeInfo:getEnumValInfo( name.txt ))
   end
   
end


function FormatterFilter:processDeclAlge( node, opt )

   
   local algeTypeInfo = node:get_algeType()
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
         end
      end
   end
   
end


function FormatterFilter:processNewAlgeVal( node, opt )

   
   for __index, exp in pairs( node:get_paramList() ) do
      filter( exp, self, opt:nextOpt( node ) )
   end
   
end


function FormatterFilter:processDeclClass( node, opt )

   if node:get_accessMode() == Ast.AccessMode.Pub then
      self:write( "pub " )
   end
   
   do
      local _switchExp = node:get_expType():get_kind()
      if _switchExp == Ast.TypeInfoKind.Class then
         if node:get_moduleName() then
            self:write( "module " )
         else
          
            self:write( "class " )
         end
         
      elseif _switchExp == Ast.TypeInfoKind.IF then
         self:write( "interface " )
      end
   end
   
   self:write( node:get_name().txt )
   
   local classType = node:get_expType()
   if #classType:get_itemTypeInfoList() > 0 then
      self:write( "<" )
      for index, genType in pairs( classType:get_itemTypeInfoList() ) do
         if index > 1 then
            self:write( "," )
         end
         
         self:write( genType:getTxt(  ) )
      end
      
      self:write( ">" )
   end
   
   
   do
      local moduleName = node:get_moduleName()
      if moduleName ~= nil then
         self:write( " require " )
         self:write( string.format( "%s ", moduleName.txt) )
         do
            local gluePrefix = node:get_gluePrefix()
            if gluePrefix ~= nil then
               self:write( "glue " )
               self:write( gluePrefix )
            end
         end
         
      end
   end
   
   
   if classType:hasBase(  ) or #classType:get_interfaceList() > 0 then
      self:write( " extend " )
      if classType:hasBase(  ) then
         self:write( classType:get_baseTypeInfo():getTxt(  ) )
      end
      
      if #classType:get_interfaceList() > 0 then
         self:write( "(" )
         for index, ifType in pairs( classType:get_interfaceList() ) do
            if index > 1 then
               self:write( "," )
            end
            
            self:write( ifType:getTxt(  ) )
         end
         
         self:write( ")" )
      end
      
   end
   
   self:writeln( "" )
   self:writeln( "{" )
   self:pushIndent(  )
   
   for index, field in pairs( node:get_fieldList() ) do
      filter( field, self, opt:nextOpt( node ) )
   end
   
   
   self:popIndent(  )
   self:writeln( "}" )
end


function FormatterFilter:processDeclMember( node, opt )

   do
      local _switchExp = node:get_accessMode()
      if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Pro or _switchExp == Ast.AccessMode.None or _switchExp == Ast.AccessMode.Local then
         self:write( Ast.accessMode2txt( node:get_accessMode() ) )
         self:write( " " )
      end
   end
   
   self:write( "let " )
   
   local symbol = node:get_symbolInfo()
   if symbol:get_mutable() then
      self:write( "mut " )
   end
   
   self:write( symbol:get_name() )
   self:write( ":" )
   
   filter( node:get_refType(), self, opt:nextOpt( node ) )
   
   if node:get_getterMode() == Ast.AccessMode.None then
      if node:get_setterMode() ~= Ast.AccessMode.None then
         self:write( "{none," )
         self:write( Ast.accessMode2txt( node:get_setterMode() ) )
         self:write( "}" )
      end
      
   else
    
      self:write( "{" )
      self:write( Ast.accessMode2txt( node:get_getterMode() ) )
      if not node:get_getterMutable() then
         self:write( "&" )
      end
      
      if node:get_getterRetType() ~= symbol:get_typeInfo() then
         self:write( ":" )
         self:write( node:get_getterRetType():getTxt(  ) )
      end
      
      if node:get_setterMode() ~= Ast.AccessMode.None then
         self:write( "," )
         self:write( Ast.accessMode2txt( node:get_setterMode() ) )
      end
      
      self:write( "}" )
   end
   
   self:writeln( ";" )
end


function FormatterFilter:processExpMacroExp( node, opt )

   
   local stmtList = node:get_stmtList()
   for __index, stmt in pairs( stmtList ) do
      filter( stmt, self, opt:nextOpt( node ) )
   end
   
end


function FormatterFilter:processDeclMacro( node, opt )

end


function FormatterFilter:processExpMacroStat( node, opt )

   
   for __index, node in pairs( node:get_expStrList() ) do
      filter( node, self, opt:nextOpt( node ) )
   end
   
end



function FormatterFilter:processUnwrapSet( node, opt )

   
   filter( node:get_dstExpList(), self, opt:nextOpt( node ) )
   filter( node:get_srcExpList(), self, opt:nextOpt( node ) )
   
   if node:get_unwrapBlock() then
      filter( _lune.unwrap( node:get_unwrapBlock()), self, opt:nextOpt( node ) )
   end
   
end


function FormatterFilter:processIfUnwrap( node, opt )

   self:write( "if! " )
   if #node:get_varSymList() ~= 0 then
      self:write( "let " )
      for index, varSym in pairs( node:get_varSymList() ) do
         if index > 1 then
            self:write( "," )
         end
         
         self:write( varSym:get_name() )
      end
      
      self:write( " = " )
   end
   
   filter( node:get_expList(), self, opt:nextOpt( node ) )
   filter( node:get_block(), self, opt:nextOpt( node ) )
   if node:get_nilBlock() then
      self:write( "else " )
      filter( _lune.unwrap( node:get_nilBlock()), self, opt:nextOpt( node ) )
   end
   
end


function FormatterFilter:processWhen( node, opt )

   
   local symTxt = ""
   for index, symPair in pairs( node:get_symPairList() ) do
      symTxt = string.format( "%s %s", symTxt, symPair:get_src())
   end
   
   
   filter( node:get_block(), self, opt:nextOpt( node ) )
   do
      local _exp = node:get_elseBlock()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
end


function FormatterFilter:processDeclVar( node, opt )

   do
      local _switchExp = node:get_accessMode()
      if _switchExp == Ast.AccessMode.Pub then
         self:write( "pub " )
      end
   end
   
   
   do
      local _switchExp = node:get_mode()
      if _switchExp == Nodes.DeclVarMode.Let then
         self:write( "let " )
      elseif _switchExp == Nodes.DeclVarMode.Unwrap then
         self:write( "let! " )
      end
   end
   
   
   for index, sym in pairs( node:get_symbolInfoList() ) do
      if index > 1 then
         self:write( ", " )
      end
      
      if sym:get_mutable() then
         self:write( "mut " )
      end
      
      self:write( sym:get_name() )
   end
   
   
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         self:write( " = " )
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
   do
      local _exp = node:get_unwrapBlock()
      if _exp ~= nil then
         self:write( " " )
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
   do
      local _exp = node:get_thenBlock()
      if _exp ~= nil then
         self:write( "then" )
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
   do
      local _exp = node:get_syncBlock()
      if _exp ~= nil then
         self:write( "do" )
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
   self:writeln( ";" )
end



function FormatterFilter:processDeclArg( node, opt )

   self:write( node:get_symbolInfo():get_name() )
   self:write( ":" )
   
   self:write( node:get_expType():getTxt(  ) )
end


function FormatterFilter:processDeclArgDDD( node, opt )

   self:write( "..." )
end


function FormatterFilter:processExpDDD( node, opt )

   self:write( "..." )
end


function FormatterFilter:processExpSubDDD( node, opt )

end



function FormatterFilter:processDeclForm( node, opt )

end


function FormatterFilter:processDeclFuncInfo( node, declInfo, opt )

   local funcType = node:get_expType()
   if funcType:get_accessMode() == Ast.AccessMode.Pub then
      self:write( "pub " )
   end
   
   self:write( "fn " )
   
   if opt:get_parent():get_kind() ~= Nodes.NodeKind.get_DeclClass() and funcType:get_kind() == Ast.TypeInfoKind.Method then
      local classType = funcType:get_parentInfo()
      self:write( classType:get_rawTxt() )
      self:write( "." )
   end
   
   
   do
      local nameToken = declInfo:get_name()
      if nameToken ~= nil then
         self:write( nameToken.txt )
      end
   end
   
   self:write( "(" )
   
   local argList = declInfo:get_argList()
   if #argList ~= 0 then
      self:write( " " )
   end
   
   for index, arg in pairs( argList ) do
      if index > 1 then
         self:write( ", " )
      end
      
      filter( arg, self, opt:nextOpt( node ) )
   end
   
   if #argList ~= 0 then
      self:write( " " )
   end
   
   self:write( ")" )
   
   if Ast.TypeInfo.isMut( funcType ) and declInfo:get_kind() == Nodes.FuncKind.Mtd then
      self:write( " mut" )
   end
   
   
   if #funcType:get_retTypeInfoList() ~= 0 then
      self:write( " : " )
      for index, retType in pairs( funcType:get_retTypeInfoList() ) do
         if index > 1 then
            self:write( ", " )
         end
         
         self:write( retType:getTxt(  ) )
      end
      
   end
   
   
   do
      local _exp = declInfo:get_body()
      if _exp ~= nil then
         self:write( " " )
         filter( _exp, self, opt:nextOpt( node ) )
      else
         self:writeln( ";" )
      end
   end
   
end

function FormatterFilter:processDeclFunc( node, opt )

   self:processDeclFuncInfo( node, node:get_declInfo(), opt )
end


function FormatterFilter:processDeclMethod( node, opt )

   self:processDeclFuncInfo( node, node:get_declInfo(), opt )
end


function FormatterFilter:processDeclConstr( node, opt )

   self:processDeclFuncInfo( node, node:get_declInfo(), opt )
end



function FormatterFilter:processDeclDestr( node, opt )

   self:processDeclFuncInfo( node, node:get_declInfo(), opt )
end



function FormatterFilter:processExpCallSuper( node, opt )

   do
      local expListNode = node:get_expList()
      if expListNode ~= nil then
         self:write( "super(" )
         filter( expListNode, self, opt:nextOpt( node ) )
         self:write( ")" )
      else
         self:write( "super()" )
      end
   end
   
end


function FormatterFilter:processRefType( node, opt )

   self:write( node:get_expType():getTxt(  ) )
end


function FormatterFilter:processIf( node, opt )

   local stmtList = node:get_stmtList()
   for index, stmt in pairs( stmtList ) do
      do
         local _switchExp = stmt:get_kind()
         if _switchExp == Nodes.IfKind.If then
            self:write( "if " )
         elseif _switchExp == Nodes.IfKind.ElseIf then
            self:write( "elseif " )
         elseif _switchExp == Nodes.IfKind.Else then
            self:write( "else " )
         end
      end
      
      filter( stmt:get_exp(), self, opt:nextOpt( node ) )
      self:write( " " )
      filter( stmt:get_block(), self, opt:nextOpt( node ) )
   end
   
end


function FormatterFilter:processSwitch( node, opt )

   
   filter( node:get_exp(), self, opt:nextOpt( node ) )
   local caseList = node:get_caseList()
   for __index, caseInfo in pairs( caseList ) do
      filter( caseInfo:get_expList(), self, opt:nextOpt( node ) )
      filter( caseInfo:get_block(), self, opt:nextOpt( node ) )
   end
   
   do
      local _exp = node:get_default()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
end


function FormatterFilter:processMatch( node, opt )

   
   filter( node:get_val(), self, opt:nextOpt( node ) )
   local caseList = node:get_caseList()
   for __index, caseInfo in pairs( caseList ) do
      filter( caseInfo:get_block(), self, opt:nextOpt( node ) )
   end
   
   do
      local _exp = node:get_defaultBlock()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
end


function FormatterFilter:processWhile( node, opt )

   self:write( "while " )
   filter( node:get_exp(), self, opt:nextOpt( node ) )
   self:write( " " )
   filter( node:get_block(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processRepeat( node, opt )

   
   filter( node:get_block(), self, opt:nextOpt( node ) )
   filter( node:get_exp(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processFor( node, opt )

   self:write( string.format( "for %s = ", node:get_val():get_name()) )
   filter( node:get_init(), self, opt:nextOpt( node ) )
   self:write( ", " )
   filter( node:get_to(), self, opt:nextOpt( node ) )
   do
      local _exp = node:get_delta()
      if _exp ~= nil then
         self:write( ", " )
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
   self:write( " " )
   filter( node:get_block(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processApply( node, opt )

   self:write( "apply " )
   
   for index, var in pairs( node:get_varList() ) do
      if index > 1 then
         self:write( ", " )
      end
      
      self:write( var.txt )
   end
   
   self:write( " of " )
   filter( node:get_exp(), self, opt:nextOpt( node ) )
   self:write( " " )
   filter( node:get_block(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processForeach( node, opt )

   self:write( "foreach " )
   self:write( (_lune.unwrap( node:get_val()) ).txt )
   do
      local key = node:get_key()
      if key ~= nil then
         self:write( ", " )
         self:write( key.txt )
      end
   end
   
   self:write( " in " )
   filter( node:get_exp(), self, opt:nextOpt( node ) )
   self:writeln( "" )
   filter( node:get_block(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processForsort( node, opt )

   
   local index = ""
   do
      local _exp = node:get_key()
      if _exp ~= nil then
         index = _exp.txt
      end
   end
   
   filter( node:get_exp(), self, opt:nextOpt( node ) )
   filter( node:get_block(), self, opt:nextOpt( node ) )
end



function FormatterFilter:processExpUnwrap( node, opt )

   self:write( "unwrap " )
   filter( node:get_exp(), self, opt:nextOpt( node ) )
   do
      local _exp = node:get_default()
      if _exp ~= nil then
         self:write( " default " )
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
end


local function getTypeListTxt( typeList )

   local txt = ""
   for index, typeInfo in pairs( typeList ) do
      if index > 1 then
         txt = txt .. ", "
      end
      
      txt = txt .. typeInfo:getTxt(  )
   end
   
   return txt
end

function FormatterFilter:processExpCall( node, opt )

   filter( node:get_func(), self, opt:nextOpt( node ) )
   self:write( "(" )
   do
      local _exp = node:get_argList()
      if _exp ~= nil then
         self:write( " " )
         filter( _exp, self, opt:nextOpt( node ) )
         self:write( " " )
      end
   end
   
   self:write( ")" )
end


function FormatterFilter:processExpList( node, opt )

   local expList = node:get_expList()
   for index, exp in pairs( expList ) do
      if index > 1 then
         if exp:get_expType():get_kind() ~= Ast.TypeInfoKind.Abbr then
            self:write( ", " )
         else
          
            self:write( " " )
         end
         
      end
      
      filter( exp, self, opt:nextOpt( node ) )
   end
   
end


function FormatterFilter:processExpAccessMRet( node, opt )

end


function FormatterFilter:processExpOp1( node, opt )

   self:write( node:get_op().txt )
   do
      local _switchExp = node:get_op().txt
      if _switchExp == "not" then
         self:write( " " )
      end
   end
   
   filter( node:get_exp(), self, opt:nextOpt( node ) )
end



function FormatterFilter:processExpToDDD( node, opt )

   
   filter( node:get_expList(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processExpMultiTo1( node, opt )

   
   filter( node:get_exp(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processExpCast( node, opt )

   
   filter( node:get_exp(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processExpParen( node, opt )

   
   filter( node:get_exp(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processExpOp2( node, opt )

   filter( node:get_exp1(), self, opt:nextOpt( node ) )
   self:write( string.format( " %s ", node:get_op().txt) )
   filter( node:get_exp2(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processExpNew( node, opt )

   self:write( "new " )
   
   filter( node:get_symbol(), self, opt:nextOpt( node ) )
   self:write( "(" )
   do
      local _exp = node:get_argList()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
   self:write( ")" )
end


function FormatterFilter:processExpRef( node, opt )

   self:write( node:get_token().txt )
end


function FormatterFilter:processExpRefItem( node, opt )

   filter( node:get_val(), self, opt:nextOpt( node ) )
   if node:get_nilAccess() then
      self:write( "$" )
   end
   
   do
      local _exp = node:get_index()
      if _exp ~= nil then
         self:write( "[ " )
         filter( _exp, self, opt:nextOpt( node ) )
         self:write( " ]" )
      else
         do
            local _exp = node:get_symbol()
            if _exp ~= nil then
               self:write( "." )
               self:write( _exp )
            end
         end
         
      end
   end
   
end


function FormatterFilter:processRefField( node, opt )

   filter( node:get_prefix(), self, opt:nextOpt( node ) )
   self:write( "." )
   self:write( node:get_field().txt )
end


function FormatterFilter:processExpOmitEnum( node, opt )

   self:write( "." )
   self:write( node:get_valToken().txt )
end



function FormatterFilter:processGetField( node, opt )

   filter( node:get_prefix(), self, opt:nextOpt( node ) )
   if node:get_nilAccess() then
      self:write( "$" )
   end
   
   self:write( ".$" )
   self:write( node:get_field().txt )
end



function FormatterFilter:processReturn( node, opt )

   self:write( "return" )
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         self:write( " " )
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
   self:writeln( ";" )
end



function FormatterFilter:processProvide( node, opt )

end


function FormatterFilter:processAlias( node, opt )

end


function FormatterFilter:processTestBlock( node, opt )

   
   filter( node:get_block(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processBoxing( node, opt )

   
   filter( node:get_src(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processUnboxing( node, opt )

   
   filter( node:get_src(), self, opt:nextOpt( node ) )
end


function FormatterFilter:processLiteralList( node, opt )

   self:write( "[" )
   
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         self:write( " " )
         filter( _exp, self, opt:nextOpt( node ) )
         self:write( " " )
      end
   end
   
   
   self:write( "]" )
end


function FormatterFilter:processLiteralSet( node, opt )

   self:write( "(@" )
   
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         self:write( " " )
         filter( _exp, self, opt:nextOpt( node ) )
         self:write( " " )
      end
   end
   
   
   self:write( ")" )
end


function FormatterFilter:processLiteralMap( node, opt )

   
   local pairList = node:get_pairList()
   for __index, pair in pairs( pairList ) do
      filter( pair:get_key(), self, opt:nextOpt( node ) )
      filter( pair:get_val(), self, opt:nextOpt( node ) )
   end
   
end


function FormatterFilter:processLiteralArray( node, opt )

   
   do
      local _exp = node:get_expList()
      if _exp ~= nil then
         filter( _exp, self, opt:nextOpt( node ) )
      end
   end
   
end


function FormatterFilter:processLiteralChar( node, opt )

end


function FormatterFilter:processLiteralInt( node, opt )

   self:write( node:get_token().txt )
end


function FormatterFilter:processLiteralReal( node, opt )

end


function FormatterFilter:processLiteralString( node, opt )

   self:write( node:get_token().txt )
   
   do
      local expList = node:get_expList()
      if expList ~= nil then
         self:write( " ( " )
         filter( expList, self, opt:nextOpt( node ) )
         self:write( " )" )
      end
   end
   
end


function FormatterFilter:processLiteralBool( node, opt )

   self:write( node:get_token().txt )
end


function FormatterFilter:processLiteralNil( node, opt )

   self:write( "nil" )
end


function FormatterFilter:processBreak( node, opt )

   self:writeln( "break;" )
end


function FormatterFilter:processLiteralSymbol( node, opt )

end



function FormatterFilter:processAbbr( node, opt )

   self:write( "##" )
end



return _moduleObj
