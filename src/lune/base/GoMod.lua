--lune/base/GoMod.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@GoMod'
local _lune = {}
if _lune6 then
   _lune = _lune6
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

function _lune.nilacc( val, fieldName, access, ... )
   if not val then
      return nil
   end
   if fieldName then
      local field = val[ fieldName ]
      if not field then
         return nil
      end
      if access == "item" then
         local typeId = type( field )
         if typeId == "table" then
            return field[ ... ]
         elseif typeId == "string" then
            return string.byte( field, ... )
         end
      elseif access == "call" then
         return field( ... )
      elseif access == "callmtd" then
         return field( val, ... )
      end
      return field
   end
   if access == "item" then
      local typeId = type( val )
      if typeId == "table" then
         return val[ ... ]
      elseif typeId == "string" then
         return string.byte( val, ... )
      end
   elseif access == "call" then
      return val( ... )
   elseif access == "list" then
      local list, arg = ...
      if not list then
         return nil
      end
      return val( list, arg )
   end
   error( string.format( "illegal access -- %s", access ) )
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

if not _lune6 then
   _lune6 = _lune
end


local Types = _lune.loadModule( 'lune.base.Types' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Depend = _lune.loadModule( 'lune.base.Depend' )
local Log = _lune.loadModule( 'lune.base.Log' )

local ModProjInfo = {}
_moduleObj.ModProjInfo = ModProjInfo
function ModProjInfo._setmeta( obj )
  setmetatable( obj, { __index = ModProjInfo  } )
end
function ModProjInfo._new( path, projRoot, mod, fullMod )
   local obj = {}
   ModProjInfo._setmeta( obj )
   if obj.__init then
      obj:__init( path, projRoot, mod, fullMod )
   end
   return obj
end
function ModProjInfo:__init( path, projRoot, mod, fullMod )

   self.path = path
   self.projRoot = projRoot
   self.mod = mod
   self.fullMod = fullMod
end
function ModProjInfo:get_path()
   return self.path
end
function ModProjInfo:get_projRoot()
   return self.projRoot
end
function ModProjInfo:get_mod()
   return self.mod
end
function ModProjInfo:get_fullMod()
   return self.fullMod
end


local GoModResult = {}
GoModResult._name2Val = {}
_moduleObj.GoModResult = GoModResult
function GoModResult:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "GoModResult.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function GoModResult._from( val )
   return _lune._AlgeFrom( GoModResult, val )
end

GoModResult.Found = { "Found", {{}}}
GoModResult._name2Val["Found"] = GoModResult.Found
GoModResult.NotFound = { "NotFound"}
GoModResult._name2Val["NotFound"] = GoModResult.NotFound
GoModResult.NotGo = { "NotGo"}
GoModResult._name2Val["NotGo"] = GoModResult.NotGo


local ModInfo = {}
_moduleObj.ModInfo = ModInfo
function ModInfo:getGoModPath( ver, mod )

   local pathList = {}
   
   do
      local gopath = Depend.getGOPATH(  )
      if gopath ~= nil then
         table.insert( pathList, Util.pathJoin( gopath, string.format( "src/%s", mod) ) )
         
         local gomod = ""
         for __index, aChar in pairs( {mod:byte( 1, #mod )} ) do
            if aChar ~= nil then
               if aChar >= 65 and aChar <= 90 then
                  gomod = string.format( "%s!%c", gomod, aChar - 65 + 97)
               else
                
                  gomod = string.format( "%s%c", gomod, aChar)
               end
               
            end
            
         end
         
         gomod = string.format( "%s@%s", gomod, ver)
         
         table.insert( pathList, Util.pathJoin( gopath, string.format( "pkg/mod/%s", gomod) ) )
      end
   end
   
   return pathList
end
function ModInfo:getModPathList(  )

   local list = {}
   for mod, ver in pairs( self.moduleMap ) do
      do
         local _exp = self.replaceMap[mod]
         if _exp ~= nil then
            table.insert( list, _exp )
         else
            for __index, path in ipairs( self:getGoModPath( ver, mod ) ) do
               if Depend.existFile( path ) then
                  table.insert( list, path )
                  break
               end
               
            end
            
         end
      end
      
   end
   
   return list
end
function ModInfo._new( name, moduleMap, replaceMap )
   local obj = {}
   ModInfo._setmeta( obj )
   if obj.__init then obj:__init( name, moduleMap, replaceMap ); end
   return obj
end
function ModInfo:__init(name, moduleMap, replaceMap) 
   self.name = name
   self.moduleMap = moduleMap
   self.replaceMap = replaceMap
   self.goModDir2Path = {}
   self.path2modProjInfo = {}
   self.latestModProjInfo = nil
   
   for mod, dst in pairs( replaceMap ) do
      self.goModDir2Path[dst] = mod
   end
   
   for mod, ver in pairs( moduleMap ) do
      for __index, path in ipairs( self:getGoModPath( ver, mod ) ) do
         self.goModDir2Path[path] = mod
      end
      
   end
   
end
function ModInfo:getLatestProjRoot(  )

   return _lune.nilacc( self.latestModProjInfo, 'get_projRoot', 'callmtd' )
end
function ModInfo:getLocalModulePathList( path )

   local pathList = {}
   for mod, ver in pairs( self.moduleMap ) do
      if path:find( mod, 1, true ) == 1 then
         local relativeName = path:sub( #mod + 2 )
         
         do
            local replacePath = self.replaceMap[mod]
            if replacePath ~= nil then
               
               table.insert( pathList, Util.pathJoin( replacePath, relativeName ) )
               break
            end
         end
         
         
         do
            local gopath = Depend.getGOPATH(  )
            if gopath ~= nil then
               table.insert( pathList, Util.pathJoin( gopath, string.format( "src/%s", path) ) )
               
               local gomod = ""
               for __index, aChar in pairs( {mod:byte( 1, #mod )} ) do
                  if aChar ~= nil then
                     if aChar >= 65 and aChar <= 90 then
                        gomod = string.format( "%s!%c", gomod, aChar - 65 + 97)
                     else
                      
                        gomod = string.format( "%s%c", gomod, aChar)
                     end
                     
                  end
                  
               end
               
               gomod = string.format( "%s@%s", gomod, ver)
               
               table.insert( pathList, Util.pathJoin( gopath, string.format( "pkg/mod/%s/%s", gomod, relativeName) ) )
            end
         end
         
         break
      end
      
   end
   
   return pathList
end
function ModInfo:convPath( mod, suffix )

   return mod:gsub( "^go/", "" ):gsub( "%.", "/" ):gsub( ":", "." ) .. suffix
end
function ModInfo:getProjRootPath( mod, path )

   local convPath = self:convPath( mod, ".lns" ):gsub( "github%.com/[^/]+/[^/]+/", "" )
   local projRoot = path:sub( 1, #path - #convPath )
   if projRoot ~= "/" then
      projRoot = projRoot:gsub( "/$", "" )
   end
   
   
   path = Util.parentPath( path )
   local modList = Util.splitStr( mod, "[^%.]+" )
   local startIndex = #modList
   for modIndex = 1, #modList do
      if Depend.existFile( Util.pathJoin( path, "lune.js" ) ) then
         startIndex = modIndex
         break
      end
      
      if path == projRoot then
         startIndex = modIndex
         break
      end
      
      
      path = Util.parentPath( path )
   end
   
   
   local convMod = ""
   for index = #modList - startIndex + 1, #modList do
      if convMod ~= "" then
         convMod = string.format( "%s.", convMod)
      end
      
      convMod = convMod .. modList[index]
   end
   
   return path, convMod
end
function ModInfo:convLocalModulePath( mod, suffix, baseDir )
   local __func__ = '@lune.@base.@GoMod.ModInfo.convLocalModulePath'

   if not mod:find( "^go/" ) then
      if baseDir ~= nil then
         local goModDir = nil
         goModDir = self.goModDir2Path[baseDir]
         if goModDir ~= nil then
            mod = string.format( "go/%s.%s", goModDir:gsub( "%.", ":" ):gsub( "/", "." ), mod)
         end
         
         if not goModDir then
            Log.log( Log.Level.Log, __func__, 195, function (  )
            
               return string.format( "not found baseDir -- %s", baseDir)
            end )
            
         end
         
      else
         return _lune.newAlge( GoModResult.NotGo)
      end
      
   end
   
   
   local workMod = self:convPath( mod, suffix )
   
   do
      local _exp = self.path2modProjInfo[workMod]
      if _exp ~= nil then
         self.latestModProjInfo = _exp
         return _lune.newAlge( GoModResult.Found, {_exp})
      end
   end
   
   
   local pathList = self:getLocalModulePathList( workMod )
   if not mod:find( "^go/" ) then
      if baseDir ~= nil then
         table.insert( pathList, Util.pathJoin( baseDir, workMod ) )
      end
      
   end
   
   table.insert( pathList, Util.pathJoin( "vendor", workMod ) )
   
   for __index, path in ipairs( pathList ) do
      if Depend.existFile( path ) then
         local projRoot, convMod = self:getProjRootPath( mod, path )
         local projInfo = ModProjInfo._new(path, projRoot, convMod, mod)
         self.path2modProjInfo[workMod] = projInfo
         self.latestModProjInfo = projInfo
         return _lune.newAlge( GoModResult.Found, {projInfo})
      else
       
         Log.log( Log.Level.Log, __func__, 229, function (  )
         
            return string.format( "not found %s", path)
         end )
         
      end
      
   end
   
   return _lune.newAlge( GoModResult.NotFound)
end
function ModInfo:getLuaModulePath( mod, baseDir )

   local info
   
   do
      local _matchExp = self:convLocalModulePath( mod, ".lns", baseDir )
      if _matchExp[1] == GoModResult.NotGo[1] then
      
         return mod, nil, mod
      elseif _matchExp[1] == GoModResult.NotFound[1] then
      
         return mod, nil, mod
      elseif _matchExp[1] == GoModResult.Found[1] then
         local workInfo = _matchExp[2][1]
      
         info = workInfo
      end
   end
   
   
   return info:get_mod(), info:get_projRoot(), info:get_fullMod()
end
function ModInfo._setmeta( obj )
  setmetatable( obj, { __index = ModInfo  } )
end
function ModInfo:get_name()
   return self.name
end
function ModInfo:get_moduleMap()
   return self.moduleMap
end
function ModInfo:get_replaceMap()
   return self.replaceMap
end


local BlockKind = {}
BlockKind._val2NameMap = {}
function BlockKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "BlockKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function BlockKind._from( val )
   if BlockKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
BlockKind.__allList = {}
function BlockKind.get__allList()
   return BlockKind.__allList
end

BlockKind.None = 0
BlockKind._val2NameMap[0] = 'None'
BlockKind.__allList[1] = BlockKind.None
BlockKind.Require = 1
BlockKind._val2NameMap[1] = 'Require'
BlockKind.__allList[2] = BlockKind.Require
BlockKind.Replace = 2
BlockKind._val2NameMap[2] = 'Replace'
BlockKind.__allList[3] = BlockKind.Replace


local function getReplace( map, tokenList, modIndex )

   local prevArrow = false
   for __index, token in ipairs( tokenList ) do
      if token == "=>" then
         prevArrow = true
      elseif prevArrow then
         map[tokenList[modIndex]] = token
         break
      end
      
   end
   
end

local function getGoMap(  )

   local requireMap = {}
   local replaceMap = {}
   local name = "lnsc"
   do
      local file = io.open( "go.mod" )
      if file ~= nil then
         local inBlock = BlockKind.None
         while true do
            local line = file:read( "*l" )
            if  nil == line then
               local _line = line
            
               break
            end
            
            local trimedLine = line:gsub( "^%s", "" )
            local tokenList = Util.splitStr( trimedLine, "[^%s]+" )
            do
               local _switchExp = inBlock
               if _switchExp == BlockKind.Require then
                  if line:find( "^%)" ) then
                     inBlock = BlockKind.None
                  else
                   
                     requireMap[tokenList[1]] = tokenList[2]
                  end
                  
               elseif _switchExp == BlockKind.Replace then
                  if line:find( "^%)" ) then
                     inBlock = BlockKind.None
                  else
                   
                     getReplace( replaceMap, tokenList, 1 )
                  end
                  
               elseif _switchExp == BlockKind.None then
                  if line:find( "^module%s+" ) then
                     name = tokenList[2]
                  elseif line:find( "^require%s+[^%(]" ) then
                     if #tokenList == 3 then
                        requireMap[tokenList[2]] = tokenList[3]
                     end
                     
                  elseif line:find( "^require%s+%(" ) then
                     inBlock = BlockKind.Require
                  elseif line:find( "^replace%s+[^%(]" ) then
                     getReplace( replaceMap, tokenList, 2 )
                  elseif line:find( "^replace%s+%(" ) then
                     inBlock = BlockKind.Replace
                  end
                  
               end
            end
            
         end
         
         file:close(  )
      end
   end
   
   local modInfo = ModInfo._new(name, requireMap, replaceMap)
   return modInfo
end
_moduleObj.getGoMap = getGoMap

return _moduleObj
