package baseformat

import (
	"github.com/solid-resourcepack/bbconv/bbformat"
	"github.com/ungerik/go3d/float64/vec3"
)

var origin = vec3.T{8, 8, 8}

func ConvertElement(parent Bone, element bbformat.Element, scale float64) *Element {
	result := &Element{
		Faces: element.Faces,
	}
	if !element.HasTextures() {
		return nil
	}
	inflate := vec3.T{float64(element.Inflate), float64(element.Inflate), float64(element.Inflate)}
	elementOrigin := toVec(element.Origin)
	centerOrigin := centralize(*elementOrigin, parent.Origin, scale)
	identifier := identifier(toVec(element.Rotation))
	rotationOrigin := rotationOrigin(identifier, centerOrigin, parent.Origin, scale)

	// Convert From and To
	from := centralize(*toVec(element.From), parent.Origin, scale)
	from.Add(&rotationOrigin)
	from.Add(&origin)
	from.Sub(&inflate)
	result.From = from

	to := centralize(*toVec(element.To), parent.Origin, scale)
	to.Add(&rotationOrigin)
	to.Add(&origin)
	to.Add(&inflate)
	result.To = to

	// Convert Rotation
	var convertedRotation Rotation
	if element.Rotation != nil {
		rotation := toVec(element.Rotation)
		if rotation != nil {
			rotation.Sub(&identifier)
			if *rotation != vec3.Zero {
				convertedRotation = convertRotation(element)
				convertedRotationOrigin := centerOrigin.Added(&rotationOrigin)
				convertedRotationOrigin.Add(&origin)
				convertedRotation.Origin = convertedRotationOrigin
				result.Rotation = &convertedRotation
			}
		}
	}

	return result
}

func centralize(target vec3.T, groupOrigin vec3.T, scale float64) vec3.T {
	result := target.Subed(&groupOrigin)
	return result.Muled(&vec3.T{1 / scale, 1 / scale, 1 / scale})
}

func rotationOrigin(identifier vec3.T, rotationOrigin vec3.T, groupOrigin vec3.T, scale float64) vec3.T {
	centerOrigin := centralize(rotationOrigin, groupOrigin, scale)
	quaternion := ToQuaternion(&identifier)
	rotated := quaternion.RotatedVec3(&centerOrigin)
	return rotated.Subed(&centerOrigin)
}

func convertRotation(element bbformat.Element) Rotation {
	rotation := Rotation{
		Axis:    "",
		Angle:   0,
		Rescale: element.Rescale,
	}
	if element.Rotation[0] != 0 {
		rotation.Axis = "x"
		rotation.Angle = float64(element.Rotation[0])
	}
	if element.Rotation[1] != 0 {
		rotation.Axis = "y"
		rotation.Angle = float64(element.Rotation[1])
	}
	if element.Rotation[2] != 0 {
		rotation.Axis = "z"
		rotation.Angle = float64(element.Rotation[2])
	}
	return rotation
}

func toVec(points []float32) *vec3.T {
	if len(points) != 3 {
		return nil
	}
	return &vec3.T{
		float64(points[0]),
		float64(points[1]),
		float64(points[2]),
	}
}
