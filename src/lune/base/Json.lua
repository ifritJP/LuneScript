--lune/base/Json.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Json'
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
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Types = _lune.loadModule( 'lune.base.Types' )

local function getRawTxt( token )

   return token.txt:sub( 2, -2 )
end

local function getVal( parser )

   local token = parser:getTokenNoErr(  )
   do
      local _switchExp = token.kind
      if _switchExp == Parser.TokenKind.Dlmt then
         do
            local _switchExp = token.txt
            if _switchExp == "{" then
               local map = {}
               while true do
                  local key = parser:getTokenNoErr(  )
                  do
                     local _switchExp = key.kind
                     if _switchExp == Parser.TokenKind.Str then
                     elseif _switchExp == Parser.TokenKind.Dlmt then
                        do
                           local _switchExp = key.txt
                           if _switchExp == "}" then
                              return map, true
                           elseif _switchExp == "," then
                              key = parser:getTokenNoErr(  )
                              if key.kind ~= Parser.TokenKind.Str then
                                 return nil, false
                              end
                              
                           else 
                              
                                 return nil, false
                           end
                        end
                        
                     else 
                        
                           return nil, false
                     end
                  end
                  
                  if parser:getTokenNoErr(  ).txt ~= ":" then
                     return nil, false
                  end
                  
                  local val, ok = getVal( parser )
                  if not ok then
                     return nil, false
                  end
                  
                  map[getRawTxt( key )] = val
               end
               
            elseif _switchExp == "[" then
               local list = {}
               local count = 1
               while true do
                  local nextToken = parser:getTokenNoErr(  )
                  do
                     local _switchExp = nextToken.txt
                     if _switchExp == "]" then
                        return list, true
                     elseif _switchExp == "," then
                     else 
                        
                           parser:pushback(  )
                     end
                  end
                  
                  local val, ok = getVal( parser )
                  if not ok then
                     return nil, false
                  end
                  
                  list[count] = val
                  count = count + 1
               end
               
            end
         end
         
         return nil, false
      elseif _switchExp == Parser.TokenKind.Int then
         local num = tonumber( token.txt )
         if  nil == num then
            local _num = num
         
            return nil, false
         end
         
         return math.floor(num), true
      elseif _switchExp == Parser.TokenKind.Real then
         local num = tonumber( token.txt )
         if  nil == num then
            local _num = num
         
            return nil, false
         end
         
         return num, true
      elseif _switchExp == Parser.TokenKind.Str then
         return getRawTxt( token ), true
      elseif _switchExp == Parser.TokenKind.Kywd then
         do
            local _switchExp = token.txt
            if _switchExp == "true" then
               return true, true
            elseif _switchExp == "false" then
               return false, true
            elseif _switchExp == "nil" or _switchExp == "null" then
               return nil, true
            end
         end
         
         return nil, false
      end
   end
   
   return nil, false
end

local function fromStr( txt )

   local stream = Parser.TxtStream.new(txt)
   local parser = Parser.DefaultPushbackParser.new(Parser.StreamParser.new(stream, "json", false, nil))
   local val, ok = getVal( parser )
   if not ok then
      return nil, parser:getLastPos(  )
   end
   
   return val, nil
end
_moduleObj.fromStr = fromStr



return _moduleObj
