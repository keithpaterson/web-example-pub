package main

import (
	"errors"
	"fmt"
	"os"

	"webkins/service"
	"webkins/service/logger"
)

var (
	// ErrorMissingServerOrLogger - error message returned during bad setup
	ErrorMissingServer error = errors.New("missing server")
)

func main() {
	fmt.Printf("Starting server and logging")

	// Setup the logger
	logger, err := logger.RootLogger()
	if err != nil {
		fmt.Println("******** COULD NOT CREATE A LOGGER!!!!!!! ************")
		fmt.Println("\t", err)
		os.Exit(-1)
	}

	// Probably want to get the port from config
	server := service.NewService(8080)

	// This sets the logger for all endpoints
	server.SetLogger(logger)

	// Setup and start the server
	err = server.Run()

	if err != nil {
		fmt.Println("Server had a problem starting: ", err)
	}
}
