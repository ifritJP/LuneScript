// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_NodeIndexer bool
var NodeIndexer__mod__ string
var NodeIndexer_declNameSpaceNodeKindSet *LnsSet

// 164: decl @lune.@base.@NodeIndexer.Indexer.dump.comp
func Indexer_dump__comp_0_(_env *LnsEnv, node1 *Nodes_Node,node2 *Nodes_Node) bool {
    return node1.FP.Get_id(_env) < node2.FP.Get_id(_env)
}

// declaration Class -- NamespaceInfo
type NodeIndexer_NamespaceInfoMtd interface {
    GetIdTxt(_env *LnsEnv) string
    getNewId(_env *LnsEnv, arg1 LnsInt) LnsInt
    Get_childCount(_env *LnsEnv) LnsInt
    Get_nsType(_env *LnsEnv) *Ast_TypeInfo
    Get_subId(_env *LnsEnv) LnsInt
}
type NodeIndexer_NamespaceInfo struct {
    depth LnsInt
    parent LnsAny
    subId LnsInt
    childCount LnsInt
    nsType *Ast_TypeInfo
    idProvMap *LnsMap
    FP NodeIndexer_NamespaceInfoMtd
}
func NodeIndexer_NamespaceInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*NodeIndexer_NamespaceInfo).FP
}
type NodeIndexer_NamespaceInfoDownCast interface {
    ToNodeIndexer_NamespaceInfo() *NodeIndexer_NamespaceInfo
}
func NodeIndexer_NamespaceInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(NodeIndexer_NamespaceInfoDownCast)
    if ok { return work.ToNodeIndexer_NamespaceInfo() }
    return nil
}
func (obj *NodeIndexer_NamespaceInfo) ToNodeIndexer_NamespaceInfo() *NodeIndexer_NamespaceInfo {
    return obj
}
func NewNodeIndexer_NamespaceInfo(_env *LnsEnv, arg1 LnsAny, arg2 *Ast_TypeInfo) *NodeIndexer_NamespaceInfo {
    obj := &NodeIndexer_NamespaceInfo{}
    obj.FP = obj
    obj.InitNodeIndexer_NamespaceInfo(_env, arg1, arg2)
    return obj
}
func (self *NodeIndexer_NamespaceInfo) Get_subId(_env *LnsEnv) LnsInt{ return self.subId }
func (self *NodeIndexer_NamespaceInfo) Get_childCount(_env *LnsEnv) LnsInt{ return self.childCount }
func (self *NodeIndexer_NamespaceInfo) Get_nsType(_env *LnsEnv) *Ast_TypeInfo{ return self.nsType }
// 18: DeclConstr
func (self *NodeIndexer_NamespaceInfo) InitNodeIndexer_NamespaceInfo(_env *LnsEnv, parent LnsAny,nsType *Ast_TypeInfo) {
    self.parent = parent
    self.childCount = 1
    var subId LnsInt
    var depth LnsInt
    if parent != nil{
        parent_26 := parent.(*NodeIndexer_NamespaceInfo)
        subId = parent_26.childCount
        depth = parent_26.depth + 1
        parent_26.childCount = parent_26.childCount + 1
    } else {
        subId = 0
        depth = 0
    }
    self.subId = subId
    self.depth = depth
    self.nsType = nsType
    self.idProvMap = NewLnsMap( map[LnsAny]LnsAny{})
}

// 39: decl @lune.@base.@NodeIndexer.NamespaceInfo.getNewId
func (self *NodeIndexer_NamespaceInfo) getNewId(_env *LnsEnv, nodeKind LnsInt) LnsInt {
    var idProv *Ast_IdProvider
    
    {
        _idProv := self.idProvMap.Get(nodeKind)
        if _idProv == nil{
            var work *Ast_IdProvider
            work = NewAst_IdProvider(_env, 0, 100000)
            self.idProvMap.Set(nodeKind,work)
            idProv = work
        } else {
            idProv = _idProv.(*Ast_IdProvider)
        }
    }
    return idProv.FP.GetNewId(_env)
}

// 48: decl @lune.@base.@NodeIndexer.NamespaceInfo.getIdTxt
func (self *NodeIndexer_NamespaceInfo) GetIdTxt(_env *LnsEnv) string {
    {
        _parent := self.parent
        if !Lns_IsNil( _parent ) {
            parent := _parent.(*NodeIndexer_NamespaceInfo)
            if self.depth < 4{
                return _env.GetVM().String_format("%s_%s", []LnsAny{parent.FP.GetIdTxt(_env), self.FP.Get_nsType(_env).FP.Get_rawTxt(_env)})
            } else { 
                return _env.GetVM().String_format("%s%s", []LnsAny{parent.FP.GetIdTxt(_env), _env.GetVM().String_format("_%d", []LnsAny{self.subId})})
            }
        } else {
            return ""
        }
    }
// insert a dummy
    return ""
}


