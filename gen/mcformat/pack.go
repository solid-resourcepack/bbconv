package mcformat

type MCPackData struct {
	ModelName string
	Models    map[string]Model
	Links     []Link
	Textures  map[int16]string
}
