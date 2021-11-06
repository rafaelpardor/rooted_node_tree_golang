package tree_test

import (
	"fmt"

	tree "github.com/rafaelpardor/rooted_node_tree_golang"
)

func ExampleNew() {

	root := tree.New()
	fmt.Printf("%T\n", root)

	// Output:
	// *tree.node
}

func ExampleNewNode() {

	node := tree.NewNode()
	fmt.Printf("%T\n", node)

	// Output:
	// *tree.node

}

func ExampleNode() {
	node := tree.NewNode()
	fmt.Printf("%T\n", node)
	// Output:
	// *tree.node
}
