package model

import "github.com/mangenotwork/extras/common/utils"

func SplitUId(uid string) (int, int) {
	if len(uid) < 4 {
	return 0,0
	}
	return utils.Str2Int(uid[0:4]), utils.Str2Int(uid[4:len(uid)-1])
}