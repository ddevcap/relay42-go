package relay42

import (
	"fmt"
	"net/http"
)

// ContentService holds the R42 service
type ContentService service

// GetCampaigns returns content campaigns
func (service *ContentService) GetCampaigns() ([]*Campaign, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/content/external", service.r.siteId)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return []*Campaign{}, err
	}

	var campaigns []*Campaign
	err = service.r.do(req, &campaigns)

	return campaigns, err
}

// GetCampaign returns a content campaign by campaignId
func (service *ContentService) GetCampaign(campaignId string) (*Campaign, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/content/external/%s", service.r.siteId, campaignId)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return &Campaign{}, err
	}

	var campaign *Campaign
	err = service.r.do(req, &campaign)

	return campaign, err
}

// GetCampaignPlacements returns campaign placements by campaignId
func (service *ContentService) GetCampaignPlacements(campaignId string) ([]*Placement, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/content/external/%s/placements", service.r.siteId, campaignId)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return []*Placement{}, err
	}

	var placements []*Placement
	err = service.r.do(req, &placements)

	return placements, err
}
