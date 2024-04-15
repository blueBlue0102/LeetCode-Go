package leetcode

import (
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	data := []rune(s)
	index := 0
	for {
		if data[index] >= '0' && data[index] <= '9' {
			// 遇到數字了，表示需要展開遞迴
			// 紀錄當下的 index，之後需要更新 data 為遞迴後的結果，data = append(data[:numStartAt], 遞迴 result)
			numStartAt := index
			recurIndex := index
			// 取得數字
			timesString := ""
			for {
				if data[recurIndex] >= '0' && data[recurIndex] <= '9' {
					timesString += string(data[recurIndex])
					recurIndex++
				} else {
					break
				}
			}
			times, _ := strconv.Atoi(timesString)
			// 找到數字之後，開始查找這層中括號的字串內容
			bracketCount := 1
			recurIndex++
			stringStartAt := recurIndex
			stringEndAt := recurIndex
			for bracketCount > 0 {
				if data[recurIndex] == '[' {
					bracketCount++
				} else if data[recurIndex] == ']' {
					bracketCount--
					stringEndAt = recurIndex
				}
				recurIndex++
			}
			// 遞迴處理找到的字串
			recurResult := DecodeString(string(data[stringStartAt:stringEndAt]))
			data = []rune(string(data[:numStartAt]) + strings.Repeat(recurResult, times) + string(data[stringEndAt+1:]))
			index--
		}

		if index < len(data)-1 {
			index++
		} else {
			break
		}

	}
	return string(data)
}
