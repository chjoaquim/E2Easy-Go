package processor

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func GetVarsFromResponse(expectedVars map[string]string, result StepResult) map[string]string {
	globalVars := make(map[string]string, 0)
	if len(expectedVars) > 0 {
		for k, v := range expectedVars {
			evaluated := getValueFromResult(v, result)
			globalVars[k] = evaluated
		}
	}

	return globalVars
}

func getValueFromResult(varName string, result StepResult) string {
	values := strings.Split(varName, ".")
	var value string
	log.Infof("varName: %s", varName)
	switch values[0] {
	case "response": {
		switch values[1] {
		case "body": {
			if len(values) == 2 {
				return result.Message
			} else  {
				bodyJson := []byte(result.Message)
				log.Infof("in Body: %s", bodyJson)
				c := make(map[string]json.RawMessage)

				json.Unmarshal(bodyJson, &c)
				log.Infof("in C: %s", c)
				value = string(c[values[2]])
			}
			break
		}
		case "statusCode": {
			value = strconv.Itoa(result.StatusCode)
		}
		}
	}
	default:
		value = ""
	}

	return value
}
