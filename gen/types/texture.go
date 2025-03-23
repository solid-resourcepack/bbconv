package types

type Texture struct {
	Path             string `json:"path"`
	Name             string `json:"name"`
	Folder           string `json:"folder"`
	Namespace        string `json:"namespace"`
	ID               string `json:"id"`
	Group            string `json:"group"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	UvWidth          int    `json:"uv_width"`
	UvHeight         int    `json:"uv_height"`
	Particle         bool   `json:"particle"`
	UseAsDefault     bool   `json:"use_as_default"`
	LayersEnabled    bool   `json:"layers_enabled"`
	SyncToProject    string `json:"sync_to_project"`
	RenderMode       string `json:"render_mode"`
	RenderSides      string `json:"render_sides"`
	FrameTime        int    `json:"frame_time"`
	FrameOrderType   string `json:"frame_order_type"`
	FrameOrder       string `json:"frame_order"`
	FrameInterpolate bool   `json:"frame_interpolate"`
	Visible          bool   `json:"visible"`
	Internal         bool   `json:"internal"`
	Saved            bool   `json:"saved"`
	UUID             string `json:"uuid"`
	Source           string `json:"source"`
}
