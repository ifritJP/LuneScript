--lune/base/Option.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Option'
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

local ModeKind = {}
_moduleObj.ModeKind = ModeKind
ModeKind._val2NameMap = {}
function ModeKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ModeKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function ModeKind._from( val )
   if ModeKind._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
ModeKind.Unknown = ''
ModeKind._val2NameMap[''] = 'Unknown'
ModeKind.Token = 'token'
ModeKind._val2NameMap['token'] = 'Token'
ModeKind.Ast = 'ast'
ModeKind._val2NameMap['ast'] = 'Ast'
ModeKind.Diag = 'diag'
ModeKind._val2NameMap['diag'] = 'Diag'
ModeKind.Complete = 'comp'
ModeKind._val2NameMap['comp'] = 'Complete'
ModeKind.Lua = 'lua'
ModeKind._val2NameMap['lua'] = 'Lua'
ModeKind.LuaMeta = 'LUA'
ModeKind._val2NameMap['LUA'] = 'LuaMeta'
ModeKind.Save = 'save'
ModeKind._val2NameMap['save'] = 'Save'
ModeKind.SaveMeta = 'SAVE'
ModeKind._val2NameMap['SAVE'] = 'SaveMeta'
ModeKind.Exec = 'exe'
ModeKind._val2NameMap['exe'] = 'Exec'
ModeKind.Glue = 'glue'
ModeKind._val2NameMap['glue'] = 'Glue'

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
   self.mode = ModeKind.Unknown
   self.scriptPath = ""
end
function Option.setmeta( obj )
  setmetatable( obj, { __index = Option  } )
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
            do
               local mode = ModeKind._from( arg )
               if mode ~= nil then
                  option.mode = mode
               else
                  Util.err( string.format( "unknown mode -- %s", arg) )
                  os.exit( 1 )
               end
            end
            
         else
          
            do
               local _switchExp = (option.mode )
               if _switchExp == ModeKind.Complete then
                  if not option.analyzeModule then
                     option.analyzeModule = arg
                  elseif not lineNo then
                     lineNo = math.floor(tonumber( arg ))
                  elseif not column then
                     column = math.floor(tonumber( arg ))
                     option.analyzePos = Parser.Position.new(_lune.unwrap( lineNo), _lune.unwrap( column))
                  end
                  
               elseif _switchExp == ModeKind.Save or _switchExp == ModeKind.SaveMeta or _switchExp == ModeKind.Glue then
                  option.outputDir = arg
               else 
                  
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
