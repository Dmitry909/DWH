package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"raw_data_gen/database"
	"raw_data_gen/generators"
	"raw_data_gen/model"
)

func addOrder(row model.AssignedOrder) {
	database.AddAssignedOrder(&row)
}

func main() {
	err := database.EstablishConnection()
	if err != nil {
		log.Printf("Error connecting to database: %s; retrying...\n", err.Error())
		return
	}

	assignOrderId := 0
	days := 300
	for day := 1; day <= days; day++ {
		ordersNumber := int(generators.NumberOfOrdersPerDay(day))
		fmt.Println(day, ordersNumber)
		executorsThisDay := generators.GenerateExecutors(day, ordersNumber)
		executorNow := 0
		timeStep := time.Duration(24*60*60/ordersNumber) * time.Second
		currentTimestamp := generators.StartTimestamp.Add(time.Duration(24*day) * time.Hour)
		for order := 0; order < ordersNumber; order++ {
			status := generators.GetRandomStatus(day)
			assignTime := currentTimestamp.Add(time.Duration(order) * timeStep)
			acquireTime := assignTime.Add(time.Duration(rand.Float64()*60) * time.Second)
			row := model.AssignedOrder{
				AssignedOrderId:             strconv.Itoa(assignOrderId),
				OrderId:                     strconv.Itoa(assignOrderId),
				ExecutorId:                  executorsThisDay[executorNow],
				ExecutionStatus:             status,
				CoinCoefficient:             0.0,
				CoinBonusAmount:             0.0,
				FinalCoinAmount:             0.0,
				ZoneName:                    "Moscow",
				HasExecutorFallbackBeenUsed: false,
				AssignTime:                  assignTime,
				FirstAcquireTime:            acquireTime,
			}
			addOrder(row)
			assignOrderId++

			currentTimestamp.Add(timeStep)
		}
	}
}
