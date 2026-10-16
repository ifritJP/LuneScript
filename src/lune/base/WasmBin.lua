--lune/base/WasmBin.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@WasmBin'
local _lune = {}
if _lune8 then
   _lune = _lune8
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
      work = 0.0 - work
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
   while work < 1.0 do
      if exp <= -1022 then
         break
      end
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
            return byte( 0x63 ) .. sleb( self.typeIdx )
         end
         return byte( 0x64 ) .. sleb( self.typeIdx )
      end
   end
end
function ValType._setmeta( obj )
  setmetatable( obj, { __index = ValType  } )
end
function ValType._new( kind, typeIdx, nullable )
   local obj = {}
   ValType._setmeta( obj )
   if obj.__init then
      obj:__init( kind, typeIdx, nullable )
   end
   return obj
end
function ValType:__init( kind, typeIdx, nullable )

   self.kind = kind
   self.typeIdx = typeIdx
   self.nullable = nullable
end
function ValType:get_kind()
   return self.kind
end
function ValType:get_typeIdx()
   return self.typeIdx
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

local Ins = {}
_moduleObj.Ins = Ins
function Ins:encode(  )

   local txt = ""
   for __index, arg in ipairs( self.args ) do
      txt = txt .. arg:encode(  )
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
         txt = txt .. self.code
      elseif _switchExp == InsKind.Seq then
         for __index, ins in ipairs( self.body ) do
            txt = txt .. ins:encode(  )
         end
      elseif _switchExp == InsKind.Block or _switchExp == InsKind.Loop then
         if self.kind == InsKind.Block then
            txt = txt .. byte( 0x02 )
         else
          
            txt = txt .. byte( 0x03 )
         end
         txt = txt .. blockType
         for __index, ins in ipairs( self.body ) do
            txt = txt .. ins:encode(  )
         end
         txt = txt .. byte( 0x0b )
      elseif _switchExp == InsKind.If then
         txt = txt .. byte( 0x04 ) .. blockType
         for __index, ins in ipairs( self.body ) do
            txt = txt .. ins:encode(  )
         end
         do
            local elseBody = self.elseBody
            if elseBody ~= nil then
               txt = txt .. byte( 0x05 )
               for __index, ins in ipairs( elseBody ) do
                  txt = txt .. ins:encode(  )
               end
            end
         end
         txt = txt .. byte( 0x0b )
      end
   end
   return txt
end
function Ins._setmeta( obj )
  setmetatable( obj, { __index = Ins  } )
end
function Ins._new( kind, code, args, body, elseBody, result )
   local obj = {}
   Ins._setmeta( obj )
   if obj.__init then
      obj:__init( kind, code, args, body, elseBody, result )
   end
   return obj
end
function Ins:__init( kind, code, args, body, elseBody, result )

   self.kind = kind
   self.code = code
   self.args = args
   self.body = body
   self.elseBody = elseBody
   self.result = result
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

local function op( code, args )

   return Ins._new(InsKind.Plain, code, args, {}, nil, nil)
end
_moduleObj.op = op

local function seq( body )

   return Ins._new(InsKind.Seq, "", {}, body, nil, nil)
end
_moduleObj.seq = seq

local function block( body, result )

   return Ins._new(InsKind.Block, "", {}, body, nil, result)
end
_moduleObj.block = block

local function loop( body, result )

   return Ins._new(InsKind.Loop, "", {}, body, nil, result)
end
_moduleObj.loop = loop

local function ifElse( cond, thenBody, elseBody, result )

   return Ins._new(InsKind.If, "", {cond}, thenBody, elseBody, result)
end
_moduleObj.ifElse = ifElse


local function i32Const( val )

   return op( byte( 0x41 ) .. sleb( val ), {} )
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

local function globalGet( index )

   return op( byte( 0x23 ) .. uleb( index ), {} )
end
_moduleObj.globalGet = globalGet

local function globalSet( index, val )

   return op( byte( 0x24 ) .. uleb( index ), {val} )
end
_moduleObj.globalSet = globalSet

local function call( funcIndex, args )

   return op( byte( 0x10 ) .. uleb( funcIndex ), args )
end
_moduleObj.call = call

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

