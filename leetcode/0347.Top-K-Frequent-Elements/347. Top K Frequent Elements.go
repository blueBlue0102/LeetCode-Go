package leetcode

func TopKFrequentElements(nums []int, k int) []int {
	bucket := make([][]int, len(nums))
	freqTable := map[int]int{}
	for _, num := range nums {
		freqTable[num]++
	}
	for num, freq := range freqTable {
		freq--
		if bucket[freq] == nil {
			bucket[freq] = make([]int, 0)
		}
		bucket[freq] = append(bucket[freq], num)
	}
	result := []int{}
	count := 0
	for i := len(bucket) - 1; i >= 0; i-- {
		if bucket[i] == nil {
			continue
		}
		result = append(result, bucket[i]...)
		count += len(bucket[i])
		if count == k {
			break
		}
	}
	return result
}
