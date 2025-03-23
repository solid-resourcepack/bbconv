package baseformat

import (
	"github.com/solid-resourcepack/bbconv/bbformat"
	"github.com/solid-resourcepack/bbconv/util"
	"github.com/ungerik/go3d/float64/vec3"
)

func ConvertElements(parent bbformat.Bone, elements []bbformat.Element) {

}

func toVec(points []float32) vec3.T {
	return vec3.T{
		float64(points[0]),
		float64(points[1]),
		float64(points[2]),
	}
}

func relativizeAll(pivot vec3.T, points []vec3.T) []vec3.T {
	return centerAll(pivotAll(pivot, points))
}

func pivotAll(origin vec3.T, points []vec3.T) []vec3.T {
	return util.Map(points, func(v vec3.T) vec3.T {
		return vec3.Sub(&v, &origin)
	})
}

func centerAll(points []vec3.T) []vec3.T {
	mcCenter := vec3.T{8, 8, 8} //The center of a minecraft item
	sub := vec3.Mul(findBiggest(points), &vec3.T{0.5, 0.5, 0.5})
	center := vec3.Sub(&mcCenter, &sub)
	return util.Map(points, func(v vec3.T) vec3.T {
		return vec3.Add(&v, &center)
	})
}

func findBiggest(points []vec3.T) *vec3.T {
	biggest := points[0]
	maxMagnitude := biggest.Length()

	for _, v := range points[1:] {
		mag := v.Length()
		if mag > maxMagnitude {
			biggest = v
			maxMagnitude = mag
		}
	}
	return &biggest
}
