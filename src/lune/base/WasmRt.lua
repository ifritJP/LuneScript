--lune/base/WasmRt.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@WasmRt'
local _lune = {}
if _lune8 then
   _lune = _lune8
end
function _lune.loadModule( mod )
   if __luneScript and not package.preload[ mod ] then
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

if not _lune8 then
   _lune8 = _lune
end
local hashI64



local W = _lune.loadModule( 'lune.base.WasmBin' )

local fBaseHash = 0
_moduleObj.fBaseHash = fBaseHash


local fListItems = 1
_moduleObj.fListItems = fListItems


local fListLen = 2
_moduleObj.fListLen = fListLen


local fMapKeys = 1
_moduleObj.fMapKeys = fMapKeys


local fMapVals = 2
_moduleObj.fMapVals = fMapVals


local fMapCount = 3
_moduleObj.fMapCount = fMapCount


local fMapUsed = 4
_moduleObj.fMapUsed = fMapUsed


local fObjVt = 1
_moduleObj.fObjVt = fObjVt


local fVtItabs = 0
_moduleObj.fVtItabs = fVtItabs


local fCloCode = 1
_moduleObj.fCloCode = fCloCode


local CloInfo = {}
_moduleObj.CloInfo = CloInfo
function CloInfo._setmeta( obj )
  setmetatable( obj, { __index = CloInfo  } )
end
function CloInfo._new( funcType, cloType, params, results )
   local obj = {}
   CloInfo._setmeta( obj )
   if obj.__init then
      obj:__init( funcType, cloType, params, results )
   end
   return obj
end
function CloInfo:__init( funcType, cloType, params, results )

   self.funcType = funcType
   self.cloType = cloType
   self.params = params
   self.results = results
end
function CloInfo:get_funcType()
   return self.funcType
end
function CloInfo:get_cloType()
   return self.cloType
end
function CloInfo:get_params()
   return self.params
end
function CloInfo:get_results()
   return self.results
end


local function lg( index )

   return W.localGet( index )
end

local function ls( index, val )

   return W.localSet( index, val )
end

local function i32c( val )

   return W.i32Const( val )
end

local function i64c( val )

   return W.i64Const( val )
end

local function append( list, items )

   for __index, item in ipairs( items ) do
      table.insert( list, item )
   end
end

local function p( code, args )

   return W.plain( code, args )
end

local function when_( cond, thenBody )

   return W.ifElse( cond, thenBody, nil, nil )
end

local function incr( index )

   return ls( index, p( W.opI32Add, {lg( index ), i32c( 1 )} ) )
end

local cModule = "lnsrt"
_moduleObj.cModule = cModule


local jsModule = "lnsjs"
_moduleObj.jsModule = jsModule



local Runtime = {}
_moduleObj.Runtime = Runtime
function Runtime._new( mod )
   local obj = {}
   Runtime._setmeta( obj )
   if obj.__init then obj:__init( mod ); end
   return obj
end
function Runtime:__init(mod) 
   self.mod = mod
   self.name2func = {}
   self.lazyFuncMap = {}
   self.sig2cloInfo = {}
   self.name2luaFunc = {}
   self.tFile = -1
   self.luaImported = false
   self.tStr = mod:addArrayType( W.fieldI8( true ) )
   self.strType = W.refType( self.tStr )
   self.tAnyArr = mod:addArrayType( W.field( W.typeAnyRef, true ) )
   self.anyArrType = W.refType( self.tAnyArr )
   local hashField = W.field( W.typeI32, true )
   self.tBase = mod:addStructType( nil, {hashField} )
   self.tBoxI64 = mod:addStructType( nil, {W.field( W.typeI64, false )} )
   self.tBoxF64 = mod:addStructType( nil, {W.field( W.typeF64, false )} )
   self.tBoxBool = mod:addStructType( nil, {W.field( W.typeI32, false )} )
   self.tList = mod:addStructType( self.tBase, {hashField, W.field( self.anyArrType, true ), W.field( W.typeI32, true )} )
   self.listType = W.refType( self.tList )
   self.tMap = mod:addStructType( self.tBase, {hashField, W.field( self.anyArrType, true ), W.field( self.anyArrType, true ), W.field( W.typeI32, true ), W.field( W.typeI32, true )} )
   self.mapType = W.refType( self.tMap )
   self.tVT = mod:addStructType( nil, {W.field( self.anyArrType, false )} )
   self.tObj = mod:addStructType( self.tBase, {hashField, W.field( W.refType( self.tVT ), false )} )
   self.objType = W.refType( self.tObj )
   self.tClo = mod:addStructType( self.tBase, {hashField} )
   self.cloType = W.refType( self.tClo )
   self.tCellI64 = mod:addStructType( nil, {W.field( W.typeI64, true )} )
   self.tCellF64 = mod:addStructType( nil, {W.field( W.typeF64, true )} )
   self.tCellI32 = mod:addStructType( nil, {W.field( W.typeI32, true )} )
   self.tCellRef = mod:addStructType( nil, {W.field( W.typeAnyRef, true )} )
   self.tLuaVal = mod:addStructType( self.tBase, {hashField, W.field( W.typeI32, false ), W.field( W.typeI32, false )} )
   mod:importMemory( _moduleObj.cModule, "memory" )
   self.fnCBuf = mod:importFunc( _moduleObj.cModule, "lnsr_buf", {W.typeI32}, {W.typeI32} )
   mod:setFuncName( self.fnCBuf, "lnsrt.buf" )
   self.fnCRes = mod:importFunc( _moduleObj.cModule, "lnsr_res", {}, {W.typeI32} )
   mod:setFuncName( self.fnCRes, "lnsrt.res" )
   self.gHashCounter = mod:addGlobal( W.typeI32, i32c( 0 ) )
   self.gTomb = mod:addGlobal( W.refType( self.tBase ), W.structNew( self.tBase, {i32c( -1 )} ) )
end
function Runtime._setmeta( obj )
  setmetatable( obj, { __index = Runtime  } )
end
function Runtime:get_tStr()
   return self.tStr
end
function Runtime:get_tAnyArr()
   return self.tAnyArr
end
function Runtime:get_tBase()
   return self.tBase
end
function Runtime:get_tBoxI64()
   return self.tBoxI64
end
function Runtime:get_tBoxF64()
   return self.tBoxF64
end
function Runtime:get_tBoxBool()
   return self.tBoxBool
end
function Runtime:get_tList()
   return self.tList
end
function Runtime:get_tMap()
   return self.tMap
end
function Runtime:get_tVT()
   return self.tVT
end
function Runtime:get_tObj()
   return self.tObj
end
function Runtime:get_tClo()
   return self.tClo
end
function Runtime:get_tCellI64()
   return self.tCellI64
end
function Runtime:get_tCellF64()
   return self.tCellF64
end
function Runtime:get_tCellI32()
   return self.tCellI32
end
function Runtime:get_tCellRef()
   return self.tCellRef
end
function Runtime:get_tLuaVal()
   return self.tLuaVal
end
function Runtime:get_strType()
   return self.strType
end
function Runtime:get_anyArrType()
   return self.anyArrType
end
function Runtime:get_listType()
   return self.listType
end
function Runtime:get_mapType()
   return self.mapType
end
function Runtime:get_objType()
   return self.objType
end
function Runtime:get_cloType()
   return self.cloType
end
function Runtime:get_luaImported()
   return self.luaImported
end

function Runtime:getCloInfo( params, results )

   local funcParams = {self.cloType}
   for __index, param in ipairs( params ) do
      table.insert( funcParams, param )
   end
   local funcType = self.mod:getFuncType( funcParams, results )
   local key = string.format( "%d", funcType)
   do
      local info = self.sig2cloInfo[key]
      if info ~= nil then
         return info
      end
   end
   local cloType = self.mod:addStructType( self.tClo, {W.field( W.typeI32, true ), W.field( W.refType( funcType ), false )} )
   local info = CloInfo._new(funcType, cloType, params, results)
   self.sig2cloInfo[key] = info
   return info
end

function Runtime:getCellType( valType )

   do
      local _switchExp = valType:get_kind()
      if _switchExp == W.ValKind.I64 then
         return self.tCellI64
      elseif _switchExp == W.ValKind.F64 then
         return self.tCellF64
      elseif _switchExp == W.ValKind.I32 then
         return self.tCellI32
      elseif _switchExp == W.ValKind.Ref then
         return self.tCellRef
      end
   end
end

function Runtime:getFunc( name )

   return self.name2func[name]
end

function Runtime:newFunc( name, params, results )

   local index, func = self.mod:addFunc( params, results )
   self.name2func[name] = index
   self.mod:setFuncName( index, "rt." .. name )
   return index, func
end

function Runtime:boxI64( val )

   return W.structNew( self.tBoxI64, {val} )
end

function Runtime:boxF64( val )

   return W.structNew( self.tBoxF64, {val} )
end

function Runtime:boxBool( val )

   return W.structNew( self.tBoxBool, {val} )
end

function Runtime:unboxI64( val )

   return W.structGet( self.tBoxI64, 0, W.refCast( self.tBoxI64, val ) )
end

function Runtime:unboxF64( val )

   return W.call( self:fnToF64(  ), {val} )
end

function Runtime:unboxBool( val )

   return W.structGet( self.tBoxBool, 0, W.refCast( self.tBoxBool, val ) )
end

function Runtime:fnToF64(  )

   do
      local index = self:getFunc( "toF64" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "toF64", {W.typeAnyRef}, {W.typeF64} )
   func:setBody( {when_( W.refTest( self.tBoxI64, lg( 0 ) ), {W.ret( {p( W.opF64ConvertI64S, {self:unboxI64( lg( 0 ) )} )} )} ), W.structGet( self.tBoxF64, 0, W.refCast( self.tBoxF64, lg( 0 ) ) )} )
   return index
end

function Runtime:strLiteral( txt )

   local dataIdx = self.mod:addData( txt )
   return W.arrayNewData( self.tStr, dataIdx, i32c( 0 ), i32c( #txt ) )
end


function Runtime:strLen( val )

   return W.arrayLen( val )
end

function Runtime:fnHostWrite(  )

   return self:hostFunc( "write", {self.strType}, {} )
end

function Runtime:fnHostWriteErr(  )

   return self:hostFunc( "write_err", {self.strType}, {} )
end

function Runtime:error( mess )

   return W.seq( {W.call( self:hostFunc( "error", {self.strType}, {} ), {mess} ), W.unreachable(  )} )
end

function Runtime:fnStrConcat(  )

   do
      local index = self:getFunc( "strConcat" )
      if index ~= nil then
         return index
      end
   end
   local strType = self.strType
   local strIdx = self.tStr
   local index, func = self:newFunc( "strConcat", {strType, strType}, {strType} )
   local lenA = func:addLocal( W.typeI32 )
   local lenB = func:addLocal( W.typeI32 )
   local work = func:addLocal( strType )
   func:setBody( {ls( lenA, W.arrayLen( lg( 0 ) ) ), ls( lenB, W.arrayLen( lg( 1 ) ) ), ls( work, W.arrayNewDefault( strIdx, p( W.opI32Add, {lg( lenA ), lg( lenB )} ) ) ), W.arrayCopy( strIdx, strIdx, lg( work ), i32c( 0 ), lg( 0 ), i32c( 0 ), lg( lenA ) ), W.arrayCopy( strIdx, strIdx, lg( work ), lg( lenA ), lg( 1 ), i32c( 0 ), lg( lenB ) ), lg( work )} )
   return index
end

function Runtime:fnStrCmp(  )

   do
      local index = self:getFunc( "strCmp" )
      if index ~= nil then
         return index
      end
   end
   local strType = self.strType
   local strIdx = self.tStr
   local index, func = self:newFunc( "strCmp", {strType, strType}, {W.typeI32} )
   local lenA = func:addLocal( W.typeI32 )
   local lenB = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   local diff = func:addLocal( W.typeI32 )
   func:setBody( {ls( lenA, W.arrayLen( lg( 0 ) ) ), ls( lenB, W.arrayLen( lg( 1 ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( lenA )} ) ), W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( lenB )} ) ), ls( diff, p( W.opI32Sub, {W.arrayGetU( strIdx, lg( 0 ), lg( pos ) ), W.arrayGetU( strIdx, lg( 1 ), lg( pos ) )} ) ), when_( lg( diff ), {W.ret( {lg( diff )} )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), p( W.opI32Sub, {lg( lenA ), lg( lenB )} )} )
   return index
end

function Runtime:fnStrEq(  )

   do
      local index = self:getFunc( "strEq" )
      if index ~= nil then
         return index
      end
   end
   local strType = self.strType
   local strIdx = self.tStr
   local index, func = self:newFunc( "strEq", {strType, strType}, {W.typeI32} )
   local len = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {when_( W.refEq( lg( 0 ), lg( 1 ) ), {W.ret( {i32c( 1 )} )} ), when_( p( W.opI32Or, {W.refIsNull( lg( 0 ) ), W.refIsNull( lg( 1 ) )} ), {W.ret( {i32c( 0 )} )} ), ls( len, W.arrayLen( lg( 0 ) ) ), when_( p( W.opI32Ne, {lg( len ), W.arrayLen( lg( 1 ) )} ), {W.ret( {i32c( 0 )} )} ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( len )} ) ), when_( p( W.opI32Ne, {W.arrayGetU( strIdx, lg( 0 ), lg( pos ) ), W.arrayGetU( strIdx, lg( 1 ), lg( pos ) )} ), {W.ret( {i32c( 0 )} )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), i32c( 1 )} )
   return index
end

function Runtime:fnStrSub(  )

   do
      local index = self:getFunc( "strSub" )
      if index ~= nil then
         return index
      end
   end
   local strType = self.strType
   local strIdx = self.tStr
   local index, func = self:newFunc( "strSub", {strType, W.typeI64, W.typeI64}, {strType} )
   local len = func:addLocal( W.typeI64 )
   local work = func:addLocal( strType )
   func:setBody( {ls( len, p( W.opI64ExtendI32U, {W.arrayLen( lg( 0 ) )} ) ), when_( p( W.opI64LtS, {lg( 1 ), i64c( 0 )} ), {ls( 1, p( W.opI64Add, {lg( len ), p( W.opI64Add, {lg( 1 ), i64c( 1 )} )} ) )} ), when_( p( W.opI64LtS, {lg( 2 ), i64c( 0 )} ), {ls( 2, p( W.opI64Add, {lg( len ), p( W.opI64Add, {lg( 2 ), i64c( 1 )} )} ) )} ), when_( p( W.opI64LtS, {lg( 1 ), i64c( 1 )} ), {ls( 1, i64c( 1 ) )} ), when_( p( W.opI64GtS, {lg( 2 ), lg( len )} ), {ls( 2, lg( len ) )} ), when_( p( W.opI64GtS, {lg( 1 ), lg( 2 )} ), {W.ret( {W.arrayNewDefault( strIdx, i32c( 0 ) )} )} ), ls( work, W.arrayNewDefault( strIdx, p( W.opI32WrapI64, {p( W.opI64Add, {p( W.opI64Sub, {lg( 2 ), lg( 1 )} ), i64c( 1 )} )} ) ) ), W.arrayCopy( strIdx, strIdx, lg( work ), i32c( 0 ), lg( 0 ), p( W.opI32WrapI64, {p( W.opI64Sub, {lg( 1 ), i64c( 1 )} )} ), W.arrayLen( lg( work ) ) ), lg( work )} )
   return index
