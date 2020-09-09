package processor

import (
	"github.com/carloshjoaquim/E2Easy-Go/file_reader"
	"github.com/carloshjoaquim/E2Easy-Go/rest"
	log "github.com/sirupsen/logrus"
	"time"
)

type StepResult struct {
	Success    bool
	Message    string
	Time       time.Duration
	StatusCode int
}

func RunStep(step file_reader.Step) StepResult {
	log.Infof("Running step: %v", step.StepName)
	var stepResult StepResult
	switch step.Method {
	case "GET":
		{
			result, err := rest.Get(step.Path, step.Headers)
			if err != nil {
				log.Errorf("Error when trying to execute a GET request %v", err)
				stepResult = getErrorResult(err)
				break
			}
			stepResult = getSuccessResult(result)
		}
	case "POST":
		{
			result, err := rest.Post(step.Path, step.Body, step.Headers)
			if err != nil {
				log.Errorf("Error when trying to execute a POST request %v", err)
				stepResult = getErrorResult(err)
				break
			}
			stepResult = getSuccessResult(result)
		}
	}

	return stepResult
}

func getErrorResult(err error) StepResult {
	return StepResult{
		Success: false,
		Message: err.Error(),
		Time:    0,
	}
}

func getSuccessResult(result *rest.CallerResponse) StepResult {
	return StepResult{
		Success: true,
		Message: string(result.Body),
		Time:    result.RequestDuration,
		StatusCode: result.StatusCode,
	}
}
