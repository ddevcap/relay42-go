package relay42

import "context"

const(
	TypeBannerView 		= "bannerView"
	TypeTypeConversion 	= "typeConversion"
	TypeEngagements 	= "engagements"
	TypeExperimentMatch = "experimentMatch"
	TypePageView 		= "pageView"
	TypeSessionStart 	= "sessionStart"
	TypeUserPreferences = "userPreferences"
	TypeSync 			= "sync"
	TypeExternalFact 	= "externalFact"
)

type InteractionStreamHandlerFunc func(interaction *Interaction, cancel context.CancelFunc)

type Interaction struct {
	*General
	*BannerView
	*Conversion
	*Engagements
	*ExperimentMatch
	*PageView
	*SessionStart
	*Sync
	*ExternalFact
}

type General struct {
	InteractionType 		string		`json:"interactionType"`
	Timestamp 				int64		`json:"timestamp"`
	TrackId 				string		`json:"trackId"`
	SiteNumber 				int			`json:"siteNumber"`
	EventId 				string		`json:"eventId"`
}

type BannerView struct {
	*General
	Identifier 				string		`json:"identifier"`
	SubIdentifier 			string		`json:"subIdentifier"`
	Referral				string		`json:"referral"`
}

type Conversion struct {
	*General
	TransactionId			string		`json:"transactionId"`
	Value					float64		`json:"value"`
	Products				interface{}	`json:"products"`
	Variables				interface{}	`json:"variables"`
}

type Engagements struct {
	*General
	Type 					string		`json:"type"`
	Content 				interface{}	`json:"content"`
	Variables				interface{}	`json:"variables"`
}

type ExperimentMatch struct {
	*General
	ExperimentGroupNumber 	int			`json:"experimentGroupNumber"`
	ExperimentNumber		int			`json:"experimentNumber"`
}

type PageView struct {
	*General
	Url 					string		`json:"url"`
	Source 					string 		`json:"source"`
	UserAgent 				string		`json:"userAgent"`
}

type SessionStart struct {
	*General
	Url 					string		`json:"url"`
	Source 					string 		`json:"source"`
	UserAgent 				string		`json:"userAgent"`
}

type UserPreferences struct {
	*General
	OptOutRemarketing		bool		`json:"optOutRemarketing"`
	OptOutAdapting			bool		`json:"optOutAdapting"`
}

type Sync struct {
	*General
	PartnerNumber 			int			`json:"partnerNumber"`
	PartnerCookie 			string		`json:"partnerCookie"`
	MergeType 				string		`json:"mergeType"`
}

type ExternalFact struct {
	*General
	Type 					string		`json:"type"`
	Variables				interface{}	`json:"variables"`
	TTL 					int			`json:"ttl"`
	ForceInsert 			bool		`json:"forceInsert"`
	OperationType			string		`json:"operationType"`
}