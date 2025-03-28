package mcformat

import (
	"github.com/solid-resourcepack/bbconv/baseformat"
	"github.com/solid-resourcepack/bbconv/bbformat"
)

type Element struct {
	From     baseformat.Vector        `json:"from"`
	To       baseformat.Vector        `json:"to"`
	Rotation *Rotation                `json:"rotation,omitempty"`
	Faces    map[bbformat.Facing]Face `json:"faces,omitempty"`
}

type Rotation struct {
	Origin baseformat.Vector `json:"origin"`
	Axis   string            `json:"axis"`
	Angle  float64           `json:"angle"`
}

type Face struct {
	UV      []float32 `json:"uv"`
	Texture string    `json:"texture"`
}
