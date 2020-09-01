package main

import (
	"encoding/json"
	"github.com/carloshjoaquim/E2Easy-Go/file_reader"
	"github.com/carloshjoaquim/E2Easy-Go/processor"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infoln("Hello E2E ! ... ")
	ConfigureFlags()

	var c file_reader.Config
	c.ReadFile()

	for _, s := range c.Steps {
		resultStep := processor.RunStep(s)
		pretty, _ := json.MarshalIndent(resultStep, "", "    ")
		log.Infof("Response: \n %+v", string(pretty))

		processor.GetVarsFromResponse(s.Vars, resultStep)

		log.Infoln("Running Tests ... ")
		testsResult := processor.ProcessTests(s.Tests)
		for _, tr := range testsResult {
			log.Infof("\nName: %v \n" +
				"Type: %v \n" +
				"Expected: %v \n" +
				"Actual: %v \n" +
				"Result: %v \n", tr.Name, tr.Type, tr.Expected, tr.Actual, tr.Result)
		}
		log.Infoln("End ofTests ... ")
	}
}
