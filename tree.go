package tree

type Node interface {
	Children() []Node
}

type Node struct {
	ChildFirst *Node
	ChildLast  *Node
	SibLeft    *Node
	SibRight   *Node
}

func (n *Node) Childer() []*Node {
	// Todo
	return nil
}
