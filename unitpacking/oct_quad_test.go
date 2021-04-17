package unitpacking_test

import (
	"fmt"
	"testing"

	"github.com/recolude/unitpacking/unitpacking"
	"github.com/stretchr/testify/assert"
)

func TestOctQuad16(t *testing.T) {
	for _, tc := range testVectors {
		unit := tc.Normalized()
		name := fmt.Sprintf("%.2f,%.2f,%.2f", unit.X(), unit.Y(), unit.Z())

		t.Run(name, func(t *testing.T) {
			packed := unitpacking.PackOctQuad16(unit)
			assert.Len(t, packed, 2)
			unpacked := unitpacking.UnpackOctQuad16(packed)

			assert.InDelta(t, unit.X(), unpacked.X(), 0.02, "X components not equal: %.2f != %.2f", unit.X(), unpacked.X())
			assert.InDelta(t, unit.Y(), unpacked.Y(), 0.02, "Y components not equal: %.2f != %.2f", unit.Y(), unpacked.Y())
			assert.InDelta(t, unit.Z(), unpacked.Z(), 0.02, "Z components not equal: %.2f != %.2f", unit.Z(), unpacked.Z())
		})
	}
}

func TestOctQuad24(t *testing.T) {
	for _, tc := range testVectors {
		unit := tc.Normalized()
		name := fmt.Sprintf("%.2f,%.2f,%.2f", unit.X(), unit.Y(), unit.Z())

		t.Run(name, func(t *testing.T) {
			packed := unitpacking.PackOctQuad24(unit)
			assert.Len(t, packed, 3)
			unpacked := unitpacking.UnpackOctQuad24(packed)

			assert.InDelta(t, unit.X(), unpacked.X(), 0.02, "X components not equal: %.2f != %.2f", unit.X(), unpacked.X())
			assert.InDelta(t, unit.Y(), unpacked.Y(), 0.02, "Y components not equal: %.2f != %.2f", unit.Y(), unpacked.Y())
			assert.InDelta(t, unit.Z(), unpacked.Z(), 0.02, "Z components not equal: %.2f != %.2f", unit.Z(), unpacked.Z())
		})
	}
}

func TestOctQuad32(t *testing.T) {
	for _, tc := range testVectors {
		unit := tc.Normalized()
		name := fmt.Sprintf("%.2f,%.2f,%.2f", unit.X(), unit.Y(), unit.Z())

		t.Run(name, func(t *testing.T) {
			packed := unitpacking.PackOctQuad32(unit)
			assert.Len(t, packed, 4)
			unpacked := unitpacking.UnpackOctQuad32(packed)

			assert.InDelta(t, unit.X(), unpacked.X(), 0.00005, "X components not equal: %.2f != %.2f", unit.X(), unpacked.X())
			assert.InDelta(t, unit.Y(), unpacked.Y(), 0.00005, "Y components not equal: %.2f != %.2f", unit.Y(), unpacked.Y())
			assert.InDelta(t, unit.Z(), unpacked.Z(), 0.00005, "Z components not equal: %.2f != %.2f", unit.Z(), unpacked.Z())
		})
	}
}
