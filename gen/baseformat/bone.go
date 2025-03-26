package baseformat

import (
	"github.com/ungerik/go3d/float64/vec3"
)

type Bone struct {
	UUID     string    `json:"-"`
	Id       string    `json:"id"`
	Key      string    `json:"model,omitempty"`
	Children []Bone    `json:"children,omitempty"`
	Visuals  []Element `json:"-"`
	Origin   vec3.T    `json:"origin"`
	Visible  bool      `json:"visible"`
	Scale    float64   `json:"scale"`
}

func (b Bone) FindBone(uuid string) *Bone {
	if b.UUID == uuid {
		return &b
	}
	if len(b.Children) == 0 {
		return nil
	}
	for _, child := range b.Children {
		if result := child.FindBone(uuid); result != nil {
			return result
		}
	}
	return nil
}
