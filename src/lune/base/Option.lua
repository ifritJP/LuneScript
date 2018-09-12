--lune/base/Option.lns
local _moduleObj = {}
if not _ENV._lune then
   _lune = {}
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

local Parser = require( 'lune.base.Parser' )

local Util = require( 'lune.base.Util' )

local version = "00.01"

_moduleObj.version = version

local Option = {}
_moduleObj.Option = Option
function Option.new(  )
  local obj = {}
  Option.setmeta( obj )
  if obj.__init then obj:__init(  ); end
return obj
end
function Option:__init() 
  self.validProf = false
  self.mode = ""
  self.scriptPath = ""
end
function Option.setmeta( obj )
  setmetatable( obj, { __index = Option  } )
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
              option.analyzePos = Parser.Position.new(_lune.unwrap( lineNo), _lune.unwrap( column))
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
    Parser.StreamParser.setStdinStream( _lune.unwrap( option.analyzeModule) )
  end
  return option
end
_moduleObj.analyze = analyze
return _moduleObj
