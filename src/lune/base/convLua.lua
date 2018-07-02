--lune/base/convLua.lns
local moduleObj = {}
local TransUnit = require( 'lune.base.TransUnit' )

local filterObj = {}
moduleObj.filterObj = filterObj
function filterObj.new( streamName, stream, exeFlag )
  local obj = {}
  setmetatable( obj, { __index = filterObj } )
  return obj.__init and obj:__init( streamName, stream, exeFlag ) or nil;
end
function filterObj:__init(streamName, stream, exeFlag) 
  self.streamName = streamName
  self.stream = stream
  self.moduleName2Info = {}
  self.exeFlag = exeFlag
  self.indent = 0
  self.curLineNo = 1
  self.className2InfoMap = {}
  return self
end

local stepIndent = 2
local builtInModuleSet = {}
builtInModuleSet["io"] = true
builtInModuleSet["string"] = true
builtInModuleSet["table"] = true
builtInModuleSet["math"] = true
builtInModuleSet["debug"] = true
builtInModuleSet["_luneScript"] = true
function filterObj:write( txt )
  if self.needIndent then
    self.stream:write( string.rep( " ", self.indent ) )
    self.needIndent = false
  end
  for cr in string.gmatch( txt, "\n" ) do
    self.curLineNo = self.curLineNo + 1
  end
  self.stream:write( txt )
end

function filterObj:setIndent( indent )
  self.indent = indent
end

function filterObj:writeln( txt, baseIndent )
  self:write( txt )
  self:write( "\n" )
  self.needIndent = true
  self.indent = baseIndent
end

filterObj[TransUnit.nodeKind.None] = function ( self, node, parent, baseIndent )
  self:writeln( "-- none", baseIndent )
end

filterObj[TransUnit.nodeKind.Import] = function ( self, node, parent, baseIndent )
  local module = node.info
  local moduleName = string.gsub( module, ".*%.", "" )
  local moduleInfo = true
  self.moduleName2Info[moduleName] = moduleInfo
  if self.exeFlag then
    self:writeln( string.format( "local %s = _luneScript.loadModule( '%s' )", moduleName, module), baseIndent )
  else 
    self:writeln( string.format( "local %s = require( '%s' )", moduleName, module), baseIndent )
  end
end

filterObj[TransUnit.nodeKind.Root] = function ( self, node, parent, baseIndent )
  self:writeln( string.format( "--%s", self.streamName), baseIndent )
  self:writeln( "local moduleObj = {}", baseIndent )
  for __index, child in pairs( node.info.childlen ) do
    child:filter( self, node, baseIndent )
    self:writeln( "", baseIndent )
  end
  self:writeln( "local _className2InfoMap = {}", baseIndent )
  self:writeln( "moduleObj._className2InfoMap = _className2InfoMap", baseIndent )
  do
    local __sorted = {}
    local __map = self.className2InfoMap
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, className in ipairs( __sorted ) do
      classInfo = __map[ className ]
      do
        self:writeln( string.format( "local _classInfo%s = {}", className), baseIndent )
        self:writeln( string.format( "_className2InfoMap.%s = _classInfo%s", className, className), baseIndent )
        do
          local __sorted = {}
          local __map = classInfo
          for __key in pairs( __map ) do
            table.insert( __sorted, __key )
          end
          table.sort( __sorted )
          for __index, methodName in ipairs( __sorted ) do
            methodInfo = __map[ methodName ]
            do
              self:writeln( string.format( "_classInfo%s.%s = {", className, methodName), baseIndent )
              self:writeln( string.format( "  name='%s', staticFlag = %s, accessMode = '%s' }", methodName, methodInfo.staticFlag, methodInfo.accessMode), baseIndent )
            end
          end
        end
        
      end
    end
  end
  
  self:writeln( "return moduleObj", baseIndent )
end

filterObj[TransUnit.nodeKind.Block] = function ( self, node, parent, baseIndent )
  local word = ""
  if node.info.kind == "if" or node.info.kind == "elseif" then
    word = "then"
  elseif node.info.kind == "else" then
    word = ""
  elseif node.info.kind == "while" then
    word = "do"
  elseif node.info.kind == "repeat" then
    word = ""
  elseif node.info.kind == "for" then
    word = "do"
  elseif node.info.kind == "apply" then
    word = "do"
  elseif node.info.kind == "foreach" then
    word = "do"
  elseif node.info.kind == "func" then
    word = ""
  elseif node.info.kind == "{" then
    word = "do"
  end
  self:writeln( word, baseIndent + stepIndent )
  for __index, statement in pairs( node.info.stmtList ) do
    statement:filter( self, node, baseIndent + stepIndent )
    self:writeln( "", baseIndent + stepIndent )
  end
  self:setIndent( baseIndent )
  if node.info.kind == "{" then
    self:write( "end", baseIndent )
  end
