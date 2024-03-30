package leetcode

func FindSmallestLetterGreaterThanTarget(letters []byte, target byte) byte {
	// Brute Force
	// for _, letter := range letters {
	// 	if letter > target {
	// 		return letter
	// 	}
	// }
	// return letters[0]

	// Binary Search
	left := 0 // left 最終會是最小的大於 target 的 value，除非 left 超出了 array(left == len(letters))
	right := len(letters) - 1

	for left <= right {
		mid := (left + right) / 2
		if letters[mid] > target {
			right = mid - 1
		} else if letters[mid] <= target {
			left = mid + 1
		}
	}

	if left == len(letters) {
		return letters[0]
	} else {
		return letters[left]
	}
}
