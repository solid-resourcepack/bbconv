package mcformat

import (
	"github.com/solid-resourcepack/bbconv/baseformat"
	"github.com/solid-resourcepack/bbconv/bbformat"
	"strconv"
)

func ExtractTextures(bone *baseformat.Bone) []int16 {
	textures := make([]int16, 0)
	for _, element := range bone.Visuals {
		textures = appendUniqueAndPresent(element, bbformat.FacingNorth, textures)
		textures = appendUniqueAndPresent(element, bbformat.FacingSouth, textures)
		textures = appendUniqueAndPresent(element, bbformat.FacingEast, textures)
		textures = appendUniqueAndPresent(element, bbformat.FacingWest, textures)
		textures = appendUniqueAndPresent(element, bbformat.FacingUp, textures)
		textures = appendUniqueAndPresent(element, bbformat.FacingDown, textures)
	}
	return textures
}

func appendUniqueAndPresent(element baseformat.Element, facing bbformat.Facing, textures []int16) []int16 {
	if element.Faces[facing].Texture != nil {
		return appendUnique(textures, int16(*element.Faces[facing].Texture))
	}
	return textures
}

func ConvertTextures(model *baseformat.Model) map[int16]string {
	result := make(map[int16]string)
	for key, value := range model.Textures {
		parsed, err := strconv.ParseFloat(key, 16)
		if err != nil {
			continue
		}
		result[int16(parsed)] = value
	}
	return result
}

func appendUnique(slice []int16, elem int16) []int16 {
	for _, v := range slice {
		if v == elem {
			return slice
		}
	}
	return append(slice, elem)
}
