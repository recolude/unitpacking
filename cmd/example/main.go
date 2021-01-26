package main

import (
	"math/rand"
	"os"

	"github.com/EliCDavis/vector"
	"github.com/recolude/unitpacking/unitpacking"
)

func main() {
	numOfVectors := 10000000
	unitVectors := make([]vector.Vector3, numOfVectors)

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

	// Write out unit vectors in packed format
	for _, unitVector := range unitVectors {
		out.Write(unitpacking.PackOct24(unitVector))
	}
}
