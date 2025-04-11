package bbformat

import (
	"encoding/json"
	"fmt"
)

type Bone struct {
	Name       string      `json:"name"`
	Origin     []float32   `json:"origin"`
	Rotation   []float32   `json:"rotation"`
	Color      int         `json:"color"`
	UUID       string      `json:"uuid"`
	MirrorUv   bool        `json:"mirror_uv"`
	IsOpen     bool        `json:"isOpen"`
	Locked     bool        `json:"locked"`
	Visibility bool        `json:"visibility"`
	AutoUv     int         `json:"autouv"`
	Selected   bool        `json:"selected"`
	Children   []BoneChild `json:"children"`
}

type BoneChild struct {
	Ref  *string
	Bone *Bone
}

func (bc *BoneChild) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an element ref
	var ref string
	if err := json.Unmarshal(data, &ref); err == nil {
		bc.Ref = &ref
		return nil
	}

	// Try to unmarshal as a *Bone struct
	var node *Bone
	if err := json.Unmarshal(data, &node); err == nil {
		bc.Bone = node
		return nil
	}

	return fmt.Errorf("invalid child format: %s", string(data))
}
