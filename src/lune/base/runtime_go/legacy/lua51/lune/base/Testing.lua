--lune/base/Testing.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Testing'
local _lune = {}
if _lune6 then
   _lune = _lune6
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
local Result = {}
_moduleObj.Result = Result
function Result:outputResult( stream )

   stream:write( string.format( "test total: %s %d (OK:%d, NG:%d)\n", self.name, self.okNum + self.ngNum, self.okNum, self.ngNum) )
end
function Result:err( mess, mod, lineNo )

   self.ngNum = self.ngNum + 1
   io.stderr:write( string.format( "error: %s:%d: %s\n", mod, lineNo, mess) )
end
function Result:isTrue( val1, val1txt, msg, mod, lineNo )

   if val1 == true then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "not true -- %s:%s:[%s]\n", msg or "", val1txt, tostring( val1)), mod, lineNo )
   return false
end
function Result:isNotTrue( val1, val1txt, msg, mod, lineNo )

   if not val1 then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "is true -- %s:%s:[%s]\n", msg or "", val1txt, tostring( val1)), mod, lineNo )
   return false
end
function Result:isNil( val1, val1txt, msg, mod, lineNo )

   if val1 == nil then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "is not nil -- %s:%s:[%s]\n", msg or "", val1txt, tostring( val1)), mod, lineNo )
   return false
end
function Result:isNotNil( val1, val1txt, msg, mod, lineNo )

   if val1 ~= nil then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "is nil -- %s:%s:[%s]\n", msg or "", val1txt, tostring( val1)), mod, lineNo )
   return false
end
function Result:checkEq( val1, val2, val1txt, val2txt, msg, mod, lineNo )

   if val1 == val2 then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "not equal -- %s:%s:[%s] != %s:[%s]\n", msg or "", val1txt, tostring( val1), val2txt, tostring( val2)), mod, lineNo )
   return false
end
function Result:checkNotEq( val1, val2, val1txt, val2txt, msg, mod, lineNo )

   if val1 ~= val2 then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "equal -- %s:%s:[%s] == %s:[%s]\n", msg or "", val1txt, tostring( val1), val2txt, tostring( val2)), mod, lineNo )
   return false
end
function Result.setmeta( obj )
  setmetatable( obj, { __index = Result  } )
end
function Result.new( name, okNum, ngNum )
   local obj = {}
   Result.setmeta( obj )
   if obj.__init then
      obj:__init( name, okNum, ngNum )
   end
   return obj
end
function Result:__init( name, okNum, ngNum )

   self.name = name
   self.okNum = okNum
   self.ngNum = ngNum
end


local Ctrl = {}
_moduleObj.Ctrl = Ctrl
function Ctrl.setmeta( obj )
  setmetatable( obj, { __index = Ctrl  } )
end
function Ctrl.new( result )
   local obj = {}
   Ctrl.setmeta( obj )
   if obj.__init then
      obj:__init( result )
   end
   return obj
end
function Ctrl:__init( result )

   self.result = result
end
function Ctrl:get_result()
   return self.result
end
function Ctrl:checkEq( ... )
   return self.result:checkEq( ... )
end

function Ctrl:checkNotEq( ... )
   return self.result:checkNotEq( ... )
end

function Ctrl:err( ... )
   return self.result:err( ... )
end

function Ctrl:isNil( ... )
   return self.result:isNil( ... )
end

function Ctrl:isNotNil( ... )
   return self.result:isNotNil( ... )
end

function Ctrl:isNotTrue( ... )
   return self.result:isNotTrue( ... )
end

function Ctrl:isTrue( ... )
   return self.result:isTrue( ... )
end

function Ctrl:outputResult( ... )
   return self.result:outputResult( ... )
end




local TestCase = {}
function TestCase.setmeta( obj )
  setmetatable( obj, { __index = TestCase  } )
end
function TestCase.new( func, result )
   local obj = {}
   TestCase.setmeta( obj )
   if obj.__init then
      obj:__init( func, result )
   end
   return obj
end
function TestCase:__init( func, result )

   self.func = func
   self.result = result
end
function TestCase:get_func()
   return self.func
end
function TestCase:get_result()
   return self.result
end


local TestModuleInfo = {}
function TestModuleInfo.new( name )
   local obj = {}
   TestModuleInfo.setmeta( obj )
   if obj.__init then obj:__init( name ); end
   return obj
end
function TestModuleInfo:__init(name) 
   self.runned = false
   self.name = name
   self.testcaseMap = {}
end
function TestModuleInfo:addCase( name, testCase )

   self.testcaseMap[name] = testCase
end
function TestModuleInfo:run(  )

   self.runned = true
   print( string.format( "module: %s %s", self.name, string.rep( "=", 30 )) )
   do
      local __sorted = {}
      local __map = self.testcaseMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, name in ipairs( __sorted ) do
         local testcase = __map[ name ]
         do
            print( string.format( "%s: %s", name, string.rep( "-", 15 )) )
            testcase:get_func()( Ctrl.new(testcase:get_result()) )
         end
      end
   end
   
end
function TestModuleInfo:outputResult( stream )

   if not self.runned then
      return 
   end
   
   print( string.format( "module: %s %s", self.name, string.rep( "=", 30 )) )
   do
      local __sorted = {}
      local __map = self.testcaseMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local testcase = __map[ __key ]
         do
            testcase:get_result():outputResult( stream )
         end
      end
   end
   
end
function TestModuleInfo.setmeta( obj )
  setmetatable( obj, { __index = TestModuleInfo  } )
end
function TestModuleInfo:get_runned()
   return self.runned
end
function TestModuleInfo:get_name()
   return self.name
end
function TestModuleInfo:get_testcaseMap()
   return self.testcaseMap
end

local testModuleMap = {}

local function registerTestcase( modName, caseName, testcase )

   local info = testModuleMap[modName]
   if  nil == info then
      local _info = info
   
      info = TestModuleInfo.new(modName)
      testModuleMap[modName] = info
   end
   
   local result = Result.new(caseName, 0, 0)
   info:addCase( caseName, TestCase.new(testcase, result) )
end
_moduleObj.registerTestcase = registerTestcase

local function run( modPath )

   do
      local __sorted = {}
      local __map = testModuleMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, name in ipairs( __sorted ) do
         local info = __map[ name ]
         do
            if name == modPath then
               info:run(  )
            end
            
         end
      end
   end
   
end
_moduleObj.run = run
local function runAll(  )

   do
      local __sorted = {}
      local __map = testModuleMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local info = __map[ __key ]
         do
            info:run(  )
         end
      end
   end
   
end
_moduleObj.runAll = runAll

local function outputAllResult( stream )

   do
      local __sorted = {}
      local __map = testModuleMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local info = __map[ __key ]
         do
            info:outputResult( stream )
         end
      end
   end
   
end
_moduleObj.outputAllResult = outputAllResult





























return _moduleObj
