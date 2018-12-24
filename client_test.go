package robinhood

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/pkg/errors"
)

func NewTestClient(responses []*http.Response) *Client {
	httpClient := mockHTTPClient(responses)
	return &Client{
		HTTPClient: &httpClient,
	}
}

type mockHTTPClient []*http.Response

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	old := *c
	if len(old) == 0 {
		return nil, fmt.Errorf("no more responses")
	}

	first := old[0]
	*c = old[1:]
	return first, nil
}

func loadResponses(filenames ...string) ([]*http.Response, error) {
	var result []*http.Response
	for _, filename := range filenames {
		resp, err := loadResponse(filename)
		if err != nil {
			return nil, errors.Wrapf(err, "error loading: %v", filename)
		}

		result = append(result, resp)
	}

	return result, nil
}

func loadResponse(filename string) (*http.Response, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	_, err = rr.Write(body)
	if err != nil {
		return nil, err
	}

	rr.WriteHeader(http.StatusOK)
	return rr.Result(), nil
}
