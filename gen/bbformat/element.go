package bbformat

type Element struct {
	Name                string          `json:"name"`
	BoxUv               bool            `json:"box_uv"`
	Rescale             bool            `json:"rescale"`
	Locked              bool            `json:"locked"`
	LightEmission       int             `json:"light_emission"`
	RenderOrder         string          `json:"render_order"`
	AllowMirrorModeling bool            `json:"allow_mirror_modeling"`
	From                []float32       `json:"from"`
	To                  []float32       `json:"to"`
	Autouv              int             `json:"autouv"`
	Inflate             float32         `json:"inflate,omitempty"`
	Color               int             `json:"color"`
	Origin              []float32       `json:"origin"`
	Rotation            []float32       `json:"rotation"`
	Faces               map[Facing]Face `json:"faces"`
	Type                string          `json:"type"`
	UUID                string          `json:"uuid"`
}

func (e *Element) HasTextures() bool {
	for _, face := range e.Faces {
		if face.Texture != nil {
			return true
		}
	}
	return false
}

type Facing string

var (
	FacingNorth Facing = "north"
	FacingSouth Facing = "south"
	FacingEast  Facing = "east"
	FacingWest  Facing = "west"
	FacingDown  Facing = "down"
	FacingUp    Facing = "up"
)

type Face struct {
	UV       []float32 `json:"uv"`
	Rotation *float32  `json:"rotation,omitempty"`
	Texture  *int      `json:"texture,omitempty"`
}
