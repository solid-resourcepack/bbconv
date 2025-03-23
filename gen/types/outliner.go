package types

type Outliner struct {
	Name       string    `json:"name"`
	Origin     []float32 `json:"origin"`
	Color      int       `json:"color"`
	UUID       string    `json:"uuid"`
	Export     bool      `json:"export"`
	MirrorUv   bool      `json:"mirror_uv"`
	IsOpen     bool      `json:"isOpen"`
	Locked     bool      `json:"locked"`
	Visibility bool      `json:"visibility"`
	Autouv     int       `json:"autouv"`
	Selected   bool      `json:"selected"`
	Children   []any     `json:"children"` //TODO: Use children or string
}

type Child struct {
	Name       string    `json:"name"`
	Origin     []float32 `json:"origin"`
	Color      int       `json:"color"`
	UUID       string    `json:"uuid"`
	Export     bool      `json:"export"`
	MirrorUv   bool      `json:"mirror_uv"`
	IsOpen     bool      `json:"isOpen"`
	Locked     bool      `json:"locked"`
	Visibility bool      `json:"visibility"`
	Autouv     int       `json:"autouv"`
	Selected   bool      `json:"selected"`
	Children   []any     `json:"children"` //TODO: Use children or string
}
