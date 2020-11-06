package processor

import (
	"fmt"
	"github.com/carloshjoaquim/E2Easy-Go/src/file_reader"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fastjson"
	"strconv"
	"strings"
	"time"
)

var (
	globalVars = make(map[string]string, 0)
)

func GetVarsFromResponse(expectedVars map[string]string, result StepResult) map[string]string {
	if len(expectedVars) > 0 {
		for k, v := range expectedVars {
			evaluated := getValueFromResult(v, result)
			if evaluated != "" {
				globalVars[k] = evaluated
			}
		}
	}

	return globalVars
}

func GetValueOfVar(varName string) string {
	cleanedVar := strings.ReplaceAll(strings.ReplaceAll(varName, "${", ""), "}", "")

	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(globalVars[cleanedVar], "{", ""), "}", ""), string('"'), "")
}

func getValueFromResult(varName string, result StepResult) string {
	varName = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(varName, "[", "."), "]", "."), "..", ".")
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
						var p fastjson.Parser
						bodyJson, err := p.Parse(result.Message)
						if err != nil {
							log.Error(err)
						}
						result := bodyJson.Get(removeSliceItems(values, 2)...)
						value = fmt.Sprintf("%v", result)
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
	testPassed := true
	for _, t := range testResult {
		testPassed = testPassed && t.Result
	}

	globalTests, _ := strconv.ParseBool(globalVars[configName])
	globalVars[configName] = fmt.Sprintf("%v", globalTests && testPassed)
	globalVars[varName] = fmt.Sprintf("%v", testPassed)
}

func InitGlobalVars(c file_reader.Config) {
	for _, s := range c.Steps {
		for n, v := range s.Vars {
			AddVar(strings.TrimSpace(n), strings.TrimSpace(v))

			if v == "${UUID()}" {
				globalVars[n] = fmt.Sprintf("%v", uuid.New())
			}
			if strings.Contains(v, "Instant.now()") {
				globalVars[n] = calculateInstantValue(v)
			}
		}
	}

	initGlobalTest(c.TestName)
}

func calculateInstantValue(command string) string {
	spt := strings.Split(command, ".")
	result := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	if len(spt) > 2 {
		if strings.Contains(spt[2], "plus") {
			calculatedTime := plusValueToTime(time.Now().UTC(), getTimeAmount(spt[2]), getTimeUnit(command))
			result = calculatedTime.Format("2006-01-02T15:04:05.000Z")
		}
		if strings.Contains(spt[2], "minus") {
			calculatedTime := plusValueToTime(time.Now().UTC(), -1*getTimeAmount(spt[2]), getTimeUnit(command))
			result = calculatedTime.Format("2006-01-02T15:04:05.000Z")
		}
	}

	return result
}

func getTimeAmount(value string) int64 {
	value = strings.ReplaceAll(value, "plus(", ",")
	value = strings.ReplaceAll(value, "minus(", ",")
	valueToAdd, _ := strconv.ParseInt(strings.Split(value,",")[1], 10, 64)
	return valueToAdd
}

func getTimeUnit(value string) string {
	if strings.Contains(value, "HOURS") {
		return "HOURS"
	}
	if strings.Contains(value, "DAYS") {
		return "DAYS"
	}
	return ""
}

func plusValueToTime(oldTime time.Time, value int64, temp string) time.Time {
	var newDate = oldTime
	switch temp {
	case "HOURS":
		{
			newDate = oldTime.Add(time.Hour * time.Duration(value))
		}
	case "DAYS":
		{
			newDate = oldTime.AddDate(0,0, int(value))
		}
	}
	return newDate
}

func initGlobalTest(configName string) {
	globalVars[fmt.Sprintf("%v.allTestsPassed", configName)] = fmt.Sprintf("%v", true)
}

func ReplaceVars(value string) string {
	if strings.Contains(value, "${") {
		for k := range globalVars {
			value = strings.ReplaceAll(value, fmt.Sprintf("${%s}", k), GetValueOfVar(k))
		}
	}
	return value
}

func AddVar(varName string, value string) {
	globalVars[varName] = value
}

func AppendVar(varName string, value string) {
	globalVars[varName] = fmt.Sprintf("%v %v", globalVars[varName], value)
}

func removeSliceItems(slice []string, n int) []string {
	i := 0
	for i < n {
		slice = append(slice[:0], slice[1:]...)
		i += 1
	}
	return slice
}
