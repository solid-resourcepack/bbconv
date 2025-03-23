package baseformat

import "github.com/ungerik/go3d/float64/quaternion"

type Animation struct {
	Loop       bool       `json:"loop"`
	Length     float64    `json:"length"`
	StartDelay float32    `json:"start_delay"`
	LoopDelay  float32    `json:"loop_delay"`
	Name       string     `json:"name"`
	Animators  []Animator `json:"animators"`
}

type Animator struct {
	Name     string             `json:"name"`
	Bone     string             `json:"bone"`
	Position []PositionKeyframe `json:"position,omitempty"`
	Rotation []RotationKeyframe `json:"rotation,omitempty"`
	Scale    []ScaleKeyframe    `json:"scale,omitempty"`
}

type PositionKeyframe struct {
	Time          float32 `json:"time"`
	Point         Point   `json:"point"`
	Interpolation string  `json:"interpolation"`
}

type RotationKeyframe struct {
	Time          float32      `json:"time"`
	Point         quaternion.T `json:"point"`
	Interpolation string       `json:"interpolation"`
}

type ScaleKeyframe struct {
	Time          float32 `json:"time"`
	Point         Point   `json:"point"`
	Interpolation string  `json:"interpolation"`
}

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}
