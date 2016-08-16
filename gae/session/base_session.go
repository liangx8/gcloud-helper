package session

import (
	"net/http"
)
type (
	baseSessionMaker struct{
		w http.ResponseWriter
		r *http.Request

	}
	baseSession map[string]interface{}
)

func (sm *baseSessionMaker)New() Session{
}
func (sm *baseSessionMaker)Get(id string) Session{
}
func (sm *baseSessionMaker)IsValid(id string) bool{
}
// Generate a Unique id string
func UniqueId() string{
}
func BaseSessionInit(){
	BuildMaker = baseBuildMaker
}
func baseBuildMaker(w http.ResponseWriter,r *http.Request) SessionMaker{
	return &baseSessionMaker{w:w,r:r}
}
