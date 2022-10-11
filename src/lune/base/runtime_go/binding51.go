// +build gopherlua

package runtimelns
//import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import _ "embed"
//go:embed legacy/lua51/lune/base/Async.lua
var lune_base_Async []byte
//go:embed legacy/lua51/lune/base/Runner.lua
var lune_base_Runner []byte
//go:embed legacy/lua51/lune/base/Str.lua
var lune_base_Str []byte
//go:embed legacy/lua51/lune/base/Types.lua
var lune_base_Types []byte
//go:embed legacy/lua51/lune/base/Ver.lua
var lune_base_Ver []byte
//go:embed legacy/lua51/lune/base/LuaMod.lua
var lune_base_LuaMod []byte
//go:embed legacy/lua51/lune/base/LuneControl.lua
var lune_base_LuneControl []byte
//go:embed legacy/lua51/lune/base/Meta.lua
var lune_base_Meta []byte
//go:embed legacy/lua51/lune/base/Testing.lua
var lune_base_Testing []byte
//go:embed legacy/lua51/lune/base/Testing.lns
var lns_lune_base_Testing string
//go:embed legacy/lua51/lune/base/TestingMacro.lua
var lune_base_TestingMacro []byte
//go:embed legacy/lua51/lune/base/Code.lua
var lune_base_Code []byte
//go:embed legacy/lua51/lune/base/Log.lua
var lune_base_Log []byte
//go:embed legacy/lua51/lune/base/LuaVer.lua
var lune_base_LuaVer []byte
//go:embed legacy/lua51/lune/base/Depend.lua
var lune_base_Depend []byte
//go:embed legacy/lua51/lune/base/Util.lua
var lune_base_Util []byte
//go:embed legacy/lua51/lune/base/frontInterface.lua
var lune_base_frontInterface []byte
//go:embed legacy/lua51/lune/base/Writer.lua
var lune_base_Writer []byte
//go:embed legacy/lua51/lune/base/AsyncParser.lua
var lune_base_AsyncParser []byte
//go:embed legacy/lua51/lune/base/Parser.lua
var lune_base_Parser []byte
//go:embed legacy/lua51/lune/base/Json.lua
var lune_base_Json []byte
//go:embed legacy/lua51/lune/base/Ast.lua
var lune_base_Ast []byte
//go:embed legacy/lua51/lune/base/Option.lua
var lune_base_Option []byte
//go:embed legacy/lua51/lune/base/GoMod.lua
var lune_base_GoMod []byte
//go:embed legacy/lua51/lune/base/Nodes.lua
var lune_base_Nodes []byte
//go:embed legacy/lua51/lune/base/NodeIndexer.lua
var lune_base_NodeIndexer []byte
//go:embed legacy/lua51/lune/base/Formatter.lua
var lune_base_Formatter []byte
//go:embed legacy/lua51/lune/base/Macro.lua
var lune_base_Macro []byte
//go:embed legacy/lua51/lune/base/TransUnitIF.lua
var lune_base_TransUnitIF []byte
//go:embed legacy/lua51/lune/base/BuiltinTransUnit.lua
var lune_base_BuiltinTransUnit []byte
//go:embed legacy/lua51/lune/base/Builtin.lua
var lune_base_Builtin []byte
//go:embed legacy/lua51/lune/base/AstInfo.lua
var lune_base_AstInfo []byte
//go:embed legacy/lua51/lune/base/Import.lua
var lune_base_Import []byte
//go:embed legacy/lua51/lune/base/TransUnit.lua
var lune_base_TransUnit []byte
//go:embed legacy/lua51/lune/base/DependLuaOnLns.lua
var lune_base_DependLuaOnLns []byte
//go:embed legacy/lua51/lune/base/convLua.lua
var lune_base_convLua []byte
//go:embed legacy/lua51/lune/base/dumpNode.lua
var lune_base_dumpNode []byte
//go:embed legacy/lua51/lune/base/convGo.lua
var lune_base_convGo []byte
//go:embed legacy/lua51/lune/base/convPython.lua
var lune_base_convPython []byte
//go:embed legacy/lua51/lune/base/OutputDepend.lua
var lune_base_OutputDepend []byte
//go:embed legacy/lua51/lune/base/Converter.lua
var lune_base_Converter []byte
//go:embed legacy/lua51/lune/base/glueFilter.lua
var lune_base_glueFilter []byte
//go:embed legacy/lua51/lune/base/front.lua
var lune_base_front []byte
//go:embed legacy/lua51/lune/base/runtime7.lua
var lune_base_runtime7 []byte
func init_bind() {
AddlnsSrcInfo( "lune.base.Async", lune_base_Async )
AddlnsSrcInfo( "lune.base.Runner", lune_base_Runner )
AddlnsSrcInfo( "lune.base.Str", lune_base_Str )
AddlnsSrcInfo( "lune.base.Types", lune_base_Types )
AddlnsSrcInfo( "lune.base.Ver", lune_base_Ver )
AddlnsSrcInfo( "lune.base.LuaMod", lune_base_LuaMod )
AddlnsSrcInfo( "lune.base.LuneControl", lune_base_LuneControl )
AddlnsSrcInfo( "lune.base.Meta", lune_base_Meta )
AddlnsSrcInfo( "lune.base.Testing", lune_base_Testing )
AddlnsSrcInfo( "lune.base.TestingMacro", lune_base_TestingMacro )
AddlnsSrcInfo( "lune.base.Code", lune_base_Code )
AddlnsSrcInfo( "lune.base.Log", lune_base_Log )
AddlnsSrcInfo( "lune.base.LuaVer", lune_base_LuaVer )
AddlnsSrcInfo( "lune.base.Depend", lune_base_Depend )
AddlnsSrcInfo( "lune.base.Util", lune_base_Util )
AddlnsSrcInfo( "lune.base.frontInterface", lune_base_frontInterface )
AddlnsSrcInfo( "lune.base.Writer", lune_base_Writer )
AddlnsSrcInfo( "lune.base.AsyncParser", lune_base_AsyncParser )
AddlnsSrcInfo( "lune.base.Parser", lune_base_Parser )
AddlnsSrcInfo( "lune.base.Json", lune_base_Json )
AddlnsSrcInfo( "lune.base.Ast", lune_base_Ast )
AddlnsSrcInfo( "lune.base.Option", lune_base_Option )
AddlnsSrcInfo( "lune.base.GoMod", lune_base_GoMod )
AddlnsSrcInfo( "lune.base.Nodes", lune_base_Nodes )
AddlnsSrcInfo( "lune.base.NodeIndexer", lune_base_NodeIndexer )
AddlnsSrcInfo( "lune.base.Formatter", lune_base_Formatter )
AddlnsSrcInfo( "lune.base.Macro", lune_base_Macro )
AddlnsSrcInfo( "lune.base.TransUnitIF", lune_base_TransUnitIF )
AddlnsSrcInfo( "lune.base.BuiltinTransUnit", lune_base_BuiltinTransUnit )
AddlnsSrcInfo( "lune.base.Builtin", lune_base_Builtin )
AddlnsSrcInfo( "lune.base.AstInfo", lune_base_AstInfo )
AddlnsSrcInfo( "lune.base.Import", lune_base_Import )
AddlnsSrcInfo( "lune.base.TransUnit", lune_base_TransUnit )
AddlnsSrcInfo( "lune.base.DependLuaOnLns", lune_base_DependLuaOnLns )
AddlnsSrcInfo( "lune.base.convLua", lune_base_convLua )
AddlnsSrcInfo( "lune.base.dumpNode", lune_base_dumpNode )
AddlnsSrcInfo( "lune.base.convGo", lune_base_convGo )
AddlnsSrcInfo( "lune.base.convPython", lune_base_convPython )
AddlnsSrcInfo( "lune.base.OutputDepend", lune_base_OutputDepend )
AddlnsSrcInfo( "lune.base.Converter", lune_base_Converter )
AddlnsSrcInfo( "lune.base.glueFilter", lune_base_glueFilter )
AddlnsSrcInfo( "lune.base.front", lune_base_front )
AddlnsSrcInfo( "lune.base.runtime7", lune_base_runtime7 )
AddlnsLnsInfo( "lune.base.Testing", lns_lune_base_Testing )
}
