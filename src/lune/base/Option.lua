--lune/base/Option.lns
local moduleObj = {}
local function _lune_nilacc( val, fieldName, access, ... )
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
function _lune_unwrap( val )
  if val == nil then
     _luneScript.error( 'unwrap val is nil' )
  end
  return val
end
function _lune_unwrapDefault( val, defval )
  if val == nil then
     return defval
  end
  return val
end

local Parser = require( 'lune.base.Parser' )

local Util = require( 'lune.base.Util' )

local version = "00.01"

moduleObj.version = version

local Option = {}
moduleObj.Option = Option
function Option.new(  )
  local obj = {}
  setmetatable( obj, { __index = Option } )
  if obj.__init then obj:__init(  ); end
return obj
end
function Option:__init() 
  self.validProf = false
  self.mode = ""
  self.scriptPath = ""
end
do
  end

local function analyze( argList )

  if #argList < 2 then
    print( [==[
usage: [-prof] src.lns mode [mode-option]
  - src.lns ast
  - src.lns comp [-i] module line column
  - src.lns [lua|LUA] 
  - src.lns [save|SAVE] output-dir
  - src.lns exe
]==] )
    os.exit( 1 )
  end
  local option = Option.new()
  
  local useStdInFlag = false
  
  local lineNo = nil
  
  local column = nil
  
  local index = 1
  
  while #argList >= index do
    local arg = argList[index]
    
    if arg:find( "^-" ) then
      do
        local _switchExp = (arg )
        if _switchExp == "-i" then
          useStdInFlag = true
        elseif _switchExp == "-prof" then
          option.validProf = true
        elseif _switchExp == "--nodebug" then
          Util.setDebugFlag( false )
        else 
          Util.log( string.format( "unknown option -- %s", arg) )
          os.exit( 1 )
        end
      end
      
    else 
      if option.scriptPath == "" then
        option.scriptPath = arg
      elseif option.mode == "" then
        option.mode = arg
      else 
        do
          local _switchExp = (option.mode )
          if _switchExp == "comp" then
            if not option.analyzeModule then
              option.analyzeModule = arg
            elseif not lineNo then
              lineNo = math.floor(tonumber( arg ))
            elseif not column then
              column = math.floor(tonumber( arg ))
              option.analyzePos = Parser.Position.new(_lune_unwrap( lineNo), _lune_unwrap( column))
            end
          elseif _switchExp == "save" or _switchExp == "SAVE" then
            option.outputDir = arg
          else 
            Util.err( string.format( "unknown option for this mode (%s) -- %s", option.mode, arg) )
          end
        end
        
      end
    end
    index = index + 1
  end
  if useStdInFlag and option.analyzeModule then
    Parser.StreamParser.setStdinStream( _lune_unwrap( option.analyzeModule) )
  end
  return option
end
moduleObj.analyze = analyze
return moduleObj
