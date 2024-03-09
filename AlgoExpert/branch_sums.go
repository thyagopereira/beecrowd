package main

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func isLeaf(node *BinaryTree) bool {
	return node.Left == nil && node.Right == nil
}

func BranchSum(node *BinaryTree, results *[]int, sum int) {

	sum += node.Value
	if isLeaf(node) {
		*results = append(*results, sum)
	} else {
		if node.Left != nil {
			BranchSum(node.Left, results, sum)
		}

		if node.Right != nil {
			BranchSum(node.Right, results, sum)
		}
	}
}

func BranchSums(root *BinaryTree) []int {
	result := []int{}

	BranchSum(root, &result, 0)
	return result
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
				Left: &BinaryTree{
					Value: 10,
					Left:  nil,
					Right: nil,
				},
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

	result := BranchSums(&root)
	fmt.Println(result)
}
