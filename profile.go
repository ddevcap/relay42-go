package relay42

import (
	"fmt"
	"net/url"
	"net/http"
	"encoding/json"
	"bytes"
	"context"
)

type ProfileService service

func (service *ProfileService) StreamInteractions(iq string, ish InteractionStreamHandlerFunc) error {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/profiles/interactions/stream", service.r.siteId)
	query := url.Values{}
	query.Set("query", iq)

	req, err := service.r.newRequest(method, path, query, nil)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	req = req.WithContext(ctx)

	if err != nil {
		return err
	}

	err = service.r.doStream(req, func(b []byte){
		interaction := &Interaction{}
		json.Unmarshal(b, interaction)
		ish(interaction, cancel)
	})

	return err
}

func (service *ProfileService) StreamPartnerInteractions(iq string, ish InteractionStreamHandlerFunc) error {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/profiles/partners/stream", service.r.siteId)
	query := url.Values{}
	query.Set("query", iq)

	req, err := service.r.newRequest(method, path, query, nil)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	req = req.WithContext(ctx)

	if err != nil {
		return err
	}

	err = service.r.doStream(req, func(b []byte){
		interaction := &Interaction{}
		json.Unmarshal(b, interaction)
		ish(interaction, cancel)
	})

	return err
}

func (service *ProfileService) AddInteractions(pt, pid string, ensureProfile bool, interactions ...Interaction) (string, error)  {
	method := http.MethodPut

	if ensureProfile {
		method = http.MethodPost
	}

	path := fmt.Sprintf("/v1/site-%d/profiles/%s/facts", service.r.siteId, pt)
	query := url.Values{}
	query.Set("partnerId", pid)

	body := map[string][]Interaction{
		"interactions":interactions,
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, err := service.r.newRequest(method, path, query, b)
	if err != nil {
		return pid, err
	}

	req.Header.Set("Content-Type", "application/json")

	return pid, service.r.do(req, nil)
}

func (service *ProfileService) DeleteMappings(pt, pid string, mappings ...*Mapping) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/v1/site-%d/profiles/%s/mappings", service.r.siteId, pt)
	query := url.Values{}
	query.Set("partnerId", pid)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(mappings)

	req, err := service.r.newRequest(method, path, query, b)
	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}

func (service *ProfileService) GetMappings(pt, pid string) ([]*Mapping, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/profiles/%s/mappings", service.r.siteId, pt)
	query := url.Values{}
	query.Set("partnerId", pid)

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return []*Mapping{}, err
	}

	var mappingMap map[int][]string
	var mappingSlice []*Mapping

	err = service.r.do(req, &mappingMap)
	if err != nil {
		return []*Mapping{}, err
	}

	for pt, pid := range mappingMap {
		mappingSlice = append(mappingSlice, &Mapping{
			PartnerId:pid,
			PartnerType:pt,
		})
	}

	return mappingSlice, err
}

func (service *ProfileService) GetPartnerMappings(pt, pid, mpt string) ([]string, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/profiles/%s/mappings/%s", service.r.siteId, pt, mpt)
	query := url.Values{}
	query.Set("partnerId", pid)

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return []string{}, err
	}

	var partnerSlice []string
	err = service.r.do(req, &partnerSlice)
	if err != nil {
		return []string{}, err
	}

	return partnerSlice, err
}

func (service *ProfileService) AddMappings(pt, pid, mergeType string, ensureProfile bool, mappings ...*Mapping) (string, error) {
	method := http.MethodPut

	if ensureProfile {
		method = http.MethodPost
	}

	path := fmt.Sprintf("/v1/site-%d/profiles/%s/mappings", service.r.siteId, pt)

	query := url.Values{}
	query.Set("partnerId", pid)
	query.Set("forceInsert", fmt.Sprintf("%t", ensureProfile))
	query.Set("mergeType", fmt.Sprintf("%t", mergeType))

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(mappings)

	req, err := service.r.newRequest(method, path, query, b)
	if err != nil {
		return pid, err
	}

	req.Header.Set("Content-Type", "application/json")

	return pid, service.r.do(req, nil)
}

func (service *ProfileService) DeleteFacts(pt, pid string, factsNames ...string) error {
	method := http.MethodDelete
	path := fmt.Sprintf("/v1/site-%d/profiles/%s/facts", service.r.siteId, pt)
	query := url.Values{}
	query.Set("partnerId", pid)

	for _, factName := range factsNames {
		query.Set("fact", factName)
	}

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return err
	}

	return service.r.do(req, nil)
}

func (service *ProfileService) GetProfileId(pt, pid string) ([]*Fact, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/profiles/%s/facts", service.r.siteId, pt)
	query := url.Values{}
	query.Set("partnerId", pid)

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return []*Fact{}, err
	}

	var facts []*Fact
	err = service.r.do(req, &facts)

	return facts, err
}

func (service *ProfileService) GetFacts(pt, pid string) ([]*Fact, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/profiles/%s/facts", service.r.siteId, pt)
	query := url.Values{}
	query.Set("partnerId", pid)

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return []*Fact{}, err
	}

	var facts []*Fact
	err = service.r.do(req, &facts)

	return facts, err
}

func (service *ProfileService) AddFacts(pt, pid string, ensureProfile bool, facts ...*Fact) (string, error) {
	method := http.MethodPut

	if ensureProfile {
		method = http.MethodPost
	}

	path := fmt.Sprintf("/v1/site-%d/profiles/%s/facts", service.r.siteId, pt)

	query := url.Values{}
	query.Set("partnerId", pid)
	query.Set("forceInsert", fmt.Sprintf("%t", ensureProfile))

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(facts)

	req, err := service.r.newRequest(method, path, query, b)
	if err != nil {
		return pid, err
	}

	req.Header.Set("Content-Type", "application/json")

	return pid, service.r.do(req, nil)
}

func (service *ProfileService) GetId(pt, pid string) (string, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/profiles/%s/profileId", service.r.siteId, pt)
	query := url.Values{}
	query.Set("partnerId", pid)

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return "", err
	}

	var id string
	err = service.r.do(req, &id)

	return id, err
}

func (service *ProfileService) GetSegments(pt, pid string) ([]*Segment, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/profiles/%s/segments", service.r.siteId, pt)
	query := url.Values{}
	query.Set("partnerId", pid)

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return []*Segment{}, err
	}

	var segments []*Segment
	err = service.r.do(req, &segments)

	return segments, err
}

func (service *ProfileService) GetSegment(pt, pid, segmentName string) (*Segment, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/profiles/%s/segments/%s", service.r.siteId, pt, segmentName)
	query := url.Values{}
	query.Set("partnerId", pid)

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return &Segment{}, err
	}

	var segment *Segment
	err = service.r.do(req, &segment)

	return segment, err
}

