package utils

import "strconv"

// Str2Int64 string -> int64
func Str2Int64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func Str2Bool(str string) bool {
	for _, v := range []string{"0","f","F","FALSE","false","False","no","否"}{
		if str == v {
			return false
		}
	}
	return true
}