end

function Runtime:fnI64ToStr(  )

   do
      local index = self:getFunc( "i64ToStr" )
      if index ~= nil then
         return index
      end
   end
   local strType = self.strType
   local strIdx = self.tStr
   local index, func = self:newFunc( "i64ToStr", {W.typeI64}, {strType} )
   local buf = func:addLocal( strType )
   local pos = func:addLocal( W.typeI32 )
   local neg = func:addLocal( W.typeI32 )
   local digit = func:addLocal( W.typeI64 )
   local work = func:addLocal( strType )
   func:setBody( {ls( buf, W.arrayNewDefault( strIdx, i32c( 20 ) ) ), ls( pos, i32c( 20 ) ), ls( neg, p( W.opI64LtS, {lg( 0 ), i64c( 0 )} ) ), W.loop( {ls( digit, p( W.opI64RemS, {lg( 0 ), i64c( 10 )} ) ), when_( lg( neg ), {ls( digit, p( W.opI64Sub, {i64c( 0 ), lg( digit )} ) )} ), ls( pos, p( W.opI32Sub, {lg( pos ), i32c( 1 )} ) ), W.arraySet( strIdx, lg( buf ), lg( pos ), p( W.opI32Add, {i32c( 0x30 ), p( W.opI32WrapI64, {lg( digit )} )} ) ), ls( 0, p( W.opI64DivS, {lg( 0 ), i64c( 10 )} ) ), W.brIf( 0, p( W.opI64Ne, {lg( 0 ), i64c( 0 )} ) )}, nil ), when_( lg( neg ), {ls( pos, p( W.opI32Sub, {lg( pos ), i32c( 1 )} ) ), W.arraySet( strIdx, lg( buf ), lg( pos ), i32c( 0x2d ) )} ), ls( work, W.arrayNewDefault( strIdx, p( W.opI32Sub, {i32c( 20 ), lg( pos )} ) ) ), W.arrayCopy( strIdx, strIdx, lg( work ), i32c( 0 ), lg( buf ), lg( pos ), W.arrayLen( lg( work ) ) ), lg( work )} )
   return index
end

function Runtime:fnHex8(  )

   do
      local index = self:getFunc( "hex8" )
      if index ~= nil then
         return index
      end
   end
   local strType = self.strType
   local strIdx = self.tStr
   local index, func = self:newFunc( "hex8", {W.typeI32}, {strType} )
   local buf = func:addLocal( strType )
   local pos = func:addLocal( W.typeI32 )
   local digit = func:addLocal( W.typeI32 )
   func:setBody( {ls( buf, W.arrayNewDefault( strIdx, i32c( 8 ) ) ), ls( pos, i32c( 8 ) ), W.loop( {ls( pos, p( W.opI32Sub, {lg( pos ), i32c( 1 )} ) ), ls( digit, p( W.opI32And, {lg( 0 ), i32c( 15 )} ) ), W.arraySet( strIdx, lg( buf ), lg( pos ), W.select( p( W.opI32Add, {i32c( 0x30 ), lg( digit )} ), p( W.opI32Add, {i32c( 0x61 - 10 ), lg( digit )} ), p( W.opI32LtU, {lg( digit ), i32c( 10 )} ) ) ), ls( 0, p( W.opI32ShrU, {lg( 0 ), i32c( 4 )} ) ), W.brIf( 0, lg( pos ) )}, nil ), lg( buf )} )
   return index
end

function Runtime:fnIdHash(  )

   do
      local index = self:getFunc( "idHash" )
      if index ~= nil then
         return index
      end
   end
   local baseType = W.refType( self.tBase )
   local index, func = self:newFunc( "idHash", {baseType}, {W.typeI32} )
   local hash = func:addLocal( W.typeI32 )
   func:setBody( {ls( hash, W.structGet( self.tBase, _moduleObj.fBaseHash, lg( 0 ) ) ), when_( p( W.opI32Eqz, {lg( hash )} ), {W.globalSet( self.gHashCounter, p( W.opI32Add, {W.globalGet( self.gHashCounter ), i32c( 1 )} ) ), ls( hash, p( W.opI32Mul, {W.globalGet( self.gHashCounter ), i32c( 0x9e3779b1 )} ) ), when_( p( W.opI32Eqz, {lg( hash )} ), {ls( hash, i32c( 1 ) )} ), W.structSet( self.tBase, _moduleObj.fBaseHash, lg( 0 ), lg( hash ) )} ), lg( hash )} )
   return index
end

function Runtime:fnToStr(  )

   do
      local index = self:getFunc( "toStr" )
      if index ~= nil then
         return index
      end
   end
   local strType = self.strType
   local index, func = self:newFunc( "toStr", {W.typeAnyRef}, {strType} )
   local concat = self:fnStrConcat(  )
   func:setBody( {when_( W.refIsNull( lg( 0 ) ), {W.ret( {self:strLiteral( "nil" )} )} ), when_( W.refTest( self.tStr, lg( 0 ) ), {W.ret( {W.refCast( self.tStr, lg( 0 ) )} )} ), when_( W.refTest( self.tBoxI64, lg( 0 ) ), {W.ret( {W.call( self:fnI64ToStr(  ), {self:unboxI64( lg( 0 ) )} )} )} ), when_( W.refTest( self.tBoxF64, lg( 0 ) ), {W.ret( {W.call( self:fnRealToStr(  ), {W.structGet( self.tBoxF64, 0, W.refCast( self.tBoxF64, lg( 0 ) ) )} )} )} ), when_( W.refTest( self.tBoxBool, lg( 0 ) ), {W.ret( {W.ifElse( self:unboxBool( lg( 0 ) ), {self:strLiteral( "true" )}, {self:strLiteral( "false" )}, strType )} )} ), when_( p( W.opI32Or, {W.refTest( self.tClo, lg( 0 ) ), W.ifElse( W.refTest( self.tLuaVal, lg( 0 ) ), {p( W.opI32Eq, {self:luaValType( lg( 0 ) ), i32c( _moduleObj.luaTFunction )} )}, {i32c( 0 )}, W.typeI32 )} ), {W.ret( {W.call( concat, {self:strLiteral( "function: 0x" ), W.call( self:fnHex8(  ), {W.call( self:fnIdHash(  ), {W.refCast( self.tBase, lg( 0 ) )} )} )} )} )} ), when_( W.refTest( self.tBase, lg( 0 ) ), {W.ret( {W.call( concat, {self:strLiteral( "table: 0x" ), W.call( self:fnHex8(  ), {W.call( self:fnIdHash(  ), {W.refCast( self.tBase, lg( 0 ) )} )} )} )} )} ), self:strLiteral( "userdata" )} )
   return index
end

function Runtime:isNum( val )

   return p( W.opI32Or, {W.refTest( self.tBoxI64, val ), W.refTest( self.tBoxF64, val )} )
end

function Runtime:fnEqAny(  )

   do
      local index = self:getFunc( "eqAny" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "eqAny", {W.typeAnyRef, W.typeAnyRef}, {W.typeI32} )
   func:setBody( {when_( W.refIsNull( lg( 0 ) ), {W.ret( {W.refIsNull( lg( 1 ) )} )} ), when_( W.refIsNull( lg( 1 ) ), {W.ret( {i32c( 0 )} )} ), when_( W.refTest( self.tStr, lg( 0 ) ), {when_( p( W.opI32Eqz, {W.refTest( self.tStr, lg( 1 ) )} ), {W.ret( {i32c( 0 )} )} ), W.ret( {W.call( self:fnStrEq(  ), {W.refCast( self.tStr, lg( 0 ) ), W.refCast( self.tStr, lg( 1 ) )} )} )} ), when_( p( W.opI32And, {self:isNum( lg( 0 ) ), self:isNum( lg( 1 ) )} ), {when_( p( W.opI32And, {W.refTest( self.tBoxI64, lg( 0 ) ), W.refTest( self.tBoxI64, lg( 1 ) )} ), {W.ret( {p( W.opI64Eq, {self:unboxI64( lg( 0 ) ), self:unboxI64( lg( 1 ) )} )} )} ), W.ret( {p( W.opF64Eq, {self:unboxF64( lg( 0 ) ), self:unboxF64( lg( 1 ) )} )} )} ), when_( W.refTest( self.tBoxBool, lg( 0 ) ), {when_( p( W.opI32Eqz, {W.refTest( self.tBoxBool, lg( 1 ) )} ), {W.ret( {i32c( 0 )} )} ), W.ret( {p( W.opI32Eq, {self:unboxBool( lg( 0 ) ), self:unboxBool( lg( 1 ) )} )} )} ), W.refEq( W.refCast( W.htEq, lg( 0 ) ), W.refCast( W.htEq, lg( 1 ) ) )} )
   return index
end

function Runtime:fnLessAny(  )

   do
      local index = self:getFunc( "lessAny" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "lessAny", {W.typeAnyRef, W.typeAnyRef}, {W.typeI32} )
   func:setBody( {when_( p( W.opI32And, {W.refTest( self.tStr, lg( 0 ) ), W.refTest( self.tStr, lg( 1 ) )} ), {W.ret( {p( W.opI32LtS, {W.call( self:fnStrCmp(  ), {W.refCast( self.tStr, lg( 0 ) ), W.refCast( self.tStr, lg( 1 ) )} ), i32c( 0 )} )} )} ), when_( p( W.opI32And, {self:isNum( lg( 0 ) ), self:isNum( lg( 1 ) )} ), {when_( p( W.opI32And, {W.refTest( self.tBoxI64, lg( 0 ) ), W.refTest( self.tBoxI64, lg( 1 ) )} ), {W.ret( {p( W.opI64LtS, {self:unboxI64( lg( 0 ) ), self:unboxI64( lg( 1 ) )} )} )} ), W.ret( {p( W.opF64Lt, {self:unboxF64( lg( 0 ) ), self:unboxF64( lg( 1 ) )} )} )} ), self:error( self:strLiteral( "attempt to compare" ) ), i32c( 0 )} )
   return index
end

function hashI64( val )

   return p( W.opI32WrapI64, {p( W.opI64Xor, {val, p( W.opI64ShrU, {val, i64c( 32 )} )} )} )
end

function Runtime:fnHashAny(  )

   do
      local index = self:getFunc( "hashAny" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "hashAny", {W.typeAnyRef}, {W.typeI32} )
   local hash = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   local len = func:addLocal( W.typeI32 )
   local strVal = func:addLocal( self.strType )
   local fval = func:addLocal( W.typeF64 )
   local ival = func:addLocal( W.typeI64 )
   func:setBody( {when_( W.refIsNull( lg( 0 ) ), {W.ret( {i32c( 0 )} )} ), when_( W.refTest( self.tStr, lg( 0 ) ), {ls( strVal, W.refCast( self.tStr, lg( 0 ) ) ), ls( hash, i32c( 0x811c9dc5 ) ), ls( len, W.arrayLen( lg( strVal ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( len )} ) ), ls( hash, p( W.opI32Mul, {p( W.opI32Xor, {lg( hash ), W.arrayGetU( self.tStr, lg( strVal ), lg( pos ) )} ), i32c( 16777619 )} ) ), incr( pos ), W.br( 0 )}, nil )}, nil ), W.ret( {lg( hash )} )} ), when_( W.refTest( self.tBoxI64, lg( 0 ) ), {ls( ival, self:unboxI64( lg( 0 ) ) ), W.ret( {hashI64( lg( ival ) )} )} ), when_( W.refTest( self.tBoxF64, lg( 0 ) ), {ls( fval, W.structGet( self.tBoxF64, 0, W.refCast( self.tBoxF64, lg( 0 ) ) ) ), when_( p( W.opI32And, {p( W.opF64Eq, {p( W.opF64Trunc, {lg( fval )} ), lg( fval )} ), p( W.opF64Lt, {p( W.opF64Abs, {lg( fval )} ), W.f64Const( 9.2e18 )} )} ), {ls( ival, p( W.opI64TruncF64S, {lg( fval )} ) ), W.ret( {hashI64( lg( ival ) )} )} ), ls( ival, p( W.opI64ReinterpretF64, {lg( fval )} ) ), W.ret( {hashI64( lg( ival ) )} )} ), when_( W.refTest( self.tBoxBool, lg( 0 ) ), {W.ret( {p( W.opI32Add, {self:unboxBool( lg( 0 ) ), i32c( 1 )} )} )} ), when_( W.refTest( self.tBase, lg( 0 ) ), {W.ret( {W.call( self:fnIdHash(  ), {W.refCast( self.tBase, lg( 0 ) )} )} )} ), i32c( 0 )} )
   return index
end

function Runtime:newMap(  )

   return W.call( self:fnMapNew(  ), {i32c( 8 )} )
end


function Runtime:fnMapNew(  )

   do
      local index = self:getFunc( "mapNew" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "mapNew", {W.typeI32}, {self.mapType} )
   func:setBody( {W.structNew( self.tMap, {i32c( 0 ), W.arrayNewDefault( self.tAnyArr, lg( 0 ) ), W.arrayNewDefault( self.tAnyArr, lg( 0 ) ), i32c( 0 ), i32c( 0 )} )} )
   return index
end


function Runtime:isTomb( val )

   return W.refEq( W.refCast( W.htEq, val ), W.globalGet( self.gTomb ) )
end

