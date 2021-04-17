package unitpacking_test

import (
	"fmt"
	"testing"

	"github.com/EliCDavis/vector"
	"github.com/recolude/unitpacking/unitpacking"
	"github.com/stretchr/testify/assert"
)

func TestQuadRecurse_SimpleTopRight(t *testing.T) {
	output := unitpacking.QuadRecurse(
		vector.NewVector2(0.5, 0.5),
		vector.NewVector2(-1, -1),
		vector.NewVector2(1, 1),
		1,
	)

	assert.Len(t, output, 1)
	assert.Equal(t, unitpacking.TopRight, output[0])
}

func TestQuadRecurse_SimpleBottomRight(t *testing.T) {
	output := unitpacking.QuadRecurse(
		vector.NewVector2(0.5, -0.5),
		vector.NewVector2(-1, -1),
		vector.NewVector2(1, 1),
		1,
	)

	assert.Len(t, output, 1)
	assert.Equal(t, unitpacking.BottomRight, output[0])
}

func TestQuadRecurse_SimpleBottomLeft(t *testing.T) {
	output := unitpacking.QuadRecurse(
		vector.NewVector2(-0.5, -0.5),
		vector.NewVector2(-1, -1),
		vector.NewVector2(1, 1),
		1,
	)

	assert.Len(t, output, 1)
	assert.Equal(t, unitpacking.BottomLeft, output[0])
}

func TestQuadRecurse_SimpleTopLeft(t *testing.T) {
	output := unitpacking.QuadRecurse(
		vector.NewVector2(-0.5, 0.5),
		vector.NewVector2(-1, -1),
		vector.NewVector2(1, 1),
		1,
	)

	assert.Len(t, output, 1)
	assert.Equal(t, unitpacking.TopLeft, output[0])
}

func TestQuadRecurse_MultipleLevels(t *testing.T) {
	output := unitpacking.QuadRecurse(
		vector.NewVector2(0.25, 0.25),
		vector.NewVector2(-1, -1),
		vector.NewVector2(1, 1),
		2,
	)

	assert.Len(t, output, 2)
	assert.Equal(t, unitpacking.TopRight, output[1])
	assert.Equal(t, unitpacking.BottomLeft, output[0])
}

func TestQuad_ByteAndBack(t *testing.T) {
	for _, tc := range quadTestVectors {
		name := fmt.Sprintf("%.2f,%.2f", tc.X(), tc.Y())

		t.Run(name, func(t *testing.T) {
			packed := unitpacking.Vec2ToByteQuad(tc)
			unpacked := unitpacking.ByteQuadToVec2(packed)

			assert.InDelta(t, tc.X(), unpacked.X(), 0.07, "X components not equal: %.2f != %.2f", tc.X(), unpacked.X())
			assert.InDelta(t, tc.Y(), unpacked.Y(), 0.07, "Y components not equal: %.2f != %.2f", tc.Y(), unpacked.Y())
		})
	}
}

func TestQuad_2BytesAndBack(t *testing.T) {
	for _, tc := range quadTestVectors {
		name := fmt.Sprintf("%.2f,%.2f", tc.X(), tc.Y())

		t.Run(name, func(t *testing.T) {
			packed := unitpacking.Vec2ToTwoByteQuad(tc)
			assert.Len(t, packed, 2)
			unpacked := unitpacking.TwoByteQuadToVec2(packed)

			assert.InDelta(t, tc.X(), unpacked.X(), 0.004, "X components not equal: %.2f != %.2f", tc.X(), unpacked.X())
			assert.InDelta(t, tc.Y(), unpacked.Y(), 0.004, "Y components not equal: %.2f != %.2f", tc.Y(), unpacked.Y())
		})
	}
}
