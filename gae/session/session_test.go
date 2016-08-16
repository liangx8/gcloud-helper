package session_test

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
)

func Test_cookie(t *testing.T){
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		cs := r.Cookies()
		if len(cs) > 0 {
			fmt.Fprintln(w,"cookies:")
			fmt.Fprintf(w,"%v\n",cs)
			fmt.Fprintln(w,"domain:",cs[0].Domain)
		} else {
			fmt.Fprintln(w,"N/A")
		}
		
	}))
	defer ts.Close()
	client := &http.Client{}
	req,err := http.NewRequest("GET",ts.URL,nil)
	req.AddCookie(&http.Cookie{
		Name:"sessionid",
		Value:"12345678",
		Path:"/",
		Domain:"www.rc-greed.com",
		Expires: time.Now(),
	})
	res,err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	body,err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	t.Errorf("%s",body)
}
