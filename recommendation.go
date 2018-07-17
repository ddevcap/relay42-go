package relay42

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// RecommendationService holds the R42 service
type RecommendationService service

// GetSimilarItems returns similar items by profileNumber, prefix, attribute and itemIds
func (service *RecommendationService) GetSimilarItems(profileNumber, prefix, attribute string, resultSize int, itemIds ...string) ([]*SimilarItem, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/recommendations/%s/items", service.r.siteID, profileNumber)
	query := url.Values{}
	query.Set("prefix", prefix)
	query.Set("attribute", attribute)
	query.Set("itemIds", strings.Join(itemIds, ","))

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return []*SimilarItem{}, err
	}

	var similarItems []*SimilarItem
	err = service.r.do(req, &similarItems)
	if err != nil {
		return []*SimilarItem{}, err
	}

	return similarItems, err
}

// GetItemSimilarItems returns item similar items by profileNumber, prefix, attribute and itemID
func (service *RecommendationService) GetItemSimilarItems(profileNumber, prefix, attribute string, resultSize int, itemID string) ([]*SimilarItem, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/recommendations/%s/%s", service.r.siteID, profileNumber, itemID)
	query := url.Values{}
	query.Set("prefix", prefix)
	query.Set("attribute", attribute)

	req, err := service.r.newRequest(method, path, query, nil)
	if err != nil {
		return []*SimilarItem{}, err
	}

	var similarItems []*SimilarItem
	err = service.r.do(req, &similarItems)
	if err != nil {
		return []*SimilarItem{}, err
	}

	return similarItems, err
}

// AddOrUpdateSimilarItems adds or updates similar items by profileNumber
func (service *RecommendationService) AddOrUpdateSimilarItems(profileNumber string, itemSimilarities ...ItemSimilarity) error {
	method := http.MethodPost
	path := fmt.Sprintf("/v1/site-%d/recommendations/%s/items", service.r.siteID, profileNumber)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(itemSimilarities)

	req, err := service.r.newRequest(method, path, nil, b)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	return service.r.do(req, nil)
}
