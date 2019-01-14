package = "lunescript"
version = "main-1"
source = {
   url = "git://github.com/ifritJP/LuneScript"
}
description = {
   summary = "LuneScript is nil safety trans-compiler for Lua.",
   detailed = [[
LuneScript is a Transcompiler for Lua.

- Learning cost is low because it is based on Lua and C syntax.
- Because LuneScript is a statically typed language, simple checks can be found at compile time by type checking.
- Minimize the effort of type declaration by type inference.
- Null safety.
- Generics (partly only) allows processing while preserving type information.
- Corresponds to class definition as grammar of language.
- Macros can realize designs that do not depend on dynamic processing such as polymorphism.
- Supports data representation compatible with JSON.
- Transformed Lua code can be operated as a single unit without assuming external libraries.
- Since the process written in LuneScript is output as is, the transcoded Lua code, There is no performance deterioration.
- Existing Lua external modules can be used from LuneScript.
- LuneScript runs on Lua and is easy to install as it requires no Lua standard modules.
- Supports Lua 5.1, 5.2, 5.3.
- LuneScript is developed by self hosting.
]],
   homepage = "https://github.com/ifritJP/LuneScript.git",
   license = "MIT"
}
dependencies = {
"lua >= 5.2, < 5.4"
}
build = {
   type = "builtin",
   modules = {
      [ "lune.base.base" ] = "src/lune/base/base.lua",
      [ "lune.base.Ver" ] = "src/lune/base/Ver.lua",
      [ "lune.base.LuaMod" ] = "src/lune/base/LuaMod.lua",
      [ "lune.base.Depend" ] = "src/lune/base/Depend.lua",
      [ "lune.base.Util" ] = "src/lune/base/Util.lua",
      [ "lune.base.LuaVer" ] = "src/lune/base/LuaVer.lua",
      [ "lune.base.frontInterface" ] = "src/lune/base/frontInterface.lua",
      [ "lune.base.Writer" ] = "src/lune/base/Writer.lua",
      [ "lune.base.Parser" ] = "src/lune/base/Parser.lua",
      [ "lune.base.Option" ] = "src/lune/base/Option.lua",
      [ "lune.base.Ast" ] = "src/lune/base/Ast.lua",
      [ "lune.base.TransUnit" ] = "src/lune/base/TransUnit.lua",
      [ "lune.base.convLua" ] = "src/lune/base/convLua.lua",
      [ "lune.base.dumpNode" ] = "src/lune/base/dumpNode.lua",
      [ "lune.base.OutputDepend" ] = "src/lune/base/OutputDepend.lua",
      [ "lune.base.front" ] = "src/lune/base/front.lua",
      [ "lune.base.glueFilter" ] = "src/lune/base/glueFilter.lua"
   },
   install = {
      bin = {
         lnsc = "src/lnsc.lua"
      }
   },
   copy_directories = { "docs" }
}
