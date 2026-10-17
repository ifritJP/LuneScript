--lune/base/WasmBin.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@WasmBin'
local _lune = {}
if _lune8 then
   _lune = _lune8
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


local function byte( val )

   return string.format( "%c", val & 0xff)
end
_moduleObj.byte = byte

local function uleb( val )

   local work = val
   local txt = ""
   while true do
      local bt = work & 0x7f
      work = work >> 7
      if work ~= 0 then
         bt = bt | 0x80
      end
      txt = txt .. byte( bt )
      if work == 0 then
         break
      end
   end
   return txt
end
_moduleObj.uleb = uleb

local function sleb( val )

   local work = val
   local txt = ""
   while true do
      local bt = work & 0x7f
      work = work >> 7
      if val < 0 then
         work = work | (-1 << 57 )
      end
      local done
      
      if (work == 0 and bt & 0x40 == 0 ) or (work == -1 and bt & 0x40 ~= 0 ) then
         done = true
      else
       
         done = false
      end
      if done then
         txt = txt .. byte( bt )
         break
      end
      txt = txt .. byte( bt | 0x80 )
   end
   return txt
end
_moduleObj.sleb = sleb

local function int64LE( val )

   local work = val
   local txt = ""
   for _1 = 1, 8 do
      txt = txt .. byte( work & 0xff )
      work = work >> 8
   end
   return txt
end
_moduleObj.int64LE = int64LE

local function f64bits( val )

   if val ~= val then
      return 0x7ff8000000000000
   end
   local work = val
   local sign = 0
   if work < 0.0 or (work == 0.0 and 1.0 / work < 0.0 ) then
      sign = 1
      work = -work
   end
   if work == 0.0 then
      return sign << 63
   end
   if work > 1.7976931348623157e308 then
      return (sign << 63 ) | (0x7ff << 52 )
   end
   local exp = 0
   while work >= 2.0 do
      work = work / 2.0
      exp = exp + 1
   end
   while work < 1.0 and exp > -1022 do
      work = work * 2.0
      exp = exp - 1
   end
   local two52 = 4503599627370496.0
   if work < 1.0 then
      local mant = math.floor((work * two52 ))
      return (sign << 63 ) | mant
   end
   local mant = math.floor(((work - 1.0 ) * two52 ))
   return (sign << 63 ) | ((exp + 1023 ) << 52 ) | mant
end
_moduleObj.f64bits = f64bits

