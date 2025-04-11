package baseformat

import (
	"fmt"
	"github.com/ungerik/go3d/float64/quaternion"
	"github.com/ungerik/go3d/float64/vec3"
)

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
	Position      Vector  `json:"position"`
	Interpolation string  `json:"interpolation"`
}

type RotationKeyframe struct {
	Time          float32    `json:"time"`
	LeftRotation  Quaternion `json:"left_rotation"`
	RightRotation Quaternion `json:"right_rotation"`
	Interpolation string     `json:"interpolation"`
}

type ScaleKeyframe struct {
	Time          float32 `json:"time"`
	Scale         Vector  `json:"scale"`
	Interpolation string  `json:"interpolation"`
}

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type Vector vec3.T

func (v Vector) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`
		{
			"x": "%f",
			"y": "%f",
            "z": "%f"
		}
	`, v[0], v[1], v[2])
	return []byte(str), nil
}

type Quaternion quaternion.T

func (q Quaternion) MarshalJSON() ([]byte, error) {
	var w float64
	if q[0] == 0 && q[1] == 0 && q[2] == 0 && q[3] == 0 {
		w = 1
	} else {
		w = q[3]
	}
	str := fmt.Sprintf(`
		{
			"x": "%f",
			"y": "%f",
            "z": "%f",
			"w": "%f"
		}
	`, q[0], q[1], q[2], w)
	return []byte(str), nil
}
