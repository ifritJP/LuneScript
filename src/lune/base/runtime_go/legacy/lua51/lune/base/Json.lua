--lune/base/Json.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Json'
local _lune = {}
if _lune5 then
   _lune = _lune5
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

if not _lune5 then
   _lune5 = _lune
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

   local parser = Parser.DefaultPushbackParser.new(Parser.StreamParser.create( _lune.newAlge( Types.ParserSrc.LnsCode, {txt,"json",nil}), false ))
   local val, ok = getVal( parser )
   if not ok then
      return nil, parser:getLastPos(  )
   end
   
   return val, nil
end
_moduleObj.fromStr = fromStr



return _moduleObj
