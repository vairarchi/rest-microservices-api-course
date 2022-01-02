package main

import (
	"github.com/vairarchi/rest-microservices-api-course/banking/logger"

	"github.com/vairarchi/rest-microservices-api-course/banking/app"
)

func main() {
	logger.Info("Starting the application at port :8000")
	app.Start()
}
