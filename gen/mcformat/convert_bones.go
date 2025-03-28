package mcformat

import (
	"fmt"
	"github.com/solid-resourcepack/bbconv/baseformat"
	"strconv"
)

func ConvertBone(bone *baseformat.Bone, model baseformat.Model, namespace string) (*Model, *Link, error) {
	usedTextures := ExtractTextures(bone)
	modelTextureMap := make(map[string]string)
	for _, texture := range usedTextures {
		id := strconv.Itoa(int(texture))
		modelTextureMap[id] = fmt.Sprintf("%s:item/%s_%v", namespace, model.Name, id)
	}

	elements, err := ConvertElements(bone.Visuals)
	if err != nil {
		return nil, nil, err
	}
	return &Model{
			Textures: modelTextureMap,
			Elements: elements,
		}, &Link{
			Model: LinkModel{
				Type:  "minecraft:model",
				Model: bone.Key,
			},
		}, nil
}
