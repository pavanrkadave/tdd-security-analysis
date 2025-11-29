package valueobjects_test

import (
	"testing"

	"github.com/pavanrkadave/tdd-security-analysis/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestMeasure_Create(t *testing.T) {
	//Arrange
	measureDescription := "Configure Firewall"

	//Act
	measure := valueobjects.NewMeasure(measureDescription)

	//Assert
	assert.NotNil(t, measure)
	assert.Equal(t, measureDescription, measure.Description())
}
