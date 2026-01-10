package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version string
	ServiceName string
	HttpPort int
}

var configurations Config

func loadConfig() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION not set in .env file")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME not set in .env file")
		os.Exit(1)
	}

	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		fmt.Println("HTTP_PORT not set in .env file")
		os.Exit(1)
	}

	httpPort, err := strconv.ParseInt(httpPortStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid HTTP_PORT value")
		os.Exit(1)
	}

	configurations = Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: int(httpPort),
	}

}

func GetConfig() Config {
	loadConfig()
	
	return configurations
}