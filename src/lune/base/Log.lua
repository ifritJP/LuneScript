--lune/base/Log.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Log'
local _lune = {}
if _lune3 then
   _lune = _lune3
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

if not _lune3 then
   _lune3 = _lune
end


local Async = _lune.loadModule( 'lune.base.Async' )
local Depend = _lune.loadModule( 'lune.base.Depend' )

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
Level.Log = 3
Level._val2NameMap[3] = 'Log'
Level.__allList[4] = Level.Log
Level.Info = 4
Level._val2NameMap[4] = 'Info'
Level.__allList[5] = Level.Info
Level.Debug = 5
Level._val2NameMap[5] = 'Debug'
Level.__allList[6] = Level.Debug
Level.Trace = 6
Level._val2NameMap[6] = 'Trace'
Level.__allList[7] = Level.Trace


local name2levelMap


do
   local work = {}
   work["fatal"] = Level.Fatal
   work["error"] = Level.Err
   work["warn"] = Level.Warn
   work["log"] = Level.Log
   work["info"] = Level.Info
   work["debug"] = Level.Debug
   work["trace"] = Level.Trace
   name2levelMap = work
end


local function str2level( txt )

   return name2levelMap[txt]
end
_moduleObj.str2level = str2level


local Control = {}
function Control:log( level, funcName, lineNo, callback )

   local logStream = io.stderr
   if level <= self.level then
      local nowClock = os.clock(  )
      
      logStream:write( string.format( "%6d:%s:%s:%d:", math.floor((nowClock * 1000 )), Level:_getTxt( level)
      , funcName, lineNo) )
      logStream:write( callback(  ) )
      logStream:write( "\n" )
   end
   
end
function Control:direct( level, funcName, lineNo, mess )

   self:log( level, funcName, lineNo, function (  )
   
      return mess
   end )
end
function Control.setmeta( obj )
  setmetatable( obj, { __index = Control  } )
end
function Control.new( level )
   local obj = {}
   Control.setmeta( obj )
   if obj.__init then
      obj:__init( level )
   end
   return obj
end
function Control:__init( level )

   self.level = level
end


local control = Control.new(Level.Err)

local function setLevel( level )

   control = Control.new(level)
   if level >= Level.Log then
      Depend.setRuntimeLog( true )
   end
   
end
_moduleObj.setLevel = setLevel

local function log( level, funcName, lineNo, callback )

   control:log( level, funcName, lineNo, callback )
end
_moduleObj.log = log

local function direct( level, funcName, lineNo, mess )

   control:direct( level, funcName, lineNo, mess )
end
_moduleObj.direct = direct







return _moduleObj
