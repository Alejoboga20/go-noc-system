package entities

type LogSeverityLevel string

const (
	LOW    LogSeverityLevel = "LOW"
	MEDIUM LogSeverityLevel = "MEDIUM"
	HIGH   LogSeverityLevel = "HIGH"
)

type LogEnity struct {
	Level      LogSeverityLevel `json:"severity_level"`
	ServiceURL string           `json:"service_url"`
	Message    string           `json:"message"`
	Origin     string           `json:"origin"`
	CreatedAt  string           `json:"created_at"`
}
