--lune/base/Option.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Option'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

function _lune._fromList( obj, list, memInfoList )
   if type( list ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( list[ index ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s[%s]", memInfo.name, key) or memInfo.name
      end
      obj[ index ] = val
   end
   return true
end
function _lune._AlgeFrom( Alge, val )
   local work = Alge._name2Val[ val[ 1 ] ]
   if not work then
      return nil
   end
   if #work == 1 then
     return work
   end
   local paramList = {}
   local result, mess = _lune._fromList( paramList, val[ 2 ], work[ 2 ] )
   if not result then
      return nil, mess
   end
   return { work[ 1 ], paramList }
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

function _lune._toStem( val )
   return val
end
function _lune._toInt( val )
   if type( val ) == "number" then
      return math.floor( val )
   end
   return nil
end
function _lune._toReal( val )
   if type( val ) == "number" then
      return val
   end
   return nil
end
function _lune._toBool( val )
   if type( val ) == "boolean" then
      return val
   end
   return nil
end
function _lune._toStr( val )
   if type( val ) == "string" then
      return val
   end
   return nil
end
function _lune._toList( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toValInfo = toValInfoList[ 1 ]
      for index, mem in ipairs( val ) do
         local memval, mess = toValInfo.func( mem, toValInfo.child )
         if memval == nil and not toValInfo.nilable then
            if mess then
              return nil, string.format( "%d.%s", index, mess )
            end
            return nil, index
         end
         tbl[ index ] = memval
      end
      return tbl
   end
   return nil
end
function _lune._toMap( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toKeyInfo = toValInfoList[ 1 ]
      local toValInfo = toValInfoList[ 2 ]
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal, valSub = toValInfo.func( mem, toValInfo.child )
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
function _lune._fromMap( obj, map, memInfoList )
   if type( map ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( map[ memInfo.name ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s.%s", memInfo.name, key) or memInfo.name
      end
      obj[ memInfo.name ] = val
   end
   return true
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
local Types = _lune.loadModule( 'lune.base.Types' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Json = _lune.loadModule( 'lune.base.Json' )
local Util = _lune.loadModule( 'lune.base.Util' )
local LuaMod = _lune.loadModule( 'lune.base.LuaMod' )
local Ver = _lune.loadModule( 'lune.base.Ver' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Depend = _lune.loadModule( 'lune.base.Depend' )
local Log = _lune.loadModule( 'lune.base.Log' )
local Ast = _lune.loadModule( 'lune.base.Ast' )




local function getBuildCount(  )

   return 6252
end


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
    
ModeKind.__allList = {}
function ModeKind.get__allList()
   return ModeKind.__allList
end

ModeKind.Unknown = ''
ModeKind._val2NameMap[''] = 'Unknown'
ModeKind.__allList[1] = ModeKind.Unknown
ModeKind.Token = 'token'
ModeKind._val2NameMap['token'] = 'Token'
ModeKind.__allList[2] = ModeKind.Token
ModeKind.Ast = 'ast'
ModeKind._val2NameMap['ast'] = 'Ast'
ModeKind.__allList[3] = ModeKind.Ast
ModeKind.Diag = 'diag'
ModeKind._val2NameMap['diag'] = 'Diag'
ModeKind.__allList[4] = ModeKind.Diag
ModeKind.Complete = 'comp'
ModeKind._val2NameMap['comp'] = 'Complete'
ModeKind.__allList[5] = ModeKind.Complete
ModeKind.Lua = 'lua'
ModeKind._val2NameMap['lua'] = 'Lua'
ModeKind.__allList[6] = ModeKind.Lua
ModeKind.LuaMeta = 'LUA'
ModeKind._val2NameMap['LUA'] = 'LuaMeta'
ModeKind.__allList[7] = ModeKind.LuaMeta
ModeKind.Save = 'save'
ModeKind._val2NameMap['save'] = 'Save'
ModeKind.__allList[8] = ModeKind.Save
ModeKind.SaveMeta = 'SAVE'
ModeKind._val2NameMap['SAVE'] = 'SaveMeta'
ModeKind.__allList[9] = ModeKind.SaveMeta
ModeKind.Exec = 'exe'
ModeKind._val2NameMap['exe'] = 'Exec'
ModeKind.__allList[10] = ModeKind.Exec
ModeKind.Glue = 'glue'
ModeKind._val2NameMap['glue'] = 'Glue'
ModeKind.__allList[11] = ModeKind.Glue
ModeKind.BootC = 'bootC'
ModeKind._val2NameMap['bootC'] = 'BootC'
ModeKind.__allList[12] = ModeKind.BootC
ModeKind.Format = 'format'
ModeKind._val2NameMap['format'] = 'Format'
ModeKind.__allList[13] = ModeKind.Format
ModeKind.Builtin = 'builtin'
ModeKind._val2NameMap['builtin'] = 'Builtin'
ModeKind.__allList[14] = ModeKind.Builtin
ModeKind.Inquire = 'inq'
ModeKind._val2NameMap['inq'] = 'Inquire'
ModeKind.__allList[15] = ModeKind.Inquire


local function getRuntimeModule(  )

   return string.format( "lune.base.runtime%d", Ver.luaModVersion)
end
_moduleObj.getRuntimeModule = getRuntimeModule

local Option = {}
_moduleObj.Option = Option
function Option.new(  )
   local obj = {}
   Option.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Option:__init() 
   self.packageName = nil
   self.testing = false
   self.convTo = nil
   self.validProf = false
   self.mode = ModeKind.Unknown
   self.scriptPath = ""
   self.useLuneModule = nil
   self.updateOnLoad = false
   self.byteCompile = false
   self.stripDebugInfo = false
   self.targetLuaVer = LuaVer.getCurVer(  )
   self.transCtrlInfo = Types.TransCtrlInfo.create_normal(  )
   self.bootPath = nil
   self.useIpairs = false
   self.analyzeModule = nil
   self.outputDir = nil
   self.analyzePos = nil
   self.dependsPath = nil
end
function Option:openDepend( relPath )

   do
      local path = self.dependsPath
      if path ~= nil then
         local filePath
         
         if relPath ~= nil then
            if path:find( "/$" ) then
               filePath = string.format( "%s%s", path, relPath)
            else
             
               filePath = string.format( "%s/%s", path, relPath)
            end
            
         else
            filePath = path
         end
         
         return (io.open( filePath, "w" ) )
      end
   end
   
   do
      local path = self.dependsPath
      if path ~= nil then
         return (io.open( path, "w" ) )
      end
   end
   
   return nil
end
function Option.setmeta( obj )
  setmetatable( obj, { __index = Option  } )
end


local function outputLuneMod( path )

   local lune_path = "runtime.lua"
   if path ~= nil then
      if path ~= "" then
         lune_path = path
      end
      
   end
   
   local fileObj = io.open( lune_path, "w" )
   if  nil == fileObj then
      local _fileObj = fileObj
   
      return string.format( "failed to open -- %s", lune_path)
   end
   
   
   fileObj:write( [==[
--[[
MIT License

Copyright (c) 2018,2019 ifritJP

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
]]
]==] )
   
   for __index, kind in ipairs( LuaMod.CodeKind.get__allList() ) do
      fileObj:write( LuaMod.getCode( kind ) )
   end
   
   
   fileObj:close(  )
   return nil
end
_moduleObj.outputLuneMod = outputLuneMod

local function analyze( argList )
   local __func__ = '@lune.@base.@Option.analyze'

   
   local function printUsage( code )
   
      print( [==[
usage:
  <type1> [-prof] [-r] src.lns mode [mode-option]
  <type2> -mklunemod path
  <type3> --version

* type1
  - src.lns [common_op] ast
  - src.lns [common_op] comp [-i] module line column
  - src.lns [common_op] inq [-i] module line column
  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] <lua|LUA>
  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] [--depends dependfile] <save|SAVE> output-dir
  - src.lns [common_op] exe

  -r: use 'require( "lune.base.runtime" )'
  -ol: output lua version. ver = 51 or 52 or 53.
  -ob: output bytecompiled-code.
      -ob0 is without debug information.
      -ob1 is with debug information.
  -langC: transcompile to c-lang.
  -langGo: transcompile to golang.
  -oc: output path of the source code transcompiled to c-lang .
  --depends: output dependfile

  common_op:
    -u: update meta and lua on load.
    -Werror: error by warrning.
    --log <mode>: set log level.
         mode: fatal, error, warn, log, info, debug, trace
    --warning-shadowing: shadowing error convert to warning.
    --compat-comment: backward compatibility to process the comment.
    --disable-checking-define-abbr: disable checking for ##.
    --uptodate <mode>: checking uptodate mode.
            mode: skip check.
            none: skip process when file is uptodate.
            touch: touch meta file when file is uptodate.  (default)
    --use-ipairs: use ipairs for foreach with List value.
    --default-lazy: set lazy-loading at default.
    --valid-luaval: enable luaval when transcompie to lua.
    --package <name>: set the package name for the go-lang.

* type2
  dir: output directory.
]==] )
      os.exit( code )
   end
   
   local option = Option.new()
   local useStdInFlag = false
   local lineNo = nil
   local column = nil
   
   local index = 1
   
   do
      local file = io.open( "lune.js", "r" )
      if file ~= nil then
         local ProjInfo = {}
         setmetatable( ProjInfo, { ifList = {Mapping,} } )
         function ProjInfo.setmeta( obj )
  setmetatable( obj, { __index = ProjInfo  } )
end
         function ProjInfo.new( cmd_option )
   local obj = {}
   ProjInfo.setmeta( obj )
   if obj.__init then
      obj:__init( cmd_option )
   end
   return obj
end
function ProjInfo:__init( cmd_option )

            self.cmd_option = cmd_option
         end
         function ProjInfo:_toMap()
  return self
end
function ProjInfo._fromMap( val )
  local obj, mes = ProjInfo._fromMapSub( {}, val )
  if obj then
     ProjInfo.setmeta( obj )
  end
  return obj, mes
end
function ProjInfo._fromStem( val )
  return ProjInfo._fromMap( val )
end

         function ProjInfo._fromMapSub( obj, val )
            local memInfo = {}
            table.insert( memInfo, { name = "cmd_option", func = _lune._toList, nilable = false, child = { { func = _lune._toStr, nilable = false, child = {} } } } )
            local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end
         
         do
            local projInfo = ProjInfo._fromStem( (Json.fromStr( file:read( "*a" ) or "" ) ) )
            if projInfo ~= nil then
               local workArgList = {}
               for __index, arg in ipairs( argList ) do
                  table.insert( workArgList, arg )
               end
               
               for __index, arg in ipairs( projInfo.cmd_option ) do
                  table.insert( workArgList, arg )
               end
               
               argList = workArgList
            end
         end
         
      end
   end
   
   
   local function getNextOp(  )
   
      if #argList <= index then
         return nil
      end
      
      index = index + 1
      return argList[index]
   end
   
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
            elseif _switchExp == "--version" then
               print( string.format( "LuneScript: version %s (%d:%s) [%s]", Ver.version, getBuildCount(  ), Depend.getLuaVersion(  ), Ver.metaVersion) )
               os.exit( 0 )
            elseif _switchExp == "--builtin" then
               do
                  local __sorted = {}
                  local __map = Ast.getBuiltInTypeIdMap(  )
                  for __key in pairs( __map ) do
                     table.insert( __sorted, __key )
                  end
                  table.sort( __sorted )
                  for __index, typeId in ipairs( __sorted ) do
                     local typeInfo = __map[ typeId ]
                     do
                        print( typeId, typeInfo:getTxt(  ) )
                     end
                  end
               end
               
               os.exit( 0 )
            elseif _switchExp == "-mklunemod" then
               local path = getNextOp(  )
               do
                  local mess = outputLuneMod( path )
                  if mess ~= nil then
                     Util.errorLog( mess )
                     os.exit( 1 )
                  end
               end
               
               os.exit( 0 )
            elseif _switchExp == "--mkbuiltin" then
               local path = getNextOp(  )
               if  nil == path then
                  local _path = path
               
                  path = "."
               end
               
               option.scriptPath = path .. "/lns_builtin.lns"
               option.mode = ModeKind.Builtin
            elseif _switchExp == "-r" then
               option.useLuneModule = getRuntimeModule(  )
            elseif _switchExp == "--runtime" then
               option.useLuneModule = getNextOp(  )
            elseif _switchExp == "-oc" then
               option.bootPath = getNextOp(  )
            elseif _switchExp == "-u" then
               option.updateOnLoad = true
            elseif _switchExp == "-Werror" then
               option.transCtrlInfo.stopByWarning = true
            elseif _switchExp == "--disable-checking-define-abbr" then
               option.transCtrlInfo.checkingDefineAbbr = false
            elseif _switchExp == "--compat-comment" then
               option.transCtrlInfo.compatComment = true
            elseif _switchExp == "--warning-shadowing" then
               option.transCtrlInfo.warningShadowing = true
            elseif _switchExp == "--valid-luaval" then
               option.transCtrlInfo.validLuaval = true
            elseif _switchExp == "--default-lazy" then
               option.transCtrlInfo.defaultLazy = true
            elseif _switchExp == "--package" then
               option.packageName = getNextOp(  )
            elseif _switchExp == "--log" then
               do
                  local txt = getNextOp(  )
                  if txt ~= nil then
                     do
                        local level = Log.str2level( txt )
                        if level ~= nil then
                           Log.setLevel( level )
                        else
                           Util.errorLog( string.format( "illegal level -- %s", txt) )
                        end
                     end
                     
                  end
               end
               
            elseif _switchExp == "--testing" then
               option.testing = true
            elseif _switchExp == "--depends" then
               option.dependsPath = getNextOp(  )
            elseif _switchExp == "--use-ipairs" then
               option.useIpairs = true
            elseif _switchExp == "--uptodate" then
               do
                  local txt = getNextOp(  )
                  if txt ~= nil then
                     do
                        local mode = Types.CheckingUptodateMode._from( txt )
                        if mode ~= nil then
                           option.transCtrlInfo.uptodateMode = mode
                        else
                           Util.errorLog( "illegal mode -- " .. txt )
                        end
                     end
                     
                  end
               end
               
            elseif _switchExp == "-langC" then
               option.convTo = Types.Conv.C
               option.transCtrlInfo.validLuaval = true
            elseif _switchExp == "-langGo" then
               option.convTo = Types.Conv.Go
               option.transCtrlInfo.validLuaval = true
            elseif _switchExp == "-ol" then
               do
                  local txt = getNextOp(  )
                  if txt ~= nil then
                     do
                        local _switchExp = txt
                        if _switchExp == "51" then
                           option.targetLuaVer = LuaVer.ver51
                        elseif _switchExp == "52" then
                           option.targetLuaVer = LuaVer.ver52
                        elseif _switchExp == "53" then
                           option.targetLuaVer = LuaVer.ver53
                        end
                     end
                     
                  end
               end
               
            elseif _switchExp == "-ob0" or _switchExp == "-ob1" then
               option.byteCompile = true
               if arg == "-ob0" then
                  option.stripDebugInfo = true
               end
               
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
               end
            end
            
         else
          
            do
               local _switchExp = (option.mode )
               if _switchExp == ModeKind.Complete or _switchExp == ModeKind.Inquire then
                  if not option.analyzeModule then
                     option.analyzeModule = arg
                  elseif not lineNo then
                     lineNo = math.floor((_lune.unwrapDefault( tonumber( arg ), 0) ))
                  elseif not column then
                     column = math.floor((_lune.unwrapDefault( tonumber( arg ), 0) ))
                     option.analyzePos = Parser.Position.new(_lune.unwrap( lineNo), _lune.unwrap( column), Util.scriptPath2Module( option.scriptPath ))
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
   
   
   if option.mode ~= ModeKind.Builtin then
      if option.scriptPath == "" or option.mode == ModeKind.Unknown then
         printUsage( (#argList == 0 or argList[1] == "" ) and 0 or 1 )
      end
      
   end
   
   
   if useStdInFlag then
      if option.analyzeModule then
         Parser.StreamParser.setStdinStream( _lune.unwrap( option.analyzeModule) )
      else
       
         if option.scriptPath ~= "" then
            Parser.StreamParser.setStdinStream( Util.scriptPath2Module( option.scriptPath ) )
         end
         
      end
      
   end
   
   
   Log.log( Log.Level.Log, __func__, 485, function (  )
   
      return string.format( "mode is '%s'", ModeKind:_getTxt( option.mode)
      )
   end )
   
   
   return option
end
_moduleObj.analyze = analyze

local function createDefaultOption( path )

   local option = Option.new()
   option.scriptPath = path
   option.useLuneModule = getRuntimeModule(  )
   option.useIpairs = true
   return option
end
_moduleObj.createDefaultOption = createDefaultOption



return _moduleObj
