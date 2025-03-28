package mcformat

import (
	"fmt"
	"github.com/solid-resourcepack/bbconv/baseformat"
	"github.com/solid-resourcepack/bbconv/bbformat"
	"strconv"
)

func ConvertElements(elements []baseformat.Element) ([]Element, error) {
	result := make([]Element, 0, len(elements))
	for _, element := range elements {
		var rotation *Rotation
		if element.Rotation != nil {
			rotation = &Rotation{
				Origin: baseformat.Vector{8, 8, 8},
				Axis:   "",
				Angle:  0,
			}
			if element.Rotation[0] != 0 {
				rotation.Axis = "x"
				rotation.Angle = element.Rotation[0]
			}
			if element.Rotation[1] != 0 {
				rotation.Axis = "y"
				rotation.Angle = element.Rotation[1]
			}
			if element.Rotation[2] != 0 {
				rotation.Axis = "z"
				rotation.Angle = element.Rotation[2]
			}
		}
		result = append(result, Element{
			From:     baseformat.Vector(element.From),
			To:       baseformat.Vector(element.To),
			Rotation: rotation,
			Faces:    convertFaces(element.Faces),
		})
	}
	return result, nil
}

func convertFaces(faces map[bbformat.Facing]bbformat.Face) map[bbformat.Facing]Face {
	facing := make(map[bbformat.Facing]Face)
	for key, value := range faces {
		facing[key] = Face{
			UV:      value.UV,
			Texture: fmt.Sprintf("#%s", strconv.Itoa(value.Texture)),
		}
	}
	return facing
}
