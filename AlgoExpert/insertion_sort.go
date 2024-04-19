package main

import "fmt"

func InsertionSort(array []int) []int {

	for i := 0; i < len(array); i++ {
		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}

	return array
}

func main() {
	input := []int{8, 5, 2, 9, 5, 6, 3}
	expected := []int{2, 3, 5, 5, 6, 8, 9}

	output := InsertionSort(input)

	fmt.Println(output)

	for i := 0; i < len(expected); i++ {
		if expected[i] != output[i] {
			panic("Wrong Answer")
		}
	}

}
