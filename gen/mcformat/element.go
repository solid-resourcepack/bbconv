package mcformat

import "github.com/solid-resourcepack/bbconv/baseformat"

type Element struct {
	From     baseformat.Vector `json:"from"`
	To       baseformat.Vector `json:"to"`
	Rotation Rotation          `json:"rotation,omitempty"`
	Faces    map[Facing]Face   `json:"faces,omitempty"`
}

type Rotation struct {
	Origin baseformat.Vector `json:"origin"`
	Axis   string            `json:"axis"`
	Angle  float64           `json:"angle"`
}

type Face struct {
	UV      []float64 `json:"uv"`
	Texture string    `json:"texture"`
}

type Facing string

var (
	FacingNorth Facing = "north"
	FacingSouth Facing = "south"
	FacingEast  Facing = "east"
	FacingWest  Facing = "west"
	FacingDown  Facing = "down"
	FacingUp    Facing = "up"
)
