package unitpacking_test

import (
	"testing"

	"github.com/recolude/unitpacking/unitpacking"
	"github.com/stretchr/testify/assert"
)

func TestClamp(t *testing.T) {
	tests := map[string]struct {
		min   float64
		max   float64
		input float64
		want  float64
	}{
		"at min boundary":       {input: 0, min: 0, max: 1, want: 0},
		"min boundary exceeded": {input: -1, min: 0, max: 1, want: 0},
		"at max boundary":       {input: 1, min: 0, max: 1, want: 1},
		"max boundary exceeded": {input: 2, min: 0, max: 1, want: 1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := unitpacking.Clamp(tc.input, tc.min, tc.max)
			assert.Equal(t, tc.want, got)
		})
	}
}
