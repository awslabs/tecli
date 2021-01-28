package aid

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// ToJSON converts a given struct to json
func ToJSON(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.Fatalf("unable to convert struct to json\n%v\n", b)
	}

	return string(b)
}
