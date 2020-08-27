package main

import (
	"encoding/json"
	"github.com/carloshjoaquim/E2Easy-Go/file_reader"
	"github.com/carloshjoaquim/E2Easy-Go/rest"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infoln("Hello E2E ! ... ")
	ConfigureFlags()

	var c file_reader.Config
	c.ReadFile()

	log.Infoln(c)
	log.Println("Trying to CALL RESTY...")

	response, err := rest.Get(c.Steps[0].Path)
	if err != nil {
		log.Error("Error whit GET")
	}

	var body interface{}
	err = json.Unmarshal(response.Body, &body)

	log.Infof("Response: %s", body)
}
