package lunar_calendar

import "fmt"

type Calendar struct {
	Year   int
	Month  int
	Day    int
	IsLeap bool
}

func (c *Calendar) String() string {
	return fmt.Sprintf("Year:%d, Month:%d, Day:%d, IsLeap:%v", c.Year, c.Month, c.Day, c.IsLeap)
}

// GetBitInt 位操作
func GetBitInt(data, length, shift int) int {
	a := data & (((1 << length) - 1) << shift)
	b := a >> shift
	return b
}

// SolarFromInt 阳历，天转年月日
func SolarFromInt(g int) *Calendar {
	// go的向下取整
	y := (10000*g + 14780) / 3652425
	//y := math.Floor(float64(y1))
	ddd := g - (365*y + y/4 - y/100 + y/400)
	if ddd < 0 {
		y -= 1
		ddd = g - (365*y + y/4 - y/100 + y/400)
	}
	mi := (100*ddd + 52) / 3060
	mm := (mi+2)%12 + 1
	y += (mi + 2) / 12
	dd := ddd - (mi*306+5)/10 + 1
	return &Calendar{
		Year:  y,
		Month: mm,
		Day:   dd,
	}
}

// SolarToInt 阳历，年月日转天
func SolarToInt(y, m, d int) int {
	m = (m + 9) % 12
	//向下取整
	y -= m / 10
	a := 365*y + y/4 - y/100 + y/400 + (m*306+5)/10 + (d - 1)
	return a
}

// IsLeapMonth 判断是否闰月
func IsLeapMonth(days, m int) bool {
	leap := GetBitInt(days, 4, 13)
	a := leap != 0 && m > leap && m == leap+1
	return a
}

func LunarToSolar(y, m, d int) *Calendar {
	//阴历转阳历
	days := LMD[y-LMD[0]]
	leap := GetBitInt(days, 4, 13)
	offset := 0
	looped := leap

	if m <= leap || leap == 0 {
		looped = m - 1
	} else {
		looped = m
	}

	var i int
	for i = 0; i < looped; i++ {
		if GetBitInt(days, 1, 12-i) == 1 {
			offset += 30
		} else {
			offset += 29
		}
	}
	offset += d
	solar11 := S11[y-S11[0]]

	_y := GetBitInt(solar11, 12, 9)
	_m := GetBitInt(solar11, 4, 5)
	_d := GetBitInt(solar11, 5, 0)
	return SolarFromInt(SolarToInt(_y, _m, _d) + offset - 1)
}

// SolarToLunar 阳历转阴历
func SolarToLunar(y, m, d int) *Calendar {
	_y, _m, _d := 0, 0, 0
	index := y - S11[0]
	data := (y << 9) | (m << 5) | d
	if S11[index] > data {
		index -= 1
	}
	solar11 := S11[index]
	_y = GetBitInt(solar11, 12, 9)
	_m = GetBitInt(solar11, 4, 5)
	_d = GetBitInt(solar11, 5, 0)
	offset := SolarToInt(y, m, d) - SolarToInt(_y, _m, _d)

	days := LMD[index]
	_y, _m = index+S11[0], 1
	offset += 1
	var dm int
	var i int
	for i = 0; i < 13; i++ {
		if GetBitInt(days, 1, 12-i) == 1 {
			dm = 30
		} else {
			dm = 29
		}

		if offset > dm {
			_m += 1
			offset -= dm
		} else {
			break
		}
	}

	_d = offset
	if IsLeapMonth(days, _m) {
		_m -= 1
	}
	return &Calendar{
		Year:   _y,
		Month:  _m,
		Day:    _d,
		IsLeap: false,
	}
}
