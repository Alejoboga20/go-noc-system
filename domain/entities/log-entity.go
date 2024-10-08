package entities

type LogSeverityLevel string

const (
	LOW    LogSeverityLevel = "LOW"
	MEDIUM LogSeverityLevel = "MEDIUM"
	HIGH   LogSeverityLevel = "HIGH"
)

type LogEnity struct {
	Level     LogSeverityLevel `json:"severity_level"`
	Message   string           `json:"message"`
	CreatedAt string           `json:"created_at"`
	Origin    string           `json:"origin"`
}
