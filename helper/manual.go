package helper

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/box"
	yaml "gopkg.in/yaml.v2"
)

// Manual mapping the fields used by a Cobra command
type Manual struct {
	Use     string `yaml:"use"`
	Example string `yaml:"example"`
	Short   string `yaml:"short"`
	Long    string `yaml:"long"`
}

// GetManual retrieve information about the given command
func GetManual(command string) (Manual, error) {
	var man Manual
	var err error
	manualBlob, status := box.Get("/manual/" + command + ".yaml")
	if status {
		err = yaml.Unmarshal(manualBlob, &man)
		if err != nil {
			return man, fmt.Errorf("unable to decode YAML file, error:\n%v\n", err)
		}
	} else {
		logrus.Fatal("unable to read manual from box")
	}

	return man, err
}
