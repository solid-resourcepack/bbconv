package baseformat

import "github.com/solid-resourcepack/bbconv/bbformat"

type Model struct {
	Name       string
	BoneTree   []Bone              `json:"bone_tree"`
	Animations []Animation         `json:"animations"`
	Textures   map[string]string   `json:"-"`
	Resolution bbformat.Resolution `json:"-"`
}
