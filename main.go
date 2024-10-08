package main

import (
	"fmt"
	"time"

	"github.com/Alejoboga20/go-noc-system/domain/usecases"
	datasources "github.com/Alejoboga20/go-noc-system/infrastructure/datasources"
	repositories "github.com/Alejoboga20/go-noc-system/infrastructure/repositories"
	services "github.com/Alejoboga20/go-noc-system/infrastructure/services"
)

func success() {
	fmt.Println("Success")
}

func failure() {
	fmt.Println("Failure")
}

var INTERVAL = 1 * time.Second

func main() {
	ticker := time.NewTicker(INTERVAL)
	defer ticker.Stop()

	httpClient := services.NewHTTPClientImplementation()
	fileSystemDataSource := datasources.NewFileSystemDataSourceImplementation()
	logRepository := repositories.NewLogRepositoryImplementation(fileSystemDataSource)

	checkService := usecases.NewCheckServiceImplementation(
		httpClient,
		logRepository,
		success,
		failure)

	for range ticker.C {
		checkService.Check("https://www.google.com/")
	}
}
