package main

import "fmt"

func SelectionSort(array []int) []int {

	for i := 0; i < len(array); i++ {
		min := i
		for j := i; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		array[i], array[min] = array[min], array[i]
	}

	return array
}

func main() {
	input := []int{8, 5, 2, 9, 5, 6, 3}
	expected := []int{2, 3, 5, 5, 6, 8, 9}

	output := SelectionSort(input)
	fmt.Println(output)

	for i, _ := range expected {
		if expected[i] != output[i] {
			panic("Wrong Answer")
		}
	}
}
