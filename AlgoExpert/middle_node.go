package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
	slow, fast := linkedList, linkedList

	for fast.Next != nil {
		slow = slow.Next

		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}

	}

	return slow
}

func main() {

	// 2, 7,3,5
	input := &LinkedList{
		Value: 2,
		Next: &LinkedList{
			Value: 7,
			Next: &LinkedList{
				Value: 3,
				Next: &LinkedList{
					Value: 5,
					Next:  nil,
				},
			},
		},
	}

	expected := &LinkedList{
		Value: 3,
		Next: &LinkedList{
			Value: 5,
			Next:  nil,
		},
	}

	output := MiddleNode(input)

	for output != nil {
		if output.Value != expected.Value {
			panic("WRONG ANSWER")
		}
		output = output.Next
		expected = expected.Next
	}

}
