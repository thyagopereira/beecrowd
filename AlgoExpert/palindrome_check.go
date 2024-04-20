package main

import "fmt"

func IsPalindrome(str string) bool {
	// Test middle out

	runes := []rune(str)
	mid := len(runes) / 2

	var i, j int
	if len(runes)%2 != 0 {
		i, j = mid, mid
	} else {
		i, j = mid-1, mid
	}

	for i >= 0 && j < len(str) {
		if str[i] != str[j] {
			return false
		}

		i--
		j++
	}

	return true
}

func main() {
	input := "abbas"
	output := IsPalindrome(input)

	fmt.Println(output)

	// Test
	expected := false
	if expected != output {
		panic("Wrong answer")
	}

}
