package session

import (
	"net/http"
)
type (
	baseSessionMaker struct{
		w http.ResponseWriter
		r *http.Request
		sessionPool chan map[string]interface{}
	}
	baseSession struct{
		getId func() string
	}
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
	pool :=make(chan map[string]*baseSession,1)
	pool<- make(map[string]*baseSession)
	return &baseSessionMaker{w:w,r:r,sessionPool:pool}
}
var table []rune = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
