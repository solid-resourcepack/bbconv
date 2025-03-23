package types

type Element struct {
	Name                string    `json:"name"`
	BoxUv               bool      `json:"box_uv"`
	Rescale             bool      `json:"rescale"`
	Locked              bool      `json:"locked"`
	LightEmission       int       `json:"light_emission"`
	RenderOrder         string    `json:"render_order"`
	AllowMirrorModeling bool      `json:"allow_mirror_modeling"`
	From                []float32 `json:"from"`
	To                  []float32 `json:"to"`
	Autouv              int       `json:"autouv"`
	Color               int       `json:"color"`
	Origin              []float32 `json:"origin"`
	Faces               Faces     `json:"faces"`
	Type                string    `json:"type"`
	UUID                string    `json:"uuid"`
}

type Faces struct {
	North UV `json:"north"`
	East  UV `json:"east"`
	South UV `json:"south"`
	West  UV `json:"west"`
	Up    UV `json:"up"`
	Down  UV `json:"down"`
}

type UV struct {
	Uv       []int   `json:"uv"`
	Rotation float32 `json:"rotation,omitempty"`
	Texture  int     `json:"texture,omitempty"`
}
