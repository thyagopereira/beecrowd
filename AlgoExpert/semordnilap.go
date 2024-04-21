package main

import (
	"fmt"
	"reflect"
)

func reverseWord(word string) string {
	reversed := ""

	for i := (len(word) - 1); i >= 0; i-- {
		reversed += string(word[i])
	}

	return reversed
}

func Semordnilap(words []string) [][]string {
	// Write your code here.

	aux := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		if _, ok := aux[words[i]]; !ok {
			aux[words[i]] = true
		}
	}

	answer := make([][]string, 0)
	for k, _ := range aux {
		runSemor := make([]string, 0)
		reversed := reverseWord(k)

		if reversed != k {
			if _, ok := aux[reversed]; ok {
				runSemor = append(runSemor, k)
				delete(aux, k)
				runSemor = append(runSemor, reversed)
				delete(aux, reversed)
			}
		}

		if len(runSemor) != 0 {
			answer = append(answer, runSemor)
		}

	}

	return answer
}

func main() {
	input := []string{"diaper", "abc", "test",
		"cba", "repaid"}

	expected := [][]string{
		{"diaper", "repaid"},
		{"abc", "cba"},
	}

	output := Semordnilap(input)
	fmt.Println(output)

	if !reflect.DeepEqual(expected, output) {
		panic("Wrong Answer")
	}

}
