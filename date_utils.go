package utils

import "time"

const (
	apiDateLaylout = "2006-01-02T15:04:05.000Z"
	apiDBLaylout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLaylout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLaylout)
}
