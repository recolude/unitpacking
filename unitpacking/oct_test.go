package unitpacking_test

import (
	"fmt"
	"testing"

	"github.com/EliCDavis/vector"
	"github.com/recolude/unitpacking/unitpacking"
	"github.com/stretchr/testify/assert"
)

func TestOctPack32(t *testing.T) {
	tests := []vector.Vector3{
		vector.NewVector3(0, 1, 2),
		vector.NewVector3(1, 0, 0),
		vector.NewVector3(0, 1, 0),
		vector.NewVector3(0, 0, 1),
		vector.NewVector3(-1, 0, 0),
		vector.NewVector3(0, -1, 0),
		vector.NewVector3(0, 0, -1),
		vector.NewVector3(-1, -1, -1),
		vector.NewVector3(-1, 1, -1),
		vector.NewVector3(-0.997605826445425, 0.06365823804882093, -0.027023022974122023),
	}

	for _, tc := range tests {
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

func TestOctPack24(t *testing.T) {
	tests := []vector.Vector3{
		vector.NewVector3(0, 1, 2),
		vector.NewVector3(1, 0, 0),
		vector.NewVector3(0, 1, 0),
		vector.NewVector3(0, 0, 1),
		vector.NewVector3(-1, 0, 0),
		vector.NewVector3(0, -1, 0),
		vector.NewVector3(0, 0, -1),
		vector.NewVector3(-1, -1, -1),
		vector.NewVector3(-1, 1, -1),
		vector.NewVector3(-0.997605826445425, 0.06365823804882093, -0.027023022974122023),
	}

	for _, tc := range tests {
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
