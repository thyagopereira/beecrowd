package main

import (
	"fmt"
	"math"
)

func FindThreeLargestNumbers(array []int) []int {
	small, mid, large := int(math.Inf(-1)), int(math.Inf(-1)), int(math.Inf(-1))

	for i := 0; i < len(array); i++ {
		if array[i] > small {
			if array[i] > mid {
				if array[i] > large {
					tmpMid := large
					tmpSmall := mid
					large = array[i]
					mid = tmpMid
					small = tmpSmall
				} else {
					tmpSmall := mid
					mid = array[i]
					small = tmpSmall
				}
			} else {
				small = array[i]
			}
		}

	}

	return []int{small, mid, large}
}

func main() {
	input := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	expected := []int{18, 141, 541}

	output := FindThreeLargestNumbers(input)
	fmt.Println(output)

	for i := 0; i < len(output); i++ {
		if output[i] != expected[i] {
			panic("WRONG ANSWER")
		}
	}

}
