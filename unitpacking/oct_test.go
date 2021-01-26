package unitpacking_test

import (
	"fmt"
	"testing"

	"github.com/recolude/unitpacking/unitpacking"
	"github.com/stretchr/testify/assert"
)

func TestOct32(t *testing.T) {
	for _, tc := range testVectors {
		unit := tc.Normalized()
		name := fmt.Sprintf("%.2f,%.2f,%.2f", unit.X(), unit.Y(), unit.Z())

		t.Run(name, func(t *testing.T) {
			packed := unitpacking.PackOct32(unit)
			assert.Len(t, packed, 4)
			unpacked := unitpacking.UnpackOct32(packed)

			assert.InDelta(t, unit.X(), unpacked.X(), 0.01, "X components not equal: %.2f != %.2f", unit.X(), unpacked.X())
			assert.InDelta(t, unit.Y(), unpacked.Y(), 0.01, "Y components not equal: %.2f != %.2f", unit.Y(), unpacked.Y())
			assert.InDelta(t, unit.Z(), unpacked.Z(), 0.01, "Z components not equal: %.2f != %.2f", unit.Z(), unpacked.Z())
		})
	}
}

func TestOct24(t *testing.T) {
	for _, tc := range testVectors {
		unit := tc.Normalized()
		name := fmt.Sprintf("%.2f,%.2f,%.2f", unit.X(), unit.Y(), unit.Z())

		t.Run(name, func(t *testing.T) {
			packed := unitpacking.PackOct24(unit)
			assert.Len(t, packed, 3)
			unpacked := unitpacking.UnpackOct24(packed)

			assert.InDelta(t, unit.X(), unpacked.X(), 0.01, "X components not equal: %.2f != %.2f", unit.X(), unpacked.X())
			assert.InDelta(t, unit.Y(), unpacked.Y(), 0.01, "Y components not equal: %.2f != %.2f", unit.Y(), unpacked.Y())
			assert.InDelta(t, unit.Z(), unpacked.Z(), 0.01, "Z components not equal: %.2f != %.2f", unit.Z(), unpacked.Z())
		})
	}
}

func TestOct16(t *testing.T) {
	for _, tc := range testVectors {
		unit := tc.Normalized()
		name := fmt.Sprintf("%.2f,%.2f,%.2f", unit.X(), unit.Y(), unit.Z())

		t.Run(name, func(t *testing.T) {
			packed := unitpacking.PackOct16(unit)
			assert.Len(t, packed, 2)
			unpacked := unitpacking.UnpackOct16(packed)

			assert.InDelta(t, unit.X(), unpacked.X(), 0.02, "X components not equal: %.2f != %.2f", unit.X(), unpacked.X())
			assert.InDelta(t, unit.Y(), unpacked.Y(), 0.02, "Y components not equal: %.2f != %.2f", unit.Y(), unpacked.Y())
			assert.InDelta(t, unit.Z(), unpacked.Z(), 0.02, "Z components not equal: %.2f != %.2f", unit.Z(), unpacked.Z())
		})
	}
}
