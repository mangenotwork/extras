package utils

import "strconv"

// Str2Int64 string -> int
func Str2Int(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	return 0
}

// Str2Int64 string -> int32
func Str2Int32(str string) int32 {
	if i, err := strconv.ParseInt(str, 10, 32); err == nil {
		return int32(i)
	}
	return 0
}

// Str2Int64 string -> int64
func Str2Int64(str string) int64 {
	if i, err := strconv.ParseInt(str, 10, 64); err == nil {
		return i
	}
	return 0
}

func Str2Bool(str string) bool {
	for _, v := range []string{"0","f","F","FALSE","false","False","no","否"}{
		if str == v {
			return false
		}
	}
	return true
}

func Int642Str(i int64) string {
	return strconv.FormatInt(i,10)
}