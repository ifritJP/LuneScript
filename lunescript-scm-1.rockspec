package = "LuneScript"
version = "scm-1"
source = {
	url = "git://github.com/ifritJP/LuneScript"
}
description = {
	homepage = "*** please enter a project homepage ***",
	license = "MIT*"
}
dependencies = {}
build = {
	type = "builtin",
	install = {
		bin = {
			lune = "src/lune/base/base.lua"
		},
		lua = {
			["lune.base.Parser"] = "src/lune/base/Parser.lns",
			["lune.base.TransUnit"] = "src/lune/base/TransUnit.lns",
			["lune.base.Util"] = "src/lune/base/Util.lns",
			["lune.base.convLua"] = "src/lune/base/convLua.lns",
			["lune.base.dumpNode"] = "src/lune/base/dumpNode.lns",
			["primal.Parser"] = "src/primal/Parser.lua",
			["primal.TransUnit"] = "src/primal/TransUnit.lua",
			["primal.base"] = "src/primal/base.lua",
			["primal.convLua"] = "src/primal/convLua.lua",
			["primal.dumpNode"] = "src/primal/dumpNode.lua",
		}
	},
	modules = {
		["lune.base.Parser"] = "src/lune/base/Parser.lua",
		["lune.base.TransUnit"] = "src/lune/base/TransUnit.lua",
		["lune.base.Util"] = "src/lune/base/Util.lua",
		["lune.base.base"] = "src/lune/base/base.lua",
		["lune.base.convLua"] = "src/lune/base/convLua.lua",
		["lune.base.dumpNode"] = "src/lune/base/dumpNode.lua",
		["primal.Parser"] = "src/primal/Parser.lua",
		["primal.TransUnit"] = "src/primal/TransUnit.lua",
		["primal.base"] = "src/primal/base.lua",
		["primal.convLua"] = "src/primal/convLua.lua",
		["primal.dumpNode"] = "src/primal/dumpNode.lua",
	}
}
