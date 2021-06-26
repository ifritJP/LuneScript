package = "lunescript51"
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
"lua = 5.1"
}
build = {
   type = "builtin",
   modules = {
      [ "lune.base.runtime" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/runtime.lua",
      [ "lune.base.runtime5" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/runtime5.lua",
      [ "lune.base.base" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/base.lua",
      [ "lune.base.Str" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Str.lua",
      [ "lune.base.Types" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Types.lua",
      [ "lune.base.Async" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Async.lua",
      [ "lune.base.Runner" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Runner.lua",
      [ "lune.base.Ver" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Ver.lua",
      [ "lune.base.Meta" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Meta.lua",
      [ "lune.base.LuneControl" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/LuneControl.lua",
      [ "lune.base.LuaMod" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/LuaMod.lua",
      [ "lune.base.Testing" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Testing.lua",
      [ "lune.base.Code" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Code.lua",
      [ "lune.base.Log" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Log.lua",
      [ "lune.base.Depend" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Depend.lua",
      [ "lune.base.Util" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Util.lua",
      [ "lune.base.LuaVer" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/LuaVer.lua",
      [ "lune.base.frontInterface" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/frontInterface.lua",
      [ "lune.base.Writer" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Writer.lua",
      [ "lune.base.AsyncParser" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/AsyncParser.lua",
      [ "lune.base.Parser" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Parser.lua",
      [ "lune.base.Json" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Json.lua",
      [ "lune.base.Option" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Option.lua",
      [ "lune.base.GoMod" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/GoMod.lua",
      [ "lune.base.Ast" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Ast.lua",
      [ "lune.base.Nodes" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Nodes.lua",
      [ "lune.base.NodeIndexer" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/NodeIndexer.lua",
      [ "lune.base.Macro" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Macro.lua",
      [ "lune.base.TransUnitIF" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/TransUnitIF.lua",
      [ "lune.base.BuiltinTransUnit" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/BuiltinTransUnit.lua",
      [ "lune.base.Builtin" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Builtin.lua",
      [ "lune.base.AstInfo" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/AstInfo.lua",
      [ "lune.base.Import" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Import.lua",
      [ "lune.base.TransUnit" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/TransUnit.lua",
      [ "lune.base.convLua" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/convLua.lua",
      [ "lune.base.DependLuaOnLns" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/DependLuaOnLns.lua",
      [ "lune.base.convGo" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/convGo.lua",
      [ "lune.base.dumpNode" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/dumpNode.lua",
      [ "lune.base.Formatter" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/Formatter.lua",
      [ "lune.base.OutputDepend" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/OutputDepend.lua",
      [ "lune.base.front" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/front.lua",
      [ "lune.base.glueFilter" ] = "src/lune/base/runtime_go/legacy/lua51/lune/base/glueFilter.lua",
   },
   install = {
      bin = {
         lnsc = "src/lnsc.lua"
      }
   },
   copy_directories = { "docs" }
}
