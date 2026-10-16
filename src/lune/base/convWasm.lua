--lune/base/convWasm.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@convWasm'
local _lune = {}
if _lune8 then
   _lune = _lune8
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



local Ast = _lune.loadModule( 'lune.base.Ast' )

local Nodes = _lune.loadModule( 'lune.base.Nodes' )

local Util = _lune.loadModule( 'lune.base.Util' )

local AstInfo = _lune.loadModule( 'lune.base.AstInfo' )

local Builtin = _lune.loadModule( 'lune.base.Builtin' )

local W = _lune.loadModule( 'lune.base.WasmBin' )


local Opt = {}
_moduleObj.Opt = Opt
function Opt._new( parent )
   local obj = {}
   Opt._setmeta( obj )
   if obj.__init then obj:__init( parent ); end
   return obj
end
function Opt:__init(parent) 
   self.parent = parent
end
function Opt._setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
end


local Option = {}
_moduleObj.Option = Option
function Option._new(  )
   local obj = {}
   Option._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Option:__init() 
end
function Option._setmeta( obj )
  setmetatable( obj, { __index = Option  } )
end

local opI32Eqz = 0x45

local opI32Eq = 0x46

local opI32Ne = 0x47

local opI32LtS = 0x48

local opI32GtS = 0x4a

local opI32LeS = 0x4c

local opI32GeS = 0x4e

local opI32Add = 0x6a

local opI32Sub = 0x6b

local opI32And = 0x71

local opI32Or = 0x72

local opI64Eqz = 0x50

local opI64Eq = 0x51

local opI64Ne = 0x52

local opI64LtS = 0x53

local opI64GtS = 0x55

local opI64LeS = 0x57

local opI64GeS = 0x59

local opI64Add = 0x7c

local opI64Sub = 0x7d

local opI64Mul = 0x7e

local opI64DivS = 0x7f

local opI64RemS = 0x81

local opI64And = 0x83

local opI64Or = 0x84

local opI64Xor = 0x85

local opI64Shl = 0x86

local opI64ShrU = 0x88

local opF64Eq = 0x61

local opF64Ne = 0x62

local opF64Lt = 0x63

local opF64Gt = 0x64

local opF64Le = 0x65

local opF64Ge = 0x66

local opF64Neg = 0x9a

local opF64Floor = 0x9c

local opF64Add = 0xa0

local opF64Sub = 0xa1

local opF64Mul = 0xa2

local opF64Div = 0xa3

local opI64ExtendI32U = 0xad

local opF64ConvertI64S = 0xb9

local opI64TruncF64S = 0xb0

local opRefNull = 0xd0

local strCmpOpMap = {["~="] = opI32Ne, ["<"] = opI32LtS, ["<="] = opI32LeS, [">"] = opI32GtS, [">="] = opI32GeS}

local i64CmpOpMap = {["=="] = opI64Eq, ["~="] = opI64Ne, ["<"] = opI64LtS, ["<="] = opI64LeS, [">"] = opI64GtS, [">="] = opI64GeS}

local f64CmpOpMap = {["=="] = opF64Eq, ["~="] = opF64Ne, ["<"] = opF64Lt, ["<="] = opF64Le, [">"] = opF64Gt, [">="] = opF64Ge}

local gcArrayNewData = 0x09

local gcArrayNewDefault = 0x07

local gcArrayGetU = 0x0d

local gcArrayLen = 0x0f

local gcArrayCopy = 0x11

local Kind = {}
Kind._val2NameMap = {}
function Kind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Kind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function Kind._from( val )
   if Kind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
Kind.__allList = {}
function Kind.get__allList()
   return Kind.__allList
end

Kind.Int = 0
Kind._val2NameMap[0] = 'Int'
Kind.__allList[1] = Kind.Int
Kind.Real = 1
Kind._val2NameMap[1] = 'Real'
Kind.__allList[2] = Kind.Real
Kind.Bool = 2
Kind._val2NameMap[2] = 'Bool'
Kind.__allList[3] = Kind.Bool
Kind.Str = 3
Kind._val2NameMap[3] = 'Str'
Kind.__allList[4] = Kind.Str
Kind.Nil = 4
Kind._val2NameMap[4] = 'Nil'
Kind.__allList[5] = Kind.Nil
Kind.Void = 5
Kind._val2NameMap[5] = 'Void'
Kind.__allList[6] = Kind.Void


local function plain( code, args )

   return W.op( W.byte( code ), args )
end


local convFilter = {}
setmetatable( convFilter, { __index = Nodes.Filter } )
function convFilter._new( streamName, stream, ast )
   local obj = {}
   convFilter._setmeta( obj )
   if obj.__init then obj:__init( streamName, stream, ast ); end
   return obj
end
function convFilter:__init(streamName, stream, ast) 
   Nodes.Filter.__init( self,true, ast:get_exportInfo():get_moduleTypeInfo(), ast:get_exportInfo():get_moduleTypeInfo():get_scope())
   self.streamName = streamName
   self.stream = stream
   self.processInfo = ast:get_exportInfo():get_processInfo():clone(  )
   self.builtinFuncs = ast:get_builtinFunc()
   self.moduleScope = _lune.unwrap( ast:get_exportInfo():get_moduleTypeInfo():get_scope())
   local mod = W.Module._new()
   self.strTypeIdx = mod:addByteArrayType(  )
   self.strType = W.ValType._new(W.ValKind.Ref, self.strTypeIdx, true)
   self.fnPrintI64 = mod:importFunc( "lns", "print_i64", {W.typeI64}, {} )
   self.fnPrintF64 = mod:importFunc( "lns", "print_f64", {W.typeF64}, {} )
   self.fnPrintBool = mod:importFunc( "lns", "print_bool", {W.typeI32}, {} )
   self.fnPrintStr = mod:importFunc( "lns", "print_str", {self.strType}, {} )
   self.fnPrintNil = mod:importFunc( "lns", "print_nil", {}, {} )
   self.fnPrintSep = mod:importFunc( "lns", "print_sep", {}, {} )
   self.fnPrintEnd = mod:importFunc( "lns", "print_end", {}, {} )
   self.mod = mod
   self.fnStrConcat = 0
   self.fnStrEq = 0
   self.fnStrCmp = 0
   self.fnDivI64 = 0
   self.fnModI64 = 0
   self.fnModF64 = 0
   self.fnShiftL = 0
   self.fnShiftR = 0
   self.sym2func = {}
   self.sym2global = {}
   self.sym2local = {}
   self.curFunc = nil
   self.depth = 0
   self.breakDepthList = {}
