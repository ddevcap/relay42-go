package relay42

import (
	"net/http"
	"fmt"
)

type CustomerJourneyService service

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

func (service *CustomerJourneyService) RemoveProfileFromJourney(journeyId, profileId string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s/%s", service.r.siteId, journeyId, profileId)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}

func (service *CustomerJourneyService) RemoveProfileFromJourneys(profileId string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/v1/site-%d/customer/journey/%s", service.r.siteId, profileId)

	req, err := service.r.newRequest(method, path, nil, nil)
	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}