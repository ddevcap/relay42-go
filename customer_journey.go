package relay42

import (
	"fmt"
	"net/http"
)

// CustomerJourneyService holds the R42 service
type CustomerJourneyService service

// GetCurrentJourneyStep returns a profile journey step by journeyId and profileId
func (service *ContentService) GetCurrentJourneyStep(journeyId, profileId string) (*JourneyStep, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s/%s", service.r.siteId, journeyId, profileId)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return &JourneyStep{}, err
	}

	var journeyStep *JourneyStep
	err = service.r.do(req, &journeyStep)

	return journeyStep, err
}

// GetJourneySteps returns profile journey steps by profileId
func (service *ContentService) GetJourneySteps(profileId string) ([]*JourneyStep, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s", service.r.siteId, profileId)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return []*JourneyStep{}, err
	}

	var journeySteps []*JourneyStep
	err = service.r.do(req, &journeySteps)

	return journeySteps, err
}

// RemoveProfileFromJourney removes a profile from a journey step by journeyId and profileId
func (service *CustomerJourneyService) RemoveProfileFromJourney(journeyId, profileId string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s/%s", service.r.siteId, journeyId, profileId)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}

// RemoveProfileFromJourneys removes a profile from all journeys by profileId
func (service *CustomerJourneyService) RemoveProfileFromJourneys(profileId string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s", service.r.siteId, profileId)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}