end

filterObj[TransUnit.nodeKind.StmtExp] = function ( self, node, parent, baseIndent )
  node.info:filter( self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.DeclClass] = function ( self, node, parent, baseIndent )
  local classInfo = {}
  local className = node.info.name.txt
  self.className2InfoMap[className] = classInfo
  self:writeln( string.format( "local %s = {}", className ), baseIndent )
  if node.info.accessMode == "pub" then
    self:writeln( string.format( "moduleObj.%s = %s", className, className ), baseIndent )
  end
  for __index, field in pairs( node.info.fieldList ) do
    field:filter( self, node, baseIndent )
  end
end

filterObj[TransUnit.nodeKind.DeclMember] = function ( self, node, parent, baseIndent )
end

filterObj[TransUnit.nodeKind.DeclConstr] = function ( self, node, parent, baseIndent )
  local className = node.info.className.txt
  self:write( string.format( "function %s.new( ", className ) )
  local argTxt = ""
  for index, arg in pairs( node.info.argList ) do
    if index > 1 then
      self:write( ", " )
      argTxt = argTxt .. ", "
    end
    arg:filter( self, node, baseIndent )
    argTxt = argTxt .. arg.info.name.txt
  end
  self:writeln( " )", baseIndent + stepIndent )
  self:writeln( "local obj = {}", baseIndent + stepIndent )
  self:writeln( string.format( "setmetatable( obj, { __index = %s } )", className ), baseIndent + stepIndent )
  self:writeln( string.format( "return obj.__init and obj:__init( %s ) or nil;", argTxt ), baseIndent )
  self:writeln( "end", baseIndent )
  self:write( string.format( "function %s:__init(%s) ", className, argTxt ) )
  node.info.body:filter( self, node, baseIndent )
  self:writeln( "end", baseIndent )
end

filterObj[TransUnit.nodeKind.DeclMethod] = function ( self, node, parent, baseIndent )
  local classInfo = self.className2InfoMap[node.info.className.txt]
  local delimit = ":"
  if node.info.staticFlag then
    delimit = "."
  end
  local methodName = node.info.name.txt
  self:write( string.format( "function %s%s%s( ", node.info.className.txt, delimit, methodName) )
  classInfo[methodName] = {["staticFlag"] = node.info.staticFlag, ["accessMode"] = node.info.accessMode}
  for index, arg in pairs( node.info.argList ) do
    if index > 1 then
      self:write( ", " )
    end
    arg:filter( self, node, baseIndent )
  end
  self:write( " )", baseIndent )
  node.info.body:filter( self, node, baseIndent )
  self:writeln( "end", baseIndent )
end

