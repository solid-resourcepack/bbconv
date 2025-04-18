package mcformat

import (
	"fmt"
	"github.com/solid-resourcepack/bbconv/baseformat"
	"github.com/solid-resourcepack/bbconv/bbformat"
	"math"
	"strconv"
)

func ConvertElements(elements []baseformat.Element, model baseformat.Model) ([]Element, error) {
	result := make([]Element, 0, len(elements))
	for _, element := range elements {
		var rotation *Rotation
		if element.Rotation != nil {
			rotation = &Rotation{
				Origin:  element.Rotation.Origin.Slice(),
				Axis:    element.Rotation.Axis,
				Angle:   element.Rotation.Angle,
				Rescale: element.Rotation.Rescale,
			}
		}
		result = append(result, Element{
			From:     element.From.Slice(),
			To:       element.To.Slice(),
			Rotation: rotation,
			Faces:    convertFaces(element.Faces, model),
		})
	}
	return result, nil
}

func convertFaces(faces map[bbformat.Facing]bbformat.Face, model baseformat.Model) map[bbformat.Facing]Face {
	facing := make(map[bbformat.Facing]Face)
	for key, value := range faces {
		if value.Texture == nil {
			continue
		}
		facing[key] = Face{
			UV:      mapUv(value.UV, model),
			Texture: fmt.Sprintf("#%s", strconv.Itoa(*value.Texture)),
		}
	}
	return facing
}

func mapUv(UV []float32, model baseformat.Model) []float32 {
	mappedUv := UV
	resU := 16.0 / float32(model.Resolution.Width)
	resV := 16.0 / float32(model.Resolution.Height)
	for i, uv := range UV {
		if math.Mod(float64(i), 2) == 0 {
			mappedUv[i] = uv * resU
		} else {
			mappedUv[i] = uv * resV
		}

	}
	return mappedUv
}
