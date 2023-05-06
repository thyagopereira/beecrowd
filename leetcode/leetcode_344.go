package leetcode

func reverseString(s []byte) {
	//  Sempre trocar i com j - O(n)

	i, j := 0, len(s)-1

	for i < j {
		s[i], s[j] = s[j], s[i]
		j--
		i++
	}
}
