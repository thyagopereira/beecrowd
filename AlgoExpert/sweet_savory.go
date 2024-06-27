package main

import (
	"fmt"
	"math"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func preProcess(dishes []int) ([]int, []int) {
	sweet := make([]int, 0)
	savory := make([]int, 0)
	for i := 0; i < len(dishes); i++ {
		if dishes[i] < 0 {
			sweet = append(sweet, dishes[i])
		} else {
			savory = append(savory, dishes[i])
		}
	}

	sort.Reverse(sort.IntSlice(sweet))
	sort.Ints(savory)

	return sweet, savory
}

func SweetAndSavory(dishes []int, target int) []int {
	// Must be sorted
	sweet, savory := preProcess(dishes)

	bestPair := []int{0, 0}
	bestDiff := math.MaxInt64

	i, j := 0, 0
	for i < len(sweet) && j < len(savory) {
		current := sweet[i] + savory[j]

		if current <= target {
			currentDiff := target - current
			if currentDiff < bestDiff {
				bestPair[0] = sweet[i]
				bestPair[1] = savory[j]
				bestDiff = currentDiff
			}
			j += 1
		} else {
			i += 1
		}

	}

	return bestPair

}

func main() {

	// -25,-7, -4, 2, 5, 12, 100
	// diff

	inputDishes := []int{2, 5, -4, -7, 12, 100, -25}
	inputTarget := 7

	output := SweetAndSavory(inputDishes, inputTarget)
	expectedOutput := []int{-7, 12}

	fmt.Println(output)
	for i := 0; i < len(expectedOutput); i++ {
		if output[i] != expectedOutput[i] {
			panic("Wrong Answer")
		}
	}

}
