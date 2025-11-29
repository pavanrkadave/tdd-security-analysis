package valueobjects_test

import (
	"testing"

	"github.com/pavanrkadave/tdd-security-analysis/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestLikelihood_AllValues_HaveCorrectString(t *testing.T) {
	tests := []struct {
		name       string
		likelihood valueobjects.Likelihood
		want       string
	}{
		{"Check Rare Exists", valueobjects.Rare, "RARE"},
		{"Check Occasional Exists", valueobjects.Occasional, "OCCASIONAL"},
		{"Check Frequent Exists", valueobjects.Frequent, "FREQUENT"},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			assert.Equal(t, testItem.want, testItem.likelihood.String())
		})
	}
}
