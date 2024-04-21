package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x-y < 0 {
		return x
	}
	return y
}

func FirstNonRepeatingCharacter(str string) int {

	counter := make(map[string][]int)

	for i := 0; i < len(str); i++ {
		counter[string(str[i])] = append(counter[string(str[i])], i)
	}

	var minor int = int(math.Inf(1))
	for _, v := range counter {
		if len(v) == 1 {
			minor = min(v[0], minor)

		}
	}

	if minor == int(math.Inf(1)) {
		return -1
	}

	return minor
}

func main() {
	input := "abcdcaf"
	expected := 1

	output := FirstNonRepeatingCharacter(input)
	fmt.Println(output)

	if output != expected {
		panic("Wrong answer")
	}
}
