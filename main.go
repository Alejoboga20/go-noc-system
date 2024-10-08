package main

import (
	"log"
	"os"
	"strconv"
	"strings"
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

func main() {
	var webServices []string
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	for _, env := range os.Environ() {
		keyValuePair := strings.Split(env, "=")
		key := keyValuePair[0]
		value := keyValuePair[1]

		if strings.HasPrefix(key, "WEB_SERVICE_URL_") && value != "" {
			webServices = append(webServices, value)
		}
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
		for _, serviceUrl := range webServices {
			checkService.Check(serviceUrl)
		}
	}
}
