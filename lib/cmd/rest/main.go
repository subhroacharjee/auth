package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/subhroacharjee/auth/lib/controller"
	"github.com/subhroacharjee/auth/lib/routes"
	"github.com/subhroacharjee/auth/lib/services/migrations"
	"github.com/subhroacharjee/auth/lib/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DEV  = "dev"
	PROD = "prod"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = DEV
	}

	if env == DEV {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9000"
	}
	postgresDSN, err := util.GetDbDSN()
	if err != nil {
		return err
	}
	db, err := gorm.Open(postgres.Open(*postgresDSN), &gorm.Config{})
	if err = migrations.MigrateModels(db); err != nil {
		return err
	}
	controller := controller.NewController(controller.ResolverOptions{
		DB: db,
	})

	router := routes.NewRouter(*controller)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	server.ListenAndServe()

	return nil
}
