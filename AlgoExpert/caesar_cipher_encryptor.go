package main

import "fmt"

func getLetters() map[int]string {
	return map[int]string{
		97:  "a",
		98:  "b",
		99:  "c",
		100: "d",
		101: "e",
		102: "f",
		103: "g",
		104: "h",
		105: "i",
		106: "j",
		107: "k",
		108: "l",
		109: "m",
		110: "n",
		111: "o",
		112: "p",
		113: "q",
		114: "r",
		115: "s",
		116: "t",
		117: "u",
		118: "v",
		119: "w",
		120: "x",
		121: "y",
		122: "z",
	}
}

const startPoint int = 96

func CaesarCipherEncryptor(str string, key int) string {
	letters := getLetters()

	var answer string = ""

	for _, r := range str {

		runeCode := (int(r) + key)

		for runeCode > 122 {
			runeCode = (runeCode % 122) + startPoint
		}

		answer += letters[runeCode]
	}

	return answer

}

func main() {
	strInput := "xyz"
	key := 2

	expected := "zab"
	output := CaesarCipherEncryptor(strInput, key)
	fmt.Println(output)

	if expected != output {
		panic("Wrong Answer")
	}
}
