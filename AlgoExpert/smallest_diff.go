package main

import (
	"fmt"
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	// To transpass the fact that  int(math.Inf(1)) is < 0
	const INF = 1000000000000
	sort.Ints(array1)
	sort.Ints(array2)

	idxOne, idxTwo := 0, 0
	smallestPair := []int{array1[idxOne], array2[idxTwo]}
	currentDiff := int(math.Inf(1))
	smallestDiff := INF

	checknewDiff := func() {
		if currentDiff < smallestDiff {
			smallestDiff = currentDiff
			smallestPair[0] = array1[idxOne]
			smallestPair[1] = array2[idxTwo]
		}
	}

	for idxOne < len(array1) && idxTwo < len(array2) {
		if array1[idxOne] < array2[idxTwo] {
			currentDiff = array2[idxTwo] - array1[idxOne]
			checknewDiff()
			idxOne += 1
		} else if array2[idxTwo] < array1[idxOne] {
			currentDiff = array1[idxOne] - array2[idxTwo]
			checknewDiff()
			idxTwo += 1
		} else {
			return []int{array1[idxOne], array2[idxTwo]}
		}

	}

	return smallestPair
}

func main() {
	arrayOne := []int{240, 124, 86, 111, 2, 84, 954, 27, 89}
	arrayTwo := []int{1, 3, 954, 19, 8}

	expected := []int{954, 954}
	output := SmallestDifference(arrayOne, arrayTwo)
	fmt.Println(output)
	for i := 0; i < len(output); i++ {
		if expected[i] != output[i] {
			panic("Wrong answer")
		}
	}
}
