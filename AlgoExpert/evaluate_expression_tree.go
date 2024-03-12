package main

import (
	"fmt"
)

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func EvaluateExpressionTree(tree *BinaryTree) int {
	if tree.Value > 0 {
		return tree.Value
	} else {
		switch tree.Value {
		case -1:
			return EvaluateExpressionTree(tree.Left) + EvaluateExpressionTree(tree.Right)
		case -2:
			return EvaluateExpressionTree(tree.Left) - EvaluateExpressionTree(tree.Right)
		case -3:
			return EvaluateExpressionTree(tree.Left) / EvaluateExpressionTree(tree.Right)
		case -4:
			return EvaluateExpressionTree(tree.Left) * EvaluateExpressionTree(tree.Right)
		default:
			return 0
		}
	}
}

func main() {

	root := BinaryTree{
		Value: -1,
		Left: &BinaryTree{
			Value: -2,
			Left: &BinaryTree{
				Value: -4,
				Left: &BinaryTree{
					Value: 2,
					Left:  nil,
					Right: nil,
				},
				Right: &BinaryTree{
					Value: 3,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BinaryTree{
				Value: 2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryTree{
			Value: -3,
			Left: &BinaryTree{
				Value: 8,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryTree{
				Value: 3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := EvaluateExpressionTree(&root)
	fmt.Println(result)
}
