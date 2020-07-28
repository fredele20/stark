package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	port int
}

func readConfig() Config {
	portString := os.Getenv("PORT")

	if portString == "" {
		portString = "8000"
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		panic(fmt.Sprintf("Could not parse %s to int", portString))
	}

	return Config{port:port}
}