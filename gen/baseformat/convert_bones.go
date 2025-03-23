package baseformat

import (
	"errors"
	"fmt"
	"github.com/solid-resourcepack/bbconv/bbformat"
)

func SingularBoneTree(bone *bbformat.Bone, elements []bbformat.Element) (*Bone, error) {
	parentBone := &Bone{
		Id:       bone.Name,
		Children: []Bone{},
	}

	for _, child := range bone.Children {
		switch child := child.(type) {
		case string:
			if err := appendElement(child, parentBone, elements); err != nil {
				return nil, err
			}
		case *bbformat.Bone:
			if err := appendPartTree(parentBone, child, elements); err != nil {
				return nil, err
			}
		// TODO: do we really need this?
		case map[string]interface{}:
			parsed, err := bbformat.MapToBone(child)
			if err != nil {
				return nil, err
			}
			if err := appendPartTree(parentBone, parsed, elements); err != nil {
				return nil, err
			}
		default:
			fmt.Printf("Not expected Type: %T\n", child)
			return nil, errors.New("child is not a Bone or string reference")
		}
	}
	return parentBone, nil
}

func appendPartTree(parentBone *Bone, bone *bbformat.Bone, elements []bbformat.Element) error {
	subTree, err := SingularBoneTree(bone, elements)
	if err != nil {
		return err
	}
	newBone := Bone{
		Id:       subTree.Id,
		Children: subTree.Children,
		Visuals:  subTree.Visuals,
	}
	parentBone.Children = append(parentBone.Children, newBone)
	return nil
}

func ConvertBones(outliners []bbformat.Bone, elements []bbformat.Element) ([]Bone, error) {
	if len(outliners) < 1 {
		return nil, errors.New("not at least one outliner provided")
	}
	result := make([]Bone, 0, len(outliners))
	for _, bone := range outliners {
		tree, err := SingularBoneTree(&bone, elements)
		if err != nil {
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

func appendElement(id string, parentBone *Bone, elements []bbformat.Element) error {
	element, ok := findElement(id, elements)
	if !ok {
		return errors.New("child element not found")
	}
	parentBone.Visuals = append(parentBone.Visuals, Element{
		From: toVec(element.From),
		To:   toVec(element.To),
	})
	return nil
}
