package sessions

import (
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"github.com/rocwong/neko"
)

func NewRediStore(size int, network, address, password string, keyPairs ...[]byte) (neko.SessionStore, error) {
	store, err := redistore.NewRediStore(size, network, address, password, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &rediStore{store}, nil
}

type rediStore struct {
	*redistore.RediStore
}

func (c *rediStore) Options(options neko.SessionOptions) {
	c.RediStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
	}
}
