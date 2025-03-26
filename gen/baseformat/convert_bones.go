package baseformat

import (
	"errors"
	"github.com/solid-resourcepack/bbconv/bbformat"
	"log"
)

func SingularBoneTree(bone *bbformat.Bone, elements []bbformat.Element, baseId string) (*Bone, error) {
	parentBone := &Bone{
		Id:       bone.Name,
		Key:      baseId + ConvertToKey(bone.Name),
		Origin:   *toVec(bone.Origin),
		UUID:     bone.UUID,
		Visible:  bone.Visibility,
		Children: []Bone{},
	}

	for _, child := range bone.Children {
		if (child.Ref != nil) == (child.Bone != nil) {
			return nil, errors.New("child is both ref and bone or nothing of both")
		}
		if child.Ref != nil {
			if err := appendElement(*child.Ref, parentBone, bone, elements); err != nil {
				return nil, err
			}
		}
		if child.Bone != nil {
			if err := appendPartTree(parentBone, child.Bone, elements, baseId); err != nil {
				return nil, err
			}
		}
	}
	ResizeVisuals(parentBone)
	return parentBone, nil
}

func appendPartTree(parentBone *Bone, bone *bbformat.Bone, elements []bbformat.Element, baseId string) error {
	subTree, err := SingularBoneTree(bone, elements, baseId)
	if err != nil {
		return err
	}
	parentBone.Children = append(parentBone.Children, *subTree)
	return nil
}

func ConvertBones(outliners []bbformat.Bone, elements []bbformat.Element, baseId string) ([]Bone, error) {
	if len(outliners) < 1 {
		return nil, errors.New("not at least one outliner provided")
	}
	result := make([]Bone, 0, len(outliners))
	for _, bone := range outliners {
		if !bone.Visibility {
			continue
		}
		tree, err := SingularBoneTree(&bone, elements, baseId)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		result = append(result, *tree)
	}
	return result, nil
}

func findElement(elementId string, elements []bbformat.Element) (*bbformat.Element, bool) {
	for _, element := range elements {
		if element.UUID == elementId {
			return &element, true
		}
	}
	return nil, false
}

func appendElement(id string, parentBone *Bone, bbParentBone *bbformat.Bone, elements []bbformat.Element) error {
	if !parentBone.Visible {
		return nil
	}
	element, ok := findElement(id, elements)
	if !ok {
		return errors.New("child element not found")
	}
	parentBone.Visuals = append(parentBone.Visuals, ConvertElement(*bbParentBone, *element))
	return nil
}
