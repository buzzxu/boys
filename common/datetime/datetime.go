package datetime

import "time"

const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

func NowUnix() int64 {
	return time.Now().Unix()
}

func UnixToTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}

func NowFormat(fmtStr string) string {
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	return time.Now().Format(fmtStr)
}

func UnixFormat(sec int64, fmtStr string) string {
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	return time.Unix(sec, 0).Format(fmtStr)
}

func Unix(datetime string) int64 {
	return ToUnix("2006-01-02 15:04:05", datetime)
}

func ToUnix(fmt, value string) int64 {
	return ToUnixLocation(fmt, value, "Asia/Shanghai")
}

func ToUnixLocation(fmt, value, localtion string) int64 {
	var loc *time.Location
	if localtion != "" {
		loc, _ = time.LoadLocation(localtion)
	} else {
		loc = time.Local
	}
	if fmt == "" {
		fmt = "2006-01-02 15:04:05"
	}
	t, _ := time.ParseInLocation(fmt, value, loc)
	return t.Unix()
}
