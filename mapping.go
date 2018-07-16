package relay42

// Const for all merge types
const (
	MergeTypeNoMerge         = "NO_MERGE"
	MergeTypeSite            = "SITE"
	MergeTypeMergeNoOverride = "MERGE_NO_OVERRIDE"
)

// Mapping holds mapping data
type Mapping struct {
	PartnerID   []string `json:"partnerId"`
	PartnerType int      `json:"partnerType"`
}