function Runtime:fnMapFind(  )

   do
      local index = self:getFunc( "mapFind" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "mapFind", {self.mapType, W.typeAnyRef}, {W.typeI32} )
   local keys = func:addLocal( self.anyArrType )
   local mask = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   local key = func:addLocal( W.typeAnyRef )
   func:setBody( {ls( keys, W.structGet( self.tMap, _moduleObj.fMapKeys, lg( 0 ) ) ), ls( mask, p( W.opI32Sub, {W.arrayLen( lg( keys ) ), i32c( 1 )} ) ), ls( pos, p( W.opI32And, {W.call( self:fnHashAny(  ), {lg( 1 )} ), lg( mask )} ) ), W.loop( {ls( key, W.arrayGet( self.tAnyArr, lg( keys ), lg( pos ) ) ), when_( W.refIsNull( lg( key ) ), {W.ret( {i32c( -1 )} )} ), when_( p( W.opI32Eqz, {self:isTomb( lg( key ) )} ), {when_( W.call( self:fnEqAny(  ), {lg( key ), lg( 1 )} ), {W.ret( {lg( pos )} )} )} ), ls( pos, p( W.opI32And, {p( W.opI32Add, {lg( pos ), i32c( 1 )} ), lg( mask )} ) ), W.br( 0 )}, nil ), W.unreachable(  )} )
   return index
end

function Runtime:fnMapGet(  )

   do
      local index = self:getFunc( "mapGet" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "mapGet", {self.mapType, W.typeAnyRef}, {W.typeAnyRef} )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {when_( W.refIsNull( lg( 1 ) ), {W.ret( {W.refNull( W.htAny )} )} ), ls( pos, W.call( self:fnMapFind(  ), {lg( 0 ), lg( 1 )} ) ), when_( p( W.opI32LtS, {lg( pos ), i32c( 0 )} ), {W.ret( {W.refNull( W.htAny )} )} ), W.arrayGet( self.tAnyArr, W.structGet( self.tMap, _moduleObj.fMapVals, lg( 0 ) ), lg( pos ) )} )
   return index
end

function Runtime:fnMapSet(  )

   do
      local index = self:getFunc( "mapSet" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "mapSet", {self.mapType, W.typeAnyRef, W.typeAnyRef}, {} )
   local pos = func:addLocal( W.typeI32 )
   local keys = func:addLocal( self.anyArrType )
   local mask = func:addLocal( W.typeI32 )
   local key = func:addLocal( W.typeAnyRef )
   local cap = func:addLocal( W.typeI32 )
   local map = 0
   func:setBody( {when_( W.refIsNull( lg( 1 ) ), {self:error( self:strLiteral( "table index is nil" ) )} ), ls( pos, W.call( self:fnMapFind(  ), {lg( map ), lg( 1 )} ) ), when_( W.refIsNull( lg( 2 ) ), {when_( p( W.opI32GeS, {lg( pos ), i32c( 0 )} ), {W.arraySet( self.tAnyArr, W.structGet( self.tMap, _moduleObj.fMapKeys, lg( map ) ), lg( pos ), W.globalGet( self.gTomb ) ), W.arraySet( self.tAnyArr, W.structGet( self.tMap, _moduleObj.fMapVals, lg( map ) ), lg( pos ), W.refNull( W.htAny ) ), W.structSet( self.tMap, _moduleObj.fMapCount, lg( map ), p( W.opI32Sub, {W.structGet( self.tMap, _moduleObj.fMapCount, lg( map ) ), i32c( 1 )} ) )} ), W.ret( {} )} ), when_( p( W.opI32GeS, {lg( pos ), i32c( 0 )} ), {W.arraySet( self.tAnyArr, W.structGet( self.tMap, _moduleObj.fMapVals, lg( map ) ), lg( pos ), lg( 2 ) ), W.ret( {} )} ), ls( cap, W.arrayLen( W.structGet( self.tMap, _moduleObj.fMapKeys, lg( map ) ) ) ), when_( p( W.opI32GeS, {p( W.opI32Mul, {p( W.opI32Add, {W.structGet( self.tMap, _moduleObj.fMapUsed, lg( map ) ), i32c( 1 )} ), i32c( 4 )} ), p( W.opI32Mul, {lg( cap ), i32c( 3 )} )} ), {W.call( self:fnMapRehash(  ), {lg( map ), p( W.opI32Mul, {lg( cap ), i32c( 2 )} )} )} ), ls( keys, W.structGet( self.tMap, _moduleObj.fMapKeys, lg( map ) ) ), ls( mask, p( W.opI32Sub, {W.arrayLen( lg( keys ) ), i32c( 1 )} ) ), ls( pos, p( W.opI32And, {W.call( self:fnHashAny(  ), {lg( 1 )} ), lg( mask )} ) ), W.block( {W.loop( {ls( key, W.arrayGet( self.tAnyArr, lg( keys ), lg( pos ) ) ), W.brIf( 1, W.refIsNull( lg( key ) ) ), W.brIf( 1, self:isTomb( lg( key ) ) ), ls( pos, p( W.opI32And, {p( W.opI32Add, {lg( pos ), i32c( 1 )} ), lg( mask )} ) ), W.br( 0 )}, nil )}, nil ), when_( W.refIsNull( lg( key ) ), {W.structSet( self.tMap, _moduleObj.fMapUsed, lg( map ), p( W.opI32Add, {W.structGet( self.tMap, _moduleObj.fMapUsed, lg( map ) ), i32c( 1 )} ) )} ), W.arraySet( self.tAnyArr, lg( keys ), lg( pos ), lg( 1 ) ), W.arraySet( self.tAnyArr, W.structGet( self.tMap, _moduleObj.fMapVals, lg( map ) ), lg( pos ), lg( 2 ) ), W.structSet( self.tMap, _moduleObj.fMapCount, lg( map ), p( W.opI32Add, {W.structGet( self.tMap, _moduleObj.fMapCount, lg( map ) ), i32c( 1 )} ) )} )
   return index
end

function Runtime:fnMapRehash(  )

   do
      local index = self:getFunc( "mapRehash" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "mapRehash", {self.mapType, W.typeI32}, {} )
   local oldKeys = func:addLocal( self.anyArrType )
   local oldVals = func:addLocal( self.anyArrType )
   local pos = func:addLocal( W.typeI32 )
   local key = func:addLocal( W.typeAnyRef )
   local map = 0
   func:setBody( {ls( oldKeys, W.structGet( self.tMap, _moduleObj.fMapKeys, lg( map ) ) ), ls( oldVals, W.structGet( self.tMap, _moduleObj.fMapVals, lg( map ) ) ), when_( p( W.opI32LtS, {p( W.opI32Mul, {W.structGet( self.tMap, _moduleObj.fMapCount, lg( map ) ), i32c( 2 )} ), W.arrayLen( lg( oldKeys ) )} ), {ls( 1, W.arrayLen( lg( oldKeys ) ) )} ), W.structSet( self.tMap, _moduleObj.fMapKeys, lg( map ), W.arrayNewDefault( self.tAnyArr, lg( 1 ) ) ), W.structSet( self.tMap, _moduleObj.fMapVals, lg( map ), W.arrayNewDefault( self.tAnyArr, lg( 1 ) ) ), W.structSet( self.tMap, _moduleObj.fMapCount, lg( map ), i32c( 0 ) ), W.structSet( self.tMap, _moduleObj.fMapUsed, lg( map ), i32c( 0 ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), W.arrayLen( lg( oldKeys ) )} ) ), ls( key, W.arrayGet( self.tAnyArr, lg( oldKeys ), lg( pos ) ) ), when_( p( W.opI32Eqz, {W.refIsNull( lg( key ) )} ), {when_( p( W.opI32Eqz, {self:isTomb( lg( key ) )} ), {W.call( self:fnMapSet(  ), {lg( map ), lg( key ), W.arrayGet( self.tAnyArr, lg( oldVals ), lg( pos ) )} )} )} ), incr( pos ), W.br( 0 )}, nil )}, nil )} )
   return index
end

function Runtime:fnMapNext(  )

   do
      local index = self:getFunc( "mapNext" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "mapNext", {self.mapType, W.typeI32}, {W.typeI32} )
   local keys = func:addLocal( self.anyArrType )
   local key = func:addLocal( W.typeAnyRef )
   func:setBody( {ls( keys, W.structGet( self.tMap, _moduleObj.fMapKeys, lg( 0 ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( 1 ), W.arrayLen( lg( keys ) )} ) ), ls( key, W.arrayGet( self.tAnyArr, lg( keys ), lg( 1 ) ) ), when_( p( W.opI32Eqz, {W.refIsNull( lg( key ) )} ), {when_( p( W.opI32Eqz, {self:isTomb( lg( key ) )} ), {W.ret( {lg( 1 )} )} )} ), incr( 1 ), W.br( 0 )}, nil )}, nil ), i32c( -1 )} )
   return index
end


function Runtime:mapKeyAt( map, pos )

   return W.arrayGet( self.tAnyArr, W.structGet( self.tMap, _moduleObj.fMapKeys, map ), pos )
end

function Runtime:mapValAt( map, pos )

   return W.arrayGet( self.tAnyArr, W.structGet( self.tMap, _moduleObj.fMapVals, map ), pos )
end

function Runtime:fnMapKeys(  )

   do
      local index = self:getFunc( "mapKeys" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "mapKeys", {self.mapType}, {self.listType} )
   local list = func:addLocal( self.listType )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {ls( list, W.call( self:fnListNew(  ), {W.structGet( self.tMap, _moduleObj.fMapCount, lg( 0 ) )} ) ), W.block( {W.loop( {ls( pos, W.call( self:fnMapNext(  ), {lg( 0 ), lg( pos )} ) ), W.brIf( 1, p( W.opI32LtS, {lg( pos ), i32c( 0 )} ) ), W.call( self:fnListPush(  ), {lg( list ), self:mapKeyAt( lg( 0 ), lg( pos ) )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), lg( list )} )
   return index
end

function Runtime:fnMapClone(  )

   do
      local index = self:getFunc( "mapClone" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "mapClone", {self.mapType}, {self.mapType} )
   local map = func:addLocal( self.mapType )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {ls( map, self:newMap(  ) ), W.block( {W.loop( {ls( pos, W.call( self:fnMapNext(  ), {lg( 0 ), lg( pos )} ) ), W.brIf( 1, p( W.opI32LtS, {lg( pos ), i32c( 0 )} ) ), W.call( self:fnMapSet(  ), {lg( map ), self:mapKeyAt( lg( 0 ), lg( pos ) ), self:mapValAt( lg( 0 ), lg( pos ) )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), lg( map )} )
   return index
end

function Runtime:fnSetOp( mode )

   local name = string.format( "setOp%d", mode)
   do
      local index = self:getFunc( name )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( name, {self.mapType, self.mapType}, {self.mapType} )
   local pos = func:addLocal( W.typeI32 )
   local key = func:addLocal( W.typeAnyRef )
   local body = {}
   if mode == 1 then
      table.insert( body, W.block( {W.loop( {ls( pos, W.call( self:fnMapNext(  ), {lg( 1 ), lg( pos )} ) ), W.brIf( 1, p( W.opI32LtS, {lg( pos ), i32c( 0 )} ) ), ls( key, self:mapKeyAt( lg( 1 ), lg( pos ) ) ), W.call( self:fnMapSet(  ), {lg( 0 ), lg( key ), lg( key )} ), incr( pos ), W.br( 0 )}, nil )}, nil ) )
   else
    
      local cond = p( W.opI32GeS, {W.call( self:fnMapFind(  ), {lg( 1 ), lg( key )} ), i32c( 0 )} )
      if mode == 0 then
         cond = p( W.opI32Eqz, {cond} )
      end
      table.insert( body, W.block( {W.loop( {ls( pos, W.call( self:fnMapNext(  ), {lg( 0 ), lg( pos )} ) ), W.brIf( 1, p( W.opI32LtS, {lg( pos ), i32c( 0 )} ) ), ls( key, self:mapKeyAt( lg( 0 ), lg( pos ) ) ), when_( cond, {W.call( self:fnMapSet(  ), {lg( 0 ), lg( key ), W.refNull( W.htAny )} )} ), incr( pos ), W.br( 0 )}, nil )}, nil ) )
   end
   table.insert( body, lg( 0 ) )
   func:setBody( body )
   return index
end

function Runtime:fnListNew(  )

   do
      local index = self:getFunc( "listNew" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "listNew", {W.typeI32}, {self.listType} )
   func:setBody( {when_( p( W.opI32LtS, {lg( 0 ), i32c( 4 )} ), {ls( 0, i32c( 4 ) )} ), W.structNew( self.tList, {i32c( 0 ), W.arrayNewDefault( self.tAnyArr, lg( 0 ) ), i32c( 0 )} )} )
   return index
end

function Runtime:fnListPush(  )

   do
      local index = self:getFunc( "listPush" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "listPush", {self.listType, W.typeAnyRef}, {} )
   local items = func:addLocal( self.anyArrType )
   local len = func:addLocal( W.typeI32 )
   local work = func:addLocal( self.anyArrType )
   func:setBody( {ls( items, W.structGet( self.tList, _moduleObj.fListItems, lg( 0 ) ) ), ls( len, W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) ) ), when_( p( W.opI32GeS, {lg( len ), W.arrayLen( lg( items ) )} ), {ls( work, W.arrayNewDefault( self.tAnyArr, p( W.opI32Add, {p( W.opI32Mul, {lg( len ), i32c( 2 )} ), i32c( 4 )} ) ) ), W.arrayCopy( self.tAnyArr, self.tAnyArr, lg( work ), i32c( 0 ), lg( items ), i32c( 0 ), lg( len ) ), ls( items, lg( work ) ), W.structSet( self.tList, _moduleObj.fListItems, lg( 0 ), lg( items ) )} ), W.arraySet( self.tAnyArr, lg( items ), lg( len ), lg( 1 ) ), W.structSet( self.tList, _moduleObj.fListLen, lg( 0 ), p( W.opI32Add, {lg( len ), i32c( 1 )} ) )} )
   return index
end

function Runtime:fnListAppend(  )

   do
      local index = self:getFunc( "listAppend" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "listAppend", {self.listType, self.listType}, {} )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {when_( W.refIsNull( lg( 1 ) ), {W.ret( {} )} ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), W.structGet( self.tList, _moduleObj.fListLen, lg( 1 ) )} ) ), W.call( self:fnListPush(  ), {lg( 0 ), W.arrayGet( self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, lg( 1 ) ), lg( pos ) )} ), incr( pos ), W.br( 0 )}, nil )}, nil )} )
   return index
end

function Runtime:fnPrintList(  )

   do
      local index = self:getFunc( "printList" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "printList", {self.listType}, {} )
   local pos = func:addLocal( W.typeI32 )
   local tab = self:strLiteral( "\t" )
   func:setBody( {W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) )} ) ), when_( lg( pos ), {W.call( self:fnHostWrite(  ), {tab} )} ), W.call( self:fnHostWrite(  ), {W.call( self:fnToStr(  ), {W.arrayGet( self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, lg( 0 ) ), lg( pos ) )} )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), W.call( self:fnHostWrite(  ), {self:strLiteral( "\n" )} )} )
   return index
