package relay42

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SegmentService service

type Segment struct {
	SiteNumber  int         `json:"siteNumber"`
	Number      int         `json:"segmentNumber"`
	SegmentName string      `json:"segmentName"`
	Parameters  interface{} `json:"parameters"`
}

type SegmentStreamItem struct {
	Added    []interface{} `json:"added"`
	Removed  []interface{} `json:"removed"`
	Modified []interface{} `json:"modified"`
}

type SegmentStreamHandlerFunc func(segmentStreamItem *SegmentStreamItem, cancel context.CancelFunc)

func (service *SegmentService) Stream(partnerType string, segmentStreamHandle SegmentStreamHandlerFunc) error {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/segments/stream", service.r.siteId)
	query := url.Values{}
	query.Set("partnerType", partnerType)

	req, err := service.r.newRequest(method, path, query, nil)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	req = req.WithContext(ctx)

	if err != nil {
		return err
	}

	err = service.r.doStream(req, func(b []byte) {
		streamItem := &SegmentStreamItem{}
		json.Unmarshal(b, streamItem)
		segmentStreamHandle(streamItem, cancel)
	})

	return err
}

func (service *SegmentService) StreamSegment(partnerType, segmentNumber string, segmentStreamHandle SegmentStreamHandlerFunc) error {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/segments/stream/%s_%s", service.r.siteId, service.r.siteId, segmentNumber)
	query := url.Values{}
	query.Set("partnerType", partnerType)

	req, err := service.r.newRequest(method, path, query, nil)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	req = req.WithContext(ctx)

	if err != nil {
		return err
	}

	err = service.r.doStream(req, func(b []byte) {
		streamItem := &SegmentStreamItem{}
		json.Unmarshal(b, streamItem)
		segmentStreamHandle(streamItem, cancel)
	})

	return err
}
