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
      [ "lune.base.runtime" ] = "src/lune/base/runtime.lua",
      [ "lune.base.runtime3" ] = "src/lune/base/runtime3.lua",
      [ "lune.base.base" ] = "src/lune/base/base.lua",
      [ "lune.base.Str" ] = "src/lune/base/Str.lua",
      [ "lune.base.Types" ] = "src/lune/base/Types.lua",
      [ "lune.base.Async" ] = "src/lune/base/Async.lua",
      [ "lune.base.Runner" ] = "src/lune/base/Runner.lua",
      [ "lune.base.Ver" ] = "src/lune/base/Ver.lua",
      [ "lune.base.Meta" ] = "src/lune/base/Meta.lua",
      [ "lune.base.LuneControl" ] = "src/lune/base/LuneControl.lua",
      [ "lune.base.LuaMod" ] = "src/lune/base/LuaMod.lua",
      [ "lune.base.Testing" ] = "src/lune/base/Testing.lua",
      [ "lune.base.Code" ] = "src/lune/base/Code.lua",
      [ "lune.base.Log" ] = "src/lune/base/Log.lua",
      [ "lune.base.Depend" ] = "src/lune/base/Depend.lua",
      [ "lune.base.Util" ] = "src/lune/base/Util.lua",
      [ "lune.base.LuaVer" ] = "src/lune/base/LuaVer.lua",
      [ "lune.base.frontInterface" ] = "src/lune/base/frontInterface.lua",
      [ "lune.base.Writer" ] = "src/lune/base/Writer.lua",
      [ "lune.base.AsyncParser" ] = "src/lune/base/AsyncParser.lua",
      [ "lune.base.Parser" ] = "src/lune/base/Parser.lua",
      [ "lune.base.Json" ] = "src/lune/base/Json.lua",
      [ "lune.base.Option" ] = "src/lune/base/Option.lua",
      [ "lune.base.GoMod" ] = "src/lune/base/GoMod.lua",
      [ "lune.base.Ast" ] = "src/lune/base/Ast.lua",
      [ "lune.base.Nodes" ] = "src/lune/base/Nodes.lua",
      [ "lune.base.Macro" ] = "src/lune/base/Macro.lua",
      [ "lune.base.TransUnitIF" ] = "src/lune/base/TransUnitIF.lua",
      [ "lune.base.SimpleTransUnit" ] = "src/lune/base/SimpleTransUnit.lua",
      [ "lune.base.Builtin" ] = "src/lune/base/Builtin.lua",
      [ "lune.base.Import" ] = "src/lune/base/Import.lua",
      [ "lune.base.TransUnit" ] = "src/lune/base/TransUnit.lua",
      [ "lune.base.DependLuaOnLns" ] = "src/lune/base/DependLuaOnLns.lua",
      [ "lune.base.convLua" ] = "src/lune/base/convLua.lua",
      [ "lune.base.convGo" ] = "src/lune/base/convGo.lua",
      [ "lune.base.dumpNode" ] = "src/lune/base/dumpNode.lua",
      [ "lune.base.Formatter" ] = "src/lune/base/Formatter.lua",
      [ "lune.base.OutputDepend" ] = "src/lune/base/OutputDepend.lua",
      [ "lune.base.front" ] = "src/lune/base/front.lua",
      [ "lune.base.glueFilter" ] = "src/lune/base/glueFilter.lua",
   },
   install = {
      bin = {
         lnsc = "src/lnsc.lua"
      }
   },
   copy_directories = { "docs" }
}