end

function Runtime:fnListGet(  )

   do
      local index = self:getFunc( "listGet" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "listGet", {self.listType, W.typeI64}, {W.typeAnyRef} )
   func:setBody( {when_( p( W.opI32Or, {p( W.opI64LtS, {lg( 1 ), i64c( 1 )} ), p( W.opI64GtS, {lg( 1 ), p( W.opI64ExtendI32U, {W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) )} )} )} ), {W.ret( {W.refNull( W.htAny )} )} ), W.arrayGet( self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, lg( 0 ) ), p( W.opI32WrapI64, {p( W.opI64Sub, {lg( 1 ), i64c( 1 )} )} ) )} )
   return index
end

function Runtime:fnListSet(  )

   do
      local index = self:getFunc( "listSet" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "listSet", {self.listType, W.typeI64, W.typeAnyRef}, {} )
   local len = func:addLocal( W.typeI64 )
   func:setBody( {ls( len, p( W.opI64ExtendI32U, {W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) )} ) ), when_( p( W.opI64Eq, {lg( 1 ), p( W.opI64Add, {lg( len ), i64c( 1 )} )} ), {W.call( self:fnListPush(  ), {lg( 0 ), lg( 2 )} ), W.ret( {} )} ), when_( p( W.opI32Or, {p( W.opI64LtS, {lg( 1 ), i64c( 1 )} ), p( W.opI64GtS, {lg( 1 ), lg( len )} )} ), {self:error( self:strLiteral( "list index out of range" ) )} ), W.arraySet( self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, lg( 0 ) ), p( W.opI32WrapI64, {p( W.opI64Sub, {lg( 1 ), i64c( 1 )} )} ), lg( 2 ) )} )
   return index
end

function Runtime:fnListRemove(  )

   do
      local index = self:getFunc( "listRemove" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "listRemove", {self.listType, W.typeI64}, {W.typeAnyRef} )
   local len = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   local items = func:addLocal( self.anyArrType )
   local val = func:addLocal( W.typeAnyRef )
   func:setBody( {ls( len, W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) ) ), when_( p( W.opI32Eqz, {lg( len )} ), {W.ret( {W.refNull( W.htAny )} )} ), ls( items, W.structGet( self.tList, _moduleObj.fListItems, lg( 0 ) ) ), W.ifElse( p( W.opI64LeS, {lg( 1 ), i64c( 0 )} ), {ls( pos, p( W.opI32Sub, {lg( len ), i32c( 1 )} ) )}, {ls( pos, p( W.opI32WrapI64, {p( W.opI64Sub, {lg( 1 ), i64c( 1 )} )} ) )}, nil ), when_( p( W.opI32GeU, {lg( pos ), lg( len )} ), {W.ret( {W.refNull( W.htAny )} )} ), ls( val, W.arrayGet( self.tAnyArr, lg( items ), lg( pos ) ) ), W.arrayCopy( self.tAnyArr, self.tAnyArr, lg( items ), lg( pos ), lg( items ), p( W.opI32Add, {lg( pos ), i32c( 1 )} ), p( W.opI32Sub, {p( W.opI32Sub, {lg( len ), lg( pos )} ), i32c( 1 )} ) ), ls( len, p( W.opI32Sub, {lg( len ), i32c( 1 )} ) ), W.arraySet( self.tAnyArr, lg( items ), lg( len ), W.refNull( W.htAny ) ), W.structSet( self.tList, _moduleObj.fListLen, lg( 0 ), lg( len ) ), lg( val )} )
   return index
end

function Runtime:fnListClone(  )

   do
      local index = self:getFunc( "listClone" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "listClone", {self.listType}, {self.listType} )
   local list = func:addLocal( self.listType )
   local len = func:addLocal( W.typeI32 )
   func:setBody( {ls( len, W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) ) ), ls( list, W.call( self:fnListNew(  ), {lg( len )} ) ), W.arrayCopy( self.tAnyArr, self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, lg( list ) ), i32c( 0 ), W.structGet( self.tList, _moduleObj.fListItems, lg( 0 ) ), i32c( 0 ), lg( len ) ), W.structSet( self.tList, _moduleObj.fListLen, lg( list ), lg( len ) ), lg( list )} )
   return index
end

function Runtime:getLessCloInfo(  )

   return self:getCloInfo( {W.typeAnyRef, W.typeAnyRef}, {W.typeI32} )
end

function Runtime:fnListSort(  )

   do
      local index = self:getFunc( "listSort" )
      if index ~= nil then
         return index
      end
   end
   local lessInfo = self:getLessCloInfo(  )
   local index, func = self:newFunc( "listSort", {self.listType, W.refType( lessInfo:get_cloType() )}, {} )
   local items = func:addLocal( self.anyArrType )
   local len = func:addLocal( W.typeI32 )
   local gap = func:addLocal( W.typeI32 )
   local i = func:addLocal( W.typeI32 )
   local j = func:addLocal( W.typeI32 )
   local tmp = func:addLocal( W.typeAnyRef )
   local cmp = 1
   local less = W.ifElse( W.refIsNull( lg( cmp ) ), {W.call( self:fnLessAny(  ), {lg( tmp ), W.arrayGet( self.tAnyArr, lg( items ), p( W.opI32Sub, {lg( j ), lg( gap )} ) )} )}, {W.callRef( lessInfo:get_funcType(), {lg( cmp ), lg( tmp ), W.arrayGet( self.tAnyArr, lg( items ), p( W.opI32Sub, {lg( j ), lg( gap )} ) ), W.structGet( lessInfo:get_cloType(), _moduleObj.fCloCode, lg( cmp ) )} )}, W.typeI32 )
   func:setBody( {ls( items, W.structGet( self.tList, _moduleObj.fListItems, lg( 0 ) ) ), ls( len, W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) ) ), ls( gap, p( W.opI32ShrU, {lg( len ), i32c( 1 )} ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32Eqz, {lg( gap )} ) ), ls( i, lg( gap ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( i ), lg( len )} ) ), ls( tmp, W.arrayGet( self.tAnyArr, lg( items ), lg( i ) ) ), ls( j, lg( i ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32LtS, {lg( j ), lg( gap )} ) ), W.brIf( 1, p( W.opI32Eqz, {less} ) ), W.arraySet( self.tAnyArr, lg( items ), lg( j ), W.arrayGet( self.tAnyArr, lg( items ), p( W.opI32Sub, {lg( j ), lg( gap )} ) ) ), ls( j, p( W.opI32Sub, {lg( j ), lg( gap )} ) ), W.br( 0 )}, nil )}, nil ), W.arraySet( self.tAnyArr, lg( items ), lg( j ), lg( tmp ) ), incr( i ), W.br( 0 )}, nil )}, nil ), ls( gap, p( W.opI32ShrU, {lg( gap ), i32c( 1 )} ) ), W.br( 0 )}, nil )}, nil )} )
   return index
end

local luaFormat = 0
_moduleObj.luaFormat = luaFormat


local luaFind = 1
_moduleObj.luaFind = luaFind


local luaGsub = 2
_moduleObj.luaGsub = luaGsub


local luaTostring = 3
_moduleObj.luaTostring = luaTostring


local luaGmatch = 4
_moduleObj.luaGmatch = luaGmatch


local luaMatch = 5
_moduleObj.luaMatch = luaMatch


local luaTonumber = 6
_moduleObj.luaTonumber = luaTonumber


local luaByte = 7
_moduleObj.luaByte = luaByte


local luaRep = 8
_moduleObj.luaRep = luaRep


local luaLower = 9
_moduleObj.luaLower = luaLower


local luaUpper = 10
_moduleObj.luaUpper = luaUpper


local luaReverse = 11
_moduleObj.luaReverse = luaReverse


local luaLen = 12
_moduleObj.luaLen = luaLen


local luaChar = 13
_moduleObj.luaChar = luaChar


local luaGmatchAll = 14
_moduleObj.luaGmatchAll = luaGmatchAll


local luaReplace = 15
_moduleObj.luaReplace = luaReplace


local luaSub = 16
_moduleObj.luaSub = luaSub


local luaRandom = 17
_moduleObj.luaRandom = luaRandom


local luaRandomseed = 18
_moduleObj.luaRandomseed = luaRandomseed


local luaLoad = 19
_moduleObj.luaLoad = luaLoad


local luaRequire = 20
_moduleObj.luaRequire = luaRequire


local luaPackagePath = 21
_moduleObj.luaPackagePath = luaPackagePath


local luaSearchpath = 22
_moduleObj.luaSearchpath = luaSearchpath


local luaDump = 23
_moduleObj.luaDump = luaDump


local luaGlobals = 24
_moduleObj.luaGlobals = luaGlobals


local luaRunLuaOnLns = 25
_moduleObj.luaRunLuaOnLns = luaRunLuaOnLns


local luaPreload = 26
_moduleObj.luaPreload = luaPreload


local luaTTable = 5
_moduleObj.luaTTable = luaTTable


local luaTFunction = 6
_moduleObj.luaTFunction = luaTFunction


function Runtime:luaFunc( name, params, results )

   do
      local funcId = self.name2luaFunc[name]
      if funcId ~= nil then
         return funcId
      end
   end
   self.luaImported = true
   local funcId = self.mod:importFunc( _moduleObj.cModule, "lnsl_" .. name, params, results )
   self.mod:setFuncName( funcId, "lua." .. name )
   self.name2luaFunc[name] = funcId
   return funcId
end

function Runtime:fnLuaPushStr(  )

   do
      local index = self:getFunc( "luaPushStr" )
      if index ~= nil then
         return index
      end
   end
   local fnBuf = self:luaFunc( "buf", {W.typeI32}, {W.typeI32} )
   local fnPushStr = self:luaFunc( "push_str", {W.typeI32}, {} )
   local index, func = self:newFunc( "luaPushStr", {self.strType}, {} )
   local len = func:addLocal( W.typeI32 )
   local addr = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {ls( len, W.arrayLen( lg( 0 ) ) ), ls( addr, W.call( fnBuf, {lg( len )} ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( len )} ) ), W.i32Store8( p( W.opI32Add, {lg( addr ), lg( pos )} ), W.arrayGetU( self.tStr, lg( 0 ), lg( pos ) ) ), incr( pos ), W.br( 0 )}, nil )}, nil ), W.call( fnPushStr, {lg( len )} )} )
   return index
end

function Runtime:fnLuaToStr(  )

   do
      local index = self:getFunc( "luaToStr" )
      if index ~= nil then
         return index
      end
   end
   local fnToStr = self:luaFunc( "to_str", {W.typeI32}, {W.typeI32} )
   local fnStrLen = self:luaFunc( "str_len", {}, {W.typeI32} )
   local index, func = self:newFunc( "luaToStr", {W.typeI32}, {self.strType} )
   local addr = func:addLocal( W.typeI32 )
   local work = func:addLocal( self.strType )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {ls( addr, W.call( fnToStr, {lg( 0 )} ) ), ls( work, W.arrayNewDefault( self.tStr, W.call( fnStrLen, {} ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), W.arrayLen( lg( work ) )} ) ), W.arraySet( self.tStr, lg( work ), lg( pos ), W.i32Load8U( p( W.opI32Add, {lg( addr ), lg( pos )} ) ) ), incr( pos ), W.br( 0 )}, nil )}, nil ), lg( work )} )
   return index
end

