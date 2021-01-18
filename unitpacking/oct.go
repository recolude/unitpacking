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

// PackOct24 maps a unit vector to a 2D UV of a octahedron, and then writes the
// 2D coordinates to 3 bytes, 12bits per coordinate.
func PackOct24(v vector.Vector3) []byte {
	uvCords := MapToOctUV(v)

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
// 3D unit sphere coordinates
func UnpackOct24(b []byte) vector.Vector3 {
	everything := uint(b[0]) | (uint(b[1]) << 8) | (uint(b[2]) << 16)
	rawY := (int)((everything) & 0b111111111111)
	rawX := (int)(everything >> 12)

	cleanedX := clamp((float64(rawX)-2048.0)/2047.0, -1.0, 1.0)
	cleanedY := clamp((float64(rawY)-2048.0)/2047.0, -1.0, 1.0)

	return FromOctUV(vector.NewVector2(cleanedX, cleanedY))
}

// MapToOctUV converts a 3D sphere's coordinates to a 2D octahedron UV
func MapToOctUV(v vector.Vector3) vector.Vector2 {
	// Project the sphere onto the octahedron, and then onto the xy plane
	// vec2 p = v.xy * (1.0 / (abs(v.x) + abs(v.y) + abs(v.z)));
	p := vector.NewVector2(v.X(), v.Y()).MultByConstant(1.0 / (math.Abs(v.X()) + math.Abs(v.Y()) + math.Abs(v.Z())))
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
