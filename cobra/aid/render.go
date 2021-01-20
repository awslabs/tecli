package aid

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/awslabs/tfe-cli/cobra/model"
	"github.com/awslabs/tfe-cli/helper"
	gomplateV3 "github.com/hairyhenderson/gomplate/v3"
	"github.com/sirupsen/logrus"
)

// BuildTemplate build the given template located under tfe-cli/ directory (without the .tmpl extension)
func BuildTemplate(name string) error {
	var inputFiles = []string{}
	var outputFiles = []string{}

	if helper.FileExists("tfe-cli/" + name + ".tmpl") {
		inputFiles = append(inputFiles, "tfe-cli/"+name+".tmpl")
		outputFiles = append(outputFiles, strings.ToUpper(name)+".md")
	}

	var config gomplateV3.Config
	config.InputFiles = inputFiles
	config.OutputFiles = outputFiles

	dataSources := []string{}
	if helper.FileExists("tfe-cli/" + name + ".yaml") {
		dataSources = append(dataSources, "db=./tfe-cli/"+name+".yaml")
	}

	config.DataSources = dataSources

	err := gomplateV3.RunTemplates(&config)
	if err != nil {
		logrus.Fatalf("Gomplate.RunTemplates() failed with %s\n", err)
	}

	return err
}

func writeInputs() error {
	variables, err := os.Open("variables.tf")
	if err != nil {
		logrus.Fatal(err)
	}
	defer variables.Close()

	// create INPUTS.md
	inputs, err := os.OpenFile("INPUTS.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Println(err)
	}
	defer inputs.Close()

	if _, err := inputs.WriteString("| Name | Description | Type | Default | Required |\n|------|-------------|:----:|:-----:|:-----:|\n"); err != nil {
		logrus.Println(err)
	}

	var varName, varType, varDescription, varDefault string
	varRequired := "no"

	// startBlock := false
	scanner := bufio.NewScanner(variables)
	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if len(line) > 0 {
			if strings.Contains(line, "variable") && strings.Contains(line, "{") {
				out, found := helper.GetStringBetweenDoubleQuotes(line)
				if found {
					varName = out
				}

			}

			if strings.Contains(line, "type") && strings.Contains(line, "=") {
				slc := helper.GetStringTrimmed(line, "=")
				if slc[0] == "type" {
					varType = slc[1]
					if strings.Contains(varType, "({") {
						slc = helper.GetStringTrimmed(varType, "({")
						varType = slc[0]
					}
				}
			}

			if strings.Contains(line, "description") && strings.Contains(line, "=") {
				slc := helper.GetStringTrimmed(line, "=")
				if slc[0] == "description" {
					out, found := helper.GetStringBetweenDoubleQuotes(slc[1])
					if found {
						varDescription = out
					}
				}
			}

			if strings.Contains(line, "default") && strings.Contains(line, "=") {
				slc := helper.GetStringTrimmed(line, "=")
				if slc[0] == "default" {
					varDefault = slc[1]
					if strings.Contains(varDefault, "{") {
						varDefault = "<map>"
					}
				}
			}

			// end of the variable declaration
			if strings.Contains(line, "}") && len(line) == 1 {
				if len(varName) > 0 && len(varType) > 0 && len(varDescription) > 0 {

					var result string
					if len(varDefault) == 0 {
						varRequired = "yes"
						result = fmt.Sprintf("| %s | %s | %s | %s | %s |\n", varName, varDescription, varType, varDefault, varRequired)
					} else {
						result = fmt.Sprintf("| %s | %s | %s | `%s` | %s |\n", varName, varDescription, varType, varDefault, varRequired)
					}

					if _, err := inputs.WriteString(result); err != nil {
						logrus.Println(err)
					}
					varName, varType, varDescription, varDefault, varRequired = "", "", "", "", "no"
				}
			}

		}

	}

	if err := scanner.Err(); err != nil {
		logrus.Fatal(err)

	}
	return err
}

func writeOutputs() error {
	outputs, err := os.Open("outputs.tf")
	if err != nil {
		logrus.Fatal(err)
	}
	defer outputs.Close()

	// create INPUTS.md
	outs, err := os.OpenFile("OUTPUTS.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Println(err)
	}
	defer outs.Close()

	if _, err := outs.WriteString("| Name | Description |\n|------|-------------|\n"); err != nil {
		logrus.Println(err)
	}

	var outName, outDescription string

	scanner := bufio.NewScanner(outputs)
	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if len(line) > 0 {
			if strings.Contains(line, "output") && strings.Contains(line, "{") {
				out, found := helper.GetStringBetweenDoubleQuotes(line)
				if found {
					outName = out
				}
			}

			if strings.Contains(line, "description") && strings.Contains(line, "=") {
				slc := helper.GetStringTrimmed(line, "=")
				if slc[0] == "description" {
					out, found := helper.GetStringBetweenDoubleQuotes(slc[1])
					if found {
						outDescription = out
					}
				}
			}

			// end of the output declaration
			if strings.Contains(line, "}") && len(line) == 1 {
				if len(outName) > 0 && len(outDescription) > 0 {

					result := fmt.Sprintf("| %s | %s | |\n", outName, outDescription)

					if _, err := outs.WriteString(result); err != nil {
						logrus.Println(err)
					}
					outName, outDescription = "", ""
				}
			}

		}

	}

	if err := scanner.Err(); err != nil {
		logrus.Fatal(err)

	}
	return err
}

// UpdateReadMeLogoURL TODO ...
func UpdateReadMeLogoURL(readme model.ReadMe, response model.UnsplashRandomPhotoResponse) error {
	readme.Logo.URL = response.Urls.Regular
	err := WriteInterfaceToFile(readme, helper.BuildPath("tfe-cli/readme.yaml"))
	if err != nil {
		return fmt.Errorf("unable to save new readme template\n%v", err)
	}

	return nil
}