function Runtime:fnLuaPushAny(  )

   do
      local index = self:getFunc( "luaPushAny" )
      if index ~= nil then
         return index
      end
   end
   local fnPushNil = self:luaFunc( "push_nil", {}, {} )
   local fnPushInt = self:luaFunc( "push_int", {W.typeI64}, {} )
   local fnPushReal = self:luaFunc( "push_real", {W.typeF64}, {} )
   local fnPushBool = self:luaFunc( "push_bool", {W.typeI32}, {} )
   local fnPushRef = self:luaFunc( "push_ref", {W.typeI32}, {} )
   local fnNewTable = self:luaFunc( "new_table", {}, {} )
   local fnSetTable = self:luaFunc( "set_table", {W.typeI32}, {} )
   local index, func = self:newFunc( "luaPushAny", {W.typeAnyRef}, {} )
   local list = func:addLocal( self.listType )
   local map = func:addLocal( self.mapType )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {when_( W.refIsNull( lg( 0 ) ), {W.call( fnPushNil, {} ), W.ret( {} )} ), when_( W.refTest( self.tLuaVal, lg( 0 ) ), {W.call( fnPushRef, {self:luaValRef( lg( 0 ) )} ), W.ret( {} )} ), when_( W.refTest( self.tList, lg( 0 ) ), {ls( list, W.refCast( self.tList, lg( 0 ) ) ), W.call( fnNewTable, {} ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), W.structGet( self.tList, _moduleObj.fListLen, lg( list ) )} ) ), incr( pos ), W.call( fnPushInt, {p( W.opI64ExtendI32S, {lg( pos )} )} ), W.call( index, {W.arrayGet( self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, lg( list ) ), p( W.opI32Sub, {lg( pos ), i32c( 1 )} ) )} ), W.call( fnSetTable, {i32c( -3 )} ), W.br( 0 )}, nil )}, nil ), W.ret( {} )} ), when_( W.refTest( self.tMap, lg( 0 ) ), {ls( map, W.refCast( self.tMap, lg( 0 ) ) ), W.call( fnNewTable, {} ), W.block( {W.loop( {ls( pos, W.call( self:fnMapNext(  ), {lg( map ), lg( pos )} ) ), W.brIf( 1, p( W.opI32LtS, {lg( pos ), i32c( 0 )} ) ), W.call( index, {self:mapKeyAt( lg( map ), lg( pos ) )} ), W.call( index, {self:mapValAt( lg( map ), lg( pos ) )} ), W.call( fnSetTable, {i32c( -3 )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), W.ret( {} )} ), when_( W.refTest( self.tBoxI64, lg( 0 ) ), {W.call( fnPushInt, {self:unboxI64( lg( 0 ) )} ), W.ret( {} )} ), when_( W.refTest( self.tBoxF64, lg( 0 ) ), {W.call( fnPushReal, {W.structGet( self.tBoxF64, 0, W.refCast( self.tBoxF64, lg( 0 ) ) )} ), W.ret( {} )} ), when_( W.refTest( self.tBoxBool, lg( 0 ) ), {W.call( fnPushBool, {self:unboxBool( lg( 0 ) )} ), W.ret( {} )} ), W.call( self:fnLuaPushStr(  ), {W.call( self:fnToStr(  ), {lg( 0 )} )} )} )
   return index
end

function Runtime:fnLuaGetAny(  )

   do
      local index = self:getFunc( "luaGetAny" )
      if index ~= nil then
         return index
      end
   end
   local fnKind = self:luaFunc( "kind", {W.typeI32}, {W.typeI32} )
   local fnToInt = self:luaFunc( "to_int", {W.typeI32}, {W.typeI64} )
   local fnToReal = self:luaFunc( "to_real", {W.typeI32}, {W.typeF64} )
   local fnToBool = self:luaFunc( "to_bool", {W.typeI32}, {W.typeI32} )
   local fnRef = self:luaFunc( "ref", {W.typeI32}, {W.typeI32} )
   local fnType = self:luaFunc( "type", {W.typeI32}, {W.typeI32} )
   local fnRegister = self:jsFunc( "luaval_new", {W.typeAnyRef, W.typeI32}, {} )
   local index, func = self:newFunc( "luaGetAny", {W.typeI32}, {W.typeAnyRef} )
   local kind = func:addLocal( W.typeI32 )
   local luaVal = func:addLocal( W.refType( self.tLuaVal ) )
   local function isKind( val )
   
      return p( W.opI32Eq, {lg( kind ), i32c( val )} )
   end
   func:setBody( {ls( kind, W.call( fnKind, {lg( 0 )} ) ), when_( isKind( 0 ), {W.ret( {W.refNull( W.htAny )} )} ), when_( isKind( 1 ), {W.ret( {self:boxI64( W.call( fnToInt, {lg( 0 )} ) )} )} ), when_( isKind( 2 ), {W.ret( {self:boxF64( W.call( fnToReal, {lg( 0 )} ) )} )} ), when_( isKind( 3 ), {W.ret( {self:boxBool( W.call( fnToBool, {lg( 0 )} ) )} )} ), when_( isKind( 4 ), {W.ret( {W.call( self:fnLuaToStr(  ), {lg( 0 )} )} )} ), ls( luaVal, W.structNew( self.tLuaVal, {i32c( 0 ), W.call( fnRef, {lg( 0 )} ), W.call( fnType, {lg( 0 )} )} ) ), W.call( fnRegister, {lg( luaVal ), self:luaValRef( lg( luaVal ) )} ), lg( luaVal )} )
   return index
end

function Runtime:fnLuaCall(  )

   do
      local index = self:getFunc( "luaCall" )
      if index ~= nil then
         return index
      end
   end
   local fnPushBuiltin = self:luaFunc( "push_builtin", {W.typeI32}, {} )
   local index, func = self:newFunc( "luaCall", {W.typeI32, self.listType}, {self.listType} )
   func:setBody( {W.call( fnPushBuiltin, {lg( 0 )} ), W.call( self:fnLuaCallPushed(  ), {lg( 1 )} )} )
   return index
end

function Runtime:fnLuaCallVal(  )

   do
      local index = self:getFunc( "luaCallVal" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "luaCallVal", {W.typeAnyRef, self.listType}, {self.listType} )
   func:setBody( {W.call( self:fnLuaPushAny(  ), {lg( 0 )} ), W.call( self:fnLuaCallPushed(  ), {lg( 1 )} )} )
   return index
end

function Runtime:fnLuaCallPushed(  )

   do
      local index = self:getFunc( "luaCallPushed" )
      if index ~= nil then
         return index
      end
   end
   local fnCall = self:luaFunc( "call", {W.typeI32}, {W.typeI32} )
   local fnSettop = self:luaFunc( "settop", {W.typeI32}, {} )
   local index, func = self:newFunc( "luaCallPushed", {self.listType}, {self.listType} )
   local argNum = func:addLocal( W.typeI32 )
   local num = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   local result = func:addLocal( self.listType )
   local mess = func:addLocal( self.strType )
   func:setBody( {ls( argNum, W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( argNum )} ) ), W.call( self:fnLuaPushAny(  ), {W.arrayGet( self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, lg( 0 ) ), lg( pos ) )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), ls( num, W.call( fnCall, {lg( argNum )} ) ), when_( p( W.opI32LtS, {lg( num ), i32c( 0 )} ), {ls( mess, W.call( self:fnLuaToStr(  ), {i32c( -1 )} ) ), W.call( fnSettop, {i32c( 0 )} ), self:error( lg( mess ) )} ), ls( result, W.call( self:fnListNew(  ), {lg( num )} ) ), ls( pos, i32c( 1 ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GtS, {lg( pos ), lg( num )} ) ), W.call( self:fnListPush(  ), {lg( result ), W.call( self:fnLuaGetAny(  ), {lg( pos )} )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), W.call( fnSettop, {i32c( 0 )} ), lg( result )} )
   return index
end

function Runtime:listItem( list, index )

   return W.arrayGet( self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, list ), i32c( index ) )
end

function Runtime:fnLuaGetAtLazy(  )

   do
      local index = self:getFunc( "luaGetAtLazy" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "luaGetAtLazy", {W.typeAnyRef, W.typeAnyRef}, {W.typeAnyRef} )
   self.lazyFuncMap["luaGetAtLazy"] = func
   return index
end

function Runtime:fnLuaCallValLazy(  )

   do
      local index = self:getFunc( "luaCallValLazy" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "luaCallValLazy", {W.typeAnyRef, self.listType}, {self.listType} )
   self.lazyFuncMap["luaCallValLazy"] = func
   return index
end

function Runtime:finalizeLazyLua(  )

   local lazyList = {{"luaGetAtLazy", "getAt"}, {"luaCallValLazy", "callVal"}}
   for __index, item in ipairs( lazyList ) do
      do
         local func = self.lazyFuncMap[item[1]]
         if  nil == func then
            local _func = func
         
            goto __continue1
         end
         
         local body = {}
         if self.luaImported then
            local target
            
            if item[2] == "getAt" then
               target = self:fnLuaGetAt(  )
            else
             
               target = self:fnLuaCallVal(  )
            end
            table.insert( body, W.call( target, {lg( 0 ), lg( 1 )} ) )
         else
          
            table.insert( body, W.unreachable(  ) )
         end
         func:setBody( body )
      end
      ::__continue1::
   end
end

function Runtime:luaValRef( val )

   return W.structGet( self.tLuaVal, 1, W.refCast( self.tLuaVal, val ) )
end

function Runtime:luaValType( val )

   return W.structGet( self.tLuaVal, 2, W.refCast( self.tLuaVal, val ) )
end

function Runtime:fnLuaGetAt(  )

   do
      local index = self:getFunc( "luaGetAt" )
      if index ~= nil then
         return index
      end
   end
   local fnPushRef = self:luaFunc( "push_ref", {W.typeI32}, {} )
   local fnGetTable = self:luaFunc( "get_table", {W.typeI32}, {} )
   local fnSettop = self:luaFunc( "settop", {W.typeI32}, {} )
   local index, func = self:newFunc( "luaGetAt", {W.typeAnyRef, W.typeAnyRef}, {W.typeAnyRef} )
   local result = func:addLocal( W.typeAnyRef )
   func:setBody( {W.call( fnPushRef, {self:luaValRef( lg( 0 ) )} ), W.call( self:fnLuaPushAny(  ), {lg( 1 )} ), W.call( fnGetTable, {i32c( -2 )} ), ls( result, W.call( self:fnLuaGetAny(  ), {i32c( -1 )} ) ), W.call( fnSettop, {i32c( 0 )} ), lg( result )} )
   return index
end

function Runtime:fnLuaLen(  )

   do
      local index = self:getFunc( "luaLen" )
      if index ~= nil then
         return index
      end
   end
   local fnPushRef = self:luaFunc( "push_ref", {W.typeI32}, {} )
   local fnLen = self:luaFunc( "len", {W.typeI32}, {W.typeI64} )
   local fnSettop = self:luaFunc( "settop", {W.typeI32}, {} )
   local index, func = self:newFunc( "luaLen", {W.typeAnyRef}, {W.typeI64} )
   local result = func:addLocal( W.typeI64 )
   func:setBody( {W.call( fnPushRef, {self:luaValRef( lg( 0 ) )} ), ls( result, W.call( fnLen, {i32c( -1 )} ) ), W.call( fnSettop, {i32c( 0 )} ), lg( result )} )
   return index
end

function Runtime:fnLuaToList(  )

   do
      local index = self:getFunc( "luaToList" )
      if index ~= nil then
         return index
      end
   end
   local fnPushRef = self:luaFunc( "push_ref", {W.typeI32}, {} )
   local fnPushInt = self:luaFunc( "push_int", {W.typeI64}, {} )
   local fnGetTable = self:luaFunc( "get_table", {W.typeI32}, {} )
   local fnSettop = self:luaFunc( "settop", {W.typeI32}, {} )
   local index, func = self:newFunc( "luaToList", {W.typeAnyRef}, {self.listType} )
   local list = func:addLocal( self.listType )
   local pos = func:addLocal( W.typeI64 )
   local val = func:addLocal( W.typeAnyRef )
   func:setBody( {ls( list, W.call( self:fnListNew(  ), {i32c( 4 )} ) ), W.call( fnPushRef, {self:luaValRef( lg( 0 ) )} ), W.block( {W.loop( {ls( pos, p( W.opI64Add, {lg( pos ), i64c( 1 )} ) ), W.call( fnPushInt, {lg( pos )} ), W.call( fnGetTable, {i32c( 1 )} ), ls( val, W.call( self:fnLuaGetAny(  ), {i32c( 2 )} ) ), W.call( fnSettop, {i32c( 1 )} ), W.brIf( 1, W.refIsNull( lg( val ) ) ), W.call( self:fnListPush(  ), {lg( list ), lg( val )} ), W.br( 0 )}, nil )}, nil ), W.call( fnSettop, {i32c( 0 )} ), lg( list )} )
   return index
end

function Runtime:fnLuaTableToMap(  )

   do
      local index = self:getFunc( "luaTableToMap" )
      if index ~= nil then
         return index
      end
   end
   local fnPushNil = self:luaFunc( "push_nil", {}, {} )
   local fnNext = self:luaFunc( "next", {W.typeI32}, {W.typeI32} )
   local fnGettop = self:luaFunc( "gettop", {}, {W.typeI32} )
   local fnSettop = self:luaFunc( "settop", {W.typeI32}, {} )
   local fnType = self:luaFunc( "type", {W.typeI32}, {W.typeI32} )
   local index, func = self:newFunc( "luaTableToMap", {W.typeI32, W.typeI32}, {self.mapType} )
   local map = func:addLocal( self.mapType )
   local top = func:addLocal( W.typeI32 )
   local key = func:addLocal( W.typeAnyRef )
   local val = func:addLocal( W.typeAnyRef )
   local function getVal( pos )
   
      return W.ifElse( p( W.opI32And, {lg( 1 ), p( W.opI32Eq, {W.call( fnType, {pos} ), i32c( _moduleObj.luaTTable )} )} ), {W.call( index, {pos, i32c( 1 )} )}, {W.call( self:fnLuaGetAny(  ), {pos} )}, W.typeAnyRef )
   end
   func:setBody( {ls( map, self:newMap(  ) ), W.call( fnPushNil, {} ), W.block( {W.loop( {W.brIf( 1, p( W.opI32Eqz, {W.call( fnNext, {lg( 0 )} )} ) ), ls( top, W.call( fnGettop, {} ) ), ls( key, getVal( p( W.opI32Sub, {lg( top ), i32c( 1 )} ) ) ), ls( val, getVal( lg( top ) ) ), W.call( self:fnMapSet(  ), {lg( map ), lg( key ), lg( val )} ), W.call( fnSettop, {p( W.opI32Sub, {lg( top ), i32c( 1 )} )} ), W.br( 0 )}, nil )}, nil ), lg( map )} )
   return index
end

function Runtime:fnLuaToMap(  )

   do
      local index = self:getFunc( "luaToMap" )
      if index ~= nil then
         return index
      end
   end
   local fnPushRef = self:luaFunc( "push_ref", {W.typeI32}, {} )
   local fnSettop = self:luaFunc( "settop", {W.typeI32}, {} )
   local index, func = self:newFunc( "luaToMap", {W.typeAnyRef, W.typeI32}, {self.mapType} )
   local map = func:addLocal( self.mapType )
   func:setBody( {W.call( fnPushRef, {self:luaValRef( lg( 0 ) )} ), ls( map, W.call( self:fnLuaTableToMap(  ), {i32c( 1 ), lg( 1 )} ) ), W.call( fnSettop, {i32c( 0 )} ), lg( map )} )
   return index
end

function Runtime:fnLuaExpand(  )

   do
      local index = self:getFunc( "luaExpand" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "luaExpand", {W.typeAnyRef}, {W.typeAnyRef} )
   func:setBody( {when_( W.refTest( self.tLuaVal, lg( 0 ) ), {when_( p( W.opI32Eq, {self:luaValType( lg( 0 ) ), i32c( _moduleObj.luaTTable )} ), {W.ret( {W.call( self:fnLuaToMap(  ), {lg( 0 ), i32c( 1 )} )} )} )} ), lg( 0 )} )
   return index
end

function Runtime:fnLuaLoadfile(  )

   do
      local index = self:getFunc( "luaLoadfile" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "luaLoadfile", {self.strType}, {self.listType} )
   local file = func:addLocal( self.objType )
   local mess = func:addLocal( self.strType )
   local txt = func:addLocal( self.strType )
   local args = func:addLocal( self.listType )
   func:setBody( {W.call( self:fnIoOpen(  ), {lg( 0 ), W.refNull( self.tStr )} ), ls( mess, W.nop(  ) ), ls( file, W.nop(  ) ), ls( args, W.call( self:fnListNew(  ), {i32c( 2 )} ) ), W.ifElse( W.refIsNull( lg( file ) ), {W.call( self:fnListPush(  ), {lg( args ), W.refNull( W.htAny )} ), W.call( self:fnListPush(  ), {lg( args ), W.call( self:fnStrConcat(  ), {self:strLiteral( "cannot open " ), lg( mess )} )} ), W.ret( {lg( args )} )}, nil, nil ), ls( txt, W.call( self:fnFileRead(  ), {lg( file ), self:strLiteral( "a" )} ) ), W.call( self:fnFileClose(  ), {lg( file )} ), W.call( self:fnListPush(  ), {lg( args ), lg( txt )} ), W.call( self:fnListPush(  ), {lg( args ), W.refNull( W.htAny )} ), W.call( self:fnListPush(  ), {lg( args ), W.call( self:fnStrConcat(  ), {self:strLiteral( "@" ), lg( 0 )} )} ), W.call( self:fnLuaCall(  ), {i32c( _moduleObj.luaLoad ), lg( args )} )} )
   return index
end

function Runtime:fnRealToStr(  )

   do
      local index = self:getFunc( "realToStr" )
      if index ~= nil then
         return index
      end
   end
   local fnPushBuiltin = self:luaFunc( "push_builtin", {W.typeI32}, {} )
   local fnPushReal = self:luaFunc( "push_real", {W.typeF64}, {} )
   local fnCall = self:luaFunc( "call", {W.typeI32}, {W.typeI32} )
   local fnSettop = self:luaFunc( "settop", {W.typeI32}, {} )
   local index, func = self:newFunc( "realToStr", {W.typeF64}, {self.strType} )
   local work = func:addLocal( self.strType )
   func:setBody( {W.call( fnPushBuiltin, {i32c( _moduleObj.luaTostring )} ), W.call( fnPushReal, {lg( 0 )} ), W.drop( W.call( fnCall, {i32c( 1 )} ) ), ls( work, W.call( self:fnLuaToStr(  ), {i32c( 1 )} ) ), W.call( fnSettop, {i32c( 0 )} ), lg( work )} )
   return index
end

function Runtime:fnFormatLua(  )

   do
      local index = self:getFunc( "formatLua" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "formatLua", {self.strType, self.listType}, {self.strType} )
   local args = func:addLocal( self.listType )
   func:setBody( {ls( args, W.call( self:fnListNew(  ), {i32c( 8 )} ) ), W.call( self:fnListPush(  ), {lg( args ), lg( 0 )} ), W.call( self:fnListAppend(  ), {lg( args ), lg( 1 )} ), W.refCast( self.tStr, self:listItem( W.call( self:fnLuaCall(  ), {i32c( _moduleObj.luaFormat ), lg( args )} ), 0 ) )} )
   return index
end

function Runtime:fnStrJoin(  )

   do
      local index = self:getFunc( "strJoin" )
      if index ~= nil then
         return index
      end
   end
   local strIdx = self.tStr
   local index, func = self:newFunc( "strJoin", {self.listType}, {self.strType} )
   local num = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   local total = func:addLocal( W.typeI32 )
   local work = func:addLocal( self.strType )
   local item = func:addLocal( self.strType )
   local function itemAt(  )
   
      return W.refCast( strIdx, W.arrayGet( self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, lg( 0 ) ), lg( pos ) ) )
   end
   func:setBody( {ls( num, W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( num )} ) ), ls( total, p( W.opI32Add, {lg( total ), W.arrayLen( itemAt(  ) )} ) ), incr( pos ), W.br( 0 )}, nil )}, nil ), ls( work, W.arrayNewDefault( strIdx, lg( total ) ) ), ls( total, i32c( 0 ) ), ls( pos, i32c( 0 ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( num )} ) ), ls( item, itemAt(  ) ), W.arrayCopy( strIdx, strIdx, lg( work ), lg( total ), lg( item ), i32c( 0 ), W.arrayLen( lg( item ) ) ), ls( total, p( W.opI32Add, {lg( total ), W.arrayLen( lg( item ) )} ) ), incr( pos ), W.br( 0 )}, nil )}, nil ), lg( work )} )
   return index
