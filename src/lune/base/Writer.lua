--lune/base/Writer.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Writer'
local _lune = {}
if _lune0 then
   _lune = _lune0
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

if not _lune0 then
   _lune0 = _lune
end
local Util = _lune.loadModule( 'lune.base.Util' )
local Writer = {}
_moduleObj.Writer = Writer
function Writer.setmeta( obj )
  setmetatable( obj, { __index = Writer  } )
end
function Writer.new(  )
   local obj = {}
   Writer.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Writer:__init(  )

end

local XML = {}
setmetatable( XML, { ifList = {Writer,} } )
_moduleObj.XML = XML
function XML.new( stream )
   local obj = {}
   XML.setmeta( obj )
   if obj.__init then obj:__init( stream ); end
   return obj
end
function XML:__init(stream) 
   self.stream = stream
   self.elementList = {}
   self.depth = 0
end
function XML.convertXmlTxt( val )

   if val == "" then
      return ""
   end
   
   if type( val ) == "number" then
      return string.format( "%g", val)
   end
   
   local txt = string.format( "%s", val)
   txt = string.gsub( txt, '&', "&amp;" )
   txt = string.gsub( txt, '>', "&gt;" )
   txt = string.gsub( txt, '<', "&lt;" )
   txt = string.gsub( txt, '"', "&quot;" )
   txt = string.gsub( txt, "'", "&apos;" )
   return txt
end
function XML:startElement( name )

   table.insert( self.elementList, name )
   self.stream:write( string.format( '<%s>', name) )
   self.depth = self.depth + 1
end
function XML:startParent( name, arrayFlag )

   self:startElement( name )
end
function XML:endElement(  )

   local name = table.remove( self.elementList )
   self.stream:write( string.format( '</%s>', name) )
   self.depth = self.depth - 1
   if self.depth == 0 then
      self.stream:write( '\n' )
   elseif self.depth < 0 then
      Util.err( "illegal depth" )
   end
   
end
function XML:writeValue( val )

   self.stream:write( XML.convertXmlTxt( val ) )
end
function XML:write( name, val )

   self:startElement( name )
   self:writeValue( val )
   self:endElement(  )
end
function XML:fin(  )

end
function XML.setmeta( obj )
  setmetatable( obj, { __index = XML  } )
end

local JsonLayer = {}
function JsonLayer.setmeta( obj )
  setmetatable( obj, { __index = JsonLayer  } )
end
function JsonLayer.new( state, arrayFlag, name, madeByArrayFlag, elementNameSet, parentFlag, openElement )
   local obj = {}
   JsonLayer.setmeta( obj )
   if obj.__init then
      obj:__init( state, arrayFlag, name, madeByArrayFlag, elementNameSet, parentFlag, openElement )
   end
   return obj
end
function JsonLayer:__init( state, arrayFlag, name, madeByArrayFlag, elementNameSet, parentFlag, openElement )

   self.state = state
   self.arrayFlag = arrayFlag
   self.name = name
   self.madeByArrayFlag = madeByArrayFlag
   self.elementNameSet = elementNameSet
   self.parentFlag = parentFlag
   self.openElement = openElement
end

local JSON = {}
setmetatable( JSON, { ifList = {Writer,} } )
_moduleObj.JSON = JSON
function JSON:startLayer( arrayFlag, madeByArrayFlag )

   local info = JsonLayer.new('none', arrayFlag, self.prevName, madeByArrayFlag, {}, true, false)
   table.insert( self.layerQueue, info )
   self.stream:write( arrayFlag and "[" or "{" )
end
function JSON.new( stream )
   local obj = {}
   JSON.setmeta( obj )
   if obj.__init then obj:__init( stream ); end
   return obj
end
function JSON:__init(stream) 
   self.stream = stream
   self.layerQueue = {}
   self.prevName = ""
   self:startLayer( false, false )
