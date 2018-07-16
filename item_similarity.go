package relay42

// ItemSimilarity hold item similarity data
type ItemSimilarity struct {
	ID               string        `json:"itemId"`
	ItemSimilarities []SimilarItem `json:"itemSimilarities"`
}
