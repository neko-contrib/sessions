package sessions

import (
	"github.com/gorilla/sessions"
	"github.com/rocwong/neko"
)

type cookieStore struct {
	*sessions.CookieStore
}

func (c *cookieStore) Options(options neko.SessionOptions) {
	c.CookieStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
	}
}

// NewCookieStore returns a new CookieStore.
func NewCookieStore(keyPairs ...[]byte) neko.SessionStore {
	return &cookieStore{sessions.NewCookieStore(keyPairs...)}
}
