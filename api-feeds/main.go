package main

import (
	//	"context"
	"database/sql"
	"fmt"

	//	"net/http"
	//	"os"
	//	"tes/service"
	//	"tes/service/repository"

	//	"github.com/go-kit/kit/log"
	//	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	connStr := "postgres://postgres:Kgdwfjrc123@localhost:5432/feeds?sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Now we are connected to POSTGRESQL DATABASE.")
}

func main() {
	initDB()

	// var logger log.Logger
	// {
	// 	logger = log.NewLogfmtLogger(os.Stdout)
	// 	logger = log.NewSyncLogger(logger)
	// 	logger = log.With(logger,
	// 		"service", "usermgmt",
	// 		"time", log.DefaultTimestampUTC,
	// 		"caller", log.DefaultCaller,
	// 	)
	// }

	// var srv service.Service
	// {
	// 	svcRepository := repository.NewRepo(db, logger)

	// 	srv = service.NewService(
	// 		svcRepository,
	// 		// configs,
	// 		logger,
	// 	)
	// }
	// ctx := context.Background()

	// endpoints := service.MakeFeedsEndpoints(srv)

	// errChan := make(chan error)

	// go func() {
	// 	handler := service.NewHTTPServer(ctx, endpoints)
	// 	errChan <- http.ListenAndServe(":8080", handler)

	// }()

	// level.Error(logger).Log("exit", <-errChan)
}
