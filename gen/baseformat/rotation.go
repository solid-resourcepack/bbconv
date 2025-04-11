package baseformat

import (
	"github.com/ungerik/go3d/float64/quaternion"
	"github.com/ungerik/go3d/float64/vec3"
	"log"
	"math"
)

var rotationDegree = 22.5

var validRotationDegrees = []float64{
	0,
	rotationDegree,
	rotationDegree * 2,
	-rotationDegree,
	-rotationDegree * 2,
}

func identifier(rotation *vec3.T) vec3.T {
	if rotation == nil {
		return vec3.Zero
	}
	if checkValidRotationVector(*rotation) {
		return vec3.Zero
	}
	log.Println("invalid rotation vector", *rotation)
	return *rotation
}

func checkValidRotationVector(rotation vec3.T) bool {
	var i = 0
	if rotation[0] != 0 {
		i++
	}
	if rotation[1] != 0 {
		i++
	}
	if rotation[2] != 0 {
		i++
	}
	return i == 1 && checkValidDegree(rotation[0]) && checkValidDegree(rotation[1]) && checkValidDegree(rotation[2])
}

func checkValidDegree(degree float64) bool {
	for _, validDegree := range validRotationDegrees {
		if degree == validDegree {
			return true
		}
	}
	return false
}

// Degrees to Radians
func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func ToQuaternion(v *vec3.T) quaternion.T {
	x, y, z := degToRad(v[0]), degToRad(v[1]), degToRad(v[2])

	// Compute half angles
	cx, sx := math.Cos(x/2), math.Sin(x/2)
	cy, sy := math.Cos(y/2), math.Sin(y/2)
	cz, sz := math.Cos(z/2), math.Sin(z/2)

	return quaternion.T{
		sx*cy*cz - cx*sy*sz,
		cx*sy*cz + sx*cy*sz,
		cx*cy*sz - sx*sy*cz,
		cx*cy*cz + sx*sy*sz,
	}
}
