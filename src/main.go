package runner

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

	for _, s := range c.Steps {
		if processor.SatisfiesCondition(&s) {
			resultStep := processor.RunStep(s)
			processor.GetVarsFromResponse(s.Vars, resultStep)
			log.Infoln("Running Tests ... ")
			testsResult := processor.ProcessTests(s.Tests)
			processor.AddTestVar(fmt.Sprintf("%s.tests", s.StepName), testsResult, c.TestName)

			for _, tr := range testsResult {
				log.Infof("\nName: %v \n"+
					"Type: %v \n"+
					"Expected: %v \n"+
					"Actual: %v \n"+
					"Result: %v \n", tr.Name, tr.Type, tr.Expected, tr.Actual, tr.Result)
			}
		} else {
			continue
		}
	}
}
