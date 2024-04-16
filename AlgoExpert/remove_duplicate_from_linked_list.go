package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	var head, nextDistinct *LinkedList
	head = linkedList
	current := linkedList

	for current != nil {
		nextDistinct = current.Next
		for nextDistinct != nil && nextDistinct.Value == current.Value {
			nextDistinct = nextDistinct.Next
		}

		current.Next = nextDistinct
		current = nextDistinct
	}

	return head

}

func main() {

	// 1 - 1 - 3 - 4 - 4- 4 - 5 - 6 - 6
	input := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 1,
			Next: &LinkedList{
				Value: 3,
				Next: &LinkedList{
					Value: 4,
					Next: &LinkedList{
						Value: 4,
						Next: &LinkedList{
							Value: 4,
							Next: &LinkedList{
								Value: 5,
								Next: &LinkedList{
									Value: 6,
									Next: &LinkedList{
										Value: 6,
										Next:  nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	//1 - 3 - 4 - 5 - 6
	expected := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 3,
			Next: &LinkedList{
				Value: 4,
				Next: &LinkedList{
					Value: 5,
					Next: &LinkedList{
						Value: 6,
						Next:  nil,
					},
				},
			},
		},
	}
	output := RemoveDuplicatesFromLinkedList(input)

	for output != nil {
		if output.Value != expected.Value {
			panic("Wrong answer")
		}
		fmt.Println(output.Value)
		output = output.Next
		expected = expected.Next
	}
}
