package valueobjects_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreat_Creation(t *testing.T) {
	//Arrange
	threatName := "Power Outage"

	//Action
	threat := valueobjects.NewThreat(threatName)

	//Assertion
	assert.NotNil(t, threat)
	assert.Equal(t, threatName, threat.Name())
}
