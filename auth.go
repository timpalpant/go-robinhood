package robinhood

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// DefaultClientID is the OAuth2 ID used by the robinhood.com website.
const DefaultClientID = "c82SH0WZOsabOXGP2sxqcj34FxkvfnWRZBKlBjFS"

var ErrMFARequired = errors.New("two-factor auth (2FA) code required and not supplied")

// OAuth implements oauth2.TokenSource for Robinhood.
type OAuth struct {
	Endpoint, ClientID, Username, Password, MFA string
}

// OAuthResponse is the JSON-encoded response returned from OAuth2 token requests.
type OAuthResponse struct {
	oauth2.Token
	ExpiresIn   int    `json:"expires_in"`
	MFARequired bool   `json:"mfa_required"`
	MFAType     string `json:"mfa_type"`
}

func (o *OAuth) Token() (*oauth2.Token, error) {
	req, err := o.Request()
	if err != nil {
		return nil, errors.Wrap(err, "could not create request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not post login")
	}
	defer resp.Body.Close()

	var authResp OAuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode token")
	}

	if authResp.MFARequired {
		return nil, ErrMFARequired
	}

	token := &authResp.Token
	token.Expiry = time.Now().Add(time.Duration(authResp.ExpiresIn) * time.Second)
	return token, nil
}

func (o *OAuth) Request() (*http.Request, error) {
	ep := o.Endpoint
	if ep == "" {
		ep = Endpoint + "/oauth2/token"
	}

	clientID := o.ClientID
	if clientID == "" {
		clientID = DefaultClientID
	}

	u, err := url.Parse(ep)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Add("expires_in", fmt.Sprint(24*time.Hour/time.Second))
	q.Add("client_id", clientID)
	q.Add("grant_type", "password")
	q.Add("scope", "internal")
	u.RawQuery = q.Encode()

	authParams := url.Values{}
	authParams.Set("username", o.Username)
	authParams.Set("password", o.Password)
	if o.MFA != "" {
		authParams.Set("mfa_code", o.MFA)
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(authParams.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	return req, nil
}
