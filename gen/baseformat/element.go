package baseformat

import (
	"github.com/solid-resourcepack/bbconv/bbformat"
	"github.com/ungerik/go3d/float64/vec3"
)

type Element struct {
	From     vec3.T
	To       vec3.T
	Faces    map[bbformat.Facing]bbformat.Face
	Rotation *Rotation
}

type Rotation struct {
	Origin  vec3.T  `json:"origin"`
	Axis    string  `json:"axis"`
	Angle   float64 `json:"angle"`
	Rescale bool    `json:"rescale"`
}