end
function convFilter._setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end

function convFilter:error( node, mess )

   Util.err( string.format( "%s:%d:%d: wasm: %s", self.streamName, node:get_pos().lineNo, node:get_pos().column, mess) )
end


function convFilter:notSupport( node )

   self:error( node, string.format( "not support -- %s", Nodes.getNodeKindName( node:get_kind() )) )
end


function convFilter:getKind( node, typeInfo )

   if typeInfo:equals( self.processInfo, Ast.builtinTypeInt ) then
      return Kind.Int
   end
   if typeInfo:equals( self.processInfo, Ast.builtinTypeReal ) then
      return Kind.Real
   end
   if typeInfo:equals( self.processInfo, Ast.builtinTypeBool ) then
      return Kind.Bool
   end
   if typeInfo:equals( self.processInfo, Ast.builtinTypeString ) then
      return Kind.Str
   end
   if typeInfo:equals( self.processInfo, Ast.builtinTypeNil ) then
      return Kind.Nil
   end
   self:error( node, string.format( "not support type -- %s", typeInfo:getTxt(  )) )
end


function convFilter:getValType( node, kind )

   do
      local _switchExp = kind
      if _switchExp == Kind.Int then
         return W.typeI64
      elseif _switchExp == Kind.Real then
         return W.typeF64
      elseif _switchExp == Kind.Bool then
         return W.typeI32
      elseif _switchExp == Kind.Str then
         return self.strType
      end
   end
   self:error( node, string.format( "not support type -- %s", Kind:_getTxt( kind)
   ) )
end

function convFilter:getDefaultVal( kind )

   do
      local _switchExp = kind
      if _switchExp == Kind.Real then
         return W.f64Const( 0.0 )
      elseif _switchExp == Kind.Bool then
         return W.i32Const( 0 )
      elseif _switchExp == Kind.Str then
         return W.op( W.byte( opRefNull ) .. W.sleb( self.strTypeIdx ), {} )
      end
   end
   return W.i64Const( 0 )
end

function convFilter:getCurFunc( node )

   local func = self.curFunc
   if  nil == func then
      local _func = func
   
      self:error( node, "no function" )
   end
   
   return func
end

