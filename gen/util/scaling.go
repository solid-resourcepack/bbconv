package util

import (
	"math"

	"github.com/ungerik/go3d/float64/vec3"
)

type Rectangle3D struct {
	Min vec3.T
	Max vec3.T
}

func GetScalingFactor(r1, r2 Rectangle3D) float64 {
	realMin := vec3.Min(&r1.Min, &vec3.T{8, 8, 8}) // TODO: use public constant for origin vector (8,8,8)
	realMax := vec3.Min(&r1.Max, &vec3.T{8, 8, 8}) // TODO: use public constant for origin vector (8,8,8)

	center1 := vec3.T{
		(realMin[0] + realMax[0]) / 2,
		(realMin[1] + realMax[1]) / 2,
		(realMin[2] + realMax[2]) / 2,
	}

	scale := 1.0 // Start with the highest possible scale

	// Check scaling needed for both min and max bounds
	for i := 0; i < 3; i++ {
		selfMin := math.Abs(realMin[i] - center1[i]) // Distance from center to min
		selfMax := math.Abs(realMax[i] - center1[i]) // Distance from center to max

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
