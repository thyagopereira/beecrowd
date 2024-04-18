package main

import "fmt"

func BinarySearch(array []int, target int) int {
	b, e := 0, len(array)-1

	for b <= e {
		mid := e - b

		if array[mid] == target {
			return mid
		}

		if array[mid] < target {
			b = mid + 1
		}

		if array[mid] > target {
			e = mid - 1
		}
	}

	return -1
}

func main() {
	input := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 73

	output := BinarySearch(input, target)
	expected := len(input) - 1

	fmt.Println(output)
	if output != expected {
		panic("WRONG ANSWER")
	}
}