function convFilter:setupRuntime(  )

   local i64 = W.typeI64
   local f64 = W.typeF64
   local i32 = W.typeI32
   local strType = self.strType
   local strIdx = W.uleb( self.strTypeIdx )
   do
      local index, func = self.mod:addFunc( {i64, i64}, {i64} )
      local q = func:addLocal( i64 )
      func:setBody( {W.localSet( q, plain( opI64DivS, {W.localGet( 0 ), W.localGet( 1 )} ) ), W.ifElse( plain( opI32And, {plain( opI64Ne, {plain( opI64RemS, {W.localGet( 0 ), W.localGet( 1 )} ), W.i64Const( 0 )} ), plain( opI32Ne, {plain( opI64LtS, {W.localGet( 0 ), W.i64Const( 0 )} ), plain( opI64LtS, {W.localGet( 1 ), W.i64Const( 0 )} )} )} ), {W.localSet( q, plain( opI64Sub, {W.localGet( q ), W.i64Const( 1 )} ) )}, nil, nil ), W.localGet( q )} )
      self.fnDivI64 = index
   end
   do
      local index, func = self.mod:addFunc( {i64, i64}, {i64} )
      local r = func:addLocal( i64 )
      func:setBody( {W.localSet( r, plain( opI64RemS, {W.localGet( 0 ), W.localGet( 1 )} ) ), W.ifElse( plain( opI32And, {plain( opI64Ne, {W.localGet( r ), W.i64Const( 0 )} ), plain( opI32Ne, {plain( opI64LtS, {W.localGet( r ), W.i64Const( 0 )} ), plain( opI64LtS, {W.localGet( 1 ), W.i64Const( 0 )} )} )} ), {W.localSet( r, plain( opI64Add, {W.localGet( r ), W.localGet( 1 )} ) )}, nil, nil ), W.localGet( r )} )
      self.fnModI64 = index
   end
   do
      local index, func = self.mod:addFunc( {f64, f64}, {f64} )
      func:setBody( {plain( opF64Sub, {W.localGet( 0 ), plain( opF64Mul, {plain( opF64Floor, {plain( opF64Div, {W.localGet( 0 ), W.localGet( 1 )} )} ), W.localGet( 1 )} )} )} )
      self.fnModF64 = index
   end
   do
      local indexL, funcL = self.mod:addFunc( {i64, i64}, {i64} )
      local indexR, funcR = self.mod:addFunc( {i64, i64}, {i64} )
      local function makeBody( code, revCode )
      
         return {W.ifElse( plain( opI32Or, {plain( opI64LeS, {W.localGet( 1 ), W.i64Const( -64 )} ), plain( opI64GeS, {W.localGet( 1 ), W.i64Const( 64 )} )} ), {W.ret( {W.i64Const( 0 )} )}, nil, nil ), W.ifElse( plain( opI64LtS, {W.localGet( 1 ), W.i64Const( 0 )} ), {W.ret( {plain( revCode, {W.localGet( 0 ), plain( opI64Sub, {W.i64Const( 0 ), W.localGet( 1 )} )} )} )}, nil, nil ), plain( code, {W.localGet( 0 ), W.localGet( 1 )} )}
      end
      funcL:setBody( makeBody( opI64Shl, opI64ShrU ) )
      funcR:setBody( makeBody( opI64ShrU, opI64Shl ) )
      self.fnShiftL = indexL
      self.fnShiftR = indexR
   end
   do
      local index, func = self.mod:addFunc( {strType, strType}, {strType} )
      local lenA = func:addLocal( i32 )
      local lenB = func:addLocal( i32 )
      local work = func:addLocal( strType )
      func:setBody( {W.localSet( lenA, W.gcOp( gcArrayLen, "", {W.localGet( 0 )} ) ), W.localSet( lenB, W.gcOp( gcArrayLen, "", {W.localGet( 1 )} ) ), W.localSet( work, W.gcOp( gcArrayNewDefault, strIdx, {plain( opI32Add, {W.localGet( lenA ), W.localGet( lenB )} )} ) ), W.gcOp( gcArrayCopy, strIdx .. strIdx, {W.localGet( work ), W.i32Const( 0 ), W.localGet( 0 ), W.i32Const( 0 ), W.localGet( lenA )} ), W.gcOp( gcArrayCopy, strIdx .. strIdx, {W.localGet( work ), W.localGet( lenA ), W.localGet( 1 ), W.i32Const( 0 ), W.localGet( lenB )} ), W.localGet( work )} )
      self.fnStrConcat = index
   end
   do
      local index, func = self.mod:addFunc( {strType, strType}, {i32} )
      local lenA = func:addLocal( i32 )
      local lenB = func:addLocal( i32 )
      local pos = func:addLocal( i32 )
      local diff = func:addLocal( i32 )
      func:setBody( {W.localSet( lenA, W.gcOp( gcArrayLen, "", {W.localGet( 0 )} ) ), W.localSet( lenB, W.gcOp( gcArrayLen, "", {W.localGet( 1 )} ) ), W.block( {W.loop( {W.brIf( 1, plain( opI32GeS, {W.localGet( pos ), W.localGet( lenA )} ) ), W.brIf( 1, plain( opI32GeS, {W.localGet( pos ), W.localGet( lenB )} ) ), W.localSet( diff, plain( opI32Sub, {W.gcOp( gcArrayGetU, strIdx, {W.localGet( 0 ), W.localGet( pos )} ), W.gcOp( gcArrayGetU, strIdx, {W.localGet( 1 ), W.localGet( pos )} )} ) ), W.ifElse( W.localGet( diff ), {W.ret( {W.localGet( diff )} )}, nil, nil ), W.localSet( pos, plain( opI32Add, {W.localGet( pos ), W.i32Const( 1 )} ) ), W.br( 0 )}, nil )}, nil ), plain( opI32Sub, {W.localGet( lenA ), W.localGet( lenB )} )} )
      self.fnStrCmp = index
   end
   do
      local index, func = self.mod:addFunc( {strType, strType}, {i32} )
      func:setBody( {plain( opI32Eqz, {W.call( self.fnStrCmp, {W.localGet( 0 ), W.localGet( 1 )} )} )} )
      self.fnStrEq = index
   end
   do
      local index, func = self.mod:addFunc( {strType}, {i32} )
      func:setBody( {W.gcOp( gcArrayLen, "", {W.localGet( 0 )} )} )
      self.mod:exportFunc( "lns_str_len", index )
   end
   do
      local index, func = self.mod:addFunc( {strType, i32}, {i32} )
      func:setBody( {W.gcOp( gcArrayGetU, strIdx, {W.localGet( 0 ), W.localGet( 1 )} )} )
      self.mod:exportFunc( "lns_str_get", index )
   end
end

