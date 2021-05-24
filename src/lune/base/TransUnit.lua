--lune/base/TransUnit.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@TransUnit'
local _lune = {}
if _lune3 then
   _lune = _lune3
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

if not _lune3 then
   _lune3 = _lune
end





local Types = _lune.loadModule( 'lune.base.Types' )
local Meta = _lune.loadModule( 'lune.base.Meta' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Writer = _lune.loadModule( 'lune.base.Writer' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Option = _lune.loadModule( 'lune.base.Option' )
local Code = _lune.loadModule( 'lune.base.Code' )
local Log = _lune.loadModule( 'lune.base.Log' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )
local Macro = _lune.loadModule( 'lune.base.Macro' )
local TransUnitIF = _lune.loadModule( 'lune.base.TransUnitIF' )
local Builtin = _lune.loadModule( 'lune.base.Builtin' )
local Import = _lune.loadModule( 'lune.base.Import' )



local DeclClassMode = TransUnitIF.DeclClassMode
local DeclFuncMode = {}
DeclFuncMode._val2NameMap = {}
function DeclFuncMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "DeclFuncMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function DeclFuncMode._from( val )
   if DeclFuncMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
DeclFuncMode.__allList = {}
function DeclFuncMode.get__allList()
   return DeclFuncMode.__allList
end

DeclFuncMode.Func = 0
DeclFuncMode._val2NameMap[0] = 'Func'
DeclFuncMode.__allList[1] = DeclFuncMode.Func
DeclFuncMode.Class = 1
DeclFuncMode._val2NameMap[1] = 'Class'
DeclFuncMode.__allList[2] = DeclFuncMode.Class
DeclFuncMode.Module = 2
DeclFuncMode._val2NameMap[2] = 'Module'
DeclFuncMode.__allList[3] = DeclFuncMode.Module
DeclFuncMode.Glue = 3
DeclFuncMode._val2NameMap[3] = 'Glue'
DeclFuncMode.__allList[4] = DeclFuncMode.Glue


local ExpSymbolMode = {}
ExpSymbolMode._val2NameMap = {}
function ExpSymbolMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ExpSymbolMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ExpSymbolMode._from( val )
   if ExpSymbolMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ExpSymbolMode.__allList = {}
function ExpSymbolMode.get__allList()
   return ExpSymbolMode.__allList
end

ExpSymbolMode.Symbol = 0
ExpSymbolMode._val2NameMap[0] = 'Symbol'
ExpSymbolMode.__allList[1] = ExpSymbolMode.Symbol
ExpSymbolMode.Fn = 1
ExpSymbolMode._val2NameMap[1] = 'Fn'
ExpSymbolMode.__allList[2] = ExpSymbolMode.Fn
ExpSymbolMode.Field = 2
ExpSymbolMode._val2NameMap[2] = 'Field'
ExpSymbolMode.__allList[3] = ExpSymbolMode.Field
ExpSymbolMode.FieldNil = 3
ExpSymbolMode._val2NameMap[3] = 'FieldNil'
ExpSymbolMode.__allList[4] = ExpSymbolMode.FieldNil
ExpSymbolMode.Get = 4
ExpSymbolMode._val2NameMap[4] = 'Get'
ExpSymbolMode.__allList[5] = ExpSymbolMode.Get
ExpSymbolMode.GetNil = 5
ExpSymbolMode._val2NameMap[5] = 'GetNil'
ExpSymbolMode.__allList[6] = ExpSymbolMode.GetNil


local AnalyzeMode = {}
_moduleObj.AnalyzeMode = AnalyzeMode
AnalyzeMode._val2NameMap = {}
function AnalyzeMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "AnalyzeMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function AnalyzeMode._from( val )
   if AnalyzeMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
AnalyzeMode.__allList = {}
function AnalyzeMode.get__allList()
   return AnalyzeMode.__allList
end

AnalyzeMode.Compile = 0
AnalyzeMode._val2NameMap[0] = 'Compile'
AnalyzeMode.__allList[1] = AnalyzeMode.Compile
AnalyzeMode.Diag = 1
AnalyzeMode._val2NameMap[1] = 'Diag'
AnalyzeMode.__allList[2] = AnalyzeMode.Diag
AnalyzeMode.Complete = 2
AnalyzeMode._val2NameMap[2] = 'Complete'
AnalyzeMode.__allList[3] = AnalyzeMode.Complete
AnalyzeMode.Inquire = 3
AnalyzeMode._val2NameMap[3] = 'Inquire'
AnalyzeMode.__allList[4] = AnalyzeMode.Inquire


local DefaultAsyncMode = {}
DefaultAsyncMode._val2NameMap = {}
function DefaultAsyncMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "DefaultAsyncMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function DefaultAsyncMode._from( val )
   if DefaultAsyncMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
DefaultAsyncMode.__allList = {}
function DefaultAsyncMode.get__allList()
   return DefaultAsyncMode.__allList
end

DefaultAsyncMode.NoAsync = 0
DefaultAsyncMode._val2NameMap[0] = 'NoAsync'
DefaultAsyncMode.__allList[1] = DefaultAsyncMode.NoAsync
DefaultAsyncMode.AsyncFunc = 1
DefaultAsyncMode._val2NameMap[1] = 'AsyncFunc'
DefaultAsyncMode.__allList[2] = DefaultAsyncMode.AsyncFunc
DefaultAsyncMode.AsyncAll = 2
DefaultAsyncMode._val2NameMap[2] = 'AsyncAll'
DefaultAsyncMode.__allList[3] = DefaultAsyncMode.AsyncAll


local AccessSymPos = {}
function AccessSymPos.setmeta( obj )
  setmetatable( obj, { __index = AccessSymPos  } )
end
function AccessSymPos.new( symbol, pos )
   local obj = {}
   AccessSymPos.setmeta( obj )
   if obj.__init then
      obj:__init( symbol, pos )
   end
   return obj
end
function AccessSymPos:__init( symbol, pos )

   self.symbol = symbol
   self.pos = pos
end
function AccessSymPos:get_symbol()
   return self.symbol
end
function AccessSymPos:get_pos()
   return self.pos
end


local RefAccessSymPos = {}
function RefAccessSymPos.setmeta( obj )
  setmetatable( obj, { __index = RefAccessSymPos  } )
end
function RefAccessSymPos.new( symbol, pos )
   local obj = {}
   RefAccessSymPos.setmeta( obj )
   if obj.__init then
      obj:__init( symbol, pos )
   end
   return obj
end
function RefAccessSymPos:__init( symbol, pos )

   self.symbol = symbol
   self.pos = pos
end
function RefAccessSymPos:get_symbol()
   return self.symbol
end
function RefAccessSymPos:get_pos()
   return self.pos
end


local function clearThePosForModToRef( scope, moduleScope )

   local list = {}
   scope:filterSymbolTypeInfo( scope, moduleScope, Ast.ScopeAccess.Normal, function ( symInfo )
   
      if symInfo:get_kind() == Ast.SymbolKind.Var then
         table.insert( list, RefAccessSymPos.new(symInfo, symInfo:get_posForModToRef()) )
         symInfo:set_posForModToRef( nil )
      end
      
      return true
   end )
   return list
end

local TentativeSymbol = {}
function TentativeSymbol.new( parent, scope, moduleScope, loopFlag, refAccessSymPosList )
   local obj = {}
   TentativeSymbol.setmeta( obj )
   if obj.__init then obj:__init( parent, scope, moduleScope, loopFlag, refAccessSymPosList ); end
   return obj
end
function TentativeSymbol:__init(parent, scope, moduleScope, loopFlag, refAccessSymPosList) 
   self.symbolSet = {}
   self.oldSymbolSet = nil
   self.parent = parent
   self.scope = scope
   self.skipFlag = false
   self.loopFlag = loopFlag
   self.accessSymList = {}
   self.initSymSet = {}
   
   self.accessSymSet = {}
   self.modSymbolSet = {}
   
   local list
   
   if refAccessSymPosList ~= nil then
      list = refAccessSymPosList
   else
      if loopFlag then
         list = clearThePosForModToRef( scope, moduleScope )
      else
       
         list = {}
      end
      
   end
   
   self.sym2posForModToRef = list
end
function TentativeSymbol:modSym( accessSym )

   self.modSymbolSet[accessSym:getOrg(  )]= true
end
function TentativeSymbol:addAccessSym( accessSym )

   self.accessSymSet[accessSym:getOrg(  )]= true
end
function TentativeSymbol:addAccessSymPos( accessSym )

   table.insert( self.accessSymList, accessSym )
end
function TentativeSymbol:clearAccessSym(  )

   if #self.accessSymList ~= 0 then
      self.accessSymList = {}
   end
   
end
function TentativeSymbol:checkAndExclude( symbolInfo )

   symbolInfo = symbolInfo:getOrg(  )
   if _lune._Set_has(self.symbolSet, symbolInfo ) then
      self.symbolSet[symbolInfo]= nil
      return true
   end
   
   return false
end
function TentativeSymbol:regist( symbolInfo, pos )

   self.symbolSet[symbolInfo:getOrg(  )]= true
   self.initSymSet[symbolInfo:getOrg(  )]= true
   symbolInfo:set_hasValueFlag( true )
   
   if self:get_scope():isInnerOf( symbolInfo:get_scope() ) then
      if not symbolInfo:get_mutable() then
         local work = self
         while true do
            do
               local _exp = work.parent
               if _exp ~= nil then
                  if work.scope == symbolInfo:get_scope() then
                     break
                  end
                  
                  if work.loopFlag then
                     return false
                  end
                  
                  work = _exp
               else
                  break
               end
            end
            
         end
         
      end
      
   end
   
   
   return true
end
function TentativeSymbol:skip(  )

   for symbolInfo, __val in pairs( self.symbolSet ) do
      symbolInfo:clearValue(  )
   end
   
   self.skipFlag = true
end
function TentativeSymbol:merge( finishFlag )

   if self.skipFlag then
      self.skipFlag = false
      do
         local other = self.oldSymbolSet
         if other ~= nil then
            self.symbolSet = _lune._Set_clone(other )
         end
      end
      
      if finishFlag then
         for symbolInfo, __val in pairs( self.symbolSet ) do
            symbolInfo:updateValue( symbolInfo:get_posForLatestMod() )
         end
         
      end
      
      return self.oldSymbolSet ~= nil
   end
   
   do
      local other = self.oldSymbolSet
      if other ~= nil then
         local mergedSet = _lune._Set_and(_lune._Set_clone(self.symbolSet ), other )
         if finishFlag then
            for symbolInfo, __val in pairs( _lune._Set_sub(_lune._Set_or(_lune._Set_clone(self.symbolSet ), other ), mergedSet ) ) do
               symbolInfo:clearValue(  )
            end
            
         else
          
            for symbolInfo, __val in pairs( _lune._Set_or(_lune._Set_clone(self.symbolSet ), other ) ) do
               symbolInfo:clearValue(  )
            end
            
         end
         
         self.symbolSet = mergedSet
      else
         if not finishFlag then
            for symbolInfo, __val in pairs( self.symbolSet ) do
               symbolInfo:clearValue(  )
            end
            
         end
         
      end
   end
   
   return true
end
function TentativeSymbol:syncPos(  )

   if self.loopFlag then
      for __index, info in ipairs( self.sym2posForModToRef ) do
         local symbol = info:get_symbol()
         if symbol:get_posForModToRef() then
            symbol:set_posForModToRef( symbol:get_posForLatestMod() )
         else
          
            symbol:set_posForModToRef( info:get_pos() )
         end
         
      end
      
   end
   
end
function TentativeSymbol:finish( complete )

   self:syncPos(  )
   self:merge( true )
   do
      local parent = self.parent
      if parent ~= nil then
         if complete then
            for symbolInfo, __val in pairs( self.symbolSet ) do
               if symbolInfo:get_hasValueFlag() then
                  if parent.scope:isInnerOf( symbolInfo:get_scope() ) then
                     parent.symbolSet[symbolInfo]= true
                  end
                  
               end
               
            end
            
         else
          
            for symbolInfo, __val in pairs( self.symbolSet ) do
               symbolInfo:clearValue(  )
            end
            
         end
         
         
         
         
         for symbolInfo, __val in pairs( self.initSymSet ) do
            if symbolInfo:get_scope() ~= self.scope then
               parent.initSymSet[symbolInfo]= true
            end
            
         end
         
         
         for symbolInfo, __val in pairs( self.accessSymSet ) do
            if symbolInfo:get_scope() ~= self.scope then
               parent.accessSymSet[symbolInfo]= true
            end
            
         end
         
         
         for symbolInfo, __val in pairs( self.modSymbolSet ) do
            if symbolInfo:get_scope() ~= self.scope then
               parent.modSymbolSet[symbolInfo]= true
            end
            
         end
         
         
         return parent
      end
   end
   
   
   return nil
end
function TentativeSymbol:newSet( scope )

   if self:merge( false ) then
      self.oldSymbolSet = self.symbolSet
   end
   
   self.symbolSet = {}
   self.scope = scope
end
function TentativeSymbol.setmeta( obj )
  setmetatable( obj, { __index = TentativeSymbol  } )
end
function TentativeSymbol:get_initSymSet()
   return self.initSymSet
end
function TentativeSymbol:get_scope()
   return self.scope
end
function TentativeSymbol:get_accessSymList()
   return self.accessSymList
end


local AnalyzingState = {}
AnalyzingState._val2NameMap = {}
function AnalyzingState:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "AnalyzingState.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function AnalyzingState._from( val )
   if AnalyzingState._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
AnalyzingState.__allList = {}
function AnalyzingState.get__allList()
   return AnalyzingState.__allList
end

AnalyzingState.Other = 0
AnalyzingState._val2NameMap[0] = 'Other'
AnalyzingState.__allList[1] = AnalyzingState.Other
AnalyzingState.Constructor = 1
AnalyzingState._val2NameMap[1] = 'Constructor'
AnalyzingState.__allList[2] = AnalyzingState.Constructor
AnalyzingState.InitBlock = 2
AnalyzingState._val2NameMap[2] = 'InitBlock'
AnalyzingState.__allList[3] = AnalyzingState.InitBlock
AnalyzingState.ClassMethod = 3
AnalyzingState._val2NameMap[3] = 'ClassMethod'
AnalyzingState.__allList[4] = AnalyzingState.ClassMethod
AnalyzingState.Func = 4
AnalyzingState._val2NameMap[4] = 'Func'
AnalyzingState.__allList[5] = AnalyzingState.Func



local ClosureFun = {}
function ClosureFun:check(  )

   do
      local _exp = _lune.nilacc( self.symbol:get_typeInfo():get_scope(), 'get_closureSymList', 'callmtd' )
      if _exp ~= nil then
         if #_exp > 0 then
            self.fromScope:setClosure( self.symbol )
            return true
         end
         
      end
   end
   
   return false
end
function ClosureFun.checkList( list )

   local workList = list
   local remainList = {}
   while true do
      local update = false
      for __index, closureFun in ipairs( workList ) do
         if closureFun:check(  ) then
            update = true
         else
          
            table.insert( remainList, closureFun )
         end
         
      end
      
      if not update then
         break
      end
      
      workList = remainList
      remainList = {}
   end
   
end
function ClosureFun.setmeta( obj )
  setmetatable( obj, { __index = ClosureFun  } )
end
function ClosureFun.new( symbol, fromScope )
   local obj = {}
   ClosureFun.setmeta( obj )
   if obj.__init then
      obj:__init( symbol, fromScope )
   end
   return obj
end
function ClosureFun:__init( symbol, fromScope )

   self.symbol = symbol
   self.fromScope = fromScope
end


local AccessSymbolSet = {}
function AccessSymbolSet.new(  )
   local obj = {}
   AccessSymbolSet.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function AccessSymbolSet:__init() 
   self.accessSym2Pos = {}
end
function AccessSymbolSet:add( symbol )

   self.accessSym2Pos[symbol] = symbol:get_posForLatestMod()
end
function AccessSymbolSet:applyPos( excludeSymList )

   local set = excludeSymList
   if  nil == set then
      local _set = set
   
      set = {}
   end
   
   for symbol, pos in pairs( self.accessSym2Pos ) do
      if not _lune._Set_has(set, symbol:getOrg(  ) ) then
         symbol:set_posForModToRef( pos )
      end
      
   end
   
end
function AccessSymbolSet.setmeta( obj )
  setmetatable( obj, { __index = AccessSymbolSet  } )
end
function AccessSymbolSet:get_accessSym2Pos()
   return self.accessSym2Pos
end


local AccessSymbolSetQueue = {}
function AccessSymbolSetQueue.new(  )
   local obj = {}
   AccessSymbolSetQueue.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function AccessSymbolSetQueue:__init() 
   self.queue = {}
end
function AccessSymbolSetQueue:push(  )

   table.insert( self.queue, AccessSymbolSet.new() )
end
function AccessSymbolSetQueue:pop( symbolList )

   self.queue[#self.queue]:applyPos( symbolList )
   table.remove( self.queue )
end
function AccessSymbolSetQueue:add( symbol )

   self.queue[#self.queue]:add( symbol )
end
function AccessSymbolSetQueue:getMap(  )

   return self.queue[#self.queue]:get_accessSym2Pos()
end
function AccessSymbolSetQueue.setmeta( obj )
  setmetatable( obj, { __index = AccessSymbolSetQueue  } )
end


local TransUnit = {}
setmetatable( TransUnit, { ifList = {TransUnitIF.TransUnitIF,Parser.PushbackParser,} } )
_moduleObj.TransUnit = TransUnit
function TransUnit:get_scope(  )

   return self.scope
end
function TransUnit.new( moduleId, importModuleInfo, macroEval, analyzeModule, mode, pos, targetLuaVer, ctrl_info )
   local obj = {}
   TransUnit.setmeta( obj )
   if obj.__init then obj:__init( moduleId, importModuleInfo, macroEval, analyzeModule, mode, pos, targetLuaVer, ctrl_info ); end
   return obj
end
function TransUnit:__init(moduleId, importModuleInfo, macroEval, analyzeModule, mode, pos, targetLuaVer, ctrl_info) 
   self.class2defaultAsyncMode = {}
   self.defaultAsyncMode = DefaultAsyncMode.NoAsync
   self.importedAliasMap = {}
   self.importModuleSet = {}
   self.processInfo = Ast.createProcessInfo( ctrl_info.validCheckingMutable, ctrl_info.validLuaval, ctrl_info.validAstDetailError )
   self.accessSymbolSetQueue = AccessSymbolSetQueue.new()
   self.advertisedTypeSet = {}
   self.closureFunList = {}
   self.scopeAccess = Ast.ScopeAccess.Normal
   self.macroCtrl = Macro.MacroCtrl.new(macroEval)
   self.typeNameCtrl = Ast.defaultTypeNameCtrl
   self.protoClassMap = {}
   self.analyzingStateQueue = {}
   self.ctrl_info = ctrl_info
   self.ignoreToCheckSymbol_ = false
   self.moduleId = moduleId
   self.helperInfo = frontInterface.LuneHelperInfo.new()
   self.targetLuaVer = targetLuaVer
   self.importModuleInfo = importModuleInfo
   self.protoFuncMap = {}
   self.loopScopeQueue = {}
   self.has__func__Symbol = {}
   self.nodeManager = Nodes.NodeManager.new()
   self.macroScope = nil
   self.validMutControl = true
   self.modifier = TransUnitIF.Modifier.new(self.validMutControl, self.processInfo)
   self.moduleName = ""
   self.moduleType = Ast.headTypeInfo
   self.parser = Parser.DefaultPushbackParser.new(Parser.DummyParser.new())
   self.subfileList = {}
   self.globalScope = Ast.Scope.new(self.processInfo, Ast.rootScope, false, nil)
   
   self.scope = Ast.Scope.new(self.processInfo, self.globalScope, true, nil)
   self.topScope = self.scope
   self.moduleScope = self.scope
   
   self.tentativeSymbol = TentativeSymbol.new(nil, self.globalScope, self.moduleScope, false, nil)
   
   self.typeId2ClassMap = {}
   self.typeInfo2ClassNode = {}
   self.commentCtrl = Parser.CommentCtrl.new()
   self.errMessList = {}
   self.warnMessList = {}
   self.analyzeMode = _lune.unwrapDefault( mode, AnalyzeMode.Compile)
   self.analyzePos = _lune.unwrapDefault( pos, self:createPosition( 0, 0 ))
   self.analyzeModule = _lune.unwrapDefault( analyzeModule, "")
   self.provideNode = nil
   
   self.importCtrl = nil
end
function TransUnit:getLatestPos(  )

   return self.parser:get_currentToken().pos
end
function TransUnit:pushAnalyzingState( state )

   table.insert( self.analyzingStateQueue, state )
end
function TransUnit:popAnalyzingState(  )

   if #self.analyzingStateQueue == 0 then
      self:error( "underflow analyzingStateQueue" )
   end
   
   table.remove( self.analyzingStateQueue )
end
function TransUnit:inAnalyzingState( state )

   if #self.analyzingStateQueue > 0 then
      return self.analyzingStateQueue[#self.analyzingStateQueue] == state
   end
   
   return false
end
function TransUnit:addErrMess( pos, mess )

   if mess:find( "type mismatch.*<- &" ) then
      mess = mess .. ". if your code is the old style, use the opiton '--legacy-mutable-control'."
   end
   
   table.insert( self.errMessList, string.format( "%s:%d:%d: error: %s", pos.streamName, pos.lineNo, pos.column, mess) )
end
function TransUnit:addWarnMess( pos, mess )

   table.insert( self.warnMessList, string.format( "%s:%d:%d: warning: %s", self.parser:getStreamName(  ), pos.lineNo, pos.column, mess) )
end
function TransUnit:addWarnErrMess( pos, err, mess )

   if err then
      self:addErrMess( pos, mess )
   else
    
      self:addWarnMess( pos, mess )
   end
   
end
function TransUnit:pushScope( classFlag, baseInfo, interfaceList )

   self.scope = Ast.TypeInfo.createScope( self.processInfo, self.scope, classFlag, baseInfo, interfaceList )
   return self.scope
end
function TransUnit:popScope(  )

   self.scope = self.scope:get_parent(  )
end
function TransUnit:prepareTentativeSymbol( scope, loopFlag, list )

   self.tentativeSymbol = TentativeSymbol.new(self.tentativeSymbol, scope, self.moduleScope, loopFlag, list)
end
function TransUnit:checkAccesSym(  )

   for __index, accessSym in ipairs( self.tentativeSymbol:get_accessSymList() ) do
      local symbolInfo = accessSym:get_symbol()
      if not symbolInfo:get_hasValueFlag() then
         self:addErrMess( accessSym:get_pos(), string.format( "This can't access variable have no value -- %s", symbolInfo:get_name()) )
      end
      
   end
   
   
   self.tentativeSymbol:clearAccessSym(  )
end
function TransUnit:finishTentativeSymbol( complete )

   self:checkAccesSym(  )
   local tentativeSymbol = self.tentativeSymbol
   self.tentativeSymbol = _lune.unwrap( tentativeSymbol:finish( complete ))
   
   do
      local errSymMap = {}
      for symbolInfo, __val in pairs( tentativeSymbol:get_initSymSet() ) do
         if tentativeSymbol:get_scope():get_parent() == symbolInfo:get_scope() then
            if not symbolInfo:get_hasValueFlag() then
               errSymMap[symbolInfo:get_name()] = symbolInfo
            end
            
         end
         
      end
      
      do
         local __sorted = {}
         local __map = errSymMap
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, __key in ipairs( __sorted ) do
            local symbolInfo = __map[ __key ]
            do
               self:addErrMess( self.parser:getLastPos(  ), string.format( "There is the case no initialized value for '%s'", symbolInfo:get_name()) )
            end
         end
      end
      
   end
   
   
   if tentativeSymbol:get_scope():get_validCheckingUnaccess() then
      do
         local __sorted = {}
         local __map = tentativeSymbol:get_scope():get_symbol2SymbolInfoMap()
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, __key in ipairs( __sorted ) do
            local symbolInfo = __map[ __key ]
            do
               if (not symbolInfo:get_posForModToRef() or symbolInfo:get_posForModToRef() ~= symbolInfo:get_posForLatestMod() ) and (symbolInfo:get_kind() == Ast.SymbolKind.Var or symbolInfo:get_kind() == Ast.SymbolKind.Fun ) and not Ast.isPubToExternal( symbolInfo:get_accessMode() ) and not symbolInfo:get_name():find( "^_" ) then
                  do
                     local pos = symbolInfo:get_posForLatestMod() or symbolInfo:get_pos()
                     if pos ~= nil then
                        self:addWarnMess( pos, string.format( "'%s' var isn't accessed", symbolInfo:get_name()) )
                     end
                  end
                  
               end
               
            end
         end
      end
      
   end
   
end
function TransUnit:mergeTentativeSymbol( scope )

   self:checkAccesSym(  )
   self.tentativeSymbol:newSet( scope )
end
function TransUnit:getCurrentClass(  )

   local typeInfo = Ast.headTypeInfo
   local scope = self.scope
   repeat 
      do
         local _exp = scope:get_ownerTypeInfo()
         if _exp ~= nil then
            if _exp:get_kind() == Ast.TypeInfoKind.Class or _exp:get_kind() == Ast.TypeInfoKind.Module or _exp:get_kind() == Ast.TypeInfoKind.IF then
               return _exp
            end
            
         end
      end
      
      scope = scope:get_parent()
   until scope:isRoot(  )
   return typeInfo
end
function TransUnit:getCurrentNamespaceTypeInfo(  )

   return self.scope:getNamespaceTypeInfo(  )
end
function TransUnit:getCurrentNamespaceScope(  )

   return _lune.unwrap( self:getCurrentNamespaceTypeInfo(  ):get_scope())
end
function TransUnit:pushModule( processInfo, externalFlag, name, mutable )

   local typeInfo = Ast.headTypeInfo
   
   local modName
   
   if name:find( "^@" ) then
      modName = name
   else
    
      modName = string.format( "@%s", name)
   end
   
   
   do
      local _exp = self.scope:getTypeInfoChild( modName )
      if _exp ~= nil then
         typeInfo = _exp
         self.scope = _lune.unwrap( Ast.getScope( typeInfo ))
      else
         local _
         local parentInfo = self:getCurrentNamespaceTypeInfo(  )
         local parentScope = self.scope
         local scope = self:pushScope( true )
         typeInfo = processInfo:createModule( scope, parentInfo, externalFlag, modName, mutable )
         
         local _508, existSym = parentScope:addClass( processInfo, modName, nil, typeInfo )
         if existSym ~= nil then
            self:addErrMess( self.parser:getLastPos(  ), string.format( "module symbols exist -- %s.%s -- %s.%s", existSym:get_namespaceTypeInfo():getFullName( self.typeNameCtrl, parentScope, false ), existSym:get_name(), parentInfo:getFullName( self.typeNameCtrl, parentScope, false ), modName) )
         end
         
      end
   end
   
   if not self.typeId2ClassMap[typeInfo:get_typeId()] then
      local namespace = Nodes.NamespaceInfo.new(modName, self.scope, typeInfo)
      self.typeId2ClassMap[typeInfo:get_typeId()] = namespace
   end
   
   return typeInfo
end
function TransUnit:popModule(  )

   self:popScope(  )
end
function TransUnit:pushExtModule( externalFlag, name, accessMode, pos, lazy, lang, requirePath )

   local typeInfo = Ast.headTypeInfo
   
   local modName = name
   do
      local _exp = self.scope:getTypeInfoChild( modName )
      if _exp ~= nil then
         self:addErrMess( pos, string.format( "multiple define -- %s", name) )
         return _exp
      else
         local parentInfo = self:getCurrentNamespaceTypeInfo(  )
         local parentScope = self.scope
         local scope = self:pushScope( true )
         typeInfo = self.processInfo:createExtModule( scope, parentInfo, externalFlag, accessMode, name, lang, requirePath )
         
         parentScope:addExtModule( self.processInfo, name, pos, typeInfo, lazy, lang )
      end
   end
   
   if not self.typeId2ClassMap[typeInfo:get_typeId()] then
      local namespace = Nodes.NamespaceInfo.new(modName, self.scope, typeInfo)
      self.typeId2ClassMap[typeInfo:get_typeId()] = namespace
   end
   
   return typeInfo
end
function TransUnit:pushClassScope( errPos, classTypeInfo )

   if self.scope ~= _lune.nilacc( classTypeInfo:get_scope(), 'get_parent', 'callmtd' ) then
      local classParentName
      
      local classParentTypeId
      
      do
         local parentType = _lune.nilacc( _lune.nilacc( classTypeInfo:get_scope(), 'get_parent', 'callmtd' ), 'get_ownerTypeInfo', 'callmtd' )
         if parentType ~= nil then
            classParentName = parentType:getFullName( self.typeNameCtrl, _lune.unwrap( parentType:get_scope()), false )
            classParentTypeId = parentType:get_typeId()
         else
            classParentName = "nil"
            classParentTypeId = Ast.dummyIdInfo
         end
      end
      
      self:addErrMess( errPos, string.format( "This class does not exist in this scope. -- %s -- %s(%d), %s(%d)", classTypeInfo:getTxt(  ), _lune.nilacc( self.scope:get_ownerTypeInfo(), 'getFullName', 'callmtd' , self.typeNameCtrl, self.scope, false ), _lune.nilacc( _lune.nilacc( self.scope:get_ownerTypeInfo(), 'get_typeId', 'callmtd' ), "id" ) or -1, classParentName, classParentTypeId.id) )
   end
   
   self.scope = _lune.unwrap( Ast.getScope( classTypeInfo ))
end
function TransUnit:pushClass( processInfo, errPos, mode, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace )

   local typeInfo = Ast.headTypeInfo
   do
      local _exp = self.scope:getTypeInfo( name, self.scope, true, Ast.ScopeAccess.Normal )
      if _exp ~= nil then
         typeInfo = _exp
      end
   end
   
   
   if typeInfo ~= Ast.headTypeInfo then
      if _lune.nilacc( typeInfo:get_scope(), 'get_parent', 'callmtd' ) ~= self.scope then
         self:addErrMess( errPos, string.format( "multiple class(%s)", typeInfo:getTxt( self.typeNameCtrl )) )
         self:error( "stop by error" )
      end
      
   end
   
   if typeInfo ~= Ast.headTypeInfo then
      if typeInfo:get_abstractFlag() ~= abstractFlag then
         self:addErrMess( errPos, string.format( "mismatch class(%s) abstract for prototpye", typeInfo:getTxt( self.typeNameCtrl )) )
      end
      
      if typeInfo:get_accessMode() ~= accessMode then
         self:addErrMess( errPos, string.format( "mismatch class(%s) accessmode(%s) for prototpye accessmode(%s)", typeInfo:getTxt( self.typeNameCtrl ), Ast.AccessMode:_getTxt( accessMode)
         , Ast.AccessMode:_getTxt( typeInfo:get_accessMode())
         ) )
      end
      
      if baseInfo ~= nil then
         if typeInfo:get_baseTypeInfo() ~= baseInfo then
            self:addErrMess( errPos, string.format( "mismatch class(%s) base class(%s) for prototpye base class(%s)", typeInfo:getTxt( self.typeNameCtrl ), baseInfo:getTxt(  ), typeInfo:get_baseTypeInfo():getTxt(  )) )
         end
         
      else
         if typeInfo:hasBase(  ) then
            self:addErrMess( errPos, string.format( "mismatch class(%s) base class(None) for prototpye base class(%s)", typeInfo:getTxt( self.typeNameCtrl ), typeInfo:get_baseTypeInfo():getTxt(  )) )
         end
         
      end
      
      
      local function compareList( protoList, typeList, message )
      
         if #protoList == #typeList then
            for index, protoType in ipairs( protoList ) do
               if protoType ~= typeList[index] then
                  self:addErrMess( errPos, string.format( "mismatch class(%s) %s(%s) for prototpye %s(%s)", typeInfo:getTxt( self.typeNameCtrl ), message, typeList[index]:getTxt( self.typeNameCtrl ), message, protoType:getTxt(  )) )
               end
               
            end
            
         else
          
            self:addErrMess( errPos, string.format( "mismatch class(%s) %s(%d) for prototpye %s(%d)", typeInfo:getTxt( self.typeNameCtrl ), message, #typeList, message, #protoList) )
         end
         
      end
      
      compareList( typeInfo:get_interfaceList(), _lune.unwrapDefault( interfaceList, {}), "interface" )
      
      compareList( typeInfo:get_itemTypeInfoList(), _lune.unwrapDefault( genTypeList, {}), "generics" )
      
      self.scope = _lune.unwrap( Ast.getScope( typeInfo ))
      do
         local _switchExp = (typeInfo:get_kind() )
         if _switchExp == Ast.TypeInfoKind.Class then
            if mode == TransUnitIF.DeclClassMode.Interface then
               self:addErrMess( errPos, string.format( "define interface already -- %s", name) )
               Util.printStackTrace(  )
            end
            
         elseif _switchExp == Ast.TypeInfoKind.IF then
            if mode ~= TransUnitIF.DeclClassMode.Interface then
               self:addErrMess( errPos, string.format( "define class already -- %s", name) )
               Util.printStackTrace(  )
            end
            
         end
      end
      
   else
    
      local parentInfo = self:getCurrentNamespaceTypeInfo(  )
      
      local parentScope = self.scope
      local scope = self:pushScope( true, baseInfo, interfaceList )
      local workGenTypeList
      
      if genTypeList ~= nil then
         workGenTypeList = genTypeList
      else
         workGenTypeList = {}
      end
      
      
      typeInfo = processInfo:createClass( mode ~= TransUnitIF.DeclClassMode.Interface, abstractFlag, scope, baseInfo, interfaceList, workGenTypeList, parentInfo, externalFlag, accessMode, name )
      
      parentScope:addClassLazy( processInfo, name, errPos, typeInfo, mode == TransUnitIF.DeclClassMode.LazyModule )
   end
   
   if genTypeList ~= nil then
      for __index, genType in ipairs( genTypeList ) do
         self.scope:addAlternate( processInfo, accessMode, genType:get_txt(), errPos, genType )
      end
      
   end
   
   
   local namespace = defNamespace
   if  nil == namespace then
      local _namespace = namespace
   
      namespace = Nodes.NamespaceInfo.new(name, self.scope, typeInfo)
   end
   
   self.typeId2ClassMap[typeInfo:get_typeId(  )] = namespace
   return typeInfo
end
function TransUnit:popClass(  )

   self:popScope(  )
end
function TransUnit:isTargetTokenPos( txt, pos )

   if self.analyzePos.lineNo == pos.lineNo and self.analyzePos.column >= pos.column and self.analyzePos.column <= pos.column + #txt then
      return true
   end
   
   return false
end
function TransUnit:isTargetToken( token )

   return self:isTargetTokenPos( token.txt, token.pos )
end
function TransUnit:dumpSymbolType( accessMode, name, typeInfo )

   local writer = Writer.JSON.new(io.stdout)
   writer:startParent( "lunescript", false )
   writer:startParent( "inquire", false )
   writer:write( "access", Ast.accessMode2txt( accessMode ) )
   writer:write( "name", name )
   writer:write( "type", typeInfo:getTxt( self.typeNameCtrl ) )
   writer:write( "typeKind", Ast.TypeInfoKind:_getTxt( typeInfo:get_kind())
    )
   writer:write( "static", string.format( "%s", typeInfo:get_staticFlag()) )
   writer:write( "display", typeInfo:get_display_stirng_with( typeInfo:get_rawTxt(), nil ) )
   writer:endElement(  )
   writer:endElement(  )
   writer:fin(  )
   os.exit( 0 )
end
function TransUnit:errorShadowingOp( pos, symbolInfo, errFlag )

   if symbolInfo ~= nil then
      local symPos = symbolInfo:get_pos()
      if symPos ~= nil then
         local mess = string.format( "This symbol is shadowed from %d:%d -- %s", pos.lineNo, pos.column, symbolInfo:get_name())
         self:addWarnErrMess( symPos, errFlag, mess )
      end
      
      local mess = string.format( "shadowing symbol of %s -- %s", symPos and string.format( "%s:%s", _lune.nilacc( symPos, "lineNo" ), _lune.nilacc( symPos, "column" )) or "external", symbolInfo:get_name())
      self:addWarnErrMess( pos, errFlag, mess )
   end
   
end
function TransUnit:errorShadowing( pos, symbolInfo )

   self:errorShadowingOp( pos, symbolInfo, not self.ctrl_info.warningShadowing )
end
function TransUnit:checkShadowing( pos, name, scope )

   if name == "_" then
      return 
   end
   
   local symbolInfo = self.scope:getSymbolTypeInfo( name, scope, self.moduleScope, self.scopeAccess )
   if  nil == symbolInfo then
      local _symbolInfo = symbolInfo
   
      return 
   end
   
   
   self:errorShadowing( pos, symbolInfo )
end
function TransUnit:addLocalVar( pos, argFlag, canBeLeft, name, typeInfo, mutable, allowShadow )

   if not allowShadow then
      self:checkShadowing( pos, name, self.scope )
      
   end
   
   
   if self.analyzeMode == AnalyzeMode.Inquire and self:isTargetTokenPos( name, pos ) then
      self:dumpSymbolType( Ast.AccessMode.Local, name, typeInfo )
   end
   
   
   return _lune.unwrap( self.scope:addLocalVar( self.processInfo, argFlag, canBeLeft, name, pos, typeInfo, mutable ))
end
function TransUnit.setmeta( obj )
  setmetatable( obj, { __index = TransUnit  } )
end
function TransUnit:get_errMessList()
   return self.errMessList
end
function TransUnit:get_warnMessList()
   return self.warnMessList
end
function TransUnit:get_importedAliasMap()
   return self.importedAliasMap
end


local op2levelMap = {}
local op1levelMap = {}
local opTopLevel = 0

do
   local opLevelBase = 0
   local function regOpLevel( opnum, opList )
   
      opLevelBase = opLevelBase + 1
      if opnum == 1 then
         for __index, op in pairs( opList ) do
            op1levelMap[op] = opLevelBase
         end
         
      else
       
         for __index, op in pairs( opList ) do
            op2levelMap[op] = opLevelBase
         end
         
      end
      
   end
   regOpLevel( 2, {"="} )
   regOpLevel( 2, {"or"} )
   regOpLevel( 2, {"and"} )
   regOpLevel( 2, {"<", ">", "<=", ">=", "~=", "=="} )
   regOpLevel( 2, {"|"} )
   regOpLevel( 2, {"~"} )
   regOpLevel( 2, {"&"} )
   regOpLevel( 2, {"|<<", "|>>"} )
   regOpLevel( 2, {".."} )
   regOpLevel( 2, {"+", "-"} )
   regOpLevel( 2, {"*", "/", "//", "%"} )
   regOpLevel( 1, {"`", ",,", ",,,", ",,,,"} )
   regOpLevel( 1, {"not", "#", "-", "~"} )
   regOpLevel( 1, {"^"} )
   
   opTopLevel = opLevelBase + 1
end


local quotedChar2Code = {}
quotedChar2Code['a'] = 7
quotedChar2Code['b'] = 8
quotedChar2Code['t'] = 9
quotedChar2Code['n'] = 10
quotedChar2Code['v'] = 11
quotedChar2Code['f'] = 12
quotedChar2Code['r'] = 13
quotedChar2Code['\\'] = 92
quotedChar2Code['"'] = 34
quotedChar2Code["'"] = 39

function TransUnit:createModifier( typeInfo, mutMode )

   if not self.validMutControl then
      return typeInfo
   end
   
   return self.processInfo:createModifier( typeInfo, mutMode )
end


function TransUnit:createExtType( pos, typeInfo )

   do
      local _matchExp = self.processInfo:createLuaval( typeInfo, true )
      if _matchExp[1] == Ast.LuavalResult.OK[1] then
         local work = _matchExp[2][1]
         local _ = _matchExp[2][2]
      
         return work
      elseif _matchExp[1] == Ast.LuavalResult.Err[1] then
         local err = _matchExp[2][1]
      
         self:addErrMess( pos, string.format( "not support -- %s", err) )
         return typeInfo
      end
   end
   
end


local BuiltinFuncType = Builtin.BuiltinFuncType
_moduleObj.BuiltinFuncType = BuiltinFuncType
local builtinFunc = Builtin.BuiltinFuncType.new()
local function getBuiltinFunc(  )

   return builtinFunc
end
_moduleObj.getBuiltinFunc = getBuiltinFunc

local function isStrFormFunc( typeInfo )

   if typeInfo:get_srcTypeInfo() == builtinFunc.string_format then
      return true
   end
   
   return false
end
_moduleObj.isStrFormFunc = isStrFormFunc

function TransUnit:checkThreading( pos )

   local curClass = self:getCurrentClass(  )
   if curClass ~= Ast.headTypeInfo then
      if curClass:isInheritFrom( self.processInfo, builtinFunc.lnsthread_, nil ) then
         return true
      end
      
   end
   
   if _lune._Set_has(self.helperInfo.pragmaSet, _lune.newAlge( LuneControl.Pragma.use_async) ) then
      self:addErrMess( pos, "This operation need perform on thread" )
   end
   
   return false
end


function TransUnit:errorAt( pos, mess )

   self:addErrMess( pos, mess )
   
   for __index, errmess in ipairs( self.errMessList ) do
      Util.errorLog( errmess )
   end
   
   for __index, warnmess in ipairs( self.warnMessList ) do
      Util.errorLog( warnmess )
   end
   
   if self.macroCtrl:get_analyzeInfo():get_mode() ~= Nodes.MacroMode.None then
      print( "------ near code -----", Nodes.MacroMode:_getTxt( self.macroCtrl:get_analyzeInfo():get_mode())
       )
      print( self.parser:getNearCode(  ) )
      print( "------" )
   end
   
   
   Util.err( "has error" )
end


function TransUnit:error( mess )

   self:errorAt( self.parser:getLastPos(  ), mess )
end


function TransUnit:createNoneNode( pos )

   return Nodes.NoneNode.create( self.nodeManager, pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone} )
end


function TransUnit:pushbackToken( token )

   self.parser:pushbackToken( token )
end


function TransUnit:newPushback( tokenList )

   self.parser:newPushback( tokenList )
end


function TransUnit:getStreamName(  )

   return self.parser:getStreamName(  )
end


function TransUnit:createPosition( lineNo, column )

   return self.parser:createPosition( lineNo, column )
end


function TransUnit:getTokenNoErr(  )

   local token
   
   
   local commentList = {}
   local workToken = self.parser:getTokenNoErr(  )
   while workToken.kind == Parser.TokenKind.Cmnt do
      table.insert( commentList, workToken )
      workToken = self.parser:getTokenNoErr(  )
   end
   
   if workToken.kind ~= Parser.TokenKind.Eof then
      token = workToken
      if self.macroCtrl:get_analyzeInfo():get_mode() ~= Nodes.MacroMode.None then
         token = self.macroCtrl:expandMacroVal( self.typeNameCtrl, self.scope, self, token )
      end
      
      if not self.ctrl_info.compatComment then
         self.commentCtrl:addDirect( commentList )
      end
      
   else
    
      token = Parser.getEofToken(  )
      self.commentCtrl:addDirect( commentList )
   end
   
   
   if #token:get_commentList() > 0 then
      self.commentCtrl:add( token )
   end
   
   
   return token
end


function TransUnit:getToken( allowEof )

   local token = self:getTokenNoErr(  )
   if token == Parser.getEofToken(  ) then
      if allowEof then
         return Parser.getEofToken(  )
      end
      
      self:error( "EOF" )
   end
   
   
   
   return token
end


function TransUnit:pushback(  )

   self.parser:pushback(  )
end


function TransUnit:pushbackStr( name, statement, pos )

   self.parser:pushbackStr( name, statement, pos )
end


local SymbolMode = {}
SymbolMode._val2NameMap = {}
function SymbolMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "SymbolMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function SymbolMode._from( val )
   if SymbolMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
SymbolMode.__allList = {}
function SymbolMode.get__allList()
   return SymbolMode.__allList
end

SymbolMode.Must_ = 0
SymbolMode._val2NameMap[0] = 'Must_'
SymbolMode.__allList[1] = SymbolMode.Must_
SymbolMode.MustNot_ = 1
SymbolMode._val2NameMap[1] = 'MustNot_'
SymbolMode.__allList[2] = SymbolMode.MustNot_
SymbolMode.MustNot_Or_ = 2
SymbolMode._val2NameMap[2] = 'MustNot_Or_'
SymbolMode.__allList[3] = SymbolMode.MustNot_Or_


local specialSymbolSet = {["__init"] = true, ["__free"] = true, ["__"] = true, ["_exp"] = true}
local builtinKeywordSet = {["self"] = true, ["super"] = true}
function TransUnit:checkSymbol( token, mode )

   if token.kind ~= Parser.TokenKind.Symb and token.kind ~= Parser.TokenKind.Kywd and token.kind ~= Parser.TokenKind.Type then
      self:addErrMess( token.pos, string.format( "illegal symbol -- '%s'", token.txt) )
   end
   
   local frontChar = string.byte( token.txt, 1 )
   if mode == SymbolMode.Must_ and frontChar ~= 95 then
      self:addErrMess( token.pos, string.format( "macro name must begin with '_' -- '%s'", token.txt) )
   elseif mode ~= SymbolMode.Must_ and frontChar == 95 then
      if not self.ignoreToCheckSymbol_ then
         if mode == SymbolMode.MustNot_Or_ and token.txt == "_" then
            
         elseif not _lune._Set_has(specialSymbolSet, token.txt ) then
            self:addErrMess( token.pos, string.format( "symbol must not begin with '_' -- '%s'", token.txt) )
         end
         
      end
      
   elseif _lune._Set_has(builtinKeywordSet, token.txt ) then
      self:addErrMess( token.pos, string.format( "this symbol is special keyword -- %s", token.txt) )
   elseif Parser.isLuaKeyword( token.txt ) or Parser.isOp2( token.txt ) or Parser.isOp1( token.txt ) then
      self:addErrMess( token.pos, string.format( "this symbol is lua keyword -- %s", token.txt) )
   end
   
   return token
end


function TransUnit:getSymbolToken( mode )

   return self:checkSymbol( self:getToken(  ), mode )
end


function TransUnit:checkToken( token, txt )

   if token.txt ~= txt then
      self:error( string.format( "not found -- %s. expects %s", txt, token.txt) )
   end
   
   return token
end


function TransUnit:checkNextToken( txt )

   return self:checkToken( self:getToken(  ), txt )
end

function TransUnit:getContinueToken(  )

   local token = self:getToken(  )
   return token, token.consecutive
end


function TransUnit:analyzeStatementList( stmtList, firstSwitchingParser, termTxt )

   local breakKind = Nodes.BreakKind.None
   if #stmtList > 0 then
      breakKind = stmtList[#stmtList]:getBreakKind( Nodes.CheckBreakMode.Normal )
   end
   
   
   local parser2lastLineMap = {}
   local function getLastLineNo(  )
   
      do
         local lastLineNo = parser2lastLineMap[self.parser]
         if lastLineNo ~= nil then
            return lastLineNo
         end
      end
      
      return self.parser:getLastPos(  ).lineNo
   end
   local function setLastLineNo( lineNo )
   
      parser2lastLineMap[self.parser] = lineNo
   end
   
   local lastStatement = nil
   local lastLineNo = getLastLineNo(  )
   local function setTailComment( statement )
   
      local blank
      
      local commentList = self.commentCtrl:get_commentList()
      if #commentList > 0 then
         if lastStatement ~= nil then
            local tailComment = nil
            for __index, comment in ipairs( commentList ) do
               if comment.pos.lineNo == lastStatement:get_pos().lineNo then
                  if not tailComment then
                     lastStatement:set_tailComment( comment )
                     tailComment = comment
                  else
                   
                     
                  end
                  
               end
               
            end
            
            if tailComment then
               table.remove( commentList, 1 )
            end
            
         end
         
      end
      
      if #commentList > 0 then
         blank = commentList[1].pos.lineNo - commentList[1]:getLineCount(  ) - lastLineNo
      else
       
         if statement ~= nil then
            blank = statement:get_pos().lineNo - lastLineNo
         else
            blank = self.parser:getLastPos(  ).lineNo - lastLineNo
         end
         
      end
      
      return blank
   end
   
   while true do
      lastLineNo = getLastLineNo(  )
      local statement = self:analyzeStatement( termTxt )
      if statement ~= nil then
         if breakKind ~= Nodes.BreakKind.None then
            if statement:get_kind() ~= Nodes.NodeKind.get_BlankLine() then
               self:addErrMess( statement:get_pos(), string.format( "This statement is not reached -- %s", Nodes.BreakKind:_getTxt( breakKind)
               ) )
            end
            
         end
         
         local blank = setTailComment( statement )
         if blank > 1 and not firstSwitchingParser then
            table.insert( stmtList, Nodes.BlankLineNode.create( self.nodeManager, self:createPosition( lastLineNo + 1, 0 ), self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, blank - 1 ) )
         end
         
         setLastLineNo( self.parser:getLastPos(  ).lineNo )
         if firstSwitchingParser then
            firstSwitchingParser = false
         end
         
         
         table.insert( stmtList, statement )
         lastStatement = statement
         if statement:get_kind() ~= Nodes.NodeKind.get_BlankLine() then
            breakKind = statement:getBreakKind( Nodes.CheckBreakMode.Normal )
         end
         
         
         statement:addComment( self.commentCtrl:get_commentList() )
         self.commentCtrl:clear(  )
      else
         setTailComment( nil )
         break
      end
      
   end
   
   return lastStatement, lastLineNo
end


function TransUnit:analyzeStatementListSubfile( stmtList )

   local statement = self:analyzeStatement(  )
   
   do
      local _exp = statement
      if _exp ~= nil then
         if _exp:get_kind() ~= Nodes.NodeKind.get_Subfile() then
            self:error( "subfile must have 'subfile' declaration at top." )
         end
         
      else
         self:error( "subfile must have 'subfile' declaration at top." )
      end
   end
   
   
   return (self:analyzeStatementList( stmtList, true ) )
end


function TransUnit:supportLang( lang )

   for pragma, __val in pairs( self.helperInfo.pragmaSet ) do
      do
         local _matchExp = pragma
         if _matchExp[1] == LuneControl.Pragma.limit_conv_code[1] then
            local codeSet = _matchExp[2][1]
         
            return _lune._Set_has(codeSet, lang )
         end
      end
      
   end
   
   return true
end


function TransUnit:analyzeLuneControl( firstToken )

   local node = nil
   local nextToken = self:getToken(  )
   
   local pragma
   
   do
      local _switchExp = (nextToken.txt )
      if _switchExp == "disable_mut_control" then
         self.validMutControl = false
         self.modifier:set_validMutControl( false )
         pragma = _lune.newAlge( LuneControl.Pragma.disable_mut_control)
      elseif _switchExp == "ignore_symbol_" then
         self.ignoreToCheckSymbol_ = true
         pragma = _lune.newAlge( LuneControl.Pragma.ignore_symbol_)
      elseif _switchExp == "load__lune_module" then
         pragma = _lune.newAlge( LuneControl.Pragma.load__lune_module)
      elseif _switchExp == "limit_conv_code" then
         local codeSet = {}
         while true do
            local token = self:getToken(  )
            if token.txt == ";" then
               self:pushback(  )
               break
            end
            
            do
               local code = LuneControl.Code._from( token.txt )
               if code ~= nil then
                  codeSet[code]= true
               else
                  self:addErrMess( token.pos, string.format( "illegal code -- '%s'", token.txt) )
               end
            end
            
         end
         
         pragma = _lune.newAlge( LuneControl.Pragma.limit_conv_code, {codeSet})
      elseif _switchExp == "use_async" then
         pragma = _lune.newAlge( LuneControl.Pragma.use_async)
      elseif _switchExp == "run_async_pipe" then
         if not _lune._Set_has(self.helperInfo.pragmaSet, _lune.newAlge( LuneControl.Pragma.use_async) ) then
            self:addErrMess( nextToken.pos, "must set '_lune_control use_async'" )
         end
         
         
         local nowMethod = self:getCurrentNamespaceTypeInfo(  )
         local nowClass = nowMethod:get_parentInfo()
         local valid = false
         if nowMethod:get_kind() == Ast.TypeInfoKind.Method and Ast.isClass( nowClass ) then
            do
               local loopMethod = _lune.nilacc( nowClass:get_scope(), 'getTypeInfoChild', 'callmtd' , "loop" )
               if loopMethod ~= nil then
                  if loopMethod:get_kind() == Ast.TypeInfoKind.Method and #loopMethod:get_argTypeInfoList() == 0 then
                     valid = true
                  end
                  
               end
            end
            
         end
         
         if valid then
            pragma = _lune.newAlge( LuneControl.Pragma.run_async_pipe)
         else
          
            self:addErrMess( nextToken.pos, "this option only use in method of the class have loop method." )
            return nil
         end
         
      elseif _switchExp == "run_async_runner" then
         local nowMethod = self:getCurrentNamespaceTypeInfo(  )
         local nowClass = self:getCurrentClass(  )
         if nowMethod:get_kind() == Ast.TypeInfoKind.Method and nowClass:isInheritFrom( self.processInfo, Ast.builtinTypeRunner, nil ) then
            pragma = _lune.newAlge( LuneControl.Pragma.run_async_runner)
         else
          
            self:addErrMess( nextToken.pos, "this option only use in the class inherit the __Runner." )
            return nil
         end
         
      elseif _switchExp == "default_async_func" then
         pragma = _lune.newAlge( LuneControl.Pragma.default_async_func)
         self.defaultAsyncMode = DefaultAsyncMode.AsyncFunc
      elseif _switchExp == "default_async_all" then
         pragma = _lune.newAlge( LuneControl.Pragma.default_async_all)
         self.defaultAsyncMode = DefaultAsyncMode.AsyncAll
      else 
         
            self:addErrMess( nextToken.pos, string.format( "unknown option -- %s", nextToken.txt) )
            self:checkNextToken( ";" )
            return nil
      end
   end
   
   
   node = Nodes.LuneControlNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, pragma )
   self.helperInfo.pragmaSet[pragma]= true
   
   self:checkNextToken( ";" )
   return node
end


local TentativeMode = {}
TentativeMode._val2NameMap = {}
function TentativeMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "TentativeMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function TentativeMode._from( val )
   if TentativeMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
TentativeMode.__allList = {}
function TentativeMode.get__allList()
   return TentativeMode.__allList
end

TentativeMode.Ignore = 0
TentativeMode._val2NameMap[0] = 'Ignore'
TentativeMode.__allList[1] = TentativeMode.Ignore
TentativeMode.Loop = 1
TentativeMode._val2NameMap[1] = 'Loop'
TentativeMode.__allList[2] = TentativeMode.Loop
TentativeMode.Simple = 2
TentativeMode._val2NameMap[2] = 'Simple'
TentativeMode.__allList[3] = TentativeMode.Simple
TentativeMode.Start = 3
TentativeMode._val2NameMap[3] = 'Start'
TentativeMode.__allList[4] = TentativeMode.Start
TentativeMode.Merge = 4
TentativeMode._val2NameMap[4] = 'Merge'
TentativeMode.__allList[5] = TentativeMode.Merge
TentativeMode.Finish = 5
TentativeMode._val2NameMap[5] = 'Finish'
TentativeMode.__allList[6] = TentativeMode.Finish


function TransUnit:analyzeBlock( blockKind, tentativeMode, scope, refAccessSymPosList )

   local token = self:checkNextToken( "{" )
   
   local backScope = self.scope
   if scope ~= nil then
      self.scope = scope
   else
      self:pushScope( false )
   end
   
   local blockScope = self.scope
   
   blockScope:addIgnoredVar( self.processInfo )
   
   do
      local _switchExp = tentativeMode
      if _switchExp == TentativeMode.Loop then
         self:prepareTentativeSymbol( self.scope, true, refAccessSymPosList )
      elseif _switchExp == TentativeMode.Simple or _switchExp == TentativeMode.Start or _switchExp == TentativeMode.Ignore then
         self:prepareTentativeSymbol( self.scope, false, nil )
      elseif _switchExp == TentativeMode.Merge or _switchExp == TentativeMode.Finish then
         self:mergeTentativeSymbol( self.scope )
      end
   end
   
   
   local loopFlag = false
   do
      local _switchExp = blockKind
      if _switchExp == Nodes.BlockKind.For or _switchExp == Nodes.BlockKind.Apply or _switchExp == Nodes.BlockKind.While or _switchExp == Nodes.BlockKind.Repeat or _switchExp == Nodes.BlockKind.Foreach then
         loopFlag = true
         table.insert( self.loopScopeQueue, self.scope )
      end
   end
   
   
   local stmtList = {}
   self:analyzeStatementList( stmtList, false, "}" )
   
   self:checkNextToken( "}" )
   
   if loopFlag then
      table.remove( self.loopScopeQueue )
   end
   
   
   if scope ~= nil then
      self.scope = backScope
   else
      self:popScope(  )
   end
   
   local node = Nodes.BlockNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, blockKind, blockScope, stmtList )
   
   if node:getBreakKind( Nodes.CheckBreakMode.Normal ) ~= Nodes.BreakKind.None then
      self.tentativeSymbol:skip(  )
   end
   
   
   if blockKind ~= Nodes.BlockKind.Repeat then
      do
         local _switchExp = tentativeMode
         if _switchExp == TentativeMode.Simple or _switchExp == TentativeMode.Finish then
            self:finishTentativeSymbol( true )
         elseif _switchExp == TentativeMode.Ignore or _switchExp == TentativeMode.Loop then
            self:finishTentativeSymbol( false )
         end
      end
      
   end
   
   
   return node
end


function TransUnit:analyzeImportFor( pos, modulePath, assignName, assigned, lazyLoad )

   local backupScope = self.scope
   self.scope = self.topScope
   
   self.processInfo:switchIdProvier( Ast.IdType.Ext )
   
   local importObj = self.importCtrl
   if  nil == importObj then
      local _importObj = importObj
   
      
      importObj = Import.Import.new(self, self.importModuleInfo, self.moduleType, builtinFunc, self.globalScope, self.macroCtrl, self.typeNameCtrl, self.importedAliasMap, self.validMutControl)
      self.importCtrl = importObj
   end
   
   local importProcessInfo = self.processInfo:newUser(  )
   importProcessInfo:switchIdProvier( Ast.IdType.Ext )
   local moduleInfo = importObj:processImport( importProcessInfo, modulePath )
   importProcessInfo:switchIdProvier( Ast.IdType.Base )
   
   self.processInfo:switchIdProvier( Ast.IdType.Base )
   
   self.scope = backupScope
   
   local provideInfo = moduleInfo:get_exportInfo():get_provideInfo()
   local moduleTypeInfo = provideInfo:get_typeInfo()
   self.scope:addModule( moduleTypeInfo, moduleInfo:assign( assignName ) )
   
   local moduleSymbolKind = provideInfo:get_symbolKind()
   local moduleSymbolInfo, shadowing = self.scope:add( self.processInfo, moduleSymbolKind, false, false, assignName, pos, moduleTypeInfo, Ast.AccessMode.Local, true, provideInfo:get_mutable() and Ast.MutMode.Mut or Ast.MutMode.IMut, true, lazyLoad ~= Nodes.LazyLoad.Off )
   
   if moduleSymbolInfo ~= nil then
      return Nodes.ImportNode.create( self.nodeManager, pos, self.macroCtrl:isInAnalyzeArgMode(  ), {moduleTypeInfo}, modulePath, lazyLoad, assignName, assigned, moduleSymbolInfo, moduleTypeInfo )
   end
   
   
   if shadowing ~= nil then
      self:errorShadowingOp( pos, shadowing, shadowing:get_typeInfo() ~= moduleTypeInfo )
   end
   
   return self:createNoneNode( pos )
end


function TransUnit:analyzeImport( opeToken )

   local lazyLoad
   
   if self:getToken(  ).txt == "." then
      local modeToken = self:getToken(  )
      do
         local _switchExp = modeToken.txt
         if _switchExp == "l" then
            lazyLoad = Nodes.LazyLoad.On
         elseif _switchExp == "d" then
            lazyLoad = Nodes.LazyLoad.Off
         else 
            
               lazyLoad = Nodes.LazyLoad.Off
               self:error( string.format( "illegal import mode -- %s", modeToken.txt) )
         end
      end
      
   else
    
      self:pushback(  )
      if self.ctrl_info.defaultLazy then
         lazyLoad = Nodes.LazyLoad.Auto
      else
       
         lazyLoad = Nodes.LazyLoad.Off
      end
      
   end
   
   if lazyLoad ~= Nodes.LazyLoad.Off then
      self.helperInfo.useLazyLoad = true
   end
   
   
   local moduleToken = self:getToken(  )
   local modulePath = moduleToken.txt
   local nextToken = moduleToken
   
   while true do
      nextToken = self:getToken(  )
      do
         local _switchExp = nextToken.txt
         if _switchExp == "." or _switchExp == "/" or _switchExp == ":" then
            local demilit = nextToken.txt
            nextToken = self:getToken(  )
            moduleToken = nextToken
            modulePath = string.format( "%s%s%s", modulePath, demilit, moduleToken.txt)
         else 
            
               break
         end
      end
      
   end
   
   local assignName = moduleToken
   local assigned
   
   if nextToken.txt == "as" then
      assignName = self:getSymbolToken( SymbolMode.MustNot_ )
      nextToken = self:getToken(  )
      assigned = true
   else
    
      assigned = false
   end
   
   self:checkToken( nextToken, ";" )
   
   local node = self:analyzeImportFor( opeToken.pos, modulePath, assignName.txt, assigned, lazyLoad )
   self.importModuleSet[node:get_expType()]= true
   
   return node
end


function TransUnit:analyzeTestCase( firstToken )

   local newScope = self:pushScope( false )
   
   local importNode = self:analyzeImportFor( firstToken.pos, "lune.base.Testing", "__t", false, Nodes.LazyLoad.Off )
   
   local nameToken = self:getSymbolToken( SymbolMode.MustNot_ )
   
   self:checkNextToken( "(" )
   
   local ctrlToken = self:getSymbolToken( SymbolMode.MustNot_ )
   local ctrlName = ctrlToken.txt
   self:checkNextToken( ")" )
   
   local moduleType = importNode:get_expType()
   local ctrlType = _lune.nilacc( moduleType:get_scope(), 'getTypeInfoChild', 'callmtd' , "Ctrl" )
   if  nil == ctrlType then
      local _ctrlType = ctrlType
   
      self:error( "not found Testing.Ctrl class" )
   end
   
   self:addLocalVar( ctrlToken.pos, true, false, ctrlToken.txt, ctrlType, Ast.MutMode.IMut, false )
   
   self.scopeAccess = Ast.ScopeAccess.Full
   
   local block = self:analyzeBlock( Nodes.BlockKind.Test, TentativeMode.Ignore, newScope, nil )
   
   self.scopeAccess = Ast.ScopeAccess.Normal
   
   self:popScope(  )
   
   return Nodes.TestCaseNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, nameToken, importNode, ctrlName, block )
end


function TransUnit:analyzeTest( firstToken )

   local nextToken = self:getToken(  )
   if nextToken.txt ~= "{" then
      self:pushback(  )
      return self:analyzeTestCase( firstToken )
   end
   
   self:checkToken( nextToken, "{" )
   
   local stmtList = {}
   self:analyzeStatementList( stmtList, false, "}" )
   self:checkNextToken( "}" )
   
   return Nodes.TestBlockNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, stmtList )
end


function TransUnit:analyzeSubfile( token )

   if self.scope ~= self.moduleScope then
      self:error( "'module' must be top scope." )
   end
   
   
   local mode = self:getToken(  )
   
   local moduleName = ""
   while true do
      local nextToken = self:getToken(  )
      if nextToken.txt == ";" then
         break
      end
      
      if moduleName == "" then
         moduleName = nextToken.txt
      else
       
         moduleName = string.format( "%s%s", moduleName, nextToken.txt)
      end
      
   end
   
   local usePath = nil
   if moduleName == "" then
      self:addErrMess( token.pos, "illegal subfile" )
   else
    
      if mode.txt == "use" then
         usePath = moduleName
         if frontInterface.searchModule( moduleName ) then
            table.insert( self.subfileList, moduleName )
         else
          
            self:addErrMess( token.pos, string.format( "not found subfile -- %s", moduleName) )
         end
         
      elseif mode.txt == "owner" then
         if frontInterface.getLuaModulePath( self.moduleName ) ~= moduleName then
            self:addErrMess( token.pos, string.format( "illegal owner module -- %s, %s", moduleName, self.moduleName) )
         end
         
      else
       
         self:addErrMess( mode.pos, string.format( "illegal module mode -- %s", mode.txt) )
      end
      
   end
   
   return Nodes.SubfileNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, usePath )
end


function TransUnit:analyzeEnvLock( token )

   return Nodes.EnvNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, self:analyzeBlock( Nodes.BlockKind.Env, TentativeMode.Simple ) )
end


function TransUnit:analyzeIf( token )

   local nextToken, continueFlag = self:getContinueToken(  )
   if continueFlag and nextToken.txt == "!" then
      return self:analyzeIfUnwrap( token )
   end
   
   self:pushback(  )
   
   local list = {}
   local ifExp = self:analyzeExpOneRVal( false, false )
   table.insert( list, Nodes.IfStmtInfo.new(Nodes.IfKind.If, ifExp, self:analyzeBlock( Nodes.BlockKind.If, TentativeMode.Start )) )
   
   local function checkCond( condExp )
   
      do
         local _switchExp = condExp:get_expType():get_kind()
         if _switchExp == Ast.TypeInfoKind.Nilable or _switchExp == Ast.TypeInfoKind.Stem then
            
         elseif _switchExp == Ast.TypeInfoKind.Prim then
            if not condExp:get_expType():equals( self.processInfo, Ast.builtinTypeBool ) then
               self:addErrMess( condExp:get_pos(), string.format( "This exp never be false -- %s", condExp:get_expType():getTxt(  )) )
            end
            
         else 
            
               self:addErrMess( condExp:get_pos(), string.format( "This exp never be false -- %s", condExp:get_expType():getTxt(  )) )
         end
      end
      
   end
   checkCond( ifExp )
   
   nextToken = self:getToken( true )
   if nextToken.txt == "elseif" then
      while nextToken.txt == "elseif" do
         local condExp = self:analyzeExpOneRVal( false, false )
         table.insert( list, Nodes.IfStmtInfo.new(Nodes.IfKind.ElseIf, condExp, self:analyzeBlock( Nodes.BlockKind.Elseif, TentativeMode.Merge )) )
         checkCond( condExp )
         nextToken = self:getToken( true )
      end
      
   end
   
   
   if nextToken.txt == "else" then
      table.insert( list, Nodes.IfStmtInfo.new(Nodes.IfKind.Else, self:createNoneNode( nextToken.pos ), self:analyzeBlock( Nodes.BlockKind.Else, TentativeMode.Finish )) )
   else
    
      self:finishTentativeSymbol( false )
      self:pushback(  )
   end
   
   
   return Nodes.IfNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, list )
end

function TransUnit:processCaseDefault( firstToken, caseKind, nextToken, hasCase )

   local keyword = firstToken.txt:gsub( "_", "" )
   local fullKeyword = string.format( "_%s", keyword)
   if firstToken.txt == fullKeyword and caseKind ~= Nodes.CaseKind.MustFull then
      self:addErrMess( firstToken.pos, string.format( "This '%s' hasn't enough 'case' condition.", keyword) )
   end
   
   
   local defaultBlock = nil
   local failSafeDefault = false
   if nextToken.txt == "default" or nextToken.txt == "_default" then
      if firstToken.txt == fullKeyword then
         self:addErrMess( nextToken.pos, string.format( "'_%s' can't have default.", keyword) )
      end
      
      
      if nextToken.txt == "_default" then
         failSafeDefault = true
      elseif caseKind == Nodes.CaseKind.Full then
         self:addWarnMess( nextToken.pos, string.format( "This '%s' has full case. This 'default' is no reach.", keyword) )
      end
      
      defaultBlock = self:analyzeBlock( Nodes.BlockKind.Default, not hasCase and TentativeMode.Simple or TentativeMode.Finish )
   else
    
      if hasCase then
         self:finishTentativeSymbol( caseKind ~= Nodes.CaseKind.Lack )
      end
      
      self:pushback(  )
   end
   
   self:checkNextToken( "}" )
   
   if not hasCase then
      self:addWarnMess( firstToken.pos, string.format( "'%s' should have 'case' blocks.", keyword) )
      if defaultBlock then
         self:addErrMess( firstToken.pos, string.format( "'%s' must have 'case' blocks when have 'default' block.", keyword) )
      end
      
   end
   
   
   return defaultBlock, failSafeDefault
end


function TransUnit:analyzeSwitch( firstToken )

   local exp = self:analyzeExpOneRVal( false, false )
   
   self:checkNextToken( "{" )
   
   local caseList = {}
   
   local condObjSet = {}
   local condSymIdSet = {}
   local hasNilCond = false
   
   local nextToken = self:getToken(  )
   local firstFlag = true
   while (nextToken.txt == "case" ) do
      self:checkToken( nextToken, "case" )
      local condexpList = self:analyzeExpList( false, false, false, nil, {exp:get_expType()}, true )
      local condBock = self:analyzeBlock( Nodes.BlockKind.Switch, firstFlag and TentativeMode.Start or TentativeMode.Merge )
      if firstFlag then
         firstFlag = false
      end
      
      
      for __index, condExp in ipairs( condexpList:get_expList() ) do
         if condExp:get_expType():get_nilable() then
            if hasNilCond then
               self:addWarnMess( condExp:get_pos(), "multiple case with nil or nilable" )
            else
             
               hasNilCond = true
            end
            
         end
         
         do
            local condLiteral = condExp:getLiteral(  )
            if condLiteral ~= nil then
               local literalObj = Nodes.getLiteralObj( condLiteral )
               if literalObj ~= nil then
                  if _lune._Set_has(condObjSet, literalObj ) then
                     self:addErrMess( condExp:get_pos(), string.format( "multiple case exp -- %s", literalObj) )
                  else
                   
                     condObjSet[literalObj]= true
                  end
                  
               end
               
            else
               do
                  local expRef = _lune.__Cast( condExp, 3, Nodes.ExpRefNode )
                  if expRef ~= nil then
                     local symInfo = expRef:get_symbolInfo()
                     if _lune._Set_has(condSymIdSet, symInfo:get_symbolId() ) then
                        self:addErrMess( condExp:get_pos(), string.format( "multiple case exp -- %s", symInfo:get_name()) )
                     else
                      
                        condSymIdSet[symInfo:get_symbolId()]= true
                     end
                     
                  end
               end
               
            end
         end
         
         if not exp:get_expType():canEvalWith( self.processInfo, condExp:get_expType(), Ast.CanEvalType.Equal, {} ) then
            self:addErrMess( condExp:get_pos(), string.format( "case exp unmatch type -- %s <- %s", exp:get_expType():getTxt(  ), condExp:get_expType():getTxt(  )) )
         end
         
      end
      
      
      table.insert( caseList, Nodes.CaseInfo.new(condexpList, condBock) )
      nextToken = self:getToken(  )
   end
   
   
   local caseKind
   
   do
      local enumType = _lune.__Cast( exp:get_expType():get_srcTypeInfo():get_aliasSrc(), 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         local miss = false
         for __index, enumVal in pairs( enumType:get_name2EnumValInfo() ) do
            if not _lune._Set_has(condObjSet, Ast.getEnumLiteralVal( enumVal:get_val() ) ) then
               miss = true
               break
            end
            
         end
         
         if not miss then
            if firstToken.txt == "_switch" then
               caseKind = Nodes.CaseKind.MustFull
            else
             
               caseKind = Nodes.CaseKind.Full
            end
            
         else
          
            caseKind = Nodes.CaseKind.Lack
         end
         
      else
         caseKind = Nodes.CaseKind.Lack
         if firstToken.txt == "_switch" then
            self:addErrMess( exp:get_pos(), "The condition of '_switch' must be enum." )
         end
         
      end
   end
   
   
   local defaultBlock, failSafeDefault = self:processCaseDefault( firstToken, caseKind, nextToken, #caseList ~= 0 )
   
   return Nodes.SwitchNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, exp, caseList, defaultBlock, caseKind, failSafeDefault )
end


function TransUnit:analyzeMatch( firstToken )

   local exp = self:analyzeExpOneRVal( false, false )
   
   local algeTypeInfo = _lune.__Cast( exp:get_expType():get_srcTypeInfo(), 3, Ast.AlgeTypeInfo )
   if  nil == algeTypeInfo then
      local _algeTypeInfo = algeTypeInfo
   
      self:error( string.format( "match must have alge value -- %s", exp:get_expType():getTxt(  )) )
   end
   
   
   self:checkNextToken( "{" )
   
   local caseList = {}
   local algeValNameSet = {}
   
   local nextToken = self:getToken(  )
   local firstFlag = true
   while (nextToken.txt == "case" ) do
      self:checkNextToken( "." )
      local valNameToken = self:getToken(  )
      
      self:checkAlgeComp( valNameToken, algeTypeInfo )
      
      local valInfo = algeTypeInfo:getValInfo( valNameToken.txt )
      if  nil == valInfo then
         local _valInfo = valInfo
      
         self:error( string.format( "not found val -- %s", valNameToken.txt) )
      end
      
      if _lune._Set_has(algeValNameSet, valNameToken.txt ) then
         self:addErrMess( valNameToken.pos, string.format( "multiple val -- %s", valNameToken.txt) )
      end
      
      algeValNameSet[valNameToken.txt]= true
      
      local valParamNameList = {}
      nextToken = self:getToken(  )
      local blockScope = self:pushScope( false )
      if nextToken.txt == "(" then
         for __index, paramType in ipairs( valInfo:get_typeList() ) do
            local paramName = self:getSymbolToken( SymbolMode.MustNot_Or_ )
            self:checkShadowing( paramName.pos, paramName.txt, self.scope )
            
            local workType = paramType
            if Ast.TypeInfo.isMut( paramType ) and not Ast.TypeInfo.isMut( exp:get_expType() ) then
               workType = self:createModifier( workType, Ast.MutMode.IMut )
            end
            
            table.insert( valParamNameList, self:addLocalVar( paramName.pos, false, false, paramName.txt, workType, Ast.MutMode.IMut ) )
            nextToken = self:getToken(  )
            if nextToken.txt ~= "," then
               break
            end
            
         end
         
         self:checkToken( nextToken, ")" )
      else
       
         self:pushback(  )
      end
      
      if #valParamNameList ~= #valInfo:get_typeList() then
         self:addErrMess( valNameToken.pos, string.format( "unmatch param -- %d != %d", #valParamNameList, #valInfo:get_typeList()) )
      end
      
      local mode = firstFlag and TentativeMode.Start or TentativeMode.Merge
      
      local block = self:analyzeBlock( Nodes.BlockKind.Match, mode, blockScope, nil )
      if firstFlag then
         firstFlag = false
      end
      
      self:popScope(  )
      local valRefNode = Nodes.ExpRefNode.create( self.nodeManager, valNameToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {valInfo:get_algeTpye()}, valInfo:get_symbolInfo() )
      local matchCase = Nodes.MatchCase.new(valInfo, valRefNode, valParamNameList, block)
      
      table.insert( caseList, matchCase )
      nextToken = self:getToken(  )
   end
   
   
   local caseKind
   
   if _lune._Set_len(algeValNameSet ) == algeTypeInfo:get_valInfoNum() then
      if firstToken.txt == "_match" then
         caseKind = Nodes.CaseKind.MustFull
      else
       
         caseKind = Nodes.CaseKind.Full
      end
      
   else
    
      caseKind = Nodes.CaseKind.Lack
   end
   
   
   local defaultBlock, failSafeDefault = self:processCaseDefault( firstToken, caseKind, nextToken, #caseList ~= 0 )
   
   return Nodes.MatchNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, exp, algeTypeInfo, caseList, defaultBlock, caseKind, failSafeDefault )
end


function TransUnit:analyzeWhile( token )

   
   local refAccessSymPosList = clearThePosForModToRef( self.scope, self.moduleScope )
   
   local cond = self:analyzeExpOneRVal( false, false )
   local infinit = false
   if cond:get_expType() == Ast.builtinTypeBool then
      do
         local literal = cond:getLiteral(  )
         if literal ~= nil then
            do
               local _matchExp = literal
               if _matchExp[1] == Nodes.Literal.Bool[1] then
                  local val = _matchExp[2][1]
               
                  infinit = val
               end
            end
            
         end
      end
      
   elseif not cond:get_expType():get_nilable() then
      infinit = true
   end
   
   
   return Nodes.WhileNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, cond, infinit, self:analyzeBlock( Nodes.BlockKind.While, TentativeMode.Loop, nil, refAccessSymPosList ) )
end


function TransUnit:analyzeRepeat( token )

   local scope = self:pushScope( false )
   local node = Nodes.RepeatNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, self:analyzeBlock( Nodes.BlockKind.Repeat, TentativeMode.Loop, scope, nil ), self:analyzeExpOneRVal( false, false ) )
   self:finishTentativeSymbol( false )
   
   self:popScope(  )
   
   self:checkNextToken( ";" )
   return node
end


function TransUnit:analyzeFor( firstToken )

   
   local scope = self:pushScope( false )
   
   local val = self:getToken(  )
   if val.kind ~= Parser.TokenKind.Symb then
      self:error( "not symbol" )
   end
   
   self:checkNextToken( "=" )
   local exp1 = self:analyzeExpOneRVal( false, false )
   if not Ast.isNumberType( exp1:get_expType() ) then
      self:addErrMess( exp1:get_pos(), string.format( "exp1 is not number -- %s", exp1:get_expType():getTxt(  )) )
   end
   
   self:checkNextToken( "," )
   local exp2 = self:analyzeExpOneRVal( false, false )
   if not Ast.isNumberType( exp2:get_expType() ) then
      self:addErrMess( exp2:get_pos(), string.format( "exp2 is not number -- %s", exp2:get_expType():getTxt(  )) )
   end
   
   local token = self:getToken(  )
   local exp3 = nil
   if token.txt == "," then
      exp3 = self:analyzeExpOneRVal( false, false )
      do
         local _exp = exp3
         if _exp ~= nil then
            if not Ast.isNumberType( _exp:get_expType() ) then
               self:addErrMess( _exp:get_pos(), string.format( "exp is not number -- %s", _exp:get_expType():getTxt(  )) )
            end
            
         end
      end
      
   else
    
      self:pushback(  )
   end
   
   
   if exp1:get_expType() == Ast.builtinTypeInt and exp2:get_expType() == Ast.builtinTypeReal or _lune.nilacc( exp3, 'get_expType', 'callmtd' ) == Ast.builtinTypeReal then
      exp1 = Nodes.ExpCastNode.create( self.nodeManager, exp1:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeReal}, exp1, Ast.builtinTypeReal, Nodes.CastKind.Force )
   end
   
   
   local symbolInfo = self:addLocalVar( val.pos, false, true, val.txt, exp1:get_expType(), Ast.MutMode.IMut )
   
   local node = Nodes.ForNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, self:analyzeBlock( Nodes.BlockKind.For, TentativeMode.Loop, scope, nil ), symbolInfo, exp1, exp2, exp3 )
   self:popScope(  )
   
   return node
end


function TransUnit:analyzeApply( token )

   local scope = self:pushScope( false )
   local varList = {}
   local nextToken = Parser.getEofToken(  )
   repeat 
      local var = self:getSymbolToken( SymbolMode.MustNot_Or_ )
      if var.kind ~= Parser.TokenKind.Symb then
         self:error( "illegal symbol" )
      end
      
      table.insert( varList, var )
      nextToken = self:getToken(  )
   until nextToken.txt ~= ","
   self:checkToken( nextToken, "of" )
   
   local expListNode = self:analyzeExpList( false, false, false )
   
   local itFunc = Ast.builtinTypeNone
   
   local expTypeList = expListNode:get_expTypeList()
   if #expTypeList < 3 then
      self:addErrMess( expListNode:get_pos(), string.format( "apply must have 3 values -- %s", #expTypeList) )
   else
    
      itFunc = expTypeList[1]:get_extedType()
   end
   
   
   local itemTypeList = {}
   local defaultItemType = Ast.builtinTypeStem_
   
   local readyFlag = false
   do
      local callNode = _lune.__Cast( Nodes.getUnwraped( expListNode:get_expList()[1] ), 3, Nodes.ExpCallNode )
      if callNode ~= nil then
         local callFuncType = callNode:get_func():get_expType()
         if callFuncType:equals( self.processInfo, builtinFunc.str_gmatch ) or callFuncType:equals( self.processInfo, builtinFunc.string_gmatch ) then
            table.insert( itemTypeList, Ast.builtinTypeString )
            defaultItemType = Ast.builtinTypeString:get_nilableTypeInfo()
            readyFlag = true
         end
         
      end
   end
   
   
   if not readyFlag then
      do
         local _switchExp = itFunc:get_kind()
         if _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.FormFunc or _switchExp == Ast.TypeInfoKind.Form then
         else 
            
               self:addErrMess( expListNode:get_pos(), string.format( "The 1st value must be iterator function. -- %s", itFunc:getTxt(  )) )
         end
      end
      
      
      if #itFunc:get_argTypeInfoList() ~= 2 then
         self:addErrMess( expListNode:get_pos(), string.format( "iterator function must has two arguments. -- %s", itFunc:getTxt(  )) )
      else
       
         local arg2Type = itFunc:get_argTypeInfoList()[2]
         if not arg2Type:get_nilable() then
            self:addErrMess( expListNode:get_pos(), string.format( "the 2nd argument of iterator function must be nilable. -- %s", itFunc:getTxt(  )) )
         end
         
      end
      
      
      if #itFunc:get_retTypeInfoList() == 0 then
         self:addErrMess( expListNode:get_pos(), "iterator function must return value." )
      else
       
         local iteRetType = itFunc:get_retTypeInfoList()[1]
         if not iteRetType:get_nilable() then
            self:addErrMess( expListNode:get_pos(), "iterator function must return nilable type at 1st." )
         end
         
      end
      
      
      for index, itemType in ipairs( itFunc:get_retTypeInfoList() ) do
         local workType = itemType
         if index == 1 then
            if itemType:get_nilable() then
               workType = workType:get_nonnilableType()
            end
            
         end
         
         table.insert( itemTypeList, workType )
      end
      
   end
   
   
   local varSymList = {}
   for index, var in ipairs( varList ) do
      local itemType = defaultItemType
      if index <= #itemTypeList then
         itemType = itemTypeList[index]
      end
      
      table.insert( varSymList, self:addLocalVar( var.pos, false, true, var.txt, itemType, Ast.MutMode.IMut ) )
   end
   
   
   local block = self:analyzeBlock( Nodes.BlockKind.Apply, TentativeMode.Loop, scope, nil )
   self:popScope(  )
   
   return Nodes.ApplyNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, varSymList, expListNode, block )
end


function TransUnit:convToExtTypeList( pos, typeInfo, list )

   if typeInfo:get_nonnilableType():get_kind() ~= Ast.TypeInfoKind.Ext then
      return list
   end
   
   local newList, mess = Ast.convToExtTypeList( self.processInfo, list )
   if newList ~= nil then
      return newList
   end
   
   self:addErrMess( pos, mess )
   return list
end


function TransUnit:analyzeForeach( token, sortFlag )

   local scope = self:pushScope( false )
   local mainSymToken = Parser.getEofToken(  )
   local subSymToken = nil
   local mainSym
   
   local subSym = nil
   local nextToken = Parser.getEofToken(  )
   for index = 1, 2 do
      local symbol = self:getToken(  )
      if symbol.kind ~= Parser.TokenKind.Symb then
         self:error( "illegal symbol" )
      end
      
      if index == 1 then
         mainSymToken = symbol
      else
       
         subSymToken = symbol
      end
      
      nextToken = self:getToken(  )
      if nextToken.txt ~= "," then
         break
      end
      
   end
   
   self:checkToken( nextToken, "in" )
   
   local exp = self:analyzeExpOneRVal( false, false )
   local function checkSortType( sortKeyType )
   
      if sortFlag then
         
         do
            local _switchExp = sortKeyType:get_srcTypeInfo():get_extedType()
            if _switchExp == Ast.builtinTypeString or _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeReal or _switchExp == Ast.builtinTypeStem then
            else 
               
                  self:addErrMess( exp:get_pos(), string.format( "This type can't use forsort -- %s", sortKeyType:getTxt(  )) )
            end
         end
         
      end
      
   end
   
   local expType = exp:get_expType():get_extedType()
   local itemTypeInfoList = self:convToExtTypeList( token.pos, exp:get_expType(), expType:get_itemTypeInfoList() )
   do
      local _switchExp = expType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Map then
         mainSym = self:addLocalVar( mainSymToken.pos, false, true, mainSymToken.txt, itemTypeInfoList[2], Ast.MutMode.IMut )
         do
            local _exp = subSymToken
            if _exp ~= nil then
               subSym = self:addLocalVar( _exp.pos, false, true, _exp.txt, itemTypeInfoList[1], Ast.MutMode.IMut )
            end
         end
         
         checkSortType( itemTypeInfoList[1] )
      elseif _switchExp == Ast.TypeInfoKind.Set then
         if subSymToken ~= nil then
            self:addErrMess( subSymToken.pos, "Set can't use index" )
         end
         
         mainSym = self:addLocalVar( mainSymToken.pos, false, true, mainSymToken.txt, itemTypeInfoList[1], Ast.MutMode.IMut )
         checkSortType( itemTypeInfoList[1] )
      elseif _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         if sortFlag then
            self:addErrMess( exp:get_pos(), string.format( "'%s' doesn't support forsort.", Ast.TypeInfoKind:_getTxt( expType:get_kind())
            ) )
         end
         
         mainSym = self:addLocalVar( mainSymToken.pos, false, true, mainSymToken.txt, itemTypeInfoList[1], Ast.MutMode.IMut )
         do
            local _exp = subSymToken
            if _exp ~= nil then
               subSym = self:addLocalVar( _exp.pos, false, false, _exp.txt, Ast.builtinTypeInt, Ast.MutMode.IMut )
            end
         end
         
      else 
         
            self:errorAt( exp:get_pos(), string.format( "unknown kind type of exp for foreach -- %s", expType:getTxt(  )) )
      end
   end
   
   
   local seqSym = nil
   do
      local refNode = _lune.__Cast( exp, 3, Nodes.ExpRefNode )
      if refNode ~= nil then
         local seqSymbol = refNode:get_symbolInfo()
         if seqSymbol:get_mutable() or Ast.TypeInfo.isMut( seqSymbol:get_typeInfo() ) then
            scope:addOverrideImut( self.processInfo, seqSymbol )
            seqSym = seqSymbol:get_name()
         end
         
      end
   end
   
   
   local block = self:analyzeBlock( Nodes.BlockKind.Foreach, TentativeMode.Loop, scope, nil )
   
   if seqSym ~= nil then
      scope:remove( seqSym )
   end
   
   
   self:popScope(  )
   
   local threading
   
   if exp:get_expType():get_kind() == Ast.TypeInfoKind.Ext then
      threading = self:checkThreading( token.pos )
   else
    
      threading = false
   end
   
   
   if sortFlag then
      return Nodes.ForsortNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, mainSym, subSym, exp, threading, block, sortFlag )
   else
    
      return Nodes.ForeachNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, mainSym, subSym, exp, threading, block )
   end
   
end


function TransUnit:analyzeProvide( firstToken )

   local token = self:getSymbolToken( SymbolMode.MustNot_ )
   local symbolNode = self:analyzeExpSymbol( firstToken, token, ExpSymbolMode.Symbol, nil, true, false )
   self:checkNextToken( ";" )
   
   local symbolInfoList = symbolNode:getSymbolInfo(  )
   if #symbolInfoList ~= 1 then
      self:error( "'provide' must be symbol." )
   end
   
   local symbolInfo = symbolInfoList[1]
   
   local node = Nodes.ProvideNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, symbolInfo )
   if self.provideNode then
      self:addErrMess( firstToken.pos, "multiple provide" )
   end
   
   self.provideNode = node
   
   if symbolInfo:get_accessMode() ~= Ast.AccessMode.Pub then
      self:addErrMess( firstToken.pos, string.format( "provide variable must be 'pub'.  -- %s", symbolInfo:get_accessMode()) )
   end
   
   
   return node
end


function TransUnit:analyzeScope( firstToken )

   
   local nextToken = self:getToken(  )
   local scopeKind
   
   do
      local _switchExp = nextToken.txt
      if _switchExp == "root" then
         scopeKind = Nodes.ScopeKind.Root
      else 
         
            self:error( string.format( "illegal scope kind. -- %s", nextToken.txt) )
      end
   end
   
   
   local symList = {}
   nextToken = self:getToken(  )
   if nextToken.txt == "(" then
      nextToken = self:getToken(  )
      while nextToken.txt ~= ")" do
         local symbolNode = self:analyzeExpSymbol( nextToken, nextToken, ExpSymbolMode.Symbol, nil, true, false )
         local workSymList = symbolNode:getSymbolInfo(  )
         if #workSymList > 0 then
            table.insert( symList, workSymList[1] )
         end
         
         nextToken = self:getToken(  )
         do
            local _switchExp = nextToken.txt
            if _switchExp == ")" then
            elseif _switchExp == "," then
               nextToken = self:getToken(  )
            else 
               
                  self:error( string.format( "illegal token: expects ')' or ',' but -- %s", nextToken.txt) )
            end
         end
         
      end
      
   else
    
      self:pushback(  )
   end
   
   local bakScope = self.scope
   self.scope = self.topScope
   self:pushScope( false )
   for __index, symInfo in ipairs( symList ) do
      self.scope:addAlias( self.processInfo, symInfo:get_name(), nextToken.pos, false, symInfo:get_accessMode(), symInfo:get_typeInfo():get_parentInfo(), symInfo )
   end
   
   
   local block = self:analyzeBlock( Nodes.BlockKind.Block, TentativeMode.Simple )
   local node = Nodes.ScopeNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, scopeKind, self.scope, symList, block )
   self.scope = bakScope
   
   return node
end


function TransUnit:analyzeRefType( accessMode, allowDDD, parentPub )

   local firstToken = self:getToken(  )
   local token = firstToken
   local mutMode
   
   do
      local _switchExp = token.txt
      if _switchExp == "&" then
         mutMode = Ast.MutMode.IMut
         token = self:getToken(  )
      elseif _switchExp == "allmut" then
         mutMode = Ast.MutMode.AllMut
         token = self:getToken(  )
      else 
         
            mutMode = nil
      end
   end
   
   local name
   
   if token.txt == "..." then
      local dddSym = _lune.unwrap( self.moduleScope:getSymbolInfo( "...", self.moduleScope, true, Ast.ScopeAccess.Normal ))
      name = Nodes.ExpRefNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {dddSym:get_typeInfo()}, Ast.AccessSymbolInfo.new(self.processInfo, dddSym, _lune.newAlge( Ast.OverrideMut.None), true) )
   else
    
      name = self:analyzeExpSymbol( firstToken, token, ExpSymbolMode.Symbol, nil, true, false )
      local symbolList = name:getSymbolInfo(  )
      if #symbolList > 0 then
         local symbol = symbolList[1]
         if symbol:get_kind() ~= Ast.SymbolKind.Typ then
            self:addErrMess( name:get_pos(), string.format( "illegal type -- %s", symbol:get_name()) )
         end
         
      else
       
         self:addErrMess( name:get_pos(), string.format( "illegal symbol node -- %s", Nodes.getNodeKindName( name:get_kind() )) )
      end
      
   end
   
   return self:analyzeRefTypeWithSymbol( accessMode, allowDDD, mutMode, name, parentPub )
end


function TransUnit:analyzeRefTypeWithSymbol( accessMode, allowDDD, mutMode, symbolNode, parentPub )

   local typeInfo = symbolNode:get_expType()
   if typeInfo:get_kind() == Ast.TypeInfoKind.Set and not self.helperInfo.useSet then
      self.helperInfo.useSet = true
   end
   
   
   do
      local aliasType = _lune.__Cast( typeInfo, 3, Ast.AliasTypeInfo )
      if aliasType ~= nil then
         local aliasSrc = aliasType:get_aliasSrcTypeInfo()
         if not _lune._Set_has(self.importModuleSet, aliasSrc:getModule(  ) ) and self.moduleType:get_parentInfo() ~= aliasSrc:getModule(  ):get_parentInfo() then
            self:addErrMess( symbolNode:get_pos(), string.format( "must import '%s' for this alias -- %s", aliasSrc:getModule(  ):getFullName( self.typeNameCtrl, self.scope, false ), symbolNode:getSymbolInfo(  )[1]:get_name()) )
         end
         
      end
   end
   
   
   if parentPub and Ast.isPubToExternal( accessMode ) and not Ast.isPubToExternal( typeInfo:get_accessMode() ) then
      self:addErrMess( symbolNode:get_pos(), string.format( "This type must be public. -- %s", typeInfo:getTxt(  )) )
   end
   
   
   local continueToken, continueFlag = self:getContinueToken(  )
   local token = continueToken
   if continueFlag and token.txt == "!" then
      typeInfo = typeInfo:get_nilableTypeInfo(  )
      token = self:getToken(  )
   end
   
   
   local itemNodeList = {}
   local arrayMode = "no"
   while true do
      if token.txt == '[' or token.txt == '[@' then
         if token.txt == '[' then
            arrayMode = "list"
            typeInfo = self.processInfo:createList( accessMode, self:getCurrentClass(  ), {typeInfo}, Ast.MutMode.Mut )
         else
          
            arrayMode = "array"
            typeInfo = self.processInfo:createArray( accessMode, self:getCurrentClass(  ), {typeInfo}, Ast.MutMode.Mut )
         end
         
         token = self:getToken(  )
         if token.txt ~= ']' then
            self:pushback(  )
            self:checkNextToken( ']' )
         end
         
         table.insert( itemNodeList, symbolNode )
      elseif token.txt == "<" then
         local genericList = {}
         local nextToken = Parser.getEofToken(  )
         repeat 
            local typeExp = self:analyzeRefType( accessMode, false, parentPub )
            table.insert( itemNodeList, typeExp )
            table.insert( genericList, typeExp:get_expType() )
            nextToken = self:getToken(  )
         until nextToken.txt ~= ","
         self:checkToken( nextToken, '>' )
         
         local function checkAlternateTypeCount( count )
         
            if #genericList ~= count then
               self:addErrMess( symbolNode:get_pos(), string.format( "generic type count is unmatch. -- %d", #genericList) )
               return false
            end
            
            return true
         end
         
         do
            local _switchExp = typeInfo:get_kind()
            if _switchExp == Ast.TypeInfoKind.Map then
               if #genericList ~= 2 then
                  self:addErrMess( symbolNode:get_pos(), "Key or value type is unknown" )
                  typeInfo = self.processInfo:createMap( accessMode, self:getCurrentClass(  ), Ast.builtinTypeStem, Ast.builtinTypeStem, Ast.MutMode.Mut )
               else
                
                  typeInfo = self.processInfo:createMap( accessMode, self:getCurrentClass(  ), genericList[1], genericList[2], Ast.MutMode.Mut )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.List then
               if checkAlternateTypeCount( 1 ) then
                  typeInfo = self.processInfo:createList( accessMode, self:getCurrentClass(  ), genericList, Ast.MutMode.Mut )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Array then
               if checkAlternateTypeCount( 1 ) then
                  typeInfo = self.processInfo:createArray( accessMode, self:getCurrentClass(  ), genericList, Ast.MutMode.Mut )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Set then
               if checkAlternateTypeCount( 1 ) then
                  typeInfo = self.processInfo:createSet( accessMode, self:getCurrentClass(  ), genericList, Ast.MutMode.Mut )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.DDD then
               if checkAlternateTypeCount( 1 ) then
                  typeInfo = self.processInfo:createDDD( genericList[1], false, false )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
               if checkAlternateTypeCount( #typeInfo:get_itemTypeInfoList() ) then
                  for __index, itemType in ipairs( genericList ) do
                     if itemType:get_nilable() then
                        self:addErrMess( symbolNode:get_pos(), string.format( "can't use nilable type -- %s", itemType:getTxt(  )) )
                     end
                     
                  end
                  
                  for index, itemType in ipairs( typeInfo:get_itemTypeInfoList() ) do
                     if itemType:hasBase(  ) and not genericList[index]:isInheritFrom( self.processInfo, itemType:get_baseTypeInfo() ) then
                        self:addErrMess( symbolNode:get_pos(), string.format( "'%s' of %s doesn't inherit '%s'", itemType:getTxt(  ), typeInfo:getTxt(  ), itemType:get_baseTypeInfo():getTxt(  )) )
                     end
                     
                  end
                  
                  typeInfo = self.processInfo:createGeneric( typeInfo, genericList, self.moduleType )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Box then
               if checkAlternateTypeCount( 1 ) then
                  typeInfo = self.processInfo:createBox( accessMode, genericList[1] )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Ext then
               if checkAlternateTypeCount( 1 ) then
                  typeInfo = self:createExtType( symbolNode:get_pos(), genericList[1] )
               end
               
            else 
               
                  self:error( string.format( "not support generic: %s", typeInfo:getTxt(  ) ) )
            end
         end
         
      else
       
         self:pushback(  )
         break
      end
      
      token = self:getToken(  )
   end
   
   if token.txt == "!" then
      typeInfo = typeInfo:get_nilableTypeInfo(  )
      self:getToken(  )
   end
   
   
   if not allowDDD then
      if typeInfo:get_kind() == Ast.TypeInfoKind.DDD then
         self:addErrMess( symbolNode:get_pos(), string.format( "invalid type. -- '%s'", typeInfo:getTxt(  )) )
      end
      
   end
   
   
   if mutMode ~= nil then
      if typeInfo:get_mutMode() ~= mutMode then
         typeInfo = self:createModifier( typeInfo, mutMode )
      end
      
   end
   
   
   if typeInfo:get_kind() == Ast.TypeInfoKind.Module then
      self:addErrMess( symbolNode:get_pos(), string.format( "module can't use as Type. -- %s", typeInfo:getTxt(  )) )
   end
   
   
   return Nodes.RefTypeNode.create( self.nodeManager, symbolNode:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, symbolNode, itemNodeList, mutMode, arrayMode )
end


function TransUnit:analyzeDeclArgList( accessMode, scope, argList, parentPub )

   local nextToken = Parser.noneToken
   local hasDDDFlag = false
   repeat 
      nextToken = self:getToken(  )
      if nextToken.txt == ")" then
         break
      end
      
      
      if hasDDDFlag then
         self:addErrMess( nextToken.pos, "Argument exists after '...'." )
      end
      
      
      local mutable = Ast.MutMode.IMut
      if nextToken.txt == "mut" then
         mutable = Ast.MutMode.Mut
         nextToken = self:getToken(  )
      end
      
      local argName = nextToken
      if argName.txt == "..." then
         hasDDDFlag = true
         
         local workToken, flag = self:getContinueToken(  )
         self:pushback(  )
         
         local dddTypeInfo = Ast.builtinTypeDDD
         if flag and workToken.txt == "<" then
            self:pushbackToken( nextToken )
            local refTypeNode = self:analyzeRefType( accessMode, true, parentPub )
            dddTypeInfo = refTypeNode:get_expType()
         end
         
         
         table.insert( argList, Nodes.DeclArgDDDNode.create( self.nodeManager, argName.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {dddTypeInfo} ) )
         scope:addLocalVar( self.processInfo, true, true, argName.txt, argName.pos, dddTypeInfo, Ast.MutMode.IMut )
      else
       
         argName = self:checkSymbol( argName, SymbolMode.MustNot_ )
         
         self:checkShadowing( argName.pos, argName.txt, scope )
         self:checkNextToken( ":" )
         local refType = self:analyzeRefType( accessMode, false, parentPub )
         if refType:get_expType():get_kind() == Ast.TypeInfoKind.Class and #refType:get_expType():get_itemTypeInfoList() > 0 then
            local argType = refType:get_expType():get_srcTypeInfo()
            if not Ast.isGenericType( argType ) then
               self:addErrMess( refType:get_pos(), string.format( "can't use this type without <T>. please use %s.", argType:getTxt(  )) )
            end
            
         end
         
         
         do
            local symbolInfo = scope:addLocalVar( self.processInfo, true, true, argName.txt, argName.pos, refType:get_expType(), mutable )
            if symbolInfo ~= nil then
               local arg = Nodes.DeclArgNode.create( self.nodeManager, argName.pos, self.macroCtrl:isInAnalyzeArgMode(  ), refType:get_expTypeList(), argName, symbolInfo, refType )
               
               table.insert( argList, arg )
            end
         end
         
      end
      
      nextToken = self:getToken(  )
   until nextToken.txt ~= ","
   
   self:checkToken( nextToken, ")" )
   
   return nextToken
end


function TransUnit:checkOverrideMethod( overrideType, typeInfo )

   local accessMode = typeInfo:get_accessMode()
   local funcName = typeInfo:get_rawTxt()
   local altTypeList = typeInfo:get_itemTypeInfoList()
   local alt2typeMap = typeInfo:get_parentInfo():createAlt2typeMap( false )
   local errList = {}
   
   local function addErr( mess )
   
      local fullName = string.format( "%s.%s", typeInfo:get_parentInfo():get_rawTxt(), typeInfo:get_rawTxt())
      table.insert( errList, string.format( "%s: %s: %s -- %s", fullName, mess, typeInfo:get_display_stirng(), typeInfo:get_display_stirng()) )
   end
   
   if self.ctrl_info.validAsyncCtrl then
      if overrideType:get_asyncMode() ~= typeInfo:get_asyncMode() then
         addErr( string.format( "mismatch asyncMode --  %s, %s", Ast.Async:_getTxt( overrideType:get_asyncMode())
         , Ast.Async:_getTxt( typeInfo:get_asyncMode())
         ) )
      end
      
   end
   
   
   if overrideType:get_accessMode(  ) ~= accessMode then
      local mess = string.format( "mismatch override accessMode -- %s,%s", Ast.AccessMode:_getTxt( overrideType:get_accessMode(  ))
      , Ast.AccessMode:_getTxt( accessMode)
      )
      addErr( mess )
   end
   
   if overrideType:get_staticFlag(  ) ~= typeInfo:get_staticFlag() then
      addErr( "mismatch override staticFlag -- " .. funcName )
   end
   
   if overrideType:get_kind(  ) ~= Ast.TypeInfoKind.Method then
      addErr( string.format( "mismatch override kind -- %s, %d", funcName, overrideType:get_kind(  )) )
   end
   
   if overrideType:get_mutMode() ~= typeInfo:get_mutMode() then
      addErr( string.format( "mismatch mutable -- %s", funcName) )
   end
   
   
   if #overrideType:get_itemTypeInfoList() ~= #altTypeList then
      local mess = string.format( "mismatch altTypeList -- %d, %d", #overrideType:get_itemTypeInfoList(), #altTypeList)
      addErr( mess )
   else
    
      for index, alterType in ipairs( overrideType:get_itemTypeInfoList() ) do
         alt2typeMap[alterType] = altTypeList[index]
      end
      
   end
   
   
   local matchFlag, err = overrideType:canEvalWith( self.processInfo, typeInfo, Ast.CanEvalType.SetEq, alt2typeMap )
   if not matchFlag then
      if err ~= nil then
         addErr( string.format( "mismatch method type -- %s", err) )
      else
         addErr( "mismatch method type" )
      end
      
   end
   
   
   for index, retType in ipairs( overrideType:get_retTypeInfoList() ) do
      if #typeInfo:get_retTypeInfoList() >= index then
         if retType:get_nonnilableType():get_kind() == Ast.TypeInfoKind.Alternate and typeInfo:get_retTypeInfoList()[index]:get_nonnilableType():get_kind() ~= Ast.TypeInfoKind.Alternate then
            local mess = string.format( "not support to override the method has generics at return type. -- %s", funcName)
            addErr( mess )
         end
         
      end
      
   end
   
   
   return errList
end


function TransUnit:checkOverriededMethodOfAllClass(  )

   local function process( pos, alt2typeMap, classScope, superScope )
   
      superScope:filterTypeInfoField( true, classScope, self.scopeAccess, function ( symbolInfo )
      
         if symbolInfo:get_name() == "__init" then
            return true
         end
         
         local implimented = true
         if symbolInfo:get_typeInfo():get_kind() == Ast.TypeInfoKind.Method then
            do
               local impMethodType = classScope:getTypeInfoField( symbolInfo:get_name(), true, classScope, self.scopeAccess )
               if impMethodType ~= nil then
                  if impMethodType == symbolInfo:get_typeInfo() then
                     if symbolInfo:get_typeInfo():get_abstractFlag() and classScope ~= superScope then
                        implimented = false
                     end
                     
                  else
                   
                     for __index, err in ipairs( self:checkOverrideMethod( symbolInfo:get_typeInfo(), impMethodType ) ) do
                        self:addErrMess( pos, err )
                     end
                     
                  end
                  
               else
                  implimented = false
               end
            end
            
         end
         
         if not implimented then
            self:addErrMess( pos, string.format( "not implements method -- %s.%s at %s", _lune.nilacc( superScope:get_ownerTypeInfo(), 'getTxt', 'callmtd'  ), symbolInfo:get_name(), _lune.nilacc( classScope:get_ownerTypeInfo(), 'getTxt', 'callmtd'  )) )
         end
         
         return true
      end )
   end
   
   local typeId2DeclClassNode = {}
   for classTypeInfo, classNode in pairs( self.typeInfo2ClassNode ) do
      typeId2DeclClassNode[classTypeInfo:get_typeId().id] = classNode
   end
   
   
   do
      local __sorted = {}
      local __map = typeId2DeclClassNode
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local classNode = __map[ __key ]
         do
            local classTypeInfo = classNode:get_expType()
            local workTypeInfo = classTypeInfo
            local alt2typeMap = classTypeInfo:createAlt2typeMap( false )
            repeat 
               if not classTypeInfo:get_abstractFlag() then
                  if workTypeInfo ~= Ast.headTypeInfo then
                     process( classNode:get_pos(), alt2typeMap, _lune.unwrap( classTypeInfo:get_scope()), _lune.unwrap( workTypeInfo:get_scope()) )
                  end
                  
               end
               
               for __index, ifType in ipairs( workTypeInfo:get_interfaceList() ) do
                  if ifType ~= Ast.builtinTypeMapping then
                     process( classNode:get_pos(), alt2typeMap, _lune.unwrap( classTypeInfo:get_scope()), _lune.unwrap( ifType:get_scope()) )
                  end
                  
               end
               
               workTypeInfo = workTypeInfo:get_baseTypeInfo()
            until workTypeInfo == Ast.headTypeInfo
         end
      end
   end
   
end


local ASTInfo = {}
_moduleObj.ASTInfo = ASTInfo
function ASTInfo.setmeta( obj )
  setmetatable( obj, { __index = ASTInfo  } )
end
function ASTInfo.new( node, exportInfo, streamName )
   local obj = {}
   ASTInfo.setmeta( obj )
   if obj.__init then
      obj:__init( node, exportInfo, streamName )
   end
   return obj
end
function ASTInfo:__init( node, exportInfo, streamName )

   self.node = node
   self.exportInfo = exportInfo
   self.streamName = streamName
end
function ASTInfo:get_node()
   return self.node
end
function ASTInfo:get_exportInfo()
   return self.exportInfo
end
function ASTInfo:get_streamName()
   return self.streamName
end


function TransUnit:createAST( parser, macroFlag, moduleName )
   local __func__ = '@lune.@base.@TransUnit.TransUnit.createAST'

   Log.log( Log.Level.Log, __func__, 536, function (  )
      local __func__ = '@lune.@base.@TransUnit.TransUnit.createAST.<anonymous>'
   
      return string.format( "%s start -- %s", __func__, parser:getStreamName(  ))
   end )
   
   
   self.moduleName = _lune.unwrapDefault( moduleName, "")
   
   do
      self.scope = Ast.rootScope
      local builtin = Builtin.Builtin.new(self, self.targetLuaVer, self.ctrl_info)
      builtinFunc = builtin:registBuiltInScope(  )
      self.scope = self.topScope
   end
   
   
   local moduleTypeInfo = Ast.headTypeInfo
   local moduleSymboInfo = nil
   
   if moduleName ~= nil then
      for txt in string.gmatch( frontInterface.getLuaModulePath( moduleName ), '[^%.]+' ) do
         moduleTypeInfo = self:pushModule( self.processInfo, false, txt, true )
      end
      
   end
   
   self.moduleScope = self.scope
   self.moduleType = moduleTypeInfo
   self.moduleScope:addVar( self.processInfo, Ast.AccessMode.Global, "__mod__", nil, Ast.builtinTypeString, Ast.MutMode.IMut, true )
   
   self.typeNameCtrl = Ast.TypeNameCtrl.new(moduleTypeInfo)
   
   self.parser = Parser.DefaultPushbackParser.new(parser)
   
   self.scope:addIgnoredVar( self.processInfo )
   
   local ast
   
   local globalSymbolList = {}
   
   local lastStatement = nil
   if macroFlag then
      ast = self:analyzeBlock( Nodes.BlockKind.Macro, TentativeMode.Ignore )
   else
    
      local children = {}
      local lastLineNo
      
      lastStatement, lastLineNo = self:analyzeStatementList( children, false )
      
      local statement = Nodes.BlankLineNode.create( self.nodeManager, self:createPosition( lastLineNo + 1, 0 ), self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, 0 )
      statement:addComment( self.commentCtrl:get_commentList() )
      self.commentCtrl:clear(  )
      table.insert( children, statement )
      
      local token = self:getTokenNoErr(  )
      if token ~= Parser.getEofToken(  ) then
         self:error( string.format( "%s:%d:%d:(%s) not eof -- %s", self.parser:getStreamName(  ), token.pos.lineNo, token.pos.column, Types.TokenKind:_getTxt( token.kind)
         , token.txt) )
      end
      
      
      for __index, subModule in ipairs( self.subfileList ) do
         local file = frontInterface.searchModule( subModule )
         if  nil == file then
            local _file = file
         
            self:error( string.format( "not found subfile -- %s", subModule) )
         end
         
         
         if self.scope ~= self.moduleScope then
            self:error( "scope does not close" )
         end
         
         
         local subParser = Parser.StreamParser.create( file, false, subModule )
         if  nil == subParser then
            local _subParser = subParser
         
            self:error( string.format( "open error -- %s", file) )
         end
         
         
         self.parser = Parser.DefaultPushbackParser.new(subParser)
         
         lastStatement = self:analyzeStatementListSubfile( children )
         
         token = self:getTokenNoErr(  )
         if token ~= Parser.getEofToken(  ) then
            Util.err( string.format( "unknown:%d:%d:(%s) %s", token.pos.lineNo, token.pos.column, Types.TokenKind:_getTxt( token.kind)
            , token.txt) )
         end
         
      end
      
      
      self:checkOverriededMethodOfAllClass(  )
      
      local rootNode = Nodes.RootNode.create( self.nodeManager, self:createPosition( 0, 0 ), self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, children, self.moduleScope, self.macroCtrl:get_useModuleMacroSet(), self.moduleId, self.processInfo, moduleTypeInfo, nil, self.helperInfo, self.nodeManager, _lune.unwrapDefault( _lune.nilacc( self.importCtrl, 'get_importModule2ModuleInfo', 'callmtd' ), {}), self.macroCtrl:get_typeId2MacroInfo(), self.typeId2ClassMap )
      ast = rootNode
      do
         local _exp = self.provideNode
         if _exp ~= nil then
            if lastStatement ~= _exp then
               self:addErrMess( _exp:get_pos(), "'provide' must be last." )
            end
            
            rootNode:set_provide( _exp )
            moduleSymboInfo = _exp:get_symbol()
         end
      end
      
      
      ClosureFun.checkList( self.closureFunList )
      
      for __index, node in ipairs( children ) do
         do
            local workNode = _lune.__Cast( node, 3, Nodes.DeclVarNode )
            if workNode ~= nil then
               for __index, symbolInfo in ipairs( workNode:get_symbolInfoList() ) do
                  if symbolInfo:get_accessMode() == Ast.AccessMode.Global then
                     table.insert( globalSymbolList, symbolInfo )
                  end
                  
               end
               
            end
         end
         
         do
            local workNode = _lune.__Cast( node, 3, Nodes.DeclFuncNode )
            if workNode ~= nil then
               if workNode:get_declInfo():get_accessMode() == Ast.AccessMode.Global then
                  do
                     local symbolInfo = workNode:get_declInfo():get_symbol()
                     if symbolInfo ~= nil then
                        table.insert( globalSymbolList, symbolInfo )
                     end
                  end
                  
               end
               
            end
         end
         
      end
      
   end
   
   
   if moduleName ~= nil then
      for _1797 in string.gmatch( moduleName, '[^%.]+' ) do
         self:popModule(  )
      end
      
   end
   
   
   local function createId2proto( map )
   
      local id2proto = {}
      for protoType, _1804 in pairs( map ) do
         id2proto[protoType:get_typeId().id] = protoType
      end
      
      return id2proto
   end
   
   do
      local __sorted = {}
      local __map = createId2proto( self.protoFuncMap )
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local protoType = __map[ __key ]
         do
            self:addErrMess( _lune.unwrap( self.protoFuncMap[protoType]), string.format( "This function doesn't have body. -- %s", protoType:getTxt(  )) )
         end
      end
   end
   
   do
      local __sorted = {}
      local __map = createId2proto( self.protoClassMap )
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local protoType = __map[ __key ]
         do
            self:addErrMess( _lune.unwrap( self.protoClassMap[protoType]), string.format( "This class doesn't have body. -- %s", protoType:getTxt(  )) )
         end
      end
   end
   
   
   for __index, mess in ipairs( self.warnMessList ) do
      Util.errorLog( mess )
   end
   
   if #self.errMessList > 0 then
      for __index, mess in ipairs( self.errMessList ) do
         Util.errorLog( mess )
      end
      
      Util.err( "has error" )
   end
   
   if self.ctrl_info.stopByWarning and #self.warnMessList > 0 then
      Util.err( "has error" )
   end
   
   
   do
      local _switchExp = self.analyzeMode
      if _switchExp == AnalyzeMode.Diag or _switchExp == AnalyzeMode.Complete or _switchExp == AnalyzeMode.Inquire then
         os.exit( 0 )
      end
   end
   
   
   local provideInfo
   
   if moduleSymboInfo ~= nil then
      provideInfo = frontInterface.ModuleProvideInfo.new(moduleSymboInfo:get_typeInfo(), moduleSymboInfo:get_kind(), moduleSymboInfo:get_mutable())
   else
      provideInfo = frontInterface.ModuleProvideInfo.new(moduleTypeInfo, Ast.SymbolKind.Typ, false)
   end
   
   
   local exportInfo = Nodes.ExportInfo.new(moduleTypeInfo, provideInfo, self.processInfo, globalSymbolList, self.macroCtrl:get_declMacroInfoMap())
   
   return ASTInfo.new(ast, exportInfo, parser:getStreamName(  ))
end


function TransUnit:analyzeDeclMacroSub( accessMode, firstToken, nameToken, macroScope, parentType, workArgList )

   if self.macroCtrl:get_isDeclaringMacro() then
      self:error( "can't declare macro in the macro." )
   end
   
   
   self.macroCtrl:startDecl(  )
   
   local pubFlag = false
   do
      local _switchExp = accessMode
      if _switchExp == Ast.AccessMode.Pub then
         pubFlag = true
      elseif _switchExp == Ast.AccessMode.Local or _switchExp == Ast.AccessMode.None then
      else 
         
            self:addErrMess( firstToken.pos, string.format( "macro not support this access mode. -- %s", Ast.AccessMode:_getTxt( accessMode)
            ) )
      end
   end
   
   local argList = {}
   local argTypeList = {}
   for __index, argNode in ipairs( workArgList ) do
      do
         local _exp = _lune.__Cast( argNode, 3, Nodes.DeclArgNode )
         if _exp ~= nil then
            table.insert( argList, _exp )
         else
            self:error( "macro argument can not use '...'." )
         end
      end
      
      local argType = argNode:get_expType()
      table.insert( argTypeList, argType )
   end
   
   
   local nextToken = self:getToken(  )
   
   local retTypeList
   
   if nextToken.txt == ":" then
      retTypeList = self:analyzeRefType( accessMode, true, false ):get_expTypeList()
      self:checkNextToken( "{" )
   else
    
      retTypeList = {}
      self:checkToken( nextToken, "{" )
   end
   
   
   nextToken = self:getToken(  )
   
   local stmtNode = nil
   if nextToken.txt == "{" then
      self.macroScope = macroScope
      macroScope:set_validCheckingUnaccess( false )
      
      local funcType = self.processInfo:createFunc( false, true, nil, Ast.TypeInfoKind.Func, Ast.headTypeInfo, false, true, true, Ast.AccessMode.Global, "_lnsLoad", Ast.Async.Async, nil, {Ast.builtinTypeString, Ast.builtinTypeString}, {Ast.builtinTypeStem}, false )
      macroScope:addLocalVar( self.processInfo, false, false, "_lnsLoad", nil, funcType, Ast.MutMode.IMut )
      
      local macroLocalVarType = self.processInfo:createMap( Ast.AccessMode.Local, self.moduleType, Ast.builtinTypeString, Ast.builtinTypeStem, Ast.MutMode.Mut )
      macroScope:addLocalVar( self.processInfo, false, true, "__var", nil, macroLocalVarType, Ast.MutMode.IMut )
      
      local stmtList = {}
      self:prepareTentativeSymbol( self.scope, false, nil )
      self:analyzeStatementList( stmtList, false, "}" )
      
      stmtNode = Nodes.BlockNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, Nodes.BlockKind.Macro, macroScope, stmtList )
      
      self:checkNextToken( "}" )
      self:finishTentativeSymbol( true )
      
      self.macroScope = nil
   else
    
      self:pushback(  )
   end
   
   local tokenList = {}
   local braceCount = 0
   
   while true do
      nextToken = self:getToken(  )
      if nextToken.txt == "{" then
         braceCount = braceCount + 1
      elseif nextToken.txt == "}" then
         if braceCount == 0 then
            break
         end
         
         braceCount = braceCount - 1
      end
      
      table.insert( tokenList, nextToken )
   end
   
   
   local typeInfo = self.processInfo:createFunc( false, false, macroScope, Ast.TypeInfoKind.Macro, parentType, false, false, true, accessMode, nameToken.txt, Ast.Async.Async, nil, argTypeList, retTypeList )
   local declMacroInfo = Nodes.DeclMacroInfo.new(pubFlag, nameToken, argList, stmtNode, tokenList)
   local node = Nodes.DeclMacroNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, declMacroInfo )
   
   self.macroCtrl:regist( self.processInfo, node, macroScope )
   
   return node
end


function TransUnit:analyzeDeclMacro( accessMode, firstToken )

   local _
   
   local nameToken = self:getSymbolToken( SymbolMode.Must_ )
   self:checkNextToken( "(" )
   
   local scope = Ast.TypeInfo.createScope( self.processInfo, self.topScope, false, nil, nil )
   local workArgList = {}
   self:analyzeDeclArgList( accessMode, scope, workArgList, false )
   local parentInfo = self:getCurrentNamespaceTypeInfo(  )
   local backScope = self.scope
   self.scope = scope
   
   self.scope:addIgnoredVar( self.processInfo )
   
   local node = self:analyzeDeclMacroSub( accessMode, firstToken, nameToken, scope, parentInfo, workArgList )
   self.scope = backScope
   
   local _1877, existSym = self.scope:addMacro( self.processInfo, nameToken.pos, node:get_expType(), accessMode )
   if existSym then
      self:addErrMess( nameToken.pos, string.format( "multiple define symbol -- %s", nameToken.txt) )
   end
   
   
   return node
end


function TransUnit:analyzeExtend( accessMode, firstPos )

   local baseRef = nil
   local interfaceList = {}
   local ifAlt2typeMap = {}
   local ifRefList = {}
   
   local nextToken = self:getToken(  )
   if nextToken.txt ~= "(" then
      self:pushback(  )
      local workBaseRefType = self:analyzeRefType( accessMode, false, Ast.isPubToExternal( accessMode ) )
      baseRef = workBaseRefType
      local baseType = workBaseRefType:get_expType()
      if baseType:get_kind() ~= Ast.TypeInfoKind.Class then
         self:addErrMess( workBaseRefType:get_pos(), string.format( "%s is not class.", baseType:getTxt(  )) )
      end
      
      if Ast.isPubToExternal( accessMode ) and not Ast.isPubToExternal( baseType:get_accessMode() ) then
         self:addErrMess( workBaseRefType:get_pos(), string.format( "%s can't be external symbol.", baseType:getTxt(  )) )
      end
      
      nextToken = self:getToken(  )
   end
   
   if nextToken.txt == "(" then
      while true do
         nextToken = self:getToken(  )
         if nextToken.txt == ")" then
            break
         end
         
         self:pushback(  )
         local ifTypeNode = self:analyzeRefType( accessMode, false, Ast.isPubToExternal( accessMode ) )
         table.insert( ifRefList, ifTypeNode )
         local ifType = ifTypeNode:get_expType()
         if ifType:get_kind() ~= Ast.TypeInfoKind.IF then
            self:error( string.format( "%s is not interface -- %d", ifType:getTxt(  ), ifType:get_kind()) )
         end
         
         
         if Ast.isGenericType( ifType ) then
            for altType, genType in pairs( ifType:createAlt2typeMap( false ) ) do
               ifAlt2typeMap[altType] = genType
            end
            
         end
         
         
         table.insert( interfaceList, ifType )
         if Ast.isPubToExternal( accessMode ) and not Ast.isPubToExternal( ifType:get_accessMode() ) then
            self:addErrMess( ifTypeNode:get_pos(), string.format( "%s can't be external symbol.", ifType:getTxt(  )) )
         end
         
         
         nextToken = self:getToken(  )
         if nextToken.txt ~= "," then
            if nextToken.txt == ")" then
               break
            end
            
            self:error( "illegal token" )
         end
         
      end
      
      nextToken = self:getToken(  )
   end
   
   local symbol2TypeInfo = {}
   for __index, ifType in ipairs( interfaceList ) do
      _lune.nilacc( ifType:get_scope(), 'filterTypeInfoField', 'callmtd' , true, self.scope, self.scopeAccess, function ( symbolInfo )
      
         if symbolInfo:get_kind() == Ast.SymbolKind.Mtd then
            do
               local ifFuncType = symbol2TypeInfo[symbolInfo:get_name()]
               if ifFuncType ~= nil then
                  local ret, mess = ifFuncType:canEvalWith( self.processInfo, symbolInfo:get_typeInfo(), Ast.CanEvalType.SetOp, ifAlt2typeMap )
                  if not ret then
                     self:addErrMess( firstPos, string.format( "mismatch method type -- %s.%s, %s.%s\n%s", symbolInfo:get_typeInfo():get_parentInfo():getTxt(  ), symbolInfo:get_name(), ifFuncType:get_parentInfo():getTxt(  ), ifFuncType:getTxt(  ), mess) )
                  end
                  
               else
                  symbol2TypeInfo[symbolInfo:get_name()] = symbolInfo:get_typeInfo()
               end
            end
            
         end
         
         return true
      end )
   end
   
   
   local baseTypeInfo = nil
   if baseRef ~= nil then
      baseTypeInfo = baseRef:get_expType()
   end
   
   return nextToken, baseTypeInfo, interfaceList, ifAlt2typeMap, Nodes.ClassInheritInfo.new(baseRef, ifRefList)
end


function TransUnit:analyzePushClass( mode, abstractFlag, firstToken, name, allowMultiple, requirePath, moduleLang, accessMode, altTypeList )

   if Ast.isPubToExternal( accessMode ) and self.moduleScope ~= self.scope then
      self:addErrMess( firstToken.pos, "The public class must declare at top scope." )
   end
   
   local tempScope = self:pushScope( false )
   for __index, altType in ipairs( altTypeList ) do
      tempScope:addAlternate( self.processInfo, accessMode, altType:get_rawTxt(), name.pos, altType )
   end
   
   
   local nextToken = self:getToken(  )
   local baseTypeInfo = nil
   local interfaceList = nil
   local inheritInfo
   
   if nextToken.txt == "extend" then
      local _
      nextToken, baseTypeInfo, interfaceList, _, inheritInfo = self:analyzeExtend( accessMode, firstToken.pos )
      
      if baseTypeInfo ~= nil then
         do
            local initTypeInfo = _lune.nilacc( baseTypeInfo:get_scope(), 'getTypeInfoChild', 'callmtd' , "__init" )
            if initTypeInfo ~= nil then
               if initTypeInfo:get_accessMode() == Ast.AccessMode.Pri then
                  self:addErrMess( firstToken.pos, "The access mode of '__init' is 'pri'." )
               end
               
            end
         end
         
         
         if baseTypeInfo:isInheritFrom( self.processInfo, builtinFunc.lnsthread_, nil ) and not _lune._Set_has(self.helperInfo.pragmaSet, _lune.newAlge( LuneControl.Pragma.use_async) ) then
            self:addErrMess( nextToken.pos, "must set '_lune_control use_async'" )
         end
         
      end
      
   else
    
      inheritInfo = Nodes.ClassInheritInfo.new(nil, {})
   end
   
   
   self:popScope(  )
   
   local classTypeInfo
   
   do
      local _switchExp = mode
      if _switchExp == TransUnitIF.DeclClassMode.Module or _switchExp == TransUnitIF.DeclClassMode.LazyModule then
         local parentScope = self.scope
         classTypeInfo = self:pushExtModule( false, name.txt, accessMode, name.pos, mode == TransUnitIF.DeclClassMode.LazyModule, _lune.unwrap( moduleLang), (_lune.unwrap( requirePath) ):getExcludedDelimitTxt(  ) )
      elseif _switchExp == TransUnitIF.DeclClassMode.Class or _switchExp == TransUnitIF.DeclClassMode.Interface then
         classTypeInfo = self:pushClass( self.processInfo, firstToken.pos, mode, abstractFlag, baseTypeInfo, interfaceList, altTypeList, false, name.txt, allowMultiple, accessMode )
      end
   end
   
   
   return nextToken, classTypeInfo, inheritInfo
end


function TransUnit:analyzeDeclAlternateType( belongClassFlag, token, accessMode )

   local altTypeList = {}
   local nextToken = token
   local altNameSet = {}
   local altIndex = 0
   while true do
      altIndex = altIndex + 1
      local genericSymToken = self:getSymbolToken( SymbolMode.MustNot_ )
      if self.scope:getTypeInfo( genericSymToken.txt, self.scope, false, self.scopeAccess ) then
         self:addErrMess( genericSymToken.pos, string.format( "shadowing Type -- %s", genericSymToken.txt) )
      else
       
         if _lune._Set_has(altNameSet, genericSymToken.txt ) then
            self:addErrMess( genericSymToken.pos, string.format( "multiple Type -- %s", genericSymToken.txt) )
         else
          
            altNameSet[genericSymToken.txt]= true
         end
         
      end
      
      local workToken = self:getToken(  )
      if workToken.txt == "!" then
         self:addErrMess( workToken.pos, "not support nilable" )
         workToken = self:getToken(  )
      end
      
      local baseTypeInfo = nil
      local interfaceList = {}
      if workToken.txt == ":" then
         workToken, baseTypeInfo, interfaceList = self:analyzeExtend( accessMode, token.pos )
      end
      
      
      local altType = self.processInfo:createAlternate( belongClassFlag, altIndex, genericSymToken.txt, accessMode, self.moduleType, baseTypeInfo, interfaceList )
      table.insert( altTypeList, altType )
      
      if workToken.txt == ">" then
         nextToken = self:getToken(  )
         break
      end
      
      self:checkToken( workToken, "," )
   end
   
   return nextToken, altTypeList
end


function TransUnit:analyzeDeclProto( accessMode, firstToken )

   local nextToken = self:getToken(  )
   local abstractFlag = false
   if nextToken.txt == "abstract" then
      abstractFlag = true
      nextToken = self:getToken(  )
   end
   
   
   if nextToken.txt == "class" or nextToken.txt == "interface" then
      local name = self:getSymbolToken( SymbolMode.MustNot_ )
      local altTypeList = {}
      do
         local workToken = self:getToken(  )
         if workToken.txt == "<" then
            workToken, altTypeList = self:analyzeDeclAlternateType( true, workToken, accessMode )
         end
         
         self:pushbackToken( workToken )
      end
      
      
      if accessMode == Ast.AccessMode.Local then
         accessMode = Ast.AccessMode.Pri
      end
      
      
      local declMode
      
      if nextToken.txt == "class" then
         declMode = DeclClassMode.Class
      else
       
         declMode = TransUnitIF.DeclClassMode.Interface
         abstractFlag = true
      end
      
      local classTypeInfo
      
      local inheritInfo
      
      nextToken, classTypeInfo, inheritInfo = self:analyzePushClass( declMode, abstractFlag, firstToken, name, false, nil, nil, accessMode, altTypeList )
      
      self.protoClassMap[classTypeInfo] = firstToken.pos
      
      self:popClass(  )
      self:checkToken( nextToken, ";" )
      
      return Nodes.ProtoClassNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {classTypeInfo}, name, inheritInfo )
   end
   
   self:error( "illegal proto" )
end


function TransUnit:analyzeDeclEnum( accessMode, firstToken )

   local _
   if Ast.isPubToExternal( accessMode ) and self.scope ~= self.moduleScope and _lune.nilacc( self.scope:get_ownerTypeInfo(), 'get_kind', 'callmtd' ) ~= Ast.TypeInfoKind.Class then
      self:addErrMess( firstToken.pos, "can't external at inner scope." )
   end
   
   
   local name = self:getSymbolToken( SymbolMode.MustNot_ )
   
   self:checkNextToken( "{" )
   
   local valueList = {}
   local scope = self:pushScope( true )
   
   local workEnumTypeInfo = nil
   
   local nextToken = self:getToken(  )
   local number = 0.0
   local prevValTypeInfo = Ast.headTypeInfo
   local valTypeInfo = Ast.headTypeInfo
   while nextToken.txt ~= "}" do
      local valName = self:checkSymbol( nextToken, SymbolMode.MustNot_ )
      
      nextToken = self:getToken(  )
      
      local enumVal = _lune.newAlge( Ast.EnumLiteral.Real, {number})
      do
         local _switchExp = (prevValTypeInfo )
         if _switchExp == Ast.builtinTypeReal then
         elseif _switchExp == Ast.builtinTypeInt or _switchExp == Ast.headTypeInfo then
            enumVal = _lune.newAlge( Ast.EnumLiteral.Int, {math.floor(number)})
         end
      end
      
      
      if nextToken.txt == "=" then
         local exp = self:analyzeExpOneRVal( false, false )
         local literal, mess = exp:getLiteral(  )
         if literal ~= nil then
            do
               local _matchExp = literal
               if _matchExp[1] == Nodes.Literal.Int[1] then
                  local val = _matchExp[2][1]
               
                  enumVal = _lune.newAlge( Ast.EnumLiteral.Int, {val})
                  number = val * 1.0
                  valTypeInfo = Ast.builtinTypeInt
               elseif _matchExp[1] == Nodes.Literal.Real[1] then
                  local val = _matchExp[2][1]
               
                  enumVal = _lune.newAlge( Ast.EnumLiteral.Real, {val})
                  number = val
                  valTypeInfo = Ast.builtinTypeReal
               elseif _matchExp[1] == Nodes.Literal.Str[1] then
                  local val = _matchExp[2][1]
               
                  enumVal = _lune.newAlge( Ast.EnumLiteral.Str, {val})
                  valTypeInfo = Ast.builtinTypeString
               else 
                  
                     self:error( string.format( "illegal enum val -- %s", Nodes.Literal:_getTxt( literal)
                     ) )
               end
            end
            
         else
            self:error( string.format( "illegal enum val -- %s", mess) )
         end
         
         
         nextToken = self:getToken(  )
      else
       
         do
            local _switchExp = (prevValTypeInfo )
            if _switchExp == Ast.headTypeInfo then
               valTypeInfo = Ast.builtinTypeInt
            elseif _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeReal then
               valTypeInfo = prevValTypeInfo
            else 
               
                  self:addErrMess( valName.pos, string.format( "illegal enum val type -- %s", valTypeInfo:getTxt(  )) )
            end
         end
         
      end
      
      if prevValTypeInfo ~= Ast.headTypeInfo and prevValTypeInfo ~= valTypeInfo then
         self:addErrMess( valName.pos, string.format( "multiple enum val type. %s, %s", valTypeInfo:getTxt(  ), prevValTypeInfo:getTxt(  )) )
      end
      
      prevValTypeInfo = valTypeInfo
      
      if not workEnumTypeInfo then
         workEnumTypeInfo = self.processInfo:createEnum( scope, self:getCurrentNamespaceTypeInfo(  ), false, accessMode, name.txt, valTypeInfo )
      end
      
      
      if workEnumTypeInfo ~= nil then
         local evalValSym = _lune.unwrap( scope:addEnumVal( self.processInfo, valName.txt, valName.pos, workEnumTypeInfo ))
         
         local enumValInfo = Ast.EnumValInfo.new(valName.txt, enumVal, evalValSym)
         table.insert( valueList, valName )
         
         workEnumTypeInfo:addEnumValInfo( enumValInfo )
      end
      
      
      if nextToken.txt == "}" then
         break
      end
      
      self:checkToken( nextToken, "," )
      nextToken = self:getToken(  )
      number = number + 1
   end
   
   
   local enumTypeInfo = workEnumTypeInfo
   if  nil == enumTypeInfo then
      local _enumTypeInfo = enumTypeInfo
   
      enumTypeInfo = self.processInfo:createEnum( scope, self:getCurrentNamespaceTypeInfo(  ), false, accessMode, name.txt, Ast.builtinTypeNone )
   end
   
   
   self:popScope(  )
   
   local _2051, shadowing = self.scope:addEnum( self.processInfo, accessMode, name.txt, name.pos, enumTypeInfo )
   self:errorShadowing( name.pos, shadowing )
   
   return Nodes.DeclEnumNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {enumTypeInfo}, enumTypeInfo, accessMode, name, valueList, scope )
end


function TransUnit:analyzeDeclAlge( accessMode, firstToken )

   local _
   self.helperInfo.useAlge = true
   
   local name = self:getSymbolToken( SymbolMode.MustNot_ )
   
   self:checkNextToken( "{" )
   
   local scope = self.scope
   local algeScope = self:pushScope( true )
   
   local algeTypeInfo = self.processInfo:createAlge( algeScope, self:getCurrentNamespaceTypeInfo(  ), false, accessMode, name.txt )
   local _2063, shadowing = scope:addAlge( self.processInfo, accessMode, name.txt, name.pos, algeTypeInfo )
   self:errorShadowing( name.pos, shadowing )
   
   local algeValList = {}
   
   local nextToken = self:getToken(  )
   while nextToken.txt ~= "}" do
      local valName = self:checkSymbol( nextToken, SymbolMode.MustNot_ )
      if algeTypeInfo:getValInfo( valName.txt ) then
         self:addErrMess( valName.pos, string.format( "multiple symbole -- %s", valName.txt) )
      end
      
      
      nextToken = self:getToken(  )
      
      local paramList = {}
      
      local typeInfoList = {}
      if nextToken.txt == "(" then
         while true do
            local paramNameToken
            
            local workToken1 = self:getToken(  )
            local workToken2 = self:getToken(  )
            if workToken2.txt ~= ":" then
               self:pushback(  )
               self:pushback(  )
               paramNameToken = nil
            else
             
               paramNameToken = workToken1
            end
            
            
            local typeNode = self:analyzeRefType( Ast.AccessMode.Pub, false, Ast.isPubToExternal( accessMode ) )
            if self.protoClassMap[typeNode:get_expType()] then
               self:addErrMess( typeNode:get_pos(), string.format( "can't use the prototype class -- %s", typeNode:get_expType():getTxt(  )) )
            end
            
            table.insert( typeInfoList, typeNode:get_expType() )
            table.insert( paramList, Nodes.AlgeValParamInfo.new(paramNameToken, typeNode) )
            
            nextToken = self:getToken(  )
            if nextToken.txt ~= "," then
               self:checkToken( nextToken, ")" )
               nextToken = self:getToken(  )
               break
            end
            
         end
         
      end
      
      
      local workAlgeValSym, algeValSymShadow = algeScope:addAlgeVal( self.processInfo, valName.txt, valName.pos, algeTypeInfo )
      self:errorShadowing( valName.pos, algeValSymShadow )
      local algeValSym = _lune.unwrap( (workAlgeValSym or shadowing ))
      local algeValInfo = Ast.AlgeValInfo.new(valName.txt, typeInfoList, algeTypeInfo, algeValSym)
      algeTypeInfo:addValInfo( algeValInfo )
      
      table.insert( algeValList, Nodes.DeclAlgeValInfo.new(algeValSym, paramList) )
      
      if nextToken.txt == "}" then
         break
      end
      
      self:checkToken( nextToken, "," )
      nextToken = self:getToken(  )
   end
   
   
   self:popScope(  )
   
   return Nodes.DeclAlgeNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {algeTypeInfo}, accessMode, algeTypeInfo, name, algeValList, algeScope )
end


function TransUnit:analyzeAlias( accessMode, firstToken )

   if self.scope ~= self.moduleScope then
      self:addErrMess( firstToken.pos, "alias must use at top scope." )
   end
   
   
   local newToken = self:getToken(  )
   self:checkNextToken( "=" )
   
   local srcToken = self:getToken(  )
   
   local symbolNode = self:analyzeExpSymbol( srcToken, srcToken, ExpSymbolMode.Symbol, nil, true, false )
   
   local newTypeInfo = Ast.builtinTypeNone
   local symbolInfoList = symbolNode:getSymbolInfo(  )
   
   local newSymbolInfo = Ast.dummySymbol
   
   if #symbolInfoList >= 1 then
      local symbolInfo = symbolInfoList[1]
      if newToken.txt:find( "^_" ) and not srcToken.txt:find( "^_" ) or not newToken.txt:find( "^_" ) and srcToken.txt:find( "^_" ) then
         self:addErrMess( firstToken.pos, string.format( "alias symbol unmatch. %s %s", newToken.txt, newToken.txt) )
      else
       
         do
            local _switchExp = symbolInfo:get_kind()
            if _switchExp == Ast.SymbolKind.Typ or _switchExp == Ast.SymbolKind.Fun then
               local aliasSymbolInfo, shadowing = self.scope:addAlias( self.processInfo, newToken.txt, newToken.pos, false, accessMode, self.moduleType, symbolInfo )
               if aliasSymbolInfo ~= nil then
                  newTypeInfo = aliasSymbolInfo:get_typeInfo()
                  newSymbolInfo = aliasSymbolInfo
               else
                  self:errorShadowing( newToken.pos, shadowing )
               end
               
            else 
               
                  self:addErrMess( firstToken.pos, string.format( "can alias symbol -- %s. (%s)", srcToken.txt, Ast.SymbolKind:_getTxt( symbolInfo:get_kind())
                  ) )
            end
         end
         
      end
      
   else
    
      self:addErrMess( firstToken.pos, string.format( "not found symbold -- %s", srcToken.txt) )
   end
   
   self:checkNextToken( ";" )
   
   return Nodes.AliasNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {newTypeInfo}, newSymbolInfo, symbolNode, newTypeInfo )
end


function TransUnit:analyzeRetTypeList( pubToExtFlag, accessMode, token, parentPub )

   local retTypeInfoList = {}
   local retTypeNodeList = {}
   if token.txt == ":" then
      local hasDDDFlag = false
      while true do
         local refTypeNode = self:analyzeRefType( accessMode, true, parentPub )
         if hasDDDFlag then
            self:addErrMess( refTypeNode:get_pos(), "Type exists after '...'." )
         end
         
         local retType = refTypeNode:get_expType()
         if retType:get_kind() == Ast.TypeInfoKind.DDD then
            hasDDDFlag = true
         end
         
         if pubToExtFlag and not Ast.isPubToExternal( retType:get_accessMode() ) then
            self:addErrMess( refTypeNode:get_pos(), string.format( "this is not public type -- %s", retType:getTxt(  )) )
         end
         
         
         table.insert( retTypeInfoList, retType )
         table.insert( retTypeNodeList, refTypeNode )
         token = self:getToken(  )
         if token.txt ~= "," then
            break
         end
         
      end
      
   end
   
   return retTypeInfoList, token, retTypeNodeList
end


function TransUnit:getDefaultAsync( kind, classTypeInfo, asyncMode )

   if asyncMode ~= nil then
      return asyncMode
   end
   
   
   local function process( defaultAsyncMode )
   
      do
         local _switchExp = defaultAsyncMode
         if _switchExp == DefaultAsyncMode.AsyncAll or _switchExp == DefaultAsyncMode.AsyncFunc then
            return Ast.Async.Async
         elseif _switchExp == DefaultAsyncMode.NoAsync then
            return Ast.Async.Noasync
         end
      end
      
   end
   
   if classTypeInfo ~= nil then
      do
         local _exp = self.class2defaultAsyncMode[classTypeInfo]
         if _exp ~= nil then
            return process( _exp )
         end
      end
      
   end
   
   
   do
      local _switchExp = kind
      if _switchExp == Ast.TypeInfoKind.Method then
         if self.defaultAsyncMode == DefaultAsyncMode.AsyncAll then
            return Ast.Async.Async
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.FormFunc then
         return process( self.defaultAsyncMode )
      end
   end
   
   return Ast.Async.Noasync
end


function TransUnit:getMutableAsync( token )

   local mutable = false
   local asyncMode = nil
   
   while true do
      if token.txt == "mut" then
         mutable = true
         token = self:getToken(  )
      elseif token.txt == "__async" then
         asyncMode = Ast.Async.Async
         token = self:getToken(  )
      elseif token.txt == "__noasync" then
         asyncMode = Ast.Async.Noasync
         token = self:getToken(  )
      elseif token.txt == "__trans" then
         asyncMode = Ast.Async.Transient
         token = self:getToken(  )
      else
       
         break
      end
      
   end
   
   return token, mutable, asyncMode
end


function TransUnit:analyzeDeclForm( accessMode, firstToken )

   local _
   local name = self:getSymbolToken( SymbolMode.MustNot_ )
   
   if self.scope ~= self.moduleScope and Ast.isPubToExternal( accessMode ) then
      self:addErrMess( firstToken.pos, string.format( "You must declare this form at the outside scope. -- %s", name.txt) )
   end
   
   
   self:checkNextToken( "(" )
   local argList = {}
   local funcBodyScope = self:pushScope( false )
   
   local nextToken = self:analyzeDeclArgList( accessMode, funcBodyScope, argList, Ast.isPubToExternal( accessMode ) )
   
   self:checkToken( nextToken, ")" )
   
   local asyncMode
   
   nextToken, _, asyncMode = self:getMutableAsync( self:getToken(  ) )
   
   local retTypeList = {}
   local retNodeList
   
   retTypeList, nextToken, retNodeList = self:analyzeRetTypeList( Ast.isPubToExternal( accessMode ), accessMode, nextToken, Ast.isPubToExternal( accessMode ) )
   
   self:checkToken( nextToken, ";" )
   
   self:popScope(  )
   
   local argTypeInfoList = {}
   for __index, argNode in ipairs( argList ) do
      table.insert( argTypeInfoList, argNode:get_expType() )
   end
   
   
   local formType = self.processInfo:createFunc( false, false, nil, Ast.TypeInfoKind.FormFunc, self:getCurrentNamespaceTypeInfo(  ), false, false, true, accessMode, name.txt, self:getDefaultAsync( Ast.TypeInfoKind.FormFunc, self:getCurrentClass(  ), asyncMode ), nil, argTypeInfoList, retTypeList, false )
   
   local formSymbol, shadowing = self.scope:addForm( self.processInfo, name.pos, formType, accessMode )
   self:errorShadowing( name.pos, shadowing )
   
   local declFuncInfo = Nodes.DeclFuncInfo.new(Nodes.FuncKind.Form, nil, nil, name, formSymbol or shadowing, argList, false, accessMode, asyncMode, nil, retTypeList, retNodeList, false, false)
   
   return Nodes.DeclFormNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {formType}, declFuncInfo )
end


function TransUnit:analyzeDecl( accessMode, staticFlag, firstToken, token )

   if not staticFlag then
      if token.txt == "static" then
         staticFlag = true
         token = self:getToken(  )
      end
      
   end
   
   local overrideFlag = false
   if token.txt == "override" then
      overrideFlag = true
      token = self:getToken(  )
   end
   
   
   local abstractFlag = false
   if token.txt == "abstract" then
      abstractFlag = true
      token = self:getToken(  )
   end
   
   
   if token.txt == "let" then
      return self:analyzeDeclVar( Nodes.DeclVarMode.Let, accessMode, firstToken )
   elseif token.txt == "fn" then
      local nextToken = self:getToken(  )
      self:pushback(  )
      if nextToken.kind == Parser.TokenKind.Symb or Ast.isPubToExternal( accessMode ) or staticFlag or overrideFlag or abstractFlag then
         return self:analyzeDeclFunc( DeclFuncMode.Func, abstractFlag, overrideFlag, accessMode, staticFlag, nil, firstToken, nil )
      end
      
   elseif token.txt == "class" then
      return self:analyzeDeclClass( abstractFlag, accessMode, firstToken, TransUnitIF.DeclClassMode.Class )
   elseif token.txt == "interface" then
      return self:analyzeDeclClass( true, accessMode, firstToken, TransUnitIF.DeclClassMode.Interface )
   elseif token.txt == "module" then
      return self:analyzeDeclClass( false, accessMode, firstToken, TransUnitIF.DeclClassMode.Module )
   elseif token.txt == "proto" then
      return self:analyzeDeclProto( accessMode, firstToken )
   elseif token.txt == "macro" then
      return self:analyzeDeclMacro( accessMode, firstToken )
   elseif token.txt == "enum" then
      return self:analyzeDeclEnum( accessMode, firstToken )
   elseif token.txt == "alge" then
      return self:analyzeDeclAlge( accessMode, firstToken )
   elseif token.txt == "form" then
      return self:analyzeDeclForm( accessMode, firstToken )
   elseif token.txt == "alias" then
      return self:analyzeAlias( accessMode, firstToken )
   elseif token.txt == "__test" then
      return self:analyzeTest( firstToken )
   end
   
   
   return nil
end


function TransUnit:checkPublic( pos, typeInfo )

   local checkedTypeSet = {}
   local function checkPub( workType )
   
      if _lune._Set_has(checkedTypeSet, workType ) then
         return 
      end
      
      checkedTypeSet[workType]= true
      if workType:get_kind() ~= Ast.TypeInfoKind.Array and workType:get_kind() ~= Ast.TypeInfoKind.List and workType:get_kind() ~= Ast.TypeInfoKind.Set and workType:get_kind() ~= Ast.TypeInfoKind.Map and not Ast.isPubToExternal( workType:get_accessMode() ) then
         self:addErrMess( pos, string.format( "not public this type -- %s", workType:getTxt(  )) )
      else
       
         for __index, itemType in ipairs( workType:get_itemTypeInfoList() ) do
            checkPub( itemType )
         end
         
      end
      
   end
   checkPub( typeInfo )
end


function TransUnit:analyzeDeclMember( classTypeInfo, accessMode, staticFlag, firstToken )
   local __func__ = '@lune.@base.@TransUnit.TransUnit.analyzeDeclMember'

   local nextToken = self:getToken(  )
   local mutMode = Ast.MutMode.IMut
   do
      local _switchExp = nextToken.txt
      if _switchExp == "mut" then
         mutMode = Ast.MutMode.Mut
         nextToken = self:getToken(  )
      elseif _switchExp == "allmut" then
         mutMode = Ast.MutMode.AllMut
         nextToken = self:getToken(  )
      end
   end
   
   local varName = self:checkSymbol( nextToken, SymbolMode.MustNot_ )
   local token = self:getToken(  )
   local refType = self:analyzeRefType( accessMode, false, Ast.isPubToExternal( classTypeInfo:get_accessMode() ) )
   token = self:getToken(  )
   
   if refType:get_expType():get_asyncMode() == Ast.Async.Transient then
      self:addErrMess( refType:get_pos(), string.format( "can't hold with the type of __trans. -- %s", varName.txt) )
   end
   
   local getterMode = Ast.AccessMode.None
   local getterRetType = refType:get_expType()
   local getterToken = nil
   local getterMutable = Ast.MutMode.Mut
   local setterMode = Ast.AccessMode.None
   local setterToken = nil
   if token.txt == "{" then
      
      local function analyzeAccessorMode(  )
      
         local retType = Ast.headTypeInfo
         local mode = Ast.AccessMode.None
         local accessorToken = self:getToken(  )
         local workToken = accessorToken
         do
            local _switchExp = workToken.txt
            if _switchExp == "pub" or _switchExp == "pri" or _switchExp == "pro" or _switchExp == "local" then
               mode = _lune.unwrap( Ast.txt2AccessMode( workToken.txt ))
               workToken = self:getToken(  )
               if workToken.txt == "&" then
                  getterMutable = Ast.MutMode.IMut
                  workToken = self:getToken(  )
               end
               
               if workToken.txt == ":" then
                  local typeNode = self:analyzeRefType( mode, false, Ast.isPubToExternal( classTypeInfo:get_accessMode() ) )
                  retType = typeNode:get_expType()
                  workToken = self:getToken(  )
               end
               
            elseif _switchExp == "non" then
               workToken = self:getToken(  )
            else 
               
                  self:addErrMess( workToken.pos, string.format( "access mode is invalid -- %s", workToken.txt) )
            end
         end
         
         return mode, retType, accessorToken, workToken
      end
      
      do
         local workRetType
         
         getterMode, workRetType, getterToken, nextToken = analyzeAccessorMode(  )
         if workRetType ~= Ast.headTypeInfo then
            if not workRetType:canEvalWith( self.processInfo, getterRetType, Ast.CanEvalType.SetOp, classTypeInfo:createAlt2typeMap( false ) ) then
               self:addErrMess( firstToken.pos, string.format( "getter type mismatch -- %s <- %s", workRetType:getTxt(  ), getterRetType:getTxt(  )) )
            end
            
            getterRetType = workRetType
         end
         
      end
      
      if nextToken.txt == "," then
         local dummyRetType
         
         setterMode, dummyRetType, setterToken, nextToken = analyzeAccessorMode(  )
         if setterMode ~= Ast.AccessMode.None and mutMode == Ast.MutMode.IMut then
            self:addErrMess( varName.pos, string.format( "This member can't have setter, this member is immutable. -- %s", varName.txt) )
         end
         
         Log.log( Log.Level.Debug, __func__, 1853, function (  )
         
            return string.format( "%s", dummyRetType)
         end )
         
      end
      
      self:checkToken( nextToken, "}" )
      token = self:getToken(  )
   end
   
   
   self:checkToken( token, ";" )
   
   local typeInfo = refType:get_expType()
   if self.ctrl_info.legacyMutableControl then
      if Ast.TypeInfo.isMut( typeInfo ) and typeInfo:get_mutMode() ~= mutMode then
         typeInfo = self:createModifier( typeInfo, mutMode )
      end
      
      if Ast.TypeInfo.isMut( getterRetType ) and getterRetType:get_mutMode() ~= mutMode then
         getterRetType = self:createModifier( getterRetType, mutMode )
      end
      
   else
    
      if Ast.TypeInfo.isMut( getterRetType ) then
         if mutMode == Ast.MutMode.AllMut or getterMutable == Ast.MutMode.AllMut then
            getterRetType = self:createModifier( getterRetType, Ast.MutMode.AllMut )
         elseif getterMutable == Ast.MutMode.IMut then
            getterRetType = self:createModifier( getterRetType, Ast.MutMode.IMut )
         end
         
      end
      
   end
   
   
   if Ast.isPubToExternal( classTypeInfo:get_accessMode() ) then
      if Ast.isPubToExternal( accessMode ) or Ast.isPubToExternal( setterMode ) then
         self:checkPublic( refType:get_pos(), typeInfo )
      end
      
      if Ast.isPubToExternal( getterMode ) then
         self:checkPublic( refType:get_pos(), getterRetType )
      end
      
   end
   
   
   local symbolInfo, shadowing = self.scope:addMember( self.processInfo, varName.txt, varName.pos, typeInfo, accessMode, staticFlag, mutMode )
   
   local workSym = _lune.unwrap( (symbolInfo or shadowing ))
   if shadowing ~= nil then
      self:errorShadowing( varName.pos, shadowing )
   end
   
   
   return Nodes.DeclMemberNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, varName, refType, workSym, classTypeInfo, staticFlag, accessMode, getterMutable ~= Ast.MutMode.IMut, getterMode, getterToken, getterRetType, setterMode, setterToken )
end


function TransUnit:analyzeDeclMethod( classTypeInfo, declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, firstToken, name )

   local node = self:analyzeDeclFunc( declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, classTypeInfo, name, name )
   return node
end


function TransUnit:addDefaultConstructor( pos, classTypeInfo, classScope, memberNodeList, methodNameSet, oldFlag )

   if classScope:getTypeInfoChild( "__init" ) then
      self:addErrMess( pos, "already declare __init()." )
   end
   
   local argTypeList = {}
   
   if classTypeInfo:get_baseTypeInfo() ~= Ast.headTypeInfo then
      local superScope = _lune.unwrap( classTypeInfo:get_baseTypeInfo():get_scope())
      local superTypeInfo = _lune.unwrap( superScope:getTypeInfoChild( "__init" ))
      for __index, argType in ipairs( superTypeInfo:get_argTypeInfoList() ) do
         if oldFlag then
            if not argType:get_nilable() then
               self:addErrMess( pos, "not found '__init' decl." )
            end
            
         else
          
            table.insert( argTypeList, argType )
         end
         
      end
      
   end
   
   
   for __index, memberNode in ipairs( memberNodeList ) do
      if not memberNode:get_staticFlag() then
         table.insert( argTypeList, memberNode:get_expType() )
      end
      
   end
   
   
   if Ast.isPubToExternal( classTypeInfo:get_accessMode() ) then
      for __index, memberType in ipairs( argTypeList ) do
         if not Ast.isPubToExternal( memberType:get_accessMode() ) then
            self:addErrMess( pos, string.format( "The type must be 'pub' becaue using in __init(). -- %s:%s", memberType:getTxt(  ), Ast.AccessMode:_getTxt( memberType:get_accessMode())
            ) )
         end
         
      end
      
   end
   
   local ctorScope = self:pushScope( false )
   local initTypeInfo = self.processInfo:createFunc( false, false, ctorScope, Ast.TypeInfoKind.Method, classTypeInfo, true, false, false, Ast.AccessMode.Pub, "__init", Ast.Async.Async, nil, argTypeList, {} )
   if oldFlag then
      ctorScope:addVar( self.processInfo, Ast.AccessMode.Pri, "", nil, Ast.headTypeInfo, Ast.MutMode.IMut, true )
   end
   
   self:popScope(  )
   classScope:addMethod( self.processInfo, pos, initTypeInfo, Ast.AccessMode.Pub, false, false )
   methodNameSet["__init"]= true
   for __index, memberNode in ipairs( memberNodeList ) do
      if not memberNode:get_staticFlag() then
         memberNode:get_symbolInfo():updateValue( memberNode:get_symbolInfo():get_posForLatestMod() )
      end
      
   end
   
end

function TransUnit:analyzeFuncBlock( analyzingState, firstToken, classTypeInfo, funcTypeInfo, funcName, funcBodyScope, retTypeInfoList )

   if not funcTypeInfo:get_staticFlag() then
      if classTypeInfo then
         do
            local overrideType = self.scope:get_parent():getTypeInfoField( funcName, false, funcBodyScope, self.scopeAccess )
            if overrideType ~= nil then
               if not overrideType:get_abstractFlag() then
                  funcBodyScope:add( self.processInfo, Ast.SymbolKind.Fun, false, false, "super", nil, overrideType, Ast.AccessMode.Local, false, Ast.MutMode.IMut, true, false )
               end
               
            end
         end
         
      end
      
   end
   
   
   self:pushAnalyzingState( analyzingState )
   
   local body = self:analyzeBlock( Nodes.BlockKind.Func, TentativeMode.Ignore, funcBodyScope, nil )
   
   self:popAnalyzingState(  )
   
   if #retTypeInfoList ~= 0 then
      local breakKind = body:getBreakKind( Nodes.CheckBreakMode.Return )
      if retTypeInfoList[1] ~= Ast.builtinTypeNeverRet then
         do
            local _switchExp = breakKind
            if _switchExp == Nodes.BreakKind.Return or _switchExp == Nodes.BreakKind.NeverRet then
            else 
               
                  self:addErrMess( firstToken.pos, "This funcion doesn't have return." )
            end
         end
         
      else
       
         if breakKind ~= Nodes.BreakKind.NeverRet then
            self:addErrMess( firstToken.pos, string.format( "This funcion must be never return. -- %s", Nodes.BreakKind:_getTxt( breakKind)
            ) )
         end
         
      end
      
   end
   
   return body
end


function TransUnit:addAccessor( memberNode, methodNameSet, classScope, classTypeInfo )

   local memberType = memberNode:get_expType()
   local memberName = memberNode:get_name()
   local getterName = "get_" .. memberName.txt
   local accessMode = memberNode:get_getterMode()
   local typeKind
   
   if memberNode:get_staticFlag() then
      typeKind = Ast.TypeInfoKind.Func
   else
    
      typeKind = Ast.TypeInfoKind.Method
   end
   
   if accessMode ~= Ast.AccessMode.None then
      if classScope:getTypeInfoChild( getterName ) then
         self:addErrMess( memberName.pos, string.format( "exist -- %s.%s", classTypeInfo:get_rawTxt(), getterName) )
      else
       
         local mutable = memberNode:get_getterMutable()
         local getterMemberType = memberNode:get_getterRetType()
         if Ast.TypeInfo.isMut( getterMemberType ) and not mutable then
            getterMemberType = self:createModifier( getterMemberType, Ast.MutMode.IMut )
         end
         
         local retTypeInfo = self.processInfo:createFunc( false, false, self:pushScope( false ), typeKind, classTypeInfo, false, false, memberNode:get_staticFlag(), accessMode, getterName, self:getDefaultAsync( typeKind, classTypeInfo, nil ), nil, {}, {getterMemberType} )
         self:popScope(  )
         
         classScope:addMethod( self.processInfo, memberName.pos, retTypeInfo, accessMode, memberNode:get_staticFlag(), false )
         methodNameSet[getterName]= true
      end
      
   end
   
   local setterName = "set_" .. memberName.txt
   accessMode = memberNode:get_setterMode()
   if memberNode:get_setterMode() ~= Ast.AccessMode.None then
      if classScope:getTypeInfoChild( setterName ) then
         self:addErrMess( memberName.pos, string.format( "exist -- %s.%s", classTypeInfo:get_rawTxt(), setterName) )
      else
       
         local mutable
         
         if memberNode:get_symbolInfo():get_mutMode() ~= Ast.MutMode.AllMut then
            mutable = true
         else
          
            mutable = false
         end
         
         classScope:addMethod( self.processInfo, memberName.pos, self.processInfo:createFunc( false, false, self:pushScope( false ), typeKind, classTypeInfo, false, false, memberNode:get_staticFlag(), accessMode, setterName, self:getDefaultAsync( typeKind, classTypeInfo, nil ), nil, {memberType}, nil, mutable ), accessMode, memberNode:get_staticFlag(), true )
         self:popScope(  )
         methodNameSet[setterName]= true
      end
      
   end
   
end


function TransUnit:analyzeClassBody( hasProto, classAccessMode, firstToken, mode, gluePrefix, classTypeInfo, name, moduleLang, moduleName, lazyLoad, nextToken, inheritInfo )

   local memberName2Node = {}
   local allStmtList = {}
   local declStmtList = {}
   local fieldList = {}
   local memberList = {}
   local methodNameSet = {}
   local initBlockInfo = Nodes.ClassInitBlockInfo.new()
   local advertiseList = {}
   local trustList = {}
   local uninitMemberList = {}
   local node = Nodes.DeclClassNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {classTypeInfo}, classAccessMode, name, inheritInfo, hasProto, gluePrefix, moduleName, moduleLang, lazyLoad, false, allStmtList, declStmtList, fieldList, memberList, self.scope, initBlockInfo, advertiseList, trustList, uninitMemberList, {} )
   self.typeInfo2ClassNode[classTypeInfo] = node
   local alreadyCtorFlag = false
   local hasInitBlock = false
   local hasStaticMember = false
   local classScope = self.scope
   
   local function processLet( token, staticFlag, accessMode, alreadyFlag )
   
      if staticFlag then
         hasStaticMember = true
      end
      
      if mode == TransUnitIF.DeclClassMode.Interface then
         self:addErrMess( token.pos, "interface can not have member" )
      end
      
      if not staticFlag and alreadyFlag then
         self:addErrMess( token.pos, "member can't declare after '__init' method." )
      elseif staticFlag and hasInitBlock then
         self:addErrMess( token.pos, "static member can't declare after '__init' block." )
      end
      
      local memberNode = self:analyzeDeclMember( classTypeInfo, accessMode, staticFlag, token )
      table.insert( allStmtList, memberNode )
      table.insert( fieldList, memberNode )
      table.insert( memberList, memberNode )
      memberName2Node[memberNode:get_name().txt] = memberNode
      self:addAccessor( memberNode, methodNameSet, classScope, classTypeInfo )
   end
   local function checkInitializeMember( staticFlag, pos )
   
      for memberName, memberNode in pairs( memberName2Node ) do
         if memberNode:get_staticFlag() == staticFlag then
            local symbolInfo = _lune.unwrap( self.scope:getSymbolInfoChild( memberName ))
            local typeInfo = symbolInfo:get_typeInfo()
            if not symbolInfo:get_hasValueFlag() then
               local msg
               
               if staticFlag then
                  msg = string.format( "Set member -- %s", memberName)
               else
                
                  msg = string.format( "Set member -- %s.%s", name.txt, memberName)
               end
               
               if not typeInfo:get_nilable() then
                  self:addErrMess( _lune.unwrapDefault( pos, memberNode:get_pos()), msg )
               else
                
                  table.insert( uninitMemberList, symbolInfo )
                  self:addWarnMess( _lune.unwrapDefault( pos, memberNode:get_pos()), msg )
               end
               
            end
            
         end
         
      end
      
   end
   
   local function processFn( token, staticFlag, accessMode, abstractFlag, overrideFlag )
   
      local nameToken = self:getSymbolToken( SymbolMode.MustNot_ )
      local declFuncMode = DeclFuncMode.Class
      if mode == TransUnitIF.DeclClassMode.Module or mode == TransUnitIF.DeclClassMode.LazyModule then
         if gluePrefix then
            declFuncMode = DeclFuncMode.Glue
         else
          
            declFuncMode = DeclFuncMode.Module
         end
         
      end
      
      
      if nameToken.txt == "__init" then
         for __index, symbolInfo in pairs( self.scope:get_symbol2SymbolInfoMap() ) do
            if not symbolInfo:get_staticFlag() then
               symbolInfo:clearValue(  )
            end
            
         end
         
      end
      
      
      local methodNode = self:analyzeDeclMethod( classTypeInfo, declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, token, nameToken )
      table.insert( allStmtList, methodNode )
      table.insert( fieldList, methodNode )
      methodNameSet[nameToken.txt]= true
      if nameToken.txt == "__init" then
         alreadyCtorFlag = true
         
         checkInitializeMember( false, methodNode:get_pos() )
      end
      
   end
   
   local function processInitBlock( token )
   
      
      if _lune.nilacc( classTypeInfo:get_scope(), 'get_parent', 'callmtd' ) ~= self.moduleScope then
         self:addErrMess( token.pos, "The '__init' block only support at the top level classes." )
      end
      
      
      if mode ~= TransUnitIF.DeclClassMode.Class then
         self:error( string.format( "%s can not have __init block.", mode) )
      end
      
      hasInitBlock = true
      for __index, symbolInfo in pairs( self.scope:get_symbol2SymbolInfoMap() ) do
         if symbolInfo:get_staticFlag() then
            symbolInfo:clearValue(  )
         end
         
      end
      
      
      local parentScope = self.scope
      local initBlockScope = self:pushScope( false )
      self:prepareTentativeSymbol( initBlockScope, false, nil )
      
      local ininame = "___init"
      local funcTypeInfo = self.processInfo:createFunc( false, false, initBlockScope, Ast.TypeInfoKind.Func, classTypeInfo, false, false, true, Ast.AccessMode.Pri, ininame, Ast.Async.Noasync, nil, nil, nil, false )
      local funcSym, shadowing = parentScope:addFunc( self.processInfo, token.pos, funcTypeInfo, Ast.AccessMode.Pri, true, true )
      
      local block = self:analyzeFuncBlock( AnalyzingState.InitBlock, token, classTypeInfo, funcTypeInfo, ininame, initBlockScope, funcTypeInfo:get_retTypeInfoList() )
      
      local info = Nodes.DeclFuncInfo.new(Nodes.FuncKind.InitBlock, classTypeInfo, node, token, _lune.unwrap( (funcSym or shadowing )), {}, true, Ast.AccessMode.Pri, nil, block, {}, {}, false, false)
      local initBlockNode = Nodes.DeclMethodNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {funcTypeInfo}, info )
      initBlockInfo:set_func( initBlockNode )
      table.insert( allStmtList, initBlockNode )
      self:popScope(  )
      self:finishTentativeSymbol( false )
      
   end
   
   local function processAdvertise(  )
   
      local memberToken = self:getSymbolToken( SymbolMode.MustNot_ )
      nextToken = self:getToken(  )
      local prefix = ""
      if nextToken.txt ~= ";" and nextToken.txt ~= "{" then
         prefix = nextToken.txt
         nextToken = self:getToken(  )
      end
      
      self:checkToken( nextToken, ";" )
      local memberNode = memberName2Node[memberToken.txt]
      if  nil == memberNode then
         local _memberNode = memberNode
      
         self:error( string.format( "not found member -- %s", memberToken.txt) )
      end
      
      local advInfo = Nodes.AdvertiseInfo.new(memberNode, prefix, memberToken.pos)
      table.insert( advertiseList, advInfo )
      table.insert( allStmtList, Nodes.DeclAdvertiseNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, advInfo ) )
      
      self.advertisedTypeSet[memberNode:get_expType():get_srcTypeInfo():get_genSrcTypeInfo()]= true
   end
   
   local function processEnum( token, accessMode )
   
      if accessMode ~= Ast.AccessMode.Pri and (classAccessMode == Ast.AccessMode.Pri or classAccessMode == Ast.AccessMode.Local ) then
         self:addErrMess( token.pos, string.format( "unmatch access mode, class('%s') and enum('%s')", Ast.AccessMode:_getTxt( classAccessMode)
         , Ast.AccessMode:_getTxt( accessMode)
         ) )
      end
      
      local enumNode = self:analyzeDeclEnum( accessMode, token )
      table.insert( allStmtList, enumNode )
      table.insert( declStmtList, enumNode )
   end
   
   local function processLuneControl(  )
   
      nextToken = self:getToken(  )
      
      local pragma
      
      do
         local _switchExp = nextToken.txt
         if _switchExp == "default__init" then
            pragma = _lune.newAlge( LuneControl.Pragma.default__init)
            
            alreadyCtorFlag = true
            
            self:addDefaultConstructor( nextToken.pos, classTypeInfo, self.scope, memberList, methodNameSet, false )
         elseif _switchExp == "default__init_old" then
            pragma = _lune.newAlge( LuneControl.Pragma.default__init_old)
            
            alreadyCtorFlag = true
            self:addDefaultConstructor( nextToken.pos, classTypeInfo, self.scope, memberList, methodNameSet, true )
            node:setHasOldCtor(  )
         elseif _switchExp == "default_async_this_class" then
            pragma = _lune.newAlge( LuneControl.Pragma.default_async_this_class)
            self.class2defaultAsyncMode[self:getCurrentNamespaceTypeInfo(  )] = DefaultAsyncMode.AsyncAll
         elseif _switchExp == "default_noasync_this_class" then
            pragma = _lune.newAlge( LuneControl.Pragma.default_noasync_this_class)
            self.class2defaultAsyncMode[self:getCurrentNamespaceTypeInfo(  )] = DefaultAsyncMode.NoAsync
         else 
            
               self:error( string.format( "unknown option -- %s", nextToken.txt) )
         end
      end
      
      self:checkNextToken( ";" )
      
      local ctrlNode = Nodes.LuneControlNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, pragma )
      self.helperInfo.pragmaSet[pragma]= true
      table.insert( allStmtList, ctrlNode )
   end
   
   local function processClassFields( inMacro )
   
      while true do
         local token = self:getToken( inMacro )
         if token.kind == Parser.TokenKind.Eof or token.txt == "}" then
            break
         end
         
         local accessMode = Ast.txt2AccessMode( token.txt )
         if  nil == accessMode then
            local _accessMode = accessMode
         
            accessMode = Ast.AccessMode.Pri
         else
            
               token = self:getToken(  )
         end
         
         if mode == TransUnitIF.DeclClassMode.Interface and accessMode ~= Ast.AccessMode.Pub then
            self:addErrMess( token.pos, "interface's fields must be 'pub'." )
         end
         
         
         local staticFlag = false
         if token.txt == "static" then
            staticFlag = true
            token = self:getToken(  )
         end
         
         local overrideFlag = false
         if token.txt == "override" then
            overrideFlag = true
            token = self:getToken(  )
         end
         
         local abstractFlag = false
         if token.txt == "abstract" then
            abstractFlag = true
            token = self:getToken(  )
         elseif mode == TransUnitIF.DeclClassMode.Interface then
            abstractFlag = true
         end
         
         
         if token.txt == "let" then
            processLet( token, staticFlag, accessMode, alreadyCtorFlag )
         elseif token.txt == "fn" then
            processFn( token, staticFlag, accessMode, abstractFlag, overrideFlag )
         elseif token.txt == "__init" then
            processInitBlock( token )
         elseif token.txt == "advertise" then
            processAdvertise(  )
         elseif token.txt == ";" then
         elseif token.txt == "enum" then
            processEnum( token, accessMode )
         elseif token.txt == "_lune_control" then
            processLuneControl(  )
         else
          
            do
               local symbolInfo = self.scope:getSymbolInfo( token.txt, self.scope, false, self.scopeAccess )
               if symbolInfo ~= nil then
                  if symbolInfo:get_kind() == Ast.SymbolKind.Mac then
                     
                     self:checkNextToken( "(" )
                     
                     local alt2typeMap, argList = self:prepareExpCall( token.pos, symbolInfo:get_typeInfo(), {}, Ast.headTypeInfo )
                     
                     self:evalMacroOp( token, symbolInfo:get_typeInfo(), argList, function (  )
                     
                        processClassFields( true )
                     end )
                     
                     self:checkNextToken( ";" )
                  else
                   
                     self:error( "illegal field" )
                  end
                  
               else
                  self:error( "illegal field" )
               end
            end
            
         end
         
      end
      
   end
   
   processClassFields( false )
   
   do
      local _switchExp = mode
      if _switchExp == TransUnitIF.DeclClassMode.Module or _switchExp == TransUnitIF.DeclClassMode.LazyModule then
      else 
         
            if hasStaticMember and not hasInitBlock then
               self:addErrMess( node:get_pos(), string.format( "This class (%s) need __init block for initialize static members.", name.txt) )
            end
            
            
            checkInitializeMember( true, nil )
      end
   end
   
   
   return node, nextToken, methodNameSet
end


function TransUnit:analyzeDeclClass( classAbstructFlag, classAccessMode, firstToken, mode )

   local _
   if mode == TransUnitIF.DeclClassMode.Module then
      if self:getToken(  ).txt == "." then
         do
            local _switchExp = self:getToken(  ).txt
            if _switchExp == "l" then
               mode = TransUnitIF.DeclClassMode.LazyModule
            elseif _switchExp == "d" then
               mode = TransUnitIF.DeclClassMode.Module
            end
         end
         
      else
       
         self:pushback(  )
         if self.ctrl_info.defaultLazy then
            mode = TransUnitIF.DeclClassMode.LazyModule
         end
         
      end
      
   end
   
   if mode == TransUnitIF.DeclClassMode.LazyModule then
      self.helperInfo.useLazyRequire = true
   end
   
   
   do
      local _switchExp = mode
      if _switchExp == TransUnitIF.DeclClassMode.Module or _switchExp == TransUnitIF.DeclClassMode.LazyModule then
      else 
         
            do
               local _switchExp = self:getCurrentNamespaceTypeInfo(  ):get_kind()
               if _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Module then
               elseif _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Method then
                  do
                     local _switchExp = classAccessMode
                     if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
                        self:addErrMess( firstToken.pos, "Class can't declare on here." )
                     end
                  end
                  
               else 
                  
                     self:addErrMess( firstToken.pos, "Class can't declare on here." )
               end
            end
            
      end
   end
   
   
   local name = self:getSymbolToken( SymbolMode.MustNot_ )
   local altTypeList = {}
   do
      local nextToken = self:getToken(  )
      if nextToken.txt == "<" then
         nextToken, altTypeList = self:analyzeDeclAlternateType( true, nextToken, classAccessMode )
      end
      
      self:pushbackToken( nextToken )
      
      if #altTypeList > 0 and mode ~= TransUnitIF.DeclClassMode.Class then
         self:addErrMess( name.pos, string.format( "Only class can use the generics. -- %s ", name.txt) )
      end
      
   end
   
   
   if classAccessMode == Ast.AccessMode.Local then
      classAccessMode = Ast.AccessMode.Pri
   end
   
   
   local moduleName = nil
   local gluePrefix = nil
   local moduleLang = nil
   if mode == TransUnitIF.DeclClassMode.Module or mode == TransUnitIF.DeclClassMode.LazyModule then
      self:checkNextToken( "require" )
      moduleName = self:getToken(  )
      local nextToken = self:getToken(  )
      if nextToken.txt == "of" then
         local langToken = self:getToken(  )
         if langToken.kind ~= Parser.TokenKind.Str then
            self:error( string.format( "it's not a string -- %s", langToken.txt) )
         end
         
         local langIdToken = langToken:getExcludedDelimitTxt(  )
         if langIdToken ~= "" then
            for __index, langId in ipairs( Types.Lang.get__allList() ) do
               do
                  local _exp = Types.Lang._from( langId )
                  if _exp ~= nil then
                     local ldName = Types.Lang:_getTxt( _exp)
                     :gsub( ".*%.", "" )
                     if ldName == langIdToken then
                        moduleLang = _exp
                        break
                     end
                     
                  end
               end
               
            end
            
            if moduleLang == nil then
               self:errorAt( langToken.pos, string.format( "invalid lang -- %s", langToken.txt) )
            end
            
         else
          
            moduleLang = Types.Lang.Same
         end
         
         
         nextToken = self:getToken(  )
      end
      
      if nextToken.txt == "glue" then
         gluePrefix = self:getToken(  ):getExcludedDelimitTxt(  )
      else
       
         self:pushback(  )
      end
      
   end
   
   
   local existSymbolInfo = self.scope:getSymbolTypeInfo( name.txt, self.scope, self.scope, self.scopeAccess )
   
   local nextToken, classTypeInfo, inheritInfo = self:analyzePushClass( mode, classAbstructFlag, firstToken, name, true, moduleName, moduleLang or Types.Lang.Same, classAccessMode, altTypeList )
   
   local hasProto
   
   if self.protoClassMap[classTypeInfo] then
      self.protoClassMap[classTypeInfo] = nil
      hasProto = true
   else
    
      hasProto = false
      if existSymbolInfo then
         self:addErrMess( name.pos, string.format( "already declare symbol -- %s", name.txt) )
      end
      
   end
   
   
   local classScope = self.scope
   
   self:checkToken( nextToken, "{" )
   
   local mapType = self.processInfo:createMap( Ast.AccessMode.Pub, classTypeInfo, Ast.builtinTypeString, self:createModifier( Ast.builtinTypeStem, Ast.MutMode.IMut ), Ast.MutMode.IMut )
   if classTypeInfo:isInheritFrom( self.processInfo, Ast.builtinTypeMapping, nil ) then
      self.helperInfo.hasMappingClassDef = true
      
      if classTypeInfo:get_baseTypeInfo() ~= Ast.headTypeInfo and not classTypeInfo:get_baseTypeInfo():isInheritFrom( self.processInfo, Ast.builtinTypeMapping, nil ) then
         self:addErrMess( firstToken.pos, string.format( "must extend Mapping at %s", classTypeInfo:get_baseTypeInfo():getTxt(  )) )
      end
      
      local toMapFuncTypeInfo = self.processInfo:createFunc( false, false, nil, Ast.TypeInfoKind.Method, classTypeInfo, true, false, false, Ast.AccessMode.Pub, "_toMap", Ast.Async.Async, nil, {}, {mapType}, false )
      classScope:addMethod( self.processInfo, nil, toMapFuncTypeInfo, Ast.AccessMode.Pub, false, false )
   end
   
   
   local lazyLoad
   
   do
      local _switchExp = mode
      if _switchExp == TransUnitIF.DeclClassMode.LazyModule then
         lazyLoad = Nodes.LazyLoad.On
      elseif _switchExp == TransUnitIF.DeclClassMode.Module or _switchExp == TransUnitIF.DeclClassMode.Class or _switchExp == TransUnitIF.DeclClassMode.Interface then
         lazyLoad = Nodes.LazyLoad.Off
      end
   end
   
   
   local node, _2619, methodNameSet = self:analyzeClassBody( hasProto, classAccessMode, firstToken, mode, gluePrefix, classTypeInfo, name, moduleLang, moduleName, lazyLoad, nextToken, inheritInfo )
   local ctorAccessMode = Ast.AccessMode.Pub
   do
      local ctorTypeInfo = classScope:getTypeInfoChild( "__init" )
      if ctorTypeInfo ~= nil then
         ctorAccessMode = ctorTypeInfo:get_accessMode()
      else
         self:addDefaultConstructor( firstToken.pos, classTypeInfo, classScope, node:get_memberList(), methodNameSet, false )
      end
   end
   
   
   for __index, advertiseInfo in ipairs( node:get_advertiseList() ) do
      local memberType = advertiseInfo:get_member():get_expType()
      do
         local _switchExp = memberType:get_kind()
         if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
            for __index, mtdName in ipairs( Ast.getAllMethodName( memberType, Ast.MethodKind.Object ):get_list() ) do
               local scope = _lune.unwrap( memberType:get_scope())
               local child = _lune.unwrap( scope:getTypeInfoField( mtdName, true, scope, Ast.ScopeAccess.Normal ))
               if child:get_accessMode() ~= Ast.AccessMode.Pri then
                  local childName = advertiseInfo:get_prefix() .. child:getTxt(  )
                  if not _lune._Set_has(methodNameSet, childName ) then
                     local impMtdType = self.processInfo:createAdvertiseMethodFrom( classTypeInfo, child )
                     classScope:addMethod( self.processInfo, advertiseInfo:get_pos(), impMtdType, child:get_accessMode(), child:get_staticFlag(), false )
                  end
                  
               end
               
            end
            
         else 
            
               self:error( string.format( "advertise member type is illegal -- %s", advertiseInfo:get_member():get_name()) )
         end
      end
      
   end
   
   
   if classTypeInfo:isInheritFrom( self.processInfo, Ast.builtinTypeMapping, nil ) then
      local checkedTypeMap = {}
      for __index, memberNode in ipairs( node:get_memberList() ) do
         local memberType = memberNode:get_expType()
         if not Ast.NormalTypeInfo.isAvailableMapping( self.processInfo, memberType, checkedTypeMap ) then
            self:addErrMess( memberNode:get_pos(), string.format( "member type is not Mapping -- %s", memberType:getTxt(  )) )
         elseif memberType:get_kind() == Ast.TypeInfoKind.IF then
            self:addErrMess( memberNode:get_pos(), string.format( "Mapping class has not the interface type member. -- %s", memberNode:get_name().txt) )
         elseif memberType:get_abstractFlag() then
            self:addErrMess( memberNode:get_pos(), string.format( "Mapping class has not the abstract class member. -- %s", memberNode:get_name().txt) )
         end
         
      end
      
      local fromMapFuncTypeInfo = self.processInfo:createFunc( false, false, nil, Ast.TypeInfoKind.Func, classTypeInfo, true, false, true, Ast.AccessMode.Pub, "_fromMap", Ast.Async.Async, nil, {mapType:get_nilableTypeInfo()}, {classTypeInfo:get_nilableTypeInfo(), Ast.builtinTypeString:get_nilableTypeInfo()}, true )
      classScope:addMethod( self.processInfo, nil, fromMapFuncTypeInfo, ctorAccessMode, true, false )
      local fromStemFuncTypeInfo = self.processInfo:createFunc( false, false, nil, Ast.TypeInfoKind.Func, classTypeInfo, true, false, true, Ast.AccessMode.Pub, "_fromStem", Ast.Async.Async, nil, {Ast.builtinTypeStem_}, {classTypeInfo:get_nilableTypeInfo(), Ast.builtinTypeString:get_nilableTypeInfo()}, true )
      classScope:addMethod( self.processInfo, nil, fromStemFuncTypeInfo, ctorAccessMode, true, false )
   end
   
   
   if classTypeInfo:isInheritFrom( self.processInfo, Ast.builtinTypeAsyncItem, nil ) then
      if not classTypeInfo:isInheritFrom( self.processInfo, Ast.builtinTypeMapping, nil ) then
         self:addErrMess( firstToken.pos, "__AsyncItem implemented class must inherit Mapping." )
      end
      
      local pipeType = self.processInfo:createGeneric( builtinFunc.__pipe_, {classTypeInfo}, self.moduleType )
      local createPipeFuncTypeInfo = self.processInfo:createFunc( false, false, nil, Ast.TypeInfoKind.Func, classTypeInfo, true, false, true, Ast.AccessMode.Pub, "_createPipe", Ast.Async.Async, nil, {Ast.builtinTypeInt}, {pipeType:get_nilableTypeInfo()}, true )
      classScope:addMethod( self.processInfo, nil, createPipeFuncTypeInfo, Ast.AccessMode.Pub, true, false )
   end
   
   
   self:popClass(  )
   
   return node
end


function TransUnit:addMethod( classTypeInfo, methodNode, name )

   local classNodeInfo = _lune.unwrap( self.typeInfo2ClassNode[classTypeInfo])
   
   classNodeInfo:get_outerMethodSet()[name]= true
   table.insert( classNodeInfo:get_fieldList(), methodNode )
end


function TransUnit:processAddFunc( isFunc, parentScope, name, typeInfo, alt2typeMap )

   local accessMode = typeInfo:get_accessMode()
   if accessMode == Ast.AccessMode.Global then
      parentScope = self.globalScope
   end
   
   
   local hasPrototype
   
   do
      local prottype = parentScope:getTypeInfoChild( typeInfo:get_rawTxt() )
      if prottype ~= nil then
         local argTypeList = typeInfo:get_argTypeInfoList()
         local retTypeInfoList = typeInfo:get_retTypeInfoList()
         local matched = true
         do
            local matchFlag, err = Ast.TypeInfo.checkMatchType( self.processInfo, prottype:get_argTypeInfoList(), argTypeList, false, nil, alt2typeMap )
            if matchFlag ~= Ast.MatchType.Match then
               self:addErrMess( name.pos, "mismatch functype param: " .. err )
               matched = false
            end
            
         end
         
         do
            local matchFlag, err = Ast.TypeInfo.checkMatchType( self.processInfo, prottype:get_retTypeInfoList(), retTypeInfoList, false, nil, alt2typeMap )
            if matchFlag ~= Ast.MatchType.Match then
               self:addErrMess( name.pos, "mismatch functype ret: " .. err )
               matched = false
            end
            
         end
         
         do
            local matchFlag, err = typeInfo:canEvalWith( self.processInfo, prottype, Ast.CanEvalType.SetOp, alt2typeMap )
            if not matchFlag then
               if err then
                  self:addErrMess( name.pos, string.format( "mismatch functype -- %s", err) )
               else
                
                  self:addErrMess( name.pos, string.format( "mismatch functype -- %s / %s", typeInfo:get_display_stirng(), prottype:get_display_stirng()) )
               end
               
               matched = false
            end
            
         end
         
         do
            if prottype:get_asyncMode() ~= typeInfo:get_asyncMode() then
               self:addErrMess( name.pos, string.format( "mismatch async -- %s / %s", Ast.Async:_getTxt( prottype:get_asyncMode())
               , Ast.Async:_getTxt( typeInfo:get_asyncMode())
               ) )
               matched = false
            end
            
         end
         
         
         if self.protoFuncMap[prottype] then
            hasPrototype = true
            self.protoFuncMap[prottype] = nil
         else
          
            hasPrototype = false
            if not prottype:get_autoFlag() then
               self:addErrMess( name.pos, string.format( "multiple define -- %s", name.txt) )
            end
            
         end
         
         if matched then
            (_lune.unwrap( _lune.__Cast( prottype, 3, Ast.NormalTypeInfo )) ):switchScopeTo( _lune.unwrap( typeInfo:get_scope()) )
            typeInfo = prottype
         end
         
      else
         hasPrototype = false
      end
   end
   
   
   if typeInfo:get_kind() == Ast.TypeInfoKind.Method and typeInfo:get_accessMode() ~= Ast.AccessMode.Pri then
      local classType = typeInfo:get_parentInfo()
      if _lune._Set_has(self.advertisedTypeSet, classType ) and not hasPrototype then
         self:addErrMess( name.pos, string.format( "This class(%s) is used by advertise. You must declare the prototype of this method.", classType:getTxt(  )) )
      end
      
   end
   
   
   local staticFlag = typeInfo:get_staticFlag()
   local mutable = Ast.TypeInfo.isMut( typeInfo )
   
   local funcSym, shadowing
   
   if isFunc then
      funcSym, shadowing = parentScope:addFunc( self.processInfo, name.pos, typeInfo, accessMode, staticFlag, mutable )
      self:errorShadowing( name.pos, shadowing )
   else
    
      funcSym, shadowing = parentScope:addMethod( self.processInfo, name.pos, typeInfo, accessMode, staticFlag, mutable )
   end
   
   return _lune.unwrap( (funcSym or shadowing ))
end


local CantOverrideMethods = {["__init"] = true, ["__free"] = true}

function TransUnit:analyzeDeclFunc( declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, classTypeInfo, firstToken, name )

   local token = self:getToken(  )
   do
      local _exp = name
      if _exp ~= nil then
         if _exp.txt ~= "__main" then
            name = self:checkSymbol( _exp, SymbolMode.MustNot_ )
         end
         
         if declFuncMode == DeclFuncMode.Func and _exp.txt == "main" then
            self:addWarnMess( _exp.pos, "LuneScript's main function is __main." )
         end
         
      else
         if token.txt ~= "(" then
            if token.txt ~= "__main" then
               name = self:checkSymbol( token, SymbolMode.MustNot_ )
            else
             
               name = token
            end
            
            token = self:getToken(  )
         end
         
      end
   end
   
   if not name and (Ast.isPubToExternal( accessMode ) or abstractFlag or overrideFlag or staticFlag ) then
      self:addErrMess( firstToken.pos, "The anonymous function must be local." )
   end
   
   local needPopFlag = false
   if token.txt == "." then
      needPopFlag = true
      
      if name ~= nil then
         local className = name.txt
         
         classTypeInfo = self.scope:getTypeInfoChild( className )
         
         if classTypeInfo ~= nil then
            self:pushClassScope( name.pos, classTypeInfo )
         else
            self:error( string.format( "not found class -- %s", className) )
         end
         
      else
         self:error( "can't use '.' for any function name" )
      end
      
      
      name = self:getSymbolToken( SymbolMode.MustNot_ )
      token = self:getToken(  )
      
   end
   
   
   local isCtorFlag = false
   local kind = Nodes.NodeKind.get_DeclConstr()
   local typeKind = Ast.TypeInfoKind.Func
   if classTypeInfo ~= nil then
      if not staticFlag then
         typeKind = Ast.TypeInfoKind.Method
      end
      
      do
         local _switchExp = (_lune.unwrap( name) ).txt
         if _switchExp == "__init" then
            isCtorFlag = true
            kind = Nodes.NodeKind.get_DeclConstr()
            for __index, symbolInfo in pairs( self.scope:get_symbol2SymbolInfoMap() ) do
               if not symbolInfo:get_staticFlag() then
                  symbolInfo:clearValue(  )
               end
               
            end
            
         elseif _switchExp == "__free" then
            kind = Nodes.NodeKind.get_DeclDestr()
            if not self.targetLuaVer:get_canUseMetaGc() then
               self:addErrMess( firstToken.pos, "this lua version is not support __free." )
            end
            
         else 
            
               kind = Nodes.NodeKind.get_DeclMethod()
         end
      end
      
   else
      kind = Nodes.NodeKind.get_DeclFunc()
      if not staticFlag then
         staticFlag = true
      end
      
   end
   
   
   local orgStaticFlag = staticFlag
   if declFuncMode == DeclFuncMode.Module then
      staticFlag = true
   end
   
   
   local parentScope = self.scope
   local funcBodyScope = self:pushScope( false )
   
   local altTypeList = {}
   if token.txt == "<" then
      token, altTypeList = self:analyzeDeclAlternateType( false, token, accessMode )
      for __index, altType in ipairs( altTypeList ) do
         funcBodyScope:addAlternate( self.processInfo, accessMode, altType:get_rawTxt(), firstToken.pos, altType )
      end
      
   end
   
   
   self:checkToken( token, "(" )
   
   local parentPub
   
   if classTypeInfo ~= nil then
      parentPub = Ast.isPubToExternal( classTypeInfo:get_accessMode() )
   else
      parentPub = Ast.isPubToExternal( accessMode )
   end
   
   
   local argList = {}
   token = self:analyzeDeclArgList( accessMode, funcBodyScope, argList, parentPub )
   
   self:checkToken( token, ")" )
   token = self:getToken(  )
   
   local mutable
   
   local asyncMode
   
   
   token, mutable, asyncMode = self:getMutableAsync( token )
   
   local pubToExtFlag = Ast.isPubToExternal( accessMode )
   
   local argTypeList = {}
   for __index, argNode in ipairs( argList ) do
      table.insert( argTypeList, argNode:get_expType() )
   end
   
   
   local alt2typeMap = Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
   if classTypeInfo ~= nil then
      alt2typeMap = classTypeInfo:createAlt2typeMap( false )
      
      if kind == Nodes.NodeKind.get_DeclMethod() or kind == Nodes.NodeKind.get_DeclConstr() or kind == Nodes.NodeKind.get_DeclDestr() then
         local workClass = classTypeInfo
         if kind == Nodes.NodeKind.get_DeclConstr() or kind == Nodes.NodeKind.get_DeclDestr() then
            mutable = true
         end
         
         
         if not Ast.isPubToExternal( workClass:get_accessMode() ) then
            pubToExtFlag = false
         end
         
         
         if Ast.TypeInfo.isMut( workClass ) and not mutable then
            workClass = self:createModifier( workClass, Ast.MutMode.IMut )
         end
         
         if not staticFlag then
            self.scope:add( self.processInfo, Ast.SymbolKind.Arg, false, true, "self", nil, workClass, Ast.AccessMode.Pri, false, mutable and Ast.MutMode.Mut or Ast.MutMode.IMut, true, false )
         end
         
         
         if not workClass:get_abstractFlag() and abstractFlag then
            self:addErrMess( firstToken.pos, "no abstract class does not have abstract method" )
         end
         
      end
      
      
      if classTypeInfo:isInheritFrom( self.processInfo, Ast.builtinTypeRunner ) then
         for index, argNode in ipairs( argList ) do
            if Ast.isMutableType( argNode:get_expType() ) then
               self:addErrMess( argNode:get_pos(), string.format( "__Runner can't have the mutable argument. -- %d: %s", index, argNode:get_expType():getTxt(  )) )
            end
            
         end
         
      end
      
   end
   
   
   local retTypeInfoList
   
   local retTypeNodeList
   
   retTypeInfoList, token, retTypeNodeList = self:analyzeRetTypeList( pubToExtFlag, accessMode, token, parentPub )
   
   local namespaceInfo = self:getCurrentNamespaceTypeInfo(  )
   
   local funcName = ""
   if name ~= nil then
      funcName = name.txt
      
      if kind == Nodes.NodeKind.get_DeclFunc() then
         do
            local _switchExp = accessMode
            if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
               if parentScope ~= self.moduleScope then
                  self:addErrMess( firstToken.pos, "'global' or 'pub' function must exist top scope." )
               end
               
            end
         end
         
      end
      
   end
   
   
   local typeInfo
   
   local funcSym
   
   do
      local workTypeInfo = self.processInfo:createFunc( abstractFlag, false, funcBodyScope, typeKind, namespaceInfo, false, false, staticFlag, accessMode, funcName, self:getDefaultAsync( typeKind, classTypeInfo or self:getCurrentClass(  ), asyncMode ), altTypeList, argTypeList, retTypeInfoList, mutable )
      
      if name ~= nil then
         local workSym = self:processAddFunc( kind == Nodes.NodeKind.get_DeclFunc(), funcBodyScope:get_parent(), name, workTypeInfo, alt2typeMap )
         typeInfo = workSym:get_typeInfo()
         funcSym = workSym
         
         if name.txt == "__main" then
            if #typeInfo:get_argTypeInfoList() ~= 1 or typeInfo:get_argTypeInfoList()[1]:get_kind() ~= Ast.TypeInfoKind.List or typeInfo:get_argTypeInfoList()[1]:get_itemTypeInfoList()[1] ~= Ast.builtinTypeString or #typeInfo:get_retTypeInfoList() ~= 1 or typeInfo:get_retTypeInfoList()[1] ~= Ast.builtinTypeInt then
               local mess = string.format( "'__main' function's type has to be __main( argList:List<str> ) : int -- %s", typeInfo:get_display_stirng())
               self:addErrMess( name.pos, mess )
            end
            
         end
         
      else
         typeInfo = workTypeInfo
         funcSym = nil
      end
      
   end
   
   
   if overrideFlag then
      if not name then
         self:addErrMess( firstToken.pos, "can't override anonymous func" )
      end
      
      
      if _lune._Set_has(CantOverrideMethods, funcName ) then
         self:addErrMess( firstToken.pos, string.format( "This method can't override -- %s", funcName) )
      end
      
      do
         local overrideType = self.scope:get_parent():getTypeInfoField( funcName, false, funcBodyScope, self.scopeAccess )
         if overrideType ~= nil then
            for __index, err in ipairs( self:checkOverrideMethod( overrideType, typeInfo ) ) do
               self:addErrMess( firstToken.pos, err )
            end
            
         else
            self:addErrMess( firstToken.pos, "not found override -- " .. funcName )
         end
      end
      
   else
    
      if name ~= nil then
         if not _lune._Set_has(CantOverrideMethods, name.txt ) then
            if self.scope:get_parent():getTypeInfoField( name.txt, false, funcBodyScope, Ast.ScopeAccess.Full ) then
               self:addErrMess( firstToken.pos, "mismatch override --" .. funcName )
            else
             
               do
                  local ifFunc = self.scope:get_parent():getSymbolInfoIfField( name.txt, funcBodyScope, Ast.ScopeAccess.Full )
                  if ifFunc ~= nil then
                     if not ifFunc:get_typeInfo():canEvalWith( self.processInfo, typeInfo, Ast.CanEvalType.SetEq, alt2typeMap ) then
                        self:addErrMess( firstToken.pos, string.format( "mismatch method type -- %s", funcName) )
                     end
                     
                  end
               end
               
            end
            
         end
         
      end
      
   end
   
   
   local node = self:createNoneNode( firstToken.pos )
   local body = nil
   if token.txt == ";" then
      if declFuncMode == DeclFuncMode.Module or declFuncMode == DeclFuncMode.Glue then
         
      else
       
         if not abstractFlag then
            self.protoFuncMap[typeInfo] = firstToken.pos
         end
         
         if _lune.nilacc( classTypeInfo, 'get_kind', 'callmtd' ) == Ast.TypeInfoKind.IF then
            
         else
          
            if kind == Nodes.NodeKind.get_DeclMethod() then
               kind = Nodes.NodeKind.get_ProtoMethod()
            end
            
         end
         
      end
      
   else
    
      if abstractFlag then
         self:addErrMess( token.pos, "abstract method can't have body." )
      end
      
      
      self:pushback(  )
      
      local analyzingState
      
      if isCtorFlag then
         analyzingState = AnalyzingState.Constructor
      elseif staticFlag and classTypeInfo then
         analyzingState = AnalyzingState.ClassMethod
      else
       
         analyzingState = AnalyzingState.Func
      end
      
      
      funcBodyScope:addLocalVar( self.processInfo, false, false, "__func__", firstToken.pos, Ast.builtinTypeString, Ast.MutMode.IMut )
      
      local workBody = self:analyzeFuncBlock( analyzingState, firstToken, classTypeInfo, typeInfo, funcName, funcBodyScope, typeInfo:get_retTypeInfoList() )
      body = workBody
      
      if isCtorFlag then
         if classTypeInfo ~= nil then
            if classTypeInfo:get_baseTypeInfo() ~= Ast.headTypeInfo then
               local needCall = true
               for __index, stmt in ipairs( workBody:get_stmtList() ) do
                  do
                     local _switchExp = stmt:get_kind()
                     if _switchExp == Nodes.nodeKindEnum.ExpCallSuperCtor then
                        needCall = false
                     elseif _switchExp == Nodes.nodeKindEnum.BlankLine then
                     else 
                        
                           break
                     end
                  end
                  
               end
               
               if needCall then
                  self:addErrMess( workBody:get_pos(), "__init must call super() with first." )
               end
               
            end
            
         end
         
      end
      
   end
   
   
   local function createDeclFuncInfo( funcKind )
   
      local classDeclNode
      
      if classTypeInfo ~= nil then
         classDeclNode = self.typeInfo2ClassNode[classTypeInfo]
      else
         classDeclNode = nil
      end
      
      
      return Nodes.DeclFuncInfo.new(funcKind, classTypeInfo, classDeclNode, name, funcSym, argList, orgStaticFlag, accessMode, asyncMode, body, retTypeInfoList, retTypeNodeList, _lune._Set_has(self.has__func__Symbol, typeInfo ), overrideFlag)
   end
   
   do
      local _switchExp = (kind )
      if _switchExp == Nodes.NodeKind.get_DeclConstr() then
         local info = createDeclFuncInfo( Nodes.FuncKind.Ctor )
         node = Nodes.DeclConstrNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, info )
      elseif _switchExp == Nodes.NodeKind.get_DeclDestr() then
         local info = createDeclFuncInfo( Nodes.FuncKind.Dstr )
         node = Nodes.DeclDestrNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, info )
      elseif _switchExp == Nodes.NodeKind.get_DeclMethod() then
         local info = createDeclFuncInfo( Nodes.FuncKind.Mtd )
         node = Nodes.DeclMethodNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, info )
      elseif _switchExp == Nodes.NodeKind.get_ProtoMethod() then
         local info = createDeclFuncInfo( Nodes.FuncKind.Mtd )
         node = Nodes.ProtoMethodNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, info )
      elseif _switchExp == Nodes.NodeKind.get_DeclFunc() then
         local info = createDeclFuncInfo( Nodes.FuncKind.Func )
         node = Nodes.DeclFuncNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, info )
      else 
         
            self:error( string.format( "illegal kind -- %d", kind) )
      end
   end
   
   self.has__func__Symbol[typeInfo]= nil
   
   self:popScope(  )
   
   if needPopFlag then
      self:addMethod( _lune.unwrap( classTypeInfo), node, funcName )
      self:popClass(  )
   end
   
   
   return node
end


function TransUnit:createExpListNode( orgExpList, newExpList )

   local newExpTypeList = {}
   for __index, expNode in ipairs( newExpList ) do
      table.insert( newExpTypeList, expNode:get_expType() )
   end
   
   if #newExpList[#newExpList]:get_expTypeList() > 1 then
      self:addErrMess( orgExpList:get_pos(), string.format( "illegal exp -- %d", #newExpList[#newExpList]:get_expTypeList()) )
   end
   
   do
      local mRetIndex = _lune.nilacc( orgExpList:get_mRetExp(), 'get_index', 'callmtd' )
      if mRetIndex ~= nil then
         if mRetIndex > #newExpList then
            self:addErrMess( orgExpList:get_pos(), string.format( "over index -- %d", mRetIndex) )
         end
         
      end
   end
   
   return Nodes.ExpListNode.create( self.nodeManager, orgExpList:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), newExpTypeList, newExpList, orgExpList:get_mRetExp(), orgExpList:get_followOn() )
end


function TransUnit:checkLiteralEmptyCollection( pos, symbolName, expType )

   for __index, itemType in ipairs( expType:get_itemTypeInfoList() ) do
      if itemType == Ast.builtinTypeNone then
         self:addErrMess( pos, string.format( "must set the item type of Collection. -- %s:%s", symbolName, expType:get_srcTypeInfo():getTxt( self.typeNameCtrl )) )
         break
      end
      
   end
   
end


function TransUnit:accessSymbol( symbolInfo, canLeftExp )

   if symbolInfo:get_kind() == Ast.SymbolKind.Fun then
      self.scope:accessSymbol( self.moduleScope, symbolInfo )
      symbolInfo:set_posForModToRef( symbolInfo:get_posForLatestMod() )
      
      if self.scope:isClosureAccess( self.moduleScope, symbolInfo ) then
         table.insert( self.closureFunList, ClosureFun.new(symbolInfo, self.scope) )
      end
      
      do
         local scope = symbolInfo:get_typeInfo():get_scope()
         if scope ~= nil then
            for __index, symInfo in ipairs( scope:get_closureSymList() ) do
               symInfo:set_posForModToRef( symInfo:get_posForLatestMod() )
            end
            
         end
      end
      
   else
    
      self.scope:accessSymbol( self.moduleScope, symbolInfo )
      if canLeftExp then
         self.accessSymbolSetQueue:add( symbolInfo )
      else
       
         symbolInfo:set_posForModToRef( symbolInfo:get_posForLatestMod() )
      end
      
   end
   
end


local LetVarInfo = {}
function LetVarInfo.setmeta( obj )
  setmetatable( obj, { __index = LetVarInfo  } )
end
function LetVarInfo.new( mutable, varName, varType )
   local obj = {}
   LetVarInfo.setmeta( obj )
   if obj.__init then
      obj:__init( mutable, varName, varType )
   end
   return obj
end
function LetVarInfo:__init( mutable, varName, varType )

   self.mutable = mutable
   self.varName = varName
   self.varType = varType
end


function TransUnit:analyzeInitExp( firstPos, accessMode, unwrapFlag, letVarList, typeInfoList )

   local expList = nil
   local expectTypeList = {}
   for __index, varInfo in ipairs( letVarList ) do
      table.insert( expectTypeList, _lune.unwrapDefault( _lune.nilacc( varInfo.varType, 'get_expType', 'callmtd' ), Ast.builtinTypeNone) )
   end
   
   
   expList = self:analyzeExpList( false, false, false, nil, expectTypeList )
   if not expList then
      self:error( "expList is nil" )
   end
   
   local orgExpTypeList = {}
   if expList ~= nil then
      
      if unwrapFlag then
         local hasNilable = false
         for index, _2934 in ipairs( letVarList ) do
            if expList:getExpTypeAt( index ):get_nilable() then
               hasNilable = true
               break
            end
            
         end
         
         if not hasNilable then
            self:addWarnMess( firstPos, "has no nilable" )
         end
         
      end
      
      
      local workList = expList
      local updateExpList = false
      local newExpList = {}
      for index, exp in ipairs( workList:get_expList() ) do
         table.insert( newExpList, exp )
         if not exp:canBeRight( self.processInfo ) then
            self:addErrMess( exp:get_pos(), string.format( "this node(%d) can not be r-value. -- %s", index, Nodes.getNodeKindName( exp:get_kind() )) )
         end
         
      end
      
      local expTypeList = {}
      for index, expType in ipairs( workList:get_expTypeList() ) do
         if index == #workList:get_expTypeList() and expType:get_kind() == Ast.TypeInfoKind.DDD then
            local dddItemType = Ast.builtinTypeStem_
            if #expType:get_itemTypeInfoList() > 0 then
               dddItemType = expType:get_itemTypeInfoList()[1]:get_nilableTypeInfo()
            end
            
            for subIndex = index, #letVarList do
               local argType = typeInfoList[subIndex]
               local checkType = dddItemType
               if unwrapFlag then
                  checkType = dddItemType:get_nonnilableType()
               end
               
               if not argType:equals( self.processInfo, Ast.builtinTypeEmpty ) and not argType:canEvalWith( self.processInfo, checkType, Ast.CanEvalType.SetOp, {} ) then
                  self:addErrMess( firstPos, string.format( "unmatch value type (index = %d) %s <- %s", subIndex, argType:getTxt( self.typeNameCtrl ), dddItemType:getTxt(  )) )
               end
               
               table.insert( expTypeList, checkType )
               table.insert( orgExpTypeList, dddItemType )
            end
            
         else
          
            local expTypeInfo = expType
            if expType:get_kind() == Ast.TypeInfoKind.DDD then
               local itemList = expType:get_itemTypeInfoList()
               if #itemList > 0 then
                  expTypeInfo = itemList[1]
               else
                
                  expTypeInfo = Ast.builtinTypeStem_
               end
               
            end
            
            table.insert( orgExpTypeList, expTypeInfo )
            if expTypeInfo == Ast.builtinTypeNil and index <= #typeInfoList then
               orgExpTypeList[index] = typeInfoList[index]:get_nilableTypeInfo()
            end
            
            if unwrapFlag and expTypeInfo:get_nilable() then
               expTypeInfo = expTypeInfo:get_nonnilableType()
            end
            
            
            if index <= #typeInfoList then
               local varType = typeInfoList[index]
               local alt2typeMap = Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
               if varType:get_kind() == Ast.TypeInfoKind.Box then
                  alt2typeMap = varType:createAlt2typeMap( true )
               end
               
               Ast.CanEvalCtrlTypeInfo.setupNeedAutoBoxing( alt2typeMap, self.processInfo )
               
               if not varType:equals( self.processInfo, Ast.builtinTypeEmpty ) and not (unwrapFlag and expTypeInfo:equals( self.processInfo, Ast.builtinTypeNil ) ) then
                  local canEval, mess = varType:canEvalWith( self.processInfo, expTypeInfo, Ast.CanEvalType.SetOp, alt2typeMap )
                  if not canEval then
                     self:addErrMess( firstPos, string.format( "unmatch value type (index:%d) %s <- %s%s", index, varType:getTxt( self.typeNameCtrl ), expTypeInfo:getTxt( self.typeNameCtrl ), mess and string.format( " -- %s", mess) or "") )
                  end
                  
               end
               
               if varType:get_kind() == Ast.TypeInfoKind.Box then
                  typeInfoList[index] = self.processInfo:createBox( accessMode, expTypeInfo )
               end
               
               if Ast.CanEvalCtrlTypeInfo.canAutoBoxing( varType, expTypeInfo ) then
                  updateExpList = true
                  local exp = newExpList[index]
                  newExpList[index] = Nodes.BoxingNode.create( self.nodeManager, exp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {varType}, exp )
                  if not Ast.CanEvalCtrlTypeInfo.finishNeedAutoBoxing( alt2typeMap, 1 ) then
                     self:addErrMess( exp:get_pos(), string.format( "auto boxing error %s <- %s", varType:getTxt(  ), expTypeInfo:getTxt(  )) )
                  end
                  
               else
                
                  if not Ast.CanEvalCtrlTypeInfo.finishNeedAutoBoxing( alt2typeMap, 0 ) then
                     self:addErrMess( newExpList[index]:get_pos(), string.format( "illegal auto boxing error %s <- %s", varType:getTxt(  ), expTypeInfo:getTxt(  )) )
                  end
                  
               end
               
            end
            
            table.insert( expTypeList, expTypeInfo )
         end
         
      end
      
      if updateExpList then
         workList = self:createExpListNode( workList, newExpList )
      end
      
      
      do
         local alt2typeMap = Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
         
         do
            local _exp = self:checkImplicitCast( alt2typeMap, true, typeInfoList, workList, function ( dstType, expNode )
            
               return nil
            end )
            if _exp ~= nil then
               workList = _exp
            end
         end
         
      end
      
      
      for index, varType in ipairs( typeInfoList ) do
         if index > #expTypeList then
            if not varType:get_nilable() then
               self:addErrMess( firstPos, string.format( "unmatch value type (index:%d) %s <- nil", index, varType:getTxt( self.typeNameCtrl )) )
            end
            
         end
         
      end
      
      for index, typeInfo in ipairs( expTypeList ) do
         if #typeInfoList < index or typeInfoList[index]:equals( self.processInfo, Ast.builtinTypeEmpty ) then
            local workPos
            
            local workType
            
            local workName
            
            if index <= #letVarList then
               local letVar = letVarList[index]
               workPos = letVar.varName.pos
               workName = letVar.varName.txt
               if Ast.TypeInfo.isMut( typeInfo ) and not Ast.isMutable( letVar.mutable ) then
                  workType = self:createModifier( typeInfo, Ast.MutMode.IMutRe )
               else
                
                  workType = typeInfo
               end
               
               
               if self.analyzeMode == AnalyzeMode.Inquire and self:isTargetToken( letVar.varName ) then
                  self:dumpSymbolType( accessMode, letVar.varName.txt, workType )
               end
               
            else
             
               workType = typeInfo
               workPos = firstPos
               workName = ""
            end
            
            typeInfoList[index] = workType
            
            do
               local _switchExp = workType:get_kind()
               if _switchExp == Ast.TypeInfoKind.Func then
                  if #expTypeList ~= 1 or workType:get_rawTxt() ~= "" then
                     self:addErrMess( firstPos, string.format( "must set the type of variable for function. -- %s", workName) )
                  end
                  
               elseif _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Map then
                  self:checkLiteralEmptyCollection( workPos, workName, workType )
               end
            end
            
         end
         
      end
      
      return typeInfoList, letVarList, orgExpTypeList, workList
   end
   
   
   return typeInfoList, letVarList, orgExpTypeList, nil
end

function TransUnit:analyzeLetAndInitExp( firstPos, letFlag, initMutable, accessMode, unwrapFlag )

   local typeInfoList = {}
   local letVarList = {}
   
   local nextToken = Parser.getEofToken(  )
   
   if letFlag then
      repeat 
         local mutable = initMutable
         nextToken = self:getToken(  )
         if nextToken.txt == "mut" then
            mutable = Ast.MutMode.Mut
            nextToken = self:getToken(  )
         end
         
         local varName = self:checkSymbol( nextToken, SymbolMode.MustNot_Or_ )
         nextToken = self:getToken(  )
         local typeInfo = Ast.builtinTypeEmpty
         if nextToken.txt == ":" then
            local refType = self:analyzeRefType( accessMode, false, Ast.isPubToExternal( accessMode ) )
            table.insert( letVarList, LetVarInfo.new(mutable, varName, refType) )
            typeInfo = refType:get_expType()
            nextToken = self:getToken(  )
         else
          
            table.insert( letVarList, LetVarInfo.new(Ast.isMutable( mutable ) and mutable or Ast.MutMode.IMutRe, varName, nil) )
         end
         
         if not typeInfo:equals( self.processInfo, Ast.builtinTypeEmpty ) and Ast.TypeInfo.isMut( typeInfo ) and not Ast.isMutable( mutable ) and self.ctrl_info.legacyMutableControl then
            typeInfo = self:createModifier( typeInfo, Ast.MutMode.IMutRe )
         end
         
         table.insert( typeInfoList, typeInfo )
      until nextToken.txt ~= ","
   else
    
      while true do
         local symbolToken = self:getToken(  )
         nextToken = self:getToken(  )
         
         local verSym = self.scope:getSymbolTypeInfo( symbolToken.txt, self.scope, self.moduleScope, self.scopeAccess )
         if verSym ~= nil then
            table.insert( letVarList, LetVarInfo.new(verSym:get_mutMode(), symbolToken, nil) )
            table.insert( typeInfoList, verSym:get_typeInfo() )
         else
            self:addErrMess( symbolToken.pos, string.format( "not found symbol -- %s", symbolToken.txt) )
         end
         
         if nextToken.txt ~= "," then
            break
         end
         
      end
      
   end
   
   
   if nextToken.txt ~= "=" then
      self:pushback(  )
      local orgExpTypeList = {}
      return typeInfoList, letVarList, orgExpTypeList, nil
   end
   
   return self:analyzeInitExp( firstPos, accessMode, unwrapFlag, letVarList, typeInfoList )
end


function TransUnit:analyzeDeclVar( mode, accessMode, firstToken )

   local unwrapFlag = false
   local token, continueFlag = self:getContinueToken(  )
   if continueFlag and token.txt == "!" then
      unwrapFlag = true
   else
    
      self:pushback(  )
      if mode ~= Nodes.DeclVarMode.Let then
         Util.log( "need '!'" )
      end
      
   end
   
   
   if accessMode == Ast.AccessMode.Pub then
      if self.scope ~= self.moduleScope then
         self:addErrMess( firstToken.pos, "'pub' variable must exist top scope." )
      end
      
   end
   
   
   local typeInfoList, letVarList, orgExpTypeList, expList = self:analyzeLetAndInitExp( firstToken.pos, mode == Nodes.DeclVarMode.Let, mode == Nodes.DeclVarMode.Sync and Ast.MutMode.Mut or Ast.MutMode.IMut, accessMode, unwrapFlag )
   
   if mode == Nodes.DeclVarMode.Let and #typeInfoList == 1 then
      if expList ~= nil then
         local typeInfo = typeInfoList[1]
         local letVaInfo = letVarList[1]
         if #expList:get_expList() == 1 and typeInfo:get_kind() == Ast.TypeInfoKind.Func then
            local valExp = expList:get_expList()[1]
            do
               local macroExp = _lune.__Cast( valExp, 3, Nodes.ExpMacroExpNode )
               if macroExp ~= nil then
                  if macroExp:get_expType():get_kind() == Ast.TypeInfoKind.Func and #macroExp:get_stmtList() == 1 then
                     valExp = macroExp:get_stmtList()[1]
                  end
                  
               end
            end
            
            do
               local declNode = _lune.__Cast( valExp, 3, Nodes.DeclFuncNode )
               if declNode ~= nil then
                  if not declNode:get_declInfo():get_name() then
                     if Ast.isMutable( letVaInfo.mutable ) then
                        self:addErrMess( letVaInfo.varName.pos, string.format( "Any function can't be mutable. -- %s", letVaInfo.varName.txt) )
                     end
                     
                     local letVarInfo = letVarList[1]
                     local newTypeInfo = self.processInfo:createFunc( typeInfo:get_abstractFlag(), false, typeInfo:get_scope(), typeInfo:get_kind(), typeInfo:get_parentInfo(), false, false, typeInfo:get_staticFlag(), accessMode, letVarInfo.varName.txt, typeInfo:get_asyncMode(), typeInfo:get_itemTypeInfoList(), typeInfo:get_argTypeInfoList(), typeInfo:get_retTypeInfoList(), Ast.TypeInfo.isMut( typeInfo ) )
                     local funcSym = self:processAddFunc( true, self.scope, letVarInfo.varName, newTypeInfo, Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false ) )
                     self.nodeManager:delNode( declNode )
                     
                     local declInfo = Nodes.DeclFuncInfo.createFrom( declNode:get_declInfo(), letVarInfo.varName, funcSym )
                     return Nodes.DeclFuncNode.create( self.nodeManager, declNode:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {newTypeInfo}, declInfo )
                  end
                  
               end
            end
            
         end
         
      end
      
   end
   
   
   local syncScope = self.scope
   if mode == Nodes.DeclVarMode.Sync then
      syncScope = self:pushScope( false )
   end
   
   
   local symbolInfoList = {}
   
   local varList = {}
   local syncSymbolList = {}
   for index, letVarInfo in ipairs( letVarList ) do
      local varName = letVarInfo.varName
      local typeInfo = typeInfoList[index]
      local varInfo = Nodes.VarInfo.new(varName, letVarInfo.varType, typeInfo)
      table.insert( varList, varInfo )
      
      if Ast.isPubToExternal( accessMode ) then
         self:checkPublic( varName.pos, typeInfo )
      end
      
      
      if not letVarInfo.varType and typeInfo:equals( self.processInfo, Ast.builtinTypeNil ) then
         self:addErrMess( varName.pos, string.format( 'need type -- %s', varName.txt) )
      end
      
      if mode == Nodes.DeclVarMode.Sync then
         do
            local symInfo = self.scope:getSymbolInfo( varName.txt, self.scope, true, self.scopeAccess )
            if symInfo ~= nil then
               table.insert( syncSymbolList, symInfo )
            end
         end
         
      end
      
      if mode == Nodes.DeclVarMode.Let or mode == Nodes.DeclVarMode.Sync then
         if mode == Nodes.DeclVarMode.Let then
            self:checkShadowing( varName.pos, varName.txt, self.scope )
            
         end
         
         local orgExpType = Ast.builtinTypeStem_
         if not unwrapFlag then
            orgExpType = Ast.builtinTypeEmpty
         end
         
         if index <= #orgExpTypeList then
            orgExpType = orgExpTypeList[index]
         end
         
         
         local hasValue = false
         if not unwrapFlag and orgExpType ~= Ast.builtinTypeEmpty or unwrapFlag and not orgExpType:get_nilable() then
            hasValue = true
         end
         
         self.scope:addVar( self.processInfo, accessMode, varName.txt, varName.pos, typeInfo, letVarInfo.mutable, hasValue )
         
         if typeInfo:get_asyncMode() == Ast.Async.Transient then
            self:addErrMess( varName.pos, string.format( "can't set the __trans type -- index:%d, %s", index, typeInfo:getTxt(  )) )
         end
         
      end
      
      table.insert( symbolInfoList, _lune.unwrap( self.scope:getSymbolInfo( varName.txt, self.scope, true, self.scopeAccess )) )
   end
   
   if mode ~= Nodes.DeclVarMode.Sync and self.macroScope then
      self.macroCtrl:registVar( symbolInfoList )
   end
   
   
   local unwrapBlock = nil
   local thenBlock = nil
   if unwrapFlag then
      local scope = self:pushScope( false )
      for index, letVarInfo in ipairs( letVarList ) do
         if letVarInfo.varName.txt ~= "_" then
            self:addLocalVar( letVarInfo.varName.pos, false, true, "_" .. letVarInfo.varName.txt, orgExpTypeList[index], Ast.MutMode.IMut )
         end
         
      end
      
      
      unwrapBlock = self:analyzeBlock( Nodes.BlockKind.LetUnwrap, TentativeMode.Start, scope, nil )
      self:popScope(  )
      if unwrapBlock ~= nil then
         do
            local _switchExp = mode
            if _switchExp == Nodes.DeclVarMode.Let or _switchExp == Nodes.DeclVarMode.Sync then
               local breakKind = unwrapBlock:getBreakKind( Nodes.CheckBreakMode.Normal )
               for __index, symbolInfo in ipairs( symbolInfoList ) do
                  if breakKind ~= Nodes.BreakKind.None then
                     self.tentativeSymbol:checkAndExclude( symbolInfo )
                     symbolInfo:updateValue( symbolInfo:get_pos() )
                  else
                   
                     if symbolInfo:get_name() ~= "_" and not self.tentativeSymbol:checkAndExclude( symbolInfo ) then
                        if not symbolInfo:get_hasValueFlag() then
                           self:addErrMess( unwrapBlock:get_pos(), "This variable isn't set -- " .. (symbolInfo:get_name() ) )
                        end
                        
                     end
                     
                  end
                  
               end
               
            elseif _switchExp == Nodes.DeclVarMode.Unwrap then
               for __index, symbolInfo in ipairs( symbolInfoList ) do
                  symbolInfo:updateValue( firstToken.pos )
               end
               
            end
         end
         
      end
      
      
      token = self:getToken( true )
      if token.txt == "then" then
         thenBlock = self:analyzeBlock( Nodes.BlockKind.LetUnwrapThenDo, TentativeMode.Finish, scope, nil )
      else
       
         self:pushback(  )
         self:finishTentativeSymbol( true )
      end
      
   end
   
   
   local syncBlock = nil
   if mode == Nodes.DeclVarMode.Sync then
      self:checkNextToken( "do" )
      syncBlock = self:analyzeBlock( Nodes.BlockKind.LetUnwrapThenDo, TentativeMode.Simple, syncScope, nil )
      self:popScope(  )
   end
   
   
   self:checkNextToken( ";" )
   
   local node = Nodes.DeclVarNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, mode, accessMode, false, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncSymbolList, syncBlock )
   
   return node
end

function TransUnit:analyzeIfUnwrap( firstToken )

   local nextToken = self:getToken(  )
   local typeInfoList = {}
   local varNameList = {}
   local expList
   
   local varList = {}
   
   local workTypeInfoList, letVarList, workExpList
   
   if nextToken.txt == "let" then
      local _
      workTypeInfoList, letVarList, _, workExpList = self:analyzeLetAndInitExp( firstToken.pos, true, Ast.MutMode.IMut, Ast.AccessMode.Local, true )
   else
    
      local _
      self:pushback(  )
      local tmpTypeInfoList = {Ast.builtinTypeEmpty}
      local tmpLetVarList = {LetVarInfo.new(Ast.MutMode.Mut, Parser.Token.new(Parser.TokenKind.Symb, "_exp", firstToken.pos, false), nil)}
      workTypeInfoList, letVarList, _, workExpList = self:analyzeInitExp( firstToken.pos, Ast.AccessMode.Local, true, tmpLetVarList, tmpTypeInfoList )
   end
   
   typeInfoList = workTypeInfoList
   if workExpList ~= nil then
      expList = workExpList
   else
      self:addErrMess( nextToken.pos, "if! let has illegal init val." )
      self:error( "system error" )
   end
   
   for __index, varInfo in ipairs( letVarList ) do
      table.insert( varNameList, varInfo.varName )
   end
   
   
   local scope = self:pushScope( false )
   
   for index, expType in ipairs( typeInfoList ) do
      if index > #varNameList then
         break
      end
      
      local varName = varNameList[index]
      table.insert( varList, self:addLocalVar( varName.pos, false, true, varName.txt, expType, Ast.MutMode.IMut ) )
   end
   
   
   local block = self:analyzeBlock( Nodes.BlockKind.IfUnwrap, TentativeMode.Start, scope, nil )
   
   self:popScope(  )
   
   local elseBlock = nil
   nextToken = self:getToken( true )
   if nextToken.txt == "else" then
      elseBlock = self:analyzeBlock( Nodes.BlockKind.Else, TentativeMode.Finish )
   else
    
      self:finishTentativeSymbol( false )
      self:pushback(  )
   end
   
   
   local hasCond = false
   for index, expNode in ipairs( expList:get_expList() ) do
      if index ~= #expList:get_expList() then
         if Ast.isConditionalbe( self.processInfo, expNode:get_expType() ) then
            hasCond = true
            break
         end
         
      else
       
         for __index, expType in ipairs( expNode:get_expTypeList() ) do
            if Ast.isConditionalbe( self.processInfo, expType ) then
               hasCond = true
               break
            end
            
         end
         
      end
      
   end
   
   if not hasCond then
      self:addErrMess( firstToken.pos, "This condition never be false" )
   end
   
   
   for __index, varSym in ipairs( varList ) do
      do
         local _switchExp = varSym:get_name()
         if _switchExp == "_" or _switchExp == "_exp" then
         else 
            
               if not varSym:get_posForModToRef() then
                  self:addWarnMess( _lune.unwrap( varSym:get_pos()), string.format( "This symbol has no referer -- %s", varSym:get_name()) )
               end
               
         end
      end
      
   end
   
   
   return Nodes.IfUnwrapNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, varList, expList, block, elseBlock )
end


function TransUnit:analyzeWhen( firstToken )

   local nextToken, continueFlag = self:getContinueToken(  )
   if not (continueFlag and nextToken.txt == "!" ) then
      self:pushback(  )
      self:addErrMess( nextToken.pos, "'when' need '!'" )
   end
   
   
   local symListNode = self:analyzeExpList( false, false, false )
   
   local scope = self:pushScope( false )
   local symPairList = {}
   
   for __index, expNode in ipairs( symListNode:get_expList() ) do
      do
         local refNode = _lune.__Cast( expNode, 3, Nodes.ExpRefNode )
         if refNode ~= nil then
            if expNode:get_expType():get_nilable() then
               local symbolInfo = refNode:get_symbolInfo()
               local newSymbolInfo = self:addLocalVar( firstToken.pos, false, expNode:canBeLeft(  ), refNode:get_symbolInfo():get_name(), expNode:get_expType():get_nonnilableType(), Ast.MutMode.IMut, true )
               table.insert( symPairList, Nodes.UnwrapSymbolPair.new(symbolInfo, newSymbolInfo) )
            else
             
               self:addErrMess( expNode:get_pos(), string.format( "This type isn't nilable. -- %s", expNode:get_expType():getTxt(  )) )
            end
            
         else
            self:addErrMess( expNode:get_pos(), "'when' support only local variables or arguments." )
         end
      end
      
   end
   
   
   local block = self:analyzeBlock( Nodes.BlockKind.When, TentativeMode.Start, scope, nil )
   
   self:popScope(  )
   
   local elseBlock = nil
   nextToken = self:getToken( true )
   if nextToken.txt == "else" then
      elseBlock = self:analyzeBlock( Nodes.BlockKind.Else, TentativeMode.Finish )
   else
    
      self:finishTentativeSymbol( false )
      self:pushback(  )
   end
   
   
   return Nodes.WhenNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, symPairList, block, elseBlock )
end

function TransUnit:MultiTo1( exp )

   if #exp:get_expTypeList() > 1 then
      exp = Nodes.ExpMultiTo1Node.create( self.nodeManager, exp:get_pos(), exp:get_macroArgFlag(), {exp:get_expType()}, exp )
   else
    
      do
         local dddType = _lune.__Cast( exp:get_expType(), 3, Ast.DDDTypeInfo )
         if dddType ~= nil then
            exp = Nodes.ExpMultiTo1Node.create( self.nodeManager, exp:get_pos(), exp:get_macroArgFlag(), {dddType:get_typeInfo():get_nilableTypeInfo()}, exp )
         end
      end
      
   end
   
   return exp
end

function TransUnit:analyzeExpOneRVal( allowNoneType, skipOp2Flag, opLevel, expectType )

   local exp = self:analyzeExp( allowNoneType, skipOp2Flag, false, opLevel, expectType )
   exp = self:MultiTo1( exp )
   
   if expectType ~= nil then
      do
         local _switchExp = expectType:get_kind()
         if _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.Class then
            local expOrgType = exp:get_expType():get_nonnilableType():get_srcTypeInfo()
            local exceptOrgType = expectType:get_nonnilableType():get_srcTypeInfo()
            if expOrgType:isInheritFrom( self.processInfo, exceptOrgType ) then
               exp = Nodes.ExpCastNode.create( self.nodeManager, exp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {expectType}, exp, expectType, Nodes.CastKind.Implicit )
            end
            
         end
      end
      
   end
   
   
   return exp
end

function TransUnit:createExpList( pos, expTypeList, expList, followOn, abbrNode )

   local workList = {}
   local mRetExp = nil
   if #expList > 0 then
      for index, exp in ipairs( expList ) do
         if expTypeList[index] ~= Ast.builtinTypeMultiExp and Nodes.hasMultiValNode( exp ) then
            if index ~= #expList then
               local firstType
               
               do
                  local dddType = _lune.__Cast( exp:get_expType(), 3, Ast.DDDTypeInfo )
                  if dddType ~= nil then
                     firstType = dddType:get_typeInfo():get_nilableTypeInfo()
                  else
                     firstType = exp:get_expType()
                  end
               end
               
               table.insert( workList, Nodes.ExpMultiTo1Node.create( self.nodeManager, exp:get_pos(), exp:get_macroArgFlag(), {firstType}, exp ) )
            else
             
               mRetExp = Nodes.MRetExp.new(exp, index)
               for listIndex, expType in ipairs( exp:get_expTypeList() ) do
                  if listIndex ~= 1 then
                     table.insert( expTypeList, expType )
                  end
                  
               end
               
               for retIndex, retType in ipairs( exp:get_expTypeList() ) do
                  if retIndex == 1 then
                     table.insert( workList, Nodes.ExpMRetNode.create( self.nodeManager, exp:get_pos(), exp:get_macroArgFlag(), {retType}, exp ) )
                  else
                   
                     table.insert( workList, Nodes.ExpAccessMRetNode.create( self.nodeManager, exp:get_pos(), exp:get_macroArgFlag(), {retType}, exp, retIndex ) )
                  end
                  
               end
               
            end
            
         else
          
            table.insert( workList, exp )
         end
         
      end
      
   end
   
   
   if abbrNode ~= nil then
      table.insert( workList, abbrNode )
   end
   
   
   return Nodes.ExpListNode.create( self.nodeManager, pos, self.macroCtrl:isInAnalyzeArgMode(  ), expTypeList, workList, mRetExp, followOn )
end

function TransUnit:analyzeExpList( allowNoneType, skipOp2Flag, canLeftExp, expNode, expectTypeList, contExpect )

   local expList = {}
   local pos = nil
   local expTypeList = {}
   if expNode ~= nil then
      pos = expNode:get_pos()
      table.insert( expList, expNode )
      table.insert( expTypeList, expNode:get_expType() )
   end
   
   
   local index = 1
   local abbrNode = nil
   local followOn = false
   repeat 
      local expectType = nil
      local allowNoneTypeOne = allowNoneType
      if expectTypeList ~= nil then
         if #expectTypeList > 0 then
            local checkIndex = index
            if index > #expectTypeList and contExpect then
               checkIndex = #expectTypeList
            end
            
            if checkIndex <= #expectTypeList and expectTypeList[checkIndex] ~= Ast.builtinTypeNone then
               local worktype = expectTypeList[checkIndex]
               expectType = worktype
               if worktype == Ast.builtinTypeExp then
                  allowNoneTypeOne = true
               end
               
            end
            
         end
         
      end
      
      
      local exp
      
      
      if self.macroCtrl:get_analyzeInfo():get_mode() == Nodes.MacroMode.AnalyzeArg and (expectType == Ast.builtinTypeExp or expectType == Ast.builtinTypeMultiExp ) then
         self.macroCtrl:switchMacroMode(  )
         exp = self:analyzeExp( allowNoneTypeOne, skipOp2Flag, canLeftExp, 0, expectType )
         self.macroCtrl:restoreMacroMode(  )
      else
       
         exp = self:analyzeExp( allowNoneTypeOne, skipOp2Flag, canLeftExp, 0, expectType )
      end
      
      
      if not allowNoneTypeOne and not exp:canBeRight( self.processInfo ) then
         self:addErrMess( exp:get_pos(), string.format( "This arg can't be r-value. -- %s", Nodes.getNodeKindName( exp:get_kind() )) )
      end
      
      if not pos then
         pos = exp:get_pos()
      end
      
      if expectType == Ast.builtinTypeExp or expectType == Ast.builtinTypeMultiExp then
         exp = Nodes.ExpMacroArgExpNode.create( self.nodeManager, exp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), exp:get_expTypeList(), Macro.nodeToCodeTxt( exp, self.moduleType ) )
         table.insert( expTypeList, _lune.unwrap( expectType) )
      else
       
         table.insert( expTypeList, exp:get_expType() )
      end
      
      table.insert( expList, exp )
      local token = self:getToken( true )
      
      if token.txt == "**" then
         if not Nodes.hasMultiValNode( exp ) then
            self:addErrMess( exp:get_pos(), string.format( "This arg(%d) doesn't have multiple value. It must not use '**'", index) )
         end
         
         followOn = true
         
         token = self:getToken( true )
      end
      
      
      if token.txt == "##" then
         if exp:get_expType():get_kind() == Ast.TypeInfoKind.DDD then
            self:addErrMess( token.pos, "'##' can't use with '...'" )
         end
         
         abbrNode = Nodes.AbbrNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeAbbr} )
         self:getToken(  )
         break
      end
      
      
      index = index + 1
      if self.macroCtrl:get_analyzeInfo():equalsArgTypeList( expectTypeList ) then
         self.macroCtrl:get_analyzeInfo():nextArg(  )
      end
      
   until token.txt ~= ","
   
   local expListNode = self:createExpList( _lune.unwrapDefault( pos, self:createPosition( 0, 0 )), expTypeList, expList, followOn, abbrNode )
   
   if not allowNoneType then
      for expIndex, expType in ipairs( expTypeList ) do
         if expType == Ast.builtinTypeNone then
            self:addErrMess( _lune.unwrapDefault( pos, self:createPosition( 0, 0 )), string.format( "The type of exp(%d) is None!!", expIndex) )
         end
         
      end
      
   end
   
   
   self:pushback(  )
   
   return expListNode
end


function TransUnit:checkSymbolHavingValue( pos, symbolInfoList )

   for __index, symbolInfo in ipairs( symbolInfoList ) do
      if symbolInfo:get_kind() == Ast.SymbolKind.Var then
         if not symbolInfo:get_hasValueFlag() then
            self:addErrMess( pos, string.format( "this variable have no value. -- %s", symbolInfo:get_name()) )
         end
         
      end
      
   end
   
end


function TransUnit:analyzeExpRefItem( token, exp, nilAccess )

   self:checkSymbolHavingValue( exp:get_pos(), exp:getSymbolInfo(  ) )
   
   local expType = exp:get_expType()
   
   expType = expType:get_extedType()
   if nilAccess then
      if not exp:get_expType():get_nilable() then
         self:addWarnMess( token.pos, string.format( "This is not nilable. -- %s", expType:getTxt(  )) )
         nilAccess = false
      else
       
         expType = expType:get_nonnilableType()
      end
      
   elseif exp:get_expType():get_nilable() then
      self:addErrMess( token.pos, string.format( "could not access with []. Use '$[]'-- %s", exp:get_expType():getTxt(  )) )
   end
   
   
   local expectItemType = nil
   local typeInfo = Ast.builtinTypeStem_
   local indexTypeInfo = Ast.builtinTypeInt
   if expType:get_kind() == Ast.TypeInfoKind.Map then
      local itemTypeList = expType:get_itemTypeInfoList(  )
      typeInfo = itemTypeList[2]
      indexTypeInfo = itemTypeList[1]
      expectItemType = itemTypeList[1]
      if not typeInfo:equals( self.processInfo, Ast.builtinTypeStem_ ) and not typeInfo:get_nilable() then
         typeInfo = typeInfo:get_nilableTypeInfo()
      end
      
   elseif expType:get_kind() == Ast.TypeInfoKind.Array or expType:get_kind() == Ast.TypeInfoKind.List then
      typeInfo = expType:get_itemTypeInfoList(  )[1]
   elseif expType:equals( self.processInfo, Ast.builtinTypeString ) then
      typeInfo = Ast.builtinTypeInt
   elseif expType:equals( self.processInfo, Ast.builtinTypeStem ) then
      indexTypeInfo = Ast.builtinTypeStem
      typeInfo = Ast.builtinTypeStem_
   else
    
      self:addErrMess( exp:get_pos(), string.format( "could not access with []. -- %s", exp:get_expType():getTxt(  )) )
   end
   
   
   if nilAccess then
      self.helperInfo.useNilAccess = true
      if not typeInfo:get_nilable() then
         typeInfo = typeInfo:get_nilableTypeInfo()
      end
      
   end
   
   
   if Ast.TypeInfo.isMut( typeInfo ) then
      if expType:get_mutMode() == Ast.MutMode.IMutRe then
         typeInfo = self:createModifier( typeInfo, Ast.MutMode.IMutRe )
      end
      
   end
   
   
   if Ast.isExtType( exp:get_expType():get_nonnilableType() ) then
      typeInfo = self:createExtType( exp:get_pos(), typeInfo )
   end
   
   
   local indexExp = self:analyzeExpOneRVal( false, false, nil, expectItemType )
   if not indexExp:canBeRight( self.processInfo ) then
      self:addErrMess( indexExp:get_pos(), "This node can't use index" )
   end
   
   if not indexTypeInfo:canEvalWith( self.processInfo, indexExp:get_expType(), Ast.CanEvalType.SetOp, {} ) then
      self:addErrMess( indexExp:get_pos(), string.format( "unmatch index type -- %s, %s", indexTypeInfo:getTxt(  ), indexExp:get_expType():getTxt(  )) )
   end
   
   
   self:checkNextToken( "]" )
   
   if expType:get_kind() == Ast.TypeInfoKind.Array or expType:get_kind() == Ast.TypeInfoKind.List then
      do
         local indexLit = indexExp:getLiteral(  )
         if indexLit ~= nil then
            do
               local _matchExp = indexLit
               if _matchExp[1] == Nodes.Literal.Int[1] then
                  local val = _matchExp[2][1]
               
                  if val <= 0 then
                     self:addWarnMess( indexExp:get_pos(), string.format( "index <= -1 (%d)", val) )
                  end
                  
               end
            end
            
         end
      end
      
   end
   
   
   local threading = false
   if nilAccess then
      threading = self:checkThreading( token.pos )
   end
   
   return Nodes.ExpRefItemNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, exp, nilAccess, threading, nil, indexExp )
end


function TransUnit:checkImplicitCast( alt2typeMap, validCastToGen, dstTypeList, expListNode, callback )

   local expNodeList = expListNode:get_expList()
   local newExpNodeList = {}
   local expTypeList = {}
   
   local hasModNode = false
   
   local function process( index, dstType, expNode, workNode )
   
      do
         local repNode = callback( dstType, expNode )
         if repNode ~= nil then
            if not hasModNode then
               hasModNode = true
            end
            
            workNode = repNode
         else
            if expNode:get_expType() ~= Ast.builtinTypeNil and dstType ~= Ast.builtinTypeEmpty and not dstType:get_srcTypeInfo():equals( self.processInfo, expNode:get_expType():get_srcTypeInfo() ) and not dstType:get_nonnilableType():get_srcTypeInfo():equals( self.processInfo, expNode:get_expType():get_srcTypeInfo() ) then
               if expNode:get_kind() ~= Nodes.NodeKind.get_Abbr() then
                  if not hasModNode then
                     hasModNode = true
                  end
                  
                  if dstType:get_kind() == Ast.TypeInfoKind.DDD then
                     local argList = {}
                     local argTypeList = {}
                     for workIndex = index, #expNodeList do
                        local appNode = expNodeList[workIndex]
                        table.insert( argList, appNode )
                        if workIndex == #expNodeList then
                           for __index, expType in ipairs( appNode:get_expTypeList() ) do
                              table.insert( argTypeList, expType )
                           end
                           
                        else
                         
                           table.insert( argTypeList, appNode:get_expType() )
                        end
                        
                     end
                     
                     
                     local mRetExp
                     
                     do
                        local workMRetExp = expListNode:get_mRetExp()
                        if workMRetExp ~= nil then
                           mRetExp = Nodes.MRetExp.new(workMRetExp:get_exp(), workMRetExp:get_index() - index + 1)
                           
                           if #argList == 0 then
                              return Nodes.ExpSubDDDNode.create( self.nodeManager, expNode:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), workMRetExp:get_exp():get_expTypeList(), workMRetExp:get_exp(), index - workMRetExp:get_index() ), true
                           end
                           
                        else
                           mRetExp = nil
                        end
                     end
                     
                     local newExpListNode = Nodes.ExpListNode.create( self.nodeManager, expNode:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), argTypeList, argList, mRetExp, expListNode:get_followOn() )
                     
                     workNode = Nodes.ExpToDDDNode.create( self.nodeManager, expNode:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {dstType}, newExpListNode )
                     return workNode, true
                  else
                   
                     local castType
                     
                     if validCastToGen then
                        castType = _lune.unwrapDefault( alt2typeMap[dstType], dstType)
                     else
                      
                        castType = dstType
                     end
                     
                     workNode = Nodes.ExpCastNode.create( self.nodeManager, expNode:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {expNode:get_expType()}, expNode, castType, Nodes.CastKind.Implicit )
                  end
                  
               end
               
            end
            
         end
      end
      
      return workNode, false
   end
   
   for index, expNode in ipairs( expNodeList ) do
      local workNode = expNode
      local stopFlag = false
      if #dstTypeList >= index then
         if index == #expNodeList and expNode:get_expType():get_kind() == Ast.TypeInfoKind.DDD then
            for dstIndex = index, #dstTypeList do
               workNode = expNode
               workNode, stopFlag = process( dstIndex, dstTypeList[dstIndex], expNode, workNode )
               table.insert( newExpNodeList, workNode )
               table.insert( expTypeList, workNode:get_expType() )
            end
            
            break
         else
          
            workNode, stopFlag = process( index, dstTypeList[index], expNode, workNode )
         end
         
      end
      
      
      table.insert( newExpNodeList, workNode )
      table.insert( expTypeList, workNode:get_expType() )
      if stopFlag then
         break
      end
      
   end
   
   
   if not hasModNode then
      return nil
   end
   
   
   local newMRetExp = nil
   do
      local mRetExp = expListNode:get_mRetExp()
      if mRetExp ~= nil then
         if mRetExp:get_index() <= #dstTypeList and dstTypeList[mRetExp:get_index()]:get_kind() ~= Ast.TypeInfoKind.DDD then
            newMRetExp = mRetExp
         end
         
      end
   end
   
   
   return Nodes.ExpListNode.create( self.nodeManager, expListNode:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), expTypeList, newExpNodeList, newMRetExp, expListNode:get_followOn() )
end


function TransUnit:checkMatchType( message, pos, dstTypeList, expListNode, allowDstShort, warnForFollow, workAlt2typeMap )

   local expNodeList = _lune.nilacc( expListNode, 'get_expList', 'callmtd' )
   if  nil == expNodeList then
      local _expNodeList = expNodeList
   
      expNodeList = {}
   end
   
   local warnForFollowSrcIndex = nil
   local expTypeList = {}
   local workExpNodeList = expNodeList
   local hasAbbr = false
   if #expNodeList > 0 then
      if expNodeList[#expNodeList]:get_kind() == Nodes.NodeKind.get_Abbr() then
         hasAbbr = true
         local workList = {}
         for __index, node in ipairs( expNodeList ) do
            table.insert( workList, node )
         end
         
         table.remove( workList )
         workExpNodeList = workList
      end
      
   end
   
   local realExpNum = -1
   
   for index, expNode in ipairs( workExpNodeList ) do
      if realExpNum == -1 and expNode:get_kind() == Nodes.NodeKind.get_ExpAccessMRet() then
         realExpNum = index - 1
      end
      
      if index == #workExpNodeList and _lune.__Cast( expNode, 3, Nodes.ExpMacroArgExpNode ) == nil then
         for __index, expType in ipairs( expNode:get_expTypeList() ) do
            table.insert( expTypeList, expType )
         end
         
      else
       
         table.insert( expTypeList, expNode:get_expType() )
      end
      
   end
   
   if realExpNum == -1 then
      realExpNum = #workExpNodeList
   end
   
   
   if warnForFollow and #expTypeList > realExpNum then
      warnForFollowSrcIndex = realExpNum + 1
   end
   
   if hasAbbr then
      table.insert( expTypeList, Ast.builtinTypeAbbr )
   end
   
   
   local alt2typeMap
   
   if workAlt2typeMap ~= nil then
      alt2typeMap = workAlt2typeMap
   else
      alt2typeMap = Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
   end
   
   Ast.CanEvalCtrlTypeInfo.setupNeedAutoBoxing( alt2typeMap, self.processInfo )
   
   local result, mess = Ast.TypeInfo.checkMatchType( self.processInfo, dstTypeList, expTypeList, allowDstShort, warnForFollowSrcIndex, alt2typeMap )
   do
      local _switchExp = result
      if _switchExp == Ast.MatchType.Error then
         self:addErrMess( pos, string.format( "%s: %s", message, mess) )
      elseif _switchExp == Ast.MatchType.Warn then
         if not self.ctrl_info.checkingDefineAbbr and Code.isMessageOf( Code.ID.nothing_define_abbr, mess ) then
         else
          
            self:addWarnMess( pos, string.format( "%s: %s", message, mess) )
         end
         
      end
   end
   
   
   if expListNode ~= nil then
      local autoBoxingCount = 0
      local hasImplictCast = false
      local newExpListNode
      
      
      if result ~= Ast.MatchType.Error then
         do
            local workList = self:checkImplicitCast( alt2typeMap, false, dstTypeList, expListNode, function ( dstType, expNode )
            
               if Ast.CanEvalCtrlTypeInfo.canAutoBoxing( dstType, expNode:get_expType() ) then
                  autoBoxingCount = autoBoxingCount + 1
                  
                  return Nodes.BoxingNode.create( self.nodeManager, expNode:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {dstType}, expNode )
               end
               
               return nil
            end )
            if workList ~= nil then
               newExpListNode = workList
               hasImplictCast = true
            else
               newExpListNode = nil
            end
         end
         
      else
       
         newExpListNode = nil
      end
      
      
      if autoBoxingCount > 0 then
         if not Ast.CanEvalCtrlTypeInfo.finishNeedAutoBoxing( alt2typeMap, autoBoxingCount ) then
            self:addErrMess( pos, string.format( "illegal auto boxing error -- %d", autoBoxingCount) )
         end
         
         
         return result, alt2typeMap, newExpListNode, expTypeList
      elseif Ast.CanEvalCtrlTypeInfo.hasNeedAutoBoxing( alt2typeMap ) then
         self:addErrMess( pos, "not support auto boxing" )
      end
      
      
      if hasImplictCast then
         return result, alt2typeMap, newExpListNode, expTypeList
      end
      
   end
   
   if not Ast.CanEvalCtrlTypeInfo.finishNeedAutoBoxing( alt2typeMap, 0 ) then
      self:addErrMess( pos, "can't auto boxing error" )
   end
   
   
   return result, alt2typeMap, nil, expTypeList
end


function TransUnit:checkMatchValType( pos, funcTypeInfo, expList, genericTypeList, genericsClass )

   local _
   local argTypeList = funcTypeInfo:get_argTypeInfoList()
   if funcTypeInfo:get_kind() == Ast.TypeInfoKind.Ext then
      local extTypeList, mess = Ast.convToExtTypeList( self.processInfo, argTypeList )
      if extTypeList ~= nil then
         argTypeList = extTypeList
      else
         self:addErrMess( pos, string.format( "not support argType on Luaval -- %s", mess) )
      end
      
   end
   
   
   do
      local _switchExp = funcTypeInfo
      if _switchExp == builtinFunc.list_insert or _switchExp == builtinFunc.set_add or _switchExp == builtinFunc.set_del then
         
      elseif _switchExp == builtinFunc.list_sort then
         local alt2typeMap = Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
         local callback = self.processInfo:createFunc( false, false, nil, Ast.TypeInfoKind.Func, Ast.headTypeInfo, false, false, true, Ast.AccessMode.Pri, "sort", Ast.Async.Async, nil, {genericTypeList[1], genericTypeList[1]}, {Ast.builtinTypeBool}, false )
         argTypeList = {callback:get_nilableTypeInfo()}
      elseif _switchExp == builtinFunc.list_remove then
      end
   end
   
   
   local warnForFollow = true
   
   if expList ~= nil then
      if expList:get_followOn() then
         warnForFollow = false
      end
      
   end
   
   
   local alt2typeMap
   
   if funcTypeInfo:get_kind() == Ast.TypeInfoKind.Method then
      if genericsClass ~= nil then
         if funcTypeInfo:get_rawTxt() == "__init" then
            alt2typeMap = genericsClass:createAlt2typeMap( true )
         else
          
            if #funcTypeInfo:get_itemTypeInfoList() == 0 then
               alt2typeMap = genericsClass:createAlt2typeMap( false )
            else
             
               alt2typeMap = genericsClass:createAlt2typeMap( true )
               for __index, itemType in ipairs( genericsClass:get_itemTypeInfoList() ) do
                  if itemType:get_kind() == Ast.TypeInfoKind.Alternate and not alt2typeMap[itemType] then
                     alt2typeMap[itemType] = Ast.builtinTypeNone
                  end
                  
               end
               
            end
            
         end
         
      else
         self:error( "none class" )
      end
      
   else
    
      alt2typeMap = Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( #funcTypeInfo:get_itemTypeInfoList() > 0 )
   end
   
   local matchResult, _3524, newExpNodeList = self:checkMatchType( funcTypeInfo:getTxt(  ), pos, argTypeList, expList, false, warnForFollow, alt2typeMap )
   
   if expList and newExpNodeList then
      return matchResult, alt2typeMap, newExpNodeList
   end
   
   return matchResult, alt2typeMap, expList
end


function TransUnit:analyzeListItems( firstPos, nextToken, termTxt, expectTypeList )

   local expList = nil
   local itemCommonType = _lune.newAlge( Ast.CommonType.Normal, {Ast.builtinTypeNone})
   
   if nextToken.txt ~= termTxt then
      self:pushback(  )
      expList = self:analyzeExpList( false, false, false, nil, expectTypeList, expectTypeList ~= nil )
      self:checkNextToken( termTxt )
      local nodeList = (_lune.unwrap( expList) ):get_expList()
      for __index, exp in ipairs( nodeList ) do
         itemCommonType = Ast.TypeInfo.getCommonTypeCombo( self.processInfo, itemCommonType, _lune.newAlge( Ast.CommonType.Normal, {exp:get_expType()}), Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false ) )
      end
      
   end
   
   
   local itemTypeInfo
   
   do
      local _matchExp = itemCommonType
      if _matchExp[1] == Ast.CommonType.Normal[1] then
         local info = _matchExp[2][1]
      
         itemTypeInfo = info
      elseif _matchExp[1] == Ast.CommonType.Combine[1] then
         local info = _matchExp[2][1]
      
         itemTypeInfo = info:get_typeInfo( self.processInfo )
      end
   end
   
   
   if itemTypeInfo:get_kind() == Ast.TypeInfoKind.DDD then
      if #itemTypeInfo:get_itemTypeInfoList() > 0 then
         itemTypeInfo = itemTypeInfo:get_itemTypeInfoList()[1]
      else
       
         itemTypeInfo = Ast.builtinTypeStem_
      end
      
   end
   
   
   if not expectTypeList then
      local _
      local expTypeList = {}
      if expList ~= nil then
         for index, expNode in ipairs( expList:get_expList() ) do
            if index == #expList:get_expList() then
               if expNode:get_expType():get_kind() == Ast.TypeInfoKind.DDD then
                  table.insert( expTypeList, expNode:get_expType() )
               else
                
                  for _3559 = 1, #expNode:get_expTypeList() do
                     table.insert( expTypeList, itemTypeInfo )
                  end
                  
               end
               
            else
             
               table.insert( expTypeList, itemTypeInfo )
            end
            
         end
         
      end
      
      local _3562, _3563, workExpList = self:checkMatchType( "List constructor", firstPos, expTypeList, expList, false, false, nil )
      if workExpList ~= nil then
         expList = workExpList
      end
      
   end
   
   return expList, itemTypeInfo
end


function TransUnit:analyzeListConst( token, expectType )

   local nextToken = self:getToken(  )
   
   local expectTypeList = nil
   if _lune.nilacc( expectType, 'get_kind', 'callmtd' ) == Ast.TypeInfoKind.List then
      do
         local itemTypeInfoList = _lune.nilacc( expectType, 'get_itemTypeInfoList', 'callmtd' )
         if itemTypeInfoList ~= nil then
            expectTypeList = {itemTypeInfoList[1]}
         end
      end
      
   end
   
   
   local expList, itemTypeInfo = self:analyzeListItems( token.pos, nextToken, "]", expectTypeList )
   
   local typeInfoList = {Ast.builtinTypeNone}
   if token.txt == '[' then
      typeInfoList = {self.processInfo:createList( Ast.AccessMode.Local, self:getCurrentClass(  ), {itemTypeInfo}, Ast.MutMode.Mut )}
      return Nodes.LiteralListNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), typeInfoList, expList )
   else
    
      typeInfoList = {self.processInfo:createArray( Ast.AccessMode.Local, self:getCurrentClass(  ), {itemTypeInfo}, Ast.MutMode.Mut )}
      return Nodes.LiteralArrayNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), typeInfoList, expList )
   end
   
end


function TransUnit:analyzeSetConst( token, expectType )

   
   self.helperInfo.useSet = true
   
   local nextToken = self:getToken(  )
   
   local expectTypeList = nil
   if _lune.nilacc( expectType, 'get_kind', 'callmtd' ) == Ast.TypeInfoKind.Set then
      do
         local itemTypeInfoList = _lune.nilacc( expectType, 'get_itemTypeInfoList', 'callmtd' )
         if itemTypeInfoList ~= nil then
            expectTypeList = {itemTypeInfoList[1]}
         end
      end
      
   end
   
   
   local expList, itemTypeInfo = self:analyzeListItems( token.pos, nextToken, ")", expectTypeList )
   
   if itemTypeInfo:get_nilable() then
      if expList ~= nil then
         for __index, exp in ipairs( expList:get_expList() ) do
            local expType = exp:get_expType()
            if expType:get_nilable() then
               self:addErrMess( exp:get_pos(), string.format( "'Set' object can't store nilable. -- %s", expType:getTxt(  )) )
            end
            
         end
         
      end
      
   end
   
   
   local typeInfoList = {Ast.builtinTypeNone}
   typeInfoList = {self.processInfo:createSet( Ast.AccessMode.Local, self:getCurrentClass(  ), {itemTypeInfo}, Ast.MutMode.Mut )}
   return Nodes.LiteralSetNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), typeInfoList, expList )
end


function TransUnit:analyzeMapConst( token, expectType )

   local nextToken = self:getToken(  )
   local map = {}
   local pairList = {}
   local keyTypeInfo = Ast.builtinTypeNone
   local valTypeInfo = Ast.builtinTypeNone
   
   local function getMapKeyValType( pos, keyFlag, typeInfo, expType )
   
      if expType:get_nilable() then
         if keyFlag then
            self:addErrMess( pos, string.format( "map key can't set a nilable -- %s", expType:getTxt(  )) )
         end
         
         if expType:equals( self.processInfo, Ast.builtinTypeNil ) then
            return typeInfo
         end
         
         expType = expType:get_nonnilableType()
      end
      
      return Ast.TypeInfo.getCommonType( self.processInfo, typeInfo, expType, Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false ) )
   end
   
   local expectKeyType = nil
   local expectValType = nil
   if _lune.nilacc( expectType, 'get_kind', 'callmtd' ) == Ast.TypeInfoKind.Map then
      do
         local itemTypeInfoList = _lune.nilacc( expectType, 'get_itemTypeInfoList', 'callmtd' )
         if itemTypeInfoList ~= nil then
            expectKeyType = itemTypeInfoList[1]
            expectValType = itemTypeInfoList[2]
         end
      end
      
   end
   
   
   while true do
      if nextToken.txt == "}" then
         break
      end
      
      self:pushback(  )
      
      local key = self:analyzeExpOneRVal( false, false, nil, expectKeyType )
      keyTypeInfo = getMapKeyValType( key:get_pos(), true, keyTypeInfo, key:get_expType() )
      
      self:checkNextToken( ":" )
      
      local val = self:analyzeExpOneRVal( false, false, nil, expectValType )
      valTypeInfo = getMapKeyValType( val:get_pos(), false, valTypeInfo, val:get_expType() )
      table.insert( pairList, Nodes.PairItem.new(key, val) )
      map[key] = val
      nextToken = self:getToken(  )
      if nextToken.txt ~= "," then
         break
      end
      
      nextToken = self:getToken(  )
   end
   
   
   local typeInfo = self.processInfo:createMap( Ast.AccessMode.Local, self:getCurrentClass(  ), keyTypeInfo, valTypeInfo, Ast.MutMode.Mut )
   
   self:checkToken( nextToken, "}" )
   return Nodes.LiteralMapNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, map, pairList )
end


function TransUnit:evalMacroOp( firstToken, macroTypeInfo, expList, evalMacroCallback )

   local parser, mess = self.macroCtrl:evalMacroOp( self.parser:getStreamName(  ), firstToken, macroTypeInfo, expList )
   
   local bakParser = self.parser
   
   if parser ~= nil then
      self.parser = Parser.DefaultPushbackParser.new(parser)
   else
      self:error( _lune.unwrap( mess) )
   end
   
   
   self.macroCtrl:startExpandMode( firstToken.pos.lineNo, macroTypeInfo, evalMacroCallback )
   local nextToken = self:getTokenNoErr(  )
   
   self.parser = bakParser
   
   if nextToken ~= Parser.getEofToken(  ) then
      self:addErrMess( firstToken.pos, string.format( "remain macro expand-statement token -- '%s'(%d:%d)", nextToken.txt, nextToken.pos.lineNo, nextToken.pos.column) )
      if not macroTypeInfo:get_externalFlag() then
         self:addErrMess( nextToken.pos, string.format( "remain macro expand-statement token -- '%s'", nextToken.txt) )
      end
      
   end
   
end


function TransUnit:evalMacro( firstToken, macroTypeInfo, expList )

   local stmtList = {}
   
   self:evalMacroOp( firstToken, macroTypeInfo, expList, function (  )
   
      if #macroTypeInfo:get_retTypeInfoList() == 0 then
         self:analyzeStatementList( stmtList, true, "}" )
      else
       
         table.insert( stmtList, self:analyzeExp( false, false, false ) )
      end
      
   end )
   local expTypeList = macroTypeInfo:get_retTypeInfoList()
   if #macroTypeInfo:get_retTypeInfoList() > 0 then
      local macroRetTypeList = macroTypeInfo:get_retTypeInfoList()
      if #stmtList == 1 then
         local node = stmtList[1]
         local retType = macroRetTypeList[1]
         if retType:equals( self.processInfo, Ast.builtinTypeMultiExp ) then
            expTypeList = node:get_expTypeList()
         elseif retType:equals( self.processInfo, Ast.builtinTypeExp ) then
            if #node:get_expTypeList() == 1 then
               expTypeList = node:get_expTypeList()
            else
             
               self:addErrMess( firstToken.pos, "__exp can't return multiple values. use __exps." )
            end
            
         elseif #node:get_expTypeList() == 1 then
            if retType:equals( self.processInfo, node:get_expType() ) then
               expTypeList = node:get_expTypeList()
            else
             
               self:addErrMess( firstToken.pos, string.format( "mismatch type -- %s != %s", macroRetTypeList[1]:getTxt(  ), node:get_expType():getTxt(  )) )
            end
            
         else
          
            self:addErrMess( firstToken.pos, "macro can't return multiple values." )
         end
         
      else
       
         self:addErrMess( firstToken.pos, "macro to return value must be one statemnt." )
      end
      
   else
    
      expTypeList = {Ast.builtinTypeNone}
   end
   
   
   return Nodes.ExpMacroExpNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), expTypeList, macroTypeInfo, stmtList )
end

local function findForm( format )

   local remain = format:gsub( "%%%%", "" )
   local opList = {}
   
   while true do
      local pos, endPos = nil, nil
      do
         local index, endIndex = remain:find( "^%%%-?[%d]*%a" )
         if index ~= nil and  endIndex ~= nil then
            pos, endPos = index, endIndex
         else
            do
               local index, endIndex = remain:find( "[^%%]%%%-?[%d]*%a" )
               if index ~= nil and  endIndex ~= nil then
                  pos, endPos = index + 1, endIndex
               end
            end
            
         end
      end
      
      
      if pos ~= nil and endPos ~= nil then
         local op = remain:sub( pos, endPos )
         table.insert( opList, op )
         remain = remain:sub( endPos + 1 )
      else
         break
      end
      
   end
   
   return opList
end
_moduleObj.findForm = findForm

local FormType = {}
_moduleObj.FormType = FormType
FormType._val2NameMap = {}
function FormType:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "FormType.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function FormType._from( val )
   if FormType._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
FormType.__allList = {}
function FormType.get__allList()
   return FormType.__allList
end

FormType.Match = 0
FormType._val2NameMap[0] = 'Match'
FormType.__allList[1] = FormType.Match
FormType.NeedConv = 1
FormType._val2NameMap[1] = 'NeedConv'
FormType.__allList[2] = FormType.NeedConv
FormType.Unmatch = 2
FormType._val2NameMap[2] = 'Unmatch'
FormType.__allList[3] = FormType.Unmatch

local function isMatchStringFormatType( opKind, argType, luaVer )

   do
      local enumType = _lune.__Cast( argType:get_srcTypeInfo():get_aliasSrc(), 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         argType = enumType:get_valTypeInfo()
      end
   end
   
   do
      local _switchExp = string.byte( opKind, #opKind )
      if _switchExp == 115 then
         if argType:get_srcTypeInfo() ~= Ast.builtinTypeString then
            if not luaVer:get_canFormStem2Str() then
               return FormType.NeedConv, Ast.builtinTypeString
            end
            
         end
         
      elseif _switchExp == 113 then
         if argType:get_srcTypeInfo() ~= Ast.builtinTypeString then
            return FormType.Unmatch, Ast.builtinTypeString
         end
         
      elseif _switchExp == 65 or _switchExp == 97 or _switchExp == 69 or _switchExp == 101 or _switchExp == 102 or _switchExp == 71 or _switchExp == 103 then
         if argType:get_srcTypeInfo() ~= Ast.builtinTypeReal then
            return FormType.Unmatch, Ast.builtinTypeReal
         end
         
      else 
         
            if argType:get_srcTypeInfo() ~= Ast.builtinTypeInt and argType:get_srcTypeInfo() ~= Ast.builtinTypeChar then
               return FormType.Unmatch, Ast.builtinTypeInt
            end
            
      end
   end
   
   return FormType.Match, Ast.builtinTypeNone
end
_moduleObj.isMatchStringFormatType = isMatchStringFormatType

function TransUnit:checkStringFormat( pos, formatTxt, argTypeList )

   local opList = findForm( formatTxt )
   if #opList ~= #argTypeList then
      self:addErrMess( pos, string.format( "argument number is mismatch -- %d != %d (%s)", #opList, #argTypeList, formatTxt:sub( 1, 20 )) )
      return 
   end
   
   
   for index, op in ipairs( opList ) do
      local argType = argTypeList[index]
      local match, reqType = isMatchStringFormatType( op, argType, self.targetLuaVer )
      if match == FormType.Unmatch then
         local mess = string.format( "type must be %s -- %s", reqType:getTxt(  ), argType:getTxt(  ))
         self:addErrMess( pos, string.format( "argument(%d) %s", index, mess) )
      end
      
   end
   
end

function TransUnit:prepareExpCall( position, funcTypeInfo, genericTypeList, genericsClass )

   
   if funcTypeInfo:get_kind() == Ast.TypeInfoKind.Macro then
      self.macroCtrl:startAnalyzeArgMode( funcTypeInfo )
   end
   
   
   local work = self:getToken(  )
   local argList = nil
   if work.txt ~= ")" then
      self:pushback(  )
      argList = self:analyzeExpList( false, false, false, nil, funcTypeInfo:get_argTypeInfoList() )
      self:checkNextToken( ")" )
      
      if argList ~= nil then
         for __index, argNode in ipairs( argList:get_expList() ) do
            if not argNode:canBeRight( self.processInfo ) and argNode:get_kind() ~= Nodes.NodeKind.get_Abbr() then
               self:addErrMess( argNode:get_pos(), string.format( "this node can't be r-value. -- %s", Nodes.getNodeKindName( argNode:get_kind() )) )
            end
            
         end
         
      end
      
   end
   
   local matchResult, alt2typeMap, workArgList = self:checkMatchValType( position, funcTypeInfo, argList, genericTypeList, genericsClass )
   
   if funcTypeInfo:get_kind() == Ast.TypeInfoKind.Macro and matchResult == Ast.MatchType.Error then
      self:error( string.format( "unmatch macro arguments. -- %s", funcTypeInfo:getTxt(  )) )
   end
   
   
   if funcTypeInfo:get_kind() == Ast.TypeInfoKind.Macro then
      self.macroCtrl:finishMacroMode(  )
   end
   
   
   return alt2typeMap, workArgList
end

function TransUnit:checkArgForStringForm( firstToken, argList )

   local formArgTypeList = {}
   local formatTxt = ""
   if #argList:get_expList() > 0 then
      local argNode = argList:get_expList()[1]
      if argNode:get_kind() ~= Nodes.NodeKind.get_LiteralString() then
         return 
      end
      
      do
         local literal = argNode:getLiteral(  )
         if literal ~= nil then
            do
               local _matchExp = literal
               if _matchExp[1] == Nodes.Literal.Str[1] then
                  local val = _matchExp[2][1]
               
                  formatTxt = val
               end
            end
            
         end
      end
      
   end
   
   if #argList:get_expList() > 1 then
      do
         local toDDDNode = _lune.__Cast( argList:get_expList()[2], 3, Nodes.ExpToDDDNode )
         if toDDDNode ~= nil then
            for __index, workNode in ipairs( toDDDNode:get_expList():get_expList() ) do
               table.insert( formArgTypeList, workNode:get_expType() )
            end
            
         else
            self:error( string.format( "illegal node -- %s", Nodes.getNodeKindName( argList:get_expList()[2]:get_kind() )) )
         end
      end
      
   end
   
   self:checkStringFormat( firstToken.pos, formatTxt, formArgTypeList )
end

function TransUnit:checkArgForSort( firstToken, genericTypeList, argList )

   if #argList:get_expTypeList() > 0 then
      local callback = argList:get_expTypeList()[1]
      if callback == Ast.builtinTypeAbbr then
         return 
      end
      
      if #callback:get_retTypeInfoList() ~= 1 then
         self:addErrMess( firstToken.pos, string.format( "The callback's to return value of sort() must have a value. -- %d", #callback:get_retTypeInfoList()) )
         return 
      end
      
      if not Ast.builtinTypeBool:equals( self.processInfo, callback:get_retTypeInfoList()[1] ) then
         self:addErrMess( firstToken.pos, string.format( "The callback's return type of sort() must be bool. -- '%s'", callback:get_retTypeInfoList()[1]:getTxt(  )) )
      end
      
      if #callback:get_argTypeInfoList() ~= 2 then
         self:addErrMess( firstToken.pos, string.format( "The callback's argument must have 2 arguments. -- '%s'", callback:get_display_stirng()) )
      end
      
      if #genericTypeList == 1 then
         for index, argType in ipairs( callback:get_argTypeInfoList() ) do
            if not genericTypeList[1]:equals( self.processInfo, argType ) then
               self:addErrMess( firstToken.pos, string.format( "The callback's argument(%d) type must be -- '%s'", index, genericTypeList[1]:getTxt(  )) )
            end
            
         end
         
      else
       
         self:addErrMess( firstToken.pos, "The generics of the list is illegal" )
      end
      
   end
   
end


function TransUnit:processFunc( firstToken, nextToken, refFieldNode, funcExp, funcType, alt2typeMap, genericTypeList, genericsClass, argList )

   local funcSymbol
   
   local symbolInfoList = funcExp:getSymbolInfo(  )
   if #symbolInfoList > 0 then
      local symbol = symbolInfoList[1]
      if symbol:get_kind() == Ast.SymbolKind.Typ then
         self:addErrMess( funcExp:get_pos(), string.format( "can't call any Type. -- %s", symbol:get_name()) )
      end
      
      funcSymbol = symbol
   else
    
      funcSymbol = nil
   end
   
   
   local funcTypeInfo = funcExp:get_expType():get_srcTypeInfo()
   local nilAccess
   
   
   if nextToken.txt == "$(" then
      if funcTypeInfo:get_nilable() then
         funcTypeInfo = funcTypeInfo:get_nonnilableType()
         nilAccess = true
      else
       
         self:addWarnMess( funcExp:get_pos(), string.format( "This is not nilable. -- %s", funcTypeInfo:getTxt(  )) )
         nilAccess = false
      end
      
   else
    
      nilAccess = false
   end
   
   
   do
      local _switchExp = (funcTypeInfo:get_kind() )
      if _switchExp == Ast.TypeInfoKind.Method or _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
      else 
         
            do
               local extType = _lune.__Cast( funcTypeInfo, 3, Ast.ExtTypeInfo )
               if extType ~= nil then
                  do
                     local _switchExp = (extType:get_extedType():get_kind() )
                     if _switchExp == Ast.TypeInfoKind.Method or _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
                     else 
                        
                           self:error( string.format( "can't call the type -- %s, %s", funcTypeInfo:getTxt(  ), Ast.TypeInfoKind:_getTxt( funcTypeInfo:get_kind())
                           ) )
                     end
                  end
                  
               else
                  self:error( string.format( "can't call the type -- %s, %s", funcTypeInfo:getTxt(  ), Ast.TypeInfoKind:_getTxt( funcTypeInfo:get_kind())
                  ) )
               end
            end
            
      end
   end
   
   local retTypeInfoList = {}
   for index, retType in ipairs( funcTypeInfo:get_retTypeInfoList() ) do
      table.insert( retTypeInfoList, retType )
      do
         local applyType = retType:applyGeneric( self.processInfo, alt2typeMap, self.moduleType )
         if applyType ~= nil then
            retTypeInfoList[index] = applyType
         else
            if funcTypeInfo == builtinFunc.list_remove then
               retTypeInfoList[index] = genericTypeList[1]:get_nilableTypeInfo()
            elseif funcTypeInfo:get_kind() == Ast.TypeInfoKind.Func and (funcTypeInfo:get_rawTxt() == "_fromMap" or funcTypeInfo:get_rawTxt() == "_fromStem" ) and genericsClass:isInheritFrom( self.processInfo, Ast.builtinTypeMapping, alt2typeMap ) then
               retTypeInfoList[index] = genericsClass:get_nilableTypeInfo()
            else
             
               self:addErrMess( firstToken.pos, string.format( "not support generics yet. -- %s", retType:getTxt(  )) )
            end
            
         end
      end
      
   end
   
   
   if refFieldNode ~= nil then
      if funcTypeInfo:equals( self.processInfo, builtinFunc.list_unpack ) or funcTypeInfo:equals( self.processInfo, builtinFunc.array_unpack ) then
         local prefixType = refFieldNode:get_prefix():get_expType()
         if #prefixType:get_itemTypeInfoList() > 0 then
            local dddType = self.processInfo:createDDD( prefixType:get_itemTypeInfoList()[1], false, false )
            retTypeInfoList = {}
            table.insert( retTypeInfoList, dddType )
         end
         
      end
      
   end
   
   
   if nilAccess then
      local retList = {}
      for __index, retType in ipairs( retTypeInfoList ) do
         if retType:get_nilable() then
            table.insert( retList, retType )
         else
          
            table.insert( retList, retType:get_nilableTypeInfo() )
         end
         
      end
      
      retTypeInfoList = retList
      self.helperInfo.useNilAccess = true
   end
   
   
   local errorFuncFlag = false
   if #retTypeInfoList > 0 then
      local retType = retTypeInfoList[1]
      if retType == Ast.builtinTypeNeverRet then
         errorFuncFlag = true
      end
      
   end
   
   
   if argList ~= nil then
      do
         local _switchExp = funcTypeInfo
         if _switchExp == builtinFunc.string_format then
            self:checkArgForStringForm( firstToken, argList )
         elseif _switchExp == builtinFunc.list_sort or _switchExp == builtinFunc.array_sort then
            self:checkArgForSort( firstToken, genericTypeList, argList )
         end
      end
      
   end
   
   
   if funcTypeInfo:equals( self.processInfo, builtinFunc.lns__kind ) then
      do
         local expList = _lune.nilacc( argList, 'get_expList', 'callmtd' )
         if expList ~= nil then
            if #expList > 0 then
               return Nodes.LuneKindNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeInt}, expList[1] )
            end
            
         end
      end
      
      return Nodes.LuneKindNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeInt}, self:createNoneNode( firstToken.pos ) )
   end
   
   
   if funcSymbol ~= nil then
      if funcSymbol:get_name() == "super" then
         return Nodes.ExpCallSuperNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), retTypeInfoList, funcSymbol:get_typeInfo():get_parentInfo(), funcSymbol:get_typeInfo(), argList )
      end
      
   end
   
   if funcType:get_kind() == Ast.TypeInfoKind.Ext then
      local work, err = Ast.convToExtTypeList( self.processInfo, retTypeInfoList )
      if work ~= nil then
         retTypeInfoList = work
      else
         self:addErrMess( firstToken.pos, err )
      end
      
   end
   
   
   local threading
   
   if nilAccess or funcType:get_kind() == Ast.TypeInfoKind.Ext or _lune._Set_has(builtinFunc:get_needThreadingTypes(), funcType:get_nonnilableType():get_srcTypeInfo() ) then
      threading = self:checkThreading( firstToken.pos )
      if threading and funcType:get_kind() == Ast.TypeInfoKind.Ext then
         self:addErrMess( firstToken.pos, "not support to use Luaval on thread." )
      end
      
   else
    
      threading = false
   end
   
   return Nodes.ExpCallNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), retTypeInfoList, funcExp, errorFuncFlag, nilAccess, threading, argList )
end


function TransUnit:checkCallAsync( funcTypeInfo, pos )

   if funcTypeInfo:get_asyncMode() == Ast.Async.Noasync then
      local curType = self:getCurrentNamespaceTypeInfo(  )
      do
         local _switchExp = curType:get_kind()
         if _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Method then
            do
               local _switchExp = curType:get_asyncMode()
               if _switchExp == Ast.Async.Async or _switchExp == Ast.Async.Transient then
                  self:addErrMess( pos, string.format( "can't access noasync function in async. -- %s on %s", funcTypeInfo:getTxt(  ), curType:getTxt(  )) )
               end
            end
            
         end
      end
      
      if self:getCurrentClass(  ):isInheritFrom( self.processInfo, Ast.builtinTypeRunner ) then
         self:addErrMess( pos, string.format( "can't access noasync function in __Runner. -- %s", funcTypeInfo:getTxt(  )) )
      end
      
   end
   
end


function TransUnit:analyzeExpCall( firstToken, funcExp, nextToken )

   self:checkSymbolHavingValue( funcExp:get_pos(), funcExp:getSymbolInfo(  ) )
   
   local funcTypeInfo = funcExp:get_expType():get_nonnilableType()
   
   self:checkCallAsync( funcTypeInfo, funcExp:get_effectivePos() )
   
   local genericTypeList = funcTypeInfo:get_itemTypeInfoList()
   local refFieldNode = nil
   local genericsClass = Ast.headTypeInfo
   do
      local refField = _lune.__Cast( funcExp, 3, Nodes.RefFieldNode )
      if refField ~= nil then
         refFieldNode = refField
         local classType = refField:get_prefix():get_expType()
         genericsClass = classType
         
         if funcTypeInfo:get_kind() == Ast.TypeInfoKind.Method then
            genericTypeList = classType:get_itemTypeInfoList()
         end
         
      end
   end
   
   
   local alt2typeMap, argList = self:prepareExpCall( funcExp:get_pos(), funcTypeInfo, genericTypeList, genericsClass )
   
   if funcTypeInfo:equals( self.processInfo, builtinFunc.list_insert ) then
      if argList ~= nil then
         if argList:get_expType():get_nilable() then
            self:addErrMess( argList:get_pos(), "list can't insert nilable" )
         end
         
      end
      
   end
   
   if funcTypeInfo:equals( self.processInfo, builtinFunc.set_add ) then
      if argList ~= nil then
         if argList:get_expType():get_nilable() then
            self:addErrMess( argList:get_pos(), "set can't add nilable" )
         end
         
      end
      
   elseif funcTypeInfo:equals( self.processInfo, builtinFunc.list_remove ) then
      if #genericTypeList > 0 then
         if genericTypeList[1]:get_nilable() then
            self:addWarnMess( funcExp:get_pos(), "remove() is dangerous for nilable's list." )
         end
         
      end
      
   end
   
   
   if funcTypeInfo:get_rawTxt() == "" then
      self:addErrMess( funcExp:get_pos(), "can't directly call the declared function." )
   end
   
   
   local exp
   
   if funcTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Macro then
      exp = self:evalMacro( firstToken, funcTypeInfo, argList )
      if self:checkThreading( firstToken.pos ) then
         self:addErrMess( firstToken.pos, "not support to use a macro" )
      end
      
   else
    
      exp = self:processFunc( firstToken, nextToken, refFieldNode, funcExp, funcTypeInfo, alt2typeMap, genericTypeList, genericsClass, argList )
   end
   
   
   return exp, self:getTokenNoErr(  )
end


function TransUnit:analyzeExpCast( firstToken, opTxt, exp )

   local castTypeNode = self:analyzeRefType( Ast.AccessMode.Local, false, false )
   local castType = castTypeNode:get_expType()
   
   if exp:get_expType():get_kind() == Ast.TypeInfoKind.Ext and castType:get_kind() ~= Ast.TypeInfoKind.Ext and castType:get_kind() ~= Ast.TypeInfoKind.Stem then
      if opTxt == "@@" or opTxt == "@@=" then
         castType = self:createModifier( castType, Ast.MutMode.IMut )
      end
      
      castType = self:createExtType( castTypeNode:get_pos(), castType )
   end
   
   
   if castType:get_kind() == Ast.TypeInfoKind.Form and exp:get_expType():get_nonnilableType():get_kind() == Ast.TypeInfoKind.Stem then
      if self:supportLang( LuneControl.Code.C ) then
         self:addWarnMess( castTypeNode:get_pos(), "not support cast from stem to form for transcompiling to c-lang." )
      end
      
   end
   
   
   local expType = exp:get_expType()
   
   if opTxt == "@@@" or opTxt == "@@=" then
      if #castType:get_itemTypeInfoList() > 0 then
         self:addErrMess( castTypeNode:get_pos(), string.format( "not support cast for generics class yet -- %s", castType:getTxt(  )) )
      end
      
      do
         local _switchExp = castType:get_kind()
         if _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Prim then
         else 
            
               if opTxt ~= "@@=" then
                  self:addErrMess( castTypeNode:get_pos(), string.format( "not support cast -- %s", castType:getTxt(  )) )
               end
               
         end
      end
      
      
      if opTxt == "@@=" then
         local orgExpType = expType:get_extedType()
         local orgCastType = castType:get_extedType()
         if orgCastType:get_kind() ~= Ast.TypeInfoKind.IF and orgCastType:get_kind() ~= Ast.TypeInfoKind.Class then
            self:addErrMess( castTypeNode:get_pos(), string.format( "'@@=' cast must be class or interface. -- %s", castType:getTxt(  )) )
         end
         
         if orgExpType:get_srcTypeInfo() ~= Ast.builtinTypeStem and orgExpType:get_kind() ~= Ast.TypeInfoKind.IF and orgExpType:get_kind() ~= Ast.TypeInfoKind.Class then
            self:addErrMess( castTypeNode:get_pos(), string.format( "'@@=' cast must be class or interface. -- %s", castType:getTxt(  )) )
         end
         
         if not Ast.isStruct( orgCastType ) then
            self:addErrMess( castTypeNode:get_pos(), string.format( "'@@=' cast type can't use class has method -- %s", castType:getTxt(  )) )
         end
         
      end
      
   else
    
      if castType ~= Ast.builtinTypeString and (castType:get_kind() == Ast.TypeInfoKind.IF or castType:get_kind() == Ast.TypeInfoKind.Class ) then
         self:addWarnMess( castTypeNode:get_pos(), string.format( "use '@@@' cast for class or interface. -- %s", castType:getTxt(  )) )
      end
      
   end
   
   
   if opTxt ~= "@@@" and expType:get_nilable() and not castType:get_nilable() then
      self:addErrMess( firstToken.pos, string.format( "can't cast from nilable to not nilable  -- %s->%s", expType:getTxt(  ), castType:getTxt(  )) )
   end
   
   if not Ast.TypeInfo.isMut( expType ) and Ast.TypeInfo.isMut( castType ) then
      castType = self:createModifier( castType, Ast.MutMode.IMut )
   end
   
   
   if castType:canEvalWith( self.processInfo, expType, Ast.CanEvalType.SetOp, {} ) then
      if castType:get_kind() ~= Ast.TypeInfoKind.Ext and not (castType:get_kind() == Ast.TypeInfoKind.Stem and Ast.isExtType( expType ) ) then
         self:addWarnMess( castTypeNode:get_pos(), string.format( "This cast isn't need. (%s <- %s)", castType:getTxt( self.typeNameCtrl ), expType:getTxt( self.typeNameCtrl )) )
      end
      
   elseif not expType:canEvalWith( self.processInfo, castType, Ast.CanEvalType.SetOp, {} ) then
      if not Ast.isNumberType( expType ) or not Ast.isNumberType( castType ) then
         self:addErrMess( castTypeNode:get_pos(), string.format( "This type can't cast. (%s <- %s)", castType:getTxt( self.typeNameCtrl ), expType:getTxt( self.typeNameCtrl )) )
      end
      
   end
   
   
   if opTxt == "@@@" then
      castType = castType:get_nilableTypeInfo()
   end
   
   
   return Nodes.ExpCastNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {castType}, self.nodeManager:MultiTo1( exp ), castType, opTxt ~= "@@@" and Nodes.CastKind.Force or Nodes.CastKind.Normal )
end


function TransUnit:analyzeExpCont( firstToken, exp, skipFlag, canLeftExp )

   local nextToken = self:getToken( true )
   if nextToken.kind == Parser.TokenKind.Eof then
      return exp
   end
   
   
   if not skipFlag then
      repeat 
         local matchFlag = false
         if nextToken.txt == "[" or nextToken.txt == "$[" then
            matchFlag = true
            exp = self:analyzeExpRefItem( nextToken, exp, nextToken.txt == "$[" )
            nextToken = self:getToken(  )
         end
         
         if nextToken.txt == "(" or nextToken.txt == "$(" then
            matchFlag = true
            
            exp, nextToken = self:analyzeExpCall( firstToken, exp, nextToken )
         end
         
      until not matchFlag
      
      do
         local _switchExp = nextToken.txt
         if _switchExp == "@@" or _switchExp == "@@@" or _switchExp == "@@=" then
            exp = self:analyzeExpCast( firstToken, nextToken.txt, exp )
            nextToken = self:getToken(  )
         end
      end
      
   end
   
   
   do
      local _switchExp = nextToken.txt
      if _switchExp == "." then
         return self:analyzeExpSymbol( firstToken, self:getToken(  ), ExpSymbolMode.Field, exp, skipFlag, canLeftExp )
      elseif _switchExp == "$." then
         return self:analyzeExpSymbol( firstToken, self:getToken(  ), ExpSymbolMode.FieldNil, exp, skipFlag, canLeftExp )
      elseif _switchExp == ".$" then
         return self:analyzeExpSymbol( firstToken, self:getToken(  ), ExpSymbolMode.Get, exp, skipFlag, canLeftExp )
      elseif _switchExp == "$.$" then
         return self:analyzeExpSymbol( firstToken, self:getToken(  ), ExpSymbolMode.GetNil, exp, skipFlag, canLeftExp )
      end
   end
   
   
   self:pushback(  )
   return exp
end


function TransUnit:analyzeAccessClassField( classTypeInfo, mode, token )

   do
      local _switchExp = classTypeInfo:get_kind(  )
      if _switchExp == Ast.TypeInfoKind.List then
         classTypeInfo = Ast.builtinTypeList
      elseif _switchExp == Ast.TypeInfoKind.Array then
         classTypeInfo = Ast.builtinTypeArray
      elseif _switchExp == Ast.TypeInfoKind.Set then
         classTypeInfo = Ast.builtinTypeSet
      end
   end
   
   local className = classTypeInfo:getTxt(  )
   local classScope = classTypeInfo:get_scope()
   if  nil == classScope then
      local _classScope = classScope
   
      self:error( string.format( "not found field: %s, %s", className, classTypeInfo) )
   end
   
   
   local symbolInfo = nil
   local fieldTypeInfo = nil
   local getterFlag = false
   if mode == ExpSymbolMode.Get or mode == ExpSymbolMode.GetNil then
      local fieldSymbolInfo = classScope:getSymbolInfo( string.format( "get_%s", token.txt), self.scope, false, self.scopeAccess )
      if fieldSymbolInfo ~= nil then
         if (fieldSymbolInfo:get_kind(  ) == Ast.SymbolKind.Mtd or fieldSymbolInfo:get_kind(  ) == Ast.SymbolKind.Fun ) then
            local retTypeList = fieldSymbolInfo:get_typeInfo():get_retTypeInfoList(  )
            symbolInfo = fieldSymbolInfo
            if #retTypeList > 0 then
               do
                  local applyedType = retTypeList[1]:applyGeneric( self.processInfo, classTypeInfo:createAlt2typeMap( false ), self.moduleType )
                  if applyedType ~= nil then
                     fieldTypeInfo = applyedType
                  else
                     fieldTypeInfo = retTypeList[1]
                  end
               end
               
            end
            
            if #fieldSymbolInfo:get_typeInfo():get_argTypeInfoList() > 0 then
               self:addErrMess( token.pos, string.format( "can't use '$' with -- %s", fieldSymbolInfo:get_typeInfo():getTxt(  )) )
            end
            
            
            getterFlag = true
         end
         
      end
      
   end
   
   if not symbolInfo then
      symbolInfo = classScope:getSymbolInfoField( token.txt, true, self.scope, self.scopeAccess )
      if not symbolInfo then
         symbolInfo = classScope:getSymbolInfoIfField( token.txt, self.scope, self.scopeAccess )
      end
      
      if symbolInfo ~= nil then
         fieldTypeInfo = symbolInfo:get_typeInfo()
      end
      
   end
   
   if not fieldTypeInfo then
      for name, val in pairs( classScope:get_symbol2SymbolInfoMap() ) do
         Util.log( string.format( "debug: %s, %s", name, val) )
      end
      
      self:error( string.format( "not found field typeInfo: %s.%s -- %s", className, token.txt, Ast.TypeInfoKind:_getTxt( classTypeInfo:get_kind())
      ) )
   end
   
   local typeInfo = _lune.unwrapDefault( fieldTypeInfo, Ast.builtinTypeNone)
   
   if symbolInfo ~= nil then
      if self:inAnalyzingState( AnalyzingState.InitBlock ) or self:inAnalyzingState( AnalyzingState.ClassMethod ) then
         local errorMess = nil
         if self.protoFuncMap[symbolInfo:get_typeInfo()] then
            errorMess = string.format( "It can't call prototype function from static -- %s", symbolInfo:get_name())
         end
         
         if errorMess ~= nil then
            self:addErrMess( token.pos, errorMess )
         end
         
      elseif self:inAnalyzingState( AnalyzingState.Constructor ) then
         local errorMess = nil
         if self.protoFuncMap[symbolInfo:get_typeInfo()] then
            errorMess = "It can't call prototype function from '__init'"
         else
          
            if symbolInfo:get_typeInfo():get_kind() == Ast.TypeInfoKind.Method and symbolInfo:get_scope() == classScope then
               for __index, val in pairs( classScope:get_symbol2SymbolInfoMap() ) do
                  if val:get_kind() == Ast.SymbolKind.Mbr and not val:get_staticFlag() then
                     if not val:get_hasValueFlag() and not val:get_typeInfo():get_nilable() then
                        errorMess = string.format( "Set member(%s) before to access the method-- %s", val:get_name(), symbolInfo:get_name())
                        break
                     end
                     
                  end
                  
               end
               
            end
            
         end
         
         if errorMess ~= nil then
            self:addErrMess( token.pos, errorMess )
         end
         
      end
      
   end
   
   
   return typeInfo, symbolInfo, getterFlag
end


function TransUnit:dumpComp( writer, pattern, symbolInfo, getterFlag )

   local symbol = symbolInfo:get_name()
   if pattern == "" or symbol:find( pattern ) then
      if getterFlag then
         writer:startParent( "candidate", false )
         local typeInfo = symbolInfo:get_typeInfo()
         writer:write( "type", string.format( "%s", Ast.SymbolKind:_getTxt( symbolInfo:get_kind())
         ) )
         do
            local _switchExp = (symbolInfo:get_kind() )
            if _switchExp == Ast.SymbolKind.Mtd or _switchExp == Ast.SymbolKind.Fun or _switchExp == Ast.SymbolKind.Mac then
               writer:write( "displayTxt", string.format( "$%s", (typeInfo:get_rawTxt():gsub( "^get_", "" ) )) )
            elseif _switchExp == Ast.SymbolKind.Mbr then
               writer:write( "displayTxt", string.format( "$%s: %s", symbolInfo:get_name(), typeInfo:getTxt(  )) )
            end
         end
         
      else
       
         writer:startParent( "candidate", false )
         local typeInfo = symbolInfo:get_typeInfo()
         writer:write( "type", string.format( "%s", Ast.SymbolKind:_getTxt( symbolInfo:get_kind())
         ) )
         do
            local _switchExp = (symbolInfo:get_kind() )
            if _switchExp == Ast.SymbolKind.Fun or _switchExp == Ast.SymbolKind.Mtd or _switchExp == Ast.SymbolKind.Mac then
               writer:write( "displayTxt", typeInfo:get_display_stirng_with( symbolInfo:get_name(), nil ) )
            elseif _switchExp == Ast.SymbolKind.Mbr or _switchExp == Ast.SymbolKind.Var or _switchExp == Ast.SymbolKind.Arg then
               local name = symbolInfo:get_name()
               do
                  local algeTypeInfo = _lune.__Cast( typeInfo, 3, Ast.AlgeTypeInfo )
                  if algeTypeInfo ~= nil then
                     do
                        local valInfo = algeTypeInfo:getValInfo( name )
                        if valInfo ~= nil then
                           if #valInfo:get_typeList() > 0 then
                              name = string.format( "%s(", name)
                              for index, itemType in ipairs( valInfo:get_typeList() ) do
                                 if index > 1 then
                                    name = name .. ","
                                 end
                                 
                                 name = name .. itemType:get_display_stirng()
                              end
                              
                              name = name .. ")"
                           end
                           
                        end
                     end
                     
                  end
               end
               
               writer:write( "displayTxt", string.format( "%s: %s", name, typeInfo:get_display_stirng()) )
            elseif _switchExp == Ast.SymbolKind.Typ then
               writer:write( "displayTxt", string.format( "%s", (typeInfo:get_display_stirng():gsub( "@", "" ) )) )
            end
         end
         
      end
      
      writer:endElement(  )
   end
   
   return true
end


function TransUnit:dumpFieldComp( writer, isPrefixType, prefixTypeInfo, pattern, getterPattern )

   local typeInfo = prefixTypeInfo
   local scope = typeInfo:get_scope()
   if  nil == scope then
      local _scope = scope
   
      return 
   end
   
   
   scope:filterTypeInfoField( true, self.scope, self.scopeAccess, function ( symbolInfo )
   
      if (isPrefixType ) then
         if not symbolInfo:get_staticFlag() and not symbolInfo:get_typeInfo():get_staticFlag() and symbolInfo:get_kind() ~= Ast.SymbolKind.Typ then
            return true
         end
         
      elseif symbolInfo:get_staticFlag() then
         return true
      end
      
      local symbol = symbolInfo:get_name()
      if symbol ~= "__init" and symbol ~= "__free" and symbol ~= "self" then
         if getterPattern ~= nil then
            if symbolInfo:get_kind() == Ast.SymbolKind.Mtd or symbolInfo:get_kind() == Ast.SymbolKind.Fun then
               local retList = symbolInfo:get_typeInfo():get_retTypeInfoList()
               if #retList == 1 then
                  return self:dumpComp( writer, getterPattern, symbolInfo, true )
               end
               
            end
            
            return true
         end
         
         return self:dumpComp( writer, pattern, symbolInfo, false )
      end
      
      return true
   end )
end


function TransUnit:dumpSymbolComp( writer, scope, pattern )

   scope:filterSymbolTypeInfo( scope, self.moduleScope, self.scopeAccess, function ( symbolInfo )
   
      return self:dumpComp( writer, pattern, symbolInfo, false )
   end )
end




function TransUnit:checkComp( token, callback )

   if self.analyzeMode == AnalyzeMode.Complete and self:isTargetToken( token ) then
      local currentModule = self.parser:getStreamName(  ):gsub( "%.lns", "" )
      currentModule = currentModule:gsub( ".*/", "" )
      local target = self.analyzeModule:gsub( "[^%.]+%.", "" )
      if currentModule == target then
         local jsonWriter = Writer.JSON.new(io.stdout)
         jsonWriter:startParent( "lunescript", false )
         local prefix = token.txt:gsub( "lune$", "" )
         jsonWriter:write( "prefix", prefix )
         jsonWriter:startParent( "candidateList", true )
         
         callback( jsonWriter, prefix )
         
         jsonWriter:endElement(  )
         jsonWriter:endElement(  )
         jsonWriter:fin(  )
         os.exit( 0 )
      end
      
   end
   
end


function TransUnit:checkFieldComp( getterFlag, token, prefixExp )

   if self.analyzeMode ~= AnalyzeMode.Complete then
      return 
   end
   
   local prefixSymbolInfoList = prefixExp:getSymbolInfo(  )
   local prefixSymbolInfo = nil
   if #prefixSymbolInfoList == 1 then
      prefixSymbolInfo = prefixSymbolInfoList[1]
   end
   
   
   self:checkComp( token, function ( jsonWriter, prefix )
   
      local getterPattern = nil
      if getterFlag then
         getterPattern = "^get_" .. prefix
      end
      
      local isPrefixType = false
      do
         local _exp = prefixSymbolInfo
         if _exp ~= nil then
            isPrefixType = _exp:get_kind() == Ast.SymbolKind.Typ
         end
      end
      
      self:dumpFieldComp( jsonWriter, isPrefixType, prefixExp:get_expType(), prefix == "" and "" or "^" .. prefix, getterPattern )
   end )
end


function TransUnit:checkEnumComp( token, enumTypeInfo )

   if self.analyzeMode ~= AnalyzeMode.Complete then
      return 
   end
   
   self:checkComp( token, function ( jsonWriter, prefix )
   
      local scope = enumTypeInfo:get_scope()
      if  nil == scope then
         local _scope = scope
      
         return 
      end
      
      
      local pattern = prefix == "" and "" or "^" .. prefix
      scope:filterTypeInfoField( true, self.scope, self.scopeAccess, function ( symbolInfo )
      
         if symbolInfo:get_kind() == Ast.SymbolKind.Mbr then
            return self:dumpComp( jsonWriter, pattern, symbolInfo, false )
         end
         
         return true
      end )
   end )
end


function TransUnit:checkAlgeComp( token, algeTypeInfo )

   if self.analyzeMode ~= AnalyzeMode.Complete then
      return 
   end
   
   self:checkComp( token, function ( jsonWriter, prefix )
   
      self:dumpFieldComp( jsonWriter, true, algeTypeInfo, prefix == "" and "" or "^" .. prefix, nil )
   end )
end


function TransUnit:checkSymbolComp( token )

   self:checkComp( token, function ( jsonWriter, prefix )
   
      self:dumpSymbolComp( jsonWriter, self.scope, prefix == "" and "" or "^" .. prefix )
   end )
end


function TransUnit:analyzeExpField( firstToken, fieldToken, mode, prefixExp )

   if #prefixExp:get_expTypeList() > 1 then
      prefixExp = Nodes.ExpMultiTo1Node.create( self.nodeManager, prefixExp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {prefixExp:get_expType()}, prefixExp )
   end
   
   
   local threading
   
   local accessNil = false
   if mode == ExpSymbolMode.FieldNil or mode == ExpSymbolMode.GetNil then
      accessNil = true
      if not prefixExp:get_expType():get_nilable() then
         self:addWarnMess( prefixExp:get_pos(), string.format( "This is not nilable. -- %s", prefixExp:get_expType():getTxt(  )) )
      end
      
      threading = self:checkThreading( firstToken.pos )
   else
    
      threading = false
   end
   
   if self.macroCtrl:get_analyzeInfo():isAnalyzingSymArg(  ) then
      if accessNil then
         self.helperInfo.useNilAccess = true
      end
      
      
      return Nodes.RefFieldNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeSymbol}, fieldToken, nil, accessNil, threading, prefixExp )
      
   end
   
   
   local typeInfo = Ast.builtinTypeStem_
   local prefixExpType = prefixExp:get_expType()
   if accessNil then
      if prefixExpType:get_nilable() then
         prefixExpType = prefixExpType:get_nonnilableType()
         
         if prefixExpType:get_srcTypeInfo():get_kind() == Ast.TypeInfoKind.Box then
            self:addErrMess( prefixExp:get_pos(), "Nilable can't support '$.' access yet" )
         end
         
      else
       
         accessNil = false
      end
      
   end
   
   
   local extFlag
   
   if Ast.isExtType( prefixExpType ) then
      extFlag = true
      prefixExpType = prefixExpType:get_extedType()
   else
    
      extFlag = false
   end
   
   
   self:checkFieldComp( mode == ExpSymbolMode.Get or mode == ExpSymbolMode.GetNil, fieldToken, prefixExp )
   
   if accessNil then
      self.helperInfo.useNilAccess = true
      do
         local _switchExp = prefixExpType:get_kind(  )
         if _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.Alge then
            self:addErrMess( firstToken.pos, string.format( "%s does not support $.", prefixExpType:getTxt( nil )) )
         end
      end
      
   end
   
   
   local prefixSymbolInfoList = prefixExp:getSymbolInfo(  )
   self:checkSymbolHavingValue( prefixExp:get_pos(), prefixSymbolInfoList )
   
   local getterTypeInfo = nil
   local symbolInfo = nil
   do
      local _switchExp = prefixExpType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Module or _switchExp == Ast.TypeInfoKind.ExtModule or _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Box or _switchExp == Ast.TypeInfoKind.Alternate then
         local getterFlag = false
         typeInfo, symbolInfo, getterFlag = self:analyzeAccessClassField( prefixExpType, mode, fieldToken )
         if getterFlag then
            do
               local _exp = symbolInfo
               if _exp ~= nil then
                  getterTypeInfo = _exp:get_typeInfo()
               end
            end
            
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.Alge then
         local scope = _lune.unwrap( prefixExpType:get_scope())
         local fieldName = fieldToken.txt
         local symbolInfoList = prefixExp:getSymbolInfo(  )
         local isTypeSymbol = false
         if #symbolInfoList > 0 then
            if symbolInfoList[1]:get_kind() == Ast.SymbolKind.Typ then
               isTypeSymbol = true
            end
            
         end
         
         if mode == ExpSymbolMode.Get then
            local moduleType = prefixExpType:getModule(  )
            if not moduleType:equals( self.processInfo, self.moduleType ) and not self.scope:getModuleInfo( moduleType ) then
               if not self.importedAliasMap[prefixExpType] then
                  self:addErrMess( fieldToken.pos, string.format( "need to import module -- %s", prefixExpType:getModule(  ):getTxt(  )) )
               end
               
            end
            
            fieldName = "get_" .. fieldName
            do
               local funcSymbol = scope:getSymbolInfoChild( fieldName )
               if funcSymbol ~= nil then
                  symbolInfo = funcSymbol
                  local funcType = funcSymbol:get_typeInfo()
                  if funcType:get_staticFlag() ~= isTypeSymbol then
                     self:addErrMess( prefixExp:get_pos(), string.format( "Can't access -- %s, %s", fieldName, isTypeSymbol) )
                  end
                  
                  
                  local retTypeList = funcType:get_retTypeInfoList()
                  if #retTypeList == 0 then
                     self:addErrMess( fieldToken.pos, string.format( "The func (%s) doesn't return value.", funcType:getTxt(  )) )
                  else
                   
                     typeInfo = retTypeList[1]
                  end
                  
               else
                  self:addErrMess( fieldToken.pos, string.format( "not found -- %s.", fieldName) )
                  typeInfo = Ast.builtinTypeNone
               end
            end
            
            
            getterTypeInfo = Ast.headTypeInfo
         else
          
            do
               local _exp = scope:getTypeInfoChild( fieldName )
               if _exp ~= nil then
                  typeInfo = _exp
                  if typeInfo:get_kind() == Ast.TypeInfoKind.Enum or typeInfo:get_kind() == Ast.TypeInfoKind.Alge then
                     if not isTypeSymbol then
                        self:addErrMess( fieldToken.pos, string.format( "can't access field -- %s", fieldToken.txt) )
                     end
                     
                  end
                  
               else
                  self:addErrMess( fieldToken.pos, string.format( "not found field -- %s", fieldToken.txt) )
                  typeInfo = Ast.builtinTypeInt
               end
            end
            
         end
         
      elseif _switchExp == Ast.TypeInfoKind.Map then
         local work = prefixExpType:get_itemTypeInfoList()[1]
         if not work:equals( self.processInfo, Ast.builtinTypeString ) then
            self:addErrMess( fieldToken.pos, string.format( "map key type is not str. (%s)", work:getTxt(  )) )
         end
         
         typeInfo = prefixExpType:get_itemTypeInfoList()[2]
         if not typeInfo:get_nilable() then
            typeInfo = typeInfo:get_nilableTypeInfo()
         end
         
         if extFlag then
            typeInfo = self:createExtType( fieldToken.pos, typeInfo )
         end
         
         return Nodes.ExpRefItemNode.create( self.nodeManager, fieldToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, prefixExp, accessNil, threading, fieldToken.txt, nil )
      else 
         
            if prefixExpType:equals( self.processInfo, Ast.builtinTypeStem ) then
               typeInfo = Ast.builtinTypeStem_
               if extFlag then
                  typeInfo = self:createExtType( fieldToken.pos, typeInfo )
               end
               
               return Nodes.ExpRefItemNode.create( self.nodeManager, fieldToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, prefixExp, accessNil, threading, fieldToken.txt, nil )
            else
             
               self:error( string.format( "illegal type -- %s, %s", prefixExpType:getTxt(  ), Ast.TypeInfoKind:_getTxt( prefixExpType:get_kind(  ))
               ) )
            end
            
      end
   end
   
   
   if not symbolInfo then
      local prefixScope = prefixExpType:get_scope()
      do
         local _exp = prefixScope
         if _exp ~= nil then
            symbolInfo = _exp:getSymbolInfoField( fieldToken.txt, true, self.scope, self.scopeAccess )
         end
      end
      
   end
   
   
   if symbolInfo ~= nil then
      if #prefixSymbolInfoList == 1 then
         local prefixSymbolInfo = prefixSymbolInfoList[1]
         if prefixSymbolInfo:get_kind() == Ast.SymbolKind.Typ then
            if prefixSymbolInfo:get_typeInfo():get_kind() ~= Ast.TypeInfoKind.Module and not symbolInfo:get_staticFlag() and symbolInfo:get_kind() ~= Ast.SymbolKind.Typ then
               self:addErrMess( fieldToken.pos, string.format( "Type can't access this symbol. -- %s", symbolInfo:get_name()) )
            end
            
         elseif symbolInfo:get_staticFlag() and symbolInfo:get_typeInfo():get_kind() ~= Ast.TypeInfoKind.Method then
            self:addErrMess( fieldToken.pos, string.format( "can't access this symbol. -- %s", fieldToken.txt) )
         end
         
      end
      
      
      
      if not Ast.TypeInfo.isMut( prefixExpType ) and not symbolInfo:get_staticFlag() and symbolInfo:get_kind() == Ast.SymbolKind.Mtd and symbolInfo:get_mutable() then
         self:addErrMess( fieldToken.pos, string.format( "can't access mutable method. -- %s.%s", prefixExpType:getTxt(  ), fieldToken.txt) )
      end
      
      
      if symbolInfo:get_typeInfo():get_mutMode() == Ast.MutMode.AllMut then
         local curType = self:getCurrentNamespaceTypeInfo(  )
         do
            local _switchExp = curType:get_kind()
            if _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Method then
               do
                  local _switchExp = curType:get_asyncMode()
                  if _switchExp == Ast.Async.Async or _switchExp == Ast.Async.Transient then
                     self:addErrMess( fieldToken.pos, string.format( "can't access allmut type's field(%s) in async function.", symbolInfo:get_name()) )
                  end
               end
               
            end
         end
         
      end
      
   end
   
   
   local accessSymbolInfo = nil
   local symbolMutMode = typeInfo:get_mutMode()
   if symbolInfo ~= nil then
      local workSymInfo = Ast.AccessSymbolInfo.new(self.processInfo, symbolInfo, _lune.newAlge( Ast.OverrideMut.Prefix, {prefixExpType}), not accessNil)
      if not getterTypeInfo then
         typeInfo = workSymInfo:get_typeInfo()
      end
      
      accessSymbolInfo = workSymInfo
      do
         local _switchExp = mode
         if _switchExp == ExpSymbolMode.Field or _switchExp == ExpSymbolMode.FieldNil then
            symbolMutMode = symbolInfo:get_mutMode()
         end
      end
      
   end
   
   
   if accessNil then
      if not typeInfo:get_nilable() then
         typeInfo = typeInfo:get_nilableTypeInfo()
      end
      
      self.helperInfo.useNilAccess = true
   end
   
   
   if not Ast.TypeInfo.isMut( prefixExpType ) then
      if self.ctrl_info.legacyMutableControl then
         if symbolMutMode == Ast.MutMode.Mut then
            typeInfo = self:createModifier( typeInfo, Ast.MutMode.IMut )
         end
         
      else
       
         if typeInfo:get_mutMode() == Ast.MutMode.Mut and symbolMutMode ~= Ast.MutMode.AllMut then
            typeInfo = self:createModifier( typeInfo, Ast.MutMode.IMut )
         end
         
      end
      
   end
   
   
   if typeInfo:equals( self.processInfo, builtinFunc.list_unpack ) or typeInfo:equals( self.processInfo, builtinFunc.array_unpack ) then
      self.helperInfo.useUnpack = true
   end
   
   
   do
      local expRef = _lune.__Cast( prefixExp, 3, Nodes.ExpRefNode )
      if expRef ~= nil then
         local prefixSym = expRef:get_symbolInfo()
         local prefixType = prefixSym:get_typeInfo()
         if prefixSym:get_kind() == Ast.SymbolKind.Typ and prefixType:get_kind() == Ast.TypeInfoKind.Class and #prefixType:get_itemTypeInfoList() > 0 and not Ast.isGenericType( prefixType ) and not self.scope:isInnerOf( _lune.unwrap( prefixType:get_scope()) ) then
            local accessErr = false
            if typeInfo:get_kind() == Ast.TypeInfoKind.Func then
               local altSet = {}
               for __index, argType in ipairs( typeInfo:get_argTypeInfoList() ) do
                  local orgType = argType:get_nonnilableType():get_srcTypeInfo()
                  if orgType:get_kind() == Ast.TypeInfoKind.Alternate then
                     altSet[orgType]= true
                  end
                  
               end
               
               for __index, itemType in ipairs( prefixType:get_itemTypeInfoList() ) do
                  if not _lune._Set_has(altSet, itemType:get_nonnilableType():get_srcTypeInfo() ) then
                     accessErr = true
                     break
                  end
                  
               end
               
            else
             
               accessErr = true
            end
            
            if accessErr then
               self:addErrMess( prefixExp:get_pos(), string.format( "can't access this class(%s) without '<>'.", prefixType:getTxt(  )) )
            end
            
         end
         
      end
   end
   
   
   if extFlag then
      typeInfo = self:createExtType( firstToken.pos, typeInfo )
      
      if self:checkThreading( firstToken.pos ) then
         self:addErrMess( firstToken.pos, "not support to use Luaval on thread." )
      end
      
   end
   
   
   do
      local _exp = getterTypeInfo
      if _exp ~= nil then
         return Nodes.GetFieldNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, fieldToken, accessSymbolInfo, accessNil, threading, prefixExp, _exp )
         
      else
         return Nodes.RefFieldNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, fieldToken, accessSymbolInfo, accessNil, threading, prefixExp )
      end
   end
   
end


function TransUnit:analyzeNewAlge( firstToken, algeTypeInfo, prefix )

   local symbolToken = self:getSymbolToken( SymbolMode.MustNot_ )
   
   self:checkAlgeComp( symbolToken, algeTypeInfo )
   
   do
      local valInfo = algeTypeInfo:getValInfo( symbolToken.txt )
      if valInfo ~= nil then
         local argList = {}
         local argListNode
         
         if #valInfo:get_typeList() > 0 then
            self:checkNextToken( "(" )
            argListNode = self:analyzeExpList( false, false, false, nil, valInfo:get_typeList(), nil )
            argList = (_lune.unwrap( argListNode) ):get_expList()
            self:checkNextToken( ")" )
         else
          
            argListNode = nil
         end
         
         
         do
            local _4312, _4313, newExpNodeList = self:checkMatchType( "call", symbolToken.pos, valInfo:get_typeList(), argListNode, false, true, nil )
            if newExpNodeList ~= nil then
               argList = newExpNodeList:get_expList()
            end
         end
         
         
         if algeTypeInfo:get_externalFlag() and not self.scope:getModuleInfo( algeTypeInfo:getModule(  ):get_srcTypeInfo() ) then
            local fullname = algeTypeInfo:getFullName( self.typeNameCtrl, self.scope, true )
            self:addErrMess( firstToken.pos, string.format( "This module not import -- %s", fullname) )
         end
         
         
         return Nodes.NewAlgeValNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {algeTypeInfo}, symbolToken, prefix, algeTypeInfo, valInfo, argList )
      else
         local dummySymbol = _lune.nilacc( algeTypeInfo:get_parentInfo():get_scope(), 'getSymbolInfoChild', 'callmtd' , algeTypeInfo:get_rawTxt() )
         self:addErrMess( symbolToken.pos, string.format( "not found Alge -- %s", symbolToken.txt) )
         return Nodes.NewAlgeValNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {algeTypeInfo}, symbolToken, prefix, algeTypeInfo, Ast.AlgeValInfo.new("", {}, algeTypeInfo, _lune.unwrap( dummySymbol)), {} )
      end
   end
   
end


function TransUnit:checkAsyncSymbol( symbolInfo, pos )

   local curNs = self:getCurrentNamespaceTypeInfo(  )
   local warn = false
   if curNs:get_asyncMode() == Ast.Async.Async and symbolInfo:get_name() ~= "self" then
      do
         local _switchExp = symbolInfo:get_kind()
         if _switchExp == Ast.SymbolKind.Arg or _switchExp == Ast.SymbolKind.Mbr or _switchExp == Ast.SymbolKind.Var then
            if not Ast.isPrimitive( symbolInfo:get_typeInfo() ) then
               warn = true
            end
            
         end
      end
      
   end
   
   if warn then
      do
         local _switchExp = curNs:get_kind()
         if _switchExp == Ast.TypeInfoKind.Func then
            if symbolInfo:get_namespaceTypeInfo() ~= curNs and curNs:get_asyncMode() ~= Ast.Async.Transient then
               if Ast.TypeInfo.isMut( symbolInfo:get_typeInfo() ) then
                  self:addErrMess( pos, string.format( "can't access the mutable type's symbol(%s) from async (%s).", symbolInfo:get_name(), symbolInfo:get_typeInfo():getTxt(  )) )
               end
               
            end
            
         elseif _switchExp == Ast.TypeInfoKind.Method then
            if symbolInfo:get_namespaceTypeInfo() ~= curNs then
               if Ast.TypeInfo.isMut( symbolInfo:get_typeInfo() ) then
                  self:addErrMess( pos, string.format( "can't access the mutable type's symbol(%s) from async (%s).", symbolInfo:get_name(), symbolInfo:get_typeInfo():getTxt(  )) )
               end
               
            end
            
         end
      end
      
   end
   
end


function TransUnit:analyzeExpSymbol( firstToken, symbolToken, mode, prefixExp, skipFlag, canLeftExp )

   local exp
   
   
   if mode == ExpSymbolMode.Field or mode == ExpSymbolMode.Get or mode == ExpSymbolMode.FieldNil or mode == ExpSymbolMode.GetNil then
      if prefixExp ~= nil then
         exp = self:analyzeExpField( firstToken, symbolToken, mode, prefixExp )
         
         local expType = exp:get_expType()
         if prefixExp:get_expType():isModule(  ) then
            do
               local algeType = _lune.__Cast( expType, 3, Ast.AlgeTypeInfo )
               if algeType ~= nil then
                  local nextToken = self:getToken(  )
                  if nextToken.txt == "." then
                     return self:analyzeNewAlge( firstToken, algeType, exp )
                  end
                  
                  self:pushback(  )
               end
            end
            
         end
         
      else
         Util.err( "prefixExp is nil" )
      end
      
   elseif mode == ExpSymbolMode.Symbol then
      if self.macroCtrl:get_analyzeInfo():isAnalyzingSymArg(  ) then
         exp = Nodes.LiteralSymbolNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeSymbol}, symbolToken )
      else
       
         self:checkSymbolComp( symbolToken )
         local symbolInfo = self.scope:getSymbolTypeInfo( symbolToken.txt, self.scope, self.moduleScope, self.scopeAccess )
         if  nil == symbolInfo then
            local _symbolInfo = symbolInfo
         
            if self.analyzeMode ~= AnalyzeMode.Diag then
               local work = self.scope
               while true do
                  print( work, self.globalScope, Ast.rootScope )
                  if work == work:get_parent() then
                     break
                  end
                  
                  work = work:get_parent()
               end
               
               
               self.scope:filterSymbolTypeInfo( self.scope, self.moduleScope, self.scopeAccess, function ( workSymbolInfo )
               
                  print( "sym", workSymbolInfo:get_name() )
                  return true
               end )
            end
            
            self:error( "not found type -- " .. symbolToken.txt )
         end
         
         
         self:accessSymbol( symbolInfo, canLeftExp )
         
         self:checkAsyncSymbol( symbolInfo, firstToken.pos )
         
         local typeInfo = symbolInfo:get_typeInfo()
         
         do
            local _switchExp = symbolInfo:get_kind()
            if _switchExp == Ast.SymbolKind.Typ then
               do
                  local algeType = _lune.__Cast( typeInfo, 3, Ast.AlgeTypeInfo )
                  if algeType ~= nil then
                     local nextToken = self:getToken(  )
                     if nextToken.txt == "." then
                        return self:analyzeNewAlge( firstToken, algeType, nil )
                     end
                     
                     self:pushback(  )
                  end
               end
               
            elseif _switchExp == Ast.SymbolKind.Var then
               self.tentativeSymbol:addAccessSym( symbolInfo )
               if not symbolInfo:get_hasValueFlag() then
                  local nsTypeInfo = self:getCurrentNamespaceTypeInfo(  )
                  if not symbolInfo:get_scope():isInnerOf( _lune.unwrap( nsTypeInfo:get_scope()) ) then
                     self.tentativeSymbol:addAccessSymPos( AccessSymPos.new(symbolInfo, firstToken.pos) )
                     
                  end
                  
               end
               
            end
         end
         
         
         if typeInfo:equals( self.processInfo, Ast.builtinTypeSymbol ) then
            skipFlag = true
         end
         
         if typeInfo:equals( self.processInfo, builtinFunc.lns__load ) then
            self.helperInfo.useLoad = true
         end
         
         
         do
            local _switchExp = symbolToken.txt
            if _switchExp == "__func__" then
               local funcTypeInfo = self:getCurrentNamespaceTypeInfo(  )
               self.has__func__Symbol[funcTypeInfo]= true
            elseif _switchExp == "_G" or _switchExp == "_ENV" then
               local valid = false
               for pragma, __val in pairs( self.helperInfo.pragmaSet ) do
                  do
                     local _matchExp = pragma
                     if _matchExp[1] == LuneControl.Pragma.limit_conv_code[1] then
                        local codeSet = _matchExp[2][1]
                     
                        if _lune._Set_len(codeSet ) == 1 and _lune._Set_has(codeSet, LuneControl.Code.Lua ) then
                           valid = true
                           break
                        end
                        
                     end
                  end
                  
               end
               
               if not valid then
                  self:addErrMess( firstToken.pos, "'_G' and '_ENV' only can access with transcompiling to lua." )
               end
               
            elseif _switchExp == "_" then
               if not canLeftExp then
                  self:addErrMess( firstToken.pos, "It can't access the symbol '_'." )
               end
               
            end
         end
         
         
         if typeInfo:get_nonnilableType():get_kind() == Ast.TypeInfoKind.Ext and self:checkThreading( symbolToken.pos ) then
            self:addErrMess( symbolToken.pos, "not support to use Luaval on thread." )
         end
         
         
         exp = Nodes.ExpRefNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, Ast.AccessSymbolInfo.new(self.processInfo, symbolInfo, _lune.newAlge( Ast.OverrideMut.None), true) )
      end
      
   elseif mode == ExpSymbolMode.Fn then
      exp = self:analyzeDeclFunc( DeclFuncMode.Func, false, false, Ast.AccessMode.Local, false, nil, symbolToken, nil )
   else
    
      self:error( string.format( "illegal mode -- %s", mode) )
   end
   
   
   if self.analyzeMode == AnalyzeMode.Inquire and self:isTargetToken( symbolToken ) then
      local accessMode = Ast.AccessMode.None
      for __index, symbolInfo in ipairs( exp:getSymbolInfo(  ) ) do
         accessMode = symbolInfo:get_accessMode()
      end
      
      self:dumpSymbolType( accessMode, symbolToken.txt, exp:get_expType() )
   end
   
   
   return self:analyzeExpCont( firstToken, exp, skipFlag, canLeftExp )
end

function TransUnit:analyzeExpOpSet( exp, opeToken, expectTypeList )

   local _
   exp:setLValue(  )
   
   if not exp:canBeLeft(  ) then
      self:addErrMess( exp:get_pos(), string.format( "this node can not be l-value. -- %s", Nodes.getNodeKindName( exp:get_kind() )) )
   end
   
   
   local function process( lValNode )
   
      local refItemNode = _lune.__Cast( lValNode, 3, Nodes.ExpRefItemNode )
      if  nil == refItemNode then
         local _refItemNode = refItemNode
      
         return nil
      end
      
      do
         local _switchExp = refItemNode:get_val():get_expType():get_kind()
         if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Map then
            return refItemNode
         end
      end
      
      return nil
   end
   
   local listRefItemNode = nil
   do
      local symNodeList = _lune.__Cast( exp, 3, Nodes.ExpListNode )
      if symNodeList ~= nil then
         for __index, symNode in ipairs( symNodeList:get_expList() ) do
            do
               local refItemNode = process( symNode )
               if refItemNode ~= nil then
                  listRefItemNode = refItemNode
                  if #symNodeList:get_expList() > 1 then
                     self:addErrMess( refItemNode:get_pos(), "When left-value includes 'list[i]', left-value must be single." )
                  end
                  
                  break
               end
            end
            
         end
         
      else
         listRefItemNode = process( exp )
      end
   end
   
   
   local expList = self:analyzeExpList( false, false, false, nil, expectTypeList )
   for index, expType in ipairs( expList:get_expTypeList() ) do
      if expType:get_asyncMode() == Ast.Async.Transient then
         self:addErrMess( expList:get_pos(), string.format( "can't set the __trans type -- index:%d, %s", index, expType:getTxt(  )) )
      end
      
   end
   
   
   if not expList:canBeRight( self.processInfo ) then
      self:addErrMess( expList:get_pos(), string.format( "this node can not be r-value. -- %s", Nodes.getNodeKindName( expList:get_kind() )) )
   end
   
   
   local _4434, _4435, workList, expTypeList = self:checkMatchType( "= operator", opeToken.pos, exp:get_expTypeList(), expList, true, false, nil )
   if workList ~= nil then
      expList = workList
   end
   
   
   local initSymSet = {}
   local symbolList = {}
   
   for index, symbolInfo in ipairs( exp:getSymbolInfo(  ) ) do
      table.insert( symbolList, symbolInfo )
      if not symbolInfo:get_mutable() and symbolInfo:get_hasValueFlag() then
         if self.validMutControl then
            self:addErrMess( opeToken.pos, string.format( "this is not mutable variable. -- %s", symbolInfo:get_name()) )
         end
         
      end
      
      self.tentativeSymbol:modSym( symbolInfo )
      if index <= #expTypeList and not symbolInfo:get_hasValueFlag() then
         do
            local _switchExp = symbolInfo:get_kind()
            if _switchExp == Ast.SymbolKind.Var then
               if symbolInfo:get_typeInfo() == Ast.builtinTypeEmpty then
                  local expType = expTypeList[index]
                  do
                     local _switchExp = expType:get_kind()
                     if _switchExp == Ast.TypeInfoKind.DDD then
                        if #expType:get_itemTypeInfoList() > 0 then
                           expType = expType:get_itemTypeInfoList()[1]:get_nilableTypeInfo()
                        end
                        
                     elseif _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array or _switchExp == Ast.TypeInfoKind.Set or _switchExp == Ast.TypeInfoKind.Map then
                        local workPos
                        
                        if index <= #expList:get_expList() then
                           workPos = expList:get_expList()[index]:get_pos()
                        else
                         
                           workPos = opeToken.pos
                        end
                        
                        self:checkLiteralEmptyCollection( workPos, symbolInfo:get_name(), expType )
                     end
                  end
                  
                  symbolInfo:set_typeInfo( expType )
               end
               
               if not self.tentativeSymbol:regist( symbolInfo, exp:get_pos() ) then
                  self:addErrMess( opeToken.pos, string.format( "can't access in this scope. -- %s", symbolInfo:get_name()) )
               end
               
               initSymSet[symbolInfo]= true
            elseif _switchExp == Ast.SymbolKind.Mbr then
               initSymSet[symbolInfo]= true
            end
         end
         
      end
      
      if symbolInfo:get_kind() ~= Ast.SymbolKind.Var or self.scope:getNamespaceTypeInfo(  ) == symbolInfo:get_scope():getNamespaceTypeInfo(  ) then
         symbolInfo:updateValue( exp:get_pos() )
      end
      
   end
   
   
   if listRefItemNode ~= nil then
      local index
      
      do
         local indexNode = listRefItemNode:get_index()
         if indexNode ~= nil then
            index = _lune.newAlge( Nodes.IndexVal.NodeIdx, {indexNode})
         else
            index = _lune.newAlge( Nodes.IndexVal.SymIdx, {_lune.unwrap( listRefItemNode:get_symbol())})
         end
      end
      
      return Nodes.ExpSetItemNode.create( self.nodeManager, exp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, listRefItemNode:get_val(), index, expList )
   end
   
   
   return Nodes.ExpSetValNode.create( self.nodeManager, exp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, exp, expList, symbolList, initSymSet )
end


function TransUnit:analyzeExpOpEquals( pos, opToken, exp1, exp2 )

   local exp1Type = exp1:get_expType()
   local exp2Type = exp2:get_expType()
   
   if (not exp1Type:canEvalWith( self.processInfo, exp2Type, Ast.CanEvalType.SetOp, {} ) and not exp2Type:canEvalWith( self.processInfo, exp1Type, Ast.CanEvalType.SetOp, {} ) ) then
      self:addErrMess( opToken.pos, string.format( "not compatible type '%s' or '%s'", exp1Type:getTxt( self.typeNameCtrl ), exp2Type:getTxt( self.typeNameCtrl )) )
      return exp1, exp2
   end
   
   do
      local _exp = _lune.__Cast( exp1, 3, Nodes.NewAlgeValNode )
      if _exp ~= nil then
         if #_exp:get_paramList() > 0 then
            self:addErrMess( exp1:get_pos(), "can't compare alge." )
            return exp1, exp2
         end
         
      end
   end
   
   do
      local _exp = _lune.__Cast( exp2, 3, Nodes.NewAlgeValNode )
      if _exp ~= nil then
         if #_exp:get_paramList() > 0 then
            self:addErrMess( exp2:get_pos(), "can't compare alge." )
            return exp1, exp2
         end
         
      end
   end
   
   if exp1Type:equals( self.processInfo, Ast.builtinTypeBool ) and exp2Type:equals( self.processInfo, Ast.builtinTypeBool ) and (exp1:get_kind() == Nodes.NodeKind.get_LiteralBool() or exp2:get_kind() == Nodes.NodeKind.get_LiteralBool() ) then
      self:addWarnMess( exp1:get_pos(), "this operation is deprecated." )
      return exp1, exp2
   end
   
   
   local function getType( typeInfo )
   
      local workType = typeInfo:get_nonnilableType():get_srcTypeInfo()
      if workType:get_kind() == Ast.TypeInfoKind.Alternate and workType:hasBase(  ) then
         return workType:get_baseTypeInfo()
      end
      
      return workType
   end
   
   local nonNilType1 = getType( exp1Type )
   local nonNilType2 = getType( exp2Type )
   
   if nonNilType1 ~= nonNilType2 then
      if nonNilType1:get_kind() == Ast.TypeInfoKind.Class or nonNilType1:get_kind() == Ast.TypeInfoKind.IF then
         if nonNilType2:get_kind() == Ast.TypeInfoKind.Class or nonNilType2:get_kind() == Ast.TypeInfoKind.IF then
            if nonNilType1:isInheritFrom( self.processInfo, nonNilType2 ) then
               exp1 = Nodes.ExpCastNode.create( self.nodeManager, exp1:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {exp2Type:get_nonnilableType()}, exp1, exp2Type, Nodes.CastKind.Implicit )
            else
             
               exp2 = Nodes.ExpCastNode.create( self.nodeManager, exp2:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {exp1Type:get_nonnilableType()}, exp2, exp1Type, Nodes.CastKind.Implicit )
            end
            
         else
          
            exp1 = Nodes.ExpCastNode.create( self.nodeManager, exp1:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {exp1Type}, exp1, Ast.builtinTypeStem, Nodes.CastKind.Implicit )
         end
         
      else
       
         if nonNilType2:get_kind() == Ast.TypeInfoKind.Class or nonNilType2:get_kind() == Ast.TypeInfoKind.IF then
            exp2 = Nodes.ExpCastNode.create( self.nodeManager, exp2:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {exp2Type}, exp2, Ast.builtinTypeStem, Nodes.CastKind.Implicit )
         end
         
      end
      
   end
   
   
   return exp1, exp2
end


function TransUnit:analyzeExpOp2( firstToken, exp, prevOpLevel )

   while true do
      local opToken = self:getTokenNoErr(  )
      local opTxt = opToken.txt
      
      if opToken.txt == "@@" or opToken.txt == "@@@" or opToken.txt == "@@=" then
         exp = self:analyzeExpCast( firstToken, opTxt, exp )
      elseif opToken.kind == Parser.TokenKind.Ope then
         if Parser.isOp2( opTxt ) then
            if opTxt ~= "=" and not exp:canBeRight( self.processInfo ) then
               self:addErrMess( exp:get_pos(), string.format( "This can't evaluate for '%s' -- %s", opTxt, Nodes.getNodeKindName( exp:get_kind() )) )
            end
            
            local opLevel = op2levelMap[opTxt]
            if  nil == opLevel then
               local _opLevel = opLevel
            
               self:error( string.format( "unknown op -- %s %s", opTxt, prevOpLevel) )
            end
            
            do
               local _exp = prevOpLevel
               if _exp ~= nil then
                  if opLevel <= _exp then
                     self:pushback(  )
                     return exp
                  end
                  
               end
            end
            
            
            local expectTypeList = {}
            for __index, exp1Type in ipairs( exp:get_expTypeList() ) do
               local prefixExpType = exp1Type
               if prefixExpType:get_nilable() then
                  prefixExpType = prefixExpType:get_nonnilableType()
               end
               
               
               local expectType = Ast.builtinTypeNone
               do
                  local _exp = _lune.__Cast( prefixExpType:get_srcTypeInfo():get_aliasSrc(), 3, Ast.EnumTypeInfo )
                  if _exp ~= nil then
                     expectType = _exp
                  end
               end
               
               do
                  local _exp = _lune.__Cast( prefixExpType:get_srcTypeInfo(), 3, Ast.AlgeTypeInfo )
                  if _exp ~= nil then
                     expectType = _exp
                  end
               end
               
               table.insert( expectTypeList, expectType )
            end
            
            if #expectTypeList == 0 then
               self:addErrMess( exp:get_pos(), string.format( "This exp have no value -- %s", Nodes.getNodeKindName( exp:get_kind() )) )
               table.insert( expectTypeList, Ast.builtinTypeNone )
            end
            
            
            if opTxt == "=" then
               return self:analyzeExpOpSet( exp, opToken, expectTypeList )
            end
            
            
            exp = self:MultiTo1( exp )
            
            local exp2 = self:analyzeExp( false, false, false, opLevel, expectTypeList[1] )
            exp2 = self:MultiTo1( exp2 )
            
            if not exp2:canBeRight( self.processInfo ) then
               self:addErrMess( exp2:get_pos(), string.format( "This can't evaluate for '%s' -- %s", opTxt, Nodes.getNodeKindName( exp2:get_kind() )) )
            end
            
            
            local retType = Ast.builtinTypeNone
            
            if not exp2:canBeRight( self.processInfo ) then
               self:addErrMess( exp2:get_pos(), string.format( "this node can not be r-value. -- %s", Nodes.getNodeKindName( exp2:get_kind() )) )
            end
            
            
            local exp1Type = exp:get_expType()
            local exp2Type = exp2:get_expType()
            
            if exp1Type:get_kind() == Ast.TypeInfoKind.DDD then
               do
                  local dddType = _lune.__Cast( exp1Type, 3, Ast.DDDTypeInfo )
                  if dddType ~= nil then
                     exp = Nodes.ExpMultiTo1Node.create( self.nodeManager, exp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {dddType:get_typeInfo():get_nilableTypeInfo()}, exp )
                  end
               end
               
            end
            
            if exp2Type:get_kind() == Ast.TypeInfoKind.DDD then
               do
                  local dddType = _lune.__Cast( exp2Type, 3, Ast.DDDTypeInfo )
                  if dddType ~= nil then
                     exp2 = Nodes.ExpMultiTo1Node.create( self.nodeManager, exp2:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {dddType:get_typeInfo():get_nilableTypeInfo()}, exp2 )
                  end
               end
               
            end
            
            
            do
               local _switchExp = opTxt
               if _switchExp == "or" then
                  local is3op
                  
                  do
                     local opExpType = _lune.__Cast( exp1Type, 3, Ast.AndExpTypeInfo )
                     if opExpType ~= nil then
                        exp1Type = opExpType:get_exp2()
                        is3op = true
                     else
                        is3op = false
                     end
                  end
                  
                  
                  if not exp1Type:equals( self.processInfo, Ast.builtinTypeBool ) and not exp1Type:equals( self.processInfo, Ast.builtinTypeStem ) and not exp1Type:get_nilable() then
                     if _lune.nilacc( _lune.nilacc( (_lune.__Cast( exp, 3, Nodes.ExpOp2Node ) ), 'get_op', 'callmtd' ), "txt" ) == "and" then
                        
                     else
                      
                        self:addWarnMess( exp:get_pos(), string.format( "this value never be 'false' -- %s", exp1Type:getTxt(  )) )
                     end
                     
                  end
                  
                  
                  if exp1Type:equals( self.processInfo, exp2Type ) then
                     retType = exp1Type
                  elseif exp1Type:canEvalWith( self.processInfo, exp2Type, Ast.CanEvalType.SetOp, {} ) then
                     retType = exp1Type
                  elseif exp2Type:canEvalWith( self.processInfo, exp1Type, Ast.CanEvalType.SetOp, {} ) then
                     retType = exp2Type
                  elseif exp2Type:equals( self.processInfo, Ast.builtinTypeNil ) then
                     if is3op or exp1Type:equals( self.processInfo, Ast.builtinTypeBool ) then
                        retType = exp1Type:get_nilableTypeInfo()
                     else
                      
                        retType = exp1Type
                     end
                     
                  elseif exp1Type:equals( self.processInfo, Ast.builtinTypeNil ) then
                     retType = exp2Type
                  else
                   
                     if exp1Type:get_nilable() and exp2Type:get_nilable() then
                        retType = Ast.builtinTypeStem_
                     elseif exp2Type:get_nilable() then
                        retType = Ast.builtinTypeStem_
                     elseif exp1Type:get_nilable() then
                        retType = Ast.builtinTypeStem
                     else
                      
                        retType = Ast.builtinTypeStem
                     end
                     
                  end
                  
                  if retType:get_nilable() and not exp2Type:get_nilable() then
                     retType = retType:get_nonnilableType()
                  end
                  
               elseif _switchExp == "and" then
                  local workToken = self:getToken(  )
                  self:pushback(  )
                  
                  if not exp1Type:equals( self.processInfo, Ast.builtinTypeBool ) and not exp1Type:equals( self.processInfo, Ast.builtinTypeStem ) and not exp1Type:get_nilable() then
                     self:addWarnMess( exp:get_pos(), "this value never be 'false'" )
                  elseif exp2:get_kind() == Nodes.NodeKind.get_LiteralBool() then
                     do
                        local literal = exp2:getLiteral(  )
                        if literal ~= nil then
                           if not Nodes.getLiteralObj( literal ) then
                              self:addErrMess( exp2:get_pos(), "this value never be 'true'" )
                           end
                           
                        end
                     end
                     
                  end
                  
                  if exp1Type:get_nilable() then
                     if exp2Type:get_nilable() then
                        retType = exp2Type
                     else
                      
                        retType = exp2Type:get_nilableTypeInfo()
                     end
                     
                  elseif exp1Type:equals( self.processInfo, Ast.builtinTypeBool ) or exp2Type:equals( self.processInfo, Ast.builtinTypeBool ) then
                     if exp1Type:canEvalWith( self.processInfo, exp2Type, Ast.CanEvalType.SetOp, {} ) then
                        retType = exp1Type
                     elseif exp2Type:canEvalWith( self.processInfo, exp1Type, Ast.CanEvalType.SetOp, {} ) then
                        retType = exp2Type
                     else
                      
                        if exp2Type:get_nilable() then
                           retType = Ast.builtinTypeStem_
                        else
                         
                           retType = Ast.builtinTypeStem
                        end
                        
                     end
                     
                  elseif exp1Type:equals( self.processInfo, Ast.builtinTypeStem ) then
                     retType = Ast.builtinTypeStem
                  else
                   
                     retType = exp2Type
                  end
                  
                  retType = Ast.AndExpTypeInfo.new(self.processInfo, exp1Type, exp2Type, retType)
               elseif _switchExp == "<" or _switchExp == ">" or _switchExp == "<=" or _switchExp == ">=" then
                  if Ast.builtinTypeString:canEvalWith( self.processInfo, exp1Type, Ast.CanEvalType.SetOp, {} ) and Ast.builtinTypeString:canEvalWith( self.processInfo, exp2Type, Ast.CanEvalType.SetOp, {} ) or (Ast.builtinTypeInt:canEvalWith( self.processInfo, exp1Type, Ast.CanEvalType.Comp, {} ) or Ast.builtinTypeReal:canEvalWith( self.processInfo, exp1Type, Ast.CanEvalType.Comp, {} ) ) and (Ast.builtinTypeInt:canEvalWith( self.processInfo, exp2Type, Ast.CanEvalType.Comp, {} ) or Ast.builtinTypeReal:canEvalWith( self.processInfo, exp2Type, Ast.CanEvalType.Comp, {} ) ) then
                     
                  else
                   
                     self:addErrMess( opToken.pos, string.format( "no numeric type '%s' or '%s'", exp1Type:getTxt( self.typeNameCtrl ), exp2Type:getTxt( self.typeNameCtrl )) )
                  end
                  
                  retType = Ast.builtinTypeBool
               elseif _switchExp == "~=" or _switchExp == "==" then
                  exp, exp2 = self:analyzeExpOpEquals( firstToken.pos, opToken, exp, exp2 )
                  retType = Ast.builtinTypeBool
               elseif _switchExp == "^" or _switchExp == "|" or _switchExp == "~" or _switchExp == "&" or _switchExp == "|<<" or _switchExp == "|>>" then
                  if self.targetLuaVer:get_hasBitOp() == LuaVer.BitOp.Cant then
                     self:addErrMess( opToken.pos, "this lua version can't use bit operand." )
                  end
                  
                  
                  if not Ast.builtinTypeInt:canEvalWith( self.processInfo, exp1Type, Ast.CanEvalType.Logical, {} ) or not Ast.builtinTypeInt:canEvalWith( self.processInfo, exp2Type, Ast.CanEvalType.Logical, {} ) then
                     self:addErrMess( opToken.pos, string.format( "no int type '%s' or '%s'", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
                  end
                  
                  retType = Ast.builtinTypeInt
               elseif _switchExp == ".." then
                  if not exp1Type:equals( self.processInfo, Ast.builtinTypeString ) or not exp2Type:equals( self.processInfo, Ast.builtinTypeString ) then
                     self:addErrMess( opToken.pos, string.format( "no string type '%s' or '%s'", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
                  end
                  
                  retType = Ast.builtinTypeString
               elseif _switchExp == "+" or _switchExp == "-" or _switchExp == "*" or _switchExp == "/" or _switchExp == "%" then
                  if (not Ast.builtinTypeInt:canEvalWith( self.processInfo, exp1Type, Ast.CanEvalType.Math, {} ) and not Ast.builtinTypeReal:canEvalWith( self.processInfo, exp1Type, Ast.CanEvalType.Math, {} ) ) or (not Ast.builtinTypeInt:canEvalWith( self.processInfo, exp2Type, Ast.CanEvalType.Math, {} ) and not Ast.builtinTypeReal:canEvalWith( self.processInfo, exp2Type, Ast.CanEvalType.Math, {} ) ) then
                     self:addErrMess( opToken.pos, string.format( "no numeric type '%s' or '%s'", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
                  end
                  
                  
                  if exp1Type:equals( self.processInfo, Ast.builtinTypeReal ) or exp2Type:equals( self.processInfo, Ast.builtinTypeReal ) then
                     retType = Ast.builtinTypeReal
                     if exp1Type:equals( self.processInfo, Ast.builtinTypeInt ) then
                        exp = Nodes.ExpCastNode.create( self.nodeManager, exp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeReal}, exp, Ast.builtinTypeReal, Nodes.CastKind.Implicit )
                     elseif exp2Type:equals( self.processInfo, Ast.builtinTypeInt ) then
                        exp2 = Nodes.ExpCastNode.create( self.nodeManager, exp2:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeReal}, exp2, Ast.builtinTypeReal, Nodes.CastKind.Implicit )
                     end
                     
                  else
                   
                     retType = Ast.builtinTypeInt
                  end
                  
               else 
                  
                     self:error( "unknown op " .. opTxt )
               end
            end
            
            
            local threading
            
            do
               local _switchExp = opTxt
               if _switchExp == "and" or _switchExp == "or" then
                  threading = self:checkThreading( firstToken.pos )
               else 
                  
                     threading = false
               end
            end
            
            
            exp = Nodes.ExpOp2Node.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {retType}, opToken, threading, self.nodeManager:MultiTo1( exp ), self.nodeManager:MultiTo1( exp2 ) )
         else
          
            self:error( "illegal op" )
         end
         
      else
       
         self:pushback(  )
         return exp
      end
      
   end
   
end


function TransUnit:analyzeExpMacroStat( firstToken )

   local expStrList = {}
   
   self:checkNextToken( "{" )
   
   local braceCount = 0
   local prevToken = firstToken
   
   local errMessList = {}
   
   while true do
      local token = self:getToken(  )
      
      if token.txt == ",," or token.txt == ",,," or token.txt == ",,,," then
         local exp = self:analyzeExp( false, true, false, _lune.unwrap( op1levelMap[token.txt]) )
         
         local literalStr = self.macroCtrl:expandSymbol( self, token, exp, self.nodeManager, errMessList )
         for __index, errMess in ipairs( errMessList ) do
            self:addErrMess( errMess.pos, errMess.mess )
         end
         
         
         table.insert( expStrList, literalStr )
      else
       
         if token.txt == "{" then
            braceCount = braceCount + 1
         elseif token.txt == "}" then
            if braceCount == 0 then
               break
            end
            
            braceCount = braceCount - 1
         end
         
         
         local format = "' %s'"
         local consecutive
         
         if prevToken == firstToken or token.consecutive then
            format = "'%s'"
            consecutive = true
         else
          
            consecutive = false
         end
         
         local newToken = Parser.Token.new(token.kind, string.format( format, token.txt ), token.pos, consecutive)
         local literalStr = Nodes.LiteralStringNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeString}, newToken, nil, nil, false )
         table.insert( expStrList, literalStr )
      end
      
      prevToken = token
   end
   
   
   return Nodes.ExpMacroStatNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeStat}, expStrList )
end


function TransUnit:analyzeSuper( firstToken )

   
   self:checkNextToken( "(" )
   
   local nextToken = self:getToken(  )
   local expList = nil
   if nextToken.txt ~= ")" then
      self:pushback(  )
      expList = self:analyzeExpList( false, false, false )
      self:checkNextToken( ")" )
   end
   
   
   self:checkNextToken( ";" )
   
   local classType = self:getCurrentClass(  )
   
   local currentFunc = self:getCurrentNamespaceTypeInfo(  )
   if currentFunc:get_kind() == Ast.TypeInfoKind.Method then
      local superType = classType:get_baseTypeInfo(  )
      if superType:equals( self.processInfo, Ast.headTypeInfo ) then
         self:addErrMess( firstToken.pos, "This class doesn't have super-class." )
      else
       
         if currentFunc:get_rawTxt() == "__init" then
            local superScope = superType:get_scope()
            if  nil == superScope then
               local _superScope = superScope
            
               self:error( "not found super scope" )
            end
            
            local superCtorType = superScope:getTypeInfoChild( "__init" )
            if  nil == superCtorType then
               local _superCtorType = superCtorType
            
               self:error( "not found super '__init'" )
            end
            
            self:checkMatchValType( firstToken.pos, superCtorType, expList, {}, classType )
            return Nodes.ExpCallSuperCtorNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, superType, superCtorType, expList )
         else
          
            do
               local superFunc = (_lune.unwrap( superType:get_scope()) ):getTypeInfoField( currentFunc:get_rawTxt(), true, self.scope, self.scopeAccess )
               if superFunc ~= nil then
                  if superFunc:get_abstractFlag() then
                     self:addErrMess( firstToken.pos, "super is abstract." )
                  end
                  
                  self:checkMatchValType( firstToken.pos, superFunc, expList, {}, classType )
                  local exp = Nodes.ExpCallSuperNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, superType, superFunc, expList )
                  return Nodes.StmtExpNode.create( self.nodeManager, exp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), exp:get_expTypeList(), exp )
               end
            end
            
            self:addErrMess( firstToken.pos, "this is not override method." )
            return self:createNoneNode( firstToken.pos )
         end
         
      end
      
   end
   
   
   self:addErrMess( firstToken.pos, "super can't call here." )
   return self:createNoneNode( firstToken.pos )
end


function TransUnit:analyzeUnwrap( firstToken )

   local nextToken, continueFlag = self:getContinueToken(  )
   
   if not continueFlag or nextToken.txt ~= "!" then
      self:pushback(  )
      self:pushbackToken( firstToken )
      local exp = self:analyzeExp( false, false, false )
      self:checkNextToken( ";" )
      if not exp:get_expType():get_nilable() then
         self:addErrMess( exp:get_pos(), "this value is not nilable." )
      end
      
      return Nodes.StmtExpNode.create( self.nodeManager, nextToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, exp )
   end
   
   
   self:pushback(  )
   return self:analyzeDeclVar( Nodes.DeclVarMode.Unwrap, Ast.AccessMode.Local, firstToken )
end


function TransUnit:analyzeExpUnwrap( firstToken )

   local expNode = self:analyzeExpOneRVal( false, true )
   local nextToken = self:getToken(  )
   local insNode = nil
   if nextToken.txt == "default" then
      insNode = self:analyzeExpOneRVal( false, false )
   else
    
      self:pushback(  )
   end
   
   
   local unwrapType = Ast.builtinTypeStem_
   
   local expType = expNode:get_expType()
   if not expType:get_nilable() then
      unwrapType = expType
      self:addErrMess( expNode:get_pos(), string.format( "this exp is not nilable -- %s", expType:getTxt(  )) )
   elseif expType:get_kind() == Ast.TypeInfoKind.DDD then
      if #expType:get_itemTypeInfoList() > 0 then
         unwrapType = expType:get_itemTypeInfoList()[1]:get_nonnilableType()
      else
       
         unwrapType = Ast.builtinTypeStem
      end
      
   else
    
      unwrapType = expType:get_nonnilableType()
   end
   
   
   if insNode ~= nil then
      local insType = insNode:get_expType()
      
      if insType:get_nilable() then
         self:addErrMess( insNode:get_pos(), string.format( "default can't use nilable -- %s", insType:getTxt(  )) )
      end
      
      
      local alt2type = Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
      if not unwrapType:canEvalWith( self.processInfo, insType, Ast.CanEvalType.SetOp, alt2type ) then
         if not insType:canEvalWith( self.processInfo, unwrapType, Ast.CanEvalType.SetOp, alt2type ) then
            unwrapType = Ast.builtinTypeStem
            
         else
          
            unwrapType = insType
         end
         
      end
      
   end
   
   
   self.helperInfo.useUnwrapExp = true
   
   if Ast.isExtType( expType:get_nonnilableType() ) then
      do
         local _matchExp = self.processInfo:createLuaval( unwrapType, false )
         if _matchExp[1] == Ast.LuavalResult.OK[1] then
            local work = _matchExp[2][1]
            local _ = _matchExp[2][2]
         
            unwrapType = work
         elseif _matchExp[1] == Ast.LuavalResult.Err[1] then
            local err = _matchExp[2][1]
         
            self:addErrMess( firstToken.pos, err )
         end
      end
      
   end
   
   
   return Nodes.ExpUnwrapNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {unwrapType}, expNode, insNode )
end

function TransUnit:analyzeStrConst( firstToken, token )

   local exp
   
   
   local nextToken = self:getToken( true )
   if nextToken.kind ~= Parser.TokenKind.Eof then
      local param
      
      local dddParam
      
      if nextToken.txt == "(" then
         local _
         local argNodeList = self:analyzeExpList( false, false, false )
         param = argNodeList
         
         local _4710, _4711, workExpList = self:checkMatchType( "str constructor", firstToken.pos, {Ast.builtinTypeDDD}, argNodeList, false, false, nil )
         if workExpList ~= nil then
            dddParam = workExpList
         else
            dddParam = nil
         end
         
         
         self:checkNextToken( ")" )
         nextToken = self:getToken( true )
         
         if param ~= nil then
            self:checkStringFormat( token.pos, token.txt, param:get_expTypeList() )
         end
         
      else
       
         param = nil
         dddParam = nil
      end
      
      
      local threading
      
      if param then
         threading = self:checkThreading( firstToken.pos )
      else
       
         threading = false
      end
      
      
      local workExp = Nodes.LiteralStringNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeString}, token, param, dddParam, threading )
      if nextToken.txt == "[" or nextToken.txt == "$[" then
         exp = self:analyzeExpRefItem( nextToken, workExp, nextToken.txt == "$[" )
      else
       
         exp = workExp
         if nextToken.kind ~= Parser.TokenKind.Eof then
            self:pushback(  )
         end
         
      end
      
   else
    
      exp = Nodes.LiteralStringNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeString}, token, nil, nil, false )
   end
   
   return exp
end


function TransUnit:analyzeExp( allowNoneType, skipOp2Flag, canLeftExp, prevOpLevel, expectType )

   local firstToken = self:getToken(  )
   
   local function processsExpectExp( token, orgExpectType )
   
      do
         local enumTypeInfo = _lune.__Cast( orgExpectType:get_srcTypeInfo():get_aliasSrc(), 3, Ast.EnumTypeInfo )
         if enumTypeInfo ~= nil then
            local nextToken = self:getToken(  )
            self:checkEnumComp( nextToken, enumTypeInfo )
            
            do
               local valInfo = enumTypeInfo:getEnumValInfo( nextToken.txt )
               if valInfo ~= nil then
                  
                  local aliasType = nil
                  local expType = enumTypeInfo
                  
                  aliasType = self.importedAliasMap[enumTypeInfo]
                  if aliasType ~= nil then
                     expType = aliasType
                  end
                  
                  if not self.moduleType:equals( self.processInfo, orgExpectType:getModule(  ) ) then
                     if not _lune._Set_has(self.importModuleSet, orgExpectType:getModule(  ) ) then
                        if not aliasType then
                           local fullname = orgExpectType:getFullName( self.typeNameCtrl, self.scope, true )
                           self:addErrMess( token.pos, string.format( "need to import module -- %s (%s)", fullname, enumTypeInfo:getTxt(  )) )
                        end
                        
                     end
                     
                  end
                  
                  local exp = Nodes.ExpOmitEnumNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {expType}, nextToken, valInfo, aliasType, enumTypeInfo )
                  return self:analyzeExpCont( firstToken, exp, false, canLeftExp )
               end
            end
            
            
            self:error( string.format( "illegal enum val -- %s.%s", orgExpectType:getTxt(  ), nextToken.txt) )
         end
      end
      
      do
         local algeTyepInfo = _lune.__Cast( orgExpectType:get_srcTypeInfo(), 3, Ast.AlgeTypeInfo )
         if algeTyepInfo ~= nil then
            return self:analyzeNewAlge( firstToken, algeTyepInfo, nil )
         end
      end
      
      
      self:error( string.format( "illegal type for '.' -- %s", orgExpectType:getTxt(  )) )
   end
   
   local function processsNewExp( token )
   
      local _
      local exp = self:analyzeRefType( Ast.AccessMode.Local, false, false )
      
      local classTypeInfo = exp:get_expType()
      do
         local _switchExp = classTypeInfo:get_kind()
         if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
            if classTypeInfo:equals( self.processInfo, Ast.builtinTypeString ) then
               self:error( string.format( "'new' can't use this type -- %s", classTypeInfo:getTxt(  )) )
            end
            
         else 
            
               self:error( string.format( "'new' can't use this type -- %s", classTypeInfo:getTxt(  )) )
         end
      end
      
      
      if classTypeInfo:get_externalFlag() then
         do
            local _switchExp = classTypeInfo:get_accessMode()
            if _switchExp == Ast.AccessMode.Pri or _switchExp == Ast.AccessMode.Local then
               self:addErrMess( token.pos, string.format( "Can't access -- %s", Ast.AccessMode:_getTxt( classTypeInfo:get_accessMode())
               ) )
            end
         end
         
      end
      
      if classTypeInfo:get_abstractFlag() then
         self:addErrMess( token.pos, "abstract class can't new" )
      end
      
      
      local classScope = classTypeInfo:get_scope(  )
      local initTypeInfo = (_lune.unwrap( classScope) ):getTypeInfoChild( "__init" )
      if  nil == initTypeInfo then
         local _initTypeInfo = initTypeInfo
      
         self:error( "not found __init" )
      end
      
      
      self:checkCallAsync( initTypeInfo, token.pos )
      
      self:checkNextToken( "(" )
      local nextToken = self:getToken(  )
      local argList = nil
      
      if nextToken.txt ~= ")" then
         self:pushback(  )
         argList = self:analyzeExpList( false, false, false, nil, initTypeInfo:get_argTypeInfoList() )
         self:checkNextToken( ")" )
      end
      
      
      if initTypeInfo:get_accessMode() == Ast.AccessMode.Pub or (initTypeInfo:get_accessMode() == Ast.AccessMode.Pro and self.scope:getClassTypeInfo(  ):isInheritFrom( self.processInfo, classTypeInfo, nil ) ) or (self.scope:getClassTypeInfo(  ) == classTypeInfo ) or (initTypeInfo:get_accessMode() == Ast.AccessMode.Local and initTypeInfo:getModule(  ) == self.moduleType ) then
         
      else
       
         self:addErrMess( token.pos, string.format( "can't access to __init of %s", classTypeInfo:getTxt(  )) )
      end
      
      
      local _4779, alt2type, newArgList = self:checkMatchValType( exp:get_pos(), initTypeInfo, argList, classTypeInfo:get_itemTypeInfoList(), classTypeInfo )
      
      if #classTypeInfo:get_itemTypeInfoList() > 0 then
         if classTypeInfo:get_itemTypeInfoList()[1]:get_kind() == Ast.TypeInfoKind.Alternate then
            local genTypeList = {}
            local detect = true
            for __index, altType in ipairs( classTypeInfo:get_itemTypeInfoList() ) do
               do
                  local _exp = alt2type[altType]
                  if _exp ~= nil then
                     table.insert( genTypeList, _exp )
                  else
                     self:addErrMess( token.pos, string.format( "Can't new generic class. -- %s", classTypeInfo:getTxt(  )) )
                     detect = false
                     break
                  end
               end
               
            end
            
            
            if detect then
               classTypeInfo = self.processInfo:createGeneric( classTypeInfo, genTypeList, self.moduleType )
            end
            
         end
         
      end
      
      
      exp = Nodes.ExpNewNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {classTypeInfo}, exp, initTypeInfo, newArgList )
      exp = self:analyzeExpCont( firstToken, exp, false, canLeftExp )
      
      return exp
   end
   
   local function processOp1( token )
   
      if token.txt == "`" then
         return self:analyzeExpMacroStat( token ), false
      end
      
      
      local exp = self:analyzeExpOneRVal( false, true, _lune.unwrap( op1levelMap[token.txt]) )
      local typeInfo = Ast.builtinTypeNone
      local macroExpFlag = false
      local expType = exp:get_expType()
      
      if expType:get_kind() == Ast.TypeInfoKind.DDD then
         self:addErrMess( exp:get_pos(), string.format( "... can't evaluate for '%s'.", token.txt) )
      end
      
      
      do
         local _switchExp = (token.txt )
         if _switchExp == "-" then
            if not expType:equals( self.processInfo, Ast.builtinTypeInt ) and not expType:equals( self.processInfo, Ast.builtinTypeReal ) then
               self:addErrMess( token.pos, string.format( 'unmatch type for "-" -- %s', expType:getTxt(  )) )
            end
            
            typeInfo = expType
         elseif _switchExp == "#" then
            if expType:get_extedType():get_kind() ~= Ast.TypeInfoKind.List and expType:get_extedType():get_kind() ~= Ast.TypeInfoKind.Array and not Ast.builtinTypeString:canEvalWith( self.processInfo, expType, Ast.CanEvalType.SetOp, {} ) then
               self:addErrMess( token.pos, string.format( 'unmatch type for "#" -- %s', expType:getTxt(  )) )
            end
            
            typeInfo = Ast.builtinTypeInt
         elseif _switchExp == "not" then
            typeInfo = Ast.builtinTypeBool
            
            if not expType:get_nilable() and not expType:equals( self.processInfo, Ast.builtinTypeBool ) and not expType:equals( self.processInfo, Ast.builtinTypeStem ) and expType:get_kind() ~= Ast.TypeInfoKind.DDD then
               self:addErrMess( token.pos, "this 'not' operand never be false" )
            end
            
         elseif _switchExp == ",," then
            macroExpFlag = true
            typeInfo = expType
         elseif _switchExp == ",,," then
            macroExpFlag = true
            if not expType:equals( self.processInfo, Ast.builtinTypeString ) then
               self:error( "unmatch ,,, type, need string type" )
            end
            
            typeInfo = Ast.builtinTypeSymbol
         elseif _switchExp == ",,,," then
            macroExpFlag = true
            if not expType:equals( self.processInfo, Ast.builtinTypeSymbol ) then
               self:error( "unmatch ,,, type, need symbol type" )
            end
            
            typeInfo = Ast.builtinTypeString
         elseif _switchExp == "`" then
            typeInfo = Ast.builtinTypeNone
         elseif _switchExp == "~" then
            if not expType:equals( self.processInfo, Ast.builtinTypeInt ) then
               self:addErrMess( token.pos, string.format( 'unmatch type for "~" -- %s', expType:getTxt(  )) )
            end
            
            typeInfo = Ast.builtinTypeInt
         else 
            
               self:error( "unknown op1" )
         end
      end
      
      
      if macroExpFlag then
         local nextToken = self:getToken( true )
         if nextToken.txt ~= "~~" then
            self:pushback(  )
         end
         
      end
      
      
      exp = Nodes.ExpOp1Node.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {typeInfo}, token, self.macroCtrl:get_analyzeInfo():get_mode(), self.nodeManager:MultiTo1( exp ) )
      return self:analyzeExpOp2( firstToken, exp, prevOpLevel ), true
   end
   
   local token = firstToken
   local exp = self:createNoneNode( firstToken.pos )
   
   if token.txt == "##" then
      if allowNoneType then
         self:addErrMess( token.pos, "illeal syntax -- ##" )
      end
      
      return Nodes.AbbrNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeAbbr} )
   end
   
   
   if token.kind == Parser.TokenKind.Dlmt then
      if token.txt == "." then
         if expectType ~= nil then
            local orgExpectType = expectType
            if orgExpectType:get_nilable() then
               orgExpectType = orgExpectType:get_nonnilableType()
            end
            
            exp = processsExpectExp( token, orgExpectType )
         else
            self:error( "illegal '.'" )
         end
         
      elseif token.txt == '[' or token.txt == '[@' then
         exp = self:analyzeListConst( token, expectType )
      elseif token.txt == '(@' then
         exp = self:analyzeSetConst( token, expectType )
      elseif token.txt == '{' then
         exp = self:analyzeMapConst( token, expectType )
      elseif token.txt == "(" then
         exp = self:analyzeExp( false, false, false )
         self:checkNextToken( ")" )
         if not exp:canBeRight( self.processInfo ) then
            self:addErrMess( exp:get_pos(), string.format( "can't be r-value in paren. -- %s", Nodes.getNodeKindName( exp:get_kind() )) )
         end
         
         exp = Nodes.ExpParenNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {exp:get_expType()}, exp )
         exp = self:analyzeExpCont( firstToken, exp, false, canLeftExp )
      end
      
   end
   
   
   if token.txt == "new" then
      exp = processsNewExp( token )
   end
   
   
   if token.kind == Parser.TokenKind.Ope and Parser.isOp1( token.txt ) then
      local workExp, fin = processOp1( token )
      if fin then
         return workExp
      end
      
      exp = workExp
   end
   
   
   if token.kind == Parser.TokenKind.Int then
      exp = Nodes.LiteralIntNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeInt}, token, math.floor((_lune.unwrapDefault( tonumber( token.txt ), 0) )) )
   elseif token.kind == Parser.TokenKind.Real then
      exp = Nodes.LiteralRealNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeReal}, token, (_lune.unwrapDefault( tonumber( token.txt ), 0.0) ) )
   elseif token.kind == Parser.TokenKind.Char then
      local num
      
      if #token.txt == 1 then
         num = string.byte( token.txt, 1 )
      else
       
         num = _lune.unwrap( quotedChar2Code[token.txt:sub( 2, 2 )])
      end
      
      exp = Nodes.LiteralCharNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeChar}, token, num )
   elseif token.kind == Parser.TokenKind.Str then
      exp = self:analyzeStrConst( firstToken, token )
   elseif token.kind == Parser.TokenKind.Symb and token.txt == "__line__" then
      local lineNo = token.pos.lineNo
      if self.macroCtrl:get_analyzeInfo():get_mode() ~= Nodes.MacroMode.None then
         lineNo = self.macroCtrl:get_macroCallLineNo()
      end
      
      exp = Nodes.LiteralIntNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeInt}, Parser.Token.new(Parser.TokenKind.Int, string.format( "%d", lineNo), token.pos, false, nil), token.pos.lineNo )
   elseif token.kind == Parser.TokenKind.Kywd and token.txt == "fn" then
      exp = self:analyzeExpSymbol( firstToken, token, ExpSymbolMode.Fn, nil, false, false )
   elseif token.kind == Parser.TokenKind.Kywd and token.txt == "unwrap" then
      exp = self:analyzeExpUnwrap( token )
   elseif token.kind == Parser.TokenKind.Symb then
      exp = self:analyzeExpSymbol( firstToken, token, ExpSymbolMode.Symbol, nil, false, canLeftExp )
      local symbolInfoList = exp:getSymbolInfo(  )
      if #symbolInfoList == 1 then
         local symbolInfo = symbolInfoList[1]
         if symbolInfo:get_kind() == Ast.SymbolKind.Typ then
            exp = self:analyzeRefTypeWithSymbol( Ast.AccessMode.Local, false, nil, exp, false )
            local workToken = self:getToken(  )
            if workToken.txt == "." then
               exp = self:analyzeExpSymbol( firstToken, self:getToken(  ), ExpSymbolMode.Field, exp, false, canLeftExp )
            else
             
               self:pushback(  )
            end
            
         end
         
      end
      
   elseif token.kind == Parser.TokenKind.Type then
      local symbolTypeInfo = Ast.getSym2builtInTypeMap(  )[token.txt]
      if  nil == symbolTypeInfo then
         local _symbolTypeInfo = symbolTypeInfo
      
         self:error( string.format( "unknown type -- %s", token.txt) )
      end
      
      exp = Nodes.ExpRefNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, Ast.AccessSymbolInfo.new(self.processInfo, symbolTypeInfo, _lune.newAlge( Ast.OverrideMut.None), false) )
   elseif token.kind == Parser.TokenKind.Kywd and (token.txt == "true" or token.txt == "false" ) then
      exp = Nodes.LiteralBoolNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeBool}, token )
   elseif token.kind == Parser.TokenKind.Kywd and (token.txt == "nil" or token.txt == "null" ) then
      exp = Nodes.LiteralNilNode.create( self.nodeManager, firstToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNil} )
   end
   
   
   if exp:get_kind() == Nodes.NodeKind.get_None() then
      self:error( "illegal exp" )
   end
   
   
   if skipOp2Flag then
      return exp
   end
   
   
   return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
end

function TransUnit:analyzeReturn( token )

   local expList = nil
   local funcTypeInfo = self:getCurrentNamespaceTypeInfo(  )
   if funcTypeInfo == Ast.headTypeInfo or (funcTypeInfo:get_kind() ~= Ast.TypeInfoKind.Func and funcTypeInfo:get_kind() ~= Ast.TypeInfoKind.Method ) then
      self:addErrMess( token.pos, "'return' could not use here" )
   end
   
   
   local nextToken = self:getToken(  )
   
   local retTypeList = funcTypeInfo:get_retTypeInfoList()
   if nextToken.txt ~= ";" then
      self:pushback(  )
      expList = self:analyzeExpList( false, false, false, nil, retTypeList )
      self:checkNextToken( ";" )
   end
   
   
   if funcTypeInfo:getTxt(  ) == "__init" then
      self:addErrMess( token.pos, "__init method can't return" )
   end
   
   
   do
      local workList = expList
      if workList ~= nil then
         do
            local _4883, _4884, newExpNodeList = self:checkMatchType( "return", token.pos, retTypeList, workList, false, not workList:get_followOn(), nil )
            if newExpNodeList ~= nil then
               expList = newExpNodeList
            end
         end
         
      else
         if #retTypeList ~= 0 then
            if retTypeList[1]:get_kind() ~= Ast.TypeInfoKind.DDD then
               self:addErrMess( token.pos, string.format( "no return value -- need values(%d)", #retTypeList) )
            end
            
         end
         
      end
   end
   
   
   return Nodes.ReturnNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), retTypeList, expList )
end


function TransUnit:analyzeStatement( termTxt )

   self.commentCtrl:push(  )
   
   local token = self:getTokenNoErr(  )
   local statement = nil
   if token.kind == Parser.TokenKind.Sheb then
      statement = Nodes.ShebangNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, token.txt )
   end
   
   if token == Parser.getEofToken(  ) then
      return nil
   end
   
   
   
   if not statement then
      statement = self:analyzeDecl( Ast.AccessMode.Local, false, token, token )
   end
   
   
   if not statement then
      if token.txt == termTxt then
         if #self.commentCtrl:get_commentList() > 0 then
            local commentToken = self.commentCtrl:get_commentList()[1]
            statement = Nodes.BlankLineNode.create( self.nodeManager, commentToken.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, 0 )
         end
         
         self:pushback(  )
         self.commentCtrl:pop(  )
         return statement
      elseif Ast.txt2AccessMode( token.txt ) then
         local accessMode = Ast.txt2AccessMode( token.txt )
         if  nil == accessMode then
            local _accessMode = accessMode
         
            accessMode = Ast.AccessMode.Pri
         end
         
         local staticFlag = (token.txt == "static" )
         local nextToken = token
         if token.txt ~= "static" then
            nextToken = self:getToken(  )
         end
         
         statement = self:analyzeDecl( accessMode, staticFlag, token, nextToken )
         if not statement then
            self:addErrMess( nextToken.pos, string.format( "This token is illegal -- %s", nextToken.txt) )
         end
         
      elseif token.txt == "{" then
         self:pushback(  )
         statement = self:analyzeBlock( Nodes.BlockKind.Block, TentativeMode.Simple )
      elseif token.txt == "super" then
         statement = self:analyzeSuper( token )
      elseif token.txt == "__envLock" then
         statement = self:analyzeEnvLock( token )
      elseif token.txt == "if" then
         statement = self:analyzeIf( token )
      elseif token.txt == "when" then
         statement = self:analyzeWhen( token )
      elseif token.txt == "switch" or token.txt == "_switch" then
         statement = self:analyzeSwitch( token )
      elseif token.txt == "match" or token.txt == "_match" then
         statement = self:analyzeMatch( token )
      elseif token.txt == "while" then
         statement = self:analyzeWhile( token )
      elseif token.txt == "repeat" then
         statement = self:analyzeRepeat( token )
      elseif token.txt == "for" then
         statement = self:analyzeFor( token )
      elseif token.txt == "apply" then
         statement = self:analyzeApply( token )
      elseif token.txt == "foreach" then
         statement = self:analyzeForeach( token, false )
      elseif token.txt == "forsort" then
         statement = self:analyzeForeach( token, true )
      elseif token.txt == "return" then
         statement = self:analyzeReturn( token )
      elseif token.txt == "break" then
         self:checkNextToken( ";" )
         statement = Nodes.BreakNode.create( self.nodeManager, token.pos, self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone} )
         if #self.loopScopeQueue == 0 then
            self:addErrMess( token.pos, "no loop syntax." )
         end
         
      elseif token.txt == "unwrap" then
         statement = self:analyzeUnwrap( token )
      elseif token.txt == "sync" then
         statement = self:analyzeDeclVar( Nodes.DeclVarMode.Sync, Ast.AccessMode.Local, token )
      elseif token.txt == "__scope" then
         statement = self:analyzeScope( token )
      elseif token.txt == "import" then
         statement = self:analyzeImport( token )
      elseif token.txt == "subfile" then
         statement = self:analyzeSubfile( token )
      elseif token.txt == "_lune_control" then
         do
            local _exp = self:analyzeLuneControl( token )
            if _exp ~= nil then
               statement = _exp
            else
               statement = self:createNoneNode( token.pos )
            end
         end
         
      elseif token.txt == "provide" then
         statement = self:analyzeProvide( token )
      elseif token.txt == ";" then
         statement = self:createNoneNode( token.pos )
      elseif token.txt == ",," or token.txt == ",,," or token.txt == ",,,," then
         self:error( string.format( "illegal macro op -- %s", token.txt) )
      else
       
         self:pushback(  )
         
         self.accessSymbolSetQueue:push(  )
         
         local exp = self:analyzeExp( true, false, true )
         local nextToken = self:getToken(  )
         if nextToken.txt == "," then
            local expList = self:analyzeExpList( true, true, true, exp )
            exp = self:analyzeExpOp2( token, expList, nil )
            nextToken = self:getToken(  )
         end
         
         self:checkToken( nextToken, ";" )
         
         do
            local setNode = _lune.__Cast( exp, 3, Nodes.ExpSetValNode )
            if setNode ~= nil then
               local set = {}
               for __index, symbol in ipairs( setNode:get_LeftSymList() ) do
                  set[symbol:getOrg(  )]= true
               end
               
               self.accessSymbolSetQueue:pop( set )
            else
               self.accessSymbolSetQueue:pop( nil )
            end
         end
         
         statement = Nodes.StmtExpNode.create( self.nodeManager, exp:get_pos(), self.macroCtrl:isInAnalyzeArgMode(  ), {Ast.builtinTypeNone}, exp )
      end
      
   end
   
   
   if statement ~= nil then
      if not statement:canBeStatement(  ) then
         self:addErrMess( statement:get_pos(), string.format( "This node can't be statement. -- %s", Nodes.getNodeKindName( statement:get_kind() )) )
      end
      
   end
   
   self.commentCtrl:pop(  )
   
   return statement
end


return _moduleObj
