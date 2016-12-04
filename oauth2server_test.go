package oauth2server

import (
	"testing"
	"net/http"
)

func TestMain(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
    if (err != nil) {
        t.Fatal(err)
    }else if(req == nil){
    	t.Fatal(err)
    }
}