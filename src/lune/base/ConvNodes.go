// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_ConvNodes bool
var ConvNodes__mod__ string
var ConvNodes_nodeKind2createFromFunc *LnsMap
type ConvNodes_NodeVisitor func (arg1 *ConvNodes_Node,arg2 *ConvNodes_Node,arg3 string,arg4 LnsInt) LnsInt
type ConvNodes_createFromFunc_1063_ func (arg1 *Nodes_Node) *ConvNodes_Node
// for 76
func ConvNodes_convExp242(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 77
func ConvNodes_convExp255(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 59: decl @lune.@base.@ConvNodes.Node_createFromNode
func ConvNodes_Node_createFromNode_1075_(node *Nodes_Node) *ConvNodes_Node {
    return ConvNodes_Node_createFrom(node)
}




// 1: decl @lune.@base.@ConvNodes.ExpListNode_createFromNode
func ConvNodes_ExpListNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpListNode {
    var node *Nodes_ExpListNode
    
    {
        _node := Nodes_ExpListNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpListNode"}))
        } else {
            node = _node.(*Nodes_ExpListNode)
        }
    }
    var createexpList func() *LnsList
    createexpList = func() *LnsList {
        var list *LnsList
        list = node.FP.Get_expList()
        var expList *LnsList
        expList = NewLnsList([]LnsAny{})
        for _, _exp := range( list.Items ) {
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            expList.Insert(ConvNodes_Node2Stem(ConvNodes_Node_createFrom(exp)))
        }
        return expList
    }
    return NewConvNodes_ExpListNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpListNode", node, createexpList())
}

func ConvNodes__anonymous_1185_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpListNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.ExpNewNode_createFromNode
func ConvNodes_ExpNewNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpNewNode {
    var node *Nodes_ExpNewNode
    
    {
        _node := Nodes_ExpNewNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpNewNode"}))
        } else {
            node = _node.(*Nodes_ExpNewNode)
        }
    }
    var createargList func() LnsAny
    createargList = func() LnsAny {
        var child *Nodes_ExpListNode
        
        {
            _child := node.FP.Get_argList()
            if _child == nil{
                return nil
            } else {
                child = _child.(*Nodes_ExpListNode)
            }
        }
        return ConvNodes_ExpListNode_createFromNode(&child.Nodes_Node)
    }
    return NewConvNodes_ExpNewNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpNewNode", node, createargList())
}

func ConvNodes__anonymous_1233_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpNewNode_createFromNode(node).ConvNodes_Node
}


// 1: decl @lune.@base.@ConvNodes.ExpUnwrapNode_createFromNode
func ConvNodes_ExpUnwrapNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpUnwrapNode {
    var node *Nodes_ExpUnwrapNode
    
    {
        _node := Nodes_ExpUnwrapNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpUnwrapNode"}))
        } else {
            node = _node.(*Nodes_ExpUnwrapNode)
        }
    }
    var createexp func() *ConvNodes_Node
    createexp = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_exp()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    var createdefault func() LnsAny
    createdefault = func() LnsAny {
        var child *Nodes_Node
        
        {
            _child := node.FP.Get_default()
            if _child == nil{
                return nil
            } else {
                child = _child.(*Nodes_Node)
            }
        }
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpUnwrapNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpUnwrapNode", node, createexp(), createdefault())
}

func ConvNodes__anonymous_1290_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpUnwrapNode_createFromNode(node).ConvNodes_Node
}
// 1: decl @lune.@base.@ConvNodes.ExpRefNode_createFromNode
func ConvNodes_ExpRefNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpRefNode {
    var node *Nodes_ExpRefNode
    
    {
        _node := Nodes_ExpRefNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpRefNode"}))
        } else {
            node = _node.(*Nodes_ExpRefNode)
        }
    }
    return NewConvNodes_ExpRefNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpRefNode", node)
}

func ConvNodes__anonymous_1328_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpRefNode_createFromNode(node).ConvNodes_Node
}


// 1: decl @lune.@base.@ConvNodes.ExpOp2Node_createFromNode
func ConvNodes_ExpOp2Node_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpOp2Node {
    var node *Nodes_ExpOp2Node
    
    {
        _node := Nodes_ExpOp2NodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpOp2Node"}))
        } else {
            node = _node.(*Nodes_ExpOp2Node)
        }
    }
    var createexp1 func() *ConvNodes_Node
    createexp1 = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_exp1()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    var createexp2 func() *ConvNodes_Node
    createexp2 = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_exp2()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpOp2Node(node.FP.Get_id(), node.FP.Get_pos(), "ExpOp2Node", node, createexp1(), createexp2())
}

func ConvNodes__anonymous_1384_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpOp2Node_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.ExpCastNode_createFromNode
func ConvNodes_ExpCastNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpCastNode {
    var node *Nodes_ExpCastNode
    
    {
        _node := Nodes_ExpCastNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpCastNode"}))
        } else {
            node = _node.(*Nodes_ExpCastNode)
        }
    }
    var createexp func() *ConvNodes_Node
    createexp = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_exp()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpCastNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpCastNode", node, createexp())
}

func ConvNodes__anonymous_1431_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpCastNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.ExpToDDDNode_createFromNode
func ConvNodes_ExpToDDDNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpToDDDNode {
    var node *Nodes_ExpToDDDNode
    
    {
        _node := Nodes_ExpToDDDNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpToDDDNode"}))
        } else {
            node = _node.(*Nodes_ExpToDDDNode)
        }
    }
    var createexpList func() *ConvNodes_ExpListNode
    createexpList = func() *ConvNodes_ExpListNode {
        var child *Nodes_ExpListNode
        child = node.FP.Get_expList()
        return ConvNodes_ExpListNode_createFromNode(&child.Nodes_Node)
    }
    return NewConvNodes_ExpToDDDNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpToDDDNode", node, createexpList())
}

func ConvNodes__anonymous_1478_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpToDDDNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.ExpSubDDDNode_createFromNode
func ConvNodes_ExpSubDDDNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpSubDDDNode {
    var node *Nodes_ExpSubDDDNode
    
    {
        _node := Nodes_ExpSubDDDNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpSubDDDNode"}))
        } else {
            node = _node.(*Nodes_ExpSubDDDNode)
        }
    }
    var createsrc func() *ConvNodes_Node
    createsrc = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_src()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpSubDDDNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpSubDDDNode", node, createsrc())
}

func ConvNodes__anonymous_1525_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpSubDDDNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.ExpOp1Node_createFromNode
func ConvNodes_ExpOp1Node_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpOp1Node {
    var node *Nodes_ExpOp1Node
    
    {
        _node := Nodes_ExpOp1NodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpOp1Node"}))
        } else {
            node = _node.(*Nodes_ExpOp1Node)
        }
    }
    var createexp func() *ConvNodes_Node
    createexp = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_exp()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpOp1Node(node.FP.Get_id(), node.FP.Get_pos(), "ExpOp1Node", node, createexp())
}

func ConvNodes__anonymous_1572_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpOp1Node_createFromNode(node).ConvNodes_Node
}


// 1: decl @lune.@base.@ConvNodes.ExpRefItemNode_createFromNode
func ConvNodes_ExpRefItemNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpRefItemNode {
    var node *Nodes_ExpRefItemNode
    
    {
        _node := Nodes_ExpRefItemNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpRefItemNode"}))
        } else {
            node = _node.(*Nodes_ExpRefItemNode)
        }
    }
    var createval func() *ConvNodes_Node
    createval = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_val()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    var createindex func() LnsAny
    createindex = func() LnsAny {
        var child *Nodes_Node
        
        {
            _child := node.FP.Get_index()
            if _child == nil{
                return nil
            } else {
                child = _child.(*Nodes_Node)
            }
        }
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpRefItemNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpRefItemNode", node, createval(), createindex())
}

func ConvNodes__anonymous_1628_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpRefItemNode_createFromNode(node).ConvNodes_Node
}


// 1: decl @lune.@base.@ConvNodes.ExpCallNode_createFromNode
func ConvNodes_ExpCallNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpCallNode {
    var node *Nodes_ExpCallNode
    
    {
        _node := Nodes_ExpCallNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpCallNode"}))
        } else {
            node = _node.(*Nodes_ExpCallNode)
        }
    }
    var createfunc func() *ConvNodes_Node
    createfunc = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_func()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    var createargList func() LnsAny
    createargList = func() LnsAny {
        var child *Nodes_ExpListNode
        
        {
            _child := node.FP.Get_argList()
            if _child == nil{
                return nil
            } else {
                child = _child.(*Nodes_ExpListNode)
            }
        }
        return ConvNodes_ExpListNode_createFromNode(&child.Nodes_Node)
    }
    return NewConvNodes_ExpCallNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpCallNode", node, createfunc(), createargList())
}

func ConvNodes__anonymous_1684_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpCallNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.ExpMRetNode_createFromNode
func ConvNodes_ExpMRetNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpMRetNode {
    var node *Nodes_ExpMRetNode
    
    {
        _node := Nodes_ExpMRetNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpMRetNode"}))
        } else {
            node = _node.(*Nodes_ExpMRetNode)
        }
    }
    var createmRet func() *ConvNodes_Node
    createmRet = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_mRet()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpMRetNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpMRetNode", node, createmRet())
}

func ConvNodes__anonymous_1731_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpMRetNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.ExpAccessMRetNode_createFromNode
func ConvNodes_ExpAccessMRetNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpAccessMRetNode {
    var node *Nodes_ExpAccessMRetNode
    
    {
        _node := Nodes_ExpAccessMRetNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpAccessMRetNode"}))
        } else {
            node = _node.(*Nodes_ExpAccessMRetNode)
        }
    }
    var createmRet func() *ConvNodes_Node
    createmRet = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_mRet()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpAccessMRetNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpAccessMRetNode", node, createmRet())
}

func ConvNodes__anonymous_1778_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpAccessMRetNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.ExpMultiTo1Node_createFromNode
func ConvNodes_ExpMultiTo1Node_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpMultiTo1Node {
    var node *Nodes_ExpMultiTo1Node
    
    {
        _node := Nodes_ExpMultiTo1NodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpMultiTo1Node"}))
        } else {
            node = _node.(*Nodes_ExpMultiTo1Node)
        }
    }
    var createexp func() *ConvNodes_Node
    createexp = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_exp()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpMultiTo1Node(node.FP.Get_id(), node.FP.Get_pos(), "ExpMultiTo1Node", node, createexp())
}

func ConvNodes__anonymous_1825_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpMultiTo1Node_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.ExpParenNode_createFromNode
func ConvNodes_ExpParenNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpParenNode {
    var node *Nodes_ExpParenNode
    
    {
        _node := Nodes_ExpParenNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpParenNode"}))
        } else {
            node = _node.(*Nodes_ExpParenNode)
        }
    }
    var createexp func() *ConvNodes_Node
    createexp = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_exp()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_ExpParenNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpParenNode", node, createexp())
}

func ConvNodes__anonymous_1872_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpParenNode_createFromNode(node).ConvNodes_Node
}
// 1: decl @lune.@base.@ConvNodes.ExpOmitEnumNode_createFromNode
func ConvNodes_ExpOmitEnumNode_createFromNode(workNode *Nodes_Node) *ConvNodes_ExpOmitEnumNode {
    var node *Nodes_ExpOmitEnumNode
    
    {
        _node := Nodes_ExpOmitEnumNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "ExpOmitEnumNode"}))
        } else {
            node = _node.(*Nodes_ExpOmitEnumNode)
        }
    }
    return NewConvNodes_ExpOmitEnumNode(node.FP.Get_id(), node.FP.Get_pos(), "ExpOmitEnumNode", node)
}

func ConvNodes__anonymous_1910_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_ExpOmitEnumNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.RefFieldNode_createFromNode
func ConvNodes_RefFieldNode_createFromNode(workNode *Nodes_Node) *ConvNodes_RefFieldNode {
    var node *Nodes_RefFieldNode
    
    {
        _node := Nodes_RefFieldNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "RefFieldNode"}))
        } else {
            node = _node.(*Nodes_RefFieldNode)
        }
    }
    var createprefix func() *ConvNodes_Node
    createprefix = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_prefix()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_RefFieldNode(node.FP.Get_id(), node.FP.Get_pos(), "RefFieldNode", node, createprefix())
}

