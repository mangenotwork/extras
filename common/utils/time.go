package utils

import (
	"log"
	"time"
)

// Timestamp2Time 时间戳到时间
// use : Timestamp2Time(time.Now().Unix(), "20060102 15:04:05")
func Timestamp2Time(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

// 格式为  2006-01-02 15:04:05
func Str2Timestamp(str string) int64 {
	var (
		stamp time.Time
		err error
	)
	if stamp, err = time.ParseInLocation("2006-01-02 15:04:05", str, time.Local); err != nil {
		if stamp, err = time.ParseInLocation("2006-01-02", str, time.Local); err != nil {
			if stamp, err = time.ParseInLocation("20060102 15:04:05", str, time.Local); err != nil {
				if stamp, err = time.ParseInLocation("20060102", str, time.Local); err != nil {
					log.Println("Str2Timestamp err = ", err)
				}
			}
		}
	}
	return stamp.Unix()
}

func Str2TimestampStr(str string) string {
	return Int642Str(Str2Timestamp(str))
}