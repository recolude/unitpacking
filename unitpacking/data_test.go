package unitpacking_test

import "github.com/EliCDavis/vector"

var testVectors []vector.Vector3 = []vector.Vector3{
	vector.NewVector3(0, 1, 2),
	vector.NewVector3(1, 0, 0),
	vector.NewVector3(1, 1, 0),
	vector.NewVector3(1, 1, 1),
	vector.NewVector3(0, 1, 0),
	vector.NewVector3(0, 0, 1),
	vector.NewVector3(-1, 0, 0),
	vector.NewVector3(0, -1, 0),
	vector.NewVector3(0, 0, -1),
	vector.NewVector3(-1, -1, -1),
	vector.NewVector3(-1, -1, 0),
	vector.NewVector3(-1, 1, -1),
	vector.NewVector3(-1, 1, -1),
	vector.NewVector3(-0.997605826445425, 0.06365823804882093, -0.027023022974122023),
	vector.NewVector3(0.7180684556508264, -0.6958747397502424, -0.011663600506254649),
}
