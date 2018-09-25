--lune/base/front.lns
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

local frontInterface = require( 'lune.base.frontInterface' )

local Parser = require( 'lune.base.Parser' )

local convLua = require( 'lune.base.convLua' )

local TransUnit = require( 'lune.base.TransUnit' )

local Util = require( 'lune.base.Util' )

local Option = require( 'lune.base.Option' )

local dumpNode = require( 'lune.base.dumpNode' )

local glueFilter = require( 'lune.base.glueFilter' )

local Depend = require( 'lune.base.Depend' )

function _luneGetLocal( varName )
  local index = 1
  
  while true do
    local name, val = debug.getlocal( 3, index )
    
    if name == varName then
      return val
    end
    if not name then
      break
    end
    index = index + 1
  end
  error( "not found -- " .. varName )
end

function _luneSym2Str( val )
  do
    local _exp = val
    if _exp ~= nil then
    
        if type( _exp ) ~= "table" then
          return string.format( "%s", _exp )
        end
        local txt = ""
        
        for __index, item in pairs( _exp ) do
          txt = txt .. item
        end
        return txt
      end
  end
  
  return nil
end

local Front = {}
-- none
-- none
-- none
-- none
function Front.setmeta( obj )
  setmetatable( obj, { __index = Front  } )
end
function Front.new( option, loadedMap, loadedMetaMap, convertedMap )
  local obj = {}
  Front.setmeta( obj )
  if obj.__init then
    obj:__init( option, loadedMap, loadedMetaMap, convertedMap )
  end        
  return obj 
end         
function Front:__init( option, loadedMap, loadedMetaMap, convertedMap ) 

self.option = option
  self.loadedMap = loadedMap
  self.loadedMetaMap = loadedMetaMap
  self.convertedMap = convertedMap
  end
do
  end

function Front:error( message )
  Util.errorLog( message )
  Util.printStackTrace(  )
  os.exit( 1 )
end

function Front:loadLua( path )
  local chunk, err = loadfile( path )
  
  do
    local _exp = err
    if _exp ~= nil then
    
        Util.errorLog( _exp )
      end
  end
  
  do
    local _exp = chunk
    if _exp ~= nil then
    
        return _exp(  )
      end
  end
  
  error( "failed to error" )
end

local function createPaser( path, mod )
  local parser = Parser.StreamParser.create( path, false, mod )
  
  do
    local _exp = parser
    if _exp ~= nil then
    
        return _exp
      end
  end
  
  error( "failed to open " .. path )
end

function Front:createAst( path, mod, analyzeModule, analyzeMode, pos )
  local transUnit = TransUnit.TransUnit.new(convLua.MacroEvalImp.new(self.option.mode), analyzeModule, analyzeMode, pos)
  
  return transUnit:createAST( createPaser( path, mod ), false, mod )
end

local function convert( ast, streamName, stream, metaStream, convMode, inMacro )
  local conv = convLua.createFilter( streamName, stream, metaStream, convMode, inMacro, ast:get_moduleTypeInfo(), ast:get_moduleSymbolKind() )
  
  ast:get_node():processFilter( conv, nil, 0 )
end

local function loadFromTxt( txt )
  local chunk, err = load( txt )
  
  do
    local _exp = err
    if _exp ~= nil then
    
        print( _exp )
      end
  end
  
  do
    local _exp = chunk
    if _exp ~= nil then
    
        return _exp(  )
      end
  end
  
  error( "failed to error" )
end

function Front:loadFile( path, mod, onlyMeta )
  local ast = self:createAst( path, mod, nil, TransUnit.AnalyzeMode.Compile )
  
  local stream = Util.memStream.new()
  
  local metaStream = Util.memStream.new()
  
  convert( ast, path, stream, metaStream, convLua.ConvMode.Exec, false )
  local meta = loadFromTxt( metaStream:get_txt() )
  
  if onlyMeta then
    return meta, stream:get_txt()
  end
  return meta, loadFromTxt( stream:get_txt() )
end

function Front:searchModule( mod )
  local lnsSearchPath = package.path
  
  lnsSearchPath = string.gsub( lnsSearchPath, "%.lua", ".lns" )
  return package.searchpath( mod, lnsSearchPath )
end

function Front:searchLuaFile( moduleFullName, addSearchPath )
  local luaSearchPath = package.path
  
  do
    local _exp = addSearchPath
    if _exp ~= nil then
    
        luaSearchPath = string.format( "%s/?.lua;%s", addSearchPath, package.path )
      end
  end
  
  return package.searchpath( moduleFullName, luaSearchPath )
