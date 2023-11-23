package money

// IntCentsToFloatYuan 整数分转浮点元
func IntCentsToFloatYuan(cents int64) float64 {
	return float64(cents) / 100.0
}

// FloatYuanToIntCents 浮点元转整数分
func FloatYuanToIntCents(yuan float64) int64 {
	return int64(yuan * 100)
}
