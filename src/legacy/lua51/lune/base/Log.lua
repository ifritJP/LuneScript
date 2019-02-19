--lune/base/Log.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Log'
if not _lune then
   _lune = {}
end
local Level = {}
_moduleObj.Level = Level
Level._val2NameMap = {}
function Level:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Level.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function Level._from( val )
   if Level._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
Level.__allList = {}
function Level.get__allList()
   return Level.__allList
end

Level.Fatal = 0
Level._val2NameMap[0] = 'Fatal'
Level.__allList[1] = Level.Fatal
Level.Err = 1
Level._val2NameMap[1] = 'Err'
Level.__allList[2] = Level.Err
Level.Warn = 2
Level._val2NameMap[2] = 'Warn'
Level.__allList[3] = Level.Warn
Level.Info = 3
Level._val2NameMap[3] = 'Info'
Level.__allList[4] = Level.Info
Level.Debug = 4
Level._val2NameMap[4] = 'Debug'
Level.__allList[5] = Level.Debug
Level.Trace = 5
Level._val2NameMap[5] = 'Trace'
Level.__allList[6] = Level.Trace

local name2levelMap = {}
name2levelMap["fatal"] = Level.Fatal
name2levelMap["error"] = Level.Err
name2levelMap["warn"] = Level.Warn
name2levelMap["info"] = Level.Info
name2levelMap["debug"] = Level.Debug
name2levelMap["trace"] = Level.Trace
local function str2level( txt )

   return name2levelMap[txt]
end
_moduleObj.str2level = str2level
local outputLevel = Level.Err
local function setLevel( level )

   outputLevel = level
end
_moduleObj.setLevel = setLevel

local function log( level, funcName, lineNo, callback )

   if level <= outputLevel then
      io.stderr:write( string.format( "%s:%s:%d:", Level:_getTxt( level)
      , funcName, lineNo) )
      io.stderr:write( callback(  ) )
      io.stderr:write( "\n" )
   end
   
end
_moduleObj.log = log

return _moduleObj
