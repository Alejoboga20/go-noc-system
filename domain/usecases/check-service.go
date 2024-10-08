package usecases

import (
	"fmt"
	"time"

	"github.com/Alejoboga20/go-noc-system/domain/entities"
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
		log := entities.LogEnity{
			Message:    fmt.Sprintf("Service is down: %s", err.Error()),
			Level:      entities.HIGH,
			ServiceURL: serviceUrl,
			Origin:     "CheckService",
			CreatedAt:  time.Now().Format(time.RFC3339),
		}
		checkService.LogRepository.SaveLog(log)
		return
	}

	log := entities.LogEnity{
		Message:    "Service is up",
		Level:      entities.LOW,
		ServiceURL: serviceUrl,
		Origin:     "CheckService",
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	checkService.LogRepository.SaveLog(log)
	checkService.OnSuccess()
}
