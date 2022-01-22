package main

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	start "urlShortGen/cmd/internal"
	"urlShortGen/cmd/internal/api"
	"urlShortGen/cmd/internal/database"
	"urlShortGen/cmd/internal/database/inmemory"
	"urlShortGen/cmd/internal/database/pg"
	"urlShortGen/cmd/internal/service"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if err := initConfig(); err != nil {
		errorLog.Fatalf("error loading viper config: %s", err.Error())
	}

	var db database.UrlList
	mode, exists := os.LookupEnv("MEMORY_MODE")
	if exists && mode == "postgres" {
		var err error
		db, err = operations.NewPostgresDB(operations.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		})
		if err != nil {
			errorLog.Fatalf("startup error inmemory mode: %s\n", err.Error())
		}
	} else {
		db = inmemory.NewInMemoryDB()
	}

	requests := api.NewRequest(service.NewService(database.NewArchive(db), viper.GetInt("uniquestr.len"), []rune(viper.GetString("uniquestr.chars"))), errorLog, infoLog)

	srv := new(start.Server)
	infoLog.Printf("Wait for command on :%s\n", viper.GetString("port"))
	go func() {
		if err := srv.Run(viper.GetString("port"), requests.Routes()); err != nil {
			errorLog.Fatal(err)
		}
	}()

	closeSig := make(chan os.Signal, 1)
	signal.Notify(closeSig, os.Interrupt)
	<-closeSig

	if err := srv.Shutdown(context.Background()); err != nil {
		errorLog.Fatalf("error shutdown server: %s\n", err.Error())
	}

	switch db.(type) {
	case *operations.PostgresDB:
		postgresDB, ok := db.(*operations.PostgresDB)
		if !ok {
			errorLog.Fatalln("failed to convert")
		}
		if err := postgresDB.Close(); err != nil {
			errorLog.Fatalf("db connection closed error: %s\n", err.Error())
		}
	}
}

func initConfig() error {
	viper.AddConfigPath("env")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