func ConvNodes__anonymous_1957_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_RefFieldNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.GetFieldNode_createFromNode
func ConvNodes_GetFieldNode_createFromNode(workNode *Nodes_Node) *ConvNodes_GetFieldNode {
    var node *Nodes_GetFieldNode
    
    {
        _node := Nodes_GetFieldNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "GetFieldNode"}))
        } else {
            node = _node.(*Nodes_GetFieldNode)
        }
    }
    var createprefix func() *ConvNodes_Node
    createprefix = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_prefix()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_GetFieldNode(node.FP.Get_id(), node.FP.Get_pos(), "GetFieldNode", node, createprefix())
}

func ConvNodes__anonymous_2004_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_GetFieldNode_createFromNode(node).ConvNodes_Node
}
// 1: decl @lune.@base.@ConvNodes.DeclFuncNode_createFromNode
func ConvNodes_DeclFuncNode_createFromNode(workNode *Nodes_Node) *ConvNodes_DeclFuncNode {
    var node *Nodes_DeclFuncNode
    
    {
        _node := Nodes_DeclFuncNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "DeclFuncNode"}))
        } else {
            node = _node.(*Nodes_DeclFuncNode)
        }
    }
    return NewConvNodes_DeclFuncNode(node.FP.Get_id(), node.FP.Get_pos(), "DeclFuncNode", node)
}

func ConvNodes__anonymous_2042_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_DeclFuncNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.NewAlgeValNode_createFromNode
func ConvNodes_NewAlgeValNode_createFromNode(workNode *Nodes_Node) *ConvNodes_NewAlgeValNode {
    var node *Nodes_NewAlgeValNode
    
    {
        _node := Nodes_NewAlgeValNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "NewAlgeValNode"}))
        } else {
            node = _node.(*Nodes_NewAlgeValNode)
        }
    }
    var createparamList func() *LnsList
    createparamList = func() *LnsList {
        var list *LnsList
        list = node.FP.Get_paramList()
        var expList *LnsList
        expList = NewLnsList([]LnsAny{})
        for _, _exp := range( list.Items ) {
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            expList.Insert(ConvNodes_Node2Stem(ConvNodes_Node_createFrom(exp)))
        }
        return expList
    }
    return NewConvNodes_NewAlgeValNode(node.FP.Get_id(), node.FP.Get_pos(), "NewAlgeValNode", node, createparamList())
}

func ConvNodes__anonymous_2103_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_NewAlgeValNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.LuneKindNode_createFromNode
func ConvNodes_LuneKindNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LuneKindNode {
    var node *Nodes_LuneKindNode
    
    {
        _node := Nodes_LuneKindNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LuneKindNode"}))
        } else {
            node = _node.(*Nodes_LuneKindNode)
        }
    }
    var createexp func() *ConvNodes_Node
    createexp = func() *ConvNodes_Node {
        var child *Nodes_Node
        child = node.FP.Get_exp()
        return ConvNodes_Node_createFromNode_1075_(child)
    }
    return NewConvNodes_LuneKindNode(node.FP.Get_id(), node.FP.Get_pos(), "LuneKindNode", node, createexp())
}

func ConvNodes__anonymous_2150_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LuneKindNode_createFromNode(node).ConvNodes_Node
}
// 1: decl @lune.@base.@ConvNodes.LiteralCharNode_createFromNode
func ConvNodes_LiteralCharNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralCharNode {
    var node *Nodes_LiteralCharNode
    
    {
        _node := Nodes_LiteralCharNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LiteralCharNode"}))
        } else {
            node = _node.(*Nodes_LiteralCharNode)
        }
    }
    return NewConvNodes_LiteralCharNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralCharNode", node)
}

func ConvNodes__anonymous_2188_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LiteralCharNode_createFromNode(node).ConvNodes_Node
}
// 1: decl @lune.@base.@ConvNodes.LiteralIntNode_createFromNode
func ConvNodes_LiteralIntNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralIntNode {
    var node *Nodes_LiteralIntNode
    
    {
        _node := Nodes_LiteralIntNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LiteralIntNode"}))
        } else {
            node = _node.(*Nodes_LiteralIntNode)
        }
    }
    return NewConvNodes_LiteralIntNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralIntNode", node)
}

func ConvNodes__anonymous_2226_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LiteralIntNode_createFromNode(node).ConvNodes_Node
}
// 1: decl @lune.@base.@ConvNodes.LiteralRealNode_createFromNode
func ConvNodes_LiteralRealNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralRealNode {
    var node *Nodes_LiteralRealNode
    
    {
        _node := Nodes_LiteralRealNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LiteralRealNode"}))
        } else {
            node = _node.(*Nodes_LiteralRealNode)
        }
    }
    return NewConvNodes_LiteralRealNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralRealNode", node)
}

func ConvNodes__anonymous_2264_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LiteralRealNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.LiteralArrayNode_createFromNode
func ConvNodes_LiteralArrayNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralArrayNode {
    var node *Nodes_LiteralArrayNode
    
    {
        _node := Nodes_LiteralArrayNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LiteralArrayNode"}))
        } else {
            node = _node.(*Nodes_LiteralArrayNode)
        }
    }
    var createexpList func() LnsAny
    createexpList = func() LnsAny {
        var child *Nodes_ExpListNode
        
        {
            _child := node.FP.Get_expList()
            if _child == nil{
                return nil
            } else {
                child = _child.(*Nodes_ExpListNode)
            }
        }
        return ConvNodes_ExpListNode_createFromNode(&child.Nodes_Node)
    }
    return NewConvNodes_LiteralArrayNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralArrayNode", node, createexpList())
}

func ConvNodes__anonymous_2311_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LiteralArrayNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.LiteralListNode_createFromNode
func ConvNodes_LiteralListNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralListNode {
    var node *Nodes_LiteralListNode
    
    {
        _node := Nodes_LiteralListNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LiteralListNode"}))
        } else {
            node = _node.(*Nodes_LiteralListNode)
        }
    }
    var createexpList func() LnsAny
    createexpList = func() LnsAny {
        var child *Nodes_ExpListNode
        
        {
            _child := node.FP.Get_expList()
            if _child == nil{
                return nil
            } else {
                child = _child.(*Nodes_ExpListNode)
            }
        }
        return ConvNodes_ExpListNode_createFromNode(&child.Nodes_Node)
    }
    return NewConvNodes_LiteralListNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralListNode", node, createexpList())
}

func ConvNodes__anonymous_2358_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LiteralListNode_createFromNode(node).ConvNodes_Node
}

// 1: decl @lune.@base.@ConvNodes.LiteralSetNode_createFromNode
func ConvNodes_LiteralSetNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralSetNode {
    var node *Nodes_LiteralSetNode
    
    {
        _node := Nodes_LiteralSetNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LiteralSetNode"}))
        } else {
            node = _node.(*Nodes_LiteralSetNode)
        }
    }
    var createexpList func() LnsAny
    createexpList = func() LnsAny {
        var child *Nodes_ExpListNode
        
        {
            _child := node.FP.Get_expList()
            if _child == nil{
                return nil
            } else {
                child = _child.(*Nodes_ExpListNode)
            }
        }
        return ConvNodes_ExpListNode_createFromNode(&child.Nodes_Node)
    }
    return NewConvNodes_LiteralSetNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralSetNode", node, createexpList())
}

func ConvNodes__anonymous_2405_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LiteralSetNode_createFromNode(node).ConvNodes_Node
}
// 297: decl @lune.@base.@ConvNodes.LiteralMapNode_createFromNode
func ConvNodes_LiteralMapNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralMapNode {
    __func__ := "@lune.@base.@ConvNodes.LiteralMapNode_createFromNode"
    var node *Nodes_LiteralMapNode
    
    {
        _node := Nodes_LiteralMapNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s -- %s", []LnsAny{workNode.FP.Get_kind(), __func__}))
        } else {
            node = _node.(*Nodes_LiteralMapNode)
        }
    }
    var pairList *LnsList
    pairList = NewLnsList([]LnsAny{})
    for _, _item := range( node.FP.Get_pairList().Items ) {
        item := _item.(Nodes_PairItemDownCast).ToNodes_PairItem()
        var key *ConvNodes_Node
        key = ConvNodes_Node_createFromNode_1075_(item.FP.Get_key())
        var val *ConvNodes_Node
        val = ConvNodes_Node_createFromNode_1075_(item.FP.Get_val())
        pairList.Insert(ConvNodes_PairItem2Stem(NewConvNodes_PairItem(key, val)))
    }
    return NewConvNodes_LiteralMapNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralMapNode", node, pairList)
}

func ConvNodes__anonymous_2476_(node *Nodes_Node) *ConvNodes_Node {
    return ConvNodes_Node_createFromNode_1075_(node)
}

// 1: decl @lune.@base.@ConvNodes.LiteralStringNode_createFromNode
func ConvNodes_LiteralStringNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralStringNode {
    var node *Nodes_LiteralStringNode
    
    {
        _node := Nodes_LiteralStringNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LiteralStringNode"}))
        } else {
            node = _node.(*Nodes_LiteralStringNode)
        }
    }
    var createdddParam func() LnsAny
    createdddParam = func() LnsAny {
        var child *Nodes_ExpListNode
        
        {
            _child := node.FP.Get_dddParam()
            if _child == nil{
                return nil
            } else {
                child = _child.(*Nodes_ExpListNode)
            }
        }
        return ConvNodes_ExpListNode_createFromNode(&child.Nodes_Node)
    }
    return NewConvNodes_LiteralStringNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralStringNode", node, createdddParam())
}

func ConvNodes__anonymous_2523_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LiteralStringNode_createFromNode(node).ConvNodes_Node
}
// 1: decl @lune.@base.@ConvNodes.LiteralBoolNode_createFromNode
func ConvNodes_LiteralBoolNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralBoolNode {
    var node *Nodes_LiteralBoolNode
    
    {
        _node := Nodes_LiteralBoolNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LiteralBoolNode"}))
        } else {
            node = _node.(*Nodes_LiteralBoolNode)
        }
    }
    return NewConvNodes_LiteralBoolNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralBoolNode", node)
}

func ConvNodes__anonymous_2561_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LiteralBoolNode_createFromNode(node).ConvNodes_Node
}
// 1: decl @lune.@base.@ConvNodes.LiteralSymbolNode_createFromNode
func ConvNodes_LiteralSymbolNode_createFromNode(workNode *Nodes_Node) *ConvNodes_LiteralSymbolNode {
    var node *Nodes_LiteralSymbolNode
    
    {
        _node := Nodes_LiteralSymbolNodeDownCastF(workNode.FP)
        if _node == nil{
            Util_err(Lns_getVM().String_format("illegal node -- %s, %s", []LnsAny{Nodes_getNodeKindName(workNode.FP.Get_kind()), "LiteralSymbolNode"}))
        } else {
            node = _node.(*Nodes_LiteralSymbolNode)
        }
    }
    return NewConvNodes_LiteralSymbolNode(node.FP.Get_id(), node.FP.Get_pos(), "LiteralSymbolNode", node)
}

func ConvNodes__anonymous_2599_(node *Nodes_Node) *ConvNodes_Node {
    return &ConvNodes_LiteralSymbolNode_createFromNode(node).ConvNodes_Node
}
func convertNodes____anonymous_2886_(convNode *ConvNodes_Node,convParent *ConvNodes_Node,convRelation string,convDepth LnsInt) LnsInt {
    Lns_print([]LnsAny{Lns_getVM().String_format("%s: conved %s: %d:%d", []LnsAny{Lns_getVM().String_rep(" ", convDepth * 3) + convRelation, convNode.FP.Get_kind(), convNode.FP.Get_pos().LineNo, convNode.FP.Get_pos().Column})})
    return Nodes_NodeVisitMode__Child
}
func convertNodes___anonymous_2883_(node *Nodes_Node,parent *Nodes_Node,relation string,depth LnsInt) LnsInt {
    Lns_print([]LnsAny{Lns_getVM().String_format("%s: %s: %d:%d", []LnsAny{Lns_getVM().String_rep(" ", depth * 3) + relation, Nodes_getNodeKindName(node.FP.Get_kind()), node.FP.Get_pos().LineNo, node.FP.Get_pos().Column})})
    {
        _createFunc := ConvNodes_nodeKind2createFromFunc.Items[node.FP.Get_kind()]
        if _createFunc != nil {
            createFunc := _createFunc.(ConvNodes_createFromFunc_1063_)
            var workNode *ConvNodes_Node
            workNode = createFunc(node)
            Lns_print([]LnsAny{Lns_getVM().String_format("%s: conved %s: %d:%d", []LnsAny{Lns_getVM().String_rep(" ", depth * 3) + relation, workNode.FP.Get_kind(), workNode.FP.Get_pos().LineNo, workNode.FP.Get_pos().Column})})
            workNode.FP.Visit(ConvNodes_NodeVisitor(convertNodes____anonymous_2886_), depth)
            if node.FP.Get_kind() == Nodes_NodeKind_get_DeclFunc(){
                return Nodes_NodeVisitMode__Child
            }
            return Nodes_NodeVisitMode__Next
        } else {
            return Nodes_NodeVisitMode__Child
        }
    }
// insert a dummy
    return 0
}
// 339: decl @lune.@base.@ConvNodes.convertNodes
func ConvNodes_convertNodes(targetNode *Nodes_Node) {
    targetNode.FP.Visit(Nodes_NodeVisitor(convertNodes___anonymous_2883_), 0)
}