local function gcOp( sub, imm, args )

   return op( byte( 0xfb ) .. uleb( sub ) .. imm, args )
end
_moduleObj.gcOp = gcOp

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
end
function Func:addLocal( valType )

   table.insert( self.localTypes, valType )
   return self.paramNum + #self.localTypes - 1
end
function Func:setBody( body )

   self.body = body
end
function Func:encode(  )

   local localList = {}
   for __index, localType in ipairs( self.localTypes ) do
      table.insert( localList, uleb( 1 ) .. localType:encode(  ) )
   end
   local txt = vec( localList )
   for __index, ins in ipairs( self.body ) do
      txt = txt .. ins:encode(  )
   end
   txt = txt .. byte( 0x0b )
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
   self.importFuncNum = 0
   self.funcList = {}
   self.globalList = {}
   self.exportList = {}
   self.dataList = {}
end
function Module:addType( encoded )

   table.insert( self.typeList, encoded )
   return #self.typeList - 1
end
function Module:addByteArrayType(  )

   return self:addType( byte( 0x5e ) .. byte( 0x78 ) .. byte( 0x01 ) )
end
function Module:getFuncType( params, results )

   local paramList = {}
   for __index, param in ipairs( params ) do
      table.insert( paramList, param:encode(  ) )
   end
   local resultList = {}
   for __index, result in ipairs( results ) do
      table.insert( resultList, result:encode(  ) )
   end
   local encoded = byte( 0x60 ) .. vec( paramList ) .. vec( resultList )
   do
      local index = self.funcTypeMap[encoded]
      if index ~= nil then
         return index
      end
   end
   local index = self:addType( encoded )
   self.funcTypeMap[encoded] = index
   return index
end
function Module:importFunc( modName, funcName, params, results )

   if #self.funcList ~= 0 then
      error( "import after func" )
   end
   local typeIdx = self:getFuncType( params, results )
   table.insert( self.importList, name( modName ) .. name( funcName ) .. byte( 0x00 ) .. uleb( typeIdx ) )
   self.importFuncNum = self.importFuncNum + 1
   return self.importFuncNum - 1
end
function Module:addFunc( params, results )

   local typeIdx = self:getFuncType( params, results )
   local func = Func._new(typeIdx, #params)
   table.insert( self.funcList, func )
   return self.importFuncNum + #self.funcList - 1, func
end
function Module:addGlobal( valType, init )

   table.insert( self.globalList, valType:encode(  ) .. byte( 0x01 ) .. init:encode(  ) .. byte( 0x0b ) )
   return #self.globalList - 1
end
function Module:exportFunc( exportName, funcIndex )

   table.insert( self.exportList, name( exportName ) .. byte( 0x00 ) .. uleb( funcIndex ) )
end
function Module:addData( data )

   table.insert( self.dataList, byte( 0x01 ) .. uleb( #data ) .. data )
   return #self.dataList - 1
end
function Module:encode(  )

   local txt = string.format( "%sasm%s", byte( 0 ), byte( 1 ) .. byte( 0 ) .. byte( 0 ) .. byte( 0 ))
   txt = txt .. section( 1, vec( self.typeList ) )
   if #self.importList > 0 then
      txt = txt .. section( 2, vec( self.importList ) )
   end
   local funcTypeList = {}
   local codeList = {}
   for __index, func in ipairs( self.funcList ) do
      table.insert( funcTypeList, uleb( func:get_typeIdx() ) )
      table.insert( codeList, func:encode(  ) )
   end
   txt = txt .. section( 3, vec( funcTypeList ) )
   if #self.globalList > 0 then
      txt = txt .. section( 6, vec( self.globalList ) )
   end
   if #self.exportList > 0 then
      txt = txt .. section( 7, vec( self.exportList ) )
   end
   if #self.dataList > 0 then
      txt = txt .. section( 12, uleb( #self.dataList ) )
   end
   txt = txt .. section( 10, vec( codeList ) )
   if #self.dataList > 0 then
      txt = txt .. section( 11, vec( self.dataList ) )
   end
   return txt
end
function Module._setmeta( obj )
  setmetatable( obj, { __index = Module  } )
end
function Module:get_importFuncNum()
   return self.importFuncNum
end


return _moduleObj
