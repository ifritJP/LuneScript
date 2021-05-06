--lune/base/Builtin.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Builtin'
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
local Util = _lune.loadModule( 'lune.base.Util' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Log = _lune.loadModule( 'lune.base.Log' )

local TransUnitIF = _lune.loadModule( 'lune.base.TransUnitIF' )



local Builtin = {}
_moduleObj.Builtin = Builtin
function Builtin.new( transUnitIF, targetLuaVer, ctrl_info )
   local obj = {}
   Builtin.setmeta( obj )
   if obj.__init then obj:__init( transUnitIF, targetLuaVer, ctrl_info ); end
   return obj
end
function Builtin:__init(transUnitIF, targetLuaVer, ctrl_info) 
   self.transUnitIF = transUnitIF
   self.targetLuaVer = targetLuaVer
   self.ctrl_info = ctrl_info
   self.processInfo = Ast.getRootProcessInfo(  )
   self.modifier = TransUnitIF.Modifier.new(true, self.processInfo)
end
function Builtin.setmeta( obj )
  setmetatable( obj, { __index = Builtin  } )
end





local BuiltinFuncType = {}
_moduleObj.BuiltinFuncType = BuiltinFuncType
function BuiltinFuncType.new(  )
   local obj = {}
   BuiltinFuncType.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function BuiltinFuncType:__init() 
   self.lns_ = Ast.headTypeInfo
   self.lns__fcall = Ast.headTypeInfo
   self.lns__fcall_sym = Ast.dummySymbol
   self.lns__kind = Ast.headTypeInfo
   self.lns__kind_sym = Ast.dummySymbol
   self.lns__load = Ast.headTypeInfo
   self.lns__load_sym = Ast.dummySymbol
   self.lns_collectgarbage = Ast.headTypeInfo
   self.lns_collectgarbage_sym = Ast.dummySymbol
   self.lns_error = Ast.headTypeInfo
   self.lns_error_sym = Ast.dummySymbol
   self.lns_expandLuavalMap = Ast.headTypeInfo
   self.lns_expandLuavalMap_sym = Ast.dummySymbol
   self.lns_loadfile = Ast.headTypeInfo
   self.lns_loadfile_sym = Ast.dummySymbol
   self.lns_print = Ast.headTypeInfo
   self.lns_print_sym = Ast.dummySymbol
   self.lns_require = Ast.headTypeInfo
   self.lns_require_sym = Ast.dummySymbol
   self.lns_tonumber = Ast.headTypeInfo
   self.lns_tonumber_sym = Ast.dummySymbol
   self.lns_tostring = Ast.headTypeInfo
   self.lns_tostring_sym = Ast.dummySymbol
   self.lns_type = Ast.headTypeInfo
   self.lns_type_sym = Ast.dummySymbol
   self.istream_ = Ast.headTypeInfo
   self.istream___attrib = Ast.headTypeInfo
   self.istream___attrib_sym = Ast.dummySymbol
   self.istream_close = Ast.headTypeInfo
   self.istream_close_sym = Ast.dummySymbol
   self.istream_read = Ast.headTypeInfo
   self.istream_read_sym = Ast.dummySymbol
   self.ostream_ = Ast.headTypeInfo
   self.ostream___attrib = Ast.headTypeInfo
   self.ostream___attrib_sym = Ast.dummySymbol
   self.ostream_close = Ast.headTypeInfo
   self.ostream_close_sym = Ast.dummySymbol
   self.ostream_flush = Ast.headTypeInfo
   self.ostream_flush_sym = Ast.dummySymbol
   self.ostream_write = Ast.headTypeInfo
   self.ostream_write_sym = Ast.dummySymbol
   self.luastream_ = Ast.headTypeInfo
   self.luastream___attrib = Ast.headTypeInfo
   self.luastream___attrib_sym = Ast.dummySymbol
   self.luastream_close = Ast.headTypeInfo
   self.luastream_close_sym = Ast.dummySymbol
   self.luastream_flush = Ast.headTypeInfo
   self.luastream_flush_sym = Ast.dummySymbol
   self.luastream_read = Ast.headTypeInfo
   self.luastream_read_sym = Ast.dummySymbol
   self.luastream_seek = Ast.headTypeInfo
   self.luastream_seek_sym = Ast.dummySymbol
   self.luastream_write = Ast.headTypeInfo
   self.luastream_write_sym = Ast.dummySymbol
   self.mapping_ = Ast.headTypeInfo
   self.mapping___attrib = Ast.headTypeInfo
   self.mapping___attrib_sym = Ast.dummySymbol
   self.mapping__toMap = Ast.headTypeInfo
   self.mapping__toMap_sym = Ast.dummySymbol
   self.__pipe_ = Ast.headTypeInfo
   self.__pipe_get = Ast.headTypeInfo
   self.__pipe_get_sym = Ast.dummySymbol
   self.__pipe_put = Ast.headTypeInfo
   self.__pipe_put_sym = Ast.dummySymbol
   self.lnsthread_ = Ast.headTypeInfo
   self.lnsthread___init = Ast.headTypeInfo
   self.lnsthread___init_sym = Ast.dummySymbol
   self.lnsthread_loop = Ast.headTypeInfo
   self.lnsthread_loop_sym = Ast.dummySymbol
   self.io_ = Ast.headTypeInfo
   self.io___attrib = Ast.headTypeInfo
   self.io___attrib_sym = Ast.dummySymbol
   self.io_open = Ast.headTypeInfo
   self.io_open_sym = Ast.dummySymbol
   self.io_popen = Ast.headTypeInfo
   self.io_popen_sym = Ast.dummySymbol
   self.io_stderr = Ast.headTypeInfo
   self.io_stderr_sym = Ast.dummySymbol
   self.io_stdin = Ast.headTypeInfo
   self.io_stdin_sym = Ast.dummySymbol
   self.io_stdout = Ast.headTypeInfo
   self.io_stdout_sym = Ast.dummySymbol
   self.package_ = Ast.headTypeInfo
   self.package___attrib = Ast.headTypeInfo
   self.package___attrib_sym = Ast.dummySymbol
   self.package_path = Ast.headTypeInfo
   self.package_path_sym = Ast.dummySymbol
   self.package_searchpath = Ast.headTypeInfo
   self.package_searchpath_sym = Ast.dummySymbol
   self.os_ = Ast.headTypeInfo
   self.os___attrib = Ast.headTypeInfo
   self.os___attrib_sym = Ast.dummySymbol
   self.os_clock = Ast.headTypeInfo
   self.os_clock_sym = Ast.dummySymbol
   self.os_date = Ast.headTypeInfo
   self.os_date_sym = Ast.dummySymbol
   self.os_difftime = Ast.headTypeInfo
   self.os_difftime_sym = Ast.dummySymbol
   self.os_exit = Ast.headTypeInfo
   self.os_exit_sym = Ast.dummySymbol
   self.os_remove = Ast.headTypeInfo
   self.os_remove_sym = Ast.dummySymbol
   self.os_rename = Ast.headTypeInfo
   self.os_rename_sym = Ast.dummySymbol
   self.os_time = Ast.headTypeInfo
   self.os_time_sym = Ast.dummySymbol
   self.string_ = Ast.headTypeInfo
   self.string___attrib = Ast.headTypeInfo
   self.string___attrib_sym = Ast.dummySymbol
   self.string_byte = Ast.headTypeInfo
   self.string_byte_sym = Ast.dummySymbol
   self.string_dump = Ast.headTypeInfo
   self.string_dump_sym = Ast.dummySymbol
   self.string_find = Ast.headTypeInfo
   self.string_find_sym = Ast.dummySymbol
   self.string_format = Ast.headTypeInfo
   self.string_format_sym = Ast.dummySymbol
   self.string_gmatch = Ast.headTypeInfo
   self.string_gmatch_sym = Ast.dummySymbol
   self.string_gsub = Ast.headTypeInfo
   self.string_gsub_sym = Ast.dummySymbol
   self.string_lower = Ast.headTypeInfo
   self.string_lower_sym = Ast.dummySymbol
   self.string_rep = Ast.headTypeInfo
   self.string_rep_sym = Ast.dummySymbol
   self.string_reverse = Ast.headTypeInfo
   self.string_reverse_sym = Ast.dummySymbol
   self.string_sub = Ast.headTypeInfo
   self.string_sub_sym = Ast.dummySymbol
   self.string_upper = Ast.headTypeInfo
   self.string_upper_sym = Ast.dummySymbol
   self.str_ = Ast.headTypeInfo
   self.str___attrib = Ast.headTypeInfo
   self.str___attrib_sym = Ast.dummySymbol
   self.str_byte = Ast.headTypeInfo
   self.str_byte_sym = Ast.dummySymbol
   self.str_find = Ast.headTypeInfo
   self.str_find_sym = Ast.dummySymbol
   self.str_format = Ast.headTypeInfo
   self.str_format_sym = Ast.dummySymbol
   self.str_gmatch = Ast.headTypeInfo
   self.str_gmatch_sym = Ast.dummySymbol
   self.str_gsub = Ast.headTypeInfo
   self.str_gsub_sym = Ast.dummySymbol
   self.str_lower = Ast.headTypeInfo
   self.str_lower_sym = Ast.dummySymbol
   self.str_rep = Ast.headTypeInfo
   self.str_rep_sym = Ast.dummySymbol
   self.str_reverse = Ast.headTypeInfo
   self.str_reverse_sym = Ast.dummySymbol
   self.str_sub = Ast.headTypeInfo
   self.str_sub_sym = Ast.dummySymbol
   self.str_upper = Ast.headTypeInfo
   self.str_upper_sym = Ast.dummySymbol
   self.list_ = Ast.headTypeInfo
   self.list___less = Ast.headTypeInfo
   self.list___less_sym = Ast.dummySymbol
   self.list_insert = Ast.headTypeInfo
   self.list_insert_sym = Ast.dummySymbol
   self.list_remove = Ast.headTypeInfo
   self.list_remove_sym = Ast.dummySymbol
   self.list_sort = Ast.headTypeInfo
   self.list_sort_sym = Ast.dummySymbol
   self.list_unpack = Ast.headTypeInfo
   self.list_unpack_sym = Ast.dummySymbol
   self.array_ = Ast.headTypeInfo
   self.array___less = Ast.headTypeInfo
   self.array___less_sym = Ast.dummySymbol
   self.array_sort = Ast.headTypeInfo
   self.array_sort_sym = Ast.dummySymbol
   self.array_unpack = Ast.headTypeInfo
   self.array_unpack_sym = Ast.dummySymbol
   self.set_ = Ast.headTypeInfo
   self.set_add = Ast.headTypeInfo
   self.set_add_sym = Ast.dummySymbol
   self.set_and = Ast.headTypeInfo
   self.set_and_sym = Ast.dummySymbol
   self.set_clone = Ast.headTypeInfo
   self.set_clone_sym = Ast.dummySymbol
   self.set_del = Ast.headTypeInfo
   self.set_del_sym = Ast.dummySymbol
   self.set_has = Ast.headTypeInfo
   self.set_has_sym = Ast.dummySymbol
   self.set_len = Ast.headTypeInfo
   self.set_len_sym = Ast.dummySymbol
   self.set_or = Ast.headTypeInfo
   self.set_or_sym = Ast.dummySymbol
   self.set_sub = Ast.headTypeInfo
   self.set_sub_sym = Ast.dummySymbol
   self.math_ = Ast.headTypeInfo
   self.math___attrib = Ast.headTypeInfo
   self.math___attrib_sym = Ast.dummySymbol
   self.math_random = Ast.headTypeInfo
   self.math_random_sym = Ast.dummySymbol
   self.math_randomseed = Ast.headTypeInfo
   self.math_randomseed_sym = Ast.dummySymbol
   self.debug_ = Ast.headTypeInfo
   self.debug___attrib = Ast.headTypeInfo
   self.debug___attrib_sym = Ast.dummySymbol
   self.debug_getinfo = Ast.headTypeInfo
   self.debug_getinfo_sym = Ast.dummySymbol
   self.debug_getlocal = Ast.headTypeInfo
   self.debug_getlocal_sym = Ast.dummySymbol
   self.nilable_ = Ast.headTypeInfo
   self.nilable_val = Ast.headTypeInfo
   self.nilable_val_sym = Ast.dummySymbol
   
   
   
   self.allSymbol = {}
   self.allClass = {}
   self.allSymbolSet = {}
   self.allFuncTypeSet = {}
   self.needThreadingTypes = {}
end
function BuiltinFuncType:register( symbolInfo )

   table.insert( self.allSymbol, symbolInfo )
   self.allSymbolSet[symbolInfo]= true
end
function BuiltinFuncType:registerClass( classInfo )

   table.insert( self.allClass, classInfo )
   do
      local _switchExp = classInfo:get_rawTxt()
      if _switchExp == '' then
         self.lns_ = classInfo
      elseif _switchExp == 'iStream' then
         self.istream_ = classInfo
      elseif _switchExp == 'oStream' then
         self.ostream_ = classInfo
      elseif _switchExp == 'luaStream' then
         self.luastream_ = classInfo
      elseif _switchExp == 'Mapping' then
         self.mapping_ = classInfo
      elseif _switchExp == '__pipe' then
         self.__pipe_ = classInfo
      elseif _switchExp == 'LnsThread' then
         self.lnsthread_ = classInfo
      elseif _switchExp == 'io' then
         self.io_ = classInfo
      elseif _switchExp == 'package' then
         self.package_ = classInfo
      elseif _switchExp == 'os' then
         self.os_ = classInfo
      elseif _switchExp == 'string' then
         self.string_ = classInfo
      elseif _switchExp == 'str' then
         self.str_ = classInfo
      elseif _switchExp == 'List' then
         self.list_ = classInfo
      elseif _switchExp == 'Array' then
         self.array_ = classInfo
      elseif _switchExp == 'Set' then
         self.set_ = classInfo
      elseif _switchExp == 'math' then
         self.math_ = classInfo
      elseif _switchExp == 'debug' then
         self.debug_ = classInfo
      elseif _switchExp == 'Nilable' then
         self.nilable_ = classInfo
      end
   end
   
end
function BuiltinFuncType.setmeta( obj )
  setmetatable( obj, { __index = BuiltinFuncType  } )
end
function BuiltinFuncType:get_allSymbol()
   return self.allSymbol
end
function BuiltinFuncType:get_allClass()
   return self.allClass
end
function BuiltinFuncType:get_allFuncTypeSet()
   return self.allFuncTypeSet
end
function BuiltinFuncType:get_allSymbolSet()
   return self.allSymbolSet
end
function BuiltinFuncType:get_needThreadingTypes()
   return self.needThreadingTypes
end


local builtinFunc = BuiltinFuncType.new()

local function getBuiltinFunc(  )

   return builtinFunc
end
_moduleObj.getBuiltinFunc = getBuiltinFunc

local function setupBuiltinTypeInfo( name, fieldName, symInfo )

   local typeInfo = symInfo:get_typeInfo()
   local function process_(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '_fcall' then
            builtinFunc.lns__fcall = typeInfo
            builtinFunc.lns__fcall_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == '_kind' then
            builtinFunc.lns__kind = typeInfo
            builtinFunc.lns__kind_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == '_load' then
            builtinFunc.lns__load = typeInfo
            builtinFunc.lns__load_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'collectgarbage' then
            builtinFunc.lns_collectgarbage = typeInfo
            builtinFunc.lns_collectgarbage_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'error' then
            builtinFunc.lns_error = typeInfo
            builtinFunc.lns_error_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'expandLuavalMap' then
            builtinFunc.lns_expandLuavalMap = typeInfo
            builtinFunc.lns_expandLuavalMap_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'loadfile' then
            builtinFunc.lns_loadfile = typeInfo
            builtinFunc.lns_loadfile_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'print' then
            builtinFunc.lns_print = typeInfo
            builtinFunc.lns_print_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'require' then
            builtinFunc.lns_require = typeInfo
            builtinFunc.lns_require_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'tonumber' then
            builtinFunc.lns_tonumber = typeInfo
            builtinFunc.lns_tonumber_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'tostring' then
            builtinFunc.lns_tostring = typeInfo
            builtinFunc.lns_tostring_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'type' then
            builtinFunc.lns_type = typeInfo
            builtinFunc.lns_type_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_iStream(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.istream___attrib = typeInfo
            builtinFunc.istream___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'close' then
            builtinFunc.istream_close = typeInfo
            builtinFunc.istream_close_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'read' then
            builtinFunc.istream_read = typeInfo
            builtinFunc.istream_read_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_oStream(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.ostream___attrib = typeInfo
            builtinFunc.ostream___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'close' then
            builtinFunc.ostream_close = typeInfo
            builtinFunc.ostream_close_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'flush' then
            builtinFunc.ostream_flush = typeInfo
            builtinFunc.ostream_flush_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'write' then
            builtinFunc.ostream_write = typeInfo
            builtinFunc.ostream_write_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_luaStream(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.luastream___attrib = typeInfo
            builtinFunc.luastream___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'close' then
            builtinFunc.luastream_close = typeInfo
            builtinFunc.luastream_close_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'flush' then
            builtinFunc.luastream_flush = typeInfo
            builtinFunc.luastream_flush_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'read' then
            builtinFunc.luastream_read = typeInfo
            builtinFunc.luastream_read_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'seek' then
            builtinFunc.luastream_seek = typeInfo
            builtinFunc.luastream_seek_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'write' then
            builtinFunc.luastream_write = typeInfo
            builtinFunc.luastream_write_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_Mapping(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.mapping___attrib = typeInfo
            builtinFunc.mapping___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == '_toMap' then
            builtinFunc.mapping__toMap = typeInfo
            builtinFunc.mapping__toMap_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process___pipe(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == 'get' then
            builtinFunc.__pipe_get = typeInfo
            builtinFunc.__pipe_get_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'put' then
            builtinFunc.__pipe_put = typeInfo
            builtinFunc.__pipe_put_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_LnsThread(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__init' then
            builtinFunc.lnsthread___init = typeInfo
            builtinFunc.lnsthread___init_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'loop' then
            builtinFunc.lnsthread_loop = typeInfo
            builtinFunc.lnsthread_loop_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_io(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.io___attrib = typeInfo
            builtinFunc.io___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'open' then
            builtinFunc.io_open = typeInfo
            builtinFunc.io_open_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'popen' then
            builtinFunc.io_popen = typeInfo
            builtinFunc.io_popen_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'stderr' then
            builtinFunc.io_stderr = typeInfo
            builtinFunc.io_stderr_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'stdin' then
            builtinFunc.io_stdin = typeInfo
            builtinFunc.io_stdin_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'stdout' then
            builtinFunc.io_stdout = typeInfo
            builtinFunc.io_stdout_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_package(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.package___attrib = typeInfo
            builtinFunc.package___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'path' then
            builtinFunc.package_path = typeInfo
            builtinFunc.package_path_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'searchpath' then
            builtinFunc.package_searchpath = typeInfo
            builtinFunc.package_searchpath_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_os(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.os___attrib = typeInfo
            builtinFunc.os___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'clock' then
            builtinFunc.os_clock = typeInfo
            builtinFunc.os_clock_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'date' then
            builtinFunc.os_date = typeInfo
            builtinFunc.os_date_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'difftime' then
            builtinFunc.os_difftime = typeInfo
            builtinFunc.os_difftime_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'exit' then
            builtinFunc.os_exit = typeInfo
            builtinFunc.os_exit_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'remove' then
            builtinFunc.os_remove = typeInfo
            builtinFunc.os_remove_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'rename' then
            builtinFunc.os_rename = typeInfo
            builtinFunc.os_rename_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'time' then
            builtinFunc.os_time = typeInfo
            builtinFunc.os_time_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_string(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.string___attrib = typeInfo
            builtinFunc.string___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'byte' then
            builtinFunc.string_byte = typeInfo
            builtinFunc.string_byte_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'dump' then
            builtinFunc.string_dump = typeInfo
            builtinFunc.string_dump_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'find' then
            builtinFunc.string_find = typeInfo
            builtinFunc.string_find_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'format' then
            builtinFunc.string_format = typeInfo
            builtinFunc.string_format_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'gmatch' then
            builtinFunc.string_gmatch = typeInfo
            builtinFunc.string_gmatch_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'gsub' then
            builtinFunc.string_gsub = typeInfo
            builtinFunc.string_gsub_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'lower' then
            builtinFunc.string_lower = typeInfo
            builtinFunc.string_lower_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'rep' then
            builtinFunc.string_rep = typeInfo
            builtinFunc.string_rep_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'reverse' then
            builtinFunc.string_reverse = typeInfo
            builtinFunc.string_reverse_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'sub' then
            builtinFunc.string_sub = typeInfo
            builtinFunc.string_sub_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'upper' then
            builtinFunc.string_upper = typeInfo
            builtinFunc.string_upper_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_str(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.str___attrib = typeInfo
            builtinFunc.str___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'byte' then
            builtinFunc.str_byte = typeInfo
            builtinFunc.str_byte_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'find' then
            builtinFunc.str_find = typeInfo
            builtinFunc.str_find_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'format' then
            builtinFunc.str_format = typeInfo
            builtinFunc.str_format_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'gmatch' then
            builtinFunc.str_gmatch = typeInfo
            builtinFunc.str_gmatch_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'gsub' then
            builtinFunc.str_gsub = typeInfo
            builtinFunc.str_gsub_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'lower' then
            builtinFunc.str_lower = typeInfo
            builtinFunc.str_lower_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'rep' then
            builtinFunc.str_rep = typeInfo
            builtinFunc.str_rep_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'reverse' then
            builtinFunc.str_reverse = typeInfo
            builtinFunc.str_reverse_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'sub' then
            builtinFunc.str_sub = typeInfo
            builtinFunc.str_sub_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'upper' then
            builtinFunc.str_upper = typeInfo
            builtinFunc.str_upper_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_List(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__less' then
            builtinFunc.list___less = typeInfo
            builtinFunc.list___less_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'insert' then
            builtinFunc.list_insert = typeInfo
            builtinFunc.list_insert_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'remove' then
            builtinFunc.list_remove = typeInfo
            builtinFunc.list_remove_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'sort' then
            builtinFunc.list_sort = typeInfo
            builtinFunc.list_sort_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'unpack' then
            builtinFunc.list_unpack = typeInfo
            builtinFunc.list_unpack_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_Array(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__less' then
            builtinFunc.array___less = typeInfo
            builtinFunc.array___less_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'sort' then
            builtinFunc.array_sort = typeInfo
            builtinFunc.array_sort_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'unpack' then
            builtinFunc.array_unpack = typeInfo
            builtinFunc.array_unpack_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_Set(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == 'add' then
            builtinFunc.set_add = typeInfo
            builtinFunc.set_add_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'and' then
            builtinFunc.set_and = typeInfo
            builtinFunc.set_and_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'clone' then
            builtinFunc.set_clone = typeInfo
            builtinFunc.set_clone_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'del' then
            builtinFunc.set_del = typeInfo
            builtinFunc.set_del_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'has' then
            builtinFunc.set_has = typeInfo
            builtinFunc.set_has_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'len' then
            builtinFunc.set_len = typeInfo
            builtinFunc.set_len_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'or' then
            builtinFunc.set_or = typeInfo
            builtinFunc.set_or_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'sub' then
            builtinFunc.set_sub = typeInfo
            builtinFunc.set_sub_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_math(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.math___attrib = typeInfo
            builtinFunc.math___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'random' then
            builtinFunc.math_random = typeInfo
            builtinFunc.math_random_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'randomseed' then
            builtinFunc.math_randomseed = typeInfo
            builtinFunc.math_randomseed_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_debug(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == '__attrib' then
            builtinFunc.debug___attrib = typeInfo
            builtinFunc.debug___attrib_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'getinfo' then
            builtinFunc.debug_getinfo = typeInfo
            builtinFunc.debug_getinfo_sym = symInfo
            builtinFunc:register( symInfo )
         elseif _switchExp == 'getlocal' then
            builtinFunc.debug_getlocal = typeInfo
            builtinFunc.debug_getlocal_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   local function process_Nilable(  )
   
      do
         local _switchExp = fieldName
         if _switchExp == 'val' then
            builtinFunc.nilable_val = typeInfo
            builtinFunc.nilable_val_sym = symInfo
            builtinFunc:register( symInfo )
         end
      end
      
   end
   
   
   
   do
      local _switchExp = name
      if _switchExp == '' then
         process_(  )
      elseif _switchExp == 'iStream' then
         process_iStream(  )
      elseif _switchExp == 'oStream' then
         process_oStream(  )
      elseif _switchExp == 'luaStream' then
         process_luaStream(  )
      elseif _switchExp == 'Mapping' then
         process_Mapping(  )
      elseif _switchExp == '__pipe' then
         process___pipe(  )
      elseif _switchExp == 'LnsThread' then
         process_LnsThread(  )
      elseif _switchExp == 'io' then
         process_io(  )
      elseif _switchExp == 'package' then
         process_package(  )
      elseif _switchExp == 'os' then
         process_os(  )
      elseif _switchExp == 'string' then
         process_string(  )
      elseif _switchExp == 'str' then
         process_str(  )
      elseif _switchExp == 'List' then
         process_List(  )
      elseif _switchExp == 'Array' then
         process_Array(  )
      elseif _switchExp == 'Set' then
         process_Set(  )
      elseif _switchExp == 'math' then
         process_math(  )
      elseif _switchExp == 'debug' then
         process_debug(  )
      elseif _switchExp == 'Nilable' then
         process_Nilable(  )
      end
   end
   
end

local function getBuiltInInfo(  )

   return {{[""] = {["_fcall"] = {["arg"] = {"form", "&..."}, ["ret"] = {""}}, ["_kind"] = {["arg"] = {"stem!"}, ["ret"] = {"int"}}, ["_load"] = {["arg"] = {"str", "stem!"}, ["ret"] = {"Luaval<form>!", "str!"}}, ["collectgarbage"] = {["arg"] = {}, ["ret"] = {}}, ["error"] = {["arg"] = {"str"}, ["ret"] = {"__"}}, ["expandLuavalMap"] = {["arg"] = {"Luaval<&stem>!"}, ["ret"] = {"&stem!"}}, ["loadfile"] = {["arg"] = {"str"}, ["ret"] = {"Luaval<form>!", "str!"}}, ["print"] = {["arg"] = {"&..."}, ["ret"] = {}}, ["require"] = {["arg"] = {"str"}, ["ret"] = {"Luaval<stem>"}}, ["tonumber"] = {["arg"] = {"str", "int!"}, ["ret"] = {"real!"}}, ["tostring"] = {["arg"] = {"&stem"}, ["ret"] = {"str"}}, ["type"] = {["arg"] = {"&stem!"}, ["ret"] = {"str"}}}}, {["iStream"] = {["__attrib"] = {["type"] = {"interface"}}, ["close"] = {["arg"] = {}, ["ret"] = {}, ["type"] = {"mut"}}, ["read"] = {["arg"] = {"stem!"}, ["ret"] = {"str!"}, ["type"] = {"mut"}}}}, {["oStream"] = {["__attrib"] = {["type"] = {"interface"}}, ["close"] = {["arg"] = {}, ["ret"] = {}, ["type"] = {"mut"}}, ["flush"] = {["arg"] = {}, ["ret"] = {}, ["type"] = {"mut"}}, ["write"] = {["arg"] = {"str"}, ["ret"] = {"stem!", "str!"}, ["type"] = {"mut"}}}}, {["luaStream"] = {["__attrib"] = {["inplements"] = {"iStream", "oStream"}, ["type"] = {"interface"}}, ["close"] = {["arg"] = {}, ["ret"] = {}, ["type"] = {"mut"}}, ["flush"] = {["arg"] = {}, ["ret"] = {}, ["type"] = {"mut"}}, ["read"] = {["arg"] = {"stem!"}, ["ret"] = {"str!"}, ["type"] = {"mut"}}, ["seek"] = {["arg"] = {"str", "int"}, ["ret"] = {"int!", "str!"}, ["type"] = {"mut"}}, ["write"] = {["arg"] = {"str"}, ["ret"] = {"stem!", "str!"}, ["type"] = {"mut"}}}}, {["Mapping"] = {["__attrib"] = {["type"] = {"interface"}}, ["_toMap"] = {["arg"] = {}, ["ret"] = {}, ["type"] = {"method"}}}}, {["__pipe<T>"] = {["get"] = {["arg"] = {}, ["ret"] = {"T!"}, ["type"] = {"mut"}}, ["put"] = {["arg"] = {"T!"}, ["ret"] = {}, ["type"] = {"mut"}}}}, {["LnsThread"] = {["__init"] = {["arg"] = {}, ["ret"] = {}, ["type"] = {"method"}}, ["loop"] = {["arg"] = {}, ["ret"] = {}, ["type"] = {"abstract"}}}}, {["io"] = {["__attrib"] = {["type"] = {"module"}}, ["open"] = {["arg"] = {"str", "str!"}, ["ret"] = {"luaStream!", "str!"}}, ["popen"] = {["arg"] = {"str"}, ["ret"] = {"Luaval<luaStream>!"}}, ["stderr"] = {["type"] = {"var"}, ["typeInfo"] = {"oStream"}}, ["stdin"] = {["type"] = {"var"}, ["typeInfo"] = {"iStream"}}, ["stdout"] = {["type"] = {"var"}, ["typeInfo"] = {"oStream"}}}}, {["package"] = {["__attrib"] = {["type"] = {"module"}}, ["path"] = {["type"] = {"var"}, ["typeInfo"] = {"str"}}, ["searchpath"] = {["arg"] = {"str", "str"}, ["ret"] = {"str!"}}}}, {["os"] = {["__attrib"] = {["type"] = {"module"}}, ["clock"] = {["arg"] = {}, ["ret"] = {"real"}}, ["date"] = {["arg"] = {"str!", "stem!"}, ["ret"] = {"Luaval<stem>!"}}, ["difftime"] = {["arg"] = {"&stem", "&stem"}, ["ret"] = {"int"}}, ["exit"] = {["arg"] = {"int!"}, ["ret"] = {"__"}}, ["remove"] = {["arg"] = {"str"}, ["ret"] = {"bool!", "str!"}}, ["rename"] = {["arg"] = {"str", "str"}, ["ret"] = {"stem!", "str!"}}, ["time"] = {["arg"] = {"stem!"}, ["ret"] = {"stem!"}}}}, {["string"] = {["__attrib"] = {["type"] = {"module"}}, ["byte"] = {["arg"] = {"str", "int!", "int!"}, ["ret"] = {"...<int>"}}, ["dump"] = {["arg"] = {"&Luaval<form>", "bool!"}, ["ret"] = {"str"}}, ["find"] = {["arg"] = {"str", "str", "int!", "bool!"}, ["ret"] = {"...<int>"}}, ["format"] = {["arg"] = {"str", "..."}, ["ret"] = {"str"}}, ["gmatch"] = {["arg"] = {"str", "str"}, ["ret"] = {"Luaval<form>", "stem!", "stem!"}}, ["gsub"] = {["arg"] = {"str", "str", "str"}, ["ret"] = {"str", "int"}}, ["lower"] = {["arg"] = {"str"}, ["ret"] = {"str"}}, ["rep"] = {["arg"] = {"str", "int"}, ["ret"] = {"str"}}, ["reverse"] = {["arg"] = {"str"}, ["ret"] = {"str"}}, ["sub"] = {["arg"] = {"str", "int", "int!"}, ["ret"] = {"str"}}, ["upper"] = {["arg"] = {"str"}, ["ret"] = {"str"}}}}, {["str"] = {["__attrib"] = {["inplements"] = {"Mapping"}}, ["byte"] = {["arg"] = {"int!", "int!"}, ["ret"] = {"...<int!>"}, ["type"] = {"method"}}, ["find"] = {["arg"] = {"str", "int!", "bool!"}, ["ret"] = {"...<int>"}, ["type"] = {"method"}}, ["format"] = {["arg"] = {"&..."}, ["ret"] = {"str"}, ["type"] = {"method"}}, ["gmatch"] = {["arg"] = {"str"}, ["ret"] = {"Luaval<form>", "stem!", "stem!"}, ["type"] = {"method"}}, ["gsub"] = {["arg"] = {"str", "str"}, ["ret"] = {"str", "int"}, ["type"] = {"method"}}, ["lower"] = {["arg"] = {}, ["ret"] = {"str"}, ["type"] = {"method"}}, ["rep"] = {["arg"] = {"int"}, ["ret"] = {"str"}, ["type"] = {"method"}}, ["reverse"] = {["arg"] = {}, ["ret"] = {"str"}, ["type"] = {"method"}}, ["sub"] = {["arg"] = {"int", "int!"}, ["ret"] = {"str"}, ["type"] = {"method"}}, ["upper"] = {["arg"] = {}, ["ret"] = {"str"}, ["type"] = {"method"}}}}, {["List<T>"] = {["__less"] = {["arg"] = {"T", "T"}, ["ret"] = {"bool"}, ["type"] = {"formfunc"}}, ["insert"] = {["arg"] = {"T"}, ["ret"] = {""}, ["type"] = {"mut"}}, ["remove"] = {["arg"] = {"int!"}, ["ret"] = {"T!"}, ["type"] = {"mut"}}, ["sort"] = {["arg"] = {"__less!"}, ["ret"] = {}, ["type"] = {"mut"}}, ["unpack"] = {["arg"] = {}, ["ret"] = {"..."}, ["type"] = {"method"}}}}, {["Array<T>"] = {["__less"] = {["arg"] = {"T", "T"}, ["ret"] = {"bool"}, ["type"] = {"formfunc"}}, ["sort"] = {["arg"] = {"__less!"}, ["ret"] = {}, ["type"] = {"mut"}}, ["unpack"] = {["arg"] = {}, ["ret"] = {"..."}, ["type"] = {"method"}}}}, {["Set<T>"] = {["add"] = {["arg"] = {"T"}, ["ret"] = {}, ["type"] = {"mut"}}, ["and"] = {["arg"] = {"&Set<T>"}, ["ret"] = {"Set<T>"}, ["type"] = {"mut"}}, ["clone"] = {["arg"] = {}, ["ret"] = {"Set<T>"}, ["type"] = {"method"}}, ["del"] = {["arg"] = {"T"}, ["ret"] = {}, ["type"] = {"mut"}}, ["has"] = {["arg"] = {"T"}, ["ret"] = {"bool"}, ["type"] = {"method"}}, ["len"] = {["arg"] = {}, ["ret"] = {"int"}, ["type"] = {"method"}}, ["or"] = {["arg"] = {"&Set<T>"}, ["ret"] = {"Set<T>"}, ["type"] = {"mut"}}, ["sub"] = {["arg"] = {"&Set<T>"}, ["ret"] = {"Set<T>"}, ["type"] = {"mut"}}}}, {["math"] = {["__attrib"] = {["type"] = {"module"}}, ["random"] = {["arg"] = {"int!", "int!"}, ["ret"] = {"real"}}, ["randomseed"] = {["arg"] = {"int"}, ["ret"] = {}}}}, {["debug"] = {["__attrib"] = {["type"] = {"module"}}, ["getinfo"] = {["arg"] = {"int"}, ["ret"] = {"Map<str,stem>!"}}, ["getlocal"] = {["arg"] = {"int", "int"}, ["ret"] = {"str!", "stem!"}}}}, {["Nilable<_T>"] = {["val"] = {["arg"] = {}, ["ret"] = {"_T!"}, ["type"] = {"method"}}}}}
end


function Builtin:getTypeInfo( typeName )

   local _
   
   do
      local _switchExp = typeName
      if _switchExp == "_T" then
         return Ast.builtinTypeBox:get_boxingType()
      elseif _switchExp == "_T!" then
         return Ast.builtinTypeBox:get_boxingType():get_nilableTypeInfo()
      end
   end
   
   
   local function getTypeInfoFromScope( scope, symbol, genTypeList )
   
      local typeInfo = scope:getTypeInfo( symbol, scope, false, Ast.ScopeAccess.Normal )
      if  nil == typeInfo then
         local _typeInfo = typeInfo
      
         return nil
      end
      
      local validGenType = false
      for __index, genType in ipairs( genTypeList ) do
         if genType:get_kind() ~= Ast.TypeInfoKind.Alternate then
            validGenType = true
            break
         end
         
      end
      
      if validGenType then
         do
            local _switchExp = typeInfo:get_kind()
            if _switchExp == Ast.TypeInfoKind.Map then
               if #genTypeList ~= 2 then
                  Util.err( string.format( "illegal map param -- %d", #genTypeList) )
               end
               
               local keyType = genTypeList[1]
               local valType = genTypeList[2]
               return self.processInfo:createMap( Ast.AccessMode.Pub, typeInfo:get_parentInfo(), keyType, valType, typeInfo:get_mutMode() )
            elseif _switchExp == Ast.TypeInfoKind.Ext then
               if #genTypeList ~= 1 then
                  Util.err( string.format( "illegal param -- %d", #genTypeList) )
               end
               
               do
                  local _matchExp = self.processInfo:createLuaval( genTypeList[1], true )
                  if _matchExp[1] == Ast.LuavalResult.OK[1] then
                     local workType = _matchExp[2][1]
                     local _ = _matchExp[2][2]
                  
                     if self.ctrl_info.validLuaval then
                        return workType
                     end
                     
                     return genTypeList[1]
                  elseif _matchExp[1] == Ast.LuavalResult.Err[1] then
                     local mess = _matchExp[2][1]
                  
                     Util.err( mess )
                  end
               end
               
            elseif _switchExp == Ast.TypeInfoKind.DDD then
               if #genTypeList ~= 1 then
                  Util.err( string.format( "illegal map param -- %d", #genTypeList) )
               end
               
               return self.processInfo:createDDD( genTypeList[1], true, false )
            else 
               
                  Util.err( string.format( "not support type -- %s", typeInfo:getTxt(  )) )
            end
         end
         
      end
      
      
      return typeInfo
   end
   
   local mutable = true
   if typeName:find( "^&" ) then
      mutable = false
      typeName = typeName:gsub( "^&", "" )
   end
   
   local genTypeList = {}
   local _3803, endIndex = typeName:find( "[%w%.]+<" )
   local suffix = ""
   if endIndex ~= nil then
      local genTypeName = typeName:sub( endIndex + 1 )
      while true do
         do
            local tailIndex = genTypeName:find( "[,>]" )
            if tailIndex ~= nil then
               local genType = self:getTypeInfo( genTypeName:sub( 1, tailIndex - 1 ) )
               table.insert( genTypeList, genType )
               genTypeName = genTypeName:sub( tailIndex + 1 )
            else
               suffix = genTypeName:sub( 1 )
               break
            end
         end
         
      end
      
      typeName = typeName:sub( 1, endIndex - 1 ) .. suffix
   end
   
   
   local typeInfo = Ast.headTypeInfo
   if typeName:find( "!$" ) then
      local orgTypeName = typeName:gsub( "!$", "" )
      do
         local _exp = getTypeInfoFromScope( self.transUnitIF:get_scope(), orgTypeName, genTypeList )
         if _exp ~= nil then
            typeInfo = _exp
         else
            Util.err( string.format( "not found builtin -- %s", orgTypeName) )
         end
      end
      
      typeInfo = typeInfo:get_nilableTypeInfo()
   else
    
      do
         local _exp = getTypeInfoFromScope( self.transUnitIF:get_scope(), typeName, genTypeList )
         if _exp ~= nil then
            typeInfo = _exp
         else
            Util.err( string.format( "not found builtin -- %s", typeName) )
         end
      end
      
   end
   
   if mutable then
      return typeInfo
   end
   
   typeInfo = self.modifier:createModifier( typeInfo, Ast.MutMode.IMut )
   return typeInfo
end


function Builtin:processField( name, fieldName, info, parentInfo )

   if self.targetLuaVer:isSupport( string.format( "%s.%s", name, fieldName) ) then
      if _lune.nilacc( info['type'], nil, 'item', 1) == "var" then
         local symbol = _lune.unwrap( self.transUnitIF:get_scope():add( self.processInfo, Ast.SymbolKind.Var, false, true, fieldName, nil, self:getTypeInfo( _lune.unwrap( _lune.nilacc( info['typeInfo'], nil, 'item', 1)) ), Ast.AccessMode.Pub, true, Ast.MutMode.Mut, true, false ))
         setupBuiltinTypeInfo( name, fieldName, symbol )
      else
       
         local argTypeList = {}
         for __index, argType in ipairs( _lune.unwrap( info["arg"]) ) do
            table.insert( argTypeList, self:getTypeInfo( argType ) )
         end
         
         
         local retTypeList = {}
         for __index, retType in ipairs( _lune.unwrap( info["ret"]) ) do
            local retTypeInfo = self:getTypeInfo( retType )
            table.insert( retTypeList, retTypeInfo )
         end
         
         
         local funcType = _lune.nilacc( info['type'], nil, 'item', 1)
         local mutable
         
         
         local staticFlag
         
         local kind
         
         local symbolKind
         
         local abstractFlag
         
         local accessMode
         
         do
            local _switchExp = funcType
            if _switchExp == "method" or _switchExp == "mut" then
               staticFlag = false
               kind = Ast.TypeInfoKind.Method
               symbolKind = Ast.SymbolKind.Mtd
               abstractFlag = false
               accessMode = Ast.AccessMode.Pub
               mutable = funcType == "mut"
            elseif _switchExp == "abstract" then
               staticFlag = false
               kind = Ast.TypeInfoKind.Method
               symbolKind = Ast.SymbolKind.Mtd
               abstractFlag = true
               accessMode = Ast.AccessMode.Pro
               mutable = true
            elseif _switchExp == "macro" then
               staticFlag = true
               kind = Ast.TypeInfoKind.Macro
               symbolKind = Ast.SymbolKind.Mac
               abstractFlag = false
               accessMode = Ast.AccessMode.Pub
               mutable = false
            elseif _switchExp == "formfunc" then
               staticFlag = true
               kind = Ast.TypeInfoKind.FormFunc
               symbolKind = Ast.SymbolKind.Typ
               abstractFlag = false
               accessMode = Ast.AccessMode.Pub
               mutable = false
            else 
               
                  staticFlag = true
                  kind = Ast.TypeInfoKind.Func
                  symbolKind = Ast.SymbolKind.Fun
                  abstractFlag = false
                  accessMode = Ast.AccessMode.Pub
                  mutable = false
            end
         end
         
         
         self.transUnitIF:pushScope( false )
         
         local typeInfo = self.processInfo:createFunc( abstractFlag, true, self.transUnitIF:get_scope(), kind, parentInfo, false, true, staticFlag, accessMode, fieldName, nil, argTypeList, retTypeList, mutable )
         
         self.transUnitIF:popScope(  )
         
         builtinFunc:get_allFuncTypeSet()[typeInfo]= true
         
         Ast.addBuiltin( typeInfo )
         if typeInfo:get_nilableTypeInfo() ~= Ast.headTypeInfo then
            Ast.addBuiltin( typeInfo:get_nilableTypeInfo() )
         end
         
         local symInfo = _lune.unwrap( self.transUnitIF:get_scope():add( self.processInfo, symbolKind, false, kind == Ast.TypeInfoKind.Func, fieldName, nil, typeInfo, accessMode, staticFlag, mutable and Ast.MutMode.Mut or Ast.MutMode.IMut, true, false ))
         
         setupBuiltinTypeInfo( name, fieldName, symInfo )
      end
      
   end
   
end


local readyBuiltin = false
function Builtin:registBuiltInScope(  )

   if readyBuiltin then
      return builtinFunc
   end
   
   readyBuiltin = true
   
   local builtInInfo = getBuiltInInfo(  )
   
   local builtinModuleName2Scope = {}
   
   local mapType = self.processInfo:createMap( Ast.AccessMode.Pub, Ast.headTypeInfo, Ast.builtinTypeString, Ast.builtinTypeStem, Ast.MutMode.Mut )
   self.transUnitIF:get_scope():addVar( self.processInfo, Ast.AccessMode.Global, "_ENV", nil, mapType, Ast.MutMode.IMutRe, true )
   self.transUnitIF:get_scope():addVar( self.processInfo, Ast.AccessMode.Global, "_G", nil, mapType, Ast.MutMode.IMutRe, true )
   self.transUnitIF:get_scope():addVar( self.processInfo, Ast.AccessMode.Global, "__line__", nil, Ast.builtinTypeInt, Ast.MutMode.IMut, true )
   local function processCopyAlterList( alterList, typeList )
   
      for __index, typeInfo in ipairs( typeList ) do
         table.insert( alterList, _lune.unwrap( _lune.__Cast( typeInfo, 3, Ast.AlternateTypeInfo )) )
      end
      
   end
   
   for __index, builtinClassInfo in ipairs( builtInInfo ) do
      for className, name2FieldInfo in pairs( builtinClassInfo ) do
         local name = className
         local genTypeList = {}
         
         do
            local _switchExp = className
            if _switchExp == "List<T>" then
               name = "List"
               processCopyAlterList( genTypeList, Ast.builtinTypeList:get_itemTypeInfoList() )
            elseif _switchExp == "Array<T>" then
               name = "Array"
               processCopyAlterList( genTypeList, Ast.builtinTypeArray:get_itemTypeInfoList() )
            elseif _switchExp == "Set<T>" then
               name = "Set"
               processCopyAlterList( genTypeList, Ast.builtinTypeSet:get_itemTypeInfoList() )
            elseif _switchExp == "Nilable<_T>" then
               name = "Nilable"
               processCopyAlterList( genTypeList, Ast.builtinTypeBox:get_itemTypeInfoList() )
            else 
               
                  if className:find( "<" ) then
                     name = ""
                     for token in className:gmatch( "[^<>,%s]+" ) do
                        if #name == 0 then
                           name = token
                        else
                         
                           table.insert( genTypeList, self.processInfo:createAlternate( true, #genTypeList + 1, token, Ast.AccessMode.Pri, Ast.headTypeInfo ) )
                        end
                        
                     end
                     
                  end
                  
            end
         end
         
         local parentInfo = Ast.headTypeInfo
         if name ~= "" then
            local classKind = TransUnitIF.DeclClassMode.Class
            do
               local types = _lune.nilacc( name2FieldInfo['__attrib'], nil, 'item', 'type')
               if types ~= nil then
                  if #types > 0 then
                     do
                        local _switchExp = types[1]
                        if _switchExp == "interface" then
                           classKind = TransUnitIF.DeclClassMode.Interface
                        elseif _switchExp == "module" then
                           classKind = TransUnitIF.DeclClassMode.Module
                        end
                     end
                     
                  end
                  
               end
            end
            
            local interfaceList = {}
            do
               local _exp = _lune.nilacc( name2FieldInfo['__attrib'], nil, 'item', 'inplements')
               if _exp ~= nil then
                  for __index, ifname in ipairs( _exp ) do
                     local ifType = self:getTypeInfo( ifname )
                     table.insert( interfaceList, ifType )
                  end
                  
               end
            end
            
            do
               local _switchExp = classKind
               if _switchExp == TransUnitIF.DeclClassMode.Class or _switchExp == TransUnitIF.DeclClassMode.Interface then
                  local declMode
                  
                  if classKind == TransUnitIF.DeclClassMode.Class then
                     declMode = TransUnitIF.DeclClassMode.Class
                  else
                   
                     declMode = TransUnitIF.DeclClassMode.Interface
                  end
                  
                  parentInfo = self.transUnitIF:pushClass( self.processInfo, self.transUnitIF:getLatestPos(  ), declMode, false, nil, interfaceList, genTypeList, true, name, true, Ast.AccessMode.Pub )
                  builtinFunc:registerClass( parentInfo )
               elseif _switchExp == TransUnitIF.DeclClassMode.Module then
                  parentInfo = self.transUnitIF:pushModule( self.processInfo, true, name, true )
                  
                  self.transUnitIF:get_scope():get_parent():add( self.processInfo, Ast.SymbolKind.Typ, false, false, name, nil, parentInfo, Ast.AccessMode.Local, true, Ast.MutMode.Mut, true, false )
               end
            end
            
            
            Ast.addBuiltin( parentInfo )
            Ast.addBuiltin( parentInfo:get_nilableTypeInfo() )
         end
         
         if not builtinModuleName2Scope[name] then
            if name ~= "" and self:getTypeInfo( name ) then
               builtinModuleName2Scope[name] = self.transUnitIF:get_scope()
            end
            
            do
               local __sorted = {}
               local __map = name2FieldInfo
               for __key in pairs( __map ) do
                  table.insert( __sorted, __key )
               end
               table.sort( __sorted )
               for __index, fieldName in ipairs( __sorted ) do
                  local info = __map[ fieldName ]
                  do
                     do
                        local _switchExp = fieldName
                        if _switchExp == "__attrib" then
                        else 
                           
                              self:processField( name, fieldName, info, parentInfo )
                        end
                     end
                     
                  end
               end
            end
            
         end
         
         if name ~= "" then
            self.transUnitIF:popClass(  )
         end
         
      end
      
   end
   
   local threadSafeSet = {[builtinFunc.lns_error] = true, [builtinFunc.lns_print] = true, [builtinFunc.lns_type] = true, [builtinFunc.lns_tonumber] = true, [builtinFunc.io_open] = true, [builtinFunc.set_has] = true, [builtinFunc.set_add] = true}
   
   for typeInfo, __val in pairs( builtinFunc:get_allFuncTypeSet() ) do
      if not _lune._Set_has(threadSafeSet, typeInfo ) then
         builtinFunc:get_needThreadingTypes()[typeInfo]= true
      end
      
   end
   
   
   return builtinFunc
end


return _moduleObj
