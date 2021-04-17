package unitpacking

import "github.com/EliCDavis/vector"

// PackOctQuad16 maps a unit vector to a 2D UV of a octahedron, and then
// encodes the 2D coordinates inside a quad tree.
func PackOctQuad16(v vector.Vector3) []byte {
	return Vec2ToTwoByteQuad(MapToOctUV(v))
}

// UnpackOctQuad16 builds a 2D coordinate from the encoded quadtree and then
// converts the 2D octahedron UV to 3D unit sphere coordinates.
func UnpackOctQuad16(b []byte) vector.Vector3 {
	return FromOctUV(TwoByteQuadToVec2(b))
}

// PackOctQuad24 maps a unit vector to a 2D UV of a octahedron, and then
// encodes the 2D coordinates inside a quad tree.
func PackOctQuad24(v vector.Vector3) []byte {
	return Vec2ToThreeByteQuad(MapToOctUV(v))
}

// UnpackOctQuad24 builds a 2D coordinate from the encoded quadtree and then
// converts the 2D octahedron UV to 3D unit sphere coordinates.
func UnpackOctQuad24(b []byte) vector.Vector3 {
	return FromOctUV(ThreeByteQuadToVec2(b))
}

// PackOctQuad32 maps a unit vector to a 2D UV of a octahedron, and then
// encodes the 2D coordinates inside a quad tree.
func PackOctQuad32(v vector.Vector3) []byte {
	return Vec2ToFourByteQuad(MapToOctUV(v))
}

// UnpackOctQuad32 builds a 2D coordinate from the encoded quadtree and then
// converts the 2D octahedron UV to 3D unit sphere coordinates.
func UnpackOctQuad32(b []byte) vector.Vector3 {
	return FromOctUV(FourByteQuadToVec2(b))
}
