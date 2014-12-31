package sessions

import (
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/rocwong/neko"
	"log"
	"net/http"
)

var appName string

func Sessions(name string, store neko.SessionStore) neko.HandlerFunc {
	return func(ctx *neko.Context) {
		appName = ctx.Engine.AppName
		sess := &session{name: name, request: ctx.Req, store: store, written: false, writer: ctx.Writer}
		ctx.Session = sess
		ctx.Writer.Before(func(neko.ResponseWriter) {
			if sess.Written() {
				err := sess.Session().Save(ctx.Req, ctx.Writer)
				if err != nil {
					log.Printf("[%s] SESSION ERROR! %s\n", appName, err)
				}
			}
		})
		defer context.Clear(ctx.Req)
		ctx.Next()
	}
}

type session struct {
	name    string
	request *http.Request
	store   neko.SessionStore
	session *sessions.Session
	written bool
	writer  http.ResponseWriter
}

func (s *session) Get(key interface{}) interface{} {
	return s.Session().Values[key]
}

func (s *session) Set(key interface{}, val interface{}) {
	s.Session().Values[key] = val
	s.written = true
}

func (s *session) Delete(key interface{}) {
	delete(s.Session().Values, key)
	s.written = true
}

func (s *session) Clear() {
	for key := range s.Session().Values {
		s.Delete(key)
	}
}

func (s *session) AddFlash(value interface{}, vars ...string) {
	s.Session().AddFlash(value, vars...)
	s.written = true
}

func (s *session) Flashes(vars ...string) []interface{} {
	s.written = true
	return s.Session().Flashes(vars...)
}

func (s *session) Options(options neko.SessionOptions) {
	s.Session().Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
	}
}

func (s *session) Session() *sessions.Session {
	if s.session == nil {
		var err error
		if s.session, err = s.store.Get(s.request, s.name); err != nil {
			log.Printf("[%s] SESSION ERROR! %s\n", appName, err)
		}
	}

	return s.session
}

func (s *session) Written() bool {
	return s.written
}
