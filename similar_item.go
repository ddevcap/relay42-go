package relay42

// SimilarItem holds similar item data
type SimilarItem struct {
	ID         string `json:"similarItemId"`
	Similarity int    `json:"similarity"`
}
