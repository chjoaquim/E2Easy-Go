package rest

import (
	"github.com/carloshjoaquim/E2Easy-Go/src/file_reader"
	"github.com/carloshjoaquim/E2Easy-Go/src/processor"
	resty "github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

var (
	restClient = resty.New().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetTimeout(getTimeout() * time.Millisecond)
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

func Post(path string, body string, headers []file_reader.Headers) (*CallerResponse, error) {
	setHeaders(headers)
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

func Put(path string, body string, headers []file_reader.Headers) (*CallerResponse, error) {
	setHeaders(headers)
	response, err := restClient.
		SetHostURL(path).R().
		SetBody(body).
		Put("")

	if err != nil {
		log.Error("Error when trying to perform a PUT request due to: %s", err)
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

func Get(path string, headers []file_reader.Headers) (*CallerResponse, error) {
	setHeaders(headers)
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

func setHeaders(headers []file_reader.Headers) {
	if len(headers) > 0 {
		for _, header := range headers {
			restClient.SetHeader(header.Name, header.Value)
		}
	}
}

func getTimeout() time.Duration {
	timeout := processor.GetValueOfVar("timeout")
	if timeout != "" {
		conv,_ := strconv.ParseInt(timeout, 10, 64)
		return time.Duration(conv)
	} else {
		return 1000
	}
}
