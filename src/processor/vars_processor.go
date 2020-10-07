package processor

import (
	"encoding/json"
	"strconv"
	"strings"
)

var (
	globalVars = make(map[string]string, 0)
)

func GetVarsFromResponse(expectedVars map[string]string, result StepResult) map[string]string {
	if len(expectedVars) > 0 {
		for k, v := range expectedVars {
			evaluated := getValueFromResult(v, result)
			globalVars[k] = evaluated
		}
	}

	return globalVars
}

func GetValueOfVar(varName string) string {
	cleanedVar := strings.ReplaceAll(strings.ReplaceAll(varName, "${", ""), "}", "")

	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(globalVars[cleanedVar], "{", ""), "}", ""), string('"'), "")
}

func getValueFromResult(varName string, result StepResult) string {
	values := strings.Split(varName, ".")
	var value string
	switch values[0] {
	case "response": {
		switch values[1] {
		case "body": {
			if len(values) == 2 {
				return result.Message
			} else  {
				bodyJson := []byte(result.Message)
				c := make(map[string]json.RawMessage)

				json.Unmarshal(bodyJson, &c)
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

