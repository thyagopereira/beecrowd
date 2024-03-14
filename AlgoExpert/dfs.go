package main

import "fmt"

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) isLeaf() bool {
	if len(n.Children) == 0 {
		return true
	}
	return false
}

func (n *Node) DFS(array *[]string) {
	*array = append(*array, n.Name)

	for i := 0; i < len(n.Children); i++ {
		if !n.isLeaf() {
			n.Children[i].DFS(array)
		}
	}
	return
}

func (n *Node) DepthFirstSearch(array []string) []string {
	n.DFS(&array)
	return array
}

func main() {
	root := &Node{
		Name: "A",
		Children: []*Node{
			{
				Name: "B",
				Children: []*Node{
					{
						Name:     "E",
						Children: []*Node{},
					},
					{
						Name: "F",
						Children: []*Node{
							{
								Name:     "I",
								Children: []*Node{},
							},
							{
								Name:     "J",
								Children: []*Node{},
							},
						},
					},
				},
			},
			{
				Name:     "C",
				Children: []*Node{},
			},
			{
				Name: "D",
				Children: []*Node{
					{
						Name: "G",
						Children: []*Node{
							{
								Name:     "K",
								Children: []*Node{},
							},
						},
					},
					{
						Name:     "H",
						Children: []*Node{},
					},
				},
			},
		},
	}

	input := make([]string, 0)
	out := root.DepthFirstSearch(input)
	fmt.Println(out)
}
