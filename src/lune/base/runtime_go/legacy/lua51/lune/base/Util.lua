--lune/base/Util.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Util'
local _lune = {}
if _lune7 then
   _lune = _lune7
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

if not _lune7 then
   _lune7 = _lune
end


local Depend = _lune.loadModule( 'lune.base.Depend' )
local Log = _lune.loadModule( 'lune.base.Log' )
local Str = _lune.loadModule( 'lune.base.Str' )

local consoleOStream = io.stdout
local errStream = io.stderr

local ConsoleAdapter = {}
setmetatable( ConsoleAdapter, { ifList = {oStream,} } )
function ConsoleAdapter:write( txt )

   self.writer( txt )
   return self, nil
end
function ConsoleAdapter:flush(  )

end
function ConsoleAdapter:close(  )

end
function ConsoleAdapter._setmeta( obj )
  setmetatable( obj, { __index = ConsoleAdapter  } )
end
function ConsoleAdapter._new( writer )
   local obj = {}
   ConsoleAdapter._setmeta( obj )
   if obj.__init then
      obj:__init( writer )
   end
   return obj
end
function ConsoleAdapter:__init( writer )

   self.writer = writer
end


local function setConsoleOStream( stream )

   do
      consoleOStream = stream
   end
   
end
_moduleObj.setConsoleOStream = setConsoleOStream

local function setConsoleOStreamWithWriter( writer )

   setConsoleOStream( ConsoleAdapter._new(writer) )
end
_moduleObj.setConsoleOStreamWithWriter = setConsoleOStreamWithWriter

local function println( ... )

   local list = {...}
   do
      for index, arg in pairs( list ) do
         consoleOStream:write( string.format( "%s", tostring( arg)) )
         if index ~= #list then
            consoleOStream:write( "\t" )
         else
          
            consoleOStream:write( "\n" )
         end
         
      end
      
   end
   
end
_moduleObj.println = println

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

   do
      errStream:write( message .. "\n" )
   end
   
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
local function splitStr( txt, pattern )

   local list = {}
   do
      for token in string.gmatch( txt, pattern ) do
         table.insert( list, token )
      end
      
   end
   
   return list
end
_moduleObj.splitStr = splitStr

local function splitModule( modPath )

   return splitStr( modPath, '[^%./:]+' )
end
_moduleObj.splitModule = splitModule

local OrderedSet = {}
_moduleObj.OrderedSet = OrderedSet
function OrderedSet._new(  )
   local obj = {}
   OrderedSet._setmeta( obj )
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
      return true
   end
   
   return false
end
function OrderedSet:clone(  )

   local obj = OrderedSet._new()
   for __index, val in ipairs( self.list ) do
      obj.set[val]= true
      table.insert( obj.list, val )
   end
   
   return obj
end
function OrderedSet:has( val )

   return _lune._Set_has(self.set, val )
