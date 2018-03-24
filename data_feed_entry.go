package relay42

type DataFeedEntry struct {
	Key		string 					`json:"key"`
	TTL		int 					`json:"ttl"`
	Values	map[string]interface{}	`json:"values"`
}