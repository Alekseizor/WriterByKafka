package app

import (
	"WriterByKafka/internal/app/config"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	config *config.Config
}

func New() (*Application, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	return &Application{
		config: cfg,
	}, nil
}

func (a *Application) Run() error {
	log.Println("application start running")
	a.StartServer()
	log.Println("application shut down")

	return nil
}
