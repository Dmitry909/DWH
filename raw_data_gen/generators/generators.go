package generators

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

var coefsByWeekday = map[time.Weekday]float64{
	time.Monday:    1,
	time.Tuesday:   1.05,
	time.Wednesday: 1.05,
	time.Thursday:  1.1,
	time.Friday:    1.2,
	time.Saturday:  0.9,
	time.Sunday:    0.8,
}

var StartTimestamp time.Time = time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

func NumberOfOrdersPerDay(day int) float64 {
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
	ts := StartTimestamp.Add(time.Duration(24*day) * time.Hour)
	unexpectedCoeff := 1.0
	if day*241%40 == 26 {
		unexpectedCoeff = 0.2
	}
	return baseValue * coefsByWeekday[ts.Weekday()] * unexpectedCoeff
}

func GenerateExecutors(day, ordersNumber int) []string {
	rand.Seed(int64(day))
	randomCoeff := 0.9 + rand.Float64()*0.2
	baseAmount := float64(ordersNumber/2) * randomCoeff
	if day*239%40 == 26 {
		baseAmount *= 2
	}
	intBaseAmount := int(baseAmount)
	executors := make([]string, intBaseAmount)
	for i := 0; i < intBaseAmount; i++ {
		executors[i] = strconv.Itoa(i)
	}
	return executors
}

func CancelRate(day int) (float64, float64) {
	if day*157%30 == 26 {
		return 0.05, 0.45
	}
	return 0.05, 0.15
}

func GetRandomStatus(day int) string {
	notFinishedRate, cancelRate := CancelRate(day)
	status := "completed"
	rnd := rand.Float64()
	if rnd < notFinishedRate {
		status = "assigned"
	} else if rnd < cancelRate {
		status = "cancelled"
	}
	return status
}
