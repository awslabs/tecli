package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFFEClient(t *testing.T) {
	client := GetTFEClient()
	assert.NotNil(t, client)
}
