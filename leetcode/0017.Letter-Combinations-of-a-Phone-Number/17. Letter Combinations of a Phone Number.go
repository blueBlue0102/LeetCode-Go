package leetcode

var buttonMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func LetterCombinationsofaPhoneNumber(digits string) []string {
	return Enumerate(&[]string{}, digits, "")
}

func Enumerate(result *[]string, digits string, data string) []string {
	// 因 digits 已限定在 2~9，所以使用 len 來計算 digits 的長度
	if len(digits) == 1 {
		for _, str := range buttonMap[string(digits[0])] {
			*result = append(*result, data+str)
		}
	} else if len(digits) > 1 {
		for _, str := range buttonMap[string(digits[0])] {
			Enumerate(result, digits[1:], data+str)
		}
	}
	return *result
}
