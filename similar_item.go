package relay42

// SimilarItem holds similar item data
type SimilarItem struct {
	Id         string `json:"similarItemId"`
	Similarity int    `json:"similarity"`
}
