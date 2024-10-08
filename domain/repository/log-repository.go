package repository

import "github.com/Alejoboga20/go-noc-system/domain/entities"

type LogRepository interface {
	GetLogs(severityLevel entities.LogSeverityLevel) []entities.LogEnity
	SaveLog(log entities.LogEnity)
}
