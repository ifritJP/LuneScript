// +build !gopherlua

package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import _ "embed"
//go:embed Async.lua
var lune_base_Async []byte
//go:embed Runner.lua
var lune_base_Runner []byte
//go:embed Str.lua
var lune_base_Str []byte
//go:embed Types.lua
var lune_base_Types []byte
//go:embed Ver.lua
var lune_base_Ver []byte
//go:embed LuaMod.lua
var lune_base_LuaMod []byte
//go:embed LuneControl.lua
var lune_base_LuneControl []byte
//go:embed Meta.lua
var lune_base_Meta []byte
//go:embed Testing.lua
var lune_base_Testing []byte
//go:embed Testing.lns
var lns_lune_base_Testing string
//go:embed TestingMacro.lua
var lune_base_TestingMacro []byte
//go:embed Code.lua
var lune_base_Code []byte
//go:embed Log.lua
var lune_base_Log []byte
//go:embed LuaVer.lua
var lune_base_LuaVer []byte
//go:embed Depend.lua
var lune_base_Depend []byte
//go:embed Util.lua
var lune_base_Util []byte
//go:embed frontInterface.lua
var lune_base_frontInterface []byte
//go:embed Writer.lua
var lune_base_Writer []byte
//go:embed AsyncTokenizer.lua
var lune_base_AsyncTokenizer []byte
//go:embed Tokenizer.lua
var lune_base_Tokenizer []byte
//go:embed Json.lua
var lune_base_Json []byte
//go:embed Ast.lua
var lune_base_Ast []byte
//go:embed Option.lua
var lune_base_Option []byte
//go:embed GoMod.lua
var lune_base_GoMod []byte
//go:embed Nodes.lua
var lune_base_Nodes []byte
//go:embed NodeIndexer.lua
var lune_base_NodeIndexer []byte
//go:embed Formatter.lua
var lune_base_Formatter []byte
//go:embed Macro.lua
var lune_base_Macro []byte
//go:embed TransUnitIF.lua
var lune_base_TransUnitIF []byte
//go:embed BuiltinTransUnit.lua
var lune_base_BuiltinTransUnit []byte
//go:embed Builtin.lua
var lune_base_Builtin []byte
//go:embed AstInfo.lua
var lune_base_AstInfo []byte
//go:embed Import.lua
var lune_base_Import []byte
//go:embed TransUnit.lua
var lune_base_TransUnit []byte
//go:embed DependLuaOnLns.lua
var lune_base_DependLuaOnLns []byte
//go:embed convLua.lua
var lune_base_convLua []byte
//go:embed dumpNode.lua
var lune_base_dumpNode []byte
//go:embed convGo.lua
var lune_base_convGo []byte
//go:embed convPython.lua
var lune_base_convPython []byte
//go:embed OutputDepend.lua
var lune_base_OutputDepend []byte
//go:embed Converter.lua
var lune_base_Converter []byte
//go:embed glueFilter.lua
var lune_base_glueFilter []byte
//go:embed front.lua
var lune_base_front []byte
//go:embed runtime8.lua
var lune_base_runtime8 []byte
func InitBinding() {
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
AddlnsSrcInfo( "lune.base.AsyncTokenizer", lune_base_AsyncTokenizer )
AddlnsSrcInfo( "lune.base.Tokenizer", lune_base_Tokenizer )
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
AddlnsSrcInfo( "lune.base.runtime8", lune_base_runtime8 )
AddlnsLnsInfo( "lune.base.Testing", lns_lune_base_Testing )
}
