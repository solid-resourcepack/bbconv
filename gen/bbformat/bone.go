package bbformat

import (
	"errors"
)

type Bone struct {
	Name       string    `json:"name"`
	Origin     []float32 `json:"origin"`
	Color      int       `json:"color"`
	UUID       string    `json:"uuid"`
	MirrorUv   bool      `json:"mirror_uv"`
	IsOpen     bool      `json:"isOpen"`
	Locked     bool      `json:"locked"`
	Visibility bool      `json:"visibility"`
	AutoUv     int       `json:"autouv"`
	Selected   bool      `json:"selected"`
	Children   []any     `json:"children"` //TODO: Use children or string
}

// TODO: do we really need this?
func MapToBone(data map[string]interface{}) (*Bone, error) {
	bone := &Bone{}

	// Extract and assert the types for each field
	if name, ok := data["name"].(string); ok {
		bone.Name = name
	} else {
		return nil, errors.New("missing or invalid 'name' field")
	}

	if origin, ok := data["origin"].([]interface{}); ok {
		// Convert []interface{} to []float32
		bone.Origin = make([]float32, len(origin))
		for i, v := range origin {
			if f, ok := v.(float64); ok {
				bone.Origin[i] = float32(f)
			} else {
				return nil, errors.New("invalid type for 'origin' elements, expected float64")
			}
		}
	} else {
		return nil, errors.New("missing or invalid 'origin' field")
	}

	if color, ok := data["color"].(float64); ok {
		bone.Color = int(color)
	} else {
		return nil, errors.New("missing or invalid 'color' field")
	}

	// Repeat similar checks for other fields, for example:
	if uuid, ok := data["uuid"].(string); ok {
		bone.UUID = uuid
	}

	// Add checks for other fields...

	// For nested fields like Children (which is []any), you may need additional logic to handle types
	if children, ok := data["children"].([]interface{}); ok {
		// Example handling: assuming children are either strings or Bone objects
		for _, child := range children {
			switch v := child.(type) {
			case string:
				bone.Children = append(bone.Children, v)
			case map[string]interface{}:
				// Recursively convert map to Bone
				childBone, err := MapToBone(v)
				if err != nil {
					return nil, err
				}
				bone.Children = append(bone.Children, childBone)
			default:
				return nil, errors.New("invalid child type")
			}
		}
	}

	return bone, nil
}
