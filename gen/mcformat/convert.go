package mcformat

import (
	"github.com/solid-resourcepack/bbconv/baseformat"
	"log"
)

func BaseToMc(model *baseformat.Model, namespace string) (MCPackData, error) {
	textures := ConvertTextures(model)
	models := make(map[string]Model, 0)
	links := make([]Link, 0)
	for _, bone := range model.BoneTree {
		parsedModels, parsedLinks, err := appendBoneTree(bone, *model, namespace)
		if err != nil {
			log.Println(err)
			continue
		}
		for key, value := range parsedModels {
			models[key] = value
		}
		links = append(parsedLinks, parsedLinks...)
	}
	return MCPackData{
		ModelName: model.Name,
		Models:    models,
		Links:     links,
		Textures:  textures,
	}, nil
}

func appendBoneTree(bone baseformat.Bone, baseModel baseformat.Model, namespace string) (map[string]Model, []Link, error) {
	models := make(map[string]Model)
	links := make([]Link, 0)
	model, link, err := parseBone(bone, baseModel, namespace)
	if err != nil {
		return nil, nil, err
	}
	if model != nil && link != nil {
		models[bone.Key] = *model
		links = append(links, *link)
	}
	for _, bones := range bone.Children {
		childModels, childLinks, err := appendBoneTree(bones, baseModel, namespace)
		if err != nil {
			return nil, nil, err
		}
		for key, value := range childModels {
			models[key] = value
		}
		links = append(links, childLinks...)
	}
	return models, links, nil
}

func parseBone(bone baseformat.Bone, model baseformat.Model, namespace string) (*Model, *Link, error) {
	if !bone.Visible || len(bone.Visuals) == 0 {
		return nil, nil, nil
	}
	return ConvertBone(&bone, model, namespace)
}