// declaration Class -- Filter
type ConvNodes_FilterMtd interface {
    ProcessCallExt(arg1 *ConvNodes_CallExtNode, arg2 LnsAny)
    ProcessCondPopVal(arg1 *ConvNodes_CondPopValNode, arg2 LnsAny)
    ProcessCondSetVal(arg1 *ConvNodes_CondSetValNode, arg2 LnsAny)
    ProcessDeclFunc(arg1 *ConvNodes_DeclFuncNode, arg2 LnsAny)
    ProcessExpAccessMRet(arg1 *ConvNodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(arg1 *ConvNodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCast(arg1 *ConvNodes_ExpCastNode, arg2 LnsAny)
    ProcessExpList(arg1 *ConvNodes_ExpListNode, arg2 LnsAny)
    ProcessExpMRet(arg1 *ConvNodes_ExpMRetNode, arg2 LnsAny)
    ProcessExpMultiTo1(arg1 *ConvNodes_ExpMultiTo1Node, arg2 LnsAny)
    ProcessExpNew(arg1 *ConvNodes_ExpNewNode, arg2 LnsAny)
    ProcessExpOmitEnum(arg1 *ConvNodes_ExpOmitEnumNode, arg2 LnsAny)
    ProcessExpOp1(arg1 *ConvNodes_ExpOp1Node, arg2 LnsAny)
    ProcessExpOp2(arg1 *ConvNodes_ExpOp2Node, arg2 LnsAny)
    ProcessExpParen(arg1 *ConvNodes_ExpParenNode, arg2 LnsAny)
    ProcessExpRef(arg1 *ConvNodes_ExpRefNode, arg2 LnsAny)
    ProcessExpRefItem(arg1 *ConvNodes_ExpRefItemNode, arg2 LnsAny)
    ProcessExpSubDDD(arg1 *ConvNodes_ExpSubDDDNode, arg2 LnsAny)
    ProcessExpToDDD(arg1 *ConvNodes_ExpToDDDNode, arg2 LnsAny)
    ProcessExpUnwrap(arg1 *ConvNodes_ExpUnwrapNode, arg2 LnsAny)
    ProcessGetAt(arg1 *ConvNodes_GetAtNode, arg2 LnsAny)
    ProcessGetField(arg1 *ConvNodes_GetFieldNode, arg2 LnsAny)
    ProcessLiteralArray(arg1 *ConvNodes_LiteralArrayNode, arg2 LnsAny)
    ProcessLiteralBool(arg1 *ConvNodes_LiteralBoolNode, arg2 LnsAny)
    ProcessLiteralChar(arg1 *ConvNodes_LiteralCharNode, arg2 LnsAny)
    ProcessLiteralInt(arg1 *ConvNodes_LiteralIntNode, arg2 LnsAny)
    ProcessLiteralList(arg1 *ConvNodes_LiteralListNode, arg2 LnsAny)
    ProcessLiteralMap(arg1 *ConvNodes_LiteralMapNode, arg2 LnsAny)
    ProcessLiteralReal(arg1 *ConvNodes_LiteralRealNode, arg2 LnsAny)
    ProcessLiteralSet(arg1 *ConvNodes_LiteralSetNode, arg2 LnsAny)
    ProcessLiteralString(arg1 *ConvNodes_LiteralStringNode, arg2 LnsAny)
    ProcessLiteralSymbol(arg1 *ConvNodes_LiteralSymbolNode, arg2 LnsAny)
    ProcessLuneKind(arg1 *ConvNodes_LuneKindNode, arg2 LnsAny)
    ProcessNewAlgeVal(arg1 *ConvNodes_NewAlgeValNode, arg2 LnsAny)
    ProcessNilAccFin(arg1 *ConvNodes_NilAccFinNode, arg2 LnsAny)
    ProcessNilAccFinCall(arg1 *ConvNodes_NilAccFinCallNode, arg2 LnsAny)
    ProcessNilAccLast(arg1 *ConvNodes_NilAccLastNode, arg2 LnsAny)
    ProcessNilAccPush(arg1 *ConvNodes_NilAccPushNode, arg2 LnsAny)
    ProcessRefField(arg1 *ConvNodes_RefFieldNode, arg2 LnsAny)
}
type ConvNodes_Filter struct {
    FP ConvNodes_FilterMtd
}
func ConvNodes_Filter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_Filter).FP
}
type ConvNodes_FilterDownCast interface {
    ToConvNodes_Filter() *ConvNodes_Filter
}
func ConvNodes_FilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_FilterDownCast)
    if ok { return work.ToConvNodes_Filter() }
    return nil
}
func (obj *ConvNodes_Filter) ToConvNodes_Filter() *ConvNodes_Filter {
    return obj
}
func (self *ConvNodes_Filter) InitConvNodes_Filter() {
}








































// declaration Class -- Node
type ConvNodes_NodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_Node struct {
    id LnsInt
    pos *Parser_Position
    kind string
    FP ConvNodes_NodeMtd
}
func ConvNodes_Node2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_Node).FP
}
type ConvNodes_NodeDownCast interface {
    ToConvNodes_Node() *ConvNodes_Node
}
func ConvNodes_NodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_NodeDownCast)
    if ok { return work.ToConvNodes_Node() }
    return nil
}
func (obj *ConvNodes_Node) ToConvNodes_Node() *ConvNodes_Node {
    return obj
}
func (self *ConvNodes_Node) InitConvNodes_Node(arg1 LnsInt, arg2 *Parser_Position, arg3 string) {
    self.id = arg1
    self.pos = arg2
    self.kind = arg3
}
func (self *ConvNodes_Node) Get_id() LnsInt{ return self.id }
func (self *ConvNodes_Node) Get_pos() *Parser_Position{ return self.pos }
func (self *ConvNodes_Node) Get_kind() string{ return self.kind }
// 41: decl @lune.@base.@ConvNodes.Node.processFilter
func (self *ConvNodes_Node) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
}


// 52: decl @lune.@base.@ConvNodes.Node.createFrom
func ConvNodes_Node_createFrom(node *Nodes_Node) *ConvNodes_Node {
    var _func ConvNodes_createFromFunc_1063_
    
    {
        _func := ConvNodes_nodeKind2createFromFunc.Items[node.FP.Get_kind()]
        if _func == nil{
            Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{Nodes_getNodeKindName(node.FP.Get_kind())}))
        } else {
            _func = _func.(ConvNodes_createFromFunc_1063_)
        }
    }
    return _func(node)
}


// declaration Class -- ExpNode
type ConvNodes_ExpNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpNode struct {
    ConvNodes_Node
    FP ConvNodes_ExpNodeMtd
}
func ConvNodes_ExpNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpNode).FP
}
type ConvNodes_ExpNodeDownCast interface {
    ToConvNodes_ExpNode() *ConvNodes_ExpNode
}
func ConvNodes_ExpNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpNodeDownCast)
    if ok { return work.ToConvNodes_ExpNode() }
    return nil
}
func (obj *ConvNodes_ExpNode) ToConvNodes_ExpNode() *ConvNodes_ExpNode {
    return obj
}
func (self *ConvNodes_ExpNode) InitConvNodes_ExpNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string) {
    self.ConvNodes_Node.InitConvNodes_Node( arg1,arg2,arg3)
}

