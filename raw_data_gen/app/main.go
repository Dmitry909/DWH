package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"raw_data_gen/generators"
)

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

func main() {
	assignOrderId := 0
	for day := 1; day <= 300; day++ {
		ordersNumber := int(generators.NumberOfOrdersPerDay(day))
		executorsThisDay := generators.GenerateExecutors(day, ordersNumber)
		executorNow := 0
		timeStep := time.Duration(24*60*60/ordersNumber) * time.Second
		currentTimestamp := generators.StartTimestamp.Add(time.Duration(24*day) * time.Hour)
		for order := 0; order < ordersNumber; order++ {
			status := generators.GetRandomStatus(day)
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
