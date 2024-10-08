package main

import (
	"fmt"

	"github.com/Alejoboga20/go-noc-system/domain/usecases"
	"github.com/Alejoboga20/go-noc-system/infrastructure"
)

func success() {
	fmt.Println("Success")
}

func failure() {
	fmt.Println("Failure")
}

func main() {
	httpClient := infrastructure.NewHTTPClientImplementation()
	checkService := usecases.NewCheckServiceImplementation(
		httpClient,
		nil,
		success,
		failure)

	checkService.Check("https://www.google.com/")
}
