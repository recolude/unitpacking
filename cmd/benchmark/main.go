package main

import (
	"bufio"
	"bytes"
	"compress/flate"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/EliCDavis/mango"
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
			avgErrOut = fmt.Sprintf("%.6f", *e.avgError)
		}

		n, err := fmt.Fprintf(out, "| %s | %s | %s | %s | %d b | %d b | %.4f\n", ds.set, e.method, runtimeOut, avgErrOut, e.uncomressed, e.compressed, float64(e.uncomressed)/float64(e.compressed))
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
	results := make([]runResultEntry, len(methods)+1)
	results[0] = runbaseline(unitVectors)
	for i, m := range methods {
		results[i+1] = runBenchEnry(unitVectors, m)
	}
	return dataset{
		set:     name,
		entries: results,
	}
}

func getDatasetPathsFromDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	validFiles := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.ToLower(filepath.Ext(file.Name())) == ".obj" {
			validFiles = append(validFiles, file.Name())
		}
	}

	return validFiles, nil
}

func calcFlatNormals(m mango.Mesh) []vector.Vector3 {

	normals := make([]vector.Vector3, len(m.Vertices()))

	for _, tri := range m.Triangles() {
		// normalize(cross(B-A, C-A))
		normalized := m.Vertices()[tri.P2()].Sub(m.Vertices()[tri.P1()]).Cross(m.Vertices()[tri.P3()].Sub(m.Vertices()[tri.P1()]))
		normals[tri.P1()] = normalized
		normals[tri.P2()] = normalized
		normals[tri.P3()] = normalized
	}

	return normals
}

func calcSmoothNormals(m mango.Mesh) []vector.Vector3 {
	normals := make([]vector.Vector3, len(m.Vertices()))
	for _, tri := range m.Triangles() {
		// normalize(cross(B-A, C-A))
		normalized := m.Vertices()[tri.P2()].Sub(m.Vertices()[tri.P1()]).Cross(m.Vertices()[tri.P3()].Sub(m.Vertices()[tri.P1()]))
		normals[tri.P1()] = normals[tri.P1()].Add(normalized)
		normals[tri.P2()] = normals[tri.P2()].Add(normalized)
		normals[tri.P3()] = normals[tri.P3()].Add(normalized)
	}
	return normals
}

func firstWords(value string, count int) (string, string) {
	// Loop over all indexes in the string.
	for i := range value {
		// If we encounter a space, reduce the count.
		if value[i] == ' ' {
			count--
			// When no more words required, return a substring.
			if count == 0 {
				return value[0:i], value[i+1:]
			}
		}
	}
	// Return the entire string.
	return value, ""
}

func strToVector(str string) (*vector.Vector3, error) {
	components := strings.Split(str, " ")

	if len(components) != 3 {
		return nil, errors.New("unable to parse: " + str)
	}

	xParse, err := strconv.ParseFloat(components[0], 64)
	if err != nil {
		return nil, errors.New("unable to parse X componenent: " + components[0])
	}

	yParse, err := strconv.ParseFloat(components[1], 64)
	if err != nil {
		return nil, errors.New("unable to parse Y componenent: " + components[1])
	}

	zParse, err := strconv.ParseFloat(components[2], 64)
	if err != nil {
		return nil, errors.New("unable to parse Z componenent: " + components[2])
	}

	v := vector.NewVector3(xParse, yParse, zParse)
	return &v, nil
}

func strToFaceIndexes(str string) (int, int, int, error) {
	components := strings.Split(str, " ")

	if len(components) != 3 {
		return -1, -1, -1, fmt.Errorf("unable to parse: (%s)", str)
	}

	v1Components := strings.Split(components[0], "/")
	v1Parse, err := strconv.Atoi(v1Components[0])
	if err != nil {
		return -1, -1, -1, errors.New("unable to parse X componenent: " + v1Components[0])
	}

	v2Components := strings.Split(components[1], "/")
	v2Parse, err := strconv.Atoi(v2Components[0])
	if err != nil {
		return -1, -1, -1, errors.New("unable to parse Y componenent: " + v2Components[1])
	}

	v3Components := strings.Split(components[2], "/")
	v3Parse, err := strconv.Atoi(v3Components[0])
	if err != nil {
		return -1, -1, -1, errors.New("unable to parse Z componenent: " + v3Components[0])
	}

	return v1Parse, v2Parse, v3Parse, nil
}