filterObj[TransUnit.nodeKind.DeclVar] = function ( self, node, parent, baseIndent )
  if node.info.accessMode ~= "global" then
    self:write( "local " )
  end
  local varName = ""
  for index, var in pairs( node.info.varList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( var.name.txt )
  end
  if node.info.expList then
    self:write( " = " )
    node.info.expList:filter( self, node, baseIndent )
  end
  if node.info.accessMode == "pub" then
    self:writeln( "", baseIndent )
    for index, var in pairs( node.info.varList ) do
      self:writeln( string.format( "moduleObj.%s = %s", var.name.txt, var.name.txt ), baseIndent )
    end
  end
end

filterObj[TransUnit.nodeKind.DeclArg] = function ( self, node, parent, baseIndent )
  self:write( string.format( "%s", node.info.name.txt ) )
end

filterObj[TransUnit.nodeKind.DeclArgDDD] = function ( self, node, parent, baseIndent )
  self:write( "..." )
end

filterObj[TransUnit.nodeKind.ExpDDD] = function ( self, node, parent, baseIndent )
  self:write( "..." )
end

filterObj[TransUnit.nodeKind.DeclFunc] = function ( self, node, parent, baseIndent )
  local nameToken = node.info.name
  local name = nameToken and nameToken.txt or ""
  local letTxt = ""
  if node.info.accessMode ~= "global" and #name ~= 0 then
    letTxt = "local "
  end
  self:write( string.format( "%sfunction %s( ", letTxt, name ) )
  for index, arg in pairs( node.info.argList ) do
    if index > 1 then
      self:write( ", " )
    end
    arg:filter( self, node, baseIndent )
  end
  self:write( " )", baseIndent )
  node.info.body:filter( self, node, baseIndent )
  self:writeln( "end", baseIndent )
  if node.info.accessMode == "pub" then
    self:write( string.format( "moduleObj.%s = %s", name, name) )
  end
end

filterObj[TransUnit.nodeKind.RefType] = function ( self, node, parent, baseIndent )
  self:write( (node.info.refFlag and "&" or "" ) .. (node.info.mutFlag and "mut " or "" ) .. node.info.name.txt )
  if node.info.array == "array" then
    self:write( "[@]" )
  elseif node.info.array == "list" then
    self:write( "[]" )
  end
end

filterObj[TransUnit.nodeKind.If] = function ( self, node, parent, baseIndent )
  for index, val in pairs( node.info ) do
    if index == 1 then
      self:write( "if " )
      val.exp:filter( self, node, baseIndent )
    elseif val.kind == "elseif" then
      self:write( "elseif " )
      val.exp:filter( self, node, baseIndent )
    else 
      self:write( "else" )
    end
    self:write( " " )
    val.block:filter( self, node, baseIndent )
  end
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.While] = function ( self, node, parent, baseIndent )
  self:write( "while " )
  node.info.exp:filter( self, node, baseIndent )
  self:write( " " )
  node.info.block:filter( self, node, baseIndent )
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.Repeat] = function ( self, node, parent, baseIndent )
  self:write( "repeat " )
  node.info.block:filter( self, node, baseIndent )
  self:write( "until " )
  node.info.exp:filter( self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.For] = function ( self, node, parent, baseIndent )
  self:write( string.format( "for %s = ", node.info.val.txt ) )
  node.info.init:filter( self, node, baseIndent )
  self:write( ", " )
  node.info.to:filter( self, node, baseIndent )
  if node.info.delta then
    self:write( ", " )
    node.info.delta:filter( self, node, baseIndent )
  end
  self:write( " " )
  node.info.block:filter( self, node, baseIndent )
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.Apply] = function ( self, node, parent, baseIndent )
  self:write( "for " )
  for index, var in pairs( node.info.varList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( var.txt )
  end
  self:write( " in " )
  node.info.exp:filter( self, node, baseIndent )
  self:write( " " )
  node.info.block:filter( self, node, baseIndent )
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.Foreach] = function ( self, node, parent, baseIndent )
  self:write( "for " )
  self:write( node.info.key and node.info.key.txt or "__index" )
  self:write( ", " )
  self:write( node.info.val.txt )
  self:write( " in pairs( " )
  node.info.exp:filter( self, node, baseIndent )
  self:write( " ) " )
  node.info.block:filter( self, node, baseIndent )
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.Forsort] = function ( self, node, parent, baseIndent )
  self:writeln( "do", baseIndent + stepIndent )
  self:writeln( "local __sorted = {}", baseIndent + stepIndent )
  self:write( "local __map = " )
  node.info.exp:filter( self, node, baseIndent + stepIndent )
  self:writeln( "", baseIndent + stepIndent )
  self:writeln( "for __key in pairs( __map ) do", baseIndent + stepIndent * 2 )
  self:writeln( "table.insert( __sorted, __key )", baseIndent + stepIndent )
  self:writeln( "end", baseIndent + stepIndent )
  self:writeln( "table.sort( __sorted )", baseIndent + stepIndent )
  self:write( "for __index, " )
  local key = node.info.key and node.info.key.txt or "__key"
  self:write( key )
  self:writeln( " in ipairs( __sorted ) do", baseIndent + stepIndent * 2 )
  self:writeln( string.format( "%s = __map[ %s ]", node.info.val.txt, key ), baseIndent + stepIndent * 2 )
  node.info.block:filter( self, node, baseIndent + stepIndent * 2 )
  self:writeln( "end", baseIndent + stepIndent )
  self:writeln( "end", baseIndent )
  self:writeln( "end", baseIndent )
end

filterObj[TransUnit.nodeKind.ExpCall] = function ( self, node, parent, baseIndent )
  node.info.func:filter( self, node, baseIndent )
  self:write( "( " )
  if node.info.argList then
    node.info.argList:filter( self, node, baseIndent )
  end
  self:write( " )" )
end

filterObj[TransUnit.nodeKind.ExpList] = function ( self, node, parent, baseIndent )
  for index, exp in pairs( node.info ) do
    if index > 1 then
      self:write( ", " )
    end
    exp:filter( self, node, baseIndent )
  end
end

filterObj[TransUnit.nodeKind.ExpOp1] = function ( self, node, parent, baseIndent )
  local op = node.info.op.txt
  if op == "not" then
    op = op .. " "
  end
  self:write( op )
  node.info.exp:filter( self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.ExpCast] = function ( self, node, parent, baseIndent )
  node.info.exp:filter( self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.ExpParen] = function ( self, node, parent, baseIndent )
  self:write( "(" )
  node.info:filter( self, node, baseIndent )
  self:write( " )" )
