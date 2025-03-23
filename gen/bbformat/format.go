package bbformat

type Model struct {
	Meta            Meta        `json:"meta"`
	Name            string      `json:"name"`
	ModelIdentifier string      `json:"model_identifier"`
	Resolution      Resolution  `json:"resolution"`
	Elements        []Element   `json:"elements"`
	Animations      []Animation `json:"animations"`
	Outliner        []Outliner  `json:"outliner"`
	Textures        []Texture   `json:"textures"`
}

type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Meta struct {
	FormatVersion string `json:"format_version"`
	ModelFormat   string `json:"model_format"`
	BoxUv         bool   `json:"box_uv"`
}