func importOBJ(objStream io.Reader) (mango.Mesh, error) {
	if objStream == nil {
		return mango.NewEmptyMesh(), errors.New("Need a reader to read obj from")
	}

	vertices := make([]vector.Vector3, 0)
	faces := make([]mango.Tri, 0)

	scanner := bufio.NewScanner(objStream)
	for scanner.Scan() {
		ln := scanner.Text()
		firstWord, rest := firstWords(ln, 1)
		if firstWord == "v" {
			vector, err := strToVector(strings.TrimSpace(rest))
			if err != nil {
				return mango.NewEmptyMesh(), err
			}
			vertices = append(vertices, *vector)
		}

		if firstWord == "f" {
			v1, v2, v3, err := strToFaceIndexes(strings.TrimSpace(rest))
			if err != nil {
				return mango.NewEmptyMesh(), err
			}

			faces = append(faces, mango.NewTri(v1-1, v2-1, v3-1))
		}
	}

	if scanner.Err() != nil {
		return mango.NewEmptyMesh(), scanner.Err()
	}

	return mango.NewMesh(vertices, faces), nil
}

func loadModel(filePath string) (mango.Mesh, error) {
	dat, err := os.Open(filePath)
	if err != nil {
		return mango.NewEmptyMesh(), err
	}
	return importOBJ(dat)
}

func runbaseline(unitVectors []vector.Vector3) runResultEntry {
	compressedOut := bytes.Buffer{}
	comressedWriter, err := flate.NewWriter(&compressedOut, 9)
	if err != nil {
		panic(err)
	}

	b := make([]byte, 4)
	for _, v := range unitVectors {
		binary.LittleEndian.PutUint32(b, math.Float32bits(float32(v.X())))
		comressedWriter.Write(b)

		binary.LittleEndian.PutUint32(b, math.Float32bits(float32(v.Y())))
		comressedWriter.Write(b)

		binary.LittleEndian.PutUint32(b, math.Float32bits(float32(v.Z())))
		comressedWriter.Write(b)
	}

	comressedWriter.Flush()

	return runResultEntry{
		method:      "Baseline",
		compressed:  compressedOut.Len(),
		uncomressed: len(unitVectors) * 12,
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

	pathToLoadFrom := "../../../common-3d-test-models/data"
	// pathToLoadFrom := os.Args[1]
	availableFiles, err := getDatasetPathsFromDir(pathToLoadFrom)
	if err != nil {
		panic(err)
	}

	unitWriters := []unitWriter{
		alg24Writer{os.Stdout},
		coarse24Writer{os.Stdout},
		oct24Writer{os.Stdout},
	}

	fmt.Fprintln(os.Stdout, "| Dataset | Method | Runtime | Average Error | Uncompressed | Compressed | Compression Ratio |")
	fmt.Fprintln(os.Stdout, "|-|-|-|-|-|-|-|")

	runDataset(unitVectors, "10 million random", unitWriters).Write(os.Stdout)
	for _, f := range availableFiles {
		model, err := loadModel(filepath.Join(pathToLoadFrom, f))
		if err != nil {
			continue
		}
		datasetName := filepath.Base(f)

		_, err = runDataset(calcFlatNormals(model), datasetName+" Flat", unitWriters).Write(os.Stdout)
		if err != nil {
			panic(err)
		}
		_, err = runDataset(calcSmoothNormals(model), datasetName+" Smooth", unitWriters).Write(os.Stdout)
		if err != nil {
			panic(err)
		}
	}

}
