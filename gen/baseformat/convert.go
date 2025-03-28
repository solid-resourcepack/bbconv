package baseformat

import (
	"errors"
	"github.com/solid-resourcepack/bbconv/bbformat"
	"log"
	"strings"
)

func BBToBase(bbModel *bbformat.Model, namespace string) *Model {
	name, err := ConvertName(bbModel)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	baseId, err := convertBaseId(bbModel, namespace)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	bones, err := ConvertBones(bbModel.Outliner, bbModel.Elements, *baseId)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	textures, err := ConvertTextures(bbModel.Textures)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	animations, err := ConvertAnimations(bones, bbModel.Animations)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Model{
		*name,
		bones,
		animations,
		textures,
	}
}

func convertBaseId(bbModel *bbformat.Model, namespace string) (*string, error) {
	name, err := ConvertName(bbModel)
	if err != nil {
		return nil, err
	}
	result := namespace + ":item/" + *name + "_"
	return &result, nil
}

func ConvertName(bbModel *bbformat.Model) (*string, error) {
	key := ConvertToKey(bbModel.ModelIdentifier)
	if len(key) == 0 {
		key = ConvertToKey(bbModel.Name)
	}
	if len(key) == 0 {
		return nil, errors.New("no model name provided")
	}
	return &key, nil
}

func ConvertToKey(s string) string {
	return strings.Replace(strings.ToLower(strings.TrimSpace(s)), " ", "_", -1)
}
