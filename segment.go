package relay42

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// SegmentService holds the R42 service
type SegmentService service

// Segment holds segment data
type Segment struct {
	SiteNumber  int         `json:"siteNumber"`
	Number      int         `json:"segmentNumber"`
	SegmentName string      `json:"segmentName"`
	Parameters  interface{} `json:"parameters"`
}

// SegmentStreamItem holds segment stream item data
type SegmentStreamItem struct {
	Added    []interface{} `json:"added"`
	Removed  []interface{} `json:"removed"`
	Modified []interface{} `json:"modified"`
}

// SegmentStreamHandlerFunc defines the segment stream handler type
type SegmentStreamHandlerFunc func(segmentStreamItem *SegmentStreamItem, cancel context.CancelFunc)

// Stream creates a stream by partnerType
func (service *SegmentService) Stream(partnerType string, segmentStreamHandle SegmentStreamHandlerFunc) error {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/segments/stream", service.r.siteID)
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

// StreamSegment creates a segment stream by partnerType and segmentNumber
func (service *SegmentService) StreamSegment(partnerType, segmentNumber string, segmentStreamHandle SegmentStreamHandlerFunc) error {
	method := http.MethodGet
	path := fmt.Sprintf("/v1/site-%d/segments/stream/%d_%s", service.r.siteID, service.r.siteID, segmentNumber)
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
