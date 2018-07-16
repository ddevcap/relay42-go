package relay42

// Placement holds the placement data
type Placement struct {
	Id     string `json:"placementId"`
	Number int    `json:"placementNumber"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