// declaration Class -- ExpListNode
type ConvNodes_ExpListNodeMtd interface {
    Get_expList() *LnsList
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpListNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpListNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpListNode
    expList *LnsList
    FP ConvNodes_ExpListNodeMtd
}
func ConvNodes_ExpListNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpListNode).FP
}
type ConvNodes_ExpListNodeDownCast interface {
    ToConvNodes_ExpListNode() *ConvNodes_ExpListNode
}
func ConvNodes_ExpListNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpListNodeDownCast)
    if ok { return work.ToConvNodes_ExpListNode() }
    return nil
}
func (obj *ConvNodes_ExpListNode) ToConvNodes_ExpListNode() *ConvNodes_ExpListNode {
    return obj
}
func NewConvNodes_ExpListNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpListNode, arg5 *LnsList) *ConvNodes_ExpListNode {
    obj := &ConvNodes_ExpListNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpListNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpListNode) InitConvNodes_ExpListNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpListNode, arg5 *LnsList) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.expList = arg5
}
func (self *ConvNodes_ExpListNode) Get_orgNode() *Nodes_ExpListNode{ return self.orgNode }
func (self *ConvNodes_ExpListNode) Get_expList() *LnsList{ return self.expList }
// 1: decl @lune.@base.@ConvNodes.ExpListNode.processFilter
func (self *ConvNodes_ExpListNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpList(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpListNode.visit
func (self *ConvNodes_ExpListNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.expList
        for _, _child := range( list.Items ) {
            child := _child.(ConvNodes_NodeDownCast).ToConvNodes_Node()
            if _switch1206 := visitor(child, &self.ConvNodes_Node, "expList", depth); _switch1206 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch1206 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}


// declaration Class -- ExpNewNode
type ConvNodes_ExpNewNodeMtd interface {
    Get_argList() LnsAny
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpNewNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpNewNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpNewNode
    argList LnsAny
    FP ConvNodes_ExpNewNodeMtd
}
func ConvNodes_ExpNewNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpNewNode).FP
}
type ConvNodes_ExpNewNodeDownCast interface {
    ToConvNodes_ExpNewNode() *ConvNodes_ExpNewNode
}
func ConvNodes_ExpNewNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpNewNodeDownCast)
    if ok { return work.ToConvNodes_ExpNewNode() }
    return nil
}
func (obj *ConvNodes_ExpNewNode) ToConvNodes_ExpNewNode() *ConvNodes_ExpNewNode {
    return obj
}
func NewConvNodes_ExpNewNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpNewNode, arg5 LnsAny) *ConvNodes_ExpNewNode {
    obj := &ConvNodes_ExpNewNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpNewNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpNewNode) InitConvNodes_ExpNewNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpNewNode, arg5 LnsAny) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.argList = arg5
}
func (self *ConvNodes_ExpNewNode) Get_orgNode() *Nodes_ExpNewNode{ return self.orgNode }
func (self *ConvNodes_ExpNewNode) Get_argList() LnsAny{ return self.argList }
// 1: decl @lune.@base.@ConvNodes.ExpNewNode.processFilter
func (self *ConvNodes_ExpNewNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpNew(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpNewNode.visit
func (self *ConvNodes_ExpNewNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.argList
            if _child != nil {
                child := _child.(*ConvNodes_ExpListNode)
                if _switch1490 := visitor(&child.ConvNodes_Node, &self.ConvNodes_Node, "argList", depth); _switch1490 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch1490 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- ExpUnwrapNode
type ConvNodes_ExpUnwrapNodeMtd interface {
    Get_default() LnsAny
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpUnwrapNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpUnwrapNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpUnwrapNode
    exp *ConvNodes_Node
    _default LnsAny
    FP ConvNodes_ExpUnwrapNodeMtd
}
func ConvNodes_ExpUnwrapNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpUnwrapNode).FP
}
type ConvNodes_ExpUnwrapNodeDownCast interface {
    ToConvNodes_ExpUnwrapNode() *ConvNodes_ExpUnwrapNode
}
func ConvNodes_ExpUnwrapNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpUnwrapNodeDownCast)
    if ok { return work.ToConvNodes_ExpUnwrapNode() }
    return nil
}
func (obj *ConvNodes_ExpUnwrapNode) ToConvNodes_ExpUnwrapNode() *ConvNodes_ExpUnwrapNode {
    return obj
}
func NewConvNodes_ExpUnwrapNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpUnwrapNode, arg5 *ConvNodes_Node, arg6 LnsAny) *ConvNodes_ExpUnwrapNode {
    obj := &ConvNodes_ExpUnwrapNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpUnwrapNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *ConvNodes_ExpUnwrapNode) InitConvNodes_ExpUnwrapNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpUnwrapNode, arg5 *ConvNodes_Node, arg6 LnsAny) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.exp = arg5
    self._default = arg6
}
func (self *ConvNodes_ExpUnwrapNode) Get_orgNode() *Nodes_ExpUnwrapNode{ return self.orgNode }
func (self *ConvNodes_ExpUnwrapNode) Get_exp() *ConvNodes_Node{ return self.exp }
func (self *ConvNodes_ExpUnwrapNode) Get_default() LnsAny{ return self._default }
// 1: decl @lune.@base.@ConvNodes.ExpUnwrapNode.processFilter
func (self *ConvNodes_ExpUnwrapNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpUnwrap(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpUnwrapNode.visit
func (self *ConvNodes_ExpUnwrapNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch1763 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch1763 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch1763 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self._default
            if _child != nil {
                child := _child.(*ConvNodes_Node)
                if _switch1816 := visitor(child, &self.ConvNodes_Node, "default", depth); _switch1816 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch1816 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- ExpRefNode
type ConvNodes_ExpRefNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpRefNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpRefNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpRefNode
    FP ConvNodes_ExpRefNodeMtd
}
func ConvNodes_ExpRefNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpRefNode).FP
}
type ConvNodes_ExpRefNodeDownCast interface {
    ToConvNodes_ExpRefNode() *ConvNodes_ExpRefNode
}
func ConvNodes_ExpRefNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpRefNodeDownCast)
    if ok { return work.ToConvNodes_ExpRefNode() }
    return nil
}
func (obj *ConvNodes_ExpRefNode) ToConvNodes_ExpRefNode() *ConvNodes_ExpRefNode {
    return obj
}
func NewConvNodes_ExpRefNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpRefNode) *ConvNodes_ExpRefNode {
    obj := &ConvNodes_ExpRefNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpRefNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_ExpRefNode) InitConvNodes_ExpRefNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpRefNode) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
}
func (self *ConvNodes_ExpRefNode) Get_orgNode() *Nodes_ExpRefNode{ return self.orgNode }
// 1: decl @lune.@base.@ConvNodes.ExpRefNode.processFilter
func (self *ConvNodes_ExpRefNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpRef(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpRefNode.visit
func (self *ConvNodes_ExpRefNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- ExpOp2Node
type ConvNodes_ExpOp2NodeMtd interface {
    Get_exp1() *ConvNodes_Node
    Get_exp2() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpOp2Node
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpOp2Node struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpOp2Node
    exp1 *ConvNodes_Node
    exp2 *ConvNodes_Node
    FP ConvNodes_ExpOp2NodeMtd
}
func ConvNodes_ExpOp2Node2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpOp2Node).FP
}
type ConvNodes_ExpOp2NodeDownCast interface {
    ToConvNodes_ExpOp2Node() *ConvNodes_ExpOp2Node
}
func ConvNodes_ExpOp2NodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpOp2NodeDownCast)
    if ok { return work.ToConvNodes_ExpOp2Node() }
    return nil
}
func (obj *ConvNodes_ExpOp2Node) ToConvNodes_ExpOp2Node() *ConvNodes_ExpOp2Node {
    return obj
}
func NewConvNodes_ExpOp2Node(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpOp2Node, arg5 *ConvNodes_Node, arg6 *ConvNodes_Node) *ConvNodes_ExpOp2Node {
    obj := &ConvNodes_ExpOp2Node{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpOp2Node(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *ConvNodes_ExpOp2Node) InitConvNodes_ExpOp2Node(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpOp2Node, arg5 *ConvNodes_Node, arg6 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.exp1 = arg5
    self.exp2 = arg6
}
func (self *ConvNodes_ExpOp2Node) Get_orgNode() *Nodes_ExpOp2Node{ return self.orgNode }
func (self *ConvNodes_ExpOp2Node) Get_exp1() *ConvNodes_Node{ return self.exp1 }
func (self *ConvNodes_ExpOp2Node) Get_exp2() *ConvNodes_Node{ return self.exp2 }
// 1: decl @lune.@base.@ConvNodes.ExpOp2Node.processFilter
func (self *ConvNodes_ExpOp2Node) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpOp2(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpOp2Node.visit
func (self *ConvNodes_ExpOp2Node) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp1
        if _switch2275 := visitor(child, &self.ConvNodes_Node, "exp1", depth); _switch2275 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch2275 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *ConvNodes_Node
        child = self.exp2
        if _switch2329 := visitor(child, &self.ConvNodes_Node, "exp2", depth); _switch2329 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch2329 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpCastNode
type ConvNodes_ExpCastNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpCastNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpCastNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpCastNode
    exp *ConvNodes_Node
    FP ConvNodes_ExpCastNodeMtd
}
func ConvNodes_ExpCastNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpCastNode).FP
}
type ConvNodes_ExpCastNodeDownCast interface {
    ToConvNodes_ExpCastNode() *ConvNodes_ExpCastNode
}
func ConvNodes_ExpCastNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpCastNodeDownCast)
    if ok { return work.ToConvNodes_ExpCastNode() }
    return nil
}
func (obj *ConvNodes_ExpCastNode) ToConvNodes_ExpCastNode() *ConvNodes_ExpCastNode {
    return obj
}
func NewConvNodes_ExpCastNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpCastNode, arg5 *ConvNodes_Node) *ConvNodes_ExpCastNode {
    obj := &ConvNodes_ExpCastNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpCastNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpCastNode) InitConvNodes_ExpCastNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpCastNode, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.exp = arg5
}
func (self *ConvNodes_ExpCastNode) Get_orgNode() *Nodes_ExpCastNode{ return self.orgNode }
func (self *ConvNodes_ExpCastNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.ExpCastNode.processFilter
func (self *ConvNodes_ExpCastNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCast(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpCastNode.visit
func (self *ConvNodes_ExpCastNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch2603 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch2603 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch2603 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpToDDDNode
type ConvNodes_ExpToDDDNodeMtd interface {
    Get_expList() *ConvNodes_ExpListNode
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpToDDDNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpToDDDNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpToDDDNode
    expList *ConvNodes_ExpListNode
    FP ConvNodes_ExpToDDDNodeMtd
}
func ConvNodes_ExpToDDDNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpToDDDNode).FP
}
type ConvNodes_ExpToDDDNodeDownCast interface {
    ToConvNodes_ExpToDDDNode() *ConvNodes_ExpToDDDNode
}
func ConvNodes_ExpToDDDNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpToDDDNodeDownCast)
    if ok { return work.ToConvNodes_ExpToDDDNode() }
    return nil
}
func (obj *ConvNodes_ExpToDDDNode) ToConvNodes_ExpToDDDNode() *ConvNodes_ExpToDDDNode {
    return obj
}
func NewConvNodes_ExpToDDDNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpToDDDNode, arg5 *ConvNodes_ExpListNode) *ConvNodes_ExpToDDDNode {
    obj := &ConvNodes_ExpToDDDNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpToDDDNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpToDDDNode) InitConvNodes_ExpToDDDNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpToDDDNode, arg5 *ConvNodes_ExpListNode) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.expList = arg5
}
func (self *ConvNodes_ExpToDDDNode) Get_orgNode() *Nodes_ExpToDDDNode{ return self.orgNode }
func (self *ConvNodes_ExpToDDDNode) Get_expList() *ConvNodes_ExpListNode{ return self.expList }
// 1: decl @lune.@base.@ConvNodes.ExpToDDDNode.processFilter
func (self *ConvNodes_ExpToDDDNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpToDDD(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpToDDDNode.visit
func (self *ConvNodes_ExpToDDDNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_ExpListNode
        child = self.expList
        if _switch2857 := visitor(&child.ConvNodes_Node, &self.ConvNodes_Node, "expList", depth); _switch2857 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch2857 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpSubDDDNode
type ConvNodes_ExpSubDDDNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpSubDDDNode
    Get_pos() *Parser_Position
    Get_src() *ConvNodes_Node
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpSubDDDNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpSubDDDNode
    src *ConvNodes_Node
    FP ConvNodes_ExpSubDDDNodeMtd
}
func ConvNodes_ExpSubDDDNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpSubDDDNode).FP
}
type ConvNodes_ExpSubDDDNodeDownCast interface {
    ToConvNodes_ExpSubDDDNode() *ConvNodes_ExpSubDDDNode
}
func ConvNodes_ExpSubDDDNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpSubDDDNodeDownCast)
    if ok { return work.ToConvNodes_ExpSubDDDNode() }
    return nil
}
func (obj *ConvNodes_ExpSubDDDNode) ToConvNodes_ExpSubDDDNode() *ConvNodes_ExpSubDDDNode {
    return obj
}
func NewConvNodes_ExpSubDDDNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpSubDDDNode, arg5 *ConvNodes_Node) *ConvNodes_ExpSubDDDNode {
    obj := &ConvNodes_ExpSubDDDNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpSubDDDNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpSubDDDNode) InitConvNodes_ExpSubDDDNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpSubDDDNode, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.src = arg5
}
func (self *ConvNodes_ExpSubDDDNode) Get_orgNode() *Nodes_ExpSubDDDNode{ return self.orgNode }
func (self *ConvNodes_ExpSubDDDNode) Get_src() *ConvNodes_Node{ return self.src }
// 1: decl @lune.@base.@ConvNodes.ExpSubDDDNode.processFilter
func (self *ConvNodes_ExpSubDDDNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpSubDDD(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpSubDDDNode.visit
func (self *ConvNodes_ExpSubDDDNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.src
        if _switch3112 := visitor(child, &self.ConvNodes_Node, "src", depth); _switch3112 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch3112 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpOp1Node
type ConvNodes_ExpOp1NodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpOp1Node
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpOp1Node struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpOp1Node
    exp *ConvNodes_Node
    FP ConvNodes_ExpOp1NodeMtd
}
func ConvNodes_ExpOp1Node2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpOp1Node).FP
}
type ConvNodes_ExpOp1NodeDownCast interface {
    ToConvNodes_ExpOp1Node() *ConvNodes_ExpOp1Node
}
func ConvNodes_ExpOp1NodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpOp1NodeDownCast)
    if ok { return work.ToConvNodes_ExpOp1Node() }
    return nil
}
func (obj *ConvNodes_ExpOp1Node) ToConvNodes_ExpOp1Node() *ConvNodes_ExpOp1Node {
    return obj
}
func NewConvNodes_ExpOp1Node(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpOp1Node, arg5 *ConvNodes_Node) *ConvNodes_ExpOp1Node {
    obj := &ConvNodes_ExpOp1Node{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpOp1Node(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpOp1Node) InitConvNodes_ExpOp1Node(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpOp1Node, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.exp = arg5
}
func (self *ConvNodes_ExpOp1Node) Get_orgNode() *Nodes_ExpOp1Node{ return self.orgNode }
func (self *ConvNodes_ExpOp1Node) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.ExpOp1Node.processFilter
func (self *ConvNodes_ExpOp1Node) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpOp1(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpOp1Node.visit
func (self *ConvNodes_ExpOp1Node) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch3365 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch3365 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch3365 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpRefItemNode
type ConvNodes_ExpRefItemNodeMtd interface {
    Get_id() LnsInt
    Get_index() LnsAny
    Get_kind() string
    Get_orgNode() *Nodes_ExpRefItemNode
    Get_pos() *Parser_Position
    Get_val() *ConvNodes_Node
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpRefItemNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpRefItemNode
    val *ConvNodes_Node
    index LnsAny
    FP ConvNodes_ExpRefItemNodeMtd
}
func ConvNodes_ExpRefItemNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpRefItemNode).FP
}
type ConvNodes_ExpRefItemNodeDownCast interface {
    ToConvNodes_ExpRefItemNode() *ConvNodes_ExpRefItemNode
}
func ConvNodes_ExpRefItemNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpRefItemNodeDownCast)
    if ok { return work.ToConvNodes_ExpRefItemNode() }
    return nil
}
func (obj *ConvNodes_ExpRefItemNode) ToConvNodes_ExpRefItemNode() *ConvNodes_ExpRefItemNode {
    return obj
}
func NewConvNodes_ExpRefItemNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpRefItemNode, arg5 *ConvNodes_Node, arg6 LnsAny) *ConvNodes_ExpRefItemNode {
    obj := &ConvNodes_ExpRefItemNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpRefItemNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *ConvNodes_ExpRefItemNode) InitConvNodes_ExpRefItemNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpRefItemNode, arg5 *ConvNodes_Node, arg6 LnsAny) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.val = arg5
    self.index = arg6
}
func (self *ConvNodes_ExpRefItemNode) Get_orgNode() *Nodes_ExpRefItemNode{ return self.orgNode }
func (self *ConvNodes_ExpRefItemNode) Get_val() *ConvNodes_Node{ return self.val }
func (self *ConvNodes_ExpRefItemNode) Get_index() LnsAny{ return self.index }
// 1: decl @lune.@base.@ConvNodes.ExpRefItemNode.processFilter
func (self *ConvNodes_ExpRefItemNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpRefItem(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpRefItemNode.visit
func (self *ConvNodes_ExpRefItemNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.val
        if _switch3629 := visitor(child, &self.ConvNodes_Node, "val", depth); _switch3629 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch3629 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.index
            if _child != nil {
                child := _child.(*ConvNodes_Node)
                if _switch3682 := visitor(child, &self.ConvNodes_Node, "index", depth); _switch3682 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch3682 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- ExpCallNode
type ConvNodes_ExpCallNodeMtd interface {
    Get_argList() LnsAny
    Get_func() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpCallNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpCallNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpCallNode
    _func *ConvNodes_Node
    argList LnsAny
    FP ConvNodes_ExpCallNodeMtd
}
func ConvNodes_ExpCallNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpCallNode).FP
}
type ConvNodes_ExpCallNodeDownCast interface {
    ToConvNodes_ExpCallNode() *ConvNodes_ExpCallNode
}
func ConvNodes_ExpCallNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpCallNodeDownCast)
    if ok { return work.ToConvNodes_ExpCallNode() }
    return nil
}
func (obj *ConvNodes_ExpCallNode) ToConvNodes_ExpCallNode() *ConvNodes_ExpCallNode {
    return obj
}
func NewConvNodes_ExpCallNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpCallNode, arg5 *ConvNodes_Node, arg6 LnsAny) *ConvNodes_ExpCallNode {
    obj := &ConvNodes_ExpCallNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpCallNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *ConvNodes_ExpCallNode) InitConvNodes_ExpCallNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpCallNode, arg5 *ConvNodes_Node, arg6 LnsAny) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self._func = arg5
    self.argList = arg6
}
func (self *ConvNodes_ExpCallNode) Get_orgNode() *Nodes_ExpCallNode{ return self.orgNode }
func (self *ConvNodes_ExpCallNode) Get_func() *ConvNodes_Node{ return self._func }
func (self *ConvNodes_ExpCallNode) Get_argList() LnsAny{ return self.argList }
// 1: decl @lune.@base.@ConvNodes.ExpCallNode.processFilter
func (self *ConvNodes_ExpCallNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCall(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpCallNode.visit
func (self *ConvNodes_ExpCallNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self._func
        if _switch3974 := visitor(child, &self.ConvNodes_Node, "func", depth); _switch3974 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch3974 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.argList
            if _child != nil {
                child := _child.(*ConvNodes_ExpListNode)
                if _switch4028 := visitor(&child.ConvNodes_Node, &self.ConvNodes_Node, "argList", depth); _switch4028 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch4028 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- ExpMRetNode
type ConvNodes_ExpMRetNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_mRet() *ConvNodes_Node
    Get_orgNode() *Nodes_ExpMRetNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpMRetNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpMRetNode
    mRet *ConvNodes_Node
    FP ConvNodes_ExpMRetNodeMtd
}
func ConvNodes_ExpMRetNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpMRetNode).FP
}
type ConvNodes_ExpMRetNodeDownCast interface {
    ToConvNodes_ExpMRetNode() *ConvNodes_ExpMRetNode
}
func ConvNodes_ExpMRetNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpMRetNodeDownCast)
    if ok { return work.ToConvNodes_ExpMRetNode() }
    return nil
}
func (obj *ConvNodes_ExpMRetNode) ToConvNodes_ExpMRetNode() *ConvNodes_ExpMRetNode {
    return obj
}
func NewConvNodes_ExpMRetNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpMRetNode, arg5 *ConvNodes_Node) *ConvNodes_ExpMRetNode {
    obj := &ConvNodes_ExpMRetNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpMRetNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpMRetNode) InitConvNodes_ExpMRetNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpMRetNode, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.mRet = arg5
}
func (self *ConvNodes_ExpMRetNode) Get_orgNode() *Nodes_ExpMRetNode{ return self.orgNode }
func (self *ConvNodes_ExpMRetNode) Get_mRet() *ConvNodes_Node{ return self.mRet }
// 1: decl @lune.@base.@ConvNodes.ExpMRetNode.processFilter
func (self *ConvNodes_ExpMRetNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMRet(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpMRetNode.visit
func (self *ConvNodes_ExpMRetNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.mRet
        if _switch4311 := visitor(child, &self.ConvNodes_Node, "mRet", depth); _switch4311 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch4311 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpAccessMRetNode
type ConvNodes_ExpAccessMRetNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_mRet() *ConvNodes_Node
    Get_orgNode() *Nodes_ExpAccessMRetNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpAccessMRetNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpAccessMRetNode
    mRet *ConvNodes_Node
    FP ConvNodes_ExpAccessMRetNodeMtd
}
func ConvNodes_ExpAccessMRetNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpAccessMRetNode).FP
}
type ConvNodes_ExpAccessMRetNodeDownCast interface {
    ToConvNodes_ExpAccessMRetNode() *ConvNodes_ExpAccessMRetNode
}
func ConvNodes_ExpAccessMRetNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpAccessMRetNodeDownCast)
    if ok { return work.ToConvNodes_ExpAccessMRetNode() }
    return nil
}
func (obj *ConvNodes_ExpAccessMRetNode) ToConvNodes_ExpAccessMRetNode() *ConvNodes_ExpAccessMRetNode {
    return obj
}
func NewConvNodes_ExpAccessMRetNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpAccessMRetNode, arg5 *ConvNodes_Node) *ConvNodes_ExpAccessMRetNode {
    obj := &ConvNodes_ExpAccessMRetNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpAccessMRetNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpAccessMRetNode) InitConvNodes_ExpAccessMRetNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpAccessMRetNode, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.mRet = arg5
}
func (self *ConvNodes_ExpAccessMRetNode) Get_orgNode() *Nodes_ExpAccessMRetNode{ return self.orgNode }
func (self *ConvNodes_ExpAccessMRetNode) Get_mRet() *ConvNodes_Node{ return self.mRet }
// 1: decl @lune.@base.@ConvNodes.ExpAccessMRetNode.processFilter
func (self *ConvNodes_ExpAccessMRetNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpAccessMRet(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpAccessMRetNode.visit
func (self *ConvNodes_ExpAccessMRetNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.mRet
        if _switch4564 := visitor(child, &self.ConvNodes_Node, "mRet", depth); _switch4564 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch4564 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpMultiTo1Node
type ConvNodes_ExpMultiTo1NodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpMultiTo1Node
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpMultiTo1Node struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpMultiTo1Node
    exp *ConvNodes_Node
    FP ConvNodes_ExpMultiTo1NodeMtd
}
func ConvNodes_ExpMultiTo1Node2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpMultiTo1Node).FP
}
type ConvNodes_ExpMultiTo1NodeDownCast interface {
    ToConvNodes_ExpMultiTo1Node() *ConvNodes_ExpMultiTo1Node
}
func ConvNodes_ExpMultiTo1NodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpMultiTo1NodeDownCast)
    if ok { return work.ToConvNodes_ExpMultiTo1Node() }
    return nil
}
func (obj *ConvNodes_ExpMultiTo1Node) ToConvNodes_ExpMultiTo1Node() *ConvNodes_ExpMultiTo1Node {
    return obj
}
func NewConvNodes_ExpMultiTo1Node(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpMultiTo1Node, arg5 *ConvNodes_Node) *ConvNodes_ExpMultiTo1Node {
    obj := &ConvNodes_ExpMultiTo1Node{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpMultiTo1Node(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpMultiTo1Node) InitConvNodes_ExpMultiTo1Node(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpMultiTo1Node, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.exp = arg5
}
func (self *ConvNodes_ExpMultiTo1Node) Get_orgNode() *Nodes_ExpMultiTo1Node{ return self.orgNode }
func (self *ConvNodes_ExpMultiTo1Node) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.ExpMultiTo1Node.processFilter
func (self *ConvNodes_ExpMultiTo1Node) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMultiTo1(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpMultiTo1Node.visit
func (self *ConvNodes_ExpMultiTo1Node) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch4817 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch4817 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch4817 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpParenNode
type ConvNodes_ExpParenNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpParenNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpParenNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpParenNode
    exp *ConvNodes_Node
    FP ConvNodes_ExpParenNodeMtd
}
func ConvNodes_ExpParenNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpParenNode).FP
}
type ConvNodes_ExpParenNodeDownCast interface {
    ToConvNodes_ExpParenNode() *ConvNodes_ExpParenNode
}
func ConvNodes_ExpParenNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpParenNodeDownCast)
    if ok { return work.ToConvNodes_ExpParenNode() }
    return nil
}
func (obj *ConvNodes_ExpParenNode) ToConvNodes_ExpParenNode() *ConvNodes_ExpParenNode {
    return obj
}
func NewConvNodes_ExpParenNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpParenNode, arg5 *ConvNodes_Node) *ConvNodes_ExpParenNode {
    obj := &ConvNodes_ExpParenNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpParenNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_ExpParenNode) InitConvNodes_ExpParenNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpParenNode, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.exp = arg5
}
func (self *ConvNodes_ExpParenNode) Get_orgNode() *Nodes_ExpParenNode{ return self.orgNode }
func (self *ConvNodes_ExpParenNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.ExpParenNode.processFilter
func (self *ConvNodes_ExpParenNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpParen(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpParenNode.visit
func (self *ConvNodes_ExpParenNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch5070 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch5070 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch5070 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpOmitEnumNode
type ConvNodes_ExpOmitEnumNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_ExpOmitEnumNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_ExpOmitEnumNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_ExpOmitEnumNode
    FP ConvNodes_ExpOmitEnumNodeMtd
}
func ConvNodes_ExpOmitEnumNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_ExpOmitEnumNode).FP
}
type ConvNodes_ExpOmitEnumNodeDownCast interface {
    ToConvNodes_ExpOmitEnumNode() *ConvNodes_ExpOmitEnumNode
}
func ConvNodes_ExpOmitEnumNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_ExpOmitEnumNodeDownCast)
    if ok { return work.ToConvNodes_ExpOmitEnumNode() }
    return nil
}
func (obj *ConvNodes_ExpOmitEnumNode) ToConvNodes_ExpOmitEnumNode() *ConvNodes_ExpOmitEnumNode {
    return obj
}
func NewConvNodes_ExpOmitEnumNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpOmitEnumNode) *ConvNodes_ExpOmitEnumNode {
    obj := &ConvNodes_ExpOmitEnumNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_ExpOmitEnumNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_ExpOmitEnumNode) InitConvNodes_ExpOmitEnumNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_ExpOmitEnumNode) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
}
func (self *ConvNodes_ExpOmitEnumNode) Get_orgNode() *Nodes_ExpOmitEnumNode{ return self.orgNode }
// 1: decl @lune.@base.@ConvNodes.ExpOmitEnumNode.processFilter
func (self *ConvNodes_ExpOmitEnumNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpOmitEnum(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.ExpOmitEnumNode.visit
func (self *ConvNodes_ExpOmitEnumNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- RefFieldNode
type ConvNodes_RefFieldNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_RefFieldNode
    Get_pos() *Parser_Position
    Get_prefix() *ConvNodes_Node
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_RefFieldNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_RefFieldNode
    prefix *ConvNodes_Node
    FP ConvNodes_RefFieldNodeMtd
}
func ConvNodes_RefFieldNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_RefFieldNode).FP
}
type ConvNodes_RefFieldNodeDownCast interface {
    ToConvNodes_RefFieldNode() *ConvNodes_RefFieldNode
}
func ConvNodes_RefFieldNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_RefFieldNodeDownCast)
    if ok { return work.ToConvNodes_RefFieldNode() }
    return nil
}
func (obj *ConvNodes_RefFieldNode) ToConvNodes_RefFieldNode() *ConvNodes_RefFieldNode {
    return obj
}
func NewConvNodes_RefFieldNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_RefFieldNode, arg5 *ConvNodes_Node) *ConvNodes_RefFieldNode {
    obj := &ConvNodes_RefFieldNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_RefFieldNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_RefFieldNode) InitConvNodes_RefFieldNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_RefFieldNode, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.prefix = arg5
}
func (self *ConvNodes_RefFieldNode) Get_orgNode() *Nodes_RefFieldNode{ return self.orgNode }
func (self *ConvNodes_RefFieldNode) Get_prefix() *ConvNodes_Node{ return self.prefix }
// 1: decl @lune.@base.@ConvNodes.RefFieldNode.processFilter
func (self *ConvNodes_RefFieldNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRefField(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.RefFieldNode.visit
func (self *ConvNodes_RefFieldNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.prefix
        if _switch5490 := visitor(child, &self.ConvNodes_Node, "prefix", depth); _switch5490 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch5490 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- GetFieldNode
type ConvNodes_GetFieldNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_GetFieldNode
    Get_pos() *Parser_Position
    Get_prefix() *ConvNodes_Node
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_GetFieldNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_GetFieldNode
    prefix *ConvNodes_Node
    FP ConvNodes_GetFieldNodeMtd
}
func ConvNodes_GetFieldNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_GetFieldNode).FP
}
type ConvNodes_GetFieldNodeDownCast interface {
    ToConvNodes_GetFieldNode() *ConvNodes_GetFieldNode
}
func ConvNodes_GetFieldNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_GetFieldNodeDownCast)
    if ok { return work.ToConvNodes_GetFieldNode() }
    return nil
}
func (obj *ConvNodes_GetFieldNode) ToConvNodes_GetFieldNode() *ConvNodes_GetFieldNode {
    return obj
}
func NewConvNodes_GetFieldNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_GetFieldNode, arg5 *ConvNodes_Node) *ConvNodes_GetFieldNode {
    obj := &ConvNodes_GetFieldNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_GetFieldNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_GetFieldNode) InitConvNodes_GetFieldNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_GetFieldNode, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.prefix = arg5
}
func (self *ConvNodes_GetFieldNode) Get_orgNode() *Nodes_GetFieldNode{ return self.orgNode }
func (self *ConvNodes_GetFieldNode) Get_prefix() *ConvNodes_Node{ return self.prefix }
// 1: decl @lune.@base.@ConvNodes.GetFieldNode.processFilter
func (self *ConvNodes_GetFieldNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessGetField(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.GetFieldNode.visit
func (self *ConvNodes_GetFieldNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.prefix
        if _switch5743 := visitor(child, &self.ConvNodes_Node, "prefix", depth); _switch5743 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch5743 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- DeclFuncNode
type ConvNodes_DeclFuncNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_DeclFuncNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_DeclFuncNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_DeclFuncNode
    FP ConvNodes_DeclFuncNodeMtd
}
func ConvNodes_DeclFuncNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_DeclFuncNode).FP
}
type ConvNodes_DeclFuncNodeDownCast interface {
    ToConvNodes_DeclFuncNode() *ConvNodes_DeclFuncNode
}
func ConvNodes_DeclFuncNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_DeclFuncNodeDownCast)
    if ok { return work.ToConvNodes_DeclFuncNode() }
    return nil
}
func (obj *ConvNodes_DeclFuncNode) ToConvNodes_DeclFuncNode() *ConvNodes_DeclFuncNode {
    return obj
}
func NewConvNodes_DeclFuncNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_DeclFuncNode) *ConvNodes_DeclFuncNode {
    obj := &ConvNodes_DeclFuncNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_DeclFuncNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_DeclFuncNode) InitConvNodes_DeclFuncNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_DeclFuncNode) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
}
func (self *ConvNodes_DeclFuncNode) Get_orgNode() *Nodes_DeclFuncNode{ return self.orgNode }
// 1: decl @lune.@base.@ConvNodes.DeclFuncNode.processFilter
func (self *ConvNodes_DeclFuncNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclFunc(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.DeclFuncNode.visit
func (self *ConvNodes_DeclFuncNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- NewAlgeValNode
type ConvNodes_NewAlgeValNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_NewAlgeValNode
    Get_paramList() *LnsList
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_NewAlgeValNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_NewAlgeValNode
    paramList *LnsList
    FP ConvNodes_NewAlgeValNodeMtd
}
func ConvNodes_NewAlgeValNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_NewAlgeValNode).FP
}
type ConvNodes_NewAlgeValNodeDownCast interface {
    ToConvNodes_NewAlgeValNode() *ConvNodes_NewAlgeValNode
}
func ConvNodes_NewAlgeValNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_NewAlgeValNodeDownCast)
    if ok { return work.ToConvNodes_NewAlgeValNode() }
    return nil
}
func (obj *ConvNodes_NewAlgeValNode) ToConvNodes_NewAlgeValNode() *ConvNodes_NewAlgeValNode {
    return obj
}
func NewConvNodes_NewAlgeValNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_NewAlgeValNode, arg5 *LnsList) *ConvNodes_NewAlgeValNode {
    obj := &ConvNodes_NewAlgeValNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_NewAlgeValNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_NewAlgeValNode) InitConvNodes_NewAlgeValNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_NewAlgeValNode, arg5 *LnsList) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.paramList = arg5
}
func (self *ConvNodes_NewAlgeValNode) Get_orgNode() *Nodes_NewAlgeValNode{ return self.orgNode }
func (self *ConvNodes_NewAlgeValNode) Get_paramList() *LnsList{ return self.paramList }
// 1: decl @lune.@base.@ConvNodes.NewAlgeValNode.processFilter
func (self *ConvNodes_NewAlgeValNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessNewAlgeVal(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.NewAlgeValNode.visit
func (self *ConvNodes_NewAlgeValNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.paramList
        for _, _child := range( list.Items ) {
            child := _child.(ConvNodes_NodeDownCast).ToConvNodes_Node()
            if _switch6167 := visitor(child, &self.ConvNodes_Node, "paramList", depth); _switch6167 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch6167 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}


// declaration Class -- LuneKindNode
type ConvNodes_LuneKindNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LuneKindNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LuneKindNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LuneKindNode
    exp *ConvNodes_Node
    FP ConvNodes_LuneKindNodeMtd
}
func ConvNodes_LuneKindNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LuneKindNode).FP
}
type ConvNodes_LuneKindNodeDownCast interface {
    ToConvNodes_LuneKindNode() *ConvNodes_LuneKindNode
}
func ConvNodes_LuneKindNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LuneKindNodeDownCast)
    if ok { return work.ToConvNodes_LuneKindNode() }
    return nil
}
func (obj *ConvNodes_LuneKindNode) ToConvNodes_LuneKindNode() *ConvNodes_LuneKindNode {
    return obj
}
func NewConvNodes_LuneKindNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LuneKindNode, arg5 *ConvNodes_Node) *ConvNodes_LuneKindNode {
    obj := &ConvNodes_LuneKindNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LuneKindNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_LuneKindNode) InitConvNodes_LuneKindNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LuneKindNode, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.exp = arg5
}
func (self *ConvNodes_LuneKindNode) Get_orgNode() *Nodes_LuneKindNode{ return self.orgNode }
func (self *ConvNodes_LuneKindNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.LuneKindNode.processFilter
func (self *ConvNodes_LuneKindNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLuneKind(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LuneKindNode.visit
func (self *ConvNodes_LuneKindNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch6450 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch6450 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch6450 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- LiteralCharNode
type ConvNodes_LiteralCharNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralCharNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralCharNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralCharNode
    FP ConvNodes_LiteralCharNodeMtd
}
func ConvNodes_LiteralCharNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralCharNode).FP
}
type ConvNodes_LiteralCharNodeDownCast interface {
    ToConvNodes_LiteralCharNode() *ConvNodes_LiteralCharNode
}
func ConvNodes_LiteralCharNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralCharNodeDownCast)
    if ok { return work.ToConvNodes_LiteralCharNode() }
    return nil
}
func (obj *ConvNodes_LiteralCharNode) ToConvNodes_LiteralCharNode() *ConvNodes_LiteralCharNode {
    return obj
}
func NewConvNodes_LiteralCharNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralCharNode) *ConvNodes_LiteralCharNode {
    obj := &ConvNodes_LiteralCharNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralCharNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_LiteralCharNode) InitConvNodes_LiteralCharNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralCharNode) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
}
func (self *ConvNodes_LiteralCharNode) Get_orgNode() *Nodes_LiteralCharNode{ return self.orgNode }
// 1: decl @lune.@base.@ConvNodes.LiteralCharNode.processFilter
func (self *ConvNodes_LiteralCharNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralChar(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralCharNode.visit
func (self *ConvNodes_LiteralCharNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- LiteralIntNode
type ConvNodes_LiteralIntNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralIntNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralIntNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralIntNode
    FP ConvNodes_LiteralIntNodeMtd
}
func ConvNodes_LiteralIntNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralIntNode).FP
}
type ConvNodes_LiteralIntNodeDownCast interface {
    ToConvNodes_LiteralIntNode() *ConvNodes_LiteralIntNode
}
func ConvNodes_LiteralIntNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralIntNodeDownCast)
    if ok { return work.ToConvNodes_LiteralIntNode() }
    return nil
}
func (obj *ConvNodes_LiteralIntNode) ToConvNodes_LiteralIntNode() *ConvNodes_LiteralIntNode {
    return obj
}
func NewConvNodes_LiteralIntNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralIntNode) *ConvNodes_LiteralIntNode {
    obj := &ConvNodes_LiteralIntNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralIntNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_LiteralIntNode) InitConvNodes_LiteralIntNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralIntNode) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
}
func (self *ConvNodes_LiteralIntNode) Get_orgNode() *Nodes_LiteralIntNode{ return self.orgNode }
// 1: decl @lune.@base.@ConvNodes.LiteralIntNode.processFilter
func (self *ConvNodes_LiteralIntNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralInt(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralIntNode.visit
func (self *ConvNodes_LiteralIntNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- LiteralRealNode
type ConvNodes_LiteralRealNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralRealNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralRealNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralRealNode
    FP ConvNodes_LiteralRealNodeMtd
}
func ConvNodes_LiteralRealNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralRealNode).FP
}
type ConvNodes_LiteralRealNodeDownCast interface {
    ToConvNodes_LiteralRealNode() *ConvNodes_LiteralRealNode
}
func ConvNodes_LiteralRealNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralRealNodeDownCast)
    if ok { return work.ToConvNodes_LiteralRealNode() }
    return nil
}
func (obj *ConvNodes_LiteralRealNode) ToConvNodes_LiteralRealNode() *ConvNodes_LiteralRealNode {
    return obj
}
func NewConvNodes_LiteralRealNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralRealNode) *ConvNodes_LiteralRealNode {
    obj := &ConvNodes_LiteralRealNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralRealNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_LiteralRealNode) InitConvNodes_LiteralRealNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralRealNode) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
}
func (self *ConvNodes_LiteralRealNode) Get_orgNode() *Nodes_LiteralRealNode{ return self.orgNode }
// 1: decl @lune.@base.@ConvNodes.LiteralRealNode.processFilter
func (self *ConvNodes_LiteralRealNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralReal(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralRealNode.visit
func (self *ConvNodes_LiteralRealNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- LiteralArrayNode
type ConvNodes_LiteralArrayNodeMtd interface {
    Get_expList() LnsAny
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralArrayNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralArrayNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralArrayNode
    expList LnsAny
    FP ConvNodes_LiteralArrayNodeMtd
}
func ConvNodes_LiteralArrayNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralArrayNode).FP
}
type ConvNodes_LiteralArrayNodeDownCast interface {
    ToConvNodes_LiteralArrayNode() *ConvNodes_LiteralArrayNode
}
func ConvNodes_LiteralArrayNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralArrayNodeDownCast)
    if ok { return work.ToConvNodes_LiteralArrayNode() }
    return nil
}
func (obj *ConvNodes_LiteralArrayNode) ToConvNodes_LiteralArrayNode() *ConvNodes_LiteralArrayNode {
    return obj
}
func NewConvNodes_LiteralArrayNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralArrayNode, arg5 LnsAny) *ConvNodes_LiteralArrayNode {
    obj := &ConvNodes_LiteralArrayNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralArrayNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_LiteralArrayNode) InitConvNodes_LiteralArrayNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralArrayNode, arg5 LnsAny) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.expList = arg5
}
func (self *ConvNodes_LiteralArrayNode) Get_orgNode() *Nodes_LiteralArrayNode{ return self.orgNode }
func (self *ConvNodes_LiteralArrayNode) Get_expList() LnsAny{ return self.expList }
// 1: decl @lune.@base.@ConvNodes.LiteralArrayNode.processFilter
func (self *ConvNodes_LiteralArrayNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralArray(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralArrayNode.visit
func (self *ConvNodes_LiteralArrayNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*ConvNodes_ExpListNode)
                if _switch7204 := visitor(&child.ConvNodes_Node, &self.ConvNodes_Node, "expList", depth); _switch7204 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch7204 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- LiteralListNode
type ConvNodes_LiteralListNodeMtd interface {
    Get_expList() LnsAny
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralListNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralListNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralListNode
    expList LnsAny
    FP ConvNodes_LiteralListNodeMtd
}
func ConvNodes_LiteralListNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralListNode).FP
}
type ConvNodes_LiteralListNodeDownCast interface {
    ToConvNodes_LiteralListNode() *ConvNodes_LiteralListNode
}
func ConvNodes_LiteralListNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralListNodeDownCast)
    if ok { return work.ToConvNodes_LiteralListNode() }
    return nil
}
func (obj *ConvNodes_LiteralListNode) ToConvNodes_LiteralListNode() *ConvNodes_LiteralListNode {
    return obj
}
func NewConvNodes_LiteralListNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralListNode, arg5 LnsAny) *ConvNodes_LiteralListNode {
    obj := &ConvNodes_LiteralListNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralListNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_LiteralListNode) InitConvNodes_LiteralListNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralListNode, arg5 LnsAny) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.expList = arg5
}
func (self *ConvNodes_LiteralListNode) Get_orgNode() *Nodes_LiteralListNode{ return self.orgNode }
func (self *ConvNodes_LiteralListNode) Get_expList() LnsAny{ return self.expList }
// 1: decl @lune.@base.@ConvNodes.LiteralListNode.processFilter
func (self *ConvNodes_LiteralListNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralList(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralListNode.visit
func (self *ConvNodes_LiteralListNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*ConvNodes_ExpListNode)
                if _switch7466 := visitor(&child.ConvNodes_Node, &self.ConvNodes_Node, "expList", depth); _switch7466 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch7466 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- LiteralSetNode
type ConvNodes_LiteralSetNodeMtd interface {
    Get_expList() LnsAny
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralSetNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralSetNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralSetNode
    expList LnsAny
    FP ConvNodes_LiteralSetNodeMtd
}
func ConvNodes_LiteralSetNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralSetNode).FP
}
type ConvNodes_LiteralSetNodeDownCast interface {
    ToConvNodes_LiteralSetNode() *ConvNodes_LiteralSetNode
}
func ConvNodes_LiteralSetNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralSetNodeDownCast)
    if ok { return work.ToConvNodes_LiteralSetNode() }
    return nil
}
func (obj *ConvNodes_LiteralSetNode) ToConvNodes_LiteralSetNode() *ConvNodes_LiteralSetNode {
    return obj
}
func NewConvNodes_LiteralSetNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralSetNode, arg5 LnsAny) *ConvNodes_LiteralSetNode {
    obj := &ConvNodes_LiteralSetNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralSetNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_LiteralSetNode) InitConvNodes_LiteralSetNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralSetNode, arg5 LnsAny) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.expList = arg5
}
func (self *ConvNodes_LiteralSetNode) Get_orgNode() *Nodes_LiteralSetNode{ return self.orgNode }
func (self *ConvNodes_LiteralSetNode) Get_expList() LnsAny{ return self.expList }
// 1: decl @lune.@base.@ConvNodes.LiteralSetNode.processFilter
func (self *ConvNodes_LiteralSetNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralSet(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralSetNode.visit
func (self *ConvNodes_LiteralSetNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*ConvNodes_ExpListNode)
                if _switch7728 := visitor(&child.ConvNodes_Node, &self.ConvNodes_Node, "expList", depth); _switch7728 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch7728 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- PairItem
type ConvNodes_PairItemMtd interface {
    Get_key() *ConvNodes_Node
    Get_val() *ConvNodes_Node
}
type ConvNodes_PairItem struct {
    key *ConvNodes_Node
    val *ConvNodes_Node
    FP ConvNodes_PairItemMtd
}
func ConvNodes_PairItem2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_PairItem).FP
}
type ConvNodes_PairItemDownCast interface {
    ToConvNodes_PairItem() *ConvNodes_PairItem
}
func ConvNodes_PairItemDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_PairItemDownCast)
    if ok { return work.ToConvNodes_PairItem() }
    return nil
}
func (obj *ConvNodes_PairItem) ToConvNodes_PairItem() *ConvNodes_PairItem {
    return obj
}
func NewConvNodes_PairItem(arg1 *ConvNodes_Node, arg2 *ConvNodes_Node) *ConvNodes_PairItem {
    obj := &ConvNodes_PairItem{}
    obj.FP = obj
    obj.InitConvNodes_PairItem(arg1, arg2)
    return obj
}
func (self *ConvNodes_PairItem) InitConvNodes_PairItem(arg1 *ConvNodes_Node, arg2 *ConvNodes_Node) {
    self.key = arg1
    self.val = arg2
}
func (self *ConvNodes_PairItem) Get_key() *ConvNodes_Node{ return self.key }
func (self *ConvNodes_PairItem) Get_val() *ConvNodes_Node{ return self.val }

