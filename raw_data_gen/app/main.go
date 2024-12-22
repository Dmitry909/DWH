package main

import (
	"fmt"
	"math"
	"time"
)

var startTimestamp time.Time = time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
var coefsByWeekday = map[time.Weekday]float64{
	time.Monday:    1,
	time.Tuesday:   1.05,
	time.Wednesday: 1.05,
	time.Thursday:  1.1,
	time.Friday:    1.2,
	time.Saturday:  0.9,
	time.Sunday:    0.8,
}

func numberOfOrdersPerDay(day int) float64 {
	if day < 0 || day > 300 {
		panic("Wrong day passed to numberOfOrdersPerDay")
	}
	var baseValue float64 = 0
	if day <= 30 {
		baseValue = float64(3 * day)
	} else if day <= 167 {
		baseValue = (math.Sqrt(float64(-4*day*day+1040*day-27199)) + 160) / 2
	} else {
		baseValue = math.Sqrt(float64(day-166)) + 173
	}
	ts := startTimestamp.Add(time.Duration(24*day) * time.Hour)
	unexpectedCoeff := 1.0
	if day*241%40 == 26 {
		unexpectedCoeff = 0.2
	}
	return baseValue * coefsByWeekday[ts.Weekday()] * unexpectedCoeff
}

func main() {
	for day := 0; day <= 300; day++ {
		ordersNumber := int64(numberOfOrdersPerDay(day))
		fmt.Println(day, ordersNumber)
	}
}
