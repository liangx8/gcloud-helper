package memcache
/*
import (
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"

	"github.com/liangx8/spark/session"
)
type (
	memSession struct{
	}
	sessionMaker struct{
		ctx context.Context
		
	}
)
// SessionMaker implements
func (sm *sessionMaker)New() session.Session{
}
func (sm *sessionMaker)Get(id string)session.Session{
}
func (sm *sessionMaker)IsValid(id string)bool{
	return true
}
// Session implement
func (se *memSession)Get(key string, ptr interface{})bool{
}
func (se *memSession)Put(key string, ptr interface{}){
}
func (se *memSession)Id()string{
}

func SessionInit(){
	session.BuildMaker= func(r *http.Request) session.SessionMaker{
		return &sessionMaker{ctx:appengine.NewContext(r)}
	}
}
func sessionById(ctx context.Context,id string) *memSession{
	
}

*/
