package robinhood

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang/glog"
	"github.com/google/go-querystring/query"
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
	HTTPClient
}

func NewClient(tokenSource oauth2.TokenSource) *Client {
	return &Client{
		HTTPClient: oauth2.NewClient(context.Background(), tokenSource),
	}
}

func (c *Client) getJSON(rawURL string, request interface{}, result interface{}) error {
	url, err := buildFullURL(rawURL, request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	return c.do(req, result)
}

func (c *Client) postForm(url string, request interface{}, result interface{}) error {
	values, err := query.Values(request)
	if err != nil {
		return err
	}

	reader := strings.NewReader(values.Encode())
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	return c.do(req, result)
}

func (c *Client) do(req *http.Request, result interface{}) error {
	glog.V(2).Infof("%v: %v", req.Method, req.URL)
	resp, err := c.HTTPClient.Do(req)
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

func buildFullURL(baseURL string, request interface{}) (string, error) {
	values, err := query.Values(request)
	if err != nil {
		return "", err
	}

	urlParsed, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	urlParsed.RawQuery = values.Encode()
	return urlParsed.String(), nil
}
