package bbformat

type Animation struct {
	UUID           string              `json:"uuid"`
	Name           string              `json:"name"`
	Loop           string              `json:"loop"`
	Override       bool                `json:"override"`
	Length         float64             `json:"length"`
	Snapping       int                 `json:"snapping"`
	Selected       bool                `json:"selected"`
	AnimTimeUpdate string              `json:"anim_time_update"`
	BlendWeight    string              `json:"blend_weight"`
	StartDelay     string              `json:"start_delay"`
	LoopDelay      string              `json:"loop_delay"`
	Animators      map[string]Animator `json:"animators"`
}

type Animator struct {
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Keyframes []Keyframe `json:"keyframes"`
}

type Keyframe struct {
	Channel       string           `json:"channel"`
	DataPoints    []map[string]any `json:"data_points"`
	UUID          string           `json:"uuid"`
	Time          float32          `json:"time"`
	Color         int              `json:"color"`
	Interpolation string           `json:"interpolation"`
}
