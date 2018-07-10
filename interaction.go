package relay42

import (
	"encoding/json"
	"golang.org/x/net/context"
)

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

type InteractionStreamHandlerFunc func(interaction *Interaction, cancel context.CancelFunc)

type Interaction struct {
	General
	InteractionJSON json.RawMessage
}

func (i *Interaction) UnmarshalJSON(b []byte) error {
	g := &General{}
	err := json.Unmarshal(b, g)

	i.InteractionType = g.InteractionType
	i.EventId = g.EventId
	i.SiteNumber = g.SiteNumber
	i.Timestamp = g.Timestamp
	i.TrackId = g.TrackId
	i.InteractionJSON = b

	return err
}

func (i *Interaction) ToBannerView() *BannerView {
	e := &BannerView{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

func (i *Interaction) ToConversion() *Conversion {
	e := &Conversion{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

func (i *Interaction) ToEngagement() *Engagement {
	e := &Engagement{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

func (i *Interaction) ToExperimentMatch() *ExperimentMatch {
	e := &ExperimentMatch{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

func (i *Interaction) ToPageView() *PageView {
	e := &PageView{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

func (i *Interaction) ToSessionStart() *SessionStart {
	e := &SessionStart{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

func (i *Interaction) ToUserPreferences() *UserPreferences {
	e := &UserPreferences{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

func (i *Interaction) ToSync() *Sync {
	e := &Sync{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}
func (i *Interaction) ToExternalFact() *ExternalFact {
	e := &ExternalFact{}
	json.Unmarshal(i.InteractionJSON, e)
	return e
}

type General struct {
	InteractionType string `json:"interactionType"`
	Timestamp       int64  `json:"timestamp"`
	TrackId         string `json:"trackId"`
	SiteNumber      int    `json:"siteNumber"`
	EventId         string `json:"eventId"`
}

type BannerView struct {
	General
	Identifier    string `json:"identifier"`
	SubIdentifier string `json:"subIdentifier"`
	Referral      string `json:"referral"`
}

type Conversion struct {
	General
	TransactionId string      `json:"transactionId"`
	Value         float64     `json:"value"`
	Products      interface{} `json:"products"`
	Variables     interface{} `json:"variables"`
}

type Engagement struct {
	General
	Type      string      `json:"type"`
	Content   interface{} `json:"content"`
	Variables interface{} `json:"variables"`
}

type ExperimentMatch struct {
	General
	ExperimentGroupNumber int `json:"experimentGroupNumber"`
	ExperimentNumber      int `json:"experimentNumber"`
}

type PageView struct {
	General
	Url       string `json:"url"`
	Source    string `json:"source"`
	UserAgent string `json:"userAgent"`
}

type SessionStart struct {
	General
	Url       string `json:"url"`
	Source    string `json:"source"`
	UserAgent string `json:"userAgent"`
}

type UserPreferences struct {
	General
	OptOutRemarketing bool `json:"optOutRemarketing"`
	OptOutAdapting    bool `json:"optOutAdapting"`
}

type Sync struct {
	General
	PartnerNumber int    `json:"partnerNumber"`
	PartnerCookie string `json:"partnerCookie"`
	MergeType     string `json:"mergeType"`
}

type ExternalFact struct {
	General
	Type          string      `json:"type"`
	Variables     interface{} `json:"variables"`
	TTL           int         `json:"ttl"`
	ForceInsert   bool        `json:"forceInsert"`
	OperationType string      `json:"operationType"`
}
