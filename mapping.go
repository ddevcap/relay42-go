package relay42

const (
	MergeTypeNoMerge 			= "NO_MERGE"
	MergeTypeSite				= "SITE"
	MergeTypeMergeNoOverride 	= "MERGE_NO_OVERRIDE"
)

type Mapping struct {
	PartnerId 		[]string	`json:"partnerId"`
	PartnerType 	int			`json:"partnerType"`
}