package valueobjects_test

import (
	"testing"

	"github.com/pavanrkadave/tdd-security-analysis/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestImpact_AllValues_HaveCorrectString(t *testing.T) {
	tests := []struct {
		name   string
		impact valueobjects.Impact
		want   string
	}{
		{"Check Marginal Exists", valueobjects.Marginal, "MARGINAL"},
		{"Check Moderate Exists", valueobjects.Moderate, "MODERATE"},
		{"Check Critical Exists", valueobjects.Critical, "CRITICAL"},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			assert.Equal(t, testItem.want, testItem.impact.String())
		})
	}
}
