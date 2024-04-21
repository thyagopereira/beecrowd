package main

import "fmt"

func GenerateDocument(characters string, document string) bool {
	freqChar := make(map[string]int)
	for _, r := range characters {
		freqChar[string(r)] += 1
	}

	freqDocument := make(map[string]int)
	for _, r := range document {
		freqDocument[string(r)] += 1
	}

	for _, k := range document {
		if _, ok := freqChar[string(k)]; !ok {
			return false
		} else if freqChar[string(k)] < freqDocument[string(k)] {
			return false
		}
	}

	return true
}

func main() {
	characters := "Bste!hetsi ogEAxpelrt x "
	document := "AlgoExpert is the Best!"

	expected := true
	output := GenerateDocument(characters, document)

	fmt.Println(output)

	if output != expected {
		panic("Wrong answer")
	}
}
