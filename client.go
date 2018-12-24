package robinhood

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

const (
	Endpoint = "https://api.robinhood.com"
)

// HTTPClient is an interface for the methods we use of http.Client.
// It enables using a mock HTTP client for tests.
type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

type Client struct {
	httpClient HTTPClient
}

func NewClient(tokenSource oauth2.TokenSource) *Client {
	return &Client{
		httpClient: oauth2.NewClient(context.Background(), tokenSource),
	}
}

func (c *Client) getJSON(url string, result interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	return c.do(req, result)
}

func (c *Client) do(req *http.Request, result interface{}) error {
	glog.V(2).Infof("%v: %v", req.Method, req.URL)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		apiError := APIError{StatusCode: resp.StatusCode}
		if err := dec.Decode(&apiError.Errors); err != nil {
			return errors.Wrapf(err, "%v: error reading body", resp.Status)
		}

		return apiError
	}

	return dec.Decode(result)
}
