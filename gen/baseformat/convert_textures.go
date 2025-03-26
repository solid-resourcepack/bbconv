package baseformat

import (
	"github.com/solid-resourcepack/bbconv/bbformat"
)

func ConvertTextures(textures []bbformat.Texture) (map[string]string, error) {
	result := make(map[string]string)
	for _, texture := range textures {
		result[texture.ID] = texture.Source
	}
	return result, nil
}
