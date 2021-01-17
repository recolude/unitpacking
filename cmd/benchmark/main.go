package main

import (
	"math/rand"
	"os"

	"github.com/EliCDavis/vector"
	"github.com/recolude/unitpacking/unitpacking"
)

func main() {
	numOfVectors := 1000000
	unitVectors := make([]vector.Vector3, numOfVectors)

	for i := 0; i < numOfVectors; i++ {
		unitVectors[i] = vector.NewVector3(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		).Normalized()
	}

	algOut, err := os.Create("alg.out")
	if err != nil {
		panic(err)
	}

	coarseOut, err := os.Create("coarse.out")
	if err != nil {
		panic(err)
	}

	for _, v := range unitVectors {
		_, err = algOut.Write(unitpacking.PackAlg24(v))
		if err != nil {
			panic(err)
		}

		_, err = coarseOut.Write(unitpacking.PackCoarse24(v))
		if err != nil {
			panic(err)
		}
	}

}
