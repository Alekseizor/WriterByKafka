package main

import (
	"WriterByKafka/internal/pkg/app"
	"context"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

//@title Marketplace API
//@version 1.0
//@description API Server for Marketplace

// @contact.name Aleksei
// @contact.url https://vk.com/alekseyzorkin
// @contact.email zorkin22@bk.ru

// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	ctx := context.Background()
	log.Println("app start")

	application, err := app.New()
	if err != nil {
		log.Printf("cant create application: %s", err)
		os.Exit(2)
	}
	err = application.Run()
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t run application")
		log.Printf("can`t run application: %s", err)

		os.Exit(2)
	}

}