end

function Front:checkUptodateMeta( metaPath, addSearchPath )
  local meta = self:loadLua( metaPath )
  
  for moduleFullName, dependInfo in pairs( meta._dependModuleMap ) do
    do
      local moduleLuaPath = self:searchLuaFile( moduleFullName, addSearchPath )
      if moduleLuaPath ~= nil then
      
          local moduleLnsPath = moduleLuaPath:gsub( "%.lua$", ".lns" )
          
          if not Util.getReadyCode( moduleLnsPath, metaPath ) then
            return nil
          end
        end
    end
    
  end
  return meta
end

function Front:loadModule( mod )
  if self.loadedMap[mod] == nil then
    do
      local _exp = self.convertedMap[mod]
      if _exp ~= nil then
      
          if not self.loadedMetaMap[mod] then
            error( string.format( "nothing meta -- %s", mod) )
          end
          local info = {}
          
          info['mod'] = loadFromTxt( _exp )
          info['meta'] = self.loadedMetaMap[mod]
          self.loadedMap[mod] = info
        else
      
          do
            local lnsPath = self:searchModule( mod )
            if lnsPath ~= nil then
            
                local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
                
                do
                  local _exp = self.option.outputDir
                  if _exp ~= nil then
                  
                      luaPath = self:searchLuaFile( mod, _exp )
                    end
                end
                
                local loadVal = nil
                
                do
                  local _exp = luaPath
                  if _exp ~= nil then
                  
                      if Util.getReadyCode( lnsPath, _exp ) then
                        local metaPath = string.gsub( _exp, "%.lua$", ".meta" )
                        
                        if Util.getReadyCode( lnsPath, metaPath ) then
                          loadVal = self:loadLua( _exp )
                          local meta = self:checkUptodateMeta( metaPath, self.option.outputDir )
                          
                          if meta then
                            local info = {}
                            
                            info['mod'] = loadVal
                            info['meta'] = meta
                            self.loadedMap[mod] = info
                          end
                        end
                      end
                    end
                end
                
                if loadVal == nil then
                  local meta, workVal = self:loadFile( lnsPath, mod, false )
                  
                  local info = {}
                  
                  info['mod'] = workVal
                  info['meta'] = meta
                  self.loadedMap[mod] = info
                end
              end
          end
          
        end
    end
    
  end
  do
    local _exp = self.loadedMap[mod]
    if _exp ~= nil then
    
        return _lune.unwrap( _exp['mod']), _lune.unwrap( _exp['meta'])
      end
  end
  
  error( string.format( "load error, %s", mod) )
end

function Front:loadMeta( mod )
  if self.loadedMetaMap[mod] == nil then
    do
      local _exp = self.loadedMap[mod]
      if _exp ~= nil then
      
          self.loadedMetaMap[mod] = _exp['meta']
        else
      
          do
            local lnsPath = self:searchModule( mod )
            if lnsPath ~= nil then
            
                local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
                
                do
                  local _exp = self.option.outputDir
                  if _exp ~= nil then
                  
                      luaPath = self:searchLuaFile( mod, _exp )
                    end
                end
                
                local meta = nil
                
                do
                  local _exp = luaPath
                  if _exp ~= nil then
                  
                      if Util.getReadyCode( lnsPath, _exp ) then
                        local metaPath = string.gsub( _exp, "%.lua$", ".meta" )
                        
                        if Util.getReadyCode( lnsPath, metaPath ) then
                          meta = self:checkUptodateMeta( metaPath, self.option.outputDir )
                          if meta then
                            self.loadedMetaMap[mod] = meta
                          end
                        end
                      end
                    end
                end
                
                if meta == nil then
                  local metawork, luaTxt = self:loadFile( lnsPath, mod, true )
                  
                  self.loadedMetaMap[mod] = metawork
                  self.convertedMap[mod] = luaTxt
                end
              end
          end
          
        end
    end
    
  end
  do
    local _exp = self.loadedMetaMap[mod]
    if _exp ~= nil then
    
        return _lune.unwrap( _exp)
      end
  end
  
  error( string.format( "load meta error, %s", mod) )
end

