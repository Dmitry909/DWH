package database

import (
	"os"
	"raw_data_gen/model"

	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	host        = "rc1b-xy6y7apt7j7jdpc7.mdb.yandexcloud.net,rc1d-opy4a78yulgzu7z2.mdb.yandexcloud.net"
	port        = 6432
	user        = "user1"
	password    = "NgdXRLUNn67d8tR"
	dbname      = "db1"
	sslrootcert = "/home/dmitry/hse/4year/DWH/raw_data_gen/database/root.crt"
	timeZone    = "Europe/Moscow"
)

var db *pgxpool.Pool

func EstablishConnection() (err error) {
	connString := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=require sslrootcert=%s target_session_attrs=read-write",
		host, port, dbname, user, password, sslrootcert)

	db, err = pgxpool.New(context.Background(), connString)
	return err
}

func AddAssignedOrder(assignedOrder *model.AssignedOrder) (bool, error) {
	_, err := db.Exec(context.Background(), `
		INSERT INTO assigned_orders (
				AssignedOrderId,
				OrderId,
				ExecutorId,
				ExecutionStatus,
				CoinCoefficient,
				CoinBonusAmount,
				FinalCoinAmount,
				ZoneName,
				HasExecutorFallbackBeenUsed,
				AssignTime,
				FirstAcquireTime
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		assignedOrder.AssignedOrderId,
		assignedOrder.OrderId,
		assignedOrder.ExecutorId,
		assignedOrder.ExecutionStatus,
		assignedOrder.CoinCoefficient,
		assignedOrder.CoinBonusAmount,
		assignedOrder.FinalCoinAmount,
		assignedOrder.ZoneName,
		assignedOrder.HasExecutorFallbackBeenUsed,
		assignedOrder.AssignTime,
		assignedOrder.FirstAcquireTime,
	)
	if err != nil {
		fmt.Println("db error:", err)
		os.Exit(1)
	}
	return true, err
}
