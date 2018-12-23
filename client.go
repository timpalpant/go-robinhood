package robinhood

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

const (
	Endpoint = "https://api.robinhood.com"
)

type Client struct {
	*http.Client
}

func NewClient(username, password string) *Client {
	return &Client{
		Client: &http.Client{},
	}
}

func (c *Client) GetInstruments() ([]*Instrument, error) {
	url := Endpoint + "/instruments/"
	var result []*Instrument
	for url != "" {
		var resp struct {
			Results []*Instrument
			Next    string
		}

		if err := c.getJSON(url, &resp); err != nil {
			return nil, err
		}

		result = append(result, resp.Results...)
		url = resp.Next
		if len(resp.Results) == 0 {
			break
		}
	}

	return result, nil
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
	resp, err := c.Client.Do(req)
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