// declaration Class -- LiteralMapNode
type ConvNodes_LiteralMapNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralMapNode
    Get_pairList() *LnsList
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralMapNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralMapNode
    pairList *LnsList
    FP ConvNodes_LiteralMapNodeMtd
}
func ConvNodes_LiteralMapNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralMapNode).FP
}
type ConvNodes_LiteralMapNodeDownCast interface {
    ToConvNodes_LiteralMapNode() *ConvNodes_LiteralMapNode
}
func ConvNodes_LiteralMapNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralMapNodeDownCast)
    if ok { return work.ToConvNodes_LiteralMapNode() }
    return nil
}
func (obj *ConvNodes_LiteralMapNode) ToConvNodes_LiteralMapNode() *ConvNodes_LiteralMapNode {
    return obj
}
func NewConvNodes_LiteralMapNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralMapNode, arg5 *LnsList) *ConvNodes_LiteralMapNode {
    obj := &ConvNodes_LiteralMapNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralMapNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_LiteralMapNode) InitConvNodes_LiteralMapNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralMapNode, arg5 *LnsList) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.pairList = arg5
}
func (self *ConvNodes_LiteralMapNode) Get_orgNode() *Nodes_LiteralMapNode{ return self.orgNode }
func (self *ConvNodes_LiteralMapNode) Get_pairList() *LnsList{ return self.pairList }
// 1: decl @lune.@base.@ConvNodes.LiteralMapNode.processFilter
func (self *ConvNodes_LiteralMapNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralMap(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralMapNode.visit
func (self *ConvNodes_LiteralMapNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- LiteralStringNode
type ConvNodes_LiteralStringNodeMtd interface {
    Get_dddParam() LnsAny
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralStringNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralStringNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralStringNode
    dddParam LnsAny
    FP ConvNodes_LiteralStringNodeMtd
}
func ConvNodes_LiteralStringNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralStringNode).FP
}
type ConvNodes_LiteralStringNodeDownCast interface {
    ToConvNodes_LiteralStringNode() *ConvNodes_LiteralStringNode
}
func ConvNodes_LiteralStringNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralStringNodeDownCast)
    if ok { return work.ToConvNodes_LiteralStringNode() }
    return nil
}
func (obj *ConvNodes_LiteralStringNode) ToConvNodes_LiteralStringNode() *ConvNodes_LiteralStringNode {
    return obj
}
func NewConvNodes_LiteralStringNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralStringNode, arg5 LnsAny) *ConvNodes_LiteralStringNode {
    obj := &ConvNodes_LiteralStringNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralStringNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_LiteralStringNode) InitConvNodes_LiteralStringNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralStringNode, arg5 LnsAny) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
    self.dddParam = arg5
}
func (self *ConvNodes_LiteralStringNode) Get_orgNode() *Nodes_LiteralStringNode{ return self.orgNode }
func (self *ConvNodes_LiteralStringNode) Get_dddParam() LnsAny{ return self.dddParam }
// 1: decl @lune.@base.@ConvNodes.LiteralStringNode.processFilter
func (self *ConvNodes_LiteralStringNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralString(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralStringNode.visit
func (self *ConvNodes_LiteralStringNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.dddParam
            if _child != nil {
                child := _child.(*ConvNodes_ExpListNode)
                if _switch8219 := visitor(&child.ConvNodes_Node, &self.ConvNodes_Node, "dddParam", depth); _switch8219 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch8219 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- LiteralBoolNode
type ConvNodes_LiteralBoolNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralBoolNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralBoolNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralBoolNode
    FP ConvNodes_LiteralBoolNodeMtd
}
func ConvNodes_LiteralBoolNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralBoolNode).FP
}
type ConvNodes_LiteralBoolNodeDownCast interface {
    ToConvNodes_LiteralBoolNode() *ConvNodes_LiteralBoolNode
}
func ConvNodes_LiteralBoolNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralBoolNodeDownCast)
    if ok { return work.ToConvNodes_LiteralBoolNode() }
    return nil
}
func (obj *ConvNodes_LiteralBoolNode) ToConvNodes_LiteralBoolNode() *ConvNodes_LiteralBoolNode {
    return obj
}
func NewConvNodes_LiteralBoolNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralBoolNode) *ConvNodes_LiteralBoolNode {
    obj := &ConvNodes_LiteralBoolNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralBoolNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_LiteralBoolNode) InitConvNodes_LiteralBoolNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralBoolNode) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
}
func (self *ConvNodes_LiteralBoolNode) Get_orgNode() *Nodes_LiteralBoolNode{ return self.orgNode }
// 1: decl @lune.@base.@ConvNodes.LiteralBoolNode.processFilter
func (self *ConvNodes_LiteralBoolNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralBool(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralBoolNode.visit
func (self *ConvNodes_LiteralBoolNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- LiteralSymbolNode
type ConvNodes_LiteralSymbolNodeMtd interface {
    Get_id() LnsInt
    Get_kind() string
    Get_orgNode() *Nodes_LiteralSymbolNode
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_LiteralSymbolNode struct {
    ConvNodes_ExpNode
    orgNode *Nodes_LiteralSymbolNode
    FP ConvNodes_LiteralSymbolNodeMtd
}
func ConvNodes_LiteralSymbolNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_LiteralSymbolNode).FP
}
type ConvNodes_LiteralSymbolNodeDownCast interface {
    ToConvNodes_LiteralSymbolNode() *ConvNodes_LiteralSymbolNode
}
func ConvNodes_LiteralSymbolNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_LiteralSymbolNodeDownCast)
    if ok { return work.ToConvNodes_LiteralSymbolNode() }
    return nil
}
func (obj *ConvNodes_LiteralSymbolNode) ToConvNodes_LiteralSymbolNode() *ConvNodes_LiteralSymbolNode {
    return obj
}
func NewConvNodes_LiteralSymbolNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralSymbolNode) *ConvNodes_LiteralSymbolNode {
    obj := &ConvNodes_LiteralSymbolNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_LiteralSymbolNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_LiteralSymbolNode) InitConvNodes_LiteralSymbolNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *Nodes_LiteralSymbolNode) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.orgNode = arg4
}
func (self *ConvNodes_LiteralSymbolNode) Get_orgNode() *Nodes_LiteralSymbolNode{ return self.orgNode }
// 1: decl @lune.@base.@ConvNodes.LiteralSymbolNode.processFilter
func (self *ConvNodes_LiteralSymbolNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralSymbol(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.LiteralSymbolNode.visit
func (self *ConvNodes_LiteralSymbolNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- NilAccPushNode
type ConvNodes_NilAccPushNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_NilAccPushNode struct {
    ConvNodes_ExpNode
    exp *ConvNodes_Node
    FP ConvNodes_NilAccPushNodeMtd
}
func ConvNodes_NilAccPushNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_NilAccPushNode).FP
}
type ConvNodes_NilAccPushNodeDownCast interface {
    ToConvNodes_NilAccPushNode() *ConvNodes_NilAccPushNode
}
func ConvNodes_NilAccPushNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_NilAccPushNodeDownCast)
    if ok { return work.ToConvNodes_NilAccPushNode() }
    return nil
}
func (obj *ConvNodes_NilAccPushNode) ToConvNodes_NilAccPushNode() *ConvNodes_NilAccPushNode {
    return obj
}
func NewConvNodes_NilAccPushNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) *ConvNodes_NilAccPushNode {
    obj := &ConvNodes_NilAccPushNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_NilAccPushNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_NilAccPushNode) InitConvNodes_NilAccPushNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.exp = arg4
}
func (self *ConvNodes_NilAccPushNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.NilAccPushNode.processFilter
func (self *ConvNodes_NilAccPushNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessNilAccPush(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.NilAccPushNode.visit
func (self *ConvNodes_NilAccPushNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch8812 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch8812 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch8812 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- NilAccFinNode
type ConvNodes_NilAccFinNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_NilAccFinNode struct {
    ConvNodes_ExpNode
    exp *ConvNodes_Node
    FP ConvNodes_NilAccFinNodeMtd
}
func ConvNodes_NilAccFinNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_NilAccFinNode).FP
}
type ConvNodes_NilAccFinNodeDownCast interface {
    ToConvNodes_NilAccFinNode() *ConvNodes_NilAccFinNode
}
func ConvNodes_NilAccFinNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_NilAccFinNodeDownCast)
    if ok { return work.ToConvNodes_NilAccFinNode() }
    return nil
}
func (obj *ConvNodes_NilAccFinNode) ToConvNodes_NilAccFinNode() *ConvNodes_NilAccFinNode {
    return obj
}
func NewConvNodes_NilAccFinNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) *ConvNodes_NilAccFinNode {
    obj := &ConvNodes_NilAccFinNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_NilAccFinNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_NilAccFinNode) InitConvNodes_NilAccFinNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.exp = arg4
}
func (self *ConvNodes_NilAccFinNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.NilAccFinNode.processFilter
func (self *ConvNodes_NilAccFinNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessNilAccFin(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.NilAccFinNode.visit
func (self *ConvNodes_NilAccFinNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch8949 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch8949 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch8949 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- NilAccFinCallNode
type ConvNodes_NilAccFinCallNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_NilAccFinCallNode struct {
    ConvNodes_ExpNode
    exp *ConvNodes_Node
    FP ConvNodes_NilAccFinCallNodeMtd
}
func ConvNodes_NilAccFinCallNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_NilAccFinCallNode).FP
}
type ConvNodes_NilAccFinCallNodeDownCast interface {
    ToConvNodes_NilAccFinCallNode() *ConvNodes_NilAccFinCallNode
}
func ConvNodes_NilAccFinCallNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_NilAccFinCallNodeDownCast)
    if ok { return work.ToConvNodes_NilAccFinCallNode() }
    return nil
}
func (obj *ConvNodes_NilAccFinCallNode) ToConvNodes_NilAccFinCallNode() *ConvNodes_NilAccFinCallNode {
    return obj
}
func NewConvNodes_NilAccFinCallNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) *ConvNodes_NilAccFinCallNode {
    obj := &ConvNodes_NilAccFinCallNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_NilAccFinCallNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_NilAccFinCallNode) InitConvNodes_NilAccFinCallNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.exp = arg4
}
func (self *ConvNodes_NilAccFinCallNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.NilAccFinCallNode.processFilter
func (self *ConvNodes_NilAccFinCallNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessNilAccFinCall(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.NilAccFinCallNode.visit
func (self *ConvNodes_NilAccFinCallNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch9086 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch9086 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9086 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- NilAccLastNode
type ConvNodes_NilAccLastNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_NilAccLastNode struct {
    ConvNodes_ExpNode
    exp *ConvNodes_Node
    FP ConvNodes_NilAccLastNodeMtd
}
func ConvNodes_NilAccLastNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_NilAccLastNode).FP
}
type ConvNodes_NilAccLastNodeDownCast interface {
    ToConvNodes_NilAccLastNode() *ConvNodes_NilAccLastNode
}
func ConvNodes_NilAccLastNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_NilAccLastNodeDownCast)
    if ok { return work.ToConvNodes_NilAccLastNode() }
    return nil
}
func (obj *ConvNodes_NilAccLastNode) ToConvNodes_NilAccLastNode() *ConvNodes_NilAccLastNode {
    return obj
}
func NewConvNodes_NilAccLastNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) *ConvNodes_NilAccLastNode {
    obj := &ConvNodes_NilAccLastNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_NilAccLastNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_NilAccLastNode) InitConvNodes_NilAccLastNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.exp = arg4
}
func (self *ConvNodes_NilAccLastNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.NilAccLastNode.processFilter
func (self *ConvNodes_NilAccLastNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessNilAccLast(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.NilAccLastNode.visit
func (self *ConvNodes_NilAccLastNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch9223 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch9223 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9223 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- CallExtNode
type ConvNodes_CallExtNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_CallExtNode struct {
    ConvNodes_ExpNode
    exp *ConvNodes_Node
    FP ConvNodes_CallExtNodeMtd
}
func ConvNodes_CallExtNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_CallExtNode).FP
}
type ConvNodes_CallExtNodeDownCast interface {
    ToConvNodes_CallExtNode() *ConvNodes_CallExtNode
}
func ConvNodes_CallExtNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_CallExtNodeDownCast)
    if ok { return work.ToConvNodes_CallExtNode() }
    return nil
}
func (obj *ConvNodes_CallExtNode) ToConvNodes_CallExtNode() *ConvNodes_CallExtNode {
    return obj
}
func NewConvNodes_CallExtNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) *ConvNodes_CallExtNode {
    obj := &ConvNodes_CallExtNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_CallExtNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_CallExtNode) InitConvNodes_CallExtNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.exp = arg4
}
func (self *ConvNodes_CallExtNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.CallExtNode.processFilter
func (self *ConvNodes_CallExtNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessCallExt(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.CallExtNode.visit
func (self *ConvNodes_CallExtNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch9360 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch9360 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9360 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- CondPopValNode
type ConvNodes_CondPopValNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_CondPopValNode struct {
    ConvNodes_ExpNode
    exp *ConvNodes_Node
    FP ConvNodes_CondPopValNodeMtd
}
func ConvNodes_CondPopValNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_CondPopValNode).FP
}
type ConvNodes_CondPopValNodeDownCast interface {
    ToConvNodes_CondPopValNode() *ConvNodes_CondPopValNode
}
func ConvNodes_CondPopValNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_CondPopValNodeDownCast)
    if ok { return work.ToConvNodes_CondPopValNode() }
    return nil
}
func (obj *ConvNodes_CondPopValNode) ToConvNodes_CondPopValNode() *ConvNodes_CondPopValNode {
    return obj
}
func NewConvNodes_CondPopValNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) *ConvNodes_CondPopValNode {
    obj := &ConvNodes_CondPopValNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_CondPopValNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_CondPopValNode) InitConvNodes_CondPopValNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.exp = arg4
}
func (self *ConvNodes_CondPopValNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.CondPopValNode.processFilter
func (self *ConvNodes_CondPopValNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessCondPopVal(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.CondPopValNode.visit
func (self *ConvNodes_CondPopValNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch9497 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch9497 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9497 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- CondSetValNode
type ConvNodes_CondSetValNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_CondSetValNode struct {
    ConvNodes_ExpNode
    exp *ConvNodes_Node
    FP ConvNodes_CondSetValNodeMtd
}
func ConvNodes_CondSetValNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_CondSetValNode).FP
}
type ConvNodes_CondSetValNodeDownCast interface {
    ToConvNodes_CondSetValNode() *ConvNodes_CondSetValNode
}
func ConvNodes_CondSetValNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_CondSetValNodeDownCast)
    if ok { return work.ToConvNodes_CondSetValNode() }
    return nil
}
func (obj *ConvNodes_CondSetValNode) ToConvNodes_CondSetValNode() *ConvNodes_CondSetValNode {
    return obj
}
func NewConvNodes_CondSetValNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) *ConvNodes_CondSetValNode {
    obj := &ConvNodes_CondSetValNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_CondSetValNode(arg1, arg2, arg3, arg4)
    return obj
}
func (self *ConvNodes_CondSetValNode) InitConvNodes_CondSetValNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.exp = arg4
}
func (self *ConvNodes_CondSetValNode) Get_exp() *ConvNodes_Node{ return self.exp }
// 1: decl @lune.@base.@ConvNodes.CondSetValNode.processFilter
func (self *ConvNodes_CondSetValNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessCondSetVal(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.CondSetValNode.visit
func (self *ConvNodes_CondSetValNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch9634 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch9634 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9634 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- GetAtNode
type ConvNodes_GetAtNodeMtd interface {
    Get_exp() *ConvNodes_Node
    Get_id() LnsInt
    Get_index() *ConvNodes_Node
    Get_kind() string
    Get_pos() *Parser_Position
    ProcessFilter(arg1 *ConvNodes_Filter, arg2 LnsAny)
    Visit(arg1 ConvNodes_NodeVisitor, arg2 LnsInt) bool
}
type ConvNodes_GetAtNode struct {
    ConvNodes_ExpNode
    exp *ConvNodes_Node
    index *ConvNodes_Node
    FP ConvNodes_GetAtNodeMtd
}
func ConvNodes_GetAtNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvNodes_GetAtNode).FP
}
type ConvNodes_GetAtNodeDownCast interface {
    ToConvNodes_GetAtNode() *ConvNodes_GetAtNode
}
func ConvNodes_GetAtNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvNodes_GetAtNodeDownCast)
    if ok { return work.ToConvNodes_GetAtNode() }
    return nil
}
func (obj *ConvNodes_GetAtNode) ToConvNodes_GetAtNode() *ConvNodes_GetAtNode {
    return obj
}
func NewConvNodes_GetAtNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node, arg5 *ConvNodes_Node) *ConvNodes_GetAtNode {
    obj := &ConvNodes_GetAtNode{}
    obj.FP = obj
    obj.ConvNodes_ExpNode.FP = obj
    obj.ConvNodes_Node.FP = obj
    obj.InitConvNodes_GetAtNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvNodes_GetAtNode) InitConvNodes_GetAtNode(arg1 LnsInt, arg2 *Parser_Position, arg3 string, arg4 *ConvNodes_Node, arg5 *ConvNodes_Node) {
    self.ConvNodes_ExpNode.InitConvNodes_ExpNode( arg1,arg2,arg3)
    self.exp = arg4
    self.index = arg5
}
func (self *ConvNodes_GetAtNode) Get_exp() *ConvNodes_Node{ return self.exp }
func (self *ConvNodes_GetAtNode) Get_index() *ConvNodes_Node{ return self.index }
// 1: decl @lune.@base.@ConvNodes.GetAtNode.processFilter
func (self *ConvNodes_GetAtNode) ProcessFilter(filter *ConvNodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessGetAt(self, opt)
}

