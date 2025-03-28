package mcformat

import (
	"github.com/solid-resourcepack/bbconv/bbformat"
)

type Element struct {
	From     []float64                `json:"from"`
	To       []float64                `json:"to"`
	Rotation *Rotation                `json:"rotation,omitempty"`
	Faces    map[bbformat.Facing]Face `json:"faces,omitempty"`
}

type Rotation struct {
	Origin  []float64 `json:"origin"`
	Axis    string    `json:"axis"`
	Angle   float64   `json:"angle"`
	Rescale bool      `json:"rescale"`
}

type Face struct {
	UV      []float32 `json:"uv"`
	Texture string    `json:"texture"`
}
