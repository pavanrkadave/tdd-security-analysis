package valueobjects_test

import (
	"testing"

	"github.com/pavanrkadave/tdd-security-analysis/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestLikelihood_Rare_Exists(t *testing.T) {
	likelihood := valueobjects.Rare

	assert.Equal(t, "RARE", likelihood.String())
}
