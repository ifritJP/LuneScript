--lune/base/MetaLoader.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@MetaLoader'
local _lune = {}
if _lune2 then
   _lune = _lune2
end
function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
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
local Str = _lune.loadModule( 'lune.base.Str' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Meta = _lune.loadModule( 'lune.base.Meta' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Depend = _lune.loadModule( 'lune.base.Depend' )

local SymbolInfo = {}
function SymbolInfo.setmeta( obj )
  setmetatable( obj, { __index = SymbolInfo  } )
end
function SymbolInfo.new( name, val )
   local obj = {}
   SymbolInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, val )
   end
   return obj
end
function SymbolInfo:__init( name, val )

   self.name = name
   self.val = val
end
function SymbolInfo:get_name()
   return self.name
end
function SymbolInfo:get_val()
   return self.val
end


local Scope = {}
function Scope.new( parent )
   local obj = {}
   Scope.setmeta( obj )
   if obj.__init then obj:__init( parent ); end
   return obj
end
function Scope:__init(parent) 
   self.parent = parent
   self.name2val = {}
end
function Scope:add( name, val )

   local symbolInfo = SymbolInfo.new(name, val)
   self.name2val[name] = symbolInfo
   return symbolInfo
end
function Scope:get( name )

   do
      local _exp = self.name2val[name]
      if _exp ~= nil then
         return _exp
      end
   end
   
   do
      local _exp = self.parent
      if _exp ~= nil then
         return _exp:get( name )
      end
   end
   
   return nil
end
function Scope.setmeta( obj )
  setmetatable( obj, { __index = Scope  } )
end
function Scope:get_parent()
   return self.parent
end


local quotedChar2Code = {['a'] = 7, ['b'] = 8, ['t'] = 9, ['n'] = 10, ['v'] = 11, ['f'] = 12, ['r'] = 13, ['\\'] = 92, ['"'] = 34, ["'"] = 39}

local Loader = {}
function Loader.new( parser )
   local obj = {}
   Loader.setmeta( obj )
   if obj.__init then obj:__init( parser ); end
   return obj
end
function Loader:__init(parser) 
   self.parser = parser
   self.curScope = Scope.new(nil)
end
function Loader:analyzeTable(  )

   local map = {}
   local index = 1
   while true do
      local token = self.parser:getTokenNoErr(  )
      if token.kind == Parser.TokenKind.Eof then
         return nil, "eof"
      end
      
      do
         local _switchExp = token.txt
         if _switchExp == "}" then
            return map, nil
         elseif _switchExp == "," then
         else 
            
               local item = nil
               do
                  local _switchExp = token.kind
                  if _switchExp == Parser.TokenKind.Symb then
                     local nextToken = self.parser:getTokenNoErr(  )
                     self.parser:pushbackToken( nextToken )
                     if nextToken.txt == "=" then
                        item = token.txt
                     else
                      
                        return nil, string.format( "not support -- %s", token.txt)
                     end
                     
                  elseif _switchExp == Parser.TokenKind.Dlmt then
                     if token.txt == "{" then
                        local err
                        
                        item, err = self:analyzeTable(  )
                        if err then
                           return nil, err
                        end
                        
                     else
                      
                        return nil, string.format( "not support -- %s", token.txt)
                     end
                     
                  else 
                     
                        self.parser:pushback(  )
                        local err
                        
                        item, err = self:analyzeExp(  )
                        if err then
                           return nil, err
                        end
                        
                  end
               end
               
               if not item then
                  self.parser:pushbackToken( token )
               end
               
               local nextToken = self.parser:getTokenNoErr(  )
               if nextToken.txt == "=" then
                  local val, work = self:analyzeExp(  )
                  if work then
                     return nil, work
                  end
                  
                  if item ~= nil and val ~= nil then
                     map[item] = val
                  end
                  
               else
                
                  self.parser:pushback(  )
                  if item ~= nil then
                     map[index] = item
                  end
                  
                  index = index + 1
               end
               
         end
      end
      
   end
   
end
function Loader:analyzeExp(  )

   while true do
      local token = self.parser:getTokenNoErr(  )
      if token.kind == Parser.TokenKind.Eof then
         return nil, "eof"
      end
      
      local txt = token.txt
      do
         local _switchExp = token.kind
         if _switchExp == Parser.TokenKind.Dlmt then
            if txt == "{" then
               return self:analyzeTable(  )
            end
            
            return nil, string.format( "illegal delimit -- %s", txt)
         elseif _switchExp == Parser.TokenKind.Char then
            if #txt == 1 then
               return string.byte( txt, 1 ), nil
            end
            
            return _lune.unwrap( quotedChar2Code[txt:sub( 2, 2 )]), nil
         elseif _switchExp == Parser.TokenKind.Int then
            return math.floor((_lune.unwrapDefault( tonumber( txt ), 0) )), nil
         elseif _switchExp == Parser.TokenKind.Real then
            return _lune.unwrapDefault( tonumber( txt ), 0.0), nil
         elseif _switchExp == Parser.TokenKind.Str then
            if string.byte( txt, 1 ) == 96 then
               return txt:sub( 4, -4 ), nil
            end
            
            return txt:sub( 2, -2 ), nil
         elseif _switchExp == Parser.TokenKind.Symb then
            do
               local symbolInfo = self.curScope:get( txt )
               if symbolInfo ~= nil then
                  return symbolInfo:get_val(), nil
               end
            end
            
            return nil, string.format( "not found -- %s", txt)
         elseif _switchExp == Parser.TokenKind.Kywd then
            do
               local _switchExp = token.txt
               if _switchExp == "true" then
                  return true, nil
               elseif _switchExp == "false" then
                  return false, nil
               end
            end
            
            return nil, string.format( "illegal keyword -- %s", token.txt)
         elseif _switchExp == Parser.TokenKind.Cmnt then
         elseif _switchExp == Parser.TokenKind.Eof or _switchExp == Parser.TokenKind.Type then
            return nil, string.format( "illegal kind -- %s", Parser.TokenKind:_getTxt( token.kind)
            )
         elseif _switchExp == Parser.TokenKind.Ope then
            return nil, string.format( "not support -- %s", txt)
         end
      end
      
   end
   
end
function Loader:processDeclVar(  )

   local token = self.parser:getTokenNoErr(  )
   if token.kind ~= Parser.TokenKind.Symb then
      return string.format( "no synbol -- %s", token.txt)
   end
   
   local nextToken = self.parser:getTokenNoErr(  )
   if nextToken.txt ~= "=" then
      self.parser:pushback(  )
      return nil
   end
   
   local val, err = self:analyzeExp(  )
   if err then
      return err
   end
   
   
   self.curScope:add( token.txt, val )
   return nil
end
function Loader:processExpSym( symToken )

   local symbolInfo = self.curScope:get( symToken.txt )
   if  nil == symbolInfo then
      local _symbolInfo = symbolInfo
   
      return string.format( "not found -- %s", symToken.txt)
   end
   
   local nextToken = self.parser:getTokenNoErr(  )
   local exp = symbolInfo:get_val()
   local index = nil
   
   while true do
      if nextToken.kind == Parser.TokenKind.Eof then
         return "eof"
      end
      
      do
         local _switchExp = nextToken.txt
         if _switchExp == "." then
            local fieldToken = self.parser:getTokenNoErr(  )
            if exp ~= nil then
               index = fieldToken.txt
            else
               return "nil access"
            end
            
         elseif _switchExp == "[" then
            local indexExp, err = self:analyzeExp(  )
            if exp ~= nil then
               index = indexExp
            else
               return string.format( "illegal index -- %s", err)
            end
            
            local closeToken = self.parser:getTokenNoErr(  )
            if closeToken.txt ~= "]" then
               return string.format( "illegal token -- %s", closeToken.txt)
            end
            
         elseif _switchExp == "=" then
            local val, err = self:analyzeExp(  )
            if err then
               return err
            end
            
            if index ~= nil and exp ~= nil then
               local map = exp
               map[index] = val
            end
            
            return nil
         end
      end
      
      nextToken = self.parser:getTokenNoErr(  )
   end
   
end
function Loader:processStmt(  )

   while true do
      local token = self.parser:getTokenNoErr(  )
      if token.kind == Parser.TokenKind.Eof then
         return nil, "eof"
      end
      
      do
         local _switchExp = token.kind
         if _switchExp == Parser.TokenKind.Symb then
            do
               local _exp = self:processExpSym( token )
               if _exp ~= nil then
                  return nil, _exp
               end
            end
            
         elseif _switchExp == Parser.TokenKind.Kywd then
            do
               local _switchExp = token.txt
               if _switchExp == "local" then
                  do
                     local err = self:processDeclVar(  )
                     if err ~= nil then
                        return nil, err
                     end
                  end
                  
               elseif _switchExp == "do" then
                  self.curScope = Scope.new(self.curScope)
               elseif _switchExp == "end" then
                  do
                     local parent = self.curScope:get_parent()
                     if parent ~= nil then
                        self.curScope = parent
                     else
                        return nil, "underflow scope"
                     end
                  end
                  
               elseif _switchExp == "return" then
                  return self:analyzeExp(  )
               end
            end
            
         end
      end
      
   end
   
end
function Loader:process(  )

   local val, err = self:processStmt(  )
   if err ~= nil then
      local pos = self.parser:getTokenNoErr(  ).pos
      return nil, string.format( "%d:%d:%s", pos.lineNo, pos.column, err)
   end
   
   return val, err
end
function Loader.setmeta( obj )
  setmetatable( obj, { __index = Loader  } )
end


local function loadFromStream( path, stream )

   local streamParser = Parser.StreamParser.new(stream, path, true)
   local parser = Parser.DefaultPushbackParser.new(streamParser)
   local loader = Loader.new(parser)
   
   local item, err = loader:process(  )
   if err then
      return nil, err
   end
   
   
   
   
   if item ~= nil then
      local map = item
      
      
      local formatVersion
      
      do
         local workVal = map["__formatVersion"]
         if workVal ~= nil then
            formatVersion = workVal
         else
            return nil, string.format( "error -- %s", 'formatVersion')
         end
      end
      
      
      
      local enableTest
      
      do
         local workVal = map["__enableTest"]
         if workVal ~= nil then
            enableTest = workVal
         else
            return nil, string.format( "error -- %s", 'enableTest')
         end
      end
      
      
      
      local buildId
      
      do
         local workVal = map["__buildId"]
         if workVal ~= nil then
            buildId = workVal
         else
            return nil, string.format( "error -- %s", 'buildId')
         end
      end
      
      
      
      local typeId2ClassInfoMap
      
      do
         local workVal = map["__typeId2ClassInfoMap"]
         if workVal ~= nil then
            typeId2ClassInfoMap = workVal
         else
            return nil, string.format( "error -- %s", 'typeId2ClassInfoMap')
         end
      end
      
      
      
      local typeInfoList
      
      do
         local workVal = map["__typeInfoList"]
         if workVal ~= nil then
            typeInfoList = workVal
         else
            return nil, string.format( "error -- %s", 'typeInfoList')
         end
      end
      
      
      
      local varName2InfoMap
      
      do
         local workVal = map["__varName2InfoMap"]
         if workVal ~= nil then
            varName2InfoMap = workVal
         else
            return nil, string.format( "error -- %s", 'varName2InfoMap')
         end
      end
      
      
      
      local moduleTypeId
      
      do
         local workVal = map["__moduleTypeId"]
         if workVal ~= nil then
            moduleTypeId = math.floor(workVal)
         else
            return nil, string.format( "error -- %s", 'moduleTypeId')
         end
      end
      
      
      
      local moduleSymbolKind
      
      do
         local workVal = map["__moduleSymbolKind"]
         if workVal ~= nil then
            moduleSymbolKind = math.floor(workVal)
         else
            return nil, string.format( "error -- %s", 'moduleSymbolKind')
         end
      end
      
      
      
      local moduleMutable
      
      do
         local workVal = map["__moduleMutable"]
         if workVal ~= nil then
            moduleMutable = workVal
         else
            return nil, string.format( "error -- %s", 'moduleMutable')
         end
      end
      
      
      
      local dependModuleMap
      
      do
         local workVal = map["__dependModuleMap"]
         if workVal ~= nil then
            dependModuleMap = workVal
         else
            return nil, string.format( "error -- %s", 'dependModuleMap')
         end
      end
      
      
      
      local dependIdMap
      
      do
         local workVal = map["__dependIdMap"]
         if workVal ~= nil then
            dependIdMap = workVal
         else
            return nil, string.format( "error -- %s", 'dependIdMap')
         end
      end
      
      
      
      local macroName2InfoMap
      
      do
         local workVal = map["__macroName2InfoMap"]
         if workVal ~= nil then
            macroName2InfoMap = workVal
         else
            return nil, string.format( "error -- %s", 'macroName2InfoMap')
         end
      end
      
      
      
      local subModuleMap
      
      do
         local workVal = map["__subModuleMap"]
         if workVal ~= nil then
            subModuleMap = workVal
         else
            return nil, string.format( "error -- %s", 'subModuleMap')
         end
      end
      
      
      
      local hasTest
      
      do
         local workVal = map["__hasTest"]
         if workVal ~= nil then
            hasTest = workVal
         else
            return nil, string.format( "error -- %s", 'hasTest')
         end
      end
      
      
      
      return Meta._MetaInfo.new(formatVersion, enableTest, buildId, typeId2ClassInfoMap, typeInfoList, varName2InfoMap, moduleTypeId, moduleSymbolKind, moduleMutable, dependModuleMap, dependIdMap, macroName2InfoMap, hasTest, subModuleMap), nil
   end
   
   
   return nil, "error"
end
_moduleObj.loadFromStream = loadFromStream

local function formToMeta( func, err )

   if func ~= nil then
      local obj = func(  )
      local val = (obj )
      if  nil == val then
         local _val = val
      
         return nil, "error"
      end
      
      return val, nil
   else
      return nil, err
   end
   
end

local function loadFromCode( path, code )

   if not Depend.useMetaLoader(  ) then
      return formToMeta( _lune.loadstring52( code, nil ), nil )
   end
   
   return loadFromStream( path, Parser.TxtStream.new(code) )
end
_moduleObj.loadFromCode = loadFromCode

local function load( path )

   if not Depend.useMetaLoader(  ) then
      return formToMeta( loadfile( path ) )
   end
   
   
   local stream = io.open( path, "r" )
   if  nil == stream then
      local _stream = stream
   
      return nil, string.format( "open error -- %s", path)
   end
   
   return loadFromStream( path, stream )
end
_moduleObj.load = load





return _moduleObj
