package util

import (
	"math"
	"net/http"
	"time"
)

var cst *time.Location

func init() {
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}

	// 默认设置为中国时区
	time.Local = cst
}

// CSTLayoutString 格式化时间
// 返回 "2006-01-02 15:04:05" 格式的时间
func CSTLayoutString() string {
	return time.Now().In(cst).Format(time.DateTime)
}

// ParseCSTInLocation 格式化时间
func ParseCSTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(time.DateTime, date, cst)
}

// CSTLayoutStringToUnix 返回 unix 时间戳
// 2020-01-24 21:11:11 => 1579871471
func CSTLayoutStringToUnix(cstLayoutString string) (int64, error) {
	stamp, err := time.ParseInLocation(time.DateTime, cstLayoutString, cst)
	if err != nil {
		return 0, err
	}
	return stamp.Unix(), nil
}

// GMTLayoutString 格式化时间
// 返回 "Mon, 02 Jan 2006 15:04:05 GMT" 格式的时间
func GMTLayoutString() string {
	return time.Now().In(cst).Format(http.TimeFormat)
}

// ParseGMTInLocation 格式化时间
func ParseGMTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(http.TimeFormat, date, cst)
}

// BeginOfDay 当天 0 点
func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay 当天最后时刻
func EndOfDay(t time.Time) time.Time {
	return BeginOfTomorrow(t).Add(-time.Nanosecond)
}

// BeginOfYesterday 昨天 0 点
func BeginOfYesterday(t time.Time) time.Time {
	return BeginOfDay(t.AddDate(0, 0, -1))
}

// EndOfYesterday 昨天最后时刻
func EndOfYesterday(t time.Time) time.Time {
	return EndOfDay(t.AddDate(0, 0, -1))
}

// BeginOfTomorrow 明天 0 点
func BeginOfTomorrow(t time.Time) time.Time {
	return BeginOfDay(t.AddDate(0, 0, 1))
}

// EndOfTomorrow 明天 0 点
func EndOfTomorrow(t time.Time) time.Time {
	return EndOfDay(t.AddDate(0, 0, 1))
}

// BeginOfSecond 0 毫秒
func BeginOfSecond(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

// EndOfSecond 最后一毫秒
func EndOfSecond(t time.Time) time.Time {
	return BeginOfSecond(t).Add(time.Second - time.Nanosecond)
}

// BeginOfMinute 0 秒
func BeginOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute 最后一秒
func EndOfMinute(t time.Time) time.Time {
	return BeginOfMinute(t).Add(time.Minute - time.Nanosecond)
}

// BeginOfHour 0 分
func BeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour 最后一分
func EndOfHour(t time.Time) time.Time {
	return BeginOfHour(t).Add(time.Hour - time.Nanosecond)
}

// BeginOfWeek 本周一 0 点
func BeginOfWeek(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	return BeginOfDay(t).AddDate(0, 0, offset)
}

// EndOfWeek 本周末最后一刻
func EndOfWeek(t time.Time) time.Time {
	return BeginOfNextWeek(t).Add(-time.Nanosecond)
}

// BeginOfLastWeek 上周一 0 点
func BeginOfLastWeek(t time.Time) time.Time {
	return BeginOfWeek(t.AddDate(0, 0, -7))
}

// EndOfLastWeek 上周一最后一刻
func EndOfLastWeek(t time.Time) time.Time {
	return EndOfWeek(t.AddDate(0, 0, -7))
}

// BeginOfNextWeek 下周一 0 点
func BeginOfNextWeek(t time.Time) time.Time {
	return BeginOfWeek(t.AddDate(0, 0, 7))
}

// EndOfNextWeek 下周一最后一刻
func EndOfNextWeek(t time.Time) time.Time {
	return EndOfWeek(t.AddDate(0, 0, 7))
}

// BeginOfMonth 当月第一天 0 点
func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 当月最后一刻
func EndOfMonth(t time.Time) time.Time {
	return BeginOfNextMonth(t).Add(-time.Nanosecond)
}

// BeginOfLastMonth 上月第一天 0 点
func BeginOfLastMonth(t time.Time) time.Time {
	return BeginOfMonth(BeginOfMonth(t).AddDate(0, 0, -1))
}

// EndOfLastMonth 上月最后一刻
func EndOfLastMonth(t time.Time) time.Time {
	return BeginOfMonth(t).Add(-time.Nanosecond)
}

// BeginOfNextMonth 下月第一天 0 点
func BeginOfNextMonth(t time.Time) time.Time {
	return BeginOfMonth(BeginOfMonth(t).AddDate(0, 0, 31))
}

// EndOfNextMonth 下月最后一刻
func EndOfNextMonth(t time.Time) time.Time {
	return BeginOfMonth(BeginOfMonth(t).AddDate(0, 0, 62)).Add(-time.Nanosecond)
}

// GetMonthDays 当月天数
func GetMonthDays(t time.Time) int {
	return int(BeginOfNextMonth(t).Sub(BeginOfMonth(t)).Hours() / 24)
}

// BeginOfYear 本年第一天 0 点
func BeginOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 本年最后一刻
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// IsLeapYear 判断是否为闰年
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInYear 返回年份天数
func DaysInYear(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// SubInLocation 计算时间差
func SubInLocation(ts time.Time) float64 {
	return math.Abs(time.Now().In(cst).Sub(ts).Seconds())
}
