package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"maker-checker/cmd/maker-checker/config"
	"maker-checker/logger"
	"maker-checker/store/psql"
	httptransport "maker-checker/transport/http"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func getLogProfile() (*zap.Logger, error) {
	if config.Profile == "prod" {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}

func main() {
	v := viper.New()
	err := config.Init(v)
	if err != nil {
		panic(err)
	}

	if err = config.Load(v); err != nil {
		panic(err)
	}

	logBackend, err := getLogProfile()
	if err != nil {
		panic(err)
	}

	l := logger.NewContextLogger(logBackend.Sugar())
	log := l.With("service", "maker-checker")

	defer func() {
		_ = log.Sync()
	}()

	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConfig := &psql.Config{
		User:               config.DbUser,
		Password:           config.DbPassword,
		Host:               config.DbHost,
		Port:               config.DbPort,
		DbName:             config.DbName,
		SkipMigrations:     config.SkipMigrations,
		MigrationsLocation: config.MigrationsLocation,
	}

	psql.RunMigrations(mainCtx, log, dbConfig)

	db, err := psql.NewConnection(&psql.Config{
		User:     config.DbUser,
		Password: config.DbPassword,
		Host:     config.DbHost,
		Port:     config.DbPort,
		DbName:   config.DbName,
	})

	if err != nil {
		log.Fatalw(mainCtx, "failed to setup db", "err", err)
	}

	var a httptransport.ApiHandler
	{
		a = httptransport.NewHandler(log, db)
		a = httptransport.LoggingMiddleware(log)(a)
	}

	var h http.Handler
	{
		h = httptransport.MakeHTTPHandler(log, a)
	}

	errs := make(chan error)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-sig)
	}()

	go func() {
		log.Infow(mainCtx, "Starting HTTP server", "port", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	log.Info(mainCtx, "exit", <-errs)
}
