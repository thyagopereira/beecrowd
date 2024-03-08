package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func closestValueSearch(node *BST, target, closest int) int {

	closestDiff := absInt(closest - target)
	closestActual := absInt(node.Value - target)

	var close int
	if closestDiff < closestActual {
		close = closest
	} else {
		close = node.Value
	}

	if target >= node.Value {
		if node.Right != nil {
			return closestValueSearch(node.Right, target, close)
		}
	}

	if target <= node.Value {
		if node.Left != nil {
			return closestValueSearch(node.Left, target, close)
		}
	}

	return close
}

func (tree *BST) FindClosestValue(target int) int {

	return closestValueSearch(tree, target, tree.Value)
}

func main() {

	root := BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left: &BST{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BST{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BST{
			Value: 15,
			Left: &BST{
				Value: 13,
				Left:  nil,
				Right: &BST{
					Value: 14,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BST{
				Value: 22,
				Right: nil,
				Left:  nil,
			},
		},
	}

	result := root.FindClosestValue(12)
	fmt.Println(result)
}
