package relay42

import (
	"encoding/json"
	"golang.org/x/net/context"
)

// Const for all interaction types
const (
	TypeBannerView      = "bannerView"
	TypeTypeConversion  = "typeConversion"
	TypeEngagement      = "engagement"
	TypeExperimentMatch = "experimentMatch"
	TypePageView        = "pageView"
	TypeSessionStart    = "sessionStart"
	TypeUserPreferences = "userPreferences"
	TypeSync            = "sync"
	TypeExternalFact    = "externalFact"
)

// InteractionStreamHandlerFunc defines the interaction handler type
type InteractionStreamHandlerFunc func(interaction *Interaction, cancel context.CancelFunc)

// Interaction holds the interaction general data and raw json
type Interaction struct {
	General
	InteractionJSON json.RawMessage
}

// UnmarshalJSON for custom json marshaling
func (i *Interaction) UnmarshalJSON(b []byte) error {
	g := &General{}
	err := json.Unmarshal(b, g)

	i.InteractionType = g.InteractionType
	i.EventID = g.EventID
	i.SiteNumber = g.SiteNumber
	i.Timestamp = g.Timestamp
	i.TrackID = g.TrackID
	i.InteractionJSON = b

	return err
}

// ToBannerView returns a BannerView from an interaction
func (i *Interaction) ToBannerView() *BannerView {
	e := &BannerView{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

// ToConversion returns a Conversion from an interaction
func (i *Interaction) ToConversion() *Conversion {
	e := &Conversion{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

// ToEngagement returns an Engagement from an interaction
func (i *Interaction) ToEngagement() *Engagement {
	e := &Engagement{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

// ToExperimentMatch returns an ExperimentMatch from an interaction
func (i *Interaction) ToExperimentMatch() *ExperimentMatch {
	e := &ExperimentMatch{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

// ToPageView returns a PageView from an interaction
func (i *Interaction) ToPageView() *PageView {
	e := &PageView{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

// ToSessionStart returns a SessionStart from an interaction
func (i *Interaction) ToSessionStart() *SessionStart {
	e := &SessionStart{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

// ToUserPreferences returns a UserPreferences from an interaction
func (i *Interaction) ToUserPreferences() *UserPreferences {
	e := &UserPreferences{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

// ToSync returns a Sync from an interaction
func (i *Interaction) ToSync() *Sync {
	e := &Sync{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

// ToExternalFact returns an ExternalFact from an interaction
func (i *Interaction) ToExternalFact() *ExternalFact {
	e := &ExternalFact{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

// General holds interaction general data
type General struct {
	InteractionType string `json:"interactionType"`
	Timestamp       int64  `json:"timestamp"`
	TrackID         string `json:"trackId"`
	SiteNumber      int    `json:"siteNumber"`
	EventID         string `json:"eventId"`
}

// BannerView holds banner view data
type BannerView struct {
	General
	Identifier    string `json:"identifier"`
	SubIdentifier string `json:"subIdentifier"`
	Referral      string `json:"referral"`
}

// Conversion holds conversion data
type Conversion struct {
	General
	TransactionID string      `json:"transactionId"`
	Value         float64     `json:"value"`
	Products      interface{} `json:"products"`
	Variables     interface{} `json:"variables"`
}

// Engagement holds engagement data
type Engagement struct {
	General
	Type      string      `json:"type"`
	Content   interface{} `json:"content"`
	Variables interface{} `json:"variables"`
}

// ExperimentMatch holds experiment match data
type ExperimentMatch struct {
	General
	ExperimentGroupNumber int `json:"experimentGroupNumber"`
	ExperimentNumber      int `json:"experimentNumber"`
}

// PageView page view view data
type PageView struct {
	General
	URL       string `json:"url"`
	Source    string `json:"source"`
	UserAgent string `json:"userAgent"`
}

// SessionStart session start data
type SessionStart struct {
	General
	URL       string `json:"url"`
	Source    string `json:"source"`
	UserAgent string `json:"userAgent"`
}

// UserPreferences holds user preferences data
type UserPreferences struct {
	General
	OptOutRemarketing bool `json:"optOutRemarketing"`
	OptOutAdapting    bool `json:"optOutAdapting"`
}

// Sync holds sync data
type Sync struct {
	General
	PartnerNumber int    `json:"partnerNumber"`
	PartnerCookie string `json:"partnerCookie"`
	MergeType     string `json:"mergeType"`
}

// ExternalFact holds external fact data
type ExternalFact struct {
	General
	Type          string      `json:"type"`
	Variables     interface{} `json:"variables"`
	TTL           int         `json:"ttl"`
	ForceInsert   bool        `json:"forceInsert"`
	OperationType string      `json:"operationType"`
}
