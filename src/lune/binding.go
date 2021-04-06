// +build !gopherlua
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import _ "embed"
import "C"
//go:embed base/Async.luac
var lune_base_Async []byte
//go:embed base/Str.luac
var lune_base_Str []byte
//go:embed base/Types.luac
var lune_base_Types []byte
//go:embed base/Ver.luac
var lune_base_Ver []byte
//go:embed base/LuaMod.luac
var lune_base_LuaMod []byte
//go:embed base/LuneControl.luac
var lune_base_LuneControl []byte
//go:embed base/Meta.luac
var lune_base_Meta []byte
//go:embed base/Testing.luac
var lune_base_Testing []byte
//go:embed base/Code.luac
var lune_base_Code []byte
//go:embed base/Log.luac
var lune_base_Log []byte
//go:embed base/LuaVer.luac
var lune_base_LuaVer []byte
//go:embed base/Depend.luac
var lune_base_Depend []byte
//go:embed base/Util.luac
var lune_base_Util []byte
//go:embed base/frontInterface.luac
var lune_base_frontInterface []byte
//go:embed base/Writer.luac
var lune_base_Writer []byte
//go:embed base/AsyncParser.luac
var lune_base_AsyncParser []byte
//go:embed base/Parser.luac
var lune_base_Parser []byte
//go:embed base/Json.luac
var lune_base_Json []byte
//go:embed base/Ast.luac
var lune_base_Ast []byte
//go:embed base/Option.luac
var lune_base_Option []byte
//go:embed base/GoMod.luac
var lune_base_GoMod []byte
//go:embed base/Nodes.luac
var lune_base_Nodes []byte
//go:embed base/Formatter.luac
var lune_base_Formatter []byte
//go:embed base/Macro.luac
var lune_base_Macro []byte
//go:embed base/TransUnit.luac
var lune_base_TransUnit []byte
//go:embed base/DependLuaOnLns.luac
var lune_base_DependLuaOnLns []byte
//go:embed base/convLua.luac
var lune_base_convLua []byte
//go:embed base/dumpNode.luac
var lune_base_dumpNode []byte
//go:embed base/convCC.luac
var lune_base_convCC []byte
//go:embed base/convGo.luac
var lune_base_convGo []byte
//go:embed base/OutputDepend.luac
var lune_base_OutputDepend []byte
//go:embed base/glueFilter.luac
var lune_base_glueFilter []byte
//go:embed base/front.luac
var lune_base_front []byte
//go:embed base/runtime3.luac
var lune_base_runtime3 []byte
func Setup() {
AddlnsSrcInfo( "lune.base.Async", lune_base_Async )
AddlnsSrcInfo( "lune.base.Str", lune_base_Str )
AddlnsSrcInfo( "lune.base.Types", lune_base_Types )
AddlnsSrcInfo( "lune.base.Ver", lune_base_Ver )
AddlnsSrcInfo( "lune.base.LuaMod", lune_base_LuaMod )
AddlnsSrcInfo( "lune.base.LuneControl", lune_base_LuneControl )
AddlnsSrcInfo( "lune.base.Meta", lune_base_Meta )
AddlnsSrcInfo( "lune.base.Testing", lune_base_Testing )
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
AddlnsSrcInfo( "lune.base.Formatter", lune_base_Formatter )
AddlnsSrcInfo( "lune.base.Macro", lune_base_Macro )
AddlnsSrcInfo( "lune.base.TransUnit", lune_base_TransUnit )
AddlnsSrcInfo( "lune.base.DependLuaOnLns", lune_base_DependLuaOnLns )
AddlnsSrcInfo( "lune.base.convLua", lune_base_convLua )
AddlnsSrcInfo( "lune.base.dumpNode", lune_base_dumpNode )
AddlnsSrcInfo( "lune.base.convCC", lune_base_convCC )
AddlnsSrcInfo( "lune.base.convGo", lune_base_convGo )
AddlnsSrcInfo( "lune.base.OutputDepend", lune_base_OutputDepend )
AddlnsSrcInfo( "lune.base.glueFilter", lune_base_glueFilter )
AddlnsSrcInfo( "lune.base.front", lune_base_front )
AddlnsSrcInfo( "lune.base.runtime3", lune_base_runtime3 )
}
