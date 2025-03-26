package baseformat

import (
	"github.com/solid-resourcepack/bbconv/bbformat"
	"log"
)

func BBToBase(bbModel *bbformat.Model) *Model {
	bones, err := ConvertBones(bbModel.Outliner, bbModel.Elements)
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
		bbModel.Name,
		bones,
		animations,
		textures,
	}
}
