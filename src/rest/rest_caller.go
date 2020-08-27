package rest

import (
	resty "github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	restClient = resty.New().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetTimeout(10000 * time.Millisecond)
)

type CallerResponse struct {
	Body            []byte
	Status          string
	StatusCode      int
	RequestDuration time.Duration
}

func GetRestClient() *http.Client {
	return restClient.GetClient()
}

func Post(path string, body string) (*CallerResponse, error) {
	response, err := restClient.
		SetHostURL(path).R().
		SetBody(body).
		Post("")

	if err != nil {
		log.Error("Error when trying to perform a POST request due to: %s", err)
		return nil, err
	}

	callerResponse := &CallerResponse{
		Body:            response.Body(),
		Status:          response.Status(),
		StatusCode:      response.StatusCode(),
		RequestDuration: response.Time(),
	}

	return callerResponse, nil
}

func Get(path string) (*CallerResponse, error) {
	response, err := restClient.
		SetHostURL(path).R().
		Get("")

	if err != nil {
		log.Error("Error when trying to perform a GET request due to: %s", err)
		return nil, err
	}

	callerResponse := &CallerResponse{
		Body:            response.Body(),
		Status:          response.Status(),
		StatusCode:      response.StatusCode(),
		RequestDuration: response.Time(),
	}

	return callerResponse, nil
}