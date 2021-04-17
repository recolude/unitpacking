package main

import (
	"compress/flate"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/EliCDavis/vector"
	"github.com/recolude/unitpacking/unitpacking"
)

func main() {

	width := 512
	height := 512
	numOfVectors := width * height
	unitVectors := make([]vector.Vector3, numOfVectors)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Generate a bunch of unit vectors
	for i := 0; i < numOfVectors; i++ {
		unitVectors[i] = vector.NewVector3(
			(rand.Float64()*2.0)-1.0,
			(rand.Float64()*2.0)-1.0,
			(rand.Float64()*2.0)-1.0,
		).Normalized()
	}

	out, err := os.Create("example.data")
	if err != nil {
		panic(err)
	}

	comressedWriter, err := flate.NewWriter(out, 9)
	if err != nil {
		panic(err)
	}

	// Write out unit vectors in packed format
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			data := unitpacking.PackOct24(unitVectors[y+(x*width)])
			comressedWriter.Write(data)
			img.Set(x, y, color.RGBA{
				R: data[0],
				G: data[1],
				B: data[2],
				A: 255,
			})
		}
	}

	comressedWriter.Flush()

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)

}