end
function OrderedSet:removeLast(  )

   if #self.list == 0 then
      err( "empty" )
   end
   
   self.set[self.list[#self.list]]= nil
   table.remove( self.list )
end
function OrderedSet._setmeta( obj )
  setmetatable( obj, { __index = OrderedSet  } )
end
function OrderedSet:get_list()
   return self.list
end


local OrderdMap = {}
_moduleObj.OrderdMap = OrderdMap
function OrderdMap._new(  )
   local obj = {}
   OrderdMap._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function OrderdMap:__init() 
   self.map = {}
   self.keyList = {}
end
function OrderdMap:clear(  )

   self.map = {}
   self.keyList = {}
end
function OrderdMap:add( key, val, overwrite )

   if self.map[key] then
      if overwrite then
         self.map[key] = val
      end
      
      return 
   end
   
   self.map[key] = val
   table.insert( self.keyList, key )
end
function OrderdMap._setmeta( obj )
  setmetatable( obj, { __index = OrderdMap  } )
end
function OrderdMap:get_map()
   return self.map
end
function OrderdMap:get_keyList()
   return self.keyList
end


local memStream = {}
setmetatable( memStream, { ifList = {oStream,} } )
_moduleObj.memStream = memStream
function memStream._new(  )
   local obj = {}
   memStream._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function memStream:__init() 
   self.txt = Str.Builder._new()
end
function memStream:get_txt(  )

   self.txt:flush(  )
   return self.txt:get_txt()
end
function memStream:write( val )

   self.txt:add( val )
   return self, nil
end
function memStream:close(  )

   self.txt:flush(  )
end
function memStream:flush(  )

   self.txt:flush(  )
end
function memStream._setmeta( obj )
  setmetatable( obj, { __index = memStream  } )
end


local TxtStream = {}
setmetatable( TxtStream, { ifList = {iStream,} } )
_moduleObj.TxtStream = TxtStream
function TxtStream._new( txt )
   local obj = {}
   TxtStream._setmeta( obj )
   if obj.__init then obj:__init( txt ); end
   return obj
end
function TxtStream:__init(txt) 
   self.txt = txt
   self.start = 1
   self.eof = false
   self.lineList = Str.getLineList( self.txt )
   self.lineNo = 1
end
function TxtStream:getSubstring( fromLineNo, toLineNo )

   local txt = ""
   local to = _lune.unwrapDefault( toLineNo, #self.lineList + 1)
   for index = fromLineNo, to - 1 do
      if index < 1 or index > #self.lineList then
         break
      end
      
      txt = string.format( "%s%s", txt, self.lineList[index])
   end
   
   return txt
end
function TxtStream:read( mode )

   if mode ~= '*l' then
      err( string.format( "not support -- %s", tostring( mode)) )
   end
   
   if self.lineNo > #self.lineList then
      return nil
   end
   
   self.lineNo = self.lineNo + 1
   local line = self.lineList[self.lineNo - 1]
   if Str.endsWith( line, "\n" ) then
      return line:sub( 1, #line - 1 )
   end
   
   return line
end
function TxtStream:close(  )

end
function TxtStream._setmeta( obj )
  setmetatable( obj, { __index = TxtStream  } )
end
function TxtStream:get_txt()
   return self.txt
end
function TxtStream:get_lineNo()
   return self.lineNo
end


local NullOStream = {}
setmetatable( NullOStream, { ifList = {oStream,} } )
_moduleObj.NullOStream = NullOStream
function NullOStream:write( val )

   return self, nil
end
function NullOStream:close(  )

end
function NullOStream:flush(  )

end
function NullOStream._setmeta( obj )
  setmetatable( obj, { __index = NullOStream  } )
end
function NullOStream._new(  )
   local obj = {}
   NullOStream._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function NullOStream:__init(  )

end


local SourceStream = {}
_moduleObj.SourceStream = SourceStream
function SourceStream._setmeta( obj )
  setmetatable( obj, { __index = SourceStream  } )
end
function SourceStream._new(  )
   local obj = {}
   SourceStream._setmeta( obj )
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
function SimpleSourceOStream._new( stream, headStream, stepIndent )
   local obj = {}
   SimpleSourceOStream._setmeta( obj )
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
   self.indentSpace = ""
end
function SimpleSourceOStream:get_indent(  )

   if #self.indentQueue > 0 then
      return self.indentQueue[#self.indentQueue]
   end
   
   return 0
end
function SimpleSourceOStream:writeRaw( txt )

   if self.needIndent then
      
      self.nowStream:write( self.indentSpace )
      self.needIndent = false
   end
   
   self.nowStream:write( txt )
end
function SimpleSourceOStream:write( txt )

   if not txt:find( "\n", 1, true ) then
      self:writeRaw( txt )
      return 
   end
   
   local stream = self.nowStream
   for __index, line in ipairs( Str.getLineList( txt ) ) do
      if self.needIndent then
         
         stream:write( self.indentSpace )
         self.needIndent = false
      end
      
      
      if #line > 0 and string.byte( line, #line ) == 10 then
         self.curLineNo = self.curLineNo + 1
      end
      
      stream:write( line )
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
   if indent > #SimpleSourceOStream.indentSpaceList then
      err( string.format( "overflow indent -- %d", indent) )
   end
   
   self.indentSpace = SimpleSourceOStream.indentSpaceList[indent + 1]
end
function SimpleSourceOStream:popIndent(  )

   if #self.indentQueue == 0 then
      err( "self.indentQueue == 0" )
   end
   
   table.remove( self.indentQueue )
   self.indentSpace = SimpleSourceOStream.indentSpaceList[self:get_indent() + 1]
end
function SimpleSourceOStream:switchToHeader(  )

   self.nowStream = self.headStream
end
function SimpleSourceOStream:returnToSource(  )

   self.nowStream = self.srcStream
end
function SimpleSourceOStream._setmeta( obj )
  setmetatable( obj, { __index = SimpleSourceOStream  } )
end
function SimpleSourceOStream:get_stepIndent()
   return self.stepIndent
end
do
   local list = {}
   local txt = ""
   for _1 = 1, 100 do
      table.insert( list, txt )
      txt = txt .. " "
   end
   
   SimpleSourceOStream.indentSpaceList = list
end


local function log( message )

   if debugFlag then
      errorLog( message )
   end
   
end
_moduleObj.log = log

local function printStackTrace(  )

   errorLog( Depend.getStackTrace(  ) )
end
_moduleObj.printStackTrace = printStackTrace
local function getReadyCode( depPath, tgtPath )
   local __func__ = '@lune.@base.@Util.getReadyCode'

   local tgtTime, depTime = Depend.getFileLastModifiedTime( tgtPath ), Depend.getFileLastModifiedTime( depPath )
   if  nil == tgtTime or  nil == depTime then
      local _tgtTime = tgtTime
      local _depTime = depTime
   
      
      return false
   end
   
   if tgtTime >= depTime then
      return true
   end
   
   Log.log( Log.Level.Warn, __func__, 402, function (  )
   
      return string.format( "not ready %g < %g : %s, %s", tgtTime, depTime, tgtPath, depPath)
   end )
   
   return false
end
_moduleObj.getReadyCode = getReadyCode

local function scriptPath2Module( path )

   if path:find( "^/" ) then
      err( "script must be relative-path -- " .. path )
   end
   
   local mod = path:gsub( "^./", "" ):gsub( "/", "." )
   return (string.gsub( mod, "%.lns$", "" ) )
end
_moduleObj.scriptPath2Module = scriptPath2Module

local function scriptPath2ModuleFromProjDir( path, projDir )

   if projDir ~= nil then
      local workpath
      
      if not projDir:find( "/$" ) then
         workpath = projDir .. "/"
      else
       
         workpath = projDir
      end
      
      path = path:gsub( "^" .. workpath, "" )
   end
   
   return scriptPath2Module( path )
end
_moduleObj.scriptPath2ModuleFromProjDir = scriptPath2ModuleFromProjDir

local function pathJoin( dir, path )

   if path:find( "^/" ) then
      return path
   end
   
   if dir:find( "/$" ) then
      return string.format( "%s%s", dir, path)
   end
   
   return string.format( "%s/%s", dir, path)
end
_moduleObj.pathJoin = pathJoin

local function parentPath( path )

   if path:find( "/$" ) then
      path = path:gsub( "/$", "" )
   end
   
   return (path:gsub( "/[^/]+$", "" ) )
end
_moduleObj.parentPath = parentPath

local function searchProjDir( dir )

   local work = dir
   while work ~= "" do
      if Depend.existFile( pathJoin( work, "lune.js" ) ) then
         return work
      end
      
      local parent = work:gsub( "/[^/]+$", "" )
      if parent == work then
         return nil
      end
      
      work = parent
   end
   
   return nil
end
_moduleObj.searchProjDir = searchProjDir



return _moduleObj