// 225: decl @lune.@base.@ConvNodes.GetAtNode.visit
func (self *ConvNodes_GetAtNode) Visit(visitor ConvNodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *ConvNodes_Node
        child = self.exp
        if _switch9782 := visitor(child, &self.ConvNodes_Node, "exp", depth); _switch9782 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9782 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *ConvNodes_Node
        child = self.index
        if _switch9836 := visitor(child, &self.ConvNodes_Node, "index", depth); _switch9836 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9836 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


func Lns_ConvNodes_init() {
    if init_ConvNodes { return }
    init_ConvNodes = true
    ConvNodes__mod__ = "@lune.@base.@ConvNodes"
    Lns_InitMod()
    Lns_Parser_init()
    Lns_Nodes_init()
    Lns_Util_init()
    ConvNodes_nodeKind2createFromFunc = NewLnsMap( map[LnsAny]LnsAny{})
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpList(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1185_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpNew(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1233_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpUnwrap(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1290_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpRef(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1328_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpOp2(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1384_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpCast(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1431_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpToDDD(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1478_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpSubDDD(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1525_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpOp1(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1572_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpRefItem(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1628_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpCall(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1684_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpMRet(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1731_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpAccessMRet(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1778_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpMultiTo1(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1825_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpParen(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1872_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_ExpOmitEnum(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1910_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_RefField(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_1957_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_GetField(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2004_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_DeclFunc(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2042_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_NewAlgeVal(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2103_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LuneKind(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2150_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralChar(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2188_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralInt(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2226_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralReal(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2264_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralArray(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2311_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralList(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2358_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralSet(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2405_))
    
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralMap(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2476_))
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralString(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2523_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralBool(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2561_))
    
    ConvNodes_nodeKind2createFromFunc.Set(Nodes_NodeKind_get_LiteralSymbol(),ConvNodes_createFromFunc_1063_(ConvNodes__anonymous_2599_))
    
    
    
    
    
    
    
    
    
    Lns_LuaVer_init()
    Lns_Option_init()
    Lns_convLua_init()
}
func init() {
    init_ConvNodes = false
}
