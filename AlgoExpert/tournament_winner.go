package main

import "fmt"

const HOME_TEAM_WON = 1

func main() {
	input := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}

	results := []int{0, 0, 1}
	winner := TournamentWinner(input, results)
	fmt.Println(winner)
}

func TournamentWinner(competitions [][]string, results []int) string {
	res := make(map[string]int)

	for i := 0; i < len(competitions); i++ {
		if results[i] == 0 {
			res[competitions[i][1]] += 1
		} else {
			res[competitions[i][0]] += 1
		}
	}

	win, count := "", 0
	for k, v := range res {
		if v >= count {
			win = k
			count = v
		}
	}

	return win
}
