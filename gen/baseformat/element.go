package baseformat

import (
	"github.com/solid-resourcepack/bbconv/bbformat"
	"github.com/ungerik/go3d/float64/vec3"
)

type Element struct {
	From  vec3.T
	To    vec3.T
	Faces bbformat.Faces
}
