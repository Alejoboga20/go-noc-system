package datasources

import "github.com/Alejoboga20/go-noc-system/domain/entities"

type LogDataSource interface {
	GetLogs(severityLevel entities.LogSeverityLevel) []entities.LogEnity
	SaveLog(log string)
}
