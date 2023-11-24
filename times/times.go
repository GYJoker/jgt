package times

import (
	"fmt"
	"strings"
	"time"
)

// GetTimePointer 获取时间指针
func GetTimePointer(t time.Time) *time.Time {
	return &t
}

// GetNowPointer 获取当前时间指针
func GetNowPointer() *time.Time {
	t := time.Now()
	return GetTimePointer(t)
}

// GetSystemStartTime 获取系统开始时间
func GetSystemStartTime() time.Time {
	return ParseTime("2023-01-01 00:00:00")
}

// FormatTime 格式化时间
func FormatTime(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// FormatSystemStrTime 格式化时间
func FormatSystemStrTime(t string) string {
	if t == "" {
		return ""
	}

	parse, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return ""
	}

	return FormatTime(&parse)
}

var TimeLocation, _ = time.LoadLocation("Asia/Shanghai") // GetTimePointer 获取时间指针

// ParseTime 解析时间
func ParseTime(t string) time.Time {
	split := strings.Split(t, " ")
	if len(split) != 2 {
		return time.Time{}
	}

	dateFormat := checkUpdateDateFormat(split[0])
	if dateFormat == "" {
		return time.Time{}
	}

	timeFormat := checkUpdateDateTimeFormat(split[1])
	if timeFormat == "" {
		return time.Time{}
	}

	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", dateFormat, timeFormat), TimeLocation)
	return tt
}

// ParseTTTime 解析时间  2023-09-03T10:27:28+08:00
func ParseTTTime(t string) time.Time {
	tt, _ := time.Parse(time.RFC3339, t)
	return tt
}

// ParseDate 解析日期
func ParseDate(t string) time.Time {
	format := checkUpdateDateFormat(t)
	if format == "" {
		return time.Time{}
	}
	tt, _ := time.ParseInLocation("2006-01-02", format, TimeLocation)
	return tt
}

// OffsetMinuteTime 偏移时间 -- 分钟
func OffsetMinuteTime(t time.Time, offset int) time.Time {
	return t.Add(time.Duration(offset) * time.Minute)
}

// OffsetHourTime 偏移时间 -- 小时
func OffsetHourTime(t time.Time, offset int) time.Time {
	return t.Add(time.Duration(offset) * time.Hour)
}

// OffsetDayTime 偏移时间 -- 天
func OffsetDayTime(t time.Time, offset int) time.Time {
	return addDate(t, 0, 0, offset)
}

// OffsetWeekTime 偏移时间 -- 周
func OffsetWeekTime(t time.Time, offset int) time.Time {
	return addDate(t, 0, 0, offset*7)
}

// OffsetMonthTime 偏移时间 -- 月
func OffsetMonthTime(t time.Time, offset int) time.Time {
	return addDate(t, 0, offset, 0)
}

// OffsetYearTime 偏移时间 -- 年
func OffsetYearTime(t time.Time, offset int) time.Time {
	return addDate(t, offset, 0, 0)
}

// BeginOfDay 一天开始时间
func BeginOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay 一天结束时间
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

// BeginOfWeek 一周开始时间
func BeginOfWeek(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	return BeginOfDay(OffsetDayTime(t, offset))
}

// EndOfWeek 一周结束时间
func EndOfWeek(t time.Time) time.Time {
	offset := int(time.Sunday - t.Weekday())
	if offset < 0 {
		offset = 6
	}
	return EndOfDay(OffsetDayTime(t, offset-1))
}

// BeginOfMonth 一个月开始时间
func BeginOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 一个月结束时间
func EndOfMonth(t time.Time) time.Time {
	return EndOfDay(BeginOfMonth(OffsetMonthTime(t, 1)).AddDate(0, 0, -1))
}

// BeginOfYear 一年开始时间
func BeginOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 一年结束时间
func EndOfYear(t time.Time) time.Time {
	return EndOfDay(BeginOfYear(OffsetYearTime(t, 1)).AddDate(0, 0, -1))
}

// DiffDays 两个时间相差天数
func DiffDays(t1, t2 time.Time) int {
	diff := int(t1.Sub(t2).Hours() / 24)
	if t1.After(t2) && diff > 0 {
		diff = -diff
	}
	if t1.Before(t2) && diff < 0 {
		diff = -diff
	}
	return diff
}

func addDate(t time.Time, years, months, days int) time.Time {
	// limit month
	if months >= 12 || months <= 12 {
		years += months / 12
		months = months % 12
	}

	// get datetime parts
	ye := t.Year()
	mo := t.Month()
	da := t.Day()
	hr := t.Hour()
	mi := t.Minute()
	se := t.Second()
	ns := t.Nanosecond()
	lo := t.Location()

	// log.Printf("input: %d - %d - %d\n", ye, mo, da)
	// log.Printf("delta: %d - %d - %d\n", years, months, days)

	// years
	ye += years

	// months
	mo += time.Month(months)
	if mo > 12 {
		mo -= 12
		ye++
	} else if mo < 1 {
		mo += 12
		ye--
	}

	// after adding month, we should adjust day of month value
	if da <= 28 {
		// nothing to change
	} else if da == 29 {
		if mo == 2 {
			if !isLeapYear(ye) {
				da = 28
			}
		}
		// else, OK

	} else if da == 30 {
		if mo == 2 {
			if isLeapYear(ye) {
				da = 29
			} else {
				da = 28
			}
		}
		// else, OK

	} else if da == 31 {
		switch mo {
		case 2:
			if isLeapYear(ye) {
				da = 29
			} else {
				da = 28
			}
		case 1, 3, 5, 7, 8, 10, 12:
			da = 31
		case 4, 6, 9, 11:
			da = 30
		}
	}

	// date
	da += days

	// return
	return time.Date(ye, mo, da, hr, mi, se, ns, lo)
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

// checkUpdateDateFormat 检查修改日期格式为 2006-01-02
func checkUpdateDateFormat(date string) string {
	if date == "" {
		return ""
	}

	split := strings.Split(date, "-")
	if len(split) != 3 {
		return ""
	}

	if len(split[1]) == 1 {
		split[1] = "0" + split[1]
	}
	if len(split[2]) == 1 {
		split[2] = "0" + split[2]
	}

	return strings.Join(split, "-")
}

// checkUpdateDateTimeFormat 检查修改日期格式为 15:04:05
func checkUpdateDateTimeFormat(date string) string {
	if date == "" {
		return ""
	}

	split := strings.Split(date, ":")
	if len(split) != 3 {
		return ""
	}

	if len(split[0]) == 1 {
		split[0] = "0" + split[0]
	}
	if len(split[1]) == 1 {
		split[1] = "0" + split[1]
	}
	if len(split[2]) == 1 {
		split[2] = "0" + split[2]
	}

	return strings.Join(split, ":")
}