// declaration Class -- Index
type NodeIndexer_IndexMtd interface {
    GetIdTxt(_env *LnsEnv) string
    Get_index(_env *LnsEnv) LnsInt
    get_nsInfo(_env *LnsEnv) *NodeIndexer_NamespaceInfo
}
type NodeIndexer_Index struct {
    nsInfo *NodeIndexer_NamespaceInfo
    index LnsInt
    FP NodeIndexer_IndexMtd
}
func NodeIndexer_Index2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*NodeIndexer_Index).FP
}
type NodeIndexer_IndexDownCast interface {
    ToNodeIndexer_Index() *NodeIndexer_Index
}
func NodeIndexer_IndexDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(NodeIndexer_IndexDownCast)
    if ok { return work.ToNodeIndexer_Index() }
    return nil
}
func (obj *NodeIndexer_Index) ToNodeIndexer_Index() *NodeIndexer_Index {
    return obj
}
func NewNodeIndexer_Index(_env *LnsEnv, arg1 *NodeIndexer_NamespaceInfo, arg2 LnsInt) *NodeIndexer_Index {
    obj := &NodeIndexer_Index{}
    obj.FP = obj
    obj.InitNodeIndexer_Index(_env, arg1, arg2)
    return obj
}
func (self *NodeIndexer_Index) get_nsInfo(_env *LnsEnv) *NodeIndexer_NamespaceInfo{ return self.nsInfo }
func (self *NodeIndexer_Index) Get_index(_env *LnsEnv) LnsInt{ return self.index }
// 78: DeclConstr
func (self *NodeIndexer_Index) InitNodeIndexer_Index(_env *LnsEnv, nsInfo *NodeIndexer_NamespaceInfo,index LnsInt) {
    self.nsInfo = nsInfo
    self.index = index
}

// 83: decl @lune.@base.@NodeIndexer.Index.getIdTxt
func (self *NodeIndexer_Index) GetIdTxt(_env *LnsEnv) string {
    return _env.GetVM().String_format("%s_%d", []LnsAny{self.nsInfo.FP.GetIdTxt(_env), self.index})
}


// declaration Class -- Indexer
type NodeIndexer_IndexerMtd interface {
    Dump(_env *LnsEnv)
    GetIndex(_env *LnsEnv, arg1 *Nodes_Node) *NodeIndexer_Index
    Get_node2Index(_env *LnsEnv) *LnsMap
    Start(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsSet)
    visit(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsInt)
}
type NodeIndexer_Indexer struct {
    nsInfo *NodeIndexer_NamespaceInfo
    processInfo *Ast_ProcessInfo
    node2Index *LnsMap
    nsType2nsInfo *LnsMap
    targetNodeKindSet *LnsSet
    FP NodeIndexer_IndexerMtd
}
func NodeIndexer_Indexer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*NodeIndexer_Indexer).FP
}
type NodeIndexer_IndexerDownCast interface {
    ToNodeIndexer_Indexer() *NodeIndexer_Indexer
}
func NodeIndexer_IndexerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(NodeIndexer_IndexerDownCast)
    if ok { return work.ToNodeIndexer_Indexer() }
    return nil
}
func (obj *NodeIndexer_Indexer) ToNodeIndexer_Indexer() *NodeIndexer_Indexer {
    return obj
}
func NewNodeIndexer_Indexer(_env *LnsEnv, arg1 *Ast_ProcessInfo) *NodeIndexer_Indexer {
    obj := &NodeIndexer_Indexer{}
    obj.FP = obj
    obj.InitNodeIndexer_Indexer(_env, arg1)
    return obj
}
func (self *NodeIndexer_Indexer) Get_node2Index(_env *LnsEnv) *LnsMap{ return self.node2Index }
// 104: DeclConstr
func (self *NodeIndexer_Indexer) InitNodeIndexer_Indexer(_env *LnsEnv, processInfo *Ast_ProcessInfo) {
    self.nsType2nsInfo = NewLnsMap( map[LnsAny]LnsAny{})
    self.processInfo = processInfo
    self.nsInfo = NewNodeIndexer_NamespaceInfo(_env, nil, Ast_headTypeInfo)
    self.node2Index = NewLnsMap( map[LnsAny]LnsAny{})
    self.targetNodeKindSet = NewLnsSet([]LnsAny{})
}

// 113: decl @lune.@base.@NodeIndexer.Indexer.getIndex
func (self *NodeIndexer_Indexer) GetIndex(_env *LnsEnv, node *Nodes_Node) *NodeIndexer_Index {
    return Lns_unwrap( self.node2Index.Get(node)).(*NodeIndexer_Index)
}

