package unitpacking

import "github.com/EliCDavis/vector"

// Quadrant2D represents a quadrant in a 2D space.
type Quadrant2D int

const (
	TopLeft Quadrant2D = iota
	TopRight
	BottomLeft
	BottomRight
)

// QuadRecurse recursively builds a quad tree based on the given Vector2. The
// tree's depth is determined by the number of levels passed in.
func QuadRecurse(in, min, max vector.Vector2, levels int) []Quadrant2D {
	if levels <= 0 {
		return nil
	}

	midX := ((max.X() + min.X()) / 2)
	midY := ((max.Y() + min.Y()) / 2)

	var dir Quadrant2D
	var newMinX float64
	var newMinY float64
	var newMaxX float64
	var newMaxY float64

	if in.X() < midX {
		newMinX = min.X()
		newMaxX = midX

		if in.Y() < midY {
			newMinY = min.Y()
			newMaxY = midY
			dir = BottomLeft
		} else {
			newMinY = midY
			newMaxY = max.Y()
			dir = TopLeft
		}
	} else {
		newMinX = midX
		newMaxX = max.X()

		if in.Y() < midY {
			newMinY = min.Y()
			newMaxY = midY
			dir = BottomRight
		} else {
			newMinY = midY
			newMaxY = max.Y()
			dir = TopRight
		}
	}

	return append(
		QuadRecurse(
			in,
			vector.NewVector2(newMinX, newMinY),
			vector.NewVector2(newMaxX, newMaxY),
			levels-1,
		),
		dir,
	)
}

// Vec2ToByteQuad creates a quadtree of depth 4 and encodes itself into a
// single byte
func Vec2ToByteQuad(v vector.Vector2) byte {
	results := QuadRecurse(
		v,
		vector.NewVector2(-1, -1),
		vector.NewVector2(1, 1),
		4,
	)
	return byte(results[0]) | (byte(results[1]) << 2) | (byte(results[2]) << 4) | (byte(results[3]) << 6)
}

// ByteToVec2 calculates a Vector2 based on the encoded quadtree inside the
// byte.
func ByteQuadToVec2(b byte) vector.Vector2 {
	directions := make([]Quadrant2D, 4)
	directions[3] = Quadrant2D(b & 0b11)
	directions[2] = Quadrant2D((b >> 2) & 0b11)
	directions[1] = Quadrant2D((b >> 4) & 0b11)
	directions[0] = Quadrant2D((b >> 6) & 0b11)
	return recalc(directions)
}

// Vec2ToTwoByteQuad creates a quadtree of depth 8 and encodes itself in two
// bytes
func Vec2ToTwoByteQuad(v vector.Vector2) []byte {
	results := QuadRecurse(
		v,
		vector.NewVector2(-1, -1),
		vector.NewVector2(1, 1),
		8,
	)
	return []byte{
		byte(results[0]) | (byte(results[1]) << 2) | (byte(results[2]) << 4) | (byte(results[3]) << 6),
		byte(results[4]) | (byte(results[5]) << 2) | (byte(results[6]) << 4) | (byte(results[7]) << 6),
	}
}

// Vec2ToThreeByteQuad creates a quadtree of depth 12 and encodes itself in
// three bytes
func Vec2ToThreeByteQuad(v vector.Vector2) []byte {
	results := QuadRecurse(
		v,
		vector.NewVector2(-1, -1),
		vector.NewVector2(1, 1),
		12,
	)
	return []byte{
		byte(results[0]) | (byte(results[1]) << 2) | (byte(results[2]) << 4) | (byte(results[3]) << 6),
		byte(results[4]) | (byte(results[5]) << 2) | (byte(results[6]) << 4) | (byte(results[7]) << 6),
		byte(results[8]) | (byte(results[9]) << 2) | (byte(results[10]) << 4) | (byte(results[11]) << 6),
	}
}

// ThreeByteQuadToVec2 calculates a Vector2 based on the encoded quadtree inside
// the 3 bytes
func ThreeByteQuadToVec2(b []byte) vector.Vector2 {
	directions := make([]Quadrant2D, 12)

	directions[11] = Quadrant2D(b[0] & 0b11)
	directions[10] = Quadrant2D((b[0] >> 2) & 0b11)
	directions[9] = Quadrant2D((b[0] >> 4) & 0b11)
	directions[8] = Quadrant2D((b[0] >> 6) & 0b11)
	directions[7] = Quadrant2D(b[1] & 0b11)
	directions[6] = Quadrant2D((b[1] >> 2) & 0b11)
	directions[5] = Quadrant2D((b[1] >> 4) & 0b11)
	directions[4] = Quadrant2D((b[1] >> 6) & 0b11)
	directions[3] = Quadrant2D(b[2] & 0b11)
	directions[2] = Quadrant2D((b[2] >> 2) & 0b11)
	directions[1] = Quadrant2D((b[2] >> 4) & 0b11)
	directions[0] = Quadrant2D((b[2] >> 6) & 0b11)

	return recalc(directions)
}

