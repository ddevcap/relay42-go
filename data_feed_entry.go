package relay42

// DataFeedEntry holds data feed entry data
type DataFeedEntry struct {
	Key    string                 `json:"key"`
	TTL    int                    `json:"ttl"`
	Values map[string]interface{} `json:"values"`
}
