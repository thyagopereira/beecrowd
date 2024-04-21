package main

import (
	"fmt"
	stringsLib "strings"
)

func CommonCharacters(strings []string) []string {

	strSet := make(map[string]int)
	for i := 0; i < len(strings); i++ {
		for _, r := range strings[i] {
			if _, ok := strSet[string(r)]; !ok {
				strSet[string(r)] = 0
			}
		}
	}

	// Agora strSet é um HashSet que contém como chave, apenas as letras sem repetição
	for k, _ := range strSet {
		for i := 0; i < len(strings); i++ {
			if stringsLib.Contains(strings[i], k) {
				strSet[k] += 1
			}
		}
	}

	commonChar := make([]string, 0)
	for k, _ := range strSet {
		if strSet[k] == len(strings) {
			commonChar = append(commonChar, k)
		}
	}

	return commonChar

}

func main() {
	input := []string{"abc", "bcd", "cbad"}
	expected := []string{"b", "c"}

	output := CommonCharacters(input)
	fmt.Println(output)

	for i, _ := range expected {
		if output[i] != expected[i] {
			panic("Wrong Answer")
		}
	}
}
