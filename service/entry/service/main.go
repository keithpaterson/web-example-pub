package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"webkins/service"
	"webkins/service/logger"
)

const (
	defaultPort = 8080
	portEnvKey  = "BODKINS_PORT"
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

	// Get the port we want to use for the server, default to 8080
	port := getPort()

	// Probably want to get the port from config
	logger.Infow("CreateService", "Port", port)
	fmt.Println("Create service on port", port)
	server := service.NewService(port)

	// This sets the logger for all endpoints
	server.SetLogger(logger)

	// Setup and start the server
	err = server.Run()

	if err != nil {
		fmt.Println("Server had a problem starting: ", err)
	}
}

func getPort() int {
	// need some protection here so that prod can't be fiddled with, but for now read from env/command-line
	//   probably best to disable this with a 'prod' build tag (or only enable it with a 'dev' build tag)
	// prefer command-line to env
	for _, arg := range os.Args[1:] {
		if port, err := strconv.Atoi(arg); err == nil {
			return port
		}
		fmt.Println("WARNING: argument", arg, "could not be converted to a port number, so it was ignored")
		continue
	}
	if value, exists := os.LookupEnv(portEnvKey); exists {
		if port, err := strconv.Atoi(value); err == nil {
			return port
		}
	}
	return defaultPort
}
