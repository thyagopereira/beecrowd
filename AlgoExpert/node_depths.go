package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func isLeaf(node *BinaryTree) bool {
	return node.Left == nil && node.Right == nil
}

func SumDepths(node *BinaryTree, h int, sum *int) {

	*sum += h
	if !isLeaf(node) {
		if node.Left != nil {
			SumDepths(node.Left, h+1, sum)
		}

		if node.Right != nil {
			SumDepths(node.Right, h+1, sum)
		}
	}
}

func NodeDepths(root *BinaryTree) int {

	sum := 0
	SumDepths(root, 0, &sum)
	return sum
}

func main() {
	root := BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
					Left:  nil,
					Right: nil,
				},
				Right: &BinaryTree{
					Value: 9,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BinaryTree{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryTree{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
		},
	}

	result := NodeDepths(&root)
	fmt.Println(result)
}
