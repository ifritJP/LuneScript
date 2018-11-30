local version = tonumber( _VERSION:gsub( "^[^%d]+", "" ), nil )
if version < 5.2 then
   io.stderr:write(
      string.format(
	 "LuneScript doesn't support this lua version(%s). %s\n",
	 version, "please use the version >= 5.2." ) )
   os.exit( 1 )
end

local lua = arg[-1]

local dir = arg[0]:gsub( "setup.lua$", "" )
if dir == "" then
   dir = "./"
end

local modDirList = { "<cancel>" }
for path in string.gmatch( package.path, "[^;]+" ) do
   if path:find( "/%?%.lua$" ) and not path:find( "^%./" ) then
      table.insert( modDirList, ( path:gsub( "/%?%.lua$", "" ) ) )
   end
end

for index, path in ipairs( modDirList ) do
   print( index, path )
end

if #modDirList == 1 then
   print( "it not find lua module directory!" )
   os.exit( 1 )
end

io.stdout:write( string.format( "install path? [default 2]: ") )

local dirNum = ""
if arg[ 1 ] ~= "-d" then
   dirNum = io.stdin:read( '*l' )
end
if dirNum == "" then
   dirNum = "2"
end
local modDir = modDirList[ tonumber( dirNum ) ]
if not modDir or modDir == "<cancel>" then
   print( "canceled" )
   os.exit( 1 )
end

local process = io.popen( string.format( "which %s", arg[ -1 ] ) )

lua = process:read( '*l' ):gsub( '\n', '' )
process:close()

local mkobj = io.open( "lune.mk", "w" )
mkobj:write( string.format( "LUA=%s\nLUA_MOD_DIR=%s\n", lua, modDir ) )
mkobj:close()

print( [[

generated lune.mk file.
please execute following command.

$ sudo make install
    or 
$ sudo make uninstall]] )
