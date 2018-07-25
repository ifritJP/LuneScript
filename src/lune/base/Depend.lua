--lune/base/Depend.lns
local moduleObj = {}
local function getFileLastModifiedTime( path )
  local file = io.open( path )
  
  do
    local _exp = file
    if _exp then
    
        _exp:close(  )
      else
    
        return nil
      end
  end
  
  local stream = io.popen( string.format( "stat -c '%%Y' %s", path) )
  
  do
    local _exp = stream
    if _exp then
    
        local val = _exp:read( '*a' )
        
        do
          local _exp = val
          if _exp then
          
              return tonumber( _exp )
            end
        end
        
      end
  end
  
  return nil
end
moduleObj.getFileLastModifiedTime = getFileLastModifiedTime
return moduleObj
