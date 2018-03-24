package relay42

type Segment struct {
	SiteNumber 		int				`json:"siteNumber"`
	Number 			int 			`json:"segmentNumber"`
	SegmentName 	string 			`json:"segmentName"`
	Parameters 		interface{} 	`json:"parameters"`
}