package usecases

import (
	"fmt"

	"github.com/Alejoboga20/go-noc-system/domain/repository"
	"github.com/Alejoboga20/go-noc-system/domain/services"
)

type CheckService interface {
	Check(serviceUrl string)
}

type CheckServiceImplementation struct {
	HTTPClient    services.HTTPClient
	LogRepository repository.LogRepository
	OnSuccess     func()
	OnFailure     func()
}

func NewCheckServiceImplementation(httpClient services.HTTPClient, logRepository repository.LogRepository, onSuccess func(), onFailure func()) CheckService {
	return &CheckServiceImplementation{
		HTTPClient:    httpClient,
		LogRepository: logRepository,
		OnSuccess:     onSuccess,
		OnFailure:     onFailure,
	}
}

func (checkService *CheckServiceImplementation) Check(serviceUrl string) {
	err := checkService.HTTPClient.GET(serviceUrl)

	if err != nil {
		fmt.Println("Error: ", err)
		checkService.OnFailure()
		return
	}

	checkService.OnSuccess()
}
