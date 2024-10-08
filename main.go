package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Alejoboga20/go-noc-system/domain/usecases"
	datasources "github.com/Alejoboga20/go-noc-system/infrastructure/datasources"
	repositories "github.com/Alejoboga20/go-noc-system/infrastructure/repositories"
	services "github.com/Alejoboga20/go-noc-system/infrastructure/services"
	"github.com/joho/godotenv"
)

func success() {

}

func failure() {

}

var WEB_SERVICE_URL = "https://www.google.com"

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	intervalStr := os.Getenv("PING_INTERVAL")
	filePath := os.Getenv("FILE_PATH")

	pingInterval, err := strconv.Atoi(intervalStr)

	if err != nil {
		log.Fatal("Error parsing PING_INTERVAL")
	}

	ticker := time.NewTicker(time.Duration(pingInterval) * time.Second)
	defer ticker.Stop()

	httpClient := services.NewHTTPClientImplementation()
	fileSystemDataSource := datasources.NewFileSystemDataSourceImplementation(filePath)
	logRepository := repositories.NewLogRepositoryImplementation(fileSystemDataSource)

	checkService := usecases.NewCheckServiceImplementation(
		httpClient,
		logRepository,
		success,
		failure)

	for range ticker.C {
		checkService.Check(WEB_SERVICE_URL)
	}
}
