package baseformat

import (
	"errors"
	"fmt"
	"github.com/solid-resourcepack/bbconv/bbformat"
	"github.com/ungerik/go3d/float64/vec3"
	"strconv"
	"strings"
)

func ConvertAnimations(tree []Bone, animations []bbformat.Animation) ([]Animation, error) {
	result := make([]Animation, len(animations))
	for i, animation := range animations {
		animators, err := ConvertAnimators(tree, animation.Animators)
		if err != nil {
			return nil, err
		}
		result[i] = Animation{
			Loop:       false,
			Length:     animation.Length,
			StartDelay: 0,
			LoopDelay:  0,
			Name:       animation.Name,
			Animators:  animators,
		}
	}
	return result, nil
}

func ConvertAnimators(tree []Bone, animators map[string]bbformat.Animator) ([]Animator, error) {
	result := make([]Animator, 0, len(animators))
	for key, animator := range animators {
		bone := findBone(tree, key)
		if bone == nil {
			return nil, errors.New(fmt.Sprintf("bone %s not found", key))
		}
		position, rotation, scale, err := ConvertKeyframes(animator.Keyframes)
		if err != nil {
			return nil, err
		}
		result = append(result, Animator{
			Name:     animator.Name,
			Bone:     bone.Id,
			Position: position,
			Rotation: rotation,
			Scale:    scale,
		})
	}
	return result, nil
}

func findBone(tree []Bone, uuid string) *Bone {
	if len(tree) == 0 {
		return nil
	}
	for _, bone := range tree {
		if result := bone.FindBone(uuid); result != nil {
			return result
		}
	}
	return nil
}

func ConvertDataPoint(keyframe bbformat.Keyframe) (*vec3.T, error) {
	if len(keyframe.DataPoints) < 1 {
		return nil, errors.New("no dataPoints found")
	}
	x, err := toFloat(keyframe.DataPoints[0]["x"])
	if err != nil {
		return nil, err
	}
	y, err := toFloat(keyframe.DataPoints[0]["y"])
	if err != nil {
		return nil, err
	}
	z, err := toFloat(keyframe.DataPoints[0]["z"])
	if err != nil {
		return nil, err
	}
	return &vec3.T{
		*x,
		*y,
		*z,
	}, nil
}

func toFloat(data any) (*float64, error) {
	switch v := data.(type) {
	case string:
		if len(v) == 0 {
			return nil, errors.New("no data provided")
		}
		result, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return nil, err
		}
		return &result, nil
	case float64:
		return &v, nil
	default:
		return nil, errors.New("no data provided")
	}
}

func ConvertKeyframes(keyframes []bbformat.Keyframe) ([]PositionKeyframe, []RotationKeyframe, []ScaleKeyframe, error) {
	positionKeyframes := make([]PositionKeyframe, 0)
	rotationKeyframes := make([]RotationKeyframe, 0)
	scaleKeyframes := make([]ScaleKeyframe, 0)
	for _, keyframe := range keyframes {
		data, err := ConvertDataPoint(keyframe)
		if err != nil {
			continue
		}
		switch keyframe.Channel {
		case bbformat.KeyFrameTypeRotation:
			rotation := ToQuaternion(data)
			rotationKeyframes = append(rotationKeyframes, RotationKeyframe{
				Time:          keyframe.Time,
				LeftRotation:  Quaternion(rotation),
				RightRotation: Quaternion(rotation.Inverted()),
				Interpolation: keyframe.Interpolation,
			})
		case bbformat.KeyFrameTypePosition:
			positionKeyframes = append(positionKeyframes, PositionKeyframe{
				Time:          keyframe.Time,
				Interpolation: keyframe.Interpolation,
				Position:      Vector(*data),
			})
		case bbformat.KeyFrameTypeScale:
			scaleKeyframes = append(scaleKeyframes, ScaleKeyframe{
				Time:          keyframe.Time,
				Interpolation: keyframe.Interpolation,
				Scale:         Vector(*data),
			})
		}

	}
	return positionKeyframes, rotationKeyframes, scaleKeyframes, nil
}
