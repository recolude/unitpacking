package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/EliCDavis/vector"
	"github.com/recolude/unitpacking/unitpacking"
)

type dataset struct {
	set     string
	entries []runResultEntry
}

func (ds dataset) Write(out io.Writer) (int, error) {
	writtenCount := 0

	for _, e := range ds.entries {

		runtimeOut := "N/A"
		if e.duration != nil {
			runtimeOut = fmt.Sprintf("%s", *e.duration)
		}

		avgErrOut := "N/A"
		if e.avgError != nil {
			avgErrOut = fmt.Sprintf("%.4f", *e.avgError)
		}

		n, err := fmt.Fprintf(out, "| %s | %s | %s | %s | %d | %d | %.4f\n", ds.set, e.method, runtimeOut, avgErrOut, e.uncomressed, e.compressed, float64(e.uncomressed)/float64(e.compressed))
		writtenCount += n
		if err != nil {
			return writtenCount, err
		}
	}

	return writtenCount, nil
}

type runResultEntry struct {
	method      string
	compressed  int
	uncomressed int
	duration    *time.Duration
	avgError    *float64
}

func (rre runResultEntry) compressionRatio() float64 {
	return float64(rre.uncomressed) / float64(rre.compressed)
}

type unitWriter interface {
	pack(v vector.Vector3) []byte
	unpack(b []byte) vector.Vector3
	method() string
}

type alg24Writer struct{ out io.Writer }

func (alg24w alg24Writer) method() string                 { return "alg24" }
func (alg24w alg24Writer) pack(v vector.Vector3) []byte   { return unitpacking.PackAlg24(v) }
func (alg24w alg24Writer) unpack(b []byte) vector.Vector3 { return unitpacking.UnpackAlg24(b) }

type oct24Writer struct{ out io.Writer }

func (oct24w oct24Writer) method() string                 { return "oct24" }
func (oct24w oct24Writer) pack(v vector.Vector3) []byte   { return unitpacking.PackOct24(v) }
func (oct24w oct24Writer) unpack(b []byte) vector.Vector3 { return unitpacking.UnpackOct24(b) }

type coarse24Writer struct{ out io.Writer }

func (coarse24w coarse24Writer) method() string                 { return "coarse24" }
func (coarse24w coarse24Writer) pack(v vector.Vector3) []byte   { return unitpacking.PackCoarse24(v) }
func (coarse24w coarse24Writer) unpack(b []byte) vector.Vector3 { return unitpacking.UnpackCoarse24(b) }

func runBenchEnry(unitVectors []vector.Vector3, uw unitWriter) runResultEntry {
	accErr := 0.0

	// Just time it...
	start := time.Now()
	for _, v := range unitVectors {
		uw.unpack(uw.pack(v))
	}
	duration := time.Since(start)

	// Now calculate error and compression
	out := bytes.Buffer{}
	compressedOut := bytes.Buffer{}
	comressedWriter, err := flate.NewWriter(&compressedOut, 9)
	if err != nil {
		panic(err)
	}
	for _, v := range unitVectors {
		unpacked := uw.unpack(uw.pack(v))
		out.Write(uw.pack(v))
		comressedWriter.Write(uw.pack(v))
		accErr += math.Abs(v.X() - unpacked.X())
		accErr += math.Abs(v.Y() - unpacked.Y())
		accErr += math.Abs(v.Z() - unpacked.Z())
	}

	comressedWriter.Flush()

	avgErr := accErr / float64(len(unitVectors)*3)
	return runResultEntry{
		method:      uw.method(),
		compressed:  compressedOut.Len(),
		uncomressed: out.Len(),
		avgError:    &avgErr,
		duration:    &duration,
	}
}

func runDataset(unitVectors []vector.Vector3, name string, methods []unitWriter) dataset {
	results := make([]runResultEntry, len(methods))
	for i, m := range methods {
		results[i] = runBenchEnry(unitVectors, m)
	}
	return dataset{
		set:     name,
		entries: results,
	}
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

	setsToRun := map[string]struct {
		input []vector.Vector3
	}{
		"10 million random": {input: unitVectors},
	}

	unitWriters := []unitWriter{
		alg24Writer{os.Stdout},
		coarse24Writer{os.Stdout},
		oct24Writer{os.Stdout},
	}

	fmt.Fprintln(os.Stdout, "| Dataset | Method | Runtime | Average Error | Uncompressed | Compressed | Compression Ratio |")
	fmt.Fprintln(os.Stdout, "|-|-|-|-|-|-|-|")

	for name, set := range setsToRun {
		_, err := runDataset(set.input, name, unitWriters).Write(os.Stdout)
		if err != nil {
			panic(err)
		}
	}

}
