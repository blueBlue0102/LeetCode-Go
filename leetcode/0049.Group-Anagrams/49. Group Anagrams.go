package leetcode

import "bytes"

func GroupAnagrams(strs []string) [][]string {
	result := [][]string{}
	hashTable := map[string][]string{}
	for _, str := range strs {
		key := getHashTableKey(str)
		if hashTable[key] == nil {
			hashTable[key] = []string{}
		}
		hashTable[key] = append(hashTable[key], str)
	}

	for _, val := range hashTable {
		result = append(result, val)
	}
	return result
}

func getHashTableKey(str string) string {
	runeArr := make([]rune, 26)
	for _, runeVal := range str {
		runeArr[runeVal-'a']++
	}

	var buffer bytes.Buffer
	for _, runeVal := range runeArr {
		buffer.WriteRune(runeVal)
		// 用於分隔
		// 不然的話，1211....，分辨不出來 a 有 121 個還是 1 個還是怎樣
		buffer.WriteString("#")
	}
	return buffer.String()
}
