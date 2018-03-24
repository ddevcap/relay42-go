package relay42

import (
	"io"
	"fmt"
	"log"
	"net/url"
	"net/http"
	"encoding/json"
	"github.com/moul/http2curl"
	"net/http/httputil"
	"bufio"
	"bytes"
)

type Relay42 struct {
	client      	*http.Client
	username		string
	password 		string
	siteId			int
	Debug       	bool
	QueryParams 	map[string]string
	Headers     	map[string]string
	BaseURL     	string

	Content			*ContentService
	DataFeed		*DataFeedService
	Profile			*ProfileService
	CustomerJourney *CustomerJourneyService
}

type service struct {
	r *Relay42
}

func NewClient(username, password string) *Relay42 {
	r := &Relay42{
		client:http.DefaultClient,
		username: username,
		password: password,
		BaseURL: "https://api.relay42.com:443/",
	}

	r.Content 			= &ContentService{r: r}
	r.DataFeed			= &DataFeedService{r: r}
	r.Profile 			= &ProfileService{r: r}
	r.CustomerJourney 	= &CustomerJourneyService{r: r}

	return r
}

func (r *Relay42) Site(siteId int) {
	r.siteId = siteId
}

func (r *Relay42) newRequest(method, path string, query url.Values, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(r.BaseURL)

	if err != nil {
		return nil, err
	}

	for key, value := range r.QueryParams {
		query.Set(key, value)
	}

	u.Path = path
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(r.username, r.password)

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	return req, nil
}


func (r *Relay42) do(req *http.Request, v interface{}) error {
	if r.Debug == true {
		command, _ := http2curl.GetCurlCommand(req)
		fmt.Println(command)
	}

	res, err := r.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusBadRequest {
		if v != nil {
			defer res.Body.Close()
			err = json.NewDecoder(res.Body).Decode(v)
			if err != nil {
				return err
			}
		}
		return nil
	}

	return r.handleError(req, res)
}

func (r *Relay42) doStream(req *http.Request, handlerFunc func([]byte)) error {
	if r.Debug == true {
		command, _ := http2curl.GetCurlCommand(req)
		fmt.Println(command)
	}

	res, err := r.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusBadRequest {
		defer res.Body.Close()
		var reader *bufio.Reader
		reader = bufio.NewReader(res.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				return err
			}
			line = bytes.TrimSpace(line)
			if len(line) == 0 {
				continue
			}
			handlerFunc(line)
		}
		return nil
	}

	return r.handleError(req, res)
}

func (r *Relay42) handleError(req *http.Request, res *http.Response) error {
	if r.Debug == true {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%q", dump)
	}

	var e ErrorResponse
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&e)
	if err != nil {
		return err
	}

	apiError := APIError{
		req: req,
		res: res,
		err: &e,
	}

	switch apiError.err.ErrorCode {

	case http.StatusBadRequest:
		return BadRequestError{apiError}
		break

	case http.StatusUnauthorized:
		return UnauthorizedError{apiError}
		break

	case http.StatusForbidden:
		return ForbiddenError{apiError}
		break

	case http.StatusInternalServerError:
		return InternalServerError{apiError}
		break

	}
	return e
}