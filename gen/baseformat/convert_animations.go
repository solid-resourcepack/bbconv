package baseformat

import "github.com/solid-resourcepack/bbconv/bbformat"

func ConvertAnimations(animations []bbformat.Animation) ([]Animation, error) {
	return make([]Animation, len(animations)), nil
}
