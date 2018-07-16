package relay42

import "encoding/json"

// Fact hold the fact data
type Fact struct {
	CreationTime int         `json:"creation_time"`
	Name         string      `json:"name"`
	TTL          int         `json:"original_ttl"`
	Properties   interface{} `json:"parameters"`
}

// MarshalJSON for custom json marshaling
func (f Fact) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name       string      `json:"factName"`
		TTL        int         `json:"factTtl"`
		Properties interface{} `json:"properties"`
	}{
		Name:       f.Name,
		TTL:        f.TTL,
		Properties: f.Properties,
	})
}
