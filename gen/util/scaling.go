package util

import (
	"github.com/ungerik/go3d/float64/vec3"
	"math"
)

type Rectangle3D struct {
	Min vec3.T
	Max vec3.T
}

func GetScalingFactor(r1, r2 Rectangle3D) float64 {
	center1 := vec3.T{
		(r1.Min[0] + r1.Max[0]) / 2,
		(r1.Min[1] + r1.Max[1]) / 2,
		(r1.Min[2] + r1.Max[2]) / 2,
	}

	scale := 1.0 // Start with the highest possible scale

	// Check scaling needed for both min and max bounds
	for i := 0; i < 3; i++ {
		selfMin := math.Abs(r1.Min[i] - center1[i]) // Distance from center to min
		selfMax := math.Abs(r1.Max[i] - center1[i]) // Distance from center to max

		boundingMin := math.Abs(r2.Min[i] - center1[i]) // Allowed distance to min in r2
		boundingMax := math.Abs(r2.Max[i] - center1[i]) // Allowed distance to max in r2

		if selfMin > 0 {
			scale = math.Min(scale, boundingMin/selfMin)
		}
		if selfMax > 0 {
			scale = math.Min(scale, boundingMax/selfMax)
		}
	}

	return scale
}

func center(r Rectangle3D) *vec3.T {
	return r.Min.Add(&r.Max).Scale(0.5)
}