// 117: decl @lune.@base.@NodeIndexer.Indexer.visit
func (self *NodeIndexer_Indexer) visit(_env *LnsEnv, targetNode *Nodes_Node,dept LnsInt) {
    targetNode.FP.Visit(_env, Nodes_NodeVisitor(func(_env *LnsEnv, node *Nodes_Node,parent *Nodes_Node,relation string,depth LnsInt) LnsInt {
        if self.targetNodeKindSet.Has(node.FP.Get_kind(_env)){
            self.node2Index.Set(node,NewNodeIndexer_Index(_env, self.nsInfo, self.nsInfo.FP.getNewId(_env, node.FP.Get_kind(_env))))
        }
        if NodeIndexer_declNameSpaceNodeKindSet.Has(node.FP.Get_kind(_env)){
            var bakNsInfo *NodeIndexer_NamespaceInfo
            bakNsInfo = self.nsInfo
            var parentNsInfo *NodeIndexer_NamespaceInfo
            parentNsInfo = self.nsInfo
            {
                _declMethodNode := Nodes_DeclMethodNodeDownCastF(node.FP)
                if !Lns_IsNil( _declMethodNode ) {
                    declMethodNode := _declMethodNode.(*Nodes_DeclMethodNode)
                    if declMethodNode.FP.Get_declInfo(_env).FP.Get_outsizeOfClass(_env){
                        parentNsInfo = Lns_unwrap( self.nsType2nsInfo.Get(node.FP.Get_expType(_env).FP.Get_parentInfo(_env))).(*NodeIndexer_NamespaceInfo)
                    }
                }
            }
            self.nsInfo = NewNodeIndexer_NamespaceInfo(_env, parentNsInfo, node.FP.Get_expType(_env))
            self.nsType2nsInfo.Set(node.FP.Get_expType(_env),self.nsInfo)
            self.FP.visit(_env, node, depth + 1)
            self.nsInfo = bakNsInfo
            return Nodes_NodeVisitMode__Next
        }
        return Nodes_NodeVisitMode__Child
    }), 0, NewLnsSet([]LnsAny{}))
}

// 152: decl @lune.@base.@NodeIndexer.Indexer.start
func (self *NodeIndexer_Indexer) Start(_env *LnsEnv, targetNode *Nodes_Node,targetNodeKindSet *LnsSet) {
    self.targetNodeKindSet = targetNodeKindSet
    self.FP.visit(_env, targetNode, 0)
}

// 159: decl @lune.@base.@NodeIndexer.Indexer.dump
func (self *NodeIndexer_Indexer) Dump(_env *LnsEnv) {
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    for _node, _ := range( self.node2Index.Items ) {
        node := _node.(Nodes_NodeDownCast).ToNodes_Node()
        list.Insert(Nodes_Node2Stem(node))
    }
    list.Sort(_env, LnsItemKindStem, LnsComp(func ( _env *LnsEnv, val1, val2 LnsAny ) bool {return Indexer_dump__comp_0_( _env, val1.(Nodes_NodeDownCast).ToNodes_Node(), val2.(Nodes_NodeDownCast).ToNodes_Node())}))
    for _, _node := range( list.Items ) {
        node := _node.(Nodes_NodeDownCast).ToNodes_Node()
        var index *NodeIndexer_Index
        index = Lns_unwrap( self.node2Index.Get(node)).(*NodeIndexer_Index)
        Lns_print([]LnsAny{_env.GetVM().String_format("%d:%d:", []LnsAny{node.FP.Get_pos(_env).LineNo, node.FP.Get_pos(_env).Column}), index.FP.GetIdTxt(_env), Nodes_getNodeKindName(_env, node.FP.Get_kind(_env))})
    }
}


func Lns_NodeIndexer_init(_env *LnsEnv) {
    if init_NodeIndexer { return }
    init_NodeIndexer = true
    NodeIndexer__mod__ = "@lune.@base.@NodeIndexer"
    Lns_InitMod()
    Lns_Nodes_init(_env)
    Lns_Ast_init(_env)
    NodeIndexer_declNameSpaceNodeKindSet = NewLnsSet([]LnsAny{Nodes_NodeKind_get_DeclClass(_env), Nodes_NodeKind_get_DeclConstr(_env), Nodes_NodeKind_get_DeclFunc(_env), Nodes_NodeKind_get_ProtoMethod(_env), Nodes_NodeKind_get_DeclMethod(_env), Nodes_NodeKind_get_DeclEnum(_env), Nodes_NodeKind_get_DeclAlge(_env), Nodes_NodeKind_get_DeclMacro(_env)})
}
func init() {
    init_NodeIndexer = false
}
