package services

type HTTPClient interface {
	GET(url string) error
}
