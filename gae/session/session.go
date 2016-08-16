package session

import (
	"net/http"
)

type (
	Session interface{

		Get(key string,ptr interface{}) error
		Put(key string,ptr interface{}) error
		Id() string
	}
	SessionMaker interface{
		New() Session
		Get(string) Session
		//is session id valid ?
		IsValid(string) bool
	}
)

var BuildMaker func(*http.Request) SessionMaker
func Get(req *http.Request) Session{
	maker := BuildMaker(req)
	id := obtainId(req)
	if id == "" {
		return maker.New()
	}
	s := maker.Get(id)
	if s != nil {
		return s
	}
	return maker.new()
}
func obtainId(req *http.Request)string{
	return ""
}