function Front:exec(  )
  local mod = string.gsub( self.option.scriptPath, "/", "." )
  
  mod = string.gsub( mod, "%.lns$", "" )
  do
    local _switchExp = self.option.mode
    if _switchExp == Option.ModeKind.Token then
      local parser = createPaser( self.option.scriptPath, mod )
      
      while true do
        local token = parser:getToken(  )
        
            if  nil == token then
              local _token = token
              
              break
            end
          
        print( token.kind, token.pos.lineNo, token.pos.column, token.txt )
      end
    elseif _switchExp == Option.ModeKind.Ast then
      Util.profile( self.option.validProf, function (  )
        local ast = self:createAst( self.option.scriptPath, mod, nil, TransUnit.AnalyzeMode.Compile )
        
        ast:get_node():processFilter( dumpNode.dumpFilter.new(), "", 0 )
      end
      , self.option.scriptPath .. ".profi" )
    elseif _switchExp == Option.ModeKind.Diag then
      Util.setErrorCode( 0 )
      self:createAst( self.option.scriptPath, mod, nil, TransUnit.AnalyzeMode.Diag )
    elseif _switchExp == Option.ModeKind.Complete then
      self:createAst( self.option.scriptPath, mod, self.option.analyzeModule, TransUnit.AnalyzeMode.Complete, self.option.analyzePos )
    elseif _switchExp == Option.ModeKind.Glue then
      local ast = self:createAst( self.option.scriptPath, mod, nil, TransUnit.AnalyzeMode.Compile )
      
      local glue = glueFilter.glueFilter.new(self.option.outputDir)
      
      ast:get_node():processFilter( glue )
    elseif _switchExp == Option.ModeKind.Lua or _switchExp == Option.ModeKind.LuaMeta then
      local ast = self:createAst( self.option.scriptPath, mod, nil, TransUnit.AnalyzeMode.Compile )
      
      local convMode = convLua.ConvMode.Convert
      
      if self.option.mode == Option.ModeKind.LuaMeta then
        convMode = convLua.ConvMode.ConvMeta
      end
      convert( ast, self.option.scriptPath, io.stdout, io.stdout, convMode, false )
    elseif _switchExp == Option.ModeKind.Save or _switchExp == Option.ModeKind.SaveMeta then
      Util.profile( self.option.validProf, function (  )
        local ast = self:createAst( self.option.scriptPath, mod, nil, TransUnit.AnalyzeMode.Compile )
        
        local luaPath = self.option.scriptPath:gsub( "%.lns$", ".lua" )
        
        local metaPath = self.option.scriptPath:gsub( "%.lns$", ".meta" )
        
        if self.option.outputDir then
          local filename = mod:gsub( "%.", "/" )
          
          luaPath = string.format( "%s/%s.lua", self.option.outputDir, filename )
          metaPath = string.format( "%s/%s.meta", self.option.outputDir, filename )
        end
        if luaPath ~= self.option.scriptPath then
          local fileObj = io.open( luaPath, "w" )
          
              if  nil == fileObj then
                local _fileObj = fileObj
                
                error( string.format( "write open error -- %s", luaPath) )
              end
            
          local stream = fileObj
          
          local metaFileObj = nil
          
          local metaStream = stream
          
          local convMode = convLua.ConvMode.Convert
          
          if self.option.mode == "SAVE" then
            convMode = convLua.ConvMode.ConvMeta
            do
              local _exp = io.open( metaPath, "w" )
              if _exp ~= nil then
              
                  metaStream = _exp
                else
              
                  error( string.format( "write open error -- %s", metaPath) )
                end
            end
            
          end
          convert( ast, self.option.scriptPath, stream, metaStream, convMode, false )
          fileObj:close(  )
          do
            local _exp = metaFileObj
            if _exp ~= nil then
            
                _exp:close(  )
              end
          end
          
        end
      end
      , self.option.scriptPath .. ".profi" )
    elseif _switchExp == Option.ModeKind.Exec then
      self:loadModule( mod )
    else 
      print( "illegal mode" )
    end
  end
  
end

local function exec( args )
  local version = tonumber( _VERSION:gsub( "^[^%d]+", "" ), nil )
  
  if version < 5.2 then
    io.stderr:write( string.format( "LuneScript doesn't support this lua version(%s). %s\n", version, "please use the version after 5.2." ) )
    os.exit( 1 )
  end
  local option = Option.analyze( args )
  
  local front = Front.new(option, {}, {}, {})
  
  frontInterface.setFront( front )
  front:exec(  )
end
_moduleObj.exec = exec
return _moduleObj
