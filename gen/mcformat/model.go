package mcformat

type Model struct {
	Textures map[string]string `json:"textures,omitempty"`
	Elements []Element         `json:"elements"`
}
