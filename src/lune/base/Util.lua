--lune/base/Util.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Util'
local _lune = {}
if _lune1 then
   _lune = _lune1
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

if not _lune1 then
   _lune1 = _lune
end
local Depend = _lune.loadModule( 'lune.base.Depend' )
local debugFlag = true
local function setDebugFlag( flag )

   debugFlag = flag
end
_moduleObj.setDebugFlag = setDebugFlag
local errorCode = 1
local function setErrorCode( code )

   errorCode = code
end
_moduleObj.setErrorCode = setErrorCode
local function errorLog( message )

   io.stderr:write( message .. "\n" )
end
_moduleObj.errorLog = errorLog
local function err( message )

   if debugFlag then
      error( message )
   end
   
   errorLog( message )
   os.exit( errorCode )
end
_moduleObj.err = err
local OrderedSet = {}
_moduleObj.OrderedSet = OrderedSet
function OrderedSet.new(  )
   local obj = {}
   OrderedSet.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function OrderedSet:__init() 
   self.set = {}
   self.list = {}
end
function OrderedSet:add( val )

   if not _lune._Set_has(self.set, val ) then
      self.set[val]= true
      table.insert( self.list, val )
   end
   
end
function OrderedSet:clone(  )

   local obj = OrderedSet.new()
   for __index, val in pairs( self.list ) do
      obj.set[val]= true
      table.insert( obj.list, val )
   end
   
   return obj
end
function OrderedSet.setmeta( obj )
  setmetatable( obj, { __index = OrderedSet  } )
end
function OrderedSet:get_list()
   return self.list
end

local memStream = {}
setmetatable( memStream, { ifList = {oStream,} } )
_moduleObj.memStream = memStream
function memStream.new(  )
   local obj = {}
   memStream.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function memStream:__init() 
   self.txt = ""
end
function memStream:write( val )

   self.txt = self.txt .. val
   return self, nil
end
function memStream:close(  )

end
function memStream:flush(  )

end
function memStream.setmeta( obj )
  setmetatable( obj, { __index = memStream  } )
end
function memStream:get_txt()
   return self.txt
end

local SourceStream = {}
_moduleObj.SourceStream = SourceStream
function SourceStream.setmeta( obj )
  setmetatable( obj, { __index = SourceStream  } )
end
function SourceStream.new(  )
   local obj = {}
   SourceStream.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function SourceStream:__init(  )

end

local SimpleSourceOStream = {}
setmetatable( SimpleSourceOStream, { ifList = {SourceStream,} } )
_moduleObj.SimpleSourceOStream = SimpleSourceOStream
function SimpleSourceOStream.new( stream, headStream, stepIndent )
   local obj = {}
   SimpleSourceOStream.setmeta( obj )
   if obj.__init then obj:__init( stream, headStream, stepIndent ); end
   return obj
end
function SimpleSourceOStream:__init(stream, headStream, stepIndent) 
   self.srcStream = stream
   self.nowStream = stream
   self.headStream = _lune.unwrapDefault( headStream, stream)
   self.needIndent = true
   self.curLineNo = 0
   self.stepIndent = stepIndent
   self.indentQueue = {0}
end
function SimpleSourceOStream:get_indent(  )

   if #self.indentQueue > 0 then
      return self.indentQueue[#self.indentQueue]
   end
   
   return 0
end
function SimpleSourceOStream:writeRaw( txt )

   local stream = self.nowStream
   if self.needIndent then
      stream:write( string.rep( " ", self:get_indent() ) )
      self.needIndent = false
   end
   
   for cr in string.gmatch( txt, "\n" ) do
      self.curLineNo = self.curLineNo + 1
   end
   
   stream:write( txt )
end
function SimpleSourceOStream:write( txt )

   while true do
      do
         local index = string.find( txt, "\n" )
         if index ~= nil then
            self:writeRaw( txt:sub( 1, index ) )
            txt = txt:sub( index + 1 )
         else
            break
         end
      end
      
   end
   
   if #txt > 0 then
      self:writeRaw( txt )
   end
   
end
function SimpleSourceOStream:writeln( txt )

   self:write( txt )
   self:write( "\n" )
   self.needIndent = true
end
function SimpleSourceOStream:pushIndent( newIndent )

   local indent = _lune.unwrapDefault( newIndent, self:get_indent() + self.stepIndent)
   table.insert( self.indentQueue, indent )
end
function SimpleSourceOStream:popIndent(  )

   if #self.indentQueue == 0 then
      err( "self.indentQueue == 0" )
   end
   
   table.remove( self.indentQueue )
end
function SimpleSourceOStream:switchToHeader(  )

   self.nowStream = self.headStream
end
function SimpleSourceOStream:returnToSource(  )

   self.nowStream = self.srcStream
end
function SimpleSourceOStream.setmeta( obj )
  setmetatable( obj, { __index = SimpleSourceOStream  } )
end

local function log( message )

   if debugFlag then
      errorLog( message )
   end
   
end
_moduleObj.log = log
local function printStackTrace(  )

   for level = 2, 8 do
      do
         local debugInfo = debug.getinfo( level )
         if debugInfo ~= nil then
            errorLog( string.format( "-- %s %s", debugInfo['short_src'], debugInfo['currentline']) )
         end
      end
      
   end
   
end
_moduleObj.printStackTrace = printStackTrace
local function profile( validTest, func, path )

   if not validTest then
      return func(  )
   end
   
   local ProFi = require( 'ProFi' )
   ProFi:start(  )
   local result = func(  )
   ProFi:stop(  )
   ProFi:writeReport( path )
   return result
end
_moduleObj.profile = profile
local function getReadyCode( lnsPath, luaPath )

   local luaTime, lnsTime = Depend.getFileLastModifiedTime( luaPath ), Depend.getFileLastModifiedTime( lnsPath )
   if  nil == luaTime or  nil == lnsTime then
      local _luaTime = luaTime
      local _lnsTime = lnsTime
   
      return false
   end
   
   return luaTime >= lnsTime
end
_moduleObj.getReadyCode = getReadyCode
local function existFile( path )

   local fileObj = io.open( path )
   if  nil == fileObj then
      local _fileObj = fileObj
   
      return false
   end
   
   fileObj:close(  )
   return true
end
_moduleObj.existFile = existFile
return _moduleObj
