package mcformat

import (
	"github.com/solid-resourcepack/bbconv/baseformat"
	"github.com/solid-resourcepack/bbconv/bbformat"
	"strconv"
)

func ExtractTextures(bone *baseformat.Bone) []int16 {
	textures := make([]int16, 0)
	for _, element := range bone.Visuals {
		textures = appendIfNotExists(textures, int16(element.Faces[bbformat.FacingNorth].Texture))
		textures = appendIfNotExists(textures, int16(element.Faces[bbformat.FacingSouth].Texture))
		textures = appendIfNotExists(textures, int16(element.Faces[bbformat.FacingEast].Texture))
		textures = appendIfNotExists(textures, int16(element.Faces[bbformat.FacingWest].Texture))
		textures = appendIfNotExists(textures, int16(element.Faces[bbformat.FacingUp].Texture))
		textures = appendIfNotExists(textures, int16(element.Faces[bbformat.FacingDown].Texture))
	}
	return textures
}

func ConvertTextures(model *baseformat.Model) map[int16]string {
	result := map[int16]string{}
	for key, value := range model.Textures {
		parsed, err := strconv.ParseFloat(key, 16)
		if err != nil {
			continue
		}
		result[int16(parsed)] = value
	}
	return result
}

func appendIfNotExists(slice []int16, elem int16) []int16 {
	for _, v := range slice {
		if v == elem {
			return slice
		}
	}
	return append(slice, elem)
}
