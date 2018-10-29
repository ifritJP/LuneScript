local lua = arg[-1]

local dir = arg[0]:gsub( "setup.lua$", "" )
if dir == "" then
   dir = "./"
end

local target = "install"
if arg[1] == "uninstall" then
   target = "uninstall"
end

local process = io.popen( string.format( "which %s", arg[ -1 ] ) )

lua = process:read( '*l' ):gsub( '\n', '' )
process:close()

local command = string.format( "make -C %s LUA=%s %s", dir, lua, target ) 
print( command )

os.execute( command )
