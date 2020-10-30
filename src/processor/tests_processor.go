package processor

import (
	"github.com/carloshjoaquim/E2Easy-Go/src/file_reader"
	"strconv"
	"strings"
)

type TestInterface interface {
	verifyEquals(interface{}, interface{}) bool
	verifyContains(string, string) bool
	verifyNil(interface{}) bool
	verifyNotNil(interface{}) bool
}

type TestResult struct {
	Name     string
	Type     string
	Expected string
	Actual   string
	Result   bool
}

func ProcessTests(tests []file_reader.Tests) []TestResult {
	testResults := make([]TestResult, 0)

	for _, t := range tests {
		switch strings.ToUpper(t.Type) {
		case "EQUALS":
			{
				testResults = append(testResults, TestResult{
					Name:     t.Name,
					Type:     t.Type,
					Expected: t.Expected,
					Actual:   GetValueOfVar(t.Actual),
					Result:   verifyEquals(t.Expected, GetValueOfVar(t.Actual)),
				})

			}
		case "CONTAINS":
			{
				testResults = append(testResults, TestResult{
					Name:     t.Name,
					Type:     t.Type,
					Expected: t.Expected,
					Actual:   GetValueOfVar(t.Actual),
					Result:   verifyContains(t.Expected, GetValueOfVar(t.Actual)),
				})
			}
		case "NOT_NIL":
			{
				testResults = append(testResults, TestResult{
					Name:     t.Name,
					Type:     t.Type,
					Expected: "NOT_NIL",
					Actual:   GetValueOfVar(t.Actual),
					Result:   verifyNotNil(GetValueOfVar(t.Actual)),
				})
			}
		case "NIL":
			{
				testResults = append(testResults, TestResult{
					Name:     t.Name,
					Type:     t.Type,
					Expected: "NIL",
					Actual:   GetValueOfVar(t.Actual),
					Result:   verifyNil(GetValueOfVar(t.Actual)),
				})
			}
		}
	}

	return testResults
}

func verifyEquals(expected interface{}, actual interface{}) bool {
	return expected == actual
}

func verifyContains(expected string, actual string) bool {
	return strings.Contains(actual, expected)
}

func verifyNil(actual interface{}) bool {
	return actual == nil
}

func verifyNotNil(actual interface{}) bool {
	return actual != nil
}

func SatisfiesCondition(s *file_reader.Step) bool {
	if s.Condition == "" {
		return true
	}

	if strings.Contains(s.Condition, "not") {
		replaced := strings.ReplaceAll(s.Condition, "not", "")
		replaced = strings.TrimSpace(ReplaceVars(replaced))
		boolValue, _ := strconv.ParseBool(replaced)
		return !boolValue
	} else {
		replaced := strings.TrimSpace(ReplaceVars(s.Condition))
		boolValue, _ := strconv.ParseBool(replaced)
		return boolValue
	}
}
