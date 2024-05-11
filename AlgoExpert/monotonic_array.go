package main

import "fmt"

func IsMonotonic(array []int) bool {

	if len(array) < 2 {
		return true
	}

	p, n := 0, 1
	decreasing, increasing := false, false

	for n < len(array) {
		if array[n] < array[p] {
			decreasing = true
		} else if array[n] > array[p] {
			increasing = true
		}
		p++
		n++
	}

	return !(decreasing && increasing)

}

func main() {
	input := []int{1, 5, 10}
	output := IsMonotonic(input)

	fmt.Println(output)
	expected := true
	if expected != output {
		panic("Wrong answer")
	}

}
