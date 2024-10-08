package main

import (
	"fmt"
	"time"

	"github.com/Alejoboga20/go-noc-system/domain/usecases"
	"github.com/Alejoboga20/go-noc-system/infrastructure"
)

func success() {
	fmt.Println("Success")
}

func failure() {
	fmt.Println("Failure")
}

var INTERVAL = 10 * time.Second

func main() {
	ticker := time.NewTicker(INTERVAL)
	defer ticker.Stop()

	httpClient := infrastructure.NewHTTPClientImplementation()
	checkService := usecases.NewCheckServiceImplementation(
		httpClient,
		nil,
		success,
		failure)

	for range ticker.C {
		checkService.Check("https://www.google.com/")
	}
}
