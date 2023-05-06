package leetcode

func reverse(w []byte) []byte {
	i, j := 0, len(w)-1

	for i <= j {
		w[i], w[j] = w[j], w[i]
		i++
		j--
	}

	return w
}

func reverseWords(s string) string {
	var aux []byte
	var ans []byte

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {

			ans = append(ans, reverse(aux)...)
			ans = append(ans, []byte(" ")...)
			aux = nil
		} else {
			aux = append(aux, s[i])
		}

	}

	ans = append(ans, reverse(aux)...)

	return string(ans)
}
