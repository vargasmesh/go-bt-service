package tree

type Less[T any] func(t1, t2 T) bool

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type Tree[T any] struct {
	Root *Node[T]
	less Less[T]
}

func New[T any](less Less[T]) *Tree[T] {
	return &Tree[T]{
		less: less,
	}
}

func (t *Tree[T]) Insert(value T) *Tree[T] {
	if t.Root == nil {
		t.Root = &Node[T]{Value: value}
		return t
	}
	t.Root.insert(value, t.less)
	return t
}

func (n *Node[T]) insert(value T, less Less[T]) {
	if less(value, n.Value) {
		if n.Left == nil {
			n.Left = &Node[T]{Value: value}
		} else {
			n.Left.insert(value, less)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node[T]{Value: value}
		} else {
			n.Right.insert(value, less)
		}
	}
}

func InOrder[T any](n *Node[T], f func(*Node[T])) {
	if n == nil {
		return
	}
	InOrder(n.Left, f)
	f(n)
	InOrder(n.Right, f)
}

func PreOrder[T any](n *Node[T], f func(*Node[T])) {
	if n == nil {
		return
	}
	f(n)
	PreOrder(n.Left, f)
	PreOrder(n.Right, f)
}
