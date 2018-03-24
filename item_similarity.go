package relay42

type ItemSimilarity struct {
	Id               string        `json:"itemId"`
	ItemSimilarities []SimilarItem `json:"itemSimilarities"`
}
