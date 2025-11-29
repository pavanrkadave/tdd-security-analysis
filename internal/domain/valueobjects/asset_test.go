package valueobjects_test

import (
	"testing"

	"github.com/pavanrkadave/tdd-security-analysis/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestAsset_Creation(t *testing.T) {
	//Arrange
	assetName := "Database"

	//Action
	asset := valueobjects.NewAsset(assetName)

	//Assert
	assert.NotNil(t, asset)
	assert.Equal(t, assetName, asset.Name())
}
