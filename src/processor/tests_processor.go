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
				var actual = GetValueOfVar(t.Actual)
				testResults = append(testResults, TestResult{
					Name:     t.Name,
					Type:     t.Type,
					Expected: t.Expected,
					Actual:   actual,
					Result:   verifyEquals(t.Expected, actual),
				})

			}
		case "CONTAINS":
			{
				var actual = GetValueOfVar(t.Actual)
				testResults = append(testResults, TestResult{
					Name:     t.Name,
					Type:     t.Type,
					Expected: t.Expected,
					Actual:   actual,
					Result:   verifyContains(t.Expected, actual),
				})
			}
		case "NOT_NIL":
			{
				var actual = GetValueOfVar(t.Actual)
				testResults = append(testResults, TestResult{
					Name:     t.Name,
					Type:     t.Type,
					Expected: "NOT_NIL",
					Actual:   actual,
					Result:   verifyNotNil(actual),
				})
			}
		case "NIL":
			{
				var actual = GetValueOfVar(t.Actual)
				testResults = append(testResults, TestResult{
					Name:     t.Name,
					Type:     t.Type,
					Expected: "NIL",
					Actual:   actual,
					Result:   verifyNil(actual),
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
	return actual != nil && actual != "<nil>"
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
