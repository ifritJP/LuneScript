// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_AstInfo bool
var AstInfo__mod__ string
// declaration Class -- ASTInfo
type AstInfo_ASTInfoMtd interface {
    Get_builtinFunc(_env *LnsEnv) *Builtin_BuiltinFuncType
    Get_exportInfo(_env *LnsEnv) *FrontInterface_ExportInfo
    Get_node(_env *LnsEnv) *Nodes_Node
    Get_streamName(_env *LnsEnv) string
}
type AstInfo_ASTInfo struct {
    node *Nodes_Node
    exportInfo *FrontInterface_ExportInfo
    streamName string
    builtinFunc *Builtin_BuiltinFuncType
    FP AstInfo_ASTInfoMtd
}
func AstInfo_ASTInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*AstInfo_ASTInfo).FP
}
type AstInfo_ASTInfoDownCast interface {
    ToAstInfo_ASTInfo() *AstInfo_ASTInfo
}
func AstInfo_ASTInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(AstInfo_ASTInfoDownCast)
    if ok { return work.ToAstInfo_ASTInfo() }
    return nil
}
func (obj *AstInfo_ASTInfo) ToAstInfo_ASTInfo() *AstInfo_ASTInfo {
    return obj
}
func NewAstInfo_ASTInfo(_env *LnsEnv, arg1 *Nodes_Node, arg2 *FrontInterface_ExportInfo, arg3 string, arg4 *Builtin_BuiltinFuncType) *AstInfo_ASTInfo {
    obj := &AstInfo_ASTInfo{}
    obj.FP = obj
    obj.InitAstInfo_ASTInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *AstInfo_ASTInfo) InitAstInfo_ASTInfo(_env *LnsEnv, arg1 *Nodes_Node, arg2 *FrontInterface_ExportInfo, arg3 string, arg4 *Builtin_BuiltinFuncType) {
    self.node = arg1
    self.exportInfo = arg2
    self.streamName = arg3
    self.builtinFunc = arg4
}
func (self *AstInfo_ASTInfo) Get_node(_env *LnsEnv) *Nodes_Node{ return self.node }
func (self *AstInfo_ASTInfo) Get_exportInfo(_env *LnsEnv) *FrontInterface_ExportInfo{ return self.exportInfo }
func (self *AstInfo_ASTInfo) Get_streamName(_env *LnsEnv) string{ return self.streamName }
func (self *AstInfo_ASTInfo) Get_builtinFunc(_env *LnsEnv) *Builtin_BuiltinFuncType{ return self.builtinFunc }

func Lns_AstInfo_init(_env *LnsEnv) {
    if init_AstInfo { return }
    init_AstInfo = true
    AstInfo__mod__ = "@lune.@base.@AstInfo"
    Lns_InitMod()
    Lns_Nodes_init(_env)
    Lns_frontInterface_init(_env)
    Lns_Builtin_init(_env)
}
func init() {
    init_AstInfo = false
}
