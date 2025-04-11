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
		Scale:    getScale(*bone, elements),
	}
	if bone.Rotation != nil {
		raw := toVec(bone.Rotation)
		display := toDisplay(*raw)
		quatRotation := ToQuaternion(&display)
		parentBone.LeftRotation = Quaternion(quatRotation)
		parentBone.RightRotation = Quaternion(quatRotation.Inverted())
	}

	for _, child := range bone.Children {
		if (child.Ref != nil) == (child.Bone != nil) {
			return nil, errors.New("child is both ref and bone or nothing of both")
		}
		if child.Ref != nil {
			if err := appendElement(*child.Ref, parentBone, elements); err != nil {
				return nil, err
			}
		}
		if child.Bone != nil {
			if err := appendPartTree(parentBone, child.Bone, elements, baseId); err != nil {
				return nil, err
			}
		}
	}
	if len(parentBone.Visuals) == 0 {
		parentBone.Visible = false
	}
	boneOrigin := toVec(bone.Origin).Scaled(1 / 16.0)
	parentBone.Origin = toDisplay(boneOrigin)
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

func appendElement(id string, parentBone *Bone, elements []bbformat.Element) error {
	element, ok := findElement(id, elements)
	if !ok {
		return errors.New("child element not found")
	}
	result := ConvertElement(*parentBone, *element, parentBone.Scale)
	if result == nil {
		log.Println("Will not display element", element.UUID, "on bone", parentBone.Id)
		return nil
	}
	parentBone.Visuals = append(parentBone.Visuals, *result)
	return nil
}

func getScale(bone bbformat.Bone, elements []bbformat.Element) float64 {
	result := 1.0
	for _, child := range bone.Children {
		if child.Ref != nil {
			element, ok := findElement(*child.Ref, elements)
			if !ok {
				continue
			}
			from := toVec(element.From)
			to := toVec(element.To)
			size := to.Sub(from).Length()
			if size > result {
				result = size
			}
		}
	}
	return result / 16.0
}
