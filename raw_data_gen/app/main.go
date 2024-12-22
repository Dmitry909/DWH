package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
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

func addOrder(
	assignOrderId string,
	orderId string,
	executorId string,
	status string,
	coinCoeff float64,
	coinBonusAmount float64,
	finalCoinAmount float64,
	usedExecutorFallback bool,
	assignTime time.Time,
	acquireTime time.Time,
) {
	fmt.Println("addOrder called:", assignOrderId, orderId, executorId, status, coinCoeff, coinBonusAmount, finalCoinAmount, usedExecutorFallback, assignTime, acquireTime)
}

func generateExecutors(day, ordersNumber int) []string {
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

func cancelRate(day int) (float64, float64) {
	if day*157%30 == 26 {
		return 0.05, 0.45
	}
	return 0.05, 0.15
}

func main() {
	assignOrderId := 0
	for day := 1; day <= 300; day++ {
		ordersNumber := int(numberOfOrdersPerDay(day))
		executorsThisDay := generateExecutors(day, ordersNumber)
		executorNow := 0
		notFinishedRate, cancelRate := cancelRate(day)
		timeStep := time.Duration(24*60*60/ordersNumber) * time.Second
		currentTimestamp := startTimestamp.Add(time.Duration(24*day) * time.Hour)
		for order := 0; order < ordersNumber; order++ {
			status := "completed"
			rnd := rand.Float64()
			if rnd < notFinishedRate {
				status = "assigned"
			} else if rnd < cancelRate {
				status = "cancelled"
			}
			assignTime := currentTimestamp.Add(time.Duration(order) * timeStep)
			acquireTime := assignTime.Add(time.Duration(rand.Float64()*60) * time.Second)
			addOrder(
				strconv.Itoa(assignOrderId),
				strconv.Itoa(assignOrderId),
				executorsThisDay[executorNow],
				status,
				0.0,
				0.0,
				0.0,
				false,
				assignTime,
				acquireTime,
			)
			assignOrderId++

			currentTimestamp.Add(timeStep)
		}
	}
}
