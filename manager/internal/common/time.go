package common

import "time"

const (
	TimeLayOut = "2006-01-02 15:04:05"
)

var (
	Location, _ = time.LoadLocation("Asia/Shanghai")
)

func TimeParseAsCST(timeStr string) (time.Time, error) {
	return time.ParseInLocation(TimeLayOut, timeStr, Location)
}
