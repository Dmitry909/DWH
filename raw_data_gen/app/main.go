package main

import (
	"raw_data_gen/database"
	"raw_data_gen/generators"
	"raw_data_gen/model"

	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/dnlo/struct2csv"
	"github.com/spf13/pflag"
)

var (
	needCSV      = pflag.Bool("need-csv", true, "is need csv for DWH")
	needPushData = pflag.Bool("need-push-data", false, "is need push data to main DB")
	daysCount    = pflag.Int("days-count", 1000, "days of generated data")
)

func addOrder(row model.AssignedOrder) {
	database.AddAssignedOrder(&row)
}

func main() {
	pflag.Parse()

	if *needPushData {
		err := database.EstablishConnection()
		if err != nil {
			log.Printf("Error connecting to database: %s; retrying...\n", err.Error())
			return
		}
	}

	assignOrderId := 0

	buff := &bytes.Buffer{}
	enc := struct2csv.NewWriter(buff)
	err := enc.WriteColNames(model.DWHOrder{})
	if err != nil {
		panic(err)
	}

	for day := 1; day <= *daysCount; day++ {
		ordersNumber := int(generators.NumberOfOrdersPerDay(day))
		fmt.Println(day, ordersNumber)

		executorsThisDay := generators.GenerateExecutors(day, ordersNumber)
		executorNow := rand.Intn(len(executorsThisDay))
		timeStep := time.Duration(24*60*60/ordersNumber) * time.Second
		currentTimestamp := generators.StartTimestamp.Add(time.Duration(24*day) * time.Hour)

		for order := 0; order < ordersNumber; order++ {
			orderId := strconv.Itoa(assignOrderId)
			status := generators.GetRandomStatus(day)
			assignTime := currentTimestamp.Add(time.Duration(order) * timeStep)
			acquireTime := assignTime.Add(time.Duration(rand.Float64()*60) * time.Second)
			zoneId := generators.GetRandomZoneId(day)
			baseAmount, bonusAmount := generators.GetRandomAmount(day)
			completedTime := generators.GetRandomCompletedTime(acquireTime, baseAmount+bonusAmount)

			row := model.AssignedOrder{
				AssignedOrderId:             strconv.Itoa(assignOrderId),
				OrderId:                     orderId,
				ExecutorId:                  executorsThisDay[executorNow].Id,
				ExecutorRating:              executorsThisDay[executorNow].Rating,
				ExecutionStatus:             status,
				CoinCoefficient:             0.0,
				CoinBonusAmount:             bonusAmount,
				FinalCoinAmount:             baseAmount + bonusAmount,
				ZoneId:                      zoneId,
				HasExecutorFallbackBeenUsed: generators.GetRandomFallback(day),
				AssignTime:                  assignTime,
				FirstAcquireTime:            acquireTime,
				CompletedTime:               completedTime,
			}

			if *needCSV {
				err := enc.WriteStruct(model.DWHOrder{
					AcquireSeconds:  int32(completedTime.Sub(acquireTime).Seconds()),
					OrderId:         orderId,
					ExecutorId:      executorsThisDay[executorNow].Id,
					ExecutorRating:  executorsThisDay[executorNow].Rating,
					ExecutionStatus: status,
					CoinBonusAmount: bonusAmount,
					FinalCoinAmount: baseAmount + bonusAmount,
					ZoneId:          zoneId,
					AssignTime:      assignTime.String(),
				})
				if err != nil {
					panic(err)
				}
			}

			if *needPushData {
				addOrder(row)
			}
			assignOrderId++

			currentTimestamp.Add(timeStep)
		}
	}

	if *needCSV {
		fo, err := os.Create("output.csv")
		if err != nil {
			panic(err)
		}
		defer fo.Close()
		fo.Write(buff.Bytes())
	}
}
