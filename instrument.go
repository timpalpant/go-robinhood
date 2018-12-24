package robinhood

// List all known instruments.
func (c *Client) ListAllInstruments() ([]*Instrument, error) {
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
