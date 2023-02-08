--lune/base/Option.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Option'
local _lune = {}
if _lune8 then
   _lune = _lune8
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

if not _lune8 then
   _lune8 = _lune
end


local Types = _lune.loadModule( 'lune.base.Types' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local AsyncParser = _lune.loadModule( 'lune.base.AsyncParser' )
local Json = _lune.loadModule( 'lune.base.Json' )
local Util = _lune.loadModule( 'lune.base.Util' )
local LuaMod = _lune.loadModule( 'lune.base.LuaMod' )
local Ver = _lune.loadModule( 'lune.base.Ver' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Depend = _lune.loadModule( 'lune.base.Depend' )
local Log = _lune.loadModule( 'lune.base.Log' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Builtin = _lune.loadModule( 'lune.base.Builtin' )



local function getBuildCount(  )

   return 13330
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
ModeKind.MkMain = 'mkmain'
ModeKind._val2NameMap['mkmain'] = 'MkMain'
ModeKind.__allList[16] = ModeKind.MkMain
ModeKind.Shebang = 'shebang'
ModeKind._val2NameMap['shebang'] = 'Shebang'
ModeKind.__allList[17] = ModeKind.Shebang
ModeKind.BuildAst = 'buildAst'
ModeKind._val2NameMap['buildAst'] = 'BuildAst'
ModeKind.__allList[18] = ModeKind.BuildAst
ModeKind.Indexer = 'indexer'
ModeKind._val2NameMap['indexer'] = 'Indexer'
ModeKind.__allList[19] = ModeKind.Indexer
ModeKind.GoMod = 'gomod'
ModeKind._val2NameMap['gomod'] = 'GoMod'
ModeKind.__allList[20] = ModeKind.GoMod


local function getRuntimeModule(  )

   return string.format( "lune.base.runtime%d", Ver.luaModVersion)
end
_moduleObj.getRuntimeModule = getRuntimeModule

local Int2strMode = {}
_moduleObj.Int2strMode = Int2strMode
Int2strMode._val2NameMap = {}
function Int2strMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Int2strMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function Int2strMode._from( val )
   if Int2strMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
Int2strMode.__allList = {}
function Int2strMode.get__allList()
   return Int2strMode.__allList
end

Int2strMode.Int2strModeDepend = 0
Int2strMode._val2NameMap[0] = 'Int2strModeDepend'
Int2strMode.__allList[1] = Int2strMode.Int2strModeDepend
Int2strMode.Int2strModeNeed0 = 1
Int2strMode._val2NameMap[1] = 'Int2strModeNeed0'
Int2strMode.__allList[2] = Int2strMode.Int2strModeNeed0
Int2strMode.Int2strModeUnneed0 = 2
Int2strMode._val2NameMap[2] = 'Int2strModeUnneed0'
Int2strMode.__allList[3] = Int2strMode.Int2strModeUnneed0

local RuntimeOpt = {}
_moduleObj.RuntimeOpt = RuntimeOpt
function RuntimeOpt._new(  )
   local obj = {}
   RuntimeOpt._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function RuntimeOpt:__init() 
   self.int2strMode = Int2strMode.Int2strModeDepend
end
function RuntimeOpt._setmeta( obj )
  setmetatable( obj, { __index = RuntimeOpt  } )
end
function RuntimeOpt:get_int2strMode()
   return self.int2strMode
end


local Option = {}
_moduleObj.Option = Option
function Option._new(  )
   local obj = {}
   Option._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Option:__init() 
   self.sortGenerateCode = true
   self.dumpDebugAst = false
   self.legacyNewName = false
   self.stdinFile = nil
   self.validPostBuild = true
   self.enableRunner = true
   self.addEnvArg = true
   self.projDir = nil
   self.runtimeOpt = RuntimeOpt._new()
   self.shebangArgList = {}
   self.mainModule = ""
   self.appName = nil
   self.packageName = nil
   self.testing = false
   self.convTo = nil
   self.noLua = false
   self.validProf = false
   self.mode = ModeKind.Unknown
   self.scriptPath = ""
   self.batchList = {}
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
   self.convGoRunnerNum = false and 4 or 0
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
function Option._setmeta( obj )
  setmetatable( obj, { __index = Option  } )
end
function Option:get_runtimeOpt()
   return self.runtimeOpt
end
function Option:get_projDir()
   return self.projDir
end
function Option:get_addEnvArg()
   return self.addEnvArg
end
function Option:get_enableRunner()
   return self.enableRunner
end
function Option:get_validPostBuild()
   return self.validPostBuild
end
function Option:get_stdinFile()
   return self.stdinFile
end
function Option:set_stdinFile( stdinFile )
   self.stdinFile = stdinFile
end
function Option:get_legacyNewName()
   return self.legacyNewName
end
function Option:get_sortGenerateCode()
   return self.sortGenerateCode
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
   
      Util.println( [==[
usage:
  <type1> [-prof] [-r] src.lns mode [mode-option]
  <type2> -mklunemod path
  <type3> -shebang path
  <type4> --version

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
  -langPython: transcompile to python.
  -noLua: no transcompile to lua.
  -oc: output path of the source code transcompiled to c-lang .
  --depends: output dependfile
  --int2str mode: mode of int to str.
     - depend: depends the lua version.
     - need0: with '.0'.
     - unneed0: without '.0'.

  common_op:
    --testing: enable test.
    --projDir <dir>: set the project dir.
    -u: update meta and lua on load.
    -Werror: error by warrning.
    --log <mode>: set log level.
         mode: fatal, error, warn, log, info, debug, trace
    --warning-shadowing: shadowing error convert to warning.
    --compat-comment: backward compatibility to process the comment.
    --disable-checking-define-abbr: disable checking for ##.
    --uptodate <mode>: checking uptodate mode.
            force: skip check for target lns file.
            forceAll: skip check for all.
            none: skip process when file is uptodate.
            touch: touch meta file when file is uptodate.  (default)
    --use-ipairs: use ipairs for foreach with List value.
    --default-lazy: set lazy-loading at default.
    --valid-luaval: enable luaval when transcompie to lua.
    --package <name>: set the package name for the go-lang.
    --app <name>: set the application name for the go-lang.
    --debug-dump-ast: dump ast for debuging.

    compati_op:
      --legacyNewName: use the legacy new method name for lua.



* type2
  path: output file path.
]==] )
      os.exit( code )
   end
   
   local option = Option._new()
   local useStdInFlag = false
   local lineNo = nil
   local column = nil
   
   local index = 1
   
   do
      local file = io.open( "lune.js", "r" )
      if file ~= nil then
         
         local ProjInfo = {}
         setmetatable( ProjInfo, { ifList = {Mapping,} } )
         function ProjInfo._setmeta( obj )
  setmetatable( obj, { __index = ProjInfo  } )
end
         function ProjInfo._new( cmd_option )
   local obj = {}
   ProjInfo._setmeta( obj )
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
     ProjInfo._setmeta( obj )
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
               for __index, arg in ipairs( projInfo.cmd_option ) do
                  table.insert( workArgList, arg )
               end
               
               for __index, arg in ipairs( argList ) do
                  table.insert( workArgList, arg )
               end
               
               argList = workArgList
            end
         end
         
         file:close(  )
      end
   end
   
   
   local function getNextOp(  )
   
      if #argList <= index then
         return nil
      end
      
      index = index + 1
      return argList[index]
   end
   local function getNextOpNonNil(  )
   
      do
         local nextOp = getNextOp(  )
         if nextOp ~= nil then
            return nextOp
         end
      end
      
      printUsage( 1 )
   end
   
   local function getNextOpInt(  )
   
      do
         local num = tonumber( getNextOpNonNil(  ) )
         if num ~= nil then
            return math.floor(num)
         end
      end
      
      printUsage( 1 )
   end
   
   Util.setDebugFlag( false )
   
   local uptodateOpt = nil
   while #argList >= index do
      local arg = argList[index]
      
      if arg:find( "^-" ) then
         if option.mode ~= ModeKind.Shebang then
            do
               local _switchExp = (arg )
               if _switchExp == "-i" then
                  useStdInFlag = true
               elseif _switchExp == "--parserPipeSize" then
                  AsyncParser.setDefaultPipeSize( getNextOpInt(  ) )
               elseif _switchExp == "-prof" then
                  option.validProf = true
               elseif _switchExp == "--noEnvArg" then
                  option.addEnvArg = false
               elseif _switchExp == "--noUseWaiter" then
                  option.transCtrlInfo.useWaiter = false
                  option.sortGenerateCode = false
               elseif _switchExp == "--debug-dump-ast" then
                  option.dumpDebugAst = true
               elseif _switchExp == "--unsortGenerateCode" then
                  option.sortGenerateCode = false
               elseif _switchExp == "--disableMultiPhaseAst" then
                  option.transCtrlInfo.validMultiPhaseTransUnit = false
               elseif _switchExp == "--disableMultiThreadAst" then
                  option.transCtrlInfo.threadPerUnitThread = 0
               elseif _switchExp == "--threadPerUnitThread" then
                  option.transCtrlInfo.threadPerUnitThread = getNextOpInt(  )
                  if option.transCtrlInfo.threadPerUnitThread < 0 then
                     printUsage( 1 )
                  end
                  
               elseif _switchExp == "--convGoRunnerNum" then
                  option.convGoRunnerNum = getNextOpInt(  )
                  if option.convGoRunnerNum < 0 then
                     printUsage( 1 )
                  end
                  
               elseif _switchExp == "--macroAsyncParseStmtLen" then
                  option.transCtrlInfo.macroAsyncParseStmtLen = getNextOpInt(  )
                  if option.transCtrlInfo.macroAsyncParseStmtLen < 0 then
                     printUsage( 1 )
                  end
                  
               elseif _switchExp == "--legacyNewName" then
                  option.legacyNewName = true
               elseif _switchExp == "--enableMacroAsync" then
                  option.transCtrlInfo.validMacroAsync = true
               elseif _switchExp == "--disableRunner" then
                  option.enableRunner = false
               elseif _switchExp == "--disablePostBuild" then
                  option.validPostBuild = false
               elseif _switchExp == "--enableAsyncCtl" then
                  option.transCtrlInfo.validAsyncCtrl = true
               elseif _switchExp == "--defaultAsync" then
                  option.transCtrlInfo.defaultAsync = true
               elseif _switchExp == "--limitThread" then
                  local nextArg = getNextOp(  )
                  if  nil == nextArg then
                     local _nextArg = nextArg
                  
                     printUsage( 1 )
                  end
                  
                  local num = tonumber( nextArg )
                  if  nil == num then
                     local _num = num
                  
                     printUsage( 1 )
                  end
                  
                  Depend.setRuntimeThreadLimit( math.floor(num) )
               elseif _switchExp == "--nodebug" then
                  Util.setDebugFlag( false )
               elseif _switchExp == "--debug" then
                  Util.setDebugFlag( true )
               elseif _switchExp == "-shebang" then
                  option.mode = ModeKind.Shebang
               elseif _switchExp == "--version" then
                  Util.println( string.format( "LuneScript: version %s (%d:Lua%s) [%s]", Ver.version, getBuildCount(  ), Depend.getLuaVersion(  ), Ver.metaVersion) )
                  os.exit( 0 )
               elseif _switchExp == "--projDir" then
                  option.projDir = getNextOp(  )
               elseif _switchExp == "--builtin" then
                  local builtin = Builtin.Builtin._new(option.targetLuaVer, option.transCtrlInfo)
                  local builtinFunc = builtin:registBuiltInScope(  )
                  
                  do
                     local __sorted = {}
                     local __map = Ast.getBuiltInTypeIdMap(  )
                     for __key in pairs( __map ) do
                        table.insert( __sorted, __key )
                     end
                     table.sort( __sorted )
                     for __index, typeId in ipairs( __sorted ) do
                        local builtinTypeInfo = __map[ typeId ]
                        do
                           Util.println( typeId, builtinTypeInfo:get_typeInfo():getTxt(  ) )
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
               elseif _switchExp == "--disable-checking-mutable" then
                  option.transCtrlInfo.validCheckingMutable = false
               elseif _switchExp == "--legacy-mutable-control" then
                  option.transCtrlInfo.legacyMutableControl = true
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
               elseif _switchExp == "--int2str" then
                  local opt = getNextOp(  )
                  do
                     local _switchExp = opt
                     if _switchExp == "depend" then
                        option:get_runtimeOpt().int2strMode = Int2strMode.Int2strModeDepend
                     elseif _switchExp == "need0" then
                        option:get_runtimeOpt().int2strMode = Int2strMode.Int2strModeNeed0
                     elseif _switchExp == "unneed0" then
                        option:get_runtimeOpt().int2strMode = Int2strMode.Int2strModeUnneed0
                     else 
                        
                           Util.errorLog( string.format( "unknown mode -- %s", opt) )
                           os.exit( 1 )
                     end
                  end
                  
               elseif _switchExp == "--app" then
                  do
                     local _exp = getNextOp(  )
                     if _exp ~= nil then
                        option.appName = _exp
                     end
                  end
                  
               elseif _switchExp == "--main" then
                  do
                     local _exp = getNextOp(  )
                     if _exp ~= nil then
                        option.mainModule = _exp
                     end
                  end
                  
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
                  option.transCtrlInfo.testing = true
               elseif _switchExp == "--enableTestBlock" then
                  option.transCtrlInfo.testing = true
               elseif _switchExp == "--depends" then
                  option.dependsPath = getNextOp(  )
               elseif _switchExp == "--use-ipairs" then
                  option.useIpairs = true
               elseif _switchExp == "--uptodate" then
                  uptodateOpt = getNextOp(  )
               elseif _switchExp == "-noLua" then
                  option.noLua = true
               elseif _switchExp == "-langC" then
                  option.convTo = Types.Lang.C
                  option.transCtrlInfo.validLuaval = true
               elseif _switchExp == "-langGo" then
                  option.convTo = Types.Lang.Go
                  option.transCtrlInfo.validLuaval = true
                  option.transCtrlInfo.validAsyncCtrl = true
               elseif _switchExp == "-langPython" then
                  option.convTo = Types.Lang.Python
                  option.transCtrlInfo.validLuaval = false
                  option.transCtrlInfo.validAsyncCtrl = false
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
                  
                     Util.log( string.format( "unknown option -- '%s'", arg) )
                     os.exit( 1 )
               end
            end
            
         else
          
            if #option.shebangArgList == 0 then
               table.insert( option.shebangArgList, option.scriptPath )
            end
            
            table.insert( option.shebangArgList, arg )
         end
         
      else
       
         if option.scriptPath == "" then
            option.scriptPath = arg
            if option.mode == ModeKind.Shebang then
               table.insert( option.shebangArgList, option.scriptPath )
            end
            
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
                     option.analyzePos = Parser.Position._new(_lune.unwrap( lineNo), _lune.unwrap( column), Util.scriptPath2Module( option.scriptPath ))
                  end
                  
               elseif _switchExp == ModeKind.Save or _switchExp == ModeKind.SaveMeta or _switchExp == ModeKind.Glue then
                  option.outputDir = arg
               elseif _switchExp == ModeKind.Shebang then
                  if #option.shebangArgList == 0 then
                     table.insert( option.shebangArgList, option.scriptPath )
                  end
                  
                  table.insert( option.shebangArgList, arg )
               else 
                  
                     option.outputDir = arg
               end
            end
            
         end
         
      end
      
      
      index = index + 1
   end
   
   
   if uptodateOpt ~= nil then
      do
         local _switchExp = uptodateOpt
         if _switchExp == "force" then
            option.transCtrlInfo.uptodateMode = _lune.newAlge( Types.CheckingUptodateMode.Force1, {Util.scriptPath2Module( option.scriptPath )})
         elseif _switchExp == "forceAll" then
            option.transCtrlInfo.uptodateMode = _lune.newAlge( Types.CheckingUptodateMode.ForceAll)
         elseif _switchExp == "normal" then
            option.transCtrlInfo.uptodateMode = _lune.newAlge( Types.CheckingUptodateMode.Normal)
         elseif _switchExp == "touch" then
            option.transCtrlInfo.uptodateMode = _lune.newAlge( Types.CheckingUptodateMode.Touch)
         else 
            
               Util.errorLog( "illegal mode -- " .. uptodateOpt )
         end
      end
      
   end
   
   
   if option.mode ~= ModeKind.Builtin then
      if option.scriptPath == "" or option.mode == ModeKind.Unknown then
         printUsage( (#argList == 0 or argList[1] == "" ) and 0 or 1 )
      end
      
   end
   
   
   if useStdInFlag then
      local code = io.stdin:read( "*a" )
      if  nil == code then
         local _code = code
      
         Util.err( "read error from stdin." )
      end
      
      if option.analyzeModule then
         option.stdinFile = Types.StdinFile._new(_lune.unwrap( option.analyzeModule), code)
      else
       
         if option.scriptPath ~= "" then
            option.stdinFile = Types.StdinFile._new(Util.scriptPath2Module( option.scriptPath ), code)
         end
         
      end
      
   end
   
   
   if option.scriptPath == "@-" then
      while true do
         local line = io.stdin:read( "*l" )
         if  nil == line then
            local _line = line
         
            break
         end
         
         if #line > 0 then
            table.insert( option.batchList, line )
         end
         
      end
      
   end
   
   
   Log.log( Log.Level.Log, __func__, 770, function (  )
   
      return string.format( "mode is '%s'", ModeKind:_getTxt( option.mode)
      )
   end )
   
   
   return option
end
_moduleObj.analyze = analyze

local function createDefaultOption( pathList, projDir )

   local option = Option._new()
   if #pathList == 1 then
      option.scriptPath = pathList[1]
   else
    
      option.scriptPath = "@-"
      for __index, path in ipairs( pathList ) do
         table.insert( option.batchList, path )
      end
      
   end
   
   
   option.useLuneModule = getRuntimeModule(  )
   option.useIpairs = true
   if projDir ~= nil then
      if projDir ~= "/" then
         if not projDir:find( "/$" ) then
            option.projDir = projDir .. "/"
         else
          
            option.projDir = projDir
         end
         
      end
      
   end
   
   return option
end
_moduleObj.createDefaultOption = createDefaultOption



return _moduleObj