// TwoByteQuadToVec2 calculates a Vector2 based on the encoded quadtree inside
// the 2 bytes
func TwoByteQuadToVec2(b []byte) vector.Vector2 {
	directions := make([]Quadrant2D, 8)

	directions[7] = Quadrant2D(b[0] & 0b11)
	directions[6] = Quadrant2D((b[0] >> 2) & 0b11)
	directions[5] = Quadrant2D((b[0] >> 4) & 0b11)
	directions[4] = Quadrant2D((b[0] >> 6) & 0b11)
	directions[3] = Quadrant2D(b[1] & 0b11)
	directions[2] = Quadrant2D((b[1] >> 2) & 0b11)
	directions[1] = Quadrant2D((b[1] >> 4) & 0b11)
	directions[0] = Quadrant2D((b[1] >> 6) & 0b11)

	return recalc(directions)
}

// Vec2ToFourByteQuad creates a quadtree of depth 16 and encodes itself in
// 4 bytes
func Vec2ToFourByteQuad(v vector.Vector2) []byte {
	results := QuadRecurse(
		v,
		vector.NewVector2(-1, -1),
		vector.NewVector2(1, 1),
		16,
	)
	return []byte{
		byte(results[0]) | (byte(results[1]) << 2) | (byte(results[2]) << 4) | (byte(results[3]) << 6),
		byte(results[4]) | (byte(results[5]) << 2) | (byte(results[6]) << 4) | (byte(results[7]) << 6),
		byte(results[8]) | (byte(results[9]) << 2) | (byte(results[10]) << 4) | (byte(results[11]) << 6),
		byte(results[12]) | (byte(results[13]) << 2) | (byte(results[14]) << 4) | (byte(results[15]) << 6),
	}
}

// FourByteQuadToVec2 calculates a Vector2 based on the encoded quadtree inside
// the 4 bytes
func FourByteQuadToVec2(b []byte) vector.Vector2 {
	directions := make([]Quadrant2D, 16)

	directions[15] = Quadrant2D(b[0] & 0b11)
	directions[14] = Quadrant2D((b[0] >> 2) & 0b11)
	directions[13] = Quadrant2D((b[0] >> 4) & 0b11)
	directions[12] = Quadrant2D((b[0] >> 6) & 0b11)

	directions[11] = Quadrant2D(b[1] & 0b11)
	directions[10] = Quadrant2D((b[1] >> 2) & 0b11)
	directions[9] = Quadrant2D((b[1] >> 4) & 0b11)
	directions[8] = Quadrant2D((b[1] >> 6) & 0b11)

	directions[7] = Quadrant2D(b[2] & 0b11)
	directions[6] = Quadrant2D((b[2] >> 2) & 0b11)
	directions[5] = Quadrant2D((b[2] >> 4) & 0b11)
	directions[4] = Quadrant2D((b[2] >> 6) & 0b11)

	directions[3] = Quadrant2D(b[3] & 0b11)
	directions[2] = Quadrant2D((b[3] >> 2) & 0b11)
	directions[1] = Quadrant2D((b[3] >> 4) & 0b11)
	directions[0] = Quadrant2D((b[3] >> 6) & 0b11)

	return recalc(directions)
}

func recalc(directions []Quadrant2D) vector.Vector2 {
	multiplyer := 0.5
	outVec := vector.Vector2Zero()
	for _, v := range directions {
		switch v {
		case TopRight:
			outVec = outVec.Add(vector.NewVector2(multiplyer, multiplyer))
			break

		case TopLeft:
			outVec = outVec.Add(vector.NewVector2(-multiplyer, multiplyer))
			break

		case BottomLeft:
			outVec = outVec.Add(vector.NewVector2(-multiplyer, -multiplyer))
			break

		case BottomRight:
			outVec = outVec.Add(vector.NewVector2(multiplyer, -multiplyer))
			break
		}
		multiplyer /= 2.0
	}

	return outVec
}
