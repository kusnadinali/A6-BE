package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"be-feeds/service"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
)

const (
	host       = "103.157.96.115"
	port       = 5432
	dbuser     = "shimano"
	dbpassword = "milkyway"
	dbname     = "db_shimano"
)

var db *sql.DB

func initDB() {
	dbSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, dbuser, dbpassword, dbname)

	var err error

	db, err = sql.Open("postgres", dbSource)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Now We Are Connected To PostgreSQL Database.")
}

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "feeds",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	initDB()
	flag.Parse()
	ctx := context.Background()

	var srv service.Service
	{
		svcRepository := service.NewRepo(db, logger)

		srv = service.NewService(svcRepository, logger)
	}

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := service.MakeFeedsEndpoints(srv)

	go func() {
		fmt.Println("Listening On Port", *httpAddr)
		handler := service.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errChan)
}
