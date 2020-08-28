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

		vars := processor.GetVarsFromResponse(s.Vars, resultStep)
		log.Infof("Vars: \n %+v", vars)
	}
}
