package file_reader

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	TestName string `yaml:"name"`
	Steps    []Step `yaml:"steps"`
}

type Step struct {
	StepName  string            `yaml:"name"`
	Condition string            `yaml:"condition,omitempty"`
	Path      string            `yaml:"path"`
	Method    string            `yaml:"method"`
	Body      string            `yaml:body,omitempty`
	Headers   []Headers         `yaml:headers,omitempty`
	Vars      map[string]string `yaml:"vars,omitempty"`
	Tests     []Tests           `yaml:"tests,omitempty"`
}

type Headers struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Tests struct {
	Name     string `yaml:"name"`
	Expected string `yaml:"expected"`
	Actual   string `yaml:"actual"`
	Type     string `yaml:"type"`
}

func (c *Config) ReadFile() *Config {
	filePath := flag.Lookup("file").Value.String()
	ymlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("Error when trying to read yaml file. %s", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(ymlFile, &c)
	if err != nil {
		log.Error("Error when trying to unmarshal yaml file.")
	}

	return c
}
