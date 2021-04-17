package unitpacking

import (
	"math"

	"github.com/EliCDavis/vector"
)

// Returns +-1
func signNotZero(v vector.Vector2) vector.Vector2 {
	x := 1.0
	if v.X() < 0.0 {
		x = -1.0
	}

	y := 1.0
	if v.Y() < 0.0 {
		y = -1.0
	}

	return vector.NewVector2(x, y)
}

func multVect(a, b vector.Vector2) vector.Vector2 {
	return vector.NewVector2(
		a.X()*b.X(),
		a.Y()*b.Y(),
	)
}

// PackOct32 maps a unit vector to a 2D UV of a octahedron, and then writes the
// 2D coordinates to 4 bytes, 2 bytes per coordinate.
func PackOct32(v vector.Vector3) []byte {
	uvCords := MapToOctUVPrecise(v, 32)

	// 2 ^ 16 = 65,536;
	x := uint(math.Floor(uvCords.X()*32767) + 32768)
	y := uint(math.Floor(uvCords.Y()*32767) + 32768)
	everything := (x << 16) | y

	return []byte{
		(byte)(everything & 0xFF),
		(byte)((everything >> 8) & 0xFF),
		(byte)((everything >> 16) & 0xFF),
		(byte)((everything >> 24) & 0xFF),
	}
}

// UnpackOct32 reads in two 16bit numbers and converts from 2D octahedron UV to
// 3D unit sphere coordinates.
func UnpackOct32(b []byte) vector.Vector3 {
	everything := uint(b[0]) | (uint(b[1]) << 8) | (uint(b[2]) << 16) | (uint(b[3]) << 24)
	rawY := (int)((everything) & 0xFFFF)
	rawX := (int)(everything >> 16)

	cleanedX := Clamp((float64(rawX)-32768.0)/32767.0, -1.0, 1.0)
	cleanedY := Clamp((float64(rawY)-32768.0)/32767.0, -1.0, 1.0)

	return FromOctUV(vector.NewVector2(cleanedX, cleanedY))
}

// PackOct24 maps a unit vector to a 2D UV of a octahedron, and then writes the
// 2D coordinates to 3 bytes, 12bits per coordinate.
func PackOct24(v vector.Vector3) []byte {
	uvCords := MapToOctUVPrecise(v, 24)

	// 2 ^ 12 = 4,096;
	x := uint(math.Floor(uvCords.X()*2047) + 2048)
	y := uint(math.Floor(uvCords.Y()*2047) + 2048)
	everything := (x << 12) | y

	return []byte{
		(byte)(everything & 0xFF),
		(byte)((everything >> 8) & 0xFF),
		(byte)((everything >> 16) & 0xFF),
	}
}

// UnpackOct24 reads in two 12bit numbers and converts from 2D octahedron UV to
// 3D unit sphere coordinates.
func UnpackOct24(b []byte) vector.Vector3 {
	everything := uint(b[0]) | (uint(b[1]) << 8) | (uint(b[2]) << 16)
	rawY := (int)((everything) & 0b111111111111)
	rawX := (int)(everything >> 12)

	cleanedX := Clamp((float64(rawX)-2048.0)/2047.0, -1.0, 1.0)
	cleanedY := Clamp((float64(rawY)-2048.0)/2047.0, -1.0, 1.0)

	return FromOctUV(vector.NewVector2(cleanedX, cleanedY))
}

// PackOct16 maps a unit vector to a 2D UV of a octahedron, and then writes the
// 2D coordinates to 2 bytes, 8bits per coordinate.
func PackOct16(v vector.Vector3) []byte {
	uvCords := MapToOctUVPrecise(v, 16)

	// 2 ^ 8 = 256;
	x := uint(math.Floor(uvCords.X()*127) + 128)
	y := uint(math.Floor(uvCords.Y()*127) + 128)
	everything := (x << 8) | y

	return []byte{
		(byte)(everything & 0xFF),
		(byte)((everything >> 8) & 0xFF),
	}
}

// UnpackOct16 reads in two 8bit numbers and converts from 2D octahedron UV to
// 3D unit sphere coordinates.
func UnpackOct16(b []byte) vector.Vector3 {
	everything := uint(b[0]) | (uint(b[1]) << 8)
	rawY := (int)((everything) & 0b11111111)
	rawX := (int)(everything >> 8)

	cleanedX := Clamp((float64(rawX)-128.0)/127.0, -1.0, 1.0)
	cleanedY := Clamp((float64(rawY)-128.0)/127.0, -1.0, 1.0)

	return FromOctUV(vector.NewVector2(cleanedX, cleanedY))
}

// MapToOctUVPrecise brute force finds an optimal UV coordinate that minimizes
// rounding error.
func MapToOctUVPrecise(v vector.Vector3, n int) vector.Vector2 {
	s := MapToOctUV(v) // Remap to the square

	// Each snorm’s max value interpreted as an integer,
	// e.g., 127.0 for snorm8
	M := float64(int(1)<<((n/2)-1)) - 1.0

	// Remap components to snorm(n/2) precision...with floor instead
	// of round (see equation 1)
	s = floorVec2(clampVec2(s, -1.0, 1.0).MultByConstant(M)).MultByConstant(1.0 / M)
	bestRepresentation := s
	highestCosine := FromOctUV(s).Dot(v)

	// Test all combinations of floor and ceil and keep the best.
	// Note that at +/- 1, this will exit the square... but that
	// will be a worse encoding and never win.
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			// This branch will be evaluated at compile time
			if (i != 0) || (j != 0) {
				// Offset the bit pattern (which is stored in floating
				// point!) to effectively change the rounding mode
				// (when i or j is 0: floor, when it is one: ceiling)
				candidate := vector.NewVector2(float64(i), float64(j)).MultByConstant(1 / M).Add(s)
				cosine := FromOctUV(candidate).Dot(v)
				if cosine > highestCosine {
					bestRepresentation = candidate
					highestCosine = cosine
				}
			}
		}
	}

	return bestRepresentation
}

// MapToOctUV converts a 3D sphere's coordinates to a 2D octahedron UV.
func MapToOctUV(v vector.Vector3) vector.Vector2 {
	// Project the sphere onto the octahedron, and then onto the xy plane
	// vec2 p = v.xy * (1.0 / (abs(v.x) + abs(v.y) + abs(v.z)));
	p := vector.
		NewVector2(v.X(), v.Y()).
		MultByConstant(1.0 / (math.Abs(v.X()) + math.Abs(v.Y()) + math.Abs(v.Z())))
	if v.Z() > 0 {
		return p
	}

	// Reflect the folds of the lower hemisphere over the diagonals
	// return ((1.0 - math.Abs(p.yx)) * signNotZero(p))
	return multVect(signNotZero(p), vector.NewVector2(1.0-math.Abs(p.Y()), 1.0-math.Abs(p.X())))
}

// FromOctUV converts a 2D octahedron UV coordinate to a point on a 3D sphere.
func FromOctUV(e vector.Vector2) vector.Vector3 {
	// vec3 v = vec3(e.xy, 1.0 - abs(e.x) - abs(e.y));
	v := vector.NewVector3(e.X(), e.Y(), 1.0-math.Abs(e.X())-math.Abs(e.Y()))

	// if (v.z < 0) v.xy = (1.0 - abs(v.yx)) * signNotZero(v.xy);
	if v.Z() < 0 {
		n := multVect(vector.NewVector2(1.0-math.Abs(v.Y()), 1.0-math.Abs(v.X())), signNotZero(vector.NewVector2(v.X(), v.Y())))
		v = v.SetX(n.X()).SetY(n.Y())
	}

	return v.Normalized()
}
