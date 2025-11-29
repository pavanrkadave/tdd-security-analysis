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

func TestLikelihood_Occasional_Exists(t *testing.T) {
	likelihood := valueobjects.Occasional

	assert.Equal(t, "OCCASIONAL", likelihood.String())
}

func TestLikelihood_Frequent_Exists(t *testing.T) {
	likelihood := valueobjects.Frequent

	assert.Equal(t, "FREQUENT", likelihood.String())
}
