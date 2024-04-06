package leetcode

func HappyNumber(n int) bool {
	dict := make(map[int]bool)

	for {
		if n == 1 {
			return true
		}
		sum := 0
		for num := n; num > 0; num /= 10 {
			digit := num % 10
			sum += digit * digit
		}
		if dict[sum] {
			return false
		}
		dict[sum] = true
		n = sum
	}
}
