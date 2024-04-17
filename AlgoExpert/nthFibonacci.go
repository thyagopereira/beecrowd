package main

func GetNthFib(n int) int {
	a := 0
	b := 1

	if n == 1 {
		return a
	}

	if n == 2 {
		return b
	}

	var c int
	for i := 0; i < n-2; i++ {
		c = a + b
		a = b
		b = c
	}

	return c

}

func main() {

	n := 6
	output := GetNthFib(n)

	if output != 5 {
		panic("WRONG ANSWER")
	}
}