end
function JSON:getLayerInfo(  )

   if #self.layerQueue == 0 then
      return nil
   end
   
   return self.layerQueue[#self.layerQueue]
end
function JSON:endLayer(  )

   if #self.layerQueue == 0 then
      Util.err( "illegal depth" )
   end
   
   while #self.layerQueue > 0 do
      local info = _lune.unwrap( self:getLayerInfo(  ))
      if info.arrayFlag then
         self.stream:write( ']' )
      else
       
         self.stream:write( '}' )
      end
      
      table.remove( self.layerQueue )
      local parentInfo = self:getLayerInfo(  )
      if not _lune.nilacc( parentInfo, "madeByArrayFlag" ) then
         break
      end
      
   end
   
   if #self.layerQueue == 0 then
      self.stream:write( '\n' )
   end
   
end
function JSON:equalLayerState( state )

   return self.layerQueue[#self.layerQueue].state == state
end
function JSON:isArrayLayer(  )

   return self.layerQueue[#self.layerQueue].arrayFlag
end
function JSON:setLayerState( state )

   self.layerQueue[#self.layerQueue].state = state
end
function JSON:getLayerName(  )

   return self.layerQueue[#self.layerQueue].name
end
function JSON:addElementName( name )

   local info = _lune.unwrap( self:getLayerInfo(  ))
   local nameSet = info.elementNameSet
   if not info.arrayFlag and _lune._Set_has(nameSet, name ) then
      Util.err( "exist same name: " .. name )
   end
   
   nameSet[name]= true
end
function JSON:startParent( name, arrayFlag )

   self:addElementName( name )
   if self:equalLayerState( 'termed' ) or self:equalLayerState( 'named' ) or self:equalLayerState( 'valued' ) then
      self.stream:write( "," )
   elseif self:equalLayerState( 'none' ) then
      
   end
   
   local parentInfo = self:getLayerInfo(  )
   if not arrayFlag and _lune.nilacc( parentInfo, "arrayFlag" ) then
      self:startLayer( false, true )
   end
   
   self.stream:write( string.format( '"%s": ', name) )
   self:startLayer( arrayFlag, false )
   self.prevName = name
end
function JSON:startElement( name )

   self:addElementName( name )
   if self:equalLayerState( 'termed' ) then
      self.stream:write( "," )
   elseif self:equalLayerState( 'named' ) then
      Util.err( 'illegal layer state' )
   elseif self:equalLayerState( 'none' ) then
      
   end
   
   if self:isArrayLayer(  ) then
      self:startLayer( false, true )
   end
   
   local info = _lune.unwrap( self:getLayerInfo(  ))
   if info.openElement then
      Util.err( 'illegal openElement' )
   end
   
   info.openElement = true
   self.stream:write( string.format( '"%s": ', name) )
   self:setLayerState( 'named' )
   self.prevName = name
end
function JSON:endElement(  )

   if self:equalLayerState( 'none' ) or self:equalLayerState( 'termed' ) then
      self:endLayer(  )
   elseif self:equalLayerState( 'valued' ) then
      local info = _lune.unwrap( self:getLayerInfo(  ))
      if info.openElement then
         info.openElement = false
      end
      
      if _lune.nilacc( self:getLayerInfo(  ), "madeByArrayFlag" ) then
         self:endLayer(  )
      end
      
   else
    
      Util.err( string.format( 'illegal layer state %s', self:getLayerName(  )) )
   end
   
   self:setLayerState( 'termed' )
end
function JSON.convertJsonTxt( txt )

   if txt == "" then
      return ""
   end
   
   txt = string.gsub( txt, '"', '\\"' )
   txt = string.gsub( txt, '\\', '\\\\' )
   txt = string.gsub( txt, '\n', '\\n' )
   return txt
end
function JSON:writeValue( val )

   local txt = ""
   local typeId = type( val )
   if typeId == "number" then
      txt = string.format( "%g", val)
   elseif typeId == "boolean" then
      txt = val and "true" or "false"
   else
    
      txt = string.format( '"%s"', JSON.convertJsonTxt( string.format( '%s', val) ))
   end
   
   self.stream:write( txt )
   self:setLayerState( 'valued' )
end
function JSON:write( name, val )

   self:startElement( name )
   self:writeValue( val )
   self:endElement(  )
end
function JSON:fin(  )

   if self:equalLayerState( 'none' ) or self:equalLayerState( 'termed' ) then
      self:endLayer(  )
   else
    
      Util.err( 'illegal' )
   end
   
end
function JSON.setmeta( obj )
  setmetatable( obj, { __index = JSON  } )
end

return _moduleObj
