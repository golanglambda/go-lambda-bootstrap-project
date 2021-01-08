package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golanglambda/go-lambda-bootstrap-project/internal"
	log "github.com/sirupsen/logrus"
)

func main() {

	//Load environment config
	cfg, err := internal.NewConfig()
	if err != nil {
		panic(fmt.Errorf("Unable to load config: %v", err))
	}

	//Set log configuration
	log.SetFormatter(&log.JSONFormatter{})
	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = log.InfoLevel //Default level
	}
	log.SetLevel(level)

	//Bootstrap Handler
	h, err := internal.BootstrapHandler(cfg)
	if err != nil {
		log.Fatalf("lambda function initialization failed: %v", err)
	}

	//call handler function
	lambda.Start(h.Execute)
}
