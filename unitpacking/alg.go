package unitpacking

import (
	"math"

	"github.com/EliCDavis/vector"
)

// PackAlg24 converts the x y z components of a normalized vector into a  3
// bytes for efficient transport. Uses trig to pack X into 12 bytes, Y into 11
// bytes, and 1 to denote sign of Z.
func PackAlg24(v vector.Vector3) []byte {
	// 2 ^ 12 = 4,096;
	x := uint(math.Floor(v.X()*2047) + 2048)

	// 2^11 = 2048;
	y := uint(math.Floor(v.Y()*1023) + 1024)

	// Single byte as to whether or not Z was original positive or
	// negative
	zPositive := uint(0)
	if v.Z() >= 0.0 {
		zPositive = 1
	}

	// Combine everything into one number
	everything := (x << 12) | (y << 1) | zPositive

	// Piece out that number
	return []byte{
		(byte)(everything & 0xFF),
		(byte)((everything >> 8) & 0xFF),
		(byte)((everything >> 16) & 0xFF),
	}
}

// UnpackAlg24 will take a previously packed vector and extract it out of 3
// bytes
func UnpackAlg24(b []byte) vector.Vector3 {
	everything := uint(b[0]) | (uint(b[1]) << 8) | (uint(b[2]) << 16)
	rawZ := (everything & 1) == 1
	rawY := (int)((everything >> 1) & 0b11111111111)
	rawX := (int)(everything >> 12)

	cleanedX := clamp((float64(rawX)-2048.0)/2047.0, -1.0, 1.0)
	cleanedY := clamp((float64(rawY)-1024.0)/1023.0, -1.0, 1.0)
	cleanedZ := math.Sqrt(1.0 - (cleanedX * cleanedX) - (cleanedY * cleanedY))
	if !rawZ {
		cleanedZ *= -1
	}

	return vector.NewVector3(cleanedX, cleanedY, cleanedZ)
}
