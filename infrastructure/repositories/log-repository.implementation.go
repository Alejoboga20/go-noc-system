package infrastructure

import (
	"github.com/Alejoboga20/go-noc-system/domain/datasources"
	"github.com/Alejoboga20/go-noc-system/domain/entities"
	"github.com/Alejoboga20/go-noc-system/domain/repository"
)

type LogRepositoryImplementation struct {
	DataSource datasources.LogDataSource
}

func NewLogRepositoryImplementation(dataSource datasources.LogDataSource) repository.LogRepository {
	return &LogRepositoryImplementation{
		DataSource: dataSource,
	}
}

func (logRepository *LogRepositoryImplementation) GetLogs(severityLevel entities.LogSeverityLevel) []entities.LogEnity {
	logs := logRepository.DataSource.GetLogs(severityLevel)

	return logs
}

func (logRepository *LogRepositoryImplementation) SaveLog(log entities.LogEnity) {
	logRepository.DataSource.SaveLog(log)
}