end

function Runtime:fnFormat(  )

   do
      local index = self:getFunc( "format" )
      if index ~= nil then
         return index
      end
   end
   local strIdx = self.tStr
   local index, func = self:newFunc( "format", {self.strType, self.listType}, {self.strType} )
   local len = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   local start = func:addLocal( W.typeI32 )
   local argPos = func:addLocal( W.typeI32 )
   local code = func:addLocal( W.typeI32 )
   local arg = func:addLocal( W.typeAnyRef )
   local pieces = func:addLocal( self.listType )
   local fallback = W.ret( {W.call( self:fnFormatLua(  ), {lg( 0 ), lg( 1 )} )} )
   local function pushSeg( from, to )
   
      return W.call( self:fnListPush(  ), {lg( pieces ), W.call( self:fnStrSub(  ), {lg( 0 ), p( W.opI64ExtendI32S, {from} ), p( W.opI64ExtendI32S, {to} )} )} )
   end
   func:setBody( {ls( len, W.arrayLen( lg( 0 ) ) ), ls( pieces, W.call( self:fnListNew(  ), {i32c( 8 )} ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( len )} ) ), when_( p( W.opI32Eq, {W.arrayGetU( strIdx, lg( 0 ), lg( pos ) ), i32c( 0x25 )} ), {when_( p( W.opI32GeS, {p( W.opI32Add, {lg( pos ), i32c( 1 )} ), lg( len )} ), {fallback} ), ls( code, W.arrayGetU( strIdx, lg( 0 ), p( W.opI32Add, {lg( pos ), i32c( 1 )} ) ) ), W.ifElse( p( W.opI32Eq, {lg( code ), i32c( 0x25 )} ), {pushSeg( p( W.opI32Add, {lg( start ), i32c( 1 )} ), p( W.opI32Add, {lg( pos ), i32c( 1 )} ) )}, {when_( p( W.opI32GeS, {lg( argPos ), W.structGet( self.tList, _moduleObj.fListLen, lg( 1 ) )} ), {fallback} ), ls( arg, W.arrayGet( self.tAnyArr, W.structGet( self.tList, _moduleObj.fListItems, lg( 1 ) ), lg( argPos ) ) ), incr( argPos ), pushSeg( p( W.opI32Add, {lg( start ), i32c( 1 )} ), lg( pos ) ), W.ifElse( p( W.opI32Eq, {lg( code ), i32c( 0x73 )} ), {when_( W.refTest( self.tLuaVal, lg( arg ) ), {fallback} ), W.call( self:fnListPush(  ), {lg( pieces ), W.call( self:fnToStr(  ), {lg( arg )} )} )}, {when_( p( W.opI32Eqz, {W.refTest( self.tBoxI64, lg( arg ) )} ), {fallback} ), W.ifElse( p( W.opI32Eq, {lg( code ), i32c( 0x64 )} ), {W.call( self:fnListPush(  ), {lg( pieces ), W.call( self:fnI64ToStr(  ), {self:unboxI64( lg( arg ) )} )} )}, {when_( p( W.opI32Ne, {lg( code ), i32c( 0x63 )} ), {fallback} ), W.call( self:fnListPush(  ), {lg( pieces ), W.arrayNewFixed( self.tStr, {p( W.opI32WrapI64, {self:unboxI64( lg( arg ) )} )} )} )}, nil )}, nil )}, nil ), ls( pos, p( W.opI32Add, {lg( pos ), i32c( 2 )} ) ), ls( start, lg( pos ) ), W.br( 1 )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), pushSeg( p( W.opI32Add, {lg( start ), i32c( 1 )} ), lg( len ) ), W.call( self:fnStrJoin(  ), {lg( pieces )} )} )
   return index
end

function Runtime:bufLen( buf )

   return p( W.opI32Or, {p( W.opI32Or, {W.arrayGetU( self.tStr, buf, i32c( 0 ) ), p( W.opI32Shl, {W.arrayGetU( self.tStr, buf, i32c( 1 ) ), i32c( 8 )} )} ), p( W.opI32Or, {p( W.opI32Shl, {W.arrayGetU( self.tStr, buf, i32c( 2 ) ), i32c( 16 )} ), p( W.opI32Shl, {W.arrayGetU( self.tStr, buf, i32c( 3 ) ), i32c( 24 )} )} )} )
end

function Runtime:setBufLen( buf, len )

   local list = {}
   for index = 0, 3 do
      table.insert( list, W.arraySet( self.tStr, buf, i32c( index ), p( W.opI32ShrU, {len, i32c( index * 8 )} ) ) )
   end
   return W.seq( list )
end

function Runtime:newBuf(  )

   return W.arrayNewDefault( self.tStr, i32c( 4 + 256 ) )
end

function Runtime:bufLenI64( buf )

   return p( W.opI64ExtendI32U, {self:bufLen( buf )} )
end

function Runtime:clearBuf( buf )

   return self:setBufLen( buf, i32c( 0 ) )
end

function Runtime:fnBufAdd(  )

   do
      local index = self:getFunc( "bufAdd" )
      if index ~= nil then
         return index
      end
   end
   local strIdx = self.tStr
   local index, func = self:newFunc( "bufAdd", {self.strType, self.strType}, {self.strType} )
   local len = func:addLocal( W.typeI32 )
   local need = func:addLocal( W.typeI32 )
   local cap = func:addLocal( W.typeI32 )
   local work = func:addLocal( self.strType )
   func:setBody( {ls( len, self:bufLen( lg( 0 ) ) ), ls( need, p( W.opI32Add, {p( W.opI32Add, {lg( len ), W.arrayLen( lg( 1 ) )} ), i32c( 4 )} ) ), when_( p( W.opI32GtU, {lg( need ), W.arrayLen( lg( 0 ) )} ), {ls( cap, p( W.opI32Shl, {W.arrayLen( lg( 0 ) ), i32c( 1 )} ) ), when_( p( W.opI32GtU, {lg( need ), lg( cap )} ), {ls( cap, lg( need ) )} ), ls( work, W.arrayNewDefault( strIdx, lg( cap ) ) ), W.arrayCopy( strIdx, strIdx, lg( work ), i32c( 0 ), lg( 0 ), i32c( 0 ), p( W.opI32Add, {lg( len ), i32c( 4 )} ) ), ls( 0, lg( work ) )} ), W.arrayCopy( strIdx, strIdx, lg( 0 ), p( W.opI32Add, {lg( len ), i32c( 4 )} ), lg( 1 ), i32c( 0 ), W.arrayLen( lg( 1 ) ) ), self:setBufLen( lg( 0 ), p( W.opI32Sub, {lg( need ), i32c( 4 )} ) ), lg( 0 )} )
   return index
end

function Runtime:fnBufToStr(  )

   do
      local index = self:getFunc( "bufToStr" )
      if index ~= nil then
         return index
      end
   end
   local strIdx = self.tStr
   local index, func = self:newFunc( "bufToStr", {self.strType}, {self.strType} )
   local work = func:addLocal( self.strType )
   func:setBody( {ls( work, W.arrayNewDefault( strIdx, self:bufLen( lg( 0 ) ) ) ), W.arrayCopy( strIdx, strIdx, lg( work ), i32c( 0 ), lg( 0 ), i32c( 4 ), W.arrayLen( lg( work ) ) ), lg( work )} )
   return index
end

function Runtime:fnStrLineList(  )

   do
      local index = self:getFunc( "strLineList" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "strLineList", {self.strType}, {self.listType} )
   local len = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   local start = func:addLocal( W.typeI32 )
   local list = func:addLocal( self.listType )
   local function pushSeg( to )
   
      return W.call( self:fnListPush(  ), {lg( list ), W.call( self:fnStrSub(  ), {lg( 0 ), p( W.opI64ExtendI32S, {p( W.opI32Add, {lg( start ), i32c( 1 )} )} ), p( W.opI64ExtendI32S, {to} )} )} )
   end
   func:setBody( {ls( len, W.arrayLen( lg( 0 ) ) ), ls( list, W.call( self:fnListNew(  ), {i32c( 8 )} ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( len )} ) ), incr( pos ), when_( p( W.opI32Eq, {W.arrayGetU( self.tStr, lg( 0 ), p( W.opI32Sub, {lg( pos ), i32c( 1 )} ) ), i32c( 0x0a )} ), {pushSeg( lg( pos ) ), ls( start, lg( pos ) )} ), W.br( 0 )}, nil )}, nil ), when_( p( W.opI32LtS, {lg( start ), lg( len )} ), {pushSeg( lg( len ) )} ), lg( list )} )
   return index
end

function Runtime:fnStrFind(  )

   do
      local index = self:getFunc( "strFind" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "strFind", {self.strType, self.strType, W.typeI64, W.typeI32}, {self.listType} )
   local args = func:addLocal( self.listType )
   local len = func:addLocal( W.typeI64 )
   local found = func:addLocal( W.typeI32 )
   local result = func:addLocal( self.listType )
   func:setBody( {when_( p( W.opI32Or, {lg( 3 ), p( W.opI32Eqz, {W.call( self:fnHasLuaSpecials(  ), {lg( 1 )} )} )} ), {ls( len, p( W.opI64ExtendI32U, {W.arrayLen( lg( 0 ) )} ) ), ls( 2, W.call( self:fnPosRelat(  ), {lg( 2 ), lg( len )} ) ), when_( p( W.opI64LtS, {lg( 2 ), i64c( 1 )} ), {ls( 2, i64c( 1 ) )} ), ls( result, W.call( self:fnListNew(  ), {i32c( 2 )} ) ), when_( p( W.opI64GtS, {lg( 2 ), p( W.opI64Add, {lg( len ), i64c( 1 )} )} ), {W.call( self:fnListPush(  ), {lg( result ), W.refNull( W.htAny )} ), W.ret( {lg( result )} )} ), ls( found, W.call( self:fnStrIndexOf(  ), {lg( 0 ), lg( 1 ), p( W.opI32WrapI64, {p( W.opI64Sub, {lg( 2 ), i64c( 1 )} )} )} ) ), W.ifElse( p( W.opI32LtS, {lg( found ), i32c( 0 )} ), {W.call( self:fnListPush(  ), {lg( result ), W.refNull( W.htAny )} )}, {W.call( self:fnListPush(  ), {lg( result ), self:boxI64( p( W.opI64ExtendI32S, {p( W.opI32Add, {lg( found ), i32c( 1 )} )} ) )} ), W.call( self:fnListPush(  ), {lg( result ), self:boxI64( p( W.opI64ExtendI32S, {p( W.opI32Add, {lg( found ), W.arrayLen( lg( 1 ) )} )} ) )} )}, nil ), W.ret( {lg( result )} )} ), ls( args, W.call( self:fnListNew(  ), {i32c( 4 )} ) ), W.call( self:fnListPush(  ), {lg( args ), lg( 0 )} ), W.call( self:fnListPush(  ), {lg( args ), lg( 1 )} ), W.call( self:fnListPush(  ), {lg( args ), self:boxI64( lg( 2 ) )} ), W.call( self:fnListPush(  ), {lg( args ), self:boxBool( lg( 3 ) )} ), W.call( self:fnLuaCall(  ), {i32c( _moduleObj.luaFind ), lg( args )} )} )
   return index
end