end

filterObj[TransUnit.nodeKind.ExpOp2] = function ( self, node, parent, baseIndent )
  node.info.exp1:filter( self, node, baseIndent )
  self:write( " " .. node.info.op.txt .. " " )
  node.info.exp2:filter( self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.ExpRef] = function ( self, node, parent, baseIndent )
  self:write( node.info.txt )
end

filterObj[TransUnit.nodeKind.ExpRefItem] = function ( self, node, parent, baseIndent )
  if node.info.val.kind == TransUnit.nodeKind.LiteralString then
    self:write( "string.byte( " )
    node.info.val:filter( self, node, baseIndent )
    self:write( ", " )
    node.info.index:filter( self, node, baseIndent )
    self:write( " )" )
  else 
    node.info.val:filter( self, node, baseIndent )
    self:write( "[" )
    node.info.index:filter( self, node, baseIndent )
    self:write( "]" )
  end
end

filterObj[TransUnit.nodeKind.RefField] = function ( self, node, parent, baseIndent )
  node.info.prefix:filter( self, node, baseIndent )
  local delimit = "."
  if parent.kind == TransUnit.nodeKind.ExpCall then
    local prefixSymbol = node.info.prefix.info.txt
    if node.info.prefix.kind == TransUnit.nodeKind.ExpRef and (builtInModuleSet[prefixSymbol] or self.moduleName2Info[prefixSymbol] or self.className2InfoMap[prefixSymbol] ) then
      delimit = "."
    else 
      delimit = ":"
    end
  end
  self:write( delimit .. node.info.field.txt )
end

filterObj[TransUnit.nodeKind.Return] = function ( self, node, parent, baseIndent )
  self:write( "return " )
  node.info:filter( self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.LiteralList] = function ( self, node, parent, baseIndent )
  self:write( "{" )
  node.info:filter( self, node, baseIndent )
  self:write( "}" )
end

filterObj[TransUnit.nodeKind.LiteralMap] = function ( self, node, parent, baseIndent )
  self:write( "{" )
  for index, pair in pairs( node.info.pairList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( "[" )
    pair.key:filter( self, node, baseIndent )
    self:write( "] = " )
    pair.val:filter( self, node, baseIndent )
    index = index + 1
  end
  self:write( "}" )
end

filterObj[TransUnit.nodeKind.LiteralArray] = function ( self, node, parent, baseIndent )
  self:write( "{" )
  node.info:filter( self, node, baseIndent )
  self:write( "}" )
end

filterObj[TransUnit.nodeKind.LiteralChar] = function ( self, node, parent, baseIndent )
  self:write( string.format( "%g", node.info.num ) )
end

filterObj[TransUnit.nodeKind.LiteralInt] = function ( self, node, parent, baseIndent )
  self:write( string.format( "%d", node.info.num ) )
end

filterObj[TransUnit.nodeKind.LiteralReal] = function ( self, node, parent, baseIndent )
  self:write( string.format( "%s", node.info.num ) )
end

filterObj[TransUnit.nodeKind.LiteralString] = function ( self, node, parent, baseIndent )
  local txt = node.info.token.txt
  if string.find( txt, '^```' ) then
    txt = '[==[' .. txt:sub( 4, -4 ) .. ']==]'
  end
  if #node.info.argList > 0 then
    self:write( string.format( "string.format( %s, ", txt ) )
    for index, val in pairs( node.info.argList ) do
      if index > 1 then
        self:write( ", " )
      end
      val:filter( self, node, baseIndent )
    end
    self:write( ")" )
  else 
    self:write( txt )
  end
end

filterObj[TransUnit.nodeKind.LiteralBool] = function ( self, node, parent, baseIndent )
  self:write( node.info.txt )
end

filterObj[TransUnit.nodeKind.LiteralNil] = function ( self, node, parent, baseIndent )
  self:write( "nil" )
end

filterObj[TransUnit.nodeKind.Break] = function ( self, node, parent, baseIndent )
  self:write( "break" )
end

local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
local _classInfofilterObj = {}
_className2InfoMap.filterObj = _classInfofilterObj
_classInfofilterObj.setIndent = {
  name='setIndent', staticFlag = false, accessMode = 'pri' }
_classInfofilterObj.write = {
  name='write', staticFlag = false, accessMode = 'pri' }
_classInfofilterObj.writeln = {
  name='writeln', staticFlag = false, accessMode = 'pri' }
return moduleObj
