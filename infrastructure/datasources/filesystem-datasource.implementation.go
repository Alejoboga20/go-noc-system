package infrastructure

import (
	"fmt"

	"github.com/Alejoboga20/go-noc-system/domain/datasources"
	"github.com/Alejoboga20/go-noc-system/domain/entities"
)

type FileSystemDataSourceImplementation struct{}

func NewFileSystemDataSourceImplementation() datasources.LogDataSource {
	return &FileSystemDataSourceImplementation{}
}

func (fsds *FileSystemDataSourceImplementation) GetLogs(severityLevel entities.LogSeverityLevel) []entities.LogEnity {
	fmt.Println("Getting logs from file system")
	return []entities.LogEnity{}
}

func (fsds *FileSystemDataSourceImplementation) SaveLog(log string) {
	fmt.Println("Saving logs from file system")
	// Save log into file system
}
