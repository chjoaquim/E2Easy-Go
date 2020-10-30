package processor

import (
	"encoding/json"
	"fmt"
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
	case "response":
		{
			switch values[1] {
			case "body":
				{
					if len(values) == 2 {
						return result.Message
					} else {
						bodyJson := []byte(result.Message)
						c := make(map[string]interface{})
						json.Unmarshal(bodyJson, &c)

						parsed, _ := findMapVar(c, removeSliceItems(values, 2)...)
						value = fmt.Sprintf("%s", parsed)
					}
					break
				}
			case "statusCode":
				{
					value = strconv.Itoa(result.StatusCode)
				}
			}
		}
	default:
		value = ""
	}
	return value
}

func AddTestVar(varName string, testResult []TestResult, configName string) {
	testError := false
	for _, t := range testResult {
		testError = testError && t.Result
	}

	globalTests, _ := strconv.ParseBool(globalVars[configName])
	globalVars[configName] = fmt.Sprintf("%v", globalTests && testError)
	globalVars[varName] = fmt.Sprintf("%v", testError)
}

func InitGlobalTest(configName string) {
	globalVars[configName] = fmt.Sprintf("%v", true)
}

func ReplaceVars(value string) string {
	if strings.Contains(value, "${") {
		for k := range globalVars {
			value = strings.ReplaceAll(value, fmt.Sprintf("${%s}", k), GetValueOfVar(k))
		}
	}
	return value
}

func AppendVar(varName string, value string) {
	globalVars[varName] = globalVars[varName] + value
}

func removeSliceItems(slice []string, n int) []string {
	i := 0
	for i < n {
		slice = append(slice[:0], slice[1:]...)
		i += 1
	}
	return slice
}

func findMapVar(m map[string]interface{}, ks ...string) (rval interface{}, err error) {
	var ok bool

	if len(ks) == 0 {
		return nil, fmt.Errorf("%s needs at least one key", m)
	}
	if rval, ok = m[ks[0]]; !ok {
		return nil, fmt.Errorf("key not found; remaining keys: %v", ks)
	} else if len(ks) == 1 { // we've reached the final key
		return rval, nil
	} else if m, ok = rval.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("malformed structure at %#v", rval)
	} else { // 1+ more keys
		return findMapVar(m, ks[1:]...)
	}
}
