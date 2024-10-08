package main

import (
	"time"

	"github.com/Alejoboga20/go-noc-system/domain/usecases"
	datasources "github.com/Alejoboga20/go-noc-system/infrastructure/datasources"
	repositories "github.com/Alejoboga20/go-noc-system/infrastructure/repositories"
	services "github.com/Alejoboga20/go-noc-system/infrastructure/services"
)

func success() {

}

func failure() {

}

var INTERVAL = 1 * time.Second
var filePath = "./logs"

func main() {
	ticker := time.NewTicker(INTERVAL)
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
		checkService.Check("https://www.googles.com/")
	}
}
