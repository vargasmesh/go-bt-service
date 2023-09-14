package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(value int) *Tree {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return t
	}
	t.Root.insert(value)
	return t
}

func (n *Node) insert(value int) {
	if value > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.insert(value)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.insert(value)
		}
	}
}

func InOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	InOrder(n.Left, f)
	f(n)
	InOrder(n.Right, f)
}

func PreOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	f(n)
	PreOrder(n.Left, f)
	PreOrder(n.Right, f)
}
