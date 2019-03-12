--lune/base/Option.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Option'
if not _lune then
   _lune = {}
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

local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local LuaMod = _lune.loadModule( 'lune.base.LuaMod' )
local Ver = _lune.loadModule( 'lune.base.Ver' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Log = _lune.loadModule( 'lune.base.Log' )

local function getBuildCount(  )

   return 566
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

local CheckingUptodateMode = {}
_moduleObj.CheckingUptodateMode = CheckingUptodateMode
CheckingUptodateMode._val2NameMap = {}
function CheckingUptodateMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CheckingUptodateMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function CheckingUptodateMode._from( val )
   if CheckingUptodateMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
CheckingUptodateMode.__allList = {}
function CheckingUptodateMode.get__allList()
   return CheckingUptodateMode.__allList
end

CheckingUptodateMode.Force = 'force'
CheckingUptodateMode._val2NameMap['force'] = 'Force'
CheckingUptodateMode.__allList[1] = CheckingUptodateMode.Force
CheckingUptodateMode.Normal = 'none'
CheckingUptodateMode._val2NameMap['none'] = 'Normal'
CheckingUptodateMode.__allList[2] = CheckingUptodateMode.Normal
CheckingUptodateMode.Touch = 'touch'
CheckingUptodateMode._val2NameMap['touch'] = 'Touch'
CheckingUptodateMode.__allList[3] = CheckingUptodateMode.Touch

local TransCtrlInfo = {}
_moduleObj.TransCtrlInfo = TransCtrlInfo
function TransCtrlInfo.setmeta( obj )
  setmetatable( obj, { __index = TransCtrlInfo  } )
end
function TransCtrlInfo.new( checkingDefineAbbr, stopByWarning, uptodateMode )
   local obj = {}
   TransCtrlInfo.setmeta( obj )
   if obj.__init then
      obj:__init( checkingDefineAbbr, stopByWarning, uptodateMode )
   end        
   return obj 
end         
function TransCtrlInfo:__init( checkingDefineAbbr, stopByWarning, uptodateMode ) 

   self.checkingDefineAbbr = checkingDefineAbbr
   self.stopByWarning = stopByWarning
   self.uptodateMode = uptodateMode
end

function TransCtrlInfo.create_normal(  )

   return TransCtrlInfo.new(true, false, CheckingUptodateMode.Touch)
end

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
   self.useLuneModule = false
   self.updateOnLoad = false
   self.byteCompile = false
   self.stripDebugInfo = false
   self.targetLuaVer = LuaVer.curVer
   self.transCtrlInfo = TransCtrlInfo.create_normal(  )
end
function Option:openDepend(  )

   do
      local path = self.dependsPath
      if path ~= nil then
         return io.open( path, "w" )
      end
   end
   
   return nil
end
function Option.setmeta( obj )
  setmetatable( obj, { __index = Option  } )
end

local function outputLuneMod( dir )

   local path = "_lune.lua"
   do
      local _exp = dir
      if _exp ~= nil then
         if _exp ~= "" then
            path = _exp:gsub( "/$", "" ) .. "/" .. path
         end
         
      end
   end
   
   local fileObj = io.open( path, "w" )
   if  nil == fileObj then
      local _fileObj = fileObj
   
      return string.format( "failed to open -- %s", path)
   end
   
   for __index, kind in pairs( LuaMod.CodeKind.get__allList() ) do
      fileObj:write( LuaMod.getCode( kind ) )
   end
   
   fileObj:close(  )
   return nil
end
_moduleObj.outputLuneMod = outputLuneMod
local function analyze( argList )

   local function printUsage( code )
   
      print( [==[
usage:
  <type1> [-prof] [-r] src.lns mode [mode-option]
  <type2> -mklunemod dir
  <type3> --version

* type1
  - src.lns [common_op] ast
  - src.lns [common_op] comp [-i] module line column
  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] <lua|LUA>
  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] [--depends dependfile] <save|SAVE> output-dir
  - src.lns [common_op] exe

  -r: use 'require( "lune.base._lune" )'
  -ol: output lua version. ver = 51 or 52 or 53.
  -ob: output bytecompiled-code.
      -ob0 is without debug information. 
      -ob1 is with debug information. 
  --depends: output dependfile

  common_op:
    -u: update meta and lua on load.
    -Werror: error by warrning.
    --log <mode>: set log level.
         mode: fatal, error, warn, info, debug, trace
    --disable-checking-define-abbr: disable checking for ##.
    --uptodate <mode>: checking uptodate mode.
            mode: skip check.
            none: skip process when file is uptodate.
            touch: touch meta file when file is uptodate.  (default)

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
               print( string.format( "LuneScript: version %s.%d (%s)", Ver.version, getBuildCount(  ), _VERSION) )
               os.exit( 0 )
            elseif _switchExp == "-mklunemod" then
               local path = (#argList > index ) and argList[index + 1] or nil
               do
                  local mess = outputLuneMod( path )
                  if mess ~= nil then
                     Util.errorLog( mess )
                     os.exit( 1 )
                  end
               end
               
               os.exit( 0 )
            elseif _switchExp == "-r" then
               option.useLuneModule = true
            elseif _switchExp == "-u" then
               option.updateOnLoad = true
            elseif _switchExp == "-Werror" then
               option.transCtrlInfo.stopByWarning = true
            elseif _switchExp == "--disable-checking-define-abbr" then
               option.transCtrlInfo.checkingDefineAbbr = false
            elseif _switchExp == "--log" then
               if #argList > index then
                  local txt = argList[index + 1]
                  do
                     local level = Log.str2level( txt )
                     if level ~= nil then
                        Log.setLevel( level )
                     else
                        Util.errorLog( string.format( "illegal level -- %s", txt) )
                     end
                  end
                  
               end
               
               index = index + 1
            elseif _switchExp == "--depends" then
               if #argList > index then
                  option.dependsPath = argList[index + 1]
               end
               
               index = index + 1
            elseif _switchExp == "--uptodate" then
               if #argList > index then
                  do
                     local mode = CheckingUptodateMode._from( argList[index + 1] )
                     if mode ~= nil then
                        option.transCtrlInfo.uptodateMode = mode
                     else
                        Util.errorLog( "illegal mode -- " .. argList[index + 1] )
                     end
                  end
                  
               end
               
               index = index + 1
            elseif _switchExp == "-ol" then
               if #argList > index then
                  do
                     local _switchExp = argList[index + 1]
                     if _switchExp == "51" then
                        option.targetLuaVer = LuaVer.ver51
                     elseif _switchExp == "52" then
                        option.targetLuaVer = LuaVer.ver52
                     elseif _switchExp == "53" then
                        option.targetLuaVer = LuaVer.ver53
                     end
                  end
                  
               end
               
               index = index + 1
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
   
   if option.scriptPath == "" or option.mode == ModeKind.Unknown then
      printUsage( (#argList == 0 or argList[1] == "" ) and 0 or 1 )
   end
   
   if useStdInFlag and option.analyzeModule then
      Parser.StreamParser.setStdinStream( _lune.unwrap( option.analyzeModule) )
   end
   
   return option
end
_moduleObj.analyze = analyze
return _moduleObj
