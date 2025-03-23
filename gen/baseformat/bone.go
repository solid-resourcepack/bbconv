package baseformat

type Bone struct {
	Id       string    `json:"id"`
	Key      string    `json:"model,omitempty"`
	Children []Bone    `json:"children,omitempty"`
	Visuals  []Element `json:"-"`
}