function Runtime:fnPosRelat(  )

   do
      local index = self:getFunc( "posRelat" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "posRelat", {W.typeI64, W.typeI64}, {W.typeI64} )
   func:setBody( {when_( p( W.opI64GeS, {lg( 0 ), i64c( 0 )} ), {W.ret( {lg( 0 )} )} ), when_( p( W.opI64GtS, {p( W.opI64Sub, {i64c( 0 ), lg( 0 )} ), lg( 1 )} ), {W.ret( {i64c( 0 )} )} ), p( W.opI64Add, {p( W.opI64Add, {lg( 1 ), lg( 0 )} ), i64c( 1 )} )} )
   return index
end

function Runtime:fnHasLuaSpecials(  )

   do
      local index = self:getFunc( "hasLuaSpecials" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "hasLuaSpecials", {self.strType}, {W.typeI32} )
   local pos = func:addLocal( W.typeI32 )
   local code = func:addLocal( W.typeI32 )
   local checks = {}
   for __index, ch in ipairs( {0x5e, 0x24, 0x2a, 0x2b, 0x3f, 0x2e, 0x28, 0x5b, 0x25, 0x2d} ) do
      table.insert( checks, when_( p( W.opI32Eq, {lg( code ), i32c( ch )} ), {W.ret( {i32c( 1 )} )} ) )
   end
   local loopBody = {W.brIf( 1, p( W.opI32GeS, {lg( pos ), W.arrayLen( lg( 0 ) )} ) ), ls( code, W.arrayGetU( self.tStr, lg( 0 ), lg( pos ) ) )}
   for __index, check in ipairs( checks ) do
      table.insert( loopBody, check )
   end
   table.insert( loopBody, incr( pos ) )
   table.insert( loopBody, W.br( 0 ) )
   func:setBody( {W.block( {W.loop( loopBody, nil )}, nil ), i32c( 0 )} )
   return index
end

function Runtime:fnStrIndexOf(  )

   do
      local index = self:getFunc( "strIndexOf" )
      if index ~= nil then
         return index
      end
   end
   local strIdx = self.tStr
   local index, func = self:newFunc( "strIndexOf", {self.strType, self.strType, W.typeI32}, {W.typeI32} )
   local last = func:addLocal( W.typeI32 )
   local ptnLen = func:addLocal( W.typeI32 )
   local first = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {ls( ptnLen, W.arrayLen( lg( 1 ) ) ), ls( last, p( W.opI32Sub, {W.arrayLen( lg( 0 ) ), lg( ptnLen )} ) ), when_( p( W.opI32Eqz, {lg( ptnLen )} ), {W.ret( {lg( 2 )} )} ), ls( first, W.arrayGetU( strIdx, lg( 1 ), i32c( 0 ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GtS, {lg( 2 ), lg( last )} ) ), when_( p( W.opI32Eq, {W.arrayGetU( strIdx, lg( 0 ), lg( 2 ) ), lg( first )} ), {ls( pos, i32c( 1 ) ), W.block( {W.loop( {when_( p( W.opI32GeS, {lg( pos ), lg( ptnLen )} ), {W.ret( {lg( 2 )} )} ), W.brIf( 1, p( W.opI32Ne, {W.arrayGetU( strIdx, lg( 0 ), p( W.opI32Add, {lg( 2 ), lg( pos )} ) ), W.arrayGetU( strIdx, lg( 1 ), lg( pos ) )} ) ), incr( pos ), W.br( 0 )}, nil )}, nil )} ), incr( 2 ), W.br( 0 )}, nil )}, nil ), i32c( -1 )} )
   return index
end

function Runtime:getNativeStrFunc( luaId )

   local name
   
   do
      local _switchExp = luaId
      if _switchExp == _moduleObj.luaByte then
         name = "strByte"
      elseif _switchExp == _moduleObj.luaRep then
         name = "strRep"
      elseif _switchExp == _moduleObj.luaLower then
         name = "strLower"
      elseif _switchExp == _moduleObj.luaUpper then
         name = "strUpper"
      elseif _switchExp == _moduleObj.luaReverse then
         name = "strReverse"
      elseif _switchExp == _moduleObj.luaLen then
         name = "strLen"
      elseif _switchExp == _moduleObj.luaReplace then
         name = "strReplace"
      elseif _switchExp == _moduleObj.luaFind then
         name = "strFindL"
      elseif _switchExp == _moduleObj.luaSub then
         name = "strSubL"
      else 
         
            return nil
      end
   end
   do
      local index = self:getFunc( name )
      if index ~= nil then
         return index
      end
   end
   local strIdx = self.tStr
   local index, func = self:newFunc( name, {self.listType}, {self.listType} )
   local txt = func:addLocal( self.strType )
   local len = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   local work = func:addLocal( self.strType )
   local result = func:addLocal( self.listType )
   local num1 = func:addLocal( W.typeI64 )
   local num2 = func:addLocal( W.typeI64 )
   local str1 = func:addLocal( self.strType )
   local str2 = func:addLocal( self.strType )
   local found = func:addLocal( W.typeI32 )
   local code = func:addLocal( W.typeI32 )
   local fallback = W.ret( {W.call( self:fnLuaCall(  ), {i32c( luaId ), lg( 0 )} )} )
   local argNum = W.structGet( self.tList, _moduleObj.fListLen, lg( 0 ) )
   local function arg( argIdx )
   
      return W.ifElse( p( W.opI32GtS, {argNum, i32c( argIdx )} ), {self:listItem( lg( 0 ), argIdx )}, {W.refNull( W.htAny )}, W.typeAnyRef )
   end
   local function optInt( argIdx, dst, def )
   
      return W.ifElse( W.refIsNull( arg( argIdx ) ), {ls( dst, def )}, {when_( p( W.opI32Eqz, {W.refTest( self.tBoxI64, arg( argIdx ) )} ), {fallback} ), ls( dst, self:unboxI64( arg( argIdx ) ) )}, nil )
   end
   local function needStr( argIdx, dst )
   
      return W.seq( {when_( p( W.opI32Eqz, {W.refTest( strIdx, arg( argIdx ) )} ), {fallback} ), ls( dst, W.refCast( strIdx, arg( argIdx ) ) )} )
   end
   local function mapBytes( conv )
   
      return {ls( work, W.arrayNewDefault( strIdx, lg( len ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( len )} ) ), ls( code, W.arrayGetU( strIdx, lg( txt ), lg( pos ) ) ), W.arraySet( strIdx, lg( work ), lg( pos ), conv ), incr( pos ), W.br( 0 )}, nil )}, nil )}
   end
   local function shiftRange( from, to, diff )
   
      return W.select( p( W.opI32Add, {lg( code ), i32c( diff )} ), lg( code ), p( W.opI32LeU, {p( W.opI32Sub, {lg( code ), i32c( from )} ), i32c( to - from )} ) )
   end
   local function ret1( val )
   
      return {ls( result, W.call( self:fnListNew(  ), {i32c( 1 )} ) ), W.call( self:fnListPush(  ), {lg( result ), val} ), lg( result )}
   end
   local body = {needStr( 0, txt ), ls( len, W.arrayLen( lg( txt ) ) )}
   do
      local _switchExp = luaId
      if _switchExp == _moduleObj.luaByte then
         table.insert( body, optInt( 1, num1, i64c( 1 ) ) )
         table.insert( body, ls( num1, W.call( self:fnPosRelat(  ), {lg( num1 ), p( W.opI64ExtendI32U, {lg( len )} )} ) ) )
         table.insert( body, optInt( 2, num2, lg( num1 ) ) )
         table.insert( body, ls( num2, W.call( self:fnPosRelat(  ), {lg( num2 ), p( W.opI64ExtendI32U, {lg( len )} )} ) ) )
         table.insert( body, when_( p( W.opI64LtS, {lg( num1 ), i64c( 1 )} ), {ls( num1, i64c( 1 ) )} ) )
         table.insert( body, when_( p( W.opI64GtS, {lg( num2 ), p( W.opI64ExtendI32U, {lg( len )} )} ), {ls( num2, p( W.opI64ExtendI32U, {lg( len )} ) )} ) )
         table.insert( body, ls( result, W.call( self:fnListNew(  ), {i32c( 1 )} ) ) )
         table.insert( body, W.block( {W.loop( {W.brIf( 1, p( W.opI64GtS, {lg( num1 ), lg( num2 )} ) ), W.call( self:fnListPush(  ), {lg( result ), self:boxI64( p( W.opI64ExtendI32U, {W.arrayGetU( strIdx, lg( txt ), p( W.opI32WrapI64, {p( W.opI64Sub, {lg( num1 ), i64c( 1 )} )} ) )} ) )} ), ls( num1, p( W.opI64Add, {lg( num1 ), i64c( 1 )} ) ), W.br( 0 )}, nil )}, nil ) )
         table.insert( body, lg( result ) )
      elseif _switchExp == _moduleObj.luaRep then
         table.insert( body, when_( p( W.opI32Eqz, {W.refTest( self.tBoxI64, arg( 1 ) )} ), {fallback} ) )
         table.insert( body, ls( num1, self:unboxI64( arg( 1 ) ) ) )
         table.insert( body, W.ifElse( W.refIsNull( arg( 2 ) ), {ls( str1, W.arrayNewDefault( strIdx, i32c( 0 ) ) )}, {needStr( 2, str1 )}, nil ) )
         table.insert( body, when_( p( W.opI64LeS, {lg( num1 ), i64c( 0 )} ), {W.ret( {W.seq( ret1( W.arrayNewDefault( strIdx, i32c( 0 ) ) ) )} )} ) )
         table.insert( body, ls( num2, p( W.opI64Sub, {p( W.opI64Mul, {lg( num1 ), p( W.opI64ExtendI32U, {p( W.opI32Add, {lg( len ), W.arrayLen( lg( str1 ) )} )} )} ), p( W.opI64ExtendI32U, {W.arrayLen( lg( str1 ) )} )} ) ) )
         table.insert( body, when_( p( W.opI32Or, {p( W.opI64GtS, {lg( num1 ), i64c( 0x10000000 )} ), p( W.opI64GtS, {lg( num2 ), i64c( 0x10000000 )} )} ), {fallback} ) )
         table.insert( body, ls( work, W.arrayNewDefault( strIdx, p( W.opI32WrapI64, {lg( num2 )} ) ) ) )
         table.insert( body, W.block( {W.loop( {W.arrayCopy( strIdx, strIdx, lg( work ), lg( pos ), lg( txt ), i32c( 0 ), lg( len ) ), ls( pos, p( W.opI32Add, {lg( pos ), lg( len )} ) ), ls( num1, p( W.opI64Sub, {lg( num1 ), i64c( 1 )} ) ), W.brIf( 1, p( W.opI64Eqz, {lg( num1 )} ) ), W.arrayCopy( strIdx, strIdx, lg( work ), lg( pos ), lg( str1 ), i32c( 0 ), W.arrayLen( lg( str1 ) ) ), ls( pos, p( W.opI32Add, {lg( pos ), W.arrayLen( lg( str1 ) )} ) ), W.br( 0 )}, nil )}, nil ) )
         append( body, ret1( lg( work ) ) )
      elseif _switchExp == _moduleObj.luaLower then
         append( body, mapBytes( shiftRange( 0x41, 0x5a, 0x20 ) ) )
         append( body, ret1( lg( work ) ) )
      elseif _switchExp == _moduleObj.luaUpper then
         append( body, mapBytes( shiftRange( 0x61, 0x7a, -0x20 ) ) )
         append( body, ret1( lg( work ) ) )
      elseif _switchExp == _moduleObj.luaReverse then
         table.insert( body, ls( work, W.arrayNewDefault( strIdx, lg( len ) ) ) )
         table.insert( body, W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( len )} ) ), W.arraySet( strIdx, lg( work ), lg( pos ), W.arrayGetU( strIdx, lg( txt ), p( W.opI32Sub, {p( W.opI32Sub, {lg( len ), lg( pos )} ), i32c( 1 )} ) ) ), incr( pos ), W.br( 0 )}, nil )}, nil ) )
         append( body, ret1( lg( work ) ) )
      elseif _switchExp == _moduleObj.luaLen then
         append( body, ret1( self:boxI64( p( W.opI64ExtendI32U, {lg( len )} ) ) ) )
      elseif _switchExp == _moduleObj.luaFind then
         table.insert( body, needStr( 1, str1 ) )
         table.insert( body, optInt( 2, num1, i64c( 1 ) ) )
         table.insert( body, ls( code, W.ifElse( W.refIsNull( arg( 3 ) ), {i32c( 0 )}, {W.ifElse( W.refTest( self.tBoxBool, arg( 3 ) ), {self:unboxBool( arg( 3 ) )}, {i32c( 1 )}, W.typeI32 )}, W.typeI32 ) ) )
         table.insert( body, W.call( self:fnStrFind(  ), {lg( txt ), lg( str1 ), lg( num1 ), lg( code )} ) )
      elseif _switchExp == _moduleObj.luaSub then
         table.insert( body, when_( p( W.opI32Eqz, {W.refTest( self.tBoxI64, arg( 1 ) )} ), {fallback} ) )
         table.insert( body, ls( num1, self:unboxI64( arg( 1 ) ) ) )
         table.insert( body, optInt( 2, num2, i64c( -1 ) ) )
         append( body, ret1( W.call( self:fnStrSub(  ), {lg( txt ), lg( num1 ), lg( num2 )} ) ) )
      elseif _switchExp == _moduleObj.luaReplace then
         table.insert( body, needStr( 1, str1 ) )
         table.insert( body, needStr( 2, str2 ) )
         table.insert( body, when_( p( W.opI32Eqz, {W.arrayLen( lg( str1 ) )} ), {fallback} ) )
         table.insert( body, ls( result, W.call( self:fnListNew(  ), {i32c( 8 )} ) ) )
         table.insert( body, W.block( {W.loop( {ls( found, W.call( self:fnStrIndexOf(  ), {lg( txt ), lg( str1 ), lg( pos )} ) ), W.brIf( 1, p( W.opI32LtS, {lg( found ), i32c( 0 )} ) ), W.call( self:fnListPush(  ), {lg( result ), W.call( self:fnStrSub(  ), {lg( txt ), p( W.opI64ExtendI32S, {p( W.opI32Add, {lg( pos ), i32c( 1 )} )} ), p( W.opI64ExtendI32S, {lg( found )} )} )} ), W.call( self:fnListPush(  ), {lg( result ), lg( str2 )} ), ls( pos, p( W.opI32Add, {lg( found ), W.arrayLen( lg( str1 ) )} ) ), W.br( 0 )}, nil )}, nil ) )
         table.insert( body, W.call( self:fnListPush(  ), {lg( result ), W.call( self:fnStrSub(  ), {lg( txt ), p( W.opI64ExtendI32S, {p( W.opI32Add, {lg( pos ), i32c( 1 )} )} ), p( W.opI64ExtendI32S, {lg( len )} )} )} ) )
         table.insert( body, ls( work, W.call( self:fnStrJoin(  ), {lg( result )} ) ) )
         append( body, ret1( lg( work ) ) )
      end
   end
   func:setBody( body )
   return index
