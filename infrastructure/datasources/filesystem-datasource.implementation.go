package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Alejoboga20/go-noc-system/domain/datasources"
	"github.com/Alejoboga20/go-noc-system/domain/entities"
)

type FileSystemDataSourceImplementation struct {
	FilePath string
}

func NewFileSystemDataSourceImplementation(filePath string) datasources.LogDataSource {
	return &FileSystemDataSourceImplementation{
		FilePath: filePath,
	}
}

func (fsds *FileSystemDataSourceImplementation) GetLogs(severityLevel entities.LogSeverityLevel) []entities.LogEnity {
	fmt.Println("Getting logs from file system")
	return []entities.LogEnity{}
}

func (fsds *FileSystemDataSourceImplementation) SaveLog(log entities.LogEnity) {
	if err := os.MkdirAll(fsds.FilePath, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	filePathBySeverity := filepath.Join(fsds.FilePath, fmt.Sprintf("%s.json", log.Level))
	file, err := os.OpenFile(filePathBySeverity, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	logEntryJson, err := json.Marshal(log)

	if err != nil {
		fmt.Println("Error marshalling log")
		return
	}

	_, err = file.Write(append(logEntryJson, '\n'))

	if err != nil {
		fmt.Println("Error writing log")
		return
	}

	fmt.Println("Log saved")
}
