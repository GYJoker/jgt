package times

import "time"

// GetTimePointer 获取时间指针
func GetTimePointer(t time.Time) *time.Time {
	return &t
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

// ParseTime 解析时间
func ParseTime(t string) time.Time {
	tt, _ := time.Parse("2006-01-02 15:04:05", t)
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
