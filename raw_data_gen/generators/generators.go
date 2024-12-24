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

type Executor struct {
	Id     string
	Rating float32
}

var StartTimestamp time.Time = time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

func NumberOfOrdersPerDay(day int) float64 {
	if day < 0 {
		panic("Wrong day passed to numberOfOrdersPerDay")
	}

	var baseValue float64 = 0
	baseDay := day % 300
	if baseDay <= 30 {
		baseValue = float64(3*baseDay) + float64(rand.Int()%10)
	} else if baseDay <= 167 {
		baseValue = (math.Sqrt(float64(-4*baseDay*baseDay+1040*baseDay-27199)) + 160) / 2
	} else {
		baseValue = math.Sqrt(float64(baseDay-166)) + 173
	}

	ts := StartTimestamp.Add(time.Duration(24*day) * time.Hour)
	unexpectedCoeff := 1.0

	if baseDay*241%40 == 26 {
		unexpectedCoeff = 0.2
	}

	return baseValue * coefsByWeekday[ts.Weekday()] * unexpectedCoeff
}

func GenerateExecutors(day, ordersNumber int) []Executor {
	rand.Seed(int64(day))
	randomCoeff := 0.9 + rand.Float64()*0.2
	baseAmount := float64(ordersNumber/2) * randomCoeff

	if day*239%40 == 26 {
		baseAmount *= 2
	}

	intBaseAmount := int(baseAmount)
	executors := make([]Executor, intBaseAmount)

	for i := 0; i < intBaseAmount; i++ {
		executors[i].Id = strconv.Itoa(i)
		executors[i].Rating = rand.Float32()*(5.0-4.0) + 4.0
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
	notFinishedRate, _ := CancelRate(day)
	status := "completed"
	rnd := rand.Float64()
	if rnd < notFinishedRate {
		status = "cancelled"
	}
	return status
}

func GetRandomAmount(day int) (int32, int32) {
	if rand.Int()%10 == 0 {
		return rand.Int31n(1500-189) + 189, rand.Int31n(500-200) + 200
	}
	return rand.Int31n(1500-189) + 189, 0
}

func GetRandomZoneId(day int) string {
	return strconv.Itoa(rand.Int() % 132)
}

func GetRandomFallback(day int) bool {
	return rand.Int()%100 == 0
}

func GetRandomCompletedTime(start time.Time, amount int32) time.Time {
	return start.Add(time.Duration(amount/3)*time.Minute + time.Duration(rand.Int()%10))
}
