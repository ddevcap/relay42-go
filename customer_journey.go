package relay42

import (
	"fmt"
	"net/http"
)

// CustomerJourneyService holds the R42 service
type CustomerJourneyService service

// GetCurrentJourneyStep returns a profile journey step by journeyID and profileID
func (service *ContentService) GetCurrentJourneyStep(journeyID, profileID string) (*JourneyStep, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s/%s", service.r.siteId, journeyID, profileID)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return &JourneyStep{}, err
	}

	var journeyStep *JourneyStep
	err = service.r.do(req, &journeyStep)

	return journeyStep, err
}

// GetJourneySteps returns profile journey steps by profileID
func (service *ContentService) GetJourneySteps(profileID string) ([]*JourneyStep, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s", service.r.siteId, profileID)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return []*JourneyStep{}, err
	}

	var journeySteps []*JourneyStep
	err = service.r.do(req, &journeySteps)

	return journeySteps, err
}

// RemoveProfileFromJourney removes a profile from a journey step by journeyID and profileID
func (service *CustomerJourneyService) RemoveProfileFromJourney(journeyID, profileID string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s/%s", service.r.siteId, journeyID, profileID)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}

// RemoveProfileFromJourneys removes a profile from all journeys by profileID
func (service *CustomerJourneyService) RemoveProfileFromJourneys(profileID string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s", service.r.siteId, profileID)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}
