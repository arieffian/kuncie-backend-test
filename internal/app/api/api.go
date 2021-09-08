package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/arieffian/kuncie-backend-test/internal/config"
	"github.com/arieffian/kuncie-backend-test/internal/connectors"
	"github.com/arieffian/kuncie-backend-test/internal/mutations"
	"github.com/arieffian/kuncie-backend-test/internal/queries"
	helper "github.com/arieffian/kuncie-backend-test/pkg/helpers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	log "github.com/sirupsen/logrus"
)

var (
	// userLogger instance of logrus logger
	apiLogger = log.WithField("go", "API")

	// graphql schema
	Schema graphql.Schema
)

func Start() {
	configureLogging()
	log.Infof("Starting api service")
	InitializeDB()

	Schema, err := graphql.NewSchema(InitializeGraphQL())

	if err != nil {
		apiLogger.Errorln(err)
		panic("cannot init schema")
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &Schema,
	})

	var wait time.Duration

	address := fmt.Sprintf("%s:%s", config.Get("server.host"), config.Get("server.port"))
	log.Info("Server binding to ", address)

	srv := &http.Server{
		Addr:         address,
		Handler:      httpHandler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Infof("Shutting down")
	os.Exit(0)
}

func InitializeDB() {
	if config.Get("db.type") == "mysql" {
		log.Warnf("Using MYSQL")

		queries.TransactionRepo = connectors.GetMySQLDBInstance()

		mutations.TransactionRepo = connectors.GetMySQLDBInstance()
		mutations.ProductRepo = connectors.GetMySQLDBInstance()
		mutations.UserRepo = connectors.GetMySQLDBInstance()
	} else {
		apiLogger.Fatal("unknown database type")
		panic(fmt.Sprintf("unknown database type %s. Correct your configuration 'db.type' or env-var 'AAA_DB_TYPE'. allowed values are INMEMORY or MYSQL", config.Get("db.type")))
	}

}

func configureLogging() {
	lLevel := config.Get("server.log.level")
	fmt.Println("Setting log level to ", lLevel)
	switch strings.ToUpper(lLevel) {
	default:
		fmt.Println("Unknown level [", lLevel, "]. Log level set to ERROR")
		log.SetLevel(log.ErrorLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	}

	lType := config.Get("log.type")
	fmt.Println("Setting log type to ", lType)
	if lType == "FILE" {
		helper.InitLogRotate()
		logFile := helper.GetFileLog()
		if logFile != nil {
			log.SetOutput(logFile)
		}
	}
}

func InitializeGraphQL() graphql.SchemaConfig {
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	}

	return schemaConfig
}
