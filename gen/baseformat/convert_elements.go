package baseformat

import (
	"github.com/solid-resourcepack/bbconv/bbformat"
	"github.com/solid-resourcepack/bbconv/util"
	"github.com/ungerik/go3d/float64/vec3"
	"math"
)

var minPoint = vec3.T{-16, -16, -16}
var maxPoint = vec3.T{32, 32, 32}
var origin = vec3.T{8, 8, 8}
var pointBounds = util.Rectangle3D{
	Min: minPoint,
	Max: maxPoint,
}

func ConvertElement(parent bbformat.Bone, element bbformat.Element) Element {
	pivot := toVec(parent.Origin)
	vecs := make([]vec3.T, 2)
	vecs[0] = toVec(element.From)
	vecs[1] = toVec(element.To)
	vecs = pivotAll(pivot, vecs)
	vecs = originAll(vecs)
	return Element{
		Faces: element.Faces,
		From:  vecs[0],
		To:    vecs[1],
	}
}

func ResizeVisuals(bone *Bone) {
	elements := bone.Visuals
	if len(elements) == 0 {
		return
	}
	points := make([]vec3.T, len(elements)*2)
	for i := 0; i < len(elements); i++ {
		points[i] = elements[i].From
		points[i+len(elements)] = elements[i].To
	}
	minVec, maxVec := CalculateBounds(points)
	scale := CalculateScale(minVec, maxVec)
	Resize(bone, minVec, maxVec, scale)
}

func CalculateBounds(points []vec3.T) (vec3.T, vec3.T) {
	// Initialize min/max bounds from the first vector
	minVec := points[0]
	maxVec := points[0]

	// Compute the bounding box of the given points
	for _, p := range points {
		minVec[0] = math.Min(minVec[0], p[0])
		minVec[1] = math.Min(minVec[1], p[1])
		minVec[2] = math.Min(minVec[2], p[2])

		maxVec[0] = math.Max(maxVec[0], p[0])
		maxVec[1] = math.Max(maxVec[1], p[1])
		maxVec[2] = math.Max(maxVec[2], p[2])
	}
	return minVec, maxVec
}

func CalculateScale(minVec, maxVec vec3.T) float64 {
	scale := util.GetScalingFactor(util.Rectangle3D{
		Min: minVec,
		Max: maxVec,
	}, pointBounds)
	return scale
}

func Resize(bone *Bone, minSize vec3.T, maxSize vec3.T, scaleFactor float64) {
	elements := bone.Visuals

	// Compute the center of the entire bounding box
	add := vec3.Add(&minSize, &maxSize)
	center := add.Scale(0.5)
	for i, b := range elements {
		elements[i].From = scalePoint(&b.From, center, scaleFactor)
		elements[i].To = scalePoint(&b.To, center, scaleFactor)
	}
	bone.Scale = scaleFactor
}

// Scale a point relative to the center
func scalePoint(point, center *vec3.T, scaleFactor float64) vec3.T {
	offset := vec3.Sub(point, center)
	result := offset.Scale(scaleFactor)
	return vec3.Add(center, result)
}

func toVec(points []float32) vec3.T {
	return vec3.T{
		float64(points[0]),
		float64(points[1]),
		float64(points[2]),
	}
}

func pivotAll(origin vec3.T, points []vec3.T) []vec3.T {
	return util.Map(points, func(v vec3.T) vec3.T {
		return vec3.Sub(&v, &origin)
	})
}

func originAll(points []vec3.T) []vec3.T {
	return util.Map(points, func(v vec3.T) vec3.T {
		return vec3.Add(&v, &origin)
	})
}
