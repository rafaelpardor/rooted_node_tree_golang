package tree

type Node interface {
	Children() []Node
}

type node struct {
	parent *node
	first  *node
	last   *node
	left   *node
	right  *node
}

// New is a new shorted way to write NewNode
func New() *node { return NewNode() }

// NewNode returns a struct that fulfills the Node Interface
func NewNode() *node {
	return new(node)
}

// Children return a slice of all inmediate children, not grantchildren.
func (n *node) Children() []*node {
	// TODO
	return nil
}