end

function Runtime:fnStrGsub(  )

   do
      local index = self:getFunc( "strGsub" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "strGsub", {self.strType, self.strType, self.strType}, {self.strType, W.typeI64} )
   local result = func:addLocal( self.listType )
   local args = func:addLocal( self.listType )
   func:setBody( {ls( args, W.call( self:fnListNew(  ), {i32c( 4 )} ) ), W.call( self:fnListPush(  ), {lg( args ), lg( 0 )} ), W.call( self:fnListPush(  ), {lg( args ), lg( 1 )} ), W.call( self:fnListPush(  ), {lg( args ), lg( 2 )} ), ls( result, W.call( self:fnLuaCall(  ), {i32c( _moduleObj.luaGsub ), lg( args )} ) ), W.refCast( self.tStr, self:listItem( lg( result ), 0 ) ), self:unboxI64( self:listItem( lg( result ), 1 ) )} )
   return index
end

function Runtime:getFileType(  )

   if self.tFile < 0 then
      self.tFile = self.mod:addStructType( self.tObj, {W.field( W.typeI32, true ), W.field( W.refType( self.tVT ), false ), W.field( W.typeI32, false )} )
   end
   return self.tFile
end

function Runtime:hostFunc( name, params, results )

   do
      local index = self:getFunc( "host." .. name )
      if index ~= nil then
         return index
      end
   end
   local function isStr( valType )
   
      return valType == self.strType
   end
   local cParams = {}
   for __index, param in ipairs( params ) do
      if isStr( param ) then
         table.insert( cParams, W.typeI32 )
         table.insert( cParams, W.typeI32 )
      else
       
         table.insert( cParams, param )
      end
   end
   local strResult = #results == 1 and isStr( results[1] )
   local cResults = strResult and {W.typeI32} or results
   local cFunc = self.mod:importFunc( _moduleObj.cModule, "lnsr_" .. name, cParams, cResults )
   self.mod:setFuncName( cFunc, "lnsrt." .. name )
   local index, func = self:newFunc( "host." .. name, params, results )
   local addr = func:addLocal( W.typeI32 )
   local len = func:addLocal( W.typeI32 )
   local body = {}
   local args = {}
   local total = i32c( 0 )
   for pos, param in ipairs( params ) do
      if isStr( param ) then
         total = p( W.opI32Add, {total, W.arrayLen( lg( pos - 1 ) )} )
      end
   end
   table.insert( body, ls( addr, W.call( self.fnCBuf, {total} ) ) )
   local offset = lg( addr )
   for pos, param in ipairs( params ) do
      if isStr( param ) then
         table.insert( args, offset )
         table.insert( args, W.arrayLen( lg( pos - 1 ) ) )
         table.insert( body, W.call( self:fnCopyToC(  ), {lg( pos - 1 ), offset} ) )
         offset = p( W.opI32Add, {offset, W.arrayLen( lg( pos - 1 ) )} )
      else
       
         table.insert( args, lg( pos - 1 ) )
      end
   end
   if strResult then
      table.insert( body, ls( len, W.call( cFunc, args ) ) )
      table.insert( body, W.ifElse( p( W.opI32LtS, {lg( len ), i32c( 0 )} ), {W.refNull( self.tStr )}, {W.call( self:fnStrFromC(  ), {W.call( self.fnCRes, {} ), lg( len )} )}, self.strType ) )
   else
    
      table.insert( body, W.call( cFunc, args ) )
   end
   func:setBody( body )
   return index
end

function Runtime:jsFunc( name, params, results )

   do
      local index = self:getFunc( "js." .. name )
      if index ~= nil then
         return index
      end
   end
   local index = self.mod:importFunc( _moduleObj.jsModule, name, params, results )
   self.mod:setFuncName( index, "js." .. name )
   self.name2func["js." .. name] = index
   return index
end

function Runtime:fnCopyToC(  )

   do
      local index = self:getFunc( "copyToC" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "copyToC", {self.strType, W.typeI32}, {} )
   local len = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {ls( len, W.arrayLen( lg( 0 ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( len )} ) ), W.i32Store8( p( W.opI32Add, {lg( 1 ), lg( pos )} ), W.arrayGetU( self.tStr, lg( 0 ), lg( pos ) ) ), incr( pos ), W.br( 0 )}, nil )}, nil )} )
   return index
end

function Runtime:fnStrFromC(  )

   do
      local index = self:getFunc( "strFromC" )
      if index ~= nil then
         return index
      end
   end
   local index, func = self:newFunc( "strFromC", {W.typeI32, W.typeI32}, {self.strType} )
   local work = func:addLocal( self.strType )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {ls( work, W.arrayNewDefault( self.tStr, lg( 1 ) ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( 1 )} ) ), W.arraySet( self.tStr, lg( work ), lg( pos ), W.i32Load8U( p( W.opI32Add, {lg( 0 ), lg( pos )} ) ) ), incr( pos ), W.br( 0 )}, nil )}, nil ), lg( work )} )
   return index
end

function Runtime:fileFd( obj )

   local tFile = self:getFileType(  )
   return W.structGet( tFile, 2, W.refCast( tFile, obj ) )
end

function Runtime:fnIoOpen(  )

   do
      local index = self:getFunc( "ioOpen" )
      if index ~= nil then
         return index
      end
   end
   local fnOpen = self:hostFunc( "file_open", {self.strType, self.strType}, {W.typeI32} )
   local fnError = self:hostFunc( "file_error", {}, {self.strType} )
   local tFile = self:getFileType(  )
   local index, func = self:newFunc( "ioOpen", {self.strType, self.strType}, {self.objType, self.strType} )
   local fd = func:addLocal( W.typeI32 )
   func:setBody( {when_( W.refIsNull( lg( 1 ) ), {ls( 1, self:strLiteral( "r" ) )} ), ls( fd, W.call( fnOpen, {lg( 0 ), lg( 1 )} ) ), when_( p( W.opI32LtS, {lg( fd ), i32c( 0 )} ), {W.ret( {W.refNull( self.tObj ), W.call( fnError, {} )} )} ), W.structNew( tFile, {i32c( 0 ), W.structNew( self.tVT, {W.refNull( self.tAnyArr )} ), lg( fd )} ), W.refNull( self.tStr )} )
   return index
end

function Runtime:fnIoPopen(  )

   do
      local index = self:getFunc( "ioPopen" )
      if index ~= nil then
         return index
      end
   end
   local fnPopen = self:hostFunc( "file_popen", {self.strType}, {W.typeI32} )
   local tFile = self:getFileType(  )
   local index, func = self:newFunc( "ioPopen", {self.strType}, {self.objType} )
   local fd = func:addLocal( W.typeI32 )
   func:setBody( {ls( fd, W.call( fnPopen, {lg( 0 )} ) ), when_( p( W.opI32LtS, {lg( fd ), i32c( 0 )} ), {W.ret( {W.refNull( self.tObj )} )} ), W.structNew( tFile, {i32c( 0 ), W.structNew( self.tVT, {W.refNull( self.tAnyArr )} ), lg( fd )} )} )
   return index
end

function Runtime:fnIoStdin(  )

   do
      local index = self:getFunc( "ioStdin" )
      if index ~= nil then
         return index
      end
   end
   local fnStdin = self:hostFunc( "file_stdin", {}, {W.typeI32} )
   local tFile = self:getFileType(  )
   local index, func = self:newFunc( "ioStdin", {}, {self.objType} )
   func:setBody( {W.structNew( tFile, {i32c( 0 ), W.structNew( self.tVT, {W.refNull( self.tAnyArr )} ), W.call( fnStdin, {} )} )} )
   return index
end

function Runtime:fnArgList(  )

   do
      local index = self:getFunc( "argList" )
      if index ~= nil then
         return index
      end
   end
   local fnCount = self:hostFunc( "arg_count", {}, {W.typeI32} )
   local fnGet = self:hostFunc( "arg_get", {W.typeI32}, {self.strType} )
   local index, func = self:newFunc( "argList", {}, {self.listType} )
   local list = func:addLocal( self.listType )
   local num = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI32 )
   func:setBody( {ls( num, W.call( fnCount, {} ) ), ls( list, W.call( self:fnListNew(  ), {lg( num )} ) ), W.block( {W.loop( {W.brIf( 1, p( W.opI32GeS, {lg( pos ), lg( num )} ) ), W.call( self:fnListPush(  ), {lg( list ), W.call( fnGet, {lg( pos )} )} ), incr( pos ), W.br( 0 )}, nil )}, nil ), lg( list )} )
   return index
end

function Runtime:fnOsExit(  )

   return self:hostFunc( "os_exit", {W.typeI64}, {} )
end

function Runtime:fnOsFileOp( name )

   do
      local index = self:getFunc( "os_" .. name )
      if index ~= nil then
         return index
      end
   end
   local params = {self.strType}
   if name == "rename" then
      table.insert( params, self.strType )
   end
   local fnHost = self:hostFunc( "os_" .. name, params, {self.strType} )
   local index, func = self:newFunc( "os_" .. name, params, {W.typeAnyRef, self.strType} )
   local mess = func:addLocal( self.strType )
   local args = {}
   for pos, _1 in ipairs( params ) do
      table.insert( args, lg( pos - 1 ) )
   end
   func:setBody( {ls( mess, W.call( fnHost, args ) ), when_( W.refIsNull( lg( mess ) ), {W.ret( {self:boxBool( i32c( 1 ) ), W.refNull( self.tStr )} )} ), W.refNull( W.htAny ), lg( mess )} )
   return index
end

function Runtime:fnOsTime(  )

   return self:hostFunc( "os_time", {}, {W.typeI64} )
end

function Runtime:fnFileMtime(  )

   return self:hostFunc( "file_mtime", {self.strType}, {W.typeF64} )
end

function Runtime:fnOsClock(  )

   return self:hostFunc( "os_clock", {}, {W.typeF64} )
end

function Runtime:isFile( obj )

   return W.refTest( self:getFileType(  ), obj )
end

function Runtime:fnFileRead(  )

   do
      local index = self:getFunc( "fileRead" )
      if index ~= nil then
         return index
      end
   end
   local fnRead = self:hostFunc( "file_read", {W.typeI32, W.typeI32, W.typeI64}, {self.strType} )
   local index, func = self:newFunc( "fileRead", {W.typeAnyRef, W.typeAnyRef}, {self.strType} )
   local kind = func:addLocal( W.typeI32 )
   local num = func:addLocal( W.typeI64 )
   local fmt = func:addLocal( self.strType )
   local ch = func:addLocal( W.typeI32 )
   func:setBody( {ls( kind, i32c( 1 ) ), when_( W.refTest( self.tBoxI64, lg( 1 ) ), {ls( kind, i32c( 3 ) ), ls( num, self:unboxI64( lg( 1 ) ) )} ), when_( W.refTest( self.tStr, lg( 1 ) ), {ls( fmt, W.refCast( self.tStr, lg( 1 ) ) ), ls( ch, W.arrayGetU( self.tStr, lg( fmt ), i32c( 0 ) ) ), when_( p( W.opI32Eq, {lg( ch ), i32c( 0x2a )} ), {ls( ch, W.arrayGetU( self.tStr, lg( fmt ), i32c( 1 ) ) )} ), when_( p( W.opI32Eq, {lg( ch ), i32c( 0x61 )} ), {ls( kind, i32c( 0 ) )} ), when_( p( W.opI32Eq, {lg( ch ), i32c( 0x4c )} ), {ls( kind, i32c( 2 ) )} )} ), W.call( fnRead, {self:fileFd( lg( 0 ) ), lg( kind ), lg( num )} )} )
   return index
end

function Runtime:fnFileWrite(  )

   do
      local index = self:getFunc( "fileWrite" )
      if index ~= nil then
         return index
      end
   end
   local fnWrite = self:hostFunc( "file_write", {W.typeI32, self.strType}, {} )
   local index, func = self:newFunc( "fileWrite", {W.typeAnyRef, self.strType}, {} )
   func:setBody( {W.call( fnWrite, {self:fileFd( lg( 0 ) ), lg( 1 )} )} )
   return index
end

function Runtime:fnFileSeek(  )

   do
      local index = self:getFunc( "fileSeek" )
      if index ~= nil then
         return index
      end
   end
   local fnSeek = self:hostFunc( "file_seek", {W.typeI32, W.typeI32, W.typeI64}, {W.typeI64} )
   local index, func = self:newFunc( "fileSeek", {W.typeAnyRef, self.strType, W.typeI64}, {self.listType} )
   local whence = func:addLocal( W.typeI32 )
   local ch = func:addLocal( W.typeI32 )
   local pos = func:addLocal( W.typeI64 )
   local result = func:addLocal( self.listType )
   func:setBody( {ls( whence, i32c( 1 ) ), ls( ch, W.arrayGetU( self.tStr, lg( 1 ), i32c( 0 ) ) ), when_( p( W.opI32Eq, {lg( ch ), i32c( 0x73 )} ), {ls( whence, i32c( 0 ) )} ), when_( p( W.opI32Eq, {lg( ch ), i32c( 0x65 )} ), {ls( whence, i32c( 2 ) )} ), ls( pos, W.call( fnSeek, {self:fileFd( lg( 0 ) ), lg( whence ), lg( 2 )} ) ), ls( result, W.call( self:fnListNew(  ), {i32c( 2 )} ) ), W.ifElse( p( W.opI64LtS, {lg( pos ), i64c( 0 )} ), {W.call( self:fnListPush(  ), {lg( result ), W.refNull( W.htAny )} ), W.call( self:fnListPush(  ), {lg( result ), self:strLiteral( "Invalid argument" )} )}, {W.call( self:fnListPush(  ), {lg( result ), self:boxI64( lg( pos ) )} )}, nil ), lg( result )} )
   return index
end

function Runtime:fnFileClose(  )

   do
      local index = self:getFunc( "fileClose" )
      if index ~= nil then
         return index
      end
   end
   local fnClose = self:hostFunc( "file_close", {W.typeI32}, {} )
   local index, func = self:newFunc( "fileClose", {W.typeAnyRef}, {} )
   func:setBody( {W.call( fnClose, {self:fileFd( lg( 0 ) )} )} )
   return index
end


return _moduleObj