local function name( txt )

   return uleb( #txt ) .. txt
end
_moduleObj.name = name

local function vec( list )

   local txt = uleb( #list )
   for __index, item in ipairs( list ) do
      txt = txt .. item
   end
   return txt
end
_moduleObj.vec = vec

local function section( id, body )

   return byte( id ) .. uleb( #body ) .. body
end
_moduleObj.section = section

local ValKind = {}
_moduleObj.ValKind = ValKind
ValKind._val2NameMap = {}
function ValKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ValKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ValKind._from( val )
   if ValKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ValKind.__allList = {}
function ValKind.get__allList()
   return ValKind.__allList
end

ValKind.I32 = 0
ValKind._val2NameMap[0] = 'I32'
ValKind.__allList[1] = ValKind.I32
ValKind.I64 = 1
ValKind._val2NameMap[1] = 'I64'
ValKind.__allList[2] = ValKind.I64
ValKind.F64 = 2
ValKind._val2NameMap[2] = 'F64'
ValKind.__allList[3] = ValKind.F64
ValKind.Ref = 3
ValKind._val2NameMap[3] = 'Ref'
ValKind.__allList[4] = ValKind.Ref

local htAny = -18
_moduleObj.htAny = htAny


local htEq = -19
_moduleObj.htEq = htEq


local htI31 = -20
_moduleObj.htI31 = htI31


local htStruct = -21
_moduleObj.htStruct = htStruct


local htArray = -22
_moduleObj.htArray = htArray


local htFunc = -16
_moduleObj.htFunc = htFunc


local htExtern = -17
_moduleObj.htExtern = htExtern


local htNone = -15
_moduleObj.htNone = htNone


local ValType = {}
_moduleObj.ValType = ValType
function ValType:encode(  )

   do
      local _switchExp = self.kind
      if _switchExp == ValKind.I32 then
         return byte( 0x7f )
      elseif _switchExp == ValKind.I64 then
         return byte( 0x7e )
      elseif _switchExp == ValKind.F64 then
         return byte( 0x7c )
      elseif _switchExp == ValKind.Ref then
         if self.nullable then
            return byte( 0x63 ) .. sleb( self.heapType )
         end
         return byte( 0x64 ) .. sleb( self.heapType )
      end
   end
end
function ValType:isRef(  )

   return self.kind == ValKind.Ref
end
function ValType:equalsType( other )

   return self.kind == other.kind and self.heapType == other.heapType and self.nullable == other.nullable
end
function ValType._setmeta( obj )
  setmetatable( obj, { __index = ValType  } )
end
function ValType._new( kind, heapType, nullable )
   local obj = {}
   ValType._setmeta( obj )
   if obj.__init then
      obj:__init( kind, heapType, nullable )
   end
   return obj
end
function ValType:__init( kind, heapType, nullable )

   self.kind = kind
   self.heapType = heapType
   self.nullable = nullable
end
function ValType:get_kind()
   return self.kind
end
function ValType:get_heapType()
   return self.heapType
end
function ValType:get_nullable()
   return self.nullable
end


local typeI32 = ValType._new(ValKind.I32, 0, false)
_moduleObj.typeI32 = typeI32


local typeI64 = ValType._new(ValKind.I64, 0, false)
_moduleObj.typeI64 = typeI64


local typeF64 = ValType._new(ValKind.F64, 0, false)
_moduleObj.typeF64 = typeF64


local function refType( heapType )

   return ValType._new(ValKind.Ref, heapType, true)
end
_moduleObj.refType = refType

local typeAnyRef = refType( _moduleObj.htAny )
_moduleObj.typeAnyRef = typeAnyRef


local typeEqRef = refType( _moduleObj.htEq )
_moduleObj.typeEqRef = typeEqRef


local FieldType = {}
_moduleObj.FieldType = FieldType
function FieldType:encode(  )

   local txt
   
   do
      local valType = self.valType
      if valType ~= nil then
         txt = valType:encode(  )
      else
         txt = byte( 0x78 )
      end
   end
   if self.mutable then
      return txt .. byte( 0x01 )
   end
   return txt .. byte( 0x00 )
end
function FieldType._setmeta( obj )
  setmetatable( obj, { __index = FieldType  } )
end
function FieldType._new( valType, packedI8, mutable )
   local obj = {}
   FieldType._setmeta( obj )
   if obj.__init then
      obj:__init( valType, packedI8, mutable )
   end
   return obj
end
function FieldType:__init( valType, packedI8, mutable )

   self.valType = valType
   self.packedI8 = packedI8
   self.mutable = mutable
end
function FieldType:get_valType()
   return self.valType
end
function FieldType:get_packedI8()
   return self.packedI8
end
function FieldType:get_mutable()
   return self.mutable
end


local function field( valType, mutable )

   return FieldType._new(valType, false, mutable)
end
_moduleObj.field = field

local function fieldI8( mutable )

   return FieldType._new(nil, true, mutable)
end
_moduleObj.fieldI8 = fieldI8

local opUnreachable = 0x00
_moduleObj.opUnreachable = opUnreachable


local opNop = 0x01
_moduleObj.opNop = opNop


local opDrop = 0x1a
_moduleObj.opDrop = opDrop


local opSelect = 0x1b
_moduleObj.opSelect = opSelect


local opI32Eqz = 0x45
_moduleObj.opI32Eqz = opI32Eqz


local opI32Eq = 0x46
_moduleObj.opI32Eq = opI32Eq


local opI32Ne = 0x47
_moduleObj.opI32Ne = opI32Ne


local opI32LtS = 0x48
_moduleObj.opI32LtS = opI32LtS


local opI32LtU = 0x49
_moduleObj.opI32LtU = opI32LtU


local opI32GtS = 0x4a
_moduleObj.opI32GtS = opI32GtS


local opI32GtU = 0x4b
_moduleObj.opI32GtU = opI32GtU


local opI32LeS = 0x4c
_moduleObj.opI32LeS = opI32LeS


local opI32LeU = 0x4d
_moduleObj.opI32LeU = opI32LeU


local opI32GeS = 0x4e
_moduleObj.opI32GeS = opI32GeS


local opI32GeU = 0x4f
_moduleObj.opI32GeU = opI32GeU


local opI64Eqz = 0x50
_moduleObj.opI64Eqz = opI64Eqz


local opI64Eq = 0x51
_moduleObj.opI64Eq = opI64Eq


local opI64Ne = 0x52
_moduleObj.opI64Ne = opI64Ne


local opI64LtS = 0x53
_moduleObj.opI64LtS = opI64LtS


local opI64LtU = 0x54
_moduleObj.opI64LtU = opI64LtU


local opI64GtS = 0x55
_moduleObj.opI64GtS = opI64GtS


local opI64GtU = 0x56
_moduleObj.opI64GtU = opI64GtU


local opI64LeS = 0x57
_moduleObj.opI64LeS = opI64LeS


local opI64LeU = 0x58
_moduleObj.opI64LeU = opI64LeU


local opI64GeS = 0x59
_moduleObj.opI64GeS = opI64GeS


local opI64GeU = 0x5a
_moduleObj.opI64GeU = opI64GeU


local opF64Eq = 0x61
_moduleObj.opF64Eq = opF64Eq


local opF64Ne = 0x62
_moduleObj.opF64Ne = opF64Ne


local opF64Lt = 0x63
_moduleObj.opF64Lt = opF64Lt


local opF64Gt = 0x64
_moduleObj.opF64Gt = opF64Gt


local opF64Le = 0x65
_moduleObj.opF64Le = opF64Le


local opF64Ge = 0x66
_moduleObj.opF64Ge = opF64Ge


local opI32Add = 0x6a
_moduleObj.opI32Add = opI32Add


local opI32Sub = 0x6b
_moduleObj.opI32Sub = opI32Sub


local opI32Mul = 0x6c
_moduleObj.opI32Mul = opI32Mul


local opI32DivS = 0x6d
_moduleObj.opI32DivS = opI32DivS


local opI32DivU = 0x6e
_moduleObj.opI32DivU = opI32DivU


local opI32RemS = 0x6f
_moduleObj.opI32RemS = opI32RemS


local opI32RemU = 0x70
_moduleObj.opI32RemU = opI32RemU


local opI32And = 0x71
_moduleObj.opI32And = opI32And


local opI32Or = 0x72
_moduleObj.opI32Or = opI32Or


local opI32Xor = 0x73
_moduleObj.opI32Xor = opI32Xor


local opI32Shl = 0x74
_moduleObj.opI32Shl = opI32Shl


local opI32ShrS = 0x75
_moduleObj.opI32ShrS = opI32ShrS


local opI32ShrU = 0x76
_moduleObj.opI32ShrU = opI32ShrU


local opI64Add = 0x7c
_moduleObj.opI64Add = opI64Add


local opI64Sub = 0x7d
_moduleObj.opI64Sub = opI64Sub


local opI64Mul = 0x7e
_moduleObj.opI64Mul = opI64Mul


local opI64DivS = 0x7f
_moduleObj.opI64DivS = opI64DivS


local opI64DivU = 0x80
_moduleObj.opI64DivU = opI64DivU


local opI64RemS = 0x81
_moduleObj.opI64RemS = opI64RemS


local opI64RemU = 0x82
_moduleObj.opI64RemU = opI64RemU


local opI64And = 0x83
_moduleObj.opI64And = opI64And


local opI64Or = 0x84
_moduleObj.opI64Or = opI64Or


local opI64Xor = 0x85
_moduleObj.opI64Xor = opI64Xor


local opI64Shl = 0x86
_moduleObj.opI64Shl = opI64Shl


local opI64ShrS = 0x87
_moduleObj.opI64ShrS = opI64ShrS


local opI64ShrU = 0x88
_moduleObj.opI64ShrU = opI64ShrU


local opF64Abs = 0x99
_moduleObj.opF64Abs = opF64Abs


local opF64Neg = 0x9a
_moduleObj.opF64Neg = opF64Neg


local opF64Ceil = 0x9b
_moduleObj.opF64Ceil = opF64Ceil


local opF64Floor = 0x9c
_moduleObj.opF64Floor = opF64Floor


local opF64Trunc = 0x9d
_moduleObj.opF64Trunc = opF64Trunc


local opF64Sqrt = 0x9f
_moduleObj.opF64Sqrt = opF64Sqrt


local opF64Add = 0xa0
_moduleObj.opF64Add = opF64Add


local opF64Sub = 0xa1
_moduleObj.opF64Sub = opF64Sub


local opF64Mul = 0xa2
_moduleObj.opF64Mul = opF64Mul


local opF64Div = 0xa3
_moduleObj.opF64Div = opF64Div


local opI32WrapI64 = 0xa7
_moduleObj.opI32WrapI64 = opI32WrapI64


local opI64ExtendI32S = 0xac
_moduleObj.opI64ExtendI32S = opI64ExtendI32S


local opI64ExtendI32U = 0xad
_moduleObj.opI64ExtendI32U = opI64ExtendI32U


local opI64TruncF64S = 0xb0
_moduleObj.opI64TruncF64S = opI64TruncF64S


local opF64ConvertI64S = 0xb9
_moduleObj.opF64ConvertI64S = opF64ConvertI64S


local opI64ReinterpretF64 = 0xbd
_moduleObj.opI64ReinterpretF64 = opI64ReinterpretF64


local opF64ReinterpretI64 = 0xbf
_moduleObj.opF64ReinterpretI64 = opF64ReinterpretI64


local opRefIsNull = 0xd1
_moduleObj.opRefIsNull = opRefIsNull


local opRefEq = 0xd3
_moduleObj.opRefEq = opRefEq


local InsKind = {}
_moduleObj.InsKind = InsKind
InsKind._val2NameMap = {}
function InsKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "InsKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function InsKind._from( val )
   if InsKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
InsKind.__allList = {}
function InsKind.get__allList()
   return InsKind.__allList
end

InsKind.Plain = 0
InsKind._val2NameMap[0] = 'Plain'
InsKind.__allList[1] = InsKind.Plain
InsKind.Seq = 1
InsKind._val2NameMap[1] = 'Seq'
InsKind.__allList[2] = InsKind.Seq
InsKind.Block = 2
InsKind._val2NameMap[2] = 'Block'
InsKind.__allList[3] = InsKind.Block
InsKind.Loop = 3
InsKind._val2NameMap[3] = 'Loop'
InsKind.__allList[4] = InsKind.Loop
InsKind.If = 4
InsKind._val2NameMap[4] = 'If'
InsKind.__allList[5] = InsKind.If
InsKind.FuncRef = 5
InsKind._val2NameMap[5] = 'FuncRef'
InsKind.__allList[6] = InsKind.FuncRef

local Ins = {}
_moduleObj.Ins = Ins
function Ins:encodeTo( stream, funcIndexList )

   for __index, arg in ipairs( self.args ) do
      arg:encodeTo( stream, funcIndexList )
   end
   local blockType
   
   do
      local result = self.result
      if result ~= nil then
         blockType = result:encode(  )
      else
         blockType = byte( 0x40 )
      end
   end
   do
      local _switchExp = self.kind
      if _switchExp == InsKind.Plain then
         stream:write( self.code )
      elseif _switchExp == InsKind.FuncRef then
         stream:write( self.code .. uleb( funcIndexList[self.funcId + 1] ) )
      elseif _switchExp == InsKind.Seq then
         for __index, ins in ipairs( self.body ) do
            ins:encodeTo( stream, funcIndexList )
         end
      elseif _switchExp == InsKind.Block or _switchExp == InsKind.Loop then
         if self.kind == InsKind.Block then
            stream:write( byte( 0x02 ) )
         else
          
            stream:write( byte( 0x03 ) )
         end
         stream:write( blockType )
         for __index, ins in ipairs( self.body ) do
            ins:encodeTo( stream, funcIndexList )
         end
         stream:write( byte( 0x0b ) )
      elseif _switchExp == InsKind.If then
         stream:write( byte( 0x04 ) .. blockType )
         for __index, ins in ipairs( self.body ) do
            ins:encodeTo( stream, funcIndexList )
         end
         do
            local elseBody = self.elseBody
            if elseBody ~= nil then
               stream:write( byte( 0x05 ) )
               for __index, ins in ipairs( elseBody ) do
                  ins:encodeTo( stream, funcIndexList )
               end
            end
         end
         stream:write( byte( 0x0b ) )
      end
   end
end
function Ins._setmeta( obj )
  setmetatable( obj, { __index = Ins  } )
end
function Ins._new( kind, code, args, body, elseBody, result, funcId )
   local obj = {}
   Ins._setmeta( obj )
   if obj.__init then
      obj:__init( kind, code, args, body, elseBody, result, funcId )
   end
   return obj
end
function Ins:__init( kind, code, args, body, elseBody, result, funcId )

   self.kind = kind
   self.code = code
   self.args = args
   self.body = body
   self.elseBody = elseBody
   self.result = result
   self.funcId = funcId
end
function Ins:get_kind()
   return self.kind
end
function Ins:get_code()
   return self.code
end
function Ins:get_args()
   return self.args
end
function Ins:get_body()
   return self.body
end
function Ins:get_elseBody()
   return self.elseBody
end
function Ins:get_result()
   return self.result
end
function Ins:get_funcId()
   return self.funcId
end

local function op( code, args )

   return Ins._new(InsKind.Plain, code, args, {}, nil, nil, 0)
end
_moduleObj.op = op

local function plain( code, args )

   return Ins._new(InsKind.Plain, byte( code ), args, {}, nil, nil, 0)
end
_moduleObj.plain = plain

local function seq( body )

   return Ins._new(InsKind.Seq, "", {}, body, nil, nil, 0)
end
_moduleObj.seq = seq

local function block( body, result )

   return Ins._new(InsKind.Block, "", {}, body, nil, result, 0)
end
_moduleObj.block = block

local function loop( body, result )

   return Ins._new(InsKind.Loop, "", {}, body, nil, result, 0)
end
_moduleObj.loop = loop

local function ifElse( cond, thenBody, elseBody, result )

   return Ins._new(InsKind.If, "", {cond}, thenBody, elseBody, result, 0)
end
_moduleObj.ifElse = ifElse


local function i32Const( val )

   local work = val & 0xffffffff
   if work >= 0x80000000 then
      work = work - 0x100000000
   end
   return op( byte( 0x41 ) .. sleb( work ), {} )
end
_moduleObj.i32Const = i32Const

local function i64Const( val )

   return op( byte( 0x42 ) .. sleb( val ), {} )
end
_moduleObj.i64Const = i64Const

local function f64Const( val )

   return op( byte( 0x44 ) .. int64LE( f64bits( val ) ), {} )
end
_moduleObj.f64Const = f64Const

local function localGet( index )

   return op( byte( 0x20 ) .. uleb( index ), {} )
end
_moduleObj.localGet = localGet

local function localSet( index, val )

   return op( byte( 0x21 ) .. uleb( index ), {val} )
end
_moduleObj.localSet = localSet

local function localTee( index, val )

   return op( byte( 0x22 ) .. uleb( index ), {val} )
end
_moduleObj.localTee = localTee

local function globalGet( index )

   return op( byte( 0x23 ) .. uleb( index ), {} )
end
_moduleObj.globalGet = globalGet

local function globalSet( index, val )

   return op( byte( 0x24 ) .. uleb( index ), {val} )
end
_moduleObj.globalSet = globalSet

local function call( funcId, args )

   return Ins._new(InsKind.FuncRef, byte( 0x10 ), args, {}, nil, nil, funcId)
end
_moduleObj.call = call

local function callRef( funcTypeIdx, args )

   return op( byte( 0x14 ) .. uleb( funcTypeIdx ), args )
end
_moduleObj.callRef = callRef

local function br( depth )

   return op( byte( 0x0c ) .. uleb( depth ), {} )
end
_moduleObj.br = br

local function brIf( depth, cond )

   return op( byte( 0x0d ) .. uleb( depth ), {cond} )
end
_moduleObj.brIf = brIf

local function ret( vals )

   return op( byte( 0x0f ), vals )
end
_moduleObj.ret = ret

local function drop( val )

   return op( byte( 0x1a ), {val} )
end
_moduleObj.drop = drop

local function unreachable(  )

   return op( byte( 0x00 ), {} )
end
_moduleObj.unreachable = unreachable

local function nop(  )

   return op( "", {} )
end
_moduleObj.nop = nop

local function select( val1, val2, cond )

   return op( byte( 0x1b ), {val1, val2, cond} )
end
_moduleObj.select = select

local function refNull( heapType )

   return op( byte( 0xd0 ) .. sleb( heapType ), {} )
end
_moduleObj.refNull = refNull

local function refIsNull( val )

   return op( byte( 0xd1 ), {val} )
end
_moduleObj.refIsNull = refIsNull

local function refEq( val1, val2 )

   return op( byte( 0xd3 ), {val1, val2} )
end
_moduleObj.refEq = refEq

local function gcOp( sub, imm, args )

   return op( byte( 0xfb ) .. uleb( sub ) .. imm, args )
end
_moduleObj.gcOp = gcOp

local function structNew( typeIdx, args )

   return gcOp( 0x00, uleb( typeIdx ), args )
end
_moduleObj.structNew = structNew

local function structNewDefault( typeIdx )

   return gcOp( 0x01, uleb( typeIdx ), {} )
end
_moduleObj.structNewDefault = structNewDefault

local function structGet( typeIdx, fieldIdx, obj )

   return gcOp( 0x02, uleb( typeIdx ) .. uleb( fieldIdx ), {obj} )
end
_moduleObj.structGet = structGet

local function structSet( typeIdx, fieldIdx, obj, val )

   return gcOp( 0x05, uleb( typeIdx ) .. uleb( fieldIdx ), {obj, val} )
end
_moduleObj.structSet = structSet

local function arrayNew( typeIdx, init, len )

   return gcOp( 0x06, uleb( typeIdx ), {init, len} )
end
_moduleObj.arrayNew = arrayNew

local function arrayNewDefault( typeIdx, len )

   return gcOp( 0x07, uleb( typeIdx ), {len} )
end
_moduleObj.arrayNewDefault = arrayNewDefault

local function arrayNewFixed( typeIdx, vals )

   return gcOp( 0x08, uleb( typeIdx ) .. uleb( #vals ), vals )
end
_moduleObj.arrayNewFixed = arrayNewFixed

local function arrayNewData( typeIdx, dataIdx, offset, len )

   return gcOp( 0x09, uleb( typeIdx ) .. uleb( dataIdx ), {offset, len} )
end
_moduleObj.arrayNewData = arrayNewData

local function arrayGet( typeIdx, arr, index )

   return gcOp( 0x0b, uleb( typeIdx ), {arr, index} )
end
_moduleObj.arrayGet = arrayGet

local function arrayGetU( typeIdx, arr, index )

   return gcOp( 0x0d, uleb( typeIdx ), {arr, index} )
end
_moduleObj.arrayGetU = arrayGetU

local function arraySet( typeIdx, arr, index, val )

   return gcOp( 0x0e, uleb( typeIdx ), {arr, index, val} )
end
_moduleObj.arraySet = arraySet

local function arrayLen( arr )

   return gcOp( 0x0f, "", {arr} )
end
_moduleObj.arrayLen = arrayLen

local function arrayCopy( dstType, srcType, dst, dstIndex, src, srcIndex, len )

   return gcOp( 0x11, uleb( dstType ) .. uleb( srcType ), {dst, dstIndex, src, srcIndex, len} )
end
_moduleObj.arrayCopy = arrayCopy

local function refTest( heapType, val )

   return gcOp( 0x14, sleb( heapType ), {val} )
end
_moduleObj.refTest = refTest

local function refCast( heapType, val )

   return gcOp( 0x17, sleb( heapType ), {val} )
end
_moduleObj.refCast = refCast

local function i32Load8U( addr )

   return op( byte( 0x2d ) .. uleb( 0 ) .. uleb( 0 ), {addr} )
end
_moduleObj.i32Load8U = i32Load8U

local function i32Store8( addr, val )

   return op( byte( 0x3a ) .. uleb( 0 ) .. uleb( 0 ), {addr, val} )
end
_moduleObj.i32Store8 = i32Store8

local memStream = {}
setmetatable( memStream, { ifList = {oStream,} } )
_moduleObj.memStream = memStream
function memStream._new(  )
   local obj = {}
   memStream._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function memStream:__init() 
   self.list = {}
   self.work = ""
end
function memStream:get_txt(  )

   self:flush(  )
   local txt = ""
   for __index, item in ipairs( self.list ) do
      txt = txt .. item
   end
   self.list = {txt}
   return txt
end
function memStream:write( val )

   self.work = self.work .. val
   if #self.work > 4000 then
      self:flush(  )
   end
   return self, nil
end
function memStream:close(  )

   self:flush(  )
end
function memStream:flush(  )

   if #self.work > 0 then
      table.insert( self.list, self.work )
      self.work = ""
   end
end
function memStream._setmeta( obj )
  setmetatable( obj, { __index = memStream  } )
end

local Func = {}
_moduleObj.Func = Func
function Func._new( typeIdx, paramNum )
   local obj = {}
   Func._setmeta( obj )
   if obj.__init then obj:__init( typeIdx, paramNum ); end
   return obj
end
function Func:__init(typeIdx, paramNum) 
   self.typeIdx = typeIdx
   self.paramNum = paramNum
   self.localTypes = {}
   self.body = {}
   self.hasBody = false
end
function Func:addLocal( valType )

   table.insert( self.localTypes, valType )
   return self.paramNum + #self.localTypes - 1
end
function Func:setBody( body )

   self.body = body
   self.hasBody = true
end
function Func:encode( funcIndexList )

   local localList = {}
   for __index, localType in ipairs( self.localTypes ) do
      table.insert( localList, uleb( 1 ) .. localType:encode(  ) )
   end
   local stream = memStream._new()
   stream:write( vec( localList ) )
   for __index, ins in ipairs( self.body ) do
      ins:encodeTo( stream, funcIndexList )
   end
   if not self.hasBody then
      stream:write( byte( 0x00 ) )
   end
   stream:write( byte( 0x0b ) )
   local txt = stream:get_txt()
   return uleb( #txt ) .. txt
end
function Func._setmeta( obj )
  setmetatable( obj, { __index = Func  } )
end
function Func:get_typeIdx()
   return self.typeIdx
end
function Func:get_localTypes()
   return self.localTypes
end
function Func:get_body()
   return self.body
end
function Func:get_paramNum()
   return self.paramNum
end
function Func:get_hasBody()
   return self.hasBody
end

local function encodeFuncType( params, results )

   local paramList = {}
   for __index, param in ipairs( params ) do
      table.insert( paramList, param:encode(  ) )
   end
   local resultList = {}
   for __index, result in ipairs( results ) do
      table.insert( resultList, result:encode(  ) )
   end
   return byte( 0x60 ) .. vec( paramList ) .. vec( resultList )
end

local Module = {}
_moduleObj.Module = Module
function Module._new(  )
   local obj = {}
   Module._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Module:__init() 
   self.typeList = {}
   self.funcTypeMap = {}
   self.importList = {}
   self.importTypeList = {}
   self.importIdList = {}
   self.importMemList = {}
   self.funcList = {}
   self.funcIdList = {}
   self.id2func = {}
   self.nextFuncId = 0
   self.globalList = {}
   self.globalInitList = {}
   self.exportNameList = {}
   self.exportIdList = {}
   self.dataList = {}
   self.refFuncSet = {}
   self.refFuncList = {}
   self.funcNameMap = {}
end
function Module:setFuncName( funcId, funcName )

   self.funcNameMap[funcId] = funcName
end
function Module:reserveType(  )

   table.insert( self.typeList, "" )
   return #self.typeList - 1
end
function Module:defineType( index, superIdx, comp )

   local superList = {}
   if superIdx ~= nil then
      table.insert( superList, uleb( superIdx ) )
   end
   self.typeList[index + 1] = byte( 0x50 ) .. vec( superList ) .. comp
end
function Module:defineStructType( index, superIdx, fields )

   local list = {}
   for __index, fieldType in ipairs( fields ) do
      table.insert( list, fieldType:encode(  ) )
   end
   self:defineType( index, superIdx, byte( 0x5f ) .. vec( list ) )
end
function Module:addStructType( superIdx, fields )

   local index = self:reserveType(  )
   self:defineStructType( index, superIdx, fields )
   return index
end
function Module:addArrayType( fieldType )

   local index = self:reserveType(  )
   self:defineType( index, nil, byte( 0x5e ) .. fieldType:encode(  ) )
   return index
end
function Module:getFuncType( params, results )

   local encoded = encodeFuncType( params, results )
   do
      local index = self.funcTypeMap[encoded]
      if index ~= nil then
         return index
      end
   end
   local index = self:reserveType(  )
   self:defineType( index, nil, encoded )
   self.funcTypeMap[encoded] = index
   return index
end
function Module:importFunc( modName, funcName, params, results )

   table.insert( self.importList, name( modName ) .. name( funcName ) )
   table.insert( self.importTypeList, encodeFuncType( params, results ) )
   local funcId = self.nextFuncId
   self.nextFuncId = self.nextFuncId + 1
   table.insert( self.importIdList, funcId )
   return funcId
end
function Module:importMemory( modName, memName )

   table.insert( self.importMemList, name( modName ) .. name( memName ) .. byte( 0x02 ) .. byte( 0x00 ) .. uleb( 0 ) )
end
function Module:addFunc( params, results )

   local typeIdx = self:getFuncType( params, results )
   return self:addFuncWithType( typeIdx, #params )
end
function Module:addFuncWithType( typeIdx, paramNum )

   local func = Func._new(typeIdx, paramNum)
   local funcId = self.nextFuncId
   self.nextFuncId = self.nextFuncId + 1
   table.insert( self.funcList, func )
   table.insert( self.funcIdList, funcId )
   self.id2func[funcId] = func
   return funcId, func
end
function Module:getFunc( funcId )

   return _lune.unwrap( self.id2func[funcId])
end
function Module:refFunc( funcId )

   if not _lune._Set_has(self.refFuncSet, funcId ) then
      self.refFuncSet[funcId]= true
      table.insert( self.refFuncList, funcId )
   end
   return Ins._new(InsKind.FuncRef, byte( 0xd2 ), {}, {}, nil, nil, funcId)
end
function Module:addGlobal( valType, init )

   table.insert( self.globalList, valType )
   table.insert( self.globalInitList, init )
   return #self.globalList - 1
end
function Module:exportFunc( exportName, funcId )

   table.insert( self.exportNameList, exportName )
   table.insert( self.exportIdList, funcId )
end
function Module:addData( data )

   table.insert( self.dataList, byte( 0x01 ) .. uleb( #data ) .. data )
   return #self.dataList - 1
end
function Module:encode(  )

   local funcIndexList = {}
   for _1 = 1, self.nextFuncId do
      table.insert( funcIndexList, 0 )
   end
   for index, funcId in ipairs( self.importIdList ) do
      funcIndexList[funcId + 1] = index - 1
   end
   for index, funcId in ipairs( self.funcIdList ) do
      funcIndexList[funcId + 1] = #self.importIdList + index - 1
   end
   local stream = memStream._new()
   stream:write( string.format( "%sasm%s", byte( 0 ), byte( 1 ) .. byte( 0 ) .. byte( 0 ) .. byte( 0 )) )
   local typeEntryList = {byte( 0x4e ) .. vec( self.typeList )}
   local importTypeMap = {}
   local importList = {}
   for index, entry in ipairs( self.importList ) do
      local funcType = self.importTypeList[index]
      local typeIdx
      
      do
         local work = importTypeMap[funcType]
         if work ~= nil then
            typeIdx = work
         else
            typeIdx = #self.typeList + #typeEntryList - 1
            importTypeMap[funcType] = typeIdx
            table.insert( typeEntryList, funcType )
         end
      end
      table.insert( importList, entry .. byte( 0x00 ) .. uleb( typeIdx ) )
   end
   stream:write( section( 1, vec( typeEntryList ) ) )
   for __index, entry in ipairs( self.importMemList ) do
      table.insert( importList, entry )
   end
   if #importList > 0 then
      stream:write( section( 2, vec( importList ) ) )
   end
   local funcTypeList = {}
   local codeList = {}
   for __index, func in ipairs( self.funcList ) do
      table.insert( funcTypeList, uleb( func:get_typeIdx() ) )
      table.insert( codeList, func:encode( funcIndexList ) )
   end
   stream:write( section( 3, vec( funcTypeList ) ) )
   if #self.globalList > 0 then
      local globalList = {}
      for index, valType in ipairs( self.globalList ) do
         local initStream = memStream._new()
         self.globalInitList[index]:encodeTo( initStream, funcIndexList )
         table.insert( globalList, valType:encode(  ) .. byte( 0x01 ) .. initStream:get_txt() .. byte( 0x0b ) )
      end
      stream:write( section( 6, vec( globalList ) ) )
   end
   if #self.exportNameList > 0 then
      local exportList = {}
      for index, exportName in ipairs( self.exportNameList ) do
         table.insert( exportList, name( exportName ) .. byte( 0x00 ) .. uleb( funcIndexList[self.exportIdList[index] + 1] ) )
      end
      stream:write( section( 7, vec( exportList ) ) )
   end
   if #self.refFuncList > 0 then
      local idxList = {}
      for __index, funcId in ipairs( self.refFuncList ) do
         table.insert( idxList, uleb( funcIndexList[funcId + 1] ) )
      end
      stream:write( section( 9, uleb( 1 ) .. byte( 0x03 ) .. byte( 0x00 ) .. vec( idxList ) ) )
   end
   if #self.dataList > 0 then
      stream:write( section( 12, uleb( #self.dataList ) ) )
   end
   stream:write( section( 10, vec( codeList ) ) )
   if #self.dataList > 0 then
      stream:write( section( 11, vec( self.dataList ) ) )
   end
   local index2name = {}
   for funcId, funcName in pairs( self.funcNameMap ) do
      index2name[funcIndexList[funcId + 1]] = funcName
   end
   local nameList = {}
   do
      local __sorted = {}
      local __map = index2name
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, funcIndex in ipairs( __sorted ) do
         local funcName = __map[ funcIndex ]
         do
            table.insert( nameList, uleb( funcIndex ) .. name( funcName ) )
         end
      end
   end
   if #nameList > 0 then
      local funcNames = vec( nameList )
      stream:write( section( 0, name( "name" ) .. byte( 1 ) .. uleb( #funcNames ) .. funcNames ) )
   end
   return stream:get_txt()
end
function Module._setmeta( obj )
  setmetatable( obj, { __index = Module  } )
end


return _moduleObj
