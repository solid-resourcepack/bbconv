package mcformat

type Link struct {
	Model LinkModel `json:"model"`
}

type LinkModel struct {
	Type  string `json:"type"`
	Model string `json:"model"`
}
