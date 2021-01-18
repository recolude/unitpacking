package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/EliCDavis/vector"
	"github.com/recolude/unitpacking/unitpacking"
)

func fileWritingBenchmark(unitVectors []vector.Vector3) {
	algOut, err := os.Create("alg.out")
	if err != nil {
		panic(err)
	}

	coarseOut, err := os.Create("coarse.out")
	if err != nil {
		panic(err)
	}

	octOut, err := os.Create("oct.out")
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for _, v := range unitVectors {
		_, err = algOut.Write(unitpacking.PackAlg24(v))
		if err != nil {
			panic(err)
		}
	}
	duration := time.Since(start)
	fmt.Printf("Packing And Writing Alg Took %s\n", duration)

	start = time.Now()
	for _, v := range unitVectors {
		_, err = coarseOut.Write(unitpacking.PackCoarse24(v))
		if err != nil {
			panic(err)
		}
	}
	duration = time.Since(start)
	fmt.Printf("Packing And Writing Coarse Took %s\n", duration)

	start = time.Now()
	for _, v := range unitVectors {
		_, err = octOut.Write(unitpacking.PackOct24(v))
		if err != nil {
			panic(err)
		}
	}
	duration = time.Since(start)
	fmt.Printf("Packing And Writing Oct Took %s\n", duration)

}

func avgErrorBenchmark(unitVectors []vector.Vector3) {

	accErr := 0.0
	for _, v := range unitVectors {
		unpacked := unitpacking.UnpackAlg24(unitpacking.PackAlg24(v))
		accErr += math.Abs(v.X() - unpacked.X())
		accErr += math.Abs(v.Y() - unpacked.Y())
		accErr += math.Abs(v.Z() - unpacked.Z())
	}
	fmt.Printf("Avg Alg Error %f\n", accErr/float64(3*len(unitVectors)))

	accErr = 0.0
	for _, v := range unitVectors {
		unpacked := unitpacking.UnpackCoarse24(unitpacking.PackCoarse24(v))
		accErr += math.Abs(v.X() - unpacked.X())
		accErr += math.Abs(v.Y() - unpacked.Y())
		accErr += math.Abs(v.Z() - unpacked.Z())
	}
	fmt.Printf("Avg Coarse Error %f\n", accErr/float64(3*len(unitVectors)))

	accErr = 0.0
	for _, v := range unitVectors {
		unpacked := unitpacking.UnpackOct24(unitpacking.PackOct24(v))
		accErr += math.Abs(v.X() - unpacked.X())
		accErr += math.Abs(v.Y() - unpacked.Y())
		accErr += math.Abs(v.Z() - unpacked.Z())
	}
	fmt.Printf("Avg Oct Error %f\n", accErr/float64(3*len(unitVectors)))

}

func main() {
	numOfVectors := 10000000
	unitVectors := make([]vector.Vector3, numOfVectors)

	for i := 0; i < numOfVectors; i++ {
		unitVectors[i] = vector.NewVector3(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		).Normalized()
	}

	fileWritingBenchmark(unitVectors)
	avgErrorBenchmark(unitVectors)
}
