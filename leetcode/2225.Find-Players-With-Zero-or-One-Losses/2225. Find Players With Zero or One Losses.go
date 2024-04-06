package leetcode

func FindPlayersWithZeroorOneLosses(matches [][]int) [][]int {
	ans1 := []int{}
	ans2 := []int{}
	playerLooseCount := make([]int, 100001)
	for index := range playerLooseCount {
		playerLooseCount[index] = -1
	}

	for _, match := range matches {
		winner := match[0]
		looser := match[1]
		if playerLooseCount[winner] == -1 {
			playerLooseCount[winner] = 0
		}
		if playerLooseCount[looser] == -1 {
			playerLooseCount[looser] = 0
		}

		playerLooseCount[looser]++
	}

	for player, scores := range playerLooseCount {
		if scores == 0 {
			ans1 = append(ans1, player)
		} else if scores == 1 {
			ans2 = append(ans2, player)
		}
	}
	return [][]int{ans1, ans2}
}
