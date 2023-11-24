package lunar_calendar

import "testing"

func TestLunarToSolar(t *testing.T) {
	c := LunarToSolar(2023, 7, 7)
	t.Log(c)

	c = LunarToSolar(2021, 10, 2)
	t.Log(c)

	c = LunarToSolar(2024, 1, 1)
	t.Log(c)
}
