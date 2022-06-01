package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marcosvieirajr/payment/configs"
	"github.com/marcosvieirajr/payment/internal/app/usecases"
	"github.com/marcosvieirajr/payment/internal/platform/handlers"
	"github.com/marcosvieirajr/payment/internal/platform/repository"
	"github.com/sirupsen/logrus"
)

func main() {
	// loading config file
	configs := configs.Load()

	// preparing logrus log
	level, _ := logrus.ParseLevel(configs.LogLevel)
	log := logrus.New()
	log.SetLevel(level)
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// creating DB connection
	db, err := repository.NewDB(configs)
	if err != nil {
		log.Fatal(err)
	}

	// registering repositories
	accountRepository := repository.NewAccountRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	// registering gateways (optional)
	getAccountGateway := accountRepository
	createAccountGateway := accountRepository
	countAccountsByDocumentGateway := accountRepository
	createTransactionGateway := transactionRepository

	// registering use cases
	getAccountUseCase := usecases.NewGetAccountUseCase(getAccountGateway)
	createAccountUseCase := usecases.NewCreateAccountUseCase(countAccountsByDocumentGateway, createAccountGateway)
	createTransactionUseCase := usecases.NewCreateTransactionUseCase(getAccountGateway, createTransactionGateway)

	// creating handlers
	createAccountHandler := handlers.CreateAccount(createAccountUseCase)
	getAccountHandler := handlers.GetAccount(getAccountUseCase)
	createTransactionHandler := handlers.CreateTransaction(createTransactionUseCase)

	// registering routers
	engine := gin.Default()

	engine.GET("/health", handlers.CheckHealth(db))
	v1 := engine.Group("/v1")
	{
		v1.POST("/accounts", createAccountHandler)
		v1.GET("/accounts/:id", getAccountHandler)
		v1.POST("/transactions", createTransactionHandler)
	}

	// configuring HTTP server
	address := fmt.Sprintf("%v:%v", configs.HostName, configs.HostPort)
	srv := http.Server{
		Addr:         address,
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Infof("listening and serving HTTP on %s", address)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting server: %s", err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-wait
	log.Infof("received signal %v. Shutting down server gracefully ...", sig.String())
	log.Info("press Ctrl+C again to force")

	// Create a deadline to wait for. waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), configs.GracefulTimeout)
	defer cancel()

	// gracefully shutdown the server doesn't block if no connections,
	// but will otherwise wait until the timeout deadline.
	srv.Shutdown(ctx)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Info("server exiting")
	os.Exit(0)
}
