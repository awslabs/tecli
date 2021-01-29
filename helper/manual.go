package helper

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/box"
	yaml "gopkg.in/yaml.v2"
)

// Manual mapping the fields used by a Cobra command
type Manual struct {
	Use     string `yaml:"use,omitempty"`
	Example string `yaml:"example,omitempty"`
	Short   string `yaml:"short,omitempty"`
	Long    string `yaml:"long,omitempty"`
}

//GetManual returns the command's manuall
func GetManual(command string, args []string) (Manual, error) {
	var man Manual
	var err error
	manualBlob, status := box.Get("/manual/" + command + ".yaml")
	if status {
		err = yaml.Unmarshal(manualBlob, &man)
		if err != nil {
			return man, fmt.Errorf("unable to decode YAML file, error:\n%v", err)
		}
	} else {
		logrus.Fatal("unable to read manual from box")
	}

	if man.Use != "" {
		man.Use = strings.ReplaceAll(man.Use, "{{ arguments }}", "["+strings.Join(args, "|")+"]")
	}

	return man, err
}
