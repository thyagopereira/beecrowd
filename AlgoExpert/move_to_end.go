package main

func MoveElementToEnd(array []int, toMove int) []int {
	b, e := 0, len(array)-1

	for b < e {
		if array[b] == toMove {
			array[b], array[e] = array[e], array[b]
			e--
		} else {
			b++
		}
	}

	return array
}

func main() {
	input := []int{2, 1, 2, 2, 2, 3, 4, 2}
	output := MoveElementToEnd(input, 2)

	expected := []int{4, 1, 3, 2, 2, 2, 2, 2}
	for i := 0; i < len(expected); i++ {
		if output[i] != expected[i] {
			panic("Wrong answer")
		}
	}

}