function convFilter:decodeString( node, txt )

   local quote = txt:sub( 1, 1 )
   if quote == "`" then
      return txt:sub( 4, #txt - 3 )
   end
   local body = txt:sub( 2, #txt - 1 )
   local result = ""
   local pos = 1
   while pos <= #body do
      local ch = body:sub( pos, pos )
      if ch ~= "\\" then
         result = result .. ch
         pos = pos + 1
      else
       
         local next = body:sub( pos + 1, pos + 1 )
         pos = pos + 2
         do
            local _switchExp = next
            if _switchExp == "n" then
               result = result .. "\n"
            elseif _switchExp == "t" then
               result = result .. "\t"
            elseif _switchExp == "r" then
               result = result .. "\r"
            elseif _switchExp == "a" then
               result = result .. W.byte( 7 )
            elseif _switchExp == "b" then
               result = result .. W.byte( 8 )
            elseif _switchExp == "f" then
               result = result .. W.byte( 12 )
            elseif _switchExp == "v" then
               result = result .. W.byte( 11 )
            elseif _switchExp == "\\" or _switchExp == "\"" or _switchExp == "'" then
               result = result .. next
            elseif _switchExp == "x" then
               local hex = body:sub( pos, pos + 1 )
               local val = tonumber( hex, 16 )
               if  nil == val then
                  local _val = val
               
                  self:error( node, string.format( "illegal escape -- %s", txt) )
               end
               
               result = result .. W.byte( math.floor(val) )
               pos = pos + 2
            else 
               
                  local val = 0
                  local count = 0
                  pos = pos - 1
                  while count < 3 do
                     local code = body:byte( pos )
                     if  nil == code then
                        local _code = code
                     
                        break
                     end
                     
                     if code < 0x30 or code > 0x39 then
                        break
                     end
                     val = val * 10 + code - 0x30
                     count = count + 1
                     pos = pos + 1
                  end
                  if count == 0 then
                     self:error( node, string.format( "illegal escape -- %s", txt) )
                  end
                  result = result .. W.byte( val )
            end
         end
      end
   end
   return result
end

function convFilter:getExpKind( node )

   do
      local castNode = _lune.__Cast( node, 3, Nodes.ExpCastNode )
      if castNode ~= nil then
         return self:getKind( node, castNode:get_castType() )
      end
   end
   if node:get_kind() == Nodes.NodeKind.get_ExpCall() then
      local callNode = _lune.__Cast( node, 3, Nodes.ExpCallNode )
      if  nil == callNode then
         local _callNode = callNode
      
         self:error( node, "cast" )
      end
      
      local funcType = callNode:get_func():get_expType()
      local retList = funcType:get_retTypeInfoList()
      if #retList == 0 then
         return Kind.Void
      end
      if #retList > 1 then
         self:error( node, "not support multiple return values" )
      end
      return self:getKind( node, retList[1] )
   end
   return self:getKind( node, node:get_expType() )
end

function convFilter:toReal( node )

   local ins = self:lowerExp( node )
   if self:getExpKind( node ) == Kind.Int then
      return plain( opF64ConvertI64S, {ins} )
   end
   return ins
end

function convFilter:lowerOp2( node )

   local opTxt = node:get_op().txt
   local kind1 = self:getExpKind( node:get_exp1() )
   local kind2 = self:getExpKind( node:get_exp2() )
   do
      local _switchExp = opTxt
      if _switchExp == "and" or _switchExp == "or" then
         if kind1 ~= Kind.Bool or kind2 ~= Kind.Bool then
            self:error( node, string.format( "not support non bool %s", opTxt) )
         end
         local exp1 = self:lowerExp( node:get_exp1() )
         self.depth = self.depth + 1
         local exp2 = self:lowerExp( node:get_exp2() )
         self.depth = self.depth - 1
         if opTxt == "and" then
            return W.ifElse( exp1, {exp2}, {W.i32Const( 0 )}, W.typeI32 )
         end
         return W.ifElse( exp1, {W.i32Const( 1 )}, {exp2}, W.typeI32 )
      elseif _switchExp == ".." then
         return W.call( self.fnStrConcat, {self:lowerExp( node:get_exp1() ), self:lowerExp( node:get_exp2() )} )
      end
   end
   do
      local _switchExp = opTxt
      if _switchExp == "==" or _switchExp == "~=" or _switchExp == "<" or _switchExp == "<=" or _switchExp == ">" or _switchExp == ">=" then
         if kind1 == Kind.Str and kind2 == Kind.Str then
            local cmp = W.call( self.fnStrCmp, {self:lowerExp( node:get_exp1() ), self:lowerExp( node:get_exp2() )} )
            if opTxt == "==" then
               return plain( opI32Eqz, {cmp} )
            end
            local code = strCmpOpMap[opTxt]
            if  nil == code then
               local _code = code
            
               self:error( node, string.format( "illegal op -- %s", opTxt) )
            end
            
            return plain( code, {cmp, W.i32Const( 0 )} )
         end
         if kind1 == Kind.Bool and kind2 == Kind.Bool then
            local code
            
            do
               local _switchExp = opTxt
               if _switchExp == "==" then
                  code = opI32Eq
               elseif _switchExp == "~=" then
                  code = opI32Ne
               else 
                  
                     self:error( node, string.format( "not support bool %s", opTxt) )
               end
            end
            return plain( code, {self:lowerExp( node:get_exp1() ), self:lowerExp( node:get_exp2() )} )
         end
         if kind1 == Kind.Int and kind2 == Kind.Int then
            local code = i64CmpOpMap[opTxt]
            if  nil == code then
               local _code = code
            
               self:error( node, string.format( "illegal op -- %s", opTxt) )
            end
            
            return plain( code, {self:lowerExp( node:get_exp1() ), self:lowerExp( node:get_exp2() )} )
         end
         if (kind1 == Kind.Int or kind1 == Kind.Real ) and (kind2 == Kind.Int or kind2 == Kind.Real ) then
            local code = f64CmpOpMap[opTxt]
            if  nil == code then
               local _code = code
            
               self:error( node, string.format( "illegal op -- %s", opTxt) )
            end
            
            return plain( code, {self:toReal( node:get_exp1() ), self:toReal( node:get_exp2() )} )
         end
         self:error( node, string.format( "not support compare -- %s %s %s", Kind:_getTxt( kind1)
         , opTxt, Kind:_getTxt( kind2)
         ) )
      end
   end
   local resultKind = self:getKind( node, node:get_expType() )
   if resultKind == Kind.Int then
      local exp1 = self:lowerExp( node:get_exp1() )
      local exp2 = self:lowerExp( node:get_exp2() )
      do
         local _switchExp = opTxt
         if _switchExp == "+" then
            return plain( opI64Add, {exp1, exp2} )
         elseif _switchExp == "-" then
            return plain( opI64Sub, {exp1, exp2} )
         elseif _switchExp == "*" then
            return plain( opI64Mul, {exp1, exp2} )
         elseif _switchExp == "/" then
            return W.call( self.fnDivI64, {exp1, exp2} )
         elseif _switchExp == "%" then
            return W.call( self.fnModI64, {exp1, exp2} )
         elseif _switchExp == "&" then
            return plain( opI64And, {exp1, exp2} )
         elseif _switchExp == "|" then
            return plain( opI64Or, {exp1, exp2} )
         elseif _switchExp == "~" then
            return plain( opI64Xor, {exp1, exp2} )
         elseif _switchExp == "|<<" then
            return W.call( self.fnShiftL, {exp1, exp2} )
         elseif _switchExp == "|>>" then
            return W.call( self.fnShiftR, {exp1, exp2} )
         end
      end
   end
   if resultKind == Kind.Real then
      local exp1 = self:toReal( node:get_exp1() )
      local exp2 = self:toReal( node:get_exp2() )
      do
         local _switchExp = opTxt
         if _switchExp == "+" then
            return plain( opF64Add, {exp1, exp2} )
         elseif _switchExp == "-" then
            return plain( opF64Sub, {exp1, exp2} )
         elseif _switchExp == "*" then
            return plain( opF64Mul, {exp1, exp2} )
         elseif _switchExp == "/" then
            return plain( opF64Div, {exp1, exp2} )
         elseif _switchExp == "%" then
            return W.call( self.fnModF64, {exp1, exp2} )
         end
      end
   end
   self:error( node, string.format( "not support op -- %s (%s)", opTxt, Kind:_getTxt( resultKind)
   ) )
end

function convFilter:lowerOp1( node )

   local opTxt = node:get_op().txt
   local kind = self:getExpKind( node:get_exp() )
   local exp = self:lowerExp( node:get_exp() )
   do
      local _switchExp = opTxt
      if _switchExp == "-" then
         if kind == Kind.Int then
            return plain( opI64Sub, {W.i64Const( 0 ), exp} )
         end
         if kind == Kind.Real then
            return plain( opF64Neg, {exp} )
         end
      elseif _switchExp == "not" then
         if kind == Kind.Bool then
            return plain( opI32Eqz, {exp} )
         end
      elseif _switchExp == "#" then
         if kind == Kind.Str then
            return plain( opI64ExtendI32U, {W.gcOp( gcArrayLen, "", {exp} )} )
         end
      elseif _switchExp == "~" then
         if kind == Kind.Int then
            return plain( opI64Xor, {exp, W.i64Const( -1 )} )
         end
      end
   end
   self:error( node, string.format( "not support op -- %s (%s)", opTxt, Kind:_getTxt( kind)
   ) )
end

function convFilter:lowerArgs( node )

   local list = {}
   do
      local argList = node:get_argList()
      if argList ~= nil then
         if argList:get_mRetExp() then
            self:error( node, "not support multiple return values in args" )
         end
         for __index, arg in ipairs( argList:get_expList() ) do
            table.insert( list, self:lowerExp( arg ) )
         end
      end
   end
   return list
end

function convFilter:getDDDArgs( node )

   local list = {}
   do
      local argList = node:get_argList()
      if argList ~= nil then
         if argList:get_mRetExp() then
            self:error( node, "not support multiple return values in args" )
         end
         for __index, arg in ipairs( argList:get_expList() ) do
            do
               local toDDD = _lune.__Cast( arg, 3, Nodes.ExpToDDDNode )
               if toDDD ~= nil then
                  for __index, exp in ipairs( toDDD:get_expList():get_expList() ) do
                     table.insert( list, exp )
                  end
               else
                  table.insert( list, arg )
               end
            end
         end
      end
   end
   return list
end

function convFilter:lowerPrint( node )

   local list = {}
   do
      for index, arg in ipairs( self:getDDDArgs( node ) ) do
         if index > 1 then
            table.insert( list, W.call( self.fnPrintSep, {} ) )
         end
         local kind = self:getExpKind( arg )
         do
            local _switchExp = kind
            if _switchExp == Kind.Int then
               table.insert( list, W.call( self.fnPrintI64, {self:lowerExp( arg )} ) )
            elseif _switchExp == Kind.Real then
               table.insert( list, W.call( self.fnPrintF64, {self:lowerExp( arg )} ) )
            elseif _switchExp == Kind.Bool then
               table.insert( list, W.call( self.fnPrintBool, {self:lowerExp( arg )} ) )
            elseif _switchExp == Kind.Str then
               table.insert( list, W.call( self.fnPrintStr, {self:lowerExp( arg )} ) )
            elseif _switchExp == Kind.Nil then
               table.insert( list, W.call( self.fnPrintNil, {} ) )
            elseif _switchExp == Kind.Void then
               self:error( arg, "void value" )
            end
         end
      end
   end
   table.insert( list, W.call( self.fnPrintEnd, {} ) )
   return W.seq( list )
end


function convFilter:lowerCall( node )

   if node:get_func():get_expType():equals( self.processInfo, self.builtinFuncs.lns_print ) then
      return self:lowerPrint( node )
   end
   local refNode = _lune.__Cast( node:get_func(), 3, Nodes.ExpRefNode )
   if  nil == refNode then
      local _refNode = refNode
   
      self:error( node, string.format( "not support func -- %s", Nodes.getNodeKindName( node:get_func():get_kind() )) )
   end
   
   local funcIndex = self.sym2func[refNode:get_symbolInfo():getOrg(  )]
   if  nil == funcIndex then
      local _funcIndex = funcIndex
   
      self:error( node, string.format( "not support func -- %s", refNode:get_symbolInfo():get_name()) )
   end
   
   return W.call( funcIndex, self:lowerArgs( node ) )
end

function convFilter:lowerRef( node )

   local symbol = node:get_symbolInfo():getOrg(  )
   do
      local index = self.sym2local[symbol]
      if index ~= nil then
         return W.localGet( index )
      end
   end
   do
      local index = self.sym2global[symbol]
      if index ~= nil then
         return W.globalGet( index )
      end
   end
   self:error( node, string.format( "not support symbol -- %s", symbol:get_name()) )
end

function convFilter:lowerSetSym( node, symbol, val )

   local orgSym = symbol:getOrg(  )
   do
      local index = self.sym2local[orgSym]
      if index ~= nil then
         return W.localSet( index, val )
      end
   end
   do
      local index = self.sym2global[orgSym]
      if index ~= nil then
         return W.globalSet( index, val )
      end
   end
   self:error( node, string.format( "not support symbol -- %s", symbol:get_name()) )
end

function convFilter:lowerExp( node )

   do
      local _switchExp = node:get_kind()
      if _switchExp == Nodes.NodeKind.get_LiteralInt() then
         local literal = _lune.__Cast( node, 3, Nodes.LiteralIntNode )
         if  nil == literal then
            local _literal = literal
         
            self:error( node, "cast" )
         end
         
         return W.i64Const( literal:get_num() )
      elseif _switchExp == Nodes.NodeKind.get_LiteralReal() then
         local literal = _lune.__Cast( node, 3, Nodes.LiteralRealNode )
         if  nil == literal then
            local _literal = literal
         
            self:error( node, "cast" )
         end
         
         return W.f64Const( literal:get_num() )
      elseif _switchExp == Nodes.NodeKind.get_LiteralBool() then
         local literal = _lune.__Cast( node, 3, Nodes.LiteralBoolNode )
         if  nil == literal then
            local _literal = literal
         
            self:error( node, "cast" )
         end
         
         if literal:get_token().txt == "true" then
            return W.i32Const( 1 )
         end
         return W.i32Const( 0 )
      elseif _switchExp == Nodes.NodeKind.get_LiteralString() then
         local literal = _lune.__Cast( node, 3, Nodes.LiteralStringNode )
         if  nil == literal then
            local _literal = literal
         
            self:error( node, "cast" )
         end
         
         if literal:get_orgParam() then
            self:error( node, "not support format string" )
         end
         local data = self:decodeString( node, literal:get_token().txt )
         local dataIdx = self.mod:addData( data )
         return W.gcOp( gcArrayNewData, W.uleb( self.strTypeIdx ) .. W.uleb( dataIdx ), {W.i32Const( 0 ), W.i32Const( #data )} )
      elseif _switchExp == Nodes.NodeKind.get_ExpRef() then
         local refNode = _lune.__Cast( node, 3, Nodes.ExpRefNode )
         if  nil == refNode then
            local _refNode = refNode
         
            self:error( node, "cast" )
         end
         
         return self:lowerRef( refNode )
      elseif _switchExp == Nodes.NodeKind.get_ExpParen() then
         local parenNode = _lune.__Cast( node, 3, Nodes.ExpParenNode )
         if  nil == parenNode then
            local _parenNode = parenNode
         
            self:error( node, "cast" )
         end
         
         return self:lowerExp( parenNode:get_exp() )
      elseif _switchExp == Nodes.NodeKind.get_ExpOp2() then
         local op2Node = _lune.__Cast( node, 3, Nodes.ExpOp2Node )
         if  nil == op2Node then
            local _op2Node = op2Node
         
            self:error( node, "cast" )
         end
         
         return self:lowerOp2( op2Node )
      elseif _switchExp == Nodes.NodeKind.get_ExpOp1() then
         local op1Node = _lune.__Cast( node, 3, Nodes.ExpOp1Node )
         if  nil == op1Node then
            local _op1Node = op1Node
         
            self:error( node, "cast" )
         end
         
         return self:lowerOp1( op1Node )
      elseif _switchExp == Nodes.NodeKind.get_ExpCall() then
         local callNode = _lune.__Cast( node, 3, Nodes.ExpCallNode )
         if  nil == callNode then
            local _callNode = callNode
         
            self:error( node, "cast" )
         end
         
         return self:lowerCall( callNode )
      elseif _switchExp == Nodes.NodeKind.get_ExpCast() then
         local castNode = _lune.__Cast( node, 3, Nodes.ExpCastNode )
         if  nil == castNode then
            local _castNode = castNode
         
            self:error( node, "cast" )
         end
         
         local srcKind = self:getExpKind( castNode:get_exp() )
         local dstKind = self:getKind( node, castNode:get_castType() )
         local exp = self:lowerExp( castNode:get_exp() )
         if srcKind == dstKind then
            return exp
         end
         if srcKind == Kind.Int and dstKind == Kind.Real then
            return plain( opF64ConvertI64S, {exp} )
         end
         if srcKind == Kind.Real and dstKind == Kind.Int then
            return plain( opI64TruncF64S, {exp} )
         end
         self:error( node, string.format( "not support cast -- %s -> %s", Kind:_getTxt( srcKind)
         , Kind:_getTxt( dstKind)
         ) )
      end
   end
   self:notSupport( node )
end

function convFilter:lowerBlock( node )

   local list = {}
   for __index, stmt in ipairs( node:get_stmtList() ) do
      self:lowerStmt( stmt, list )
   end
   return list
end

function convFilter:enterLoop(  )

   self.depth = self.depth + 1
   table.insert( self.breakDepthList, self.depth )
   self.depth = self.depth + 1
end

function convFilter:exitLoop( body )

   self.depth = self.depth - 2
   table.remove( self.breakDepthList )
   return W.block( {W.loop( body, nil )}, nil )
end


function convFilter:lowerDeclVar( node, out )

   if node:get_unwrapFlag() or node:get_condRetInfo() then
      self:error( node, "not support let!" )
   end
   local expList = {}
   do
      local expListNode = node:get_expList()
      if expListNode ~= nil then
         if expListNode:get_mRetExp() then
            self:error( node, "not support multiple return values" )
         end
         for __index, exp in ipairs( expListNode:get_expList() ) do
            table.insert( expList, exp )
         end
      end
   end
   for index, symbol in ipairs( node:get_symbolInfoList() ) do
      local orgSym = symbol:getOrg(  )
      local kind = self:getKind( node, symbol:get_typeInfo() )
      local valType = self:getValType( node, kind )
      if self.curFunc == nil then
         self:error( node, "no function" )
      end
      if symbol:get_scope() == self.moduleScope then
         self.sym2global[orgSym] = self.mod:addGlobal( valType, self:getDefaultVal( kind ) )
      else
       
         local func = self:getCurFunc( node )
         self.sym2local[orgSym] = func:addLocal( valType )
      end
      if index <= #expList then
         table.insert( out, self:lowerSetSym( node, orgSym, self:lowerExp( expList[index] ) ) )
      end
   end
end


function convFilter:lowerIf( node )

   local function process( index )
   
      local stmt = node:get_stmtList()[index]
      if stmt:get_kind() == Nodes.IfKind.Else then
         return self:lowerBlock( stmt:get_block() )
      end
      local cond = self:lowerExp( stmt:get_exp() )
      self.depth = self.depth + 1
      local thenBody = self:lowerBlock( stmt:get_block() )
      local elseBody = nil
      if index < #node:get_stmtList() then
         elseBody = process( index + 1 )
      end
      self.depth = self.depth - 1
      return {W.ifElse( cond, thenBody, elseBody, nil )}
   end
   return W.seq( process( 1 ) )
end


function convFilter:lowerFor( node )

   if self:getKind( node, node:get_val():get_typeInfo() ) ~= Kind.Int then
      self:error( node, "not support non int for" )
   end
   local func = self:getCurFunc( node )
   local valIdx = func:addLocal( W.typeI64 )
   self.sym2local[node:get_val():getOrg(  )] = valIdx
   local toIdx = func:addLocal( W.typeI64 )
   local stepIdx = func:addLocal( W.typeI64 )
   local list = {}
   table.insert( list, W.localSet( valIdx, self:lowerExp( node:get_init() ) ) )
   table.insert( list, W.localSet( toIdx, self:lowerExp( node:get_to() ) ) )
   do
      local delta = node:get_delta()
      if delta ~= nil then
         table.insert( list, W.localSet( stepIdx, self:lowerExp( delta ) ) )
      else
         table.insert( list, W.localSet( stepIdx, W.i64Const( 1 ) ) )
      end
   end
   self:enterLoop(  )
   do
      local body = {}
      self.depth = self.depth + 1
      local exitCond = W.ifElse( plain( opI64GeS, {W.localGet( stepIdx ), W.i64Const( 0 )} ), {plain( opI64GtS, {W.localGet( valIdx ), W.localGet( toIdx )} )}, {plain( opI64LtS, {W.localGet( valIdx ), W.localGet( toIdx )} )}, W.typeI32 )
      self.depth = self.depth - 1
      table.insert( body, W.brIf( 1, exitCond ) )
      for __index, ins in ipairs( self:lowerBlock( node:get_block() ) ) do
         table.insert( body, ins )
      end
      table.insert( body, W.localSet( valIdx, plain( opI64Add, {W.localGet( valIdx ), W.localGet( stepIdx )} ) ) )
      table.insert( body, W.br( 0 ) )
      table.insert( list, self:exitLoop( body ) )
   end
   return W.seq( list )
end

function convFilter:lowerStmt( node, out )

   do
      local _switchExp = node:get_kind()
      if _switchExp == Nodes.NodeKind.get_None() or _switchExp == Nodes.NodeKind.get_BlankLine() or _switchExp == Nodes.NodeKind.get_Shebang() or _switchExp == Nodes.NodeKind.get_LuneControl() then
         return 
      elseif _switchExp == Nodes.NodeKind.get_DeclFunc() then
         if self.depth == 0 then
            return 
         end
      elseif _switchExp == Nodes.NodeKind.get_StmtExp() then
         local stmtNode = _lune.__Cast( node, 3, Nodes.StmtExpNode )
         if  nil == stmtNode then
            local _stmtNode = stmtNode
         
            self:error( node, "cast" )
         end
         
         if stmtNode:get_exp():get_kind() == Nodes.NodeKind.get_ExpSetVal() then
            self:lowerStmt( stmtNode:get_exp(), out )
            return 
         end
         local ins = self:lowerExp( stmtNode:get_exp() )
         if self:getExpKind( stmtNode:get_exp() ) == Kind.Void then
            table.insert( out, ins )
         else
          
            table.insert( out, W.drop( ins ) )
         end
         return 
      elseif _switchExp == Nodes.NodeKind.get_DeclVar() then
         local declNode = _lune.__Cast( node, 3, Nodes.DeclVarNode )
         if  nil == declNode then
            local _declNode = declNode
         
            self:error( node, "cast" )
         end
         
         self:lowerDeclVar( declNode, out )
         return 
      elseif _switchExp == Nodes.NodeKind.get_ExpSetVal() then
         local setNode = _lune.__Cast( node, 3, Nodes.ExpSetValNode )
         if  nil == setNode then
            local _setNode = setNode
         
            self:error( node, "cast" )
         end
         
         if #setNode:get_LeftSymList() ~= 1 or #setNode:get_exp2():get_expList() ~= 1 then
            self:error( node, "not support multiple assign" )
         end
         table.insert( out, self:lowerSetSym( node, setNode:get_LeftSymList()[1], self:lowerExp( setNode:get_exp2():get_expList()[1] ) ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_If() then
         local ifNode = _lune.__Cast( node, 3, Nodes.IfNode )
         if  nil == ifNode then
            local _ifNode = ifNode
         
            self:error( node, "cast" )
         end
         
         table.insert( out, self:lowerIf( ifNode ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_While() then
         local whileNode = _lune.__Cast( node, 3, Nodes.WhileNode )
         if  nil == whileNode then
            local _whileNode = whileNode
         
            self:error( node, "cast" )
         end
         
         self:enterLoop(  )
         do
            local body = {}
            if not whileNode:get_infinit() then
               table.insert( body, W.brIf( 1, plain( opI32Eqz, {self:lowerExp( whileNode:get_exp() )} ) ) )
            end
            for __index, ins in ipairs( self:lowerBlock( whileNode:get_block() ) ) do
               table.insert( body, ins )
            end
            table.insert( body, W.br( 0 ) )
            table.insert( out, self:exitLoop( body ) )
         end
         return 
      elseif _switchExp == Nodes.NodeKind.get_Repeat() then
         local repeatNode = _lune.__Cast( node, 3, Nodes.RepeatNode )
         if  nil == repeatNode then
            local _repeatNode = repeatNode
         
            self:error( node, "cast" )
         end
         
         self:enterLoop(  )
         do
            local body = {}
            for __index, ins in ipairs( self:lowerBlock( repeatNode:get_block() ) ) do
               table.insert( body, ins )
            end
            table.insert( body, W.brIf( 0, plain( opI32Eqz, {self:lowerExp( repeatNode:get_exp() )} ) ) )
            table.insert( out, self:exitLoop( body ) )
         end
         return 
      elseif _switchExp == Nodes.NodeKind.get_For() then
         local forNode = _lune.__Cast( node, 3, Nodes.ForNode )
         if  nil == forNode then
            local _forNode = forNode
         
            self:error( node, "cast" )
         end
         
         table.insert( out, self:lowerFor( forNode ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Break() then
         if #self.breakDepthList == 0 then
            self:error( node, "break outside loop" )
         end
         table.insert( out, W.br( self.depth - self.breakDepthList[#self.breakDepthList] ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Return() then
         local retNode = _lune.__Cast( node, 3, Nodes.ReturnNode )
         if  nil == retNode then
            local _retNode = retNode
         
            self:error( node, "cast" )
         end
         
         local vals = {}
         do
            local expList = retNode:get_expList()
            if expList ~= nil then
               for __index, exp in ipairs( expList:get_expList() ) do
                  table.insert( vals, self:lowerExp( exp ) )
               end
            end
         end
         table.insert( out, W.ret( vals ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Scope() then
         local scopeNode = _lune.__Cast( node, 3, Nodes.ScopeNode )
         if  nil == scopeNode then
            local _scopeNode = scopeNode
         
            self:error( node, "cast" )
         end
         
         for __index, ins in ipairs( self:lowerBlock( scopeNode:get_block() ) ) do
            table.insert( out, ins )
         end
         return 
      elseif _switchExp == Nodes.NodeKind.get_Block() then
         local blockNode = _lune.__Cast( node, 3, Nodes.BlockNode )
         if  nil == blockNode then
            local _blockNode = blockNode
         
            self:error( node, "cast" )
         end
         
         for __index, ins in ipairs( self:lowerBlock( blockNode ) ) do
            table.insert( out, ins )
         end
         return 
      end
   end
   self:notSupport( node )
end

function convFilter:registerFunc( node )

   local declInfo = node:get_declInfo()
   if declInfo:get_classTypeInfo() then
      self:error( node, "not support method" )
   end
   local symbol = declInfo:get_symbol()
   if  nil == symbol then
      local _symbol = symbol
   
      self:error( node, "not support anonymous function" )
   end
   
   local params = {}
   for __index, arg in ipairs( declInfo:get_argList() ) do
      local argNode = _lune.__Cast( arg, 3, Nodes.DeclArgNode )
      if  nil == argNode then
         local _argNode = argNode
      
         self:error( arg, string.format( "not support arg -- %s", Nodes.getNodeKindName( arg:get_kind() )) )
      end
      
      table.insert( params, self:getValType( arg, self:getKind( arg, argNode:get_symbolInfo():get_typeInfo() ) ) )
   end
   local results = {}
   for __index, retType in ipairs( declInfo:get_retTypeInfoList() ) do
      table.insert( results, self:getValType( node, self:getKind( node, retType ) ) )
   end
   local index, func = self.mod:addFunc( params, results )
   self.sym2func[symbol:getOrg(  )] = index
   return func
end

function convFilter:lowerFunc( node, func )

   local declInfo = node:get_declInfo()
   for index, arg in ipairs( declInfo:get_argList() ) do
      local argNode = _lune.__Cast( arg, 3, Nodes.DeclArgNode )
      if  nil == argNode then
         local _argNode = argNode
      
         self:error( arg, "cast" )
      end
      
      self.sym2local[argNode:get_symbolInfo():getOrg(  )] = index - 1
   end
   self.curFunc = func
   self.depth = 1
   local body = {}
   do
      local block = declInfo:get_body()
      if block ~= nil then
         for __index, ins in ipairs( self:lowerBlock( block ) ) do
            table.insert( body, ins )
         end
      end
   end
   if #declInfo:get_retTypeInfoList() > 0 then
      table.insert( body, W.unreachable(  ) )
   end
   func:setBody( body )
   self.curFunc = nil
   self.depth = 0
end


function convFilter:processRoot( node, opt )

   self:setupRuntime(  )
   local funcList = {}
   local wasmFuncList = {}
   for __index, child in ipairs( node:get_children() ) do
      if child:get_kind() == Nodes.NodeKind.get_DeclFunc() then
         local declNode = _lune.__Cast( child, 3, Nodes.DeclFuncNode )
         if  nil == declNode then
            local _declNode = declNode
         
            self:error( child, "cast" )
         end
         
         table.insert( funcList, declNode )
         table.insert( wasmFuncList, self:registerFunc( declNode ) )
      end
   end
   local mainIndex, mainFunc = self.mod:addFunc( {}, {} )
   self.curFunc = mainFunc
   self.depth = 0
   local body = {}
   for __index, child in ipairs( node:get_children() ) do
      self:lowerStmt( child, body )
   end
   mainFunc:setBody( body )
   self.curFunc = nil
   self.mod:exportFunc( "main", mainIndex )
   for index, declNode in ipairs( funcList ) do
      self:lowerFunc( declNode, wasmFuncList[index] )
   end
   self.stream:write( self.mod:encode(  ) )
end


local function createFilter( streamName, stream, ast )

   return convFilter._new(streamName, stream, ast)
end
_moduleObj.createFilter = createFilter


return _moduleObj
