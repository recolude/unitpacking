package unitpacking

import (
	"math"

	"github.com/EliCDavis/vector"
)

func normalToByte(normal float64) byte {
	if normal <= -1.0 {
		return 0
	}
	if normal >= 1.0 {
		return 255
	}
	return byte(math.Floor(normal*127.0)) + 128
}

func byteToNormal(b byte) float64 {
	return (float64(b) - 128.0) / 127.0
}

// PackCoarse24 will convert each component of the vector into a single byte,
// and will return those bytes in an array where x is at index 0, and z is at
// index 2
func PackCoarse24(v vector.Vector3) []byte {
	return []byte{
		normalToByte(v.X()),
		normalToByte(v.Y()),
		normalToByte(v.Z()),
	}
}

// UnpackCoarse24 will take a previously packed vector and extract it out of 3
// bytes
func UnpackCoarse24(in []byte) vector.Vector3 {
	return vector.NewVector3(
		byteToNormal(in[0]),
		byteToNormal(in[1]),
		byteToNormal(in[2]),
	)
}
