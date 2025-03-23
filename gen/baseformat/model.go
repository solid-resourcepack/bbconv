package baseformat

type Model struct {
	Name       string
	BoneTree   []Bone            `json:"bone_tree"`
	Animations []Animation       `json:"animations"`
	Textures   map[string]string `json:"-"`
}
