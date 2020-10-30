package e2easy

import (
	"fmt"
	"github.com/carloshjoaquim/E2Easy-Go/src/file_reader"
	"github.com/carloshjoaquim/E2Easy-Go/src/processor"
	log "github.com/sirupsen/logrus"
)

func RunE2E() {
	log.Infoln("Hello E2E ! ... ")
	ConfigureFlags()

	var c file_reader.Config
	c.ReadFile()
	processor.InitGlobalVars(c)
	for _, s := range c.Steps {
		if processor.SatisfiesCondition(&s) {
			resultStep := processor.RunStep(s)
			processor.GetVarsFromResponse(s.Vars, resultStep)
			log.Infoln("Running Tests ... ")
			testsResult := processor.ProcessTests(s.Tests)
			processor.AddTestVar(fmt.Sprintf("%s.allTestsPassed", s.StepName), testsResult, c.TestName)

			for _, tr := range testsResult {
				testLog := fmt.Sprintf("\nName: %v \n"+
					"Type: %v \n"+
					"Expected: %v \n"+
					"Actual: %v \n"+
					"Result: %v \n", tr.Name, tr.Type, tr.Expected, tr.Actual, tr.Result)

				processor.AppendVar(fmt.Sprintf("%s.tests", c.TestName), testLog)
				if tr.Result {
					processor.AppendVar(fmt.Sprintf("%s.tests.success", c.TestName), testLog)
				} else {
					processor.AppendVar(fmt.Sprintf("%s.tests.failed", c.TestName), testLog)
				}
				log.Info(testLog)
			}
		} else {
			continue
		}
	}
}
