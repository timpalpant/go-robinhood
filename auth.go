package robinhood

import (
	"errors"
)

// DefaultClientID is the OAuth2 ID used by the robinhood.com website.
const DefaultClientID = "c82SH0WZOsabOXGP2sxqcj34FxkvfnWRZBKlBjFS"

var ErrMFARequired = errors.New("Two Factor Auth (2FA) code required and not supplied")
