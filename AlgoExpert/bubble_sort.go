package main

import "fmt"

func BubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[j] > array[i] {
				tmp := array[i]
				array[i] = array[j]
				array[j] = tmp
			}
		}
	}

	return array
}

func main() {

	input := []int{8, 5, 2, 9, 5, 6, 3}
	expected := []int{2, 3, 5, 5, 6, 8, 9}

	output := BubbleSort(input)
	fmt.Println(output)

	for i := 0; i < len(expected); i++ {
		if output[i] != expected[i] {
			panic("WRONG ANSWER")
		}
	}

}
